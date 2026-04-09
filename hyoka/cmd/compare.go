package cmd

import (
"encoding/json"
"fmt"
"strings"
"time"

"github.com/ronniegeraghty/hyoka/internal/comparison"
"github.com/spf13/cobra"
)

func compareCmd() *cobra.Command {
var (
configA    string
configB    string
runA       string
runB       string
config     string
since      string
reportsDir string
format     string
topN       int
)

cmd := &cobra.Command{
Use:   "compare",
Short: "Compare evaluation results between configs, runs, or time periods",
Long: `Compare evaluation results to identify regressions and improvements.

Three comparison modes:

  Config comparison (--config-a / --config-b):
    Compare latest results for two configs across all shared prompts.

  Run comparison (--run-a / --run-b):
    Compare results from two specific evaluation runs.

  Temporal comparison (--config / --since):
    Compare a config's results before and after a date.`,
RunE: func(cmd *cobra.Command, args []string) error {
reportsDir = resolvePathFlag(cmd, "reports-dir", []string{"./reports", "../reports"})
out := cmd.OutOrStdout()

mode, err := detectMode(configA, configB, runA, runB, config, since)
if err != nil {
return err
}

switch mode {
case "configs":
cmp, err := comparison.CompareConfigs(reportsDir, configA, configB)
if err != nil {
return fmt.Errorf("config comparison: %w", err)
}
return renderConfigComparison(out, cmp, format, topN)

case "runs":
cmp, err := comparison.CompareRuns(reportsDir, runA, runB)
if err != nil {
return fmt.Errorf("run comparison: %w", err)
}
return renderRunComparison(out, cmp, format, topN)

case "temporal":
sinceTime, err := time.Parse("2006-01-02", since)
if err != nil {
return fmt.Errorf("invalid --since date %q (expected YYYY-MM-DD): %w", since, err)
}
cmp, err := comparison.TemporalDiff(reportsDir, config, sinceTime)
if err != nil {
return fmt.Errorf("temporal comparison: %w", err)
}
return renderTemporalComparison(out, cmp, format, topN)

default:
return fmt.Errorf("unknown mode %q", mode)
}
},
}

cmd.Flags().StringVar(&configA, "config-a", "", "First config name (for config comparison)")
cmd.Flags().StringVar(&configB, "config-b", "", "Second config name (for config comparison)")
cmd.Flags().StringVar(&runA, "run-a", "", "First run ID (for run comparison)")
cmd.Flags().StringVar(&runB, "run-b", "", "Second run ID (for run comparison)")
cmd.Flags().StringVar(&config, "config", "", "Config name (for temporal comparison)")
cmd.Flags().StringVar(&since, "since", "", "Cutoff date YYYY-MM-DD (for temporal comparison)")
cmd.Flags().StringVar(&reportsDir, "reports-dir", "./reports", "Directory containing evaluation reports")
cmd.Flags().StringVar(&format, "format", "table", "Output format: table or json")
cmd.Flags().IntVar(&topN, "top", 0, "Show only top N changes (0 = all)")

return cmd
}

// detectMode determines which comparison mode to use based on flag combinations.
func detectMode(configA, configB, runA, runB, config, since string) (string, error) {
hasConfigPair := configA != "" || configB != ""
hasRunPair := runA != "" || runB != ""
hasTemporal := config != "" || since != ""

set := 0
if hasConfigPair {
set++
}
if hasRunPair {
set++
}
if hasTemporal {
set++
}

if set == 0 {
return "", fmt.Errorf("specify one of:\n  --config-a + --config-b  (compare configs)\n  --run-a + --run-b        (compare runs)\n  --config + --since       (temporal comparison)")
}
if set > 1 {
return "", fmt.Errorf("only one comparison mode at a time: use --config-a/--config-b, --run-a/--run-b, or --config/--since")
}

if hasConfigPair {
if configA == "" || configB == "" {
return "", fmt.Errorf("both --config-a and --config-b are required")
}
return "configs", nil
}
if hasRunPair {
if runA == "" || runB == "" {
return "", fmt.Errorf("both --run-a and --run-b are required")
}
return "runs", nil
}
// hasTemporal
if config == "" || since == "" {
return "", fmt.Errorf("both --config and --since are required for temporal comparison")
}
return "temporal", nil
}

// ---------- Rendering ----------

func renderConfigComparison(out interface{ Write([]byte) (int, error) }, cmp *comparison.ConfigComparison, format string, topN int) error {
if format == "json" {
return writeJSON(out, cmp)
}
fmt.Fprintf(out, "Config comparison: %s vs %s\n", cmp.ConfigA, cmp.ConfigB)
renderDiffTable(out, cmp.PerPrompt, cmp.Summary, topN)
return nil
}

func renderRunComparison(out interface{ Write([]byte) (int, error) }, cmp *comparison.RunComparison, format string, topN int) error {
if format == "json" {
return writeJSON(out, cmp)
}
fmt.Fprintf(out, "Run comparison: %s vs %s\n", cmp.RunA, cmp.RunB)
renderDiffTable(out, cmp.PerPrompt, cmp.Summary, topN)
return nil
}

func renderTemporalComparison(out interface{ Write([]byte) (int, error) }, cmp *comparison.TemporalComparison, format string, topN int) error {
if format == "json" {
return writeJSON(out, cmp)
}
fmt.Fprintf(out, "Temporal comparison: %s (base: %s → latest: %s)\n", cmp.Config, cmp.BaseRun, cmp.LatestRun)
renderDiffTable(out, cmp.PerPrompt, cmp.Summary, topN)
return nil
}

func renderDiffTable(out interface{ Write([]byte) (int, error) }, diffs []comparison.PromptDiff, summary comparison.ComparisonSummary, topN int) {
sep := strings.Repeat("─", 72)
fmt.Fprintf(out, "%s\n", sep)
fmt.Fprintf(out, "%-40s %8s %8s %8s %6s\n", "Prompt", "A", "B", "Delta", "")
fmt.Fprintf(out, "%s\n", sep)

shown := diffs
if topN > 0 && topN < len(diffs) {
shown = diffs[:topN]
}

for _, d := range shown {
label := d.PromptID
if len(label) > 39 {
label = label[:36] + "..."
}

indicator := "  "
switch {
case d.OnlyInA:
indicator = "⊖"
case d.OnlyInB:
indicator = "⊕"
case d.Delta > 0.001:
indicator = "▲"
case d.Delta < -0.001:
indicator = "▼"
}

if d.OnlyInA {
fmt.Fprintf(out, "%-40s %8.3f %8s %8s %s\n", label, d.ScoreA, "—", "—", indicator)
} else if d.OnlyInB {
fmt.Fprintf(out, "%-40s %8s %8.3f %8s %s\n", label, "—", d.ScoreB, "—", indicator)
} else {
fmt.Fprintf(out, "%-40s %8.3f %8.3f %+8.3f %s\n", label, d.ScoreA, d.ScoreB, d.Delta, indicator)
}
}

if topN > 0 && topN < len(diffs) {
fmt.Fprintf(out, "\n(showing top %d of %d prompts)\n", topN, len(diffs))
}

fmt.Fprintf(out, "%s\n", sep)
fmt.Fprintf(out, "Summary: %d improved, %d regressed, %d unchanged (avg Δ %+.3f)\n",
summary.Improved, summary.Regressed, summary.Unchanged, summary.AvgDelta)

if len(summary.TopImproved) > 0 {
fmt.Fprintf(out, "\nTop improved:\n")
for _, d := range summary.TopImproved {
fmt.Fprintf(out, "  %s: %+.3f\n", d.PromptID, d.Delta)
}
}
if len(summary.TopRegressed) > 0 {
fmt.Fprintf(out, "\nTop regressed:\n")
for _, d := range summary.TopRegressed {
fmt.Fprintf(out, "  %s: %+.3f\n", d.PromptID, d.Delta)
}
}
}

func writeJSON(out interface{ Write([]byte) (int, error) }, v any) error {
data, err := json.MarshalIndent(v, "", "  ")
if err != nil {
return fmt.Errorf("marshalling JSON: %w", err)
}
_, err = out.Write(append(data, '\n'))
return err
}
