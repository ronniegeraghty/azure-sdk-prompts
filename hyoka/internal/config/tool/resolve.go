package tool

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// ProgressEmitter is an optional callback for emitting progress events during
// tool resolution. A nil emitter means callers don't care about progress and
// resolution proceeds silently (matching the pre-emission behavior).
type ProgressEmitter func(progress.ProgressEvent)

// ResolveSkills takes a list of tool entries and resolves the skill entries to
// absolute directory paths. The baseDir is used as the root for resolving
// relative local paths.
//
//   - source: local, skill_dir: false → path points to a single skill (must contain SKILL.md)
//   - source: local, skill_dir: true  → path is a directory of skills (subdirs contain SKILL.md)
//   - source: local with glob         → each glob match is treated as a single skill directory
//   - source: remote                  → dispatched to the registered Fetcher (default: npx)
//
// This is a thin wrapper around ResolveSkillsWithReporter that passes a nil
// emitter. Callers that want per-skill ToolResolutionStart/Result events on
// the progress channel should use ResolveSkillsWithReporter directly.
func ResolveSkills(ctx context.Context, entries []Entry, baseDir string) ([]string, error) {
	return ResolveSkillsWithReporter(ctx, entries, baseDir, nil)
}

// ResolveSkillsWithReporter resolves skill entries and, when emit is non-nil,
// emits a ToolResolutionStart / ToolResolutionResult pair for each skill entry
// encountered. Emission is sequential — Start and Result for one skill are
// delivered before the next skill's Start, so renderers bound by the
// "last-line only" update rule can treat each pair as a single tail line.
//
// An emit of nil silently skips progress reporting; behavior is otherwise
// identical to the pre-emission implementation.
func ResolveSkillsWithReporter(ctx context.Context, entries []Entry, baseDir string, emit ProgressEmitter) ([]string, error) {
	var dirs []string
	for _, entry := range entries {
		if entry.ResolvedType() != TypeSkill {
			continue
		}
		emitStart(emit, entry.Name, progress.ToolKindSkill)

		var resolved []string
		var err error
		switch entry.SkillSource() {
		case SourceLocal:
			resolved, err = resolveLocal(entry, baseDir)
			if err != nil {
				err = fmt.Errorf("resolving local skill %q: %w", entry.Path, err)
			} else {
				slog.Debug("Resolved local skill", "path", entry.Path, "skill_dir", entry.SkillDir, "resolved_count", len(resolved))
			}
		case SourceRemote:
			var dir string
			dir, err = FetchRemote(ctx, entry)
			if err != nil {
				err = fmt.Errorf("fetching remote skill %s/%s: %w", entry.Repo, entry.Name, err)
			} else {
				resolved = []string{dir}
			}
		default:
			err = fmt.Errorf("unknown skill source %q", entry.Source)
		}

		if err != nil {
			emitResult(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, err.Error())
			return nil, err
		}
		// A skill entry that resolves to zero directories (missing SKILL.md,
		// empty skill_dir, missing path) is a user-visible failure even though
		// the existing contract treats it as a non-fatal warning.
		if len(resolved) == 0 {
			emitResult(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, "no skill directories resolved")
		} else {
			emitResult(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusLoaded, "")
		}
		dirs = append(dirs, resolved...)
	}
	return dirs, nil
}

func emitStart(emit ProgressEmitter, name, kind string) {
	if emit == nil {
		return
	}
	emit(progress.ProgressEvent{
		Type:     progress.EventToolResolutionStart,
		ToolName: name,
		ToolKind: kind,
	})
}

func emitResult(emit ProgressEmitter, name, kind, status, reason string) {
	if emit == nil {
		return
	}
	emit(progress.ProgressEvent{
		Type:     progress.EventToolResolutionResult,
		ToolName: name,
		ToolKind: kind,
		Status:   status,
		Reason:   reason,
	})
}

// EmitMCPResolutions walks entries, emits a ToolResolutionStart/Result pair
// for every MCP entry, and returns. Unlike skills, MCP entries carry no
// filesystem state that needs resolving at config-load time — "resolution"
// here is a static validation that the entry has the fields its mode
// requires (Command for local, URL for remote). A nil emit is a no-op.
//
// This exists as a sibling of ResolveSkillsWithReporter so the engine can
// render a Tools block that includes MCP lines alongside skill lines before
// the Copilot session starts.
func EmitMCPResolutions(entries []Entry, emit ProgressEmitter) {
	if emit == nil {
		return
	}
	for _, entry := range entries {
		if entry.ResolvedType() != TypeMCP {
			continue
		}
		emitStart(emit, entry.Name, progress.ToolKindMCP)
		status := progress.ToolStatusLoaded
		reason := ""
		switch entry.ResolvedMCPType() {
		case "remote":
			if entry.URL == "" {
				status = progress.ToolStatusFailed
				reason = "remote MCP entry missing url"
			}
		default: // local
			if entry.Command == "" {
				status = progress.ToolStatusFailed
				reason = "local MCP entry missing command"
			}
		}
		emitResult(emit, entry.Name, progress.ToolKindMCP, status, reason)
	}
}

// CountSkills counts the number of actual skill subdirectories (containing
// SKILL.md) within the given resolved skill directories. Each resolved
// directory is expected to be a single skill (containing SKILL.md directly)
// since skill_dir entries are expanded into individual skill dirs during
// resolution.
func CountSkills(dirs []string) int {
	count := 0
	for _, dir := range dirs {
		if _, err := os.Stat(filepath.Join(dir, "SKILL.md")); err == nil {
			count++
		}
	}
	return count
}

// resolveLocal resolves a local skill entry to absolute directory paths.
// When entry.SkillDir is true, the path is a directory of skills — each
// subdirectory containing SKILL.md is returned. When false (default), the
// path points to a single skill directory. Glob patterns are also supported.
func resolveLocal(entry Entry, baseDir string) ([]string, error) {
	path := entry.Path
	// Make relative paths absolute based on baseDir
	if !filepath.IsAbs(path) {
		path = filepath.Join(baseDir, path)
	}

	// Check for glob characters
	if strings.ContainsAny(path, "*?[") {
		matches, err := filepath.Glob(path)
		if err != nil {
			return nil, fmt.Errorf("invalid glob pattern %q: %w", path, err)
		}
		slog.Debug("Skills glob expansion", "pattern", path, "matches", len(matches))
		// Filter to directories only
		var dirs []string
		for _, m := range matches {
			info, err := os.Stat(m)
			if err != nil {
				continue
			}
			if info.IsDir() {
				abs, absErr := filepath.Abs(m)
				if absErr != nil {
					slog.Warn("Failed to resolve absolute path", "path", m, "error", absErr)
				}
				dirs = append(dirs, abs)
			}
		}
		return dirs, nil
	}

	// Non-glob: try to resolve the path, checking candidate locations
	resolved := resolvePath(path, baseDir)
	if resolved == "" {
		slog.Warn("Skills directory does not exist", "path", entry.Path)
		return nil, nil
	}

	if entry.SkillDir {
		if len(entry.ExcludedSkills) > 0 {
			return resolveSkillDirWithExclusions(resolved, entry.ExcludedSkills)
		}
		return resolveSkillDir(resolved)
	}

	// Single skill — validate it contains SKILL.md
	if _, err := os.Stat(filepath.Join(resolved, "SKILL.md")); err != nil {
		slog.Warn("Skill directory missing SKILL.md", "path", resolved)
		return nil, nil
	}
	return []string{resolved}, nil
}

// resolvePath finds the first existing directory matching path or baseDir+path.
func resolvePath(path, baseDir string) string {
	candidates := []string{
		path,
		filepath.Join(baseDir, path),
	}
	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			abs, absErr := filepath.Abs(c)
			if absErr != nil {
				slog.Warn("Failed to resolve absolute path", "path", c, "error", absErr)
				return c
			}
			return abs
		}
	}
	return ""
}

// resolveSkillDir expands a directory of skills into individual skill dirs.
// Each subdirectory containing SKILL.md is returned. If no skills are found,
// a warning is logged and nil is returned.
func resolveSkillDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("Failed to read skill directory", "path", dir, "error", err)
		return nil, nil
	}
	var dirs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		subDir := filepath.Join(dir, e.Name())
		if _, err := os.Stat(filepath.Join(subDir, "SKILL.md")); err == nil {
			dirs = append(dirs, subDir)
		}
	}
	if len(dirs) == 0 {
		slog.Warn("Skills directory contains no skills (no subdirectories with SKILL.md)",
			"path", dir)
	}
	return dirs, nil
}

// resolveSkillDirWithExclusions is like resolveSkillDir but filters out
// subdirectories whose basenames are in the excluded list.
func resolveSkillDirWithExclusions(dir string, excluded []string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("Failed to read skill directory", "path", dir, "error", err)
		return nil, nil
	}
	excludedSet := make(map[string]bool, len(excluded))
	for _, e := range excluded {
		excludedSet[e] = true
	}
	var dirs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if excludedSet[e.Name()] {
			continue // skip excluded skills
		}
		subDir := filepath.Join(dir, e.Name())
		if _, err := os.Stat(filepath.Join(subDir, "SKILL.md")); err == nil {
			dirs = append(dirs, subDir)
		}
	}
	if len(dirs) == 0 {
		slog.Warn("Skills directory contains no skills (no subdirectories with SKILL.md)",
			"path", dir)
	}
	return dirs, nil
}

// FetchRemote fetches a remote skill via the registered Fetcher (default: git).
// Returns the directory where the skill was installed. The fetcher is chosen
// from DefaultRegistry; users can register custom fetchers via Register to
// override or extend this behavior. Honors entry.Version (which may have been
// set by ApplyVersionOverrides) and falls back to the fetcher's default when
// empty. The provided context is passed to the fetcher so cancellation and
// deadlines propagate into any HTTP/exec work it performs.
//
// The on-disk cache location is owned by internal/toolload — callers no
// longer pass a base directory.
func FetchRemote(ctx context.Context, entry Entry) (string, error) {
	f := LookupFetcher(entry)
	if f == nil {
		return "", fmt.Errorf("no fetcher registered for remote skill %q (repo=%q)", entry.Name, entry.Repo)
	}
	res, err := f.Fetch(ctx, FetchRequest{
		Entry:   entry,
		Version: entry.Version,
	})
	if err != nil {
		return "", err
	}
	return res.Dir, nil
}
