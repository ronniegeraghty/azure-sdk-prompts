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
)

// FetchRequest carries the data a Fetcher needs to materialize a remote tool
// (typically a skill) on disk.
type FetchRequest struct {
	// Entry is the tool entry being fetched (already version-overridden).
	Entry Entry
	// BaseDir is the per-eval working directory; fetchers should keep their
	// per-tool cache scoped beneath this directory unless they have a strong
	// reason to use a global cache.
	BaseDir string
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
// Caches under <baseDir>/.skills-cache/<version>/<owner>/<repo>/.
type gitFetcher struct{}

func (gitFetcher) Name() string { return defaultFetcherName }

// CanFetch returns true for any remote skill entry. The default fetcher is the
// last-resort handler; custom fetchers can pre-empt it via their own CanFetch.
func (gitFetcher) CanFetch(entry Entry) bool {
	return entry.ResolvedType() == TypeSkill && entry.SkillSource() == SourceRemote
}

func (gitFetcher) Fetch(ctx context.Context, req FetchRequest) (FetchResult, error) {
	entry := req.Entry
	versionSegment := req.Version
	if versionSegment == "" {
		versionSegment = "default"
	}

	// Parse skill spec from explicit repo + name fields.
	owner, repo, skillName := parseSkillSpec(entry.Repo, entry.Name)

	// Cache path: <baseDir>/.skills-cache/<version>/<owner>/<repo>/
	cacheDir := filepath.Join(req.BaseDir, ".skills-cache", versionSegment, owner, repo)

	// Clone or update the repo
	if err := ensureRepoCloned(ctx, owner, repo, versionSegment, cacheDir); err != nil {
		return FetchResult{}, fmt.Errorf("cloning %s/%s: %w", owner, repo, err)
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
//
// The legacy "name@skills" shorthand for microsoft/skills has been removed.
// Callers must declare the source repo explicitly via the entry's repo: field.
func parseSkillSpec(repo, name string) (owner, repoName, skillName string) {
	// Handle "name@owner/repo" format
	if idx := strings.LastIndex(name, "@"); idx > 0 {
		skillName = name[:idx]
		ownerRepo := name[idx+1:]
		parts := strings.SplitN(ownerRepo, "/", 2)
		if len(parts) == 2 {
			return parts[0], parts[1], skillName
		}
		// Malformed — return owner empty so it fails downstream with a clear path.
		return ownerRepo, "", skillName
	}

	// Standard "owner/repo" with separate name. Strip optional github.com/ prefix.
	r := strings.TrimPrefix(repo, "github.com/")
	parts := strings.SplitN(r, "/", 2)
	if len(parts) != 2 {
		// Malformed repo — return as-is, will fail downstream
		return r, "", name
	}
	return parts[0], parts[1], name
}

// ensureRepoCloned ensures the repo is cloned at cacheDir. If it already
// exists, runs git fetch and checks out the specified version. All git output
// is suppressed unless the command fails.
func ensureRepoCloned(ctx context.Context, owner, repo, version, cacheDir string) error {
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)

	_, statErr := os.Stat(filepath.Join(cacheDir, ".git"))
	if statErr == nil {
		// Repo already exists — fetch and checkout
		slog.Debug("Repo cache hit, updating", "owner", owner, "repo", repo, "version", version)
		if err := runGitQuiet(ctx, cacheDir, "fetch", "--all", "--tags"); err != nil {
			return fmt.Errorf("git fetch: %w", err)
		}
		// Checkout the version (default branch if version == "default")
		ref := version
		if version == "default" {
			ref = "HEAD"
		}
		if err := runGitQuiet(ctx, cacheDir, "checkout", ref); err != nil {
			return fmt.Errorf("git checkout %s: %w", ref, err)
		}
		return nil
	}

	// Fresh clone
	slog.Debug("Cloning repo", "owner", owner, "repo", repo, "version", version, "url", repoURL)
	parentDir := filepath.Dir(cacheDir)
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		return fmt.Errorf("creating cache dir: %w", err)
	}

	cloneArgs := []string{"clone", "--quiet"}
	if version != "default" {
		cloneArgs = append(cloneArgs, "--branch", version)
	}
	cloneArgs = append(cloneArgs, repoURL, cacheDir)

	if err := runGitQuiet(ctx, "", cloneArgs...); err != nil {
		return fmt.Errorf("git clone: %w", err)
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
