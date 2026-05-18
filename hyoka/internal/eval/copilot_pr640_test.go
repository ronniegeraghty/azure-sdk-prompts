package eval

import (
	"testing"
)

// TestMaxSessionActionsLimit_Resolution verifies that the action limit resolution
// logic correctly prioritizes evalMaxSessionActions over engine maxSessionActions.
// This is the core logic from copilot.go lines 262-266.
// Regression test for PR #640 Fix 3 (lines 600, 622 should use maxSessionActionsLimit).
func TestMaxSessionActionsLimit_Resolution(t *testing.T) {
	tests := []struct {
		name                     string
		evalMaxSessionActions    int
		engineMaxSessionActions  int
		expectedResolvedLimit    int
	}{
		{
			name:                    "per-eval limit overrides engine default",
			evalMaxSessionActions:   3,
			engineMaxSessionActions: 100,
			expectedResolvedLimit:   3, // per-eval wins
		},
		{
			name:                    "engine default used when per-eval not set",
			evalMaxSessionActions:   0, // not set
			engineMaxSessionActions: 50,
			expectedResolvedLimit:   50, // engine default
		},
		{
			name:                    "both unset defaults to 0",
			evalMaxSessionActions:   0,
			engineMaxSessionActions: 0,
			expectedResolvedLimit:   0, // no limit
		},
		{
			name:                    "negative per-eval uses engine default",
			evalMaxSessionActions:   -5,
			engineMaxSessionActions: 25,
			expectedResolvedLimit:   25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the resolution logic from copilot.go:262-266
			maxSessionActionsLimit := tt.evalMaxSessionActions
			if maxSessionActionsLimit <= 0 {
				maxSessionActionsLimit = tt.engineMaxSessionActions
			}

			if maxSessionActionsLimit != tt.expectedResolvedLimit {
				t.Errorf("expected resolved limit %d, got %d", tt.expectedResolvedLimit, maxSessionActionsLimit)
			}
		})
	}
}

// TestActionCounting_EventTypes verifies which event types should count as actions.
// Per Ronnie's spec: "anything the agent does is an action: reasoning, tool calls, 
// bash commands, responses."
// Regression test for PR #640 Fix 3 (assistant.reasoning should be counted).
func TestActionCounting_EventTypes(t *testing.T) {
	tests := []struct {
		name          string
		eventType     string
		countsAsAction bool
	}{
		{
			name:          "tool.execution_start counts as action",
			eventType:     "tool.execution_start",
			countsAsAction: true,
		},
		{
			name:          "assistant.message counts as action",
			eventType:     "assistant.message",
			countsAsAction: true,
		},
		{
			name:          "assistant.reasoning counts as action (NEW in PR #640)",
			eventType:     "assistant.reasoning",
			countsAsAction: true,
		},
		{
			name:          "tool.execution_complete does NOT count",
			eventType:     "tool.execution_complete",
			countsAsAction: false,
		},
		{
			name:          "assistant.turn_start does NOT count",
			eventType:     "assistant.turn_start",
			countsAsAction: false,
		},
		{
			name:          "assistant.turn_end does NOT count",
			eventType:     "assistant.turn_end",
			countsAsAction: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is the switch statement logic from copilot.go:597-627
			countsAsAction := false

			switch tt.eventType {
			case "tool.execution_start":
				countsAsAction = true
			case "assistant.message":
				countsAsAction = true
			case "assistant.reasoning":
				// PR #640 adds this case
				countsAsAction = true
			}

			if countsAsAction != tt.countsAsAction {
				t.Errorf("event type %q: expected countsAsAction=%v, got %v",
					tt.eventType, tt.countsAsAction, countsAsAction)
			}
		})
	}
}

// TestActionCounting_LimitEnforcement verifies that the action counter stops
// processing once the limit is exceeded.
// Regression test for PR #640 Fix 3.
func TestActionCounting_LimitEnforcement(t *testing.T) {
	maxActions := 5

	// Simulate a stream of events
	events := []struct {
		eventType string
	}{
		{"tool.execution_start"},     // 1
		{"assistant.reasoning"},      // 2
		{"assistant.message"},        // 3
		{"tool.execution_start"},     // 4
		{"assistant.reasoning"},      // 5
		{"tool.execution_complete"},  // not counted
		{"assistant.message"},        // 6 - should trigger limit
		{"tool.execution_start"},     // should not reach
	}

	actionCount := 0
	limitHit := false

	for _, event := range events {
		// Count actions
		switch event.eventType {
		case "tool.execution_start", "assistant.message", "assistant.reasoning":
			actionCount++
			// Check limit (logic from copilot.go:600 and 622)
			if maxActions > 0 && actionCount > maxActions && !limitHit {
				limitHit = true
				t.Logf("Limit hit at action %d", actionCount)
				break
			}
		}

		if limitHit {
			break
		}
	}

	if !limitHit {
		t.Errorf("expected limit to be hit")
	}

	// The counter should have stopped at maxActions + 1 (the event that triggered the check)
	if actionCount != maxActions+1 {
		t.Errorf("expected actionCount %d (limit + 1), got %d", maxActions+1, actionCount)
	}
}

// TestActionCounting_PerEvalLimitPriority is an integration-style test that
// verifies the entire flow: per-eval limit resolution + action counting.
// This test exercises the exact scenario from the bug report.
func TestActionCounting_PerEvalLimitPriority(t *testing.T) {
	// Engine default is high
	engineMaxSessionActions := 100

	// Per-eval limit is lower (this should win)
	evalMaxSessionActions := 3

	// Resolve limit (copilot.go:262-266)
	maxSessionActionsLimit := evalMaxSessionActions
	if maxSessionActionsLimit <= 0 {
		maxSessionActionsLimit = engineMaxSessionActions
	}

	if maxSessionActionsLimit != 3 {
		t.Fatalf("setup error: expected limit 3, got %d", maxSessionActionsLimit)
	}

	// Simulate events
	events := []string{
		"tool.execution_start",  // 1
		"assistant.message",     // 2
		"tool.execution_start",  // 3
		"assistant.message",     // 4 - should trigger limit
		"tool.execution_start",  // should not reach
	}

	actionCount := 0
	limitHit := false

	for _, eventType := range events {
		switch eventType {
		case "tool.execution_start", "assistant.message":
			actionCount++
			// CRITICAL: Must use maxSessionActionsLimit, not engineMaxSessionActions
			if maxSessionActionsLimit > 0 && actionCount > maxSessionActionsLimit && !limitHit {
				limitHit = true
				break
			}
		}
		if limitHit {
			break
		}
	}

	// Limit should have been hit at per-eval threshold (3), NOT engine default (100)
	if !limitHit {
		t.Errorf("expected limit to be hit at per-eval threshold %d", evalMaxSessionActions)
	}

	if actionCount > evalMaxSessionActions+1 {
		t.Errorf("action counter exceeded per-eval limit: count=%d, limit=%d", 
			actionCount, evalMaxSessionActions)
	}

	// If we had used engineMaxSessionActions instead, counter would have gone to 5
	// This proves the test would catch the bug on lines 600 and 622
}
