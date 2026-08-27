package review

import (
	"encoding/json"
	"log/slog"
	"sync"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/copilotevent"
)

// eventCollector accumulates assistant messages and review events from
// a Copilot session. It is safe for concurrent use.
type eventCollector struct {
	mu               sync.Mutex
	assistantContent string
	events           []ReviewEvent
	actionCount      int
	actionLimitHit   bool
	maxActions       int
	model            string
	cancel           func()
	toolTracker      copilotevent.ToolTracker
}

func newEventCollector(model string, maxActions int, cancel func()) *eventCollector {
	return &eventCollector{
		model:      model,
		maxActions: maxActions,
		cancel:     cancel,
	}
}

// handleEvent processes a single Copilot session event.
func (c *eventCollector) handleEvent(event copilot.SessionEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	details := copilotevent.Extract(event)
	c.toolTracker.Enrich(event, &details)

	// Count actions and enforce limit
	switch event.Type() {
	case copilot.SessionEventTypeAssistantReasoning,
		copilot.SessionEventTypeAssistantMessage,
		copilot.SessionEventTypeToolExecutionStart:
		c.actionCount++
		if c.maxActions > 0 && c.actionCount > c.maxActions && !c.actionLimitHit {
			c.actionLimitHit = true
			slog.Warn("Review action limit reached, cancelling session",
				"model", c.model, "actions", c.actionCount, "max_session_actions", c.maxActions)
			c.cancel()
		}
	}

	// Log review events at debug level for visibility during runs.
	switch event.Type() {
	case copilot.SessionEventTypeAssistantTurnStart:
		slog.Debug("Review turn started", "model", c.model)
	case copilot.SessionEventTypeAssistantTurnEnd:
		slog.Debug("Review turn ended", "model", c.model)
	case copilot.SessionEventTypeAssistantMessage:
		if details.Content != nil {
			slog.Debug("Review assistant message", "model", c.model,
				"content_len", len(*details.Content))
		}
	case copilot.SessionEventTypeToolExecutionStart:
		toolName := ""
		if details.ToolName != nil {
			toolName = *details.ToolName
		}
		slog.Debug("Review tool start", "model", c.model, "tool", toolName)
	case copilot.SessionEventTypeToolExecutionComplete:
		toolName := ""
		if details.ToolName != nil {
			toolName = *details.ToolName
		}
		slog.Debug("Review tool complete", "model", c.model, "tool", toolName)
	case copilot.SessionEventTypeAssistantUsage:
		slog.Debug("Review token usage", "model", c.model)
	}

	if event.Type() == copilot.SessionEventTypeAssistantMessage && details.Content != nil {
		c.assistantContent += *details.Content
	}

	// Capture all events for the report timeline
	evt := ReviewEvent{Type: string(event.Type())}
	if details.ToolName != nil {
		evt.ToolName = *details.ToolName
	}
	if details.Content != nil {
		evt.Content = *details.Content
	}
	if details.Arguments != nil {
		if argsBytes, err := json.Marshal(details.Arguments); err == nil {
			evt.ToolArgs = string(argsBytes)
		}
	}
	if details.Result != nil {
		if details.Result.Content != nil {
			evt.Result = *details.Result.Content
		}
	}
	if details.Error != nil {
		if details.Error.ErrorClass != nil {
			evt.Error = details.Error.ErrorClass.Message
		} else if details.Error.String != nil {
			evt.Error = *details.Error.String
		}
	}
	if details.Duration != nil {
		evt.Duration = *details.Duration
	}
	c.events = append(c.events, evt)
}

// response returns the accumulated assistant content and event timeline.
func (c *eventCollector) response() (string, []ReviewEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	events := make([]ReviewEvent, len(c.events))
	copy(events, c.events)
	return c.assistantContent, events
}
