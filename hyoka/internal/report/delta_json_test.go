package report

import (
	"encoding/json"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// TestEvalReport_WorkspaceDelta_JSONSerialization verifies workspace_delta
// field serializes correctly in EvalReport (Scenario 2.1).
func TestEvalReport_WorkspaceDelta_JSONSerialization(t *testing.T) {
	report := &EvalReport{
		SchemaVersion: CurrentSchemaVersion,
		PromptID:      "test-prompt",
		ConfigName:    "test-config",
		Success:       true,
		WorkspaceDelta: &workspace.WorkspaceDelta{
			BytesAdded:        500,
			BytesRemoved:      100,
			BytesNet:          400,
			NewFileCount:      3,
			ModifiedFileCount: 1,
			DeletedFileCount:  1,
			NewFiles: []workspace.NewFile{
				{Path: "main.py", Size: 200, Hash: "abc123"},
			},
			ModifiedFiles: []workspace.ModifiedFile{
				{Path: "lib.py", SizeBefore: 50, SizeAfter: 100, HashAfter: "def456"},
			},
			DeletedFiles: []workspace.DeletedFile{
				{Path: "old.py", OriginalSize: 100},
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	// Verify workspace_delta field is present
	jsonStr := string(jsonData)
	if !contains(jsonStr, "workspace_delta") {
		t.Errorf("workspace_delta field missing from JSON: %s", jsonStr)
	}
	if !contains(jsonStr, "bytes_added") {
		t.Errorf("bytes_added field missing from workspace_delta: %s", jsonStr)
	}
	if !contains(jsonStr, "new_files") {
		t.Errorf("new_files field missing from workspace_delta: %s", jsonStr)
	}

	// Unmarshal back
	var decoded EvalReport
	if err := json.Unmarshal(jsonData, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	// Verify delta decoded correctly
	if decoded.WorkspaceDelta == nil {
		t.Fatal("WorkspaceDelta nil after unmarshal")
	}
	if decoded.WorkspaceDelta.BytesAdded != 500 {
		t.Errorf("BytesAdded: expected 500, got %d", decoded.WorkspaceDelta.BytesAdded)
	}
	if decoded.WorkspaceDelta.NewFileCount != 3 {
		t.Errorf("NewFileCount: expected 3, got %d", decoded.WorkspaceDelta.NewFileCount)
	}
}

// TestEvalReport_MissingWorkspaceDelta_BackwardCompat verifies reports
// from older hyoka versions (no workspace_delta) decode without error (Scenario 2.2).
func TestEvalReport_MissingWorkspaceDelta_BackwardCompat(t *testing.T) {
	// JSON from older hyoka version (pre-#566)
	oldJSON := `{
		"schema_version": 2,
		"prompt_id": "test-prompt",
		"config_name": "test-config",
		"timestamp": "2024-01-01T00:00:00Z",
		"duration_seconds": 10.5,
		"prompt_metadata": {},
		"config_used": {},
		"generated_files": [],
		"event_count": 0,
		"tool_calls": [],
		"success": true
	}`

	var report EvalReport
	if err := json.Unmarshal([]byte(oldJSON), &report); err != nil {
		t.Fatalf("json.Unmarshal failed on old JSON: %v", err)
	}

	// Verify no error and WorkspaceDelta is nil (not panic)
	if report.WorkspaceDelta != nil {
		t.Errorf("expected nil WorkspaceDelta for old JSON, got %+v", report.WorkspaceDelta)
	}

	// Verify other fields decoded correctly
	if report.PromptID != "test-prompt" {
		t.Errorf("PromptID: expected test-prompt, got %s", report.PromptID)
	}
	if !report.Success {
		t.Errorf("Success: expected true, got false")
	}
}

// TestEvalReport_WorkspaceDelta_ZeroValues verifies zero-value delta
// serializes all fields (not omitted) (Scenario 2.3).
func TestEvalReport_WorkspaceDelta_ZeroValues(t *testing.T) {
	report := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "test-prompt",
		ConfigName:     "test-config",
		Success:        true,
		WorkspaceDelta: &workspace.WorkspaceDelta{}, // all zero values
	}

	jsonData, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	jsonStr := string(jsonData)

	// Verify zero values are present (not omitted)
	if !contains(jsonStr, `"bytes_added":0`) {
		t.Errorf("bytes_added missing or omitted: %s", jsonStr)
	}
	if !contains(jsonStr, `"bytes_net":0`) {
		t.Errorf("bytes_net missing or omitted: %s", jsonStr)
	}
	if !contains(jsonStr, `"new_file_count":0`) {
		t.Errorf("new_file_count missing or omitted: %s", jsonStr)
	}
}

// TestEvalReport_WorkspaceDelta_Nil verifies nil delta omits field entirely
// (saves bytes in JSON when delta not computed).
func TestEvalReport_WorkspaceDelta_Nil(t *testing.T) {
	report := &EvalReport{
		SchemaVersion:  CurrentSchemaVersion,
		PromptID:       "test-prompt",
		ConfigName:     "test-config",
		Success:        true,
		WorkspaceDelta: nil, // explicitly nil
	}

	jsonData, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	jsonStr := string(jsonData)

	// Verify workspace_delta field is omitted (omitempty behavior)
	if contains(jsonStr, "workspace_delta") {
		t.Errorf("workspace_delta should be omitted when nil, but found in JSON: %s", jsonStr)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
