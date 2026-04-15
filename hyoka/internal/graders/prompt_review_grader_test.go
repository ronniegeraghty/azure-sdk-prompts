package graders

import (
"context"
"os"
"path/filepath"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func TestPromptReviewGraderKindAndName(t *testing.T) {
g := NewPromptReviewGrader("test-review", &review.StubReviewer{}, nil)
if g.Kind() != KindPromptReview {
	t.Errorf("Kind() = %q, want %q", g.Kind(), KindPromptReview)
}
if g.Name() != "test-review" {
	t.Errorf("Name() = %q, want %q", g.Name(), "test-review")
}
}

func TestPromptReviewGraderSingleReviewer(t *testing.T) {
ws := t.TempDir()
if err := os.WriteFile(filepath.Join(ws, "main.py"), []byte("print('hello')"), 0644); err != nil {
	t.Fatal(err)
}

g := NewPromptReviewGrader("single-review", &review.StubReviewer{}, nil)
input := GraderInput{
	WorkspacePath:  ws,
	OriginalPrompt: "Create a hello world script",
	EvalCriteria:   "- Must print hello",
}

result, err := g.Grade(context.Background(), input)
if err != nil {
	t.Fatalf("Grade() error: %v", err)
}
defer g.CleanupWorkspace()

if result.Kind != KindPromptReview {
	t.Errorf("result.Kind = %q, want %q", result.Kind, KindPromptReview)
}
if !result.Pass {
	t.Error("expected review to pass (stub reviewer always passes)")
}
if result.Score <= 0 {
	t.Errorf("expected positive score, got %f", result.Score)
}
if result.ReviewDetails == nil {
	t.Fatal("expected ReviewDetails to be populated")
}
if result.ReviewDetails.OverallScore != 1 {
	t.Errorf("OverallScore = %d, want 1", result.ReviewDetails.OverallScore)
}
if result.ReviewDetails.MaxScore != 1 {
	t.Errorf("MaxScore = %d, want 1", result.ReviewDetails.MaxScore)
}
if len(result.ReviewDetails.Criteria) != 1 {
	t.Fatalf("expected 1 criterion, got %d", len(result.ReviewDetails.Criteria))
}

// Verify backward-compat fields.
if g.LastConsolidated == nil {
	t.Error("expected LastConsolidated to be set")
}
if g.LastPanel != nil {
	t.Error("expected LastPanel to be nil for single reviewer")
}
}

func TestPromptReviewGraderNoReviewer(t *testing.T) {
g := NewPromptReviewGrader("empty", nil, nil)
_, err := g.Grade(context.Background(), GraderInput{WorkspacePath: t.TempDir()})
if err == nil {
	t.Error("expected error when no reviewer configured")
}
}

func TestPromptReviewGraderCleanup(t *testing.T) {
ws := t.TempDir()
os.WriteFile(filepath.Join(ws, "f.txt"), []byte("data"), 0644)

g := NewPromptReviewGrader("cleanup-test", &review.StubReviewer{}, nil)
_, err := g.Grade(context.Background(), GraderInput{
	WorkspacePath:  ws,
	OriginalPrompt: "test",
})
if err != nil {
	t.Fatal(err)
}

reviewDir := g.LastReviewWorkDir
if reviewDir == "" {
	t.Fatal("expected LastReviewWorkDir to be set")
}

// Directory should exist before cleanup.
if _, err := os.Stat(reviewDir); os.IsNotExist(err) {
	t.Fatal("review workspace should exist before cleanup")
}

g.CleanupWorkspace()

// Directory should not exist after cleanup.
if _, err := os.Stat(reviewDir); !os.IsNotExist(err) {
	t.Error("review workspace should be removed after cleanup")
}
if g.LastReviewWorkDir != "" {
	t.Error("LastReviewWorkDir should be empty after cleanup")
}
}

func TestCriteriaScore(t *testing.T) {
tests := []struct {
	name  string
	input *review.ReviewResult
	want  float64
}{
	{"all passed", &review.ReviewResult{OverallScore: 5, MaxScore: 5}, 1.0},
	{"half passed", &review.ReviewResult{OverallScore: 3, MaxScore: 6}, 0.5},
	{"none passed", &review.ReviewResult{OverallScore: 0, MaxScore: 5}, 0.0},
	{"zero max", &review.ReviewResult{OverallScore: 0, MaxScore: 0}, 0.0},
}
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		got := criteriaScore(tt.input)
		if got != tt.want {
			t.Errorf("criteriaScore() = %f, want %f", got, tt.want)
		}
	})
}
}

func TestCopyDirContents(t *testing.T) {
src := t.TempDir()
os.MkdirAll(filepath.Join(src, "subdir"), 0755)
os.WriteFile(filepath.Join(src, "top.txt"), []byte("top"), 0644)
os.WriteFile(filepath.Join(src, "subdir", "nested.txt"), []byte("nested"), 0644)

dst := t.TempDir()
if err := copyDirContents(src, dst); err != nil {
	t.Fatal(err)
}

data, err := os.ReadFile(filepath.Join(dst, "top.txt"))
if err != nil || string(data) != "top" {
	t.Error("top.txt not copied correctly")
}
data, err = os.ReadFile(filepath.Join(dst, "subdir", "nested.txt"))
if err != nil || string(data) != "nested" {
	t.Error("nested.txt not copied correctly")
}
}
