package review

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	copilot "github.com/github/copilot-sdk/go"
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

	// Count actions and enforce limit
	switch event.Type {
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
	switch event.Type {
	case copilot.SessionEventTypeAssistantTurnStart:
		slog.Debug("Review turn started", "model", c.model)
	case copilot.SessionEventTypeAssistantTurnEnd:
		slog.Debug("Review turn ended", "model", c.model)
	case copilot.SessionEventTypeAssistantMessage:
		if event.Data.Content != nil {
			slog.Debug("Review assistant message", "model", c.model,
				"content_len", len(*event.Data.Content))
		}
	case copilot.SessionEventTypeToolExecutionStart:
		toolName := ""
		if event.Data.ToolName != nil {
			toolName = *event.Data.ToolName
		}
		slog.Debug("Review tool start", "model", c.model, "tool", toolName)
	case copilot.SessionEventTypeToolExecutionComplete:
		toolName := ""
		if event.Data.ToolName != nil {
			toolName = *event.Data.ToolName
		}
		slog.Debug("Review tool complete", "model", c.model, "tool", toolName)
	case copilot.SessionEventTypeAssistantUsage:
		slog.Debug("Review token usage", "model", c.model)
	}

	if event.Type == copilot.SessionEventTypeAssistantMessage && event.Data.Content != nil {
		c.assistantContent += *event.Data.Content
	}

	// Capture all events for the report timeline
	evt := ReviewEvent{Type: string(event.Type)}
	if event.Data.ToolName != nil {
		evt.ToolName = *event.Data.ToolName
	}
	if event.Data.Content != nil {
		evt.Content = *event.Data.Content
	}
	if event.Data.Arguments != nil {
		if argsBytes, err := json.Marshal(event.Data.Arguments); err == nil {
			evt.ToolArgs = string(argsBytes)
		}
	}
	if event.Data.Result != nil {
		if event.Data.Result.Content != nil {
			evt.Result = *event.Data.Result.Content
		}
	}
	if event.Data.Error != nil {
		if event.Data.Error.ErrorClass != nil {
			evt.Error = event.Data.Error.ErrorClass.Message
		} else if event.Data.Error.String != nil {
			evt.Error = *event.Data.Error.String
		}
	}
	if event.Data.Duration != nil {
		evt.Duration = *event.Data.Duration
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

// buildConsolidationPrompt constructs the prompt sent to the consolidator model.
func buildConsolidationPrompt(originalPrompt string, panel []ReviewResult) string {
	var b strings.Builder
	b.WriteString("You are a senior review consolidator. Multiple independent reviewers have scored the same agent output.\n")
	b.WriteString("Synthesize their feedback into a single consensus review.\n\n")

	b.WriteString("## Original Prompt\n\n")
	b.WriteString(originalPrompt)
	b.WriteString("\n\n")

	b.WriteString("## Individual Reviews\n\n")
	for i, r := range panel {
		reviewJSON, _ := json.MarshalIndent(r, "", "  ")
		fmt.Fprintf(&b, "### Reviewer %d (%s)\n```json\n%s\n```\n\n", i+1, r.Model, string(reviewJSON))
	}

	b.WriteString("## Instructions\n\n")
	b.WriteString("Produce a consensus review using the criteria-based pass/fail system. ")
	b.WriteString("For each criterion, it PASSES if the majority of reviewers marked it as passed. ")
	b.WriteString("Use the union of all criteria across reviewers. ")
	b.WriteString("Combine the best issues and strengths from all reviewers. ")
	b.WriteString("Write a summary that captures the consensus view.\n\n")
	b.WriteString("Respond with ONLY a JSON object in the same format as the individual reviews.\n")

	return b.String()
}
