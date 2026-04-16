package eval

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile creates a file at baseDir/rel with the given contents, ensuring
// the parent directory exists.
func writeFile(t *testing.T, baseDir, rel string, contents []byte) {
	t.Helper()
	full := filepath.Join(baseDir, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", full, err)
	}
	if err := os.WriteFile(full, contents, 0o644); err != nil {
		t.Fatalf("write %s: %v", full, err)
	}
}

func TestSnapshotStarterSizes(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "main.go", []byte("package main\n"))
	writeFile(t, dir, "pkg/lib.go", []byte("package pkg\n"))

	snap := snapshotStarterSizes(dir, []string{"main.go", "pkg/lib.go", "missing.go"})

	if got, want := snap["main.go"], int64(len("package main\n")); got != want {
		t.Errorf("main.go size: got %d want %d", got, want)
	}
	if got, want := snap["pkg/lib.go"], int64(len("package pkg\n")); got != want {
		t.Errorf("pkg/lib.go size: got %d want %d", got, want)
	}
	// Missing files are recorded with size 0 (safe default — worst case the
	// agent's copy is fully charged).
	if got, ok := snap["missing.go"]; !ok || got != 0 {
		t.Errorf("missing.go: got (%d, ok=%v), want (0, ok=true)", got, ok)
	}
}

func TestComputeAgentOutputSize(t *testing.T) {
	tests := []struct {
		name     string
		starter  map[string][]byte // starter files written before snapshot
		after    map[string][]byte // files present at guardrail time (may overlap/differ)
		removed  []string          // paths to remove after snapshot (simulate deletions)
		expected int64
	}{
		{
			name:     "unchanged starter produces zero agent bytes",
			starter:  map[string][]byte{"a.go": []byte("hello"), "b.go": []byte("world!!")},
			after:    map[string][]byte{"a.go": []byte("hello"), "b.go": []byte("world!!")},
			expected: 0,
		},
		{
			name:     "starter modified in place counts only delta",
			starter:  map[string][]byte{"a.go": []byte("hello")},                 // 5
			after:    map[string][]byte{"a.go": []byte("hello, extended world")}, // 21
			expected: 21 - 5,
		},
		{
			name:     "new agent file counts full size",
			starter:  map[string][]byte{"a.go": []byte("hello")},
			after:    map[string][]byte{"a.go": []byte("hello"), "new.go": []byte("abcdef")},
			expected: 6,
		},
		{
			name:     "starter shrunk does not count negative bytes",
			starter:  map[string][]byte{"a.go": []byte("longer original contents")}, // 24
			after:    map[string][]byte{"a.go": []byte("tiny")},                     // 4
			expected: 0,
		},
		{
			name:     "deleted starter contributes no bytes",
			starter:  map[string][]byte{"a.go": []byte("hello"), "b.go": []byte("world!!")},
			after:    map[string][]byte{"a.go": []byte("hello")},
			removed:  []string{"b.go"},
			expected: 0,
		},
		{
			name:     "mixed: delta + new file",
			starter:  map[string][]byte{"a.go": []byte("hi")},                                     // 2
			after:    map[string][]byte{"a.go": []byte("hi there"), "extra.txt": []byte("xxxx")}, // 8, 4
			expected: (8 - 2) + 4,
		},
		{
			name:     "zero-byte starter unchanged",
			starter:  map[string][]byte{"empty.txt": []byte("")},
			after:    map[string][]byte{"empty.txt": []byte("")},
			expected: 0,
		},
		{
			name:     "zero-byte starter grows",
			starter:  map[string][]byte{"empty.txt": []byte("")},
			after:    map[string][]byte{"empty.txt": []byte("content added")},
			expected: 13,
		},
		{
			name:     "empty starter project",
			starter:  map[string][]byte{},
			after:    map[string][]byte{"new.go": []byte("content")},
			expected: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			var starterFiles []string
			for name, body := range tc.starter {
				writeFile(t, dir, name, body)
				starterFiles = append(starterFiles, name)
			}
			snap := snapshotStarterSizes(dir, starterFiles)

			// Rewrite workspace to match "after" state.
			for name := range tc.starter {
				_ = os.Remove(filepath.Join(dir, name))
			}
			var currentFiles []string
			for name, body := range tc.after {
				writeFile(t, dir, name, body)
				currentFiles = append(currentFiles, name)
			}
			for _, r := range tc.removed {
				_ = os.Remove(filepath.Join(dir, r))
			}

			got := computeAgentOutputSize(dir, currentFiles, snap)
			if got != tc.expected {
				t.Errorf("agent output size: got %d want %d", got, tc.expected)
			}
		})
	}
}

func TestComputeAgentFileCount(t *testing.T) {
	tests := []struct {
		name     string
		snapshot map[string]int64
		current  []string
		want     int
	}{
		{
			name:     "only starter files present — zero agent files",
			snapshot: map[string]int64{"a.go": 10, "b.go": 20},
			current:  []string{"a.go", "b.go"},
			want:     0,
		},
		{
			name:     "one new file",
			snapshot: map[string]int64{"a.go": 10},
			current:  []string{"a.go", "new.go"},
			want:     1,
		},
		{
			name:     "one deleted starter",
			snapshot: map[string]int64{"a.go": 10, "b.go": 20},
			current:  []string{"a.go"},
			want:     1,
		},
		{
			name:     "new + deleted",
			snapshot: map[string]int64{"a.go": 10, "b.go": 20},
			current:  []string{"a.go", "new.go"},
			want:     2, // 1 new, 1 deleted
		},
		{
			name:     "no starter at all — every current file counts",
			snapshot: map[string]int64{},
			current:  []string{"x.go", "y.go", "z.go"},
			want:     3,
		},
		{
			name:     "zero-byte starter files still count as starter",
			snapshot: map[string]int64{"empty.txt": 0, "README.md": 0},
			current:  []string{"empty.txt", "README.md", "new.go"},
			want:     1, // only new.go
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := computeAgentFileCount(tc.current, tc.snapshot)
			if got != tc.want {
				t.Errorf("agent file count: got %d want %d", got, tc.want)
			}
		})
	}
}
