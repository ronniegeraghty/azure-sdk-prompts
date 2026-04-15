// Package report handles generation of JSON and Markdown reports.
package report

import (
	"fmt"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/pairwise"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

// CurrentSchemaVersion is the latest report schema version.
// v1 = legacy monolithic ReviewResult; v2 = grader-based.
const CurrentSchemaVersion = 2

// GraderResult holds the output from a single grader (LLM reviewer, build check, etc.).
type GraderResult struct {
	GraderName   string              `json:"grader_name"`
	GraderType   string              `json:"grader_type"` // "review", "file", "program", "prompt", "behavior", etc.
	Model        string              `json:"model,omitempty"`
	Scores       review.ReviewScores `json:"scores"`
	OverallScore int                 `json:"overall_score"`
	MaxScore     int                 `json:"max_score"`
	Summary      string              `json:"summary"`
	Issues       []string            `json:"issues,omitempty"`
	Strengths    []string            `json:"strengths,omitempty"`
	Duration     float64             `json:"duration_seconds,omitempty"`
	IsConsensus  bool                `json:"is_consensus,omitempty"`

	// Grader-system fields (populated for pluggable graders).
	Score  float64 `json:"score,omitempty"`  // 0.0–1.0 normalized
	Weight float64 `json:"weight,omitempty"` // Weight for aggregation
	Pass   *bool   `json:"pass,omitempty"`   // nil for legacy review-type graders
	Gate   bool    `json:"gate,omitempty"`   // Gate grader flag

	// Typed details — only one populated per result.
	FileDetails     *FileGraderDetail     `json:"file_details,omitempty"`
	ProgramDetails  *ProgramGraderDetail  `json:"program_details,omitempty"`
	PromptDetails   *PromptGraderDetail   `json:"prompt_details,omitempty"`
	BehaviorDetails *BehaviorGraderDetail `json:"behavior_details,omitempty"`
}

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
	AvailableTools    []string `json:"availableTools,omitempty"`
	ExcludedTools     []string `json:"excludedTools,omitempty"`
	MCPServers        []string `json:"mcpServers,omitempty"`
	MCPToolsInvoked   []string `json:"mcpToolsInvoked,omitempty"`
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
	ReviewedFiles      []ReviewedFile        `json:"reviewed_files,omitempty"`
	Review             *review.ReviewResult  `json:"review,omitempty"`
	ReviewPanel        []review.ReviewResult `json:"review_panel,omitempty"`
	GraderResults      []GraderResult        `json:"grader_results,omitempty"`
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
	Success            bool                  `json:"success"`
	Error              string                `json:"error,omitempty"`
	ErrorDetails       string                `json:"error_details,omitempty"`
	ErrorCategory      string                `json:"error_category,omitempty"` // timeout, sdk_error, generation_failure, review_failure, no_files
	FailureReason      string                `json:"failure_reason,omitempty"` // human-readable explanation of failure
	IsStub             bool                  `json:"is_stub,omitempty"`
	RerunCommand       string                `json:"rerunCommand,omitempty"`
	// Generator guardrails (#35)
	GuardrailMaxTurns          int    `json:"guardrail_max_turns,omitempty"`
	GuardrailMaxFiles          int    `json:"guardrail_max_files,omitempty"`
	GuardrailMaxOutputSize     int64  `json:"guardrail_max_output_size,omitempty"`
	GuardrailMaxSessionActions int    `json:"guardrail_max_session_actions,omitempty"`
	GuardrailAbortReason       string `json:"guardrail_abort_reason,omitempty"`
	// Action limit soft cap — generation stopped but review proceeds with partial results
	ActionLimitReached bool `json:"action_limit_reached,omitempty"`
	ActionCount        int  `json:"action_count,omitempty"`
}

// ToolLoadResult records the outcome of loading a single tool, skill, or MCP server.
type ToolLoadResult struct {
	Name    string `json:"name"`
	Status  string `json:"status"` // "loaded", "failed", "configured"
	Error   string `json:"error,omitempty"`
	Details string `json:"details,omitempty"` // e.g., command string for MCP servers
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
		if r.Gate && (r.Pass == nil || !*r.Pass) {
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
			Pass:          r.Pass != nil && *r.Pass,
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
