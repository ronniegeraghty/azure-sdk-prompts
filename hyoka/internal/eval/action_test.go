package eval

import (
	"testing"
	"time"
)

func TestNewActionTimeline_Empty(t *testing.T) {
	tl := NewActionTimeline(nil)
	if tl == nil {
		t.Fatal("expected non-nil timeline")
	}
	if len(tl.Events) != 0 {
		t.Errorf("expected 0 events, got %d", len(tl.Events))
	}
	if tl.TotalTurns != 0 {
		t.Errorf("expected 0 turns, got %d", tl.TotalTurns)
	}
	if tl.TotalDuration != 0 {
		t.Errorf("expected 0 duration, got %v", tl.TotalDuration)
	}
	if tl.Summary != (TimelineSummary{}) {
		t.Errorf("expected zero summary, got %+v", tl.Summary)
	}
}

func TestNewActionTimeline_SummaryCounts(t *testing.T) {
	events := []ActionEvent{
		{Type: "tool_call", TurnNumber: 1, Duration: 100 * time.Millisecond, Success: true},
		{Type: "file_read", TurnNumber: 1, Duration: 50 * time.Millisecond, Success: true},
		{Type: "file_write", TurnNumber: 2, Duration: 200 * time.Millisecond, Success: true},
		{Type: "bash", TurnNumber: 2, Duration: 300 * time.Millisecond, Success: true},
		{Type: "bash", TurnNumber: 3, Duration: 150 * time.Millisecond, Success: false, Error: "exit 1"},
		{Type: "tool_call", TurnNumber: 3, Duration: 75 * time.Millisecond, Success: true},
		{Type: "response", TurnNumber: 3, Duration: 10 * time.Millisecond, Success: true},
	}

	tl := NewActionTimeline(events)

	if tl.TotalTurns != 3 {
		t.Errorf("expected 3 turns, got %d", tl.TotalTurns)
	}

	wantDur := 885 * time.Millisecond
	if tl.TotalDuration != wantDur {
		t.Errorf("expected total duration %v, got %v", wantDur, tl.TotalDuration)
	}

	if tl.Summary.ToolCalls != 2 {
		t.Errorf("expected 2 tool_calls, got %d", tl.Summary.ToolCalls)
	}
	if tl.Summary.FileReads != 1 {
		t.Errorf("expected 1 file_read, got %d", tl.Summary.FileReads)
	}
	if tl.Summary.FileWrites != 1 {
		t.Errorf("expected 1 file_write, got %d", tl.Summary.FileWrites)
	}
	if tl.Summary.BashCmds != 2 {
		t.Errorf("expected 2 bash commands, got %d", tl.Summary.BashCmds)
	}
}

func TestNewActionTimeline_ResponseTypeNotCounted(t *testing.T) {
	events := []ActionEvent{
		{Type: "response", TurnNumber: 1, Duration: 10 * time.Millisecond},
	}
	tl := NewActionTimeline(events)

	if tl.Summary.ToolCalls != 0 || tl.Summary.FileReads != 0 ||
		tl.Summary.FileWrites != 0 || tl.Summary.BashCmds != 0 {
		t.Errorf("response events should not increment summary counts: %+v", tl.Summary)
	}
}

func TestTruncateField_NoTruncation(t *testing.T) {
	input := "short string"
	got := TruncateField(input, 100)
	if got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}

func TestTruncateField_ExactLimit(t *testing.T) {
	input := "exact"
	got := TruncateField(input, 5)
	if got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}

func TestTruncateField_Truncates(t *testing.T) {
	input := "this is a long string that should be truncated at some point"
	got := TruncateField(input, 30)

	if len(got) > 30 {
		t.Errorf("expected length <= 30, got %d: %q", len(got), got)
	}
	if got[len(got)-1] != ']' {
		t.Errorf("expected truncated suffix, got %q", got)
	}
}

func TestTruncateField_ZeroMax(t *testing.T) {
	got := TruncateField("hello", 0)
	if got != "hello" {
		t.Errorf("expected unchanged string for maxLen=0, got %q", got)
	}
}

func TestTruncateField_NegativeMax(t *testing.T) {
	got := TruncateField("hello", -1)
	if got != "hello" {
		t.Errorf("expected unchanged string for negative maxLen, got %q", got)
	}
}

func TestTruncateField_VerySmallMax(t *testing.T) {
	got := TruncateField("hello world", 3)
	if len(got) != 3 {
		t.Errorf("expected length 3, got %d: %q", len(got), got)
	}
}

func TestTruncateField_EmptyString(t *testing.T) {
	got := TruncateField("", 10)
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}
