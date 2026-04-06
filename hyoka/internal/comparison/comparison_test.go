package comparison

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ronniegeraghty/hyoka/internal/report"
	"github.com/ronniegeraghty/hyoka/internal/review"
)

// ---------- Helpers ----------

func boolPtr(b bool) *bool { return &b }

// writeReport writes a report.json file under reportsDir/{runID}/results/{promptID}/{config}/report.json
func writeReport(t *testing.T, reportsDir, runID, promptID, config, timestamp string, graders []report.GraderResult, rev *review.ReviewResult) {
	t.Helper()
	// Flatten config path (replace / with os separator for directory nesting).
	configDir := filepath.Join(reportsDir, runID, "results", promptID, filepath.FromSlash(config))
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatal(err)
	}

	r := report.EvalReport{
		SchemaVersion: 2,
		PromptID:      promptID,
		ConfigName:    config,
		Timestamp:     timestamp,
		Success:       true,
		GraderResults: graders,
		Review:        rev,
	}
	if len(graders) > 0 {
		r.ScoreBreakdown = report.BuildScoreBreakdown(graders)
	}

	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configDir, "report.json"), data, 0644); err != nil {
		t.Fatal(err)
	}
}

func makeGraders(scores ...float64) []report.GraderResult {
	var graders []report.GraderResult
	for i, s := range scores {
		pass := s >= 0.5
		graders = append(graders, report.GraderResult{
			GraderName: graderName(i),
			GraderType: "prompt",
			Score:      s,
			Weight:     1.0,
			Pass:       boolPtr(pass),
		})
	}
	return graders
}

func graderName(i int) string {
	names := []string{"correctness", "completeness", "best-practices", "security", "documentation"}
	if i < len(names) {
		return names[i]
	}
	return "grader-" + string(rune('a'+i))
}

// ---------- CompareConfigs ----------

func TestCompareConfigs_BasicDiff(t *testing.T) {
	dir := t.TempDir()

	// Config A: two prompts, scores 0.6 and 0.8
	writeReport(t, dir, "run-001", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.6), nil)
	writeReport(t, dir, "run-001", "prompt-beta", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.8), nil)

	// Config B: same prompts, scores 0.9 and 0.7
	writeReport(t, dir, "run-001", "prompt-alpha", "azure-mcp/opus", "2025-01-01T10:00:00Z", makeGraders(0.9), nil)
	writeReport(t, dir, "run-001", "prompt-beta", "azure-mcp/opus", "2025-01-01T10:00:00Z", makeGraders(0.7), nil)

	cmp, err := CompareConfigs(dir, "baseline/opus", "azure-mcp/opus")
	if err != nil {
		t.Fatal(err)
	}

	if cmp.ConfigA != "baseline/opus" || cmp.ConfigB != "azure-mcp/opus" {
		t.Errorf("config names mismatch: got %q / %q", cmp.ConfigA, cmp.ConfigB)
	}
	if len(cmp.PerPrompt) != 2 {
		t.Fatalf("expected 2 prompt diffs, got %d", len(cmp.PerPrompt))
	}

	// Find prompt-alpha: delta should be +0.3
	for _, pd := range cmp.PerPrompt {
		if pd.PromptID == "prompt-alpha" {
			if math.Abs(pd.Delta-0.3) > 0.01 {
				t.Errorf("prompt-alpha: expected delta ~0.3, got %f", pd.Delta)
			}
		}
		if pd.PromptID == "prompt-beta" {
			if math.Abs(pd.Delta-(-0.1)) > 0.01 {
				t.Errorf("prompt-beta: expected delta ~-0.1, got %f", pd.Delta)
			}
		}
	}

	if cmp.Summary.Improved != 1 {
		t.Errorf("expected 1 improved, got %d", cmp.Summary.Improved)
	}
	if cmp.Summary.Regressed != 1 {
		t.Errorf("expected 1 regressed, got %d", cmp.Summary.Regressed)
	}
}

func TestCompareConfigs_PicksLatestReport(t *testing.T) {
	dir := t.TempDir()

	// Two runs for config A, same prompt — older run has score 0.4, newer has 0.9
	writeReport(t, dir, "run-001", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.4), nil)
	writeReport(t, dir, "run-002", "prompt-alpha", "baseline/opus", "2025-01-02T10:00:00Z", makeGraders(0.9), nil)

	// Config B with score 0.9
	writeReport(t, dir, "run-001", "prompt-alpha", "azure-mcp/opus", "2025-01-01T10:00:00Z", makeGraders(0.9), nil)

	cmp, err := CompareConfigs(dir, "baseline/opus", "azure-mcp/opus")
	if err != nil {
		t.Fatal(err)
	}

	// Should use the latest (0.9), so delta ≈ 0
	for _, pd := range cmp.PerPrompt {
		if pd.PromptID == "prompt-alpha" {
			if math.Abs(pd.Delta) > 0.01 {
				t.Errorf("expected delta ~0 (latest used), got %f", pd.Delta)
			}
		}
	}
}

func TestCompareConfigs_DisjointPrompts(t *testing.T) {
	dir := t.TempDir()

	writeReport(t, dir, "run-001", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", makeGraders(0.7), nil)
	writeReport(t, dir, "run-001", "prompt-beta", "config-b", "2025-01-01T10:00:00Z", makeGraders(0.8), nil)

	cmp, err := CompareConfigs(dir, "config-a", "config-b")
	if err != nil {
		t.Fatal(err)
	}

	if len(cmp.PerPrompt) != 2 {
		t.Fatalf("expected 2 diffs (disjoint), got %d", len(cmp.PerPrompt))
	}

	for _, pd := range cmp.PerPrompt {
		if pd.PromptID == "prompt-alpha" && !pd.OnlyInA {
			t.Error("prompt-alpha should be OnlyInA")
		}
		if pd.PromptID == "prompt-beta" && !pd.OnlyInB {
			t.Error("prompt-beta should be OnlyInB")
		}
	}

	// Disjoint prompts don't contribute to improved/regressed counts.
	if cmp.Summary.Improved != 0 || cmp.Summary.Regressed != 0 {
		t.Errorf("disjoint prompts should not affect improved/regressed counts")
	}
}

func TestCompareConfigs_NoReports(t *testing.T) {
	dir := t.TempDir()
	_, err := CompareConfigs(dir, "no-such", "also-missing")
	if err == nil {
		t.Error("expected error for empty reports dir")
	}
}

func TestCompareConfigs_GraderDiffs(t *testing.T) {
	dir := t.TempDir()

	gradersA := []report.GraderResult{
		{GraderName: "correctness", GraderType: "prompt", Score: 0.5, Weight: 1.0, Pass: boolPtr(true)},
		{GraderName: "security", GraderType: "prompt", Score: 0.3, Weight: 1.0, Pass: boolPtr(false)},
	}
	gradersB := []report.GraderResult{
		{GraderName: "correctness", GraderType: "prompt", Score: 0.8, Weight: 1.0, Pass: boolPtr(true)},
		{GraderName: "security", GraderType: "prompt", Score: 0.9, Weight: 1.0, Pass: boolPtr(true)},
	}

	writeReport(t, dir, "run-001", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", gradersA, nil)
	writeReport(t, dir, "run-001", "prompt-alpha", "config-b", "2025-01-01T10:00:00Z", gradersB, nil)

	cmp, err := CompareConfigs(dir, "config-a", "config-b")
	if err != nil {
		t.Fatal(err)
	}

	if len(cmp.PerPrompt) != 1 {
		t.Fatalf("expected 1 prompt diff, got %d", len(cmp.PerPrompt))
	}

	pd := cmp.PerPrompt[0]
	if len(pd.GraderDiffs) != 2 {
		t.Fatalf("expected 2 grader diffs, got %d", len(pd.GraderDiffs))
	}

	for _, gd := range pd.GraderDiffs {
		switch gd.Name {
		case "correctness":
			if math.Abs(gd.Delta-0.3) > 0.01 {
				t.Errorf("correctness delta: expected 0.3, got %f", gd.Delta)
			}
		case "security":
			if math.Abs(gd.Delta-0.6) > 0.01 {
				t.Errorf("security delta: expected 0.6, got %f", gd.Delta)
			}
			if gd.PassA != false || gd.PassB != true {
				t.Errorf("security pass: expected false→true, got %v→%v", gd.PassA, gd.PassB)
			}
		}
	}
}

// ---------- CompareRuns ----------

func TestCompareRuns_BasicDiff(t *testing.T) {
	dir := t.TempDir()

	writeReport(t, dir, "20250101-100000", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.5), nil)
	writeReport(t, dir, "20250102-100000", "prompt-alpha", "baseline/opus", "2025-01-02T10:00:00Z", makeGraders(0.8), nil)

	cmp, err := CompareRuns(dir, "20250101-100000", "20250102-100000")
	if err != nil {
		t.Fatal(err)
	}

	if cmp.RunA != "20250101-100000" || cmp.RunB != "20250102-100000" {
		t.Errorf("run IDs mismatch")
	}
	if len(cmp.PerPrompt) != 1 {
		t.Fatalf("expected 1 diff, got %d", len(cmp.PerPrompt))
	}
	if math.Abs(cmp.PerPrompt[0].Delta-0.3) > 0.01 {
		t.Errorf("expected delta ~0.3, got %f", cmp.PerPrompt[0].Delta)
	}
	if cmp.Summary.Improved != 1 {
		t.Errorf("expected 1 improved, got %d", cmp.Summary.Improved)
	}
}

func TestCompareRuns_MissingRun(t *testing.T) {
	dir := t.TempDir()
	_, err := CompareRuns(dir, "nonexistent-a", "nonexistent-b")
	if err == nil {
		t.Error("expected error for missing run directories")
	}
}

func TestCompareRuns_MultipleConfigs(t *testing.T) {
	dir := t.TempDir()

	// Run A has two configs for the same prompt.
	writeReport(t, dir, "run-a", "prompt-alpha", "config-x", "2025-01-01T10:00:00Z", makeGraders(0.4), nil)
	writeReport(t, dir, "run-a", "prompt-alpha", "config-y", "2025-01-01T10:00:00Z", makeGraders(0.6), nil)

	// Run B has the same.
	writeReport(t, dir, "run-b", "prompt-alpha", "config-x", "2025-01-02T10:00:00Z", makeGraders(0.8), nil)
	writeReport(t, dir, "run-b", "prompt-alpha", "config-y", "2025-01-02T10:00:00Z", makeGraders(0.7), nil)

	cmp, err := CompareRuns(dir, "run-a", "run-b")
	if err != nil {
		t.Fatal(err)
	}

	// Each prompt+config combination is a separate entry.
	if len(cmp.PerPrompt) < 1 {
		t.Fatal("expected at least 1 prompt diff")
	}

	// All should show improvement.
	if cmp.Summary.Improved < 1 {
		t.Errorf("expected at least 1 improved, got %d", cmp.Summary.Improved)
	}
}

// ---------- TemporalDiff ----------

func TestTemporalDiff_BasicSplit(t *testing.T) {
	dir := t.TempDir()

	// Historical report (before cutoff).
	writeReport(t, dir, "run-old", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.4), nil)
	// Recent report (after cutoff).
	writeReport(t, dir, "run-new", "prompt-alpha", "baseline/opus", "2025-02-01T10:00:00Z", makeGraders(0.9), nil)

	since := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	tc, err := TemporalDiff(dir, "baseline/opus", since)
	if err != nil {
		t.Fatal(err)
	}

	if tc.Config != "baseline/opus" {
		t.Errorf("config mismatch: %q", tc.Config)
	}
	if len(tc.PerPrompt) != 1 {
		t.Fatalf("expected 1 prompt diff, got %d", len(tc.PerPrompt))
	}
	if math.Abs(tc.PerPrompt[0].Delta-0.5) > 0.01 {
		t.Errorf("expected delta ~0.5, got %f", tc.PerPrompt[0].Delta)
	}
	if tc.Summary.Improved != 1 {
		t.Errorf("expected 1 improved, got %d", tc.Summary.Improved)
	}
}

func TestTemporalDiff_NoBaseReports(t *testing.T) {
	dir := t.TempDir()
	writeReport(t, dir, "run-new", "prompt-alpha", "baseline/opus", "2025-02-01T10:00:00Z", makeGraders(0.9), nil)

	since := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	_, err := TemporalDiff(dir, "baseline/opus", since)
	if err == nil {
		t.Error("expected error when no base reports exist")
	}
}

func TestTemporalDiff_NoRecentReports(t *testing.T) {
	dir := t.TempDir()
	writeReport(t, dir, "run-old", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", makeGraders(0.5), nil)

	since := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	_, err := TemporalDiff(dir, "baseline/opus", since)
	if err == nil {
		t.Error("expected error when no recent reports exist")
	}
}

func TestTemporalDiff_NoConfig(t *testing.T) {
	dir := t.TempDir()
	since := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	_, err := TemporalDiff(dir, "nonexistent", since)
	if err == nil {
		t.Error("expected error for nonexistent config")
	}
}

// ---------- LegacyReviewScoring ----------

func TestCompareConfigs_LegacyReviewScoring(t *testing.T) {
	dir := t.TempDir()

	// Config A: legacy review-based scoring (no graders).
	writeReport(t, dir, "run-001", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", nil,
		&review.ReviewResult{OverallScore: 7, MaxScore: 10})

	// Config B: legacy review-based scoring.
	writeReport(t, dir, "run-001", "prompt-alpha", "config-b", "2025-01-01T10:00:00Z", nil,
		&review.ReviewResult{OverallScore: 9, MaxScore: 10})

	cmp, err := CompareConfigs(dir, "config-a", "config-b")
	if err != nil {
		t.Fatal(err)
	}

	if len(cmp.PerPrompt) != 1 {
		t.Fatalf("expected 1 diff, got %d", len(cmp.PerPrompt))
	}
	// 0.9 - 0.7 = 0.2
	if math.Abs(cmp.PerPrompt[0].Delta-0.2) > 0.01 {
		t.Errorf("expected delta ~0.2, got %f", cmp.PerPrompt[0].Delta)
	}
}

// ---------- Summary ----------

func TestBuildSummary_TopN(t *testing.T) {
	// Create 10 diffs: 7 improved, 3 regressed.
	var diffs []PromptDiff
	for i := 0; i < 7; i++ {
		diffs = append(diffs, PromptDiff{
			PromptID: "improved-" + string(rune('a'+i)),
			ScoreA:   0.3,
			ScoreB:   0.3 + float64(i+1)*0.05,
			Delta:    float64(i+1) * 0.05,
		})
	}
	for i := 0; i < 3; i++ {
		diffs = append(diffs, PromptDiff{
			PromptID: "regressed-" + string(rune('a'+i)),
			ScoreA:   0.8,
			ScoreB:   0.8 - float64(i+1)*0.1,
			Delta:    -float64(i+1) * 0.1,
		})
	}

	s := buildSummary(diffs)

	if s.Improved != 7 {
		t.Errorf("expected 7 improved, got %d", s.Improved)
	}
	if s.Regressed != 3 {
		t.Errorf("expected 3 regressed, got %d", s.Regressed)
	}
	if len(s.TopImproved) != 5 {
		t.Errorf("expected top 5 improved, got %d", len(s.TopImproved))
	}
	if len(s.TopRegressed) != 3 {
		t.Errorf("expected top 3 regressed, got %d", len(s.TopRegressed))
	}

	// TopImproved should be sorted descending by delta.
	if s.TopImproved[0].Delta < s.TopImproved[len(s.TopImproved)-1].Delta {
		t.Error("TopImproved should be sorted by delta descending")
	}
	// TopRegressed should be sorted ascending (most negative first).
	if s.TopRegressed[0].Delta > s.TopRegressed[len(s.TopRegressed)-1].Delta {
		t.Error("TopRegressed should be sorted by delta ascending (most negative first)")
	}
}

func TestBuildSummary_AllUnchanged(t *testing.T) {
	diffs := []PromptDiff{
		{PromptID: "p1", ScoreA: 0.5, ScoreB: 0.5, Delta: 0.0},
		{PromptID: "p2", ScoreA: 0.7, ScoreB: 0.7, Delta: 0.0},
	}
	s := buildSummary(diffs)
	if s.Unchanged != 2 {
		t.Errorf("expected 2 unchanged, got %d", s.Unchanged)
	}
	if s.AvgDelta != 0 {
		t.Errorf("expected avg delta 0, got %f", s.AvgDelta)
	}
}

// ---------- Internal Helpers ----------

func TestScoreFromReport_Graders(t *testing.T) {
	graders := makeGraders(0.6, 0.8)
	r := &report.EvalReport{GraderResults: graders}
	r.ScoreBreakdown = report.BuildScoreBreakdown(graders)

	score := scoreFromReport(r)
	// Average of 0.6, 0.8 = 0.7
	if math.Abs(score-0.7) > 0.01 {
		t.Errorf("expected score ~0.7, got %f", score)
	}
}

func TestScoreFromReport_LegacyReview(t *testing.T) {
	r := &report.EvalReport{
		Review: &review.ReviewResult{OverallScore: 8, MaxScore: 10},
	}
	score := scoreFromReport(r)
	if math.Abs(score-0.8) > 0.01 {
		t.Errorf("expected score ~0.8, got %f", score)
	}
}

func TestScoreFromReport_NoScore(t *testing.T) {
	r := &report.EvalReport{}
	score := scoreFromReport(r)
	if score != 0.0 {
		t.Errorf("expected 0.0 for empty report, got %f", score)
	}
}

func TestPromptIDFromKey(t *testing.T) {
	cases := []struct {
		key, want string
	}{
		{"prompt-alpha", "prompt-alpha"},
		{"prompt-alpha|baseline/opus", "prompt-alpha"},
		{"a|b|c", "a"},
	}
	for _, tc := range cases {
		got := promptIDFromKey(tc.key)
		if got != tc.want {
			t.Errorf("promptIDFromKey(%q) = %q, want %q", tc.key, got, tc.want)
		}
	}
}
