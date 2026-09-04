package workspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestComputeDelta_ZeroByteFiles verifies zero-byte files are tracked (Scenario 1.8).
func TestComputeDelta_ZeroByteFiles(t *testing.T) {
	before := &Snapshot{files: map[string]fileInfo{
		"__init__.py": {size: 0, hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}, // SHA-256 of empty
	}}
	after := &Snapshot{files: map[string]fileInfo{
		"__init__.py": {size: 0, hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}, // unchanged
		"empty.txt":   {size: 0, hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}}

	delta := ComputeDelta(before, after)

	if delta.NewFileCount != 1 {
		t.Errorf("expected 1 new file, got %d", delta.NewFileCount)
	}
	if len(delta.NewFiles) != 1 || delta.NewFiles[0].Path != "empty.txt" {
		t.Errorf("NewFiles incorrect: %+v", delta.NewFiles)
	}
	if delta.NewFiles[0].Size != 0 {
		t.Errorf("expected size 0, got %d", delta.NewFiles[0].Size)
	}
	if delta.BytesAdded != 0 || delta.BytesRemoved != 0 || delta.BytesNet != 0 {
		t.Errorf("expected zero bytes, got: added=%d removed=%d net=%d",
			delta.BytesAdded, delta.BytesRemoved, delta.BytesNet)
	}
	if delta.ModifiedFileCount != 0 {
		t.Errorf("unchanged __init__.py should not appear in ModifiedFiles")
	}
}

// TestComputeDelta_EmptyWorkspaceAfterAgent verifies all files deleted (Scenario 1.9).
func TestComputeDelta_EmptyWorkspaceAfterAgent(t *testing.T) {
	before := &Snapshot{files: map[string]fileInfo{
		"main.py": {size: 100, hash: "hash1"},
		"lib.py":  {size: 50, hash: "hash2"},
	}}
	after := &Snapshot{files: make(map[string]fileInfo)} // empty

	delta := ComputeDelta(before, after)

	if delta.DeletedFileCount != 2 {
		t.Errorf("expected 2 deleted files, got %d", delta.DeletedFileCount)
	}
	if delta.BytesRemoved != 150 {
		t.Errorf("expected 150 bytes removed, got %d", delta.BytesRemoved)
	}
	if delta.BytesNet != -150 {
		t.Errorf("expected net -150 bytes, got %d", delta.BytesNet)
	}
	if len(delta.DeletedFiles) != 2 {
		t.Errorf("DeletedFiles incorrect: %+v", delta.DeletedFiles)
	}
}

// TestComputeDelta_CreateAndDeleteSameFile verifies no net change when
// agent creates and deletes the same file in one session (Scenario 1.6).
func TestComputeDelta_CreateAndDeleteSameFile(t *testing.T) {
	// Both before and after are empty — agent created temp.py then deleted it
	empty := &Snapshot{files: make(map[string]fileInfo)}

	delta := ComputeDelta(empty, empty)

	if delta.NewFileCount != 0 || delta.ModifiedFileCount != 0 || delta.DeletedFileCount != 0 {
		t.Errorf("expected no changes (zero net), got: new=%d modified=%d deleted=%d",
			delta.NewFileCount, delta.ModifiedFileCount, delta.DeletedFileCount)
	}
	if delta.BytesAdded != 0 || delta.BytesRemoved != 0 || delta.BytesNet != 0 {
		t.Errorf("expected zero bytes, got: added=%d removed=%d net=%d",
			delta.BytesAdded, delta.BytesRemoved, delta.BytesNet)
	}
}

// TestComputeDelta_UnicodeFilenames verifies non-ASCII paths work (Scenario 5.5).
func TestComputeDelta_UnicodeFilenames(t *testing.T) {
	before := &Snapshot{files: make(map[string]fileInfo)}
	after := &Snapshot{files: map[string]fileInfo{
		"文档.txt":    {size: 10, hash: "hash1"},
		"résumé.md": {size: 20, hash: "hash2"},
	}}

	delta := ComputeDelta(before, after)

	if delta.NewFileCount != 2 {
		t.Errorf("expected 2 new files, got %d", delta.NewFileCount)
	}
	if delta.BytesAdded != 30 {
		t.Errorf("expected 30 bytes added, got %d", delta.BytesAdded)
	}

	// Verify paths are preserved correctly
	paths := make(map[string]bool)
	for _, f := range delta.NewFiles {
		paths[f.Path] = true
	}
	if !paths["文档.txt"] || !paths["résumé.md"] {
		t.Errorf("unicode paths not preserved correctly: %+v", delta.NewFiles)
	}
}

// TestComputeDelta_BinaryFiles verifies binary files tracked by size (Scenario 5.1).
func TestComputeDelta_BinaryFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a binary file
	binPath := filepath.Join(dir, "image.png")
	binData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A} // PNG header
	if err := os.WriteFile(binPath, binData, 0644); err != nil {
		t.Fatalf("failed to write binary file: %v", err)
	}

	before, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	// Remove file from "before" to simulate creation
	before.files = make(map[string]fileInfo)

	after, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	delta := ComputeDelta(before, after)

	if delta.NewFileCount != 1 {
		t.Errorf("expected 1 new file, got %d", delta.NewFileCount)
	}
	if delta.NewFiles[0].Size != int64(len(binData)) {
		t.Errorf("expected size %d, got %d", len(binData), delta.NewFiles[0].Size)
	}
	// Hash should be computed (no panic on binary)
	if delta.NewFiles[0].Hash == "" {
		t.Errorf("expected hash computed for binary file")
	}
}

// TestWorkspaceDelta_JSONRoundtrip verifies JSON serialization (Scenario 2.1).
func TestWorkspaceDelta_JSONRoundtrip(t *testing.T) {
	original := &WorkspaceDelta{
		BytesAdded:        350,
		BytesRemoved:      50,
		BytesNet:          300,
		NewFileCount:      2,
		ModifiedFileCount: 1,
		DeletedFileCount:  1,
		NewFiles: []NewFile{
			{Path: "main.py", Size: 100, Hash: "abc123"},
			{Path: "lib.py", Size: 200, Hash: "def456"},
		},
		ModifiedFiles: []ModifiedFile{
			{Path: "config.json", SizeBefore: 50, SizeAfter: 100, HashAfter: "ghi789"},
		},
		DeletedFiles: []DeletedFile{
			{Path: "old.py", OriginalSize: 50},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	// Verify field names (snake_case)
	jsonStr := string(jsonData)
	if !contains(jsonStr, "bytes_added") || !contains(jsonStr, "bytes_removed") || !contains(jsonStr, "bytes_net") {
		t.Errorf("JSON missing expected snake_case fields: %s", jsonStr)
	}

	// Unmarshal back
	var decoded WorkspaceDelta
	if err := json.Unmarshal(jsonData, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	// Verify round-trip equality
	if decoded.BytesAdded != original.BytesAdded {
		t.Errorf("BytesAdded: expected %d, got %d", original.BytesAdded, decoded.BytesAdded)
	}
	if decoded.BytesRemoved != original.BytesRemoved {
		t.Errorf("BytesRemoved: expected %d, got %d", original.BytesRemoved, decoded.BytesRemoved)
	}
	if decoded.BytesNet != original.BytesNet {
		t.Errorf("BytesNet: expected %d, got %d", original.BytesNet, decoded.BytesNet)
	}
	if decoded.NewFileCount != original.NewFileCount {
		t.Errorf("NewFileCount: expected %d, got %d", original.NewFileCount, decoded.NewFileCount)
	}
	if len(decoded.NewFiles) != len(original.NewFiles) {
		t.Errorf("NewFiles length: expected %d, got %d", len(original.NewFiles), len(decoded.NewFiles))
	}
	if len(decoded.ModifiedFiles) != len(original.ModifiedFiles) {
		t.Errorf("ModifiedFiles length: expected %d, got %d", len(original.ModifiedFiles), len(decoded.ModifiedFiles))
	}
	if len(decoded.DeletedFiles) != len(original.DeletedFiles) {
		t.Errorf("DeletedFiles length: expected %d, got %d", len(original.DeletedFiles), len(decoded.DeletedFiles))
	}
}

// TestWorkspaceDelta_JSONWithZeroValues verifies zero values serialize (Scenario 2.3).
func TestWorkspaceDelta_JSONWithZeroValues(t *testing.T) {
	delta := &WorkspaceDelta{} // all zero values

	jsonData, err := json.Marshal(delta)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	jsonStr := string(jsonData)

	// Verify zero values are present (not omitted)
	if !contains(jsonStr, `"bytes_added":0`) {
		t.Errorf("bytes_added missing or omitted: %s", jsonStr)
	}
	if !contains(jsonStr, `"bytes_removed":0`) {
		t.Errorf("bytes_removed missing or omitted: %s", jsonStr)
	}
	if !contains(jsonStr, `"bytes_net":0`) {
		t.Errorf("bytes_net missing or omitted: %s", jsonStr)
	}
	if !contains(jsonStr, `"new_file_count":0`) {
		t.Errorf("new_file_count missing or omitted: %s", jsonStr)
	}
}

// TestComputeDelta_PermissionChangeOnly verifies permission-only changes
// are NOT tracked as modifications (Scenario 5.6).
func TestComputeDelta_PermissionChangeOnly(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "script.sh")

	// Create file with initial permission
	content := "#!/bin/bash\necho hello\n"
	if err := os.WriteFile(filePath, []byte(content), 0755); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	before, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	// Change permission (but not content)
	if err := os.Chmod(filePath, 0644); err != nil {
		t.Fatalf("chmod failed: %v", err)
	}

	after, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	delta := ComputeDelta(before, after)

	// Permission change should NOT appear as modification (content/size unchanged)
	if delta.ModifiedFileCount != 0 {
		t.Errorf("expected no modifications (permission-only change), got %d modified files", delta.ModifiedFileCount)
	}
	if delta.BytesAdded != 0 || delta.BytesRemoved != 0 || delta.BytesNet != 0 {
		t.Errorf("expected zero bytes changed, got: added=%d removed=%d net=%d",
			delta.BytesAdded, delta.BytesRemoved, delta.BytesNet)
	}
}

// TestTakeSnapshot_HiddenFilesExcluded verifies hidden files are skipped.
func TestTakeSnapshot_HiddenFilesExcluded(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, filepath.Join(dir, "visible.txt"), "visible")
	writeFile(t, filepath.Join(dir, ".hidden"), "hidden")
	os.Mkdir(filepath.Join(dir, ".hiddendir"), 0755)
	writeFile(t, filepath.Join(dir, ".hiddendir", "nested.txt"), "nested")

	snap, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	// Only visible.txt should be in snapshot
	if len(snap.files) != 1 {
		t.Errorf("expected 1 file, got %d: %+v", len(snap.files), snap.files)
	}
	if _, ok := snap.files["visible.txt"]; !ok {
		t.Errorf("visible.txt not in snapshot")
	}
	if _, ok := snap.files[".hidden"]; ok {
		t.Errorf(".hidden should be excluded from snapshot")
	}
}

// contains is a helper to check if a string contains a substring.
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
