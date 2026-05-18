package eval

import (
	"os"
	"path/filepath"
	"testing"
)

// TestIsolateGraderWorkspace_CopiesAndIsolates verifies that
// IsolateGraderWorkspace returns a fresh directory that contains a copy of
// the source contents and that mutations to the isolated copy never leak
// back to the source.
func TestIsolateGraderWorkspace_CopiesAndIsolates(t *testing.T) {
	source := t.TempDir()
	if err := os.WriteFile(filepath.Join(source, "main.py"), []byte("print('hi')"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(source, "pkg"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "pkg", "lib.py"), []byte("x = 1"), 0644); err != nil {
		t.Fatal(err)
	}

	isolated, cleanup, err := IsolateGraderWorkspace(source)
	if err != nil {
		t.Fatalf("IsolateGraderWorkspace error: %v", err)
	}
	defer cleanup()

	if isolated == source {
		t.Fatal("isolated path must differ from source")
	}

	// Source contents must be present in the isolated copy.
	if data, err := os.ReadFile(filepath.Join(isolated, "main.py")); err != nil || string(data) != "print('hi')" {
		t.Fatalf("isolated main.py missing or wrong: %v / %q", err, data)
	}
	if data, err := os.ReadFile(filepath.Join(isolated, "pkg", "lib.py")); err != nil || string(data) != "x = 1" {
		t.Fatalf("isolated pkg/lib.py missing or wrong: %v / %q", err, data)
	}

	// Simulate a mutating grader: write a new file and overwrite an
	// existing one in the isolated copy.
	if err := os.WriteFile(filepath.Join(isolated, "side_effect.txt"), []byte("polluted"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(isolated, "main.py"), []byte("MUTATED"), 0644); err != nil {
		t.Fatal(err)
	}

	// Source must remain pristine.
	if _, err := os.Stat(filepath.Join(source, "side_effect.txt")); !os.IsNotExist(err) {
		t.Errorf("source workspace was polluted by grader mutation: %v", err)
	}
	if data, err := os.ReadFile(filepath.Join(source, "main.py")); err != nil || string(data) != "print('hi')" {
		t.Errorf("source main.py was mutated: %v / %q", err, data)
	}
}

// TestIsolateGraderWorkspace_TwoGradersDoNotCrossContaminate verifies that
// running two mutating graders in sequence — each receiving its own isolated
// copy — does not allow the second grader to see the first's mutations.
func TestIsolateGraderWorkspace_TwoGradersDoNotCrossContaminate(t *testing.T) {
	source := t.TempDir()
	if err := os.WriteFile(filepath.Join(source, "data.txt"), []byte("clean"), 0644); err != nil {
		t.Fatal(err)
	}

	// Grader A: writes a file into its isolated dir.
	dirA, cleanA, err := IsolateGraderWorkspace(source)
	if err != nil {
		t.Fatalf("isolate A: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dirA, "from_a.txt"), []byte("a"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dirA, "data.txt"), []byte("a-mutated"), 0644); err != nil {
		t.Fatal(err)
	}
	cleanA()

	// Grader B: must see a clean workspace, not Grader A's leftovers.
	dirB, cleanB, err := IsolateGraderWorkspace(source)
	if err != nil {
		t.Fatalf("isolate B: %v", err)
	}
	defer cleanB()

	if dirA == dirB {
		t.Fatal("expected distinct isolated dirs across grader iterations")
	}
	if _, err := os.Stat(filepath.Join(dirB, "from_a.txt")); !os.IsNotExist(err) {
		t.Errorf("grader B saw grader A's leftover file")
	}
	if data, err := os.ReadFile(filepath.Join(dirB, "data.txt")); err != nil || string(data) != "clean" {
		t.Errorf("grader B saw grader A's mutation: %v / %q", err, data)
	}

	// And the source still untouched after both graders ran.
	if data, err := os.ReadFile(filepath.Join(source, "data.txt")); err != nil || string(data) != "clean" {
		t.Errorf("source data.txt mutated across grader iterations: %v / %q", err, data)
	}
	if _, err := os.Stat(filepath.Join(source, "from_a.txt")); !os.IsNotExist(err) {
		t.Errorf("source was polluted by grader A: %v", err)
	}
}

// TestIsolateGraderWorkspace_CleanupRemovesDir verifies the returned cleanup
// function actually removes the isolated directory.
func TestIsolateGraderWorkspace_CleanupRemovesDir(t *testing.T) {
	source := t.TempDir()
	os.WriteFile(filepath.Join(source, "f"), []byte("x"), 0644)

	isolated, cleanup, err := IsolateGraderWorkspace(source)
	if err != nil {
		t.Fatalf("isolate: %v", err)
	}
	if _, err := os.Stat(isolated); err != nil {
		t.Fatalf("isolated dir should exist: %v", err)
	}

	cleanup()

	if _, err := os.Stat(isolated); !os.IsNotExist(err) {
		t.Errorf("cleanup should remove isolated dir: %v", err)
	}
}
