# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, internal/eval + internal/review packages
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Core Context

Agent Neo initialized as Core Engine architect. Charter: evaluation pipeline, review orchestration, criteria system, feature flags. Expertise: eval/engine, review/graders, wiring-layer design, #587 regression prevention.

### Condensed History (Phase 0–4)

**Phase 0 (2026-04-03):** Team initialization. Neo tasked with critical path WorkspaceDelta (#566) + comparison unification (#357) + hierarchical review modes.

**Phase 1–3 (2026-04-04 → 2026-04-17):** Shipped #566 WorkspaceDelta (2-day hard cap), #355/#356 multi-bucket review infrastructure, #357 comparison unification. Encountered #587 regression trap (tests pass, runtime behavior absent). Recovered via wiring-layer integration tests.

**Phase 4 (2026-04-17 → 2026-04-20):** Hierarchical criteria system (#356), comparison-result struct contract, custom tool fetcher with version override (#597), remote-skill caching. All approvals and merges completed.

**Phase 5 (2026-04-20):** Implemented --review-mode isolated flag (#580) enabling multi-bucket sessions. Tests excellent at unit level; wiring layer gap discovered by Switch (exact #587 trap). Locked out per reviewer-protocol; Tank fixed wiring tests.

**Key pattern:** Wiring-layer regression tests (integration via engine.Run with stubs) are mandatory for flag-driven feature work. Unit tests alone insufficient.

## Recent Sessions


Decision: .squad/decisions.md | Orchestration Log: .squad/orchestration-log/2026-04-17T20:53:40Z-morpheus.md

### 2026-04-20 (Phase 5 Wrap-up — Morpheus Arch Review)

**Status:** Phase 5 PR #592 approved with followups for Phase 6.

**For Neo:** Three follow-up issues (#594, #595, #596) identified for Phase 6 scope:
- #594: Remove backup test files (.backup, .test suffix)
- #595: Unify dashboard/prompts fetch pattern (frontend concern, but may affect report data structure)
- #596: Refine `isTestValue()` heuristic (affects schema validation in #369)

**Next:** Phase 6 planning will prioritize these based on dependency graph. Morpheus's review is in `.squad/reviews/phase-5-arch-review-2026-04-20T200455Z.md`.


### 2026-04-20: Issue #580 — Review Session Splitting (PR #603)

**Status:** PR #603 open against `phase-6`.

Re-implemented `--review-mode isolated` after the dead-flag revert (PR #587 nuked PR #578 because the flag had no runtime effect). This time the flag actually splits Copilot review sessions.

Design:
- `criteria.GraderEntry.Isolate` + `criteria.GraderGroup.Isolate` (group wins over per-grader)
- `criteria.BuildReviewBuckets` returns 1 bucket (combined, default) or N buckets (isolated mode)
- `review.MultiBucketReviewer` + `MultiBucketPanelReviewer` interfaces; `PanelReviewer.ReviewPanelBuckets` runs `bucket_count × panel_model_count` sessions
- `mergeBucketResults` prefixes per-bucket criterion names with `[bucket-name]` so deterministic any-fail voting across panel models stays unambiguous
- Engine warns via slog when isolated mode is requested but nothing is marked — no silent dead flag
- Combined mode is byte-identical to today (single string path preserved when `len(buckets) ≤ 1`)

## Learnings

- **Commit early and often when other agents may be active.** Mid-task another agent's worktree operation appeared to swap the branch in my main repo path, dumping all my unstaged edits to existing files. Untracked files survived; tracked-edit losses cost ~30 minutes to re-apply. Now: every passing test boundary → commit.
- **Heredoc/create tools occasionally drop content silently.** Always verify with `ls -la` / `grep` after non-trivial writes.
- **Cross-package type duplication is sometimes the right call.** `Bucket` lives in `criteria`, `review`, and `graders` because `graders → review` already exists and the engine builds buckets from `criteria`. Centralizing would create a cycle. Document the duplication and convert at boundaries.
- **Make resurrected dead flags observably alive.** Adding a `slog.Warn` when `--review-mode isolated` is requested but no graders are marked is the cheap insurance against the same regression that killed PR #578 (flag wired but no runtime effect).

## 2026-04-21: #599 — Prompt `group` property

**Branch:** `ronniegeraghty/issue-599-group-property` (off phase-6)
**Worktree:** `/home/rgeraghty/projects/hyoka-599`

Added optional top-level `group` frontmatter field. Backwards compatible (empty = ungrouped). Kebab-case validation: `^[a-z][a-z0-9]*(-[a-z0-9]+)*$`, max 64 chars. Enforced in both `validate.ValidatePromptStruct` (struct path) and `validate.validatePrompt` (CLI path — these two paths exist in parallel; remember to keep them in sync). Propagated to reports via `EvalReport.PromptMeta["group"]` (no schema bump). Tests in `prompt/group_test.go` and `validate/group_test.go`. Site work deferred to Trinity (#600). Decision: `.squad/decisions/inbox/neo-issue-599-group-property.md`.

**Note:** parser.go, types.go, validate.go in this codebase are not gofmt-clean (no leading whitespace on body code). Resisted the urge to gofmt — kept PR diff focused on the issue. If we ever decide to gofmt the codebase, do it as a separate PR.
### Session 2026-04-21 (#597 — WI-027 Tool Versioning & Custom Fetcher)

**Status:** COMPLETE
**Branch:** `ronniegeraghty/issue-597-tool-versioning`
**PR:** (filed against phase-6)

Built on WI-026 (#334) cache infrastructure. Added pluggable `tool.Fetcher`
interface in `internal/config/tool/`, process-wide `DefaultRegistry` preloaded
with the existing npx behavior wrapped as `npxFetcher`, and a
`tool_version_override` map at the ConfigFile level.

**Key files:**
- `hyoka/internal/config/tool/fetcher.go` (new) — Fetcher interface, Registry,
  npxFetcher, ValidateFetchers
- `hyoka/internal/config/tool/resolve.go` — FetchRemote rewritten to dispatch
  through the registry (one-line lookup + delegation)
- `hyoka/internal/config/tool/entry.go` — added `Version` field
- `hyoka/internal/config/config.go` — `ToolVersionOverride map`,
  `ApplyVersionOverrides()`, called from `Load`; LoadDir merges with
  conflict-detection
- `hyoka/cmd/run.go` — `tool.ValidateFetchers` pre-flight before session start
- Tests: `fetcher_test.go` (registry, custom fetcher runtime invocation,
  error propagation, version path); `version_override_test.go` (per-entry
  precedence, idempotency, YAML parse)
- `docs/configuration.md` — new "Tool Versioning & Custom Fetchers" section

**#587-trap guard:** `TestCustomFetcherInvokedAtRuntime` registers a real
mock against the real DefaultRegistry, calls ResolveSkills, asserts call
count == 1 and the version the fetcher saw matches what was configured.
Wiring is observable, not just parseable.

**Backward compat:** zero behavior change without explicit config. Default
fetcher matches every remote skill same as before; only diff is cache path
now segments by version (`.skills-cache/default/...`).

**Tests:** `go test -race ./hyoka/... -timeout 3m` clean. Vet clean.

## Session 2026-04-21 (Phase 6 Round-1: #603 Request Changes + Reviewer-Protocol Lockout)

**Mission:** PR #603 (Review session splitting, #580) test review — ended with LOCKED OUT reassignment

**Context:** #603 implements `--review-mode isolated` flag enabling multi-bucket review sessions. Unit-level tests of `BuildReviewBuckets` (14 tests in `criteria/buckets_test.go`) are excellent.

**Switch's Finding:** Wiring-layer completely untested on 4 surfaces:
1. `Engine.reviewBuckets()` — bridge from cmd-flag to engine-mode
2. `PromptReviewGrader.gradeSingle/gradeWithPanel` branch selection — determines which path fires for `len(buckets) ∈ {0, 1, 3}`
3. `mergeBucketResults` — prefixes per-bucket criterion names
4. CLI flag validation — `cmd/run.go:289-293` validates `--review-mode` value

This is exact failure mode of #587: tests pass, runtime behavior absent.

**Verdict:** REQUEST CHANGES. Per reviewer-protocol, original implementer (Neo) locked out of revisions.

**Outcome:** Tank reassigned + implemented fix (16 tests, 22 subtests). Switch re-reviewed ✅ APPROVE. Commit 04579b47.

**Status:** #603 approved pending merge. Neo locked per protocol; Tank's wiring tests land in PR.
