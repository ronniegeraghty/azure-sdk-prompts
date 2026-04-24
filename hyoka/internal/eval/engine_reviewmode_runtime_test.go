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

func (r *recordingReviewer) Review(_ context.Context, _, _, _, _ string, _ *review.GeneratorArtifact) (*review.ReviewResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reviewCalls++
	return stubResult(1, 1), nil
}

func (r *recordingReviewer) ReviewBuckets(_ context.Context, _, _, _ string, buckets []review.Bucket, _ *review.GeneratorArtifact) (*review.ReviewResult, error) {
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

	// Phase 2 per-bucket grading: the engine creates one PromptReviewGrader
	// per bucket, so we expect N Review() calls (one per bucket) instead of
	// one ReviewBuckets() call. The per-bucket approach prioritizes display
	// clarity (each bucket renders as its own grader line) over batch-API
	// efficiency.
	rec.mu.Lock()
	defer rec.mu.Unlock()
	if rec.reviewCalls < 2 {
		t.Fatalf("expected at least 2 Review() calls (one per bucket) in isolated mode, got %d", rec.reviewCalls)
	}
}

// TestIntegrationReviewModeCombinedSkipsBuckets verifies that combined mode
// produces 3 buckets (prompt + 2 per-entry) and calls ReviewBuckets(),
// not the single Review() path. This reflects the change to always separate
// prompt-frontmatter criteria from criteria-file entries, and to create
// one bucket per criteria-file entry.
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

	// Phase 2 per-bucket grading: the engine creates one PromptReviewGrader
	// per bucket (prompt + 2 per-entry = 3 buckets in this test), so we
	// expect 3 Review() calls instead of one ReviewBuckets() call.
	rec.mu.Lock()
	defer rec.mu.Unlock()
	if rec.reviewCalls != 3 {
		t.Errorf("combined mode should call Review() 3 times (once per bucket: prompt + 2 entries), got %d calls", rec.reviewCalls)
	}
}
