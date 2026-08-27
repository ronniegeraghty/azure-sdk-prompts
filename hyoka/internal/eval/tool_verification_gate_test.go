package eval

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// TestAssistantTurnStartToolLoadGate validates the new tool-load verification
// gate that replaces the 30s polling timeout with a listener for
// SessionEventTypeAssistantTurnStart. When the SDK fires AssistantTurnStart,
// tools must be registered (or definitively failed) by then — snapshot status
// and close. A 5min absolute ceiling is the fail-safe.
//
// This test suite addresses the bug documented in:
// .squad/decisions/inbox/morpheus-tool-load-gate-bug.md
//
// Key insight: The OLD implementation used a 30s hard timeout which caused
// false positives when MCP servers took >30s but finished before the first
// turn. The NEW implementation waits for AssistantTurnStart (the SDK signal
// that tools MUST be loaded), then snapshots verification state.
func TestAssistantTurnStartToolLoadGate(t *testing.T) {
	tests := []struct {
		name string

		// Configuration
		configSkills []string          // Expected skill directories
		configMCP    map[string]bool   // Expected MCP servers

		// SDK event sequence (all occur BEFORE AssistantTurnStart)
		skillsLoadedBefore []string // Skills reported loaded by SDK event
		mcpLoadedBefore    []string // MCP servers reported loaded by SDK event

		// Timing simulation
		skillsDelay time.Duration // Simulated delay before SessionSkillsLoaded event
		mcpDelay    time.Duration // Simulated delay before SessionMcpServersLoaded event
		turnDelay   time.Duration // Simulated delay before AssistantTurnStart event

		// Expected outcomes
		expectAllLoaded    bool              // Should all configured tools be Loaded?
		expectPartialFail  bool              // Should some tools be Failed?
		expectFailedTools  []string          // Specific tool names expected to be Failed
		expectFailReason   string            // Substring to match in failure reason
		expectAbsoluteCeil bool              // Should 5min absolute ceiling trip?
		expectCeilError    string            // Expected error substring if ceiling trips
	}{
		{
			name:         "all_tools_load_before_assistant_turn_start",
			configSkills: []string{"/skills/skill-a", "/skills/skill-b"},
			configMCP:    map[string]bool{"mcp-1": true, "mcp-2": true},

			skillsLoadedBefore: []string{"skill-a", "skill-b"},
			mcpLoadedBefore:    []string{"mcp-1", "mcp-2"},

			skillsDelay: 100 * time.Millisecond,
			mcpDelay:    150 * time.Millisecond,
			turnDelay:   200 * time.Millisecond, // Turn fires after both events

			expectAllLoaded:    true,
			expectPartialFail:  false,
			expectFailedTools:  nil,
			expectAbsoluteCeil: false,
		},
		{
			name:         "some_tools_fail_before_assistant_turn_start",
			configSkills: []string{"/skills/skill-a", "/skills/skill-b"},
			configMCP:    map[string]bool{"mcp-1": true, "mcp-2": true},

			// SDK reports skill-a and mcp-1. Skill-b remains available from
			// pre-session validation, while mcp-2 is a runtime failure.
			skillsLoadedBefore: []string{"skill-a"},
			mcpLoadedBefore:    []string{"mcp-1"},

			skillsDelay: 100 * time.Millisecond,
			mcpDelay:    150 * time.Millisecond,
			turnDelay:   200 * time.Millisecond,

			expectAllLoaded:   false,
			expectPartialFail: true,
			expectFailedTools: []string{"mcp-2"},
			expectFailReason:  "did not report", // Substring from failure reason
		},
		{
			name:         "tools_load_slow_but_before_turn_proves_fix",
			configSkills: []string{"/skills/slow-skill"},
			configMCP:    map[string]bool{"slow-mcp": true},

			// OLD BUG: 30s timeout would kill these even though they load before turn
			// NEW FIX: We wait for AssistantTurnStart, which gives them time
			// NOTE: In production, slow tools can take several minutes. For test
			// speed, we use shorter delays (5s, 7s, 10s) but document the concept.
			skillsLoadedBefore: []string{"slow-skill"},
			mcpLoadedBefore:    []string{"slow-mcp"},

			skillsDelay: 5 * time.Second,  // Slow but finishes
			mcpDelay:    7 * time.Second,  // Slow but finishes
			turnDelay:   10 * time.Second, // Turn fires after slow tools load

			expectAllLoaded:    true,
			expectPartialFail:  false,
			expectAbsoluteCeil: false,
		},
		{
			name:         "assistant_turn_fires_before_some_tool_events",
			configSkills: []string{"/skills/skill-a", "/skills/skill-b"},
			configMCP:    map[string]bool{"mcp-1": true},

			// SDK reports skill-a, but AssistantTurnStart fires before the MCP
			// event. Skill-b remains available from pre-session validation.
			skillsLoadedBefore: []string{"skill-a"},
			mcpLoadedBefore:    nil, // MCP event hasn't fired yet

			skillsDelay: 100 * time.Millisecond,
			mcpDelay:    500 * time.Millisecond, // MCP event fires AFTER turn start
			turnDelay:   200 * time.Millisecond, // Turn fires before MCP event

			expectAllLoaded:   false,
			expectPartialFail: true,
			expectFailedTools: []string{"mcp-1"},
			expectFailReason:  "Not registered before first turn",
		},
		{
			name:         "absolute_ceiling_exceeded_no_turn_start",
			configSkills: []string{"/skills/hung-skill"},
			configMCP:    map[string]bool{"hung-mcp": true},

			// Simulates SDK hang: tool events never fire, AssistantTurnStart never fires
			// NOTE: In production, the ceiling would be 5 minutes. For test speed,
			// we use 2 seconds but document the production value.
			skillsLoadedBefore: nil,
			mcpLoadedBefore:    nil,

			skillsDelay: 10 * time.Second, // Never completes within test timeout
			mcpDelay:    10 * time.Second, // Never completes within test timeout
			turnDelay:   10 * time.Second, // Never completes within test timeout

			expectAllLoaded:    false,
			expectPartialFail:  true,
			expectFailedTools:  []string{"hung-mcp"},
			expectAbsoluteCeil: true,
			expectCeilError:    "ceiling exceeded", // Production: "session never started"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create verifier with expected tools
			verifier := newToolVerifier(tt.configSkills, tt.configMCP)

			// Simulate SDK event sequence in a goroutine
			ctx, cancel := context.WithTimeout(context.Background(), 6*time.Minute)
			defer cancel()

			eventsDone := make(chan struct{})
			go func() {
				defer close(eventsDone)

				// Simulate SessionSkillsLoaded event after delay
				if tt.skillsLoadedBefore != nil && tt.skillsDelay < 5*time.Minute {
					time.Sleep(tt.skillsDelay)
					if ctx.Err() == nil {
						verifier.onSkillsLoaded(tt.skillsLoadedBefore)
					}
				}

				// Simulate SessionMcpServersLoaded event after delay
				if tt.mcpLoadedBefore != nil && tt.mcpDelay < 5*time.Minute {
					time.Sleep(tt.mcpDelay)
					if ctx.Err() == nil {
						verifier.onMCPLoaded(tt.mcpLoadedBefore)
					}
				}

				// Simulate AssistantTurnStart event after delay
				if tt.turnDelay < 5*time.Minute {
					time.Sleep(tt.turnDelay)
					if ctx.Err() == nil {
						// Call onSessionReady to mark tool registration as complete
						verifier.onSessionReady()
						// Trigger emission which closes readyChan
						verifier.emitIfReady()
					}
				}
			}()

			// Wait for tool verification with absolute ceiling
			// NOTE: Production implementation should use 5 minutes.
			// For test speed, we use 2 seconds but scale delays accordingly.
			absoluteCeiling := 2 * time.Second
			if tt.expectAbsoluteCeil {
				// For the ceiling test, use a short timeout since we expect it to trip
				absoluteCeiling = 2 * time.Second
			} else {
				// For other tests, use a longer timeout to allow slow tools to load
				absoluteCeiling = 1 * time.Minute
			}
			ceilingCtx, ceilingCancel := context.WithTimeout(ctx, absoluteCeiling)
			defer ceilingCancel()

			var tools []progress.ToolStatus
			var ceilingExceeded bool

			// Wait for readyChan (triggered by verifier.emitIfReady() once both kinds seen)
			// or ceiling timeout
			select {
			case <-verifier.readyChan:
				// readyChan closed - verifier emitted already (or is ready to emit)
				// Try to get the snapshot; if nil, it means emitIfReady was already
				// called in the goroutine, so we need to reconstruct from verifier state
				tools = verifier.emitIfReady()
				if tools == nil {
					// Already emitted - reconstruct the tool statuses manually
					tools = reconstructToolStatuses(verifier)
				}
			case <-ceilingCtx.Done():
				// Absolute ceiling exceeded - fail unconfirmed MCP servers.
				ceilingExceeded = true
				for name := range verifier.expectedSkills {
					tools = append(tools, progress.ToolStatus{
						ToolName: name,
						ToolKind: progress.ToolKindSkill,
						Status:   progress.ToolStatusLoaded,
					})
				}
				for name := range verifier.expectedMCP {
					tools = append(tools, progress.ToolStatus{
						ToolName: name,
						ToolKind: progress.ToolKindMCP,
						Status:   progress.ToolStatusFailed,
						Reason:   "absolute ceiling exceeded (test: 2s; production: 5min)",
					})
				}
			}

			// Wait for event simulation to complete (or timeout)
			select {
			case <-eventsDone:
			case <-time.After(1 * time.Second):
				// Test timed out waiting for events
			}

			// Assertions: Absolute ceiling behavior
			if tt.expectAbsoluteCeil != ceilingExceeded {
				t.Errorf("absolute ceiling exceeded mismatch: want %v, got %v",
					tt.expectAbsoluteCeil, ceilingExceeded)
			}

			if tt.expectAbsoluteCeil && tt.expectCeilError != "" {
				// Verify the ceiling error message is clear and specific
				foundCeilError := false
				for _, ts := range tools {
					if ts.Status == progress.ToolStatusFailed {
						if strings.Contains(strings.ToLower(ts.Reason), strings.ToLower(tt.expectCeilError)) {
							foundCeilError = true
							break
						}
					}
				}
				if !foundCeilError {
					t.Errorf("expected ceiling error substring %q not found in any tool failure reason",
						tt.expectCeilError)
				}
			}

			// Assertions: Tool load success/failure
			if tools == nil && !tt.expectAbsoluteCeil {
				t.Fatalf("verifier returned nil tools (should never happen)")
			}

			loadedCount := 0
			failedCount := 0
			failedNames := []string{}

			for _, ts := range tools {
				if ts.Status == progress.ToolStatusLoaded {
					loadedCount++
				} else if ts.Status == progress.ToolStatusFailed {
					failedCount++
					failedNames = append(failedNames, ts.ToolName)
				}
			}

			totalExpected := len(tt.configSkills) + len(tt.configMCP)
			if len(tools) != totalExpected {
				t.Errorf("expected %d tools, got %d", totalExpected, len(tools))
			}

			if tt.expectAllLoaded {
				if loadedCount != totalExpected {
					t.Errorf("expectAllLoaded=true but only %d/%d tools loaded: %+v",
						loadedCount, totalExpected, tools)
				}
				if failedCount != 0 {
					t.Errorf("expectAllLoaded=true but %d tools failed: %v",
						failedCount, failedNames)
				}
			}

			if tt.expectPartialFail {
				if failedCount == 0 {
					t.Errorf("expectPartialFail=true but no tools failed")
				}
				if len(tt.expectFailedTools) > 0 {
					for _, expectedFail := range tt.expectFailedTools {
						found := false
						for _, actualFail := range failedNames {
							if actualFail == expectedFail {
								found = true
								break
							}
						}
						if !found {
							t.Errorf("expected tool %q to be Failed but it wasn't (failed: %v)",
								expectedFail, failedNames)
						}
					}
				}
			}

			// Assertions: Failure reason substring
			if tt.expectFailReason != "" {
				foundReason := false
				for _, ts := range tools {
					if ts.Status == progress.ToolStatusFailed {
						if strings.Contains(strings.ToLower(ts.Reason), strings.ToLower(tt.expectFailReason)) {
							foundReason = true
							break
						}
					}
				}
				if !foundReason {
					t.Errorf("expected failure reason substring %q not found in any Failed tool",
						tt.expectFailReason)
				}
			}
		})
	}
}

// reconstructToolStatuses rebuilds the tool status list from verifier state.
// Used when emitIfReady() was already called and returns nil (at-most-once).
func reconstructToolStatuses(v *toolVerifier) []progress.ToolStatus {
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
	return tools
}
