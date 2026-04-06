package eval

import (
"context"
"encoding/json"
"os"
"path/filepath"
"testing"

"github.com/ronniegeraghty/hyoka/internal/config"
"github.com/ronniegeraghty/hyoka/internal/prompt"
"github.com/ronniegeraghty/hyoka/internal/report"
"github.com/ronniegeraghty/hyoka/internal/review"
)

// TestIntegrationStubEvalReviewPipeline runs the full engine pipeline with
// StubEvaluator + StubReviewer and verifies that a valid report.json and
// summary.json are written to disk.
func TestIntegrationStubEvalReviewPipeline(t *testing.T) {
outputDir := t.TempDir()
engine := NewEngineWithReviewer(&StubEvaluator{}, &review.StubReviewer{}, quietOpts(EngineOptions{
Workers:   1,
OutputDir: outputDir,
}))

prompts := []*prompt.Prompt{
{
ID:                 "integ-test-prompt",
PromptText:         "Create a sample Azure SDK client",
EvaluationCriteria: "Must compile and use DefaultAzureCredential",
Properties: map[string]string{
"service":  "identity",
"plane":    "data-plane",
"language": "python",
"category": "auth",
},
},
}
configs := []config.ToolConfig{
{
Name:      "test-config",
Generator: &config.GeneratorConfig{Model: "gpt-4"},
Reviewer:  &config.ReviewerConfig{Model: "gpt-4"},
},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("engine.Run() returned error: %v", err)
}

// --- Verify RunSummary fields ---
if summary.RunID == "" {
t.Error("expected non-empty RunID")
}
if summary.TotalEvals != 1 {
t.Errorf("expected 1 eval, got %d", summary.TotalEvals)
}
if summary.TotalPrompts != 1 {
t.Errorf("expected 1 prompt, got %d", summary.TotalPrompts)
}
if summary.TotalConfigs != 1 {
t.Errorf("expected 1 config, got %d", summary.TotalConfigs)
}
if len(summary.Results) != 1 {
t.Fatalf("expected 1 result, got %d", len(summary.Results))
}

r := summary.Results[0]
if r.PromptID != "integ-test-prompt" {
t.Errorf("expected PromptID 'integ-test-prompt', got %q", r.PromptID)
}
if r.ConfigName != "test-config" {
t.Errorf("expected ConfigName 'test-config', got %q", r.ConfigName)
}
if !r.IsStub {
t.Error("expected IsStub=true")
}
if r.Timestamp == "" {
t.Error("expected non-empty Timestamp")
}
if r.Duration <= 0 {
t.Errorf("expected positive Duration, got %f", r.Duration)
}
if r.GenerationDuration <= 0 {
t.Errorf("expected positive GenerationDuration, got %f", r.GenerationDuration)
}

// StubReviewer should populate review data
if r.Review == nil && len(r.ReviewPanel) == 0 {
t.Error("expected review data to be populated by StubReviewer")
}

// --- Verify report.json on disk ---
reportDir := report.ReportDir(outputDir, summary.RunID, prompts[0])
reportPath := filepath.Join(reportDir, "test-config", "report.json")
reportData, err := os.ReadFile(reportPath)
if err != nil {
t.Fatalf("expected report.json at %s: %v", reportPath, err)
}

var diskReport report.EvalReport
if err := json.Unmarshal(reportData, &diskReport); err != nil {
t.Fatalf("report.json is not valid JSON: %v", err)
}
if diskReport.PromptID != "integ-test-prompt" {
t.Errorf("disk report PromptID = %q, want 'integ-test-prompt'", diskReport.PromptID)
}
if diskReport.ConfigName != "test-config" {
t.Errorf("disk report ConfigName = %q, want 'test-config'", diskReport.ConfigName)
}
if !diskReport.IsStub {
t.Error("disk report: expected IsStub=true")
}
if diskReport.SchemaVersion < 1 {
t.Errorf("expected SchemaVersion >= 1, got %d", diskReport.SchemaVersion)
}

// --- Verify summary.json on disk ---
summaryPath := filepath.Join(outputDir, summary.RunID, "summary.json")
summaryData, err := os.ReadFile(summaryPath)
if err != nil {
t.Fatalf("expected summary.json at %s: %v", summaryPath, err)
}

var diskSummary report.RunSummary
if err := json.Unmarshal(summaryData, &diskSummary); err != nil {
t.Fatalf("summary.json is not valid JSON: %v", err)
}
if diskSummary.RunID != summary.RunID {
t.Errorf("summary RunID mismatch: disk=%q, memory=%q", diskSummary.RunID, summary.RunID)
}
}

// TestIntegrationMultiPromptMultiConfigWithReview verifies that the full pipeline
// fans out correctly across multiple prompts x configs and that each combination
// produces a valid report on disk.
func TestIntegrationMultiPromptMultiConfigWithReview(t *testing.T) {
outputDir := t.TempDir()
engine := NewEngineWithReviewer(&StubEvaluator{}, &review.StubReviewer{}, quietOpts(EngineOptions{
Workers:   2,
OutputDir: outputDir,
}))

prompts := []*prompt.Prompt{
{
ID:         "mp-prompt-a",
PromptText: "Create identity client",
Properties: map[string]string{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
},
{
ID:         "mp-prompt-b",
PromptText: "Create storage blob client",
Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
},
}
configs := []config.ToolConfig{
{Name: "cfg-alpha", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Reviewer: &config.ReviewerConfig{Model: "gpt-4"}},
{Name: "cfg-beta", Generator: &config.GeneratorConfig{Model: "claude-sonnet-4.5"}, Reviewer: &config.ReviewerConfig{Model: "claude-sonnet-4.5"}},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("engine.Run() returned error: %v", err)
}

if summary.TotalEvals != 4 {
t.Errorf("expected 4 evals (2x2), got %d", summary.TotalEvals)
}
if len(summary.Results) != 4 {
t.Fatalf("expected 4 results, got %d", len(summary.Results))
}

// Verify each result has a report.json on disk
for _, r := range summary.Results {
var p *prompt.Prompt
for _, pp := range prompts {
if pp.ID == r.PromptID {
p = pp
break
}
}
if p == nil {
t.Fatalf("result has unknown PromptID %q", r.PromptID)
}

reportPath := filepath.Join(report.ReportDir(outputDir, summary.RunID, p), r.ConfigName, "report.json")
data, err := os.ReadFile(reportPath)
if err != nil {
t.Errorf("missing report.json for %s/%s: %v", r.PromptID, r.ConfigName, err)
continue
}

var diskReport report.EvalReport
if err := json.Unmarshal(data, &diskReport); err != nil {
t.Errorf("invalid JSON for %s/%s: %v", r.PromptID, r.ConfigName, err)
continue
}
if diskReport.PromptID != r.PromptID {
t.Errorf("disk PromptID=%q, expected %q", diskReport.PromptID, r.PromptID)
}
if diskReport.ConfigName != r.ConfigName {
t.Errorf("disk ConfigName=%q, expected %q", diskReport.ConfigName, r.ConfigName)
}
}

// Verify summary.json
summaryPath := filepath.Join(outputDir, summary.RunID, "summary.json")
if _, err := os.Stat(summaryPath); err != nil {
t.Errorf("expected summary.json at %s: %v", summaryPath, err)
}
}

// TestIntegrationReviewerFactory verifies the factory-based reviewer path
// creates separate reviewer instances per config and generates correct reports.
func TestIntegrationReviewerFactory(t *testing.T) {
outputDir := t.TempDir()
factoryCalls := 0
factory := func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
factoryCalls++
return &review.StubReviewer{}, nil, nil
}

engine := NewEngineWithReviewerFactory(&StubEvaluator{}, factory, quietOpts(EngineOptions{
Workers:   1,
OutputDir: outputDir,
}))

prompts := []*prompt.Prompt{
{
ID:         "factory-test",
PromptText: "test",
Properties: map[string]string{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
},
}
configs := []config.ToolConfig{
{Name: "cfg-1", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Reviewer: &config.ReviewerConfig{Model: "gpt-4"}},
{Name: "cfg-2", Generator: &config.GeneratorConfig{Model: "claude-sonnet-4.5"}, Reviewer: &config.ReviewerConfig{Model: "claude-sonnet-4.5"}},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("engine.Run() returned error: %v", err)
}

if factoryCalls != 2 {
t.Errorf("expected factory called 2 times, got %d", factoryCalls)
}
if summary.TotalEvals != 2 {
t.Errorf("expected 2 evals, got %d", summary.TotalEvals)
}
}

// TestIntegrationSkipReviewStillWritesReport verifies that with SkipReview=true,
// the engine still writes valid report.json and summary.json.
func TestIntegrationSkipReviewStillWritesReport(t *testing.T) {
outputDir := t.TempDir()
engine := NewEngine(&StubEvaluator{}, quietOpts(EngineOptions{
Workers:    1,
OutputDir:  outputDir,
SkipReview: true,
}))

prompts := []*prompt.Prompt{
{
ID:         "skip-review-integ",
PromptText: "test",
Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
},
}
configs := []config.ToolConfig{
{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("engine.Run() returned error: %v", err)
}

// Verify report.json
reportPath := filepath.Join(report.ReportDir(outputDir, summary.RunID, prompts[0]), "test-config", "report.json")
data, err := os.ReadFile(reportPath)
if err != nil {
t.Fatalf("expected report.json: %v", err)
}

var diskReport report.EvalReport
if err := json.Unmarshal(data, &diskReport); err != nil {
t.Fatalf("invalid JSON: %v", err)
}
if !diskReport.IsStub {
t.Error("expected IsStub=true in report")
}

// Verify summary.json
summaryPath := filepath.Join(outputDir, summary.RunID, "summary.json")
summaryData, err := os.ReadFile(summaryPath)
if err != nil {
t.Fatalf("expected summary.json: %v", err)
}
var diskSummary report.RunSummary
if err := json.Unmarshal(summaryData, &diskSummary); err != nil {
t.Fatalf("invalid summary JSON: %v", err)
}
}
