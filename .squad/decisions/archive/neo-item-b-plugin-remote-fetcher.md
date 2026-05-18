# Neo Item B — Plugin remote fetcher

**Author:** Neo 💊
**Date:** 2026-04-27
**Status:** ✅ Implemented — awaiting Ronnie sign-off
**Spec:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md` § Item B

## What shipped

1. **New file `hyoka/internal/config/tool/plugin_fetcher.go`:**
   - `pluginFetcher` implementing `tool.Fetcher`. `Name()` is `"plugin-git"`.
   - `CanFetch` matches `entry.ResolvedType() == TypePlugin` AND
     (`entry.Source == SourceRemote` OR `entry.Source == "" && entry.Repo != ""`).
     The inferred-remote case mirrors `validatePluginEntry`'s defaulting rule
     so callers that haven't normalized yet still route correctly.
   - `Fetch` parses owner/repo, resolves `versionSegment` via
     `toolload.VersionSegment`, computes the cache dir via
     `toolload.RepoCacheDir(owner, repo, version)` (Tank's Item A helper),
     then delegates the clone to the package var `pluginCloneFn` (defaults
     to `ensureRepoCloned`). Plugin lookup runs `findPluginInRepo` and
     returns `FetchResult{Dir, Version: versionSegment}`.
   - `findPluginInRepo` and tiny `isPluginDir`/`isSkillDir`/`hasChildSkills`
     mirrors of `internal/plugin`'s helpers — duplicated locally to avoid
     a circular import (`internal/plugin` already imports `internal/toolload`,
     and `internal/config/tool` imports `internal/plugin`; pulling
     `plugin.isPluginDir` would close the loop). Item F can dedup once it
     lifts the path-list helper into `internal/plugin` as a `Candidates(...)`
     export — see "Notes for Item F" below.

2. **Registered in `DefaultRegistry`** (`fetcher.go:152`):
   ```go
   _ = r.Register(&pluginFetcher{})
   _ = r.Register(&gitFetcher{})
   ```
   The Registry's existing "default last" insertion logic (which keys off
   the literal name `"git"`) keeps `gitFetcher` at the tail regardless of
   registration order, but I register `pluginFetcher` first by intent so
   the lookup loop hits it before falling through. `pluginFetcher.Name()`
   is `"plugin-git"` (not `"git"`), so the tail-pinning logic doesn't
   touch it.

3. **`validate.go`:**
   - Threaded `ctx` into `validatePluginEntry` (was synchronous-only).
   - On cache miss for `src == SourceRemote && entry.Repo != ""`, look up
     the registered fetcher and call `Fetch`. On success, re-call
     `plugin.ResolveInstalled` (now hits the freshly-populated cache) and
     fan out children/single-skill via the same code path as the warm-cache
     case. On fetch failure, fall through to the enumerated-paths hard-fail
     and append `"Fetch attempt failed: <err>"` so operators see *why* the
     fetch didn't help.
   - The aggregated `ToolLoadError` flows through Item D's
     `JoinedError`/`SummarizeToolLoadErrors` automatically — no separate path.

4. **Tests** (`plugin_fetcher_test.go`):
   - `TestPluginFetcher_CanFetch` — table covering explicit remote, inferred
     remote, local plugin, plugin-without-repo, remote skill (gitFetcher
     territory), local skill.
   - `TestPluginFetcher_DefaultRegistry_RoutesPluginEntries` — pins the
     registry wiring so accidental reordering trips a test.
   - `TestParsePluginRepo` — owner/repo parsing for `owner/repo`,
     `github.com/owner/repo`, https, `.git`, multi-segment, malformed.
   - `TestPluginFetcher_Fetch_FindsPluginInAllPrecedenceLocations` — three
     subtests covering `.github/plugins/`, `.github/skills/`, `skills/`
     layouts.
   - `TestPluginFetcher_Fetch_PluginNotFoundInRepo` — clone succeeds, but
     the named plugin is absent. Asserts the error enumerates all three
     candidate paths.
   - `TestPluginFetcher_Fetch_CloneFails` — upstream clone error wraps
     correctly.
   - `TestPluginFetcher_Fetch_BadRepo` — empty + single-segment repo
     locators reject before clone.
   - `TestFindPluginInRepo_PrecedenceOrder` — when all three locations
     exist, `.github/plugins/` wins.
   - `TestFindPluginInRepo_ContainerLayout` — `skills/<child>/SKILL.md`
     containers are accepted (mirrors `plugin.isPluginDir`).
   - `TestValidatePluginEntry_FetchSucceedsThenResolvesChildren` — full
     end-to-end: cache miss → fetch (stubbed clone seeds cache) →
     `ResolveInstalled` hits → 2 child skills appear in the report under
     the plugin parent.

   All tests use a stub `pluginCloneFn` (package-level var) for offline
   determinism — no network, no flakiness, fast (< 0.1s for the file).

5. **Updated existing tests** in `plugin_migration_test.go`:
   - `TestValidateAndExpand_MissingRemotePlugin_EnumeratesCachePathsForRepo`
   - `TestValidateAndExpand_RemotePlugin_MissingCache_HardFails`

   Both previously asserted "cache miss = immediate hard-fail with no
   network." With Item B in place those would attempt a real git clone
   of `microsoft/skills`. Both now use `stubPluginCloneFailing` to keep
   the tests offline; they still verify the enumerated-paths contract
   on the failure path (which fires when the fetch attempt also fails).

## Decisions

### D1 — Where the fetcher lives
**Choice:** `hyoka/internal/config/tool/plugin_fetcher.go`, beside `fetcher.go`.
**Why:** The `tool.Registry` is here, the `Fetcher` interface is here, and
`gitFetcher` is here. Plugin fetching is the same shape — clone a repo,
return a directory. Splitting it into a different package (e.g.
`internal/toolload/plugin`) would force `internal/config/tool` to import
the new package just to register the fetcher in `DefaultRegistry`. Adjacent
file, same package, zero churn.

### D2 — Plugin-dir lookup precedence
**Preserved exactly** from `plugin.ResolveInstalled` (`installed.go:47-51`):

1. `<repo>/.github/plugins/<name>`
2. `<repo>/.github/skills/<name>`
3. `<repo>/skills/<name>`

Note this is **NOT** the same order as `findSkillInRepo` in `fetcher.go`
(which puts `.github/skills` first then `.github/plugins`). Skills and
plugins use different precedence — likely historical, but the validator
and the cache-lookup helper both use the plugin-first order, so the
fetcher matches them. Unifying the two orderings is out of scope for
Item B; flag for Item F to consider.

### D3 — Local helpers (`isPluginDir`, `isSkillDir`, `hasChildSkills`)
**Choice:** Duplicate the predicates locally rather than importing
`internal/plugin`.
**Why:** `internal/config/tool` already imports `internal/plugin` for
`plugin.SplitOwnerRepo`, `plugin.ResolveInstalled`, `plugin.EnumerateChildSkills`,
and `plugin.Registry`. Pulling `plugin.isPluginDir` would either require
exporting it (widening the API for one caller) or pulling the whole
predicate over. The predicates are 8 lines total; duplication is cheaper
than the API surface change. Item F is the natural moment to consolidate
when it factors out the shared `pluginCacheCandidates` helper.

### D4 — Test repo selection: mocked, not live
**Choice:** Stub `pluginCloneFn` in tests; no real public-repo integration test.
**Why:**
- `mauromedda/agent-toolkit` (which Ronnie has been using for skills) does
  not host plugins under any of the three precedence paths — it's a
  skill-style layout.
- `microsoft/skills` works for plugins but is large (clone ~30 MB) and
  slow under `go test -race`. Hitting GitHub from CI also drags in network
  flakiness.
- The fetcher's responsibility is "clone via `pluginCloneFn`, then locate
  via `findPluginInRepo`." The clone helper is Tank's territory (Item C
  is rewriting it). Mocking the clone exercises 100% of the fetcher's
  own logic — repo parsing, version segmenting, cache-dir computation,
  precedence lookup, error wrapping — without leaking responsibility for
  the clone itself into Item B's tests.

A live integration test (gated by `//go:build integration` or an env var)
would be valuable for full confidence, but is best added once Item C lands
so it tests the freshness behavior end-to-end. Filed as a follow-up.

### D5 — `pluginCloneFn` package var
**Choice:** Expose the clone helper as a swappable package var rather than
injecting it through the Fetcher struct.
**Why:** Tests need to stub it; production has exactly one implementation.
A struct field would force every call site (validator, future callers) to
plumb it through, which is API churn for zero benefit. The var is unexported
and documented as test-only.

## Verification

- `go build ./...` ✅
- `go test -race ./hyoka/internal/config/tool/... -timeout 3m` ✅ (2.6s)
- `go test -race ./hyoka/internal/{plugin,toolload,eval}/... -timeout 3m` ✅
- Pre-existing failures unchanged and out of scope:
  - `internal/report` — schema v0 panic (pre-existing)
  - `internal/serve`, `internal/comparison`, `cmd` — `boolPtr` / `Model` field (pre-existing per Tank's Item A)
  - `internal/rerender` — schema v0 panic via report (pre-existing)

**Live smoke test:** Not run. Same reasoning as Tank's Item A — burns
Copilot session credits and the test surface (mocked + table-driven) gives
high confidence. Ronnie or Switch can run a `hyoka run --service identity
--language python --config azure-mcp/...` against a config that declares
a remote plugin if a wet-finger check is wanted.

## Notes for Items E / F

- **Item E (Neo, post-session verification):** Plugin children now reach
  the cache via the same `<CacheRoot>/repos/<owner>/<repo>/default/...`
  path the verifier already inspects. No changes to E's planning.
- **Item F (Tank, dedup paths):** When you extract `pluginCacheCandidates(entry,
  cacheRoot)`, also lift `findPluginInRepo` from
  `hyoka/internal/config/tool/plugin_fetcher.go` into `internal/plugin` (or
  a new shared helper). Currently `plugin.ResolveInstalled` and
  `tool.findPluginInRepo` walk the same three-path list independently. Both
  should call one helper. The `isPluginDir`/`isSkillDir`/`hasChildSkills`
  mirrors in `plugin_fetcher.go` go away at the same time.
- **Item C (Tank, freshness):** My fetcher calls `pluginCloneFn` which
  defaults to `ensureRepoCloned`. Your edits to `ensureRepoCloned` —
  pinned-vs-default branch + `meta/<owner>/<repo>.json` + flock at the
  per-repo dir — will land underneath my fetcher transparently. No
  coordination needed unless you change `ensureRepoCloned`'s signature.
