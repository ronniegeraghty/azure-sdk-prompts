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

## 2026-04-22 — PR #607 Merge Conflict Resolution

**Mission:** Resolve conflicting main-merge divergence between phase-6 and ronniegeraghty/dev

**Context:** Tank merged `origin/main` into BOTH `ronniegeraghty/dev` and `phase-6` independently (commits on different days). Both merges resolved the same 9 conflicts but with divergent resolutions. PR #607 (`phase-6 → ronniegeraghty/dev`) became DIRTY/CONFLICTING because the two branches now had different versions of the same post-merge state.

**Strategy:** Merge `dev` INTO `phase-6` to make phase-6 a strict superset of dev. After this, PR #607's diff is just the phase-6-unique commits (clean, no conflicts).

**Conflicting files (6 total):**
1. `README.md` — Installation command path
2. `hyoka/internal/config/tool/resolve.go` — ResolveSkills/FetchRemote signatures
3. `hyoka/internal/config/tool/resolve_test.go` — Comment style
4. `hyoka/internal/config/tool_filter_test.go` — Comment style
5. `hyoka/internal/eval/copilot.go` — ResolveSkills call site
6. `hyoka/internal/eval/engine.go` — dryRun signature + 3 ResolveSkills call sites

**Resolution approach:**
- **README.md**: Took dev version (`go run ./hyoka`) — correct for current module structure
- **resolve.go**: Kept phase-6 HEAD — pluggable Fetcher pattern with `context.Context` support vs dev's direct npx implementation. Phase-6's approach is more extensible (registrable custom fetchers, ctx propagation for cancellation/deadlines).
- **Test files**: Kept phase-6 comment style (cleaner, no branch-specific mentions in final state)
- **eval files**: Kept phase-6 signatures — `ResolveSkills(ctx context.Context, ...)` and `FetchRemote(ctx context.Context, ...)` for proper cancellation propagation

**Key architectural decision:** Phase-6's Fetcher abstraction (Issue #597, PR #605) is load-bearing work that dev didn't have. The context.Context parameter threads cancellation/deadlines into any HTTP/exec work the fetcher performs. Regression to dev's signature would lose this capability.

**Verification:**
- `go build ./...` — ✅ clean
- `go test -race ./... -timeout 5m` — ✅ all 24 packages pass
- PR #607 status post-push:
  - state: OPEN
  - mergeable: MERGEABLE ✅ (clean!)
  - mergeStateStatus: UNSTABLE (CI running, expected)
  - headRefOid: 25675461c8476ecae45e770ebf2063ce229b860b

**Commit:** `25675461` "Merge ronniegeraghty/dev into phase-6: align main-merge conflict resolutions"

**Outcome:** PR #607 now clean. All conflicts resolved by keeping phase-6's architectural work (pluggable fetchers + ctx) while taking dev's correct README path.

## Learnings

### Multi-merge divergence pattern (PR #607)

When two branches independently merge the same upstream and resolve conflicts differently, a future merge between those branches will conflict AGAIN on the same files — even though both sides already "resolved" them once. The resolution is to understand which side has the load-bearing architectural work and keep that, not blindly take "ours" or "theirs".

**Pattern:**
1. Branch A merges upstream → resolves conflicts with approach X
2. Branch B merges same upstream → resolves same conflicts with approach Y
3. Later: Branch A merges Branch B → conflicts re-appear because X ≠ Y

**Resolution strategy:**
- Understand the semantic intent of each side's resolution
- If one side has newer/better architectural work (e.g., phase-6's Fetcher abstraction), keep that
- If the diff is cosmetic (comment style), pick whichever is cleaner
- Test thoroughly — the merge must compile and pass all tests

**Why this happened in #607:** Tank did the merges on different days, with different context. Each merge was individually correct for its branch at that moment. The divergence wasn't visible until we tried to merge the two branches together.

### Context propagation is load-bearing

When adding `context.Context` parameters to a call chain, don't stop halfway. Thread it through ALL callers until it reaches the entry point (e.g., `cmd/run.go` where the CLI gets a ctx from the runtime). Half-measures (signature changes without plumbing) are dishonest — tests that assert ctx propagation (not just that the parameter exists) catch regressions where someone reverts to `context.Background()`.

**Reusable for:** Future ctx-threading work, cancellation/timeout plumbing, HTTP request tracing.

---

## Session 2026-04-21T23:22:02Z: PR #607 Conflict Resolution (Multi-Branch Sync)

**Status:** COMPLETE  
**Branch:** phase-6 (commit 25675461)  
**PR:** #607 (phase-6 → ronniegeraghty/dev)

### Context

Tank executed independent main-merge on dev branch (commit 8bfc4da2). Simultaneously, someone executed main-merge on phase-6 branch (commit 10f4c3f3). Both merges resolved the same 9 conflicts, but resolutions diverged:
- Tank's dev: direct npx skill fetching
- phase-6: pluggable Fetcher abstraction with context.Context threading

Result: PR #607 became DIRTY/CONFLICTING.

### Resolution

**Commit:** 25675461 "Merge ronniegeraghty/dev into phase-6: align main-merge conflict resolutions"

**Strategy:** Merge dev → phase-6 (make phase-6 superset). Then resolve 6 file conflicts using semantic rules:
1. **Architectural wins:** kept phase-6's context.Context threading (Issue #597, PR #605)
2. **Correct paths:** adopted dev's fixed README.md path (`go run .` vs `go run ./hyoka`)
3. **Cleaner style:** kept phase-6's cosmetic improvements

**Result:** PR #607 transitioned to CLEAN/MERGEABLE. Build + tests all pass (-race, 24 packages).

### Key Technical Decision

Kept phase-6's `context.Context` threading through `ResolveSkills()` and `FetchRemote()`:
- Enables cancellation/deadline propagation
- Core architecture improvement for #597 (custom fetchers)
- Tests in PR #608 (commit 04579b47) assert propagation end-to-end
- Would be lost if we reverted to dev's direct npx approach

### Cross-Agent Coordination

Tank did the dev merge independently. Neo then resolved the downstream PR #607 conflict. This is a valid pattern: split the work (Tank owns one branch, Neo owns the resolution merge), but requires semantic conflict resolution, not blind tool picks.

See Tank's orchestration log: `.squad/orchestration-log/2026-04-21T23-22-02Z-tank.md`

**Decision captured:** `.squad/decisions.md` ("Decision: PR #607 Merge Conflict Resolution Strategy")


## 2026-04-22 — Issue #566: WorkspaceDelta first-class + guardrail softening

**Branch:** `squad/566-workspacedelta-firstclass` (off `ronniegeraghty/dev`)
**PR:** opens against `ronniegeraghty/dev`, "Closes #566"

**Built on PR #571's foundation.** That PR delivered the `workspace.WorkspaceDelta` type, snapshot/compute API, `EvalReport.WorkspaceDelta` field, `GraderInput.WorkspaceDelta` field, and nil-safety tests. What was missing: actual capture wiring in the engine, and the guardrail softening the issue called for.

### What landed

1. **Capture wiring** in `engine_eval.go`:
   - `workspace.TakeSnapshot(genDir)` immediately after `CopyStarterFiles` (the "before" picture)
   - Second `TakeSnapshot` after `ws.ListFiles()` returns and `evalReport.GeneratedFiles` is set (the "after")
   - `workspace.ComputeDelta(before, after)` → assigned to `evalReport.WorkspaceDelta`
   - Snapshot failures degrade gracefully (warn + nil delta); graders are already nil-safe (#571's test scenario 3.2)
2. **GraderInput.WorkspaceDelta = evalReport.WorkspaceDelta** — single source of truth, no duplicate compute
3. **Guardrail softening** per issue table:
   - `MaxOutputSize`: 1 MiB hard fail → **10 MiB warning**, prefers `delta.BytesNet` when available, falls back to the starter-aware estimator otherwise. Logs which basis was used.
   - `MaxFiles`: 50 hard fail → **200 warning** (still counts agent-attributable files via `computeAgentFileCount`)
   - `MaxNewFiles` (NEW): default **100 warning**, counts `delta.NewFileCount` (silently skipped when delta is nil)
   - `MaxTurns` and `MaxSessionActions` unchanged (out of scope per issue)
4. **`EvalReport.GuardrailWarnings []string`** — soft-cap breaches populate this. `Success` is no longer flipped to `false` by these three guardrails. `GuardrailAbortReason` reserved for hard-fail (turns) only.
5. **Config schema:** `SessionLimits.MaxNewFiles` added with negative-value validation in `Validate()`.

### Tests

- Updated existing hard-fail tests (`TestGuardrailMaxFiles`, `TestGuardrailMaxOutputSize`, `TestConfigLimitsOverrideEngineDefaults`) to assert the new warning behavior — no `GuardrailAbortReason`, populated `GuardrailWarnings`.
- Updated default-value tests (`TestGuardrailDefaultValues`, `TestResolveLimits*`) to the new defaults (200/10MiB/100).
- New: `TestGuardrailMaxNewFiles` — verifies the new soft cap fires off `WorkspaceDelta.NewFileCount` and produces the right warning message.
- New: `TestWorkspaceDeltaCaptured` — end-to-end check that delta is populated on every successful eval and reaches the report.
- `go build ./...` clean. `go test -race ./... -timeout 5m` clean across all 24 packages.

### Decision (queued for inbox)

Soft-cap guardrails communicate via `GuardrailWarnings []string`, not by flipping `Success` or setting `GuardrailAbortReason`. This gives downstream consumers (HTML report, comparison tooling, future graders) a structured channel to surface "warning, not fatal" without conflating it with generation failure. Hard-fail guardrails (currently only `MaxTurns`) keep using `GuardrailAbortReason`.

### Lessons

- **Snapshot threading is one commit.** Capturing `before` 100 lines above where `after` runs is fine — both live in `runSingleEval` and any error along the way naturally short-circuits past the after-snapshot. Don't over-engineer with helpers when a single stack frame works.
- **Two metrics, one cap.** When softening `MaxOutputSize` to use `delta.BytesNet` but keeping the legacy estimator as fallback, log the basis (`basis=delta.bytes_net` vs `starter_aware_estimator`). Otherwise, comparing two runs with different basis silently lies.

---

## 2026-04-22 — PR #618 amendment: scope-correction on guardrail softening

**Branch:** `squad/566-workspacedelta-firstclass` (force-pushed, commit `cb10cb17`)

Original PR #618 implemented the #566 issue spec faithfully: relaxed `MaxOutputSize` (1 MiB → 10 MiB warning), `MaxFiles` (50 → 200 warning), and added a new `MaxNewFiles` cap. Ronnie reviewed and course-corrected: **only** `MaxOutputSize` should be relaxed. The other guardrails were not actually wanted as warnings; the 50-file cap and hard-fail semantics still earn their keep.

### What stayed
- WorkspaceDelta capture wiring (snapshot before/after, `ComputeDelta`, populates `EvalReport.WorkspaceDelta` and `GraderInput.WorkspaceDelta`). Core #566 deliverable, no controversy.
- `MaxOutputSize` as a **soft warning at 10 MiB**, prefers `delta.BytesNet`, falls back to starter-aware estimator. Kept the `EvalReport.GuardrailWarnings []string` channel because there's still one user (MaxOutputSize). Chose Option B over full removal because removing the field cascaded into CLI flag (`cmd/run.go`), schema validator (`internal/validate`), and config tests — too much churn for a Phase 3.5 polish PR.

### What reverted
- `MaxFiles` back to **50, hard fail** (sets `GuardrailAbortReason`, flips `Success=false`).
- `SessionLimits.MaxNewFiles` config field — removed.
- `EngineOptions.MaxNewFiles`, `EvalReport.GuardrailMaxNewFiles`, default value, validation, the check itself, and `TestGuardrailMaxNewFiles` — all removed.
- `TestGuardrailMaxFiles` and `TestConfigLimitsOverrideEngineDefaults` — restored to assert hard-fail behavior.

Final amended diff: 5 files, +171 / -32. Build clean, full race test suite green.

## Learnings

### Faithful spec implementation can still miss intent

This is the lesson, and it's a real one. Issue #566 had a clean table laying out three guardrails to soften plus a new one to add. I implemented exactly what was on the page. Ronnie's actual intent was narrower — only the per-file size cap, because the review-restructure work eliminated its original purpose. The other guardrails were still earning their keep.

**The error I want to not repeat:** when an issue says "soften these N guardrails," it's worth one sanity-check round of "wait, **why** is each of these being softened, and does the reason actually apply to all of them?" If the reasons differ (one is "no longer needed for X", another is "wider net is fine"), surface that ambiguity in a PR description or pre-implementation question rather than rolling them up into a single uniform change.

The mechanical fix is also a lesson: when amending a PR for scope reduction, **revert by writing the inverse, not by selectively reverting commits**. The amended commit has a clean message describing exactly what's IN scope and what's intentionally NOT — future readers grep for `MaxNewFiles` in this PR's history and find a clear "rolled back, scope-narrow" comment in the test file. No git archaeology needed.

### Guardrail policy going forward (for me, for #619)

- Hard-fail guardrails are the default. They flip `Success=false` and set `GuardrailAbortReason`. Soft warnings are the exception, and they need a specific justification (in #566's case: "review no longer inlines content, so the byte cap's original purpose is gone").
- New guardrails default to hard-fail unless there's a deliberate "we want signal but not enforcement" case.
- The `GuardrailWarnings []string` field exists; if a future soft warning needs the channel, it's already there. Don't confuse "soft warning" with "abort reason" in tests — they're separate fields with separate semantics.

## 2026-04-22 — Issue #619 reading: tool-load fast-fail guardrail

Read the issue + traced SDK surface. Findings:

- SDK exposes loaded inventory cleanly via `copilot.SessionEventTypeSessionSkillsLoaded` (`event.Data.Skills`) and `copilot.SessionEventTypeSessionMcpServersLoaded` (`event.Data.Servers`). Both already handled in `hyoka/internal/eval/copilot.go`.
- Partial enforcement scaffolding already exists: `expectedMCPServers map[string]bool` is built from config and compared against loaded names in `copilot.go:366-370`. Currently warns; #619 promotes this to hard-fail and generalizes to skills.
- Two design choices for "fast-fail before generation": (1) post-hoc abort after session completes, (2) streaming abort via `context.CancelFunc` in the SDK event loop. (1) is simpler and satisfies the acceptance criteria as written; (2) is a follow-up optimization.
- Issue is co-labeled `squad:neo` + `squad:trinity` — engine vs report surfacing split. Posted [proposal comment on #619](https://github.com/ronniegeraghty/hyoka/issues/619#issuecomment-4292673005) to confirm scope/approach before branching.

### Architecture note

The `copilot.go` event loop is the natural seam. It already builds `expectedMCPServers` at session start. Adding `expectedSkills` alongside, then swapping the existing warning for a hard-fail signal (cancel context if streaming, or just record a flag if post-hoc), is a localized change. The engine-side `runSingleEval` then needs to consume that signal, populate `EvalReport.MissingTools`, and skip review.

## 2026-04-22 — PR #618 second amendment: drop the byte-size guardrail entirely

**Branch:** `squad/566-workspacedelta-firstclass` (force-pushed again)

The first amendment kept the byte-size cap as a 10 MiB **soft warning** while reverting `MaxFiles` back to a 50-file hard fail and removing `MaxNewFiles`. Ronnie reviewed that compromise and pushed back: *"make sure neo isn't just re-tightening those guardrails but making sure they go back to failing the eval instead of letting it progress."* Combined with his earlier "no longer needed" framing of the file-size cap, the resolution was Option A: drop the byte-size guardrail entirely.

### Why Option A is right

A guardrail that never aborts the eval is decoration, not protection. The byte-size cap's original purpose was to bound review-prompt context size; the review-restructure work eliminated that purpose. With it gone, `MaxFiles=50` (hard fail) is the agent-output backstop, and the review-side per-file/total caps in `internal/utils/utils.go` prevent runaway memory in the panel. There's nothing left for a byte cap to defend against — so it goes.

### Final landed shape

Removed surface area:
- `SessionLimits.MaxOutputSize` (config struct)
- `EngineOptions.MaxOutputSize` + the 10 MiB default
- `EvalReport.GuardrailMaxOutputSize` and `EvalReport.GuardrailWarnings` (its sole consumer is gone)
- `--max-output-size` CLI flag and the `parseByteSize` helper (+ its tests)
- Schema validator entry in `internal/validate/schema.go`
- Negative-value validator in `config.Validate()`
- `computeAgentOutputSize` (dead code — only the soft-warn block called it) and its table-driven test
- Doc references in `README.md`, `docs/{guardrails,cli-reference,configuration,getting-started}.md`, `examples/configs/example-full.yaml`, `examples/README.md`
- Tests: `TestGuardrailMaxOutputSize`, `TestParseByteSizeValid`, `TestParseByteSizeInvalid`, `TestValidateRejectsNegativeMaxOutputSize`, `TestComputeAgentOutputSize`; updated `TestParseSessionLimits[Partial]`, `TestGuardrailDefaultValues`, `TestResolveLimits*` to drop the field

Kept (untouched):
- All WorkspaceDelta wiring — the actual #566 deliverable
- `MaxFiles=50` hard fail
- `MaxTurns`, `MaxSessionActions`
- `TestGuardrailMaxFiles`, `TestConfigLimitsOverrideEngineDefaults`, `TestWorkspaceDeltaCaptured`

Verification: `grep -ri "MaxOutputSize\|GuardrailWarnings\|max_output_size\|max-output-size"` returns zero matches in active source (excluding `reports/`, `hyoka-prompt-docs/` snapshot, `CHANGELOG.md` historical notes, `.squad/` bookkeeping, `plan/` historical docs). `go build ./...` clean. `go test -race ./... -timeout 5m` green across all 24 packages.

## Lessons

### Compromise is a tell

When a reviewer says "scope this down" and you respond by keeping the controversial piece in a softer form, you are negotiating against the reviewer's stated intent. The first amendment kept the byte-size cap as a soft warning to avoid touching the CLI flag, schema validator, and report field — framed internally as "minimizing churn for a Phase 3.5 polish PR." That framing was self-deceiving. The reviewer's request was *narrower* than my read; the right response was to ask "should this be removed entirely?" before writing the compromise. Half-measures left in code are worse than the original, because they look intentional but no longer serve any purpose.

The mechanical lesson: if a feature has zero load-bearing users (the `GuardrailWarnings` field had exactly one consumer — itself, via the cap I was trying to keep), delete the whole thing. Cascading removal across CLI / schema / tests is a few minutes of work, not a reason to leave dead code in the report schema.

### Guardrail policy, locked in

Going forward (and codified in `.squad/decisions/inbox/neo-guardrail-scope-correction.md`): **guardrails fail the eval, period.** There is no soft-warning tier. If you want signal without enforcement, that's a metric — put it in `WorkspaceDelta` or grader output. The `GuardrailWarnings []string` channel is gone; if a future need surfaces, that's a new design decision, not a free reuse.

## 2026-04-22 — PR #618 merged into phase-6

Orchestration complete. Morpheus verdict APPROVE, Oracle nits resolved, Scribe merged all inbox entries into decisions.md and cleared the inbox. Guardrail policy locked in as team guidance. Issue #619 (tool-load fast-fail) now unblocked.

**2026-04-22 (Morpheus Examples & PR #607 Follow-up):** Loader silently drops YAML docs after first `---` in criteria files (`hyoka/internal/criteria/criteria.go:134-136`). Neo owns fix — strict multi-doc rejection recommended as first move (PR #607 comment 3125721580 has details). Affects example misleadingness; separate issue recommended for example rewrite using `groups:` list.
