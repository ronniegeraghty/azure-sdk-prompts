package eval

import "time"

// ActionEvent captures a single agent action during an evaluation session.
// This is the rich timeline event used for report rendering, distinct from
// the lightweight graders.ActionEvent used for rubric evaluation.
type ActionEvent struct {
	Timestamp  time.Time     `json:"timestamp"`
	Type       string        `json:"type"`                  // "tool_call", "file_read", "file_write", "bash", "response"
	Tool       string        `json:"tool"`                  // Tool name (e.g., "editFile", "bash")
	Input      string        `json:"input"`                 // Tool input (truncated if large)
	Output     string        `json:"output"`                // Tool output (truncated if large)
	Duration   time.Duration `json:"duration"`              // Wall-clock duration of this action
	TurnNumber int           `json:"turn_number"`           // Which conversation turn this belongs to
	Success    bool          `json:"success"`               // Whether the action completed successfully
	Error      string        `json:"error,omitempty"`       // Error message if action failed
}

// ActionTimeline aggregates all action events from a session with computed
// summary statistics. It provides the full action history for report rendering.
type ActionTimeline struct {
	Events        []ActionEvent   `json:"events"`
	TotalTurns    int             `json:"total_turns"`
	TotalDuration time.Duration   `json:"total_duration"`
	Summary       TimelineSummary `json:"summary"`
}

// TimelineSummary provides aggregate counts of action types within a timeline.
type TimelineSummary struct {
	ToolCalls  int `json:"tool_calls"`
	FileReads  int `json:"file_reads"`
	FileWrites int `json:"file_writes"`
	BashCmds   int `json:"bash_commands"`
}

// NewActionTimeline builds an ActionTimeline from a slice of events, computing
// the total turns, total duration, and per-type summary counts.
func NewActionTimeline(events []ActionEvent) *ActionTimeline {
	tl := &ActionTimeline{
		Events: events,
	}

	if len(events) == 0 {
		return tl
	}

	maxTurn := 0
	var totalDur time.Duration
	for _, e := range events {
		totalDur += e.Duration
		if e.TurnNumber > maxTurn {
			maxTurn = e.TurnNumber
		}
		switch e.Type {
		case "tool_call":
			tl.Summary.ToolCalls++
		case "file_read":
			tl.Summary.FileReads++
		case "file_write":
			tl.Summary.FileWrites++
		case "bash":
			tl.Summary.BashCmds++
		}
	}

	tl.TotalTurns = maxTurn
	tl.TotalDuration = totalDur
	return tl
}

// TruncateField truncates s to maxLen characters, appending an ellipsis
// indicator when truncation occurs. Returns s unchanged if within limit.
func TruncateField(s string, maxLen int) string {
	if maxLen <= 0 || len(s) <= maxLen {
		return s
	}
	suffix := "... [truncated]"
	if maxLen <= len(suffix) {
		return s[:maxLen]
	}
	return s[:maxLen-len(suffix)] + suffix
}
