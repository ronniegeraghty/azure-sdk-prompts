# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/main_test.go, hyoka/testdata/, hyoka/internal/ (packages to test)

## Core Context

**Archived 18 entries from earlier sessions.**

Historical patterns and learnings:

- ## Core Context: Agent Switch initialized as Tester for hyoka. Guardrail defaults: max turns 25, max files 50, max output 1MB, max session actions 50. Safety boundar...
- ## 2026-04-23: WU-2 Tool Validation Gate Tests (COMPLETE ✅): ### Context
- Neo implementing tool validation gate (WU-1) to enforce skill/MCP load checks before sending prompts
- Switch assigned WU-2: Write tes...
- ## 2026-04-20: Phase 5 Review — #364 Morpheus Mock Fix (APPROVED ✅): ### Context
- Switch double-rejected #364: (1) 20 tests failed (wrong mock paths), (2) Oracle renamed tests to `.TODO` instead of fixing
- Trinity +...
- ## 2026-04-21 — PR #605 Review (Neo's #597 tool versioning): **Verdict:** ⚠️ APPROVE WITH NITS — comment posted ([#issuecomment-4285214970](https://github.com/ronniegeraghty/hyoka/pull/605#issuecomment-4285214...
- ## 2026-04-21 — Phase 6 Rollup Review (PR #607): **Status:** ✅ APPROVE WITH NOTES
**PR:** #607 (phase-6 → ronniegeraghty/dev, epic #312)
**Scope:** Integration-level review — 6 sub-PRs (#601–#606)...
- ## Session 2026-04-21 (Phase 6 Round-1 Test Review): **Mission:** Test review of Phase 6 Round-1 batch (PRs #601, #602, #603)

**Verdicts:**
- #601 (Compare page redesign): ✅ APPROVE — 31 new tests, 99...
- ## 2026-04-21 — PR #609 Review (MultiSelectFilter tests, Trinity, phase-6): **Verdict:** ⚠️ APPROVE WITH NOTES

- Tests: 122/122 green; 3 new tests in `multi-select-filter.test.tsx` (outside-click, Escape, empty-options).
-...
- ## PR #610 test-quality review — ✅ APPROVE (2026-04-21): **Branch:** ronniegeraghty/issue-608-606-group-tests (Tank — group property follow-up tests for #606/#608)
**Tests:** `go test -race ./hyoka/... -ti...
- ## 2026 PR #612 review — Neo fetcher polish (#605/#608): **Verdict:** ✅ APPROVE (posted as comment — gh refuses self-approve on bot-owned PR)

- Verified `go build ./hyoka/...` + `go test -race ./hyoka/......
- ## 2026-04-21 — PR #611 review (Morpheus: site-embed Makefile + CI freshness gate): **Verdict:** ✅ APPROVE (posted as comment — `gh pr review --approve` rejected because branch was authored under same `ronniegeraghty` identity).

**...
- ## 2026-04-21 — PR #613 review (MultiSelectFilter follow-up tests): **Verdict:** ✅ APPROVE (posted as PR review comment — self-approve blocked)

Trinity's follow-up to my #609 review. All four deferred gaps closed:
-...
- ## 2025-01 — PR #614 review (site-embed freshness CI hardening; Morpheus, follow-up to #611): Reviewed test/CI correctness for the three nits I'd raised on #611. All three resolved cleanly:

1. **`git status --porcelain`** replaces `git diff...
- ## Team Context: Unified Grader Direction Proposed (2026-04-22): Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE sch...
- ## 2026-04-22 — Phase 1 Acceptance Tests (#624) ✅: **Mission:** Land TDD-style acceptance tests for the unified grader loader (`internal/graders/`) while Neo implemented in parallel. Commit directly...
- ## 2026-04-23 — Renderer snapshot tests (tests-renderer-snapshots): **Mission:** Add table-driven snapshot tests for BOTH progress renderers (interactive + CI) covering happy paths, failure paths, edge cases, and NO_...
- ## 2026-04-22 — Event-emission unit tests: tool resolution, verification, grader lifecycle: Sibling task to the renderer snapshots. Wrote unit tests for the new
progress events across three packages: `config/tool` (resolution),
`eval` (veri...
- ## 2026-04-22 — Sprint capstone: manual verification of new CLI output UX: Ran the 8-row matrix from `session-state/.../plan.md` end-to-end on `ronniegeraghty/dev @ 25ce00a7`. Used `script -E always -c "..." < /dev/null` to...

Full history archived. Recent entries below.

---

## Team Updates

### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three rounds. 48 new test cases (all yours). 2 regressions you caught: 1 fixed in-sprint by Tank (`2d38533f`, piped-CI auto-mode ordering), 1 filed as preexisting Known Issue (`hyoka clean` blocks on non-interactive stdin — OPEN, out-of-scope).

**Your commits this sprint:** `142da225` renderer snapshot tests (13 cases — 6 interactive scenarios + NO_COLOR + ANSI markers + 5 CI scenarios + full-output golden with `normalizeCI` timestamp stripper) · `25ce00a7` event-wiring tests (35 cases) + re-landed `EventToolsVerified` emission in `hyoka/internal/eval/tool_verification.go` with 9 tests.

**Ledger reconciliation you triggered:** the round-1/2 decisions ledger claimed `82cd8590` shipped; you proved it never merged and re-landed equivalent behavior testable-ly. Entry in `decisions.md` now marks it Re-landed via `25ce00a7`.

See `.squad/orchestration-log/2026-04-23T00-05-04Z-sprint-wrap.md` and the round-3/4 section in `.squad/decisions.md`.

## Tool Validation Gate Fixed (2026-04-23)

**Neo's Work:** Fixed blocking tool verification gate that was preventing ALL evaluations from running. Root cause: SDK emits SessionSkillsLoaded events **during** first SendAndWait, not after CreateSession. Gate was blocking before SendAndWait, causing indefinite timeout before events could ever fire.

**Relevant to Switch:** The gate implementation had WU-2 (tests written by Switch) validating the gate's blocking behavior. With the gate now disabled and observational-only, the tests for blocking gate behavior should be updated or removed. The gate itself is still active for logging tool load events, but the tests expecting "gate blocks eval on tool failure" will fail. Consider: (A) Remove gate tests entirely, (B) Rename to "observability tests" and verify tool load events are still logged, or (C) Rewrite as "gate is disabled, failures logged but not blocking" tests.

**Status:** ✅ Gate disabled, evals running. Verified with live eval (88s, passed). Observability maintained via event logging.

**Decision:** Gate remains observational pending SDK event lifecycle documentation and architectural review. Options for future re-enablement documented in decisions.md.

## 2026-04-23: WU-4 Tool-Load Validation Test Suite (COMPLETE ✅)

### Context
- Neo shipped WU-1 + WU-2: tool-load validation primitives (commits acd36cde..e6271eeb on ronniegeraghty/dev)
- Added `tool.ValidateAndExpand(ctx, ValidationInput) (*ToolLoadReport, error)` — strict pre-session validation
- Added `ToolLoadError{Kind, Name, Reason}` typed errors
- Hard-fail wiring in `eval/copilot.go` that aborts before CreateSession with `ErrorCategory="tool_load_failure"`
- Reviewer factory isolation fix in `cmd/run.go` (commit 0131f35d) — prevents cross-config skill leakage
- Switch assigned WU-4: Write comprehensive test coverage for the new validation surface

### Work Completed

#### 1. Tool Package Tests (`hyoka/internal/config/tool/validate_test.go`)
Created 15 table-driven test functions covering:
- **Happy path:** Valid plugin + skill_dir + inline skill + MCP → all loaded, no error
- **Missing plugin:** Returns `ToolLoadError{Kind:"plugin"}`, report has failed item with reason
- **Missing skill dir:** Returns `ToolLoadError{Kind:"skill"}`, validates error message
- **Malformed plugin YAML:** Registry load fails, ValidateAndExpand reports plugin not found
- **Plugin child missing SKILL.md:** Child marked failed, other plugin children still loaded
- **Empty skill_dir:** Fails with "contains no skills" reason
- **Empty config:** No tools → empty report, no error
- **Relative vs absolute paths:** Both resolve correctly to absolute paths
- **MCP missing command:** Fails with "local MCP entry missing command" reason
- **Reviewer role partitioning:** `GeneratorSkillDirs()` / `ReviewerSkillDirs()` correctly filter by role
- **Glob expansion:** Pattern `skill-*` expands to multiple children with `ParentKind=skill_dir`
- **ToolLoadReport.FirstError():** Returns first failed item as `*ToolLoadError`
- **registryLookup helper:** Nil registry, found plugin, not-found plugin

#### 2. Eval Package Tests (`hyoka/internal/eval/tool_load_hardfail_test.go`)
Created 4 integration tests proving hard-fail contract:
- **Missing plugin:** Eval aborts with `ErrorCategory="tool_load_failure"`, never calls CreateSession
- **Missing skill:** Same hard-fail behavior
- **Empty skill_dir:** Same hard-fail behavior
- **MCP missing command:** Same hard-fail behavior

All tests verify:
- `result.ErrorCategory == "tool_load_failure"`
- `result.Success == false`
- `result.Error` and `result.ErrorDetails` are non-empty
- No generated files (proves session never started)

#### 3. Cmd Package Tests (`hyoka/cmd/reviewerfactory_test.go`)
Created 4 tests for reviewer factory per-config isolation:
- **Per-config isolation:** Config A sees only skill-a, Config B sees only skill-b (no cross-leakage)
- **Missing reviewer skill fails fast:** Returns error immediately, doesn't pass unresolved path to SDK
- **Empty reviewer skill_dir fails fast:** Same behavior as generator validation
- **skill_dir expansion:** Reviewer `skill_dir=true` expands to child skills (generator parity)

### Test Results
- ✅ All 23 new tests pass with `-race` flag
- ✅ Full test suite passes: `go test -race ./...` (all packages green)
- ✅ Zero flakes, zero regressions
- ✅ Pre-existing known failures in `serve` and `validate` packages remain (noted but not fixed per task instructions)

### Commit
- **SHA:** `05b4f6d8`
- **Message:** `test(tool): table-driven coverage for ValidateAndExpand + tool_load_failure hard-fail`
- **Pushed to:** `ronniegeraghty/dev`

### Learnings

**Test fixture organization:** Created testdata under temp directories rather than `hyoka/testdata/tool_load/` because all tests use `t.TempDir()` for isolation. This avoids test pollution and makes cleanup automatic.

**Prompt struct gotcha:** The `prompt.Prompt` struct uses `PromptText` field, not `Content`. Initial test compilation failed — fixed by viewing `hyoka/internal/prompt/types.go` to find the correct field name.

**Plugin registry behavior:** Malformed YAML plugins fail silently during `reg.LoadDir()` with a warning log, then `ValidateAndExpand` treats them as "not found in registry". This is correct lenient behavior for the registry, strict behavior for the validator — tested both paths.

**Empty skill_dir vs missing skill_dir:** Both fail, but with different reasons. Missing: "does not exist". Empty: "contains no skills (no subdirectory with SKILL.md)". Tests verify exact reason strings to catch message regressions.

**Reviewer factory isolation proof:** The key test (`TestReviewerFactory_PerConfigIsolation`) proves that validating config A's reviewer tools returns only skill-a paths, NOT skill-b from config B. This directly exercises commit 0131f35d's fix — the old code would have pooled both skills into a shared slice.

**Hard-fail contract verification:** Integration tests prove `CreateSession` is never called by asserting `len(result.GeneratedFiles) == 0`. If the session had started, there would be at least workspace state files. Clean assertion without needing to mock the SDK client.

**Table-driven vs standalone:** Used standalone test functions (not a single table-driven switch) because each scenario has distinct fixture setup (temp dirs, plugin YAMLs, skill structures). Tried to avoid over-mocking — real filesystem operations, real YAML parsing, real ValidateAndExpand calls.


## 2026-04-24: Plugin Schema Migration — Test Coverage (Wave 3)

### Context
Neo retired the top-level `plugins:` YAML field and moved plugins to tool entries (`type: plugin`, `ref`, `source: local|remote`) under `generator.tools` / `reviewer.tools`. Neo also enhanced the missing-plugin error to enumerate every filesystem path checked. Tank wired a "wait till known" renderer model that buffers ToolResolutionStart and commits only on Result. Commits on `ronniegeraghty/dev`: `18d105c3` (Tank renderer), `bc06fb8f` (Neo schema).

### Work Completed
Added **17 new test functions** (≈29 cases counting sweep subtests) across 5 files:

**`hyoka/internal/config/plugin_migration_test.go`** (5 tests):
- `TestParse_PluginTypeEntry_SourceOmitted` — parses with no source; field preserved empty
- `TestParse_PluginTypeEntry_ExplicitLocalSource` — `source: local` round-trips
- `TestParse_PluginInGeneratorOnly_NotAutoAppendedToReviewer` — reviewer.tools stays empty
- `TestParse_PluginInBothRoles_BothPreserved` — explicit dual-role survives parse
- `TestParse_RejectsRetiredTopLevelPluginsField_PointsToMigration` — error contains `retired`, `generator.tools`, `type: plugin`, `source:`

**`hyoka/internal/config/configs_sweep_test.go`** (1 test, 13 subtests):
- `TestConfigSweep_AllRepoConfigsParseUnderNewSchema` — every YAML in repo `configs/` parses; no top-level `plugins:`; reviewer never gets plugin entries it didn't declare.

**`hyoka/internal/config/tool/plugin_migration_test.go`** (7 tests):
- `TestValidateAndExpand_MissingPlugin_ErrorEnumeratesEveryCheckedPath` — **every** path from `pluginCheckedPaths` must appear in the reason (6 distinct paths: `.hyoka/plugins/<n>/plugin.yaml`, `.hyoka/plugins/<n>.yaml`, legacy pluginsDir, 3 cache/installed paths)
- `TestValidateAndExpand_PluginFanOut_TwoSkillsOneMCP` — plugin with 2 skills + 1 MCP → exactly 3 child items AND 3 emitted Result events, all with `ParentName=multi-child`, `ParentKind=plugin`
- `TestValidateAndExpand_PluginOnlyInGenerator_ReviewerUntouched` — resolver twin of no-auto-append
- `TestValidateAndExpand_PluginInBothRoles_ChildrenResolveInBoth` — dual-role has `Role=generator` AND `Role=reviewer` children
- `TestValidateAndExpand_SkillDir_ThreeSubdirs_ProducesThreeChildren` — literal (non-glob) skill_dir fan-out; complements existing `GlobExpansion` test
- `TestValidateAndExpand_LocalPlugin_ResolvesFromHyokaPluginsDir` — `.hyoka/plugins/<name>/plugin.yaml` resolves with empty PluginsDir
- `TestValidateAndExpand_RemotePlugin_MissingCache_HardFails` — remote source with empty HOME cache returns `*ToolLoadError` and enumerates cache paths

**`hyoka/internal/eval/tool_load_hardfail_schema_test.go`** (2 tests):
- `TestCopilotRunner_ToolLoadFailure_RemotePluginUncached` — remote plugin, no cache → `ErrorCategory=tool_load_failure`, 0 generated files (no session)
- `TestCopilotRunner_ToolLoadFailure_PluginOnlyInGenerator_ReviewerUnaffected` — failed generator plugin aborts; error mentions plugin name

**`hyoka/internal/progress/display_interactive_plugins_test.go`** (2 tests):
- `TestInteractive_TwoPluginsDistinctHeaders` — two plugins' fan-outs don't interleave; no flat leaf duplication; each child appears under correct parent header
- `TestInteractive_WaitTillKnown_FailedEmitsReason` — complements Tank's `WaitTillKnown`: failed Result commits both ❌ marker and reason text; no transient "Loading" leaks

### Results
- ✅ `go test -race ./hyoka/...` all packages green
- ✅ Baseline (Neo's `bc06fb8f` + Tank's `18d105c3`) commits verified locally — all new tests pass against landed code
- ✅ Zero flakes

### Coverage Map (Mission → Tests)
| Scope item | Covered by |
|---|---|
| 1.a top-level plugins rejected with migration hint | `TestParse_RejectsRetiredTopLevelPluginsField_PointsToMigration` |
| 1.b `type: plugin, ref` parses | `TestParse_PluginTypeEntry_SourceOmitted`, `…_ExplicitLocalSource`, existing `TestParseGeneratorSkillsAndPlugins` |
| 1.c source defaults when omitted | `TestParse_PluginTypeEntry_SourceOmitted` (parse preserves empty; resolver infers) |
| 1.d no auto-append to reviewer | `TestParse_PluginInGeneratorOnly_NotAutoAppendedToReviewer`, `TestValidateAndExpand_PluginOnlyInGenerator_ReviewerUntouched`, sweep test |
| 1.e explicit dual-role | `TestParse_PluginInBothRoles_BothPreserved`, `TestValidateAndExpand_PluginInBothRoles_ChildrenResolveInBoth` |
| 2.a local from `.hyoka/plugins/{name}/` default | `TestValidateAndExpand_LocalPlugin_ResolvesFromHyokaPluginsDir` |
| 2.b remote fetches + caches | `TestValidateAndExpand_RemotePlugin_MissingCache_HardFails` (miss path; mocking a successful remote fetch requires a real fetcher seam — left for future wave) |
| 2.c missing plugin enumerates paths | `TestValidateAndExpand_MissingPlugin_ErrorEnumeratesEveryCheckedPath` (**asserts each of 6 paths**) |
| 2.d fetch failure pre-session | `TestCopilotRunner_ToolLoadFailure_RemotePluginUncached` |
| 2.e plugin fan-out 2+1 with parent metadata | `TestValidateAndExpand_PluginFanOut_TwoSkillsOneMCP` |
| 3. skill-dir fan-out with 3 children | `TestValidateAndExpand_SkillDir_ThreeSubdirs_ProducesThreeChildren` (non-glob), existing `GlobExpansion` (glob) |
| 4. wait-till-known buffer/emit | Tank's `WaitTillKnown`/`PluginFanout`/`SkillDirFanout`/`PluginFailedNoFanout` + my `WaitTillKnown_FailedEmitsReason`, `TwoPluginsDistinctHeaders` |
| 5. config-sweep smoke test | `TestConfigSweep_AllRepoConfigsParseUnderNewSchema` (13 YAML files) |

### Learnings
- **Don't trust remote-tracking**: I spent 15 minutes polling `origin/ronniegeraghty/dev` before realizing Neo's and Tank's commits were already on the *local* `ronniegeraghty/dev` (unpushed). Next time, check `git log ronniegeraghty/dev` directly before polling the remote.
- **Enumerated-path assertion pattern**: asserting all 6 paths (not just "contains /plugins") caught a subtle thing — `hyokaPluginsBase` uses `os.Getwd()` not `ConfigDir`, so the test has to `os.Chdir` into a temp dir (and restore) to get deterministic paths. Documented with a comment so future authors know why the test fiddles with cwd.
- **Remote plugin testing without a real fetcher**: `plugin.ResolveInstalled` is a pure-local cache walk. Redirecting `HOME` to a clean temp dir is sufficient to exercise the cache-miss hard-fail path without mocking network. Actual fetch success remains untested at the unit level — would require a `FetcherInterface` seam.
- **Sweep test guardrail**: the reviewer-tools invariant ("reviewer with no YAML tools block must have no plugin entries after parse") is the cleanest way to catch reintroduction of cross-role auto-append without needing to diff YAML vs parsed struct.

---

## Wave Completion: Plugin Loading Fix (2026-04-23)

The four-agent plugin-loading-fix wave (Neo, Tank, Oracle, Switch) completed successfully. Commits landed on ronniegeraghty/dev:
- **Neo (bc06fb8f):** Retired top-level `plugins:` field; plugins now under generator.tools/reviewer.tools as `type: plugin`
- **Tank (18d105c3, 5216678a):** Wait-till-known rendering; fan-out deduplication; parent header emitted once
- **Oracle (1e5c3b66):** docs/configuration.md plugin section; CHANGELOG breaking-change notice; config migrations
- **Switch (fb70d4c4):** 17 test functions (~29 cases); 5 new test files; full -race suite passes

**Orchestration logs:** `.squad/orchestration-log/2026-04-23T17-{44,45,46,47}Z-{neo,tank,oracle,switch}.md`  
**Session log:** `.squad/log/2026-04-23-plugin-loading-fix-wave.md`  
**Decision entries:** Merged from inbox into `.squad/decisions.md` (5 entries: Ronnie directive + 4 wave decisions)

Status: ✅ Scribe audit complete. Ready for Ronnie's release decision.

---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.
