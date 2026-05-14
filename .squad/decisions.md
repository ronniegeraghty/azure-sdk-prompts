# Active Decisions

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
