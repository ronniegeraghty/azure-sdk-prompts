# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/main_test.go, hyoka/testdata/, hyoka/internal/ (packages to test)

## Core Context

Agent Switch initialized as Tester for hyoka. Guardrail defaults: max turns 25, max files 50, max output 1MB, max session actions 50. Safety boundaries prevent real Azure provisioning (--allow-cloud to opt out). Tests run via `go test ./...` from workspace root.

### Condensed History (Phase 0–4)

**Phase 0 (2026-04-04):** Fixed flaky resourcemonitor tests with event-driven pattern.

**Phase 1–3 (2026-04-04 → 2026-04-17):** Built comprehensive test suite: SPA tests (Vitest + RTL, 41 tests), guardrail starter-awareness review (#565, #567), full Phase 4 audit. Current baseline: 64.1% Go coverage, 94.3% site tests, zero flakes.

**Phase 4 Rollup (2026-04-17 → 2026-04-20):** CLI↔site equivalence tests (marshal+unmarshal round-trip comparison). WorkspaceDelta extended tests + grader integration. Phase 4 Wave 3 consolidated review (#581, #582, #583).

**Key Pattern:** "Shared core" claims validated by DeepEqual on wire-format-roundtripped structs.

### Phase 6 CLI Invocation Convention (2026-04-21)

**Note:** As of Phase 5, main.go was moved to repo root. All test setup and example code should use:
```bash
go run . <command>     # ✅ CORRECT
```

NOT:
```bash
go run ./hyoka ...     # ❌ STALE (Phase 5 regression)
```

Oracle audited phase-6 docs and found 47 stale references — fixed in commits b5c4782c–874bedf9. Test harnesses and CI scripts should follow this convention going forward.

## Recent Sessions

## 2026-04-20: Phase 5 Review — #364 Morpheus Mock Fix (APPROVED ✅)

### Context
- Switch double-rejected #364: (1) 20 tests failed (wrong mock paths), (2) Oracle renamed tests to `.TODO` instead of fixing
- Trinity + Oracle locked out per reviewer protocol
- Morpheus stepped in (eligible, not previously rejected on #364)

### Verification
- ✅ Test files restored (`prompts-page.test.tsx`, `prompt-detail-page.test.tsx` — NOT `.TODO`)
- ✅ All 72 tests pass (`npm test -- --run`)
- ✅ Mock paths fixed: `../app/api` → `../app/data/api`
- ✅ Mock data structures match real API types (18-field PromptInfo, nested RunSummary)
- ✅ Real component bugs fixed: missing `showEnvToolsOnly` state, tool filtering split (`byToolAll` / `byToolEnv`)
- ✅ Tests are meaningful (not over-mocked — catch API contract violations)
- ✅ Live verification confirmed features work end-to-end (R150/R151/R154 all pass)

### Verdict
**✅ APPROVE** — Phase 5 ready for rollup PR.

Morpheus correctly fixed mocks, restored test coverage, and uncovered + fixed real bugs. No test coverage regressions. Live UI verification aligns with unit test claims.

**Decision:** `.squad/decisions/inbox/switch-final-364-approve.md`  
**Commit:** `4ee54be9` (merged to phase-5)

### Session 2026-04-20 (Phase 5: TDD Testing & Reviews)

**Issues:** #364, #366, #367, #368, #369 (5 total reviews)

**Phase 5 Workflow:** Pre-written TDD tests for all issues. Review all merges to shared `phase-5` branch. Enforce reviewer-protocol strictly.

**TDD Test Suites Pre-Written:**
- #364: Dashboard, Prompts Page, Prompt Detail (3 test files, 72 total tests)
- #369: Schema Validation (1 test file)
- Others: sanity checks built into implementations

**Review Cycle Summary:**

| Issue | Review 1 | Review 2 | Review 3 | Status |
|-------|----------|----------|----------|--------|
| #364 | ❌ REJECT (mock paths) | ❌ REJECT (coverage hidden) | ✅ APPROVE (Morpheus fix) | Merged |
| #366 | ✅ APPROVE | — | — | Merged |
| #367 | ✅ APPROVE | — | — | Merged |
| #368 | ✅ APPROVE | — | — | Merged |
| #369 | ❌ REJECT (incomplete schema) | ✅ APPROVE (re-review) | — | Merged |

**Critical Reviews:**

1. **#364 First Rejection:** 20 tests failed due to incorrect mock paths (`../app/api` → `../app/data/api`)
   - Locked Trinity per reviewer-protocol

2. **#364 Second Rejection:** Oracle renamed test files to `.TODO` instead of fixing mocks
   - Coverage regression — tests not fixed, hidden
   - Locked Oracle per reviewer-protocol

3. **#364 Final Approval (After Morpheus Fix):**
   - All 72 tests pass
   - Mock paths corrected, mocks match real API types
   - Component bugs fixed (state variable, tool filtering)
   - Live verification confirmed UI works

4. **#369 Rejection:** Schema definitions incomplete (missing 3 fields on PromptInfo)
   - Locked Oracle after first rejection
   - Oracle re-reviewed with Trinity's help: ✅ APPROVE

**Phase 5 Outcome:** 5 issues reviewed, 3 clean approvals, 2 issues required re-review after fixes. Reviewer-protocol escalation chain on #364 successfully resolved by Morpheus.

**Key Learning:** Reviewer-protocol is the enforcement mechanism that separates serious review from rubber-stamping. Two rejections → lock → escalate is the right pattern.

### 2026-04-20 (Phase 5 Wrap-up — Morpheus Arch Review)

**Status:** Phase 5 PR #592 approved with followups for Phase 6.

**For Switch:** Three follow-up issues (#594, #595, #596) identified for Phase 6 scope:
- #594: Remove backup test files (.backup, .test suffix)
- #595: Unify dashboard/prompts fetch pattern
- #596: Refine `isTestValue()` heuristic (affects schema validation tests in #369)

**Next:** Phase 6 planning will prioritize these based on dependency graph and test coverage strategy. Morpheus's review is in `.squad/reviews/phase-5-arch-review-2026-04-20T200455Z.md`.

### 2026-04-20 (Phase 5 Fixups — PR #592 Test Update)

**Status:** ✅ COMPLETE

**Issue:** #366 added `architecture.md` to `internalDocs` exclusion map in serve.go, but TestAPIDocsEndpoint fixture still referenced it.

**Fix:** Updated `hyoka/internal/serve/serve_test.go` — changed fixture from architecture.md to configuration.md. Preserves the 2-doc multi-listing assertion.

**Commit:** `680ba625`

**Validation:** go test ./... and go vet both pass; PR #592 CI green.

**Pattern:** Stale fixture vs. internalDocs exclusion. When a feature adds items to exclusion maps (serve.go), search for test fixtures that assert on output now excluding those items.

**Cross-agent note:** Morpheus fixed the #596 (R151 collapsible) in parallel; both committed. PR #592 now ready for Ronnie → dev merge.



### 2026-04-21 (PR #604 Review — APPROVE WITH NITS ⚠️)

**Issue:** #600 — Run-level filter system on site (Trinity, Phase 6 round 2)
**PR:** #604 → `phase-6`
**Verdict:** ⚠️ APPROVE WITH NITS → merge

#### Verification

- `git worktree add ../hyoka-pr604 origin/ronniegeraghty/issue-600-site-filters`
- `cd site && npm install && npx vitest run` → **119/119 pass** (matches PR claim)
- Read `lib/run-filters.ts`, `multi-select-filter.tsx`, `runs-page.tsx`, both new test files

#### What's solidly tested

`run-filters.ts` (16 tests): status precedence all 3 branches, OR-within / AND-across, the subtle "any eval per dimension matches at run-level" semantic, URL round-trip + unknown-status rejection + missing-`results` tolerance. `runs-page.tsx` (5 new DOM tests): filter bar render, filter→reset cycle, no-matches empty state, URL hydration via `MemoryRouter initialEntries`.

#### Nits filed (non-blocking)

`MultiSelectFilter` has live wiring with zero coverage: outside-click close, Escape close, toggle-off (the component's inline `toggle()` does NOT delegate to the unit-tested `toggleValue`), "No options" branch, and there's no component-level URL **write** round-trip test. The primitive is intended for reuse on prompts/dashboard but has no dedicated test file.

#### Why not reject

The gaps are dropdown UX micro-behaviors, not data-correctness paths. Filter model where bugs would silently corrupt output is well-covered. Trinity put heavy logic in pure functions → test density landed in the right place.

#### Learning

The "approve with nits" lane is the right call when (a) acceptance criteria all have at least one pinning test, (b) data-correctness paths are well covered, and (c) the gaps are in UX wiring whose breakage would be visually obvious. Reserve hard reject for #587/#603-class regressions where silent backend routing breaks. The nit list is itself the artifact — a follow-up issue tagged `tests` can pick up the MultiSelectFilter dedicated test file before the primitive gets reused on prompts/dashboard pages.

### 2026-04-21 (PR #606 Review — APPROVE WITH NITS ⚠️)

**Issue:** #599 — `group` property in prompt frontmatter (Neo, R102)
**PR:** #606 commit `d6f6900d` → `phase-6`
**Verdict:** ⚠️ APPROVE WITH NITS — non-blocking

#### What was verified

- Full `-race` suite green (24 packages, 3m timeout).
- Coverage: validate 81.7%, prompt 92.6%, eval 54.5% — no regression.
- Live `go run . validate` against the existing prompt library: all 89 prompts pass unchanged (backward compat).
- Adversarial injection: prompt with `group: Bad_Group!` produces clean `invalid group "..."; must be kebab-case ...` error from CLI validate path — no parser panic.
- Regex coverage matches all categories Ronnie listed (uppercase, leading digit, leading/trailing/double hyphen, special chars, empty, > 64).

#### Nits filed (non-blocking)

1. No unit test for `engine_eval.go:78-80` propagation to `PromptMeta["group"]`. Three lines, but they are the entire bridge to Trinity's #600 site filter consumption. A two-row table test would lock the contract.
2. Regex test table doesn't explicitly include `a1` or `auth-2`. Transitively covered by `abc123-def` and `auth-flows-v2` — adding explicit rows is documentation, not behavior.
3. No `json.Marshal` round-trip asserting `group,omitempty` actually omits when empty (the json-tag claim in PR description is unverified by tests).

#### Learnings

1. **"Three trivial lines = the entire user-visible contract" pattern.** The `engine_eval.go` propagation is 3 lines but it's the only path by which Trinity's #600 site filter and the whole point of issue #599 reaches the report JSON. Trivial code can be load-bearing — coverage gaps on bridge code matter more than coverage gaps on logic. Don't reject for it alone, but always nit it.
2. **Regex test tables benefit from explicit "boundary" rows even when transitively covered.** `abc123-def` proves the regex handles letter+digit. `a1` proves the *minimum-viable* letter+digit case. The latter is more debuggable when something breaks. Cheap to add, high signal in failure mode.
3. **Adversarial integration test pattern works for new validation rules:** drop a malformed prompt into the live tree, run the CLI, grep the error message. Catches "validation runs but error doesn't surface to user" — a class of bug unit tests can miss.

## 2026-04-21 — PR #605 Review (Neo's #597 tool versioning)

**Verdict:** ⚠️ APPROVE WITH NITS — comment posted ([#issuecomment-4285214970](https://github.com/ronniegeraghty/hyoka/pull/605#issuecomment-4285214970))

**Suite:** `go test -race ./hyoka/... -timeout 3m` → all 24 packages pass, EXIT 0.

**Headline guard `TestCustomFetcherInvokedAtRuntime` is real:** registers a mock against the production `DefaultRegistry`, calls `ResolveSkills` (the actual prod path used by `cmd/run.go` → `InstallSkillsAndPlugins`), asserts call count == 1 + version `v1.2.3` propagated + repo intact + returned dir. Clears the #587/#603 bar.

**Solid coverage:** custom-fetcher dispatch (happy + error), default-last ordering (`TestRegistry_DefaultStaysLast`), full per-entry-vs-map precedence matrix, YAML round-trip, backward compat (`NoMapNoOp`), registry hygiene (duplicate/nil/empty rejection).

**Nits filed (non-blocking, fast-follow):**
1. `TestValidateFetchers` only tests success — the `LookupFetcher(e) == nil` failure branch (the *whole point* of pre-flight) has zero coverage. Easy fix: unregister `npx` in test, expect wrapped error.
2. `TestNpxFetcher_VersionInPath` is misnamed — only asserts `CanFetch`/`Name`, never checks the `<version>` cache-segment path or the `repo@version` arg munging. The user-visible "toggling pins doesn't poison cache" behavior is untested. Suggest factoring path-build out of `Fetch` for unit testability.
3. `LoadDir` conflict-detection branch (`"conflicting tool_version_override for %q"`) untested.
4. Two-custom-fetchers-same-entry ordering not asserted.

**Lesson reinforced:** the #587 trap isn't only "is the runtime path tested?" — it's also "is the *failure* mode of pre-flight validation tested?" Pre-flight checks that only test success ship green when they no-op silently. Watch for this pattern in future PRs that add `Validate*` functions.

## 2026-04-21 — Phase 6 Rollup Review (PR #607)

**Status:** ✅ APPROVE WITH NOTES
**PR:** #607 (phase-6 → ronniegeraghty/dev, epic #312)
**Scope:** Integration-level review — 6 sub-PRs (#601–#606) already ✅✅ per-PR.

**Test results on integrated phase-6 branch:**
- `go test -race ./hyoka/... -timeout 3m -count=1` → **24/24 packages green**
- `cd site && npm test` → **14 files / 119 vitest cases green**
- `go build ./hyoka/...` → clean
- `go run . validate` → 12 configs, 25 graders, all valid
- `go run . run ... --dry-run` (key-vault-dp-python-crud × baseline/claude-opus-4.6) → resolves, no panic

**Feature-interaction probes:**
1. **#606 ↔ #604:** `Prompt.Group` → `PromptMeta["group"]` wiring confirmed at `engine_eval.go:79`. Filter system can consume via `EvalReport.PromptMeta`.
2. **#603 ↔ #605:** #587-trap guard intact — `slog.Warn("review-mode=isolated requested but no graders or groups are marked isolate...")` at `engine.go:282`, runtime-asserted in `engine_reviewbuckets_test.go:150`. Rebase/merge did NOT break the trap.
3. **#602 ↔ #606:** `prompt_directory` config and `group` frontmatter field are orthogonal (config-layer vs frontmatter-layer); no interaction issue. Parser handles both; validate succeeds on full prompt library.

**Non-blocking follow-ups (from per-PR reviews, still open):**
- MultiSelectFilter component has no direct unit test (logic lib `run-filters.ts` is tested at 16 cases; component rendering is covered only via integration with runs page).
- No explicit unit test asserts `PromptMeta["group"]` round-trip through EvalReport JSON (covered transitively by engine tests but not pinned).
- No cross-feature test exercises "custom fetcher + isolated review mode" — orthogonal code paths (tool-fetch happens pre-generation, review-mode post-generation), so low regression risk, but a combined integration test would be worth filing as a follow-up.

**Live eval skipped:** Dry-run covered startup/config resolution. Full live eval requires Copilot credentials and 5-10 min runtime; not necessary to gate the rollup given all unit+integration tests are green and dry-run resolves cleanly.

**Verdict:** APPROVE WITH NOTES — rollup is test-green, integration holds, no regressions. Follow-up tests listed above can land on `dev` post-merge.

## Session 2026-04-21 (Phase 6 Round-1 Test Review)

**Mission:** Test review of Phase 6 Round-1 batch (PRs #601, #602, #603)

**Verdicts:**
- #601 (Compare page redesign): ✅ APPROVE — 31 new tests, 99/99 green, edge cases covered
- #602 (Configurable prompt dir): ✅ APPROVE — 11 new tests green, backwards-compat locked
- #603 (Review session splitting): ❌ REQUEST CHANGES — wiring-layer untested on 4 surfaces; reassigned to Tank

**Pattern captured:** When re-implementing dead-flagged feature (#580 ↔ #587), wiring-layer integration tests (Engine/cmd plumbing) required as gating, not optional. Tank fixed; re-reviewed ✅ APPROVE.

**Status:** Phase 6 Round-1 test batch complete. All PRs ready to merge pending Coordinator's embedded-asset refresh (completed, a1a3c95d).

## 2026-04-21 — PR #609 Review (MultiSelectFilter tests, Trinity, phase-6)

**Verdict:** ⚠️ APPROVE WITH NOTES

- Tests: 122/122 green; 3 new tests in `multi-select-filter.test.tsx` (outside-click, Escape, empty-options).
- Quality good: correct event names (`mousedown` matches listener), ARIA-role queries, no sleep anti-patterns, no portal leaks, double-assertion on empty case.
- Non-blocking gaps flagged for follow-up: toggle/onChange behavior untested (highest-value miss for a "MultiSelect" component), summary-text branches untested, `aria-expanded`/`aria-selected` not asserted, no click-inside-stays-open counterpart. Keyboard listbox nav is a product gap to route to Trinity, not a test gap.
- Recommendation: land for targeted regression coverage; file follow-up for state/aria coverage.

## PR #610 test-quality review — ✅ APPROVE (2026-04-21)

**Branch:** ronniegeraghty/issue-608-606-group-tests (Tank — group property follow-up tests for #606/#608)
**Tests:** `go test -race ./hyoka/... -timeout 3m` → all 24 packages green.

**Files reviewed:**
- `hyoka/internal/eval/engine_group_wiring_test.go` (new, 93 LOC, 2 subtests)
- `hyoka/internal/prompt/group_json_test.go` (new, 86 LOC, 4 subtests)
- `hyoka/internal/validate/group_test.go` (+38 boundary rows)

**Verdict — wiring test passes the #587-trap bar:** Uses real `engine.Run` + `StubRunner` flowing through actual wiring point `engine_eval.go:78-80`. Asserts on observable runtime payload `r.PromptMeta["group"]`. Failure message points at the file:line so a future regression gets a pointed error. Negative case (empty Group → key absent) covers the conditional branch.

**Verdict — regex boundary rows comprehensive:** 35+ rows covering lengths (63/64/65 plain + hyphenated), whitespace classes, hyphen edge cases (only/double/triple/leading/trailing/many-single positive), digits, case, special chars (`.` `_` `+` `@` `#` `!` space), unicode (ñoño, café, emoji), embedded null byte.

**Verdict — omitempty round-trip covers both paths:** empty omitted; populated round-trips; absent unmarshals empty; explicit-empty-on-remarshal still omitted.

**Posted as comment** (gh blocked --approve since I am the PR author due to coordinator setup; verdict body still leads with ✅ per protocol).

## 2026 PR #612 review — Neo fetcher polish (#605/#608)

**Verdict:** ✅ APPROVE (posted as comment — gh refuses self-approve on bot-owned PR)

- Verified `go build ./hyoka/...` + `go test -race ./hyoka/... -timeout 3m` both green.
- `TestValidateFetchers` no-fetcher branch genuinely exercises the empty-registry path (unregisters `npxFetcher`, asserts error + substring match on repo name, restores via Cleanup).
- Rename `TestNpxFetcher_VersionInPath` → `TestNpxFetcher_CanFetchAndName` is a correctness fix — old name didn't match body.
- `TestFetchRemote_ContextPropagates` uses `context.WithValue` + unique key to prove the *same* ctx reaches `Fetch` (stronger than #587 non-nil pattern). Suggested non-blocking follow-up: cancellation variant.
- `[][]Entry → []Entry` signature change: grep confirms zero stragglers; `cmd/run.go` correctly flattens with `append(..., Tools...)`.
- All ctx threading sites (`ResolveSkills`, `FetchRemote`, `buildSessionConfig`, `Engine.dryRun`) updated, all test files updated.

## 2026-04-21 — PR #611 review (Morpheus: site-embed Makefile + CI freshness gate)

**Verdict:** ✅ APPROVE (posted as comment — `gh pr review --approve` rejected because branch was authored under same `ronniegeraghty` identity).

**Tested in worktree** `/home/rgeraghty/projects/hyoka-review-611-switch`:
- `make site-install` → `make site-embed` × 2 → `git status` clean both runs. Idempotent. ✅
- `make verify-embed` clean → exit 0.
- Drift simulation: appended a real code line to `site/src/main.tsx` → `verify-embed` exit 1 with clear actionable error. ✅
- Comment-only edit to `App.tsx` correctly produced byte-identical bundle (Vite strips comments) → no false-fail. ✅
- `.github/workflows/ci.yml` untouched. ✅

**Non-blocking findings logged on the PR:**
1. `git diff --quiet -- EMBED_DIR` misses untracked file additions — if Vite ever starts emitting a new root-level artifact, verify-embed could false-pass. Suggested `git status --porcelain -- EMBED_DIR` instead.
2. `rm -rf $(EMBED_DIR)/assets` only cleans `assets/` — stale root-level files (e.g. removed `manifest.json`) would persist undetected.
3. No `concurrency:` group on the workflow.

All three are theoretical given current Vite output (only `index.html` + `assets/`). Worth a follow-up issue, not a blocker.

## 2026-04-21 — PR #613 review (MultiSelectFilter follow-up tests)

**Verdict:** ✅ APPROVE (posted as PR review comment — self-approve blocked)

Trinity's follow-up to my #609 review. All four deferred gaps closed:
- Toggle/onChange (3 tests including controlled-state Wrapper with `toHaveBeenNthCalledWith` for sequential transitions)
- Summary branches (5 tests, all paths)
- ARIA (`aria-expanded` dynamic toggle + `aria-selected` per option)
- Inside-click (counterpart to outside-click, listbox stays mounted)

**Quality high points:**
- Real `userEvent.setup()` + `await user.click()` for actions under test
- Exact-payload assertions (`toHaveBeenCalledWith(["a","b"])`)
- `toHaveBeenNthCalledWith(1..4, …)` validates each transition, not just final state
- `fireEvent.mouseDown` retained for outside-click — correct, matches component's `mousedown` listener

**Behavioral observation validated:** Single-select renders `selected[0]` (value, not label) — confirmed in component source. Test locks current behavior with explicit `// Note:` comment. Right call as regression guard; likely worth a separate product issue (UX wants label).

**Guardrails:** component untouched (single-file diff +232 LOC), 133/133 pass clean, MultiSelectFilter 3→14 as advertised.


## 2025-01 — PR #614 review (site-embed freshness CI hardening; Morpheus, follow-up to #611)

Reviewed test/CI correctness for the three nits I'd raised on #611. All three resolved cleanly:

1. **`git status --porcelain`** replaces `git diff --quiet`. Verified by direct test: `touch hyoka/internal/serve/site/zzz-test-stray.txt` → porcelain emits `??` line → gate trips. Note: `make verify-embed` re-runs `site-embed` which wipes `EMBED_DIR/*`, so manually-injected strays get pruned before the check — but the realistic CI scenario (source change emits new asset filename, dev forgets to commit) does trip the gate after rebuild.
2. **`rm -rf $(EMBED_DIR)/*`** replaces `rm -rf $(EMBED_DIR)/assets`. PR comment grounds against vite output shape.
3. **`concurrency:`** added with standard `${{ github.workflow }}-${{ github.ref }}` + `cancel-in-progress: true`.

Idempotence verified, `go build ./hyoka/...` clean, `go test -race ./hyoka/... -timeout 3m` green (20 packages). `phase-6` removed from push triggers with clear comment. Duplication comment is honest (names the cost ~1-2 min/run, explains the trade-off).

**Non-blocking nit filed:** `EMBED_DIR := hyoka/internal/serve/site` is safe at parse time, but `make EMBED_DIR= site-embed` would expand to `rm -rf /*`. Suggested an `ifeq ($(strip $(EMBED_DIR)),)... $(error ...)` guard. Pre-existing risk, defer to follow-up.

**Verdict:** ✅ APPROVE. Posted as `--comment` (gh refused `--approve` because PR was opened under same gh account — author/reviewer isolation hit again, same as Neo on #614). Author of record is Morpheus.

**Learning:** When reviewing your own past nits, always test the ACTUAL gate semantics, not just the surface change — porcelain works, but `verify-embed`'s site-embed dependency means in-tree manual stray files get pruned before the check. The realistic CI failure mode (untracked output of a rebuild) still trips correctly.
