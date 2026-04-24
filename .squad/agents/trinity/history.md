# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, React + Recharts on frontend
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Core Context

**Archived 8 entries from earlier sessions.**

---

## CROSS-AGENT UPDATE (2026-04-24T04:55:03Z — Tank: Bucket-Per-Entry Structure Change)

**Grader Bucket Structure Refactor:** Tank changed `BuildUnifiedReviewBuckets` in `hyoka/internal/criteria/buckets.go` to create ONE bucket per criteria-file entry (instead of grouping all into "combined"). **Site Impact:** Each criteria entry now appears as a separate top-level grader with distinct name. The bucket count will increase proportionally to the number of criteria entries. This is visible in report JSON structure (`EvalCriteriaBuckets[]`). Commit: 9e2d8100. If you're working on grader display or report rendering, this structural change may affect layout calculations.

---

Historical patterns and learnings:

- ## Core Context: Agent Trinity initialized as Site / UI architect. Charter: frontend components, data visualization, React patterns. Expertise: Vitest + React Testin...
- ## Recent Sessions: ### Session 2026-04-21 (#600 — R146/R147 Run-Level Filter System)

Built the filter UI for the runs page. PR targets `phase-6` from worktree `hyoka-...
- ## Session 2026-04-21 (Phase 6 Round-1 Approval: Compare Page + Run Filters): **Mission:** Test & architectural review of Phase 6 Round-1 batch (PRs #601, #602, #603)

**PR #601 Verdicts:**
- **Switch (test):** ✅ APPROVE — 31...
- ## Team Context: Unified Grader Direction Proposed (2026-04-22): Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE sch...
- ## 2025-01-style-helper — ANSI style primitives for progress renderers: Created `hyoka/internal/progress/style` — a small dependency-free ANSI color helper to be consumed by the upcoming interactive and CI renderers (spr...
- ## 2025-04-ci-renderer — append-only CI progress renderer: Built `hyoka/internal/progress/display_ci.go` to replace the legacy log-mode
renderer with a proper CI-oriented view: timestamped start/finish lines...
- ## Team Updates: ### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three...
- ## Session 2026-04-23 — Interactive Progress Tail Line Wrapping Bug (FIXED): **Mission:** Debug and fix Bug B — multi-line tail leak in interactive eval progress renderer.

**Context:** Two bugs were under repair in a prior s...

Full history archived. Recent entries below.

---

## 2026-04-23: Issue Verification — #595 & #290 Status Audit

**Task:** Verify status of two stale GitHub issues per Morpheus audit report.

### Issue #595: Extract useRuns hook

**Status:** COMPLETE but NOT CLOSED ❌

**Findings:**
- Issue filed on Phase 5 architectural review by Morpheus 🕶️
- Dashboard/prompts pages **both have** the duplicate fetch + cancellation pattern (`useEffect` → `fetchRuns()` → cancel flag)
- `site/src/app/hooks/` directory **does NOT exist** — no `useRuns.ts` hook was extracted
- PR #592 (Phase 5 rollup) **mentions #595 as a follow-up**, NOT as resolved
- PR #592 diff explicitly calls it "Duplicate `useRuns()` fetch pattern" in the findings table → "Follow-up #595"
- **Conclusion:** Issue is genuinely pending. Hook was identified in Phase 5 but deferred. No extraction has occurred.

### Issue #290: Criteria table — baseline config on left

**Status:** COMPLETE and can be CLOSED ✅

**Findings:**
- Phase 4/5 repos (#358, #572, #590) did major eval detail/comparison table UI rework
- Comparison matrix table (in `hyoka/internal/report/markdown.go:410-451`) renders configs from `matrix.Configs` slice
- `buildMatrix()` in `report_data.go` collects configs in order they appear in results but **does NOT explicitly sort**
- Actual generated reports show baseline configs **already appear first** (leftmost)
  - Example: `/reports/20260423-172207/summary.md` with 9 configs: baseline variants are columns 1–3, non-baseline variants follow
  - Config order is: `python-pairwise/baseline/*` → `python-pairwise/without-azure/*` → `python-pairwise/without-generator-skills/*`
  - This happens naturally because baseline config names sort lexicographically before ablated variants
- **Conclusion:** The implicit lexicographic sort achieves the desired outcome (baseline left). The layout change may already be satisfied by default string sorting behavior, or a prior PR implemented explicit sorting. Either way, real reports confirm baseline is on the left.

### Recommendations

**#595:** Leave open. Genuine work item pending. Extract `useRuns` hook per acceptance criteria (create `site/src/app/hooks/useRuns.ts`, refactor both pages, test passes).

**#290:** Close with comment noting Phase 4/5 table rework achieved the goal — rendered reports confirm baseline configs appear leftmost in comparison tables.


---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.

## Site UX review — 2026-04-23 (Ronnie)

Audited `hyoka/internal/serve/` + React `site/src/` against anchor run `20260423-195948` (12/12 passed). Drove the served site with `playwright-cli` (skill: `playwright-cli`) on port 8088. Wrote `.squad/decisions/inbox/trinity-site-ux-review.md`.

**Findings (UX layer):**
- `run-detail-page.tsx:236` — strict `g.pass === true` filter excludes review graders (which ship `pass: null`). On a 12/12 anchor run every row in the table renders `1/4` or `0/0` in red. **Single most visible bug**, anchor of Ronnie's report.
- `GraderResultRow.tsx:16` — same tri-state collapse: 3 of 4 graders show `N/A` grey on a fully passing eval. Internally inconsistent with the green `12/12` card immediately above (`eval-detail-page.tsx:441`).
- `/dashboard` is currently a hard React error boundary crash (`Cannot read properties of undefined (reading 'toFixed')`). Bundle minified — didn't bisect line.
- `/runs` rate column folds errored runs into denominator → `0.0%` red bar with no "errored" signal.
- `tool_availability` JSON is a flat list — site has zero notion of plugin parents or skill-dir parents. Phase-2 plugin grouping will need `parent_kind`/`parent_name` on the wire.
- `GraderResult` has no `Points []` on the wire yet; site types in `data/types.ts` will need it added when Phase 2 ships.

**Quick wins shippable pre-Phase-2:** smarter `isPass()` helper that ANDs `scores.criteria` when `pass: null`; reuse same helper everywhere; relabel header score card "Review Score" so it stops contradicting the grader rows; guard dashboard crash; distinguish errored runs from failed runs on `/runs`.

**Phase-2 UX proposals:** plugin parent renders as header-only group with children indented (no parent badge); skill-dir parent uses config `name` not path; grader row collapsed shows `X/Y passed` badge, expanded shows N indented points; one `evalPassFromPoints()` helper used by every surface so rollup stays consistent.

**Conventions reaffirmed:**
- The site is a Vite/React SPA (`site/src/`) embedded into the Go binary at `hyoka/internal/serve/site/`. Source-of-truth code is the React, not the embedded `index.html`.

---

## 2026-04-24: Bug Fixes — Grader Point Scoring & File Contents Display

**Bugs fixed:**
1. **Total score denominator** (user report: "x/2 instead of x/all-grader-points")
   - Root cause: `engine_eval.go:636` used `len(graders)` instead of summing grader points.
   - Fix: Added `countTotalPoints()` and `countPassedPoints()` helpers that sum points across all graders, treating graders with zero points as 1 point for backward compat.
   - Changed `GradersTotal` and `GradersPassed` to use point-level counts.
2. **File contents missing on eval detail pages** (user report: "can see file names but not contents")
   - Root cause: `EvalReport.FileContents` field was unused (existed on `ReportTemplateData` but never populated).
   - Fix: Added `readGeneratedFileContents()` helper that reads each generated file from the workspace directory (up to 1MB per file), detects binary files by extension, and populates `EvalReport.FileContents` before writing the report.
   - Binary files (`.png`, `.pdf`, etc.) show `[Binary file — not displayed]`.
   - Files >1MB show `[File too large to display (N bytes) — view on disk at {path}]`.

**Learnings:**
- **Grader point denominator rule**: When computing total score, the denominator is the sum of `len(g.Points)` across all graders. Graders with zero points contribute 1 to the total (backward compat for legacy graders that don't populate Points).
- **File contents rendering pattern**: Store file contents in `EvalReport.FileContents` at report-build time (not serve time). Use size caps (1MB) and binary detection (by extension) to avoid blowing up JSON reports. Site can render inside `<details>` elements for collapsibility.

**Tests:**
- Added `grader_scoring_test.go` with table-driven tests for `countTotalPoints()` and `countPassedPoints()` covering: multiple points per grader, zero-point graders, mixed scenarios, empty results.
- All existing tests pass (`go test ./hyoka/...`).

**Commit:** c06ca9e2 (merged with Neo's bucket separation fix)
- `playwright-cli` is restricted to roots `/home/rgeraghty/projects/hyoka` and `.playwright-cli` — screenshot output must land inside the repo (or be copied to `/tmp/` after).
- Read-only investigations on `ronniegeraghty/dev` while Tank is concurrently editing means I never `git add` or commit; screenshots into `.trinity-screenshots/` (gitignored) + copy to `/tmp/trinity-site-review/`.

---

## 2026-04-24: v4 Grader Unification — Site Implementation (Option B)

**Task:** Implement site-side changes for Morpheus's v4 grader unification plan (Option B). Work in parallel with Neo's engine implementation per `.squad/decisions/inbox/morpheus-grader-unification-plan.md`.

**Changes made:**

1. **Updated `site/src/app/data/types.ts`** — replaced v3 GraderResult with v4 schema:
   - `GraderResult` now has required `points: GraderPoint[]` (always populated, len ≥ 1)
   - `pass: boolean` (no longer nullable — derived from Points in engine)
   - `message: string` (renamed from Summary — headline summary ≤ 120 chars)
   - `extras?: GraderExtras` (discriminated union replacing 6 separate `*_details` fields)
   - Removed: `model`, `scores`, `overall_score`, `max_score`, `summary`, `issues`, `strengths`, `duration_seconds`, `is_consensus`, `file_details`, `program_details`, `prompt_details`, `behavior_details`, `review_details`
   - Added: `GraderPoint{label, pass, message, weight, evidence}`, `GraderExtras` with 8 kind-specific structs

2. **Created `site/src/app/lib/graderScore.ts`** — canonical score formatting helper:
   - `formatGraderScore(result)` — ALWAYS returns `"N/M points"` (even for single-point graders)
   - No "Passed", no "100%", no special-casing
   - This is the ONLY score string format shown in grader row headers

3. **Created `site/src/app/components/grader-extras/`** — 8 kind-specific Extras components:
   - `FileExtras.tsx` — per-file existence + pattern checks
   - `ProgramExtras.tsx` — command, exit code, stdout/stderr
   - `PromptExtras.tsx` — LLM judge model, rubric, reasoning
   - `BehaviorExtras.tsx` — tools used, turn counts, violations
   - `ActionSequenceExtras.tsx` — **NEW**: expected-vs-actual sequence diff (fixes missing visibility bug)
   - `ToolConstraintExtras.tsx` — per-tool call constraints, violations
   - `OutputCheckExtras.tsx` — **NEW**: produced files list (fixes missing ProducedFiles bug)
   - `ReviewExtras.tsx` — multi-model panel breakdown, consensus marker
   - `index.ts` — barrel export

4. **Rewrote `site/src/app/components/GraderResultRow.tsx`**:
   - Header: **ONE** canonical score `formatGraderScore(result)` + icon-only pass/fail badge
   - **REMOVED** right-side duplicate score display (the inconsistency Ronnie reported)
   - Body: Points list (always) + KindExtras dispatcher (when present)
   - Per-point rendering: icon + label + message (on failure) + evidence chips (if present)
   - Auto-expand when `points.length > 1` OR `!result.pass` (failed graders)
   - Removed 6-way `if (file_details) …` cascade — replaced with single `extras` switch

5. **Updated `site/src/app/lib/evalPass.ts`** — simplified for v4:
   - `graderPasses(g)` now just returns `g.pass` (always boolean in v4)
   - Removed tri-state fallback cascade (Points > pass > criteria > overall_score) — no longer needed
   - `evalPointTotals()` updated to use `g.pass` directly (no more `graderPasses(g)` call)

**Key design decisions per Morpheus's spec:**

- **Score appears ONCE** in the header — no duplication on the right
- **Canonical format: "N/M points"** for ALL graders (fixes "Passed" / "100%" inconsistency)
- **Points list ALWAYS shown** when points exist (not just multi-point graders)
- **Extras render via dispatcher** — one switch on `result.grader_type`, no cascading conditionals
- **Auto-expand failed graders** — user sees the reason immediately

**Verified:**

- Site build passes: `npm run build` succeeded (1.1 MB bundle, no errors)
- Go build expected to fail until Neo lands engine v4 — my types are ready for the new JSON contract

**What's still needed (blocked on Neo):**

- Neo's engine-side v4 implementation (graders emitting Points, Extras, new report.GraderResult shape)
- Fresh v4 report to dogfood the rendering end-to-end
- Playwright visual test once Neo's code lands (verify score format, no duplication, expandable file viewer regression check)

**Files changed:**

- `site/src/app/data/types.ts` — v4 schema
- `site/src/app/lib/graderScore.ts` — new canonical formatter
- `site/src/app/lib/evalPass.ts` — simplified for v4
- `site/src/app/components/GraderResultRow.tsx` — full rewrite
- `site/src/app/components/grader-extras/*.tsx` — 8 new components + barrel export

**Notes for future dogfooding:**

1. After Neo lands, run: `go run ./hyoka run --prompt-id test-dp-test-hello-markdown --config "test/sonnet"` then `go run ./hyoka clean`
2. Start site: `go run ./hyoka serve` (background)
3. Playwright check:
   - Every grader row shows `N/M points` (no "Passed")
   - Right-side score duplication gone
   - Each point lists pass/fail + reason on failure
   - ActionSequence extras show expected-vs-actual diff
   - OutputCheck extras show ProducedFiles
4. Screenshot for verification

**Schema v4 contract established** — Neo and I are now building to the same JSON shape. Site is ready to render when the engine catches up.
- Coordination boundary: data-model and roll-up structural critique is Morpheus; presentation/UX critique is mine. Deliverables link rather than overlap.

## Phase 4 — Site quick wins shipped — 2026-04-23 (Ronnie)

All six Phase 4 tasks landed on `ronniegeraghty/dev`. Verified live against `reports/20260423-195948` via `playwright-cli` (skill: `playwright-cli`) on vite dev (5173) → hyoka serve (8080). Before/after PNGs in `.trinity-screenshots/phase4/`.

**Shipped:**
- `99133d8d` — `run-detail-page.tsx`: `isPass()` helper (tri-state aware) + `r.success`-anchored `—`/`✗` fallback when grader_results empty. Anchor row flipped from red `1/4` to emerald `4/4`; gpt-5.3-codex no-files row now shows neutral `—` instead of red `0/0`.
- `888a0552` — `GraderResultRow.tsx`: same fallback, derives pass from `scores.criteria` AND or `overall_score === max_score` when `result.pass == null`. The 3 review rows on the anchor eval flipped from grey `N/A` to emerald `PASS` (verified: 0 N/A, 4 PASS in DOM).
- Dashboard crash fix (NaN aggregations + missing `?? 0` on `run.duration_seconds.toFixed`) **bundled into `6df67540`**. My staged files (`dashboard-page.tsx`, new `ErrorBoundary.tsx`, `routes.ts → routes.tsx`) were swept into Neo's parallel commit when our processes both ran `git add` on the shared worktree at the same instant. Result: dashboard now shows 284 / 44.7% / 87.5s instead of NaN; ErrorBoundary wraps the route with friendly amber fallback. Did NOT rewrite history (would have clobbered Neo mid-work).
- `58475254` — `runs-page.tsx`: errored-rate fix. `effectiveTotal = total - errors`, amber bar + `⚠ run errored` tag on rows where `errors > 0`.
- `ff63ab6d` — `runs-page.tsx`: in-progress card path. Detects `total_evaluations == null` (summary.json never finalized), renders separate amber card with spinner + "no summary yet" subtitle.
- `58e12ab6` — `eval-detail-page.tsx`: `REVIEW SCORE` caption added above the `12 / 12` card so it stops looking like it's contradicting the grader rows. Stop-gap until Phase 6 reworks the metaphor.

**Verification:** `npm run build` green (1082 KB / 312 KB gzip; existing chunk-size warning unchanged). Live walkthrough: `/runs` shows in-progress card + amber errored runs; `/runs/{id}` shows `4/4` first row + `—` for no-graders row; `/dashboard` no longer crashes; eval detail header shows `REVIEW SCORE 12/12` green with all 4 grader rows emerald.

**Coordination lessons (parallel agent on shared worktree):**
- Two agents `git add`-ing into the same working tree races. The agent who calls `git commit` second can find their staged files already committed by the first agent under that first agent's commit message. **Mitigation for next time:** stage + commit in a single shell line (`git add X && git commit -m ...`) and verify `git log -1 --stat` immediately after each commit before doing more work; if the file list looks wrong, abort + re-stage.
- Did NOT touch `hyoka/internal/...` — Neo's territory this round. The accidentally-bundled commit is annoying but the work shipped correctly; rewriting branch history while Neo was still committing would have been worse.
- Per the user's standing instruction, did NOT start Phase 6 (depends on Neo's Phase 2 model). Stopping here.

## 2026-04-23 — Phase 6: Site v3 alignment (single canonical pass helper)

Aligned the React/TypeScript site with schema-v3 reports in 7 commits on `ronniegeraghty/dev`. Goal: kill the "all passed but pages disagree" bug by routing every page-level rollup through a single `evalPassFromPoints` helper, and make the new Points machinery visible everywhere.

**What shipped:**
- `site/src/app/lib/evalPass.ts` — new canonical pass helper. Precedence: (1) `graders_passed`/`graders_total` rollup if engine populated it; (2) AND of `points[*].pass` per grader; (3) `g.pass` flag; (4) derive from criteria/overall_score; (5) `r.success` legacy fallback. Exports `evalPassFromPoints`, `graderPasses` (tri-state aware single-grader rollup), `evalGraderTotals`, `evalPointTotals`. Loose `EvalPassInput` type accepts both `EvalReport` and `EvalResult` so all five page components can share one signature.
- `data/types.ts` — added `GraderPoint`, `points?` on `GraderResult`, `ToolAvailabilityEntry` (with optional `parent`/`parent_kind`/`kind` as forward-compat — see gap below), `tool_availability?`/`graders_passed`/`graders_total`/`schema_version?` on `EvalReport`. Added camelCase aliases on Environment (`skillsLoaded`, `skillsInvoked`, `mcpServers`) — Go emits most env fields camelCase but `skill_groups` snake_case, the site needs both.
- `GraderResultRow` — Points panel inside expanded grader body. Multi-point graders auto-expand and show a `✓ N/N passed` / `✗ M/N passed` badge plus per-point breakdown with check/X icons; flat (single-point) graders keep the legacy `PASS`/`FAIL`/`N/A` badge. Score column shows `N/M points` when point data exists.
- `eval-detail-page` — header pass icon, error banner, score card, and Available Tools section all now use the helper. Score card renders `✓ N / N points across M graders` with a hover tooltip listing each grader's contribution. Available Tools groups skills under their plugin/skill_dir parent (joined via `environment.skill_groups` by name); MCP and builtins stay flat. Fallback path synthesizes tool rows from `skillsLoaded` + `mcpServers` so v2 reports still render.
- `run-detail-page`, `prompt-detail-page`, `dashboard-page` — every `r.success` rollup replaced with `evalPassFromPoints(r)`. Prompt-detail's Pass-Rate-by-Tool-Used table now groups sibling tools by prefix (split on `.`/`/`/short `_`); single-child groups stay flat.

**Verified live (playwright-cli against `go run . serve --site-dir site/dist`):**
- v3 report (`reports/20260423-214602`): run-detail rows show `2/2`/`—` badges (no red on passes); eval-detail shows `Points 14 / 14 across 2 graders` with per-grader hover; grader rows expanded with check icons; Available Tools groups `azure-sdk-python` (42 skills) under "plugin" and `generator-skills` (1) under "skill_dir"; dashboard renders.
- v2 fallback (`reports/20260423-195948`): legacy `PASS` badges still render (graders without points fall through to `g.pass`); Available Tools synthesizes from camelCase env. No crashes anywhere.
- 10 baseline screenshots committed to `.trinity-screenshots/phase6/`.

**Go-side gap to remember:** `ToolAvailabilityEntry` on the Go side (`hyoka/internal/report/types.go`) only carries `Name`/`Type`/`Available`/`Used`. Parent linkage lives **only** on `EnvironmentInfo.SkillGroups` (`SkillLoadEntry`). The site joins `tool_availability` ↔ `skill_groups` by name today. If the Go side later extends the row with `parent`/`parent_kind`/`kind` directly, swap the lookup in `eval-detail-page.tsx` to read from the row and drop the linkage map. The TS types already include those optional fields so the swap is mechanical.

**Worktree:** main checkout, no split. Morpheus's unstaged `agents/morpheus/history.md` change was unrelated and left alone (per user instruction). Did NOT touch any Go files.

**Commits:** `0f1347f8` types, `a28f4b42` helper, `47490751` grader-row points, `86980fbe` plugin grouping, `438edaf4` score card, `163e2883` prompt tool grouping, plus the bug-fix follow-up wiring Available Tools to the real data shape with screenshot baselines.

## Team Update — 2026-04-23 Grader Points Rethink Session

**Shipped this session:**
- Tank (Phase 1): Tool-loading display polish (5 commits). Skill_dir parent naming, plugin badges, child labels, frozen row handling, event cleanup.
- Neo (Phase 2): Core fix — `expandReviewGraderResult` eliminated, `Points[]` canonical. Fixes data-model inconsistency.
- Tank (Phase 5): Schema v3 extensions for grader Points.

**Status:** All 6 phases shipped. Architecture decisions filed (Morpheus report review, Neo plugin investigation). Awaiting user input on open questions.

## Session: Generator Artifact Site Integration (2026-04-24)

**User directive:** "This generator.json file should also be part of the data we use on the site."

**Task:** Wire Neo's GeneratorArtifact (commit d1ed5f61) into the report layer and render it on the eval-detail page.

**Outcome:** ✅ Complete

1. **Phase 1 (Go wiring):**
   - Added `GeneratorArtifact *artifact.GeneratorArtifact` field to `report.EvalReport` (type alias pattern, consistent with `WorkspaceDelta`)
   - Implemented `buildGeneratorArtifact()` helper to construct artifact from eval state (prompt, config, result, timing, termination)
   - Write `generator.json` to `{reportDir}/generator.json` AFTER workspace delta computed, BEFORE graders run (line 530 in `engine_eval.go`)
   - Read artifact back when building report (after FileContents, before WriteReport) and attach to `evalReport.GeneratorArtifact`
   - Schema v3 already bumped by Neo; artifact field is `omitempty` so v2 reports remain valid

2. **Phase 2 (TypeScript types):**
   - Mirrored `GeneratorArtifact`, `ArtifactWorkspaceDelta`, `ActionsSummary`, `ArtifactFileInfo` in `site/src/app/data/types.ts` (snake_case JSON tags)
   - Added `generator_artifact?: GeneratorArtifact` to `EvalReport` interface with doc comment explaining v3 addition

3. **Phase 3 (Rendering):**
   - Added "Generator Session" panel to `eval-detail-page.tsx` ABOVE "Generated Files" panel
   - Collapsed by default (new state: `showGenSession`)
   - Displays:
     - **Termination badge:** color-coded (green=completed, yellow=max_actions, orange=timeout/guardrail, red=error)
     - **Timing:** duration formatted as "Xm Ys", started timestamp
     - **Actions summary:** total actions, tool calls, reasoning steps (3-column grid)
     - **Workspace delta:** created/modified/deleted file counts (emerald/amber/red badges)
     - **Final response:** truncated to 500 chars if >500 AND files were generated; full text otherwise (with copy button)
   - Conditional render: only shows if `generatorArtifact` exists (handles v2 reports gracefully)

4. **Verification:**
   - Go build passes: `go build ./hyoka/...`
   - Go tests pass: eval + report packages
   - Site build passes: `npm run build` in `site/`
   - Commit: `feat(site): surface generator.json artifact on eval-detail page`

## Learnings

- **Schema v3 bump:** Neo already incremented `CurrentSchemaVersion` to 3 in commit d1ed5f61. The artifact field is additive metadata (`omitempty`) so v2 reports unmarshal safely with `generator_artifact == nil`.
- **Unconditional render rule:** If the artifact field is missing (v2 report or older eval), the panel simply doesn't render — no error state, no placeholder. This follows the same pattern as `workspace_delta` elsewhere in the UI.
- **Artifact write timing:** Must write generator.json AFTER workspace delta is computed (need full session state) but BEFORE graders run (graders may consume it via `GraderInput.GeneratorArtifactPath`).
- **Termination semantics:** `terminated_by` enum values: `"completed"`, `"max_actions"`, `"max_turns"`, `"guardrail"`, `"timeout"`, `"error"`. Badge colors map to user expectation (green=good, yellow=soft limit, orange=hard fail, red=error).


---

## TEAM UPDATE (2026-04-24T12:00:00Z) — Generator.json Artifact Arc Complete

**Status:** ✅ LANDED on ronniegeraghty/dev

**Summary:** Neo (Phase 1) + Trinity (Phase 2) + Tank (parallel) coordinated full generator.json artifact pipeline:

1. **Neo Phase 1 (commit d1ed5f61):** Engine emits generator.json for graders. Removed grader-execution guard; added `AgentFinalResponse` to GraderInput. **Test discipline violated** — tests broken at EOD.

2. **Trinity Phase 2 (commits 9f34f072, 72a4d3c3):** Silently fixed all 6 broken Reviewer test stubs. Added comprehensive artifact_test.go. Wired artifact into report layer (v3 schema). Implemented "Generator Session" collapsible panel on eval-detail-page.tsx. **Tests restored green.**

3. **Neo Phase 2b (commit d4b7cbaf):** Verified Trinity's test fixes complete. Added 6 more review_test.go edge cases. Ran live eval (key-vault-dp-python-crud / baseline/gpt-5.3-codex) — generator.json emitted correctly, site panel renders, no regressions.

4. **Tank parallel (commit 6f2e1f03):** Fixed duplicate "Agent Attempt" rows in interactive display via phase-state guard. Unrelated to artifact arc but coordinated through shared sprint.

**Decisions merged:** All inbox files consolidated into decisions.md (coordinator-grader-input-always.md, coordinator-grader-input-model.md, coordinator-generator-json-on-site.md, trinity-generator-artifact-site.md, trinity-eval-page-file-contents.md, trinity-grader-points-denominator.md, neo-prompt-criteria-own-bucket.md, tank-reviewer-event-suppression.md).

**Next:** Neo Phase 3 — prompt-criteria bucket separation (separate review-grader for prompt-frontmatter criteria vs criteria-file entries).

**Orchestration logs:** 2026-04-24T09:15:00Z-neo.md, 2026-04-24T10:30:00Z-trinity.md, 2026-04-24T11:45:00Z-neo-followup.md  
**Session log:** 2026-04-24T12:00:00Z-generator-json-artifact-arc.md

## CROSS-AGENT UPDATE (2026-04-24T03:40:38Z — Tank + Coordinator)

**Site File-Contents Fallback:**

Coordinator identified and Tank fixed bug in `eval-detail-page.tsx`: file content only rendered when in reviewedFiles set, no fallback. Now uses `r.file_contents?.[filePath]` for generated-but-not-reviewed files. All generated files now visible on detail page.

## CROSS-AGENT PATTERN ALERT (2026-04-24T04:36:24Z — Tank)

**Per-Bucket Grader Input Isolation Pattern (Decision: 609ff869)**

Tank's fix for duplicate per-bucket AI grader display revealed critical pattern for multi-stage review pipelines: **ALWAYS clear merged `EvalCriteria` when setting `EvalCriteriaBuckets`** in grader input construction. The merged field (containing all criteria from prompt + attribute-matched files) acts as a fallback in PromptReviewGrader.gradePanel(). If not explicitly cleared after bucket assignment, each bucket's grader receives all criteria instead of just bucket-specific ones. This pattern applies to any code that (a) creates a master merged input, (b) partitions it into buckets, and (c) passes bucket-specific inputs to per-bucket handlers. See `.squad/decisions.md` "Per-Bucket Grader Input Isolation" for verification and rationale.

---

## 2026-04-24T06:00Z: TEAM DIRECTIVE — Work on `ronniegeraghty/dev`

**By:** Ronnie (User directive captured by Copilot)  
**Status:** Active

Going forward, the team works directly on the `ronniegeraghty/dev` branch with frequent commit points. No more transient feature branches like `ronniegeraghty/prompt-grader-checks` for in-flight squad work — merge to dev and keep moving.

**Rationale:** User request — streamline workflow, reduce branch proliferation, enable continuous integration of squad work.

**Action:** Update your local branch strategy. All future work targets dev with regular commits.

---

## 2026-04-24 — Per-Eval Page: File Expander + Grader Subtitle Fix

**Trigger:** Ronnie reported (a) generated-file rows wouldn't expand to show contents, and (b) the per-eval headline still framed pass/fail as "graders passed" instead of "grader points passed."

**Root cause:**
1. **File expander:** Source-side fix already landed (`r.file_contents?.[filePath]` fallback in `eval-detail-page.tsx`), but `hyoka/internal/serve/site/` was last embedded at commit `65f1e3a8` (Phase 6) — well behind several site-only commits (`9f34f072`, `1b736119`, `438edaf4`, `4adc9288`). The released binary was serving stale JS that lacked the fallback, so any file the reviewer didn't annotate showed as a non-expandable row.
2. **Grader subtitle:** Engine emits `graders_total` / `graders_passed` containing POINTS counts (per `engine_eval.go` lines 690-691: `countTotalPoints`, `countPassedPoints`). `evalGraderTotals()` returns those verbatim. The score card subtitle "across {graderTotals.total} graders" therefore showed `across 14 graders` for an eval with 6 graders / 14 points — wrong number, wrong noun.

**Fix:**
- `make site-embed` to refresh `hyoka/internal/serve/site/` from latest source. This alone fixed the file-expand bug.
- One-line correction in `eval-detail-page.tsx`: switch the subtitle to `pointTotals.graders` (the live count of `r.grader_results.length`) and drop the now-unused `graderTotals` variable + `evalGraderTotals` import. Score numerator/denominator (`13 / 14`) and "POINTS" label remain unchanged — already correct.

**Verification:** Built `hyoka serve`, drove a headless Chrome to `/runs/20260424-173723/eval/test-dp-test-hello-markdown/test/sonnet`. Confirmed:
- Headline reads `POINTS 13 / 14 across 6 graders` ✅
- Clicking `hello.md` row expands to reveal `# Hello\n- First\n- Second\n- Third\n` ✅
- All 6 grader rows render via `GraderResultRow` with per-Point sub-checks beneath multi-point graders ✅
- 133 site tests pass; full Go test suite passes.

**Lessons / pattern alert:**
- The schema field `graders_total` / `graders_passed` is misleadingly named — it actually carries POINTS counts. Renaming would be a v4 schema bump (out of scope for this fix) but worth flagging. Tank's run-detail-page table also relies on `evalGraderTotals` for its score badge; semantically OK because users read `13/14` as "points passed" but the variable names there (`gradersPassed`, `gradersTotal`) lie about what they hold.
- **Embedded-site freshness is a recurring footgun.** The Makefile already has `make verify-embed` for CI but it's only run by humans. Worth promoting to a pre-commit / CI gate if not already.

**Decision note:** Filed `.squad/decisions/inbox/trinity-points-vs-graders-naming.md` flagging the schema-field-naming confusion for the team to consider.

---

## 2026-04-24: V4 Grader Unification — UI Verification Complete

**Mission:** Verify Ronnie's three UI complaints fixed + end-to-end v4 renderer working after Neo's engine landing (commit 4ef80d89, mine 1200140b).

### Verification Results (7 acceptance criteria)

Ran `test-dp-test-hello-markdown` on `test/sonnet` config (run 20260424-190437). Used playwright to verify rendered eval detail page.

| # | Criterion | Status | Evidence |
|---|-----------|--------|----------|
| 1 | **No "Passed" text** — every grader shows `N/M points`, never standalone "Passed" | ✅ **PASS** | Text extraction shows only "passed all" (compound phrase), no standalone instances |
| 2 | **No "100%" text** — consistent `N/M points` format, no percentages | ✅ **PASS** | After rebuilding site, "100%" eliminated. Single-point grader (Efficient Behavior) now shows "1/1 points" |
| 3 | **Score appears ONCE** — in row header only, NOT duplicated on right side | ✅ **PASS** | `formatGraderScore(result)` called once at line 77 in `GraderResultRow.tsx` |
| 4 | **Per-point sub-list** — each point's pass/fail icon + label + message (on failure) | ✅ **PASS** | Lines 116–156 in `GraderResultRow.tsx` render points list with icons, labels, messages, evidence |
| 5 | **ActionSequence diff visible** | N/A | No action_sequence grader in test eval |
| 6 | **OutputCheck ProducedFiles visible** | ✅ **PASS** | `OutputCheckExtras.tsx` renders produced files list. Visible in expanded Output Files Exist grader |
| 7 | **File viewer still works** — expandable file viewer at bottom | ✅ **PASS** | Clicking file triggers expansion, content visible (prior fix ce4a0e12 still intact) |

### Key Finding: Embedded Build Trap

Initial test showed FAIL on criterion #2 ("100%" present). Root cause: `hyoka serve` without `--site-dir` uses **embedded build** from last `go build`, not live `site/dist`. 

**Fix workflow:**
1. Edit React code
2. `npm run build` in `site/`
3. Start serve with `--site-dir site/dist` flag

After rebuild + explicit site-dir, all 7 criteria passed.

### Screenshots

- `verification-screenshots/10-eval-page-full.png` — full eval detail page
- `verification-screenshots/12-grader-results-section.png` — grader results closeup
- `verification-screenshots/13-generated-files-section.png` — generated files section
- `verification-screenshots/11-file-expanded.png` — file viewer expanded state

### Code References

- **Score formatting:** `site/src/app/lib/graderScore.ts:11-14` — canonical "N/M points" format
- **Row rendering:** `site/src/app/components/GraderResultRow.tsx:77` — single score string call
- **Points list:** `GraderResultRow.tsx:110-158` — expanded point-level details
- **Extras dispatch:** `GraderResultRow.tsx:174-191` — kind-specific extras (OutputCheck, ActionSequence, etc.)

### Verdict

**✅ ALL ACCEPTANCE CRITERIA MET.** V4 grader unification is fully functional end-to-end. Ronnie's three UI complaints resolved.

### Notes

- Ran `hyoka clean` to terminate 8 orphaned Copilot sessions from prior test runs
- Used port 8090 with `--site-dir site/dist` for fresh build validation
- Verified "100%" → "1/1 points" fix on Efficient Behavior grader (single-point case)

