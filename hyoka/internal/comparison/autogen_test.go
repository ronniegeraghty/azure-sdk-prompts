package comparison

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

func TestWriteForRun_MultiConfig(t *testing.T) {
	dir := t.TempDir()
	runDir := filepath.Join(dir, "run-001")

	reports := []report.EvalReport{
		mkReport("p1", "alpha", "2025-01-01T10:00:00Z", 0.5),
		mkReport("p1", "beta", "2025-01-01T10:00:00Z", 0.8),
		mkReport("p2", "alpha", "2025-01-01T10:00:00Z", 0.6),
		mkReport("p2", "beta", "2025-01-01T10:00:00Z", 0.6),
	}

	path, err := WriteForRun(runDir, reports)
	if err != nil {
		t.Fatalf("WriteForRun failed: %v", err)
	}
	if path == "" {
		t.Fatal("expected non-empty path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading back: %v", err)
	}
	var results []ComparisonResult
	if err := json.Unmarshal(data, &results); err != nil {
		t.Fatalf("decoding: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 pair, got %d", len(results))
	}
	if results[0].LabelA != "alpha" || results[0].LabelB != "beta" {
		t.Errorf("labels: got %q/%q", results[0].LabelA, results[0].LabelB)
	}
}

func TestWriteForRun_SingleConfig_Skipped(t *testing.T) {
	dir := t.TempDir()
	runDir := filepath.Join(dir, "run-solo")

	reports := []report.EvalReport{
		mkReport("p1", "only", "2025-01-01T10:00:00Z", 0.5),
	}

	path, err := WriteForRun(runDir, reports)
	if err != nil {
		t.Fatalf("WriteForRun failed: %v", err)
	}
	if path != "" {
		t.Errorf("expected empty path for single-config run, got %q", path)
	}
	if _, err := os.Stat(filepath.Join(runDir, ComparisonsFile)); !os.IsNotExist(err) {
		t.Error("comparisons.json should not exist for single-config run")
	}
}

func TestLoadForRun_Missing(t *testing.T) {
	dir := t.TempDir()
	results, err := LoadForRun(dir)
	if err != nil {
		t.Fatalf("expected nil error for missing file, got %v", err)
	}
	if results != nil {
		t.Errorf("expected nil results, got %d entries", len(results))
	}
}

func TestLoadForRun_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	runDir := filepath.Join(dir, "run-rt")

	reports := []report.EvalReport{
		mkReport("p1", "alpha", "2025-01-01T10:00:00Z", 0.5),
		mkReport("p1", "beta", "2025-01-01T10:00:00Z", 0.9),
	}

	if _, err := WriteForRun(runDir, reports); err != nil {
		t.Fatalf("WriteForRun: %v", err)
	}
	results, err := LoadForRun(runDir)
	if err != nil {
		t.Fatalf("LoadForRun: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Kind != KindConfigs {
		t.Errorf("kind: got %q", results[0].Kind)
	}
	if results[0].Summary.Improved != 1 {
		t.Errorf("expected 1 improved, got %d", results[0].Summary.Improved)
	}
}
