package trends

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/report"
)

// TrendEntry holds data from a single historical evaluation run.
type TrendEntry struct {
	RunID      string   `json:"run_id"`
	Timestamp  string   `json:"timestamp"`
	ConfigName string   `json:"config_name"`
	PromptID   string   `json:"prompt_id"`
	Success    bool     `json:"success"`
	Duration   float64  `json:"duration_seconds"`
	Score      int      `json:"score"`
	HasReview  bool     `json:"has_review"`
	ToolCalls  []string `json:"tool_calls"`
	FileCount  int      `json:"file_count"`
	Error      string   `json:"error,omitempty"`
}

// TrendReport summarizes historical trends for a set of prompts.
type TrendReport struct {
	PromptID    string       `json:"prompt_id,omitempty"`
	Service     string       `json:"service,omitempty"`
	Language    string       `json:"language,omitempty"`
	TotalRuns   int          `json:"total_runs"`
	Entries     []TrendEntry `json:"entries"`
	GeneratedAt string       `json:"generated_at"`
}

// TrendOptions configures trend report generation.
type TrendOptions struct {
	ReportsDir string
	PromptID   string
	Service    string
	Language   string
	OutputDir  string
}

// Generate scans historical reports and produces a trend report.
func Generate(opts TrendOptions) (*TrendReport, error) {
	entries, err := scanReports(opts.ReportsDir, opts.PromptID, opts.Service, opts.Language)
	if err != nil {
		return nil, fmt.Errorf("scanning reports: %w", err)
	}

	// Sort by timestamp
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp < entries[j].Timestamp
	})

	tr := &TrendReport{
		PromptID:    opts.PromptID,
		Service:     opts.Service,
		Language:    opts.Language,
		TotalRuns:   len(entries),
		Entries:     entries,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
	}

	return tr, nil
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
	return outPath, nil
}

// WriteHTML writes the trend report as an HTML file.
func WriteHTML(tr *TrendReport, outputDir string) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("creating trends directory: %w", err)
	}

	filename := trendFilename(tr) + "-trends.html"
	outPath := filepath.Join(outputDir, filename)

	var b strings.Builder
	writeHTMLReport(&b, tr)

	if err := os.WriteFile(outPath, []byte(b.String()), 0644); err != nil {
		return "", fmt.Errorf("writing trend HTML: %w", err)
	}
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

	// Walk looking for report.json files
	err := filepath.Walk(reportsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip unreadable
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

		// Apply filters
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

		// Extract run ID from directory structure (parent of results/)
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
			entry.HasReview = true
		}

		entries = append(entries, entry)
		return nil
	})

	return entries, err
}

func writeMarkdownReport(b *strings.Builder, tr *TrendReport) {
	title := "Historical Trends"
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
	fmt.Fprintf(b, "**Generated:** %s | **Total Runs:** %d\n\n", tr.GeneratedAt, tr.TotalRuns)

	if tr.TotalRuns == 0 {
		b.WriteString("No historical data found matching the given filters.\n")
		return
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
		fmt.Fprintf(b, "| Avg Score | %.1f/10 |\n", float64(totalScore)/float64(scored))
	}
	b.WriteString("\n")

	// Config comparison
	if len(configCounts) > 1 {
		b.WriteString("## Config Comparison\n\n")
		b.WriteString("| Config | Runs | Pass Rate | Avg Score |\n")
		b.WriteString("|--------|------|-----------|----------|\n")
		for cfg, count := range configCounts {
			cp, cs, cn := 0, 0, 0
			for _, e := range tr.Entries {
				if e.ConfigName == cfg {
					if e.Success {
						cp++
					}
					if e.HasReview {
						cs += e.Score
						cn++
					}
				}
			}
			avgScore := "—"
			if cn > 0 {
				avgScore = fmt.Sprintf("%.1f/10", float64(cs)/float64(cn))
			}
			fmt.Fprintf(b, "| %s | %d | %.0f%% | %s |\n", cfg, count, pct(cp, count), avgScore)
		}
		b.WriteString("\n")
	}

	// Run history table
	b.WriteString("## Run History\n\n")
	b.WriteString("| Run ID | Config | Result | Score | Duration | Files | Tools |\n")
	b.WriteString("|--------|--------|--------|-------|----------|-------|-------|\n")
	for _, e := range tr.Entries {
		icon := "❌"
		if e.Success {
			icon = "✅"
		}
		score := "—"
		if e.HasReview {
			score = fmt.Sprintf("%d/10", e.Score)
		}
		tools := strings.Join(e.ToolCalls, ", ")
		if len(tools) > 60 {
			tools = tools[:57] + "..."
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s | %.1fs | %d | %s |\n",
			e.RunID, e.ConfigName, icon, score, e.Duration, e.FileCount, tools)
	}
	b.WriteString("\n")

	// Score trend (if we have scored entries)
	if scored > 1 {
		b.WriteString("## Score Trend\n\n")
		b.WriteString("| Timestamp | Config | Score |\n")
		b.WriteString("|-----------|--------|-------|\n")
		for _, e := range tr.Entries {
			if e.HasReview {
				fmt.Fprintf(b, "| %s | %s | %d/10 |\n", e.Timestamp, e.ConfigName, e.Score)
			}
		}
		b.WriteString("\n")

		// Detect regressions / improvements
		b.WriteString("### Changes Detected\n\n")
		prevScores := map[string]int{}
		hasChanges := false
		for _, e := range tr.Entries {
			if !e.HasReview {
				continue
			}
			key := e.PromptID + "/" + e.ConfigName
			if prev, ok := prevScores[key]; ok {
				if e.Score < prev {
					fmt.Fprintf(b, "- 📉 **Regression**: %s (%s) dropped from %d to %d\n", e.PromptID, e.ConfigName, prev, e.Score)
					hasChanges = true
				} else if e.Score > prev {
					fmt.Fprintf(b, "- 📈 **Improvement**: %s (%s) improved from %d to %d\n", e.PromptID, e.ConfigName, prev, e.Score)
					hasChanges = true
				}
			}
			prevScores[key] = e.Score
		}
		if !hasChanges {
			b.WriteString("No score changes detected between runs.\n")
		}
		b.WriteString("\n")
	}
}

func writeHTMLReport(b *strings.Builder, tr *TrendReport) {
	title := "Historical Trends"
	if tr.PromptID != "" {
		title = fmt.Sprintf("Trends: %s", tr.PromptID)
	}

	fmt.Fprintf(b, `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>%s</title>
<style>
  :root { --green: #22c55e; --red: #ef4444; --yellow: #eab308; --bg: #f8fafc; --text: #0f172a; --text-muted: #64748b; --border: #e2e8f0; --blue: #2563eb; }
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; max-width: 1200px; margin: 0 auto; padding: 2rem 1rem; color: var(--text); background: var(--bg); }
  h1 { margin: 0 0 0.25rem; } h2 { margin: 1.5rem 0 0.75rem; }
  .subtitle { color: var(--text-muted); margin-bottom: 1.5rem; }
  .stats { display: flex; gap: 1rem; flex-wrap: wrap; margin: 1.25rem 0; }
  .stat { background: #fff; border: 1px solid var(--border); border-radius: 8px; padding: 1rem 1.25rem; text-align: center; min-width: 110px; }
  .stat-value { font-size: 1.5rem; font-weight: 700; } .stat-label { font-size: 0.8rem; color: var(--text-muted); }
  table { width: 100%%; border-collapse: collapse; background: #fff; border: 1px solid var(--border); border-radius: 8px; overflow: hidden; margin-bottom: 1.5rem; }
  th { background: #f8fafc; padding: 0.6rem 0.75rem; text-align: left; font-size: 0.8rem; color: var(--text-muted); border-bottom: 2px solid var(--border); }
  td { padding: 0.6rem 0.75rem; border-bottom: 1px solid #f1f5f9; font-size: 0.85rem; }
  .pass { color: var(--green); } .fail { color: var(--red); }
  .tag { display: inline-block; background: #f3f0ff; color: #7c3aed; padding: 1px 6px; border-radius: 3px; font-size: 0.75rem; font-family: monospace; margin: 1px; }
  .change-up { color: var(--green); } .change-down { color: var(--red); }
</style>
</head>
<body>
`, title)

	fmt.Fprintf(b, "<h1>📈 %s</h1>\n", title)
	fmt.Fprintf(b, "<div class=\"subtitle\">Generated: %s | Total Runs: %d</div>\n", tr.GeneratedAt, tr.TotalRuns)

	if tr.TotalRuns == 0 {
		b.WriteString("<p>No historical data found matching the given filters.</p>\n</body>\n</html>")
		return
	}

	// Summary stats
	passed, failed, totalScore, scored := 0, 0, 0, 0
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
	}

	b.WriteString("<div class=\"stats\">\n")
	fmt.Fprintf(b, "  <div class=\"stat\"><div class=\"stat-value\">%d</div><div class=\"stat-label\">Total Runs</div></div>\n", tr.TotalRuns)
	fmt.Fprintf(b, "  <div class=\"stat\"><div class=\"stat-value pass\">%d</div><div class=\"stat-label\">Passed</div></div>\n", passed)
	fmt.Fprintf(b, "  <div class=\"stat\"><div class=\"stat-value fail\">%d</div><div class=\"stat-label\">Failed</div></div>\n", failed)
	if scored > 0 {
		fmt.Fprintf(b, "  <div class=\"stat\"><div class=\"stat-value\">%.1f</div><div class=\"stat-label\">Avg Score</div></div>\n", float64(totalScore)/float64(scored))
	}
	b.WriteString("</div>\n\n")

	// Run history table
	b.WriteString("<h2>Run History</h2>\n<table>\n<thead><tr><th>Run ID</th><th>Config</th><th>Result</th><th>Score</th><th>Duration</th><th>Files</th><th>Tools</th></tr></thead>\n<tbody>\n")
	for _, e := range tr.Entries {
		icon := "❌"
		cls := "fail"
		if e.Success {
			icon = "✅"
			cls = "pass"
		}
		score := "—"
		if e.HasReview {
			score = fmt.Sprintf("%d/10", e.Score)
		}
		var toolTags strings.Builder
		for _, t := range e.ToolCalls {
			fmt.Fprintf(&toolTags, "<span class=\"tag\">%s</span>", t)
		}
		fmt.Fprintf(b, "<tr><td>%s</td><td>%s</td><td class=\"%s\">%s</td><td>%s</td><td>%.1fs</td><td>%d</td><td>%s</td></tr>\n",
			e.RunID, e.ConfigName, cls, icon, score, e.Duration, e.FileCount, toolTags.String())
	}
	b.WriteString("</tbody>\n</table>\n\n")

	b.WriteString("</body>\n</html>")
}

func pct(n, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(n) / float64(total) * 100
}
