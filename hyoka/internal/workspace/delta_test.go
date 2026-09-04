package workspace

import (
"os"
"path/filepath"
"testing"
)

func TestTakeSnapshot(t *testing.T) {
dir := t.TempDir()

// Create test files
writeFile(t, filepath.Join(dir, "a.txt"), "content A")
writeFile(t, filepath.Join(dir, "b.txt"), "content B")
os.Mkdir(filepath.Join(dir, "subdir"), 0755)
writeFile(t, filepath.Join(dir, "subdir", "c.txt"), "content C")

snap, err := TakeSnapshot(dir)
if err != nil {
t.Fatalf("TakeSnapshot failed: %v", err)
}

if len(snap.files) != 3 {
t.Errorf("expected 3 files, got %d", len(snap.files))
}

// Verify files are present
for _, path := range []string{"a.txt", "b.txt", filepath.Join("subdir", "c.txt")} {
if _, ok := snap.files[path]; !ok {
t.Errorf("missing file %q in snapshot", path)
}
}
}

func TestComputeDelta_NewFiles(t *testing.T) {
before := &Snapshot{files: make(map[string]fileInfo)}
after := &Snapshot{files: map[string]fileInfo{
"new.txt": {size: 100, hash: "abc123"},
}}

delta := ComputeDelta(before, after)

if delta.NewFileCount != 1 {
t.Errorf("expected 1 new file, got %d", delta.NewFileCount)
}
if delta.BytesAdded != 100 {
t.Errorf("expected 100 bytes added, got %d", delta.BytesAdded)
}
if delta.BytesNet != 100 {
t.Errorf("expected net 100 bytes, got %d", delta.BytesNet)
}
if len(delta.NewFiles) != 1 || delta.NewFiles[0].Path != "new.txt" {
t.Errorf("NewFiles incorrect: %+v", delta.NewFiles)
}
}

func TestComputeDelta_DeletedFiles(t *testing.T) {
before := &Snapshot{files: map[string]fileInfo{
"deleted.txt": {size: 50, hash: "def456"},
}}
after := &Snapshot{files: make(map[string]fileInfo)}

delta := ComputeDelta(before, after)

if delta.DeletedFileCount != 1 {
t.Errorf("expected 1 deleted file, got %d", delta.DeletedFileCount)
}
if delta.BytesRemoved != 50 {
t.Errorf("expected 50 bytes removed, got %d", delta.BytesRemoved)
}
if delta.BytesNet != -50 {
t.Errorf("expected net -50 bytes, got %d", delta.BytesNet)
}
if len(delta.DeletedFiles) != 1 || delta.DeletedFiles[0].Path != "deleted.txt" {
t.Errorf("DeletedFiles incorrect: %+v", delta.DeletedFiles)
}
}

func TestComputeDelta_ModifiedFiles(t *testing.T) {
before := &Snapshot{files: map[string]fileInfo{
"modified.txt": {size: 100, hash: "hash1"},
}}
after := &Snapshot{files: map[string]fileInfo{
"modified.txt": {size: 150, hash: "hash2"}, // different hash, grew by 50
}}

delta := ComputeDelta(before, after)

if delta.ModifiedFileCount != 1 {
t.Errorf("expected 1 modified file, got %d", delta.ModifiedFileCount)
}
if delta.BytesAdded != 50 {
t.Errorf("expected 50 bytes added, got %d", delta.BytesAdded)
}
if delta.BytesNet != 50 {
t.Errorf("expected net 50 bytes, got %d", delta.BytesNet)
}
if len(delta.ModifiedFiles) != 1 {
t.Errorf("ModifiedFiles incorrect: %+v", delta.ModifiedFiles)
}
mf := delta.ModifiedFiles[0]
if mf.Path != "modified.txt" || mf.SizeBefore != 100 || mf.SizeAfter != 150 {
t.Errorf("ModifiedFile data incorrect: %+v", mf)
}
}

func TestComputeDelta_ModifiedFileShrink(t *testing.T) {
before := &Snapshot{files: map[string]fileInfo{
"shrink.txt": {size: 200, hash: "hash1"},
}}
after := &Snapshot{files: map[string]fileInfo{
"shrink.txt": {size: 50, hash: "hash2"}, // shrunk by 150
}}

delta := ComputeDelta(before, after)

if delta.ModifiedFileCount != 1 {
t.Errorf("expected 1 modified file, got %d", delta.ModifiedFileCount)
}
if delta.BytesRemoved != 150 {
t.Errorf("expected 150 bytes removed, got %d", delta.BytesRemoved)
}
if delta.BytesNet != -150 {
t.Errorf("expected net -150 bytes, got %d", delta.BytesNet)
}
}

func TestComputeDelta_MixedChanges(t *testing.T) {
before := &Snapshot{files: map[string]fileInfo{
"kept.txt":     {size: 10, hash: "h1"},
"modified.txt": {size: 100, hash: "h2"},
"deleted.txt":  {size: 50, hash: "h3"},
}}
after := &Snapshot{files: map[string]fileInfo{
"kept.txt":     {size: 10, hash: "h1"}, // unchanged
"modified.txt": {size: 120, hash: "h4"}, // grew by 20
"new.txt":      {size: 30, hash: "h5"},
}}

delta := ComputeDelta(before, after)

if delta.NewFileCount != 1 {
t.Errorf("expected 1 new file, got %d", delta.NewFileCount)
}
if delta.ModifiedFileCount != 1 {
t.Errorf("expected 1 modified file, got %d", delta.ModifiedFileCount)
}
if delta.DeletedFileCount != 1 {
t.Errorf("expected 1 deleted file, got %d", delta.DeletedFileCount)
}

// BytesAdded: 30 (new) + 20 (modified growth) = 50
if delta.BytesAdded != 50 {
t.Errorf("expected 50 bytes added, got %d", delta.BytesAdded)
}
// BytesRemoved: 50 (deleted)
if delta.BytesRemoved != 50 {
t.Errorf("expected 50 bytes removed, got %d", delta.BytesRemoved)
}
// BytesNet: 50 - 50 = 0
if delta.BytesNet != 0 {
t.Errorf("expected net 0 bytes, got %d", delta.BytesNet)
}
}

func TestComputeDelta_NoChanges(t *testing.T) {
snap := &Snapshot{files: map[string]fileInfo{
"unchanged.txt": {size: 100, hash: "hash"},
}}

delta := ComputeDelta(snap, snap)

if delta.NewFileCount != 0 || delta.ModifiedFileCount != 0 || delta.DeletedFileCount != 0 {
t.Errorf("expected no changes, got: new=%d modified=%d deleted=%d",
delta.NewFileCount, delta.ModifiedFileCount, delta.DeletedFileCount)
}
if delta.BytesAdded != 0 || delta.BytesRemoved != 0 || delta.BytesNet != 0 {
t.Errorf("expected zero bytes, got: added=%d removed=%d net=%d",
delta.BytesAdded, delta.BytesRemoved, delta.BytesNet)
}
}

func writeFile(t *testing.T, path, content string) {
t.Helper()
if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
t.Fatalf("mkdir failed: %v", err)
}
if err := os.WriteFile(path, []byte(content), 0644); err != nil {
t.Fatalf("write file failed: %v", err)
}
}
