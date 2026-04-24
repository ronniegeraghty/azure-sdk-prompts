package graders

import (
	"context"
	"fmt"
	"strings"
)

// OutputCheckGrader is a boolean grader that evaluates workspace changes
// the agent produced during a run (created or modified files, tracked via
// WorkspaceDelta) against a set of configured knobs.
//
// Semantics (locked Q7, 2026-04-22):
//   - Each configured knob yields one sub-check with pass/fail + reason.
//   - Unconfigured knobs are skipped — no implicit defaults.
//   - Sub-checks never early-exit; all run and all are reported.
//   - Overall Pass is the AND of every configured sub-check.
//   - Score is 1.0 iff Pass, else 0.0 (boolean grader, no partial credit).
//   - If no knob is configured, the grader passes trivially with score 1.0
//     and a message noting no checks were defined.
//
// "Files the agent produced" = WorkspaceDelta.NewFiles ∪ WorkspaceDelta.ModifiedFiles.
// Starter files the agent did not touch are NOT considered output.
//
// v1 knobs: min_files, max_files, require_files, forbid_files,
// require_updated, min_bytes_per_file, max_bytes_per_file.
// Globs/regex are deferred past v1.
type OutputCheckGrader struct {
	name string
	cfg  OutputCheckConfig
}

// NewOutputCheckGrader constructs an OutputCheckGrader from a parsed config.
// Returns an error only for structurally invalid config (negative values,
// inverted min/max pairs). A config with no knobs set is accepted; the
// grader will trivially pass at runtime.
func NewOutputCheckGrader(name string, cfg *OutputCheckConfig) (*OutputCheckGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("output_check grader %q: config is required", name)
	}
	if cfg.MinFiles < 0 {
		return nil, fmt.Errorf("output_check grader %q: min_files must be >= 0", name)
	}
	if cfg.MaxFiles < 0 {
		return nil, fmt.Errorf("output_check grader %q: max_files must be >= 0", name)
	}
	if cfg.MinFiles > 0 && cfg.MaxFiles > 0 && cfg.MinFiles > cfg.MaxFiles {
		return nil, fmt.Errorf("output_check grader %q: min_files (%d) > max_files (%d)",
			name, cfg.MinFiles, cfg.MaxFiles)
	}
	if cfg.MinBytesPerFile < 0 {
		return nil, fmt.Errorf("output_check grader %q: min_bytes_per_file must be >= 0", name)
	}
	if cfg.MaxBytesPerFile < 0 {
		return nil, fmt.Errorf("output_check grader %q: max_bytes_per_file must be >= 0", name)
	}
	if cfg.MinBytesPerFile > 0 && cfg.MaxBytesPerFile > 0 && cfg.MinBytesPerFile > cfg.MaxBytesPerFile {
		return nil, fmt.Errorf(
			"output_check grader %q: min_bytes_per_file (%d) > max_bytes_per_file (%d)",
			name, cfg.MinBytesPerFile, cfg.MaxBytesPerFile)
	}
	return &OutputCheckGrader{name: name, cfg: *cfg}, nil
}

// Kind returns the grader type identifier.
func (g *OutputCheckGrader) Kind() string { return KindOutputCheck }

// Name returns the human-readable name.
func (g *OutputCheckGrader) Name() string { return g.name }

// producedFile is an internal flat representation of one file the agent
// created or modified, with its final size.
type producedFile struct {
	Path     string
	Size     int64
	Modified bool // true if in ModifiedFiles, false if in NewFiles
}

// Grade evaluates every configured knob against input.WorkspaceDelta and
// returns a GraderResult whose Pass is the AND of every sub-check.
// Per v4 spec: one Point per configured knob, OutputCheckExtras carries ProducedFiles.
func (g *OutputCheckGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	// Flatten delta into produced files. A nil delta is treated as "empty":
	// the agent produced nothing. This is a deliberate choice so output_check
	// still reports meaningfully (e.g. min_files=1 fails) when the engine
	// couldn't compute a delta.
	produced := flattenDelta(input.WorkspaceDelta)
	modifiedSet := make(map[string]bool, len(produced))
	producedSet := make(map[string]bool, len(produced))
	producedFiles := make([]FileEntry, 0, len(produced))
	for _, f := range produced {
		producedSet[f.Path] = true
		if f.Modified {
			modifiedSet[f.Path] = true
		}
		producedFiles = append(producedFiles, FileEntry{Path: f.Path, Size: f.Size})
	}

	var points []GraderPoint

	// --- Run each configured sub-check. Order is stable for reporting. ---

	if g.cfg.MinFiles > 0 {
		pass := len(produced) >= g.cfg.MinFiles
		label := fmt.Sprintf("min files: ≥ %d", g.cfg.MinFiles)
		var pointMsg string
		if !pass {
			pointMsg = fmt.Sprintf("produced %d file(s), need >= %d", len(produced), g.cfg.MinFiles)
		}
		points = append(points, GraderPoint{
			Label:    label,
			Pass:     pass,
			Message:  pointMsg,
			Evidence: map[string]string{"actual": fmt.Sprintf("%d", len(produced)), "expected": fmt.Sprintf(">=%d", g.cfg.MinFiles)},
		})
	}

	if g.cfg.MaxFiles > 0 {
		pass := len(produced) <= g.cfg.MaxFiles
		label := fmt.Sprintf("max files: ≤ %d", g.cfg.MaxFiles)
		var pointMsg string
		if !pass {
			pointMsg = fmt.Sprintf("produced %d file(s), exceeds max of %d", len(produced), g.cfg.MaxFiles)
		}
		points = append(points, GraderPoint{
			Label:    label,
			Pass:     pass,
			Message:  pointMsg,
			Evidence: map[string]string{"actual": fmt.Sprintf("%d", len(produced)), "expected": fmt.Sprintf("<=%d", g.cfg.MaxFiles)},
		})
	}

	// require_files: one Point per file
	for _, p := range g.cfg.RequireFiles {
		present := producedSet[p]
		label := fmt.Sprintf("require file: %s", p)
		var pointMsg string
		if !present {
			pointMsg = fmt.Sprintf("file %s not found in output", p)
		}
		points = append(points, GraderPoint{
			Label:   label,
			Pass:    present,
			Message: pointMsg,
		})
	}

	// forbid_files: one Point per file
	for _, p := range g.cfg.ForbidFiles {
		absent := !producedSet[p]
		label := fmt.Sprintf("forbid file: %s", p)
		var pointMsg string
		if !absent {
			pointMsg = fmt.Sprintf("forbidden file %s found in output", p)
		}
		points = append(points, GraderPoint{
			Label:   label,
			Pass:    absent,
			Message: pointMsg,
		})
	}

	// require_updated: one Point per file
	for _, p := range g.cfg.RequireUpdated {
		updated := modifiedSet[p]
		label := fmt.Sprintf("require updated: %s", p)
		var pointMsg string
		if !updated {
			pointMsg = fmt.Sprintf("file %s not modified by agent", p)
		}
		points = append(points, GraderPoint{
			Label:   label,
			Pass:    updated,
			Message: pointMsg,
		})
	}

	if g.cfg.MinBytesPerFile > 0 {
		var offenders []string
		for _, f := range produced {
			if f.Size < g.cfg.MinBytesPerFile {
				offenders = append(offenders, fmt.Sprintf("%s (%d B)", f.Path, f.Size))
			}
		}
		pass := len(offenders) == 0
		// No files = fail (vacuous-truth fix)
		if len(produced) == 0 {
			pass = false
		}
		label := fmt.Sprintf("min bytes/file: ≥ %d", g.cfg.MinBytesPerFile)
		var pointMsg string
		if len(produced) == 0 {
			pointMsg = fmt.Sprintf("no produced files to check (>= %d B required)", g.cfg.MinBytesPerFile)
		} else if !pass {
			pointMsg = fmt.Sprintf("%d file(s) below %d bytes: %s", len(offenders), g.cfg.MinBytesPerFile, strings.Join(offenders, ", "))
		}
		points = append(points, GraderPoint{
			Label:   label,
			Pass:    pass,
			Message: pointMsg,
		})
	}

	if g.cfg.MaxBytesPerFile > 0 {
		var offenders []string
		for _, f := range produced {
			if f.Size > g.cfg.MaxBytesPerFile {
				offenders = append(offenders, fmt.Sprintf("%s (%d B)", f.Path, f.Size))
			}
		}
		pass := len(offenders) == 0
		label := fmt.Sprintf("max bytes/file: ≤ %d", g.cfg.MaxBytesPerFile)
		var pointMsg string
		if !pass {
			pointMsg = fmt.Sprintf("%d file(s) above %d bytes: %s", len(offenders), g.cfg.MaxBytesPerFile, strings.Join(offenders, ", "))
		}
		points = append(points, GraderPoint{
			Label:   label,
			Pass:    pass,
			Message: pointMsg,
		})
	}

	// --- Aggregate. ---
	if len(points) == 0 {
		points = []GraderPoint{{
			Label:   "no knobs configured",
			Pass:    true,
			Message: "no knobs configured — trivially passed",
		}}
	}

	msg := fmt.Sprintf("output_check passed all %d configured knob(s)", len(points))
	// Check if any failed
	failCount := 0
	for _, p := range points {
		if !p.Pass {
			failCount++
		}
	}
	if failCount > 0 {
		msg = fmt.Sprintf("output_check failed %d/%d knob(s)", failCount, len(points))
	}

	extras := &GraderExtras{
		OutputCheck: &OutputCheckExtras{
			ProducedFiles: producedFiles,
		},
	}

	return NewResult(KindOutputCheck, g.name, input.Config, points, msg, extras), nil
}

// flattenDelta returns the agent's produced files (new + modified) with
// their final sizes. A nil delta returns an empty slice.
func flattenDelta(d *WorkspaceDelta) []producedFile {
	if d == nil {
		return nil
	}
	out := make([]producedFile, 0, len(d.NewFiles)+len(d.ModifiedFiles))
	for _, f := range d.NewFiles {
		out = append(out, producedFile{Path: f.Path, Size: f.Size, Modified: false})
	}
	for _, f := range d.ModifiedFiles {
		out = append(out, producedFile{Path: f.Path, Size: f.SizeAfter, Modified: true})
	}
	return out
}
