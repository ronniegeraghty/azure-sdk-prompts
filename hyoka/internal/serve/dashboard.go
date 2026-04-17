// Package serve — dashboard API endpoints for the React SPA.
//
// These endpoints power comparison, drill-down, and trend views.
package serve

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/comparison"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/trends"
)

// registerDashboardRoutes adds dashboard-specific API routes to the mux.
// Comparison and trend handlers operate on aggregated data that we don't
// currently cache — they walk entire run directories — so cache wiring is
// limited to the per-run report loaders.
func registerDashboardRoutes(mux *http.ServeMux, opts Options, cache *fileCache) {
	mux.HandleFunc("/api/compare/configs", func(w http.ResponseWriter, r *http.Request) {
		handleAPICompareConfigs(w, r, opts.ReportsDir)
	})
	mux.HandleFunc("/api/compare/runs", func(w http.ResponseWriter, r *http.Request) {
		handleAPICompareRuns(w, r, opts.ReportsDir)
	})
	mux.HandleFunc("/api/compare/temporal", func(w http.ResponseWriter, r *http.Request) {
		handleAPICompareTemporal(w, r, opts.ReportsDir)
	})
	mux.HandleFunc("/api/trends", func(w http.ResponseWriter, r *http.Request) {
		handleAPITrends(w, r, opts.ReportsDir)
	})
	_ = cache // reserved for future per-report caching in dashboard endpoints
}

func handleAPIGraders(w http.ResponseWriter, _ *http.Request, reportsDir, runID string, cache *fileCache) {
	evals, err := loadRunEvals(reportsDir, runID, cache)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	type evalGraders struct {
		PromptID      string               `json:"prompt_id"`
		ConfigName    string               `json:"config_name"`
		GraderResults []report.GraderResult `json:"grader_results"`
	}
	var result []evalGraders
	for _, ev := range evals {
		result = append(result, evalGraders{
			PromptID:      ev.PromptID,
			ConfigName:    ev.ConfigName,
			GraderResults: ev.GraderResults,
		})
	}
	writeJSON(w, result)
}

func handleAPITimeline(w http.ResponseWriter, _ *http.Request, reportsDir, runID string, cache *fileCache) {
	evals, err := loadRunEvals(reportsDir, runID, cache)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	type evalTimeline struct {
		PromptID       string                      `json:"prompt_id"`
		ConfigName     string                      `json:"config_name"`
		ActionTimeline *report.ActionTimelineReport `json:"action_timeline"`
	}
	var result []evalTimeline
	for _, ev := range evals {
		tl := ev.ActionTimeline
		if tl == nil && len(ev.SessionEvents) > 0 {
			tl = report.BuildActionTimeline(ev.SessionEvents)
		}
		result = append(result, evalTimeline{
			PromptID:       ev.PromptID,
			ConfigName:     ev.ConfigName,
			ActionTimeline: tl,
		})
	}
	writeJSON(w, result)
}

func handleAPIScoreBreakdown(w http.ResponseWriter, _ *http.Request, reportsDir, runID string, cache *fileCache) {
	evals, err := loadRunEvals(reportsDir, runID, cache)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	type evalBreakdown struct {
		PromptID       string                 `json:"prompt_id"`
		ConfigName     string                 `json:"config_name"`
		ScoreBreakdown *report.ScoreBreakdown `json:"score_breakdown"`
	}
	var result []evalBreakdown
	for _, ev := range evals {
		sb := ev.ScoreBreakdown
		if sb == nil && len(ev.GraderResults) > 0 {
			sb = report.BuildScoreBreakdown(ev.GraderResults)
		}
		result = append(result, evalBreakdown{
			PromptID:       ev.PromptID,
			ConfigName:     ev.ConfigName,
			ScoreBreakdown: sb,
		})
	}
	writeJSON(w, result)
}

func handleAPICompareConfigs(w http.ResponseWriter, r *http.Request, reportsDir string) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	if a == "" || b == "" {
		http.Error(w, `both "a" and "b" query parameters are required`, http.StatusBadRequest)
		return
	}
	cmp, err := comparison.CompareConfigs(reportsDir, a, b)
	if err != nil {
		slog.Error("comparing configs", "error", err)
		http.Error(w, fmt.Sprintf("comparison failed: %v", err), http.StatusInternalServerError)
		return
	}
	writeJSON(w, cmp)
}

func handleAPICompareRuns(w http.ResponseWriter, r *http.Request, reportsDir string) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	if a == "" || b == "" {
		http.Error(w, `both "a" and "b" query parameters are required`, http.StatusBadRequest)
		return
	}
	if strings.Contains(a, "..") || strings.Contains(b, "..") {
		http.Error(w, "invalid run ID", http.StatusBadRequest)
		return
	}
	cmp, err := comparison.CompareRuns(reportsDir, a, b)
	if err != nil {
		slog.Error("comparing runs", "error", err)
		http.Error(w, fmt.Sprintf("comparison failed: %v", err), http.StatusInternalServerError)
		return
	}
	writeJSON(w, cmp)
}

func handleAPICompareTemporal(w http.ResponseWriter, r *http.Request, reportsDir string) {
	config := r.URL.Query().Get("config")
	sinceStr := r.URL.Query().Get("since")
	if config == "" || sinceStr == "" {
		http.Error(w, `both "config" and "since" query parameters are required`, http.StatusBadRequest)
		return
	}
	since, err := time.Parse(time.RFC3339, sinceStr)
	if err != nil {
		since, err = time.Parse("2006-01-02", sinceStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid since format (expected RFC3339 or YYYY-MM-DD): %v", err), http.StatusBadRequest)
			return
		}
	}
	cmp, err := comparison.TemporalDiff(reportsDir, config, since)
	if err != nil {
		slog.Error("temporal comparison", "error", err)
		http.Error(w, fmt.Sprintf("comparison failed: %v", err), http.StatusInternalServerError)
		return
	}
	writeJSON(w, cmp)
}

// handleAPIRunComparisons returns the auto-generated comparisons for a run
// (written by comparison.WriteForRun at end-of-run). Reads {runDir}/comparisons.json
// if present; on miss, computes on-demand by loading the run's reports so the
// site always gets a result even for runs that pre-date auto-generation.
func handleAPIRunComparisons(w http.ResponseWriter, _ *http.Request, reportsDir, runID string) {
	if strings.Contains(runID, "..") || runID == "" {
		http.Error(w, "invalid run ID", http.StatusBadRequest)
		return
	}
	runDir := filepath.Join(reportsDir, runID)
	if info, err := os.Stat(runDir); err != nil || !info.IsDir() {
		http.Error(w, "run not found", http.StatusNotFound)
		return
	}

	results, err := comparison.LoadForRun(runDir)
	if err != nil {
		slog.Error("loading run comparisons", "error", err, "run", runID)
		http.Error(w, "failed to load comparisons", http.StatusInternalServerError)
		return
	}
	if results == nil {
		// Legacy run without comparisons.json — compute on demand using the
		// same engine so the response matches what end-of-run would have
		// produced.
		reports, err := loadRunReports(runDir)
		if err != nil {
			slog.Error("loading run reports", "error", err, "run", runID)
			http.Error(w, "failed to load reports", http.StatusInternalServerError)
			return
		}
		results = comparison.AutoGenerateForRun(reports)
	}
	if results == nil {
		results = []comparison.ComparisonResult{}
	}
	writeJSON(w, results)
}

// loadRunReports walks a run directory and decodes every report.json it finds.
func loadRunReports(runDir string) ([]report.EvalReport, error) {
	var out []report.EvalReport
	err := filepath.Walk(runDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() != "report.json" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		var r report.EvalReport
		if err := json.Unmarshal(data, &r); err == nil {
			out = append(out, r)
		}
		return nil
	})
	return out, err
}

func handleAPITrends(w http.ResponseWriter, r *http.Request, reportsDir string) {
	q := r.URL.Query()
	opts := trends.TrendOptions{
		ReportsDir: reportsDir,
		PromptID:   q.Get("promptId"),
		Service:    q.Get("service"),
		Language:   q.Get("language"),
	}
	if opts.PromptID == "" && opts.Service == "" && opts.Language == "" && q.Get("config") == "" {
		http.Error(w, `at least one filter parameter is required (promptId, service, language, or config)`, http.StatusBadRequest)
		return
	}
	tr, err := trends.Generate(opts)
	if err != nil {
		slog.Error("generating trends", "error", err)
		http.Error(w, "failed to generate trends", http.StatusInternalServerError)
		return
	}
	if configFilter := q.Get("config"); configFilter != "" && tr != nil {
		for i := range tr.PromptTrends {
			filtered := make(map[string][]trends.RunResult)
			for k, v := range tr.PromptTrends[i].Configs {
				if k == configFilter {
					filtered[k] = v
				}
			}
			tr.PromptTrends[i].Configs = filtered
		}
	}
	writeJSON(w, tr)
}

func handleAPIPromptHistory(w http.ResponseWriter, r *http.Request, reportsDir string) {
	rest := strings.TrimPrefix(r.URL.Path, "/api/prompts/")
	promptID := strings.TrimSuffix(rest, "/history")
	if promptID == "" || strings.Contains(promptID, "..") {
		http.NotFound(w, r)
		return
	}
	tr, err := trends.Generate(trends.TrendOptions{
		ReportsDir: reportsDir,
		PromptID:   promptID,
	})
	if err != nil {
		slog.Error("generating prompt history", "error", err)
		http.Error(w, "failed to load prompt history", http.StatusInternalServerError)
		return
	}
	if tr == nil || len(tr.Entries) == 0 {
		writeJSON(w, []struct{}{})
		return
	}
	writeJSON(w, tr)
}

// loadRunEvals walks a run directory and returns parsed EvalReport values for
// every report.json it finds. Report bytes are served through the cache so
// repeated requests for the same run (graders, timeline, score-breakdown all
// call this) are a memory hit after the first read.
func loadRunEvals(reportsDir, runID string, cache *fileCache) ([]*report.EvalReport, error) {
	runDir := filepath.Join(reportsDir, runID)
	if _, err := os.Stat(runDir); err != nil {
		return nil, err
	}
	var evals []*report.EvalReport
	err := filepath.Walk(runDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() || info.Name() != "report.json" {
			return nil
		}
		data, readErr := cache.ReadFile(path)
		if readErr != nil {
			return nil
		}
		var ev report.EvalReport
		if json.Unmarshal(data, &ev) != nil {
			return nil
		}
		evals = append(evals, &ev)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return evals, nil
}
