package graders

import (
	"context"
	"fmt"
	"time"
)

// Grader evaluates generated code and produces a structured result.
// Each implementation handles a single concern (file check, build, LLM review, etc.).
type Grader interface {
	// Kind returns the grader kind (e.g., "file", "program", "prompt").
	Kind() string

	// Name returns the unique name of this grader instance.
	Name() string

	// Grade evaluates the generated code and returns a result.
	Grade(ctx context.Context, input *GraderInput) (*GraderResult, error)
}

// ActionEvent represents a single action taken by the agent during code generation.
type ActionEvent struct {
	Tool      string `json:"tool"`
	Arguments string `json:"arguments,omitempty"`
	Result    string `json:"result,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

// FileInfo describes a file produced by the agent.
type FileInfo struct {
	Path    string `json:"path"`
	Size    int64  `json:"size"`
	IsDir   bool   `json:"is_dir,omitempty"`
	ModTime int64  `json:"mod_time,omitempty"`
}

// SessionSummary holds metadata about the generation session.
type SessionSummary struct {
	SessionID    string        `json:"session_id"`
	Model        string        `json:"model"`
	TotalActions int           `json:"total_actions"`
	Duration     time.Duration `json:"duration"`
}

// GraderInput contains everything a grader needs to evaluate generated code.
// DM5: concrete struct rather than interface to keep grader signatures simple.
type GraderInput struct {
	WorkspacePath  string            // Path to session workspace with generated files
	ActionLog      []ActionEvent     // Agent action history
	PromptMeta     map[string]string // Prompt frontmatter properties (language, service, etc.)
	Config         map[string]any    // Grader-specific config values from YAML
	Files          []FileInfo        // Generated files listing
	SessionSummary *SessionSummary   // Optional session metadata
}

// GraderResult holds the outcome of a single grader execution.
// DM4: typed optional detail fields instead of interface{} for type safety.
type GraderResult struct {
	Kind    string  `json:"kind"`
	Name    string  `json:"name"`
	Score   float64 `json:"score"`   // 0.0–1.0 normalized
	Weight  float64 `json:"weight"`  // effective weight used in aggregation
	Pass    bool    `json:"pass"`    // whether the grader considers this passing
	Gate    bool    `json:"gate"`    // DM3: hard pass/fail overrides weighted scoring
	Message string  `json:"message"` // human-readable summary

	// Typed optional details per grader kind (DM4).
	FileDetails     *FileGraderDetails     `json:"file_details,omitempty"`
	ProgramDetails  *ProgramGraderDetails  `json:"program_details,omitempty"`
	PromptDetails   *PromptGraderDetails   `json:"prompt_details,omitempty"`
	BehaviorDetails *BehaviorGraderDetails `json:"behavior_details,omitempty"`
}

// FileGraderDetails contains results from file existence/content checks.
type FileGraderDetails struct {
	Path       string `json:"path"`
	Exists     bool   `json:"exists"`
	MatchFound bool   `json:"match_found,omitempty"` // true if pattern matched
	Pattern    string `json:"pattern,omitempty"`      // regex/glob that was tested
}

// ProgramGraderDetails contains results from running an external command.
type ProgramGraderDetails struct {
	Command  string        `json:"command"`
	ExitCode int           `json:"exit_code"`
	Stdout   string        `json:"stdout"`
	Stderr   string        `json:"stderr"`
	Duration time.Duration `json:"duration"`
}

// PromptGraderDetails contains results from an LLM-as-judge evaluation.
type PromptGraderDetails struct {
	Model     string `json:"model"`
	Rubric    string `json:"rubric"`
	Reasoning string `json:"reasoning"`
	RawScore  int    `json:"raw_score"`
	MaxScore  int    `json:"max_score"`
}

// BehaviorGraderDetails contains results from agent action log analysis.
type BehaviorGraderDetails struct {
	ToolsUsed      []string `json:"tools_used"`
	MissingTools   []string `json:"missing_tools,omitempty"`
	ForbiddenUsed  []string `json:"forbidden_used,omitempty"`
	TotalActions   int      `json:"total_actions"`
	TurnLimitHit   bool     `json:"turn_limit_hit,omitempty"`
	SequenceMatch  bool     `json:"sequence_match,omitempty"`   // for action_sequence
	ConstraintsMet bool     `json:"constraints_met,omitempty"`  // for tool_constraint
}

// AggregateResult holds the final aggregated score from all graders.
type AggregateResult struct {
	Score      float64        `json:"score"`       // 0.0–1.0 weighted average
	Pass       bool           `json:"pass"`        // true if score > 0 and no gate failures
	GateFailed bool           `json:"gate_failed"` // true if any gate grader failed
	Results    []GraderResult `json:"results"`
}

// AggregateResults computes a weighted score from a set of grader results.
// Gate semantics (DM3): if any gate grader fails, the overall result fails
// with a score of 0 regardless of other scores.
func AggregateResults(results []GraderResult) (*AggregateResult, error) {
	if len(results) == 0 {
		return nil, fmt.Errorf("no grader results to aggregate")
	}

	agg := &AggregateResult{
		Results: results,
		Pass:    true,
	}

	// Check gate failures first.
	for _, r := range results {
		if r.Gate && !r.Pass {
			agg.GateFailed = true
			agg.Pass = false
			agg.Score = 0
			return agg, nil
		}
	}

	// Compute weighted average.
	var totalWeight float64
	var weightedSum float64
	for _, r := range results {
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
