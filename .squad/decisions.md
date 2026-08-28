# Active Decisions

## Issue Triage: Quick-Win Recommendations (2026-05-22)

**By:** Morpheus  
**Status:** RECOMMENDED — awaiting team assignment and execution

**Decision:** Post-dev-branch merge prep, identified 3 high-value, low-risk quick wins from 32 open issues.

**Top 3 Recommendations:**

1. **#644: Enforce evaluation_criteria removal** (XS, ~30 min) — Neo
   - Drop loader support for deprecated field in internal/prompt (3 files)
   - Dev-branch blocker carryover; zero risk, field already non-functional

2. **#595: Extract useRuns hook** (XS, ~20 min) — Trinity
   - Refactor shared fetch+cancel pattern from dashboard/prompts pages
   - Pure React change; TypeScript ensures correctness

3. **#72: Add early auth check** (S, ~45 min) — Neo
   - Call GetAuthStatus() after client.Start() in engine init
   - Fail fast with clear messaging; single-file change

**Sequencing:** Start #644 (fastest, highest dependency), then #72 (auth), #595 in parallel (site-only).

**Honorable Mentions:** #71, #633, #75 deferred for lower effort-to-value ratio.

**Next Steps:** Backlog verification sweep after Phase 7 ships (several older issues #14, #78, #86 may be stale).

**Full Details:** See `.squad/log/2026-05-22T21-44-13Z-issue-triage-quick-wins.md`

---

## Report/Site Schema Cutover: Dual-Emit Alias Removed (2026-05-14)

**By:** Trinity  
**Status:** MERGED (commit a41238ab)

**Decision:** Removed all dual-emit alias tests and expectations from codebase. `MigrateToV3()` already hard-rejects pre-v4 reports, and `GraderResult.Checks` emits `points` (not an alias). Keeping tests or code around a never-shipped alias path only creates false failures and schema drift.

**What Changed:**
- Removed test_* in internal/report that validated unreachable alias contract
- Aligned site/src/app/data/types.ts to match current Go JSON emit shape (panel_results: score/pass/issues/strengths/criteria)
- Fixed site/src/app/shared/components to read availableTools correctly

**Rationale:** Clear the schema regression blocker from dev-branch merge review (blockers #7, #8 of 9).

---

## Dev Branch Merge Review Blockers (2026-05-14)

**Decision Context:** Six-agent parallel review of `ronniegeraghty/dev` for merge readiness. All agents found merge blockers; one approved with notes.

**Status:** REJECT — branch requires fixes before merge

**Blockers Summary:**
1. **4 failing tests in `internal/report`** — dual-emit JSON schema gap
2. **1 failing test in `internal/pairwise`** — `TestComputeCheckDiffs`
3. **Accidental binary commit** — `smoke-test-output/hyoka-neo` (14–15 MB)
4. **Test output scratch file** — `test_output.txt` at repo root
5. **Incomplete breaking change** — `evaluation_criteria` parser still accepts deprecated field
6. **Inline grader type-filtering gap** — all grader types feed into review buckets
7. **Frontend schema regression** — site reads wrong field names for `ReviewPanelResult`
8. **Frontend camelCase gap** — site doesn't read `availableTools`
9. **Missing CHANGELOG entries** — 2 bug fixes not documented

**Next Steps:**
- Fix 5 failing test cases
- Remove accidental binary/test-output commits
- Enforce `evaluation_criteria` removal in parser
- Type-filter graders in review buckets
- Update site schema types to match Go `ReviewPanelResult`
- Add camelCase `availableTools` support to site
- Add CHANGELOG entries for skill leaf-expansion and tool_availability fixes

**Agent Verdicts:**
- Morpheus (Lead): REJECT — 4 fixable blockers, <30 min to fix
- Neo (Eval): REJECT — 2 framework issues
- Tank (CLI): REJECT — 2 artifact commits (no code issues)
- Trinity (Frontend): REJECT — 2 schema regressions
- Switch (Tester): REJECT — 5 failing tests
- Oracle (Docs): APPROVE WITH NOTES — missing CHANGELOG entries

**Merge Criteria:** All blockers resolved, test suite green, CHANGELOG complete.

**Reviewed:** 2026-05-14 22:00–22:41 UTC  
**Full Details:** See `.squad/log/2026-05-14T22-41-09Z-dev-branch-merge-review.md`

---

## Inline Graders Must Encode Prompt-Specific Contract (2026-05-06)

**Decision:** Inline `graders:` on prompt files should encode the **prompt's unique evaluation contract**, not duplicate generic criteria-file graders.

**Rationale:**
- Inline graders that duplicate checks already in `criteria/language/*.yaml` or `criteria/service/*.yaml` don't earn their keep
- They create maintenance drag and obscure what makes each prompt special
- Specialized behavior belongs inline; reusable checks belong in criteria files

**Resolution:**
- If an inline grader is identical to a criteria-file grader → move it to criteria files or specialize it
- If inline grader encodes prompt-specific behavior (e.g., "this prompt must not generate .rs files") → keep it inline and document the unique contract

**Example:**
- `prompts/test/hello-markdown-with-code.prompt.md` inline graders now check Rust code-block + no extra source files (prompt-specific)
- `prompts/test/hello-yaml.prompt.yaml` inline graders check exact bullet labels + no code fence + no extra files (prompt-specific)
- Both removed duplicates that were already in `criteria/language/test.yaml`

**Follow-up:** During docs reviews or prompt creation, enforce this pattern: inline graders are **specializations**, not duplicates.
