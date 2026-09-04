package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func boolPtr(b bool) *bool { return &b }

func TestWriteReport(t *testing.T) {
	dir := t.TempDir()

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
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
	r := &EvalReport{SchemaVersion: CurrentSchemaVersion, PromptID: "test", ConfigName: "cfg"}
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
			GraderName: "claude-opus-4.6",
			GraderType: "review",
			Score:      0.8,
			Weight:     1.0,
			Pass:       true,
			Message:    "Good code",
			Checks: []GraderPoint{
				{Label: "design", Pass: true, Weight: 1.0},
			},
		},
		{
			GraderName: "consensus",
			GraderType: "review",
			Score:      0.8,
			Weight:     1.0,
			Pass:       true,
			Message:    "Consensus result",
			Checks: []GraderPoint{
				{Label: "overall", Pass: true, Weight: 1.0},
			},
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
	if parsed.GraderResults[1].GraderName != "consensus" {
		t.Errorf("expected second grader name consensus, got %q", parsed.GraderResults[1].GraderName)
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
	if results[2].GraderName != "consolidator" {
		t.Errorf("expected last grader to be consolidator (consensus), got %q", results[2].GraderName)
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
	if results[0].GraderName != "claude-sonnet" {
		t.Errorf("expected grader name claude-sonnet, got %q", results[0].GraderName)
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
		Summary: ActionSummaryReport{TotalActions: 99},
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

	if r.SchemaVersion != 2 {
		t.Errorf("MigrateToV2 should land at schema version 2, got %d", r.SchemaVersion)
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

// TestMigrateToV3Panic verifies that MigrateToV3 panics on v < 4 reports
// (hard cutover enforced as of v4). Old reports must be regenerated.
func TestMigrateToV3Panic(t *testing.T) {
	r := &EvalReport{
		PromptID:   "migrate-v3",
		ConfigName: "baseline",
		Review: &review.ReviewResult{
			Model: "test-model", OverallScore: 3, MaxScore: 5, Summary: "S",
		},
		SchemaVersion: 0, // Explicitly set to v0
	}
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("MigrateToV3 should panic on v0 reports, but it did not")
		} else {
			msg := fmt.Sprint(r)
			if !strings.Contains(msg, "no longer supported") {
				t.Errorf("expected panic message to contain 'no longer supported', got: %s", msg)
			}
		}
	}()
	
	MigrateToV3(r)
}

// --- Session Setup in Action Timeline tests (#219) ---

func TestSessionSetupIsFirstEntry(t *testing.T) {
	events := []SessionEventRecord{
		{Type: "session.skills_loaded", Content: "azure-sdk-helper"},
		{Type: "session.mcp_servers_loaded", Content: "azure-mcp"},
		{Type: "assistant.turn_start"},
		{Type: "assistant.message"},
		{Type: "assistant.turn_end"},
	}
	setup := &SessionSetupEvent{
		Tools:        []string{"create_file", "edit_file"},
		SystemPrompt: "custom (500 chars)",
	}

	tl := BuildActionTimelineWithSetup(events, setup)
	if tl == nil {
		t.Fatal("expected non-nil timeline")
	}
	if len(tl.Entries) == 0 {
		t.Fatal("expected at least one entry")
	}
	if tl.Entries[0].Type != "session_setup" {
		t.Errorf("first entry should be session_setup, got %q", tl.Entries[0].Type)
	}
	if tl.Entries[0].Index != 1 {
		t.Errorf("first entry index should be 1, got %d", tl.Entries[0].Index)
	}
	ss := tl.Entries[0].SessionSetup
	if ss == nil {
		t.Fatal("session_setup entry should have SessionSetup data")
	}
	if ss.SystemPrompt != "custom (500 chars)" {
		t.Errorf("expected system prompt status 'custom (500 chars)', got %q", ss.SystemPrompt)
	}
	if len(ss.Tools) != 2 {
		t.Errorf("expected 2 tools, got %d", len(ss.Tools))
	}
}

func TestSessionSetupMCPLoadedFromEvents(t *testing.T) {
	events := []SessionEventRecord{
		{Type: "session.mcp_servers_loaded", Content: "azure-mcp, github-mcp"},
		{Type: "assistant.turn_start"},
		{Type: "assistant.turn_end"},
	}
	setup := &SessionSetupEvent{
		MCPServers: []ToolLoadResult{
			{Name: "azure-mcp", Status: "configured", Details: "npx @azure/mcp@latest"},
		},
		SystemPrompt: "none (default)",
	}

	tl := BuildActionTimelineWithSetup(events, setup)
	if tl == nil {
		t.Fatal("expected non-nil timeline")
	}
	ss := tl.Entries[0].SessionSetup
	if ss == nil {
		t.Fatal("expected session_setup data")
	}
	if len(ss.MCPServers) != 2 {
		t.Fatalf("expected 2 MCP servers, got %d", len(ss.MCPServers))
	}
	var azureFound, githubFound bool
	for _, s := range ss.MCPServers {
		switch s.Name {
		case "azure-mcp":
			azureFound = true
			if s.Status != "configured" {
				t.Errorf("azure-mcp should retain 'configured' status, got %q", s.Status)
			}
		case "github-mcp":
			githubFound = true
			if s.Status != "loaded" {
				t.Errorf("github-mcp should have 'loaded' status, got %q", s.Status)
			}
		}
	}
	if !azureFound { t.Error("expected azure-mcp in MCP servers") }
	if !githubFound { t.Error("expected github-mcp in MCP servers") }
}

func TestSessionSetupSkillsLoadedFromEvents(t *testing.T) {
	events := []SessionEventRecord{
		{Type: "session.skills_loaded", Content: "azure-sdk-helper, code-review"},
	}
	setup := &SessionSetupEvent{
		Skills: []ToolLoadResult{
			{Name: "azure-sdk-helper", Status: "configured", Details: "local"},
		},
		SystemPrompt: "none (default)",
	}

	tl := BuildActionTimelineWithSetup(events, setup)
	if tl == nil { t.Fatal("expected non-nil timeline") }
	ss := tl.Entries[0].SessionSetup
	if len(ss.Skills) != 2 { t.Fatalf("expected 2 skills, got %d", len(ss.Skills)) }
	if ss.Skills[0].Name != "azure-sdk-helper" || ss.Skills[0].Status != "configured" {
		t.Errorf("first skill should be azure-sdk-helper/configured, got %s/%s", ss.Skills[0].Name, ss.Skills[0].Status)
	}
	if ss.Skills[1].Name != "code-review" || ss.Skills[1].Status != "loaded" {
		t.Errorf("second skill should be code-review/loaded, got %s/%s", ss.Skills[1].Name, ss.Skills[1].Status)
	}
}

func TestSessionSetupFailedLoadsIncludeErrors(t *testing.T) {
	setup := &SessionSetupEvent{
		MCPServers: []ToolLoadResult{
			{Name: "broken-server", Status: "failed", Error: "connection refused"},
		},
		Skills: []ToolLoadResult{
			{Name: "missing-skill", Status: "failed", Error: "skill not found"},
		},
		SystemPrompt: "none (default)",
	}
	tl := BuildActionTimelineWithSetup(nil, setup)
	if tl == nil { t.Fatal("expected non-nil timeline") }
	ss := tl.Entries[0].SessionSetup
	if ss.MCPServers[0].Error != "connection refused" {
		t.Errorf("expected error 'connection refused', got %q", ss.MCPServers[0].Error)
	}
	if ss.Skills[0].Error != "skill not found" {
		t.Errorf("expected error 'skill not found', got %q", ss.Skills[0].Error)
	}
}

func TestSessionSetupOmittedWhenEmpty(t *testing.T) {
	events := []SessionEventRecord{
		{Type: "assistant.turn_start"},
		{Type: "assistant.turn_end"},
	}
	tl := BuildActionTimelineWithSetup(events, nil)
	if tl == nil { t.Fatal("expected non-nil timeline") }
	for _, e := range tl.Entries {
		if e.Type == "session_setup" {
			t.Error("should not have session_setup entry when there is no setup data")
		}
	}
}

func TestBuildScoreBreakdownWeightedAverage(t *testing.T) {
	results := []GraderResult{
		{GraderName: "file_check", GraderType: "file", Score: 1.0, Weight: 1.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		{GraderName: "code_review", GraderType: "prompt", Score: 0.8, Weight: 2.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		{GraderName: "build", GraderType: "program", Score: 0.6, Weight: 1.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
	}

	sb := BuildScoreBreakdown(results)
	if sb == nil {
		t.Fatal("expected non-nil ScoreBreakdown")
	}

	if sb.GateFailed {
		t.Error("expected gate_failed to be false")
	}
	if sb.Formula != "Final Score = Σ(grader_score × weight) / Σ(weights)" {
		t.Errorf("unexpected formula: %s", sb.Formula)
	}
	if len(sb.Contributions) != 3 {
		t.Fatalf("expected 3 contributions, got %d", len(sb.Contributions))
	}

	// Expected: (1.0*1 + 0.8*2 + 0.6*1) / (1+2+1) = 3.2/4 = 0.8
	const epsilon = 0.001
	if diff := sb.FinalScore - 0.8; diff > epsilon || diff < -epsilon {
		t.Errorf("expected final score ~0.8, got %.4f", sb.FinalScore)
	}
	if diff := sb.FinalScorePct - 80.0; diff > epsilon || diff < -epsilon {
		t.Errorf("expected final score pct ~80.0, got %.1f", sb.FinalScorePct)
	}
	if diff := sb.TotalWeight - 4.0; diff > epsilon || diff < -epsilon {
		t.Errorf("expected total weight 4.0, got %.2f", sb.TotalWeight)
	}
	if diff := sb.WeightedSum - 3.2; diff > epsilon || diff < -epsilon {
		t.Errorf("expected weighted sum 3.2, got %.4f", sb.WeightedSum)
	}

	// Verify contribution percentages sum to 100%
	var totalPct float64
	for _, c := range sb.Contributions {
		totalPct += c.ContributionPct
	}
	if diff := totalPct - 100.0; diff > epsilon || diff < -epsilon {
		t.Errorf("contribution percentages should sum to 100, got %.1f", totalPct)
	}
}

func TestBuildScoreBreakdownGateFailure(t *testing.T) {
	results := []GraderResult{
		{GraderName: "file_exists", GraderType: "file", Score: 0.0, Weight: 1.0, Gate: true, Pass: false, Checks: []GraderPoint{{Label: "check", Pass: false, Weight: 1.0}}},
		{GraderName: "code_review", GraderType: "prompt", Score: 0.9, Weight: 2.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
	}

	sb := BuildScoreBreakdown(results)
	if sb == nil {
		t.Fatal("expected non-nil ScoreBreakdown")
	}

	if !sb.GateFailed {
		t.Error("expected gate_failed to be true")
	}
	if sb.FinalScore != 0 {
		t.Errorf("expected final score 0 when gate fails, got %.4f", sb.FinalScore)
	}
	if sb.FinalScorePct != 0 {
		t.Errorf("expected final score pct 0, got %.1f", sb.FinalScorePct)
	}
	if len(sb.GateFailedNames) != 1 || sb.GateFailedNames[0] != "file_exists" {
		t.Errorf("expected gate failed name 'file_exists', got %v", sb.GateFailedNames)
	}
	if sb.Formula == "Final Score = Σ(grader_score × weight) / Σ(weights)" {
		t.Error("formula should reflect gate failure override")
	}

	// All contribution percentages should be 0 when gate fails
	for _, c := range sb.Contributions {
		if c.ContributionPct != 0 {
			t.Errorf("expected 0 contribution pct on gate failure, got %.1f for %s", c.ContributionPct, c.Name)
		}
	}
}

func TestSessionSetupNilEventsWithSetup(t *testing.T) {
	setup := &SessionSetupEvent{
		SystemPrompt: "custom (100 chars)",
		Tools:        []string{"bash"},
	}
	tl := BuildActionTimelineWithSetup(nil, setup)
	if tl == nil { t.Fatal("expected non-nil timeline for setup-only data") }
	if len(tl.Entries) != 1 { t.Fatalf("expected 1 entry, got %d", len(tl.Entries)) }
	if tl.Entries[0].Type != "session_setup" {
		t.Errorf("expected session_setup, got %q", tl.Entries[0].Type)
	}
}

func TestSessionSetupJSONRoundTrip(t *testing.T) {
	dir := t.TempDir()
	setup := &SessionSetupEvent{
		MCPServers:   []ToolLoadResult{{Name: "azure-mcp", Status: "loaded"}},
		Skills:       []ToolLoadResult{{Name: "helper", Status: "loaded"}},
		Tools:        []string{"create_file", "bash"},
		SystemPrompt: "custom (200 chars)",
		StarterFiles: []string{"main.py", "requirements.txt"},
	}
	events := []SessionEventRecord{
		{Type: "assistant.turn_start"},
		{Type: "assistant.message"},
		{Type: "assistant.turn_end"},
	}
	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "setup-roundtrip",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       5.0,
		PromptMeta:     map[string]any{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.py"},
		SessionEvents:  events,
		SessionSetup:   setup,
		EventCount:     3,
		Success:        true,
	}
	p := &prompt.Prompt{
		ID:         "setup-roundtrip",
		Properties: map[string]string{"service": "identity", "plane": "data-plane", "language": "python", "category": "auth"},
	}
	reportPath, err := WriteReport(r, dir, "run-setup", p)
	if err != nil { t.Fatalf("WriteReport failed: %v", err) }
	data, err := os.ReadFile(reportPath)
	if err != nil { t.Fatalf("failed to read: %v", err) }
	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil { t.Fatalf("invalid JSON: %v", err) }
	if parsed.ActionTimeline == nil { t.Fatal("expected action_timeline in JSON") }
	if len(parsed.ActionTimeline.Entries) == 0 { t.Fatal("expected at least one timeline entry") }
	if parsed.ActionTimeline.Entries[0].Type != "session_setup" {
		t.Errorf("first entry should be session_setup, got %q", parsed.ActionTimeline.Entries[0].Type)
	}
	ss := parsed.ActionTimeline.Entries[0].SessionSetup
	if ss == nil { t.Fatal("session_setup data missing after round-trip") }
	if ss.SystemPrompt != "custom (200 chars)" { t.Errorf("system_prompt_status mismatch: %q", ss.SystemPrompt) }
	if len(ss.MCPServers) != 1 || ss.MCPServers[0].Name != "azure-mcp" { t.Errorf("MCP servers mismatch: %+v", ss.MCPServers) }
	if len(ss.StarterFiles) != 2 { t.Errorf("expected 2 starter files, got %d", len(ss.StarterFiles)) }
	if parsed.SessionSetup == nil { t.Error("expected session_setup at top level of report") }
}

func TestBuildScoreBreakdownDefaultWeight(t *testing.T) {
	results := []GraderResult{
		{GraderName: "grader_a", GraderType: "file", Score: 1.0, Weight: 0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		{GraderName: "grader_b", GraderType: "file", Score: 0.5, Weight: 0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
	}

	sb := BuildScoreBreakdown(results)
	if sb == nil {
		t.Fatal("expected non-nil ScoreBreakdown")
	}

	// Weight 0 defaults to 1.0, so: (1.0*1 + 0.5*1) / (1+1) = 0.75
	const epsilon = 0.001
	if diff := sb.FinalScore - 0.75; diff > epsilon || diff < -epsilon {
		t.Errorf("expected final score ~0.75 (default weights), got %.4f", sb.FinalScore)
	}
	for _, c := range sb.Contributions {
		if c.Weight != 1.0 {
			t.Errorf("expected effective weight 1.0, got %.2f for %s", c.Weight, c.Name)
		}
	}
}

func TestBuildScoreBreakdownNilForEmpty(t *testing.T) {
	sb := BuildScoreBreakdown(nil)
	if sb != nil {
		t.Error("expected nil for empty results")
	}
	sb = BuildScoreBreakdown([]GraderResult{})
	if sb != nil {
		t.Error("expected nil for zero-length results")
	}
}

func TestBuildScoreBreakdownNilForLegacyReview(t *testing.T) {
	// Legacy review-only results have no Score/Weight/Gate data (all zero-valued),
	// so BuildScoreBreakdown omits the section entirely.
	results := []GraderResult{
		{GraderName: "reviewer", GraderType: "review", Score: 0, Weight: 0, Checks: []GraderPoint{{Label: "review", Pass: true, Weight: 1.0}}},
	}
	sb := BuildScoreBreakdown(results)
	if sb != nil {
		t.Error("expected nil ScoreBreakdown for legacy review-only results")
	}
}

func TestScoreBreakdownJSONRoundTrip(t *testing.T) {
	dir := t.TempDir()

	graderResults := []GraderResult{
		{GraderName: "file_check", GraderType: "file", Score: 1.0, Weight: 1.0, Pass: true, Gate: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		{GraderName: "review", GraderType: "prompt", Score: 0.8, Weight: 2.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
	}

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "breakdown-test",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       10.0,
		PromptMeta:     map[string]any{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.go"},
		GraderResults:  graderResults,
		ScoreBreakdown: BuildScoreBreakdown(graderResults),
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "breakdown-test",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"},
	}

	reportPath, err := WriteReport(r, dir, "run-bd", p)
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

	if parsed.ScoreBreakdown == nil {
		t.Fatal("expected score_breakdown in JSON output")
	}
	if len(parsed.ScoreBreakdown.Contributions) != 2 {
		t.Errorf("expected 2 contributions, got %d", len(parsed.ScoreBreakdown.Contributions))
	}

	const epsilon = 0.001
	// Expected: (1.0*1 + 0.8*2) / (1+2) = 2.6/3 ≈ 0.8667
	if diff := parsed.ScoreBreakdown.FinalScore - 0.8667; diff > epsilon || diff < -epsilon {
		t.Errorf("expected final score ~0.8667, got %.4f", parsed.ScoreBreakdown.FinalScore)
	}
	if parsed.ScoreBreakdown.GateFailed {
		t.Error("gate should not be failed")
	}
}


func TestScoreBreakdownInMarkdownReport(t *testing.T) {
	dir := t.TempDir()

	graderResults := []GraderResult{
		{GraderName: "file_check", GraderType: "file", Score: 1.0, Weight: 0.5, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
		{GraderName: "review", GraderType: "prompt", Score: 0.6, Weight: 1.0, Pass: true, Checks: []GraderPoint{{Label: "check", Pass: true, Weight: 1.0}}},
	}

	r := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "md-breakdown",
		ConfigName:     "baseline",
		Timestamp:      "2024-01-15T10:00:00Z",
		Duration:       10.0,
		PromptMeta:     map[string]any{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.go"},
		GraderResults:  graderResults,
		ScoreBreakdown: BuildScoreBreakdown(graderResults),
		Success:        true,
	}

	mdPath, err := WriteMarkdownReport(r, dir, "run-md-bd", "storage", "data-plane", "go", "crud")
	if err != nil {
		t.Fatalf("WriteMarkdownReport failed: %v", err)
	}

	data, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("failed to read markdown: %v", err)
	}

	md := string(data)
	if !containsStr(md, "## Score Breakdown") {
		t.Error("Markdown report should contain '## Score Breakdown' heading")
	}
	if !containsStr(md, "file_check") {
		t.Error("Markdown report should list grader 'file_check'")
	}
	if !containsStr(md, "review") {
		t.Error("Markdown report should list grader 'review'")
	}
	if !containsStr(md, "**Final**") {
		t.Error("Markdown report should contain final score row")
	}
}

func containsStr(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (len(s) >= len(substr)) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
