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
//
// Phase 3 (2026-04-25) invariant: every grader emits ≥ 1 Point. When no knobs
// are configured, a single trivially-passing "no_knobs" Point is emitted so
// the site never falls back to its "PASS"/"100%" header. Labels are stable
// snake_case identifiers so reports are aggregable across runs; Messages
// describe what was checked in both pass and fail cases.
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
		var msg string
		if pass {
			msg = fmt.Sprintf("produced %d file(s), need >= %d", len(produced), g.cfg.MinFiles)
		} else {
			msg = fmt.Sprintf("produced %d file(s), need >= %d", len(produced), g.cfg.MinFiles)
		}
		points = append(points, GraderPoint{
			Label:    "min_files",
			Pass:     pass,
			Message:  msg,
			Evidence: map[string]string{"actual": fmt.Sprintf("%d", len(produced)), "expected": fmt.Sprintf(">=%d", g.cfg.MinFiles)},
		})
	}

	if g.cfg.MaxFiles > 0 {
		pass := len(produced) <= g.cfg.MaxFiles
		var msg string
		if pass {
			msg = fmt.Sprintf("produced %d file(s), <= %d max", len(produced), g.cfg.MaxFiles)
		} else {
			msg = fmt.Sprintf("produced %d file(s), exceeds max of %d", len(produced), g.cfg.MaxFiles)
		}
		points = append(points, GraderPoint{
			Label:    "max_files",
			Pass:     pass,
			Message:  msg,
			Evidence: map[string]string{"actual": fmt.Sprintf("%d", len(produced)), "expected": fmt.Sprintf("<=%d", g.cfg.MaxFiles)},
		})
	}

	// require_files: aggregate into a single Point per knob.
	if len(g.cfg.RequireFiles) > 0 {
		var missing []string
		for _, p := range g.cfg.RequireFiles {
			if !producedSet[p] {
				missing = append(missing, p)
			}
		}
		pass := len(missing) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("all %d required file(s) present", len(g.cfg.RequireFiles))
		} else {
			msg = fmt.Sprintf("missing %d/%d required file(s): %s",
				len(missing), len(g.cfg.RequireFiles), strings.Join(missing, ", "))
		}
		points = append(points, GraderPoint{
			Label:   "require_files",
			Pass:    pass,
			Message: msg,
		})
	}

	// forbid_files: aggregate into a single Point per knob.
	if len(g.cfg.ForbidFiles) > 0 {
		var found []string
		for _, p := range g.cfg.ForbidFiles {
			if producedSet[p] {
				found = append(found, p)
			}
		}
		pass := len(found) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("none of %d forbidden file(s) present", len(g.cfg.ForbidFiles))
		} else {
			msg = fmt.Sprintf("found forbidden file(s): %s", strings.Join(found, ", "))
		}
		points = append(points, GraderPoint{
			Label:   "forbid_files",
			Pass:    pass,
			Message: msg,
		})
	}

	// require_updated: aggregate into a single Point per knob.
	if len(g.cfg.RequireUpdated) > 0 {
		var notUpdated []string
		for _, p := range g.cfg.RequireUpdated {
			if !modifiedSet[p] {
				notUpdated = append(notUpdated, p)
			}
		}
		pass := len(notUpdated) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("all %d path(s) appear in modified files", len(g.cfg.RequireUpdated))
		} else {
			msg = fmt.Sprintf("not modified by agent: %s", strings.Join(notUpdated, ", "))
		}
		points = append(points, GraderPoint{
			Label:   "require_updated",
			Pass:    pass,
			Message: msg,
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
		var msg string
		if len(produced) == 0 {
			msg = fmt.Sprintf("no produced files to check (>= %d B required)", g.cfg.MinBytesPerFile)
		} else if pass {
			msg = fmt.Sprintf("all %d file(s) >= %d byte(s)", len(produced), g.cfg.MinBytesPerFile)
		} else {
			msg = fmt.Sprintf("%d file(s) below %d bytes: %s",
				len(offenders), g.cfg.MinBytesPerFile, strings.Join(offenders, ", "))
		}
		points = append(points, GraderPoint{
			Label:   "min_bytes_per_file",
			Pass:    pass,
			Message: msg,
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
		var msg string
		if pass {
			msg = fmt.Sprintf("all %d file(s) <= %d byte(s)", len(produced), g.cfg.MaxBytesPerFile)
		} else {
			msg = fmt.Sprintf("%d file(s) above %d bytes: %s",
				len(offenders), g.cfg.MaxBytesPerFile, strings.Join(offenders, ", "))
		}
		points = append(points, GraderPoint{
			Label:   "max_bytes_per_file",
			Pass:    pass,
			Message: msg,
		})
	}

	// --- Aggregate. ---
	// Phase 3 invariant: every grader emits ≥ 1 Point. When no knobs are
	// configured the grader still emits a single trivially-passing Point so
	// the site never falls back to its legacy "PASS"/"100%" header.
	if len(points) == 0 {
		points = []GraderPoint{{
			Label:   "no_knobs",
			Pass:    true,
			Message: "no output_check knobs configured — trivially passed",
		}}
	}

	failCount := 0
	for _, p := range points {
		if !p.Pass {
			failCount++
		}
	}
	msg := fmt.Sprintf("output_check passed all %d configured knob(s)", len(points))
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
