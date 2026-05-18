package progress

import (
	"bytes"
	"strings"
	"testing"
)

// TestInteractive_GraderPointsRendersNestedBlock exercises the Phase 2 nested
// rendering path: when len(Points) > 1, onGraderComplete writes a header line
// summarizing the aggregate (e.g. "❌ Fail (2/3)") followed by one indented
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
			{Label: "returns DefaultAzureCredential", Pass: true},
			{Label: "exposes get_secret/set_secret/delete_secret", Pass: true},
			{Label: "paginates list_secrets", Pass: false, Message: "missing pagination loop"},
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
	if !strings.Contains(out, "Fail (2/3)") {
		t.Errorf("missing aggregate badge \"Fail (2/3)\"\n--- got ---\n%s", out)
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
// formats as "✅ Pass (N/N)".
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
			{Label: "min_files", Pass: true},
			{Label: "require_files", Pass: true},
		},
	})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	if !strings.Contains(out, "Pass (2/2)") {
		t.Errorf("missing all-passed aggregate badge \"Pass (2/2)\"\n--- got ---\n%s", out)
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
			{Label: "exit code 0", Pass: true},
		},
	})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	// Single-Point case MUST now emit the aggregate badge format for
	// consistency with multi-point graders (changed from > 1 to >= 1).
	if !strings.Contains(out, "(program): ✅ Pass (1/1)") {
		t.Errorf("single-point grader should use badge format: (program): ✅ Pass (1/1)\n--- got ---\n%s", out)
	}
	// Sub-point should be rendered as an indented row
	if !strings.Contains(out, "- exit code 0:") {
		t.Errorf("single-point grader should emit nested row for the check\n--- got ---\n%s", out)
	}
	// Grader ID and kind should still be present
	if !strings.Contains(out, "build") {
		t.Errorf("missing grader ID\n--- got ---\n%s", out)
	}
}
