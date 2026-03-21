package report

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// WriteHTMLReport writes an individual evaluation report as HTML.
func WriteHTMLReport(r *EvalReport, outputDir string, runID string, service, plane, language, category string) (string, error) {
	reportDir := filepath.Join(
		outputDir, runID, "results",
		service, plane, language, category, r.ConfigName,
	)
	if err := os.MkdirAll(reportDir, 0755); err != nil {
		return "", fmt.Errorf("creating HTML report directory: %w", err)
	}

	reportPath := filepath.Join(reportDir, "report.html")

	f, err := os.Create(reportPath)
	if err != nil {
		return "", fmt.Errorf("creating HTML report file: %w", err)
	}
	defer f.Close()

	tmpl, err := template.New("report").Funcs(htmlFuncMap()).Parse(reportTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing report template: %w", err)
	}

	if err := tmpl.Execute(f, r); err != nil {
		return "", fmt.Errorf("executing report template: %w", err)
	}

	return reportPath, nil
}

// WriteSummaryHTML writes a cross-config comparison summary as HTML.
func WriteSummaryHTML(s *RunSummary, outputDir string) (string, error) {
	summaryDir := filepath.Join(outputDir, s.RunID)
	if err := os.MkdirAll(summaryDir, 0755); err != nil {
		return "", fmt.Errorf("creating summary directory: %w", err)
	}

	summaryPath := filepath.Join(summaryDir, "summary.html")

	f, err := os.Create(summaryPath)
	if err != nil {
		return "", fmt.Errorf("creating summary HTML file: %w", err)
	}
	defer f.Close()

	// Build the matrix data
	matrix := buildMatrix(s)

	tmpl, err := template.New("summary").Funcs(htmlFuncMap()).Parse(summaryTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing summary template: %w", err)
	}

	data := struct {
		Summary *RunSummary
		Matrix  *MatrixData
	}{
		Summary: s,
		Matrix:  matrix,
	}

	if err := tmpl.Execute(f, data); err != nil {
		return "", fmt.Errorf("executing summary template: %w", err)
	}

	return summaryPath, nil
}

// MatrixData holds the cross-config comparison matrix.
type MatrixData struct {
	Configs  []string
	Prompts  []string
	Cells    map[string]map[string]*MatrixCell // [promptID][configName]
}

// MatrixCell holds the data for one cell in the matrix.
type MatrixCell struct {
	Score      int
	BuildPass  bool
	HasReview  bool
	Duration   float64
	Error      string
}

func buildMatrix(s *RunSummary) *MatrixData {
	m := &MatrixData{
		Cells: make(map[string]map[string]*MatrixCell),
	}

	configSet := make(map[string]bool)
	promptSet := make(map[string]bool)

	for _, r := range s.Results {
		if !promptSet[r.PromptID] {
			promptSet[r.PromptID] = true
			m.Prompts = append(m.Prompts, r.PromptID)
		}
		if !configSet[r.ConfigName] {
			configSet[r.ConfigName] = true
			m.Configs = append(m.Configs, r.ConfigName)
		}

		if m.Cells[r.PromptID] == nil {
			m.Cells[r.PromptID] = make(map[string]*MatrixCell)
		}

		cell := &MatrixCell{
			Duration: r.Duration,
			Error:    r.Error,
		}
		if r.Build != nil {
			cell.BuildPass = r.Build.Success
		}
		if r.Review != nil {
			cell.Score = r.Review.OverallScore
			cell.HasReview = true
		}
		m.Cells[r.PromptID][r.ConfigName] = cell
	}

	return m
}

func htmlFuncMap() template.FuncMap {
	return template.FuncMap{
		"scoreColor": func(score int) string {
			switch {
			case score >= 8:
				return "#22c55e" // green
			case score >= 6:
				return "#eab308" // yellow
			case score >= 4:
				return "#f97316" // orange
			default:
				return "#ef4444" // red
			}
		},
		"buildIcon": func(pass bool) string {
			if pass {
				return "✅"
			}
			return "❌"
		},
		"join": func(items []string, sep string) string {
			return strings.Join(items, sep)
		},
		"fmtDuration": func(d float64) string {
			return fmt.Sprintf("%.1fs", d)
		},
	}
}

const reportTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Eval Report: {{.PromptID}} / {{.ConfigName}}</title>
<style>
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; max-width: 960px; margin: 2rem auto; padding: 0 1rem; color: #1a1a1a; background: #fafafa; }
  h1, h2, h3 { color: #0d1117; }
  .badge { display: inline-block; padding: 2px 8px; border-radius: 4px; font-size: 0.85em; font-weight: 600; }
  .badge-pass { background: #dcfce7; color: #166534; }
  .badge-fail { background: #fef2f2; color: #991b1b; }
  .badge-stub { background: #fef3c7; color: #92400e; }
  .card { background: white; border: 1px solid #e5e7eb; border-radius: 8px; padding: 1rem; margin: 1rem 0; }
  .scores { display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 0.5rem; }
  .score-item { text-align: center; padding: 0.5rem; border-radius: 4px; background: #f9fafb; }
  .score-value { font-size: 1.5rem; font-weight: bold; }
  .score-label { font-size: 0.8rem; color: #6b7280; }
  details { margin: 0.5rem 0; }
  summary { cursor: pointer; font-weight: 600; padding: 0.5rem 0; }
  pre { background: #f1f5f9; padding: 1rem; border-radius: 4px; overflow-x: auto; font-size: 0.85rem; }
  .meta-table { width: 100%; border-collapse: collapse; }
  .meta-table td { padding: 0.25rem 0.5rem; border-bottom: 1px solid #f3f4f6; }
  .meta-table td:first-child { font-weight: 600; width: 140px; color: #6b7280; }
  .files-list { list-style: none; padding: 0; }
  .files-list li { padding: 0.25rem 0; font-family: monospace; font-size: 0.9rem; }
  .event { padding: 0.25rem 0; font-size: 0.85rem; border-bottom: 1px solid #f3f4f6; }
  .event-type { font-weight: 600; color: #6b7280; font-size: 0.75rem; text-transform: uppercase; }
  .event-tool { color: #7c3aed; }
  .event-assistant { color: #0369a1; }
  .event-error { color: #dc2626; }
  .error-box { background: #fef2f2; border: 1px solid #fecaca; border-radius: 8px; padding: 1rem; margin: 1rem 0; }
</style>
</head>
<body>
<h1>{{.PromptID}} <span style="color:#6b7280">/ {{.ConfigName}}</span></h1>

{{if .IsStub}}
<div class="card" style="background:#fffbeb;border-color:#fde68a">
  <h2>⚠️ Stub Mode</h2>
  <p>Running in stub mode — no Copilot session was created. Results are placeholders.</p>
</div>
{{end}}

<div class="card">
  <h2>Overview</h2>
  <table class="meta-table">
    <tr><td>Status</td><td>{{if .Success}}<span class="badge badge-pass">PASS</span>{{else}}<span class="badge badge-fail">FAIL</span>{{end}}{{if .IsStub}} <span class="badge badge-stub">STUB</span>{{end}}</td></tr>
    <tr><td>Duration</td><td>{{fmtDuration .Duration}}</td></tr>
    <tr><td>Timestamp</td><td>{{.Timestamp}}</td></tr>
    <tr><td>Events</td><td>{{.EventCount}}</td></tr>
    {{if .Error}}<tr><td>Error</td><td style="color:#991b1b">{{.Error}}</td></tr>{{end}}
  </table>
</div>

{{if .Error}}
<div class="error-box">
  <h2>❌ Error Details</h2>
  <p><strong>{{.Error}}</strong></p>
  {{if .ErrorDetails}}<details open><summary>Full Error</summary><pre>{{.ErrorDetails}}</pre></details>{{end}}
</div>
{{end}}

{{if .Verification}}
<div class="card">
  <h2>Verification {{if .Verification.Pass}}<span class="badge badge-pass">PASS</span>{{else}}<span class="badge badge-fail">FAIL</span>{{end}}</h2>
  <p><strong>{{.Verification.Summary}}</strong></p>
  {{if .Verification.Reasoning}}<details open><summary>Reasoning</summary><pre>{{.Verification.Reasoning}}</pre></details>{{end}}
</div>
{{end}}

{{if .Review}}
<div class="card">
  <h2>Code Review Scores</h2>
  <div class="scores">
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.Correctness}}">{{.Review.Scores.Correctness}}</div><div class="score-label">Correctness</div></div>
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.Completeness}}">{{.Review.Scores.Completeness}}</div><div class="score-label">Completeness</div></div>
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.BestPractices}}">{{.Review.Scores.BestPractices}}</div><div class="score-label">Best Practices</div></div>
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.ErrorHandling}}">{{.Review.Scores.ErrorHandling}}</div><div class="score-label">Error Handling</div></div>
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.PackageUsage}}">{{.Review.Scores.PackageUsage}}</div><div class="score-label">Package Usage</div></div>
    <div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.CodeQuality}}">{{.Review.Scores.CodeQuality}}</div><div class="score-label">Code Quality</div></div>
    {{if .Review.Scores.ReferenceSimilarity}}<div class="score-item"><div class="score-value" style="color:{{scoreColor .Review.Scores.ReferenceSimilarity}}">{{.Review.Scores.ReferenceSimilarity}}</div><div class="score-label">Ref Similarity</div></div>{{end}}
  </div>
  <div style="text-align:center;margin-top:1rem">
    <div class="score-value" style="font-size:2rem;color:{{scoreColor .Review.OverallScore}}">{{.Review.OverallScore}}/10</div>
    <div class="score-label">Overall Score</div>
  </div>
  <p>{{.Review.Summary}}</p>
  {{if .Review.Strengths}}
  <details open><summary>Strengths</summary><ul>{{range .Review.Strengths}}<li>{{.}}</li>{{end}}</ul></details>
  {{end}}
  {{if .Review.Issues}}
  <details open><summary>Issues</summary><ul>{{range .Review.Issues}}<li>{{.}}</li>{{end}}</ul></details>
  {{end}}
</div>
{{end}}

<div class="card">
  <h2>Generated Files</h2>
  {{if .GeneratedFiles}}
  <ul class="files-list">{{range .GeneratedFiles}}<li>📄 {{.}}</li>{{end}}</ul>
  {{else}}<p>No files generated.</p>{{end}}
</div>

{{if .Build}}
<div class="card">
  <h2>Build Verification {{if .Build.Success}}<span class="badge badge-pass">PASS</span>{{else}}<span class="badge badge-fail">FAIL</span>{{end}}</h2>
  <table class="meta-table">
    <tr><td>Language</td><td>{{.Build.Language}}</td></tr>
    <tr><td>Command</td><td><code>{{.Build.Command}}</code></td></tr>
    <tr><td>Exit Code</td><td>{{.Build.ExitCode}}</td></tr>
  </table>
  {{if .Build.Stdout}}<details><summary>Stdout</summary><pre>{{.Build.Stdout}}</pre></details>{{end}}
  {{if .Build.Stderr}}<details><summary>Stderr</summary><pre>{{.Build.Stderr}}</pre></details>{{end}}
</div>
{{end}}

{{if .ToolCalls}}
<div class="card">
  <h2>Tool Calls</h2>
  <p>{{join .ToolCalls ", "}}</p>
</div>
{{end}}

{{if .SessionEvents}}
<div class="card">
  <h2>Session Transcript</h2>
  <p>{{.EventCount}} events captured</p>
  <details><summary>Show all events</summary>
  {{range .SessionEvents}}
  <div class="event">
    <span class="event-type">{{.Type}}</span>
    {{if .ToolName}}<span class="event-tool"> {{.ToolName}}</span>{{end}}
    {{if .ToolArgs}}<span style="color:#6b7280"> ({{.ToolArgs}})</span>{{end}}
    {{if .Content}}<pre style="margin:0.25rem 0;padding:0.5rem">{{.Content}}</pre>{{end}}
    {{if .Error}}<span class="event-error">{{.Error}}</span>{{end}}
  </div>
  {{end}}
  </details>
</div>
{{end}}

</body>
</html>`

const summaryTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Evaluation Summary — {{.Summary.RunID}}</title>
<style>
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; max-width: 1200px; margin: 2rem auto; padding: 0 1rem; color: #1a1a1a; background: #fafafa; }
  h1, h2 { color: #0d1117; }
  .stats { display: flex; gap: 1rem; flex-wrap: wrap; margin: 1rem 0; }
  .stat { background: white; border: 1px solid #e5e7eb; border-radius: 8px; padding: 1rem; text-align: center; min-width: 120px; }
  .stat-value { font-size: 1.5rem; font-weight: bold; }
  .stat-label { font-size: 0.8rem; color: #6b7280; }
  table { width: 100%; border-collapse: collapse; background: white; border: 1px solid #e5e7eb; border-radius: 8px; overflow: hidden; }
  th { background: #f9fafb; padding: 0.75rem; text-align: left; font-size: 0.85rem; color: #6b7280; border-bottom: 2px solid #e5e7eb; }
  td { padding: 0.75rem; border-bottom: 1px solid #f3f4f6; }
  .cell { text-align: center; }
  .cell-score { font-weight: bold; font-size: 1.1rem; }
  .cell-build { font-size: 0.85rem; }
  .cell-error { color: #991b1b; font-size: 0.8rem; }
</style>
</head>
<body>
<h1>Evaluation Summary</h1>
<p>Run: <strong>{{.Summary.RunID}}</strong> — {{.Summary.Timestamp}}</p>

<div class="stats">
  <div class="stat"><div class="stat-value">{{.Summary.TotalEvals}}</div><div class="stat-label">Evaluations</div></div>
  <div class="stat"><div class="stat-value" style="color:#22c55e">{{.Summary.Passed}}</div><div class="stat-label">Passed</div></div>
  <div class="stat"><div class="stat-value" style="color:#ef4444">{{.Summary.Failed}}</div><div class="stat-label">Failed</div></div>
  <div class="stat"><div class="stat-value" style="color:#f97316">{{.Summary.Errors}}</div><div class="stat-label">Errors</div></div>
  <div class="stat"><div class="stat-value">{{fmtDuration .Summary.Duration}}</div><div class="stat-label">Duration</div></div>
</div>

{{if .Matrix}}
<h2>Prompt × Config Matrix</h2>
<table>
  <thead>
    <tr>
      <th>Prompt</th>
      {{range .Matrix.Configs}}<th class="cell">{{.}}</th>{{end}}
    </tr>
  </thead>
  <tbody>
    {{range $prompt := .Matrix.Prompts}}
    <tr>
      <td><code>{{$prompt}}</code></td>
      {{range $config := $.Matrix.Configs}}
      <td class="cell">
        {{with index (index $.Matrix.Cells $prompt) $config}}
          {{if .Error}}<div class="cell-error">⚠️ Error</div>
          {{else}}
            {{if .HasReview}}<div class="cell-score" style="color:{{scoreColor .Score}}">{{.Score}}/10</div>{{end}}
            <div class="cell-build">{{buildIcon .BuildPass}}</div>
          {{end}}
        {{else}}<span style="color:#d1d5db">—</span>{{end}}
      </td>
      {{end}}
    </tr>
    {{end}}
  </tbody>
</table>
{{end}}

</body>
</html>`
