package graders

import (
"context"
"fmt"
)

// Grader is the core evaluation abstraction. Each grader is a single-concern
// evaluator (file check, build verification, LLM review, etc.) that scores
// one aspect of agent-generated code.
type Grader interface {
// Kind returns the grader type identifier (e.g., "file", "program", "prompt").
Kind() string
// Name returns the human-readable name of this grader instance.
Name() string
// Grade evaluates the agent output and returns a scored result.
Grade(ctx context.Context, input GraderInput) (GraderResult, error)
}

// GraderInput is a concrete struct containing everything a grader might need
// (DM5). Graders use what they need and ignore the rest.
type GraderInput struct {
WorkspacePath string         // Absolute path to the agent's output workspace
ActionLog     []ActionEvent  // Ordered list of agent actions
PromptMeta    PromptMetadata // Metadata from the prompt frontmatter
Config        GraderConfig   // The grader's own config entry
Files         []FileEntry    // Listing of files in the workspace
}

// ActionEvent represents a single agent action from the session log.
type ActionEvent struct {
Tool   string `json:"tool"`
Action string `json:"action"`
Path   string `json:"path,omitempty"`
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

// GraderResult uses typed optional fields instead of interface{} (DM4).
// Templates check `if .FileDetails` directly — no type assertions needed.
type GraderResult struct {
Kind    string  `json:"kind"`
Name    string  `json:"name"`
Score   float64 `json:"score"`   // 0.0–1.0 normalized
Weight  float64 `json:"weight"`  // Weight for aggregation
Pass    bool    `json:"pass"`    // Binary pass/fail
Gate    bool    `json:"gate"`    // If true, failure overrides weighted scoring (DM3)
Message string  `json:"message"` // Human-readable summary

// Typed details — only one populated per result (DM4).
FileDetails     *FileGraderDetails     `json:"file_details,omitempty"`
ProgramDetails  *ProgramGraderDetails  `json:"program_details,omitempty"`
PromptDetails   *PromptGraderDetails   `json:"prompt_details,omitempty"`
BehaviorDetails *BehaviorGraderDetails `json:"behavior_details,omitempty"`
}

// FileGraderDetails holds file-check specifics.
type FileGraderDetails struct {
CheckedFiles []FileCheckResult `json:"checked_files"`
}

// FileCheckResult records the outcome of a single file check.
type FileCheckResult struct {
Path           string `json:"path"`
Exists         bool   `json:"exists"`
PatternMatched *bool  `json:"pattern_matched,omitempty"` // nil if no pattern configured
Pattern        string `json:"pattern,omitempty"`
}

// ProgramGraderDetails holds program execution specifics.
type ProgramGraderDetails struct {
Command  string `json:"command"`
ExitCode int    `json:"exit_code"`
Stdout   string `json:"stdout"`
Stderr   string `json:"stderr"`
}

// PromptGraderDetails holds LLM-as-judge specifics.
type PromptGraderDetails struct {
Model     string `json:"model"`
Rubric    string `json:"rubric"`
Reasoning string `json:"reasoning"`
}

// BehaviorGraderDetails holds agent behavior analysis specifics.
type BehaviorGraderDetails struct {
ToolsUsed     []string `json:"tools_used,omitempty"`
ForbiddenUsed []string `json:"forbidden_used,omitempty"`
TurnCount     int      `json:"turn_count"`
Violations    []string `json:"violations,omitempty"`
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
