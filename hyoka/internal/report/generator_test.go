package report

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/internal/review"
)

func TestWriteReport(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		PromptID:       "test-prompt",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       12.5,
		PromptMeta:     map[string]any{"service": "storage", "language": "dotnet"},
		ConfigUsed:     map[string]any{"name": "baseline", "model": "gpt-4"},
		GeneratedFiles: []string{"Program.cs", "Storage.csproj"},
		EventCount: 15,
		ToolCalls:  []string{"create_file", "edit_file"},
		Success:    true,
	}

	p := &prompt.Prompt{
		ID:         "test-prompt",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "dotnet", "category": "authentication"},
	}








	reportPath, err := WriteReport(r, dir, "20240115-100000", p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(reportPath); err != nil {
		t.Fatalf("report file does not exist: %v", err)
	}

	// Verify JSON is valid
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON in report: %v", err)
	}

	if parsed.PromptID != "test-prompt" {
		t.Errorf("expected prompt ID 'test-prompt', got %q", parsed.PromptID)
	}
	if !parsed.Success {
		t.Error("expected success to be true")
	}
	if parsed.SchemaVersion != CurrentSchemaVersion {
		t.Errorf("expected schema version %d, got %d", CurrentSchemaVersion, parsed.SchemaVersion)
	}

	// Verify directory structure — includes prompt ID for workspace isolation
	expectedDir := filepath.Join(dir, "20240115-100000", "results", "storage", "data-plane", "dotnet", "authentication", "test-prompt", "baseline")
	if _, err := os.Stat(expectedDir); err != nil {
		t.Errorf("expected directory %s to exist", expectedDir)
	}
}

func TestWriteSummary(t *testing.T) {
	dir := t.TempDir()

	s := &RunSummary{
		RunID:        "20240115-100000",
		Timestamp:    "2024-01-15T10:00:00Z",
		TotalPrompts: 5,
		TotalConfigs: 2,
		TotalEvals:   10,
		Passed:       8,
		Failed:       1,
		Errors:       1,
		Duration:     120.5,
		Reports:      []string{"/path/to/report1.json", "/path/to/report2.json"},
	}

	summaryPath, err := WriteSummary(s, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("failed to read summary: %v", err)
	}

	var parsed RunSummary
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON in summary: %v", err)
	}

	if parsed.TotalEvals != 10 {
		t.Errorf("expected 10 total evals, got %d", parsed.TotalEvals)
	}
	if parsed.Passed != 8 {
		t.Errorf("expected 8 passed, got %d", parsed.Passed)
	}
}

func TestWriteReportInvalidDir(t *testing.T) {
	r := &EvalReport{PromptID: "test", ConfigName: "cfg"}
	p := &prompt.Prompt{ID: "test", Properties: map[string]string{"service": "svc", "plane": "plane", "language": "lang", "category": "cat"}}

	// Use a path containing characters that are invalid on both Unix and Windows.
	// On Windows, /nonexistent is treated as drive-relative and MkdirAll may succeed.
	invalidDir := filepath.Join(t.TempDir(), "not\x00valid")
	_, err := WriteReport(r, invalidDir, "run1", p)
	if err == nil {
		t.Fatal("expected error for invalid directory")
	}
}

func TestGraderResultsRoundTrip(t *testing.T) {
	dir := t.TempDir()

	graders := []GraderResult{
		{
			GraderName:   "claude-opus-4.6",
			GraderType:   "review",
			Model:        "claude-opus-4.6",
			OverallScore: 4,
			MaxScore:     5,
			Summary:      "Good code",
			Issues:       []string{"missing retry"},
			Strengths:    []string{"clean design"},
		},
		{
			GraderName:   "consensus",
			GraderType:   "review",
			OverallScore: 4,
			MaxScore:     5,
			Summary:      "Consensus result",
			IsConsensus:  true,
		},
	}

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "grader-test",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       10.0,
		PromptMeta:     map[string]any{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.go"},
		GraderResults:  graders,
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "grader-test",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
	}

	reportPath, err := WriteReport(r, dir, "run1", p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if parsed.SchemaVersion != CurrentSchemaVersion {
		t.Errorf("expected schema version %d, got %d", CurrentSchemaVersion, parsed.SchemaVersion)
	}
	if len(parsed.GraderResults) != 2 {
		t.Fatalf("expected 2 grader results, got %d", len(parsed.GraderResults))
	}
	if parsed.GraderResults[0].GraderName != "claude-opus-4.6" {
		t.Errorf("expected grader name claude-opus-4.6, got %q", parsed.GraderResults[0].GraderName)
	}
	if !parsed.GraderResults[1].IsConsensus {
		t.Error("expected second grader result to be consensus")
	}
}

func TestGraderResultsFromReviewPanel(t *testing.T) {
	consolidated := &review.ReviewResult{
		Model:        "consolidator",
		OverallScore: 3,
		MaxScore:     5,
		Summary:      "Consensus summary",
		Scores: review.ReviewScores{
			Criteria: []review.CriterionResult{
				{Name: "Builds", Passed: true, Reason: "OK"},
			},
		},
	}
	panel := []review.ReviewResult{
		{Model: "model-a", OverallScore: 4, MaxScore: 5, Summary: "Model A says good"},
		{Model: "model-b", OverallScore: 2, MaxScore: 5, Summary: "Model B says bad"},
	}

	results := GraderResultsFromReview(consolidated, panel)
	if len(results) != 3 {
		t.Fatalf("expected 3 grader results (2 panel + 1 consensus), got %d", len(results))
	}
	if results[0].GraderName != "model-a" {
		t.Errorf("expected first grader model-a, got %q", results[0].GraderName)
	}
	if !results[2].IsConsensus {
		t.Error("expected last grader result to be consensus")
	}
}

func TestGraderResultsFromSingleReview(t *testing.T) {
	single := &review.ReviewResult{
		Model:        "claude-sonnet",
		OverallScore: 5,
		MaxScore:     5,
		Summary:      "Perfect",
	}

	results := GraderResultsFromReview(single, nil)
	if len(results) != 1 {
		t.Fatalf("expected 1 grader result, got %d", len(results))
	}
	if results[0].IsConsensus {
		t.Error("single reviewer should not be marked as consensus")
	}
}

func TestBuildActionTimeline(t *testing.T) {
	boolTrue := true
	boolFalse := false
	events := []SessionEventRecord{
		{Type: "assistant.turn_start"},
		{Type: "assistant.reasoning"},
		{Type: "tool.execution_start", ToolName: "create_file", FilePath: "main.go"},
		{Type: "tool.execution_complete", ToolName: "create_file", Duration: 150.0, ToolSuccess: &boolTrue},
		{Type: "tool.execution_start", ToolName: "edit_file", MCPServerName: "editor"},
		{Type: "tool.execution_complete", ToolName: "edit_file", Duration: 80.0, ToolSuccess: &boolFalse, Error: "conflict"},
		{Type: "assistant.intent", Intent: "fixing imports"},
		{Type: "assistant.message"},
		{Type: "assistant.turn_end"},
	}

	tl := BuildActionTimeline(events)
	if tl == nil {
		t.Fatal("expected non-nil timeline")
	}

	// Verify entries (tool.execution_complete updates existing entries, doesn't create new ones)
	if len(tl.Entries) != 7 {
		t.Fatalf("expected 7 entries, got %d", len(tl.Entries))
	}

	// Verify turn_start
	if tl.Entries[0].Type != "turn_start" || tl.Entries[0].TurnNumber != 1 {
		t.Errorf("entry 0: expected turn_start turn=1, got %s turn=%d", tl.Entries[0].Type, tl.Entries[0].TurnNumber)
	}

	// Verify tool_call with completion data
	if tl.Entries[2].Type != "tool_call" || tl.Entries[2].ToolName != "create_file" {
		t.Errorf("entry 2: expected tool_call create_file, got %s %s", tl.Entries[2].Type, tl.Entries[2].ToolName)
	}
	if tl.Entries[2].Duration != 150.0 {
		t.Errorf("entry 2: expected duration 150, got %.1f", tl.Entries[2].Duration)
	}
	if tl.Entries[2].Success == nil || !*tl.Entries[2].Success {
		t.Error("entry 2: expected success=true")
	}
	if tl.Entries[2].FilePath != "main.go" {
		t.Errorf("entry 2: expected file_path main.go, got %q", tl.Entries[2].FilePath)
	}

	// Verify failed tool (entry index 3)
	if tl.Entries[3].ToolName != "edit_file" || tl.Entries[3].Error != "conflict" {
		t.Errorf("entry 3: expected edit_file with conflict error, got %s err=%s", tl.Entries[3].ToolName, tl.Entries[3].Error)
	}
	if tl.Entries[3].MCPServer != "editor" {
		t.Errorf("entry 3: expected mcp_server editor, got %q", tl.Entries[3].MCPServer)
	}

	// Verify intent (entry index 4)
	if tl.Entries[4].Type != "intent" || tl.Entries[4].Intent != "fixing imports" {
		t.Errorf("entry 4: expected intent 'fixing imports', got type=%s intent=%q", tl.Entries[4].Type, tl.Entries[4].Intent)
	}

	// Verify summary
	if tl.Summary.TotalActions != 7 {
		t.Errorf("summary: expected 7 total actions, got %d", tl.Summary.TotalActions)
	}
	if tl.Summary.TotalToolCalls != 2 {
		t.Errorf("summary: expected 2 tool calls, got %d", tl.Summary.TotalToolCalls)
	}
	if tl.Summary.TotalTurns != 1 {
		t.Errorf("summary: expected 1 turn, got %d", tl.Summary.TotalTurns)
	}
	if tl.Summary.ToolCallDuration != 230.0 {
		t.Errorf("summary: expected tool duration 230, got %.1f", tl.Summary.ToolCallDuration)
	}
	if tl.Summary.ToolSuccesses != 1 {
		t.Errorf("summary: expected 1 success, got %d", tl.Summary.ToolSuccesses)
	}
	if tl.Summary.ToolFailures != 1 {
		t.Errorf("summary: expected 1 failure, got %d", tl.Summary.ToolFailures)
	}
}

func TestBuildActionTimelineEmpty(t *testing.T) {
	tl := BuildActionTimeline(nil)
	if tl != nil {
		t.Error("expected nil timeline for empty events")
	}
	tl = BuildActionTimeline([]SessionEventRecord{})
	if tl != nil {
		t.Error("expected nil timeline for zero-length events")
	}
}

func TestActionTimelineJSONRoundTrip(t *testing.T) {
	dir := t.TempDir()

	boolTrue := true
	events := []SessionEventRecord{
		{Type: "assistant.turn_start"},
		{Type: "tool.execution_start", ToolName: "create_file"},
		{Type: "tool.execution_complete", ToolName: "create_file", Duration: 100.0, ToolSuccess: &boolTrue},
		{Type: "assistant.message"},
		{Type: "assistant.turn_end"},
	}

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "timeline-roundtrip",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       10.0,
		PromptMeta:     map[string]any{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.py"},
		SessionEvents:  events,
		EventCount:     5,
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "timeline-roundtrip",
		Properties: map[string]string{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
	}

	// WriteReport should auto-build the timeline
	reportPath, err := WriteReport(r, dir, "run-tl", p)
	if err != nil {
		t.Fatalf("WriteReport failed: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if parsed.ActionTimeline == nil {
		t.Fatal("expected action_timeline in JSON output")
	}
	if len(parsed.ActionTimeline.Entries) != 4 {
		t.Errorf("expected 4 timeline entries, got %d", len(parsed.ActionTimeline.Entries))
	}
	if parsed.ActionTimeline.Summary.TotalToolCalls != 1 {
		t.Errorf("expected 1 tool call in summary, got %d", parsed.ActionTimeline.Summary.TotalToolCalls)
	}
	if parsed.ActionTimeline.Summary.ToolCallDuration != 100.0 {
		t.Errorf("expected tool duration 100, got %.1f", parsed.ActionTimeline.Summary.ToolCallDuration)
	}

	// Verify the timeline entry details survive round-trip
	var foundToolCall bool
	for _, e := range parsed.ActionTimeline.Entries {
		if e.Type == "tool_call" && e.ToolName == "create_file" {
			foundToolCall = true
			if e.Duration != 100.0 {
				t.Errorf("tool_call duration: expected 100, got %.1f", e.Duration)
			}
			if e.Success == nil || !*e.Success {
				t.Error("tool_call: expected success=true")
			}
		}
	}
	if !foundToolCall {
		t.Error("expected to find tool_call entry for create_file")
	}
}

func TestActionTimelinePresetNotOverwritten(t *testing.T) {
	dir := t.TempDir()

	// Pre-set a timeline — WriteReport should not overwrite it.
	preset := &ActionTimelineReport{
		Entries: []ActionTimelineEntry{{Index: 1, Type: "custom"}},
		Summary: ActionTimelineSummary{TotalActions: 99},
	}

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "timeline-preset",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       5.0,
		PromptMeta:     map[string]any{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.go"},
		SessionEvents:  []SessionEventRecord{{Type: "assistant.message"}},
		ActionTimeline: preset,
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "timeline-preset",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
	}

	reportPath, err := WriteReport(r, dir, "run-preset", p)
	if err != nil {
		t.Fatalf("WriteReport failed: %v", err)
	}

	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if parsed.ActionTimeline.Summary.TotalActions != 99 {
		t.Errorf("expected preset total_actions=99, got %d", parsed.ActionTimeline.Summary.TotalActions)
	}
}

func TestMigrateToV2(t *testing.T) {
	r := &EvalReport{
		PromptID:   "migrate-test",
		ConfigName: "baseline",
		Review: &review.ReviewResult{
			Model:        "test-model",
			OverallScore: 3,
			MaxScore:     5,
			Summary:      "Test summary",
		},
		ReviewPanel: []review.ReviewResult{
			{Model: "panel-a", OverallScore: 4, MaxScore: 5, Summary: "Panel A"},
		},
	}

	MigrateToV2(r)

	if r.SchemaVersion != CurrentSchemaVersion {
		t.Errorf("expected schema version %d, got %d", CurrentSchemaVersion, r.SchemaVersion)
	}
	if len(r.GraderResults) != 2 {
		t.Fatalf("expected 2 grader results (1 panel + 1 consensus), got %d", len(r.GraderResults))
	}

	// Idempotent: calling again should not change anything
	r.GraderResults[0].GraderName = "modified"
	MigrateToV2(r)
	if r.GraderResults[0].GraderName != "modified" {
		t.Error("MigrateToV2 should be idempotent")
	}
}
