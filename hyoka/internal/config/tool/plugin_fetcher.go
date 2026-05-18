package tool

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// pluginFetcherName is the registry name for the built-in plugin fetcher.
const pluginFetcherName = "plugin-git"

// pluginCloneFn is the clone helper used by the plugin fetcher. It is a
// package-level var so tests can substitute a stub that pre-populates the
// cache dir without performing any real network I/O. Production code uses
// ensureRepoCloned (defined in fetcher.go), which Tank's Item C extends with
// pinned-vs-default freshness logic and per-repo flock — the plugin fetcher
// rides those changes for free.
var pluginCloneFn = ensureRepoCloned

// pluginFetcher is the built-in Fetcher for `type: plugin, source: remote`
// entries. It git-clones the repo into the canonical
// <CacheRoot>/repos/<owner>/<repo>/<version>/ tree and locates the plugin
// directory using plugin.FindPluginInRepo (single source of truth shared with
// plugin.ResolveInstalled and pluginCheckedPaths — see Item F).
type pluginFetcher struct{}

func (pluginFetcher) Name() string { return pluginFetcherName }

// CanFetch returns true for remote plugin entries. The default registry
// consults this fetcher before the gitFetcher (which only handles skills),
// so plugin entries always land here.
func (pluginFetcher) CanFetch(entry Entry) bool {
	if entry.ResolvedType() != TypePlugin {
		return false
	}
	// Explicit remote, OR inferred remote (no source set + repo present).
	if entry.Source == SourceRemote {
		return true
	}
	if entry.Source == "" && entry.Repo != "" {
		return true
	}
	return false
}

func (pluginFetcher) Fetch(ctx context.Context, req FetchRequest) (FetchResult, error) {
	entry := req.Entry
	if entry.Repo == "" {
		return FetchResult{}, fmt.Errorf("remote plugin %q has no repo", entry.Name)
	}
	owner, repoName := plugin.SplitOwnerRepo(entry.Repo)
	if owner == "" || repoName == "" {
		return FetchResult{}, fmt.Errorf("remote plugin %q: cannot parse repo %q (want owner/repo)", entry.Name, entry.Repo)
	}

	versionSegment := toolload.VersionSegment(req.Version)
	repoDir := toolload.RepoCacheDir(owner, repoName, versionSegment)

	if err := pluginCloneFn(ctx, owner, repoName, versionSegment, repoDir); err != nil {
		return FetchResult{}, fmt.Errorf("cloning %s/%s: %w", owner, repoName, err)
	}

	pluginDir, err := plugin.FindPluginInRepo(repoDir, entry.Name)
	if err != nil {
		return FetchResult{}, fmt.Errorf("locating plugin %q in %s/%s: %w", entry.Name, owner, repoName, err)
	}

	abs, absErr := filepath.Abs(pluginDir)
	if absErr != nil {
		slog.Warn("Failed to resolve absolute plugin path", "path", pluginDir, "error", absErr)
		abs = pluginDir
	}
	slog.Info("Resolved remote plugin", "plugin", entry.Name, "repo", entry.Repo, "version", versionSegment, "path", abs)
	return FetchResult{Dir: abs, Version: versionSegment}, nil
}

// parsePluginRepo is retained as a thin shim over plugin.SplitOwnerRepo so
// existing tests in this package continue to exercise the same parser.
// New callers should use plugin.SplitOwnerRepo directly.
func parsePluginRepo(repo string) (owner, name string) {
	return plugin.SplitOwnerRepo(repo)
}

// findPluginInRepo is a backwards-compatible shim for in-package tests.
// Production code calls plugin.FindPluginInRepo directly.
func findPluginInRepo(repoRoot, name string) (string, error) {
	return plugin.FindPluginInRepo(repoRoot, name)
}
