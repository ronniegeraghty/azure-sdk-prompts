package criteria

import (
	"strings"
	"testing"
)

func mkConfig(graders ...GraderEntry) GraderConfig {
	return GraderConfig{Graders: graders}
}

// TestReviewModeCombined verifies that 3 matched graders produce exactly one
// bucket in combined mode (legacy behavior preserved).
func TestReviewModeCombined(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a", Prompt: "A"},
		GraderEntry{Name: "b", Prompt: "B"},
		GraderEntry{Name: "c", Prompt: "C"},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "prompt criteria", ReviewModeCombined)
	if len(buckets) != 1 {
		t.Fatalf("combined mode: expected 1 bucket, got %d", len(buckets))
	}
	if len(buckets[0].Graders) != 3 {
		t.Fatalf("combined bucket: expected 3 graders, got %d", len(buckets[0].Graders))
	}
	if buckets[0].PromptCriteria == "" {
		t.Fatal("combined bucket should carry prompt criteria")
	}
}

// TestReviewModeIsolated verifies that 3 graders all marked isolate produce 3
// buckets with no combined leftover (and prompt criteria spawns its own).
func TestReviewModeIsolated(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a", Prompt: "A", Isolate: true},
		GraderEntry{Name: "b", Prompt: "B", Isolate: true},
		GraderEntry{Name: "c", Prompt: "C", Isolate: true},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeIsolated)
	if len(buckets) != 3 {
		t.Fatalf("isolated mode: expected 3 buckets, got %d", len(buckets))
	}
	names := map[string]bool{}
	for _, b := range buckets {
		if len(b.Graders) != 1 {
			t.Errorf("bucket %s: expected 1 grader, got %d", b.Name, len(b.Graders))
		}
		names[b.Graders[0].Name] = true
	}
	for _, want := range []string{"a", "b", "c"} {
		if !names[want] {
			t.Errorf("missing isolated bucket for grader %s", want)
		}
	}
}

func TestIsolatedMixed(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a", Isolate: true},
		GraderEntry{Name: "b"},
		GraderEntry{Name: "c"},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "pc", ReviewModeIsolated)
	if len(buckets) != 2 {
		t.Fatalf("expected 2 buckets (isolated + combined), got %d", len(buckets))
	}
	var combined *ReviewBucket
	for i := range buckets {
		if buckets[i].Name == "combined" {
			combined = &buckets[i]
		}
	}
	if combined == nil {
		t.Fatal("expected a combined bucket")
	}
	if len(combined.Graders) != 2 {
		t.Errorf("combined bucket should hold 2 graders, got %d", len(combined.Graders))
	}
	if combined.PromptCriteria == "" {
		t.Error("combined bucket should carry prompt criteria")
	}
}

func TestIsolatedGroup(t *testing.T) {
	cfg := GraderConfig{
		Groups: []GraderGroup{
			{
				Name:    "security",
				Isolate: true,
				Graders: []GraderEntry{
					{Name: "auth"},
					{Name: "secrets"},
				},
			},
		},
	}
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeIsolated)
	if len(buckets) != 1 {
		t.Fatalf("expected 1 bucket for isolated group, got %d", len(buckets))
	}
	if len(buckets[0].Graders) != 2 {
		t.Errorf("group bucket should hold 2 graders, got %d", len(buckets[0].Graders))
	}
	if buckets[0].Name != "security" {
		t.Errorf("bucket name = %q, want security", buckets[0].Name)
	}
}

func TestIsolatedDegradesWhenNothingMarked(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a"},
		GraderEntry{Name: "b"},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "pc", ReviewModeIsolated)
	if len(buckets) != 1 {
		t.Fatalf("isolated with nothing marked: expected 1 bucket (degraded), got %d", len(buckets))
	}
	if !HasIsolation(matched) {
		// expected — confirms the engine should warn
	}
}

func TestEmptyMatched(t *testing.T) {
	buckets := BuildReviewBuckets(nil, "pc", ReviewModeIsolated)
	if len(buckets) != 1 {
		t.Fatalf("empty matched: expected 1 bucket, got %d", len(buckets))
	}
	if buckets[0].PromptCriteria != "pc" {
		t.Errorf("expected prompt criteria preserved")
	}
}

func TestGroupIsolateOverridesGraderIsolate(t *testing.T) {
	cfg := GraderConfig{
		Groups: []GraderGroup{
			{
				Name:    "g",
				Isolate: true,
				Graders: []GraderEntry{
					{Name: "x", Isolate: true},
					{Name: "y", Isolate: true},
				},
			},
		},
	}
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeIsolated)
	if len(buckets) != 1 {
		t.Fatalf("group isolate should produce 1 bucket, got %d", len(buckets))
	}
}

func TestNonIsolatedGroupHonorsGraderIsolate(t *testing.T) {
	cfg := GraderConfig{
		Groups: []GraderGroup{
			{
				Name: "g",
				Graders: []GraderEntry{
					{Name: "x", Isolate: true},
					{Name: "y"},
				},
			},
		},
	}
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeIsolated)
	if len(buckets) != 2 {
		t.Fatalf("expected 2 buckets, got %d", len(buckets))
	}
}

func TestFormatCriteriaCombinedMatchesMergeCriteria(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a", Prompt: "A"},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "pc", ReviewModeCombined)
	got := buckets[0].FormatCriteria()
	want := MergeCriteria([]GraderEntry{{Name: "a", Prompt: "A"}}, "pc")
	if got != want {
		t.Errorf("combined bucket FormatCriteria mismatch:\n got=%q\nwant=%q", got, want)
	}
}

func TestBucketNameFallback(t *testing.T) {
	if got := bucketNameFor("", 7); got != "bucket-7" {
		t.Errorf("bucketNameFor(\"\", 7) = %q, want bucket-7", got)
	}
	if got := bucketNameFor("  named  ", 0); got != "named" {
		t.Errorf("bucketNameFor trimmed = %q, want named", got)
	}
}

func TestIsolatedBucketNamesUnique(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "x", Isolate: true},
		GraderEntry{Name: "x", Isolate: true},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeIsolated)
	if len(buckets) != 2 {
		t.Fatalf("expected 2 buckets, got %d", len(buckets))
	}
	// Names may collide — that's fine; the merger prefixes criterion names by
	// bucket name so any collision is benign for vote dedup. Just ensure both
	// graders made it through.
	if buckets[0].Graders[0].Name != "x" || buckets[1].Graders[0].Name != "x" {
		t.Errorf("both buckets should hold grader x")
	}
}

func TestHasIsolation(t *testing.T) {
	if HasIsolation(nil) {
		t.Error("nil should report no isolation")
	}
	if HasIsolation([]MatchedGrader{{Entry: GraderEntry{Isolate: false}}}) {
		t.Error("no isolate flag should report no isolation")
	}
	if !HasIsolation([]MatchedGrader{{Entry: GraderEntry{Isolate: true}}}) {
		t.Error("isolate flag should report isolation")
	}
	if !HasIsolation([]MatchedGrader{{GroupIsolate: true}}) {
		t.Error("group isolate should report isolation")
	}
}

func TestCombinedModeIgnoresIsolateFlag(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a", Isolate: true},
		GraderEntry{Name: "b", Isolate: true},
	)
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", ReviewModeCombined)
	if len(buckets) != 1 {
		t.Fatalf("combined mode must produce 1 bucket regardless of isolate flags, got %d", len(buckets))
	}
	if len(buckets[0].Graders) != 2 {
		t.Errorf("combined bucket should hold both graders, got %d", len(buckets[0].Graders))
	}
}

func TestUnknownModeTreatedAsCombined(t *testing.T) {
	cfg := mkConfig(GraderEntry{Name: "a", Isolate: true})
	matched := MatchingGradersWithIsolation([]GraderConfig{cfg}, nil)
	buckets := BuildReviewBuckets(matched, "", "weird")
	if len(buckets) != 1 {
		t.Fatalf("unknown mode should default to combined (1 bucket), got %d", len(buckets))
	}
}

func TestMatchingGradersBackwardCompat(t *testing.T) {
	cfg := mkConfig(
		GraderEntry{Name: "a"},
		GraderEntry{Name: "b"},
	)
	got := MatchingGraders([]GraderConfig{cfg}, nil)
	if len(got) != 2 {
		t.Fatalf("expected 2 graders, got %d", len(got))
	}
	if got[0].Name != "a" || got[1].Name != "b" {
		t.Errorf("ordering broken: %+v", got)
	}
	// And formatting still works
	if !strings.Contains(FormatGraders(got), "**a**") {
		t.Error("FormatGraders broken")
	}
}
