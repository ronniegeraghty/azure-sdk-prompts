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
if result.Extras.Review == nil {
	t.Fatal("expected ReviewDetails to be populated")
}

}

func TestPromptReviewGraderNoReviewer(t *testing.T) {
g := NewPromptReviewGrader("empty", nil, nil)
_, err := g.Grade(context.Background(), GraderInput{WorkspacePath: t.TempDir()})
if err == nil {
	t.Error("expected error when no reviewer configured")
}
}

func TestPromptReviewGraderRecordsWorkspacePath(t *testing.T) {
// Engine owns workspace lifecycle; the grader simply records the path
// it was handed so the engine can read reviewed files from it.
ws := t.TempDir()
os.WriteFile(filepath.Join(ws, "f.txt"), []byte("data"), 0644)

g := NewPromptReviewGrader("record-test", &review.StubReviewer{}, nil)
_, err := g.Grade(context.Background(), GraderInput{
	WorkspacePath:  ws,
	OriginalPrompt: "test",
})
if err != nil {
	t.Fatal(err)
}

if g.LastReviewWorkDir != ws {
	t.Fatalf("LastReviewWorkDir = %q, want %q (grader must use engine-provided path)", g.LastReviewWorkDir, ws)
}

// CleanupWorkspace is now a no-op; the workspace must remain intact
// for the engine to read reviewed files.
g.CleanupWorkspace()

if _, err := os.Stat(ws); os.IsNotExist(err) {
	t.Error("CleanupWorkspace must not remove engine-owned workspace")
}
if g.LastReviewWorkDir != "" {
	t.Error("CleanupWorkspace should clear LastReviewWorkDir")
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
