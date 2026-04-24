package progress

import (
	"bytes"
	"strings"
	"testing"
)

// TestInteractive_GraderPointsRendersNestedBlock exercises the Phase 2 nested
// rendering path: when len(Points) > 1, onGraderComplete writes a header line
// summarizing the aggregate (e.g. "❌ 2/3 passed") followed by one indented
// row per Point. Single- and zero-point graders should still use the legacy
// flat row.
func TestInteractive_GraderPointsRendersNestedBlock(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	id := "p/c"
	d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPhaseChange, Phase: PhaseGenerating})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "ai_review", GraderKind: "prompt_review"})
	d.HandleEvent(ProgressEvent{
		EvalID:     id,
		Type:       EventGraderComplete,
		GraderID:   "ai_review",
		GraderKind: "prompt_review",
		Result:     GraderResultFail,
		Points: []GraderPoint{
			{Name: "returns DefaultAzureCredential", Pass: true},
			{Name: "exposes get_secret/set_secret/delete_secret", Pass: true},
			{Name: "paginates list_secrets", Pass: false, Message: "missing pagination loop"},
		},
	})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()

	// Header summarizes aggregate pass/fail count.
	// Expect grader kind to display as "prompt" (mapped from "prompt_review")
	if !strings.Contains(out, "ai_review") || !strings.Contains(out, "(prompt)") {
		t.Errorf("missing grader header\n--- got ---\n%s", out)
	}
	if !strings.Contains(out, "2/3 passed") {
		t.Errorf("missing aggregate badge \"2/3 passed\"\n--- got ---\n%s", out)
	}
	// One indented row per Point.
	for _, want := range []string{
		"    - returns DefaultAzureCredential:",
		"    - exposes get_secret/set_secret/delete_secret:",
		"    - paginates list_secrets:",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing nested point row %q\n--- got ---\n%s", want, out)
		}
	}
	// The failing point should carry its message after the badge.
	if !strings.Contains(out, "missing pagination loop") {
		t.Errorf("expected failing point message in output\n--- got ---\n%s", out)
	}
}

// TestInteractive_GraderPointsAllPassedBadge verifies the all-passed aggregate
// formats as "✅ N/N passed".
func TestInteractive_GraderPointsAllPassedBadge(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	id := "p/c"
	d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPhaseChange, Phase: PhaseGenerating})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "no_secrets", GraderKind: "output_check"})
	d.HandleEvent(ProgressEvent{
		EvalID:     id,
		Type:       EventGraderComplete,
		GraderID:   "no_secrets",
		GraderKind: "output_check",
		Result:     GraderResultPass,
		Points: []GraderPoint{
			{Name: "min_files", Pass: true},
			{Name: "require_files", Pass: true},
		},
	})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	if !strings.Contains(out, "2/2 passed") {
		t.Errorf("missing all-passed aggregate badge \"2/2 passed\"\n--- got ---\n%s", out)
	}
	if !strings.Contains(out, "    - min_files:") || !strings.Contains(out, "    - require_files:") {
		t.Errorf("missing nested point rows\n--- got ---\n%s", out)
	}
}

// TestInteractive_GraderSinglePointFlatRow ensures the legacy flat-row path
// is preserved when len(Points) <= 1, so existing UX doesn't regress.
func TestInteractive_GraderSinglePointFlatRow(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	id := "p/c"
	d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPhaseChange, Phase: PhaseGenerating})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "build", GraderKind: "program"})
	d.HandleEvent(ProgressEvent{
		EvalID:     id,
		Type:       EventGraderComplete,
		GraderID:   "build",
		GraderKind: "program",
		Result:     GraderResultPass,
		Points: []GraderPoint{
			{Name: "exit code 0", Pass: true},
		},
	})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	// Single-Point case must NOT emit the aggregate badge inside a grader
	// header row (the run summary line "Summary: 1/1 passed" is unrelated
	// and exists in every transcript).
	if strings.Contains(out, "(program): ✅ 1/1 passed") || strings.Contains(out, "(program): ❌ 1/1 passed") {
		t.Errorf("single-point grader should use flat row, not aggregate badge\n--- got ---\n%s", out)
	}
	if strings.Contains(out, "    - exit code 0:") {
		t.Errorf("single-point grader should not emit nested rows\n--- got ---\n%s", out)
	}
	// But the flat row should still be present.
	if !strings.Contains(out, "build") || !strings.Contains(out, "(program)") {
		t.Errorf("missing flat grader row\n--- got ---\n%s", out)
	}
}
