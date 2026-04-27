# Tank Item A — Centralize cache root → `~/.hyoka/cache`

**Author:** Tank 📡
**Date:** 2026-04-27
**Status:** ✅ Implemented — awaiting Ronnie sign-off
**Spec:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md` § Item A

## What shipped

1. **New package `internal/toolload`** with `cacheroot.go`:
   - `CacheRoot()` — `$HYOKA_CACHE_DIR` → `$XDG_CACHE_HOME/hyoka` → `~/.hyoka/cache` → `os.TempDir()/hyoka/cache` (with slog warning). `sync.Once`.
   - `RepoCacheDir(owner, repo, version)` — single source of truth for the on-disk layout.
   - `VersionSegment(v)` — empty/"default" → "default".
   - `SetTestRoot(path) (restore func())` — exported test helper so other packages can override the cache root despite `sync.Once`. (Resets the once internally.)
   - First-call migration warning when `<cwd>/.skills-cache/` exists.

2. **`fetcher.go`:**
   - Dropped `BaseDir` field from `FetchRequest` entirely (delete, not keep-and-ignore — see decisions below).
   - `gitFetcher.Fetch` writes to `toolload.RepoCacheDir(owner, repo, version)`.

3. **`resolve.go`:**
   - Dropped `baseDir` parameter from `FetchRemote(ctx, entry)`. `ResolveSkillsWithReporter` still receives `baseDir` — but only for `resolveLocal` (relative local-skill path resolution), which is unrelated to the cache.

4. **`validate.go`:**
   - `validateSkillEntry` calls `FetchRemote(ctx, entry)` (no configDir).
   - `pluginCheckedPaths` now reads cache paths from `toolload.RepoCacheDir(...)` instead of hand-rolled `~/.hyoka/cache/default/...`.

5. **`internal/plugin/installed.go`:**
   - `ResolveInstalled` looks under `toolload.RepoCacheDir(owner, repo, "default")/...` — same path the new fetcher writes to.
   - Legacy `~/.copilot/installed-plugins/...` fallbacks preserved (Item F removes them).

## Decisions (the open questions)

### D1 — Cache layout: `repos/<owner>/<repo>/<version>/` (not `repos/<version>/<owner>/<repo>/`)
**Choice:** owner/repo first, version last.
**Why:** All versions of one repo cluster on disk — friendlier to `du`, `ls`, and humans hunting for "where did microsoft/skills go?". Morpheus's plan suggested both orderings; Ronnie's task brief explicitly called the swap out. Confirmed correct.

### D2 — `BaseDir` removal: delete, not keep-and-ignore
**Choice:** Delete the field from `FetchRequest` AND drop the `baseDir` arg from `FetchRemote`.
**Why:** Keeping a no-op field is a footgun for the next person who reads the struct and wonders "what does this control?". Test churn was minimal (one assertion in `fetcher_test.go`). Custom fetchers in tests still build because the field is gone, not silently ignored — they get a compile error if they relied on it.
**Caveat:** Anyone writing a third-party `Fetcher` outside this repo will see a breaking change. None known to exist.

### D3 — flock for concurrent fetch safety: **deferred to Item C**
**Choice:** Not implementing flock in this commit.
**Why:** Item C is "version freshness" and is already going to touch `ensureRepoCloned` to add the pinned-vs-default branch and the `meta/<owner>/<repo>.json` write. Doing flock there keeps the diff focused — Item A stays a pure cache-relocation, Item C owns all the per-repo concurrency surface. Adding flock + Windows stub now would also drag in a build-tag file for ~5 lines of payoff before Item C lands, then need re-touching when C arrives.
**Risk window:** Two concurrent `hyoka run` invocations against the same repo can race on `git fetch` between today and Item C landing. Real-world impact: `git` is generally tolerant (worst case: one of the two fetches fails, the eval that owned it surfaces a `tool_load_failure`, retry succeeds). Tank will pick this up in Item C.

### D4 — `SetTestRoot` exported helper
Tests in the `plugin` and `tool` packages need to override `CacheRoot()` because `sync.Once` makes mid-test `os.Setenv("HOME", ...)` ineffective. Rather than dropping `sync.Once` (the spec called for it), I exported a clearly-named test helper: `SetTestRoot(path) (restore func())`. Doc-marked test-only.

## Verification

- `go build ./...` ✅
- `go test -race ./hyoka/internal/{toolload,plugin,config/tool}/... -timeout 3m` ✅
- `go test -race ./... -timeout 3m` — all packages I touched are green. Pre-existing failures (unrelated to this task):
  - `internal/comparison`, `cmd` — `boolPtr` vs `bool` (pre-existing, not in my scope)
  - `internal/report`, `internal/serve` — unknown `Model` field (pre-existing)
  - `internal/rerender` — schema v0 panic (pre-existing)
  - `internal/config/tool` `TestValidateAndExpand_*` Neo's WIP from Item D (joinedToolLoadError) — unblocked by my edits, fully green again.

## Smoke test (live eval)
**Not run.** A live `hyoka run --prompt-id key-vault-dp-python-crud` requires a Copilot session and burns minutes/credits. Code trace + tests give high confidence; Ronnie or Switch can run the smoke if desired. The legacy `.skills-cache/` in cwd (4.7M, dated 2026-04-27 17:50) is still there — left untouched per spec; my migration warning will fire next run that needs a remote skill.

## What Item B / C / F should know

- **Item B (Neo, plugin fetcher):** `toolload.RepoCacheDir(owner, repo, version)` is your one-liner. The plugin fetcher should write under the **same** path the gitFetcher uses today (so `ResolveInstalled` finds it without further changes). When you add a `pluginFetcher`, give it the same canonical layout.
- **Item C (Tank, freshness):** flock is yours. Add it inside `ensureRepoCloned`. The per-repo dir for the lock should be `toolload.RepoCacheDir(owner, repo, "")`'s **parent** (so it covers all versions of the repo) — i.e. lock `<CacheRoot>/repos/<owner>/<repo>/`.
- **Item F (Tank, dedup paths):** `pluginCheckedPaths` and `ResolveInstalled` already share the canonical path via `toolload.RepoCacheDir`. F's job shrinks to extracting a `pluginCacheCandidates(entry, cacheRoot)` helper and dropping the `~/.copilot/installed-plugins/...` legacy fallback after Ronnie confirms no users depend on it.
