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

### Phase 6 CLI Invocation Convention (2026-04-21)

**Note:** As of Phase 5, main.go was moved to repo root. All documentation should use:
```bash
go run . <command>     # ✅ CORRECT
```

NOT:
```bash
go run ./hyoka ...     # ❌ STALE
```

Oracle discovered 47 stale references during phase-6 docs audit (fixed in commits b5c4782c–874bedf9). All team members should use `go run .` going forward when writing code, tests, or documentation. This applies to examples in doc files, shell scripts, and test setup code.

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

## 2026-04-21: #608 — PR #605 Fetcher Polish (PR #612)

**Branch:** `ronniegeraghty/issue-608-605-fetcher-polish` (off phase-6)
**Worktree:** `/home/rgeraghty/projects/hyoka-608-605-fetcher-polish`
**PR:** #612 (targets phase-6)

Six review follow-ups from PR #605, all in one focused commit:

1. `TestValidateFetchers` — added no-fetcher branch. Trick: the default npx fetcher accepts every remote skill, so the error path was unreachable from tests. Solution: `DefaultRegistry.Unregister(defaultFetcherName)` with `t.Cleanup` re-registration. Asserts error string names the offending repo.
2. Renamed `TestNpxFetcher_VersionInPath` → `TestNpxFetcher_CanFetchAndName`. The old name described behavior the body never asserted.
3. Deleted `SortedFetcherNames` (zero callers).
4. `FetchRemote(ctx, entry, baseDir)` and `ResolveSkills(ctx, entries, baseDir)` — threaded ctx through the eval engine (`dryRun`, `runSingleEval`, `buildSessionConfig`). New `TestFetchRemote_ContextPropagates` uses a probe fetcher that captures the ctx and asserts caller-set values survive. Guards against silent regression to `context.Background()`.
5. `npxFetcher.Fetch`: removed the `fmt.Printf` that duplicated the `slog.Info` line.
6. `ValidateFetchers([][]Entry)` → `ValidateFetchers([]Entry)`. Flat slice; call site in `cmd/run.go` now appends instead of double-nesting.

**Lessons:**
- **Unreachable error paths want a test seam, not new API.** Temporarily unregistering the default fetcher hit the branch without exporting a hook.
- **Thread ctx in one commit.** Adding ctx to `FetchRemote` alone would have been dishonest — `ResolveSkills` would still have called `context.Background()` and the plumbing would stop one level up. Signature changes that touch transitive callers are cheaper done together.
- **Signature-change tests:** adding a test that asserts ctx propagation (not just that the ctx parameter exists) converts "the caller passes ctx" into "the callee receives ctx" — the same wiring-layer discipline #603 required.

## 2026-04-22 — PR #611 architectural review (sub for Morpheus)

✅ APPROVE (posted as comment — self-approval blocked on shared `ronniegeraghty` account; Morpheus authored, Squad reviewer-author isolation triggered substitution).

PR #611 closes #608's systemic embedded-asset-freshness item with two layers: `Makefile` (`site-install`/`site-build`/`site-embed`/`verify-embed`) + `.github/workflows/site-embed-freshness.yml` (CI gate running `make verify-embed` on site-touching PRs). Verified against `.squad/skills/embedded-asset-freshness/SKILL.md` — implements exactly the two Phase-7 candidates the skill called out (make target + CI hash-diff). `site-embed` correctly avoided a `make build` umbrella that would shadow `go build`. README:222-223 makes the target discoverable. Triggers (PR `**`, push `main`/`phase-6`/`ronniegeraghty/dev`) and paths filter both correct.

Non-blocking follow-ups noted: (1) `ci.yml`'s `site-build-and-test` job already does `npm ci && npm run build` — `verify-embed` could fold in there to remove duplicate ~1-2 min of work; trade-off is losing the discrete PR status check; (2) prune `phase-6` from push triggers post-merge; (3) add `concurrency:` group.

## 2026-04-21 — PR #614 architectural review (substituting for Morpheus)

Reviewed Morpheus's systemic follow-up to my #611 nits: site-embed-freshness CI hardening (concurrency, untracked detection, wholesale prune, phase-6 trigger removal). Substituted on author/reviewer isolation.

**Verdict:** ✅ Approve with notes (posted as `--comment` — `--approve` blocked because Ronnie's gh account = PR author from GitHub's view; verdict explicit in body).

**Key calls:**
- **Keep-the-duplication decision:** agreed. Discrete required check name ("Verify embedded site bundle is fresh") gives reviewers a louder gateable signal than burying inside `site-build-and-test`. Cost is ~1–2min parallel; documentation invites future reconsideration. Sufficient — next reviewer won't re-litigate.
- **`git status --porcelain` vs `git diff --quiet`:** strictly more correct (catches untracked new asset filenames from vite content-hashing). Residual gap: `.gitignore`'d files in EMBED_DIR are invisible to both old and new check — combined with `//go:embed all:site` (the `all:` prefix embeds dotfiles/underscore-files), a gitignored file would silently ship. Pre-existing, not a regression. Defense-in-depth fix would be `--ignored` flag, but not worth it now.
- **`rm -rf $(EMBED_DIR)/*`:** aligned with skill intent. Wholesale prune handles future vite outputs (favicons/manifests) without Makefile edits. Assumption (no hand-maintained files in EMBED_DIR) is verifiable from tree and load-bearing-explicit in comment.
- **Push triggers (`main` + `ronniegeraghty/dev`):** correct for now. Inline comment self-explains the phase-6 removal and the future-pruning pattern.
- **Concurrency key (`workflow-ref`):** correctly disambiguates PR runs (unique `refs/pull/N/merge`) from branch pushes (`refs/heads/<branch>`); cross-PR isolation preserved; cancel-in-progress correctly limited to same-ref pushes.

**Follow-up I named (non-blocking):** the embedded-asset-freshness skill's "refresh procedure" prose (~line 57) is now stale — points at manual `rm -rf assets && cp -r` instead of `make site-embed`/`make verify-embed`. Hygiene PR for someone, separate from this.

**Pattern reinforced:** when a reviewer flags 3 nits and the author addresses 2 plus documents the third's tradeoff in-source, that's the right shape — don't relitigate the documented one unless a new fact emerges. The cost of the duplication (~2min CI minutes per site PR) is real but bounded; the signal-clarity benefit (named required check) is durable.
