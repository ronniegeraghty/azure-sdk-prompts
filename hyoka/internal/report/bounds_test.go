package report

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

// makeString returns a string of n bytes filled with 'x'.
func makeString(n int) string {
	return strings.Repeat("x", n)
}

func TestTruncateReport_NoOpBelowLimit(t *testing.T) {
	r := &EvalReport{
		PromptID:   "small",
		ConfigName: "cfg",
		SessionEvents: []SessionEventRecord{
			{Type: "tool.execution_complete", Content: "short content"},
		},
	}

	got := truncateReportWithLimit(r, 10*1024*1024) // 10 MB limit
	if got {
		t.Error("expected no truncation for a small report")
	}
	if r.SessionEvents[0].Content != "short content" {
		t.Error("content should be unchanged")
	}
}

func TestTruncateReport_TruncatesVerboseFields(t *testing.T) {
	big := makeString(100 * 1024) // 100 KB per field

	r := &EvalReport{
		PromptID:   "big",
		ConfigName: "cfg",
		SessionEvents: []SessionEventRecord{
			{Type: "tool.execution_complete", Content: big, ToolResult: big, ToolArgs: big},
			{Type: "assistant.message", Content: big},
		},
		ReviewedFiles: []ReviewedFile{
			{Path: "main.go", Content: big},
		},
		ErrorDetails: big,
		Review: &review.ReviewResult{
			Events: []review.ReviewEvent{
				{Content: big, Result: big, ToolArgs: big},
			},
		},
		ReviewPanel: []review.ReviewResult{
			{Events: []review.ReviewEvent{{Content: big}}},
		},
	}

	// Use a limit small enough that the report exceeds it.
	got := truncateReportWithLimit(r, 1024) // 1 KB limit
	if !got {
		t.Fatal("expected truncation to be applied")
	}

	// All large fields should be capped at maxFieldBytes + marker length.
	maxLen := maxFieldBytes + len(" "+truncatedMessage)
	check := func(name, val string) {
		t.Helper()
		if len(val) > maxLen {
			t.Errorf("%s: len %d exceeds max %d", name, len(val), maxLen)
		}
		if len(val) > maxFieldBytes && !strings.HasSuffix(val, truncatedMessage) {
			t.Errorf("%s: missing truncation marker", name)
		}
	}

	check("SessionEvents[0].Content", r.SessionEvents[0].Content)
	check("SessionEvents[0].ToolResult", r.SessionEvents[0].ToolResult)
	check("SessionEvents[0].ToolArgs", r.SessionEvents[0].ToolArgs)
	check("SessionEvents[1].Content", r.SessionEvents[1].Content)
	check("ReviewedFiles[0].Content", r.ReviewedFiles[0].Content)
	check("ErrorDetails", r.ErrorDetails)
	check("Review.Events[0].Content", r.Review.Events[0].Content)
	check("Review.Events[0].Result", r.Review.Events[0].Result)
	check("Review.Events[0].ToolArgs", r.Review.Events[0].ToolArgs)
	check("ReviewPanel[0].Events[0].Content", r.ReviewPanel[0].Events[0].Content)
}

func TestTruncateReport_LogsWarning(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelWarn}))
	slog.SetDefault(logger)
	defer slog.SetDefault(slog.Default()) // restore after test

	big := makeString(100 * 1024)
	r := &EvalReport{
		PromptID:   "warn-test",
		ConfigName: "cfg",
		SessionEvents: []SessionEventRecord{
			{Content: big},
		},
	}

	truncateReportWithLimit(r, 1024)

	logOutput := buf.String()
	if !strings.Contains(logOutput, "memory bound") {
		t.Errorf("expected warning about memory bounds, got: %s", logOutput)
	}
	if !strings.Contains(logOutput, "warn-test") {
		t.Errorf("expected prompt_id in warning, got: %s", logOutput)
	}
}

func TestWriteReport_LargeReportWrittenCorrectly(t *testing.T) {
	dir := t.TempDir()

	// Create a report with enough data that it's non-trivial.
	events := make([]SessionEventRecord, 500)
	for i := range events {
		events[i] = SessionEventRecord{
			Type:       "tool.execution_complete",
			ToolName:   "create_file",
			Content:    makeString(1024), // 1 KB each
			ToolResult: makeString(512),
		}
	}

	r := &EvalReport{
		PromptID:       "large-report",
		ConfigName:     "baseline",
		Timestamp:      "2024-06-01T00:00:00Z",
		Duration:       30.0,
		PromptMeta:     map[string]any{"service": "storage", "language": "python"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.py"},
		SessionEvents:  events,
		EventCount:     500,
		ToolCalls:      []string{"create_file"},
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "large-report",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"},
	}

	reportPath, err := WriteReport(r, dir, "run-large", p)
	if err != nil {
		t.Fatalf("WriteReport failed: %v", err)
	}

	// Read back and verify valid JSON.
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read report: %v", err)
	}

	var parsed EvalReport
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if parsed.PromptID != "large-report" {
		t.Errorf("expected prompt_id 'large-report', got %q", parsed.PromptID)
	}
	if len(parsed.SessionEvents) != 500 {
		t.Errorf("expected 500 events, got %d", len(parsed.SessionEvents))
	}
}

func TestWriteReport_TruncationAppliedAboveThreshold(t *testing.T) {
	dir := t.TempDir()

	// Build a report that exceeds MaxReportBytes when serialized.
	// Use 600 events with ~100 KB content each ≈ 60 MB raw content.
	events := make([]SessionEventRecord, 600)
	for i := range events {
		events[i] = SessionEventRecord{
			Type:       "tool.execution_complete",
			Content:    makeString(100 * 1024),
			ToolResult: makeString(100 * 1024),
		}
	}

	r := &EvalReport{
		PromptID:       "huge-report",
		ConfigName:     "baseline",
		Timestamp:      "2024-06-01T00:00:00Z",
		Duration:       60.0,
		PromptMeta:     map[string]any{"service": "storage"},
		ConfigUsed:     map[string]any{"name": "baseline"},
		GeneratedFiles: []string{"main.py"},
		SessionEvents:  events,
		EventCount:     600,
		Success:        true,
	}

	p := &prompt.Prompt{
		ID:         "huge-report",
		Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"},
	}

	reportPath, err := WriteReport(r, dir, "run-huge", p)
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

	// After truncation, individual Content fields should be capped.
	for i, ev := range parsed.SessionEvents {
		if len(ev.Content) > maxFieldBytes+len(" "+truncatedMessage)+10 {
			t.Errorf("event %d Content not truncated: len=%d", i, len(ev.Content))
		}
	}
}

func TestTruncateField(t *testing.T) {
	small := "hello"
	if truncateField(&small) {
		t.Error("should not truncate small strings")
	}
	if small != "hello" {
		t.Error("should not modify small strings")
	}

	big := makeString(maxFieldBytes + 1000)
	if !truncateField(&big) {
		t.Error("should truncate large strings")
	}
	if !strings.HasSuffix(big, truncatedMessage) {
		t.Error("truncated string should end with marker")
	}
	if len(big) > maxFieldBytes+len(" "+truncatedMessage)+10 {
		t.Errorf("truncated string too long: %d", len(big))
	}
}
