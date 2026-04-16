// Package skills resolves unified Skill entries into local directory paths
// that can be passed to the Copilot SDK session as skill directories.
package skills

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
)

// ResolveSkillDirs takes a list of tool entries and resolves the skill entries to
// absolute directory paths. The baseDir is used as the root for resolving
// relative local paths.
//
//   - source: local, skill_dir: false → path points to a single skill (must contain SKILL.md)
//   - source: local, skill_dir: true  → path is a directory of skills (subdirs contain SKILL.md)
//   - source: local with glob         → each glob match is treated as a single skill directory
//   - source: remote                  → fetches from GitHub repo via "npx skills add"
func ResolveSkillDirs(entries []config.ToolEntry, baseDir string) ([]string, error) {
	var dirs []string
	for _, entry := range entries {
		if entry.ResolvedType() != "skill" {
			continue
		}
		switch entry.SkillSource() {
		case "local":
			resolved, err := resolveLocal(entry, baseDir)
			if err != nil {
				return nil, fmt.Errorf("resolving local skill %q: %w", entry.Path, err)
			}
			slog.Debug("Resolved local skill", "path", entry.Path, "skill_dir", entry.SkillDir, "resolved_count", len(resolved))
			dirs = append(dirs, resolved...)
		case "remote":
			dir, err := fetchRemote(entry, baseDir)
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
func resolveLocal(entry config.ToolEntry, baseDir string) ([]string, error) {
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

// fetchRemote fetches a remote skill from a GitHub repo using npx skills add.
// Returns the directory where the skill was installed.
func fetchRemote(entry config.ToolEntry, baseDir string) (string, error) {
	// Determine install directory: use a skills cache dir under baseDir
	installDir := filepath.Join(baseDir, ".skills-cache", entry.Repo)
	if entry.Name != "" {
		installDir = filepath.Join(installDir, entry.Name)
	}

	if err := os.MkdirAll(installDir, 0755); err != nil {
		return "", fmt.Errorf("creating skill install dir: %w", err)
	}

	// Use npx skills add to fetch the skill
	args := []string{"skills", "add", entry.Repo, "--directory", installDir, "--yes"}
	if entry.Name != "" {
		args = append(args, "--name", entry.Name)
	}

	fmt.Printf("Fetching remote skill: %s (repo: %s)\n", entry.Name, entry.Repo)
	slog.Info("Fetching remote skill", "skill", entry.Name, "repo", entry.Repo)
	cmd := exec.Command("npx", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("npx skills add: %w", err)
	}

	abs, absErr := filepath.Abs(installDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute install path", "path", installDir, "error", absErr)
	}
	return abs, nil
}
