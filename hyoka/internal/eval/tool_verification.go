package eval

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// toolVerifier emits a deterministic bulk []progress.ToolStatus exactly once
// per eval. Skill directories are validated before session creation, so they
// are available without requiring the optional SessionSkillsLoaded event.
// MCP servers still require runtime confirmation from the SDK.
//
// Contract (mirrors .squad/decisions.md round-1-2 tool-verification wiring):
//
//  1. At-most-once per verifier instance.
//  2. Emits once configured MCP servers have produced their SDK load event, OR
//     when the first assistant turn starts. Skill-only configurations can emit
//     immediately because their directories were prevalidated.
//  3. Never emits when neither skills nor MCP are configured.
//  4. Tools sorted by (kind, name) ascending — deterministic for renderers
//     and snapshot tests.
//  5. Every configured tool appears exactly once. SDK-reported names that
//     were not configured are ignored (intent: "did what I asked for load?").
//  6. Skill match is by basename of each skill directory — the SDK reports
//     skill names that correspond to the directory that holds SKILL.md.
//
// Not safe for concurrent use. The Copilot OnEvent handler serializes every
// call under a mutex and only invokes the returned slice's consumer after
// releasing that mutex.
type toolVerifier struct {
	expectedSkills map[string]bool
	expectedMCP    map[string]bool
	loadedSkills   map[string]bool
	loadedMCP      map[string]bool
	skillsEvtSeen  bool
	mcpEvtSeen     bool
	emitted        bool
	emittedTools   []progress.ToolStatus // Cached result of first successful emit
	readyChan      chan struct{}         // Signals when verification is complete
	turnBeforeMCP  bool                  // True if turn started before MCP event
}

// newToolVerifier builds a verifier keyed on expected skills (derived from
// resolved skill directory basenames) and expected MCP server names.
func newToolVerifier(skillDirs []string, mcpNames map[string]bool) *toolVerifier {
	sk := make(map[string]bool, len(skillDirs))
	loadedSkills := make(map[string]bool, len(skillDirs))
	for _, dir := range skillDirs {
		name := filepath.Base(dir)
		if name == "" || name == "." || name == "/" || name == `\` {
			continue
		}
		sk[name] = true
		loadedSkills[name] = true
	}
	mc := make(map[string]bool, len(mcpNames))
	for n := range mcpNames {
		mc[n] = true
	}
	return &toolVerifier{
		expectedSkills: sk,
		expectedMCP:    mc,
		loadedSkills:   loadedSkills,
		loadedMCP:      make(map[string]bool),
		skillsEvtSeen:  len(sk) > 0,
		readyChan:      make(chan struct{}),
	}
}

// onSkillsLoaded records optional SDK corroboration. SDK v1 does not guarantee
// that this event fires before the first turn, or at all.
func (v *toolVerifier) onSkillsLoaded(names []string) {
	v.skillsEvtSeen = true
	for _, n := range names {
		v.loadedSkills[n] = true
	}
}

// onMCPLoaded records the SessionMcpServersLoaded SDK event, mirror of
// onSkillsLoaded for the MCP channel.
func (v *toolVerifier) onMCPLoaded(names []string) {
	v.mcpEvtSeen = true
	for _, n := range names {
		v.loadedMCP[n] = true
	}
}

// onSessionReady is called when AssistantTurnStart fires. It finalizes MCP
// verification so postSessionToolVerification does not wait forever for an MCP
// event that will never arrive. Skill events are optional and do not gate.
func (v *toolVerifier) onSessionReady() {
	if len(v.expectedMCP) > 0 && !v.mcpEvtSeen {
		v.turnBeforeMCP = true
		v.mcpEvtSeen = true // Force as seen so emitIfReady will proceed
	}
}

// emitIfReady returns the bulk ToolStatus slice on the first call after all
// configured tool-kinds' load events have been observed. Subsequent calls
// return nil (at-most-once). Returns nil when:
//   - nothing was configured to verify (both expected maps empty)
//   - not every configured kind has observed its load event yet
//   - a previous call already emitted
func (v *toolVerifier) emitIfReady() []progress.ToolStatus {
	if v.emitted {
		// Return the cached result so callers (e.g., postSessionToolVerification)
		// after an in-line emit still see the verified tool list instead of nil.
		return v.emittedTools
	}
	needSkills := len(v.expectedSkills) > 0
	needMCP := len(v.expectedMCP) > 0
	if !needSkills && !needMCP {
		return nil
	}
	if needSkills && !v.skillsEvtSeen {
		return nil
	}
	if needMCP && !v.mcpEvtSeen {
		return nil
	}

	tools := make([]progress.ToolStatus, 0, len(v.expectedSkills)+len(v.expectedMCP))
	for name := range v.expectedSkills {
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindSkill,
			Status:   progress.ToolStatusLoaded,
		})
	}
	for name := range v.expectedMCP {
		status := progress.ToolStatusFailed
		reason := ""
		if v.loadedMCP[name] {
			status = progress.ToolStatusLoaded
		} else {
			if v.turnBeforeMCP {
				reason = "Not registered before first turn"
			} else {
				reason = "SDK did not report MCP server as loaded"
			}
		}
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindMCP,
			Status:   status,
			Reason:   reason,
		})
	}
	sort.Slice(tools, func(i, j int) bool {
		if tools[i].ToolKind != tools[j].ToolKind {
			return tools[i].ToolKind < tools[j].ToolKind
		}
		return tools[i].ToolName < tools[j].ToolName
	})
	v.emitted = true
	v.emittedTools = tools
	close(v.readyChan) // Signal that verification is complete
	return tools
}

// waitForToolVerification blocks until the verifier completes or the context
// times out. Returns the verified tool statuses and any validation errors.
// This is the blocking gate that prevents evals from proceeding with failed tools.
//
// The timeout is an ABSOLUTE CEILING (default 5 minutes), not the primary gate.
// The real completion signal is the MCP load event or onSessionReady. The
// timeout only fires when a configured MCP server reaches neither signal.
func waitForToolVerification(ctx context.Context, v *toolVerifier, timeout time.Duration) ([]progress.ToolStatus, error) {
	// If nothing is configured to verify, return immediately with no error
	if len(v.expectedSkills) == 0 && len(v.expectedMCP) == 0 {
		return nil, nil
	}

	// Create timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Wait for verification to complete or timeout
	select {
	case <-v.readyChan:
		// Verification complete — collect the results
		tools := v.emitIfReady()
		if tools == nil {
			// This shouldn't happen if readyChan closed, but guard against it
			return nil, fmt.Errorf("tool verification completed but no statuses were emitted")
		}
		return tools, nil
	case <-timeoutCtx.Done():
		return nil, fmt.Errorf("tool verification timeout: session did not reach first turn within %v", timeout)
	}
}

// postSessionToolVerification waits for the SDK to confirm configured MCP
// servers after `session.SendAndWait` returned. Skills were validated before
// session creation and do not depend on SessionSkillsLoaded.
// Returns a tool.SummarizeToolLoadErrors-formatted summary of any failures
// (matching the pre-session validation format from Item D), or "" when every
// configured MCP server loaded cleanly OR when nothing was configured.
//
// Timeout semantics (Option A): The timeout is an ABSOLUTE CEILING (default
// 5 minutes) for broken sessions that never reach first turn (auth hang,
// network failure). The PRIMARY gate is onSessionReady, which fires when
// AssistantTurnStart happens and finalizes MCP registration. If the ceiling
// hits, configured MCP servers are marked Failed while prevalidated skills
// remain Loaded.
func postSessionToolVerification(ctx context.Context, v *toolVerifier, timeout time.Duration) string {
	if v == nil {
		return ""
	}
	if len(v.expectedSkills) == 0 && len(v.expectedMCP) == 0 {
		return ""
	}
	// Opportunistic flush: in production the OnEvent handler only calls
	// emitIfReady when a progressFn is registered, so an eval running
	// without a progress display would otherwise time out here even when
	// every required SDK event already fired. Calling emitIfReady ourselves
	// closes readyChan once MCP verification is ready.
	tools := v.emitIfReady()
	if tools == nil {
		var err error
		tools, err = waitForToolVerification(ctx, v, timeout)
		if err != nil {
			// Timeout — fail MCP servers that never reached a runtime signal.
			tools = expectedAsTimeoutFailures(v, timeout)
		}
	}
	var failed []*tool.ToolLoadError
	for _, t := range tools {
		if t.Status == progress.ToolStatusFailed {
			failed = append(failed, &tool.ToolLoadError{
				Kind:   t.ToolKind,
				Name:   t.ToolName,
				Reason: t.Reason,
			})
		}
	}
	return tool.SummarizeToolLoadErrors(failed)
}

// expectedAsTimeoutFailures builds the synthetic result used when runtime MCP
// verification times out. Prevalidated skills remain loaded.
func expectedAsTimeoutFailures(v *toolVerifier, timeout time.Duration) []progress.ToolStatus {
	reason := fmt.Sprintf("Session did not reach first turn within %s", timeout)
	tools := make([]progress.ToolStatus, 0, len(v.expectedSkills)+len(v.expectedMCP))
	for name := range v.expectedSkills {
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindSkill,
			Status:   progress.ToolStatusLoaded,
		})
	}
	for name := range v.expectedMCP {
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindMCP,
			Status:   progress.ToolStatusFailed,
			Reason:   reason,
		})
	}
	sort.Slice(tools, func(i, j int) bool {
		if tools[i].ToolKind != tools[j].ToolKind {
			return tools[i].ToolKind < tools[j].ToolKind
		}
		return tools[i].ToolName < tools[j].ToolName
	})
	return tools
}
