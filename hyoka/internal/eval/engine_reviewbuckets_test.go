package eval

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// captureSlog swaps the default slog logger for one that writes to buf.
// The returned restore function reinstates whatever was set before.
func captureSlog(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	buf := &bytes.Buffer{}
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	return buf, func() { slog.SetDefault(prev) }
}

// bundleWith builds a minimal *criteria.Bundle from one file's top-level
// entries, for tests that need to populate the engine's grader bundle
// without touching disk.
func bundleWith(entries ...criteria.UnifiedGraderEntry) *criteria.Bundle {
	return &criteria.Bundle{
		Configs: []criteria.UnifiedGraderConfig{{
			Graders: entries,
			Source:  "test://inline",
		}},
	}
}

// TestEngineReviewBuckets_Combined verifies that combined mode produces
// one bucket per matched entry: one for prompt-frontmatter criteria and one
// for each matched criteria-file entry.
func TestEngineReviewBuckets_Combined(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeCombined,
	}))
	e.graderBundle = bundleWith(
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "a", Prompt: "criterion A", Isolate: true},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "b", Prompt: "criterion B"},
	)

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "prompt-level criteria"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 3 {
		t.Fatalf("combined mode: expected 3 buckets (prompt + 2 per-entry), got %d", len(buckets))
	}
	// First bucket: prompt frontmatter criteria
	if buckets[0].Name != "Criteria from prompt file" {
		t.Errorf("expected first bucket 'Criteria from prompt file', got %q", buckets[0].Name)
	}
	if !strings.Contains(buckets[0].Criteria, "prompt-level criteria") {
		t.Errorf("prompt bucket missing 'prompt-level criteria': %q", buckets[0].Criteria)
	}
	// Second bucket: per-entry "a"
	if buckets[1].Name != "a" {
		t.Errorf("expected second bucket 'a', got %q", buckets[1].Name)
	}
	if !strings.Contains(buckets[1].Criteria, "criterion A") {
		t.Errorf("bucket 'a' missing 'criterion A': %q", buckets[1].Criteria)
	}
	// Third bucket: per-entry "b"
	if buckets[2].Name != "b" {
		t.Errorf("expected third bucket 'b', got %q", buckets[2].Name)
	}
	if !strings.Contains(buckets[2].Criteria, "criterion B") {
		t.Errorf("bucket 'b' missing 'criterion B': %q", buckets[2].Criteria)
	}
}

// TestEngineReviewBuckets_DefaultModeMatchesCombined verifies that an empty
// ReviewMode is treated as combined — prompt criteria get their own bucket,
// each matched entry gets its own bucket.
func TestEngineReviewBuckets_DefaultModeMatchesCombined(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	e.graderBundle = bundleWith(
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "a", Prompt: "A", Isolate: true},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "b", Prompt: "B"},
	)

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)
	if len(buckets) != 3 {
		t.Fatalf("default (empty) mode should behave like combined: expected 3 buckets, got %d", len(buckets))
	}
}

// TestEngineReviewBuckets_IsolatedWithIsolation verifies that isolated mode
// produces one bucket for prompt criteria, one per grader marked isolate, and
// a combined bucket for the rest of the criteria-file entries.
func TestEngineReviewBuckets_IsolatedWithIsolation(t *testing.T) {
	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeIsolated,
	}))
	e.graderBundle = bundleWith(
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "security", Prompt: "no hardcoded secrets", Isolate: true},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "format", Prompt: "code is formatted"},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "tests", Prompt: "tests exist"},
	)

	p := &prompt.Prompt{ID: "p1", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 3 {
		t.Fatalf("isolated mode expected 3 buckets (prompt + 1 isolated + 1 combined), got %d", len(buckets))
	}
	var sawPrompt, sawIsolated, sawCombined bool
	for _, b := range buckets {
		if b.Name == "Criteria from prompt file" {
			sawPrompt = true
			if !strings.Contains(b.Criteria, "pc") {
				t.Errorf("prompt bucket missing 'pc': %q", b.Criteria)
			}
		}
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
	if !sawPrompt {
		t.Error("expected a prompt bucket named 'Criteria from prompt file'")
	}
	if !sawIsolated {
		t.Error("expected an isolated bucket named 'security'")
	}
	if !sawCombined {
		t.Error("expected a 'combined' bucket for non-isolated graders")
	}
}

// TestEngineReviewBuckets_IsolatedDegradesWithoutIsolation verifies that
// when --review-mode isolated is requested but nothing is marked isolate,
// the engine logs a slog.Warn and produces per-entry buckets: prompt criteria
// and one bucket per criteria-file entry.
func TestEngineReviewBuckets_IsolatedDegradesWithoutIsolation(t *testing.T) {
	buf, restore := captureSlog(t)
	defer restore()

	e := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		ReviewMode: criteria.ReviewModeIsolated,
	}))
	e.graderBundle = bundleWith(
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "a", Prompt: "A"},
		criteria.UnifiedGraderEntry{Type: graders.KindPrompt, Name: "b", Prompt: "B"},
	)

	p := &prompt.Prompt{ID: "prompt-X", EvaluationCriteria: "pc"}
	buckets := e.reviewBuckets(p, nil)

	if len(buckets) != 3 {
		t.Fatalf("degraded path expected 3 buckets (prompt + 2 per-entry), got %d", len(buckets))
	}
	logs := buf.String()
	if !strings.Contains(logs, "review-mode=isolated requested but no graders or groups are marked isolate") {
		t.Errorf("expected slog.Warn about degraded isolated mode, got logs: %q", logs)
	}
	if !strings.Contains(logs, "prompt-X") {
		t.Errorf("expected warning to include prompt_id=prompt-X, got: %q", logs)
	}
}

// TestEngineReviewBuckets_NoGraders verifies the no-graders edge case.
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
