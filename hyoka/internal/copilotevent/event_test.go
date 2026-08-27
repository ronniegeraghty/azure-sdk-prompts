package copilotevent

import (
	"testing"
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

func TestToolTrackerEnrichesMatchingCompletion(t *testing.T) {
	startedAt := time.Date(2026, time.August, 27, 1, 2, 3, 0, time.UTC)
	start := copilot.SessionEvent{
		Timestamp: startedAt,
		Data: &copilot.ToolExecutionStartData{
			ToolCallID: "call-1",
			ToolName:   "view",
		},
	}
	complete := copilot.SessionEvent{
		Timestamp: startedAt.Add(42_500 * time.Microsecond),
		Data: &copilot.ToolExecutionCompleteData{
			ToolCallID: "call-1",
			Success:    true,
		},
	}

	var tracker ToolTracker
	startDetails := Extract(start)
	tracker.Enrich(start, &startDetails)
	completeDetails := Extract(complete)
	tracker.Enrich(complete, &completeDetails)

	if completeDetails.ToolName == nil || *completeDetails.ToolName != "view" {
		t.Fatalf("ToolName = %v, want view", completeDetails.ToolName)
	}
	if completeDetails.Duration == nil || *completeDetails.Duration != 42.5 {
		t.Fatalf("Duration = %v, want 42.5", completeDetails.Duration)
	}
}

func TestToolTrackerLeavesUnmatchedCompletionUnchanged(t *testing.T) {
	complete := copilot.SessionEvent{
		Timestamp: time.Now(),
		Data: &copilot.ToolExecutionCompleteData{
			ToolCallID: "unmatched",
			Success:    true,
		},
	}

	var tracker ToolTracker
	details := Extract(complete)
	tracker.Enrich(complete, &details)

	if details.ToolName != nil {
		t.Errorf("ToolName = %q, want nil", *details.ToolName)
	}
	if details.Duration != nil {
		t.Errorf("Duration = %f, want nil", *details.Duration)
	}
}
