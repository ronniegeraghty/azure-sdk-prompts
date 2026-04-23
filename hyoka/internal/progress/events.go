package progress

// EventType classifies a progress event.
type EventType int

const (
	EventStarting      EventType = iota // Eval starting, waiting for session
	EventSendingPrompt                  // Sending prompt to Copilot
	EventReasoning                      // LLM is reasoning
	EventToolStart                      // Tool call initiated
	EventToolComplete                   // Tool call finished
	EventWritingFile                    // File write tool call
	EventWaiting                        // Waiting between events
	EventPhaseChange                    // Transition between eval phases
	EventPassed                         // Eval passed
	EventFailed                         // Eval failed
	EventError                          // Eval errored

	// Tool resolution & verification (config load / post-session-start).
	EventToolResolutionStart  // A skill/plugin/MCP tool has started resolving
	EventToolResolutionResult // Resolution finished (loaded or failed)
	EventToolsVerified        // Post-session-start verification bulk result

	// Per-grader lifecycle (serialized in interactive mode).
	EventGraderStart    // A grader started running
	EventGraderComplete // A grader finished (pass/fail)

	// Session-level details captured after generation completes.
	EventSessionDetails // Files, turns, tool calls, cost summary for a session
)

// Tool kind constants used with ToolKind fields on ProgressEvent.
const (
	ToolKindSkill  = "skill"
	ToolKindPlugin = "plugin"
	ToolKindMCP    = "mcp"
)

// Tool resolution / verification status constants used with Status fields.
const (
	ToolStatusLoaded = "loaded"
	ToolStatusFailed = "failed"
)

// Tool parent-kind constants used with ParentKind fields on ProgressEvent /
// ToolStatus. An empty ParentKind means the tool has no parent container
// (it was declared directly in the config). A non-empty ParentKind groups
// the leaf under its container for display.
const (
	ToolParentKindPlugin   = "plugin"
	ToolParentKindSkillDir = "skill_dir"
)

// Grader result constants used with the Result field.
const (
	GraderResultPass = "pass"
	GraderResultFail = "fail"
)

// ToolStatus captures the post-session-start verification outcome for a single tool.
// Used as the element type of ProgressEvent.Tools on EventToolsVerified.
type ToolStatus struct {
	ToolName   string // Tool identifier (skill name, plugin name, MCP server name)
	ToolKind   string // One of ToolKindSkill, ToolKindPlugin, ToolKindMCP
	Status     string // One of ToolStatusLoaded, ToolStatusFailed
	Reason     string // Optional human-readable reason (typically for failures)
	ParentName string // Parent container identifier (plugin name or skills-dir path); empty = no parent
	ParentKind string // One of ToolParentKindPlugin, ToolParentKindSkillDir, or empty
}

// Phase identifies which stage an eval is in.
type Phase string

const (
	PhaseGenerating Phase = "generating"
	PhaseReviewing  Phase = "reviewing"
)

// ProgressEvent carries status from the eval engine or Copilot session to the display.
//
// The struct is a fat union: only fields relevant to the EventType are populated.
// Existing consumers only read EventType-specific fields, so new fields are additive
// and backward-compatible. Pointer-typed fields (e.g. Score) signal "unset" distinct
// from the zero value.
type ProgressEvent struct {
	EvalID      string    // Unique eval identifier (promptID/configName)
	PromptID    string    // Prompt ID
	ConfigName  string    // Config name
	Type        EventType // What happened
	Message     string    // Human-readable activity description
	FileCount   int       // Generated file count (for completion events)
	Phase       Phase     // Current phase (for EventPhaseChange)
	ReviewScore int       // Review score out of 10 (for EventPassed)

	// Tool resolution / verification fields.
	// Populated by EventToolResolutionStart, EventToolResolutionResult, EventToolsVerified.
	ToolName string       // Tool identifier (single-tool events)
	ToolKind string       // One of ToolKindSkill, ToolKindPlugin, ToolKindMCP
	Status   string       // One of ToolStatusLoaded, ToolStatusFailed (single-tool events)
	Reason   string       // Optional explanation, typically for failures
	Tools    []ToolStatus // Bulk verification payload for EventToolsVerified

	// Parent grouping for grouped display. When set, the renderer may
	// indent this leaf under a parent row keyed by (ParentKind, ParentName).
	// Empty ParentKind / ParentName means the leaf has no container.
	ParentName string // Parent container identifier (plugin name or skills-dir path)
	ParentKind string // One of ToolParentKindPlugin, ToolParentKindSkillDir, or empty

	// Grader lifecycle fields (EventGraderStart, EventGraderComplete).
	GraderID   string   // Grader identifier (e.g. "prompt_review", "no_secrets")
	GraderKind string   // Grader kind / model label (e.g. "claude-opus-4.6", "output_check")
	Result     string   // One of GraderResultPass, GraderResultFail (EventGraderComplete)
	Score      *float64 // Optional grader score; nil means "not reported"

	// Session summary fields (EventSessionDetails).
	Files     []string // Files written by the session
	Turns     int      // Total turns consumed
	ToolCalls int      // Total tool calls made
	Cost      float64  // Session cost in USD

	// Guardrail fields (EventFailed, EventError).
	GuardrailReason string // Populated when a guardrail terminated the run (e.g., "turn limit (25)")
}

// ProgressFunc receives progress events from evaluators.
type ProgressFunc func(ProgressEvent)

// Reporter is implemented by evaluators that support live progress updates.
type Reporter interface {
	SetProgressFunc(fn ProgressFunc)
}
