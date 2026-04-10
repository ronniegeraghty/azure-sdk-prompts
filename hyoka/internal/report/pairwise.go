package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/pairwise"
)

// PairwiseRunReport captures per-prompt and aggregate pairwise results for a run.
type PairwiseRunReport struct {
	RunID     string                     `json:"run_id"`
	Timestamp string                     `json:"timestamp"`
	Reports   []*pairwise.PairwiseReport `json:"reports"`
	Impacts   []pairwise.ToolImpact      `json:"aggregate_impacts,omitempty"`
}

// WritePairwiseReport writes pairwise results to the run output directory.
func WritePairwiseReport(runID string, reports []*pairwise.PairwiseReport, impacts []pairwise.ToolImpact, outputDir string) (string, error) {
	if len(reports) == 0 {
		return "", nil
	}

	reportDir := filepath.Join(outputDir, runID)
	if err := os.MkdirAll(reportDir, 0755); err != nil {
		return "", fmt.Errorf("creating pairwise report directory: %w", err)
	}

	path := filepath.Join(reportDir, "pairwise.json")
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("creating pairwise report file: %w", err)
	}
	defer f.Close()

	payload := PairwiseRunReport{
		RunID:     runID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Reports:   reports,
		Impacts:   impacts,
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(payload); err != nil {
		return "", fmt.Errorf("encoding pairwise report: %w", err)
	}

	return path, nil
}
