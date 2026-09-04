package review

import (
	"encoding/json"
	"testing"
)

// TestReviewResult_SkippedReviewersField verifies that the SkippedReviewers
// field exists in ReviewResult and marshals correctly with the omitempty tag.
// This is a struct-level test for PR #640 Fix 2.
// Once Neo implements the behavior in ReviewPanel/ReviewPanelBuckets, this
// test will pass AND the integration tests can be added.
func TestReviewResult_SkippedReviewersField(t *testing.T) {
	// Test 1: SkippedReviewers populated
	result := &ReviewResult{
		Model:        "test-model",
		OverallScore: 5,
		MaxScore:     10,
		Summary:      "Some checks passed",
		SkippedReviewers: []SkippedReviewer{
			{Model: "failed-model-1", Error: "timeout"},
			{Model: "failed-model-2", Error: "API error"},
		},
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("failed to marshal ReviewResult with SkippedReviewers: %v", err)
	}

	jsonStr := string(jsonData)

	// Verify the field is present in JSON
	if !contains(jsonStr, "skipped_reviewers") {
		t.Errorf("JSON should contain 'skipped_reviewers' field: %s", jsonStr)
	}

	if !contains(jsonStr, "failed-model-1") {
		t.Errorf("JSON should contain first skipped model: %s", jsonStr)
	}

	if !contains(jsonStr, "timeout") {
		t.Errorf("JSON should contain error message: %s", jsonStr)
	}

	// Unmarshal and verify round-trip
	var decoded ReviewResult
	if err := json.Unmarshal(jsonData, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(decoded.SkippedReviewers) != 2 {
		t.Errorf("expected 2 skipped reviewers after round-trip, got %d", len(decoded.SkippedReviewers))
	}

	if decoded.SkippedReviewers[0].Model != "failed-model-1" {
		t.Errorf("model name not preserved: got %q", decoded.SkippedReviewers[0].Model)
	}

	// Test 2: SkippedReviewers empty/nil with omitempty
	resultEmpty := &ReviewResult{
		Model:        "test-model",
		OverallScore: 10,
		MaxScore:     10,
		Summary:      "All passed",
		// SkippedReviewers not set (nil)
	}

	jsonEmpty, err := json.Marshal(resultEmpty)
	if err != nil {
		t.Fatalf("failed to marshal ReviewResult without SkippedReviewers: %v", err)
	}

	jsonEmptyStr := string(jsonEmpty)

	// With omitempty, the field should be omitted entirely
	if contains(jsonEmptyStr, "skipped_reviewers") {
		t.Errorf("JSON should NOT contain 'skipped_reviewers' when nil/empty (omitempty): %s", jsonEmptyStr)
	}

	// Test 3: SkippedReviewers explicitly empty slice
	resultEmptySlice := &ReviewResult{
		Model:            "test-model",
		OverallScore:     10,
		MaxScore:         10,
		Summary:          "All passed",
		SkippedReviewers: []SkippedReviewer{}, // explicitly empty
	}

	jsonEmptySlice, err := json.Marshal(resultEmptySlice)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	jsonEmptySliceStr := string(jsonEmptySlice)

	// Empty slice with omitempty should also be omitted
	if contains(jsonEmptySliceStr, "skipped_reviewers") {
		t.Errorf("JSON should NOT contain 'skipped_reviewers' when empty slice (omitempty): %s", jsonEmptySliceStr)
	}
}

// TODO: Once Neo implements ReviewPanel/ReviewPanelBuckets logic to populate
// SkippedReviewers, add integration tests here that:
//   1. Create a PanelReviewer with multiple models where one fails
//   2. Call ReviewPanel and verify SkippedReviewers is populated
//   3. Verify the error message is preserved
//   4. Test ReviewPanelBuckets with "all buckets failed" scenario
//
// For now, this struct-level test verifies the field exists and marshals correctly.

// contains checks if a string contains a substring.
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
