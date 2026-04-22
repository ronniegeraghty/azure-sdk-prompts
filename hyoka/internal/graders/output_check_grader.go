package graders

import (
	"context"
	"fmt"
)

// OutputCheckGrader verifies the agent produced files in the workspace.
//
// It is a generic, content-agnostic grader: it simply asks "did the agent
// generate output, and does that output contain bytes?" Use it to gate or
// signal whether a generation attempt produced any artifact at all, without
// hardcoding that expectation into the eval engine.
//
// Configure via OutputCheckConfig. By default, requires at least one file
// with at least one byte of content.
type OutputCheckGrader struct {
	name            string
	minFiles        int
	minBytesPerFile int64
	minTotalBytes   int64
}

// NewOutputCheckGrader constructs an OutputCheckGrader from a parsed
// OutputCheckConfig. Unset fields fall back to sensible defaults
// (min_files=1, min_bytes_per_file=1, min_total_bytes=0/disabled).
func NewOutputCheckGrader(name string, cfg *OutputCheckConfig) (*OutputCheckGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("output_check grader %q: config is required", name)
	}
	minFiles := cfg.MinFiles
	if minFiles <= 0 {
		minFiles = 1
	}
	minBytes := cfg.MinBytesPerFile
	if minBytes <= 0 {
		minBytes = 1
	}
	if cfg.MinTotalBytes < 0 {
		return nil, fmt.Errorf("output_check grader %q: min_total_bytes must be >= 0", name)
	}
	return &OutputCheckGrader{
		name:            name,
		minFiles:        minFiles,
		minBytesPerFile: minBytes,
		minTotalBytes:   cfg.MinTotalBytes,
	}, nil
}

// Kind returns the grader type identifier.
func (g *OutputCheckGrader) Kind() string { return KindOutputCheck }

// Name returns the human-readable name.
func (g *OutputCheckGrader) Name() string { return g.name }

// Grade counts files in the workspace that meet the per-file size threshold
// and reports pass/fail based on the configured minimums.
func (g *OutputCheckGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	result := GraderResult{
		Kind:   KindOutputCheck,
		Name:   g.name,
		Weight: input.Config.EffectiveWeight(),
		Gate:   input.Config.Gate,
		FileDetails: &FileGraderDetails{
			CheckedFiles: make([]FileCheckResult, 0, len(input.Files)),
		},
	}

	var qualifyingCount int
	var qualifyingBytes int64
	for _, f := range input.Files {
		hasContent := f.Size >= g.minBytesPerFile
		if hasContent {
			qualifyingCount++
			qualifyingBytes += f.Size
		}
		result.FileDetails.CheckedFiles = append(result.FileDetails.CheckedFiles, FileCheckResult{
			Path:   f.Path,
			Exists: hasContent,
		})
	}

	switch {
	case qualifyingCount < g.minFiles:
		result.Score = 0
		result.Pass = false
		result.Message = fmt.Sprintf(
			"expected at least %d file(s) with >= %d byte(s); found %d qualifying of %d total",
			g.minFiles, g.minBytesPerFile, qualifyingCount, len(input.Files),
		)
	case g.minTotalBytes > 0 && qualifyingBytes < g.minTotalBytes:
		result.Score = 0
		result.Pass = false
		result.Message = fmt.Sprintf(
			"expected at least %d total bytes across %d qualifying file(s); found %d bytes",
			g.minTotalBytes, qualifyingCount, qualifyingBytes,
		)
	default:
		result.Score = 1.0
		result.Pass = true
		result.Message = fmt.Sprintf(
			"output_check passed: %d qualifying file(s), %d total byte(s)",
			qualifyingCount, qualifyingBytes,
		)
	}

	return result, nil
}
