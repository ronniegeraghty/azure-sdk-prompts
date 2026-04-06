package report

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/prompt"
)

func TestWriteReport(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:       "test-prompt",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       12.5,
		PromptMeta:     map[string]any{"service": "storage", "language": "dotnet"},
		ConfigUsed:     map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles: []string{"Program.cs", "Storage.csproj"},
		EventCount: 15,
		ToolCalls:  []string{"create_file", "edit_file"},
		Success:    true,
	}

	p := &prompt.Prompt{
		ID:         "test-prompt",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "dotnet", "category": "authentication"},
	}








	reportPath, err := WriteReport(r, dir, "20240115-100000", p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(reportPath); err != nil {
		t.Fatalf("report file does not exist: %v", err)
	}

	// Verify JSON is valid
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON in report: %v", err)
	}

	if parsed.PromptID != "test-prompt" {
		t.Errorf("expected prompt ID 'test-prompt', got %q", parsed.PromptID)
	}
	if !parsed.Success {
		t.Error("expected success to be true")
	}

	// Verify directory structure — includes prompt ID for workspace isolation
	expectedDir := filepath.Join(dir, "20240115-100000", "results", "storage", "data-plane", "dotnet", "authentication", "test-prompt", "baseline")
	if _, err := os.Stat(expectedDir); err != nil {
		t.Errorf("expected directory %s to exist", expectedDir)
	}
}

func TestWriteSummary(t *testing.T) {
	dir := t.TempDir()

	s := &RunSummary{
		RunID:        "20240115-100000",
		Timestamp:    "2024-01-15T10:00:00Z",
		TotalPrompts: 5,
		TotalConfigs: 2,
		TotalEvals:   10,
		Passed:       8,
		Failed:       1,
		Errors:       1,
		Duration:     120.5,
		Reports:      []string{"/path/to/report1.json", "/path/to/report2.json"},
	}

	summaryPath, err := WriteSummary(s, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	var parsed RunSummary
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON in summary: %v", err)
	}

	if parsed.TotalEvals != 10 {
		t.Errorf("expected 10 total evals, got %d", parsed.TotalEvals)
	}
	if parsed.Passed != 8 {
		t.Errorf("expected 8 passed, got %d", parsed.Passed)
	}
}

func TestWriteReportInvalidDir(t *testing.T) {
	r := &EvalReport{PromptID: "test", ConfigName: "cfg"}
	p := &prompt.Prompt{ID: "test", Properties: map[string]string{"service": "svc", "plane": "plane", "language": "lang", "category": "cat"}}

	// Use a path containing characters that are invalid on both Unix and Windows.
	// On Windows, /nonexistent is treated as drive-relative and MkdirAll may succeed.
	invalidDir := filepath.Join(t.TempDir(), "not\x00valid")
	_, err := WriteReport(r, invalidDir, "run1", p)
	if err == nil {
		t.Fatal("expected error for invalid directory")
	}
}

func TestGraderResultsRoundTrip(t *testing.T) {
dir := t.TempDir()

graders := []GraderResult{
{
GraderName:   "claude-opus-4.6",
GraderType:   "review",
Model:        "claude-opus-4.6",
OverallScore: 4,
MaxScore:     5,
Summary:      "Good code",
Issues:       []string{"missing retry"},
Strengths:    []string{"clean design"},
},
{
GraderName:   "consensus",
GraderType:   "review",
OverallScore: 4,
MaxScore:     5,
Summary:      "Consensus result",
IsConsensus:  true,
},
}

r := &EvalReport{
SchemaVersion:  CurrentSchemaVersion,
PromptID:       "grader-test",
ConfigName:     "baseline",
Timestamp:      "2024-01-15T10:00:00Z",
Duration:       10.0,
PromptMeta:     map[string]any{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
ConfigUsed:     map[string]any{"name": "baseline"},
GeneratedFiles: []string{"main.go"},
GraderResults:  graders,
Success:        true,
}

p := &prompt.Prompt{
ID:         "grader-test",
Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
}

reportPath, err := WriteReport(r, dir, "run1", p)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

data, err := os.ReadFile(reportPath)
if err != nil {
t.Fatalf("failed to read: %v", err)
}

var parsed EvalReport
if err := json.Unmarshal(data, &parsed); err != nil {
t.Fatalf("invalid JSON: %v", err)
}

if parsed.SchemaVersion != CurrentSchemaVersion {
t.Errorf("expected schema version %d, got %d", CurrentSchemaVersion, parsed.SchemaVersion)
}
if len(parsed.GraderResults) != 2 {
t.Fatalf("expected 2 grader results, got %d", len(parsed.GraderResults))
}
if parsed.GraderResults[0].GraderName != "claude-opus-4.6" {
t.Errorf("expected grader name 'claude-opus-4.6', got %q", parsed.GraderResults[0].GraderName)
}
if !parsed.GraderResults[1].IsConsensus {
t.Error("expected second grader result to be consensus")
}
if len(parsed.GraderResults[0].Issues) != 1 {
t.Errorf("expected 1 issue, got %d", len(parsed.GraderResults[0].Issues))
}
}

func TestGraderResultsFromReviewPanel(t *testing.T) {
consolidated := &review.ReviewResult{
Model:        "consolidator",
OverallScore: 3,
MaxScore:     5,
Summary:      "Consensus summary",
Scores: review.ReviewScores{
Criteria: []review.CriterionResult{
{Name: "Builds", Passed: true, Reason: "OK"},
},
},
}
panel := []review.ReviewResult{
{Model: "model-a", OverallScore: 4, MaxScore: 5, Summary: "Model A says good"},
{Model: "model-b", OverallScore: 2, MaxScore: 5, Summary: "Model B says bad"},
}

results := GraderResultsFromReview(consolidated, panel)
if len(results) != 3 {
t.Fatalf("expected 3 grader results (2 panel + 1 consensus), got %d", len(results))
}
if results[0].GraderName != "model-a" {
t.Errorf("expected first grader 'model-a', got %q", results[0].GraderName)
}
if results[2].IsConsensus != true {
t.Error("expected last grader result to be consensus")
}
if results[2].GraderName != "consolidator" {
t.Errorf("expected consensus grader 'consolidator', got %q", results[2].GraderName)
}
}

func TestGraderResultsFromSingleReview(t *testing.T) {
single := &review.ReviewResult{
Model:        "claude-sonnet",
OverallScore: 5,
MaxScore:     5,
Summary:      "Perfect",
}

results := GraderResultsFromReview(single, nil)
if len(results) != 1 {
t.Fatalf("expected 1 grader result, got %d", len(results))
}
if results[0].IsConsensus {
t.Error("single reviewer should not be marked as consensus")
}
if results[0].GraderName != "claude-sonnet" {
t.Errorf("expected 'claude-sonnet', got %q", results[0].GraderName)
}
}

func TestMigrateToV2(t *testing.T) {
r := &EvalReport{
PromptID:   "migrate-test",
ConfigName: "baseline",
Review: &review.ReviewResult{
Model:        "test-model",
OverallScore: 3,
MaxScore:     5,
Summary:      "Test summary",
},
ReviewPanel: []review.ReviewResult{
{Model: "panel-a", OverallScore: 4, MaxScore: 5, Summary: "Panel A"},
},
}

MigrateToV2(r)

if r.SchemaVersion != CurrentSchemaVersion {
t.Errorf("expected schema version %d, got %d", CurrentSchemaVersion, r.SchemaVersion)
}
if len(r.GraderResults) != 2 {
t.Fatalf("expected 2 grader results (1 panel + 1 consensus), got %d", len(r.GraderResults))
}
if r.GraderResults[0].GraderName != "panel-a" {
t.Errorf("expected panel grader 'panel-a', got %q", r.GraderResults[0].GraderName)
}

// Idempotent: calling again should not change anything
r.GraderResults[0].GraderName = "modified"
MigrateToV2(r)
if r.GraderResults[0].GraderName != "modified" {
t.Error("MigrateToV2 should be idempotent and not overwrite existing grader results")
}
}
