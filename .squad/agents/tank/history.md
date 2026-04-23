# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/main.go (CLI entry), configs/ (run configs), reports/ (output), site/ (docs/serving), go.work (workspace)

## Core Context

Agent Tank initialized as Platform Dev for hyoka. Owns CLI, config, build, reports, site, and plugins. The CLI supports: `list` (show prompts), `run` (execute evaluations with filters like --service, --language, --all-configs). Smart path detection checks ./prompts then ../prompts. Fan-out confirmation prompts at >10 evals. Uses go.work workspace for multi-module builds.

## Recent Updates

📌 Team initialized on 2026-04-03

📋 **Morpheus Audit (2026-04-03):** Audit of CLI and platform layer complete. Key findings: (1) **stale path in main.go:1277** (P0) — new-prompt output references `go run ./tool/cmd/hyoka validate`, should be `go run ./hyoka validate`. (2) **main.go refactor candidate (P1)** — 1329 lines, split into cmd/ package recommended. See `.squad/decisions.md` for full list.

## Learnings

- **Agent Attempt Tail-Streaming Deleted (2026-04-23):** After four failed attempts to fix line-wrapping leaks (commits 6b3d3d48, 42ea88fb, fe6efebf, 670c5dbf), replaced the streaming tail with a three-state machine: Running, Completed, or Guardrail hit — {reason}. The single-line bounded status eliminates all wrapping math and foreign-write contamination risks. Commit b17f1ef5.

- **Config migration pattern**: When removing backward compatibility code, the best approach is to update all configs first (all 8 YAML files), then delete legacy struct fields, then remove helper methods (Normalize, Effective*), then update all call sites. This ensures compiler errors guide you to every place that needs updating.
- **Test-driven refactoring**: Large structural changes benefit from running tests after each phase (struct changes, method deletions, call site updates). The test failures become a checklist of what still needs updating.
- **Unused function cleanup**: After removing legacy fields, helper functions like resolveSkillsDirs() that worked with those fields become unused. The compiler catches these with "declared and not used" errors.


Initial setup complete. Platform is well-structured. Quick wins: fix stale path, plan main.go refactor.

### Phase 6 CLI Invocation Convention (2026-04-21)

**Note:** As of Phase 5, main.go was moved to repo root. All examples should use:
```bash
go run . <command>     # ✅ CORRECT
```

NOT:
```bash
go run ./hyoka ...     # ❌ STALE (Phase 5 regression)
```

Oracle audited phase-6 docs and found 47 stale references across 4 files — fixed in commits b5c4782c–874bedf9. Tank should ensure all new examples follow the `go run .` pattern in feature work, CLI help text, and test setup.

### Session 2026-04-04T00-05 (Morpheus Evolution Plan)

Evolution plan assigns you Phase 0 CI pipeline (P0), main.go split, YAML prompts, session limits, .hyoka directory. Read `.squad/decisions.md` for full plan. Also assigned: config validation, duplicate detection, stale path fixes.

### Session 2026-04-04T03:28 (Issue Creation)

Created 72 GitHub issues (#91–#162) for evolution plan across 5 phases:
- Phase 0 (CI pipeline): #91–#99 (9 issues)
- Phase 1 (CLI & config): #100–#119 (20 issues)
- Phase 2 (Report & skill system): #120–#137 (18 issues)
- Phase 3 (SDK & validation): #138–#143 (6 issues)
- Phase 4 (Extensibility): #144–#151 (8 issues)
- Phase 5 (Polish & tools): #152–#162 (11 issues)

All issues labeled, assigned, and staged for backlog prioritization. Backlog is fully populated.

### Session 2026-04-04T19:30 (CI Pipeline Implementation)

**Issue #91:** Created GitHub Actions CI workflow (`.github/workflows/ci.yml`)
- Triggers: All PRs + pushes to main/ronniegeraghty/dev
- Go 1.26.1 (matches go.mod), ubuntu-latest
- Steps: build → vet → test with `-race` flag
- 2-minute timeout (tests run in ~5s with race, 24× headroom)
- Race detection required from day one per D-AUTO-DM8 (concurrent code in ResourceMonitor, ProcessTracker, PanelReviewer)
- Verified all commands pass locally before pushing
- Branch: `ronniegeraghty/issue-91-ci-pipeline`
- PR: #168 → ronniegeraghty/dev

**Key learnings:**
- CI (task 0.1) is explicit blocker for Phase 1 per D-AUTO-DM17 — nothing merges until CI is green
- Race detector adds ~4-5s overhead to test runtime but catches concurrency bugs early
- Phase 0 keeps it simple: no golangci-lint yet (deferred to Phase 1)
- File path: `.github/workflows/ci.yml` (34 lines, YAML)
- **CI verified working:** First run on PR #168 passed in 1m3s (build + vet + test with -race)
### Session 2026-04-04T19:46–19:47 (Phase 0 Execution — CI + Config Migration)

**Status:** COMPLETE  
**Tasks:** 2 PRs

**PR #168 — CI Pipeline (#91):**
- Created .github/workflows/ci.yml with build, vet, test-race
- ~2-minute timeout for full test suite
- Race detection enabled to catch data races early
- CI passed on first run

**Cross-agent impact:** Enables Switch's test reliability work. Combined with Switch's event-driven pattern fix (#99, PR #167), the full test suite now passes consistently under -race detection.

**PR #171 — Config Migration (#96):**
- Refactored all 8 configs to Generator/Reviewer sub-structs
- Removed Normalize() and Effective*() getters
- Updated ~17 call sites atomically
- Net -178 lines

**Cross-agent impact:** Unblocks Neo's reviewer factory fix (#92, PR #170). Clean Generator/Reviewer schema makes per-config reviewer selection viable. Files: all 8 YAML configs, config/ package, multiple call sites.

**Learnings:** Atomic big-bang migrations work when compiler errors guide all needed updates. Test failures provide a checklist.


## 2026-04-16 — Phase 3 Merged to Dev (Neo)

Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has both Phase 3 features and starter-aware guardrail fix. All tests pass, CI green.

## Learnings — example-remote-skill PR (#573)

- `examples/configs/*.yaml` are NOT auto-loaded. The default loader reads `configs/` only (see `hyoka/cmd/run.go` + `internal/config/LoadDir`). To use an example config, pass `--config-file examples/configs/<name>.yaml` **plus** `--config <config-name>` (the `name:` field inside the YAML). Document this in the config's comment header so users copy the right invocation.
- `--dry-run` **does** execute the remote-skill fetch path (`internal/skills/fetcher.go::fetchRemote` → `npx skills add`). The "Fetching remote skill: …" line + repo clone both happen during dry-run. Good for wiring tests without eating Copilot sessions.
- Known gap: `fetchRemote` passes `--name` but not `--yes`, so `npx skills add` still prompts interactively for skill selection. Under a non-TTY (CI, piped stdin) the prompt consumes no selection and 0 skills are resolved. The dry-run still exits 0. Filing as a follow-up if it blocks real usage.
- `--all-prompts` does **not** exist. Use `--prompt-id <id>` for single, or filter flags (`--language`, `--service`, `--plane`, `--category`, `--tags`). A `--language python` filter on the full prompt set = 22 evals, which trips the >10 confirm prompt and hangs on non-TTY stdin; always narrow for dry-run smoke tests.

## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release

Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-generation, serve endpoints, hierarchical criteria, cleanup. Recommendation: **Promote dev → main and cut v0.3.1 tag.**

Decision: .squad/decisions.md | Orchestration Log: .squad/orchestration-log/2026-04-17T20:53:40Z-morpheus.md

### 2026-04-20 (Phase 5 Wrap-up — Morpheus Arch Review)

**Status:** Phase 5 PR #592 approved with followups for Phase 6.

**For Tank:** Three follow-up issues (#594, #595, #596) identified for Phase 6 scope:
- #594: Remove backup test files (.backup, .test suffix) — platform layer concern
- #595: Unify dashboard/prompts fetch pattern
- #596: Refine `isTestValue()` heuristic (affects schema validation in #369)

**Next:** Phase 6 planning will prioritize these based on dependency graph. Morpheus's review is in `.squad/reviews/phase-5-arch-review-2026-04-20T200455Z.md`.

### Session 2026-04-20 (CLI Help Scrub — Issue #364)

**Status:** COMPLETE (commit db93f408, pushed to phase-5, CI PASS)

Scrubbed code-gen-specific framing from CLI help text and Go doc comments per Oracle's README audit (2208bfcb). Replaced 18 instances across 14 files:
- `code generation quality` → `output quality`
- `generated code` → `agent output`
- `generating code` → `producing output`

Surfaces touched: root/run/new-prompt commands, graders, eval, review, config, logging, trends, serve packages.

**CI:** Build, Vet, Test + Site Build — all PASS.

**Learning:** sed is the right tool for Unicode-safe string replacement when edit tool can't match escape sequences like `\u2014` (em-dash). Always verify CLI help output end-to-end with `go run . --help` before pushing.


### Session 2026-04-21 (R77 — Configurable prompt_directory, #598)

**Status:** COMPLETE — PR pending push
**Branch:** `ronniegeraghty/issue-598-configurable-prompt-dir` → `phase-6`

**What:** Added top-level `prompt_directory:` YAML field to `ConfigFile`. New `ResolvePromptDirCandidates` and `PeekPromptDirectory` helpers in `internal/config/discovery.go`. Wired `run`, `validate`, `list`, `serve` to consult the config-driven path with priority: `--prompts` flag > config `prompt_directory:` > `.hyoka/prompts/` > `./prompts/` > `../prompts/`.

**Backwards compatible:** field is optional; absent → identical behavior to today.

**Tests:** 11 new tests in `prompt_dir_test.go` covering Load relative/absolute path resolution, LoadDir conflict detection, ResolvePromptDirCandidates ordering, PeekPromptDirectory edge cases.

**Verification:** End-to-end smoke test in `.scratch-598/` confirmed default-init path, config-driven override, and `--prompts` flag-overrides-config.

**Learnings:**
- The repo's hard-coded `prompts` discovery flowed through 4 commands (`run`, `validate`, `list`, `serve`), each with slightly different ordering of "load configs vs resolve paths". `run` had to be reordered (configs first, then prompts dir). `validate`/`list`/`serve` use `PeekPromptDirectory` to extract just the field without doing strict YAML validation, so a malformed config doesn't break commands that should be tolerant of it.
- **Concurrent worktree hazard:** Neo was running #580 in the *same* worktree (no separate `git worktree`), which kept dropping unrelated `criteria/buckets.go` and `eval/engine.go` files into my staging area. Worked around by `git restore`-ing/deleting the leftovers before each build cycle and being explicit about which files to `git add`. Future spawns should always create a worktree per agent.
- `gopkg.in/yaml.v3` doesn't strictly validate fields when target struct only declares the keys you care about (no `KnownFields(true)`), which makes `PeekPromptDirectory` safe even on configs with extra/unknown keys.
- The example config emitted by `hyoka init --with-examples` is **already broken** (flat `name:` instead of nested `configs:`). Pre-existing bug, unrelated to this issue — left for a separate fix.

## Session 2026-04-21 (Phase 6 Round-1: #602 Approval + #603 Wiring Tests Reassignment)

**Mission:** Implement wiring-layer test fixes for PR #603 (reassigned from Neo per reviewer-protocol)

**Context:** Switch requested changes on PR #603 because wiring-layer coverage gap on 4 surfaces (despite excellent unit-level `BuildReviewBuckets` coverage). Same failure mode as #587: tests pass, runtime behavior absent. Per strict reviewer-protocol, Neo locked; Tank reassigned.

**What was added:** 16 new tests (22 subtests) closing all 4 gaps:
- `internal/eval/engine_reviewbuckets_test.go` — 5 unit tests, 0%→100% line coverage; combined/isolated/degraded paths + slog warn capture
- `internal/eval/engine_reviewmode_runtime_test.go` — 2 integration tests via engine.Run with stub reviewers proving flag has runtime effect
- `internal/graders/prompt_review_grader_buckets_test.go` — 3-row table + fallback + error
- `internal/review/buckets_test.go` — prefix rules, aggregation, nil-safety
- `cmd/run_validate_test.go` — validator + flag wiring + invalid-rejection

**Coverage deltas:** review 48.6%→53.5%, graders 79.9%→82.9%, cmd 42.4%→42.6%, eval 54.5% (reviewBuckets 0%→100%).

**Switch re-review:** ✅ APPROVE. Commit 04579b47, ready to merge.

**Implication for future:** Wiring-layer regression tests (integration through engine.Run with stubs) now standard pattern for any flag-driven feature work — cheapest defense against "tests pass, behavior gone" failure mode.

**Status:** PR #603 approved. Phase 6 Round-1 test batch complete.

### Session 2026-04-21 (Phase 6 #608 — PR #606 Group Property Polish)

**Status:** COMPLETE — PR #610 → phase-6
**Branch:** `ronniegeraghty/issue-608-606-group-tests`

**What:** Three test-only additions closing coverage gaps Morpheus called out on PR #606 (group frontmatter property):

1. **Observable-wiring test** (`hyoka/internal/eval/engine_group_wiring_test.go`) — runs `engine.Run` with `StubRunner`, asserts prompt's `Group` reaches `EvalReport.PromptMeta["group"]` at engine_eval.go:78-80. Covers propagation + empty-omitted case. #587-trap pattern.
2. **Regex boundary rows** (added to `hyoka/internal/validate/group_test.go`) — ~35 new rows: 63/64/65-char limits including hyphenated forms, whitespace variants, hyphen-only, consecutive hyphens, digit-only segments, special chars, non-ASCII, emoji, null bytes.
3. **JSON omitempty round-trip** (`hyoka/internal/prompt/group_json_test.go`) — verifies `json:"group,omitempty"` on `prompt.Prompt.Group`: absent when empty, round-trips when set, cleared-group remarshal still omits.

**Verification:** `go test -race -timeout 3m ./hyoka/...` → all packages PASS.

**Learnings:**
- The #587-trap observable-wiring pattern is straightforward to apply at engine.Run level — `StubRunner` + asserting on `summary.Results[0].PromptMeta` exercises the real metadata-build code path without any generator/reviewer stubbing gymnastics. Reuse this recipe for any future "frontmatter field → report" wiring check.
- When adding boundary rows to an existing anonymous-struct table test, keeping rows terse (one-liners with trailing comments) scales better than refactoring to named rows — the existing tight style was fine for ~55 rows total.

### Session 2026-04-21 (Main Sync — dev + phase-6)

**Status:** COMPLETE  
**Branches:** `ronniegeraghty/dev` (commit 8bfc4da2), `phase-6` (commit d111c964)

**What:** Merged `origin/main` (12 commits) into both `ronniegeraghty/dev` and `phase-6` branches to sync missing commits from main. Main's tip was 7aa917a1, which had the older `hyoka/main.go` structure (pre-PR #300 restructure).

**Conflict resolution pattern:**
1. **Main.go location** — always keep the newer structure (`main.go` at repo root). Main tried to move it back to `hyoka/main.go`; rejected that and kept root structure.
2. **hyoka/internal/ paths** — rejected all main's changes in `hyoka/internal/...` since those paths don't exist on dev/phase-6 anymore (everything moved to `internal/...` in the restructure). These were stale paths from before PR #300.
3. **SkippedReviewers field** — main added `SkippedReviewers []review.SkippedReviewer` to `EvalReport` and markdown rendering. Manually ported this field to dev/phase-6 `hyoka/internal/report/types.go`.
4. **Test signature mismatches** — main added a 4th return value (`skipped []SkippedReviewer`) to `ReviewPanel()`, but dev/phase-6 still use 3-value signature. Fixed all test call sites to match our current signature.
5. **Missing test functions** — main added tests for `parseRepoSpec()` and `Branch` field validation, but these functions/fields don't exist on dev/phase-6. Disabled those tests with comments.
6. **criteria/language/rust.yaml** — took main's version (has more specific guidance on obsolete crates).

**Verification:** `go build ./... && go test -race ./... -timeout 5m` — all PASS on both branches.

**Part 3 — Docs Cleanup:** Converted all `docs/*.md` files from source-dev command form (`go run . <cmd>`) to installed-binary form (`hyoka <cmd>`). Per directive in `.squad/decisions/inbox/copilot-directive-2026-04-21T22-58-docs-installed-binary.md`, docs are for end users who installed the tool, not for contributors building from source.

**PR #607 status:** State OPEN, mergeable CONFLICTING, mergeStateStatus DIRTY. Expected — both branches now have the same main commits but via different merge paths. CI not yet triggered.

**Learnings:**
- When main and feature branches diverge on a major restructure (e.g., `hyoka/main.go` → `main.go` at root), always keep the side with the newer structure. The older structure's paths (`hyoka/internal/...`) will generate conflicts but those files don't exist anymore — reject them wholesale.
- If main adds a new field to a shared struct, hand-port it to the feature branch even if the paths differ. Build errors guide you to all the spots that need updating (e.g., `SkippedReviewers` in `EvalReport` required adding the field + updating markdown rendering).
- Test signature changes (return value count) require fixing all call sites. Use sed for bulk replacements (`sed -i 's/old_pattern/new_pattern/g'`) when the pattern is uniform across many test files.
- Disabled tests should have a clear comment explaining why and what condition would re-enable them. Keeps the intent clear for future readers.
- The `--force-with-lease` push is safe after amending a merge commit to include forgotten files — it only force-pushes if the remote still matches your pre-amendment state.
### Session 2026-04-21 (Main Sync — ronniegeraghty/dev)

**Status:** COMPLETE  
**Branch:** `ronniegeraghty/dev` (commit 8bfc4da2)

**What:** Merged `origin/main` (12 commits behind) into `ronniegeraghty/dev` to sync missing commits. Main's tip was 7aa917a1, which had the older `hyoka/main.go` structure (pre-PR #300 restructure).

**Conflict resolution:**
- **Main.go location** — kept dev's structure (`main.go` at repo root), rejected main's `hyoka/main.go`
- **hyoka/internal/ paths** — rejected all main's changes in stale `hyoka/internal/...` paths (moved to `internal/...` in restructure)
- **SkippedReviewers field** — manually ported from main to `hyoka/internal/report/types.go` + markdown rendering
- **Test signature mismatches** — fixed ReviewPanel call sites (3 return values on dev vs 4 on main)
- **Missing test functions** — disabled tests for `parseRepoSpec()` and `Branch` field (don't exist on dev)

**Verification:** `go build ./... && go test -race ./... -timeout 5m` — all PASS.

**Learnings:**
- When main and feature branches diverge on a major restructure, always keep the side with the newer structure. The older structure's paths will generate conflicts but those files don't exist anymore — reject them wholesale.
- If main adds a new field to a shared struct, hand-port it even if paths differ. Build errors guide you to all update spots.
- Use sed for bulk test fixes when pattern is uniform: `sed -i 's/old_pattern/new_pattern/g'`


---

## Session 2026-04-21T23:22:02Z: Main Sync and Docs Installed-Binary

**Status:** COMPLETE (Part A, Part B via Neo)  
**Branch:** ronniegeraghty/dev (commit 8bfc4da2)  
**User request:** Pull main into dev; switch docs/ to installed-binary command form

### Part A: Merge origin/main into dev

**Commit:** 8bfc4da2 "Merge main into ronniegeraghty/dev: pull in 12 missing commits from main"

- Resolved 9 merge conflicts independently
- Kept dev's modern structure and call signatures
- Result: dev 13 commits ahead of main
- Build ✅ Tests ✅

### Part B: Docs installed-binary conversion

**Commit:** d111c964 "docs: switch docs/ examples to installed-binary command form"

- 28 occurrences of `go run . ` → `hyoka ` in docs/getting-started.md
- Verified no other docs files had source-dev commands
- Rationale: docs/ is for users (installed binary), not contributors
- Source-dev commands live in CONTRIBUTING.md

### Cross-Agent Coordination

Neo performed Part C: resolved PR #607 conflict by merging dev into phase-6. Tank's independent dev merge diverged from phase-6's simultaneous merge on the same 9 conflicts. Neo's resolution kept phase-6's pluggable Fetcher + context.Context threading (architectural win) while adopting dev's corrected documentation paths.

**Key learning:** Multi-branch independent merges of the same upstream produce divergent resolutions. Resolution requires semantic understanding, not tool automation.

See Neo's orchestration log: `.squad/orchestration-log/2026-04-21T23-22-02Z-neo.md`

**Decisions captured:** `.squad/decisions.md` — docs installed-binary directive + PR #607 strategy


## Team Context: Unified Grader Direction Proposed (2026-04-22)

Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE schema, ONE execution path
- **Backward-compat:** Existing `criteria/*.yaml` files work without migration
- **Phased rollout:** 4 phases, zero-regression guarantee via golden-file tests
- **Your role:** Phase 1-3 implementation track (schema, execution, cleanup) — likely paired with Neo

📄 See `.squad/decisions.md` "Unified Grader Architecture Direction & Proposal" for full spec and phased plan. Awaiting team consensus and architecture sign-off.

---

## 2025 — workers default → 1

Flipped `--workers` default from `runtime.NumCPU()` (capped at 8) to `1` in `engine.go`. Updated help text in `run.go`. Kept the >8 clamp for explicit user values.

### Learnings
- `EngineOptions.Workers` default was set inside `NewEngineWithReviewerFactory` (not at flag definition time), so the CLI flag is `0` and the engine substitutes the default. Changing the substituted default is a one-line change — no flag plumbing needed.
- `runtime` package was only used for `NumCPU()` in engine.go; removing that call let me drop the import cleanly.
- This is the foundational switch for the sprint's interactive-vs-CI mode split driven by worker count.

## 2025 — --progress auto: worker-count-driven selection

Extended `--progress auto` in `hyoka/cmd/run.go` to pick "live" or "log" from the worker count. Commit `d6fd0a59`.

### Learnings
- Final decision sequence inside the `progressMode == "auto"` block reads: (1) explicit live/log/off already wins because we only enter the block for "auto"; (2) non-TTY stdout → "off" via `progress.IsTerminal(os.Stdout)`; (3) `workers > 1` → "log"; (4) default → "live". Then a post-pass downgrades "live" → "log" when debug/info logging would corrupt ANSI unless `--log-file` is set.
- Kept the downgrade as a separate step (not folded into the switch) so the intent — "live is fine except when stderr slog would clobber it" — stays legible. The `--log-file` exception from 3b9cbab9 survives intact.
- `progress.IsTerminal` is already exported from `internal/progress/display.go`, so the TTY probe didn't need a new helper.

## 2025 — fix: --progress auto suppressed CI mode on piped multi-worker runs

Reordered the `--progress auto` switch in `hyoka/cmd/run.go` so `workers>1` short-circuits before the non-TTY check. Regression from d6fd0a59: piped multi-worker runs fell through to "off" and the CI renderer never fired (Switch capstone rows 6 & 8). Extracted the resolution into a pure `resolveAutoProgress(workers, isTerminal, logLevel, logFile)` function and added `TestResolveAutoProgress` covering the 4-combo TTY×workers matrix plus the `--log-file` exception.

### Learnings
- **Case order matters in progress-mode switches**: TTY check must come *after* the worker count, not before. The CI renderer is append-only and is specifically the one that should engage in piped/CI contexts — so "non-TTY" is a single-eval-only signal, not a global suppression. My d6fd0a59 ordering accidentally inverted this.
- **Extract-to-pure-function for case-matrix testability**: factoring the resolution out of the cobra `RunE` closure into `resolveAutoProgress(workers, isTerminal, logLevel, logFile) string` made the regression guard a trivial table-driven test with no cobra/os.Stdout stubbing. Worth doing eagerly whenever a decision depends on >2 inputs.
- **The `--log-file` exception from 3b9cbab9 is fragile**: it lives as a post-pass after the main switch. Preserved it verbatim in the new pure function and added test rows for it so the next refactor can't silently drop it.

## Team Updates

### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three rounds. 48 new test cases. 2 regressions caught by Switch: 1 fixed in-sprint (yours, `2d38533f`), 1 filed as preexisting Known Issue (`hyoka clean` blocks on non-interactive stdin — OPEN, out-of-scope). Ledger discrepancy reconciled: `82cd8590` never merged → re-landed inside `25ce00a7`.

**Your commits this sprint:** `3b9cbab9` default `--workers` to 1 · `27c6a679` history · `d6fd0a59` `--progress auto` by worker count · `2d38533f` hot-fix `--progress auto` order for piped CI.

See `.squad/orchestration-log/2026-04-23T00-05-04Z-sprint-wrap.md` and the round-3/4 section in `.squad/decisions.md`.

## 2026-04-23 — feat: human-friendly slog handler for stdout/stderr

Implemented ConsoleHandler to replace structured slog output on stdout/stderr with human-friendly messages. When logs go to console (no --log-file):
- WARN: `⚠️  <msg> (key=val ...)` with dim attrs
- ERROR: `❌ <msg>` in red, attrs dim
- INFO/DEBUG: suppressed entirely (diagnostic noise on console)

When --log-file is set, structured TextHandler with timestamps is still used for diagnostics. The handler respects NO_COLOR and TTY detection via the existing progress/style package.

Added console_handler.go and console_handler_test.go with 9 table-driven tests covering enabled levels, formatting, colors, NO_COLOR, handler attrs, and WithGroup. All tests pass with -race.

**Commit:** `82fc9750` — feat(logging): human-friendly slog handler for stdout/stderr

### Learnings
- **slog.Handler interface is ergonomic**: Implementing Handle, Enabled, WithAttrs, WithGroup was straightforward. The Enabled check lets you suppress levels cleanly at the handler layer.
- **Style package integration**: The existing progress/style package with its NO_COLOR + TTY detection made color support trivial — just wrap text in styler.Red() or styler.Dim().
- **Handler selection at Setup time**: Choosing handler based on whether --log-file is set keeps the decision in one place (logging.Setup). No need for runtime handler swapping.
- **Snapshot testing with ANSI codes**: Testing ANSI sequences in strings is readable with \x1b[31m etc. inline in want strings. NO_COLOR tests verify by checking for absence of \x1b[ prefix.

### Session 2026-04-23 — CLI Output UX Sprint Round 2 (slog + renderer ownership)

**Console-friendly slog handler shipped** (2026-04-23T01:22Z, commits `82fc9750` + `727a67b0`). Interactive renderer now displays clean human-readable logs instead of JSON-style structured format.

**Scope expansion: Tank now owns `hyoka/internal/progress/`** (commits `0747aa58`, `ce9afc50`). Tank was already listed in routing.md for "progress output," but the charter now explicitly includes `display_interactive.go`, `display_ci.go`, and `style/` helpers. This correction reversed a Sprint 1 misassignment that had put CLI renderers under Trinity's scope; Tank is the CLI operator and owns all CLI-facing output.

**Related:** Neo shipped git-clone skill resolver (neo commit `cf6a7636`) to replace `npx skills add`, unblocking the interactive renderer from stdout pollution. Trinity completed agent-attempt gating (commits `0747aa58` + `ce9afc50`) and handed off future CLI renderer work to Tank — a clean scope separation: Tank = CLI, Trinity = Site/React/Reports.

## 2026-04-23 — fix: tail line wrapping with wide characters (emoji, CJK)

The interactive renderer's multi-row tail clearing logic (commit `6b3d3d48`) had the right structure but used **rune counting** instead of proper terminal **cell width** calculation. Wide characters like emoji 🔄 and ✅ occupy **2 terminal cells** but were counted as **1 rune**, causing off-by-one or more errors in truncation and row-count logic.

### Bug symptoms
When a tail line containing emoji exceeded terminal width:
- `truncateToWidth()` would truncate to N runes, thinking it fit in N cells
- But the actual cell width was N+1 or N+2 (because emoji = 2 cells)
- Line would wrap to 2 physical rows
- `tailRowCount` would compute as 1 (based on rune count)
- `rewriteTail()` would only clear 1 row, leaving the wrapped portion visible
- Result: leaked wrapped content on subsequent rewrites, appearing as scrolling multi-line tail

### Root cause
```go
// OLD (buggy) — counts runes, not cells
func visibleWidth(s string) int {
    stripped := ansiSeqRE.ReplaceAllString(s, "")
    count := 0
    for range stripped { count++ }  // ❌ 🔄 counted as 1, but occupies 2 cells
    return count
}
```

### Fix
**Commit:** `fe6efebf` — fix(progress): use proper cell width for wide characters (emoji, CJK)

- Added `github.com/mattn/go-runewidth` dependency
- `visibleWidth()` now uses `runewidth.StringWidth()` for accurate cell counting
- `truncateToWidth()` uses `runewidth.RuneWidth(r)` per-rune to check display width
- Updated tests: `🔄` = 2 cells, `✅ Loaded` = 9 cells (2 + 1 + 6)
- Added regression tests for wide-char truncation scenarios

### Verification
Tested with narrow-terminal live run (60 cols):
```bash
stty cols 60
go run . run --prompt-id identity-dp-python-default-credential --config "baseline/claude-opus-4.6"
```
Tail stayed on exactly one row throughout the entire evaluation. No leaked wrapped content.

### Learnings
- **Rune vs. cell width distinction is critical for terminal rendering**: Many Unicode characters (emoji, CJK ideographs, some symbols) take 2 terminal columns, not 1. Always use a proper wcwidth-based library like `go-runewidth` for terminal layout math.
- **Multi-row ANSI clearing requires accurate row counting**: The `\x1b[NA` (cursor up N lines) + per-row `\r\x1b[2K` (clear line) pattern only works if you know the exact number of physical rows the content occupied. Off-by-one in row count = leaked content.
- **Test against the actual terminal behavior**: The fix for Bug A (section ordering) was verified by inspection of frozen output. Bug B (wrapping) needed a **live narrow-terminal run** to reproduce — unit tests alone wouldn't have caught the cell-width bug because test assertions used rune counts too.
- **Progress package is now Tank's scope**: As of 2026-04-23 directive, `hyoka/internal/progress/` (including `display_interactive.go`, `display_ci.go`, and `style/` helpers) is owned by Tank, not Trinity. CLI-facing renderers = Tank; site/report HTML = Trinity.


### Session 2026-04-23 (Tail Leak Fix v2 — Foreign Writes + Terminal Width Edge Cases)

**Status:** COMPLETE (commit 670c5dbf)
**Issue:** Multi-line tail leak still occurring after previous runewidth fix (fe6efebf)

**Root causes identified:**
1. **Foreign writes from slog**: When --log-file NOT specified, slog ConsoleHandler writes directly to stderr. Interactive renderer writes to stdout. Both render to same TTY → slog warnings inject lines between tail writes, breaking row-count tracking.
2. **Terminal width edge case**: When tail text exactly == termWidth, cursor wrapping is terminal-dependent (immediate vs delayed), causing potential off-by-one in tailRowCount.

**Fix applied:**
1. Added `SuppressConsole` option to logging.Setup(). When interactive mode active without --log-file, run.go reconfigures logger to use io.Discard (prevents stderr corruption).
2. Modified writeTail/rewriteTail to truncate to `(termWidth - 2)` instead of `termWidth` (avoids cursor wrapping ambiguity). For narrow terminals (<12 cols), skip margin.

**Files changed:**
- `hyoka/cmd/run.go`: Import logging pkg, reconfigure when interactive mode without --log-file
- `hyoka/internal/logging/logging.go`: Add SuppressConsole + io.Discard handler branch  
- `hyoka/internal/progress/display_interactive.go`: Safety margin on truncation (2 cols)

**Verification:**
- All tests pass with -race flag
- Manual test: no warnings appear during interactive mode (when --log-file absent)
- Defensive 2-column margin prevents "exactly terminal width" edge case

**Learnings:**
- **Stdout vs stderr matter**: Even when both write to same TTY, they're separate file descriptors. A renderer tracking writes to one can't know about writes to the other.
- **The "auto" progress mode mystery**: Initially confused because user didn't pass `--progress interactive`. Found that resolveAutoProgress() in run.go auto-selects "interactive" when workers==1 && isTerminal.
- **Pre-existing downgrade logic**: run.go already had logic to downgrade interactive→ci when logLevel is debug/info without --log-file. I extended this to suppress console logging entirely (vs downgrade mode).
- **Terminal cursor wrapping is non-standard**: When you write exactly N chars to N-column terminal, cursor position is ambiguous (column N vs column 0 of next row). Only safe approach: never write exactly termWidth chars.

Related commits: fe6efebf (wide char), 6b3d3d48 (multi-row clearing), 42ea88fb (truncation base)


### Session 2026-04-23 (Agent Attempt Three-State Implementation)

**Status:** COMPLETE (commit b17f1ef5)
**Issue:** Multiple failed attempts to fix tail-streaming line-wrapping leaks

**Root decision:** After four attempts to fix the streaming tail (commits 6b3d3d48, 42ea88fb, fe6efebf, 670c5dbf), the user directed a pivot: delete the tail-streaming code path and replace with a simple three-state machine.

**Three states:**
1. **Running** — Agent session started, shows "🔄 Running" (no streaming content, no duration)
2. **Completed** — Session ended successfully, shows "✅ Completed"
3. **Guardrail hit** — Guardrail terminated the run, shows "Guardrail hit — {reason}" where reason is short-form extracted from `evalReport.GuardrailAbortReason` (e.g., "turn limit (25)", "file limit (50)")

**Implementation:**
- Added `agentAttemptState` enum to `display_interactive.go` (Running, Completed, Guardrail)
- Removed all streaming tail logic: activity counters (`agentToolCalls`, `agentActivity`), duration display, per-event tail rewrites
- Added `GuardrailReason` field to `ProgressEvent` (events.go) to propagate guardrail info from eval to display
- Wired guardrail reason extraction in `engine.go` via `extractGuardrailShortReason()` helper
- Updated `agentComplete()` to accept guardrail reason and set final state
- Removed ticker loop updates for Agent Attempt (ticker still exists but is now a no-op for agent section)
- Updated tests in `display_interactive_test.go` to expect "✅ Completed" instead of "✅ Complete" / "❌ Failed"

**Files changed:**
- `hyoka/internal/progress/events.go` — added GuardrailReason field
- `hyoka/internal/progress/display_interactive.go` — three-state machine + removed tail streaming
- `hyoka/internal/progress/display_interactive_test.go` — updated test expectations
- `hyoka/internal/eval/engine.go` — wired guardrail reason extraction and emission

**Verification:**
- `go build ./...` — clean
- `go test -race ./...` — all pass
- Manual run: `go run ./hyoka run --prompt-id key-vault-dp-python-crud --config "baseline/claude-opus-4.6"` — confirmed Agent Attempt shows "Running" then "Completed"
- `hyoka clean` — cleaned 6 orphaned sessions

**Learnings:**
- **When to give up on streaming UX:** After multiple fix attempts fail on the same root cause (line wrapping from wide chars, foreign writes, terminal width edge cases), a bounded-state display is safer than continued iteration on truncation math. The single-line three-state model eliminates an entire class of bugs by removing variability.
- **State machines > streaming for bounded-duration phases:** Agent Attempt is inherently bounded (starts, runs, ends). Streaming the "inside" is nice-to-have but risky; the states (Running, Completed, Guardrail) carry the essential info without wrapping risk.
- **Guardrail propagation via progress events:** The `evalReport.GuardrailAbortReason` was already being set in `engine_eval.go` but wasn't flowing to the display. Adding `GuardrailReason` to `ProgressEvent` and wiring it through `engine.go` made the guardrail state viable.
- **Ticker simplification:** With no dynamic content in Agent Attempt, the ticker loop became a no-op for that section. Left the ticker infrastructure in place in case we add grader progress indicators later, but removed the agent-specific refresh logic.

**Not tested end-to-end by me:**
- Triggering an actual guardrail (MaxTurns, MaxFiles) to verify the Guardrail state renders correctly. The logic is wired and the reason extraction helper is in place, but I couldn't easily trigger a guardrail in test. User should test: set MaxTurns very low (e.g., 3) via prompt frontmatter or config override, run a prompt, confirm "Guardrail hit — turn limit (3)" appears.

**Next:** User should sanity-check:
1. Run a normal prompt → confirm "Running" → "Completed" transition
2. Trigger a guardrail scenario → confirm "Guardrail hit — {reason}" displays
3. Check that no line wrapping occurs even on narrow terminals (`stty cols 60`)

