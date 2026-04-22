package graders

import (
	"strings"
	"testing"
)

func promptEntry(name, crit string, isolate bool) UnifiedGraderEntry {
	return UnifiedGraderEntry{Type: KindPrompt, Name: name, Prompt: crit, Isolate: isolate}
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
				{Type: KindPrompt, Name: "c", Prompt: "C", When: map[string]string{"category": "crud"}},
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
		{Entry: typedEntry(KindFile, "f1")},
		{Entry: promptEntry("b", "B", true)},
		{Entry: typedEntry(KindOutputCheck, "oc1")},
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
	if len(buckets) != 1 {
		t.Fatalf("combined: expected 1 bucket, got %d", len(buckets))
	}
	for _, needle := range []string{"criterion A", "criterion B", "prompt criteria"} {
		if !strings.Contains(buckets[0].Criteria, needle) {
			t.Errorf("combined bucket missing %q: %q", needle, buckets[0].Criteria)
		}
	}
}

func TestBuildUnifiedReviewBuckets_IsolatedWithIsolation(t *testing.T) {
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("security", "no secrets", true)},
		{Entry: promptEntry("format", "format ok", false)},
		{Entry: promptEntry("tests", "tests exist", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "pc", ReviewModeIsolated)
	if len(buckets) != 2 {
		t.Fatalf("isolated: expected 2 buckets, got %d", len(buckets))
	}
	var sawSec, sawCombined bool
	for _, b := range buckets {
		switch b.Name {
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
		}
	}
	if !sawSec || !sawCombined {
		t.Errorf("expected both 'security' and 'combined' buckets; sawSec=%v sawCombined=%v", sawSec, sawCombined)
	}
}

func TestBuildUnifiedReviewBuckets_IsolatedDegradesToCombined(t *testing.T) {
	prompts := []MatchedUnifiedEntry{
		{Entry: promptEntry("a", "A", false)},
		{Entry: promptEntry("b", "B", false)},
	}
	buckets := BuildUnifiedReviewBuckets(prompts, "pc", ReviewModeIsolated)
	if len(buckets) != 1 {
		t.Fatalf("degraded: expected 1 bucket, got %d", len(buckets))
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
		Type:   KindFile,
		Name:   "n",
		Weight: 0.5,
	}
	rc := e.ToRuntimeConfig()
	if rc.Kind != KindFile || rc.Name != "n" || rc.Weight != 0.5 {
		t.Errorf("runtime config unexpected: %+v", rc)
	}
}
