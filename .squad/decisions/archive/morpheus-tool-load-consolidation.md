# Tool-Load Consolidation Plan

**Author:** Morpheus 🕶️
**Date:** 2026-04-25
**Status:** Proposal — awaiting Ronnie review before fan-out to Neo / Tank
**Scope:** Remote tool loading (skills + plugins + MCP), shared cache, and pre/post-session verification gates

---

## 1. Current state mapped to Ronnie's spec

Five bullets, one per step of the target flow:

1. **Cache check** — Split-brain. Skills look at `<BaseDir>/.skills-cache/<version>/<owner>/<repo>/` (`fetcher.go:209-210`). Plugins look at `~/.hyoka/cache/default/<owner>/<repo>/...` (`installed.go:45-49`) plus `~/.copilot/installed-plugins/...` legacy paths. Two completely separate trees, neither acknowledges the other.
2. **Version freshness** — Half-implemented for skills. `ensureRepoCloned` (`fetcher.go:303-322`) calls `git fetch --all --tags` + `checkout` on **every** invocation regardless of whether `entry.Version` is pinned or "default". Plugins have **no** freshness path at all — `ResolveInstalled` is a stat-only lookup. The pinned-vs-latest distinction Ronnie wants does not exist in code.
3. **Fetch** — Only skills can fetch (`gitFetcher.Fetch` in `fetcher.go:199`). Plugins **cannot fetch** — `validatePluginEntry` (`validate.go:303`) only enumerates checked paths and fails. Ronnie's "fetch if missing" rule is unimplemented for plugins; users must pre-install via `/plugin install` or hand-clone into `~/.hyoka/cache/default/...`.
4. **Hard fail on fetch failure** — Partial. Skill fetch errors propagate via `ValidateAndExpand` → `FirstError()` → `copilot.go:185` returns a `tool_load_failure` EvalResult (the engine record path) with `error_category=tool_load_failure`. Pre-session validation IS a hard-fail at the eval boundary for the generator path. **However** `cmd/run.go:401-407` validates reviewer tools separately and `validateEntries` is sequential (`validate.go:282`), so the "wait for all tools, then fail" requirement is not met — the first failure short-circuits the rest of the entries silently.
5. **Session-level verification** — Mechanism exists but the gate is **disabled**. `toolVerifier` (`tool_verification.go`) accumulates `SessionSkillsLoaded`/`SessionMcpServersLoaded`, emits `EventToolsVerified`, and the renderer flips Loaded→Failed (`display_interactive.go:533-546`). But `waitForToolVerification` is never called — the comment block at `copilot.go:643-655` admits the gate is intentionally bypassed to avoid a deadlock (SDK only emits load events after the first message round-trip). Result: post-session verification is **cosmetic-only**. The eval keeps running even when every tool failed to load.

---

## 2. Gaps

Each gap = (a) which spec step it violates · (b) where in code · (c) user-visible symptom.

1. **`.skills-cache/` ends up in cwd, not `~/.hyoka/cache/`** — (a) Step 1. (b) `cmd/run.go:403` passes `ConfigDir: ""` to `ValidateAndExpand` for the reviewer pass; that empty string flows to `FetchRemote` → `FetchRequest.BaseDir = ""` → `filepath.Join("", ".skills-cache", ...)` produces a relative path resolved against `os.Getwd()`. (c) Ronnie sees `.skills-cache/` materialize wherever he runs `hyoka`. Bonus: for the generator path, `BaseDir` is the tmp `hyoka-config-*` dir that gets `os.RemoveAll`'d after each eval (`copilot.go:165`), so the "cache" is destroyed every run — every eval re-clones. The cache is effectively **never reused**.

2. **Plugins cannot fetch — only inspect** — (a) Steps 1, 3. (b) `installed.go:37-69` stats predetermined paths and returns `""` on miss; `validatePluginEntry` (`validate.go:423-441`) then hard-fails with an enumerated path list. There is no plugin equivalent of `gitFetcher.Fetch`. (c) Users get "plugin not found, checked: …" with instructions to run `/plugin install`. Hyoka silently outsources plugin acquisition to Copilot CLI, which means an offline-or-fresh-machine `hyoka run` will fail until the operator manually pre-installs every remote plugin.

3. **No version pinning semantics** — (a) Step 2. (b) `entry.Version` exists (`entry.go:33`) and is plumbed into `FetchRequest.Version` (`resolve.go:280`), but `ensureRepoCloned` always runs `git fetch --all --tags` then `git checkout <ref>` (`fetcher.go:310-318`). Pinned versions hit the network on every call; "default" doesn't have a freshness check (it just checks out HEAD). (c) Slow runs and unnecessary network on cached pinned repos; non-deterministic "default" behavior — a quiet upstream commit silently changes which skill code an eval ran against, with no audit trail.

4. **No version-segment for plugins** — (a) Step 2. (b) `installed.go:45` hardcodes the path segment to `default` regardless of any version field on the plugin entry. Even if we added a `version:` to a plugin entry, it would be ignored. (c) Cannot pin a plugin to a tag/branch/SHA at all today.

5. **Sequential validation short-circuits before all tools are evaluated** — (a) Step 5 (Ronnie's "do not stop at first failure"). (b) `validateEntries` is sequential (`validate.go:282-292`); each `validate*Entry` records its result and returns. `ValidateAndExpand` then calls `FirstError()` (`validate.go:185`) — but every individual entry validates fully before that, so the report DOES contain every entry. **Pre-session this is fine.** Where it fails is the post-session verifier: `waitForToolVerification` requires that BOTH `skillsEvtSeen` AND `mcpEvtSeen` fire (or only the configured one) — but the SDK doesn't always emit `SessionMcpServersLoaded` when no MCP is configured for the session, and the gate is currently bypassed entirely.

6. **Post-session verification gate is bypassed** — (a) Steps 4, 5 (post-session leg). (b) `copilot.go:643-655` explicitly disables the gate with a TODO citing #347. The verifier still emits a display event, but no code path turns "verifier reported failures" into an aborted eval. (c) An MCP server that fails to launch (e.g., azure-mcp can't authenticate) will display as Failed in the Tools block, then the eval **runs anyway** — graders score generated code that never had access to the tools the prompt was designed to require. False-positive evals.

7. **MCP pre-session validation is paper-thin** — (a) Steps 1-3 (MCP gets none of the lifecycle). (b) `validateMCPEntry` (`validate.go:723-747`) only checks `command` or `url` is non-empty. There's no fetch/install for npx-launched MCPs, no version pin, no health check. (c) Typos in `command:` only fail at session start (post-bypassed-gate), so see Gap 6.

8. **`configDir` is overloaded as both isolated config root AND fetch base** — (a) Step 1 (architectural). (b) `validateSkillEntry` calls `FetchRemote(ctx, entry, configDir)` (`validate.go:498`) where `configDir` is the throwaway `hyoka-config-*` tmp dir. The fetcher then writes its cache **inside** the dir we're about to delete. (c) Compounds Gap 1 — the cache layout is tied to a per-eval ephemeral dir instead of a stable, user-scoped location.

9. **Reviewer-side validation is a separate code path with different defaults** — (a) Step 4 (consistency). (b) `cmd/run.go:401-409` calls `ValidateAndExpand` with `ConfigDir: ""` for the reviewer factory, while `copilot.go:175` calls it with the isolated tmp dir for the generator. Two callers, two BaseDir conventions, two opportunities for the cache-location bug. (c) Reviewer skill fetches drop `.skills-cache/` straight in cwd (Gap 1's actual root cause).

---

## 3. Proposed module shape

The fix collapses the four-way divergence (skill-fetch / plugin-lookup / cache-root / verification) into a single tool-loader module that all callers reach through one function.

### 3.1 New package: `internal/toolload`

A thin orchestration layer that owns "give me the on-disk path for this remote tool, cached, fresh, hard-failed if broken."

```go
package toolload

// CacheRoot returns the single canonical cache directory:
//   $HYOKA_CACHE_DIR  (override)
//   else $XDG_CACHE_HOME/hyoka (when set)
//   else ~/.hyoka/cache
// The value is computed once per process.
func CacheRoot() string

// Resolve materializes a remote tool (skill or plugin) on disk and returns
// its absolute path. Honors entry.Version for pinning. Calls into the
// existing tool.Fetcher registry for the actual git/npx/etc work.
//
// Behavior:
//   - cached + pinned + present  → return cached path, no network
//   - cached + unpinned          → check freshness (see Freshness), refresh if stale
//   - not cached                 → fetch, populate cache, return path
//   - fetch failure              → return ToolLoadError, never partial state
func Resolve(ctx context.Context, entry tool.Entry) (Result, error)

type Result struct {
    Dir          string  // absolute path to skill/plugin dir
    Version      string  // resolved git ref ("default" → resolved SHA)
    FromCache    bool    // for logging/telemetry
    Refreshed    bool    // true if we ran git fetch this call
}

// Freshness controls how often unpinned entries re-check upstream.
// Default is RefreshAlways for "default" version (matches today's behavior);
// pinned versions are RefreshNever once present.
type Freshness int
const (
    RefreshAlways Freshness = iota
    RefreshNever
    RefreshTTL
)
```

### 3.2 Cache layout (one tree, owner-scoped)

```
${CacheRoot}/
  repos/<owner>/<repo>/<version-segment>/   # full clone, shared between skills & plugins
  meta/<owner>/<repo>.json                  # last-fetch timestamp, resolved SHA
```

Skill `Resolve` returns `repos/<owner>/<repo>/<v>/[<subpath>|<found-skill-dir>]`.
Plugin `Resolve` returns `repos/<owner>/<repo>/<v>/[.github/plugins/<name>|.github/skills/<name>|skills/<name>]`.

Same clone, two views. No more `.skills-cache/` separate tree, no more `~/.copilot/installed-plugins/` dependency.

### 3.3 Where the hard-fail decision lives

Two boundaries, one decision rule each:

**Pre-session:** `tool.ValidateAndExpand` already returns `(report, *ToolLoadError)`. Keep it. Change `FirstError()` → `AllErrors()` and have callers (`copilot.go:185`, `cmd/run.go:407`) format **every** failure into the EvalResult, not just the first. Both call sites already abort on non-nil error — that part is correct.

**Post-session:** Re-enable the gate. `copilot.go` already builds a `toolVerifier`; after `session.SendAndWait` returns (workaround for the SDK's "events fire after first round-trip" timing), call `waitForToolVerification(ctx, verifier, timeout)`. If the returned `[]ToolStatus` contains any `Failed`, return an EvalResult with `error_category=tool_load_failure` and abort grading. Concretely:

```go
// after SendAndWait succeeds
statuses, err := waitForToolVerification(genCtx, verifier, 5*time.Second)
if err != nil { /* timeout = all-fail */ }
if firstFailed := failedStatus(statuses); firstFailed != nil {
    return &EvalResult{
        Success:       false,
        ErrorCategory: "tool_load_failure",
        Error:         "post-session verification failed: " + summarize(statuses),
    }, fmt.Errorf("tool verification: %s", firstFailed.ToolName)
}
```

The verifier already waits for **all** configured kinds before emitting (`tool_verification.go:94-108`), so "wait for all verdicts before deciding" is satisfied by the existing `emitIfReady` contract — we just have to actually consume it.

### 3.4 Cache-root resolution lives in exactly one place

`toolload.CacheRoot()`. Every other caller — the existing fetcher, plugin lookup, future test helpers — calls it. Delete the `req.BaseDir` field from `FetchRequest`; it's a footgun. The Fetcher interface stays (custom fetchers are still useful), it just gets a `CacheDir string` field that the loader pre-computes.

---

## 4. Plan — work items, ordered for parallelism

Each item is independently shippable. Dependencies noted explicitly so we can fan out cleanly.

### Item A: Centralize cache root → `~/.hyoka/cache`
**Owner:** Tank
**Depends on:** nothing
**Description:** Create `internal/toolload/cacheroot.go` exposing `CacheRoot()`. Replace the `filepath.Join(req.BaseDir, ".skills-cache", ...)` line in `fetcher.go:210` with `filepath.Join(toolload.CacheRoot(), "repos", versionSegment, owner, repo)`. Drop `BaseDir` from `FetchRequest` (or keep it but ignore — Tank chooses based on test churn). Update `installed.go:45` to use `CacheRoot()` instead of `~/.hyoka/cache/default`. Update both callers in `copilot.go:175` and `cmd/run.go:401` so neither passes a stale BaseDir. Migration: on first run, if `~/.hyoka/cache/default/<...>` exists, leave it — the new layout is `repos/<...>` so they can coexist briefly; add a one-line warning if a `.skills-cache/` dir is detected in cwd. **Acceptance:** running `hyoka run` from any cwd never creates `.skills-cache/` in that cwd.

### Item B: Plugin fetch (`FetchRemotePlugin`)
**Owner:** Neo
**Depends on:** A (needs `CacheRoot`)
**Description:** Add a `pluginFetcher` to the existing `tool.Registry` (mirrors `gitFetcher`). `CanFetch` returns true for `entry.Type == TypePlugin && entry.Source == SourceRemote`. `Fetch` git-clones the repo (reusing `ensureRepoCloned`) and returns the plugin dir using the same lookup logic `installed.go` does today (`.github/plugins/<name>` → `.github/skills/<name>` → `skills/<name>`). `validatePluginEntry` calls the fetcher when the cache lookup misses, instead of immediately hard-failing. The Copilot CLI legacy `~/.copilot/installed-plugins/...` fallback stays as a transition aid. **Acceptance:** an entry `{type: plugin, source: remote, repo: microsoft/skills, name: azure-sdk-python}` resolves on a fresh machine with no manual install.

### Item C: Version freshness (pinned-vs-latest)
**Owner:** Tank
**Depends on:** A
**Description:** In `ensureRepoCloned` (`fetcher.go:303`), branch on `version`: if pinned (anything not "default") AND `<cacheDir>/.git` exists AND `git rev-parse <ref>` resolves, skip the `git fetch`. If unpinned ("default"), keep current behavior (fetch + checkout HEAD). Write a tiny `meta/<owner>/<repo>.json` recording last-fetch UTC and the resolved SHA so the report can surface "ran against SHA X fetched at T". No TTL logic in v1 — Ronnie's open question 5.2 below decides whether to add one. **Acceptance:** a second run with a pinned version does no network I/O (verifiable via `--log-level debug` showing zero git invocations).

### Item D: Pre-session hard-fail collects ALL failures
**Owner:** Neo
**Depends on:** nothing (independent of A-C)
**Description:** Add `(*ToolLoadReport).AllErrors() []*ToolLoadError`. Change `ValidateAndExpand` to return `(report, errors.Join(report.AllErrors()...))` instead of just `FirstError()`. Update `copilot.go:185` and `cmd/run.go:407` to render the joined error in `Error` / `ErrorDetails` so operators see every broken tool in one shot. `validateEntries` stays sequential (cheap and ordered output beats parallel speedup). **Acceptance:** a config with three broken tools reports all three in the failure message, not just the first.

### Item E: Re-enable post-session verification gate
**Owner:** Neo
**Depends on:** D (so the failure path is consistent)
**Description:** Remove the `// NOTE: Tool validation gate is DISABLED` block at `copilot.go:643`. Insert `waitForToolVerification(ctx, verifier, …)` AFTER `session.SendAndWait` returns successfully (per the comment, that's when the SDK emits the load events). Timeout: 30s. On timeout OR any Failed status, build an EvalResult with `error_category=tool_load_failure` and abort before grading. The existing display-only flip in `display_interactive.go:533` stays — it's the user-facing rendering. **Acceptance:** an eval with a deliberately-broken MCP entry (`command: nonexistent-binary`) errors out post-session with `tool_load_failure` instead of running graders.

### Item F: Migrate plugin path enumerations to one source of truth
**Owner:** Tank
**Depends on:** A, B
**Description:** `pluginCheckedPaths` (`validate.go:243-268`) and `ResolveInstalled` (`installed.go:37`) duplicate the path list. After A+B land, `ResolveInstalled` becomes a thin wrapper over `toolload.Resolve(entry)` (cache-only mode), and `pluginCheckedPaths` is generated from a single `pluginCacheCandidates(entry, cacheRoot)` helper used by both. Drop the `~/.copilot/installed-plugins/` legacy paths once Ronnie confirms no users depend on them (open question 5.4). **Acceptance:** removing or adding a path requires editing one function.

### Item G: Tests for each item
**Owner:** Switch (production code by Neo/Tank — Switch covers it)
**Depends on:** the production change for each item
**Description:** Per item: A → cache-root resolution unit test + integration test asserting no `.skills-cache/` in cwd. B → fetcher round-trip on a real public repo (use a stable tag of microsoft/skills). C → freshness skip test (mock git, assert no fetch when pinned + cached). D → multi-failure aggregation test. E → post-session-failure aborts grading. F → path-list dedup test. Switch follows its existing testing-patterns SKILL (table-driven, `-race`).

### Suggested rollout order

```
A ──┬─→ B ──┐
    ├─→ C   ├─→ F
    │       │
    D ──→ E ┘
```

A is the critical-path blocker (cache-root touches everything). B and C can proceed in parallel after A. D is fully independent and can land first to give Ronnie quick visual feedback that "now I see all the failures." E waits for D so the failure-rendering is consistent across pre/post boundaries. F polishes after B lands.

**Earliest parallel wave:** A + D (different files, no overlap).
**Second wave:** B + C (both depend on A).
**Third wave:** E + F.

---

## 5. Open questions for Ronnie

These need a decision before Neo/Tank start. Each has a default if Ronnie wants to defer.

1. **Cache root override:** Should `$HYOKA_CACHE_DIR` and/or `$XDG_CACHE_HOME` override `~/.hyoka/cache`? **Default:** support both, env first, XDG second, `~/.hyoka/cache` last.
2. **"Latest" definition:** Is unpinned ("default") = git default branch HEAD as of the run? **Default:** yes, with a `git fetch` every run (current behavior). Alternative: TTL-based ("only re-fetch if last fetch > 1h ago") — cheaper on hot loops but introduces a staleness window. Recommend deferring TTL to v2.
3. **Concurrent fetch safety:** Two parallel `hyoka run` processes pointing at the same cache will race on `git fetch` of the same repo. **Default:** add a `flock(2)` per `<owner>/<repo>` directory — cheap, POSIX-only (Windows uses `golang.org/x/sys/windows.LockFileEx` — there's a `windows-compatibility` skill noting this is a sharp edge). Recommend doing it now in Item A.
4. **Drop `~/.copilot/installed-plugins/` legacy?** Item F asks whether to retire the Copilot-CLI-installed fallback. **Default:** keep it for one release with a deprecation log line, drop in the next.
5. **Post-session verification timeout:** What's the ceiling? `tool_verification.go` test mentions 30s. **Default:** 30s, configurable via `--tool-verify-timeout` for slow MCP cold starts.
6. **Should fetch failures emit a per-tool `Failed` verdict in the report (granular) or a single eval-level error (current)?** Spec step 4 says both — per-tool verdict AND hard-fail. **Default:** Item D already aggregates all per-tool errors into the eval-level message; verdict rows are already in `ToolLoadReport.Items`. Confirm this satisfies Ronnie's "emit a Failed verdict for that tool" wording.

---

## 6. Out of scope (intentionally)

- **Replacing the Copilot SDK's tool-loading mechanism.** We stay on the SDK's `SessionConfig.SkillDirectories` / `MCPServers` contract. We're tightening the layer above it, not below.
- **Local skill / local plugin resolution.** `SourceLocal` paths already work and aren't part of the consolidation. They use `configDir` for relative-path resolution and that stays.
- **Reviewer-vs-generator role plumbing.** The split (`cmd/run.go` for reviewers, `copilot.go` for generators) is preserved — both just route through the new loader.
- **`ToolLoadResult` schema changes** in `report/`. The existing JSON shape consumed by the site stays. We add `RefreshedAt` / `ResolvedSHA` only if Ronnie wants the audit trail (open question 2 spillover).
- **Removing the post-session display flip.** `display_interactive.go:533` stays — it's the only thing that gives operators a real-time view today. It now becomes belt-AND-braces: display flip + engine hard-fail.
- **MCP fetch/install lifecycle** (npx pre-pull, server health check). Gap 7 is real but the answer is large enough to deserve its own proposal — flag for a follow-up after this lands.
- **Migration tool to move existing `.skills-cache/` dirs.** Detection + warning only; users can `rm -rf` the legacy dirs themselves once everything works.

---

**Bottom line:** the ground is mostly there — the verifier exists, the fetcher exists, the hard-fail plumbing exists. The work is connecting them with one cache root, one fetch path that covers both skills and plugins, and turning the disabled post-session gate back on now that we know to wait until after the first send round-trip.
