package graders

import (
	"context"
	"fmt"
	"sort"
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
func (g *OutputCheckGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	details := &OutputCheckGraderDetails{
		SubChecks: []OutputCheckSubResult{},
	}
	result := GraderResult{
		Kind:               KindOutputCheck,
		Name:               g.name,
		Weight:             input.Config.EffectiveWeight(),
		Gate:               input.Config.Gate,
		OutputCheckDetails: details,
	}

	// Flatten delta into produced files. A nil delta is treated as "empty":
	// the agent produced nothing. This is a deliberate choice so output_check
	// still reports meaningfully (e.g. min_files=1 fails) when the engine
	// couldn't compute a delta.
	produced := flattenDelta(input.WorkspaceDelta)
	modifiedSet := make(map[string]bool, len(produced))
	producedSet := make(map[string]bool, len(produced))
	producedPaths := make([]string, 0, len(produced))
	for _, f := range produced {
		producedSet[f.Path] = true
		if f.Modified {
			modifiedSet[f.Path] = true
		}
		producedPaths = append(producedPaths, f.Path)
	}
	sort.Strings(producedPaths)
	details.ProducedFiles = producedPaths

	// --- Run each configured sub-check. Order is stable for reporting. ---

	if g.cfg.MinFiles > 0 {
		pass := len(produced) >= g.cfg.MinFiles
		msg := fmt.Sprintf("min_files: need >= %d produced file(s), found %d",
			g.cfg.MinFiles, len(produced))
		if pass {
			msg = fmt.Sprintf("min_files: produced %d file(s) (>= %d required)",
				len(produced), g.cfg.MinFiles)
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "min_files", Pass: pass, Message: msg,
		})
	}

	if g.cfg.MaxFiles > 0 {
		pass := len(produced) <= g.cfg.MaxFiles
		msg := fmt.Sprintf("max_files: produced %d file(s) exceeds max of %d",
			len(produced), g.cfg.MaxFiles)
		if pass {
			msg = fmt.Sprintf("max_files: produced %d file(s) (<= %d max)",
				len(produced), g.cfg.MaxFiles)
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "max_files", Pass: pass, Message: msg,
		})
	}

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
			msg = fmt.Sprintf("require_files: all %d required file(s) present",
				len(g.cfg.RequireFiles))
		} else {
			msg = fmt.Sprintf("require_files: missing from output: %s",
				strings.Join(missing, ", "))
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "require_files", Pass: pass, Message: msg,
		})
	}

	if len(g.cfg.ForbidFiles) > 0 {
		var violated []string
		for _, p := range g.cfg.ForbidFiles {
			if producedSet[p] {
				violated = append(violated, p)
			}
		}
		pass := len(violated) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("forbid_files: none of %d forbidden path(s) present",
				len(g.cfg.ForbidFiles))
		} else {
			msg = fmt.Sprintf("forbid_files: found forbidden path(s) in output: %s",
				strings.Join(violated, ", "))
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "forbid_files", Pass: pass, Message: msg,
		})
	}

	if len(g.cfg.RequireUpdated) > 0 {
		var missing []string
		for _, p := range g.cfg.RequireUpdated {
			if !modifiedSet[p] {
				missing = append(missing, p)
			}
		}
		pass := len(missing) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("require_updated: all %d path(s) appear in modified set",
				len(g.cfg.RequireUpdated))
		} else {
			msg = fmt.Sprintf("require_updated: not modified by agent: %s",
				strings.Join(missing, ", "))
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "require_updated", Pass: pass, Message: msg,
		})
	}

	if g.cfg.MinBytesPerFile > 0 {
		var offenders []string
		for _, f := range produced {
			if f.Size < g.cfg.MinBytesPerFile {
				offenders = append(offenders,
					fmt.Sprintf("%s (%d B)", f.Path, f.Size))
			}
		}
		pass := len(offenders) == 0
		var msg string
		switch {
		case len(produced) == 0:
			// No files = no file met the minimum. Treat as fail
			// (vacuous-truth bug fix). If a config doesn't care
			// about file existence, drop min_bytes_per_file.
			pass = false
			msg = fmt.Sprintf("min_bytes_per_file: no produced files to check (>= %d required)",
				g.cfg.MinBytesPerFile)
		case pass:
			msg = fmt.Sprintf("min_bytes_per_file: all %d file(s) >= %d byte(s)",
				len(produced), g.cfg.MinBytesPerFile)
		default:
			msg = fmt.Sprintf("min_bytes_per_file: %d file(s) below %d bytes: %s",
				len(offenders), g.cfg.MinBytesPerFile, strings.Join(offenders, ", "))
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "min_bytes_per_file", Pass: pass, Message: msg,
		})
	}

	if g.cfg.MaxBytesPerFile > 0 {
		var offenders []string
		for _, f := range produced {
			if f.Size > g.cfg.MaxBytesPerFile {
				offenders = append(offenders,
					fmt.Sprintf("%s (%d B)", f.Path, f.Size))
			}
		}
		pass := len(offenders) == 0
		var msg string
		if pass {
			msg = fmt.Sprintf("max_bytes_per_file: all %d file(s) <= %d byte(s)",
				len(produced), g.cfg.MaxBytesPerFile)
		} else {
			msg = fmt.Sprintf("max_bytes_per_file: %d file(s) above %d bytes: %s",
				len(offenders), g.cfg.MaxBytesPerFile, strings.Join(offenders, ", "))
		}
		details.SubChecks = append(details.SubChecks, OutputCheckSubResult{
			Check: "max_bytes_per_file", Pass: pass, Message: msg,
		})
	}

	// --- Aggregate. ---
	if len(details.SubChecks) == 0 {
		result.Pass = true
		result.Score = 1.0
		result.Message = "output_check: no knobs configured — trivially passed"
		result.Points = []GraderPoint{{
			Name:    "no_knobs_configured",
			Pass:    true,
			Message: "no knobs configured — trivially passed",
		}}
		return result, nil
	}

	overallPass := true
	var failedLines []string
	var passedLines []string
	result.Points = make([]GraderPoint, 0, len(details.SubChecks))
	for _, sc := range details.SubChecks {
		result.Points = append(result.Points, GraderPoint{
			Name: sc.Check, Pass: sc.Pass, Message: sc.Message,
		})
		if sc.Pass {
			passedLines = append(passedLines, sc.Message)
			continue
		}
		overallPass = false
		failedLines = append(failedLines, sc.Message)
	}

	result.Pass = overallPass
	if overallPass {
		result.Score = 1.0
		result.Message = fmt.Sprintf("output_check passed (%d/%d): %s",
			len(details.SubChecks), len(details.SubChecks),
			strings.Join(passedLines, "; "))
	} else {
		result.Score = 0
		result.Message = fmt.Sprintf("output_check failed (%d/%d sub-checks failed): %s",
			len(failedLines), len(details.SubChecks),
			strings.Join(failedLines, "; "))
	}

	return result, nil
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
