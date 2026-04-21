package tool

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// ResolveSkills takes a list of tool entries and resolves the skill entries to
// absolute directory paths. The baseDir is used as the root for resolving
// relative local paths.
//
//   - source: local, skill_dir: false → path points to a single skill (must contain SKILL.md)
//   - source: local, skill_dir: true  → path is a directory of skills (subdirs contain SKILL.md)
//   - source: local with glob         → each glob match is treated as a single skill directory
//   - source: remote                  → dispatched to the registered Fetcher (default: npx)
func ResolveSkills(entries []Entry, baseDir string) ([]string, error) {
	var dirs []string
	for _, entry := range entries {
		if entry.ResolvedType() != TypeSkill {
			continue
		}
		switch entry.SkillSource() {
		case SourceLocal:
			resolved, err := resolveLocal(entry, baseDir)
			if err != nil {
				return nil, fmt.Errorf("resolving local skill %q: %w", entry.Path, err)
			}
			slog.Debug("Resolved local skill", "path", entry.Path, "skill_dir", entry.SkillDir, "resolved_count", len(resolved))
			dirs = append(dirs, resolved...)
		case SourceRemote:
			dir, err := FetchRemote(entry, baseDir)
			if err != nil {
				return nil, fmt.Errorf("fetching remote skill %s/%s: %w", entry.Repo, entry.Name, err)
			}
			dirs = append(dirs, dir)
		default:
			return nil, fmt.Errorf("unknown skill source %q", entry.Source)
		}
	}
	return dirs, nil
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

// FetchRemote fetches a remote skill via the registered Fetcher (default: npx).
// Returns the directory where the skill was installed. The fetcher is chosen
// from DefaultRegistry; users can register custom fetchers via Register to
// override or extend this behavior. Honors entry.Version (which may have been
// set by ApplyVersionOverrides) and falls back to the fetcher's default when
// empty.
func FetchRemote(entry Entry, baseDir string) (string, error) {
	f := LookupFetcher(entry)
	if f == nil {
		return "", fmt.Errorf("no fetcher registered for remote skill %q (repo=%q)", entry.Name, entry.Repo)
	}
	res, err := f.Fetch(context.Background(), FetchRequest{
		Entry:   entry,
		BaseDir: baseDir,
		Version: entry.Version,
	})
	if err != nil {
		return "", err
	}
	return res.Dir, nil
}
