package tool

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
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
// It is preloaded with the built-in npx fetcher.
var DefaultRegistry = func() *Registry {
	r := NewRegistry()
	_ = r.Register(&npxFetcher{})
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

// --- built-in npx fetcher --------------------------------------------------

const defaultFetcherName = "npx"

// npxFetcher is the built-in Fetcher backed by `npx skills add`. It preserves
// the historical behavior of FetchRemote: caches under <baseDir>/.skills-cache/
// and shells out to npx. When a Version is supplied it is appended to the repo
// arg as `repo@version`, which `npx skills add` (and the underlying git fetch)
// interprets as a branch/tag/commit ref.
type npxFetcher struct{}

func (npxFetcher) Name() string { return defaultFetcherName }

// CanFetch returns true for any remote skill entry. The default fetcher is the
// last-resort handler; custom fetchers can pre-empt it via their own CanFetch.
func (npxFetcher) CanFetch(entry Entry) bool {
	return entry.ResolvedType() == TypeSkill && entry.SkillSource() == SourceRemote
}

func (npxFetcher) Fetch(ctx context.Context, req FetchRequest) (FetchResult, error) {
	entry := req.Entry
	// Version-aware install path: keeps different versions in distinct dirs
	// so switching the override doesn't poison the cache.
	versionSegment := req.Version
	if versionSegment == "" {
		versionSegment = "default"
	}
	installDir := filepath.Join(req.BaseDir, ".skills-cache", versionSegment, entry.Repo)
	if entry.Name != "" {
		installDir = filepath.Join(installDir, entry.Name)
	}
	if err := os.MkdirAll(installDir, 0o755); err != nil {
		return FetchResult{}, fmt.Errorf("creating skill install dir: %w", err)
	}

	repoArg := entry.Repo
	if req.Version != "" {
		repoArg = entry.Repo + "@" + req.Version
	}
	args := []string{"skills", "add", repoArg, "--directory", installDir}
	if entry.Name != "" {
		args = append(args, "--name", entry.Name)
	}

	slog.Info("Fetching remote skill", "skill", entry.Name, "repo", entry.Repo, "version", versionSegment, "fetcher", defaultFetcherName)

	cmd := exec.CommandContext(ctx, "npx", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return FetchResult{}, fmt.Errorf("npx skills add: %w", err)
	}

	abs, absErr := filepath.Abs(installDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute install path", "path", installDir, "error", absErr)
		abs = installDir
	}
	return FetchResult{Dir: abs, Version: versionSegment}, nil
}
