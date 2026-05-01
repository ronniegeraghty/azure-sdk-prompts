package graders

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/artifact"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
	"context"
	"fmt"
)

// WorkspaceDelta is a type alias for the workspace delta type.
// This allows graders to access workspace.WorkspaceDelta without creating an import cycle.
type WorkspaceDelta = workspace.WorkspaceDelta

// Grader is the core evaluation abstraction. Each grader is a single-concern
// evaluator (file check, build verification, LLM review, etc.) that scores
// one aspect of agent output.
type Grader interface {
// Kind returns the grader type identifier (e.g., "file", "program", "prompt").
Kind() string
// Name returns the human-readable name of this grader instance.
Name() string
// Grade evaluates the agent output and returns a scored result.
Grade(ctx context.Context, input GraderInput) (GraderResult, error)
}

// EnvironmentTool describes one tool/skill/MCP-server entry that was wired
// into the generation environment for this eval. Used by the tool_usage
// grader to decide which env-dependent rules are applicable.
type EnvironmentTool struct {
	Name string // MCP server name or skill name
	Kind string // "mcp" or "skill"
	Repo string // for repo-sourced skills
	Path string // for local skills (used to detect generator dir)
}

// GraderInput is a concrete struct containing everything a grader might need
// (DM5). Graders use what they need and ignore the rest.
type GraderInput struct {
	WorkspacePath  string          // Absolute path to the agent's output workspace
	ActionLog      []ActionEvent   // Ordered list of agent actions
	PromptMeta     PromptMetadata  // Metadata from the prompt frontmatter
	Config         GraderConfig    // The grader's own config entry
	Files          []FileEntry     // Listing of files in the workspace
	WorkspaceDelta *WorkspaceDelta // File-level changes made by agent (#566), may be nil

	// Optional fields for graders that need review/SDK access (WI-023).
	OriginalPrompt string // Original prompt text sent to the generator
	ReferenceDir   string // Directory containing reference answers
	EvalCriteria   string // Merged evaluation criteria text

	// AgentFinalResponse is the last assistant message from the generation
	// session. For prompts that evaluate reasoning or explanations rather than
	// files (e.g., planning tasks, recommendations), this is the primary
	// artifact under review. Empty if no assistant messages were logged.
	AgentFinalResponse string

	// EvalCriteriaBuckets carries per-bucket criteria when --review-mode
	// isolated produces multiple buckets. When non-empty and length > 1,
	// review-aware graders should run one Copilot session per bucket
	// instead of using the single EvalCriteria string. Length 0 or 1
	// means use the legacy single-session path (combined mode).
	EvalCriteriaBuckets []ReviewBucket

	// GeneratorArtifactPath is the absolute path to the generator.json file.
	GeneratorArtifactPath string

	// GeneratorArtifact is a pre-parsed pointer to the generator session artifact.
	GeneratorArtifact *GeneratorArtifact

	// EnvironmentTools lists every tool/skill/MCP-server entry available
	// to the generator session. Graders that score tool use compare this
	// against actual usage signals (SkillsInvoked, MCPServersUsed).
	EnvironmentTools []EnvironmentTool

	// SkillsInvoked is the set of skill names actually invoked during
	// generation (derived from skill.invoked events).
	SkillsInvoked []string

	// MCPServersUsed is the set of MCP server names that recorded at least
	// one tool call during generation.
	MCPServersUsed []string
}

// ReviewBucket is a graders-package mirror of criteria.ReviewBucket / review.Bucket.
// It carries already-rendered criteria text plus a stable name so the engine
// can pass bucket data through GraderInput without forcing graders to import
// criteria (which would create a layering issue) or review (cycle).
type ReviewBucket struct {
	Name     string
	Criteria string               // Legacy string-based criteria (deprecated)
	Checks   []review.ReviewCheck // Id-aware checks (new path)
}

// ActionEvent represents a single agent action from the session log.
type ActionEvent struct {
Tool       string `json:"tool"`
Action     string `json:"action"`
Path       string `json:"path,omitempty"`
TurnNumber int    `json:"turn_number,omitempty"`
}

// PromptMetadata holds prompt frontmatter fields relevant to grading.
type PromptMetadata struct {
ID       string `json:"id"`
Service  string `json:"service"`
Language string `json:"language"`
Plane    string `json:"plane"`
Category string `json:"category"`
}

// FileEntry describes a file in the agent workspace.
type FileEntry struct {
Path string `json:"path"` // Relative to workspace root
Size int64  `json:"size"`
}

// GraderResult is the single shape every grader returns. Pass and Score
// are derived from Checks at construction time — they are NOT independent
// signals. Any field outside Checks is render-only and may not influence
// pass/fail.
//
// This is the canonical report shape returned by all graders. All graders
// emit a GraderResult with:
//   - Kind: the grader type (one of KindXxx constants)
//   - Name: the grader instance name from YAML config
//   - Weight: the grader's contribution weight for aggregation
//   - SourceFile: the criteria file that defined this grader
//   - SourceType: the original grader type from config (before any translation)
//   - Checks: the list of individual pass/fail checks evaluated by this grader
//   - Extras: kind-specific render-only data for display purposes
//
// Graders use the Extras field for kind-specific render-only data that does
// not influence scoring. The canonical grader types are:
//   - prompt: LLM-based review with a judge model
//   - output_check: workspace file/byte checks
//   - program: external program execution
//   - tool: unified tool-perspective checks (replaces behavior/tool_constraint/tool_usage)
//   - action_sequence: ordered action verification
//
// Deprecated grader types (emit warnings at load time):
//   - behavior → use 'tool' instead
//   - tool_constraint → use 'tool' instead
//   - tool_usage → use 'tool' instead
//   - file → use 'output_check' with require_files instead
type GraderResult struct {
	Kind    string  `json:"kind"`              // one of KindXxx
	Name    string  `json:"name"`              // YAML instance name
	Weight  float64 `json:"weight"`            // aggregation weight (from config)
	
	// Deprecated: Gate semantics are no longer enforced. All graders run and contribute
	// their score to the weighted aggregate. Use the consolidated 'tool' grader's check
	// kinds or separate explicit graders to express pass/fail requirements instead.
	Gate    bool    `json:"gate"`              // gate flag (from config)

	// Derived from Checks — see NewResult helper.
	Score   float64 `json:"score"`             // sum(check.Weight * pass) / sum(check.Weight); 0 if no checks
	Pass    bool    `json:"pass"`              // AND over Checks[i].Pass
	Message string  `json:"message"`           // headline summary (≤ ~120 chars)

	Checks  []GraderCheck  `json:"points"`           // REQUIRED, len ≥ 1; the canonical sub-checks (JSON tag unchanged for back-compat; C10 will dual-emit)
	Extras  *GraderExtras  `json:"extras,omitempty"` // kind-specific render-only payload

	// Provenance: where the grader entry was declared. Populated by the
	// engine after Grade() returns. SourceType is one of "prompt_file" or
	// "criteria_file"; SourceFile is the originating file path.
	SourceFile string `json:"source_file,omitempty"`
	SourceType string `json:"source_type,omitempty"`
}

// GraderCheck represents one binary pass/fail check inside a grader.
// A grader's overall Pass is the AND of every Check.Pass.
type GraderCheck struct {
	Label    string             `json:"label"`              // short, what was checked (e.g. "file present: src/main.py")
	Pass     bool               `json:"pass"`
	Message  string             `json:"message,omitempty"`  // why it passed/failed (the "reason" Ronnie asked for)
	Weight   float64            `json:"weight,omitempty"`   // for Score weighting; defaults to 1.0 when 0/omitted
	Evidence map[string]string  `json:"evidence,omitempty"` // tiny, optional, string-only KV (e.g. {"pattern":"^def "})
}

// Deprecated: use GraderCheck. Will be removed next release.
type GraderPoint = GraderCheck

// GraderExtras is a discriminated union carrying kind-specific render-only data.
// Only one field should be populated per result.
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
	Command    string `json:"command"`
	Args       []string `json:"args,omitempty"`
	ExitCode   int    `json:"exit_code"`
	Stdout     string `json:"stdout,omitempty"`
	Stderr     string `json:"stderr,omitempty"`
	DurationMs int64  `json:"duration_ms,omitempty"`
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
	Score     int                      `json:"score"`      // points earned
	Pass      bool                     `json:"pass"`       // all criteria passed
	Issues    []string                 `json:"issues,omitempty"`
	Strengths []string                 `json:"strengths,omitempty"`
	Criteria  []ReviewCriterionResult  `json:"criteria"`
}

// ReviewCriterionResult holds one criterion's outcome.
type ReviewCriterionResult struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Reason string `json:"reason,omitempty"`
	Weight int    `json:"weight,omitempty"` // max points for this criterion
}

// NewResult constructs a GraderResult with Pass and Score derived from Checks.
// Checks must be non-empty; panics if empty. Each check's Weight defaults to 1.0
// when zero or omitted.
//
// Invariant (Phase 3, 2026-04-25): every grader MUST emit at least one Check.
// A Checks-less GraderResult is a bug — the site's defensive "PASS"/"100%"
// fallback exists only as a safety net for legacy data on disk, never for
// freshly-generated results. Use NewErrorResult below when an internal error
// prevents a real check.
func NewResult(kind, name string, cfg GraderConfig, checks []GraderCheck, msg string, extras *GraderExtras) GraderResult {
	if len(checks) == 0 {
		panic(fmt.Sprintf("grader %s (%s) must emit at least one Check", name, kind))
	}

	// Derive Pass: AND over all Checks
	pass := true
	for _, c := range checks {
		if !c.Pass {
			pass = false
			break
		}
	}

	// Derive Score: weighted sum / total weight
	var weightedSum, totalWeight float64
	for _, c := range checks {
		w := c.Weight
		if w == 0 {
			w = 1.0
		}
		totalWeight += w
		if c.Pass {
			weightedSum += w
		}
	}
	score := 0.0
	if totalWeight > 0 {
		score = weightedSum / totalWeight
	}

	return GraderResult{
		Kind:    kind,
		Name:    name,
		Weight:  cfg.EffectiveWeight(),
		Gate:    cfg.Gate,
		Score:   score,
		Pass:    pass,
		Message: msg,
		Checks:  checks,
		Extras:  extras,
	}
}

// NewErrorResult constructs a GraderResult representing a grader that failed
// to execute (returned an error, panicked, or was skipped). It synthesizes a
// single failing "grader execution" Check so the result still satisfies the
// "every grader must emit ≥ 1 Check" invariant. Without this, the site's
// renderer can't distinguish a graderless failure from a passing grader and
// falls back to "PASS"/"100%".
//
// Use this everywhere a GraderResult is constructed outside a real Grade()
// path — currently the engine's error and panic recovery branches.
func NewErrorResult(kind, name string, cfg GraderConfig, msg string) GraderResult {
	if msg == "" {
		msg = "grader execution failed"
	}
	checks := []GraderCheck{{
		Label:   "grader executed",
		Pass:    false,
		Message: msg,
	}}
	return NewResult(kind, name, cfg, checks, msg, nil)
}

// AggregateResult holds the final aggregated score from all graders.
type AggregateResult struct {
Score      float64        `json:"score"`       // 0.0–1.0 weighted average
Pass       bool           `json:"pass"`        // true if score > 0 and no gate failures
GateFailed bool           `json:"gate_failed"` // true if any gate grader failed
Results    []GraderResult `json:"results"`
}

// AggregateResults computes a weighted score from a set of grader results.
//
// Phase 2 cutover (#625): the gate short-circuit was removed. Every grader
// runs regardless of other graders' pass/fail and every result contributes
// to the weighted score. Aggregate Pass is now "all results passed" so the
// engine can still flag failed evals in reports without special-casing
// gate graders. The GateFailed field is preserved for back-compat with
// existing report schemas but is never set to true by this function.
func AggregateResults(results []GraderResult) (*AggregateResult, error) {
	if len(results) == 0 {
		return nil, fmt.Errorf("no grader results to aggregate")
	}

	agg := &AggregateResult{
		Results: results,
		Pass:    true,
	}

	// Compute weighted average; Pass is AND of every result's Pass.
	var totalWeight float64
	var weightedSum float64
	for _, r := range results {
		if !r.Pass {
			agg.Pass = false
		}
		w := r.Weight
		if w == 0 {
			w = 1.0
		}
		totalWeight += w
		weightedSum += r.Score * w
	}

	if totalWeight > 0 {
		agg.Score = weightedSum / totalWeight
	}

	return agg, nil
}

// GeneratorArtifact is a type alias for the generator artifact type.
type GeneratorArtifact = artifact.GeneratorArtifact
