// Package report handles generation of JSON and Markdown reports.
package report

import (
	"fmt"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/artifact"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/pairwise"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// WorkspaceDelta is a type alias for the workspace delta type.
// This allows report consumers to access workspace.WorkspaceDelta
// without creating an import cycle.
type WorkspaceDelta = workspace.WorkspaceDelta

// GeneratorArtifact is a type alias for the generator artifact type.
// This allows report consumers to access artifact.GeneratorArtifact
// without creating an import cycle.
type GeneratorArtifact = artifact.GeneratorArtifact

// CurrentSchemaVersion is the latest report schema version.
//
//	v1 = legacy monolithic ReviewResult.
//	v2 = grader-based; one expanded GraderResult entry per panel member +
//	     one consensus row for prompt_review graders.
//	v3 = single-entry ai_review row carrying GraderPoints; ToolLoadResult
//	     and EnvironmentInfo carry parent linkage; pre-computed roll-ups.
//	v4 = unified Points-first GraderResult; all graders emit Points; Extras
//	     discriminated union replaces per-kind detail fields.
//
// v4 is a hard cutover: v3 reports are rejected at load time (regenerate required).
const CurrentSchemaVersion = 4

// GraderResult holds the output from a single grader (v4 schema).
// Pass and Score are derived from Points. Every grader emits at least one Point.
type GraderResult struct {
	GraderName string  `json:"grader_name"`
	GraderType string  `json:"grader_type"` // "file", "program", "prompt", "behavior", etc.
	Score      float64 `json:"score"`       // 0.0–1.0 weighted average from Points
	Weight     float64 `json:"weight"`      // Weight for aggregation
	Pass       bool    `json:"pass"`        // AND of all Points[i].Pass
	Gate       bool    `json:"gate,omitempty"`
	Message    string  `json:"message"`     // Human-readable summary

	Points []GraderPoint  `json:"points"`           // REQUIRED, len ≥ 1
	Extras *GraderExtras  `json:"extras,omitempty"` // kind-specific render-only data

	// Provenance: where the grader entry was declared.
	SourceFile string `json:"source_file,omitempty"`
	SourceType string `json:"source_type,omitempty"` // "prompt_file" or "criteria_file"
}

// GraderPoint is one binary pass/fail check inside a grader.
type GraderPoint struct {
	Label    string             `json:"label"`              // what was checked
	Pass     bool               `json:"pass"`
	Message  string             `json:"message,omitempty"`  // why it passed/failed
	Weight   float64            `json:"weight,omitempty"`   // for Score weighting; defaults to 1.0
	Evidence map[string]string  `json:"evidence,omitempty"` // optional string-only KV
}

// GraderExtras is a discriminated union carrying kind-specific render-only data.
type GraderExtras struct {
	File           *FileExtras           `json:"file,omitempty"`
	Program        *ProgramExtras        `json:"program,omitempty"`
	Prompt         *PromptExtras         `json:"prompt,omitempty"`
	Behavior       *BehaviorExtras       `json:"behavior,omitempty"`
	ActionSequence *ActionSequenceExtras `json:"action_sequence,omitempty"`
	ToolConstraint *ToolConstraintExtras `json:"tool_constraint,omitempty"`
	OutputCheck    *OutputCheckExtras    `json:"output_check,omitempty"`
	Review         *ReviewExtras         `json:"review,omitempty"`
}

// FileExtras holds file-check render-only data.
type FileExtras struct {
	Files []FileExtra `json:"files"`
}

// FileExtra records one file check's outcome.
type FileExtra struct {
	Path           string `json:"path"`
	Exists         bool   `json:"exists"`
	Pattern        string `json:"pattern,omitempty"`
	PatternMatched bool   `json:"pattern_matched,omitempty"`
	Size           int64  `json:"size,omitempty"`
}

// ProgramExtras holds program execution render-only data.
type ProgramExtras struct {
	Command    string   `json:"command"`
	Args       []string `json:"args,omitempty"`
	ExitCode   int      `json:"exit_code"`
	Stdout     string   `json:"stdout,omitempty"`
	Stderr     string   `json:"stderr,omitempty"`
	DurationMs int64    `json:"duration_ms,omitempty"`
}

// PromptExtras holds LLM-as-judge render-only data.
type PromptExtras struct {
	Model     string `json:"model"`
	Rubric    string `json:"rubric"`
	Reasoning string `json:"reasoning"`
	RawScore  int    `json:"raw_score,omitempty"`
	MaxScore  int    `json:"max_score,omitempty"`
}

// BehaviorExtras holds behavior grader render-only data.
type BehaviorExtras struct {
	ToolsUsed      []string `json:"tools_used,omitempty"`
	MissingTools   []string `json:"missing_tools,omitempty"`
	ForbiddenUsed  []string `json:"forbidden_used,omitempty"`
	TurnCount      int      `json:"turn_count,omitempty"`
	MaxTurns       int      `json:"max_turns,omitempty"`
	TotalActions   int      `json:"total_actions,omitempty"`
	TurnLimitHit   bool     `json:"turn_limit_hit,omitempty"`
	Violations     []string `json:"violations,omitempty"`
}

// ActionSequenceExtras holds action_sequence grader render-only data.
type ActionSequenceExtras struct {
	ExpectedSequence []string `json:"expected_sequence"`
	ActualSequence   []string `json:"actual_sequence"`
	MatchedActions   int      `json:"matched_actions"`
	ToolsUsed        []string `json:"tools_used,omitempty"`
	TotalActions     int      `json:"total_actions"`
}

// ToolConstraintExtras holds tool_constraint grader render-only data.
type ToolConstraintExtras struct {
	ToolsUsed      []string       `json:"tools_used,omitempty"`
	ToolCounts     map[string]int `json:"tool_counts,omitempty"`
	MissingTools   []string       `json:"missing_tools,omitempty"`
	ForbiddenUsed  []string       `json:"forbidden_used,omitempty"`
	Violations     []string       `json:"violations,omitempty"`
	ConstraintsMet bool           `json:"constraints_met"`
}

// OutputCheckExtras holds output_check grader render-only data.
type OutputCheckExtras struct {
	ProducedFiles []FileEntry `json:"produced_files"`
}

// FileEntry describes a file in the agent workspace.
type FileEntry struct {
	Path string `json:"path"` // Relative to workspace root
	Size int64  `json:"size"`
}

// ReviewExtras holds prompt_review grader render-only data.
type ReviewExtras struct {
	Model           string              `json:"model,omitempty"`
	Summary         string              `json:"summary"`
	IsConsensus     bool                `json:"is_consensus"`
	PanelResults    []ReviewPanelResult `json:"panel_results,omitempty"`
	Issues          []string            `json:"issues,omitempty"`
	Strengths       []string            `json:"strengths,omitempty"`
	DurationSeconds float64             `json:"duration_seconds,omitempty"`
}

// ReviewPanelResult holds one panel member's review result.
type ReviewPanelResult struct {
	Model     string                   `json:"model"`
	Score     int                      `json:"score"`
	Pass      bool                     `json:"pass"`
	Issues    []string                 `json:"issues,omitempty"`
	Strengths []string                 `json:"strengths,omitempty"`
	Criteria  []ReviewCriterionResult  `json:"criteria"`
}

// ReviewCriterionResult holds one criterion's outcome.
type ReviewCriterionResult struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Reason string `json:"reason,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

// --- Legacy detail types below (v3 compat; will be removed in future cleanup) ---

// FileCheckDetail records the outcome of a single file check.
type FileCheckDetail struct {
	Path           string `json:"path"`
	Exists         bool   `json:"exists"`
	PatternMatched *bool  `json:"pattern_matched,omitempty"`
	Pattern        string `json:"pattern,omitempty"`
}

// FileGraderDetail holds file-check specifics.
type FileGraderDetail struct {
	CheckedFiles []FileCheckDetail `json:"checked_files"`
}

// ProgramGraderDetail holds program execution specifics.
type ProgramGraderDetail struct {
	Command  string `json:"command"`
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
}

// PromptGraderDetail holds LLM-as-judge specifics.
type PromptGraderDetail struct {
	Model     string `json:"model"`
	Rubric    string `json:"rubric"`
	Reasoning string `json:"reasoning"`
	RawScore  int    `json:"raw_score,omitempty"`
	MaxScore  int    `json:"max_score,omitempty"`
}

// BehaviorGraderDetail holds agent behavior analysis specifics.
type BehaviorGraderDetail struct {
	ToolsUsed        []string       `json:"tools_used,omitempty"`
	MissingTools     []string       `json:"missing_tools,omitempty"`
	ForbiddenUsed    []string       `json:"forbidden_used,omitempty"`
	TurnCount        int            `json:"turn_count,omitempty"`
	MaxTurns         int            `json:"max_turns,omitempty"`
	ActualTurns      int            `json:"actual_turns,omitempty"`
	TotalActions     int            `json:"total_actions,omitempty"`
	TurnLimitHit     bool           `json:"turn_limit_hit,omitempty"`
	Violations       []string       `json:"violations,omitempty"`
	SequenceMatch    bool           `json:"sequence_match,omitempty"`
	ExpectedSequence []string       `json:"expected_sequence,omitempty"`
	ActualSequence   []string       `json:"actual_sequence,omitempty"`
	MatchedActions   int            `json:"matched_actions,omitempty"`
	ConstraintsMet   bool           `json:"constraints_met,omitempty"`
	ToolCounts       map[string]int `json:"tool_counts,omitempty"`
	// Weighted aggregation fields (#143)
	Score  float64 `json:"score,omitempty"`  // 0.0–1.0 normalized score
	Weight float64 `json:"weight,omitempty"` // Weight for aggregation (default 1.0)
	Gate   bool    `json:"gate,omitempty"`   // If true, failure overrides weighted scoring
	Pass   bool    `json:"pass,omitempty"`   // Binary pass/fail
}

// ReviewGraderDetail holds AI review grader specifics (WI-023).
type ReviewGraderDetail struct {
	Model        string                      `json:"model,omitempty"`
	OverallScore int                         `json:"overall_score"`
	MaxScore     int                         `json:"max_score"`
	Issues       []string                    `json:"issues,omitempty"`
	Strengths    []string                    `json:"strengths,omitempty"`
	IsConsensus  bool                        `json:"is_consensus,omitempty"`
	Criteria     []ReviewGraderCriterion     `json:"criteria,omitempty"`
	PanelResults []ReviewGraderPanelEntry    `json:"panel_results,omitempty"`
}

// ReviewGraderCriterion holds a single criterion result from an AI review grader.
type ReviewGraderCriterion struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Reason string `json:"reason,omitempty"`
}

// ReviewGraderPanelEntry holds one panel member's review in a review grader.
type ReviewGraderPanelEntry struct {
	Model        string                  `json:"model"`
	OverallScore int                     `json:"overall_score"`
	MaxScore     int                     `json:"max_score"`
	Summary      string                  `json:"summary"`
	Issues       []string                `json:"issues,omitempty"`
	Strengths    []string                `json:"strengths,omitempty"`
	Criteria     []ReviewGraderCriterion `json:"criteria,omitempty"`
}

// ScoreContribution describes a single grader's contribution to the final score.
type ScoreContribution struct {
	Name            string  `json:"name"`
	Kind            string  `json:"kind"`
	Score           float64 `json:"score"`
	Weight          float64 `json:"weight"`
	WeightedScore   float64 `json:"weighted_score"` // score × weight
	Gate            bool    `json:"gate,omitempty"`
	Pass            bool    `json:"pass"`
	ContributionPct float64 `json:"contribution_pct"` // (weighted_score / totalWeightedSum) × 100
}

// ScoreBreakdown shows how the final aggregated score is computed (#143).
type ScoreBreakdown struct {
	Formula         string              `json:"formula"`
	Contributions   []ScoreContribution `json:"contributions"`
	TotalWeight     float64             `json:"total_weight"`
	WeightedSum     float64             `json:"weighted_sum"`
	FinalScore      float64             `json:"final_score"`
	FinalScorePct   float64             `json:"final_score_pct"` // FinalScore × 100
	GateFailed      bool                `json:"gate_failed,omitempty"`
	GateFailedNames []string            `json:"gate_failed_names,omitempty"`
}

// SessionEventRecord is a serializable representation of a Copilot session event.
type SessionEventRecord struct {
	Type          string  `json:"type"`
	ToolName      string  `json:"tool_name,omitempty"`
	ToolArgs      string  `json:"tool_args,omitempty"`
	Content       string  `json:"content,omitempty"`
	Error         string  `json:"error,omitempty"`
	ToolResult    string  `json:"tool_result,omitempty"`
	ToolSuccess   *bool   `json:"tool_success,omitempty"`
	Duration      float64 `json:"duration_ms,omitempty"`
	MCPServerName string  `json:"mcp_server_name,omitempty"`
	MCPToolName   string  `json:"mcp_tool_name,omitempty"`
	FilePath      string  `json:"file_path,omitempty"`
	// Expanded event fields
	TurnNumber    int    `json:"turnNumber,omitempty"`
	InputTokens   int    `json:"inputTokens,omitempty"`
	OutputTokens  int    `json:"outputTokens,omitempty"`
	SkillName     string `json:"skillName,omitempty"`
	CommandText   string `json:"commandText,omitempty"`
	FileOperation string `json:"fileOperation,omitempty"`
	FileSize      int64  `json:"fileSize,omitempty"`
	SubagentID    string `json:"subagentId,omitempty"`
	IsTruncation  bool   `json:"isTruncation,omitempty"`
	Intent        string `json:"intent,omitempty"`
	WarningText   string `json:"warningText,omitempty"`
}

// ToolUsageResult holds the comparison of expected vs actual tool usage.
type ToolUsageResult struct {
	ExpectedTools []string `json:"expected_tools"`
	ActualTools   []string `json:"actual_tools"`
	MatchedTools  []string `json:"matched_tools"`
	MissingTools  []string `json:"missing_tools"`
	ExtraTools    []string `json:"extra_tools"`
	Match         bool     `json:"tool_usage_match"`
}

// ToolAvailabilityEntry records whether a specific tool/skill/MCP server
// was available to the agent and whether it was actually used during the session.
type ToolAvailabilityEntry struct {
	Name      string `json:"name"`
	Type      string `json:"type"`      // "tool", "skill", "mcp"
	Available bool   `json:"available"`
	Used      bool   `json:"used"`
}

// ReviewedFile holds an annotated code file with inline review comments.
type ReviewedFile struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// EnvironmentInfo captures session environment and configuration metadata.
type EnvironmentInfo struct {
	Model             string   `json:"model"`
	SkillDirectories  []string `json:"skillDirectories,omitempty"`
	SkillsInvoked     []string `json:"skillsInvoked,omitempty"`
	SkillsLoaded      []string `json:"skillsLoaded,omitempty"`
	// SkillGroups carries the structured view of loaded skills with parent
	// linkage (Phase 5, schema v3). It is a sibling to SkillsLoaded rather
	// than a replacement (Plan 5.5 Option B): SkillsLoaded stays as the
	// flat string list site components have consumed since v1, and
	// SkillGroups is additive so consumers that want the full topology
	// (skill name + parent plugin/skill_dir + kind) can read it without
	// breaking older site builds. Empty / nil for v2 reports and any v3
	// run that did not emit SDK skill events.
	SkillGroups       []SkillLoadEntry `json:"skill_groups,omitempty"`
	AvailableTools    []string         `json:"availableTools,omitempty"`
	ExcludedTools     []string         `json:"excludedTools,omitempty"`
	MCPServers        []string         `json:"mcpServers,omitempty"`
	MCPToolsInvoked   []string         `json:"mcpToolsInvoked,omitempty"`
	SafetyBoundaries  bool             `json:"safetyBoundaries"`
	AllowCloud        bool             `json:"allowCloud"`
	WorkingDirectory  string           `json:"workingDirectory"`
	TotalInputTokens  int              `json:"totalInputTokens,omitempty"`
	TotalOutputTokens int              `json:"totalOutputTokens,omitempty"`
	TurnCount         int              `json:"turnCount,omitempty"`
	ContextTruncated  bool             `json:"contextTruncated,omitempty"`
}

// SkillLoadEntry is one row in EnvironmentInfo.SkillGroups — a loaded
// skill plus the container (plugin or skill_dir) it came from. Kind is
// either "skill" or "mcp"; Parent / ParentKind are empty when the entry
// was declared at the top level of the config (i.e., not via a plugin).
type SkillLoadEntry struct {
	Name       string `json:"name"`
	Parent     string `json:"parent,omitempty"`
	Kind       string `json:"kind,omitempty"`        // "skill" | "mcp"
	ParentKind string `json:"parent_kind,omitempty"` // "plugin" | "skill_dir"
}

// ResourceStats holds per-eval peak resource utilization (#45).
type ResourceStats struct {
	PeakCPUPercent float64 `json:"peak_cpu_percent"`
	PeakMemoryMB   float64 `json:"peak_memory_mb"`
	SampleCount    int     `json:"sample_count"`
}

// ActionTimelineReport is the serializable form of an action timeline for JSON reports (#139).
type ActionTimelineReport struct {
	Events  []ActionEventReport   `json:"events"`
	Entries []ActionTimelineEntry `json:"entries,omitempty"`
	Summary ActionSummaryReport   `json:"summary"`
}

// ActionEventReport represents a single agent action in a report.
type ActionEventReport struct {
	Sequence   int     `json:"sequence"`
	Type       string  `json:"type"`
	Tool       string  `json:"tool,omitempty"`
	Action     string  `json:"action,omitempty"`
	Path       string  `json:"path,omitempty"`
	Input      string  `json:"input,omitempty"`
	Output     string  `json:"output,omitempty"`
	Error      string  `json:"error,omitempty"`
	Success    *bool   `json:"success,omitempty"`
	DurationMs float64 `json:"duration_ms,omitempty"`
	TurnNumber int     `json:"turn_number,omitempty"`
	MCPServer  string  `json:"mcp_server,omitempty"`
}

// ActionSummaryReport holds aggregate statistics for the timeline.
type ActionSummaryReport struct {
	TotalEvents      int            `json:"total_events"`
	TotalTurns       int            `json:"total_turns"`
	TotalActions     int            `json:"total_actions"`
	TotalToolCalls   int            `json:"total_tool_calls"`
	ToolCalls        int            `json:"tool_calls"`
	FileReads        int            `json:"file_reads"`
	FileWrites       int            `json:"file_writes"`
	BashCommands     int            `json:"bash_commands"`
	MCPCalls         int            `json:"mcp_calls"`
	Errors           int            `json:"errors"`
	ToolBreakdown    map[string]int `json:"tool_breakdown"`
	TotalDurationMs  float64        `json:"total_duration_ms,omitempty"`
	ToolCallDuration float64        `json:"tool_call_duration_ms,omitempty"`
	ToolSuccesses    int            `json:"tool_successes"`
	ToolFailures     int            `json:"tool_failures"`
}

// EvalReport contains the results of a single prompt evaluation.
type EvalReport struct {
	SchemaVersion      int                   `json:"schema_version"`
	PromptID           string                `json:"prompt_id"`
	ConfigName         string                `json:"config_name"`
	Timestamp          string                `json:"timestamp"`
	Duration           float64               `json:"duration_seconds"`
	GenerationDuration float64               `json:"generation_duration_seconds,omitempty"`
	ReviewDuration     float64               `json:"review_duration_seconds,omitempty"`
	PromptMeta         map[string]any        `json:"prompt_metadata"`
	ConfigUsed         map[string]any        `json:"config_used"`
	GeneratedFiles     []string              `json:"generated_files"`
	StarterFiles       []string              `json:"starter_files,omitempty"`
	ReviewedFiles      []ReviewedFile           `json:"reviewed_files,omitempty"`
	// FileContents maps file paths to their contents for site display.
	// Populated at report-build time from the workspace directory. Files
	// exceeding 1MB are capped with a truncation marker.
	FileContents       map[string]string     `json:"file_contents,omitempty"`
	// GeneratorArtifact captures the complete state of the generator session.
	// Populated at report-build time from generator.json if it exists.
	GeneratorArtifact  *GeneratorArtifact    `json:"generator_artifact,omitempty"`
	Review             *review.ReviewResult     `json:"review,omitempty"`
	ReviewPanel        []review.ReviewResult    `json:"review_panel,omitempty"`
	SkippedReviewers   []review.SkippedReviewer `json:"skipped_reviewers,omitempty"`
	GraderResults      []GraderResult           `json:"grader_results,omitempty"`
	ToolUsage          *ToolUsageResult        `json:"tool_usage,omitempty"`
	ToolAvailability   []ToolAvailabilityEntry `json:"tool_availability,omitempty"` // Tools available vs used (#348)
	SessionEvents      []SessionEventRecord  `json:"session_events,omitempty"`
	ActionTimeline     *ActionTimelineReport `json:"action_timeline,omitempty"` // Structured action log (#139)
	EventCount         int                   `json:"event_count"`
	ToolCalls          []string              `json:"tool_calls"`
	Environment        *EnvironmentInfo      `json:"environment,omitempty"`
	ResourceUsage      *ResourceStats        `json:"resource_usage,omitempty"`  // Per-eval resource stats (#45)
	ScoreBreakdown     *ScoreBreakdown       `json:"score_breakdown,omitempty"` // Weighted aggregation breakdown (#143)
	SessionSetup       *SessionSetupEvent    `json:"session_setup,omitempty"`   // Tool/skill/MCP loading status (#219)
	WorkspaceDelta     *WorkspaceDelta       `json:"workspace_delta,omitempty"` // File-level changes agent made (#566)
	Success            bool                  `json:"success"`
	Error              string                `json:"error,omitempty"`
	ErrorDetails       string                `json:"error_details,omitempty"`
	ErrorCategory      string                `json:"error_category,omitempty"` // timeout, sdk_error, generation_failure, review_failure
	FailureReason      string                `json:"failure_reason,omitempty"` // human-readable explanation of failure
	IsStub             bool                  `json:"is_stub,omitempty"`
	RerunCommand       string                `json:"rerunCommand,omitempty"`
	BaseConfigName     string                `json:"baseConfigName,omitempty"`  // Config name before fan-out (e.g., "python-pairwise")
	GeneratorModel     string                `json:"generatorModel,omitempty"`  // Specific model used (e.g., "claude-opus-4.6")
	PairwiseVariant    string                `json:"pairwiseVariant,omitempty"` // Pairwise variant suffix (e.g., "baseline", "without-azure", "without-azure/storage_blob_list")
	// Generator guardrails (#35) — Phase 3.5 (#566) dropped the byte-size
	// cap entirely; MaxFiles, MaxTurns, and MaxSessionActions remain as hard
	// fails that set GuardrailAbortReason.
	GuardrailMaxTurns          int    `json:"guardrail_max_turns,omitempty"`
	GuardrailMaxFiles          int    `json:"guardrail_max_files,omitempty"`
	GuardrailMaxSessionActions int    `json:"guardrail_max_session_actions,omitempty"`
	GuardrailAbortReason       string `json:"guardrail_abort_reason,omitempty"`
	// Action limit soft cap — generation stopped but review proceeds with partial results
	ActionLimitReached bool `json:"action_limit_reached,omitempty"`
	ActionCount        int  `json:"action_count,omitempty"`
	// Roll-up fields populated at engine time from the unified grader
	// aggregate (#schema_v3). The site reads these directly so it never
	// has to recompute pass/total from GraderResults — eliminates the
	// entire class of roll-up-divergence bugs by construction. omitempty:
	// zero-value (e.g., evals that never ran graders) emits no field, so
	// v2 reports round-trip unchanged.
	GradersPassed int `json:"graders_passed,omitempty"`
	GradersTotal  int `json:"graders_total,omitempty"`
}

// ToolLoadResult records the outcome of loading a single tool, skill, or MCP server.
//
// As of schema v3, ToolLoadResult also carries parent linkage so the report
// can express the same grouped relationships the live renderer shows during
// validation: a plugin or skill_dir parent emits a "container" row with
// Status empty (parents have no runtime status), and each child carries
// Kind / Parent / ParentKind back-pointers. All linkage fields are
// `omitempty` so v2 reports remain valid when unmarshaled into v3 structs
// and v3 reports without parent info round-trip identically.
type ToolLoadResult struct {
	Name       string `json:"name"`
	Status     string `json:"status,omitempty"`      // omitempty: parents omit
	Error      string `json:"error,omitempty"`
	Details    string `json:"details,omitempty"`     // e.g., command string for MCP servers
	Kind       string `json:"kind,omitempty"`        // "skill" | "mcp" | "plugin" | "skill_dir"
	Parent     string `json:"parent,omitempty"`      // container this entry is a child of
	ParentKind string `json:"parent_kind,omitempty"` // "plugin" | "skill_dir"
}

// SessionSetupEvent captures the tool/skill/MCP loading results at session start.
type SessionSetupEvent struct {
	MCPServers   []ToolLoadResult `json:"mcp_servers,omitempty"`
	Skills       []ToolLoadResult `json:"skills,omitempty"`
	Tools        []string         `json:"tools_available,omitempty"`
	SystemPrompt string           `json:"system_prompt_status"` // "custom (N chars)" or "none (default)"
	StarterFiles []string         `json:"starter_files,omitempty"`
}

// ActionTimelineEntry represents a single action in the session timeline.
type ActionTimelineEntry struct {
	Index        int                `json:"index"`
	TurnNumber   int                `json:"turn_number,omitempty"`
	Type         string             `json:"type"`
	ToolName     string             `json:"tool_name,omitempty"`
	Duration     float64            `json:"duration_ms,omitempty"`
	Success      *bool              `json:"success,omitempty"`
	Error        string             `json:"error,omitempty"`
	MCPServer    string             `json:"mcp_server,omitempty"`
	FilePath     string             `json:"file_path,omitempty"`
	Intent       string             `json:"intent,omitempty"`
	SessionSetup *SessionSetupEvent `json:"session_setup,omitempty"`
}

// ActionTimelineSummary holds aggregate statistics for the action timeline.
type ActionTimelineSummary struct {
	TotalActions     int     `json:"total_actions"`
	TotalToolCalls   int     `json:"total_tool_calls"`
	TotalTurns       int     `json:"total_turns"`
	ToolCallDuration float64 `json:"tool_call_duration_ms"`
	ToolSuccesses    int     `json:"tool_successes"`
	ToolFailures     int     `json:"tool_failures"`
}

// BuildActionTimeline derives a structured action timeline from SessionEvents.
// This is a fallback for reports that don't have a pre-built timeline from the
// eval engine (e.g., when re-rendering from JSON).
func BuildActionTimeline(events []SessionEventRecord) *ActionTimelineReport {
	return BuildActionTimelineWithSetup(events, nil)
}

// BuildActionTimelineWithSetup derives a structured action timeline from
// SessionEvents and prepends a session_setup entry when setup data is available.
func BuildActionTimelineWithSetup(events []SessionEventRecord, setup *SessionSetupEvent) *ActionTimelineReport {
	if len(events) == 0 && setup == nil {
		return nil
	}
	var reportEvents []ActionEventReport
	var entries []ActionTimelineEntry
	var summary ActionSummaryReport
	seq := 0
	index := 0
	turnNumber := 0

	// Synthesize session_setup entry from config + SDK events.
	synthesized := setup
	if synthesized == nil {
		synthesized = &SessionSetupEvent{}
	}
	// Enrich from SDK events (skills_loaded, mcp_servers_loaded).
	for _, ev := range events {
		switch ev.Type {
		case "session.skills_loaded":
			if ev.Content != "" {
				for _, name := range splitCSV(ev.Content) {
					if !hasToolLoadResult(synthesized.Skills, name) {
						synthesized.Skills = append(synthesized.Skills, ToolLoadResult{
							Name:   name,
							Status: "loaded",
						})
					}
				}
			}
		case "session.mcp_servers_loaded":
			if ev.Content != "" {
				for _, name := range splitCSV(ev.Content) {
					if !hasToolLoadResult(synthesized.MCPServers, name) {
						synthesized.MCPServers = append(synthesized.MCPServers, ToolLoadResult{
							Name:   name,
							Status: "loaded",
						})
					}
				}
			}
		}
	}

	// Only emit the entry when there is meaningful data.
	if len(synthesized.MCPServers) > 0 || len(synthesized.Skills) > 0 ||
		len(synthesized.Tools) > 0 || synthesized.SystemPrompt != "" ||
		len(synthesized.StarterFiles) > 0 {
		index++
		entries = append(entries, ActionTimelineEntry{
			Index:        index,
			Type:         "session_setup",
			SessionSetup: synthesized,
		})
		summary.TotalActions++
	}

	for _, ev := range events {
		seq++
		switch ev.Type {
		case "assistant.turn_start":
			turnNumber++
			summary.TotalTurns++
			summary.TotalActions++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "turn_start",
			})
		case "assistant.turn_end":
			summary.TotalActions++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "turn_end",
			})
		case "assistant.reasoning":
			summary.TotalActions++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "reasoning",
			})
		case "assistant.message":
			summary.TotalActions++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "message",
			})
		case "assistant.intent":
			summary.TotalActions++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "intent",
				Intent:     ev.Intent,
			})
		case "tool.execution_start":
			summary.TotalToolCalls++
			summary.TotalActions++
			summary.ToolCalls++
			index++
			entries = append(entries, ActionTimelineEntry{
				Index:      index,
				TurnNumber: turnNumber,
				Type:       "tool_call",
				ToolName:   ev.ToolName,
				MCPServer:  ev.MCPServerName,
				FilePath:   ev.FilePath,
			})
			reportEvents = append(reportEvents, ActionEventReport{
				Sequence:   seq,
				Type:       "tool_call",
				Tool:       ev.ToolName,
				Path:       ev.FilePath,
				TurnNumber: turnNumber,
				MCPServer:  ev.MCPServerName,
			})
		case "tool.execution_complete":
			summary.ToolCallDuration += ev.Duration
			summary.TotalDurationMs += ev.Duration
			// Update the matching tool_call entry with completion data.
			for i := len(entries) - 1; i >= 0; i-- {
				if entries[i].Type == "tool_call" && entries[i].ToolName == ev.ToolName && entries[i].Duration == 0 {
					entries[i].Duration = ev.Duration
					entries[i].Success = ev.ToolSuccess
					entries[i].Error = ev.Error
					break
				}
			}
			if ev.ToolSuccess != nil {
				if *ev.ToolSuccess {
					summary.ToolSuccesses++
				} else {
					summary.ToolFailures++
				}
			}
		}
	}
	summary.TotalEvents = len(events)
	return &ActionTimelineReport{
		Events:  reportEvents,
		Entries: entries,
		Summary: summary,
	}
}

// splitCSV splits a comma-separated string and trims whitespace from each element.
func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// hasToolLoadResult checks whether a name already exists in a slice of ToolLoadResult.
func hasToolLoadResult(results []ToolLoadResult, name string) bool {
	for _, r := range results {
		if r.Name == name {
			return true
		}
	}
	return false
}

// RunResourceStats holds aggregate resource utilization across all evals (#45).
type RunResourceStats struct {
	PeakCPUPercent float64 `json:"peak_cpu_percent"`
	PeakMemoryMB   float64 `json:"peak_memory_mb"`
	SessionCount   int     `json:"session_count"`
}

// RunSummary contains aggregate statistics for an evaluation run.
type RunSummary struct {
	RunID                 string                     `json:"run_id"`
	Timestamp             string                     `json:"timestamp"`
	TotalPrompts          int                        `json:"total_prompts"`
	TotalConfigs          int                        `json:"total_configs"`
	TotalEvals            int                        `json:"total_evaluations"`
	Passed                int                        `json:"passed"`
	Failed                int                        `json:"failed"`
	Errors                int                        `json:"errors"`
	Duration              float64                    `json:"duration_seconds"`
	AvgGenerationDuration float64                    `json:"avg_generation_duration_seconds,omitempty"`
	AvgReviewDuration     float64                    `json:"avg_review_duration_seconds,omitempty"`
	Reports               []string                   `json:"report_paths"`
	Results               []*EvalReport              `json:"results,omitempty"`
	Analysis              string                     `json:"analysis,omitempty"`
	ResourceUsage         *RunResourceStats          `json:"resource_usage,omitempty"`   // Aggregate resource stats (#45)
	PairwiseResults       []*pairwise.PairwiseReport `json:"pairwise_results,omitempty"` // Per-prompt pairwise impact (#123)
}

// BuildScoreBreakdown computes a ScoreBreakdown from the grader results on a report.
// It mirrors the AggregateResults logic from the graders package to show users
// exactly how the final score was derived.
func BuildScoreBreakdown(results []GraderResult) *ScoreBreakdown {
	if len(results) == 0 {
		return nil
	}

	// Skip breakdown for legacy review-only results that lack Score/Weight data.
	hasWeightedData := false
	for _, r := range results {
		if r.Score > 0 || r.Weight > 0 || r.Gate {
			hasWeightedData = true
			break
		}
	}
	if !hasWeightedData {
		return nil
	}

	sb := &ScoreBreakdown{
		Formula: "Final Score = Σ(grader_score × weight) / Σ(weights)",
	}

	// Check gate failures first.
	for _, r := range results {
		if r.Gate && !r.Pass {
			sb.GateFailed = true
			sb.GateFailedNames = append(sb.GateFailedNames, r.GraderName)
		}
	}

	var totalWeight float64
	var weightedSum float64
	for _, r := range results {
		w := r.Weight
		if w == 0 {
			w = 1.0
		}
		ws := r.Score * w
		totalWeight += w
		weightedSum += ws

		sb.Contributions = append(sb.Contributions, ScoreContribution{
			Name:          r.GraderName,
			Kind:          r.GraderType,
			Score:         r.Score,
			Weight:        w,
			WeightedScore: ws,
			Gate:          r.Gate,
			Pass:          r.Pass,
		})
	}

	sb.TotalWeight = totalWeight
	sb.WeightedSum = weightedSum
	if totalWeight > 0 {
		sb.FinalScore = weightedSum / totalWeight
	}
	if sb.GateFailed {
		sb.FinalScore = 0
		sb.Formula = fmt.Sprintf("Final Score = 0 (gate grader failed: %s)", strings.Join(sb.GateFailedNames, ", "))
	}
	sb.FinalScorePct = sb.FinalScore * 100

	// Compute contribution percentages.
	for i := range sb.Contributions {
		if weightedSum > 0 && !sb.GateFailed {
			sb.Contributions[i].ContributionPct = (sb.Contributions[i].WeightedScore / weightedSum) * 100
		}
	}

	return sb
}

// GraderResultsFromReview converts legacy ReviewResult data into []GraderResult.
// If panel is non-empty, each panel member becomes a grader entry and the
// consolidated review becomes the consensus entry.
//
// NOTE: This function is deprecated in v4 but kept for v1→v2 migration path.
// v4 reports never call this; only legacy migration uses it.
func GraderResultsFromReview(consolidated *review.ReviewResult, panel []review.ReviewResult) []GraderResult {
	if consolidated == nil {
		return nil
	}
	var results []GraderResult
	// v1→v2 migration stub: create minimal grader entries from legacy review data.
	// In v4, all new reports come from graders directly via convertGraderResults().
	for i := range panel {
		r := &panel[i]
		name := r.Model
		if name == "" {
			name = "reviewer"
		}
		// Create stub Points from legacy Scores (each criterion → one Point).
		var points []GraderPoint
		for _, crit := range r.Scores.Criteria {
			points = append(points, GraderPoint{
				Label:   crit.Name,
				Pass:    crit.Passed,
				Message: crit.Reason,
				Weight:  1.0,
			})
		}
		if len(points) == 0 {
			// Fallback: create one point from OverallScore.
			points = []GraderPoint{{
				Label:  "overall",
				Pass:   r.OverallScore >= r.MaxScore,
				Weight: 1.0,
			}}
		}
		// Derive Score from Points (matching v4 invariant).
		var totalWeight, weightedSum float64
		for _, p := range points {
			w := p.Weight
			if w == 0 {
				w = 1.0
			}
			totalWeight += w
			if p.Pass {
				weightedSum += w
			}
		}
		score := 0.0
		if totalWeight > 0 {
			score = weightedSum / totalWeight
		}

		results = append(results, GraderResult{
			GraderName: name,
			GraderType: "review",
			Score:      score,
			Weight:     1.0,
			Pass:       r.OverallScore >= r.MaxScore,
			Message:    r.Summary,
			Points:     points,
			Extras: &GraderExtras{
				Review: &ReviewExtras{
					Model:       r.Model,
					Summary:     r.Summary,
					IsConsensus: false,
					Issues:      r.Issues,
					Strengths:   r.Strengths,
				},
			},
		})
	}
	consensusName := consolidated.Model
	if consensusName == "" {
		consensusName = "consensus"
	}

	// Convert consensus result.
	var consPoints []GraderPoint
	for _, crit := range consolidated.Scores.Criteria {
		consPoints = append(consPoints, GraderPoint{
			Label:   crit.Name,
			Pass:    crit.Passed,
			Message: crit.Reason,
			Weight:  1.0,
		})
	}
	if len(consPoints) == 0 {
		consPoints = []GraderPoint{{
			Label:  "overall",
			Pass:   consolidated.OverallScore >= consolidated.MaxScore,
			Weight: 1.0,
		}}
	}

	var consTotalWeight, consWeightedSum float64
	for _, p := range consPoints {
		w := p.Weight
		if w == 0 {
			w = 1.0
		}
		consTotalWeight += w
		if p.Pass {
			consWeightedSum += w
		}
	}
	consScore := 0.0
	if consTotalWeight > 0 {
		consScore = consWeightedSum / consTotalWeight
	}

	results = append(results, GraderResult{
		GraderName: consensusName,
		GraderType: "review",
		Score:      consScore,
		Weight:     1.0,
		Pass:       consolidated.OverallScore >= consolidated.MaxScore,
		Message:    consolidated.Summary,
		Points:     consPoints,
		Extras: &GraderExtras{
			Review: &ReviewExtras{
				Model:       consolidated.Model,
				Summary:     consolidated.Summary,
				IsConsensus: len(panel) > 0,
				Issues:      consolidated.Issues,
				Strengths:   consolidated.Strengths,
			},
		},
	})
	return results
}

// MigrateToV2 upgrades a v1 (or v0) EvalReport to the v2 schema.
// It populates GraderResults from the existing Review/ReviewPanel fields
// and sets SchemaVersion. The function is idempotent.
//
// NOTE: As of v3, MigrateToV2 only ever lifts v0/v1 to v2. The v2 → v3
// jump is handled by MigrateToV3 (currently a no-op stub — see below).
func MigrateToV2(r *EvalReport) {
	if r.SchemaVersion >= 2 {
		return
	}
	if len(r.GraderResults) == 0 && r.Review != nil {
		r.GraderResults = GraderResultsFromReview(r.Review, r.ReviewPanel)
	}
	r.SchemaVersion = 2
}

// MigrateToV3 was the v2→v3 migration stub. As of v4, we enforce a hard cutover:
// v3 and earlier reports are rejected at load time.
//
// v4 introduced unified Points-first GraderResult; all graders emit Points;
// Extras is a discriminated union replacing per-kind detail fields. The v3→v4
// transformation is lossy in reverse (old detail fields cannot be reconstructed
// from v4's Extras), and the only first-party consumer (site/) is versioned
// alongside the engine. Old reports must be regenerated.
func MigrateToV3(r *EvalReport) {
	if r.SchemaVersion >= CurrentSchemaVersion {
		return
	}
	// v4 hard cutover: reject v < 4 with a clear error.
	// (Note: this function is named MigrateToV3 for historical reasons, but
	// CurrentSchemaVersion is 4 now. The name stuck from when v3 was current.)
	panic(fmt.Sprintf("report schema v%d is no longer supported; regenerate with: hyoka run ... (current schema: v%d)", r.SchemaVersion, CurrentSchemaVersion))
}
