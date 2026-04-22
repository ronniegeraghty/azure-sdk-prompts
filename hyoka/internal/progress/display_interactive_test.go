package progress

import (
	"bytes"
	"strings"
	"testing"
)

// feedInteractive builds an interactive Display pointed at buf and drives the
// provided event sequence through HandleEvent + Finish. Centralized so the
// table-driven cases below stay readable.
func feedInteractive(buf *bytes.Buffer, events []ProgressEvent) {
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: buf, Mode: ModeInteractive})
	for _, e := range events {
		d.HandleEvent(e)
	}
	d.Finish()
}

// floatPtr returns a pointer to the given float64 — convenience for embedding
// grader scores inline in test tables.
func floatPtr(v float64) *float64 { return &v }

// TestInteractive_HappyPath exercises the interactive renderer through a
// representative event stream and asserts the frozen transcript contains the
// expected section headers and final-state markers.
func TestInteractive_HappyPath(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	evalID := "p1/cfg1"
	d.HandleEvent(ProgressEvent{EvalID: evalID, PromptID: "p1", ConfigName: "cfg1", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionStart, ToolName: "azure-mcp", ToolKind: ToolKindMCP})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionResult, ToolName: "azure-mcp", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionStart, ToolName: "plugin-x", ToolKind: ToolKindPlugin})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionResult, ToolName: "plugin-x", ToolKind: ToolKindPlugin, Status: ToolStatusFailed, Reason: "not found"})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPhaseChange, Phase: PhaseGenerating})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolStart, Message: "bash"})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventWritingFile, Message: "write src/a.py"})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventSessionDetails, Files: []string{"src/a.py", "tests/t.py"}, Turns: 5, ToolCalls: 2, Cost: 0.03})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventGraderStart, GraderID: "prompt_review", GraderKind: "claude-opus-4.6"})
	score := 8.0
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventGraderComplete, GraderID: "prompt_review", GraderKind: "claude-opus-4.6", Result: GraderResultPass, Score: &score})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPassed, FileCount: 2})
	d.Finish()

	out := buf.String()
	for _, want := range []string{
		"Prompt: p1",
		"Config: cfg1",
		"Tools:",
		"azure-mcp",
		"plugin-x",
		"Failed",
		"Agent Attempt:",
		"Session Details:",
		"Graders:",
		"prompt_review",
		"Pass",
		"Summary: 1/1 passed",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n--- got ---\n%s", want, out)
		}
	}
}

// TestInteractive_NoToolsNoGraders verifies that the Tools and Graders
// section headers are omitted entirely when no events for those sections
// fire.
func TestInteractive_NoToolsNoGraders(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	evalID := "p2/cfg2"
	d.HandleEvent(ProgressEvent{EvalID: evalID, PromptID: "p2", ConfigName: "cfg2", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPhaseChange, Phase: PhaseGenerating})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	if strings.Contains(out, "Tools:") {
		t.Errorf("did not expect Tools: header; got:\n%s", out)
	}
	if strings.Contains(out, "Graders:") {
		t.Errorf("did not expect Graders: header; got:\n%s", out)
	}
}

// TestInteractive_ToolsVerifiedFlip verifies that the bulk verification event
// updates a previously "loaded" tool line when the SDK reports a failure.
// The redraw uses DECSC/DECRC, so the output buffer will contain the new
// line text — we just assert the final status appears.
func TestInteractive_ToolsVerifiedFlip(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	evalID := "p3/cfg3"
	d.HandleEvent(ProgressEvent{EvalID: evalID, PromptID: "p3", ConfigName: "cfg3", Type: EventStarting})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionStart, ToolName: "mcp-a", ToolKind: ToolKindMCP})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolResolutionResult, ToolName: "mcp-a", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPhaseChange, Phase: PhaseGenerating})
	// SDK reports mcp-a as failed — must trigger block redraw.
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventToolsVerified, Tools: []ToolStatus{{ToolName: "mcp-a", ToolKind: ToolKindMCP, Status: ToolStatusFailed, Reason: "runtime missing"}}})
	d.HandleEvent(ProgressEvent{EvalID: evalID, Type: EventPassed, FileCount: 0})
	d.Finish()

	out := buf.String()
	if !strings.Contains(out, "runtime missing") {
		t.Errorf("expected failure reason in redrawn tool block; got:\n%s", out)
	}
}

// TestInteractive_Cases is a table-driven sweep of the renderer's observable
// behavior across the scenarios listed in the tests-renderer-snapshots spec.
// The interactive renderer is animation-heavy (tail rewrites), so full-buffer
// byte-for-byte snapshots are brittle; instead each case asserts on the set
// of visible substrings that survive into the final buffer state.
func TestInteractive_Cases(t *testing.T) {
pass := floatPtr(8.0)
fail := floatPtr(3.0)

cases := []struct {
name        string
events      []ProgressEvent
wantSubs    []string // substrings that MUST appear
notWantSubs []string // substrings that must NOT appear
}{
{
name: "happy_path_one_tool_two_graders",
events: []ProgressEvent{
{EvalID: "e1", PromptID: "p1", ConfigName: "cfg1", Type: EventStarting},
{EvalID: "e1", Type: EventToolResolutionStart, ToolName: "azure-mcp", ToolKind: ToolKindMCP},
{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "azure-mcp", ToolKind: ToolKindMCP, Status: ToolStatusLoaded},
{EvalID: "e1", Type: EventPhaseChange, Phase: PhaseGenerating},
{EvalID: "e1", Type: EventToolStart, Message: "bash"},
{EvalID: "e1", Type: EventSessionDetails, Files: []string{"a.py"}, Turns: 4, ToolCalls: 1, Cost: 0.02},
{EvalID: "e1", Type: EventGraderStart, GraderID: "prompt_review", GraderKind: "claude-opus"},
{EvalID: "e1", Type: EventGraderComplete, GraderID: "prompt_review", GraderKind: "claude-opus", Result: GraderResultPass, Score: pass},
{EvalID: "e1", Type: EventGraderStart, GraderID: "no_secrets", GraderKind: "output_check"},
{EvalID: "e1", Type: EventGraderComplete, GraderID: "no_secrets", GraderKind: "output_check", Result: GraderResultPass, Score: pass},
{EvalID: "e1", Type: EventPassed, FileCount: 1},
},
wantSubs: []string{
"Prompt: p1",
"Config: cfg1",
"Tools:",
"azure-mcp",
"✅ Loaded",
"Agent Attempt:",
"✅ Complete",
"Session Details:",
"Files: a.py",
"Graders:",
"prompt_review",
"no_secrets",
"✅ Pass",
"Summary: 1/1 passed",
},
notWantSubs: []string{"❌ Fail"},
},
{
name: "tool_load_failure_at_resolution",
events: []ProgressEvent{
{EvalID: "e2", PromptID: "p2", ConfigName: "cfg2", Type: EventStarting},
{EvalID: "e2", Type: EventToolResolutionStart, ToolName: "plugin-x", ToolKind: ToolKindPlugin},
{EvalID: "e2", Type: EventToolResolutionResult, ToolName: "plugin-x", ToolKind: ToolKindPlugin, Status: ToolStatusFailed, Reason: "not found"},
{EvalID: "e2", Type: EventPhaseChange, Phase: PhaseGenerating},
{EvalID: "e2", Type: EventPassed, FileCount: 0},
},
wantSubs: []string{
"Tools:",
"plugin-x",
"❌ Failed",
"not found",
},
},
{
name: "tools_verified_flip_loaded_to_failed",
events: []ProgressEvent{
{EvalID: "e3", PromptID: "p3", ConfigName: "cfg3", Type: EventStarting},
{EvalID: "e3", Type: EventToolResolutionStart, ToolName: "mcp-a", ToolKind: ToolKindMCP},
{EvalID: "e3", Type: EventToolResolutionResult, ToolName: "mcp-a", ToolKind: ToolKindMCP, Status: ToolStatusLoaded},
{EvalID: "e3", Type: EventPhaseChange, Phase: PhaseGenerating},
{EvalID: "e3", Type: EventToolsVerified, Tools: []ToolStatus{{ToolName: "mcp-a", ToolKind: ToolKindMCP, Status: ToolStatusFailed, Reason: "runtime missing"}}},
{EvalID: "e3", Type: EventPassed, FileCount: 0},
},
wantSubs: []string{
"mcp-a",
"❌ Failed",
"runtime missing",
},
},
{
name: "grader_fail_one_pass_one_fail",
events: []ProgressEvent{
{EvalID: "e4", PromptID: "p4", ConfigName: "cfg4", Type: EventStarting},
{EvalID: "e4", Type: EventPhaseChange, Phase: PhaseGenerating},
{EvalID: "e4", Type: EventGraderStart, GraderID: "prompt_review", GraderKind: "claude-opus"},
{EvalID: "e4", Type: EventGraderComplete, GraderID: "prompt_review", GraderKind: "claude-opus", Result: GraderResultPass, Score: pass},
{EvalID: "e4", Type: EventGraderStart, GraderID: "output_check", GraderKind: "no_secrets"},
{EvalID: "e4", Type: EventGraderComplete, GraderID: "output_check", GraderKind: "no_secrets", Result: GraderResultFail, Score: fail, Message: "leaked API key"},
{EvalID: "e4", Type: EventFailed, Message: "grader output_check failed"},
},
wantSubs: []string{
"Graders:",
"prompt_review",
"✅ Pass",
"output_check",
"❌ Fail",
"leaked API key",
"grader output_check failed",
"Summary: 0/1 passed",
},
},
{
name: "error_path_generation_error",
events: []ProgressEvent{
{EvalID: "e5", PromptID: "p5", ConfigName: "cfg5", Type: EventStarting},
{EvalID: "e5", Type: EventPhaseChange, Phase: PhaseGenerating},
{EvalID: "e5", Type: EventToolStart, Message: "bash"},
{EvalID: "e5", Type: EventError, Message: "copilot session terminated"},
},
wantSubs: []string{
"Prompt: p5",
"Agent Attempt:",
"❌ Failed", // agentComplete(false) path
"copilot session terminated",
"1 errors",
},
},
}

for _, tc := range cases {
t.Run(tc.name, func(t *testing.T) {
var buf bytes.Buffer
feedInteractive(&buf, tc.events)
out := buf.String()
for _, sub := range tc.wantSubs {
if !strings.Contains(out, sub) {
t.Errorf("missing substring %q\n--- output ---\n%s", sub, out)
}
}
for _, sub := range tc.notWantSubs {
if strings.Contains(out, sub) {
t.Errorf("unexpected substring %q found\n--- output ---\n%s", sub, out)
}
}
})
}
}

// TestInteractive_ANSIMarkers exercises the raw ANSI escape sequences emitted
// by the tail-update and tools-block-redraw paths. These escapes are
// independent of the color Styler (they come from direct fmt.Fprintf against
// the constants in display_interactive.go), so they appear even when the
// writer is a bytes.Buffer.
func TestInteractive_ANSIMarkers(t *testing.T) {
// (1) Tail-update escape sequence "\r\x1b[2K" must appear when a tool
// line flips from Loading → Loaded (rewriteTail path).
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "t1", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventToolResolutionStart, ToolName: "m", ToolKind: ToolKindMCP})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventToolResolutionResult, ToolName: "m", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventPassed, FileCount: 0})
d.Finish()
tailMarker := "\r\x1b[2K"
if !strings.Contains(buf.String(), tailMarker) {
t.Errorf("expected tail-update escape %q in output; got:\n%q", tailMarker, buf.String())
}

// (2) DECSC (\x1b7) + DECRC (\x1b8) pair must bracket the tool-block
// redraw triggered by a ToolsVerified flip.
var buf2 bytes.Buffer
d2 := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf2, Mode: ModeInteractive})
d2.HandleEvent(ProgressEvent{EvalID: "t2", PromptID: "p", ConfigName: "c", Type: EventStarting})
d2.HandleEvent(ProgressEvent{EvalID: "t2", Type: EventToolResolutionStart, ToolName: "m", ToolKind: ToolKindMCP})
d2.HandleEvent(ProgressEvent{EvalID: "t2", Type: EventToolResolutionResult, ToolName: "m", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
d2.HandleEvent(ProgressEvent{EvalID: "t2", Type: EventPhaseChange, Phase: PhaseGenerating})
d2.HandleEvent(ProgressEvent{EvalID: "t2", Type: EventToolsVerified, Tools: []ToolStatus{
{ToolName: "m", ToolKind: ToolKindMCP, Status: ToolStatusFailed, Reason: "missing"},
}})
d2.HandleEvent(ProgressEvent{EvalID: "t2", Type: EventPassed, FileCount: 0})
d2.Finish()
out2 := buf2.String()
if !strings.Contains(out2, "\x1b7") {
t.Errorf("expected DECSC (\\x1b7) in redraw; got:\n%q", out2)
}
if !strings.Contains(out2, "\x1b8") {
t.Errorf("expected DECRC (\\x1b8) in redraw; got:\n%q", out2)
}
// DECSC must precede DECRC within the redraw sequence.
if i, j := strings.Index(out2, "\x1b7"), strings.LastIndex(out2, "\x1b8"); i < 0 || j < 0 || i >= j {
t.Errorf("DECSC/DECRC ordering wrong: save=%d restore=%d", i, j)
}
}

// TestInteractive_NoColorEnvDropsColor verifies that setting NO_COLOR=1
// yields output free of ANSI color codes (SGR codes). Cursor-move escapes
// (\r, \x1b[2K, \x1b7/\x1b8) still appear because those are not color — the
// styler only gates color codes. This matches the no-color.org convention
// implemented in style.detectEnabled.
func TestInteractive_NoColorEnvDropsColor(t *testing.T) {
t.Setenv("NO_COLOR", "1")
var buf bytes.Buffer
feedInteractive(&buf, []ProgressEvent{
{EvalID: "nc", PromptID: "p", ConfigName: "c", Type: EventStarting},
{EvalID: "nc", Type: EventToolResolutionStart, ToolName: "m", ToolKind: ToolKindMCP},
{EvalID: "nc", Type: EventToolResolutionResult, ToolName: "m", ToolKind: ToolKindMCP, Status: ToolStatusLoaded},
{EvalID: "nc", Type: EventGraderStart, GraderID: "g", GraderKind: "k"},
{EvalID: "nc", Type: EventGraderComplete, GraderID: "g", GraderKind: "k", Result: GraderResultPass, Score: floatPtr(9)},
{EvalID: "nc", Type: EventPassed, FileCount: 1},
})
out := buf.String()
// No SGR color codes: red/green/yellow/cyan/dim/bold must all be absent.
for _, code := range []string{"\x1b[31m", "\x1b[32m", "\x1b[33m", "\x1b[34m", "\x1b[36m", "\x1b[2m", "\x1b[1m"} {
if strings.Contains(out, code) {
t.Errorf("NO_COLOR output should not contain SGR code %q; got:\n%q", code, out)
}
}
// Plain status words still render.
if !strings.Contains(out, "✅ Loaded") {
t.Errorf("expected ✅ Loaded glyph+text; got:\n%s", out)
}
if !strings.Contains(out, "Summary: 1/1 passed") {
t.Errorf("expected summary; got:\n%s", out)
}
}
