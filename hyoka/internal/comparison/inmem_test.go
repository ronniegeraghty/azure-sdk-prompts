package comparison

import (
	"math"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// ---------- CompareReports (in-memory) ----------

func TestCompareReports_ConfigsBasic(t *testing.T) {
	reportsA := []report.EvalReport{
		mkReport("prompt-alpha", "config-a", "2025-01-01T10:00:00Z", 0.5),
		mkReport("prompt-beta", "config-a", "2025-01-01T10:00:00Z", 0.3),
	}
	reportsB := []report.EvalReport{
		mkReport("prompt-alpha", "config-b", "2025-01-01T10:00:00Z", 0.8),
		mkReport("prompt-beta", "config-b", "2025-01-01T10:00:00Z", 0.3),
	}

	got := CompareReports(KindConfigs, "config-a", "config-b", reportsA, reportsB)
	if got == nil {
		t.Fatal("CompareReports returned nil")
	}
	if got.Kind != KindConfigs {
		t.Errorf("kind: got %q", got.Kind)
	}
	if got.LabelA != "config-a" || got.LabelB != "config-b" {
		t.Errorf("labels: got %q / %q", got.LabelA, got.LabelB)
	}
	if len(got.PerPrompt) != 2 {
		t.Fatalf("expected 2 diffs, got %d", len(got.PerPrompt))
	}
	if got.Summary.Improved != 1 || got.Summary.Unchanged != 1 {
		t.Errorf("summary: improved=%d unchanged=%d", got.Summary.Improved, got.Summary.Unchanged)
	}
}

func TestCompareReports_RunsUsesCompositeKey(t *testing.T) {
	// Run A: same prompt in two configs.
	runA := []report.EvalReport{
		mkReport("prompt-alpha", "config-x", "2025-01-01T10:00:00Z", 0.4),
		mkReport("prompt-alpha", "config-y", "2025-01-01T10:00:00Z", 0.6),
	}
	runB := []report.EvalReport{
		mkReport("prompt-alpha", "config-x", "2025-01-02T10:00:00Z", 0.7),
		mkReport("prompt-alpha", "config-y", "2025-01-02T10:00:00Z", 0.6),
	}

	got := CompareReports(KindRuns, "run-a", "run-b", runA, runB)
	// Composite keys mean both config variants are carried through, but
	// diffPromptMaps collapses them by prompt ID for presentation. The
	// summary should reflect improvement from at least one pair.
	if got.Summary.Improved < 1 {
		t.Errorf("expected at least 1 improved, got %d", got.Summary.Improved)
	}
}

func TestCompareReports_DeterministicOrdering(t *testing.T) {
	reportsA := []report.EvalReport{
		mkReport("z-prompt", "a", "2025-01-01T10:00:00Z", 0.5),
		mkReport("a-prompt", "a", "2025-01-01T10:00:00Z", 0.5),
		mkReport("m-prompt", "a", "2025-01-01T10:00:00Z", 0.5),
	}
	reportsB := []report.EvalReport{
		mkReport("z-prompt", "b", "2025-01-01T10:00:00Z", 0.7),
		mkReport("a-prompt", "b", "2025-01-01T10:00:00Z", 0.7),
		mkReport("m-prompt", "b", "2025-01-01T10:00:00Z", 0.7),
	}

	got1 := CompareReports(KindConfigs, "a", "b", reportsA, reportsB)
	got2 := CompareReports(KindConfigs, "a", "b", reportsA, reportsB)

	if len(got1.PerPrompt) != 3 || len(got2.PerPrompt) != 3 {
		t.Fatalf("wrong diff count")
	}
	for i := range got1.PerPrompt {
		if got1.PerPrompt[i].PromptID != got2.PerPrompt[i].PromptID {
			t.Errorf("ordering not deterministic at index %d: %q vs %q",
				i, got1.PerPrompt[i].PromptID, got2.PerPrompt[i].PromptID)
		}
	}
	// Alphabetical order expected.
	wantOrder := []string{"a-prompt", "m-prompt", "z-prompt"}
	for i, want := range wantOrder {
		if got1.PerPrompt[i].PromptID != want {
			t.Errorf("expected %q at index %d, got %q", want, i, got1.PerPrompt[i].PromptID)
		}
	}
}

// ---------- AutoGenerateForRun ----------

func TestAutoGenerateForRun_EmitsPairs(t *testing.T) {
	tests := []struct {
		name        string
		reports     []report.EvalReport
		wantPairs   int
		wantKinds   []ComparisonKind
		wantLabels  [][2]string
	}{
		{
			name:      "single config → no pairs",
			reports:   []report.EvalReport{mkReport("p1", "only", "2025-01-01T10:00:00Z", 0.5)},
			wantPairs: 0,
		},
		{
			name: "two configs → one pair",
			reports: []report.EvalReport{
				mkReport("p1", "alpha", "2025-01-01T10:00:00Z", 0.5),
				mkReport("p1", "beta", "2025-01-01T10:00:00Z", 0.7),
			},
			wantPairs:  1,
			wantKinds:  []ComparisonKind{KindConfigs},
			wantLabels: [][2]string{{"alpha", "beta"}},
		},
		{
			name: "three configs → three pairs, alphabetical",
			reports: []report.EvalReport{
				mkReport("p1", "gamma", "2025-01-01T10:00:00Z", 0.5),
				mkReport("p1", "alpha", "2025-01-01T10:00:00Z", 0.5),
				mkReport("p1", "beta", "2025-01-01T10:00:00Z", 0.5),
			},
			wantPairs: 3,
			wantKinds: []ComparisonKind{KindConfigs, KindConfigs, KindConfigs},
			wantLabels: [][2]string{
				{"alpha", "beta"},
				{"alpha", "gamma"},
				{"beta", "gamma"},
			},
		},
		{
			name:      "empty reports",
			reports:   nil,
			wantPairs: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AutoGenerateForRun(tc.reports)
			if len(got) != tc.wantPairs {
				t.Fatalf("expected %d pairs, got %d", tc.wantPairs, len(got))
			}
			for i, cmp := range got {
				if cmp.Kind != tc.wantKinds[i] {
					t.Errorf("pair %d kind: got %q, want %q", i, cmp.Kind, tc.wantKinds[i])
				}
				if cmp.LabelA != tc.wantLabels[i][0] || cmp.LabelB != tc.wantLabels[i][1] {
					t.Errorf("pair %d labels: got %q/%q, want %q/%q",
						i, cmp.LabelA, cmp.LabelB, tc.wantLabels[i][0], tc.wantLabels[i][1])
				}
			}
		})
	}
}

func TestAutoGenerateForRun_UsesSameEngine(t *testing.T) {
	// The CLI path (disk-backed) and the auto-gen path (in-memory) must
	// produce identical results given identical inputs. This is the core
	// Morpheus gate criterion for #357.
	reports := []report.EvalReport{
		mkReport("prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", 0.5),
		mkReport("prompt-beta", "baseline/opus", "2025-01-01T10:00:00Z", 0.7),
		mkReport("prompt-alpha", "azure-mcp/opus", "2025-01-01T10:00:00Z", 0.8),
		mkReport("prompt-beta", "azure-mcp/opus", "2025-01-01T10:00:00Z", 0.7),
	}

	// Auto-generated path.
	autoGen := AutoGenerateForRun(reports)
	if len(autoGen) != 1 {
		t.Fatalf("expected 1 pair, got %d", len(autoGen))
	}
	auto := autoGen[0]

	// Direct in-memory path.
	var reportsA, reportsB []report.EvalReport
	for _, r := range reports {
		if r.ConfigName == auto.LabelA {
			reportsA = append(reportsA, r)
		} else {
			reportsB = append(reportsB, r)
		}
	}
	direct := CompareReports(KindConfigs, auto.LabelA, auto.LabelB, reportsA, reportsB)

	// Summaries must match byte-for-byte.
	if !summariesEqual(auto.Summary, direct.Summary) {
		t.Errorf("auto-gen summary %+v does not match direct %+v", auto.Summary, direct.Summary)
	}
	if len(auto.PerPrompt) != len(direct.PerPrompt) {
		t.Fatalf("per-prompt length mismatch: %d vs %d", len(auto.PerPrompt), len(direct.PerPrompt))
	}
	for i := range auto.PerPrompt {
		a, d := auto.PerPrompt[i], direct.PerPrompt[i]
		if a.PromptID != d.PromptID {
			t.Errorf("prompt %d id: %q vs %q", i, a.PromptID, d.PromptID)
		}
		if math.Abs(a.Delta-d.Delta) > 1e-9 {
			t.Errorf("prompt %d delta: %f vs %f", i, a.Delta, d.Delta)
		}
	}
}

// ---------- helpers ----------

func mkReport(promptID, config, ts string, score float64) report.EvalReport {
	pass := score >= 0.5
	r := report.EvalReport{
		SchemaVersion: 2,
		PromptID:      promptID,
		ConfigName:    config,
		Timestamp:     ts,
		Success:       pass,
		GraderResults: []report.GraderResult{
			{
				GraderName: "correctness",
				GraderType: "prompt",
				Score:      score,
				Weight:     1.0,
				Pass:       pass,
				Points: []report.GraderPoint{
					{Label: "check", Pass: pass, Weight: 1.0},
				},
			},
		},
	}
	r.ScoreBreakdown = report.BuildScoreBreakdown(r.GraderResults)
	return r
}

func summariesEqual(a, b ComparisonSummary) bool {
	return a.Improved == b.Improved &&
		a.Regressed == b.Regressed &&
		a.Unchanged == b.Unchanged &&
		math.Abs(a.AvgDelta-b.AvgDelta) < 1e-9
}
