package tool

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// FetchRequest carries the data a Fetcher needs to materialize a remote tool
// (typically a skill) on disk.
//
// Note: prior versions of this struct exposed a per-eval BaseDir field that
// callers used to scope the on-disk cache. That field was a footgun — empty
// strings from the reviewer factory caused ".skills-cache/" to land in the
// process cwd, and the per-eval tmp dir caused the generator's cache to be
// destroyed every run. Fetchers now derive their cache path from
// toolload.CacheRoot(), which is the single source of truth.
type FetchRequest struct {
	// Entry is the tool entry being fetched (already version-overridden).
	Entry Entry
	// Version is the resolved version pin. Empty string means "default" —
	// fetchers should treat this as latest / the user's normal default.
	Version string
}

// FetchResult is what a Fetcher returns once a tool is on disk.
type FetchResult struct {
	// Dir is the absolute path to the tool's installed directory (for skills,
	// the directory containing SKILL.md).
	Dir string
	// Version is the version actually installed. Defaults to "default" when
	// unspecified. Surfaced in logs so users can tell what they got.
	Version string
}

// Fetcher resolves a remote tool entry into a directory on disk. Hyoka ships
// with a default npx-based implementation; users can register additional
// fetchers via Register to add support for new sources, mirror caches, or pin
// versions in custom ways. The first fetcher whose CanFetch returns true (in
// registration order, defaults last) wins.
//
// The contract is intentionally narrow:
//   - Name must be unique within a process — used for diagnostics and to
//     detect double-registration.
//   - CanFetch must be a pure check on the entry; it must not perform I/O.
//   - Fetch must be safe to call concurrently for distinct entries; fetchers
//     that share on-disk state must serialize internally.
//   - Fetch must honor ctx cancellation when its work is non-trivial.
type Fetcher interface {
	Name() string
	CanFetch(entry Entry) bool
	Fetch(ctx context.Context, req FetchRequest) (FetchResult, error)
}

// Registry holds the active set of Fetchers. The zero value is not usable —
// always construct via NewRegistry. The default package-level registry
// (DefaultRegistry) is preloaded with the npx fetcher.
type Registry struct {
	mu       sync.RWMutex
	fetchers []Fetcher
	names    map[string]struct{}
}

// NewRegistry returns an empty registry. Most callers want DefaultRegistry.
func NewRegistry() *Registry {
	return &Registry{names: make(map[string]struct{})}
}

// Register adds a fetcher. Custom fetchers are consulted before the built-in
// default (which always lives at the tail of the list), so a user-registered
// fetcher can shadow the default when its CanFetch matches.
//
// Returns an error if a fetcher with the same Name is already registered.
func (r *Registry) Register(f Fetcher) error {
	if f == nil {
		return fmt.Errorf("nil fetcher")
	}
	name := f.Name()
	if name == "" {
		return fmt.Errorf("fetcher must have a non-empty Name")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.names[name]; ok {
		return fmt.Errorf("fetcher %q already registered", name)
	}
	r.names[name] = struct{}{}
	// Insert before any fetcher with name == defaultFetcherName so the default
	// remains last. This keeps lookup deterministic regardless of registration
	// order.
	insertAt := len(r.fetchers)
	for i, ex := range r.fetchers {
		if ex.Name() == defaultFetcherName {
			insertAt = i
			break
		}
	}
	r.fetchers = append(r.fetchers, nil)
	copy(r.fetchers[insertAt+1:], r.fetchers[insertAt:])
	r.fetchers[insertAt] = f
	return nil
}

// Unregister removes a fetcher by name. Primarily useful in tests.
func (r *Registry) Unregister(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.names, name)
	out := r.fetchers[:0]
	for _, f := range r.fetchers {
		if f.Name() != name {
			out = append(out, f)
		}
	}
	r.fetchers = out
}

// Lookup returns the first fetcher that CanFetch the entry, or nil if none
// matches (which should be impossible while the npx default is registered).
func (r *Registry) Lookup(entry Entry) Fetcher {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, f := range r.fetchers {
		if f.CanFetch(entry) {
			return f
		}
	}
	return nil
}

// Names returns registered fetcher names in lookup order. Used by diagnostics.
func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]string, len(r.fetchers))
	for i, f := range r.fetchers {
		out[i] = f.Name()
	}
	return out
}

// DefaultRegistry is the process-wide registry consulted by ResolveSkills.
// It is preloaded with the built-in git fetcher.
var DefaultRegistry = func() *Registry {
	r := NewRegistry()
	// pluginFetcher must register before gitFetcher so plugin entries route
	// to it. gitFetcher's CanFetch is skill-only, so order is safety-belt
	// rather than load-bearing — but we keep plugins first by intent.
	_ = r.Register(&pluginFetcher{})
	_ = r.Register(&gitFetcher{})
	return r
}()

// Register adds a fetcher to the default registry. Convenience wrapper around
// DefaultRegistry.Register for users who don't need their own registry.
func Register(f Fetcher) error { return DefaultRegistry.Register(f) }

// LookupFetcher returns the fetcher that would handle the given entry under
// the default registry.
func LookupFetcher(entry Entry) Fetcher { return DefaultRegistry.Lookup(entry) }

// ValidateFetchers checks that every remote skill entry in the given slice
// has a registered fetcher willing to handle it. Returns the first
// missing-fetcher error encountered. Intended to be called by the eval
// engine before any session starts so users get a fast, clear failure
// instead of a mid-eval surprise. Non-remote and non-skill entries are
// ignored.
func ValidateFetchers(entries []Entry) error {
	for _, e := range entries {
		if e.ResolvedType() != TypeSkill || e.SkillSource() != SourceRemote {
			continue
		}
		if LookupFetcher(e) == nil {
			return fmt.Errorf("no fetcher registered for remote skill %q (repo=%q)", e.Name, e.Repo)
		}
	}
	return nil
}

// --- built-in git fetcher --------------------------------------------------

const defaultFetcherName = "git"

// gitFetcher is the built-in Fetcher that clones Git repositories directly
// instead of shelling out to npx. It replaces the historical npxFetcher to
// avoid stdout pollution from npm plugin auto-install. Skill specs are parsed:
//   - "name@owner/repo" → clone owner/repo, look for skill "name"
//   - Bare "owner/repo" → clone repo, return root if no name specified
// Caches under <toolload.CacheRoot()>/repos/<owner>/<repo>/<version>/.
type gitFetcher struct{}

func (gitFetcher) Name() string { return defaultFetcherName }

// CanFetch returns true for any remote skill entry. The default fetcher is the
// last-resort handler; custom fetchers can pre-empt it via their own CanFetch.
func (gitFetcher) CanFetch(entry Entry) bool {
	return entry.ResolvedType() == TypeSkill && entry.SkillSource() == SourceRemote
}

func (gitFetcher) Fetch(ctx context.Context, req FetchRequest) (FetchResult, error) {
	entry := req.Entry
	versionSegment := toolload.VersionSegment(req.Version)

	// Parse skill spec from explicit repo + name fields.
	owner, repo, skillName, repoSubpath := parseSkillSpec(entry.Repo, entry.Name)

	// Cache path: <CacheRoot>/repos/<owner>/<repo>/<version>/. Owner/repo
	// first so all versions of one repo cluster together (Item A, Morpheus
	// 2026-04-25 plan).
	cacheDir := toolload.RepoCacheDir(owner, repo, versionSegment)

	// Clone or update the repo
	if err := ensureRepoCloned(ctx, owner, repo, versionSegment, cacheDir); err != nil {
		return FetchResult{}, fmt.Errorf("cloning %s/%s: %w", owner, repo, err)
	}

	// Explicit path: when set on a remote skill entry (either via `path:` or
	// embedded in the repo string as `<owner>/<repo>/<subpath>`), point
	// directly at the skill directory inside the repo. Skips the name-based
	// search across well-known locations (.github/skills, skills/, etc.).
	// An explicit `path:` field wins over a path embedded in `repo:`.
	subpath := entry.Path
	if subpath == "" {
		subpath = repoSubpath
	}
	if subpath != "" {
		skillDir := filepath.Join(cacheDir, filepath.FromSlash(subpath))
		if !isValidSkillDir(skillDir) {
			return FetchResult{}, fmt.Errorf("skill path %q in %s/%s does not contain SKILL.md or plugin.yaml", subpath, owner, repo)
		}
		abs, absErr := filepath.Abs(skillDir)
		if absErr != nil {
			slog.Warn("Failed to resolve absolute skill path", "path", skillDir, "error", absErr)
			abs = skillDir
		}
		slog.Info("Resolved skill via explicit path", "skill", entry.Name, "repo", entry.Repo, "subpath", subpath, "version", versionSegment, "path", abs)
		return FetchResult{Dir: abs, Version: versionSegment}, nil
	}

	// If no specific skill name, return the repo root
	if skillName == "" {
		abs, absErr := filepath.Abs(cacheDir)
		if absErr != nil {
			slog.Warn("Failed to resolve absolute path", "path", cacheDir, "error", absErr)
			abs = cacheDir
		}
		slog.Info("Resolved skill repo", "repo", entry.Repo, "version", versionSegment, "path", abs)
		return FetchResult{Dir: abs, Version: versionSegment}, nil
	}

	// Search for the named skill in common locations
	skillDir, err := findSkillInRepo(cacheDir, skillName)
	if err != nil {
		return FetchResult{}, fmt.Errorf("finding skill %q in %s/%s: %w", skillName, owner, repo, err)
	}

	abs, absErr := filepath.Abs(skillDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute skill path", "path", skillDir, "error", absErr)
		abs = skillDir
	}
	slog.Info("Resolved skill", "skill", skillName, "repo", entry.Repo, "version", versionSegment, "path", abs)
	return FetchResult{Dir: abs, Version: versionSegment}, nil
}

// parseSkillSpec parses skill specifications:
//   - If name contains "@", split at last @ to get skillname and owner/repo
//   - Otherwise, repo is "owner/repo" (or "github.com/owner/repo") and name is the skill name
//   - When repo has more than two path segments (e.g. "owner/repo/skills/python"),
//     the first two are owner/repo and the remainder is returned as subpath.
//
// The legacy "name@skills" shorthand for microsoft/skills has been removed.
// Callers must declare the source repo explicitly via the entry's repo: field.
func parseSkillSpec(repo, name string) (owner, repoName, skillName, subpath string) {
	// Handle "name@owner/repo" format
	if idx := strings.LastIndex(name, "@"); idx > 0 {
		skillName = name[:idx]
		ownerRepo := name[idx+1:]
		parts := strings.SplitN(ownerRepo, "/", 2)
		if len(parts) == 2 {
			return parts[0], parts[1], skillName, ""
		}
		// Malformed — return owner empty so it fails downstream with a clear path.
		return ownerRepo, "", skillName, ""
	}

	// Standard "owner/repo[/subpath...]" with separate name. Strip optional github.com/ prefix.
	r := strings.TrimPrefix(repo, "github.com/")
	parts := strings.SplitN(r, "/", 3)
	if len(parts) < 2 {
		// Malformed repo — return as-is, will fail downstream
		return r, "", name, ""
	}
	if len(parts) == 3 {
		return parts[0], parts[1], name, parts[2]
	}
	return parts[0], parts[1], name, ""
}

// runGit is the package-level hook used by ensureRepoCloned. Tests swap it
// out to assert call counts / arg patterns without shelling out. The default
// is runGitQuiet which actually executes git.
var runGit = runGitQuiet

// ensureRepoCloned ensures the repo is cloned at cacheDir and checked out at
// the requested version. Behavior branches on whether the version is pinned:
//
//   - Fresh (no .git): git clone + checkout. No version freshness logic.
//   - Cached + unpinned ("default"/empty): always git fetch + checkout HEAD,
//     so users iterating on a remote skill see updates immediately. (Ronnie
//     directive 2026-04-27, no TTL.)
//   - Cached + pinned (any other version): try `git rev-parse <ref>^{commit}`
//     locally first. If it resolves, skip the network and just checkout. If
//     it doesn't, fetch then checkout (the pin may be a tag/commit added to
//     the remote since we cloned).
//
// The whole operation is serialized per <owner>/<repo> via a flock on the
// parent dir so concurrent hyoka processes don't race on the same clone
// across all version subdirs. See acquireRepoLock for timeout details.
//
// All git output is suppressed unless a command fails.
func ensureRepoCloned(ctx context.Context, owner, repo, version, cacheDir string) error {
	parentDir := filepath.Dir(cacheDir)
	release, lockErr := acquireRepoLock(ctx, parentDir)
	if lockErr != nil {
		return fmt.Errorf("acquiring repo lock for %s/%s: %w", owner, repo, lockErr)
	}
	defer func() {
		if cerr := release(); cerr != nil {
			slog.Debug("releasing repo lock", "owner", owner, "repo", repo, "error", cerr)
		}
	}()

	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)

	if _, statErr := os.Stat(filepath.Join(cacheDir, ".git")); statErr != nil {
		// Fresh clone
		slog.Debug("Cloning repo", "owner", owner, "repo", repo, "version", version, "url", repoURL)
		if err := os.MkdirAll(parentDir, 0o755); err != nil {
			return fmt.Errorf("creating cache dir: %w", err)
		}
		cloneArgs := []string{"clone", "--quiet"}
		if version != "default" && version != "" {
			cloneArgs = append(cloneArgs, "--branch", version)
		}
		cloneArgs = append(cloneArgs, repoURL, cacheDir)
		if err := runGit(ctx, "", cloneArgs...); err != nil {
			return fmt.Errorf("git clone: %w", err)
		}
		return nil
	}

	// Cached repo already on disk.
	pinned := version != "" && version != "default"
	if !pinned {
		// Unpinned — keep current always-fetch behavior.
		slog.Debug("Repo cache hit (unpinned), fetching", "owner", owner, "repo", repo)
		if err := runGit(ctx, cacheDir, "fetch", "--all", "--tags"); err != nil {
			return fmt.Errorf("git fetch: %w", err)
		}
		if err := runGit(ctx, cacheDir, "checkout", "HEAD"); err != nil {
			return fmt.Errorf("git checkout HEAD: %w", err)
		}
		return nil
	}

	// Pinned + cached: try to resolve the ref locally first. If it resolves,
	// skip the network entirely — pinned refs don't move.
	if err := runGit(ctx, cacheDir, "rev-parse", "--verify", "--quiet", version+"^{commit}"); err == nil {
		slog.Debug("Repo cache hit (pinned, ref resolves locally), skipping fetch", "owner", owner, "repo", repo, "version", version)
		if err := runGit(ctx, cacheDir, "checkout", version); err != nil {
			return fmt.Errorf("git checkout %s: %w", version, err)
		}
		return nil
	}

	// Pinned but ref isn't in the local clone yet — fetch and try again.
	slog.Debug("Repo cache hit (pinned, ref missing), fetching", "owner", owner, "repo", repo, "version", version)
	if err := runGit(ctx, cacheDir, "fetch", "--all", "--tags"); err != nil {
		return fmt.Errorf("git fetch: %w", err)
	}
	if err := runGit(ctx, cacheDir, "checkout", version); err != nil {
		return fmt.Errorf("git checkout %s: %w", version, err)
	}
	return nil
}

// runGitQuiet runs a git command with stdout/stderr captured. Only surfaces
// stderr if the command exits non-zero, and even then only logs it at Debug.
func runGitQuiet(ctx context.Context, dir string, args ...string) error {
	cmd := exec.CommandContext(ctx, "git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			slog.Debug("git command failed", "args", args, "stderr", stderr.String())
		}
		return err
	}
	return nil
}

// findSkillInRepo searches for a named skill in common locations within the
// repo. Search order:
//  1. .github/skills/<name>/
//  2. .github/plugins/<name>/
//  3. .claude/skills/<name>/
//  4. .agent/skills/<name>/
//  5. skills/<name>/
//  6. Top-level plugin.yaml or marketplace.yaml pointing at custom dir
//  7. If only one skill dir exists, use it
// Returns the absolute path to the skill directory (containing SKILL.md or
// plugin.yaml).
func findSkillInRepo(repoDir, skillName string) (string, error) {
	candidates := []string{
		filepath.Join(repoDir, ".github", "skills", skillName),
		filepath.Join(repoDir, ".github", "plugins", skillName),
		filepath.Join(repoDir, ".claude", "skills", skillName),
		filepath.Join(repoDir, ".agent", "skills", skillName),
		filepath.Join(repoDir, "skills", skillName),
	}

	for _, dir := range candidates {
		if isValidSkillDir(dir) {
			return dir, nil
		}
	}

	// Check for top-level plugin.yaml or marketplace.yaml with custom skills_dir
	if customDir := checkTopLevelPluginManifest(repoDir); customDir != "" {
		skillDir := filepath.Join(repoDir, customDir, skillName)
		if isValidSkillDir(skillDir) {
			return skillDir, nil
		}
	}

	// Last resort: if there's exactly one valid skill directory, use it
	if singleSkill := findSingleSkill(repoDir); singleSkill != "" {
		return singleSkill, nil
	}

	return "", fmt.Errorf("skill %q not found in any standard location", skillName)
}

// isValidSkillDir checks if a directory contains SKILL.md or plugin.yaml
func isValidSkillDir(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "SKILL.md")); err == nil {
		return true
	}
	if _, err := os.Stat(filepath.Join(dir, "plugin.yaml")); err == nil {
		return true
	}
	return false
}

// checkTopLevelPluginManifest looks for plugin.yaml or marketplace.yaml at
// the repo root and extracts a custom skills_dir if present. Returns empty
// string if none found or no custom directory specified.
func checkTopLevelPluginManifest(repoDir string) string {
	// TODO: if needed, parse YAML and extract skills_dir field
	// For now, return empty — this is a fallback path rarely used
	return ""
}

// findSingleSkill scans the repo for skill directories. If exactly one valid
// skill directory exists, returns it. Otherwise returns empty string.
func findSingleSkill(repoDir string) string {
	var found []string
	searchDirs := []string{
		filepath.Join(repoDir, ".github", "skills"),
		filepath.Join(repoDir, ".github", "plugins"),
		filepath.Join(repoDir, ".claude", "skills"),
		filepath.Join(repoDir, ".agent", "skills"),
		filepath.Join(repoDir, "skills"),
	}

	for _, searchDir := range searchDirs {
		entries, err := os.ReadDir(searchDir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			skillDir := filepath.Join(searchDir, e.Name())
			if isValidSkillDir(skillDir) {
				found = append(found, skillDir)
			}
		}
	}

	if len(found) == 1 {
		slog.Debug("Single skill auto-selected", "path", found[0])
		return found[0]
	}
	return ""
}
