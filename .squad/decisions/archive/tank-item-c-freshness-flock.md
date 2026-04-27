# Tank Item C — Version freshness (pinned-vs-latest) + per-repo flock

**Author:** Tank 📡
**Date:** 2026-04-27
**Status:** ✅ Implemented — awaiting Ronnie sign-off
**Spec:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md` § Item C
**Directive honored:** `.squad/decisions/inbox/copilot-directive-2026-04-27T20-10Z-no-ttl.md` (no TTL)

## What shipped

1. **`ensureRepoCloned` branches on pinned vs unpinned** (`hyoka/internal/config/tool/fetcher.go`):
   - **Fresh (no `.git`):** `git clone` (+ `--branch <version>` if pinned). No fetch. Unchanged.
   - **Cached + unpinned (`""`/`default`):** `git fetch --all --tags` + `git checkout HEAD` every call. Preserves pre-existing behavior so users iterating on remote skills see updates immediately.
   - **Cached + pinned (any other version):** `git rev-parse --verify --quiet <version>^{commit}` first. If it resolves locally → just `git checkout <version>` (zero network). If it doesn't → `git fetch --all --tags` then checkout (the pin may be a tag/commit added since last clone).

2. **Per-repo flock** (`hyoka/internal/config/tool/flock_unix.go`, `flock_windows.go`):
   - `acquireRepoLock(ctx, parentDir)` opens `<parentDir>/.hyoka-lock`, holds `LOCK_EX | LOCK_NB`, polls every 500ms, gives up after 30s with a wrapped `"another hyoka process is fetching this repo (lock at <path> held >30s)"` error.
   - Locked at the **parent of the version dir** = `<CacheRoot>/repos/<owner>/<repo>/`, so all version subdirs of the same repo serialize against each other (per Tank Item A § "What Item C should know").
   - Acquired before stat/clone/fetch/checkout, released via deferred close in `ensureRepoCloned`.
   - Windows: no-op stub (`flock_windows.go`). Build-tagged. A real `LockFileEx` impl can land later if Windows usage warrants it.
   - `unix.EWOULDBLOCK` is the busy signal; any other `flock` error returns immediately.
   - `repoLockTimeout` and `repoLockPoll` are package vars so timeout-path tests run in <1s.

3. **Tests** (`hyoka/internal/config/tool/fetcher_freshness_test.go`):
   - `TestEnsureRepoCloned_PinnedCachedRefResolves_NoFetch` — 0 fetches, 1 rev-parse, 1 checkout.
   - `TestEnsureRepoCloned_PinnedCachedRefMissing_OneFetch` — rev-parse fails → 1 fetch + 1 checkout.
   - `TestEnsureRepoCloned_UnpinnedCached_AlwaysFetches` — both `""` and `"default"` always fetch.
   - `TestEnsureRepoCloned_FreshClone_CallsClone` — 1 clone, 0 fetches.
   - `TestAcquireRepoLock_SerializesConcurrentAccess` — two real flock holders, asserts non-overlapping critical sections.
   - `TestAcquireRepoLock_TimeoutReportsBusy` — timeout shrunk to 200ms; second acquirer fails with the expected error message.
   - `TestEnsureRepoCloned_FlockSerializesConcurrentEnsure` — end-to-end: two goroutines into `ensureRepoCloned` on same dir, hook adds a 50ms sleep on rev-parse, asserts non-overlap.

## Decisions (the open questions)

### D1 — Testability hook: package-level `runGit` var (mock approach), not file:// fixture
**Choice:** `var runGit = runGitQuiet` swapped per-test via `withGitHook(t, ...)`.
**Why:** A real local-bare-repo fixture (init, tag, `file://` URL) can exercise the real `runGitQuiet`, but it's slower (~hundreds of ms per case to spawn git), brittle on systems without git, and doesn't actually let us count "did we hit the network." A function var is one line, race-safe per test (each test owns the swap via `t.Cleanup`), and lets us cleanly assert call counts and inject failure for `rev-parse` to drive both branches. The flock test uses real flock + `preExistingRepo` (a fake `.git` dir is enough — the hook short-circuits actual git invocation).
**Risk:** A future reader might be tempted to abuse `runGit` from outside the package. It's lowercase + only this file uses it; if it grows, refactor to an interface. Acceptable for now.

### D2 — Flock impl: `golang.org/x/sys/unix.Flock` on POSIX, no-op on Windows
**Choice:** `LOCK_EX | LOCK_NB` polled every 500ms with a 30s deadline.
**Why:** `x/sys` was already an indirect dep, no new module surface. `Flock` on a sentinel file (`.hyoka-lock`) is the standard pattern; it's released automatically on process exit (kernel cleans up), which means a crashed hyoka can't leave a permanent lock turd. Polling instead of blocking lets us honor `ctx.Done()` and surface a friendly timeout error rather than hanging forever.
**Why not `syscall.Flock`:** On Windows there's no equivalent in `syscall`, and `x/sys/unix` keeps the build tag clean.
**Why parent dir, not version dir:** Per Ronnie's plan + Tank Item A note — covers all versions of one repo at once. Two `hyoka run`s targeting the same repo at different pins would otherwise still race on the underlying `.git/index` of sibling clones if they're sharing objects; locking at the parent eliminates the surface entirely.

### D3 — Meta file (`<CacheRoot>/meta/<owner>/<repo>.json`): **deferred**
**Choice:** Not shipping in this PR.
**Why:** Item C's diff is already ~250 LoC across 3 new files + ensureRepoCloned rewrite + 7 tests. Adding meta-file write/read with its own JSON marshaling, atomic rename, and tests would push this past the "clean reviewable change" line and overlap with Item E (post-session verification gate, Neo). Better as a follow-up where the consumer (e.g., `hyoka serve`'s "last fetched at" column) lands at the same time and justifies the format.
**Follow-up tracking:** Open a new decision note `tank-item-c-meta-deferred.md` if/when picked up. Not creating a stub now to avoid inbox noise — this section is the record.

### D4 — Branches as "pinned": treat any non-default/non-empty version as pinned
**Per spec.** Even branch names (`main`, `develop`) skip the fetch when the local clone has the ref. Trade-off: a user who pins to a branch and expects "always latest" won't get it. The escape hatch is `default` / unset — that's documented behavior. If users hit this, tighten `isPinned` to "looks like a tag/SHA" later.

### D5 — `ensureRepoCloned` still callable with empty version
**Today's callers** (gitFetcher → `versionSegment` → "default" or actual) never pass an empty string at runtime, but the function defensively treats `""` the same as `"default"` (unpinned, always-fetch). Keeps the contract symmetric with `toolload.VersionSegment`. Belt-and-suspenders.

## Verification

- `go build ./...` ✅ from `/home/rgeraghty/projects/hyoka`.
- `go test -race ./hyoka/internal/config/tool/... -timeout 3m` ✅ — all 7 new tests + existing suite pass.
- `go test -race ./... -timeout 5m` — same pre-existing failures Tank Item A documented (`cmd`, `comparison`, `report`, `serve`, `rerender`). Not in my scope; my changes did not introduce or worsen them. Confirmed they failed identically before this commit.
- Smoke test (live eval): not run (would burn Copilot credits). Code path is exercised by the freshness tests + the flock end-to-end test; high confidence.

## What Item E / F should know

- **Item E (Neo, post-session verification gate):** my flock release happens via `defer` inside `ensureRepoCloned`, so by the time `gitFetcher.Fetch` returns the lock is gone. Your gate runs after fetch — no contention. If you ever need a longer-held lock for "atomically install + verify", call a new helper that takes the lock and accepts a body func; don't reach into `ensureRepoCloned`.
- **Item F (Tank, dedup paths):** unchanged from Item A's note. The cache layout is still `<CacheRoot>/repos/<owner>/<repo>/<version>/`. The lock file `.hyoka-lock` lives at the parent — if F enumerates "what's in the cache for owner/repo," it should ignore `.hyoka-lock` (and any future `.hyoka-*` sentinels).
- **Item B (Neo, pluginFetcher — parallel):** your file is in `plugin_fetcher.go` (untracked at time of writing). My changes to `fetcher.go` only touched `ensureRepoCloned` + added the `runGit` var; no merge conflicts expected. If `pluginFetcher.Fetch` ends up calling `ensureRepoCloned` (via shared `pluginCloneFn`), it inherits the freshness branching and flock for free.

## Files touched

```
M  hyoka/internal/config/tool/fetcher.go         (ensureRepoCloned rewrite + runGit hook var)
A  hyoka/internal/config/tool/flock_unix.go      (acquireRepoLock impl, !windows)
A  hyoka/internal/config/tool/flock_windows.go   (no-op stub)
A  hyoka/internal/config/tool/fetcher_freshness_test.go  (7 tests)
```

No changes to `Entry`, `versionSegment`, `FetchRequest`, or `tool.SummarizeToolLoadErrors`. Errors from clone/fetch/checkout still bubble through `gitFetcher.Fetch` unchanged → Neo's Item D formatting is preserved.
