// Package trends provides analysis of evaluation trends over time.
// Package trends provides analysis of evaluation trends over time.
package trends

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

// TrendClassification indicates the overall trend direction.
type TrendClassification string

const (
TrendStable     TrendClassification = "stable"
TrendImproving  TrendClassification = "improving"
TrendRegressing TrendClassification = "regressing"
TrendFlaky      TrendClassification = "flaky"
TrendNew        TrendClassification = "new"
)

// TrendEntry holds data from a single historical evaluation run.
type TrendEntry struct {
RunID        string            `json:"run_id"`
Timestamp    string            `json:"timestamp"`
ConfigName   string            `json:"config_name"`
PromptID     string            `json:"prompt_id"`
Success      bool              `json:"success"`
Duration     float64           `json:"duration_seconds"`
Score        int               `json:"score"`
MaxScore     int               `json:"max_score"`
HasReview    bool              `json:"has_review"`
ToolCalls    []string          `json:"tool_calls"`
FileCount    int               `json:"file_count"`
Error        string            `json:"error,omitempty"`
Properties   map[string]string `json:"properties,omitempty"`
GraderScores []GraderScore     `json:"grader_scores,omitempty"`
}

// GraderScore holds a single grader's normalized score for trend tracking.
type GraderScore struct {
GraderName string  `json:"grader_name"`
Score      float64 `json:"score"`     // 0.0–1.0 normalized
MaxScore   int     `json:"max_score"` // integer max from the grader
}

// RunResult holds a single run's metrics for a specific prompt+config.
type RunResult struct {
RunID     string   `json:"run_id"`
Timestamp string   `json:"timestamp"`
Success   bool     `json:"success"`
Duration  float64  `json:"duration_seconds"`
FileCount int      `json:"file_count"`
Score     int      `json:"score"`
MaxScore  int      `json:"max_score"`
HasReview bool     `json:"has_review"`
ToolCalls []string `json:"tool_calls,omitempty"`
Error     string   `json:"error,omitempty"`
}

// PromptTrend holds time-series performance data for a single prompt.
type PromptTrend struct {
PromptID string                         `json:"prompt_id"`
Configs  map[string][]RunResult         `json:"configs"`
Trend    map[string]TrendClassification `json:"trend"`
}

// TrendReport summarizes historical trends for a set of prompts.
type TrendReport struct {
PromptID     string        `json:"prompt_id,omitempty"`
Service      string        `json:"service,omitempty"`
Language     string        `json:"language,omitempty"`
TotalRuns    int           `json:"total_runs"`
Entries      []TrendEntry  `json:"entries"`
PromptTrends []PromptTrend `json:"prompt_trends"`
RunIDs       []string      `json:"run_ids"`
Analysis     string        `json:"analysis,omitempty"`
GeneratedAt  string        `json:"generated_at"`
}

// TrendOptions configures trend report generation.
type TrendOptions struct {
ReportsDir string
PromptID   string
Service    string
Language   string
OutputDir  string
Analyze    bool
}

// Generate scans historical reports and produces a trend report.
func Generate(opts TrendOptions) (*TrendReport, error) {
lg := slog.With("role", "trend-analysis")
lg.Debug("Scanning reports for trends", "reports_dir", opts.ReportsDir, "prompt_id", opts.PromptID, "service", opts.Service, "language", opts.Language)
entries, err := scanReports(opts.ReportsDir, opts.PromptID, opts.Service, opts.Language)
if err != nil {
return nil, fmt.Errorf("scanning reports: %w", err)
}

lg.Debug("Trend entries found", "count", len(entries))

sort.Slice(entries, func(i, j int) bool {
return entries[i].Timestamp < entries[j].Timestamp
})

promptTrends, runIDs := buildPromptTrends(entries)

tr := &TrendReport{
PromptID:     opts.PromptID,
Service:      opts.Service,
Language:      opts.Language,
TotalRuns:    len(entries),
Entries:      entries,
PromptTrends: promptTrends,
RunIDs:       runIDs,
GeneratedAt:  time.Now().UTC().Format(time.RFC3339),
}

lg.Info("Trend report generated", "total_runs", tr.TotalRuns, "prompts", len(promptTrends), "run_ids", len(runIDs))

return tr, nil
}

// buildPromptTrends groups entries by prompt → config → chronological runs.
func buildPromptTrends(entries []TrendEntry) ([]PromptTrend, []string) {
// Collect unique run IDs in order
runIDSet := map[string]bool{}
var runIDs []string
for _, e := range entries {
if !runIDSet[e.RunID] {
runIDSet[e.RunID] = true
runIDs = append(runIDs, e.RunID)
}
}

// Group by prompt → config → runs
type key struct{ prompt, config string }
grouped := map[key][]RunResult{}
promptSet := map[string]bool{}
var promptOrder []string

for _, e := range entries {
if !promptSet[e.PromptID] {
promptSet[e.PromptID] = true
promptOrder = append(promptOrder, e.PromptID)
}
k := key{e.PromptID, e.ConfigName}
grouped[k] = append(grouped[k], RunResult{
RunID:     e.RunID,
Timestamp: e.Timestamp,
Success:   e.Success,
Duration:  e.Duration,
FileCount: e.FileCount,
Score:     e.Score,
MaxScore:  e.MaxScore,
HasReview: e.HasReview,
ToolCalls: e.ToolCalls,
Error:     e.Error,
})
}

var trends []PromptTrend
for _, pid := range promptOrder {
pt := PromptTrend{
PromptID: pid,
Configs:  map[string][]RunResult{},
Trend:    map[string]TrendClassification{},
}
for k, runs := range grouped {
if k.prompt == pid {
pt.Configs[k.config] = runs
pt.Trend[k.config] = classifyTrend(runs)
}
}
trends = append(trends, pt)
}

return trends, runIDs
}

// classifyTrend determines the trend classification for a series of runs.
func classifyTrend(runs []RunResult) TrendClassification {
if len(runs) <= 1 {
return TrendNew
}

passes, fails := 0, 0
for _, r := range runs {
if r.Success {
passes++
} else {
fails++
}
}

if fails == 0 {
// All pass — check duration trend
if len(runs) >= 3 {
mid := len(runs) / 2
firstAvg := avgDuration(runs[:mid])
secondAvg := avgDuration(runs[mid:])
if firstAvg > 0 && secondAvg < firstAvg*0.8 {
return TrendImproving
}
}
return TrendStable
}

if passes == 0 {
return TrendRegressing
}

// Previously passing, now failing → regression
if runs[0].Success && !runs[len(runs)-1].Success {
return TrendRegressing
}

return TrendFlaky
}

func avgDuration(runs []RunResult) float64 {
if len(runs) == 0 {
return 0
}
total := 0.0
for _, r := range runs {
total += r.Duration
}
return total / float64(len(runs))
}

// formatDuration returns a human-readable duration string.
func formatDuration(seconds float64) string {
if seconds <= 0 {
return "—"
}
if seconds < 60 {
return fmt.Sprintf("%.0fs", seconds)
}
if seconds < 3600 {
return fmt.Sprintf("%.1fm", seconds/60)
}
return fmt.Sprintf("%.1fh", seconds/3600)
}

// trendEmoji returns an emoji for the classification.
func trendEmoji(t TrendClassification) string {
switch t {
case TrendStable:
return "✅ stable"
case TrendImproving:
return "📈 improving"
case TrendRegressing:
return "📉 regressing"
case TrendFlaky:
return "⚠️ flaky"
case TrendNew:
return "🆕 new"
default:
return string(t)
}
}

// WriteMarkdown writes the trend report as a Markdown file.
func WriteMarkdown(tr *TrendReport, outputDir string) (string, error) {
if err := os.MkdirAll(outputDir, 0755); err != nil {
return "", fmt.Errorf("creating trends directory: %w", err)
}

filename := trendFilename(tr) + "-trends.md"
outPath := filepath.Join(outputDir, filename)

var b strings.Builder
writeMarkdownReport(&b, tr)

if err := os.WriteFile(outPath, []byte(b.String()), 0644); err != nil {
return "", fmt.Errorf("writing trend markdown: %w", err)
}
slog.Debug("Trend markdown written", "role", "trend-analysis", "path", outPath, "size", b.Len())
return outPath, nil
}

func trendFilename(tr *TrendReport) string {
if tr.PromptID != "" {
return tr.PromptID
}
parts := []string{}
if tr.Service != "" {
parts = append(parts, tr.Service)
}
if tr.Language != "" {
parts = append(parts, tr.Language)
}
if len(parts) == 0 {
return "all"
}
return strings.Join(parts, "-")
}

// scanReports walks the reports directory and extracts trend entries.
func scanReports(reportsDir, promptID, service, language string) ([]TrendEntry, error) {
var entries []TrendEntry

err := filepath.Walk(reportsDir, func(path string, info os.FileInfo, err error) error {
if err != nil {
return nil
}
if info.IsDir() || info.Name() != "report.json" {
return nil
}

data, err := os.ReadFile(path)
if err != nil {
return nil
}

var r report.EvalReport
if err := json.Unmarshal(data, &r); err != nil {
return nil
}

if promptID != "" && r.PromptID != promptID {
return nil
}
if service != "" {
svc, _ := r.PromptMeta["service"].(string)
if svc != service {
return nil
}
}
if language != "" {
lang, _ := r.PromptMeta["language"].(string)
if lang != language {
return nil
}
}

rel, _ := filepath.Rel(reportsDir, path)
parts := strings.Split(rel, string(os.PathSeparator))
runID := ""
if len(parts) > 0 {
runID = parts[0]
}

entry := TrendEntry{
RunID:      runID,
Timestamp:  r.Timestamp,
ConfigName: r.ConfigName,
PromptID:   r.PromptID,
Success:    r.Success,
Duration:   r.Duration,
ToolCalls:  r.ToolCalls,
FileCount:  len(r.GeneratedFiles),
Error:      r.Error,
}
if r.Review != nil {
entry.Score = r.Review.OverallScore
entry.MaxScore = r.Review.MaxScore
entry.HasReview = true
}

// Populate properties from prompt metadata for trend slicing.
if len(r.PromptMeta) > 0 {
props := make(map[string]string, len(r.PromptMeta))
for k, v := range r.PromptMeta {
if s, ok := v.(string); ok {
props[k] = s
}
}
entry.Properties = props
}

// Populate per-grader scores for regression detection.
for _, gr := range r.GraderResults {
entry.GraderScores = append(entry.GraderScores, GraderScore{
GraderName: gr.GraderName,
Score:      gr.Score,
MaxScore:   gr.MaxScore,
})
}

entries = append(entries, entry)
return nil
})

return entries, err
}

// --- Markdown report ---

func writeMarkdownReport(b *strings.Builder, tr *TrendReport) {
title := "Performance Trends"
if tr.PromptID != "" {
title = fmt.Sprintf("Trends: %s", tr.PromptID)
} else if tr.Service != "" || tr.Language != "" {
parts := []string{}
if tr.Service != "" {
parts = append(parts, tr.Service)
}
if tr.Language != "" {
parts = append(parts, tr.Language)
}
title = fmt.Sprintf("Trends: %s", strings.Join(parts, " / "))
}

fmt.Fprintf(b, "# %s\n\n", title)
fmt.Fprintf(b, "**Generated:** %s | **Total Evaluations:** %d\n\n", tr.GeneratedAt, tr.TotalRuns)

if tr.TotalRuns == 0 {
b.WriteString("No historical data found matching the given filters.\n")
return
}

// AI Analysis section
if tr.Analysis != "" {
b.WriteString("## 🤖 AI Analysis\n\n")
b.WriteString(tr.Analysis)
b.WriteString("\n\n---\n\n")
}

// Summary statistics
b.WriteString("## Summary\n\n")
passed, failed, totalScore, scored := 0, 0, 0, 0
configCounts := map[string]int{}
for _, e := range tr.Entries {
if e.Success {
passed++
} else {
failed++
}
if e.HasReview {
totalScore += e.Score
scored++
}
configCounts[e.ConfigName]++
}

b.WriteString("| Metric | Value |\n")
b.WriteString("|--------|-------|\n")
fmt.Fprintf(b, "| Total Evaluations | %d |\n", tr.TotalRuns)
fmt.Fprintf(b, "| Passed | %d (%.0f%%) |\n", passed, pct(passed, tr.TotalRuns))
fmt.Fprintf(b, "| Failed | %d |\n", failed)
if scored > 0 {
fmt.Fprintf(b, "| Avg Score | %.1f avg |\n", float64(totalScore)/float64(scored))
}
fmt.Fprintf(b, "| Unique Prompts | %d |\n", len(tr.PromptTrends))
fmt.Fprintf(b, "| Configs | %d |\n", len(configCounts))
b.WriteString("\n")

// Regression alerts
var regressions []string
for _, pt := range tr.PromptTrends {
for cfg, trend := range pt.Trend {
if trend == TrendRegressing {
regressions = append(regressions, fmt.Sprintf("- 📉 **%s** / `%s` — previously passing, now failing", pt.PromptID, cfg))
}
}
}
if len(regressions) > 0 {
b.WriteString("## ⚠️ Regression Alerts\n\n")
for _, r := range regressions {
b.WriteString(r + "\n")
}
b.WriteString("\n")
}

// Performance Over Time table
if len(tr.PromptTrends) > 0 && len(tr.RunIDs) > 0 {
b.WriteString("## Performance Over Time\n\n")

// Header
b.WriteString("| Prompt | Config |")
displayIDs := tr.RunIDs
if len(displayIDs) > 8 {
displayIDs = displayIDs[len(displayIDs)-8:]
}
for _, rid := range displayIDs {
short := rid
if len(short) > 10 {
short = short[:10]
}
fmt.Fprintf(b, " %s |", short)
}
b.WriteString(" Trend |\n")

// Separator
b.WriteString("|--------|--------|")
for range displayIDs {
b.WriteString("--------|")
}
b.WriteString("-------|\n")

// Rows
for _, pt := range tr.PromptTrends {
configs := sortedConfigNames(pt.Configs)
for _, cfg := range configs {
runs := pt.Configs[cfg]
runMap := map[string]RunResult{}
for _, r := range runs {
runMap[r.RunID] = r
}

fmt.Fprintf(b, "| %s | %s |", pt.PromptID, cfg)
for _, rid := range displayIDs {
if r, ok := runMap[rid]; ok {
icon := "❌"
if r.Success {
icon = "✅"
}
fmt.Fprintf(b, " %s %s |", icon, formatDuration(r.Duration))
} else {
b.WriteString(" — |")
}
}
fmt.Fprintf(b, " %s |\n", trendEmoji(pt.Trend[cfg]))
}
}
b.WriteString("\n")
}

// Config comparison
if len(configCounts) > 1 {
b.WriteString("## Config Comparison\n\n")
b.WriteString("| Config | Runs | Pass Rate | Avg Duration | Avg Score |\n")
b.WriteString("|--------|------|-----------|--------------|----------|\n")
for cfg, count := range configCounts {
cp, cs, cn := 0, 0, 0
totalDur := 0.0
for _, e := range tr.Entries {
if e.ConfigName == cfg {
if e.Success {
cp++
}
totalDur += e.Duration
if e.HasReview {
cs += e.Score
cn++
}
}
}
avgScore := "—"
if cn > 0 {
avgScore = fmt.Sprintf("%.1f avg", float64(cs)/float64(cn))
}
fmt.Fprintf(b, "| %s | %d | %.0f%% | %s | %s |\n",
cfg, count, pct(cp, count), formatDuration(totalDur/float64(count)), avgScore)
}
b.WriteString("\n")
}
}

func sortedConfigNames(configs map[string][]RunResult) []string {
names := make([]string, 0, len(configs))
for k := range configs {
names = append(names, k)
}
sort.Strings(names)
return names
}

func pct(n, total int) float64 {
if total == 0 {
return 0
}
return float64(n) / float64(total) * 100
}
