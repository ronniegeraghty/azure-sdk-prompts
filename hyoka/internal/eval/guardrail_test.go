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
