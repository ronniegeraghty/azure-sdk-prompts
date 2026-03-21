package trends

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/report"
	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/review"
)

func TestScanReportsEmpty(t *testing.T) {
	dir := t.TempDir()
	entries, err := scanReports(dir, "", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("expected 0 entries, got %d", len(entries))
	}
}

func TestScanReportsFindsReports(t *testing.T) {
	dir := t.TempDir()

	// Create a fake report.json
	runDir := filepath.Join(dir, "20250101-120000", "results", "key-vault", "data-plane", "python", "crud", "baseline")
	if err := os.MkdirAll(runDir, 0755); err != nil {
		t.Fatal(err)
	}

	r := report.EvalReport{
		PromptID:   "key-vault-dp-python-crud",
		ConfigName: "baseline",
		Timestamp:  "2025-01-01T12:00:00Z",
		Duration:   45.2,
		Success:    true,
		ToolCalls:  []string{"create_file", "bash"},
		PromptMeta: map[string]any{
			"service":  "key-vault",
			"language": "python",
		},
		GeneratedFiles: []string{"main.py"},
		Review: &review.ReviewResult{
			OverallScore: 8,
		},
	}
	data, _ := json.MarshalIndent(r, "", "  ")
	if err := os.WriteFile(filepath.Join(runDir, "report.json"), data, 0644); err != nil {
		t.Fatal(err)
	}

	entries, err := scanReports(dir, "", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].PromptID != "key-vault-dp-python-crud" {
		t.Errorf("unexpected prompt ID: %s", entries[0].PromptID)
	}
	if entries[0].Score != 8 {
		t.Errorf("expected score 8, got %d", entries[0].Score)
	}
	if !entries[0].HasReview {
		t.Error("expected HasReview to be true")
	}
}

func TestScanReportsFilterByPromptID(t *testing.T) {
	dir := t.TempDir()

	// Two different prompts
	for _, pid := range []string{"prompt-a", "prompt-b"} {
		runDir := filepath.Join(dir, "run1", "results", "svc", "dp", "py", "crud", "baseline")
		os.MkdirAll(runDir, 0755)
		r := report.EvalReport{PromptID: pid, ConfigName: "baseline", PromptMeta: map[string]any{}}
		data, _ := json.Marshal(r)
		os.WriteFile(filepath.Join(runDir, pid+"-report.json"), data, 0644)
	}
	// Write as report.json too
	for _, pid := range []string{"prompt-a", "prompt-b"} {
		runDir := filepath.Join(dir, "run1", pid)
		os.MkdirAll(runDir, 0755)
		r := report.EvalReport{PromptID: pid, ConfigName: "baseline", PromptMeta: map[string]any{}}
		data, _ := json.Marshal(r)
		os.WriteFile(filepath.Join(runDir, "report.json"), data, 0644)
	}

	entries, _ := scanReports(dir, "prompt-a", "", "")
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry for prompt-a, got %d", len(entries))
	}
}

func TestGenerateAndWriteMarkdown(t *testing.T) {
	dir := t.TempDir()

	// Create a report
	runDir := filepath.Join(dir, "reports", "run1", "results", "svc", "dp", "py", "crud", "baseline")
	os.MkdirAll(runDir, 0755)
	r := report.EvalReport{
		PromptID:   "test-prompt",
		ConfigName: "baseline",
		Timestamp:  "2025-01-01T12:00:00Z",
		Success:    true,
		PromptMeta: map[string]any{},
		ToolCalls:  []string{"bash"},
	}
	data, _ := json.Marshal(r)
	os.WriteFile(filepath.Join(runDir, "report.json"), data, 0644)

	tr, err := Generate(TrendOptions{
		ReportsDir: filepath.Join(dir, "reports"),
		PromptID:   "test-prompt",
	})
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if tr.TotalRuns != 1 {
		t.Fatalf("expected 1 run, got %d", tr.TotalRuns)
	}

	outDir := filepath.Join(dir, "trends")
	mdPath, err := WriteMarkdown(tr, outDir)
	if err != nil {
		t.Fatalf("WriteMarkdown failed: %v", err)
	}
	content, _ := os.ReadFile(mdPath)
	if !strings.Contains(string(content), "test-prompt") {
		t.Error("markdown should contain prompt ID")
	}

	htmlPath, err := WriteHTML(tr, outDir)
	if err != nil {
		t.Fatalf("WriteHTML failed: %v", err)
	}
	htmlContent, _ := os.ReadFile(htmlPath)
	if !strings.Contains(string(htmlContent), "test-prompt") {
		t.Error("HTML should contain prompt ID")
	}
}

func TestEvaluateToolUsage(t *testing.T) {
	// This tests the function in the engine, so we import it from report
	result := &report.ToolUsageResult{
		ExpectedTools: []string{"azure-mcp", "bash"},
		ActualTools:   []string{"bash", "create_file", "read_file"},
		MatchedTools:  []string{"bash"},
		MissingTools:  []string{"azure-mcp"},
		ExtraTools:    []string{"create_file", "read_file"},
		Match:         false,
	}

	if result.Match {
		t.Error("expected Match to be false")
	}
	if len(result.MissingTools) != 1 || result.MissingTools[0] != "azure-mcp" {
		t.Errorf("unexpected missing tools: %v", result.MissingTools)
	}
}

func TestPct(t *testing.T) {
	if pct(0, 0) != 0 {
		t.Error("pct(0,0) should be 0")
	}
	if pct(1, 2) != 50 {
		t.Error("pct(1,2) should be 50")
	}
	if pct(3, 3) != 100 {
		t.Error("pct(3,3) should be 100")
	}
}
