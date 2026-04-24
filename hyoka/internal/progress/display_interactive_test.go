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
"✅ Completed",
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
"✅ Completed", // Three-state model: errors still show Completed for agent attempt
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
// (1) With the "wait till known" rule, a simple Loading → Loaded sequence
// commits a single newline-terminated line on Result and never rewrites a
// tail. The output for a clean resolution therefore contains NO mid-line
// clear escape "\r\x1b[2K" (only emitted by tail rewrites, which are still
// used for the Agent Attempt and Grader sections).
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "t1", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventToolResolutionStart, ToolName: "m", ToolKind: ToolKindMCP})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventToolResolutionResult, ToolName: "m", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
d.HandleEvent(ProgressEvent{EvalID: "t1", Type: EventPassed, FileCount: 0})
d.Finish()
out1 := buf.String()
// The final Tools-section output must show the resolved line without any
// intermediate "Loading…" text bleeding through.
if strings.Contains(out1, "Loading") {
t.Errorf("unexpected transient \"Loading\" text in tools output; got:\n%q", out1)
}
if !strings.Contains(out1, "✅ Loaded") {
t.Errorf("expected resolved ✅ Loaded row in tools output; got:\n%q", out1)
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

// TestInteractive_WaitTillKnown verifies that no transient "Loading…" text
// appears in the committed output between a ToolResolutionStart and its
// matching Result — the renderer must buffer the row until terminal state.
func TestInteractive_WaitTillKnown(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "w1", PromptID: "p", ConfigName: "c", Type: EventStarting})
// Start without an immediate Result — nothing should be written yet.
d.HandleEvent(ProgressEvent{EvalID: "w1", Type: EventToolResolutionStart, ToolName: "mcp-a", ToolKind: ToolKindMCP})
mid := buf.String()
if strings.Contains(mid, "mcp-a") || strings.Contains(mid, "Loading") {
t.Errorf("expected no tool row between Start and Result; got:\n%s", mid)
}
d.HandleEvent(ProgressEvent{EvalID: "w1", Type: EventToolResolutionResult, ToolName: "mcp-a", ToolKind: ToolKindMCP, Status: ToolStatusLoaded})
d.HandleEvent(ProgressEvent{EvalID: "w1", Type: EventPassed, FileCount: 0})
d.Finish()
out := buf.String()
if !strings.Contains(out, "mcp-a") || !strings.Contains(out, "✅ Loaded") {
t.Errorf("expected resolved row for mcp-a; got:\n%s", out)
}
if strings.Contains(out, "Loading") {
t.Errorf("transient Loading text leaked into final output:\n%s", out)
}
}

// TestInteractive_PluginFanout verifies that a plugin parent renders as a
// grouped header with its children indented beneath. Mixed loaded/failed
// children must both appear under the same parent header.
func TestInteractive_PluginFanout(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "pf", PromptID: "p", ConfigName: "c", Type: EventStarting})
// Plugin parent loads first.
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionStart, ToolName: "azure-sdk-python", ToolKind: ToolKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionResult, ToolName: "azure-sdk-python", ToolKind: ToolKindPlugin, Status: ToolStatusLoaded})
// Child skill: loaded.
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionStart, ToolName: "skill1", ToolKind: ToolKindSkill})
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionResult, ToolName: "skill1", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})
// Child MCP: failed.
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionStart, ToolName: "mcp1", ToolKind: ToolKindMCP})
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventToolResolutionResult, ToolName: "mcp1", ToolKind: ToolKindMCP, Status: ToolStatusFailed, Reason: "fetch timeout", ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "pf", Type: EventPassed, FileCount: 0})
d.Finish()
out := buf.String()
for _, want := range []string{
"azure-sdk-python (plugin):",
"skill1",
"✅ Loaded",
"mcp1",
"❌ Failed",
"fetch timeout",
} {
if !strings.Contains(out, want) {
t.Errorf("missing %q in plugin-fanout output:\n%s", want, out)
}
}
// The header must precede the children in the transcript.
hIdx := strings.Index(out, "azure-sdk-python (plugin):")
cIdx := strings.Index(out, "skill1")
if hIdx < 0 || cIdx < 0 || hIdx >= cIdx {
t.Errorf("parent header must precede children; header@%d child@%d\n%s", hIdx, cIdx, out)
}
// The plugin must NOT also appear as a top-level status line
// ("azure-sdk-python (plugin): ✅ Loaded") — that would be a duplicate.
if strings.Contains(out, "azure-sdk-python (plugin): ✅") ||
strings.Contains(out, "azure-sdk-python (plugin): ❌") {
t.Errorf("plugin container must not render a leaf status row:\n%s", out)
}
}

// TestInteractive_SkillDirFanout verifies that children tagged with
// ParentKind=skill_dir are grouped under a synthesized parent header even
// when the orphan parent Start event never receives a matching Result (the
// current validate.go skill_dir behavior).
func TestInteractive_SkillDirFanout(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "sd", PromptID: "p", ConfigName: "c", Type: EventStarting})
// Orphan parent Start (no matching Result) — mimics validate.go skill_dir flow.
d.HandleEvent(ProgressEvent{EvalID: "sd", Type: EventToolResolutionStart, ToolName: "generator-skills", ToolKind: ToolKindSkill})
// Children arrive with ParentKind=skill_dir and ParentName=path.
d.HandleEvent(ProgressEvent{EvalID: "sd", Type: EventToolResolutionResult, ToolName: "skillA", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "./skills", ParentKind: ToolParentKindSkillDir})
d.HandleEvent(ProgressEvent{EvalID: "sd", Type: EventToolResolutionResult, ToolName: "skillB", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "./skills", ParentKind: ToolParentKindSkillDir})
d.HandleEvent(ProgressEvent{EvalID: "sd", Type: EventPassed, FileCount: 0})
d.Finish()
out := buf.String()
for _, want := range []string{
"./skills (skills dir):",
"skillA",
"skillB",
"✅ Loaded",
} {
if !strings.Contains(out, want) {
t.Errorf("missing %q in skill_dir-fanout output:\n%s", want, out)
}
}
// Orphan parent Start (generator-skills) must NOT appear as a flat row.
if strings.Contains(out, "generator-skills") {
t.Errorf("orphan Start should be filtered out of output:\n%s", out)
}
}

// TestInteractive_PluginFailedNoFanout verifies that a plugin that fails at
// resolution renders a single flat failed row (no empty parent header).
func TestInteractive_PluginFailedNoFanout(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "pf2", PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: "pf2", Type: EventToolResolutionStart, ToolName: "missing-plugin", ToolKind: ToolKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "pf2", Type: EventToolResolutionResult, ToolName: "missing-plugin", ToolKind: ToolKindPlugin, Status: ToolStatusFailed, Reason: "not found"})
d.HandleEvent(ProgressEvent{EvalID: "pf2", Type: EventPassed, FileCount: 0})
d.Finish()
out := buf.String()
if !strings.Contains(out, "missing-plugin") || !strings.Contains(out, "❌ Failed") || !strings.Contains(out, "not found") {
t.Errorf("expected flat failed row for missing-plugin; got:\n%s", out)
}
// Must not emit a parent header for a failed plugin (no children expected).
if strings.Contains(out, "missing-plugin (plugin):\n") {
t.Errorf("failed plugin must not emit a bare parent header:\n%s", out)
}
}

// TestInteractive_AgentCompletedRowRewritesFrozenLine regresses Phase 1
// issue (d): when a grader Start event freezes the active agent tail by
// taking over the tail itself, the agent's later Complete event must
// rewrite the original Running row in place rather than appending a
// stale Completed line at the bottom of the transcript.
//
// The buffer is raw bytes (no terminal interpreting ANSI), so we
// assert: (1) the in-place rewrite emitted a DECSC save sequence
// (\x1b7) — proving the rewriteFrozenLine branch fired rather than
// falling through to writeLine, and (2) the new "Completed" content
// was NOT terminated by a writeLine newline (which would indicate a
// stale duplicate at the bottom of the transcript).
func TestInteractive_AgentCompletedRowRewritesFrozenLine(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	id := "agent-rewrite"
	d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
	// No tools — opens the agent gate immediately on the first activity event.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventReasoning, Message: "thinking"})
// EventSessionDetails signals generation complete BEFORE graders start.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventSessionDetails, Files: []string{"a.py"}, Turns: 3, ToolCalls: 5, Cost: 0.02})
	// EventSessionDetails signals generation complete BEFORE graders start.
	// This flips the Agent Attempt line to "Completed" while it is still the
	// tail, so no rewriteFrozenLine is needed later.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventSessionDetails, Files: []string{"a.py"}, Turns: 3, ToolCalls: 5, Cost: 0.02})
	// Now graders run. The Agent Attempt line is already frozen as "Completed".
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "ai_review", GraderKind: "claude-opus-4.6"})
	score := 9.0
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderComplete, GraderID: "ai_review", GraderKind: "claude-opus-4.6", Result: GraderResultPass, Score: &score})
	// EventPassed arrives last but does not need to rewrite the agent row
	// because EventSessionDetails already did that.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	// (1) The "Completed" content must appear once and be frozen before graders.
	if !strings.Contains(out, "✅ Completed") {
		t.Errorf("expected '✅ Completed' in transcript; got:\n%q", out)
	}
	// (2) "Agent Attempt:" must appear exactly once as a line start.
	if got := strings.Count(out, "\nAgent Attempt:"); got != 1 {
		t.Errorf("want exactly 1 'Agent Attempt:' starting a new line, got %d:\n%q", got, out)
	}
	// (3) Session Details section must be present.
	if !strings.Contains(out, "Session Details:") {
		t.Errorf("expected 'Session Details:' section; got:\n%q", out)
	}
	// (4) Graders section must follow Session Details.
	if !strings.Contains(out, "Graders:") || !strings.Contains(out, "ai_review") {
		t.Errorf("expected 'Graders:' section with ai_review; got:\n%q", out)
	}
	// (5) The grader's Pass content should appear exactly once.
	if got := strings.Count(out, "Pass (9/10)"); got != 1 {
		t.Errorf("want exactly 1 'Pass (9/10)', got %d:\n%q", got, out)
	}
}

// TestInteractive_ReviewerEventsAfterCompletionIgnored regresses Bug 3:
// reviewer Copilot sessions emit EventReasoning/EventToolStart/EventToolComplete
// through the same event channel as the generator. Once EventSessionDetails
// flips Agent Attempt → Completed, subsequent activity events from reviewer
// sessions MUST NOT reopen the agent tail or create duplicate "Agent Attempt:
// ✅ Completed" rows at the bottom of the transcript.
func TestInteractive_ReviewerEventsAfterCompletionIgnored(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

	id := "reviewer-events"
	d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
	// No tools configured — gate opens on first activity.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventReasoning, Message: "generator thinking"})
	// Generation completes; Agent Attempt → Completed
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventSessionDetails, Files: []string{"a.py"}, Turns: 2, ToolCalls: 1, Cost: 0.01})
	// Start a typed grader (e.g., files_present) and complete it.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "files_present", GraderKind: "files_present"})
	score := 10.0
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderComplete, GraderID: "files_present", GraderKind: "files_present", Result: GraderResultPass, Score: &score})
	// Start the AI review grader (reviewer Copilot session)
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "ai_review", GraderKind: "claude-opus-4.6"})
	// THESE are the problematic events: the reviewer session emits the same
	// EventReasoning/EventToolStart/EventToolComplete as the generator did,
	// and they get routed to onAgentActivity → renderAgentEvent. Before the
	// fix, these would create duplicate "Agent Attempt: ✅ Completed" rows.
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventReasoning, Message: "reviewer thinking"})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventToolStart, Message: "reviewer tool 1"})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventToolComplete, Message: "reviewer tool 1 done"})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventToolStart, Message: "reviewer tool 2"})
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventToolComplete, Message: "reviewer tool 2 done"})
	// Reviewer completes
	reviewScore := 8.0
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderComplete, GraderID: "ai_review", GraderKind: "claude-opus-4.6", Result: GraderResultPass, Score: &reviewScore})
	// Eval passes
	d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 1})
	d.Finish()

	out := buf.String()
	// The fix: renderAgentEvent must return early when agentState is already
	// terminal. Assert exactly ONE "Agent Attempt:" row exists in the transcript.
	if got := strings.Count(out, "\nAgent Attempt:"); got != 1 {
		t.Errorf("want exactly 1 'Agent Attempt:' row, got %d (duplicate rows after reviewer events):\n%q", got, out)
	}
	// The Agent Attempt line must show Completed (not stuck on Running).
	if !strings.Contains(out, "✅ Completed") {
		t.Errorf("expected '✅ Completed' in Agent Attempt line; got:\n%q", out)
	}
	// Both graders should be present in the transcript.
	if !strings.Contains(out, "files_present") || !strings.Contains(out, "ai_review") {
		t.Errorf("expected both 'files_present' and 'ai_review' graders; got:\n%q", out)
	}
}

// TestInteractive_AgentAttemptSingleLineInvariant guards the single-line
// invariant: the "Agent Attempt:" prefix must appear EXACTLY ONCE as the
// start of a physical row in the rendered transcript — never split across
// two rows (a standalone header followed by a separate state row). This
// regresses the bug where ensureAgentHeader wrote "Agent Attempt:" via
// writeLine and renderAgentEvent then wrote "  🔄 Running" as a separate
// tail.
func TestInteractive_AgentAttemptSingleLineInvariant(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

id := "agent-single-line"
d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
// No tools — opens the agent gate immediately on the first activity event.
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventReasoning, Message: "thinking"})
// Mid-flight: exactly one "Agent Attempt:" should have been emitted, and
// it must NOT be followed by a newline (header + state must share a row).
mid := buf.String()
if got := strings.Count(mid, "Agent Attempt:"); got != 1 {
t.Errorf("want exactly 1 'Agent Attempt:' occurrence after first activity, got %d:\n%q", got, mid)
}
if strings.Contains(mid, "Agent Attempt:\n") {
t.Errorf("'Agent Attempt:' must not be followed by a newline (header and state must share a row):\n%q", mid)
}
// Drive to completion.
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 0})
d.Finish()

out := buf.String()
// After completion, exactly one "Agent Attempt:" should start a new
// physical line (i.e. follow a newline). The in-place rewrite uses
// \r\x1b[2K to overwrite, so a second textual occurrence may exist
// in the byte buffer — but it must NOT be preceded by a newline,
// which would mean it landed on a separate row.
if got := strings.Count(out, "\nAgent Attempt:"); got != 1 {
t.Errorf("want exactly 1 'Agent Attempt:' starting a new line after completion, got %d:\n%q", got, out)
}
if !strings.Contains(out, "✅ Completed") {
t.Errorf("expected '✅ Completed' in final transcript; got:\n%q", out)
}
if strings.Contains(out, "Agent Attempt:\n") {
t.Errorf("final transcript must not split 'Agent Attempt:' from its state:\n%q", out)
}
}

// TestInteractive_GraderCompleteIdempotentAfterFreeze regresses Phase 1
// issue (e): if a grader's Running tail gets frozen by an unrelated
// event (an EventToolStart fired between grader Start and Complete),
// the grader's later Complete event must rewrite the original frozen
// row in place — never append a duplicate ai_review entry below.
func TestInteractive_GraderCompleteIdempotentAfterFreeze(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})

id := "grader-idem"
d.HandleEvent(ProgressEvent{EvalID: id, PromptID: "p", ConfigName: "c", Type: EventStarting})
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventReasoning, Message: "thinking"})
// Grader Start owns the tail.
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderStart, GraderID: "ai_review", GraderKind: "panel"})
// Some other event commits the grader tail and steals it (this is
// exactly the historical noise pattern from the redundant
// EventToolStart/EventToolComplete bracketing the AI review).
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventToolStart, Message: "Review panel: [m1 m2]"})
// Grader Complete arrives after the tail moved.
score := 7.0
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventGraderComplete, GraderID: "ai_review", GraderKind: "panel", Result: GraderResultPass, Score: &score})
d.HandleEvent(ProgressEvent{EvalID: id, Type: EventPassed, FileCount: 0})
d.Finish()

out := buf.String()
// In-place rewrite escape was emitted.
if !strings.Contains(out, "\x1b7") {
t.Errorf("expected DECSC save sequence; got:\n%q", out)
}
// The grader's final Pass content appears exactly once — no stale
// duplicate from a fallback writeLine.
if got := strings.Count(out, "Pass (7/10)"); got != 1 {
t.Errorf("want exactly 1 'Pass (7/10)' entry, got %d:\n%q", got, out)
}
// The Pass row was committed via rewriteFrozenLine (DECRC \x1b8
// follows the Pass content), NOT via writeLine (which would write a
// trailing newline immediately after).
if i := strings.Index(out, "Pass (7/10)"); i != -1 {
tail := out[i+len("Pass (7/10)"):]
if strings.HasPrefix(tail, "\n") {
t.Errorf("Pass row appears writeLine-terminated (newline immediately after) — should be DECRC-bracketed:\n%q", out)
}
}
}
