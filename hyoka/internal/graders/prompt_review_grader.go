package graders

import (
"context"
"fmt"
"log/slog"
"os"

"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

// PromptReviewGrader wraps the AI review panel as a grader type (WI-023).
// It implements the Grader interface so the review runs as part of the
// unified grading pipeline rather than a separate phase.
type PromptReviewGrader struct {
name          string
reviewer      review.Reviewer
panelReviewer *review.PanelReviewer

// LastPanel and LastConsolidated hold the raw review results after
// Grade() completes. The engine reads these for backward-compat
// report fields (ReviewPanel, Review).
LastPanel        []review.ReviewResult
LastConsolidated *review.ReviewResult
// LastReviewWorkDir is the reviewer workspace. The engine uses it
// to capture reviewed (annotated) files after grading.
LastReviewWorkDir string
}

// NewPromptReviewGrader creates a review grader backed by the given reviewer.
// Pass either a single reviewer or a panelReviewer (not both).
func NewPromptReviewGrader(name string, singleReviewer review.Reviewer, panelReviewer *review.PanelReviewer) *PromptReviewGrader {
return &PromptReviewGrader{
	name:          name,
	reviewer:      singleReviewer,
	panelReviewer: panelReviewer,
}
}

func (g *PromptReviewGrader) Kind() string { return KindPromptReview }
func (g *PromptReviewGrader) Name() string { return g.name }

// Grade creates an isolated workspace copy, runs the review, and converts the
// result into a GraderResult. The consolidated review score maps to the grader
// score; panel member results are stored in ReviewDetails.PanelResults.
func (g *PromptReviewGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
result := GraderResult{
	Kind: KindPromptReview,
	Name: g.name,
}

// Create isolated reviewer workspace with a copy of the generated files.
reviewWorkDir, err := copyDirToTemp(input.WorkspacePath)
if err != nil {
	slog.Warn("Reviewer workspace creation failed, using original", "error", err)
	reviewWorkDir = input.WorkspacePath
} else {
	// Keep the directory alive — the engine reads reviewed files from it.
	g.LastReviewWorkDir = reviewWorkDir
}

if g.panelReviewer != nil {
	return g.gradePanel(ctx, input, reviewWorkDir, result)
}
if g.reviewer != nil {
	return g.gradeSingle(ctx, input, reviewWorkDir, result)
}
return result, fmt.Errorf("no reviewer configured")
}

// CleanupWorkspace removes the temporary reviewer workspace.
func (g *PromptReviewGrader) CleanupWorkspace() {
if g.LastReviewWorkDir != "" {
	os.RemoveAll(g.LastReviewWorkDir)
	g.LastReviewWorkDir = ""
}
}

func (g *PromptReviewGrader) gradePanel(ctx context.Context, input GraderInput, workDir string, result GraderResult) (GraderResult, error) {
models := g.panelReviewer.Models()
slog.Debug("Review grader starting panel", "models", models)

panel, consolidated, err := g.panelReviewer.ReviewPanel(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, input.EvalCriteria)
if err != nil {
	return result, fmt.Errorf("review panel failed: %w", err)
}

g.LastPanel = panel
g.LastConsolidated = consolidated

result.Pass = consolidated.Scores.AllPassed()
result.Score = criteriaScore(consolidated)
result.Message = consolidated.Summary

details := &ReviewGraderDetails{
	Model:        consolidated.Model,
	OverallScore: consolidated.OverallScore,
	MaxScore:     consolidated.MaxScore,
	Summary:      consolidated.Summary,
	Issues:       consolidated.Issues,
	Strengths:    consolidated.Strengths,
	IsConsensus:  true,
}
for _, c := range consolidated.Scores.Criteria {
	details.Criteria = append(details.Criteria, ReviewCriterion{
		Name: c.Name, Passed: c.Passed, Reason: c.Reason,
	})
}
for _, p := range panel {
	entry := ReviewPanelEntry{
		Model:        p.Model,
		OverallScore: p.OverallScore,
		MaxScore:     p.MaxScore,
		Summary:      p.Summary,
		Issues:       p.Issues,
		Strengths:    p.Strengths,
	}
	for _, c := range p.Scores.Criteria {
		entry.Criteria = append(entry.Criteria, ReviewCriterion{
			Name: c.Name, Passed: c.Passed, Reason: c.Reason,
		})
	}
	details.PanelResults = append(details.PanelResults, entry)
}
result.ReviewDetails = details
return result, nil
}

func (g *PromptReviewGrader) gradeSingle(ctx context.Context, input GraderInput, workDir string, result GraderResult) (GraderResult, error) {
slog.Debug("Review grader starting single review")

reviewResult, err := g.reviewer.Review(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, input.EvalCriteria)
if err != nil {
	return result, fmt.Errorf("review failed: %w", err)
}

g.LastConsolidated = reviewResult

result.Pass = reviewResult.Scores.AllPassed()
result.Score = criteriaScore(reviewResult)
result.Message = reviewResult.Summary

details := &ReviewGraderDetails{
	Model:        reviewResult.Model,
	OverallScore: reviewResult.OverallScore,
	MaxScore:     reviewResult.MaxScore,
	Summary:      reviewResult.Summary,
	Issues:       reviewResult.Issues,
	Strengths:    reviewResult.Strengths,
}
for _, c := range reviewResult.Scores.Criteria {
	details.Criteria = append(details.Criteria, ReviewCriterion{
		Name: c.Name, Passed: c.Passed, Reason: c.Reason,
	})
}
result.ReviewDetails = details
return result, nil
}

// criteriaScore converts a ReviewResult to a 0.0–1.0 score.
func criteriaScore(r *review.ReviewResult) float64 {
if r.MaxScore == 0 {
	return 0
}
return float64(r.OverallScore) / float64(r.MaxScore)
}

// copyDirToTemp creates a temporary copy of the source directory.
func copyDirToTemp(src string) (string, error) {
dir, err := os.MkdirTemp("", "hyoka-review-*")
if err != nil {
	return "", fmt.Errorf("creating review workspace: %w", err)
}
if err := copyDirContents(src, dir); err != nil {
	os.RemoveAll(dir)
	return "", fmt.Errorf("copying files to review workspace: %w", err)
}
return dir, nil
}

// copyDirContents copies all files from src to dst recursively.
func copyDirContents(src, dst string) error {
entries, err := os.ReadDir(src)
if err != nil {
	return err
}
for _, entry := range entries {
	srcPath := src + "/" + entry.Name()
	dstPath := dst + "/" + entry.Name()
	if entry.IsDir() {
		if err := os.MkdirAll(dstPath, 0755); err != nil {
			return err
		}
		if err := copyDirContents(srcPath, dstPath); err != nil {
			return err
		}
	} else {
		data, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dstPath, data, 0644); err != nil {
			return err
		}
	}
}
return nil
}
