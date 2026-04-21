package eval

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// captureSlog swaps the default slog logger for one that writes to buf.
// The returned restore function reinstates whatever was set before.
// Tests that depend on slog output should call it like:
//
//	buf, restore := captureSlog(t)
//	defer restore()
func captureSlog(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	buf := &bytes.Buffer{}
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	return buf, func() { slog.SetDefault(prev) }
}

// TestEngineReviewBuckets_Combined verifies that combined mode (default)
// produces exactly one review bucket containing all matched grader criteria
// merged with the prompt's own evaluation criteria — byte-identical to the
// pre-#580 single-session path.
func TestEngineReviewBuckets_Combined(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeCombined,
	}))
	e.graderConfigs = []criteria.GraderConfig{{
		Graders: []criteria.GraderEntry{
			{Name: "a", Prompt: "criterion A", Isolate: true}, // isolate flag IGNORED in combined mode
			{Name: "b", Prompt: "criterion B"},
		},
	}}

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "prompt-level criteria"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 1 {
		t.Fatalf("combined mode: expected 1 bucket, got %d", len(buckets))
	}
	if !strings.Contains(buckets[0].Criteria, "criterion A") ||
		!strings.Contains(buckets[0].Criteria, "criterion B") ||
		!strings.Contains(buckets[0].Criteria, "prompt-level criteria") {
		t.Errorf("combined bucket missing expected criteria text: %q", buckets[0].Criteria)
	}
}

// TestEngineReviewBuckets_DefaultModeMatchesCombined verifies that an empty
// ReviewMode (the zero value when --review-mode is not passed) is treated
// as combined — no isolation, exactly one bucket.
func TestEngineReviewBuckets_DefaultModeMatchesCombined(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	e.graderConfigs = []criteria.GraderConfig{{
		Graders: []criteria.GraderEntry{
			{Name: "a", Prompt: "A", Isolate: true},
			{Name: "b", Prompt: "B"},
		},
	}}

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)
	if len(buckets) != 1 {
		t.Fatalf("default (empty) mode should behave like combined: expected 1 bucket, got %d", len(buckets))
	}
}

// TestEngineReviewBuckets_IsolatedWithIsolation verifies that isolated mode
// produces one bucket per grader marked isolate plus a combined bucket for
// the rest (the wiring layer between EngineOptions.ReviewMode and
// criteria.BuildReviewBuckets — the surface a future regression would
// silently break).
func TestEngineReviewBuckets_IsolatedWithIsolation(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeIsolated,
	}))
	e.graderConfigs = []criteria.GraderConfig{{
		Graders: []criteria.GraderEntry{
			{Name: "security", Prompt: "no hardcoded secrets", Isolate: true},
			{Name: "format", Prompt: "code is formatted"},
			{Name: "tests", Prompt: "tests exist"},
		},
	}}

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 2 {
		t.Fatalf("isolated mode with 1 isolate-marked grader expected 2 buckets (1 isolated + 1 combined), got %d", len(buckets))
	}
	var sawIsolated, sawCombined bool
	for _, b := range buckets {
		if b.Name == "security" {
			sawIsolated = true
			if !strings.Contains(b.Criteria, "no hardcoded secrets") {
				t.Errorf("security bucket missing its criterion: %q", b.Criteria)
			}
		}
		if b.Name == "combined" {
			sawCombined = true
			if !strings.Contains(b.Criteria, "code is formatted") ||
				!strings.Contains(b.Criteria, "tests exist") {
				t.Errorf("combined bucket missing non-isolated criteria: %q", b.Criteria)
			}
		}
	}
	if !sawIsolated {
		t.Error("expected an isolated bucket named 'security'")
	}
	if !sawCombined {
		t.Error("expected a 'combined' bucket for non-isolated graders")
	}
}

// TestEngineReviewBuckets_IsolatedDegradesWithoutIsolation verifies the
// "observably no-op rather than silently dead" promise: when --review-mode
// isolated is requested but nothing is marked isolate, the engine logs a
// slog.Warn AND falls back to producing a single combined bucket.
//
// This is the regression check that PR #587 specifically wanted — without
// this, dropping ReviewMode from the chain would compile and pass tests.
func TestEngineReviewBuckets_IsolatedDegradesWithoutIsolation(t *testing.T) {
	buf, restore := captureSlog(t)
	defer restore()

	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeIsolated,
	}))
	e.graderConfigs = []criteria.GraderConfig{{
		Graders: []criteria.GraderEntry{
			{Name: "a", Prompt: "A"},
			{Name: "b", Prompt: "B"},
		},
	}}

	p := &prompt.Prompt{ID: "prompt-X", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 1 {
		t.Fatalf("degraded path expected 1 bucket, got %d", len(buckets))
	}
	logs := buf.String()
	if !strings.Contains(logs, "review-mode=isolated requested but no graders or groups are marked isolate") {
		t.Errorf("expected slog.Warn about degraded isolated mode, got logs: %q", logs)
	}
	if !strings.Contains(logs, "prompt-X") {
		t.Errorf("expected warning to include prompt_id=prompt-X, got: %q", logs)
	}
}

// TestEngineReviewBuckets_NoGraders verifies the no-graders edge case
// returns a single combined bucket carrying just the prompt's own criteria
// (so the legacy review path is preserved when a prompt has only inline
// criteria and no attribute-matched graders).
func TestEngineReviewBuckets_NoGraders(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	p := &prompt.Prompt{ID: "p", EvaluationCriteria: "just this"}
	buckets := e.reviewBuckets(p, nil)
	if len(buckets) != 1 {
		t.Fatalf("expected 1 bucket from prompt criteria alone, got %d", len(buckets))
	}
	if !strings.Contains(buckets[0].Criteria, "just this") {
		t.Errorf("bucket missing prompt criteria: %q", buckets[0].Criteria)
	}
}
