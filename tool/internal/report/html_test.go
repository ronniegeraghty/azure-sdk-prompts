package report

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/build"
	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/review"
)

func TestWriteHTMLReport(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:   "test-prompt",
		ConfigName: "baseline",
		Timestamp:  "2024-01-15T10:00:00Z",
		Duration:   12.5,
		PromptMeta: map[string]any{"service": "storage", "language": "dotnet"},
		ConfigUsed: map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles: []string{"Program.cs"},
		Build: &build.BuildResult{
			Language: "dotnet",
			Command:  "dotnet build",
			ExitCode: 0,
			Success:  true,
		},
		Review: &review.ReviewResult{
			Scores: review.ReviewScores{
				Correctness:   8,
				Completeness:  7,
				BestPractices: 9,
				ErrorHandling: 6,
				PackageUsage:  8,
				CodeQuality:   7,
			},
			OverallScore: 8,
			Summary:      "Good implementation",
			Issues:       []string{"Missing retry logic"},
			Strengths:    []string{"Clean code structure"},
		},
		EventCount: 15,
		ToolCalls:  []string{"create_file", "edit_file"},
		Success:    true,
	}

	reportPath, err := WriteHTMLReport(r, dir, "20240115-100000", "storage", "data-plane", "dotnet", "authentication")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	checks := []string{
		"test-prompt",
		"baseline",
		"PASS",
		"8/10",
		"Correctness",
		"Good implementation",
		"Program.cs",
		"dotnet build",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("HTML report missing %q", check)
		}
	}

	expectedDir := filepath.Join(dir, "20240115-100000", "results", "storage", "data-plane", "dotnet", "authentication", "baseline")
	if _, err := os.Stat(expectedDir); err != nil {
		t.Errorf("expected directory %s to exist", expectedDir)
	}
}

func TestWriteHTMLReportNoReview(t *testing.T) {
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

	reportPath, err := WriteHTMLReport(r, dir, "run1", "svc", "plane", "lang", "cat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "FAIL") {
		t.Error("expected FAIL badge")
	}
	if !strings.Contains(content, "timeout exceeded") {
		t.Error("expected error message in report")
	}
}

func TestWriteSummaryHTML(t *testing.T) {
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
				Build:      &build.BuildResult{Success: true},
				Review:     &review.ReviewResult{OverallScore: 8},
			},
			{
				PromptID:   "prompt-a",
				ConfigName: "azure-mcp",
				Success:    true,
				Build:      &build.BuildResult{Success: true},
				Review:     &review.ReviewResult{OverallScore: 9},
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "baseline",
				Success:    false,
				Build:      &build.BuildResult{Success: false},
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "azure-mcp",
				Success:    true,
				Build:      &build.BuildResult{Success: true},
				Review:     &review.ReviewResult{OverallScore: 7},
			},
		},
	}

	summaryPath, err := WriteSummaryHTML(s, dir)
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
		"8/10",
		"9/10",
		"7/10",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("summary HTML missing %q", check)
		}
	}
}

func TestBuildMatrix(t *testing.T) {
	s := &RunSummary{
		Results: []*EvalReport{
			{PromptID: "p1", ConfigName: "c1", Build: &build.BuildResult{Success: true}, Review: &review.ReviewResult{OverallScore: 8}},
			{PromptID: "p1", ConfigName: "c2", Build: &build.BuildResult{Success: false}},
			{PromptID: "p2", ConfigName: "c1", Error: "timeout"},
		},
	}

	m := buildMatrix(s)

	if len(m.Prompts) != 2 {
		t.Errorf("expected 2 prompts, got %d", len(m.Prompts))
	}
	if len(m.Configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(m.Configs))
	}

	cell := m.Cells["p1"]["c1"]
	if cell == nil {
		t.Fatal("expected cell for p1/c1")
	}
	if cell.Score != 8 {
		t.Errorf("expected score 8, got %d", cell.Score)
	}
	if !cell.BuildPass {
		t.Error("expected build pass")
	}

	errCell := m.Cells["p2"]["c1"]
	if errCell == nil {
		t.Fatal("expected cell for p2/c1")
	}
	if errCell.Error != "timeout" {
		t.Errorf("expected timeout error, got %q", errCell.Error)
	}
}
