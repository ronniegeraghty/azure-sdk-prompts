package graders

import (
"context"
"fmt"
"log/slog"
"regexp"
"strings"

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

// Grade runs the review and converts the result into a GraderResult. The
// engine guarantees input.WorkspacePath is an isolated copy of the generated
// workspace (see eval.IsolateGraderWorkspace), so the reviewer can write
// annotated files into it without affecting the canonical workspace or other
// graders. The consolidated review score maps to the grader score; panel
// member results are stored in ReviewDetails.PanelResults.
func (g *PromptReviewGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
	result := GraderResult{
		Kind: KindPromptReview,
		Name: g.name,
	}

	// Engine owns workspace isolation and lifecycle. Record the path so
	// the engine can read the reviewer's annotated files after Grade
	// returns.
	reviewWorkDir := input.WorkspacePath
	g.LastReviewWorkDir = reviewWorkDir

	if g.panelReviewer != nil {
		return g.gradePanel(ctx, input, reviewWorkDir, result, input.GeneratorArtifact)
	}
	if g.reviewer != nil {
		return g.gradeSingle(ctx, input, reviewWorkDir, result, input.GeneratorArtifact)
	}
	return result, fmt.Errorf("no reviewer configured")
}

// CleanupWorkspace is retained for backward compatibility with existing
// callers/tests. The engine now owns workspace lifecycle (see
// eval.IsolateGraderWorkspace), so this method is a no-op. Callers that still
// invoke it remain safe but no longer have any effect.
//
// Deprecated: workspace lifecycle is engine-owned. Do not call.
func (g *PromptReviewGrader) CleanupWorkspace() {
	// Intentionally a no-op. The engine isolates the workspace before
	// calling Grade and removes it afterwards.
	g.LastReviewWorkDir = ""
}

func (g *PromptReviewGrader) gradePanel(ctx context.Context, input GraderInput, workDir string, result GraderResult, artifact *review.GeneratorArtifact) (GraderResult, error) {
	models := g.panelReviewer.Models()
	slog.Debug("Review grader starting panel", "models", models)

	var (
		panel        []review.ReviewResult
		consolidated *review.ReviewResult
		err          error
	)
	if len(input.EvalCriteriaBuckets) > 1 {
		slog.Info("Review grader using bucketed panel review", "bucket_count", len(input.EvalCriteriaBuckets))
		panel, consolidated, err = g.panelReviewer.ReviewPanelBuckets(
			ctx, input.OriginalPrompt, workDir, input.ReferenceDir, toReviewBuckets(input.EvalCriteriaBuckets), artifact,
		)
	} else {
		criteria := input.EvalCriteria
		if criteria == "" && len(input.EvalCriteriaBuckets) == 1 {
			criteria = input.EvalCriteriaBuckets[0].Criteria
		}
		panel, consolidated, err = g.panelReviewer.ReviewPanel(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, criteria, artifact)
	}
	if err != nil {
		return result, fmt.Errorf("review panel failed: %w", err)
	}

	g.LastPanel = panel
	g.LastConsolidated = consolidated

	logCriteriaCountMismatch(g.name, input, len(consolidated.Scores.Criteria))

	// Per v4 spec: one Check per criterion, Weight = criterion max
	var checks []GraderCheck
	for i, c := range consolidated.Scores.Criteria {
		label := c.Name
		if label == "" {
			label = fmt.Sprintf("check %d", i+1)
		}
		pointMsg := ""
		if !c.Passed {
			pointMsg = c.Reason
		}
		// Default weight to 1.0 (review package doesn't track per-criterion max)
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    c.Passed,
			Message: pointMsg,
			Weight:  1.0,
		})
	}
	
	if len(checks) == 0 {
		// Fallback: review with no criteria — synthesize one point
		checks = []GraderCheck{{
			Label:   "consensus",
			Pass:    consolidated.Scores.AllPassed(),
			Message: consolidated.Summary,
			Weight:  1.0,
		}}
	}

	// Build ReviewExtras
	var panelResults []ReviewPanelResult
	for _, p := range panel {
		var criteria []ReviewCriterionResult
		for _, c := range p.Scores.Criteria {
			criteria = append(criteria, ReviewCriterionResult{
				Name:   c.Name,
				Passed: c.Passed,
				Reason: c.Reason,
				Weight: 1, // Default weight (review package doesn't track per-criterion max)
			})
		}
		panelResults = append(panelResults, ReviewPanelResult{
			Model:     p.Model,
			Score:     p.OverallScore,
			Pass:      p.Scores.AllPassed(),
			Issues:    p.Issues,
			Strengths: p.Strengths,
			Criteria:  criteria,
		})
	}

	extras := &GraderExtras{
		Review: &ReviewExtras{
			Model:       consolidated.Model,
			Summary:     consolidated.Summary,
			IsConsensus: true,
			PanelResults: panelResults,
			Issues:       consolidated.Issues,
			Strengths:    consolidated.Strengths,
		},
	}

	msg := consolidated.Summary
	return NewResult(KindPromptReview, g.name, input.Config, checks, msg, extras), nil
}

func (g *PromptReviewGrader) gradeSingle(ctx context.Context, input GraderInput, workDir string, result GraderResult, artifact *review.GeneratorArtifact) (GraderResult, error) {
	slog.Debug("Review grader starting single review")

	var (
		reviewResult *review.ReviewResult
		err          error
	)
	if len(input.EvalCriteriaBuckets) > 1 {
		if mb, ok := g.reviewer.(review.MultiBucketReviewer); ok {
			slog.Info("Review grader using bucketed single review", "bucket_count", len(input.EvalCriteriaBuckets))
			reviewResult, err = mb.ReviewBuckets(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, toReviewBuckets(input.EvalCriteriaBuckets), artifact)
		} else {
			slog.Warn("Reviewer does not support buckets; collapsing to combined criteria")
			reviewResult, err = g.reviewer.Review(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, joinCriteria(input.EvalCriteriaBuckets), artifact)
		}
	} else {
		criteria := input.EvalCriteria
		if criteria == "" && len(input.EvalCriteriaBuckets) == 1 {
			criteria = input.EvalCriteriaBuckets[0].Criteria
		}
		reviewResult, err = g.reviewer.Review(ctx, input.OriginalPrompt, workDir, input.ReferenceDir, criteria, artifact)
	}
	if err != nil {
		return result, fmt.Errorf("review failed: %w", err)
	}

	g.LastConsolidated = reviewResult

	logCriteriaCountMismatch(g.name, input, len(reviewResult.Scores.Criteria))

	// Per v4 spec: one Check per criterion, Weight = criterion max
	var checks []GraderCheck
	for i, c := range reviewResult.Scores.Criteria {
		label := c.Name
		if label == "" {
			label = fmt.Sprintf("check %d", i+1)
		}
		pointMsg := ""
		if !c.Passed {
			pointMsg = c.Reason
		}
		// Default weight to 1.0 (review package doesn't track per-criterion max)
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    c.Passed,
			Message: pointMsg,
			Weight:  1.0,
		})
	}
	
	if len(checks) == 0 {
		checks = []GraderCheck{{
			Label:   "review",
			Pass:    reviewResult.Scores.AllPassed(),
			Message: reviewResult.Summary,
			Weight:  1.0,
		}}
	}

	// Build ReviewExtras
	extras := &GraderExtras{
		Review: &ReviewExtras{
			Model:       reviewResult.Model,
			Summary:     reviewResult.Summary,
			IsConsensus: false,
			Issues:      reviewResult.Issues,
			Strengths:   reviewResult.Strengths,
		},
	}

	msg := reviewResult.Summary
	return NewResult(KindPromptReview, g.name, input.Config, checks, msg, extras), nil
}

// toReviewBuckets converts grader-package buckets to review-package buckets.
func toReviewBuckets(buckets []ReviewBucket) []review.Bucket {
	out := make([]review.Bucket, 0, len(buckets))
	for _, b := range buckets {
		out = append(out, review.Bucket{Name: b.Name, Criteria: b.Criteria})
	}
	return out
}

// joinCriteria concatenates bucket criteria into a single string for reviewers
// that do not implement MultiBucketReviewer (degraded path).
func joinCriteria(buckets []ReviewBucket) string {
	parts := make([]string, 0, len(buckets))
	for _, b := range buckets {
		if b.Criteria == "" {
			continue
		}
		parts = append(parts, b.Criteria)
	}
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += "\n\n"
		}
		out += p
	}
	return out
}

// criteriaScore converts a ReviewResult to a 0.0–1.0 score.
func criteriaScore(r *review.ReviewResult) float64 {
if r.MaxScore == 0 {
	return 0
}
return float64(r.OverallScore) / float64(r.MaxScore)
}

// numberedItemRe matches lines of the form "<n>. " (any leading whitespace).
// Used to estimate how many checks the LLM judge was asked to score.
var numberedItemRe = regexp.MustCompile(`(?m)^[ \t]*\d+\.\s+\S`)

// expectedCriteriaCount returns a best-effort count of the leaf criteria
// rendered in the bucket text. It prefers indented (nested) numbered items
// when present (the new prompt+checks rendering), and falls back to top-level
// numbered items (the legacy single-check / prompt-file rendering).
func expectedCriteriaCount(criteria string) int {
	if strings.TrimSpace(criteria) == "" {
		return 0
	}
	nested := 0
	top := 0
	for _, m := range numberedItemRe.FindAllString(criteria, -1) {
		if len(m) > 0 && (m[0] == ' ' || m[0] == '\t') {
			nested++
		} else {
			top++
		}
	}
	if nested > 0 {
		return nested
	}
	return top
}

// logCriteriaCountMismatch emits a debug log when the LLM judge returned a
// different number of criteria than the bucket text asked for. This is a
// flake-detection signal — the grader still uses whatever criteria came back.
func logCriteriaCountMismatch(graderName string, input GraderInput, returned int) {
	expected := 0
	if len(input.EvalCriteriaBuckets) > 0 {
		for _, b := range input.EvalCriteriaBuckets {
			expected += expectedCriteriaCount(b.Criteria)
		}
	} else {
		expected = expectedCriteriaCount(input.EvalCriteria)
	}
	if expected > 0 && expected != returned {
		slog.Debug("Review judge returned criterion count differs from sent",
			"grader", graderName,
			"expected", expected,
			"returned", returned,
		)
	}
}
