package report

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func TestWriteHTMLReport(t *testing.T) {
	dir := t.TempDir()

	boolTrue := true
	r := &EvalReport{
		PromptID:   "test-prompt",
		ConfigName: "baseline",
		Timestamp:  "2024-01-15T10:00:00Z",
		Duration:   12.5,
		PromptMeta: map[string]any{"service": "storage", "language": "dotnet"},
		ConfigUsed: map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles: []string{"Program.cs"},
		Review: &review.ReviewResult{
			Scores: review.ReviewScores{
				Criteria: []review.CriterionResult{
					{Name: "Code Builds", Passed: true, Reason: "Compiles successfully"},
					{Name: "Latest Packages", Passed: true, Reason: "Using latest versions"},
					{Name: "Best Practices", Passed: true, Reason: "Uses DefaultAzureCredential"},
					{Name: "Error Handling", Passed: false, Reason: "Missing retry logic"},
					{Name: "Code Quality", Passed: true, Reason: "Clean structure"},
				},
			},
			OverallScore: 4,
			MaxScore:     5,
			Summary:      "Good implementation",
			Issues:       []string{"Missing retry logic"},
			Strengths:    []string{"Clean code structure"},
		},
		GraderResults: []GraderResult{
			{GraderName: "test-grader", GraderType: "review", OverallScore: 4, MaxScore: 5, Summary: "Grader says good"},
			{GraderName: "consensus", GraderType: "review", OverallScore: 4, MaxScore: 5, Summary: "Consensus result", IsConsensus: true},
		},
		SessionEvents: []SessionEventRecord{
			{Type: "user.message", Content: "Write a dotnet storage auth sample"},
			{Type: "assistant.reasoning", Content: "I need to create an auth sample"},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"Program.cs"}`},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "File created", ToolSuccess: &boolTrue, Duration: 150.5},
			{Type: "assistant.message", Content: "Here is your sample"},
		},
		EventCount: 15,
		ToolCalls:  []string{"create_file", "edit_file"},
		Success:    true,
	}

	reportPath, err := WriteHTMLReport(r, dir, "20240115-100000", "storage", "data-plane", "dotnet", "authentication")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	checks := []string{
		"test-prompt",
		"baseline",
		"PASSED",
		"4/5",
		"Code Builds",
		"Good implementation",
		"Program.cs",
		"Generation Timeline",
		"Write a dotnet storage auth sample",
		"I need to create an auth sample",
		"Code Review",
		"Clean code structure",
		"Missing retry logic",
		"Tool call: create",
		"Back to Summary",
		"File created",
		"150ms",
		"Grader Results",
		"test-grader",
		"Grader says good",
		"Consensus result",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("HTML report missing %q", check)
		}
	}

	expectedDir := filepath.Join(dir, "20240115-100000", "results", "storage", "data-plane", "dotnet", "authentication", "test-prompt", "baseline")
	if _, err := os.Stat(expectedDir); err != nil {
		t.Errorf("expected directory %s to exist", expectedDir)
	}
}

func TestWriteHTMLReportNoReview(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:       "test-prompt",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       5.0,
		PromptMeta:     map[string]any{},
		ConfigUsed:     map[string]any{},
		GeneratedFiles: []string{},
		Success:        false,
		Error:          "timeout exceeded",
	}

	reportPath, err := WriteHTMLReport(r, dir, "run1", "svc", "plane", "lang", "cat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "FAILED") {
		t.Error("expected FAILED in report")
	}
	if !strings.Contains(content, "timeout exceeded") {
		t.Error("expected error message in report")
	}
}

func TestWriteSummaryHTML(t *testing.T) {
	dir := t.TempDir()

	s := &RunSummary{
		RunID:        "20240115-100000",
		Timestamp:    "2024-01-15T10:00:00Z",
		TotalPrompts: 2,
		TotalConfigs: 2,
		TotalEvals:   4,
		Passed:       3,
		Failed:       1,
		Errors:       0,
		Duration:     120.5,
		Results: []*EvalReport{
			{
				PromptID:   "prompt-a",
				ConfigName: "baseline",
				Success:    true,
				Review:     &review.ReviewResult{OverallScore: 4, MaxScore: 5},
			},
			{
				PromptID:   "prompt-a",
				ConfigName: "azure-mcp",
				Success:    true,
				Review:     &review.ReviewResult{OverallScore: 5, MaxScore: 5},
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "baseline",
				Success:    false,
			},
			{
				PromptID:   "prompt-b",
				ConfigName: "azure-mcp",
				Success:    true,
				Review:     &review.ReviewResult{OverallScore: 3, MaxScore: 5},
			},
		},
	}

	summaryPath, err := WriteSummaryHTML(s, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	content := string(data)
	checks := []string{
		"Evaluation Summary",
		"20240115-100000",
		"prompt-a",
		"prompt-b",
		"baseline",
		"azure-mcp",
		"4/5",
		"5/5",
		"3/5",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("summary HTML missing %q", check)
		}
	}

	// Verify the summary uses Success field: 3 passed should show ✅, 1 failed should show ❌
	// Matrix has 3 pass + 1 fail, detailed results table also has 3 pass + 1 fail
	passCount := strings.Count(content, "✅")
	failCount := strings.Count(content, "❌")
	if passCount != 6 {
		t.Errorf("expected 6 ✅ icons (3 matrix + 3 detail), got %d", passCount)
	}
	if failCount != 2 {
		t.Errorf("expected 2 ❌ icons (1 matrix + 1 detail), got %d", failCount)
	}
}

func TestWriteSummaryHTMLNoBuild(t *testing.T) {
	dir := t.TempDir()

	// Simulate results (no Build)
	s := &RunSummary{
		RunID:      "20240201-090000",
		Timestamp:  "2024-02-01T09:00:00Z",
		TotalEvals: 3,
		Passed:     2,
		Failed:     1,
		Duration:   60.0,
		Results: []*EvalReport{
			{PromptID: "p1", ConfigName: "baseline", Success: true},
			{PromptID: "p1", ConfigName: "mcp", Success: true, Review: &review.ReviewResult{OverallScore: 3, MaxScore: 5}},
			{PromptID: "p2", ConfigName: "baseline", Success: false},
		},
	}

	summaryPath, err := WriteSummaryHTML(s, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	content := string(data)
	// Matrix has 2 pass + 1 fail, detailed results table also has 2 pass + 1 fail
	passCount := strings.Count(content, "✅")
	failCount := strings.Count(content, "❌")
	if passCount != 4 {
		t.Errorf("expected 4 ✅ (2 matrix + 2 detail), got %d", passCount)
	}
	if failCount != 2 {
		t.Errorf("expected 2 ❌ (1 matrix + 1 detail), got %d", failCount)
	}
}

func TestBuildMatrix(t *testing.T) {
	s := &RunSummary{
		Results: []*EvalReport{
			{PromptID: "p1", ConfigName: "c1", Success: true, Review: &review.ReviewResult{OverallScore: 4, MaxScore: 5}},
			{PromptID: "p1", ConfigName: "c2", Success: false},
			{PromptID: "p2", ConfigName: "c1", Error: "timeout"},
		},
	}

	m := buildMatrix(s)

	if len(m.Prompts) != 2 {
		t.Errorf("expected 2 prompts, got %d", len(m.Prompts))
	}
	if len(m.Configs) != 2 {
		t.Errorf("expected 2 configs, got %d", len(m.Configs))
	}

	cell := m.Cells["p1"]["c1"]
	if cell == nil {
		t.Fatal("expected cell for p1/c1")
	}
	if cell.Score != 4 {
		t.Errorf("expected score 4, got %d", cell.Score)
	}
	if !cell.Success {
		t.Error("expected Success=true for p1/c1")
	}

	failCell := m.Cells["p1"]["c2"]
	if failCell == nil {
		t.Fatal("expected cell for p1/c2")
	}
	if failCell.Success {
		t.Error("expected Success=false for p1/c2")
	}

	errCell := m.Cells["p2"]["c1"]
	if errCell == nil {
		t.Fatal("expected cell for p2/c1")
	}
	if errCell.Error != "timeout" {
		t.Errorf("expected timeout error, got %q", errCell.Error)
	}
}

func TestBuildReportData(t *testing.T) {
	boolTrue := true
	r := &EvalReport{
		PromptID:       "test-prompt",
		GeneratedFiles: []string{"main.py", "requirements.txt"},
		SessionEvents: []SessionEventRecord{
			{Type: "session.start"},
			{Type: "user.message", Content: "Write a Python script"},
			{Type: "assistant.reasoning", Content: "I should create a script"},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"main.py"}`, MCPServerName: "fs-server"},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "File created successfully", ToolSuccess: &boolTrue, Duration: 42.5},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"requirements.txt"}`},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "File created successfully", ToolSuccess: &boolTrue, Duration: 10.0},
			{Type: "assistant.message", Content: "Done! Here are your files."},
		},
	}

	d := buildReportData(r)

	if d.Prompt != "Write a Python script" {
		t.Errorf("expected prompt from user.message, got %q", d.Prompt)
	}
	if d.Reasoning != "I should create a script" {
		t.Errorf("expected reasoning, got %q", d.Reasoning)
	}
	if d.FinalReply != "Done! Here are your files." {
		t.Errorf("expected final reply, got %q", d.FinalReply)
	}
	if len(d.ToolActions) != 2 {
		t.Errorf("expected 2 tool actions, got %d", len(d.ToolActions))
	}
	if d.ToolActions[0].Index != 1 || d.ToolActions[0].ToolName != "create" {
		t.Errorf("unexpected first tool action: %+v", d.ToolActions[0])
	}
	if d.ToolActions[0].Args != `{"path":"main.py"}` {
		t.Errorf("expected tool args, got %q", d.ToolActions[0].Args)
	}
	if d.ToolActions[0].MCPServer != "fs-server" {
		t.Errorf("expected MCP server 'fs-server', got %q", d.ToolActions[0].MCPServer)
	}
	if d.ToolActions[0].Result != "File created successfully" {
		t.Errorf("expected tool result from completion, got %q", d.ToolActions[0].Result)
	}
	if d.ToolActions[0].Success == nil || !*d.ToolActions[0].Success {
		t.Error("expected tool success=true")
	}
	if d.ToolActions[0].Duration != 42.5 {
		t.Errorf("expected duration 42.5, got %f", d.ToolActions[0].Duration)
	}
	if d.ToolActions[1].Result != "File created successfully" {
		t.Errorf("expected second tool result, got %q", d.ToolActions[1].Result)
	}
	if d.FileCount != 2 {
		t.Errorf("expected file count 2, got %d", d.FileCount)
	}

	// Verify timeline steps are built correctly
	// Expected: prompt, reasoning, tool_call, tool_call, message, complete = 6 steps
	if len(d.TimelineSteps) != 6 {
		t.Errorf("expected 6 timeline steps, got %d", len(d.TimelineSteps))
	}
	if len(d.TimelineSteps) >= 1 && d.TimelineSteps[0].StepType != "prompt" {
		t.Errorf("expected first step to be prompt, got %q", d.TimelineSteps[0].StepType)
	}
	if len(d.TimelineSteps) >= 2 && d.TimelineSteps[1].StepType != "reasoning" {
		t.Errorf("expected second step to be reasoning, got %q", d.TimelineSteps[1].StepType)
	}
	if len(d.TimelineSteps) >= 3 && d.TimelineSteps[2].StepType != "tool_call" {
		t.Errorf("expected third step to be tool_call, got %q", d.TimelineSteps[2].StepType)
	}
	if len(d.TimelineSteps) >= 3 && d.TimelineSteps[2].Duration != 42.5 {
		t.Errorf("expected tool_call duration 42.5, got %f", d.TimelineSteps[2].Duration)
	}
	if len(d.TimelineSteps) >= 6 && d.TimelineSteps[5].StepType != "complete" {
		t.Errorf("expected last step to be complete, got %q", d.TimelineSteps[5].StepType)
	}
}

func TestWriteHTMLReportPhaseTimings(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:           "timing-test",
		ConfigName:         "baseline",
		Timestamp:          "2024-01-15T10:00:00Z",
		Duration:           45.3,
		GenerationDuration: 20.1,
		ReviewDuration:     15.2,
		PromptMeta:         map[string]any{"service": "storage", "language": "go"},
		ConfigUsed:         map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles:     []string{"main.go"},
		Success:            true,
		Environment: &EnvironmentInfo{
			Model: "gpt-4",
		},
	}

	reportPath, err := WriteHTMLReport(r, dir, "20240115-100000", "storage", "data-plane", "go", "crud")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	for _, want := range []string{
		"Generation Duration",
		"Review Duration",
		"20.1s",
		"15.2s",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("HTML report missing phase timing %q", want)
		}
	}
}

func TestWriteSummaryHTMLAvgPhaseTimings(t *testing.T) {
	dir := t.TempDir()

	summary := &RunSummary{
		RunID:                 "20240115-100000",
		Timestamp:             "2024-01-15T10:00:00Z",
		TotalEvals:            2,
		Passed:                2,
		Duration:              90.0,
		AvgGenerationDuration: 18.5,
		AvgReviewDuration:     12.8,
		Results: []*EvalReport{
			{PromptID: "p1", ConfigName: "c1", Success: true, Duration: 45.0},
			{PromptID: "p2", ConfigName: "c1", Success: true, Duration: 45.0},
		},
	}

	_, err := WriteSummaryHTML(summary, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "20240115-100000", "summary.html"))
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	content := string(data)
	for _, want := range []string{
		"Avg Generation",
		"Avg Review",
		"18.5s",
		"12.8s",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("Summary HTML missing avg phase timing %q", want)
		}
	}
}

func TestEmbeddedTemplatesProduceValidHTML(t *testing.T) {
	// Verify the embedded report template executes and produces valid HTML structure.
	boolTrue := true
	data := buildReportData(&EvalReport{
		PromptID:       "embed-test",
		ConfigName:     "test-config",
		Timestamp:      "2024-06-01T12:00:00Z",
		Duration:       10.0,
		PromptMeta:     map[string]any{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "test-config"},
		GeneratedFiles: []string{"main.py"},
		Success:        true,
		Review:         &review.ReviewResult{OverallScore: 3, MaxScore: 5, Scores: review.ReviewScores{Criteria: []review.CriterionResult{{Name: "Builds", Passed: true, Reason: "ok"}}}},
		SessionEvents: []SessionEventRecord{
			{Type: "user.message", Content: "test prompt"},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"main.py"}`},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "ok", ToolSuccess: &boolTrue, Duration: 50},
		},
	})
	data.BackPath = "../summary.html"

	var buf bytes.Buffer
	if err := parsedReportTemplate.Execute(&buf, data); err != nil {
		t.Fatalf("report template execution failed: %v", err)
	}
	html := buf.String()
	if !strings.HasPrefix(html, "<!DOCTYPE html>") {
		t.Error("report template missing DOCTYPE")
	}
	if !strings.Contains(html, "</html>") {
		t.Error("report template missing closing </html>")
	}

	// Verify the embedded summary template executes and produces valid HTML structure.
	summary := &RunSummary{
		RunID:      "embed-test-run",
		Timestamp:  "2024-06-01T12:00:00Z",
		TotalEvals: 1, Passed: 1, Duration: 10.0,
		Results: []*EvalReport{{PromptID: "p1", ConfigName: "c1", Success: true, Duration: 5.0}},
	}
	matrix := buildMatrix(summary)
	stats := ComputeSummaryStats(summary)
	summaryData := struct {
		Summary *RunSummary
		Matrix  *MatrixData
		Stats   *SummaryStats
	}{Summary: summary, Matrix: matrix, Stats: stats}

	buf.Reset()
	if err := parsedSummaryTemplate.Execute(&buf, summaryData); err != nil {
		t.Fatalf("summary template execution failed: %v", err)
	}
	html = buf.String()
	if !strings.HasPrefix(html, "<!DOCTYPE html>") {
		t.Error("summary template missing DOCTYPE")
	}
	if !strings.Contains(html, "</html>") {
		t.Error("summary template missing closing </html>")
	}
}

func TestTemplateFS(t *testing.T) {
	// Verify the embedded filesystem contains the expected template files.
	entries, err := templateFS.ReadDir("templates")
	if err != nil {
		t.Fatalf("failed to read embedded templates dir: %v", err)
	}
	names := make(map[string]bool)
	for _, e := range entries {
		names[e.Name()] = true
	}
	for _, want := range []string{"report.gohtml", "summary.gohtml"} {
		if !names[want] {
			t.Errorf("embedded templates missing %q", want)
		}
	}
}

func TestWriteHTMLReportGraderDetails(t *testing.T) {
	dir := t.TempDir()
	boolTrue := true
	boolFalse := false
	patternTrue := true

	r := &EvalReport{
		PromptID:       "grader-detail-test",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-20T10:00:00Z",
		Duration:       30.0,
		PromptMeta:     map[string]any{"service": "storage", "language": "python"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.py"},
		Success:        true,
		GraderResults: []GraderResult{
			{
				GraderName: "main-exists",
				GraderType: "file",
				Summary:    "Checking required files",
				Score:      1.0,
				Weight:     0.5,
				Pass:       &boolTrue,
				Gate:       true,
				FileDetails: &FileGraderDetail{
					CheckedFiles: []FileCheckDetail{
						{Path: "main.py", Exists: true, PatternMatched: &patternTrue, Pattern: "import azure"},
						{Path: "tests.py", Exists: false},
					},
				},
			},
			{
				GraderName: "build-check",
				GraderType: "program",
				Summary:    "Build verification",
				Score:      0.0,
				Weight:     1.0,
				Pass:       &boolFalse,
				Gate:       true,
				ProgramDetails: &ProgramGraderDetail{
					Command:  "python -m py_compile main.py",
					ExitCode: 1,
					Stdout:   "Compiling...",
					Stderr:   "SyntaxError: unexpected EOF",
				},
			},
			{
				GraderName: "quality-review",
				GraderType: "prompt",
				Summary:    "LLM quality assessment",
				Score:      0.75,
				Weight:     1.0,
				Pass:       &boolTrue,
				PromptDetails: &PromptGraderDetail{
					Model:     "claude-opus-4.6",
					Rubric:    "Evaluate code quality on a 1-5 scale",
					Reasoning: "The code demonstrates good structure but lacks error handling",
					RawScore:  4,
					MaxScore:  5,
				},
			},
			{
				GraderName: "tool-usage",
				GraderType: "behavior",
				Summary:    "Agent behavior analysis",
				Score:      0.5,
				Weight:     0.5,
				Pass:       &boolFalse,
				BehaviorDetails: &BehaviorGraderDetail{
					ToolsUsed:     []string{"create", "edit"},
					MissingTools:  []string{"azure-mcp"},
					ForbiddenUsed: []string{"rm"},
					TotalActions:  15,
					MaxTurns:      25,
					ActualTurns:   12,
					Violations:    []string{"Used forbidden tool: rm"},
					ToolCounts:    map[string]int{"create": 3, "edit": 5, "rm": 1},
				},
			},
		},
	}

	reportPath, err := WriteHTMLReport(r, dir, "20240120-100000", "storage", "data-plane", "python", "crud")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	content := string(data)
	checks := []string{
		// Section header
		"Grader Results",
		"4 graders",

		// File grader
		"main-exists",
		"file",
		"main.py",
		"exists",
		"tests.py",
		"missing",
		"import azure",
		"GATE PASSED",

		// Program grader
		"build-check",
		"program",
		"python -m py_compile main.py",
		"Exit code",
		"Show stdout",
		"Compiling...",
		"Show stderr",
		"SyntaxError: unexpected EOF",
		"GATE FAILED",

		// Prompt grader
		"quality-review",
		"claude-opus-4.6",
		"Show rubric",
		"Show LLM reasoning",
		"good structure but lacks error handling",
		"4/5",

		// Behavior grader
		"tool-usage",
		"behavior",
		"azure-mcp",
		"Used forbidden tool: rm",
		"Tool call counts",

		// Score bar
		"score-bar",
		"score-bar-fill",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("HTML report missing %q", check)
		}
	}
}

func TestActionTimelineHTMLRendering(t *testing.T) {
	boolTrue := true
	boolFalse := false

	r := &EvalReport{
		PromptID:       "timeline-test",
		ConfigName:     "test-config",
		Timestamp:      "2024-06-01T12:00:00Z",
		Duration:       30.0,
		PromptMeta:     map[string]any{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "test-config"},
		GeneratedFiles: []string{"main.py"},
		Success:        true,
		SessionEvents: []SessionEventRecord{
			{Type: "user.message", Content: "test prompt"},
			{Type: "assistant.reasoning", Content: "thinking"},
			{Type: "tool.execution_start", ToolName: "create", ToolArgs: `{"path":"main.py"}`},
			{Type: "tool.execution_complete", ToolName: "create", ToolResult: "ok", ToolSuccess: &boolTrue, Duration: 150},
			{Type: "assistant.message", Content: "Done"},
		},
		ActionTimeline: BuildActionTimeline([]SessionEventRecord{
			{Type: "assistant.turn_start"},
			{Type: "assistant.reasoning"},
			{Type: "tool.execution_start", ToolName: "create", MCPServerName: "fs-server", FilePath: "main.py"},
			{Type: "tool.execution_complete", ToolName: "create", ToolSuccess: &boolTrue, Duration: 150},
			{Type: "tool.execution_start", ToolName: "bash", FilePath: "/workspace/build.sh"},
			{Type: "tool.execution_complete", ToolName: "bash", ToolSuccess: &boolFalse, Duration: 2500, Error: "build failed"},
			{Type: "assistant.intent", Intent: "fixing imports"},
			{Type: "assistant.message"},
			{Type: "assistant.turn_end"},
		}),
	}

	data := buildReportData(r)
	data.BackPath = "../summary.html"

	var buf bytes.Buffer
	if err := parsedReportTemplate.Execute(&buf, data); err != nil {
		t.Fatalf("template execution failed: %v", err)
	}
	html := buf.String()

	checks := []string{
		"Action Timeline",
		"atl-summary-bar",
		"atl-search",
		"atl-list",
		"atl-row",
		"tool calls",
		"turns",
		"tool duration",
		"succeeded",
		"create",
		"bash",
		"fs-server",
		"150ms",
		"2.5s",
		"atl-success",
		"atl-failure",
		"fixing imports",
		"build failed",
		"Filter timeline",
	}
	for _, check := range checks {
		if !strings.Contains(html, check) {
			t.Errorf("action timeline HTML missing %q", check)
		}
	}
}

func TestActionTimelineHTMLOmittedWhenNil(t *testing.T) {
	r := &EvalReport{
		PromptID:       "no-timeline",
		ConfigName:     "test-config",
		Timestamp:      "2024-06-01T12:00:00Z",
		Duration:       5.0,
		PromptMeta:     map[string]any{},
		ConfigUsed:     map[string]any{},
		GeneratedFiles: []string{},
		Success:        true,
	}

	data := buildReportData(r)
	data.BackPath = "../summary.html"

	var buf bytes.Buffer
	if err := parsedReportTemplate.Execute(&buf, data); err != nil {
		t.Fatalf("template execution failed: %v", err)
	}
	html := buf.String()

	if strings.Contains(html, "atl-summary-bar") {
		t.Error("action timeline section should not render when ActionTimeline is nil")
	}
}
