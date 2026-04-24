package artifact

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestGeneratorArtifact_RoundTrip verifies that an artifact can be marshaled
// to JSON, written to disk, read back, and deep-equal to the original.
func TestGeneratorArtifact_RoundTrip(t *testing.T) {
	original := &GeneratorArtifact{
		PromptID:       "test-prompt-id",
		EvalID:         "test-eval-id",
		ConfigName:     "test-config",
		GeneratorModel: "claude-opus-4.6",
		OriginalPrompt: "Write a Python script that uses DefaultAzureCredential",
		FinalResponse:  "Here is the implementation...",
		WorkspaceDelta: ArtifactWorkspaceDelta{
			BytesAdded:        1234,
			BytesRemoved:      567,
			BytesNet:          667,
			NewFileCount:      2,
			ModifiedFileCount: 1,
			DeletedFileCount:  0,
			CreatedFiles: []ArtifactFileInfo{
				{Path: "main.py", Size: 800},
				{Path: "utils.py", Size: 434},
			},
			ModifiedFiles: []ArtifactFileInfo{
				{Path: "README.md", Size: 150},
			},
			DeletedFiles: nil,
		},
		ActionsSummary: ActionsSummary{
			TotalActions:   10,
			ToolCalls:      8,
			ReasoningSteps: 2,
			Truncated:      false,
		},
		StartedAt:    time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC),
		EndedAt:      time.Date(2025, 1, 15, 10, 5, 30, 0, time.UTC),
		DurationMs:   330000,
		TerminatedBy: "completed",
		Error:        "",
	}

	// Write to disk
	tmpDir := t.TempDir()
	artifactPath := filepath.Join(tmpDir, "generator.json")
	if err := original.WriteToFile(artifactPath); err != nil {
		t.Fatalf("WriteToFile() error: %v", err)
	}

	// Read back
	loaded, err := LoadGeneratorArtifact(artifactPath)
	if err != nil {
		t.Fatalf("LoadGeneratorArtifact() error: %v", err)
	}

	// Deep comparison via JSON marshaling
	origJSON, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling original: %v", err)
	}
	loadedJSON, err := json.Marshal(loaded)
	if err != nil {
		t.Fatalf("marshaling loaded: %v", err)
	}

	if string(origJSON) != string(loadedJSON) {
		t.Errorf("round-trip mismatch.\nOriginal: %s\nLoaded:   %s", string(origJSON), string(loadedJSON))
	}

	// Spot-check key fields
	if loaded.PromptID != original.PromptID {
		t.Errorf("PromptID = %q, want %q", loaded.PromptID, original.PromptID)
	}
	if loaded.WorkspaceDelta.BytesAdded != original.WorkspaceDelta.BytesAdded {
		t.Errorf("BytesAdded = %d, want %d", loaded.WorkspaceDelta.BytesAdded, original.WorkspaceDelta.BytesAdded)
	}
	if loaded.ActionsSummary.TotalActions != original.ActionsSummary.TotalActions {
		t.Errorf("TotalActions = %d, want %d", loaded.ActionsSummary.TotalActions, original.ActionsSummary.TotalActions)
	}
	if len(loaded.WorkspaceDelta.CreatedFiles) != len(original.WorkspaceDelta.CreatedFiles) {
		t.Errorf("CreatedFiles count = %d, want %d", len(loaded.WorkspaceDelta.CreatedFiles), len(original.WorkspaceDelta.CreatedFiles))
	}
}

// TestGeneratorArtifact_Truncation verifies that OriginalPrompt and FinalResponse
// fields are truncated to ~16KB with a [truncated] marker when they exceed the limit.
func TestGeneratorArtifact_Truncation(t *testing.T) {
	const maxSize = 16384 // 16KB

	// Create strings larger than maxSize
	longPrompt := strings.Repeat("A", maxSize+1000)
	longResponse := strings.Repeat("B", maxSize+2000)

	truncatedPrompt := TruncateField(longPrompt, maxSize)
	truncatedResponse := TruncateField(longResponse, maxSize)

	// Verify truncation occurred
	if !strings.HasSuffix(truncatedPrompt, "[truncated]") {
		t.Error("expected truncatedPrompt to have [truncated] marker")
	}
	if !strings.HasSuffix(truncatedResponse, "[truncated]") {
		t.Error("expected truncatedResponse to have [truncated] marker")
	}

	// Verify length is approximately maxSize + len("[truncated]")
	expectedMaxLen := maxSize + len("\n[truncated]")
	if len(truncatedPrompt) > expectedMaxLen+10 {
		t.Errorf("truncatedPrompt length = %d, expected ~%d", len(truncatedPrompt), expectedMaxLen)
	}
	if len(truncatedResponse) > expectedMaxLen+10 {
		t.Errorf("truncatedResponse length = %d, expected ~%d", len(truncatedResponse), expectedMaxLen)
	}

	// Verify short strings are not truncated
	shortStr := "short string"
	notTruncated := TruncateField(shortStr, maxSize)
	if notTruncated != shortStr {
		t.Errorf("TruncateField should not modify strings under maxSize")
	}
	if strings.Contains(notTruncated, "[truncated]") {
		t.Error("short string should not have [truncated] marker")
	}
}

// TestGeneratorArtifact_WriteToFile_CreatesDirectory verifies that
// WriteToFile creates parent directories if they don't exist.
func TestGeneratorArtifact_WriteToFile_CreatesDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	nestedPath := filepath.Join(tmpDir, "subdir1", "subdir2", "artifact.json")

	artifact := &GeneratorArtifact{
		PromptID:       "test",
		OriginalPrompt: "test prompt",
	}

	if err := artifact.WriteToFile(nestedPath); err != nil {
		t.Fatalf("WriteToFile() error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(nestedPath); err != nil {
		t.Errorf("artifact file not created at %s: %v", nestedPath, err)
	}
}

// TestLoadGeneratorArtifact_FileNotExists verifies error handling when
// the artifact file doesn't exist.
func TestLoadGeneratorArtifact_FileNotExists(t *testing.T) {
	_, err := LoadGeneratorArtifact("/nonexistent/path/artifact.json")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
	if !strings.Contains(err.Error(), "reading generator artifact") {
		t.Errorf("error = %q, want 'reading generator artifact'", err.Error())
	}
}

// TestLoadGeneratorArtifact_InvalidJSON verifies error handling when
// the artifact file contains invalid JSON.
func TestLoadGeneratorArtifact_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	badPath := filepath.Join(tmpDir, "bad.json")
	if err := os.WriteFile(badPath, []byte("{invalid json"), 0644); err != nil {
		t.Fatalf("writing bad JSON: %v", err)
	}

	_, err := LoadGeneratorArtifact(badPath)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "unmarshaling generator artifact") {
		t.Errorf("error = %q, want 'unmarshaling generator artifact'", err.Error())
	}
}

// TestFromWorkspaceDelta verifies conversion from workspace.WorkspaceDelta
// to ArtifactWorkspaceDelta.
func TestFromWorkspaceDelta(t *testing.T) {
	// This is a basic smoke test since workspace package types aren't imported here
	// In real usage, this would be tested as part of the engine integration
	nilDelta := FromWorkspaceDelta(nil)
	if nilDelta.BytesAdded != 0 || len(nilDelta.CreatedFiles) != 0 {
		t.Errorf("FromWorkspaceDelta(nil) should return zero-value struct")
	}
}
