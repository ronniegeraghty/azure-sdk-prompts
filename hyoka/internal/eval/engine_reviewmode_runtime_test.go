package eval

import (
	"context"
	"sync"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

// recordingReviewer captures Reviewer / MultiBucketReviewer calls so an
// integration-style test can prove --review-mode isolated actually fires
// the bucketed code path end-to-end. This is the runtime check Switch's
// review (PR #603) called out: the previous attempt (#578) was reverted by
// #587 because the flag had no runtime effect.
type recordingReviewer struct {
	mu               sync.Mutex
	reviewCalls      int
	reviewBucketCalls int
	lastBucketCount  int
}

func (r *recordingReviewer) Review(_ context.Context, _, _, _, _ string) (*review.ReviewResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reviewCalls++
	return stubResult(1, 1), nil
}

func (r *recordingReviewer) ReviewBuckets(_ context.Context, _, _, _ string, buckets []review.Bucket) (*review.ReviewResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reviewBucketCalls++
	r.lastBucketCount = len(buckets)
	return stubResult(len(buckets), len(buckets)), nil
}

func stubResult(score, max int) *review.ReviewResult {
	return &review.ReviewResult{
		Scores: review.ReviewScores{Criteria: []review.CriterionResult{
			{Name: "stub", Passed: true},
		}},
		OverallScore: score,
		MaxScore:     max,
		Summary:      "ok",
	}
}

// TestIntegrationReviewModeIsolatedFiresBuckets is the end-to-end runtime
// regression check requested by Switch on PR #603. It runs the full engine
// pipeline (Run → evaluatePrompt → grading phase → PromptReviewGrader) with
// EngineOptions.ReviewMode = "isolated" and a grader config containing one
// isolate-marked grader, then asserts ReviewBuckets() (not Review()) was
// invoked with multiple buckets.
//
// This is the failure mode #587's revert was specifically guarding against:
// a refactor that drops ReviewMode from the chain would still pass every
// unit test in isolation but would silently bypass ReviewBuckets at runtime.
func TestIntegrationReviewModeIsolatedFiresBuckets(t *testing.T) {
	outputDir := t.TempDir()
	rec := &recordingReviewer{}
	factory := func(_ *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return rec, nil, nil
	}

	engine := NewEngineWithReviewerFactory(&StubRunner{}, factory, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		ReviewMode: criteria.ReviewModeIsolated,
	}))
	// Inject a unified Bundle directly so we don't need a CriteriaDir on disk.
	// One isolate-marked prompt grader + one regular prompt grader → 2 buckets
	// in isolated mode.
	engine.graderBundle = bundleWith(
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "security", Prompt: "no hardcoded secrets", Isolate: true},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "format", Prompt: "code is well formatted"},
	)

	prompts := []*prompt.Prompt{{
		ID:                 "isolated-runtime-check",
		PromptText:         "Demo prompt",
		EvaluationCriteria: "Must build",
		Properties: map[string]string{
			"language": "python",
			"plane":    "data-plane",
			"service":  "identity",
		},
	}}
	configs := []config.ToolConfig{{
		Name:      "test-config",
		Generator: &config.GeneratorConfig{Model: "gpt-4"},
		Reviewer:  &config.ReviewerConfig{Models: []string{"gpt-4"}},
	}}

	if _, err := engine.Run(context.Background(), prompts, configs); err != nil {
		t.Fatalf("engine.Run() error: %v", err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()
	if rec.reviewBucketCalls < 1 {
		t.Fatalf("expected ReviewBuckets() to be called at least once in isolated mode, got %d (Review calls=%d)",
			rec.reviewBucketCalls, rec.reviewCalls)
	}
	if rec.lastBucketCount < 2 {
		t.Errorf("expected at least 2 buckets passed to ReviewBuckets() (isolated + combined), got %d", rec.lastBucketCount)
	}
}

// TestIntegrationReviewModeCombinedSkipsBuckets verifies that combined mode
// produces 2 buckets (prompt + criteria files) and calls ReviewBuckets(),
// not the single Review() path. This reflects the change to always separate
// prompt-frontmatter criteria from criteria-file entries.
func TestIntegrationReviewModeCombinedSkipsBuckets(t *testing.T) {
	outputDir := t.TempDir()
	rec := &recordingReviewer{}
	factory := func(_ *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return rec, nil, nil
	}

	engine := NewEngineWithReviewerFactory(&StubRunner{}, factory, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		ReviewMode: criteria.ReviewModeCombined,
	}))
	engine.graderBundle = bundleWith(
		// Even with isolate:true, combined mode should ignore it.
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "security", Prompt: "no hardcoded secrets", Isolate: true},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "format", Prompt: "code is well formatted"},
	)

	prompts := []*prompt.Prompt{{
		ID: "combined-runtime-check", PromptText: "Demo", EvaluationCriteria: "Must build",
		Properties: map[string]string{"language": "python"},
	}}
	configs := []config.ToolConfig{{
		Name:      "test-config",
		Generator: &config.GeneratorConfig{Model: "gpt-4"},
		Reviewer:  &config.ReviewerConfig{Models: []string{"gpt-4"}},
	}}

	if _, err := engine.Run(context.Background(), prompts, configs); err != nil {
		t.Fatalf("engine.Run() error: %v", err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()
	if rec.reviewBucketCalls < 1 {
		t.Errorf("combined mode should call ReviewBuckets() (2 buckets: prompt + criteria files), got %d calls", rec.reviewBucketCalls)
	}
	if rec.lastBucketCount != 2 {
		t.Errorf("expected 2 buckets (prompt + criteria files) in combined mode, got %d", rec.lastBucketCount)
	}
}
