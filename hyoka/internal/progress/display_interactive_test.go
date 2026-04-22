package progress

import (
	"bytes"
	"strings"
	"testing"
)

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
