// Package comparison provides config-vs-config, run-vs-run, and temporal diff
// analysis for hyoka evaluation reports. It operates on EvalReport data and
// produces structured diffs suitable for display or further aggregation.
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
	PromptID    string      `json:"prompt_id"`
	ScoreA      float64     `json:"score_a"`
	ScoreB      float64     `json:"score_b"`
	Delta       float64     `json:"delta"` // B − A
	GraderDiffs []GraderDiff `json:"grader_diffs,omitempty"`
	OnlyInA     bool        `json:"only_in_a,omitempty"`
	OnlyInB     bool        `json:"only_in_b,omitempty"`
}

// ComparisonSummary aggregates high-level comparison metrics.
type ComparisonSummary struct {
	AvgDelta      float64      `json:"avg_delta"`
	Improved      int          `json:"improved"`
	Regressed     int          `json:"regressed"`
	Unchanged     int          `json:"unchanged"`
	TopImproved   []PromptDiff `json:"top_improved,omitempty"`
	TopRegressed  []PromptDiff `json:"top_regressed,omitempty"`
}

// ConfigComparison holds the result of comparing two configs across a shared prompt set.
type ConfigComparison struct {
	ConfigA   string            `json:"config_a"`
	ConfigB   string            `json:"config_b"`
	PerPrompt []PromptDiff      `json:"per_prompt"`
	Summary   ComparisonSummary `json:"summary"`
}

// RunComparison holds the result of comparing two specific runs.
type RunComparison struct {
	RunA      string            `json:"run_a"`
	RunB      string            `json:"run_b"`
	PerPrompt []PromptDiff      `json:"per_prompt"`
	Summary   ComparisonSummary `json:"summary"`
}

// TemporalComparison holds the result of comparing the latest run against a historical baseline.
type TemporalComparison struct {
	Config    string            `json:"config"`
	LatestRun string            `json:"latest_run"`
	BaseRun   string            `json:"base_run"`
	Since     time.Time         `json:"since"`
	PerPrompt []PromptDiff      `json:"per_prompt"`
	Summary   ComparisonSummary `json:"summary"`
}

// ---------- Public API ----------

// CompareConfigs loads all reports from reportsDir for the two given config
// names, matches them by prompt ID, and produces a ConfigComparison.
func CompareConfigs(reportsDir, configA, configB string) (*ConfigComparison, error) {
	lg := slog.With("role", "comparison")
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

	// Use latest report per prompt for each config.
	latestA := latestByPrompt(reportsA)
	latestB := latestByPrompt(reportsB)

	diffs := diffPromptMaps(latestA, latestB)
	summary := buildSummary(diffs)

	lg.Info("Config comparison complete",
		"config_a", configA, "config_b", configB,
		"prompts", len(diffs),
		"improved", summary.Improved,
		"regressed", summary.Regressed,
	)

	return &ConfigComparison{
		ConfigA:   configA,
		ConfigB:   configB,
		PerPrompt: diffs,
		Summary:   summary,
	}, nil
}

// CompareRuns loads all reports from two specific run directories and produces
// a RunComparison. Run IDs are the timestamp-based directory names under reportsDir.
func CompareRuns(reportsDir, runA, runB string) (*RunComparison, error) {
	lg := slog.With("role", "comparison")
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

	mapA := indexByPromptConfig(reportsA)
	mapB := indexByPromptConfig(reportsB)

	diffs := diffPromptMaps(mapA, mapB)
	summary := buildSummary(diffs)

	lg.Info("Run comparison complete",
		"run_a", runA, "run_b", runB,
		"prompts", len(diffs),
		"improved", summary.Improved,
		"regressed", summary.Regressed,
	)

	return &RunComparison{
		RunA:      runA,
		RunB:      runB,
		PerPrompt: diffs,
		Summary:   summary,
	}, nil
}

// TemporalDiff compares the latest run for a config against the most recent run
// at or before `since`. This enables "how has this config changed since date X?"
func TemporalDiff(reportsDir, config string, since time.Time) (*TemporalComparison, error) {
	lg := slog.With("role", "comparison")
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
		ts, err := time.Parse(time.RFC3339, r.Timestamp)
		if err != nil {
			// Try fallback format (some reports use a compact timestamp).
			ts, err = time.Parse("2006-01-02T15:04:05", r.Timestamp)
			if err != nil {
				continue
			}
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

	latestBase := latestByPrompt(base)
	latestRecent := latestByPrompt(recent)

	diffs := diffPromptMaps(latestBase, latestRecent)
	summary := buildSummary(diffs)

	// Determine run IDs from the partitioned reports.
	baseRun := runIDFromReport(base[len(base)-1], reportsDir)
	latestRun := runIDFromReport(recent[len(recent)-1], reportsDir)

	lg.Info("Temporal diff complete",
		"config", config, "since", since,
		"base_run", baseRun, "latest_run", latestRun,
		"prompts", len(diffs),
		"improved", summary.Improved,
		"regressed", summary.Regressed,
	)

	return &TemporalComparison{
		Config:    config,
		LatestRun: latestRun,
		BaseRun:   baseRun,
		Since:     since,
		PerPrompt: diffs,
		Summary:   summary,
	}, nil
}

// ---------- Internal helpers ----------

// loadReportsByConfig walks reportsDir and loads all report.json files
// whose ConfigName matches the given config.
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
// The returned map key is "{promptID}|{configName}" to uniquely identify entries.
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
// For run-vs-run comparison where both runs may have different configs.
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
	// Compute from grader results if present.
	if len(r.GraderResults) > 0 {
		sb := report.BuildScoreBreakdown(r.GraderResults)
		if sb != nil {
			return sb.FinalScore
		}
	}
	// Legacy review-based scoring.
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

	// Collect all grader names in deterministic order.
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
			gd.PassA = ga.Pass != nil && *ga.Pass
		}
		if gb, ok := indexB[name]; ok {
			gd.ScoreB = gb.Score
			gd.PassB = gb.Pass != nil && *gb.Pass
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
		// Strip config suffix from composite keys if present.
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

// buildSummary computes aggregate metrics and top N lists from prompt diffs.
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

	// Top improved — largest positive delta first.
	improved := filterPaired(diffs, func(d PromptDiff) bool { return d.Delta > unchangedThreshold })
	sort.Slice(improved, func(i, j int) bool { return improved[i].Delta > improved[j].Delta })
	s.TopImproved = take(improved, topN)

	// Top regressed — largest negative delta first (most negative = worst).
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

// findByPromptID finds the first entry in a map whose key starts with the given
// prompt ID (handling both bare keys and composite "promptID|config" keys).
func findByPromptID(m map[string]*report.EvalReport, promptID string) *report.EvalReport {
	// Direct lookup (simple key).
	if r, ok := m[promptID]; ok {
		return r
	}
	// Composite key scan.
	for k, r := range m {
		if promptIDFromKey(k) == promptID {
			return r
		}
	}
	return nil
}

// runIDFromReport attempts to extract the run ID from a report's timestamp.
func runIDFromReport(r report.EvalReport, _ string) string {
	ts, err := time.Parse(time.RFC3339, r.Timestamp)
	if err != nil {
		return r.Timestamp
	}
	return ts.Format("20060102-150405")
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


