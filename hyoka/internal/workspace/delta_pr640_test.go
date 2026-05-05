package workspace

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/utils"
)

// TestTakeSnapshot_BuildArtifactDirsExcluded verifies that well-known build
// artifact directories (target/, node_modules/, bin/, obj/, dist/, build/,
// vendor/, .venv/, venv/, __pycache__/, etc.) are skipped during snapshot.
// Regression test for PR #640 Fix 1.
func TestTakeSnapshot_BuildArtifactDirsExcluded(t *testing.T) {
	dir := t.TempDir()

	// Create real source files that should be included
	writeFile(t, filepath.Join(dir, "main.go"), "package main\n\nfunc main() {}")
	writeFile(t, filepath.Join(dir, "README.md"), "# Project\n")
	writeFile(t, filepath.Join(dir, "src", "lib.rs"), "fn helper() {}")

	// Create build artifact directories that should be excluded
	buildDirs := []string{
		"target",       // Rust/Cargo
		"node_modules", // Node.js/npm
		"bin",          // .NET
		"obj",          // .NET
		"dist",         // general build output
		"build",        // general build output
		"vendor",       // Go/PHP
		".venv",        // Python virtual env
		"venv",         // Python virtual env
		"__pycache__",  // Python cache
	}

	for _, buildDir := range buildDirs {
		dirPath := filepath.Join(dir, buildDir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			t.Fatalf("failed to create %s: %v", buildDir, err)
		}
		// Add dummy files inside each build artifact directory
		writeFile(t, filepath.Join(dirPath, "artifact.bin"), "binary data")
		writeFile(t, filepath.Join(dirPath, "nested", "deep.o"), "compiled object")
	}

	snap, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	// Verify real source files are included
	expectedFiles := []string{"main.go", "README.md", filepath.Join("src", "lib.rs")}
	for _, expectedFile := range expectedFiles {
		if _, ok := snap.files[expectedFile]; !ok {
			t.Errorf("expected file %q not found in snapshot", expectedFile)
		}
	}

	// Verify no files from build artifact directories are included
	for path := range snap.files {
		for _, buildDir := range buildDirs {
			// Check if any component of the path matches a build artifact dir
			if containsPathComponent(path, buildDir) {
				t.Errorf("build artifact file should be excluded: %s (from %s/)", path, buildDir)
			}
		}
	}

	// Sanity check: snapshot should have exactly 3 files (the real source files)
	if len(snap.files) != 3 {
		t.Errorf("expected 3 files in snapshot, got %d: %v", len(snap.files), snap.files)
	}
}

// TestTakeSnapshot_AllBuildArtifactDirsFromUtils verifies that the snapshot
// exclusion matches the exact list from utils.IsDefaultExcludedDir.
// This ensures the implementation stays in sync with the exclusion list.
func TestTakeSnapshot_AllBuildArtifactDirsFromUtils(t *testing.T) {
	// This test validates the contract: TakeSnapshot must skip directories
	// where utils.IsDefaultExcludedDir returns true.

	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "real.txt"), "real content")

	// All directories from utils.IsDefaultExcludedDir (as of PR #640)
	excludedDirs := []string{
		"target", "node_modules", "__pycache__", ".venv", "venv",
		"bin", "obj", "build", "dist", "out", "vendor", "packages",
		".gradle", ".cargo", "debug", "release",
	}

	for _, excluded := range excludedDirs {
		// Verify our assumptions about utils.IsDefaultExcludedDir
		if !utils.IsDefaultExcludedDir(excluded) {
			t.Fatalf("test assumption broken: utils.IsDefaultExcludedDir(%q) should be true", excluded)
		}

		dirPath := filepath.Join(dir, excluded)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			t.Fatalf("failed to create %s: %v", excluded, err)
		}
		writeFile(t, filepath.Join(dirPath, "artifact.dat"), "excluded content")
	}

	snap, err := TakeSnapshot(dir)
	if err != nil {
		t.Fatalf("TakeSnapshot failed: %v", err)
	}

	// Only real.txt should be in the snapshot
	if len(snap.files) != 1 {
		t.Errorf("expected 1 file, got %d: %v", len(snap.files), snap.files)
	}
	if _, ok := snap.files["real.txt"]; !ok {
		t.Errorf("real.txt not in snapshot")
	}

	// Verify none of the excluded files made it in
	for path := range snap.files {
		for _, excluded := range excludedDirs {
			if containsPathComponent(path, excluded) {
				t.Errorf("excluded directory %q leaked into snapshot via file: %s", excluded, path)
			}
		}
	}
}

// containsPathComponent checks if any component of path matches dir.
func containsPathComponent(path, dir string) bool {
	components := filepath.SplitList(filepath.ToSlash(path))
	for _, component := range components {
		if component == dir {
			return true
		}
	}
	// Also check using filepath split
	for p := path; p != "" && p != "." && p != "/"; p = filepath.Dir(p) {
		if filepath.Base(p) == dir {
			return true
		}
	}
	return false
}

// writeFile helper is defined in delta_test.go
