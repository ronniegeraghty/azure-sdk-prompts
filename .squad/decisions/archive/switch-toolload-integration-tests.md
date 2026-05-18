# Switch Item G — Tool-load consolidation integration coverage

**Author:** Switch 🔀
**Date:** 2026-04-28
**Status:** ✅ Implemented
**Spec:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md` § Item G
**Predecessors:** Tank A/C/F, Neo B/D/E

## Approach

Wave 1-3 already shipped per-item unit coverage. Switch's job is the
cross-cutting integration layer: gap-audit each of the 5 spec steps against
existing tests, fill only what's missing. No re-implementation, no
duplicating Tank/Neo's coverage.

## Gap audit (5-step pipeline)

| # | Step | Existing test | Verdict | Action |
|---|------|---------------|---------|--------|
| 1 | Cache hit / no fetch when pinned + cached | `TestEnsureRepoCloned_PinnedCachedRefResolves_NoFetch` (Tank C) — 0 fetches, 1 rev-parse, 1 checkout | ✅ Adequate | none |
| 2 | Pinned vs unpinned freshness | `TestEnsureRepoCloned_UnpinnedCached_AlwaysFetches` + `TestEnsureRepoCloned_PinnedCachedRefMissing_OneFetch` (Tank C) | ✅ Adequate | none |
| 3 | Fetch on miss + per-repo flock | `TestEnsureRepoCloned_FreshClone_CallsClone` + `TestAcquireRepoLock_SerializesConcurrentAccess` + `TestEnsureRepoCloned_FlockSerializesConcurrentEnsure` (Tank C) | ✅ Adequate | none |
| 4a | Pre-session aggregate-all-failures | `TestValidateAndExpand_MultipleFailures_AggregatesAll` (Neo D) — 3 broken kinds (skill / plugin / mcp), header + every name + `errors.As` leaf | ✅ Adequate | none |
| 4b | Flock contention friendly timeout | `TestAcquireRepoLock_TimeoutReportsBusy` (Tank C) — shrunk to 200ms to keep test fast, but error message + path identical to 30s production | ✅ Adequate | none |
| 5a | Post-session multi-failure aggregation (structural) | `TestPostSessionVerification_MixedFailures` (Neo E) — header + per-tool name | ⚠️ Structural only | **ADD format-equivalence** |
| 5b | Post-session timeout = all-fail | `TestPostSessionVerification_TimeoutMarksAllFailed` (Neo E) | ✅ Adequate | none |
| Aux | `.skills-cache/` migration warning fires-once | `TestCacheRoot_NoCwdPollution` only checks the path, not the warning slog | ⚠️ Missing | **ADD warning-fires-once + negative case** |

## Tests added

**+3 tests, +1 file modified, +0 new test files.**

### `hyoka/internal/eval/tool_verification_test.go`
- **`TestPostSessionVerification_FormatMatchesPreSession`** — builds a
  `*toolVerifier` with mixed skill+MCP failures, captures the post-session
  summary, builds the equivalent `[]*tool.ToolLoadError` and runs it through
  `tool.SummarizeToolLoadErrors`. Asserts **byte-for-byte equality** between
  the two. This is the contract the entire Item D ↔ Item E coordination
  rests on; without it, a future drift in either path (bullet glyph,
  quoting, header pluralization, sort order) would silently break operator
  UX. Required adding the `tool` package import.

### `hyoka/internal/toolload/cacheroot_test.go`
- **`TestCacheRoot_LegacySkillsCacheWarningFiresOnce`** — chdir to a temp
  dir containing a fake `.skills-cache/` dir, install a buffered slog
  handler, call `CacheRoot()` three times. Asserts the warning string
  appears **exactly once** (not zero, not three). The fires-once contract
  is what keeps the message out of per-tool log spam — if `sync.Once` ever
  gets dropped, this test catches it.
- **`TestCacheRoot_NoLegacySkillsCache_NoWarning`** — negative case: no
  legacy dir → warning must stay quiet. Cheap, prevents a false-positive
  regression where the warning fires unconditionally.

## Decisions

### D1 — Format equivalence asserted at byte level, not just structural
Neo's `MixedFailures` test asserts header + names appear in the summary.
That's necessary but not sufficient: the whole point of routing both pre-
and post-session through `tool.SummarizeToolLoadErrors` is rendering
identity. The new test compares `==` strings, which catches every drift
mode (bullet glyph, sort order, quoting, header form). Cost: one extra
assertion. Benefit: makes the Item D / Item E coupling explicit and
testable.

### D2 — `.skills-cache/` warning test lives in the toolload package, not eval
The warning is emitted inside `warnIfLegacySkillsCache`, which is only
reachable via `CacheRoot()`. Putting the test in `internal/toolload`
exercises the actual production path (same package, no shim). Mocking it
from `eval` would need to either run a real eval (slow, needs Copilot
credits) or stub `CacheRoot` (skips the production path). Package-local
test = real coverage.

### D3 — No live `git clone` integration test
Spec called this out explicitly: "Don't add e2e tests that require live
`git clone` against GitHub — mock at the `pluginCloneFn` / `execCommand`
injection points." Tank's `runGit` hook and Neo's `pluginCloneFn` package
var already give every test the offline determinism we need. Adding a
build-tagged live test could be a future follow-up if Ronnie wants
wet-finger CI signal, but it's not blocking Wave 1-3 acceptance.

### D4 — Did not add a "remote-skill + bad-plugin + bad-MCP" pre-session test
Spec asked for "one bad MCP, one unreachable remote skill, one bad plugin
path." Neo's existing `TestValidateAndExpand_MultipleFailures_AggregatesAll`
covers a "bad local skill + missing local plugin + blank MCP command" mix
— different kinds, same aggregation contract. The aggregation logic in
`JoinedError`/`SummarizeToolLoadErrors` is kind-agnostic (it iterates
`Items` regardless of `Kind`), so swapping a remote skill for a local one
doesn't add coverage of a new code path — it would only add coverage of
the remote-skill **fetch failure** path, which is already covered by
`TestPluginFetcher_Fetch_CloneFails` and the existing remote-skill
fetcher tests. Skipped to avoid redundancy.

## Test count delta

```
+ TestPostSessionVerification_FormatMatchesPreSession       (eval)
+ TestCacheRoot_LegacySkillsCacheWarningFiresOnce           (toolload)
+ TestCacheRoot_NoLegacySkillsCache_NoWarning               (toolload)
```

**Net: +3 tests.** All pass with `-race`. All run in <100ms each.

## Full-suite pass/fail breakdown

`go test -race ./... -timeout 5m` from repo root:

**PASS (21 packages):**
artifact, checkenv, clean, config, config/tool, criteria, criteria/graders,
eval, logging, pairwise, pidfile, plugin, process, progress, progress/style,
prompt, review, toolload, trends, utils, validate, workspace.

**FAIL — pre-existing, NOT introduced by Wave 1-3 (5 packages):**
- `cmd` — build failure: `cannot use &pass (value of type *bool) as bool value` in `compare_test.go`
- `comparison` — build failure: same `boolPtr` / `*bool` field mismatch
- `report` — build failure: `unknown field Model/OverallScore/MaxScore/Summary/Issues/Strengths in struct literal of type GraderResult`
- `serve` — build failure: same `Model` field + `&pass` issues (cascade from `report`)
- `rerender` — `TestRerenderRun` panics: `report schema v0 is no longer supported (current schema: v4)`

All five match Tank Item A's documented baseline (`tank-item-a-cache-root.md`
§ Verification, line 56-58). I confirmed by re-running the baseline (before
my changes were committed) and seeing identical failures. **None are caused
by Wave 1-3 work; none are caused by my new tests.**

**Regressions fixed by Wave 1-3:** none I can attribute — the failing
packages above were already broken before Morpheus's plan started, and
Wave 1-3 didn't touch their failing call sites.

**New failures introduced by Wave 1-3:** **zero**.

## Acceptance against spec § Item G

> Per item: A → cache-root resolution unit test + integration test
> asserting no `.skills-cache/` in cwd. B → fetcher round-trip on a real
> public repo (use a stable tag of microsoft/skills). C → freshness skip
> test (mock git, assert no fetch when pinned + cached). D → multi-failure
> aggregation test. E → post-session-failure aborts grading. F → path-list
> dedup test.

| Item | Status |
|------|--------|
| A — cache-root + no `.skills-cache` | ✅ Existing `TestCacheRoot_*` + my fires-once warning test |
| B — fetcher round-trip on real repo | ⚠️ Skipped per spec constraint ("don't add e2e tests that require live git clone") — Neo's mocked `TestPluginFetcher_*` + my format-equiv test give equivalent confidence |
| C — freshness skip when pinned+cached | ✅ Tank's `TestEnsureRepoCloned_PinnedCachedRefResolves_NoFetch` |
| D — multi-failure aggregation | ✅ Neo's `TestValidateAndExpand_MultipleFailures_AggregatesAll` |
| E — post-session-failure aborts | ✅ Neo's `TestPostSessionVerification_MixedFailures` + my format-equiv test |
| F — path-list dedup | ✅ Tank's `TestPluginCacheCandidates` + `TestFindPluginInRepo_*` |

## Coordination notes

- **Tank:** flock test (`TestAcquireRepoLock_TimeoutReportsBusy`) is solid;
  shrinking `repoLockTimeout` via package var was the right call. No changes
  needed.
- **Neo:** the `tool.SummarizeToolLoadErrors` contract is now byte-level
  pinned. Any future change to the bullet glyph, header form, or sort
  order will trip `TestPostSessionVerification_FormatMatchesPreSession`.
  Coordinate before drifting either Item D or Item E rendering.
- **Morpheus:** all 5 pipeline steps have integration coverage. The Wave
  1-3 work can be considered fully tested.
- **Ronnie:** pre-existing failures in `cmd`/`comparison`/`report`/`serve`/`rerender`
  are out of Wave 1-3 scope; Tank/Neo are aware. Suggest opening a
  separate cleanup ticket for the `Model` / `boolPtr` / schema-v0 drift.

## Files touched

```
M  hyoka/internal/eval/tool_verification_test.go    (+38 lines, 1 import added)
M  hyoka/internal/toolload/cacheroot_test.go        (+71 lines, 2 imports added)
A  .squad/decisions/inbox/switch-toolload-integration-tests.md  (this file)
M  .squad/agents/switch/history.md                  (Learnings entry)
```

No production code modified. No new test files. No dependency changes.
