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
//   - source: remote, repo "owner/name"         → fetches via "npx skills add"
//   - source: remote, repo "owner/name/subpath" → git sparse-checkout of that subpath
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

// parseRepoSpec splits a repo spec of the form "owner/name" or
// "owner/name/sub/path" into the repo slug (first two segments) and an
// optional subpath (remaining segments, joined with "/"). An empty subpath
// means "whole repo" — the caller should use the npx fetch path. A non-empty
// subpath means "fetch just this directory" — the caller should use the
// git sparse-checkout fetch path.
func parseRepoSpec(spec string) (repo, subpath string) {
	parts := strings.Split(strings.Trim(spec, "/"), "/")
	if len(parts) < 2 {
		return spec, ""
	}
	repo = parts[0] + "/" + parts[1]
	if len(parts) > 2 {
		subpath = strings.Join(parts[2:], "/")
	}
	return repo, subpath
}

// fetchRemote fetches a single skill from a GitHub repo.
//
// When entry.Repo is just "owner/name", hyoka shells out to
// "npx skills add <repo> --skill <name> ..." and consumes the resulting
// .claude/skills/<name>/ layout.
//
// When entry.Repo includes a subpath ("owner/name/sub/dir"), hyoka uses
// "git sparse-checkout" to fetch just that directory from the repo. The
// skill is expected at <subpath>/<name>/SKILL.md. This is the escape hatch
// for repos that don't publish through the skills CLI (e.g. skills living
// inside a monorepo's plugins/ tree).
//
// entry.Name is required in both cases and selects a single skill.
func fetchRemote(entry config.ToolEntry, baseDir string) (string, error) {
	if entry.Name == "" {
		return "", fmt.Errorf("remote skill from %q requires a 'name' field (the skill to install from the repo)", entry.Repo)
	}

	repo, subpath := parseRepoSpec(entry.Repo)
	if subpath != "" {
		return fetchRemoteSparse(entry, repo, subpath, baseDir)
	}
	return fetchRemoteNpx(entry, repo, baseDir)
}

// fetchRemoteNpx runs "npx skills add" to install a single skill from a
// GitHub repo. The skills CLI installs into <cwd>/.claude/skills/<skill-name>/.
// hyoka sets cmd.Dir to a per-repo cache directory so multiple configs don't
// clobber each other, then returns the absolute path to the installed skill.
func fetchRemoteNpx(entry config.ToolEntry, repo, baseDir string) (string, error) {
	// Per-repo cache dir; skills CLI will create .claude/skills/<name>/ inside it.
	cacheDir := filepath.Join(baseDir, ".skills-cache", repo)
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("creating skill cache dir: %w", err)
	}

	// Flags reference: https://www.npmjs.com/package/skills
	//   --skill <name>       select a single skill from the repo
	//   --agent claude-code  install layout hyoka consumes (.claude/skills/<name>/SKILL.md)
	//   --copy               copy files instead of symlinking (portable across machines)
	//   --yes                skip interactive confirmation prompts
	args := []string{"skills", "add", repo,
		"--skill", entry.Name,
		"--agent", "claude-code",
		"--copy",
		"--yes",
	}

	fmt.Printf("Fetching remote skill: %s (repo: %s)\n", entry.Name, repo)
	slog.Info("Fetching remote skill", "skill", entry.Name, "repo", repo, "cache_dir", cacheDir)
	cmd := exec.Command("npx", args...)
	cmd.Dir = cacheDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("npx skills add: %w", err)
	}

	installedDir := filepath.Join(cacheDir, ".claude", "skills", entry.Name)
	if _, err := os.Stat(filepath.Join(installedDir, "SKILL.md")); err != nil {
		return "", fmt.Errorf("skills add completed but SKILL.md not found at %s (skill %q may not exist in repo %q)", installedDir, entry.Name, repo)
	}

	abs, absErr := filepath.Abs(installedDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute install path", "path", installedDir, "error", absErr)
		return installedDir, nil
	}
	return abs, nil
}

// fetchRemoteSparse clones a GitHub repo with git sparse-checkout and returns
// the path to a specific skill inside a subpath. Used when entry.Repo includes
// a subpath (e.g. "microsoft/skills/.github/plugins/azure-sdk-rust/skills").
//
// The final skill dir is <cacheDir>/<subpath>/<entry.Name>/ and must contain
// SKILL.md. Subsequent runs reuse the existing clone and refresh via "git pull".
func fetchRemoteSparse(entry config.ToolEntry, repo, subpath, baseDir string) (string, error) {
	cacheDir := filepath.Join(baseDir, ".skills-cache", repo)
	skillDir := filepath.Join(cacheDir, subpath, entry.Name)
	cloneURL := fmt.Sprintf("https://github.com/%s.git", repo)

	fmt.Printf("Fetching remote skill: %s (repo: %s, subpath: %s)\n", entry.Name, repo, subpath)
	slog.Info("Fetching remote skill (sparse)", "skill", entry.Name, "repo", repo, "subpath", subpath, "cache_dir", cacheDir)

	if _, err := os.Stat(filepath.Join(cacheDir, ".git")); err != nil {
		// First-time clone: blobless, no checkout, then configure sparse-checkout.
		if err := os.MkdirAll(filepath.Dir(cacheDir), 0755); err != nil {
			return "", fmt.Errorf("creating skill cache dir: %w", err)
		}
		if err := runGit("", "clone", "--filter=blob:none", "--no-checkout", "--depth=1", cloneURL, cacheDir); err != nil {
			return "", fmt.Errorf("git clone %s: %w", repo, err)
		}
		// Non-cone mode: the subpath may start with "." (e.g. ".github/...").
		if err := runGit(cacheDir, "sparse-checkout", "init", "--no-cone"); err != nil {
			return "", fmt.Errorf("git sparse-checkout init: %w", err)
		}
	}

	// Always (re)apply the sparse pattern — cheap and lets multiple configs
	// layer different subpaths into the same clone without stomping each other.
	sparsePattern := "/" + subpath + "/"
	if err := runGit(cacheDir, "sparse-checkout", "add", sparsePattern); err != nil {
		// Fall back to `set` if `add` is unavailable (older git).
		if err := runGit(cacheDir, "sparse-checkout", "set", "--no-cone", sparsePattern); err != nil {
			return "", fmt.Errorf("git sparse-checkout add %s: %w", sparsePattern, err)
		}
	}
	if err := runGit(cacheDir, "checkout"); err != nil {
		return "", fmt.Errorf("git checkout: %w", err)
	}

	if _, err := os.Stat(filepath.Join(skillDir, "SKILL.md")); err != nil {
		return "", fmt.Errorf("sparse-checkout completed but SKILL.md not found at %s (skill %q may not exist under %q in repo %q)", skillDir, entry.Name, subpath, repo)
	}

	abs, absErr := filepath.Abs(skillDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute skill path", "path", skillDir, "error", absErr)
		return skillDir, nil
	}
	return abs, nil
}

// runGit runs `git <args>` with the given working directory (empty = inherit).
// stdout/stderr are streamed so clone progress and errors are visible.
func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
