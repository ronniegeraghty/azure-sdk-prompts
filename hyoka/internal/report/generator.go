// Package report handles generation of JSON and Markdown reports.
package report

import (
"encoding/json"
"fmt"
"log/slog"
"os"
"path/filepath"

"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// ReportDir returns the directory path for a specific evaluation report.
// The path includes the prompt ID so that different prompts sharing the same
// service/plane/language/category get isolated workspace directories.
func ReportDir(outputDir string, runID string, p *prompt.Prompt) string {
	return filepath.Join(
		outputDir, runID, "results",
		p.Service(), p.Plane(), p.Language(), p.Category(), p.ID,
	)
}

// WriteReport writes an EvalReport as JSON to the appropriate directory.
// Large reports are truncated before writing (see TruncateReport) and
// streamed directly to disk via json.NewEncoder to avoid holding the full
// serialized form in memory.
func WriteReport(r *EvalReport, outputDir string, runID string, p *prompt.Prompt) (string, error) {
	reportDir := filepath.Join(
		ReportDir(outputDir, runID, p), r.ConfigName,
	)

	if err := os.MkdirAll(reportDir, 0755); err != nil {
		return "", fmt.Errorf("creating report directory: %w", err)
	}

	// Build the action timeline from session events if not already set.
	if r.ActionTimeline == nil && len(r.SessionEvents) > 0 {
		r.ActionTimeline = BuildActionTimelineWithSetup(r.SessionEvents, r.SessionSetup)
	}

	// Ensure schema version is current.
	MigrateToV2(r)

	// Truncate verbose fields when the report is excessively large.
	TruncateReport(r)

	reportPath := filepath.Join(reportDir, "report.json")

	f, err := os.Create(reportPath)
	if err != nil {
		return "", fmt.Errorf("creating report file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(r); err != nil {
		return "", fmt.Errorf("encoding report: %w", err)
	}

	info, _ := f.Stat()
	var size int64
	if info != nil {
		size = info.Size()
	}
	slog.Debug("Report written", "path", reportPath, "size", size)
	return reportPath, nil
}

// WriteSummary writes a RunSummary as JSON, streaming directly to disk.
func WriteSummary(s *RunSummary, outputDir string) (string, error) {
	summaryDir := filepath.Join(outputDir, s.RunID)
	if err := os.MkdirAll(summaryDir, 0755); err != nil {
		return "", fmt.Errorf("creating summary directory: %w", err)
	}

	summaryPath := filepath.Join(summaryDir, "summary.json")

	f, err := os.Create(summaryPath)
	if err != nil {
		return "", fmt.Errorf("creating summary file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(s); err != nil {
		return "", fmt.Errorf("encoding summary: %w", err)
	}

	info, _ := f.Stat()
	var size int64
	if info != nil {
		size = info.Size()
	}
	slog.Debug("Summary written", "path", summaryPath, "size", size)
	return summaryPath, nil
}
