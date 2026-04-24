package criteria

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"strings"
	"testing"
)

func promptEntry(name, crit string, isolate bool) UnifiedGraderEntry {
	return UnifiedGraderEntry{Type: graders.KindPrompt, Name: name, Prompt: crit, Isolate: isolate}
}

func typedEntry(kind, name string) UnifiedGraderEntry {
	return UnifiedGraderEntry{Type: kind, Name: name, Weight: 1}
}

func TestMatchingUnifiedEntries_HonorsHierarchicalWhen(t *testing.T) {
	bundle := &Bundle{Configs: []UnifiedGraderConfig{{
		When: map[string]string{"language": "python"},
		Graders: []UnifiedGraderEntry{
			promptEntry("a", "A", false),
		},
		Groups: []UnifiedGraderGroup{{
			Name: "g",
			When: map[string]string{"plane": "data-plane"},
			Graders: []UnifiedGraderEntry{
				promptEntry("b", "B", false),
				{Type: graders.KindPrompt, Name: "c", Prompt: "C", When: map[string]string{"category": "crud"}},
			},
		}},
	}}}

	cases := []struct {
		name  string
		props map[string]string
		want  []string
	}{
		{"none", map[string]string{}, nil},
		{"python only", map[string]string{"language": "python"}, []string{"a"}},
		{"python + dp", map[string]string{"language": "python", "plane": "data-plane"}, []string{"a", "b"}},
		{"python + dp + crud", map[string]string{"language": "python", "plane": "data-plane", "category": "crud"}, []string{"a", "b", "c"}},
		{"wrong lang", map[string]string{"language": "go"}, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			matched := MatchingUnifiedEntries(bundle, tc.props)
			got := make([]string, 0, len(matched))
			for _, m := range matched {
				got = append(got, m.Entry.Name)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("index %d: got %q want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestPartitionMatched_SplitsPromptAndTyped(t *testing.T) {
	matched := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", false)},
		{Entry: typedEntry(graders.KindFile, "f1")},
		{Entry: promptEntry("b", "B", true)},
		{Entry: typedEntry(graders.KindOutputCheck, "oc1")},
	}
	prompts, typed := PartitionMatched(matched)
	if len(prompts) != 2 || prompts[0].Entry.Name != "a" || prompts[1].Entry.Name != "b" {
		t.Errorf("prompts wrong: %+v", prompts)
	}
	if len(typed) != 2 || typed[0].Entry.Name != "f1" || typed[1].Entry.Name != "oc1" {
		t.Errorf("typed wrong: %+v", typed)
	}
}

func TestHasUnifiedIsolation(t *testing.T) {
	none := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", false)},
		{Entry: promptEntry("b", "B", false)},
	}
	if HasUnifiedIsolation(none) {
		t.Error("expected false with no isolation")
	}
	withGrader := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", true)},
	}
	if !HasUnifiedIsolation(withGrader) {
		t.Error("expected true when a grader is marked isolate")
	}
	withGroup := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", false), GroupName: "g", GroupIsolate: true},
	}
	if !HasUnifiedIsolation(withGroup) {
		t.Error("expected true when a group is marked isolate")
	}
}

func TestBuildUnifiedReviewBuckets_CombinedMode(t *testing.T) {
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "criterion A", true)}, // isolate flag IGNORED in combined
		{Entry: promptEntry("b", "criterion B", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "prompt criteria", ReviewModeCombined)
	if len(buckets) != 2 {
		t.Fatalf("combined: expected 2 buckets (prompt + criteria files), got %d", len(buckets))
	}
	// First bucket should be prompt frontmatter criteria
	if buckets[0].Name != "Criteria from prompt file" {
		t.Errorf("expected first bucket name 'Criteria from prompt file', got %q", buckets[0].Name)
	}
	if !strings.Contains(buckets[0].Criteria, "prompt criteria") {
		t.Errorf("prompt bucket missing 'prompt criteria': %q", buckets[0].Criteria)
	}
	// Second bucket should be combined criteria-file entries
	if buckets[1].Name != "combined" {
		t.Errorf("expected second bucket name 'combined', got %q", buckets[1].Name)
	}
	for _, needle := range []string{"criterion A", "criterion B"} {
		if !strings.Contains(buckets[1].Criteria, needle) {
			t.Errorf("combined bucket missing %q: %q", needle, buckets[1].Criteria)
		}
	}
	// Prompt criteria should NOT be in the combined bucket
	if strings.Contains(buckets[1].Criteria, "prompt criteria") {
		t.Errorf("combined bucket should not contain prompt criteria: %q", buckets[1].Criteria)
	}
}

func TestBuildUnifiedReviewBuckets_IsolatedWithIsolation(t *testing.T) {
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("security", "no secrets", true)},
		{Entry: promptEntry("format", "format ok", false)},
		{Entry: promptEntry("tests", "tests exist", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "pc", ReviewModeIsolated)
	if len(buckets) != 3 {
		t.Fatalf("isolated: expected 3 buckets (prompt + security + combined), got %d", len(buckets))
	}
	var sawPrompt, sawSec, sawCombined bool
	for _, b := range buckets {
		switch b.Name {
		case "Criteria from prompt file":
			sawPrompt = true
			if !strings.Contains(b.Criteria, "pc") {
				t.Errorf("prompt bucket missing 'pc': %q", b.Criteria)
			}
			// Should NOT contain criteria-file entries
			if strings.Contains(b.Criteria, "no secrets") || strings.Contains(b.Criteria, "format ok") {
				t.Errorf("prompt bucket should not contain criteria-file entries: %q", b.Criteria)
			}
		case "security":
			sawSec = true
			if !strings.Contains(b.Criteria, "no secrets") {
				t.Errorf("security bucket content: %q", b.Criteria)
			}
		case "combined":
			sawCombined = true
			if !strings.Contains(b.Criteria, "format ok") || !strings.Contains(b.Criteria, "tests exist") {
				t.Errorf("combined bucket content: %q", b.Criteria)
			}
			// Should NOT contain prompt criteria
			if strings.Contains(b.Criteria, "pc") {
				t.Errorf("combined bucket should not contain prompt criteria: %q", b.Criteria)
			}
		}
	}
	if !sawPrompt || !sawSec || !sawCombined {
		t.Errorf("expected 'Criteria from prompt file', 'security', and 'combined' buckets; sawPrompt=%v sawSec=%v sawCombined=%v",
			sawPrompt, sawSec, sawCombined)
	}
}

func TestBuildUnifiedReviewBuckets_IsolatedDegradesToCombined(t *testing.T) {
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", false)},
		{Entry: promptEntry("b", "B", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "pc", ReviewModeIsolated)
	if len(buckets) != 2 {
		t.Fatalf("degraded: expected 2 buckets (prompt + combined criteria files), got %d", len(buckets))
	}
	if buckets[0].Name != "Criteria from prompt file" {
		t.Errorf("expected first bucket 'Criteria from prompt file', got %q", buckets[0].Name)
	}
	if buckets[1].Name != "combined" {
		t.Errorf("expected second bucket 'combined', got %q", buckets[1].Name)
	}
}

func TestBuildUnifiedReviewBuckets_OnlyPromptCriteria(t *testing.T) {
	// No matched criteria-file entries, only prompt frontmatter.
	buckets := BuildUnifiedReviewBuckets(nil, "prompt only", ReviewModeCombined)
	if len(buckets) != 1 {
		t.Fatalf("expected 1 bucket (prompt only), got %d", len(buckets))
	}
	if buckets[0].Name != "Criteria from prompt file" {
		t.Errorf("expected 'Criteria from prompt file', got %q", buckets[0].Name)
	}
	if !strings.Contains(buckets[0].Criteria, "prompt only") {
		t.Errorf("prompt bucket missing 'prompt only': %q", buckets[0].Criteria)
	}
}

func TestBuildUnifiedReviewBuckets_OnlyCriteriaFiles(t *testing.T) {
	// Matched criteria-file entries, no prompt frontmatter.
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "criterion A", false)},
		{Entry: promptEntry("b", "criterion B", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "", ReviewModeCombined)
	if len(buckets) != 1 {
		t.Fatalf("expected 1 bucket (criteria files only), got %d", len(buckets))
	}
	if buckets[0].Name != "combined" {
		t.Errorf("expected 'combined', got %q", buckets[0].Name)
	}
	for _, needle := range []string{"criterion A", "criterion B"} {
		if !strings.Contains(buckets[0].Criteria, needle) {
			t.Errorf("combined bucket missing %q: %q", needle, buckets[0].Criteria)
		}
	}
}

func TestBuildUnifiedReviewBuckets_EmptyInputs(t *testing.T) {
	// No matched entries, no prompt criteria → empty slice
	buckets := BuildUnifiedReviewBuckets(nil, "", ReviewModeCombined)
	if len(buckets) != 0 {
		t.Fatalf("expected 0 buckets for empty inputs, got %d", len(buckets))
	}
}

func TestMergeUnifiedCriteria_ContainsAllAndPromptCriteria(t *testing.T) {
	prompts := []UnifiedGraderEntry{
		promptEntry("a", "AA", false),
		promptEntry("b", "BB", false),
	}
	merged := MergeUnifiedCriteria(prompts, "PROMPT")
	for _, needle := range []string{"AA", "BB", "PROMPT"} {
		if !strings.Contains(merged, needle) {
			t.Errorf("merged missing %q: %q", needle, merged)
		}
	}
}

func TestUnifiedGraderEntry_ToRuntimeConfig_CopiesFields(t *testing.T) {
	e := UnifiedGraderEntry{
		Type:   graders.KindFile,
		Name:   "n",
		Weight: 0.5,
	}
	rc := e.ToRuntimeConfig()
	if rc.Kind != graders.KindFile || rc.Name != "n" || rc.Weight != 0.5 {
		t.Errorf("runtime config unexpected: %+v", rc)
	}
}
