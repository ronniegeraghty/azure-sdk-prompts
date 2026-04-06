// Package report handles generation of JSON, HTML, and Markdown reports.
package report

import (
	"github.com/ronniegeraghty/hyoka/internal/pairwise"
	"github.com/ronniegeraghty/hyoka/internal/review"
)

// CurrentSchemaVersion is the latest report schema version.
// v1 = legacy monolithic ReviewResult; v2 = grader-based.
const CurrentSchemaVersion = 2

// GraderResultEntry is a serializable representation of a single grader result.
type GraderResultEntry struct {
	Name    string  `json:"name"`
	Kind    string  `json:"kind"`
	Passed  bool    `json:"passed"`
	Score   float64 `json:"score"`
	Message string  `json:"message"`
	Weight  float64 `json:"weight"`
	Gate    bool    `json:"gate,omitempty"`
	Error   string  `json:"error,omitempty"`
}

// GraderAggregateEntry is a serializable representation of aggregated grader results.
type GraderAggregateEntry struct {
	Results     []GraderResultEntry `json:"results"`
	Score       float64             `json:"score"`
	Passed      bool                `json:"passed"`
	GatesFailed []string            `json:"gates_failed,omitempty"`
}

// GraderResult holds the output from a single grader (LLM reviewer, build check, etc.).
type GraderResult struct {
	GraderName   string              `json:"grader_name"`
	GraderType   string              `json:"grader_type"` // "review", "build", "lint", "test"
	Model        string              `json:"model,omitempty"`
	Scores       review.ReviewScores `json:"scores"`
	OverallScore int                 `json:"overall_score"`
	MaxScore     int                 `json:"max_score"`
	Summary      string              `json:"summary"`
	Issues       []string            `json:"issues,omitempty"`
	Strengths    []string            `json:"strengths,omitempty"`
	Duration     float64             `json:"duration_seconds,omitempty"`
	IsConsensus  bool                `json:"is_consensus,omitempty"`
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
	AvailableTools    []string `json:"availableTools,omitempty"`
	ExcludedTools     []string `json:"excludedTools,omitempty"`
	MCPServers        []string `json:"mcpServers,omitempty"`
	SafetyBoundaries  bool     `json:"safetyBoundaries"`
	AllowCloud        bool     `json:"allowCloud"`
	WorkingDirectory  string   `json:"workingDirectory"`
	TotalInputTokens  int      `json:"totalInputTokens,omitempty"`
	TotalOutputTokens int      `json:"totalOutputTokens,omitempty"`
	TurnCount         int      `json:"turnCount,omitempty"`
	ContextTruncated  bool     `json:"contextTruncated,omitempty"`
}

// ResourceStats holds per-eval peak resource utilization (#45).
type ResourceStats struct {
	PeakCPUPercent float64 `json:"peak_cpu_percent"`
	PeakMemoryMB   float64 `json:"peak_memory_mb"`
	SampleCount    int     `json:"sample_count"`
}

// EvalReport contains the results of a single prompt evaluation.
type EvalReport struct {
	SchemaVersion  int                   `json:"schema_version"`
	PromptID       string                `json:"prompt_id"`
	ConfigName     string                `json:"config_name"`
	Timestamp      string                `json:"timestamp"`
	Duration               float64               `json:"duration_seconds"`
	GenerationDuration     float64               `json:"generation_duration_seconds,omitempty"`
	ReviewDuration         float64               `json:"review_duration_seconds,omitempty"`
	PromptMeta     map[string]any        `json:"prompt_metadata"`
	ConfigUsed     map[string]any        `json:"config_used"`
	GeneratedFiles []string              `json:"generated_files"`
	StarterFiles   []string              `json:"starter_files,omitempty"`
	ReviewedFiles  []ReviewedFile        `json:"reviewed_files,omitempty"`
	Review         *review.ReviewResult  `json:"review,omitempty"`
	ReviewPanel    []review.ReviewResult `json:"review_panel,omitempty"`
	GraderResults  []GraderResult        `json:"grader_results_legacy,omitempty"`
	GraderAgg      *GraderAggregateEntry `json:"grader_results,omitempty"` // Pluggable grader results (#136)
	ToolUsage      *ToolUsageResult      `json:"tool_usage,omitempty"`
	SessionEvents  []SessionEventRecord  `json:"session_events,omitempty"`
	EventCount     int                   `json:"event_count"`
	ToolCalls      []string              `json:"tool_calls"`
	Environment    *EnvironmentInfo      `json:"environment,omitempty"`
	ResourceUsage  *ResourceStats        `json:"resource_usage,omitempty"` // Per-eval resource stats (#45)
	Success        bool                  `json:"success"`
	Error          string                `json:"error,omitempty"`
	ErrorDetails   string                `json:"error_details,omitempty"`
	ErrorCategory  string                `json:"error_category,omitempty"` // timeout, sdk_error, generation_failure, review_failure, no_files
	FailureReason  string                `json:"failure_reason,omitempty"` // human-readable explanation of failure
	IsStub         bool                  `json:"is_stub,omitempty"`
	RerunCommand   string                `json:"rerunCommand,omitempty"`
	// Generator guardrails (#35)
	GuardrailMaxTurns          int    `json:"guardrail_max_turns,omitempty"`
	GuardrailMaxFiles          int    `json:"guardrail_max_files,omitempty"`
	GuardrailMaxOutputSize     int64  `json:"guardrail_max_output_size,omitempty"`
	GuardrailMaxSessionActions int    `json:"guardrail_max_session_actions,omitempty"`
	GuardrailAbortReason       string `json:"guardrail_abort_reason,omitempty"`
}

// RunResourceStats holds aggregate resource utilization across all evals (#45).
type RunResourceStats struct {
	PeakCPUPercent float64 `json:"peak_cpu_percent"`
	PeakMemoryMB   float64 `json:"peak_memory_mb"`
	SessionCount   int     `json:"session_count"`
}

// RunSummary contains aggregate statistics for an evaluation run.
type RunSummary struct {
	RunID        string        `json:"run_id"`
	Timestamp    string        `json:"timestamp"`
	TotalPrompts int           `json:"total_prompts"`
	TotalConfigs int           `json:"total_configs"`
	TotalEvals   int           `json:"total_evaluations"`
	Passed       int           `json:"passed"`
	Failed       int           `json:"failed"`
	Errors       int           `json:"errors"`
	Duration              float64       `json:"duration_seconds"`
	AvgGenerationDuration float64       `json:"avg_generation_duration_seconds,omitempty"`
	AvgReviewDuration     float64       `json:"avg_review_duration_seconds,omitempty"`
	Reports      []string      `json:"report_paths"`
	Results      []*EvalReport `json:"results,omitempty"`
	Analysis     string        `json:"analysis,omitempty"`
	ResourceUsage  *RunResourceStats `json:"resource_usage,omitempty"` // Aggregate resource stats (#45)
	PairwiseResults []*pairwise.PairwiseReport `json:"pairwise_results,omitempty"` // Per-prompt pairwise impact (#123)
}

// GraderResultsFromReview converts legacy ReviewResult data into []GraderResult.
// If panel is non-empty, each panel member becomes a grader entry and the
// consolidated review becomes the consensus entry.
func GraderResultsFromReview(consolidated *review.ReviewResult, panel []review.ReviewResult) []GraderResult {
	if consolidated == nil {
		return nil
	}
	var results []GraderResult
	for i := range panel {
		r := &panel[i]
		name := r.Model
		if name == "" {
			name = "reviewer"
		}
		results = append(results, GraderResult{
			GraderName:   name,
			GraderType:   "review",
			Model:        r.Model,
			Scores:       r.Scores,
			OverallScore: r.OverallScore,
			MaxScore:     r.MaxScore,
			Summary:      r.Summary,
			Issues:       r.Issues,
			Strengths:    r.Strengths,
		})
	}
	consensusName := consolidated.Model
	if consensusName == "" {
		consensusName = "consensus"
	}
	results = append(results, GraderResult{
		GraderName:   consensusName,
		GraderType:   "review",
		Model:        consolidated.Model,
		Scores:       consolidated.Scores,
		OverallScore: consolidated.OverallScore,
		MaxScore:     consolidated.MaxScore,
		Summary:      consolidated.Summary,
		Issues:       consolidated.Issues,
		Strengths:    consolidated.Strengths,
		IsConsensus:  len(panel) > 0,
	})
	return results
}

// MigrateToV2 upgrades a v1 (or v0) EvalReport to the current v2 schema.
// It populates GraderResults from the existing Review/ReviewPanel fields
// and sets SchemaVersion. The function is idempotent.
func MigrateToV2(r *EvalReport) {
	if r.SchemaVersion >= CurrentSchemaVersion {
		return
	}
	if len(r.GraderResults) == 0 && r.Review != nil {
		r.GraderResults = GraderResultsFromReview(r.Review, r.ReviewPanel)
	}
	r.SchemaVersion = CurrentSchemaVersion
}
