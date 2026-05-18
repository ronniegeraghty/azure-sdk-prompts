package review

import (
	"strings"
	"testing"
)

// TestMergeBucketResults_NamePrefixing locks in the criterion-name
// prefixing rule that keeps per-bucket criteria distinguishable when
// the deterministic any-fail vote runs across panel models (#580).
//
// Rules under test:
//   - bucket name "" (empty) → no prefix
//   - bucket name "combined" → no prefix (legacy reserved name)
//   - any other name        → "[name] " prefix on each criterion + summary
func TestMergeBucketResults_NamePrefixing(t *testing.T) {
	tests := []struct {
		name           string
		bucketName     string
		criterionName  string
		summaryIn      string
		wantCriterion  string
		wantSummaryHas string
	}{
		{
			name:           "empty bucket name leaves criterion untouched",
			bucketName:     "",
			criterionName:  "no_secrets",
			summaryIn:      "looked clean",
			wantCriterion:  "no_secrets",
			wantSummaryHas: "looked clean",
		},
		{
			name:           "combined bucket leaves criterion untouched",
			bucketName:     "combined",
			criterionName:  "no_secrets",
			summaryIn:      "looked clean",
			wantCriterion:  "no_secrets",
			wantSummaryHas: "looked clean",
		},
		{
			name:           "named bucket prefixes criterion",
			bucketName:     "security",
			criterionName:  "no_secrets",
			summaryIn:      "looked clean",
			wantCriterion:  "[security] no_secrets",
			wantSummaryHas: "[security] looked clean",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parts := []bucketResult{{
				name: tt.bucketName,
				result: &ReviewResult{
					Scores: ReviewScores{Criteria: []CriterionResult{
						{Name: tt.criterionName, Passed: true, Reason: "ok"},
					}},
					OverallScore: 1,
					MaxScore:     1,
					Summary:      tt.summaryIn,
				},
			}}
			got := mergeBucketResults(parts)
			if len(got.Scores.Criteria) != 1 {
				t.Fatalf("want 1 criterion, got %d", len(got.Scores.Criteria))
			}
			if got.Scores.Criteria[0].Name != tt.wantCriterion {
				t.Errorf("criterion name = %q, want %q",
					got.Scores.Criteria[0].Name, tt.wantCriterion)
			}
			if !strings.Contains(got.Summary, tt.wantSummaryHas) {
				t.Errorf("summary = %q, want it to contain %q", got.Summary, tt.wantSummaryHas)
			}
		})
	}
}

// TestMergeBucketResults_AggregatesAcrossBuckets verifies that issues,
// strengths, scores, and events are concatenated/summed across buckets so
// the merged ReviewResult faithfully represents every bucket's verdict.
func TestMergeBucketResults_AggregatesAcrossBuckets(t *testing.T) {
	parts := []bucketResult{
		{
			name: "combined",
			result: &ReviewResult{
				Scores: ReviewScores{Criteria: []CriterionResult{
					{Name: "fmt", Passed: true},
				}},
				OverallScore: 1,
				MaxScore:     1,
				Summary:      "shared looked fine",
				Issues:       []string{"minor lint"},
				Strengths:    []string{"clear naming"},
				Events:       []ReviewEvent{{Type: "session.start"}},
			},
		},
		{
			name: "security",
			result: &ReviewResult{
				Scores: ReviewScores{Criteria: []CriterionResult{
					{Name: "no_secrets", Passed: false, Reason: "hardcoded key"},
				}},
				OverallScore: 0,
				MaxScore:     1,
				Summary:      "found a secret",
				Issues:       []string{"hardcoded API key"},
				Strengths:    []string{},
				Events:       []ReviewEvent{{Type: "tool.read"}, {Type: "tool.read"}},
			},
		},
	}

	got := mergeBucketResults(parts)

	if got.OverallScore != 1 || got.MaxScore != 2 {
		t.Errorf("score = %d/%d, want 1/2", got.OverallScore, got.MaxScore)
	}
	if len(got.Scores.Criteria) != 2 {
		t.Fatalf("want 2 merged criteria, got %d", len(got.Scores.Criteria))
	}
	// "combined" bucket leaves name as-is; named bucket prefixes it.
	wantNames := map[string]bool{"fmt": false, "[security] no_secrets": false}
	for _, c := range got.Scores.Criteria {
		if _, ok := wantNames[c.Name]; ok {
			wantNames[c.Name] = true
		} else {
			t.Errorf("unexpected merged criterion name %q", c.Name)
		}
	}
	for n, seen := range wantNames {
		if !seen {
			t.Errorf("missing merged criterion %q", n)
		}
	}
	if len(got.Issues) != 2 {
		t.Errorf("want 2 merged issues, got %d", len(got.Issues))
	}
	if len(got.Strengths) != 1 {
		t.Errorf("want 1 merged strength, got %d", len(got.Strengths))
	}
	if len(got.Events) != 3 {
		t.Errorf("want 3 merged events, got %d", len(got.Events))
	}
	if !strings.Contains(got.Summary, "shared looked fine") ||
		!strings.Contains(got.Summary, "[security] found a secret") {
		t.Errorf("summary missing expected fragments: %q", got.Summary)
	}
}

// TestMergeBucketResults_NilPartsAreSkipped ensures a nil result inside the
// parts slice does not panic and is silently skipped.
func TestMergeBucketResults_NilPartsAreSkipped(t *testing.T) {
	parts := []bucketResult{
		{name: "combined", result: nil},
		{name: "security", result: &ReviewResult{
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "ok", Passed: true},
			}},
			OverallScore: 1,
			MaxScore:     1,
			Summary:      "fine",
		}},
	}
	got := mergeBucketResults(parts)
	if got.OverallScore != 1 || got.MaxScore != 1 {
		t.Errorf("score = %d/%d, want 1/1", got.OverallScore, got.MaxScore)
	}
	if len(got.Scores.Criteria) != 1 {
		t.Fatalf("want 1 criterion, got %d", len(got.Scores.Criteria))
	}
	if got.Scores.Criteria[0].Name != "[security] ok" {
		t.Errorf("criterion name = %q, want [security] ok", got.Scores.Criteria[0].Name)
	}
}
