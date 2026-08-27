package eval

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// maxActionFieldLen is the maximum length for truncated input/output fields.
const maxActionFieldLen = 512

// ActionEvent represents a single captured agent action with full detail.
type ActionEvent struct {
	Sequence   int     `json:"sequence"`              // ordinal position in the timeline
	Type       string  `json:"type"`                  // classified action type
	Tool       string  `json:"tool,omitempty"`        // tool name (for tool_call, file_read, file_write, bash)
	Action     string  `json:"action,omitempty"`      // sub-action (e.g., "start", "complete")
	Path       string  `json:"path,omitempty"`        // file path when applicable
	Input      string  `json:"input,omitempty"`       // truncated tool arguments
	Output     string  `json:"output,omitempty"`      // truncated tool result
	Error      string  `json:"error,omitempty"`       // error message if failed
	Success    *bool   `json:"success,omitempty"`     // tool execution success
	DurationMs float64 `json:"duration_ms,omitempty"` // duration in milliseconds
	TurnNumber int     `json:"turn_number,omitempty"` // assistant turn number
	MCPServer  string  `json:"mcp_server,omitempty"`  // MCP server name
}

// ActionSummary holds aggregate statistics about the action timeline.
type ActionSummary struct {
	TotalEvents     int            `json:"total_events"`
	TotalTurns      int            `json:"total_turns"`
	TotalActions    int            `json:"total_actions"`
	TotalToolCalls  int            `json:"total_tool_calls"`
	ToolCalls       int            `json:"tool_calls"`
	FileReads       int            `json:"file_reads"`
	FileWrites      int            `json:"file_writes"`
	BashCommands    int            `json:"bash_commands"`
	MCPCalls        int            `json:"mcp_calls"`
	Errors          int            `json:"errors"`
	ToolBreakdown   map[string]int `json:"tool_breakdown"`
	TotalDurationMs float64        `json:"total_duration_ms,omitempty"`
	ToolSuccesses   int            `json:"tool_successes"`
	ToolFailures    int            `json:"tool_failures"`
}

// ActionTimeline holds an ordered sequence of ActionEvents with summary stats.
type ActionTimeline struct {
	Events  []ActionEvent `json:"events"`
	Summary ActionSummary `json:"summary"`
}

// NewActionTimeline creates an empty ActionTimeline ready for use.
func NewActionTimeline() *ActionTimeline {
	return &ActionTimeline{
		Events: []ActionEvent{},
		Summary: ActionSummary{
			ToolBreakdown: make(map[string]int),
		},
	}
}

// classifyEventType maps a SessionEventRecord type string to a higher-level
// action type for the timeline.
func classifyEventType(evType, toolName, mcpServer string) (actionType, action string) {
	switch evType {
	case "tool.execution_start":
		action = "start"
		switch {
		case mcpServer != "":
			actionType = "mcp_call"
		case isFileReadTool(toolName):
			actionType = "file_read"
		case isFileWriteTool(toolName):
			actionType = "file_write"
		case isBashTool(toolName):
			actionType = "bash"
		default:
			actionType = "tool_call"
		}
	case "tool.execution_complete":
		action = "complete"
		switch {
		case mcpServer != "":
			actionType = "mcp_call"
		case isFileReadTool(toolName):
			actionType = "file_read"
		case isFileWriteTool(toolName):
			actionType = "file_write"
		case isBashTool(toolName):
			actionType = "bash"
		default:
			actionType = "tool_call"
		}
	case "assistant.turn_start":
		actionType = "turn_start"
	case "assistant.turn_end":
		actionType = "turn_end"
	case "assistant.reasoning":
		actionType = "reasoning"
	case "assistant.message":
		actionType = "message"
	case "assistant.intent":
		actionType = "intent"
	case "session.workspace_file_changed":
		actionType = "file_change"
	case "command.execute":
		actionType = "bash"
		action = "start"
	case "command.completed":
		actionType = "bash"
		action = "complete"
	case "external_tool.requested":
		actionType = "mcp_call"
		action = "start"
	case "external_tool.completed":
		actionType = "mcp_call"
		action = "complete"
	case "skill.invoked":
		actionType = "skill"
	case "session.error":
		actionType = "error"
	case "session.warning":
		actionType = "warning"
	case "session.truncation":
		actionType = "truncation"
	case "session.compaction_start":
		actionType = "compaction"
		action = "start"
	case "session.compaction_complete":
		actionType = "compaction"
		action = "complete"
	case "abort":
		actionType = "abort"
	default:
		actionType = "other"
	}
	return actionType, action
}

// isFileReadTool returns true for tools that read files.
func isFileReadTool(name string) bool {
	switch name {
	case "view", "read_file", "get_file_contents", "read":
		return true
	}
	return false
}

// isBashTool returns true for shell/command tools.
func isBashTool(name string) bool {
	switch name {
	case "bash", "shell", "run_command", "execute_command":
		return true
	}
	return false
}

// BuildActionTimeline converts SessionEventRecords into a structured ActionTimeline.
func BuildActionTimeline(records []report.SessionEventRecord) *ActionTimeline {
	tl := NewActionTimeline()
	if len(records) == 0 {
		return tl
	}

	turnTracker := 0
	for i, rec := range records {
		actionType, action := classifyEventType(rec.Type, rec.ToolName, rec.MCPServerName)

		ev := ActionEvent{
			Sequence:   i,
			Type:       actionType,
			Action:     action,
			TurnNumber: rec.TurnNumber,
		}

		// Track turn number from turn_start events
		if rec.TurnNumber > 0 {
			turnTracker = rec.TurnNumber
		}
		if ev.TurnNumber == 0 && turnTracker > 0 {
			ev.TurnNumber = turnTracker
		}

		// Tool name
		// Tool name (for skill events, fall back to SkillName so text matchers
		// and tool-name lookups can find the skill).
		if rec.ToolName != "" {
			ev.Tool = rec.ToolName
		} else if actionType == "skill" && rec.SkillName != "" {
			ev.Tool = rec.SkillName
		}

		// File path
		if rec.FilePath != "" {
			ev.Path = rec.FilePath
		}

		// Input (tool args)
		if rec.ToolArgs != "" {
			ev.Input = truncateField(rec.ToolArgs, maxActionFieldLen)
		} else if rec.CommandText != "" {
			ev.Input = truncateField(rec.CommandText, maxActionFieldLen)
		}

		// Output (tool result or content)
		if rec.ToolResult != "" {
			ev.Output = truncateField(rec.ToolResult, maxActionFieldLen)
		} else if actionType == "intent" && rec.Intent != "" {
			ev.Output = truncateField(rec.Intent, maxActionFieldLen)
		} else if actionType == "warning" && rec.WarningText != "" {
			ev.Output = truncateField(rec.WarningText, maxActionFieldLen)
		} else if rec.Content != "" {
			// Includes reasoning/message text — needed by activity grader text matchers.
			ev.Output = truncateField(rec.Content, maxActionFieldLen)
		}

		// Error
		if rec.Error != "" {
			ev.Error = rec.Error
		}

		// Success
		if rec.ToolSuccess != nil {
			s := *rec.ToolSuccess
			ev.Success = &s
		}

		// Duration
		if rec.Duration > 0 {
			ev.DurationMs = rec.Duration
		}

		// MCP server
		if rec.MCPServerName != "" {
			ev.MCPServer = rec.MCPServerName
		}

		tl.Events = append(tl.Events, ev)
	}

	// Compute summary
	tl.Summary = computeSummary(tl.Events)
	return tl
}

// computeSummary aggregates statistics from action events.
func computeSummary(events []ActionEvent) ActionSummary {
	s := ActionSummary{
		TotalEvents:   len(events),
		ToolBreakdown: make(map[string]int),
	}

	turnsSeen := make(map[int]bool)
	actionEvents := 0 // non-tool actions: reasoning, message, intent, turn start/end
	for _, ev := range events {
		if ev.TurnNumber > 0 {
			turnsSeen[ev.TurnNumber] = true
		}

		// Count non-tool action events
		switch ev.Type {
		case "reasoning", "message", "intent", "turn_start", "turn_end":
			actionEvents++
		}

		switch ev.Type {
		case "tool_call":
			if ev.Action == "start" {
				s.ToolCalls++
				if ev.Tool != "" {
					s.ToolBreakdown[ev.Tool]++
				}
			}
		case "file_read":
			if ev.Action == "start" {
				s.FileReads++
				if ev.Tool != "" {
					s.ToolBreakdown[ev.Tool]++
				}
			}
		case "file_write":
			if ev.Action == "start" {
				s.FileWrites++
				if ev.Tool != "" {
					s.ToolBreakdown[ev.Tool]++
				}
			}
		case "bash":
			if ev.Action == "start" {
				s.BashCommands++
				if ev.Tool != "" {
					s.ToolBreakdown[ev.Tool]++
				}
			}
		case "mcp_call":
			if ev.Action == "start" {
				s.MCPCalls++
				if ev.Tool != "" {
					s.ToolBreakdown[ev.Tool]++
				}
			}
		}

		if ev.Error != "" {
			s.Errors++
		}

		// Track tool success/failure from "complete" events
		if ev.Action == "complete" && ev.Success != nil {
			if *ev.Success {
				s.ToolSuccesses++
			} else {
				s.ToolFailures++
			}
		}

		// Sum durations from "complete" events only to avoid double-counting
		if ev.DurationMs > 0 && ev.Action == "complete" {
			s.TotalDurationMs += ev.DurationMs
		}
	}

	s.TotalTurns = len(turnsSeen)
	// TotalToolCalls = all tool-related "start" events
	s.TotalToolCalls = s.ToolCalls + s.FileReads + s.FileWrites + s.BashCommands + s.MCPCalls
	// TotalActions = tool calls + reasoning + messages + intents + turn boundaries
	s.TotalActions = s.TotalToolCalls + actionEvents
	return s
}

// truncateField truncates a string to maxLen, appending "…" if truncated.
func truncateField(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "…"
}

// ToGraderActionLog converts the timeline into []graders.ActionEvent for
// pluggable grader input. All event types are forwarded so the activity
// grader's contains_action / excludes_action checks can match on any action
// kind (tool calls, messages, reasoning, errors, etc.).
//
// "complete" events are skipped so each logical action appears once. Events
// without an explicit start/complete pair (reasoning, message, intent,
// turn_start, etc.) are forwarded as-is.
//
// tool_call events with Tool="skill" are skipped because they're redundant:
// the individual skill name appears in the subsequent skill.invoked event.
func (tl *ActionTimeline) ToGraderActionLog() []graders.ActionEvent {
	if tl == nil {
		return nil
	}
	var out []graders.ActionEvent
	for _, ev := range tl.Events {
		if ev.Action == "complete" {
			continue
		}
		// Skip tool_call events with Tool="skill" — the individual skill name
		// appears in the subsequent skill.invoked event (Type="skill").
		if ev.Type == "tool_call" && ev.Tool == "skill" {
			continue
		}
		out = append(out, graders.ActionEvent{
			Type:       ev.Type,
			Tool:       ev.Tool,
			Action:     ev.Type, // backwards-compat: legacy callers used Action for Type
			Path:       ev.Path,
			Text:       ev.Output,
			Error:      ev.Error,
			TurnNumber: ev.TurnNumber,
			MCPServer:  ev.MCPServer,
		})
	}
	return out
}

// ToReport converts the eval ActionTimeline into the report-serializable form.
func (tl *ActionTimeline) ToReport() *report.ActionTimelineReport {
	if tl == nil {
		return nil
	}
	r := &report.ActionTimelineReport{
		Events: make([]report.ActionEventReport, len(tl.Events)),
		Summary: report.ActionSummaryReport{
			TotalEvents:     tl.Summary.TotalEvents,
			TotalTurns:      tl.Summary.TotalTurns,
			TotalActions:    tl.Summary.TotalActions,
			TotalToolCalls:  tl.Summary.TotalToolCalls,
			ToolCalls:       tl.Summary.ToolCalls,
			FileReads:       tl.Summary.FileReads,
			FileWrites:      tl.Summary.FileWrites,
			BashCommands:    tl.Summary.BashCommands,
			MCPCalls:        tl.Summary.MCPCalls,
			Errors:          tl.Summary.Errors,
			ToolBreakdown:   tl.Summary.ToolBreakdown,
			TotalDurationMs: tl.Summary.TotalDurationMs,
			ToolSuccesses:   tl.Summary.ToolSuccesses,
			ToolFailures:    tl.Summary.ToolFailures,
		},
	}
	for i, ev := range tl.Events {
		r.Events[i] = report.ActionEventReport{
			Sequence:   ev.Sequence,
			Type:       ev.Type,
			Tool:       ev.Tool,
			Action:     ev.Action,
			Path:       ev.Path,
			Input:      ev.Input,
			Output:     ev.Output,
			Error:      ev.Error,
			Success:    ev.Success,
			DurationMs: ev.DurationMs,
			TurnNumber: ev.TurnNumber,
			MCPServer:  ev.MCPServer,
		}
	}
	return r
}
