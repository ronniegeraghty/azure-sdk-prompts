package progress

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestNewDisplay_Disabled(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{
		Total:    5,
		Workers:  2,
		Writer:   &buf,
		Disabled: true,
	})
	d.HandleEvent(ProgressEvent{
		EvalID:     "test/config",
		PromptID:   "test",
		ConfigName: "config",
		Type:       EventStarting,
	})
	d.Done()
	if buf.Len() != 0 {
		t.Errorf("expected no output when disabled, got %d bytes", buf.Len())
	}
}

func TestDisplay_SlotManagement(t *testing.T) {
	var buf bytes.Buffer
	d := &Display{
		slots:     make([]slot, 2),
		total:     4,
		w:         &buf,
		disabled:  false,
		evalSlots: make(map[string]int),
		width:     120,
		startTime: time.Now(),
	}

	// Start two evals — should claim slots 0 and 1
	d.HandleEvent(ProgressEvent{
		EvalID: "p1/c1", PromptID: "p1", ConfigName: "c1",
		Type: EventStarting,
	})
	d.HandleEvent(ProgressEvent{
		EvalID: "p2/c2", PromptID: "p2", ConfigName: "c2",
		Type: EventStarting,
	})

	if len(d.evalSlots) != 2 {
		t.Errorf("expected 2 active slots, got %d", len(d.evalSlots))
	}
	if !d.slots[0].active || !d.slots[1].active {
		t.Error("expected both slots to be active")
	}

	// Complete first eval
	d.HandleEvent(ProgressEvent{
		EvalID: "p1/c1", Type: EventPassed, FileCount: 3,
	})
	if d.completed != 1 || d.passed != 1 {
		t.Errorf("expected 1 completed/passed, got %d/%d", d.completed, d.passed)
	}
	if d.slots[0].active {
		t.Error("expected slot 0 to be inactive after completion")
	}
	if !d.slots[0].completed {
		t.Error("expected slot 0 to be marked completed")
	}

	// New eval should claim the released slot 0
	d.HandleEvent(ProgressEvent{
		EvalID: "p3/c3", PromptID: "p3", ConfigName: "c3",
		Type: EventStarting,
	})
	if d.evalSlots["p3/c3"] != 0 {
		t.Errorf("expected new eval to claim slot 0, got slot %d", d.evalSlots["p3/c3"])
	}
}

func TestDisplay_EventIcons(t *testing.T) {
	var buf bytes.Buffer
	d := &Display{
		slots:     make([]slot, 1),
		total:     1,
		w:         &buf,
		disabled:  false,
		evalSlots: make(map[string]int),
		width:     120,
		startTime: time.Now(),
	}

	tests := []struct {
		evt  ProgressEvent
		icon string
	}{
		{ProgressEvent{EvalID: "p/c", PromptID: "p", ConfigName: "c", Type: EventStarting}, "⏳"},
		{ProgressEvent{EvalID: "p/c", Type: EventSendingPrompt, Message: "Sending prompt (100 chars)..."}, "→"},
		{ProgressEvent{EvalID: "p/c", Type: EventToolStart, Message: "bash → ls"}, "⚙"},
		{ProgressEvent{EvalID: "p/c", Type: EventToolComplete, Message: "bash"}, "✓"},
		{ProgressEvent{EvalID: "p/c", Type: EventReasoning, Message: "Reasoning..."}, "💭"},
		{ProgressEvent{EvalID: "p/c", Type: EventWritingFile, Message: "create → main.py"}, "📝"},
		{ProgressEvent{EvalID: "p/c", Type: EventWaiting, Message: "Waiting..."}, "⏳"},
	}

	for _, tt := range tests {
		d.HandleEvent(tt.evt)
		if d.slots[0].icon != tt.icon {
			t.Errorf("after event type %d: expected icon %q, got %q", tt.evt.Type, tt.icon, d.slots[0].icon)
		}
	}
}

func TestDisplay_CompletionCounts(t *testing.T) {
	var buf bytes.Buffer
	d := &Display{
		slots:     make([]slot, 3),
		total:     3,
		w:         &buf,
		disabled:  false,
		evalSlots: make(map[string]int),
		width:     120,
		startTime: time.Now(),
	}

	// Start 3 evals
	for i, id := range []string{"a/x", "b/y", "c/z"} {
		d.HandleEvent(ProgressEvent{
			EvalID:     id,
			PromptID:   string(rune('a' + i)),
			ConfigName: string(rune('x' + i)),
			Type:       EventStarting,
		})
	}

	d.HandleEvent(ProgressEvent{EvalID: "a/x", Type: EventPassed, FileCount: 2})
	d.HandleEvent(ProgressEvent{EvalID: "b/y", Type: EventFailed})
	d.HandleEvent(ProgressEvent{EvalID: "c/z", Type: EventError, Message: "timeout"})

	if d.passed != 1 {
		t.Errorf("expected 1 passed, got %d", d.passed)
	}
	if d.failed != 1 {
		t.Errorf("expected 1 failed, got %d", d.failed)
	}
	if d.errors != 1 {
		t.Errorf("expected 1 error, got %d", d.errors)
	}
	if d.completed != 3 {
		t.Errorf("expected 3 completed, got %d", d.completed)
	}
}

func TestDisplay_Done(t *testing.T) {
	var buf bytes.Buffer
	d := &Display{
		slots:     make([]slot, 1),
		total:     1,
		completed: 1,
		passed:    1,
		w:         &buf,
		disabled:  false,
		evalSlots: make(map[string]int),
		width:     120,
		startTime: time.Now(),
	}

	d.Done()

	output := buf.String()
	if !strings.Contains(output, "Complete: 1/1") {
		t.Errorf("expected Done output to contain 'Complete: 1/1', got %q", output)
	}
	if !strings.Contains(output, "1 passed") {
		t.Errorf("expected Done output to contain '1 passed', got %q", output)
	}
}

func TestFormatName_Short(t *testing.T) {
	d := &Display{width: 120}
	got := d.formatName("p1", "c1")
	if got != "p1/c1" {
		t.Errorf("expected 'p1/c1', got %q", got)
	}
}

func TestFormatName_Truncated(t *testing.T) {
	d := &Display{width: 120}
	got := d.formatName("very-long-prompt-id-that-exceeds-limit", "config-name")
	if len(got) > 38 {
		t.Errorf("expected name truncated to ≤38 chars, got %d: %q", len(got), got)
	}
	if !strings.HasSuffix(got, "..") {
		t.Errorf("expected truncated name to end with '..', got %q", got)
	}
}

func TestTermWidth_Default(t *testing.T) {
	w := TermWidth()
	if w <= 0 {
		t.Errorf("expected positive terminal width, got %d", w)
	}
}

func TestDisplay_ActivityTruncation(t *testing.T) {
	d := &Display{width: 80}
	maxW := d.activityWidth()
	longActivity := strings.Repeat("x", 200)
	truncated := d.truncateActivity(longActivity)
	if len(truncated) > maxW {
		t.Errorf("expected truncated activity ≤%d, got %d", maxW, len(truncated))
	}
	if !strings.HasSuffix(truncated, "...") {
		t.Errorf("expected truncated activity to end with '...', got %q", truncated)
	}
}

func TestDisplay_ClaimSlotPrefersEmpty(t *testing.T) {
	d := &Display{
		slots:     make([]slot, 3),
		evalSlots: make(map[string]int),
	}

	// Empty slots available
	idx := d.claimSlot()
	if idx != 0 {
		t.Errorf("expected first empty slot (0), got %d", idx)
	}

	// Mark slot 0 as completed, slot 1 empty
	d.slots[0] = slot{evalID: "x", completed: true}
	idx = d.claimSlot()
	if idx != 1 {
		t.Errorf("expected empty slot 1 over completed slot 0, got %d", idx)
	}

	// All have evalIDs but slot 0 is completed (inactive)
	d.slots[1] = slot{evalID: "y", active: true}
	d.slots[2] = slot{evalID: "z", active: true}
	idx = d.claimSlot()
	if idx != 0 {
		t.Errorf("expected completed slot 0, got %d", idx)
	}
}
