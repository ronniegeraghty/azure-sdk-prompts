//go:build !windows

package tool

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
)

// repoLockFilename is the per-repo sentinel locked via flock(2). Lives in
// the parent of the version subdir so it covers all versions of one repo
// at once — concurrent hyoka processes serialize on clone/fetch/checkout
// regardless of which version they want.
const repoLockFilename = ".hyoka-lock"

// repoLockTimeout is the upper bound on how long acquireRepoLock will wait
// for the lock before giving up. 30s comfortably exceeds a typical clone of
// a small skill repo on broadband; longer than this strongly suggests a
// stuck or runaway peer process.
var repoLockTimeout = 30 * time.Second

// repoLockPoll is the gap between LOCK_NB retries while waiting. Tuned for
// fast handoff without burning CPU.
var repoLockPoll = 500 * time.Millisecond

// acquireRepoLock takes an exclusive flock on a sentinel file in parentDir,
// returning a release function the caller MUST invoke. Used to serialize
// concurrent hyoka processes that may otherwise race on the same git
// clone/fetch/checkout. Returns an error wrapping "another hyoka process is
// fetching this repo" if the lock can't be acquired within repoLockTimeout.
//
// The lock is per-repo (parentDir = <CacheRoot>/repos/<owner>/<repo>/), not
// per-version, so a fetch for owner/repo@v1 blocks a fetch for owner/repo@v2
// — the underlying .git/index would otherwise still be hot.
func acquireRepoLock(ctx context.Context, parentDir string) (func() error, error) {
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		return nil, fmt.Errorf("creating lock parent dir: %w", err)
	}
	lockPath := filepath.Join(parentDir, repoLockFilename)
	f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("opening lock file %s: %w", lockPath, err)
	}

	deadline := time.Now().Add(repoLockTimeout)
	for {
		flockErr := unix.Flock(int(f.Fd()), unix.LOCK_EX|unix.LOCK_NB)
		if flockErr == nil {
			slog.Debug("acquired repo lock", "path", lockPath)
			return func() error {
				_ = unix.Flock(int(f.Fd()), unix.LOCK_UN)
				return f.Close()
			}, nil
		}
		if !errors.Is(flockErr, unix.EWOULDBLOCK) {
			_ = f.Close()
			return nil, fmt.Errorf("flock %s: %w", lockPath, flockErr)
		}
		if time.Now().After(deadline) {
			_ = f.Close()
			return nil, fmt.Errorf("another hyoka process is fetching this repo (lock at %s held >%s)", lockPath, repoLockTimeout)
		}
		select {
		case <-ctx.Done():
			_ = f.Close()
			return nil, ctx.Err()
		case <-time.After(repoLockPoll):
		}
	}
}
