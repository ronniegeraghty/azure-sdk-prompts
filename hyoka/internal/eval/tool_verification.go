package eval

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// toolVerifier accumulates SDK-reported tool loads and emits a deterministic
// bulk []progress.ToolStatus exactly once per eval, after every configured
// tool kind has produced its corresponding SDK load event.
//
// Contract (mirrors .squad/decisions.md round-1-2 tool-verification wiring):
//
//  1. At-most-once per verifier instance.
//  2. Emits once both configured kinds have fired (or the single relevant one
//     if only skills or only MCP are configured).
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
	readyChan      chan struct{} // Signals when verification is complete
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

// emitIfReady returns the bulk ToolStatus slice on the first call after all
// configured tool-kinds' load events have been observed. Subsequent calls
// return nil (at-most-once). Returns nil when:
//   - nothing was configured to verify (both expected maps empty)
//   - not every configured kind has observed its load event yet
//   - a previous call already emitted
func (v *toolVerifier) emitIfReady() []progress.ToolStatus {
	if v.emitted {
		return nil
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
			reason = "SDK did not report skill as loaded"
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
			reason = "SDK did not report MCP server as loaded"
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
	close(v.readyChan) // Signal that verification is complete
	return tools
}

// waitForToolVerification blocks until the verifier completes or the context
// times out. Returns the verified tool statuses and any validation errors.
// This is the blocking gate that prevents evals from proceeding with failed tools.
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
		return nil, fmt.Errorf("tool verification timeout: SDK did not confirm tool load within %v", timeout)
	}
}
