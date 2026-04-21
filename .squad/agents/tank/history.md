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

- **Config migration pattern**: When removing backward compatibility code, the best approach is to update all configs first (all 8 YAML files), then delete legacy struct fields, then remove helper methods (Normalize, Effective*), then update all call sites. This ensures compiler errors guide you to every place that needs updating.
- **Test-driven refactoring**: Large structural changes benefit from running tests after each phase (struct changes, method deletions, call site updates). The test failures become a checklist of what still needs updating.
- **Unused function cleanup**: After removing legacy fields, helper functions like resolveSkillsDirs() that worked with those fields become unused. The compiler catches these with "declared and not used" errors.


Initial setup complete. Platform is well-structured. Quick wins: fix stale path, plan main.go refactor.

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
