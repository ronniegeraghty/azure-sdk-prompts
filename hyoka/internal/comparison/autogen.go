package comparison

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// ComparisonsFile is the filename written in a run directory containing all
// pairwise config comparisons auto-generated for a multi-config run. The site
// run-detail view reads this file to render comparisons without recomputing.
const ComparisonsFile = "comparisons.json"

// WriteForRun auto-generates pairwise comparisons across the configs present
// in `reports` and writes them as JSON to {runDir}/comparisons.json. Returns
// the written path, nil if nothing was written (fewer than 2 configs), or an
// error on write failure.
//
// The comparisons are produced by the same engine (CompareReports) used by the
// CLI and serve API — site render and CLI output are identical by construction.
func WriteForRun(runDir string, reports []report.EvalReport) (string, error) {
	results := AutoGenerateForRun(reports)
	if len(results) == 0 {
		slog.Debug("Skipping comparisons.json (fewer than 2 configs)", "run_dir", runDir)
		return "", nil
	}

	if err := os.MkdirAll(runDir, 0755); err != nil {
		return "", fmt.Errorf("creating run directory: %w", err)
	}

	path := filepath.Join(runDir, ComparisonsFile)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("creating %s: %w", ComparisonsFile, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(results); err != nil {
		return "", fmt.Errorf("encoding comparisons: %w", err)
	}

	slog.Info("Comparisons written",
		"path", path, "pairs", len(results))
	return path, nil
}

// LoadForRun reads the auto-generated comparisons for a run, if present.
// Returns (nil, nil) if the file does not exist — callers should treat that
// as "no auto-generated comparisons available" (e.g. single-config runs).
func LoadForRun(runDir string) ([]ComparisonResult, error) {
	path := filepath.Join(runDir, ComparisonsFile)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading %s: %w", ComparisonsFile, err)
	}
	var results []ComparisonResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", ComparisonsFile, err)
	}
	return results, nil
}
