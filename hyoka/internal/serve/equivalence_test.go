package serve

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/comparison"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// TestCLISiteEquivalence_AutoGenComparisons is the Phase 4 gate criterion
// enforcement: the `hyoka compare` CLI and the site comparison page MUST
// produce identical results for the same inputs. Both paths funnel through
// comparison.CompareReports / comparison.AutoGenerateForRun — this test
// verifies that contract holds end-to-end by comparing:
//
//  1. Direct call to comparison.AutoGenerateForRun(reports) (the engine used
//     by the CLI and the auto-generator at end-of-run).
//  2. JSON payload returned by GET /api/runs/{runID}/comparisons (the path
//     the site frontend consumes).
//
// If Neo's claim that equivalence is "satisfied by construction" via the
// shared core, these two MUST be byte-identical for the same reports.
func TestCLISiteEquivalence_AutoGenComparisons(t *testing.T) {
	// Arrange: build a multi-config run so AutoGenerateForRun has something
	// to emit (requires ≥2 distinct configs).
	dir := t.TempDir()
	runID := "20260418-000000"
	runDir := filepath.Join(dir, runID)
	if err := os.MkdirAll(runDir, 0755); err != nil {
		t.Fatalf("mkdir runDir: %v", err)
	}

	pass := true
	mkReport := func(promptID, config string, score float64) report.EvalReport {
		r := report.EvalReport{
			SchemaVersion: 2,
			PromptID:      promptID,
			ConfigName:    config,
			Timestamp:     "2026-04-18T00:00:00Z",
			Success:       true,
			GraderResults: []report.GraderResult{
				{GraderName: "correctness", GraderType: "prompt", Score: score, Weight: 1.0, Pass: pass, Points: []report.GraderPoint{{Label: "check", Pass: pass, Weight: 1.0}}},
			},
		}
		r.ScoreBreakdown = report.BuildScoreBreakdown(r.GraderResults)
		return r
	}

	// Two configs × two prompts, with varied scores so the summary has
	// non-trivial improved/regressed counts.
	reports := []report.EvalReport{
		mkReport("prompt-alpha", "baseline/opus", 0.6),
		mkReport("prompt-beta", "baseline/opus", 0.8),
		mkReport("prompt-alpha", "azure-mcp/opus", 0.9),
		mkReport("prompt-beta", "azure-mcp/opus", 0.7),
	}

	// Write each report to disk under the run directory so the serve
	// handler's on-demand path (loadRunReports → AutoGenerateForRun) can
	// discover them. We deliberately do NOT write comparisons.json — that
	// forces the API to recompute via the shared engine, which is the
	// strictest form of equivalence test.
	for _, r := range reports {
		cfgDir := filepath.Join(runDir, "results", r.PromptID, filepath.FromSlash(r.ConfigName))
		if err := os.MkdirAll(cfgDir, 0755); err != nil {
			t.Fatalf("mkdir cfgDir: %v", err)
		}
		data, err := json.Marshal(r)
		if err != nil {
			t.Fatalf("marshal report: %v", err)
		}
		if err := os.WriteFile(filepath.Join(cfgDir, "report.json"), data, 0644); err != nil {
			t.Fatalf("write report: %v", err)
		}
	}

	// Path A — direct engine call (CLI path).
	cliResults := comparison.AutoGenerateForRun(reports)
	if len(cliResults) == 0 {
		t.Fatal("AutoGenerateForRun returned no results for 2-config input")
	}

	// Path B — HTTP API (site path).
	mux := buildMux(Options{ReportsDir: dir})
	req := httptest.NewRequest("GET", "/api/runs/"+runID+"/comparisons", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("GET comparisons: got %d, body=%s", rec.Code, rec.Body.String())
	}
	var apiResults []comparison.ComparisonResult
	if err := json.Unmarshal(rec.Body.Bytes(), &apiResults); err != nil {
		t.Fatalf("decode API response: %v", err)
	}

	// Assert: CLI engine output and API JSON roundtrip must be deep-equal.
	// We roundtrip the CLI slice through JSON too so any omit/marshal
	// asymmetry (e.g. Since pointers, omitempty) is normalized on both
	// sides — equivalence is measured at the wire-format level, which is
	// what the site actually consumes.
	cliJSON, err := json.Marshal(cliResults)
	if err != nil {
		t.Fatalf("marshal CLI results: %v", err)
	}
	var cliRoundtripped []comparison.ComparisonResult
	if err := json.Unmarshal(cliJSON, &cliRoundtripped); err != nil {
		t.Fatalf("unmarshal CLI results: %v", err)
	}

	if !reflect.DeepEqual(cliRoundtripped, apiResults) {
		t.Errorf("CLI and site produced divergent comparison results\nCLI:  %s\nSITE: %s",
			string(cliJSON), rec.Body.String())
	}

	// Also assert the list is non-empty and covers the single config pair
	// we seeded, so a silent return of [] doesn't trivially pass.
	if len(apiResults) != 1 {
		t.Errorf("expected 1 pairwise comparison (2 configs → 1 pair), got %d", len(apiResults))
	}
}

// TestCLISiteEquivalence_LoadedComparisonsFile verifies the on-disk path:
// when comparisons.json IS present (auto-generated at end-of-run), the API
// returns exactly what was written, which is exactly what the CLI engine
// produced. This closes the other half of the equivalence contract.
func TestCLISiteEquivalence_LoadedComparisonsFile(t *testing.T) {
	dir := t.TempDir()
	runID := "20260418-000001"
	runDir := filepath.Join(dir, runID)
	if err := os.MkdirAll(runDir, 0755); err != nil {
		t.Fatalf("mkdir runDir: %v", err)
	}

	pass := true
	mk := func(promptID, config string, score float64) report.EvalReport {
		r := report.EvalReport{
			SchemaVersion: 2, PromptID: promptID, ConfigName: config,
			Timestamp: "2026-04-18T00:00:00Z", Success: true,
			GraderResults: []report.GraderResult{
				{GraderName: "g", GraderType: "prompt", Score: score, Weight: 1.0, Pass: pass, Points: []report.GraderPoint{{Label: "check", Pass: pass, Weight: 1.0}}},
			},
		}
		r.ScoreBreakdown = report.BuildScoreBreakdown(r.GraderResults)
		return r
	}
	reports := []report.EvalReport{
		mk("p1", "cfg-a", 0.5), mk("p1", "cfg-b", 0.9),
		mk("p2", "cfg-a", 0.7), mk("p2", "cfg-b", 0.7),
	}

	// Write comparisons.json via the same writer used at end-of-run.
	if _, err := comparison.WriteForRun(runDir, reports); err != nil {
		t.Fatalf("WriteForRun: %v", err)
	}

	// CLI-equivalent: what the engine would emit right now.
	cliResults := comparison.AutoGenerateForRun(reports)

	// Site path: GET the API.
	mux := buildMux(Options{ReportsDir: dir})
	req := httptest.NewRequest("GET", "/api/runs/"+runID+"/comparisons", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("GET: %d body=%s", rec.Code, rec.Body.String())
	}
	var apiResults []comparison.ComparisonResult
	if err := json.Unmarshal(rec.Body.Bytes(), &apiResults); err != nil {
		t.Fatalf("decode: %v", err)
	}

	cliJSON, _ := json.Marshal(cliResults)
	var cliRT []comparison.ComparisonResult
	_ = json.Unmarshal(cliJSON, &cliRT)

	if !reflect.DeepEqual(cliRT, apiResults) {
		t.Errorf("on-disk comparisons diverged from engine output\nCLI:  %s\nSITE: %s",
			string(cliJSON), rec.Body.String())
	}
}
