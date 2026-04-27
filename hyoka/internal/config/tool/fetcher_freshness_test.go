package tool

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// gitRecorder is a runGit hook that captures every invocation and lets the
// test stub success/failure for specific subcommands. Used to verify the
// pinned-vs-unpinned branching in ensureRepoCloned without shelling out to
// real git.
type gitRecorder struct {
	mu    sync.Mutex
	calls [][]string
	// returns is keyed by the first arg ("fetch", "checkout", "rev-parse", ...)
	// — when set, that subcommand returns the mapped error. nil = success.
	returns map[string]error
}

func (g *gitRecorder) hook() func(ctx context.Context, dir string, args ...string) error {
	return func(_ context.Context, dir string, args ...string) error {
		g.mu.Lock()
		defer g.mu.Unlock()
		full := append([]string{dir}, args...)
		g.calls = append(g.calls, full)
		if len(args) == 0 {
			return nil
		}
		sub := args[0]
		// "clone" is special — first positional, no subcommand prefix.
		if err, ok := g.returns[sub]; ok {
			return err
		}
		return nil
	}
}

func (g *gitRecorder) countSubcommand(sub string) int {
	g.mu.Lock()
	defer g.mu.Unlock()
	n := 0
	for _, c := range g.calls {
		// c[0] is dir; c[1] is the git subcommand we passed.
		if len(c) >= 2 && c[1] == sub {
			n++
		}
	}
	return n
}

// withGitHook swaps runGit for the duration of the test and restores it.
func withGitHook(t *testing.T, hook func(ctx context.Context, dir string, args ...string) error) {
	t.Helper()
	orig := runGit
	runGit = hook
	t.Cleanup(func() { runGit = orig })
}

// preExistingRepo writes a fake .git directory so ensureRepoCloned takes the
// "cached" branch without us actually initializing a real repo. The runGit
// hook short-circuits any real git invocation.
func preExistingRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".git"), 0o755); err != nil {
		t.Fatalf("seed .git: %v", err)
	}
	return dir
}

func TestEnsureRepoCloned_PinnedCachedRefResolves_NoFetch(t *testing.T) {
	rec := &gitRecorder{returns: map[string]error{
		// rev-parse succeeds → ref is local → no fetch should happen.
	}}
	withGitHook(t, rec.hook())

	cacheDir := preExistingRepo(t)
	if err := ensureRepoCloned(context.Background(), "acme", "widgets", "v1.2.3", cacheDir); err != nil {
		t.Fatalf("ensureRepoCloned: %v", err)
	}
	if got := rec.countSubcommand("fetch"); got != 0 {
		t.Errorf("expected 0 fetch calls when pinned ref resolves locally, got %d", got)
	}
	if got := rec.countSubcommand("rev-parse"); got != 1 {
		t.Errorf("expected 1 rev-parse, got %d", got)
	}
	if got := rec.countSubcommand("checkout"); got != 1 {
		t.Errorf("expected 1 checkout, got %d", got)
	}
}

func TestEnsureRepoCloned_PinnedCachedRefMissing_OneFetch(t *testing.T) {
	rec := &gitRecorder{returns: map[string]error{
		"rev-parse": errors.New("unknown revision"),
	}}
	withGitHook(t, rec.hook())

	cacheDir := preExistingRepo(t)
	if err := ensureRepoCloned(context.Background(), "acme", "widgets", "v9.9.9", cacheDir); err != nil {
		t.Fatalf("ensureRepoCloned: %v", err)
	}
	if got := rec.countSubcommand("fetch"); got != 1 {
		t.Errorf("expected exactly 1 fetch when pinned ref missing locally, got %d", got)
	}
	if got := rec.countSubcommand("checkout"); got != 1 {
		t.Errorf("expected 1 checkout, got %d", got)
	}
}

func TestEnsureRepoCloned_UnpinnedCached_AlwaysFetches(t *testing.T) {
	for _, version := range []string{"", "default"} {
		t.Run("version="+version, func(t *testing.T) {
			rec := &gitRecorder{}
			withGitHook(t, rec.hook())

			cacheDir := preExistingRepo(t)
			if err := ensureRepoCloned(context.Background(), "acme", "widgets", version, cacheDir); err != nil {
				t.Fatalf("ensureRepoCloned: %v", err)
			}
			if got := rec.countSubcommand("fetch"); got != 1 {
				t.Errorf("unpinned must always fetch; got %d fetches", got)
			}
			if got := rec.countSubcommand("rev-parse"); got != 0 {
				t.Errorf("unpinned must NOT rev-parse; got %d", got)
			}
			if got := rec.countSubcommand("checkout"); got != 1 {
				t.Errorf("expected 1 checkout, got %d", got)
			}
		})
	}
}

func TestEnsureRepoCloned_FreshClone_CallsClone(t *testing.T) {
	rec := &gitRecorder{}
	withGitHook(t, rec.hook())

	cacheDir := filepath.Join(t.TempDir(), "owner", "repo", "v1.0.0")
	if err := ensureRepoCloned(context.Background(), "acme", "widgets", "v1.0.0", cacheDir); err != nil {
		t.Fatalf("ensureRepoCloned: %v", err)
	}
	if got := rec.countSubcommand("clone"); got != 1 {
		t.Errorf("expected 1 clone, got %d", got)
	}
	if got := rec.countSubcommand("fetch"); got != 0 {
		t.Errorf("fresh clone should not also fetch; got %d", got)
	}
}

// TestAcquireRepoLock_SerializesConcurrentAccess spawns two goroutines that
// each acquire the lock, write a timestamp to a shared file with a small
// hold, and assert their critical sections did not overlap. Real flock — no
// mocks.
func TestAcquireRepoLock_SerializesConcurrentAccess(t *testing.T) {
	parent := t.TempDir()

	type span struct {
		start, end time.Time
	}
	var spans [2]span
	const hold = 100 * time.Millisecond

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		i := i
		go func() {
			defer wg.Done()
			release, err := acquireRepoLock(context.Background(), parent)
			if err != nil {
				t.Errorf("goroutine %d: acquire: %v", i, err)
				return
			}
			spans[i].start = time.Now()
			time.Sleep(hold)
			spans[i].end = time.Now()
			if err := release(); err != nil {
				t.Errorf("goroutine %d: release: %v", i, err)
			}
		}()
	}
	wg.Wait()

	// Critical sections must not overlap. Order may vary; check both ways.
	overlap := func(a, b span) bool {
		return a.start.Before(b.end) && b.start.Before(a.end)
	}
	if overlap(spans[0], spans[1]) {
		t.Errorf("critical sections overlapped: g0=[%s,%s] g1=[%s,%s]",
			spans[0].start, spans[0].end, spans[1].start, spans[1].end)
	}
}

// TestAcquireRepoLock_TimeoutReportsBusy shrinks the timeout, holds the lock
// in a background goroutine, and asserts the second acquirer fails with the
// "another hyoka process" error.
func TestAcquireRepoLock_TimeoutReportsBusy(t *testing.T) {
	origTimeout, origPoll := repoLockTimeout, repoLockPoll
	repoLockTimeout = 200 * time.Millisecond
	repoLockPoll = 25 * time.Millisecond
	t.Cleanup(func() {
		repoLockTimeout = origTimeout
		repoLockPoll = origPoll
	})

	parent := t.TempDir()
	release, err := acquireRepoLock(context.Background(), parent)
	if err != nil {
		t.Fatalf("first acquire: %v", err)
	}

	_, err = acquireRepoLock(context.Background(), parent)
	if err == nil {
		t.Fatal("second acquire should have timed out")
	}
	if !strings.Contains(err.Error(), "another hyoka process is fetching this repo") {
		t.Errorf("unexpected error: %v", err)
	}
	if err := release(); err != nil {
		t.Errorf("release: %v", err)
	}
}

// TestEnsureRepoCloned_FlockSerializesConcurrentEnsure verifies that the
// flock inside ensureRepoCloned actually serializes — two goroutines hitting
// the same cacheDir must not run their git ops concurrently. The hook
// records entry/exit timestamps.
func TestEnsureRepoCloned_FlockSerializesConcurrentEnsure(t *testing.T) {
	cacheDir := preExistingRepo(t)

	type span struct {
		start, end time.Time
	}
	var (
		mu    sync.Mutex
		spans []span
	)
	const work = 50 * time.Millisecond

	hook := func(_ context.Context, _ string, args ...string) error {
		// Mark time only on the rev-parse — that's the first git op inside
		// the lock for the pinned-cached path. We add a sleep so concurrent
		// holders would visibly overlap if the flock were broken.
		if len(args) > 0 && args[0] == "rev-parse" {
			s := span{start: time.Now()}
			time.Sleep(work)
			s.end = time.Now()
			mu.Lock()
			spans = append(spans, s)
			mu.Unlock()
		}
		return nil
	}
	withGitHook(t, hook)

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			if err := ensureRepoCloned(context.Background(), "acme", "widgets", "v1.0.0", cacheDir); err != nil {
				t.Errorf("ensureRepoCloned: %v", err)
			}
		}()
	}
	wg.Wait()

	if len(spans) != 2 {
		t.Fatalf("expected 2 critical sections, got %d", len(spans))
	}
	if spans[0].start.Before(spans[1].end) && spans[1].start.Before(spans[0].end) {
		t.Errorf("critical sections overlapped: %+v", spans)
	}
}

