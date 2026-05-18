package graders

import (
	"context"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// stubCaller implements LLMCaller for testing.
type stubCaller struct {
	response string
	err      error
}

func (s *stubCaller) Call(_ context.Context, _ string, _ string) (string, error) {
	return s.response, s.err
}

// --- NewPromptGrader config parsing ---

func TestNewPromptGraderValidConfig(t *testing.T) {
	pg, err := NewPromptGrader("test", map[string]any{
		"model":     "claude-opus-4.6",
		"rubric":    "Evaluate correctness.",
		"max_score": 5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.Model != "claude-opus-4.6" {
		t.Errorf("model: got %q, want %q", pg.Model, "claude-opus-4.6")
	}
	if pg.Rubric != "Evaluate correctness." {
		t.Errorf("rubric: got %q", pg.Rubric)
	}
	if pg.MaxScore != 5 {
		t.Errorf("max_score: got %d, want 5", pg.MaxScore)
	}
	if pg.Name != "test" {
		t.Errorf("name: got %q, want %q", pg.Name, "test")
	}
}

func TestNewPromptGraderDefaultMaxScore(t *testing.T) {
	pg, err := NewPromptGrader("test", map[string]any{
		"model":  "claude-sonnet-4.5",
		"rubric": "Check quality.",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.MaxScore != DefaultMaxScore {
		t.Errorf("max_score: got %d, want default %d", pg.MaxScore, DefaultMaxScore)
	}
}

func TestNewPromptGraderFloat64MaxScore(t *testing.T) {
	// YAML/JSON often decodes numbers as float64
	pg, err := NewPromptGrader("test", map[string]any{
		"model":     "claude-opus-4.6",
		"rubric":    "Check.",
		"max_score": float64(20),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.MaxScore != 20 {
		t.Errorf("max_score: got %d, want 20", pg.MaxScore)
	}
}

func TestNewPromptGraderMissingModel(t *testing.T) {
	_, err := NewPromptGrader("test", map[string]any{
		"rubric": "Check quality.",
	})
	if err == nil {
		t.Fatal("expected error for missing model")
	}
	if !strings.Contains(err.Error(), "model is required") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewPromptGraderMissingRubric(t *testing.T) {
	_, err := NewPromptGrader("test", map[string]any{
		"model": "claude-opus-4.6",
	})
	if err == nil {
		t.Fatal("expected error for missing rubric")
	}
	if !strings.Contains(err.Error(), "rubric is required") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewPromptGraderZeroMaxScore(t *testing.T) {
	_, err := NewPromptGrader("test", map[string]any{
		"model":     "claude-opus-4.6",
		"rubric":    "Check.",
		"max_score": 0,
	})
	if err == nil {
		t.Fatal("expected error for zero max_score")
	}
	if !strings.Contains(err.Error(), "must be positive") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewPromptGraderNegativeMaxScore(t *testing.T) {
	_, err := NewPromptGrader("test", map[string]any{
		"model":     "claude-opus-4.6",
		"rubric":    "Check.",
		"max_score": -1,
	})
	if err == nil {
		t.Fatal("expected error for negative max_score")
	}
}

func TestNewPromptGraderInvalidMaxScoreType(t *testing.T) {
	_, err := NewPromptGrader("test", map[string]any{
		"model":     "claude-opus-4.6",
		"rubric":    "Check.",
		"max_score": "not-a-number",
	})
	if err == nil {
		t.Fatal("expected error for string max_score")
	}
	if !strings.Contains(err.Error(), "must be an integer") {
		t.Errorf("unexpected error: %v", err)
	}
}

// --- NormalizeScore ---

func TestNormalizeScore(t *testing.T) {
	tests := []struct {
		name     string
		raw, max int
		want     float64
	}{
		{"full score", 10, 10, 1.0},
		{"zero score", 0, 10, 0.0},
		{"half score", 5, 10, 0.5},
		{"custom max", 3, 5, 0.6},
		{"one of twenty", 1, 20, 0.05},
		{"zero max returns 0", 5, 0, 0.0},
		{"negative max returns 0", 5, -1, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeScore(tt.raw, tt.max)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("NormalizeScore(%d, %d) = %f, want %f", tt.raw, tt.max, got, tt.want)
			}
		})
	}
}

// --- ParseResponse ---

func TestParseResponseValid(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	resp, err := pg.ParseResponse(`{"score": 7, "reasoning": "Good code quality."}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Score != 7 {
		t.Errorf("score: got %d, want 7", resp.Score)
	}
	if resp.Reasoning != "Good code quality." {
		t.Errorf("reasoning: got %q", resp.Reasoning)
	}
}

func TestParseResponseMarkdownFenced(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	input := "```json\n{\"score\": 8, \"reasoning\": \"Well structured.\"}\n```"
	resp, err := pg.ParseResponse(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Score != 8 {
		t.Errorf("score: got %d, want 8", resp.Score)
	}
}

func TestParseResponseWithSurroundingText(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	input := `Here is my evaluation: {"score": 6, "reasoning": "Decent"} Hope this helps!`
	resp, err := pg.ParseResponse(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Score != 6 {
		t.Errorf("score: got %d, want 6", resp.Score)
	}
}

func TestParseResponseClampsHigh(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	resp, err := pg.ParseResponse(`{"score": 15, "reasoning": "over"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Score != 10 {
		t.Errorf("score should be clamped to max_score (10), got %d", resp.Score)
	}
}

func TestParseResponseClampsLow(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	resp, err := pg.ParseResponse(`{"score": -3, "reasoning": "negative"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Score != 0 {
		t.Errorf("score should be clamped to 0, got %d", resp.Score)
	}
}

func TestParseResponseNoJSON(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	_, err := pg.ParseResponse("no json here")
	if err == nil {
		t.Fatal("expected error for no JSON")
	}
	if !strings.Contains(err.Error(), "no JSON found") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestParseResponseInvalidJSON(t *testing.T) {
	pg := &PromptGrader{MaxScore: 10}
	_, err := pg.ParseResponse(`{invalid json}`)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

// --- BuildPrompt ---

func TestBuildPromptContainsRubric(t *testing.T) {
	pg := &PromptGrader{
		Model:    "claude-opus-4.6",
		Rubric:   "Check for error handling.",
		MaxScore: 10,
	}
	prompt := pg.BuildPrompt(map[string]string{
		"main.py": "print('hello')",
	})
	if !strings.Contains(prompt, "Check for error handling.") {
		t.Error("prompt should contain the rubric")
	}
	if !strings.Contains(prompt, "main.py") {
		t.Error("prompt should contain file name")
	}
	if !strings.Contains(prompt, "print('hello')") {
		t.Error("prompt should contain file content")
	}
	if !strings.Contains(prompt, "0 to 10") {
		t.Error("prompt should reference max score")
	}
}

func TestBuildPromptEmptyFiles(t *testing.T) {
	pg := &PromptGrader{MaxScore: 5, Rubric: "test"}
	prompt := pg.BuildPrompt(map[string]string{})
	if !strings.Contains(prompt, "(no files found)") {
		t.Error("prompt should indicate no files")
	}
}

// --- Grade (integration with stub) ---

func TestGradeSuccess(t *testing.T) {
	pg := &PromptGrader{
		Name:     "test-grader",
		Model:    "claude-opus-4.6",
		Rubric:   "Evaluate quality.",
		MaxScore: 10,
		Caller: &stubCaller{
			response: `{"score": 8, "reasoning": "Well-written code with good error handling."}`,
		},
	}

	// Create a temp workspace with a file
	workDir := t.TempDir()
	writeTestFile(t, workDir, "main.py", "def main():\n    print('hello')\n")

	result, err := pg.Grade(context.Background(), workDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if math.Abs(result.Score-0.8) > 1e-9 {
		t.Errorf("normalized score: got %f, want 0.8", result.Score)
	}
	if !result.Passed {
		t.Error("expected passed=true for score > 0")
	}
	if result.Details.Model != "claude-opus-4.6" {
		t.Errorf("model: got %q", result.Details.Model)
	}
	if result.Details.RawScore != 8 {
		t.Errorf("raw_score: got %d", result.Details.RawScore)
	}
	if result.Details.MaxScore != 10 {
		t.Errorf("max_score: got %d", result.Details.MaxScore)
	}
	if result.Details.Reasoning != "Well-written code with good error handling." {
		t.Errorf("reasoning: got %q", result.Details.Reasoning)
	}
}

func TestGradeZeroScore(t *testing.T) {
	pg := &PromptGrader{
		Name:     "test-grader",
		Model:    "claude-opus-4.6",
		Rubric:   "Evaluate quality.",
		MaxScore: 10,
		Caller: &stubCaller{
			response: `{"score": 0, "reasoning": "No relevant code generated."}`,
		},
	}
	workDir := t.TempDir()
	writeTestFile(t, workDir, "empty.py", "")

	result, err := pg.Grade(context.Background(), workDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Score != 0.0 {
		t.Errorf("score: got %f, want 0.0", result.Score)
	}
	if result.Passed {
		t.Error("expected passed=false for score 0")
	}
}

func TestGradeNoCaller(t *testing.T) {
	pg := &PromptGrader{
		Name:     "test-grader",
		Model:    "claude-opus-4.6",
		Rubric:   "Evaluate quality.",
		MaxScore: 10,
		Caller:   nil,
	}
	workDir := t.TempDir()
	writeTestFile(t, workDir, "main.py", "pass")

	_, err := pg.Grade(context.Background(), workDir)
	if err == nil {
		t.Fatal("expected error for nil caller")
	}
	if !strings.Contains(err.Error(), "no LLM caller configured") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGradeLLMCallError(t *testing.T) {
	pg := &PromptGrader{
		Name:     "test-grader",
		Model:    "claude-opus-4.6",
		Rubric:   "Evaluate quality.",
		MaxScore: 10,
		Caller: &stubCaller{
			err: fmt.Errorf("connection refused"),
		},
	}
	workDir := t.TempDir()
	writeTestFile(t, workDir, "main.py", "pass")

	_, err := pg.Grade(context.Background(), workDir)
	if err == nil {
		t.Fatal("expected error from LLM call failure")
	}
	if !strings.Contains(err.Error(), "LLM call failed") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGradeBadResponse(t *testing.T) {
	pg := &PromptGrader{
		Name:     "test-grader",
		Model:    "claude-opus-4.6",
		Rubric:   "Evaluate quality.",
		MaxScore: 10,
		Caller: &stubCaller{
			response: "I can't return JSON right now.",
		},
	}
	workDir := t.TempDir()
	writeTestFile(t, workDir, "main.py", "pass")

	_, err := pg.Grade(context.Background(), workDir)
	if err == nil {
		t.Fatal("expected error for unparseable response")
	}
}

// writeTestFile is a helper that creates a file in the given directory.
func writeTestFile(t *testing.T, dir, name, content string) {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0644); err != nil {
		t.Fatalf("writing test file %s: %v", p, err)
	}
}
