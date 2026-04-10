package report

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

//go:embed templates/*.gohtml
var templateFS embed.FS

// Parsed once at package init for reuse across calls.
var (
	parsedReportTemplate  = template.Must(template.New("report.gohtml").Funcs(htmlFuncMap()).ParseFS(templateFS, "templates/report.gohtml"))
	parsedSummaryTemplate = template.Must(template.New("summary.gohtml").Funcs(htmlFuncMap()).ParseFS(templateFS, "templates/summary.gohtml"))
)

// WriteHTMLReport writes an individual evaluation report as HTML.
func WriteHTMLReport(r *EvalReport, outputDir string, runID string, service, plane, language, category string) (string, error) {
	reportDir := filepath.Join(
		outputDir, runID, "results",
		service, plane, language, category, r.PromptID, r.ConfigName,
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

	data := buildReportData(r)

	// Compute back-to-summary link dynamically: count path segments from
	// runID dir to reportDir. ConfigName may contain '/' (e.g. "azure-mcp/claude-opus-4.6").
	// Segments: results/service/plane/language/category/promptID + configName parts.
	depth := 7 + strings.Count(r.ConfigName, "/")
	data.BackPath = strings.Repeat("../", depth) + "summary.html"

	// Read file contents from the generated-code directory for expandable display
	codeDir := filepath.Join(reportDir, "generated-code")
	data.FileContents = readFileContents(codeDir, r.GeneratedFiles, r.StarterFiles)

	if err := parsedReportTemplate.Execute(f, data); err != nil {
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

	// Sort results by prompt ID so same-prompt entries (different configs) are adjacent
	sort.Slice(s.Results, func(i, j int) bool {
		if s.Results[i].PromptID != s.Results[j].PromptID {
			return s.Results[i].PromptID < s.Results[j].PromptID
		}
		return s.Results[i].ConfigName < s.Results[j].ConfigName
	})

	matrix := buildMatrix(s)
	stats := ComputeSummaryStats(s)

	data := struct {
		Summary *RunSummary
		Matrix  *MatrixData
		Stats   *SummaryStats
	}{
		Summary: s,
		Matrix:  matrix,
		Stats:   stats,
	}

	if err := parsedSummaryTemplate.Execute(f, data); err != nil {
		return "", fmt.Errorf("executing summary template: %w", err)
	}

	return summaryPath, nil
}

// MatrixData holds the cross-config comparison matrix.
type MatrixData struct {
	Configs []string
	Prompts []string
	Cells   map[string]map[string]*MatrixCell // [promptID][configName]
}

// MatrixCell holds the data for one cell in the matrix.
type MatrixCell struct {
	Success    bool
	Score      int
	MaxScore   int
	HasReview  bool
	Duration   float64
	Error      string
	FileCount  int
	ToolCalls  []string
	ReportLink string
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
			Success:   r.Success,
			Duration:  r.Duration,
			Error:     r.Error,
			FileCount: len(r.GeneratedFiles),
			ToolCalls: r.ToolCalls,
		}
		if r.Review != nil {
			cell.Score = r.Review.OverallScore
			cell.MaxScore = r.Review.MaxScore
			cell.HasReview = true
		}
		// Build relative link from summary.html to individual report
		service, _ := r.PromptMeta["service"].(string)
		plane, _ := r.PromptMeta["plane"].(string)
		language, _ := r.PromptMeta["language"].(string)
		category, _ := r.PromptMeta["category"].(string)
		if service != "" && plane != "" && language != "" && category != "" {
			cell.ReportLink = strings.Join([]string{"results", service, plane, language, category, r.PromptID, r.ConfigName, "report.html"}, "/")
		}
		m.Cells[r.PromptID][r.ConfigName] = cell
	}

	return m
}

// ReportTemplateData is the enriched data passed to the individual report template.
type ReportTemplateData struct {
	*EvalReport
	Prompt         string
	Reasoning      string
	FinalReply     string
	ToolActions    []ToolAction
	TimelineSteps  []TimelineStep
	ReviewTimeline []TimelineStep            // timeline for consolidated review
	PanelTimelines map[string][]TimelineStep // model name → timeline for each panel reviewer
	FileCount      int
	FileContents   map[string]string // filename → content for expandable display
	BackPath       string            // relative path from report.html back to summary.html
}

// ToolAction represents one tool invocation extracted from session events.
type ToolAction struct {
	Index     int
	ToolName  string
	Args      string
	Result    string
	Error     string
	Success   *bool
	Duration  float64
	MCPServer string
}

// TimelineStep represents one chronological step in the agent workflow.
type TimelineStep struct {
	Index     int
	Phase     string // "generation", "review"
	StepType  string // "prompt", "reasoning", "tool_call", "message", "complete"
	Icon      string
	Title     string
	Content   string  // main content (tool result, reasoning text)
	Detail    string  // collapsible detail (tool args, full text)
	Duration  float64 // milliseconds
	Success   *bool
	Error     string
	ToolName  string
	MCPServer string
}

// buildReportData extracts structured sections from session events and
// builds a chronological timeline of agent steps.
func buildReportData(r *EvalReport) *ReportTemplateData {
	d := &ReportTemplateData{
		EvalReport: r,
		FileCount:  len(r.GeneratedFiles),
	}

	var reasoningParts []string
	var messageParts []string
	stepIndex := 0

	type pendingTool struct {
		stepIdx int
		name    string
	}
	var pendingTools []pendingTool

	for _, ev := range r.SessionEvents {
		switch ev.Type {
		case "user.message":
			if d.Prompt == "" && ev.Content != "" {
				d.Prompt = ev.Content
			}
			if ev.Content != "" {
				stepIndex++
				d.TimelineSteps = append(d.TimelineSteps, TimelineStep{
					Index:    stepIndex,
					Phase:    "generation",
					StepType: "prompt",
					Icon:     "📝",
					Title:    "Prompt sent",
					Content:  ev.Content,
				})
			}
		case "assistant.reasoning":
			if ev.Content != "" {
				reasoningParts = append(reasoningParts, ev.Content)
			}
			stepIndex++
			title := ev.Content
			if title == "" {
				title = "(thinking)"
			} else if len(title) > 80 {
				title = title[:80] + "…"
			}
			d.TimelineSteps = append(d.TimelineSteps, TimelineStep{
				Index:    stepIndex,
				Phase:    "generation",
				StepType: "reasoning",
				Icon:     "🤔",
				Title:    title,
				Content:  ev.Content,
			})
		case "tool.execution_start":
			toolName := ev.ToolName
			if toolName == "" {
				toolName = "(unknown)"
			}
			d.ToolActions = append(d.ToolActions, ToolAction{
				Index:     len(d.ToolActions) + 1,
				ToolName:  toolName,
				Args:      ev.ToolArgs,
				MCPServer: ev.MCPServerName,
			})
			stepIndex++
			toolTitle := toolName
			icon := "🔧"
			stepType := "tool_call"
			if ev.FilePath != "" {
				toolTitle += " → " + ev.FilePath
			}
			// Distinguish skill-related tool calls with dedicated icons
			if toolName == "skill" {
				icon = "📚"
				stepType = "skill"
				// Extract skill name from args JSON
				if skillArg := extractJSONField(ev.ToolArgs, "skill"); skillArg != "" {
					toolTitle = "Skill invoked: " + skillArg
				} else {
					toolTitle = "Skill invoked"
				}
			} else if toolName == "view" && strings.Contains(ev.ToolArgs, "skills") && strings.Contains(ev.ToolArgs, "references") {
				icon = "📖"
				stepType = "skill_ref"
				toolTitle = "Skill reference fetch"
				if ev.FilePath != "" {
					toolTitle += " → " + filepath.Base(ev.FilePath)
				} else if refFile := extractJSONField(ev.ToolArgs, "path"); refFile != "" {
					toolTitle += " → " + filepath.Base(refFile)
				}
			}
			step := TimelineStep{
				Index:     stepIndex,
				Phase:     "generation",
				StepType:  stepType,
				Icon:      icon,
				Title:     toolTitle,
				Detail:    ev.ToolArgs,
				ToolName:  toolName,
				MCPServer: ev.MCPServerName,
			}
			d.TimelineSteps = append(d.TimelineSteps, step)
			pendingTools = append(pendingTools, pendingTool{len(d.TimelineSteps) - 1, toolName})
		case "tool.execution_complete":
			// Update ToolActions (backward compat)
			for i := len(d.ToolActions) - 1; i >= 0; i-- {
				if d.ToolActions[i].ToolName == ev.ToolName && d.ToolActions[i].Result == "" && d.ToolActions[i].Error == "" {
					d.ToolActions[i].Result = ev.ToolResult
					d.ToolActions[i].Error = ev.Error
					d.ToolActions[i].Success = ev.ToolSuccess
					d.ToolActions[i].Duration = ev.Duration
					break
				}
			}
			// Update matching timeline step
			for i := len(pendingTools) - 1; i >= 0; i-- {
				if pendingTools[i].name == ev.ToolName {
					idx := pendingTools[i].stepIdx
					d.TimelineSteps[idx].Content = ev.ToolResult
					d.TimelineSteps[idx].Duration = ev.Duration
					d.TimelineSteps[idx].Success = ev.ToolSuccess
					if ev.Error != "" {
						d.TimelineSteps[idx].Error = ev.Error
					}
					pendingTools = append(pendingTools[:i], pendingTools[i+1:]...)
					break
				}
			}
		case "assistant.message":
			if ev.Content != "" {
				messageParts = append(messageParts, ev.Content)
			}
			stepIndex++
			title := "Agent reply"
			d.TimelineSteps = append(d.TimelineSteps, TimelineStep{
				Index:    stepIndex,
				Phase:    "generation",
				StepType: "message",
				Icon:     "💬",
				Title:    title,
				Content:  ev.Content,
			})
		case "skill.invoked":
			skillName := ev.SkillName
			if skillName == "" && ev.Content != "" {
				for _, line := range strings.Split(ev.Content, "\n") {
					line = strings.TrimSpace(line)
					if strings.HasPrefix(line, "name:") {
						skillName = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
						break
					}
				}
			}
			// Merge into the existing skill tool call timeline step if present,
			// rather than creating a duplicate entry.
			merged := false
			for i := len(d.TimelineSteps) - 1; i >= 0; i-- {
				if d.TimelineSteps[i].StepType == "skill" {
					d.TimelineSteps[i].Content = truncateStr(ev.Content, 2000)
					if skillName != "" {
						d.TimelineSteps[i].Title = "Skill loaded: " + skillName
					}
					merged = true
					break
				}
			}
			if !merged && (skillName != "" || ev.Content != "") {
				if skillName == "" {
					skillName = "(unknown)"
				}
				stepIndex++
				d.TimelineSteps = append(d.TimelineSteps, TimelineStep{
					Index:    stepIndex,
					Phase:    "generation",
					StepType: "skill",
					Icon:     "📚",
					Title:    "Skill loaded: " + skillName,
					Content:  truncateStr(ev.Content, 2000),
				})
			}
		}
	}

	// Add generation-complete step
	if len(d.TimelineSteps) > 0 {
		stepIndex++
		summary := fmt.Sprintf("%d files created", d.FileCount)
		d.TimelineSteps = append(d.TimelineSteps, TimelineStep{
			Index:    stepIndex,
			Phase:    "generation",
			StepType: "complete",
			Icon:     "✅",
			Title:    "Generation complete",
			Content:  summary,
		})
	}

	d.Reasoning = strings.Join(reasoningParts, "\n\n")
	d.FinalReply = strings.Join(messageParts, "\n\n")

	// Build review timelines from review events
	if r.Review != nil && len(r.Review.Events) > 0 {
		d.ReviewTimeline = buildReviewTimeline(r.Review.Events)
	}
	if len(r.ReviewPanel) > 0 {
		d.PanelTimelines = make(map[string][]TimelineStep, len(r.ReviewPanel))
		for _, pr := range r.ReviewPanel {
			if len(pr.Events) > 0 {
				d.PanelTimelines[pr.Model] = buildReviewTimeline(pr.Events)
			}
		}
	}

	return d
}

// buildReviewTimeline converts review events into a chronological timeline
// matching the generator timeline format.
func buildReviewTimeline(events []review.ReviewEvent) []TimelineStep {
	var steps []TimelineStep
	stepIndex := 0

	type pendingTool struct {
		stepIdx int
		name    string
	}
	var pendingTools []pendingTool

	for _, ev := range events {
		switch ev.Type {
		case "assistant.turn_start":
			// Skip — implicit from other events
		case "assistant.reasoning":
			if ev.Content != "" {
				stepIndex++
				title := ev.Content
				if len(title) > 80 {
					title = title[:80] + "…"
				}
				steps = append(steps, TimelineStep{
					Index:    stepIndex,
					Phase:    "review",
					StepType: "reasoning",
					Icon:     "🤔",
					Title:    title,
					Content:  ev.Content,
				})
			}
		case "tool.execution_start":
			if ev.ToolName != "" {
				stepIndex++
				steps = append(steps, TimelineStep{
					Index:    stepIndex,
					Phase:    "review",
					StepType: "tool_call",
					Icon:     "🔧",
					Title:    "Tool call: " + ev.ToolName,
					Detail:   ev.ToolArgs,
					ToolName: ev.ToolName,
				})
				pendingTools = append(pendingTools, pendingTool{len(steps) - 1, ev.ToolName})
			}
		case "tool.execution_complete":
			matched := false
			for i := len(pendingTools) - 1; i >= 0; i-- {
				if pendingTools[i].name == ev.ToolName {
					idx := pendingTools[i].stepIdx
					steps[idx].Content = ev.Result
					steps[idx].Duration = ev.Duration
					if ev.Error != "" {
						steps[idx].Error = ev.Error
					}
					pendingTools = append(pendingTools[:i], pendingTools[i+1:]...)
					matched = true
					break
				}
			}
			if !matched && ev.ToolName != "" {
				stepIndex++
				steps = append(steps, TimelineStep{
					Index:    stepIndex,
					Phase:    "review",
					StepType: "tool_call",
					Icon:     "🔧",
					Title:    "Tool call: " + ev.ToolName,
					Content:  ev.Result,
					Duration: ev.Duration,
					ToolName: ev.ToolName,
					Error:    ev.Error,
				})
			}
		case "assistant.message":
			if ev.Content != "" {
				stepIndex++
				steps = append(steps, TimelineStep{
					Index:    stepIndex,
					Phase:    "review",
					StepType: "message",
					Icon:     "💬",
					Title:    "Reviewer response",
					Content:  ev.Content,
				})
			}
		}
	}
	return steps
}

// readFileContents reads file contents from the code directory for display in the HTML report.
// If starterFiles is non-empty, only files NOT in the starter set are included.
// extractJSONField extracts a string field from a JSON args string.
func extractJSONField(jsonStr, field string) string {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		return ""
	}
	if v, ok := m[field].(string); ok {
		return v
	}
	return ""
}

func readFileContents(codeDir string, files []string, starterFiles []string) map[string]string {
	contents := make(map[string]string)
	starterSet := make(map[string]bool, len(starterFiles))
	for _, f := range starterFiles {
		starterSet[f] = true
	}
	for _, f := range files {
		if len(starterFiles) > 0 && starterSet[f] {
			continue // skip unchanged starter project files
		}
		path := filepath.Join(codeDir, f)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		if len(data) > 512*1024 {
			contents[f] = "(file too large to display)"
			continue
		}
		contents[f] = string(data)
	}
	return contents
}

func htmlFuncMap() template.FuncMap {
	return template.FuncMap{
		"scoreColor": func(passed, total int) string {
			if total == 0 {
				return "#6b7280" // gray
			}
			rate := float64(passed) / float64(total)
			switch {
			case rate >= 1.0:
				return "#22c55e" // green — all passed
			case rate >= 0.75:
				return "#eab308" // yellow
			case rate >= 0.5:
				return "#f97316" // orange
			default:
				return "#ef4444" // red
			}
		},
		"criterionIcon": func(passed bool) string {
			if passed {
				return "✅"
			}
			return "❌"
		},
		"statusIcon": func(success bool) string {
			if success {
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
		"fmtDurationMs": func(ms float64) string {
			if ms >= 1000 {
				return fmt.Sprintf("%.1fs", ms/1000)
			}
			return fmt.Sprintf("%.0fms", ms)
		},
		"truncate": func(s string, n int) string {
			if len(s) <= n {
				return s
			}
			return s[:n] + "…"
		},
		"isReviewLine": func(line string) bool {
			trimmed := strings.TrimSpace(line)
			return strings.Contains(trimmed, "REVIEW:")
		},
		"highlightReviewLines": func(content string) template.HTML {
			lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
			var b strings.Builder
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if strings.Contains(trimmed, "REVIEW:") {
					b.WriteString(`<span class="review-comment">`)
					b.WriteString(template.HTMLEscapeString(line))
					b.WriteString("</span>\n")
				} else {
					b.WriteString(template.HTMLEscapeString(line))
					b.WriteString("\n")
				}
			}
			return template.HTML(b.String())
		},
		"langClass": func(filename string) string {
			ext := filepath.Ext(filename)
			switch ext {
			case ".py":
				return "python"
			case ".cs":
				return "csharp"
			case ".go":
				return "go"
			case ".js":
				return "javascript"
			case ".ts":
				return "typescript"
			case ".java":
				return "java"
			case ".json":
				return "json"
			case ".yaml", ".yml":
				return "yaml"
			case ".xml":
				return "xml"
			case ".md":
				return "markdown"
			case ".sh":
				return "bash"
			case ".ps1":
				return "powershell"
			default:
				return ""
			}
		},
		"hasPrefix": strings.HasPrefix,
		"contains":  strings.Contains,
		"fileTypeSummary": func(files []string) string {
			counts := make(map[string]int)
			for _, f := range files {
				ext := filepath.Ext(f)
				if ext == "" {
					ext = "(no ext)"
				}
				counts[ext]++
			}
			type extCount struct {
				ext   string
				count int
			}
			var sorted []extCount
			for ext, count := range counts {
				sorted = append(sorted, extCount{ext, count})
			}
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].count > sorted[j].count
			})
			var parts []string
			shown := 0
			other := 0
			for _, ec := range sorted {
				if shown < 3 {
					parts = append(parts, fmt.Sprintf("%d %s", ec.count, ec.ext))
					shown++
				} else {
					other += ec.count
				}
			}
			if other > 0 {
				parts = append(parts, fmt.Sprintf("%d other", other))
			}
			return strings.Join(parts, ", ")
		},
		"boolStr": func(b *bool) string {
			if b == nil {
				return ""
			}
			if *b {
				return "✅"
			}
			return "❌"
		},
		"derefBool": func(b *bool) bool {
			return b != nil && *b
		},
		"mul": func(a, b float64) float64 {
			return a * b
		},
		"reportLink": func(r *EvalReport) string {
			service, _ := r.PromptMeta["service"].(string)
			plane, _ := r.PromptMeta["plane"].(string)
			language, _ := r.PromptMeta["language"].(string)
			category, _ := r.PromptMeta["category"].(string)
			if service == "" || plane == "" || language == "" || category == "" {
				return ""
			}
			// Use forward slashes for HTML links (filepath.Join uses OS-native backslashes on Windows)
			return strings.Join([]string{"results", service, plane, language, category, r.PromptID, r.ConfigName, "report.html"}, "/")
		},
		"impactFmt": func(v float64) string {
			if v > 0 {
				return fmt.Sprintf("+%.1f", v)
			}
			return fmt.Sprintf("%.1f", v)
		},
		"impactColor": func(v float64) string {
			if v > 0 {
				return "var(--green)"
			}
			if v < 0 {
				return "var(--red)"
			}
			return "var(--gray)"
		},
	}
}


