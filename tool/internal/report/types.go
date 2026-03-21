package report

import (
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/build"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/review"
)

// EvalReport contains the results of a single prompt evaluation.
type EvalReport struct {
PromptID       string               `json:"prompt_id"`
ConfigName     string               `json:"config_name"`
Timestamp      string               `json:"timestamp"`
Duration       float64              `json:"duration_seconds"`
PromptMeta     map[string]any       `json:"prompt_metadata"`
ConfigUsed     map[string]any       `json:"config_used"`
GeneratedFiles []string             `json:"generated_files"`
Build          *build.BuildResult   `json:"build"`
Review         *review.ReviewResult `json:"review,omitempty"`
EventCount     int                  `json:"event_count"`
ToolCalls      []string             `json:"tool_calls"`
Success        bool                 `json:"success"`
Error          string               `json:"error,omitempty"`
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
Duration     float64       `json:"duration_seconds"`
Reports      []string      `json:"report_paths"`
Results      []*EvalReport `json:"results,omitempty"`
}
