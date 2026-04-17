package graders

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// TestGraderInput_WorkspaceDeltaNil verifies graders handle nil delta gracefully (Scenario 3.2).
func TestGraderInput_WorkspaceDeltaNil(t *testing.T) {
	input := GraderInput{
		WorkspacePath:  "/tmp/workspace",
		WorkspaceDelta: nil, // Pre-#566 reports may have nil delta
		PromptMeta: PromptMetadata{
			ID:       "test-prompt",
			Service:  "test-service",
			Language: "python",
		},
	}

	// Grader must check for nil before accessing
	// This pattern demonstrates safe access:
	if input.WorkspaceDelta != nil {
		_ = input.WorkspaceDelta.NewFileCount
	}

	// Should not panic - test passes if we reach here
	if input.WorkspaceDelta != nil {
		t.Errorf("expected nil WorkspaceDelta, got %+v", input.WorkspaceDelta)
	}
}

// TestGraderInput_WorkspaceDeltaPresent verifies graders can access delta (Scenario 3.1).
func TestGraderInput_WorkspaceDeltaPresent(t *testing.T) {
	delta := &workspace.WorkspaceDelta{
		BytesAdded:   500,
		BytesRemoved: 100,
		BytesNet:     400,
		NewFileCount: 3,
		NewFiles: []workspace.NewFile{
			{Path: "main.py", Size: 200, Hash: "abc123"},
		},
	}

	input := GraderInput{
		WorkspacePath:  "/tmp/workspace",
		WorkspaceDelta: delta,
		PromptMeta: PromptMetadata{
			ID:       "test-prompt",
			Service:  "test-service",
			Language: "python",
		},
	}

	// Graders can safely access delta when present
	if input.WorkspaceDelta == nil {
		t.Fatal("expected non-nil WorkspaceDelta")
	}

	if input.WorkspaceDelta.NewFileCount != 3 {
		t.Errorf("NewFileCount: expected 3, got %d", input.WorkspaceDelta.NewFileCount)
	}
	if input.WorkspaceDelta.BytesNet != 400 {
		t.Errorf("BytesNet: expected 400, got %d", input.WorkspaceDelta.BytesNet)
	}
	if len(input.WorkspaceDelta.NewFiles) != 1 {
		t.Errorf("NewFiles: expected 1, got %d", len(input.WorkspaceDelta.NewFiles))
	}
}

// mockGrader is a test grader that safely accesses WorkspaceDelta.
type mockGrader struct{}

func (m *mockGrader) Kind() string { return "mock" }
func (m *mockGrader) Name() string { return "Mock Grader" }

func (m *mockGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
	result := GraderResult{
		Kind:    "mock",
		Name:    "Mock Grader",
		Score:   1.0,
		Weight:  1.0,
		Pass:    true,
		Message: "Mock grader passed",
	}

	// CRITICAL: Graders must nil-check WorkspaceDelta before accessing
	if input.WorkspaceDelta != nil {
		// Safe to access delta fields
		result.Message = "Mock grader with delta: " + string(rune(input.WorkspaceDelta.NewFileCount))
	}

	return result, nil
}

// TestMockGrader_NilDelta verifies mock grader handles nil delta.
func TestMockGrader_NilDelta(t *testing.T) {
	grader := &mockGrader{}

	input := GraderInput{
		WorkspacePath:  "/tmp/workspace",
		WorkspaceDelta: nil,
		PromptMeta: PromptMetadata{
			ID:       "test-prompt",
			Language: "python",
		},
	}

	result, err := grader.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade failed: %v", err)
	}

	if !result.Pass {
		t.Errorf("expected Pass=true, got false")
	}
	// Should not panic - test passes if we reach here
}

// TestMockGrader_WithDelta verifies mock grader handles non-nil delta.
func TestMockGrader_WithDelta(t *testing.T) {
	grader := &mockGrader{}

	delta := &workspace.WorkspaceDelta{
		NewFileCount: 5,
	}

	input := GraderInput{
		WorkspacePath:  "/tmp/workspace",
		WorkspaceDelta: delta,
		PromptMeta: PromptMetadata{
			ID:       "test-prompt",
			Language: "python",
		},
	}

	result, err := grader.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade failed: %v", err)
	}

	if !result.Pass {
		t.Errorf("expected Pass=true, got false")
	}
	// Should not panic - test passes if we reach here
}
