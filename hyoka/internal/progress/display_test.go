package progress

import (
"bytes"
"strings"
"testing"
)

func TestDisplay_BasicFlow(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 2, Workers: 2, Writer: &buf, Disabled: false})

d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p1", ConfigName: "c1", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "b", PromptID: "p2", ConfigName: "c2", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 3, ReviewScore: 8})
d.HandleEvent(ProgressEvent{EvalID: "b", Type: EventFailed, Message: "bad code"})
d.Finish()

out := buf.String()
if !strings.Contains(out, "Prompt: p1") {
t.Errorf("expected 'Prompt: p1' in output, got %q", out)
}
if !strings.Contains(out, "Config: c1") {
t.Errorf("expected 'Config: c1' in output, got %q", out)
}
if !strings.Contains(out, "Prompt: p2") {
t.Errorf("expected 'Prompt: p2' in output, got %q", out)
}
if !strings.Contains(out, "✅") {
t.Errorf("expected ✅ in output")
}
if !strings.Contains(out, "❌") {
t.Errorf("expected ❌ in output")
}
if !strings.Contains(out, "3 files") {
t.Errorf("expected '3 files' in output")
}
if !strings.Contains(out, "8/10") {
t.Errorf("expected '8/10' in output")
}
if !strings.Contains(out, "Summary: 1/2") {
t.Errorf("expected summary in output")
}
}

func TestDisplay_CompletedCount(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 3, Workers: 2, Writer: &buf, Disabled: false})
d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 1})
if d.CompletedEvalCount() != 1 {
t.Errorf("expected 1 completed, got %d", d.CompletedEvalCount())
}
}

func TestDisplay_ReportDir(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, ReportDir: "reports/123/"})
d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 2})
d.Finish()
if !strings.Contains(buf.String(), "reports/123/") {
t.Errorf("expected report dir in output")
}
}

func TestDisplay_Disabled(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Disabled: true})
d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 1})
d.Finish()
if buf.Len() != 0 {
t.Errorf("expected no output when disabled, got %q", buf.String())
}
}

func TestDisplay_LogMode(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 2, Workers: 2, Configs: 1, Writer: &buf, Mode: ModeLog})

d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p1", ConfigName: "c1", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderStart, GraderID: "g1"})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultPass})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 2})
d.HandleEvent(ProgressEvent{EvalID: "b", PromptID: "p2", ConfigName: "c1", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "b", Type: EventFailed, Message: "boom"})
d.Finish()

out := buf.String()
if !strings.Contains(out, "Running 2 evals") {
	t.Errorf("log/CI mode should show intro line, got %q", out)
}
if !strings.Contains(out, "p1") || !strings.Contains(out, "c1") {
	t.Errorf("log/CI mode should show prompt and config, got %q", out)
}
// Start glyph: emoji when color enabled, "START" text otherwise. Buffer
// writers always disable color, so expect the text form.
if !strings.Contains(out, "START") {
	t.Errorf("log/CI mode should show START glyph (color disabled on buffer), got %q", out)
}
if !strings.Contains(out, "PASS") {
	t.Errorf("log/CI mode should show PASS result, got %q", out)
}
if !strings.Contains(out, "FAIL") {
	t.Errorf("log/CI mode should show FAIL result, got %q", out)
}
if !strings.Contains(out, "boom") {
	t.Errorf("log/CI mode should surface failure reason, got %q", out)
}
if !strings.Contains(out, "Summary") {
	t.Errorf("log/CI mode should print summary table header, got %q", out)
}
if !strings.Contains(out, "1/2 passed") {
	t.Errorf("log/CI mode should print footer totals, got %q", out)
}
}

func TestDisplay_CIMode(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Configs: 1, Writer: &buf, Mode: ModeCI, ReportDir: "reports/x/"})
d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderStart, GraderID: "g1"})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderStart, GraderID: "g2"})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultPass})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventGraderComplete, GraderID: "g2", Result: GraderResultFail})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventFailed, Message: "graders\nfailed: g2"})
d.Finish()

out := buf.String()
if !strings.Contains(out, "1/2 graders") {
	t.Errorf("CI mode should include grader pass/total in finish line, got %q", out)
}
if !strings.Contains(out, "graders failed: g2") {
	t.Errorf("CI mode should collapse multi-line failure reason, got %q", out)
}
if !strings.Contains(out, "reports/x/") {
	t.Errorf("CI mode footer should include report path, got %q", out)
}
if !strings.Contains(out, "│") || !strings.Contains(out, "┌") {
	t.Errorf("CI mode should render unicode box-drawing summary table, got %q", out)
}
}

func TestDisplay_OffMode(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 2, Workers: 2, Writer: &buf, Mode: ModeOff})

d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "p1", ConfigName: "c1", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 1})
d.Finish()

if buf.Len() != 0 {
	t.Errorf("off mode should produce no output, got %q", buf.String())
}
}

func TestDisplay_SectionLayout(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Disabled: false})

d.HandleEvent(ProgressEvent{EvalID: "a", PromptID: "kv-crud", ConfigName: "baseline/opus", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPhaseChange, Phase: PhaseGenerating})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPhaseChange, Phase: PhaseReviewing})
d.HandleEvent(ProgressEvent{EvalID: "a", Type: EventPassed, FileCount: 5, ReviewScore: 9})
d.Finish()

out := buf.String()
if !strings.Contains(out, "Prompt: kv-crud") {
t.Errorf("expected 'Prompt: kv-crud', got %q", out)
}
if !strings.Contains(out, "Config: baseline/opus") {
t.Errorf("expected 'Config: baseline/opus', got %q", out)
}
if !strings.Contains(out, "5 files") {
t.Errorf("expected '5 files', got %q", out)
}
if !strings.Contains(out, "9/10") {
t.Errorf("expected '9/10', got %q", out)
}
}
