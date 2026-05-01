package report

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func TestWriteMarkdownReport(t *testing.T) {
	dir := t.TempDir()

	boolTrue := true
	r := &EvalReport{
		PromptID:           "test-prompt",
		ConfigName:         "baseline",
		Timestamp:          "2024-01-15T10:00:00Z",
		Duration:           12.5,
		GenerationDuration: 8.2,
		ReviewDuration:     3.1,
		PromptMeta:         map[string]any{"service": "storage", "language": "dotnet"},
		ConfigUsed:         map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles:     []string{"Program.cs"},
		Review: &review.ReviewResult{
			Scores: review.ReviewScores{
				Criteria: []review.CriterionResult{
					{Name: "Code Builds", Passed: true, Reason: "Compiles successfully"},
					{Name: "Latest Packages", Passed: true, Reason: "Using latest versions"},
					{Name: "Best Practices", Passed: true, Reason: "Uses DefaultAzureCredential"},
					{Name: "Error Handling", Passed: false, Reason: "Missing retry logic"},
					{Name: "Code Quality", Passed: true, Reason: "Clean structure"},
				},
			},
			OverallScore: 4,
			MaxScore:     5,
			Summary:      "Good implementation",
			Issues:       []string{"Missing retry logic"},
			Strengths:    []string{"Clean code structure"},
		},
		GraderResults: []GraderResult{
			{GraderName: "md-grader", GraderType: "review", Score: 0.8, Weight: 1.0, Pass: true, Message: "Grader output for MD", Points: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
			{GraderName: "consensus", GraderType: "review", Score: 0.8, Weight: 1.0, Pass: true, Message: "MD consensus", Points: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		},
		SessionEvents: []SessionEventRecord{
			{Type: "user.message", Content: "Write a dotnet storage auth sample"},
			{Type: "assistant.reasoning", Content: "I need to create an auth sample"},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"Program.cs"}`},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "File created", ToolSuccess: &boolTrue, Duration: 150.5},
			{Type: "assistant.message", Content: "Here is your sample"},
		},
		EventCount: 15,
		ToolCalls:  []string{"create_file", "edit_file"},
		Success:    true,
	}

	reportPath, err := WriteMarkdownReport(r, dir, "20240115-100000", "storage", "data-plane", "dotnet", "authentication")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	checks := []string{
		"# Evaluation Report: test-prompt",
		"baseline",
		"✅ PASSED",
		"Program.cs",
		"Write a dotnet storage auth sample",
		"I need to create an auth sample",
		"Reviewer Notes",
		"Good implementation",
		"Clean code structure",
		"Missing retry logic",
		"Tool Calls",
		"Back to Summary",
		"File created",
		"150ms",
		"Phase Timing",
		"Generation",
		"8.2s",
		"Review",
		"3.1s",
		"Grader Results",
		"md-grader",
		"consensus",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("Markdown report missing %q", check)
		}
	}
	// The legacy "Code Review (LLM-as-Judge)" criteria-passed line and the
	// per-criterion Criteria Results table are intentionally NOT rendered:
	// the prompt_review grader is now displayed under "## Grader Results"
	// and the duplicate "X/Y criteria passed" text was confusing.
	for _, antiCheck := range []string{
		"Code Review (LLM-as-Judge)",
		"criteria passed",
		"Criteria Results",
		"4/5",
		"Code Builds",
	} {
		if strings.Contains(content, antiCheck) {
			t.Errorf("Markdown report unexpectedly contains %q", antiCheck)
		}
	}
	// Graders use flat fallback (no SourceFile): grader name + pass/fail count rendered.
	if !strings.Contains(content, "md-grader (review): Pass (1/1)") {
		t.Errorf("Markdown report missing grader pass/fail line")
	}

	expectedDir := filepath.Join(dir, "20240115-100000", "results", "storage", "data-plane", "dotnet", "authentication", "test-prompt", "baseline")
	if _, err := os.Stat(expectedDir); err != nil {
		t.Errorf("expected directory %s to exist", expectedDir)
	}
}

func TestWriteMarkdownReportFailed(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:       "test-prompt",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       5.0,
		PromptMeta:     map[string]any{},
		ConfigUsed:     map[string]any{},
		GeneratedFiles: []string{},
		Success:        false,
		Error:          "timeout exceeded",
	}

	reportPath, err := WriteMarkdownReport(r, dir, "run1", "svc", "plane", "lang", "cat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "FAILED") {
		t.Error("expected FAILED in report")
	}
	if !strings.Contains(content, "timeout exceeded") {
		t.Error("expected error message in report")
	}
}

func TestWriteSummaryMarkdown(t *testing.T) {
	dir := t.TempDir()

	s := &RunSummary{
		RunID:        "20240115-100000",
		Timestamp:    "2024-01-15T10:00:00Z",
		TotalPrompts: 2,
		TotalConfigs: 2,
		TotalEvals:   4,
		Passed:       3,
		Failed:       1,
		Errors:       0,
		Duration:     120.5,
		Results: []*EvalReport{
			{
				PromptID:   "prompt-a",
				ConfigName: "baseline",
				Success:    true,
				Duration:   10.0,
				PromptMeta: map[string]any{"service": "storage", "plane": "data-plane", "language": "dotnet", "category": "auth"},
				Review:     &review.ReviewResult{OverallScore: 4, MaxScore: 5},
			},
			{
				PromptID:   "prompt-a",
				ConfigName: "azure-mcp",
				Success:    true,
				Duration:   15.0,
				PromptMeta: map[string]any{"service": "storage", "plane": "data-plane", "language": "dotnet", "category": "auth"},
				Review:     &review.ReviewResult{OverallScore: 5, MaxScore: 5},
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "baseline",
				Success:    false,
				Duration:   5.0,
				PromptMeta: map[string]any{},
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "azure-mcp",
				Success:    true,
				Duration:   12.0,
				PromptMeta: map[string]any{},
				Review:     &review.ReviewResult{OverallScore: 3, MaxScore: 5},
			},
		},
	}

	summaryPath, err := WriteSummaryMarkdown(s, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	content := string(data)
	checks := []string{
		"Evaluation Summary",
		"20240115-100000",
		"prompt-a",
		"prompt-b",
		"baseline",
		"azure-mcp",
		"4/5",
		"5/5",
		"3/5",
		"Comparison Matrix",
		"Detailed Results",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("summary markdown missing %q", check)
		}
	}

	// Verify pass/fail icons
	passCount := strings.Count(content, "✅")
	failCount := strings.Count(content, "❌")
	// Matrix: 3 pass + 1 fail, Detail table: 3 pass + 1 fail = 6 pass, 2 fail
	if passCount != 6 {
		t.Errorf("expected 6 ✅ icons (3 matrix + 3 detail), got %d", passCount)
	}
	if failCount != 2 {
		t.Errorf("expected 2 ❌ icons (1 matrix + 1 detail), got %d", failCount)
	}
}

func TestWriteMarkdownReportStub(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:       "test-prompt",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       1.0,
		PromptMeta:     map[string]any{},
		ConfigUsed:     map[string]any{},
		GeneratedFiles: []string{},
		Success:        true,
		IsStub:         true,
	}

	reportPath, err := WriteMarkdownReport(r, dir, "run1", "svc", "plane", "lang", "cat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "Stub") {
		t.Error("expected Stub indicator in report")
	}
}

func TestTruncateStr(t *testing.T) {
	short := "hello"
	if truncateStr(short, 10) != "hello" {
		t.Error("short string should not be truncated")
	}

	long := strings.Repeat("a", 100)
	result := truncateStr(long, 50)
	if len(result) <= 50 {
		t.Error("truncated string should have suffix")
	}
	if !strings.Contains(result, "truncated") {
		t.Error("truncated string should contain truncation marker")
	}
}

func TestWriteGraderResults_GroupedBySourceFile(t *testing.T) {
	results := []GraderResult{
		{
			GraderName: "Eval Criteria",
			GraderType: "prompt",
			Pass:       true,
			SourceFile: "/prompts/crud-secrets.prompt.md",
			SourceType: "prompt_file",
			Points: []GraderPoint{
				{Label: "check 1", Pass: true},
				{Label: "check 2", Pass: true},
			},
		},
		{
			GraderName: "DefaultAzureCredential Authentication",
			GraderType: "prompt",
			Pass:       false,
			SourceFile: "/criteria/python.yaml",
			SourceType: "criteria_file",
			Points: []GraderPoint{
				{Label: "Uses DefaultAzureCredential", Pass: true},
				{Label: "Uses async/await patterns", Pass: false},
			},
		},
		{
			GraderName: "Output Files Exist",
			GraderType: "output_check",
			Pass:       true,
			SourceFile: "/criteria/python.yaml",
			SourceType: "criteria_file",
			Points: []GraderPoint{
				{Label: "min_files (1)", Pass: true},
				{Label: "min_bytes_per_file (1)", Pass: true},
			},
		},
	}

	var b strings.Builder
	writeGraderResults(&b, results)
	out := b.String()

	// Level-1 file headers
	if !strings.Contains(out, "crud-secrets.prompt.md (prompt file):") {
		t.Errorf("missing prompt file header, got:\n%s", out)
	}
	if !strings.Contains(out, "python.yaml (criteria file):") {
		t.Errorf("missing criteria file header, got:\n%s", out)
	}

	// Level-2 grader lines (indented 2 spaces under file header)
	if !strings.Contains(out, "  - Eval Criteria (prompt): Pass (2/2)") {
		t.Errorf("missing Eval Criteria grader line, got:\n%s", out)
	}
	if !strings.Contains(out, "  - DefaultAzureCredential Authentication (prompt): Fail (1/2)") {
		t.Errorf("missing DefaultAzureCredential grader line, got:\n%s", out)
	}
	if !strings.Contains(out, "  - Output Files Exist (output_check): Pass (2/2)") {
		t.Errorf("missing Output Files Exist grader line, got:\n%s", out)
	}

	// Level-3 point lines (indented 6 spaces under grader)
	if !strings.Contains(out, "      - check 1: Pass") {
		t.Errorf("missing check 1 point, got:\n%s", out)
	}
	if !strings.Contains(out, "      - Uses async/await patterns: Fail") {
		t.Errorf("missing async/await point, got:\n%s", out)
	}
	if !strings.Contains(out, "      - min_files (1): Pass") {
		t.Errorf("missing min_files point, got:\n%s", out)
	}
}

func TestWriteGraderResults_FlatFallbackWhenNoSourceFile(t *testing.T) {
	results := []GraderResult{
		{
			GraderName: "my-grader",
			GraderType: "output_check",
			Pass:       true,
			// SourceFile intentionally empty
			Points: []GraderPoint{
				{Label: "file exists", Pass: true},
			},
		},
	}

	var b strings.Builder
	writeGraderResults(&b, results)
	out := b.String()

	// No file header
	if strings.Contains(out, "(prompt file)") || strings.Contains(out, "(criteria file)") {
		t.Errorf("unexpected file header in flat fallback, got:\n%s", out)
	}

	// Grader at top level (no leading spaces for file group)
	if !strings.Contains(out, "- my-grader (output_check): Pass (1/1)") {
		t.Errorf("expected flat grader line, got:\n%s", out)
	}
	if !strings.Contains(out, "    - file exists: Pass") {
		t.Errorf("expected flat point line, got:\n%s", out)
	}
}
