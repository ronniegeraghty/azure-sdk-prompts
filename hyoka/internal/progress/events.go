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

// GraderPoint mirrors graders.GraderPoint inside the progress package so
// progress events can carry per-sub-check outcomes without importing the
// graders package (which would invert the existing layering). The engine
// copies values across at emission time.
type GraderPoint struct {
	Label   string // v4: renamed from Name to match graders.GraderPoint
	Pass    bool
	Message string
}

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
	GraderID         string   // Grader identifier (e.g. "prompt_review", "no_secrets")
	GraderKind       string   // Grader kind / model label (e.g. "claude-opus-4.6", "output_check")
	GraderSourceFile string   // Source file path for grouping (populated by engine when SourceFile known)
	GraderSourceType string   // "prompt_file" or "criteria_file"; empty = unknown
	Result           string   // One of GraderResultPass, GraderResultFail (EventGraderComplete)
	Score            *float64 // Optional grader score; nil means "not reported"

	// Points carries per-sub-check outcomes for EventGraderComplete (Phase 2
	// generalization). Empty / single-element slices preserve the original
	// flat single-row render path; len(Points) > 1 triggers the nested
	// renderer block. The progress package mirrors graders.GraderPoint to
	// avoid a graders→progress import cycle; the engine copies values
	// across at emission time.
	Points []GraderPoint

	// Session summary fields (EventSessionDetails).
	Files        []string // Files written by the session
	Turns        int      // Total turns consumed
	ToolCalls    int      // Total tool calls made
	Cost         float64  // Session cost in USD (deprecated — kept for back-compat with old recordings)
	InputTokens  int      // Total input tokens consumed by the generator session
	OutputTokens int      // Total output tokens emitted by the generator session

	// Guardrail fields (EventFailed, EventError).
	GuardrailReason string // Populated when a guardrail terminated the run (e.g., "turn limit (25)")

	// Grader point totals (EventPassed, EventFailed). Zero values mean no graders were run.
	GraderPointsPassed int // Number of checks that passed across all graders
	GraderPointsTotal  int // Total number of checks across all graders
}

// ProgressFunc receives progress events from evaluators.
type ProgressFunc func(ProgressEvent)

// Reporter is implemented by evaluators that support live progress updates.
type Reporter interface {
	SetProgressFunc(fn ProgressFunc)
}
