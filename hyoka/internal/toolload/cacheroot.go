// Package toolload owns the canonical cache layout for remote tools (skills,
// plugins) fetched and reused across hyoka eval runs.
//
// Today only CacheRoot lives here; subsequent items in the tool-load
// consolidation plan (Morpheus 2026-04-25) will move resolution and
// freshness logic into this package as well. Keeping the path computation
// in one spot eliminates the historical split where skills cached under
// "<configDir>/.skills-cache/..." (which was an empty string for the
// reviewer factory, polluting cwd) and plugins under
// "~/.hyoka/cache/default/...".
package toolload

import (
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

const (
	// EnvCacheDir overrides every other source. If set and non-empty, this
	// path is used as-is for the cache root.
	EnvCacheDir = "HYOKA_CACHE_DIR"
	// EnvXDGCacheHome is honored when EnvCacheDir is unset; the cache root
	// becomes "$XDG_CACHE_HOME/hyoka".
	EnvXDGCacheHome = "XDG_CACHE_HOME"
)

var (
	cacheRootOnce sync.Once
	cacheRootVal  string
	// For testing: track whether Once has been called
	cacheRootInitialized bool
)

// CacheRoot returns the canonical cache directory for hyoka. Resolution
// order (first non-empty wins):
//
//  1. $HYOKA_CACHE_DIR
//  2. $XDG_CACHE_HOME/hyoka
//  3. ~/.hyoka/cache
//
// If both env vars are unset and os.UserHomeDir fails, falls back to
// os.TempDir() and logs a warning. The value is computed once per process.
//
// On the first call, if a legacy ".skills-cache/" directory exists in the
// current working directory, a single warning is logged so users notice it
// is safe to remove. The directory is never auto-deleted.
func CacheRoot() string {
	cacheRootOnce.Do(func() {
		cacheRootVal = resolveCacheRoot()
		cacheRootInitialized = true
		warnIfLegacySkillsCache(cacheRootVal)
	})
	return cacheRootVal
}

func resolveCacheRoot() string {
	if v := os.Getenv(EnvCacheDir); v != "" {
		return v
	}
	if v := os.Getenv(EnvXDGCacheHome); v != "" {
		return filepath.Join(v, "hyoka")
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		fallback := os.TempDir()
		slog.Warn("toolload: could not resolve user home dir, falling back to temp",
			"error", err, "fallback", fallback)
		return filepath.Join(fallback, "hyoka", "cache")
	}
	return filepath.Join(home, ".hyoka", "cache")
}

func warnIfLegacySkillsCache(newRoot string) {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	legacy := filepath.Join(cwd, ".skills-cache")
	if info, err := os.Stat(legacy); err == nil && info.IsDir() {
		slog.Warn("toolload: legacy .skills-cache/ dir detected — safe to remove; cache moved",
			"path", legacy, "new_root", newRoot)
	}
}

// VersionSegment normalizes an entry's pinned version into the on-disk
// segment used in the cache layout. Empty or "default" → "default".
func VersionSegment(version string) string {
	if version == "" {
		return "default"
	}
	return version
}

// RepoCacheDir returns the canonical on-disk directory for a remote repo
// at a given version. Layout:
//
//	<CacheRoot>/repos/<owner>/<repo>/<version-segment>/
//
// Owner/repo first so all versions of a single repo cluster together on
// disk, which is friendlier to du/ls than version-first.
func RepoCacheDir(owner, repo, version string) string {
	return filepath.Join(CacheRoot(), "repos", owner, repo, VersionSegment(version))
}

// SetTestRoot overrides CacheRoot for tests and returns a cleanup func that
// restores the prior state. Test-only — must not be called from production
// code paths. Safe to use across packages because the sync.Once is reset
// inside the cleanup.
func SetTestRoot(path string) (restore func()) {
	prevVal := cacheRootVal
	prevInitialized := cacheRootInitialized
	// Reset the Once by replacing it with a fresh instance
	cacheRootOnce = sync.Once{}
	cacheRootVal = path
	cacheRootInitialized = false
	// Ensure the once is "done" so CacheRoot returns path verbatim.
	cacheRootOnce.Do(func() {
		// Already set, just mark initialized
		cacheRootInitialized = true
	})
	return func() {
		// Restore the previous value and state
		cacheRootVal = prevVal
		cacheRootInitialized = prevInitialized
		// Reset the Once again so the old state is restored
		cacheRootOnce = sync.Once{}
		if prevInitialized {
			// Re-arm the Once as "done" with the previous value
			cacheRootOnce.Do(func() {})
		}
	}
}

// resetForTest re-arms the sync.Once so tests can assert env-driven
// behavior. Test-only, package-internal.
func resetForTest() {
	cacheRootOnce = sync.Once{}
	cacheRootVal = ""
}
