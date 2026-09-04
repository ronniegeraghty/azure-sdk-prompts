// Package comparison provides config-vs-config, run-vs-run, and temporal diff
// analysis for hyoka evaluation reports. It is the single source of truth for
// comparison logic: the `hyoka compare` CLI, the serve API, and report
// auto-generation all share this engine. See ComparisonResult (result.go) for
// the canonical payload.
package comparison

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// ---------- Types ----------

// GraderDiff holds the per-grader score delta between two evaluations.
type GraderDiff struct {
	Name   string  `json:"name"`
	ScoreA float64 `json:"score_a"`
	ScoreB float64 `json:"score_b"`
	Delta  float64 `json:"delta"` // B − A
	PassA  bool    `json:"pass_a"`
	PassB  bool    `json:"pass_b"`
}

// PromptDiff holds the per-prompt aggregate delta between two evaluations.
type PromptDiff struct {
	PromptID    string       `json:"prompt_id"`
	ScoreA      float64      `json:"score_a"`
	ScoreB      float64      `json:"score_b"`
	Delta       float64      `json:"delta"` // B − A
	GraderDiffs []GraderDiff `json:"grader_diffs,omitempty"`
	OnlyInA     bool         `json:"only_in_a,omitempty"`
	OnlyInB     bool         `json:"only_in_b,omitempty"`
}

// ComparisonSummary aggregates high-level comparison metrics.
type ComparisonSummary struct {
	AvgDelta     float64      `json:"avg_delta"`
	Improved     int          `json:"improved"`
	Regressed    int          `json:"regressed"`
	Unchanged    int          `json:"unchanged"`
	TopImproved  []PromptDiff `json:"top_improved,omitempty"`
	TopRegressed []PromptDiff `json:"top_regressed,omitempty"`
}

// ---------- Public API (disk-backed) ----------

// CompareConfigs loads all reports from reportsDir for the two given config
// names, matches them by prompt ID, and produces a ComparisonResult.
// For each config, the latest report per prompt is used.
func CompareConfigs(reportsDir, configA, configB string) (*ComparisonResult, error) {
	lg := slog.With("role", "comparison", "kind", "configs")
	lg.Debug("Comparing configs", "config_a", configA, "config_b", configB)

	reportsA, err := loadReportsByConfig(reportsDir, configA)
	if err != nil {
		return nil, fmt.Errorf("loading config %s: %w", configA, err)
	}
	reportsB, err := loadReportsByConfig(reportsDir, configB)
	if err != nil {
		return nil, fmt.Errorf("loading config %s: %w", configB, err)
	}

	if len(reportsA) == 0 && len(reportsB) == 0 {
		return nil, fmt.Errorf("no reports found for either config %q or %q", configA, configB)
	}

	result := CompareReports(KindConfigs, configA, configB, reportsA, reportsB)

	lg.Info("Config comparison complete",
		"config_a", configA, "config_b", configB,
		"prompts", len(result.PerPrompt),
		"improved", result.Summary.Improved,
		"regressed", result.Summary.Regressed,
	)
	return result, nil
}

// CompareRuns loads all reports from two specific run directories and produces
// a ComparisonResult. Run IDs are the timestamp-based directory names under
// reportsDir.
func CompareRuns(reportsDir, runA, runB string) (*ComparisonResult, error) {
	lg := slog.With("role", "comparison", "kind", "runs")
	lg.Debug("Comparing runs", "run_a", runA, "run_b", runB)

	reportsA, err := loadReportsFromRun(reportsDir, runA)
	if err != nil {
		return nil, fmt.Errorf("loading run %s: %w", runA, err)
	}
	reportsB, err := loadReportsFromRun(reportsDir, runB)
	if err != nil {
		return nil, fmt.Errorf("loading run %s: %w", runB, err)
	}

	if len(reportsA) == 0 && len(reportsB) == 0 {
		return nil, fmt.Errorf("no reports found for either run %q or %q", runA, runB)
	}

	result := CompareReports(KindRuns, runA, runB, reportsA, reportsB)

	lg.Info("Run comparison complete",
		"run_a", runA, "run_b", runB,
		"prompts", len(result.PerPrompt),
		"improved", result.Summary.Improved,
		"regressed", result.Summary.Regressed,
	)
	return result, nil
}

// TemporalDiff compares the latest run for a config against the most recent
// run at or before `since`. Produces a ComparisonResult with Kind == KindTemporal
// and Config/Since populated.
func TemporalDiff(reportsDir, config string, since time.Time) (*ComparisonResult, error) {
	lg := slog.With("role", "comparison", "kind", "temporal")
	lg.Debug("Temporal diff", "config", config, "since", since)

	all, err := loadReportsByConfig(reportsDir, config)
	if err != nil {
		return nil, fmt.Errorf("loading config %s: %w", config, err)
	}
	if len(all) == 0 {
		return nil, fmt.Errorf("no reports found for config %q", config)
	}

	// Partition into base (≤ since) and recent (> since).
	var base, recent []report.EvalReport
	for _, r := range all {
		ts, err := parseReportTimestamp(r.Timestamp)
		if err != nil {
			continue
		}
		if ts.Before(since) || ts.Equal(since) {
			base = append(base, r)
		} else {
			recent = append(recent, r)
		}
	}

	if len(base) == 0 {
		return nil, fmt.Errorf("no reports found for config %q at or before %s", config, since.Format(time.RFC3339))
	}
	if len(recent) == 0 {
		return nil, fmt.Errorf("no reports found for config %q after %s", config, since.Format(time.RFC3339))
	}

	baseRun := runIDFromReport(base[len(base)-1])
	latestRun := runIDFromReport(recent[len(recent)-1])

	result := CompareReports(KindTemporal, baseRun, latestRun, base, recent)
	result.Config = config
	sinceCopy := since
	result.Since = &sinceCopy

	lg.Info("Temporal diff complete",
		"config", config, "since", since,
		"base_run", baseRun, "latest_run", latestRun,
		"prompts", len(result.PerPrompt),
		"improved", result.Summary.Improved,
		"regressed", result.Summary.Regressed,
	)
	return result, nil
}

// ---------- Public API (in-memory) ----------

// CompareReports is the core in-memory comparison primitive. Given two slices
// of EvalReport values, it produces a ComparisonResult. This is the single
// function used by every comparison surface:
//
//   - CompareConfigs / CompareRuns / TemporalDiff call it after loading from disk.
//   - AutoGenerateForRun calls it to produce pairwise comparisons during
//     multi-config runs (attached to RunSummary.Comparisons).
//   - Callers who already have reports in memory can call it directly.
//
// For configs comparison, reports are keyed by prompt ID (latest wins). For
// runs comparison, reports are keyed by prompt+config composite key so
// multiple configs within a single run each produce their own entry.
func CompareReports(kind ComparisonKind, labelA, labelB string, reportsA, reportsB []report.EvalReport) *ComparisonResult {
	var mapA, mapB map[string]*report.EvalReport
	if kind == KindRuns {
		mapA = indexByPromptConfig(reportsA)
		mapB = indexByPromptConfig(reportsB)
	} else {
		mapA = latestByPrompt(reportsA)
		mapB = latestByPrompt(reportsB)
	}

	diffs := diffPromptMaps(mapA, mapB)
	summary := buildSummary(diffs)

	return &ComparisonResult{
		Kind:      kind,
		LabelA:    labelA,
		LabelB:    labelB,
		PerPrompt: diffs,
		Summary:   summary,
	}
}

// AutoGenerateForRun produces pairwise config comparisons for all config pairs
// present in the given reports. This is the entry point called at the end of a
// multi-config run to attach comparisons to the run summary.
//
// Returns nil when fewer than two distinct configs are present (nothing to
// compare). Pairs are emitted in deterministic alphabetical order (A<B).
func AutoGenerateForRun(reports []report.EvalReport) []ComparisonResult {
	// Group reports by config name.
	byConfig := make(map[string][]report.EvalReport)
	for _, r := range reports {
		byConfig[r.ConfigName] = append(byConfig[r.ConfigName], r)
	}
	if len(byConfig) < 2 {
		return nil
	}

	configs := make([]string, 0, len(byConfig))
	for c := range byConfig {
		configs = append(configs, c)
	}
	sort.Strings(configs)

	var out []ComparisonResult
	for i := 0; i < len(configs); i++ {
		for j := i + 1; j < len(configs); j++ {
			a, b := configs[i], configs[j]
			res := CompareReports(KindConfigs, a, b, byConfig[a], byConfig[b])
			out = append(out, *res)
		}
	}
	return out
}

// ---------- Internal helpers ----------

// loadReportsByConfig walks reportsDir and loads all report.json files whose
// ConfigName matches the given config.
func loadReportsByConfig(reportsDir, config string) ([]report.EvalReport, error) {
	var reports []report.EvalReport
	err := filepath.Walk(reportsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() != "report.json" {
			return nil
		}
		r, err := loadReport(path)
		if err != nil {
			return nil
		}
		if r.ConfigName == config {
			reports = append(reports, *r)
		}
		return nil
	})
	return reports, err
}

// loadReportsFromRun loads all report.json files from a specific run directory.
func loadReportsFromRun(reportsDir, runID string) ([]report.EvalReport, error) {
	runDir := filepath.Join(reportsDir, runID)
	info, err := os.Stat(runDir)
	if err != nil {
		return nil, fmt.Errorf("run directory %q not found: %w", runID, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%q is not a directory", runID)
	}

	var reports []report.EvalReport
	err = filepath.Walk(runDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() != "report.json" {
			return nil
		}
		r, err := loadReport(path)
		if err != nil {
			return nil
		}
		reports = append(reports, *r)
		return nil
	})
	return reports, err
}

// loadReport reads and decodes a single report.json file.
func loadReport(path string) (*report.EvalReport, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var r report.EvalReport
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// latestByPrompt selects the most recent report per prompt ID (by timestamp).
func latestByPrompt(reports []report.EvalReport) map[string]*report.EvalReport {
	result := make(map[string]*report.EvalReport)
	for i := range reports {
		r := &reports[i]
		key := r.PromptID
		existing, ok := result[key]
		if !ok || r.Timestamp > existing.Timestamp {
			result[key] = r
		}
	}
	return result
}

// indexByPromptConfig indexes reports by a composite key of promptID + configName.
// For run-vs-run comparison where both runs may span multiple configs.
func indexByPromptConfig(reports []report.EvalReport) map[string]*report.EvalReport {
	result := make(map[string]*report.EvalReport)
	for i := range reports {
		r := &reports[i]
		key := r.PromptID + "|" + r.ConfigName
		existing, ok := result[key]
		if !ok || r.Timestamp > existing.Timestamp {
			result[key] = r
		}
	}
	return result
}

// scoreFromReport computes a normalized 0.0–1.0 score from a report.
// It prefers the ScoreBreakdown.FinalScore (grader-based), falls back to
// Review.OverallScore/MaxScore (legacy), then 0.0.
func scoreFromReport(r *report.EvalReport) float64 {
	if r.ScoreBreakdown != nil {
		return r.ScoreBreakdown.FinalScore
	}
	if len(r.GraderResults) > 0 {
		sb := report.BuildScoreBreakdown(r.GraderResults)
		if sb != nil {
			return sb.FinalScore
		}
	}
	if r.Review != nil && r.Review.MaxScore > 0 {
		return float64(r.Review.OverallScore) / float64(r.Review.MaxScore)
	}
	return 0.0
}

// graderDiffs computes per-grader deltas between two reports.
func graderDiffs(a, b *report.EvalReport) []GraderDiff {
	indexA := make(map[string]*report.GraderResult)
	for i := range a.GraderResults {
		indexA[a.GraderResults[i].GraderName] = &a.GraderResults[i]
	}
	indexB := make(map[string]*report.GraderResult)
	for i := range b.GraderResults {
		indexB[b.GraderResults[i].GraderName] = &b.GraderResults[i]
	}

	// Deterministic ordering: graders from A first (in order), then new ones from B.
	seen := make(map[string]bool)
	var names []string
	for _, g := range a.GraderResults {
		if !seen[g.GraderName] {
			seen[g.GraderName] = true
			names = append(names, g.GraderName)
		}
	}
	for _, g := range b.GraderResults {
		if !seen[g.GraderName] {
			seen[g.GraderName] = true
			names = append(names, g.GraderName)
		}
	}

	var diffs []GraderDiff
	for _, name := range names {
		gd := GraderDiff{Name: name}
		if ga, ok := indexA[name]; ok {
			gd.ScoreA = ga.Score
			gd.PassA = ga.Pass
		}
		if gb, ok := indexB[name]; ok {
			gd.ScoreB = gb.Score
			gd.PassB = gb.Pass
		}
		gd.Delta = gd.ScoreB - gd.ScoreA
		diffs = append(diffs, gd)
	}
	return diffs
}

// diffPromptMaps produces PromptDiff entries from two prompt-keyed maps.
func diffPromptMaps(mapA, mapB map[string]*report.EvalReport) []PromptDiff {
	seen := make(map[string]bool)
	var keys []string
	for k := range mapA {
		pid := promptIDFromKey(k)
		if !seen[pid] {
			seen[pid] = true
			keys = append(keys, pid)
		}
	}
	for k := range mapB {
		pid := promptIDFromKey(k)
		if !seen[pid] {
			seen[pid] = true
			keys = append(keys, pid)
		}
	}
	sort.Strings(keys)

	var diffs []PromptDiff
	for _, pid := range keys {
		rA := findByPromptID(mapA, pid)
		rB := findByPromptID(mapB, pid)

		pd := PromptDiff{PromptID: pid}
		switch {
		case rA != nil && rB != nil:
			pd.ScoreA = scoreFromReport(rA)
			pd.ScoreB = scoreFromReport(rB)
			pd.Delta = pd.ScoreB - pd.ScoreA
			pd.GraderDiffs = graderDiffs(rA, rB)
		case rA != nil:
			pd.ScoreA = scoreFromReport(rA)
			pd.OnlyInA = true
		case rB != nil:
			pd.ScoreB = scoreFromReport(rB)
			pd.OnlyInB = true
		}
		diffs = append(diffs, pd)
	}
	return diffs
}

// buildSummary computes aggregate metrics and top-N lists from prompt diffs.
func buildSummary(diffs []PromptDiff) ComparisonSummary {
	const topN = 5
	const unchangedThreshold = 0.001

	var s ComparisonSummary
	var totalDelta float64
	var paired int

	for _, d := range diffs {
		if d.OnlyInA || d.OnlyInB {
			continue
		}
		paired++
		totalDelta += d.Delta
		switch {
		case d.Delta > unchangedThreshold:
			s.Improved++
		case d.Delta < -unchangedThreshold:
			s.Regressed++
		default:
			s.Unchanged++
		}
	}

	if paired > 0 {
		s.AvgDelta = totalDelta / float64(paired)
	}

	improved := filterPaired(diffs, func(d PromptDiff) bool { return d.Delta > unchangedThreshold })
	sort.Slice(improved, func(i, j int) bool { return improved[i].Delta > improved[j].Delta })
	s.TopImproved = take(improved, topN)

	regressed := filterPaired(diffs, func(d PromptDiff) bool { return d.Delta < -unchangedThreshold })
	sort.Slice(regressed, func(i, j int) bool { return regressed[i].Delta < regressed[j].Delta })
	s.TopRegressed = take(regressed, topN)

	return s
}

// promptIDFromKey strips a composite "promptID|config" key to just the promptID.
func promptIDFromKey(key string) string {
	if idx := strings.Index(key, "|"); idx >= 0 {
		return key[:idx]
	}
	return key
}

// findByPromptID finds the first entry in a map whose key starts with the
// given prompt ID (handles both bare keys and composite keys).
func findByPromptID(m map[string]*report.EvalReport, promptID string) *report.EvalReport {
	if r, ok := m[promptID]; ok {
		return r
	}
	for k, r := range m {
		if promptIDFromKey(k) == promptID {
			return r
		}
	}
	return nil
}

// runIDFromReport formats the report's timestamp as a run ID.
func runIDFromReport(r report.EvalReport) string {
	ts, err := parseReportTimestamp(r.Timestamp)
	if err != nil {
		return r.Timestamp
	}
	return ts.Format("20060102-150405")
}

// parseReportTimestamp parses a report timestamp in either RFC3339 or compact
// form.
func parseReportTimestamp(ts string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, ts); err == nil {
		return t, nil
	}
	return time.Parse("2006-01-02T15:04:05", ts)
}

// filterPaired returns only diffs that appear in both sides.
func filterPaired(diffs []PromptDiff, pred func(PromptDiff) bool) []PromptDiff {
	var out []PromptDiff
	for _, d := range diffs {
		if d.OnlyInA || d.OnlyInB {
			continue
		}
		if pred(d) {
			out = append(out, d)
		}
	}
	return out
}

// take returns up to n elements from the slice.
func take(s []PromptDiff, n int) []PromptDiff {
	if len(s) <= n {
		return s
	}
	return s[:n]
}
