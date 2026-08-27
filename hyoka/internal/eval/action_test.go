package eval

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

func boolPtr(b bool) *bool { return &b }

func TestNewActionTimeline(t *testing.T) {
	tl := NewActionTimeline()
	if tl == nil {
		t.Fatal("NewActionTimeline returned nil")
	}
	if len(tl.Events) != 0 {
		t.Errorf("expected 0 events, got %d", len(tl.Events))
	}
	if tl.Summary.ToolBreakdown == nil {
		t.Error("expected non-nil ToolBreakdown map")
	}
}

func TestBuildActionTimeline_Empty(t *testing.T) {
	tl := BuildActionTimeline(nil)
	if tl == nil {
		t.Fatal("BuildActionTimeline(nil) returned nil")
	}
	if len(tl.Events) != 0 {
		t.Errorf("expected 0 events, got %d", len(tl.Events))
	}
	if tl.Summary.TotalEvents != 0 {
		t.Errorf("expected 0 total events, got %d", tl.Summary.TotalEvents)
	}
}

func TestBuildActionTimeline_ToolCalls(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "tool.execution_start", ToolName: "view", FilePath: "/workspace/main.py"},
		{Type: "tool.execution_complete", ToolName: "view", Duration: 50, ToolSuccess: boolPtr(true)},
		{Type: "tool.execution_start", ToolName: "create", FilePath: "/workspace/out.py", ToolArgs: `{"path":"/workspace/out.py"}`},
		{Type: "tool.execution_complete", ToolName: "create", Duration: 120, ToolSuccess: boolPtr(true)},
		{Type: "tool.execution_start", ToolName: "bash", ToolArgs: `{"command":"python main.py"}`},
		{Type: "tool.execution_complete", ToolName: "bash", Duration: 300, ToolSuccess: boolPtr(true), ToolResult: "OK"},
		{Type: "assistant.turn_end", TurnNumber: 1, Duration: 500},
	}

	tl := BuildActionTimeline(records)

	if tl.Summary.TotalEvents != 8 {
		t.Errorf("expected 8 total events, got %d", tl.Summary.TotalEvents)
	}
	if tl.Summary.TotalTurns != 1 {
		t.Errorf("expected 1 turn, got %d", tl.Summary.TotalTurns)
	}
	if tl.Summary.FileReads != 1 {
		t.Errorf("expected 1 file read, got %d", tl.Summary.FileReads)
	}
	if tl.Summary.FileWrites != 1 {
		t.Errorf("expected 1 file write, got %d", tl.Summary.FileWrites)
	}
	if tl.Summary.BashCommands != 1 {
		t.Errorf("expected 1 bash command, got %d", tl.Summary.BashCommands)
	}
	if tl.Summary.Errors != 0 {
		t.Errorf("expected 0 errors, got %d", tl.Summary.Errors)
	}
	// TotalToolCalls = file_read(1) + file_write(1) + bash(1)
	if tl.Summary.TotalToolCalls != 3 {
		t.Errorf("expected 3 total tool calls, got %d", tl.Summary.TotalToolCalls)
	}
	// TotalActions = 3 tool calls + turn_start(1) + turn_end(1) = 5
	if tl.Summary.TotalActions != 5 {
		t.Errorf("expected 5 total actions, got %d", tl.Summary.TotalActions)
	}
	// All 3 tools succeeded
	if tl.Summary.ToolSuccesses != 3 {
		t.Errorf("expected 3 tool successes, got %d", tl.Summary.ToolSuccesses)
	}
	if tl.Summary.ToolFailures != 0 {
		t.Errorf("expected 0 tool failures, got %d", tl.Summary.ToolFailures)
	}

	// Check file_read event classification
	ev := tl.Events[1]
	if ev.Type != "file_read" {
		t.Errorf("expected file_read, got %s", ev.Type)
	}
	if ev.Tool != "view" {
		t.Errorf("expected tool 'view', got %s", ev.Tool)
	}
	if ev.Path != "/workspace/main.py" {
		t.Errorf("expected path /workspace/main.py, got %s", ev.Path)
	}

	// Check file_write event
	ev = tl.Events[3]
	if ev.Type != "file_write" {
		t.Errorf("expected file_write, got %s", ev.Type)
	}

	// Check bash event
	ev = tl.Events[5]
	if ev.Type != "bash" {
		t.Errorf("expected bash, got %s", ev.Type)
	}
	if ev.Input == "" {
		t.Error("expected non-empty input for bash command")
	}

	// Check complete event has duration and success
	ev = tl.Events[6]
	if ev.DurationMs != 300 {
		t.Errorf("expected 300ms duration, got %f", ev.DurationMs)
	}
	if ev.Success == nil || !*ev.Success {
		t.Error("expected success=true on bash complete")
	}

	// Total duration should sum "complete" event durations
	expectedDuration := 50.0 + 120.0 + 300.0
	if tl.Summary.TotalDurationMs != expectedDuration {
		t.Errorf("expected total duration %f, got %f", expectedDuration, tl.Summary.TotalDurationMs)
	}
}

func TestBuildActionTimeline_MCPAndErrors(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "external_tool.requested", ToolName: "azure-mcp-list", MCPServerName: "azure-sdk"},
		{Type: "external_tool.completed", ToolName: "azure-mcp-list", MCPServerName: "azure-sdk", Duration: 200, ToolSuccess: boolPtr(true)},
		{Type: "tool.execution_start", ToolName: "edit", Error: "permission denied"},
		{Type: "assistant.turn_end", TurnNumber: 1},
	}

	tl := BuildActionTimeline(records)

	if tl.Summary.MCPCalls != 1 {
		t.Errorf("expected 1 MCP call, got %d", tl.Summary.MCPCalls)
	}
	if tl.Summary.Errors != 1 {
		t.Errorf("expected 1 error, got %d", tl.Summary.Errors)
	}
	// TotalToolCalls: 1 MCP + 1 tool_call (edit) = 2
	if tl.Summary.TotalToolCalls != 2 {
		t.Errorf("expected 2 total tool calls, got %d", tl.Summary.TotalToolCalls)
	}
	// TotalActions: 2 tool calls + turn_start(1) + turn_end(1) = 4
	if tl.Summary.TotalActions != 4 {
		t.Errorf("expected 4 total actions, got %d", tl.Summary.TotalActions)
	}

	// MCP event should have server name
	ev := tl.Events[1]
	if ev.MCPServer != "azure-sdk" {
		t.Errorf("expected MCP server 'azure-sdk', got %s", ev.MCPServer)
	}
	if ev.Type != "mcp_call" {
		t.Errorf("expected mcp_call, got %s", ev.Type)
	}
}

func TestBuildActionTimeline_MultipleTurns(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view"},
		{Type: "assistant.turn_end", TurnNumber: 1},
		{Type: "assistant.turn_start", TurnNumber: 2},
		{Type: "tool.execution_start", ToolName: "create"},
		{Type: "tool.execution_complete", ToolName: "create"},
		{Type: "assistant.turn_end", TurnNumber: 2},
		{Type: "assistant.turn_start", TurnNumber: 3},
		{Type: "assistant.message", Content: "Done!"},
		{Type: "assistant.turn_end", TurnNumber: 3},
	}

	tl := BuildActionTimeline(records)

	if tl.Summary.TotalTurns != 3 {
		t.Errorf("expected 3 turns, got %d", tl.Summary.TotalTurns)
	}
	// TotalToolCalls: view(1) + create(1) = 2
	if tl.Summary.TotalToolCalls != 2 {
		t.Errorf("expected 2 total tool calls, got %d", tl.Summary.TotalToolCalls)
	}
	// TotalActions: 2 tool calls + 3 turn_start + 3 turn_end + 1 message = 9
	if tl.Summary.TotalActions != 9 {
		t.Errorf("expected 9 total actions, got %d", tl.Summary.TotalActions)
	}
	// Events in turn 2 should have TurnNumber 2
	ev := tl.Events[5] // tool.execution_start in turn 2
	if ev.TurnNumber != 2 {
		t.Errorf("expected turn number 2, got %d", ev.TurnNumber)
	}
}

func TestBuildActionTimeline_ToolBreakdown(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view"},
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view"},
		{Type: "tool.execution_start", ToolName: "create"},
		{Type: "tool.execution_complete", ToolName: "create"},
		{Type: "tool.execution_start", ToolName: "bash"},
		{Type: "tool.execution_complete", ToolName: "bash"},
	}

	tl := BuildActionTimeline(records)

	if tl.Summary.ToolBreakdown["view"] != 2 {
		t.Errorf("expected 2 view calls, got %d", tl.Summary.ToolBreakdown["view"])
	}
	if tl.Summary.ToolBreakdown["create"] != 1 {
		t.Errorf("expected 1 create call, got %d", tl.Summary.ToolBreakdown["create"])
	}
	if tl.Summary.ToolBreakdown["bash"] != 1 {
		t.Errorf("expected 1 bash call, got %d", tl.Summary.ToolBreakdown["bash"])
	}
}

func TestBuildActionTimeline_Sequence(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start"},
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view"},
	}

	tl := BuildActionTimeline(records)
	for i, ev := range tl.Events {
		if ev.Sequence != i {
			t.Errorf("event %d: expected sequence %d, got %d", i, i, ev.Sequence)
		}
	}
}

func TestBuildActionTimeline_TruncatesLargeInput(t *testing.T) {
	longArgs := make([]byte, 1000)
	for i := range longArgs {
		longArgs[i] = 'x'
	}
	records := []report.SessionEventRecord{
		{Type: "tool.execution_start", ToolName: "bash", ToolArgs: string(longArgs)},
	}

	tl := BuildActionTimeline(records)
	ev := tl.Events[0]
	if len(ev.Input) > maxActionFieldLen+3 { // +3 for "…" (3 bytes in UTF-8)
		t.Errorf("expected truncated input, got length %d", len(ev.Input))
	}
}

func TestClassifyEventType(t *testing.T) {
	tests := []struct {
		evType    string
		toolName  string
		mcpServer string
		wantType  string
		wantAct   string
	}{
		{"tool.execution_start", "view", "", "file_read", "start"},
		{"tool.execution_start", "read_file", "", "file_read", "start"},
		{"tool.execution_start", "create", "", "file_write", "start"},
		{"tool.execution_start", "edit", "", "file_write", "start"},
		{"tool.execution_start", "bash", "", "bash", "start"},
		{"tool.execution_start", "grep", "", "tool_call", "start"},
		{"tool.execution_start", "azure-documentation", "azure", "mcp_call", "start"},
		{"tool.execution_complete", "view", "", "file_read", "complete"},
		{"tool.execution_complete", "bash", "", "bash", "complete"},
		{"tool.execution_complete", "azure-documentation", "azure", "mcp_call", "complete"},
		{"assistant.turn_start", "", "", "turn_start", ""},
		{"assistant.turn_end", "", "", "turn_end", ""},
		{"assistant.reasoning", "", "", "reasoning", ""},
		{"assistant.message", "", "", "message", ""},
		{"external_tool.requested", "mcp-tool", "", "mcp_call", "start"},
		{"external_tool.completed", "mcp-tool", "", "mcp_call", "complete"},
		{"command.execute", "", "", "bash", "start"},
		{"command.completed", "", "", "bash", "complete"},
		{"session.error", "", "", "error", ""},
		{"session.truncation", "", "", "truncation", ""},
		{"session.workspace_file_changed", "", "", "file_change", ""},
		{"skill.invoked", "", "", "skill", ""},
		{"abort", "", "", "abort", ""},
		{"unknown.type", "", "", "other", ""},
	}

	for _, tc := range tests {
		gotType, gotAct := classifyEventType(tc.evType, tc.toolName, tc.mcpServer)
		if gotType != tc.wantType || gotAct != tc.wantAct {
			t.Errorf("classifyEventType(%q, %q, %q) = (%q, %q), want (%q, %q)",
				tc.evType, tc.toolName, tc.mcpServer, gotType, gotAct, tc.wantType, tc.wantAct)
		}
	}
}

func TestBuildActionTimeline_MCPToolExecutionEvent(t *testing.T) {
	records := []report.SessionEventRecord{{
		Type:          "tool.execution_start",
		ToolName:      "azure-documentation",
		MCPServerName: "azure",
	}}

	tl := BuildActionTimeline(records)
	if tl.Summary.MCPCalls != 1 {
		t.Fatalf("MCPCalls = %d, want 1", tl.Summary.MCPCalls)
	}
	if tl.Summary.ToolCalls != 0 {
		t.Fatalf("ToolCalls = %d, want 0", tl.Summary.ToolCalls)
	}
	if got := tl.Events[0].MCPServer; got != "azure" {
		t.Errorf("MCPServer = %q, want azure", got)
	}
}

func TestActionTimeline_ToGraderActionLog(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "tool.execution_start", ToolName: "view", FilePath: "/workspace/main.py"},
		{Type: "tool.execution_complete", ToolName: "view"},
		{Type: "tool.execution_start", ToolName: "create", FilePath: "/workspace/out.py"},
		{Type: "tool.execution_complete", ToolName: "create"},
		{Type: "tool.execution_start", ToolName: "bash"},
		{Type: "tool.execution_complete", ToolName: "bash"},
		{Type: "assistant.message", Content: "Done"},
		{Type: "assistant.turn_end", TurnNumber: 1},
	}

	tl := BuildActionTimeline(records)
	log := tl.ToGraderActionLog()

	// All event types are forwarded; only "complete" half of tool start/complete pairs is dropped.
	// Expected 6 events: turn_start, file_read view, file_write create, bash, message, turn_end.
	if len(log) != 6 {
		t.Fatalf("expected 6 grader events, got %d", len(log))
	}

	// Spot-check that tool events are still present and typed correctly.
	var sawView, sawCreate, sawBash, sawMessage bool
	for _, e := range log {
		switch e.Type {
		case "file_read":
			if e.Tool == "view" {
				sawView = true
			}
		case "file_write":
			if e.Tool == "create" {
				sawCreate = true
			}
		case "bash":
			sawBash = true
		case "message":
			sawMessage = true
			if e.Text != "Done" {
				t.Errorf("expected message Text 'Done', got %q", e.Text)
			}
		}
	}
	if !sawView || !sawCreate || !sawBash || !sawMessage {
		t.Errorf("missing expected events: view=%v create=%v bash=%v message=%v",
			sawView, sawCreate, sawBash, sawMessage)
	}
}

func TestActionTimeline_ToGraderActionLog_Nil(t *testing.T) {
	var tl *ActionTimeline
	if log := tl.ToGraderActionLog(); log != nil {
		t.Errorf("expected nil for nil timeline, got %v", log)
	}
}

// TestActionTimeline_ToGraderActionLog_SkillEvents verifies that tool_call
// events with Tool="skill" are filtered out, while skill.invoked events with
// individual skill names are preserved.
func TestActionTimeline_ToGraderActionLog_SkillEvents(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "tool.execution_start", ToolName: "skill"}, // Should be filtered
		{Type: "skill.invoked", SkillName: "markdown-headings"},
		{Type: "tool.execution_complete", ToolName: "skill"},
		{Type: "tool.execution_start", ToolName: "skill"}, // Should be filtered
		{Type: "skill.invoked", SkillName: "markdown-lists"},
		{Type: "tool.execution_complete", ToolName: "skill"},
		{Type: "assistant.turn_end", TurnNumber: 1},
	}

	tl := BuildActionTimeline(records)
	log := tl.ToGraderActionLog()

	// Should have: turn_start, 2 skill events (not tool_call), turn_end = 4 events
	if len(log) != 4 {
		t.Fatalf("expected 4 events, got %d", len(log))
	}

	// Verify the skill events have individual skill names, not "skill"
	var sawHeadings, sawLists bool
	for _, e := range log {
		if e.Type == "skill" {
			if e.Tool == "markdown-headings" {
				sawHeadings = true
			}
			if e.Tool == "markdown-lists" {
				sawLists = true
			}
			if e.Tool == "skill" {
				t.Errorf("found tool_call event with Tool='skill'; should be filtered")
			}
		}
	}
	if !sawHeadings || !sawLists {
		t.Errorf("missing skill events: headings=%v lists=%v", sawHeadings, sawLists)
	}
}

func TestActionTimeline_ToReport(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view", Duration: 50, ToolSuccess: boolPtr(true)},
	}

	tl := BuildActionTimeline(records)
	rpt := tl.ToReport()

	if rpt == nil {
		t.Fatal("ToReport returned nil")
	}
	if len(rpt.Events) != 3 {
		t.Errorf("expected 3 report events, got %d", len(rpt.Events))
	}
	if rpt.Summary.TotalEvents != 3 {
		t.Errorf("expected 3 total events in summary, got %d", rpt.Summary.TotalEvents)
	}
	if rpt.Summary.FileReads != 1 {
		t.Errorf("expected 1 file read in summary, got %d", rpt.Summary.FileReads)
	}
	// Verify TotalActions and TotalToolCalls are mapped
	// 1 file_read tool + turn_start(1) = TotalActions 2, TotalToolCalls 1
	if rpt.Summary.TotalToolCalls != 1 {
		t.Errorf("expected 1 total tool call in report, got %d", rpt.Summary.TotalToolCalls)
	}
	if rpt.Summary.TotalActions != 2 {
		t.Errorf("expected 2 total actions in report, got %d", rpt.Summary.TotalActions)
	}
	if rpt.Summary.ToolSuccesses != 1 {
		t.Errorf("expected 1 tool success in report, got %d", rpt.Summary.ToolSuccesses)
	}
	if rpt.Summary.ToolFailures != 0 {
		t.Errorf("expected 0 tool failures in report, got %d", rpt.Summary.ToolFailures)
	}
	// Verify field mapping
	if rpt.Events[2].DurationMs != 50 {
		t.Errorf("expected duration 50, got %f", rpt.Events[2].DurationMs)
	}
	if rpt.Events[2].Success == nil || !*rpt.Events[2].Success {
		t.Error("expected success=true in report event")
	}
}

func TestBuildActionTimeline_ToolSuccessFailure(t *testing.T) {
	records := []report.SessionEventRecord{
		{Type: "assistant.turn_start", TurnNumber: 1},
		{Type: "assistant.reasoning", Content: "Let me think..."},
		{Type: "tool.execution_start", ToolName: "view"},
		{Type: "tool.execution_complete", ToolName: "view", Duration: 50, ToolSuccess: boolPtr(true)},
		{Type: "tool.execution_start", ToolName: "edit", Error: "permission denied"},
		{Type: "tool.execution_complete", ToolName: "edit", Duration: 10, ToolSuccess: boolPtr(false)},
		{Type: "assistant.message", Content: "I hit an error"},
		{Type: "assistant.turn_end", TurnNumber: 1},
	}

	tl := BuildActionTimeline(records)

	if tl.Summary.ToolSuccesses != 1 {
		t.Errorf("expected 1 tool success, got %d", tl.Summary.ToolSuccesses)
	}
	if tl.Summary.ToolFailures != 1 {
		t.Errorf("expected 1 tool failure, got %d", tl.Summary.ToolFailures)
	}
	// TotalToolCalls: file_read(1) + tool_call/edit(1) = 2
	if tl.Summary.TotalToolCalls != 2 {
		t.Errorf("expected 2 total tool calls, got %d", tl.Summary.TotalToolCalls)
	}
	// TotalActions: 2 tool calls + turn_start(1) + turn_end(1) + reasoning(1) + message(1) = 6
	if tl.Summary.TotalActions != 6 {
		t.Errorf("expected 6 total actions, got %d", tl.Summary.TotalActions)
	}
}

func TestActionTimeline_ToReport_Nil(t *testing.T) {
	var tl *ActionTimeline
	if rpt := tl.ToReport(); rpt != nil {
		t.Errorf("expected nil for nil timeline, got %v", rpt)
	}
}

func TestTruncateField(t *testing.T) {
	short := "hello"
	if got := truncateField(short, 10); got != short {
		t.Errorf("expected %q, got %q", short, got)
	}

	long := "abcdefghij"
	got := truncateField(long, 5)
	if got != "abcde…" {
		t.Errorf("expected %q, got %q", "abcde…", got)
	}
}
