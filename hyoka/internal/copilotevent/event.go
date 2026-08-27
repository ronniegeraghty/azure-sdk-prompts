// Package copilotevent extracts common fields from Copilot SDK session events.
package copilotevent

import (
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

// Details contains the fields Hyoka records from the SDK's typed event variants.
type Details struct {
	Content       *string
	ToolName      *string
	Arguments     any
	Result        *Result
	Error         *Error
	Success       *bool
	Duration      *float64
	MCPServerName *string
	MCPToolName   *string
	Path          *string
	Operation     *string
	Intent        *string
	InputTokens   *int64
	OutputTokens  *int64
	Command       *string
	SkillName     *string
	ToolCallID    *string
	Reason        *string
	ErrorReason   *string
	Message       *string
	Skills        []Named
	Servers       []Named
}

// Result contains the textual parts of a tool execution result.
type Result struct {
	Content         *string
	DetailedContent *string
}

// Error contains the error representations recorded by older SDK releases.
type Error struct {
	ErrorClass *ErrorClass
	String     *string
}

// ErrorClass contains an error message.
type ErrorClass struct {
	Message string
}

// Named represents a named loaded SDK resource.
type Named struct {
	Name string
}

// ToolTracker restores fields that SDK v1 only exposes across a matching pair
// of tool execution events. Its zero value is ready to use.
type ToolTracker struct {
	names     map[string]string
	startedAt map[string]time.Time
}

// Enrich restores the tool name and duration on completion events.
func (t *ToolTracker) Enrich(event copilot.SessionEvent, details *Details) {
	if details.ToolCallID == nil {
		return
	}

	toolCallID := *details.ToolCallID
	switch event.Type() {
	case copilot.SessionEventTypeToolExecutionStart:
		if details.ToolName != nil {
			if t.names == nil {
				t.names = make(map[string]string)
			}
			t.names[toolCallID] = *details.ToolName
		}
		if !event.Timestamp.IsZero() {
			if t.startedAt == nil {
				t.startedAt = make(map[string]time.Time)
			}
			t.startedAt[toolCallID] = event.Timestamp
		}
	case copilot.SessionEventTypeToolExecutionComplete:
		if toolName, ok := t.names[toolCallID]; ok {
			details.ToolName = &toolName
			delete(t.names, toolCallID)
		}
		if startedAt, ok := t.startedAt[toolCallID]; ok {
			if duration := event.Timestamp.Sub(startedAt); duration >= 0 && !event.Timestamp.IsZero() {
				durationMS := float64(duration) / float64(time.Millisecond)
				details.Duration = &durationMS
			}
			delete(t.startedAt, toolCallID)
		}
	}
}

// Extract returns the common fields from a typed SDK event payload.
func Extract(event copilot.SessionEvent) Details {
	var details Details

	switch data := event.Data.(type) {
	case *copilot.AssistantIntentData:
		details.Intent = &data.Intent
	case *copilot.AssistantMessageData:
		details.Content = &data.Content
	case *copilot.AssistantReasoningData:
		details.Content = &data.Content
	case *copilot.AssistantUsageData:
		details.InputTokens = data.InputTokens
		details.OutputTokens = data.OutputTokens
	case *copilot.CommandExecuteData:
		details.Command = &data.Command
	case *copilot.ExternalToolRequestedData:
		details.Arguments = data.Arguments
		details.ToolCallID = &data.ToolCallID
		details.ToolName = &data.ToolName
	case *copilot.SessionErrorData:
		details.Content = &data.Message
		details.Error = &Error{ErrorClass: &ErrorClass{Message: data.Message}}
	case *copilot.SessionMCPServersLoadedData:
		details.Servers = make([]Named, 0, len(data.Servers))
		for _, server := range data.Servers {
			details.Servers = append(details.Servers, Named{Name: server.Name})
		}
	case *copilot.SessionSkillsLoadedData:
		details.Skills = make([]Named, 0, len(data.Skills))
		for _, skill := range data.Skills {
			details.Skills = append(details.Skills, Named{Name: skill.Name})
		}
	case *copilot.SessionWarningData:
		details.Message = &data.Message
	case *copilot.SessionWorkspaceFileChangedData:
		operation := string(data.Operation)
		details.Operation = &operation
		details.Path = &data.Path
	case *copilot.SkillInvokedData:
		details.SkillName = &data.Name
	case *copilot.SubagentCompletedData:
		details.ToolCallID = &data.ToolCallID
	case *copilot.SubagentFailedData:
		details.Error = &Error{ErrorClass: &ErrorClass{Message: data.Error}}
		details.ToolCallID = &data.ToolCallID
	case *copilot.ToolExecutionCompleteData:
		details.Success = &data.Success
		details.ToolCallID = &data.ToolCallID
		if data.Error != nil {
			details.Error = &Error{ErrorClass: &ErrorClass{Message: data.Error.Message}}
		}
		if data.Result != nil {
			details.Result = &Result{
				Content:         &data.Result.Content,
				DetailedContent: data.Result.DetailedContent,
			}
		}
	case *copilot.ToolExecutionStartData:
		details.Arguments = data.Arguments
		details.MCPServerName = data.MCPServerName
		details.MCPToolName = data.MCPToolName
		details.ToolCallID = &data.ToolCallID
		details.ToolName = &data.ToolName
	}

	return details
}
