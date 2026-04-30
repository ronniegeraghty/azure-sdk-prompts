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

// toolVerifier accumulates SDK-reported tool loads and emits a deterministic
// bulk []progress.ToolStatus exactly once per eval, after every configured
// tool kind has produced its corresponding SDK load event OR the first
// assistant turn starts (whichever comes first).
//
// Contract (mirrors .squad/decisions.md round-1-2 tool-verification wiring):
//
//  1. At-most-once per verifier instance.
//  2. Emits once both configured kinds have fired (or the single relevant one
//     if only skills or only MCP are configured), OR when the first assistant
//     turn starts (which definitively marks tool registration as complete).
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
	expectedSkills    map[string]bool
	expectedMCP       map[string]bool
	loadedSkills      map[string]bool
	loadedMCP         map[string]bool
	skillsEvtSeen     bool
	mcpEvtSeen        bool
	emitted           bool
	emittedTools      []progress.ToolStatus // Cached result of first successful emit
	readyChan         chan struct{}         // Signals when verification is complete
	turnBeforeSkills  bool          // True if turn started before skills event
	turnBeforeMCP     bool          // True if turn started before MCP event
}

// newToolVerifier builds a verifier keyed on expected skills (derived from
// resolved skill directory basenames) and expected MCP server names.
func newToolVerifier(skillDirs []string, mcpNames map[string]bool) *toolVerifier {
	sk := make(map[string]bool, len(skillDirs))
	for _, dir := range skillDirs {
		name := filepath.Base(dir)
		if name == "" || name == "." || name == "/" {
			continue
		}
		sk[name] = true
	}
	mc := make(map[string]bool, len(mcpNames))
	for n := range mcpNames {
		mc[n] = true
	}
	return &toolVerifier{
		expectedSkills: sk,
		expectedMCP:    mc,
		loadedSkills:   make(map[string]bool),
		loadedMCP:      make(map[string]bool),
		readyChan:      make(chan struct{}),
	}
}

// onSkillsLoaded records that the SessionSkillsLoaded SDK event fired and
// which skill names it reported. Safe to call even when names is empty —
// the presence of the event alone marks the skill channel as seen so a
// later emit can proceed.
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

// onSessionReady is called when AssistantTurnStart fires — the definitive
// signal that tool loading is complete. If we haven't seen tool-load events
// by the time the first turn starts, they will never arrive. This method
// forces verification to complete so postSessionToolVerification doesn't
// wait forever for events that will never come.
func (v *toolVerifier) onSessionReady() {
	// Mark which kinds have NOT been seen yet (turn happened before their event)
	if len(v.expectedSkills) > 0 && !v.skillsEvtSeen {
		v.turnBeforeSkills = true
		v.skillsEvtSeen = true // Force as seen so emitIfReady will proceed
	}
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
		status := progress.ToolStatusFailed
		reason := ""
		if v.loadedSkills[name] {
			status = progress.ToolStatusLoaded
		} else {
			if v.turnBeforeSkills {
				reason = "Not registered before first turn"
			} else {
				reason = "SDK did not report skill as loaded"
			}
		}
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindSkill,
			Status:   status,
			Reason:   reason,
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
// The real completion signal is onSessionReady (fired when AssistantTurnStart
// happens), which marks tool registration as definitively complete. The timeout
// only fires when the session is broken (auth hang, network failure, SDK bug).
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

// postSessionToolVerification waits for the SDK to confirm every configured
// skill / MCP server actually loaded after `session.SendAndWait` returned.
// Returns a tool.SummarizeToolLoadErrors-formatted summary of any failures
// (matching the pre-session validation format from Item D), or "" when
// every configured tool loaded cleanly OR when nothing was configured to
// verify.
//
// Timeout semantics (Option A): The timeout is an ABSOLUTE CEILING (default
// 5 minutes) for broken sessions that never reach first turn (auth hang,
// network failure). The PRIMARY gate is onSessionReady, which fires when
// AssistantTurnStart happens and marks tool registration as definitively
// complete. If the ceiling hits, every configured tool is marked Failed with
// reason "session did not reach first turn within <timeout>". The point of
// this gate is to eliminate false-positive evals; we'd rather hard-fail on
// a hung session than silently grade code that never had its tools.
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
	// every SDK event already fired. Calling emitIfReady ourselves closes
	// readyChan if both kinds have reported.
	tools := v.emitIfReady()
	if tools == nil {
		var err error
		tools, err = waitForToolVerification(ctx, v, timeout)
		if err != nil {
			// Timeout — synthesize a Failed entry per configured tool so the
			// summary names every tool we expected to see and didn't.
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

// expectedAsTimeoutFailures builds the synthetic Failed list used when
// waitForToolVerification times out. Sort order matches emitIfReady so
// renderers and snapshot tests see identical ordering across the
// happy/timeout paths.
func expectedAsTimeoutFailures(v *toolVerifier, timeout time.Duration) []progress.ToolStatus {
	reason := fmt.Sprintf("Session did not reach first turn within %s", timeout)
	tools := make([]progress.ToolStatus, 0, len(v.expectedSkills)+len(v.expectedMCP))
	for name := range v.expectedSkills {
		tools = append(tools, progress.ToolStatus{
			ToolName: name,
			ToolKind: progress.ToolKindSkill,
			Status:   progress.ToolStatusFailed,
			Reason:   reason,
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
