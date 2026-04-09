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

	"github.com/ronniegeraghty/hyoka/internal/config"
)

// ResolveSkillDirs takes a list of tool entries and resolves the skill entries to
// absolute directory paths. The baseDir is used as the root for resolving
// relative local paths.
//
//   - source: local  → resolves path (supports glob patterns like "./skills/generator/*")
//   - source: remote → fetches from GitHub repo via "npx skills add", returns the install dir
func ResolveSkillDirs(entries []config.ToolEntry, baseDir string) ([]string, error) {
	var dirs []string
	for _, entry := range entries {
		if entry.ResolvedType() != "skill" {
			continue
		}
		switch entry.SkillSource() {
		case "local":
			resolved, err := resolveLocal(entry.Path, baseDir)
			if err != nil {
				return nil, fmt.Errorf("resolving local skill %q: %w", entry.Path, err)
			}
			slog.Debug("Resolved local skill", "path", entry.Path, "resolved_count", len(resolved))
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

// resolveLocal resolves a local skill path (supports globs) to absolute paths.
func resolveLocal(path, baseDir string) ([]string, error) {
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
	candidates := []string{
		path,
		filepath.Join(baseDir, path),
	}
	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			abs, absErr := filepath.Abs(c)
			if absErr != nil {
				slog.Warn("Failed to resolve absolute path", "path", c, "error", absErr)
			}
			return []string{abs}, nil
		}
	}

	// Path doesn't exist yet — return absolute form anyway
	abs, absErr := filepath.Abs(path)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute path", "path", path, "error", absErr)
	}
	return []string{abs}, nil
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
	args := []string{"skills", "add", entry.Repo, "--directory", installDir}
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
