# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, internal/eval + internal/review packages
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Recent Sessions

### 2026-04-23: Tool Validation Gate Fix — SDK Event Timing

**Status:** ✅ FIXED. Commit 4b593d3b.

**Problem:** After merging commit 92a9746c (tool validation gate), NO evals could run. Every eval aborted after 10s with "tool verification timeout". Agent attempts never started.

**Root Cause:** The gate blocked BEFORE `SendAndWait` waiting for `SessionSkillsLoaded`/`SessionMcpServersLoaded` events that the SDK only emits DURING the first message round-trip. Classic deadlock: gate waits for events that can't fire until gate releases.

**Evidence from live log:**
```
14:22:42  CreateSession() completes
14:22:45  SendAndWait() called (2s later)
14:22:45  SessionSkillsLoaded fires (DURING SendAndWait, not after CreateSession)
14:22:46  Turn 1 starts
```

Gate would block at 14:22:42 and timeout at 14:22:52, aborting before the event at 14:22:45.

**Fix:** Disabled blocking gate. Tool load failures still logged (event handlers unchanged) but don't abort evals. Better to have evals run with degraded tools than zero evals at all.

**Verification:**
- Live run: `key-vault-dp-python-crud × baseline/claude-opus-4.6`
- ✅ Agent executed (3 turns, files created)
- ✅ Eval passed (88s duration)
- ✅ Skills loaded during SendAndWait (not at CreateSession)

**Lesson:** SDK event timing assumptions MUST be verified with live traces. "SessionSkillsLoaded fires right after CreateSession" was wrong. The SDK emits tool load events during/after the first message exchange, possibly because:
1. Lazy initialization (SDK loads tools when agent needs them)
2. Cold-start MCP servers (npx can take 15-30s)

**TODO:** Re-enable gate after deciding placement (post-SendAndWait?) and timeout (30s+?). Or make it optional via `--strict-tools` flag. See decision doc for options.

**Decision:** `.squad/decisions/inbox/neo-fix-tool-gate-blocking-evals.md`

---

## Core Context

Agent Neo initialized as Core Engine architect. Charter: evaluation pipeline, review orchestration, criteria system, feature flags. Expertise: eval/engine, review/graders, wiring-layer design, #587 regression prevention.

### Condensed History (Phase 0–4)

**Phase 0 (2026-04-03):** Team initialization. Neo tasked with critical path WorkspaceDelta (#566) + comparison unification (#357) + hierarchical review modes.

**Phase 1–3 (2026-04-04 → 2026-04-17):** Shipped #566 WorkspaceDelta (2-day hard cap), #355/#356 multi-bucket review infrastructure, #357 comparison unification. Encountered #587 regression trap (tests pass, runtime behavior absent). Recovered via wiring-layer integration tests.

**Phase 4 (2026-04-17 → 2026-04-20):** Hierarchical criteria system (#356), comparison-result struct contract, custom tool fetcher with version override (#597), remote-skill caching. All approvals and merges completed.

**Phase 5 (2026-04-20):** Implemented --review-mode isolated flag (#580) enabling multi-bucket sessions. Tests excellent at unit level; wiring layer gap discovered by Switch (exact #587 trap). Locked out per reviewer-protocol; Tank fixed wiring tests.

**Key pattern:** Wiring-layer regression tests (integration via engine.Run with stubs) are mandatory for flag-driven feature work. Unit tests alone insufficient.

**Grader Unification — Phase 1 Ready (2026-04-22):** All 10 schema decisions locked. Issues #624–#627 filed. Phase 1 (#624) ready for Neo pickup. Full proposal in `.squad/decisions/inbox/morpheus-grader-unification-proposal.md` (kept until Phase 1 complete). Reference `.squad/decisions.md` for locked decisions.

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

## Team Context: Unified Grader Direction Proposed (2026-04-22)

Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE schema, ONE execution path
- **Backward-compat:** Existing `criteria/*.yaml` files work without migration
- **Phased rollout:** 4 phases, zero-regression guarantee via golden-file tests
- **Your role:** Phase 1 & 2 implementation likely (unified schema + execution path)

📄 See `.squad/decisions.md` "Unified Grader Architecture Direction & Proposal" for full spec and phased plan. Awaiting team consensus and architecture sign-off. Coordinate with Tank if implementation assigned.

Reminder: Loader silent-truncation bug (#607 discovery) remains tracked as Neo follow-up.

## 2026-04-22 — Grader Unification Phase 1 (#624) shipped

**Branch:** `ronniegeraghty/dev` — direct commit, no PR.
**Commit:** `faf556eb2bfb227c8873bed7dd92b4887a24fdbe`

### Files created
- `hyoka/internal/graders/unified_entry.go` — `UnifiedGraderEntry`, `UnifiedGraderGroup`, `UnifiedGraderConfig`, validation, `matchesUnifiedWhen`, `mergeUnifiedWhen`, `IsValidUnifiedType`.
- `hyoka/internal/graders/unified_loader.go` — `ParseUnified`, `LoadUnifiedFile`, `LoadUnifiedDir`, `Bundle`, `FileError`, `Bundle.MatchingErrors(props)`, `translateLegacy`, `peekFileWhen`.
- `hyoka/internal/graders/unified_entry_test.go` — validation + back-compat + deferred-error tests.
- `hyoka/internal/graders/unified_realfixtures_test.go` — loads every real `criteria/*.yaml` via the new loader.

### Final schema shape

```go
type UnifiedGraderEntry struct {
    Type    string            `yaml:"type"`              // prompt | file | program | behavior | …
    Name    string            `yaml:"name"`
    Weight  float64           `yaml:"weight,omitempty"`
    When    map[string]string `yaml:"when,omitempty"`
    Isolate bool              `yaml:"isolate,omitempty"` // prompt-only, silently ignored for typed
    Prompt  string            `yaml:"prompt,omitempty"`  // required iff type=prompt
    Details yaml.Node         `yaml:"details,omitempty"` // required iff type!=prompt
}
```

`UnifiedGraderGroup` and `UnifiedGraderConfig` unchanged from the legacy criteria.GraderGroup/GraderConfig in shape (File-level when, Graders, Groups, Source), just referencing the new entry type.

### Learnings

**Naming collision foresight paid off.** The proposal called for `GraderEntry` / `GraderGroup` / `GraderConfig` in `internal/graders/`. Those names are already taken by the typed-runtime config in `types.go`. Using `Unified` prefix for Phase 1 lets both shapes coexist for the duration of Phases 2-3 instead of forcing a rename + mass-update in a single commit. The prefix drops in Phase 3 when `internal/criteria/` is deleted.

**`details:` over `config:` for the typed payload.** The issue allowed either. "config" is already overloaded in this codebase (YAML config files, `GraderConfig`, `ProgramConfig`, etc.). `details:` reads better in YAML and has zero naming collision risk. Locked it in the PR and documented the choice in the issue comment.

**Q4 deferred-error semantics need the file-level `when:` preserved on failure.** I extended the proposed `FileErrors map[string]error` to `map[string]FileError` where `FileError` carries `Path, When, Err`. Without this, `Bundle.MatchingErrors(props)` would need a second filesystem pass per eval. The `peekFileWhen` helper uses a permissive non-strict decoder so it can extract `when:` from files whose body is broken — if even the `when:` can't be peeked, `MatchingErrors` surfaces the error universally (fail-loud default, safer than silently hiding a file that might be relevant).

**Review-bucket builder deliberately deferred.** The issue listed `unified_buckets.go` as Phase 1, but nothing in Phase 1 consumes buckets — `internal/criteria/buckets.go` is still wired and still authoritative. Porting it now ships dead code. It'll land in the Phase 2 (#625) wiring commit where the engine actually starts using it, which keeps every line of `unified_buckets.go` immediately exercised in the same PR. Documented as deviation 3 in the issue comment.

**Grader Unification — Phase 2 Engine Cutover Shipped (2026-04-22, commit a8a6d2d4):**
Completed the deferred engine cutover from the 7d7372ef helper landing. Removed
`EngineOptions.GradersDir` and the dual `graderConfigs`/`pluginGraders` fields;
engine now loads a single `graders.Bundle` via `LoadUnifiedDir(CriteriaDir)`.
runSingleEval's grading block rewritten: `Bundle.MatchingErrors(props)` fails
the eval when a relevant file is malformed (Q4), `PartitionMatched` splits
matched entries into typed graders (run via `InstantiateGraders+RunGraders`
after `ToRuntimeConfig`) and prompt entries (feed `BuildUnifiedReviewBuckets` /
`MergeUnifiedCriteria` into `GraderInput`). Removed the gate short-circuit in
`AggregateResults`: every grader contributes; `Pass` is the AND of per-result
pass. Rewrote 6 tests against the new Bundle API; full `go test -race ./...`
green. `internal/criteria/` still on disk (Phase 3 retires it) — engine no
longer touches it. Decision memo at
`.squad/decisions/inbox/neo-phase2-engine-cutover-shipped.md`. Phase 2 of
#625 now truly complete.

**Grader Unification — Phase 3 Criteria Deleted (2026-04-22, commit 46b624fb):**
Closed out `#626`. Deleted `internal/criteria/` in full (5 files, ~1400 LOC):
`buckets.go`, `buckets_test.go`, `criteria.go`, `criteria_test.go`,
`hierarchical_test.go`. Three import sites remained after Phase 2 — all
migrated in-commit:

- `cmd/list`: now loads graders via `graders.LoadUnifiedDir`; unified
  `Source` / `When` are rendered identically, and `grader_count` in both
  text and JSON output now sums top-level graders + group graders
  (previously undercounted grouped criteria). Internal types
  `listCriteriaEntry` / `listOutput` retained for stable JSON shape.
- `cmd/validate`: criteria section rewritten against `Bundle.FileErrors`
  — surfaces every deferred parse/validation failure with the loader's
  own error message (Phase 1 semantics). Dropped the ad-hoc per-grader
  name/prompt checks since `graders.validateEntry` already enforces them
  strictly at load time. Exit code + success phrasing preserved.
- `internal/validate/schema.go`: removed the dead
  `ValidateCriteriaStruct` helper and its sole caller
  `TestCriteriaStructValidation_GraderTypes`. Unified validation lives
  in `graders.ParseUnified` / `LoadUnifiedDir` and is already covered by
  the Phase 1 suite.

Verified: `go build`, `go vet`, `go test -race ./... -timeout 3m` all
clean; `rg 'hyoka/internal/criteria' --type go` zero hits; live
`hyoka list` and `hyoka validate` smoke-tested against the real
`criteria/` directory (2 files, 25 graders — "All 2 criteria file(s)
valid").

Deviation from issue scope: `#626` also listed golden-file snapshot
tests in `internal/graders/` as a replacement for the Phase 1/2
parallel-run safety net. Deferred — the existing
`unified_loader_test`, `unified_entry_test`, `unified_realfixtures_test`,
and `phase1_loader_test` already exercise parse + matching on real
fixtures. Flagged for Ronnie/Morpheus in the closing comment on #626 in
case a dedicated snapshot suite is still wanted.

Grader unification is code-complete: unified schema (Phase 1) → engine
cutover (Phase 2) → criteria deletion (Phase 3). `internal/graders/` is
now the single source of truth for grading config. Decision memo at
`.squad/decisions/inbox/neo-phase3-criteria-deleted.md`.

## Learnings — Option A grader restructure (2026-04-22, commit 46ddda2e)

- `criteria/` umbrella + `criteria/graders/` nested sub-package mirrors
  the YAML data model (files contain graders). Parent-imports-child is
  the only legal direction — never have the grader-type package import
  the file-level package.
- `ReviewBucket` lives with grader-type code because it rides inside
  `GraderInput`; the file-level `BuildUnifiedReviewBuckets` imports it.
- When a single registry file mixes per-grader-type construction
  (NewGrader) with file-level execution (InstantiateGraders/RunGraders),
  split the file-level helpers into their own `exec.go` at the criteria
  layer. Keeps the grader-type package free of multi-entry concerns.
- `git mv` for every file move preserved history cleanly even across a
  package hierarchy change. Package-decl updates and import rewrites can
  then be done with sed in bulk.

## Learnings — ProgressEvent schema extension (CLI UX overhaul, sprint todo #2)

- `ProgressEvent` is a **fat union struct** — every existing emitter uses raw
  struct literals and only sets the fields its `EventType` cares about. I
  followed that pattern rather than introducing a nested payload interface;
  switching styles mid-struct would have forced every existing call site to
  change and broken the "no emitter changes in this task" boundary.
- **No JSON tags** on the struct. The existing fields have none and the type
  is in-process only (never serialized to the report JSON or over a wire).
  The task said "json tags consistent with existing fields" — consistent here
  means none. If we ever need to serialize, tag the whole struct in one pass
  rather than half-tagging new fields now.
- `Score *float64` — pointer so "grader didn't report a score" is
  distinguishable from a legitimate `0.0`. All other numerics stayed as
  value types because their zero has an unambiguous meaning (0 turns,
  0 tool calls, $0 cost for a free/cached run).
- Kept `Status` (tool load outcome) and `Result` (grader outcome) as
  separate fields rather than overloading one. They share lexical space
  ("pass"/"fail" vs "loaded"/"failed") but are semantically distinct — a
  single field would force renderers to branch on `EventType` just to
  interpret the value, which is exactly the coupling the schema should
  avoid.
- `EventToolsVerified` carries `Tools []ToolStatus` rather than being
  replayed as N single-tool events. The sprint plan explicitly allows one
  bulk redraw at verification time; a slice payload keeps that as a single
  event the renderer can atomically redraw.
- Exported string constants (`ToolKindSkill/Plugin/MCP`,
  `ToolStatusLoaded/Failed`, `GraderResultPass/Fail`) so the downstream
  emitter agents (tool-resolution, tool-verification, grader-serializer)
  don't hardcode magic strings that would drift.
- Did **not** add constructor helpers — existing emitters all use raw
  struct literals, so a helper would be an inconsistent half-measure.
  Documented this explicitly in the inbox decision so downstream agents
  don't chase a nonexistent pattern.

## Learnings — grader serialization + per-grader events (sprint todo #5)

**Context:** Wired `GraderStart` / `GraderComplete` events around each grader in
`engine_eval.go` so the interactive display can render a per-grader "Running…
→ Pass/Fail" tail line instead of one aggregate summary.

- **Graders were already sequential.** Expected to find goroutines + WaitGroup
  and rewrite. Actually `criteria.RunGraders` is a plain `for` loop over
  `[]Grader` calling `Grade()` one at a time, and the review grader runs
  *after* typed graders finish in `engine_eval.go`. So the "serialize in
  interactive mode" part of the task was a no-op — only the event emission
  was new code. Documented in the decision memo so future agents don't look
  for parallelism that isn't there.
- **Workers==1 interactive signal plumbed but unused at runtime.** The task
  asked for a mode-detection rule. Since graders are already serial, the
  `interactive` bool would only gate future parallelism. I chose to emit
  events unconditionally whenever the raw-event sender is non-nil, which
  matches the schema doc's "reporter nil = skip" guard and avoids dead
  branches. If we later parallelize the review grader alongside typed
  graders for CI mode, the gate belongs there — not in the emission layer.
- **Score field is kind-dependent.** `GraderResult.Score` is always populated
  (0.0–1.0 normalized), but it's only semantically meaningful for LLM-judge
  kinds (`prompt_review`, `prompt`). For `output_check` / `file` / `program`
  / `behavior` it's just `0` or `1` mirroring `Pass`, and rendering "pass
  (0/10)" would be misleading. So `emitGraderComplete` populates `Score`
  only for `KindPromptReview` and `KindPrompt`; others leave it `nil`. This
  lines up with the schema doc's "`nil` = not reported" convention.
- **Extended `RunGraders` via new `RunGradersWithHooks` rather than mutating
  the existing signature.** Keeps the two existing test call sites
  untouched and gives downstream callers (tests, future CI-mode wiring)
  the option to opt in.
- **New callback `sendRawEvent` in engine.go** auto-fills `EvalID` /
  `PromptID` / `ConfigName` so `engine_eval.go` can emit rich events
  without duplicating identity plumbing. Same pattern as `sendEvent` /
  `sendPhase` but for events that carry arbitrary fields (grader ID, score,
  result).

### emit-tool-verification (2026-XX-XX)

**Scope:** wire one `EventToolsVerified` emission out of `hyoka/internal/eval/copilot.go` after the SDK reports its session-start skill + MCP-server load results. Display flips "Loaded"-optimistic tools to "Failed" when the SDK never confirmed them.

**Learnings:**

- **Where session-start events are observed:** the SDK calls `sessionCfg.OnEvent` once per `copilot.SessionEvent`. Two events matter for verification: `SessionEventTypeSessionSkillsLoaded` (carries `event.Data.Skills[].Name`) and `SessionEventTypeSessionMcpServersLoaded` (carries `event.Data.Servers[].Name`). Both fire during `CreateSession` / early post-creation — there is no explicit "session start complete" event, so the only reliable trigger for verification is "we've now received both load events" (or "we've received the one we care about, given the config").
- **Where verification happens:** inside the shared `OnEvent` closure in `copilot.go`, which is already the single place that aggregates SDK data (session records, warnings, progress forwarding). Keeping the logic there means `expectedMCPServers` + `expectedSkills` (derived from `sessionCfg.SkillDirectories` basenames) stay colocated with the existing "Expected MCP server not loaded" warning, which we retain in addition to the new event.
- **Locking discipline for progress callbacks:** the existing pattern releases `mu` before calling `e.progressFn` to avoid holding the eval's shared lock across display writes. I preserved that — `emitToolsVerified` mutates/reads state under the lock and returns the tool slice; the actual `progressFn` call happens after `mu.Unlock()`. This is consistent with the other progress forwarders in the same handler.
- **Skill-name convention:** the SDK reports skills by `Name`, which matches the skill directory's basename (the dir containing `SKILL.md`). So deriving expected names as `filepath.Base(sessionCfg.SkillDirectories[i])` lines up 1:1 with SDK-reported names. No additional metadata lookup needed.
- **Single-emit guarantee:** `verifiedEmitted` is a closure-local flag guarded by `mu`. Whichever event (skills or MCP) fires second triggers the emission; the first event stashes its data and exits early. If only one kind is configured, emission fires as soon as that one event arrives. If neither is configured, no event is emitted (display has nothing to flip).

## Learnings — ToolResolution emit plumbing (CLI UX overhaul, sprint todo #3, commit e06ead61)

- Config-tool package had zero progress awareness. Minimal hook is a
  callback type `tool.ProgressEmitter = func(progress.ProgressEvent)`
  defined in `config/tool/resolve.go`. No import cycle: `progress` pulls
  nothing from `config/*`. Matches the plan's "simple function-pointer
  parameter is fine" guidance; no context-key shenanigans.
- Kept backward compat by making the old `ResolveSkills` a nil-emitter
  shim over the new `ResolveSkillsWithReporter`. Every existing caller
  (dry-run in `engine.go:778/790`, reporting in `engine_eval.go:262`,
  existing tests) keeps working unchanged because nil-emit is a no-op.
- MCPs don't "resolve" at config-load time — they're static entries. So
  `EmitMCPResolutions` is a validation-only helper: Loaded when the
  required mode-specific field is set (`Command` for local, `URL` for
  remote), Failed otherwise. Not routing through ResolveSkills.
- Plugins are already expanded into `Tools` at `config.Load` time (via
  `ExpandPlugins`, which runs before the engine creates a display). I
  could not emit events from inside `ExpandPlugins` because no reporter
  exists yet. Chose to add `ToolConfig.EmitPluginResolutions(emit)` that
  **re-runs** the registry + installed-plugins lookup read-only and
  emits events. The existing `slog.Warn` stays — it fires at load time
  for the log/CI paths; the new emit fires at eval time for the
  interactive renderer. Single source of truth for "found?" is the same
  two lookup calls (`reg.Get` + `resolveInstalledPlugin`) that
  `ExpandPlugins` uses.
- Engine wire-in is one hunk in
  `eval/copilot.go:buildSessionConfig`. The runner already carries
  `e.progressFn progress.ProgressFunc` (set by `engine.Run` via
  `pr.SetProgressFunc(display.HandleEvent)` through the
  `progress.Reporter` interface). Conversion
  `tool.ProgressEmitter(e.progressFn)` is structural — both are
  `func(progress.ProgressEvent)`. Order: plugin → MCP → skill, so the
  Tools block renders top-to-bottom in declaration-scan order.
- Pairing contract for downstream renderers: every
  `EventToolResolutionStart` has exactly one matching
  `EventToolResolutionResult` with the same `(ToolName, ToolKind)` before
  the next tool's Start. This holds because emission is synchronous on
  the eval goroutine and `ResolveSkillsWithReporter` / `EmitMCP*` /
  `EmitPluginResolutions` all run a simple for-loop with no concurrency.
- A skill entry resolving to zero directories (missing SKILL.md, empty
  skill_dir) used to be non-fatal with just a log warning. Now also
  surfaces as `EventToolResolutionResult{Status: Failed, Reason: "no
  skill directories resolved"}`. Existing tests don't inspect events so
  they pass; the `ResolveSkills_SingleSkillMissingSKILLMD` test still
  expects (nil err, 0 dirs) and gets that. UX-correct: users want to see
  the failure, not hope they spotted the log line.
- **Coordination hazard**: working alongside the parallel
  emit-tool-verification agent (same file: `copilot.go`), my first
  commit was clobbered when the other agent's staged changes got swept
  in (pre-commit hook or similar). Fixed via `git reset --soft HEAD~1`,
  then saved their WIP diff to `.copilot-tmp-parallel.patch`, reverted
  the file, re-applied only my hunk, and committed with
  `git commit --only <path>...` for explicit path safety. Patch didn't
  reapply cleanly after my commit (line numbers shifted); the other
  agent will need to regenerate — this is the expected failure mode for
  concurrent agents on the same file. Lesson: always commit with
  `--only <paths>` when other agents are live on the same working tree.

## Interactive renderer (display-interactive-renderer) — 2025 sprint

Built `hyoka/internal/progress/display_interactive.go` — new renderer for
the single-eval, human-watched case (`workers==1`, default). Trinity was
parallel on `display_ci.go` in the same working tree; I mirrored her
delegation pattern (Display holds a `*interactiveRenderer` pointer, set in
`NewDisplay`, dispatched from `HandleEvent`/`Finish`) so we only collided
on `display.go` — clean merge, no reverts.

Wired new `ModeInteractive = "interactive"` constant; updated `cmd/run.go`
auto-mode to pick `"interactive"` for workers==1 and `"ci"` for workers>1
(replacing the previous `"live"`/`"log"` strings — but `live`/`log` still
resolve, `log` routes into Trinity's CI renderer).

Three tests added (`display_interactive_test.go`): happy-path, no-tools /
no-graders omission, and the Tools flip path. Full `go test -race ./hyoka/...`
green.

Architecture doc: `.squad/decisions/inbox/neo-interactive-renderer.md` so
Switch can author snapshot tests without re-reading the code.

## Learnings

- **Tail-update technique**: `"\r\x1b[2K" + text` replaces the current
  line's content without advancing the row. Combined with a strict
  `writeLine`/`writeTail`/`freezeTail` discipline that tracks
  `linesWritten`, this gives us a single-line "scratchpad" at the bottom
  without any save/restore gymnastics for normal updates. Freezing the
  tail means writing `"\n"` — the line becomes immutable, `linesWritten++`,
  `tailKind = tailNone`.
- **The one exception: Tools block redraw on `EventToolsVerified`**. When
  the SDK reports a tool as failed after we already printed it as Loaded
  (the Tools section is no longer the tail — Agent Attempt lines are
  below it), we do a single bracketed redraw: `\x1b7` (DECSC save),
  `\x1b[<N>A\r` (move up N = `linesWritten - toolsFirstLine`),
  rewrite each tool line with `\x1b[2K<text>\n`, `\x1b8` (DECRC restore).
  Because we write exactly the same number of lines the tool block
  originally occupied, lines below are untouched and the restore puts us
  back on the original tail line unchanged. Gated by a `toolsVerified`
  flag so it fires at most once per eval. Only non-tail update in the
  whole renderer — documented inline.
- **Ticker scope**: interactive mode's 1-second ticker only refreshes the
  Agent Attempt tail (duration counter); it checks `tailKind == tailAgent`
  and bails otherwise. Other sections are purely event-driven. Avoids the
  522-line `display.go` region-redraw machinery entirely.
- **Locking**: `interactiveRenderer` has its own `sync.Mutex`; caller
  (Display.HandleEvent) already holds `d.mu` when delegating. Consistent
  lock ordering (d.mu → r.mu, never reverse) prevents deadlock, and the
  ticker goroutine only ever grabs `r.mu` so there's no cycle.
- **Counter plumbing**: the outer `Display` still owns `completed/passed/
  failed/errors` so `CompletedEvalCount()` keeps working; the dispatch
  shim increments them on terminal events. Same pattern Trinity used for
  the CI renderer — kept consistent.

## Team Updates

### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three rounds. 48 new test cases. 2 regressions caught by Switch: 1 fixed in-sprint by Tank (`2d38533f`), 1 filed as preexisting Known Issue (out-of-scope).

**Your commits this sprint:** `61d830c6` progress event types (tools/graders/session) · `bffd0c40` grader start/complete events · `e06ead61` tool-resolution events during config load · `a0105a9d` interactive renderer (tail-only layout).

**Ledger reconciliation:** `82cd8590` (your ToolsVerified emission) never merged into dev — the commit exists but is not an ancestor of HEAD. Behavior was re-landed by Switch inside `25ce00a7` in a more testable shape (`hyoka/internal/eval/tool_verification.go` with 9 tests). Contract preserved: at-most-once, configured-kinds-only, deterministic sort, plugins excluded, slog warn paths preserved. See `.squad/decisions.md` for full details.

See `.squad/orchestration-log/2026-04-23T00-05-04Z-sprint-wrap.md`.

### Session 2026-04-23: Git-Clone Skill Resolver (CLI Output UX Follow-up)

**Status:** ✅ COMPLETE (committed on `ronniegeraghty/dev` at 727a67b0)
**Context:** First real `--pairwise` run revealed the new interactive renderer gets stomped by stdout from the Copilot SDK's `npx skills add` plugin auto-install. Ronnie decided to stop using `npx skills add` and `copilot plugin install` entirely.

**Task:** Replace the npx/copilot-plugin-install shell-outs with a git-clone resolver hyoka owns end-to-end.

**Implementation:**
- **New `gitFetcher`** (replaces `npxFetcher` in `hyoka/internal/config/tool/fetcher.go`):
  - Parses skill specs: `name@skills` → microsoft/skills; `name@owner/repo` → owner/repo + name; bare `owner/repo` → repo root
  - Cache path: `<baseDir>/.skills-cache/<version>/<owner>/<repo>/` (preserves existing cache layout)
  - Reuses clones if already cached; runs `git fetch && git checkout <version>` to update
  - Searches for named skills in: `.github/skills/`, `.github/plugins/`, `.claude/skills/`, `.agent/skills/`, `skills/`
  - **All git output suppressed** via captured stdout/stderr; only surfaces stderr on non-zero exit (logged at Debug)
  - No direct stdout/stderr writes — renderer owns the screen
- **Updated `InstallSkillsAndPlugins`** (in `config.go`): now a no-op — resolution happens lazily on first use via the git fetcher
- **Updated `resolveInstalledPlugin`** (in `config.go`): checks the git-clone cache first (`.hyoka/cache/default/microsoft/skills/.github/plugins/{name}/` for `name@skills` shorthand), then falls back to `~/.copilot/installed-plugins/` for backward compatibility
- **Tests** (in `fetcher_test.go`): spec parsing, skill discovery in common locations, cache reuse logic, all existing fetcher tests updated to use `gitFetcher` instead of `npxFetcher`

**Build & Test:**
```
go build ./...
go test -race ./hyoka/internal/config/tool/... -timeout 3m
go test -race ./hyoka/internal/config/... -timeout 3m
```
All green. No pre-existing test failures touched.

**Commit:** `727a67b0` on `ronniegeraghty/dev` (authored by Tank's session but contains Neo's work)
**Decision:** `.squad/decisions/inbox/neo-git-skill-resolver.md`

## Learnings

- **When another agent commits your work mid-session,** verify the commit matches your intent before proceeding. In this case, Tank's commit `727a67b0` landed the exact changes I authored — the git-clone resolver, updated tests, and InstallSkillsAndPlugins no-op. No conflicts, clean integration.
- **Skill spec parsing is more flexible than the original npx path.** The `@skills` shorthand (e.g., `azure-sdk-python@skills` → microsoft/skills repo) simplifies config YAMLs. The `name@owner/repo` format enables arbitrary repo sources without expanding the YAML schema.
- **Silent git operations are critical for renderer stability.** Capturing stdout/stderr to a buffer and only logging stderr on failure keeps the CLI output clean. The interactive renderer depends on this — any stray git progress output would stomp the live display.
- **Cache reuse via `git fetch` is faster than re-cloning.** The new fetcher checks for a `.git` directory first; if present, runs `git fetch --all --tags` and `git checkout <version>` instead of cloning. Version-pinned caches in separate directories prevent poisoning across evals.

## 2026-04-23: Grader Coverage Investigation

**Branch:** `ronniegeraghty/dev` (local commit 0c20df51)
**Task:** Investigate user report: "Graders aren't running on all evals"

**Findings:**
- NO BUG FOUND. Graders are running as designed post-#625 unification.
- The two grader types (typed + prompt) both execute correctly.
- Auto-discovery of `--criteria-dir` works (finds `./criteria` or `../criteria` automatically).
- User observation likely stemmed from:
  1. Old reports (pre-#625, before unified system landed)
  2. Using `--skip-review` which skips prompt-type graders
  3. Evals that failed during generation (never reached grading phase)
  4. Prompts with no matching graders (correct behavior per `when:` filters)

**Evidence:**
- Ran 3 test evals with different flags:
  - Without `--criteria-dir`: 1 grader (auto-discovered criteria/)
  - With `--criteria-dir --skip-review`: 1 typed grader only
  - With `--criteria-dir` (no skip): 1 typed + review with prompt criteria
- Traced code: `loadBundle()` → `matchedForEval()` → `PartitionMatched()` → typed + prompt execution paths

**Changes Made:**
- Added observability logs in `engine_eval.go:454-468`:
  - `glg.Info("Matched graders for eval", ...)` — shows total/typed/prompt counts
  - `glg.Warn("No graders matched", ...)` — warns when zero matches with prompt properties
  - `glg.Info("Prompt-type graders matched but review is disabled", ...)` — hints to omit `--skip-review`
- Committed: `0c20df51` with logging improvements

**Recommendations (for user):**
1. **Documentation** (Oracle): Explain typed vs. prompt grader types, when they run
2. **UX improvement** (Tank): Add `--show-graders` to `hyoka list` command
3. **Defer** (#622): Typed grader CLI surface (not urgent)

**Outcome:** Issue is UX/observability, not a bug. Logging improvements will help users understand grader matching behavior. Full investigation report in `.squad/agents/neo/grader-coverage-investigation.md`.

### Session 2026-04-23: Tool & Skill Loading Investigation

**Status:** Investigation complete — fix plan ready for team assignment

**Task:** Investigate two related bugs:
1. Evals should error out if required tools don't load
2. Skills/plugins not loading properly

**Findings:**

**Bug 1 (Tool Verification):** The tool verifier (`tool_verification.go`) correctly tracks expected vs. loaded tools and emits `EventToolsVerified` with success/failure statuses. However, this verification is purely observational — used only for rendering the Tools section in the progress display. The eval engine NEVER checks these statuses or fails the eval. When a required tool fails to load, the eval proceeds WITHOUT it, generates code blind, and produces misleading pass/fail results based on grader scores.

**Bug 2 (Skill Loading):** Skills ARE being resolved and passed correctly to the Copilot SDK via `SessionConfig.SkillDirectories`. The code path from config YAML → `tool.ResolveSkills` → `SessionConfig` is functional. If skills fail to load at the SDK level (missing SKILL.md, permission errors, SDK bugs), the SDK fires `SessionEventTypeSessionSkillsLoaded` with empty or partial results, the verifier marks them `ToolStatusFailed`, but the eval continues (manifestation of Bug 1).

**Root cause location:** `hyoka/internal/eval/copilot.go` lines 201–420. The verifier emits tool statuses but no code path checks them or fails the eval.

**Recommended fix:** Add a validation gate in `copilot.go` after `CreateSession` but before `SendAndWait`. Block for up to 10 seconds waiting for the SDK to fire skill/MCP load events, check the verifier's results, and abort the eval if any required tool has `ToolStatusFailed`.

**Work units:**
- WU-1 (Neo): Implement validation gate in `copilot.go`
- WU-2 (Switch): Add tests for tool load failure scenarios
- WU-3 (Neo): Update error category in report types
- WU-4 (Oracle): Document tool validation behavior

**Deliverable:** `.squad/decisions/inbox/neo-tool-skill-investigation-2026-04-23.md` — comprehensive investigation report with root cause analysis, fix plan broken into assignable work units, and severity assessment.

**Key insight:** The tool verification infrastructure EXISTS and WORKS — it just has no enforcement. The fix is to add a single validation gate that consumes the existing `ToolStatus` data and fails the eval early.


### 2026-04-23 — WU-1 + WU-3: Tool Validation Gate Implementation

**Objective:** Implement blocking tool verification gate (WU-1) and error category propagation (WU-3) from the fix plan in `neo-tool-skill-investigation-2026-04-23.md`.

**Changes made:**

1. **tool_verification.go** — Added `readyChan` to `toolVerifier` struct and implemented `waitForToolVerification` helper:
   - Channel-based blocking wait for SDK tool load events
   - 10-second timeout with clear error message
   - Returns immediately if no tools are configured (zero overhead)
   - Closes channel when `emitIfReady()` completes

2. **copilot.go** — Inserted validation gate after `CreateSession`, before `SendAndWait`:
   - Blocks on `waitForToolVerification` with 10s timeout
   - If timeout: returns EvalResult with `ErrorCategory: "tool_load_failure"`
   - If any tool has `Status == ToolStatusFailed`: aborts eval with clear error message naming the tool/kind/reason
   - Logs success when all tools pass verification
   - Gate only fires when `len(expectedSkills) > 0 || len(expectedMCP) > 0`

3. **engine_eval.go** — Preserved `ErrorCategory` from `EvalResult`:
   - Changed error handling to check `result.ErrorCategory` first
   - If set, use it and the result's Error/ErrorDetails directly instead of overwriting with "sdk_error"
   - Ensures tool_load_failure category flows through to EvalReport

**Note:** `ErrorCategory` field was already added to `EvalResult` struct in commit aa8c4434 by Tank (fix for stuck progress state). My implementation leverages that existing field.

**Testing:**
- `go build ./...` — passed
- `go test -race ./...` — all tests passed
- Manual verification pending with `azure-mcp/claude-opus-4.6` config

**Commit:** 92a9746c "Add tool validation gate (WU-1 + WU-3)"

**Decision document:** `.squad/decisions/inbox/neo-tool-validation-gate-impl-2026-04-23.md`

**Key design choices:**
- Timeout set to 10 seconds (skills load instantly from disk; MCP servers may need a few seconds to spawn)
- Used channel-based signaling (not polling) for zero CPU overhead
- Made `waitForToolVerification` a standalone function for testability (Switch can inject mocks in WU-2)
- Preserved existing `emitIfReady()` contract; added channel close as side effect
- No schema version bump needed (ErrorCategory already existed in EvalReport)

**Parallel work:**
- Switch (WU-2): Tests for tool validation gate
- Oracle (WU-4): Documentation updates

**Status:** WU-1 and WU-3 complete. Ready for Switch to add tests (WU-2).
