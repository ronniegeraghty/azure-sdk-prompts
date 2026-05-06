# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Core Context

**Archived 23 entries from earlier sessions.**

Historical patterns and learnings:

- ## Core Context: Agent Morpheus initialized as Architect for hyoka. Charter: architecture reviews, live verification gates, dogfood testing. Tools: Playwright for li...
- ## 2026-04-20: Fixed #364 Test Mocks + Live Verification: **Context:** Switch double-rejected #364 due to disabled tests. Trinity and Oracle locked out.

**Actions:**
1. **Part A - Test Mock Fix:**
   - Ren...
- ## 2026-04-20: Phase 5 Architectural Review — PR #592: **Task:** Architectural review of PR #592 (phase-5 → ronniegeraghty/dev), separate from Playwright verification.

**Review Dimensions:**
1. **Plan a...
- ## 2026-04-20: Phase 5 Fixups — PR #592 R151 Gap Closure + #596 Verification: **Task:** Closure of R151 acceptance criteria (#596) discovered missing during Architect review.

**Investigation:**
- **Requirement 1:** "Pass Rate...
- ## 2026-04-21: PR #605 Architectural Review — WI-027 Tool Versioning & Custom Fetchers: **PR:** #605 (Neo) — `ronniegeraghty/issue-597-tool-versioning` → `phase-6`
**Issue:** #597 (builds on #334/WI-026 cache isolation)
**Verdict:** ⚠️...
- ## 2026-04-21: PR #606 Architectural Review — #599 `group` property (Phase 6 R2): **Author:** Neo. **Branch:** `ronniegeraghty/issue-599-group-property` → `phase-6`. **Verdict:** ⚠️ APPROVE WITH NOTES.

**Coordination context:** T...
- ## 2026-04-21: PR #607 Rollup Integration Review — Phase 6 (#312): **PR:** #607 (`phase-6 → ronniegeraghty/dev`). **Verdict:** ⚠️ **APPROVE WITH ONE BLOCKING FIX**.

**Arch pass:** Cross-feature contracts clean. `#6...
- ## Session 2026-04-21 (Phase 6 Round-1 Architectural Review): **Mission:** Architectural review of Phase 6 Round-1 batch (PRs #601, #602, #603) + live Playwright verification + embedded-asset fix

**Verdicts (a...
- ## #608 — Embedded asset freshness automation (2025): PR #611 → phase-6. Closes #608. Systemic follow-up to the #607 rollup catch.

**Two layers of defense shipped:**

1. **Makefile** (new, top-level) w...
- ## 2026-04-21 — PR #610 architectural review (Tank: #606 group property polish): **Verdict:** ✅ APPROVE (posted as comment — gh blocked self-approval since PR author is ronniegeraghty)

**Scope:** Test-only PR adding 3 files (217...
- ## 2026-04-21 — PR #612 Architectural Review: ✅ APPROVE: **PR:** #612 (Neo) — fetcher cleanups, ctx threading, signature flatten, dead code removal. Phase 6 polish off #605.

**Verdict:** APPROVE. Three ch...
- ## 2026-04 — PR #609 Review (Trinity, MultiSelectFilter tests, issue #608): ✅ APPROVE (posted as comment — gh blocks self-approve on own PR account).

Test-only PR adding 3 vitest cases to `site/src/__tests__/multi-select-fi...
- ## 2026-04-22 — PR #613 Architectural Review (Trinity, MultiSelectFilter follow-up tests): **Verdict:** ⚠️ APPROVE WITH NOTES (posted as comment — gh blocked self-approval since Ronnie is PR author).

**Scope:** 232-line test-only addition...
- ## 2026-04-21 — PR #614 (authored): site-embed freshness CI hardening: **Branch:** issue-608 follow-up worktree → **PR #614** (target `phase-6`, squash-merged at `e0b72c63`)

Systemic follow-up to my own PR #611 address...
- ## 2026-04-21 — PR #607 Final Architectural Review (Phase 6 → dev Rollup): **PR:** #607 (`phase-6` → `ronniegeraghty/dev`), rollup of six Phase 6 sub-PRs (#601, #602, #603, #604, #605, #606) plus round-3 polish (#613, #614)...
- ## 2026-04-22 — PR #618 Architectural Review (WorkspaceDelta first-class, Issue #566): **PR:** https://github.com/ronniegeraghty/hyoka/pull/618 — `squad/566-workspacedelta-firstclass` @ `2e67bc51`
**Verdict:** ✅ APPROVE (posted as `--c...
- ## Learnings: ### PR #618 merged into phase-6

The approval verdict was posted and nits addressed. Guardrail policy locked in for team — "hard-fail only, no soft-...
- ## 2026-04-22 — Validate audit of `examples/`: ### Learnings

**Validate command scope** (`hyoka/cmd/validate.go`):
- Scans three dirs: prompts (via `--prompts` flag, default `./prompts`, with `....
- ## Learnings (PR #607 hierarchical-when-example.yaml): - **Multiple group-level `when`s** in a single criteria file are expressed via the top-level `groups:` list on `GraderConfig` (`hyoka/internal/crite...
- ## Learnings (Grader Unification Architecture — #622): **Date:** 2026-04-22

**Key insight:** The grading system has two disconnected halves:
- `hyoka/internal/criteria/` — schema `GraderEntry{Name, Weig...
- ## Learnings (Proposal Recovery — 2026-04-23): **Date:** 2026-04-23

**Incident:** The 513-line grader unification proposal (`morpheus-grader-unification-proposal.md`) was deleted from the inbox...
- ## Learnings (Grader Unification — Issues Filed, 2026-04-23): **Locked schema decisions** (after Q&A with Ronnie via Copilot directives):
- **Discriminator:** flat `type` field at entry level. Prompt graders us...

Full history archived. Recent entries below.

---

## Learnings (Option A Pivot — 2026-04-22)

**Pivot:** Flat `hyoka/internal/graders/` rejected after Phases 1–3 had already shipped there. New directive (`copilot-directive-grader-package-layout.md`) locks Option A: nested `hyoka/internal/criteria/` (file-level) + `hyoka/internal/criteria/graders/` (grader-level). The package hierarchy must mirror the YAML reality — criteria files *contain* graders.

**In-flight commit resolved:** Neo's background Phase 3 spawn landed 46b624fb before replan finished — it deleted the legacy `internal/criteria/` and migrated `cmd/list` + `cmd/validate` to the flat `internal/graders/` target. The deletion is still correct; the import targets get rewritten as part of #628.

**Issues filed:**
- **#628** — Phase 3 (Option A): Restructure unified grader package to nested layout. `squad:neo`. Includes full file-by-file move map and a rename pass that drops the `Unified*` prefix Phase 1 used as a coexistence aid (`UnifiedGraderConfig` → `criteria.Config`, etc.).
- **#629** — Phase 4 (Option A follow-up): Doc + example path sweep after #628 lands. `squad:oracle`.

**Issues updated:**
- #626 (closed): comment pointing at #628.
- #627 (open): comment noting path moves; functional scope unchanged.

**Process takeaway:** When a locked directive arrives mid-rollout and some phases have already shipped, the replan commit for Phase N+1 should explicitly map each shipped commit to "still correct" vs "needs rewrite" — saves the next implementer from guessing which earlier work is safe to build on.

---

## Learnings (Per-Reviewer Vote Display Investigation — 2026-04-28)

**Date:** 2026-04-28  
**Task:** Investigate "missing per-reviewer votes" regression in grader card UI; produce scoped plan for restoration

**Key Findings:**

1. **Data is already present in JSON ✅** — No engine changes needed
   - Path: `grader_results[i].extras.review.panel_results[j].criteria[k]`
   - Go type: `hyoka/internal/criteria/graders/grader.go:216-232` (`ReviewPanelResult.Criteria`)
   - Each panel member writes full per-check data: `{name, passed, reason, weight}`
   - Verified in live report: `reports/20260428-175710/.../claude-opus-4.6/report.json`

2. **TypeScript type mismatch** — Frontend is missing the `criteria` field
   - File: `site/src/app/data/types.ts:27-34` (`ReviewPanelEntry`)
   - Type has `overall_score`, `max_score`, `summary` but **no `criteria` field**
   - JSON has `criteria: [{name, passed, reason}, ...]` but TS interface doesn't declare it
   - Root cause: `ReviewPanelEntry` predates v4 grader unification (commit `1200140b`)

3. **Rendering is incomplete** — ReviewExtras component doesn't loop over criteria
   - File: `site/src/app/components/grader-extras/ReviewExtras.tsx:68-96`
   - Current: shows `panel.model` + `panel.overall_score` + `panel.summary`
   - Missing: no rendering of `panel.criteria[]` array (per-check pass/fail + rationale)

4. **Historical note** — This is a **new feature**, not a regression restoration
   - No HTML templates ever existed with this feature (checked git history: `0ff7b486`)
   - The only old template was `templates/prompt-template.prompt.md` (a prompt file skeleton)
   - User's memory of "old reports showing per-reviewer votes" may be from manual JSON inspection

**Scoped Plan Deliverable:**

- Written to `.squad/decisions/inbox/morpheus-grader-vote-display.md`
- Three changes needed (all frontend):
  1. Add `criteria?: ReviewCriterionResult[]` to `ReviewPanelEntry` type
  2. Add per-check rendering loop in `ReviewExtras.tsx` (icons + name + reason)
  3. Add test coverage for criteria rendering
- Owner: **Trinity** (all site changes)
- Estimate: ~1 hour (types + render + test)
- No Go changes — engine already writes correct data

**Process Insight:**

When investigating a "missing feature" report:
1. **Verify data presence first** — check JSON before assuming engine bug
2. **Compare Go vs TypeScript types** — schema drift happens when one side evolves faster
3. **Check git history for "old way"** — user memory isn't always reliable; validate with commits
4. **Distinguish regression vs new feature** — affects priority and messaging

This pattern (data exists in JSON, frontend type/render missing) is a type-safety gap — TypeScript doesn't enforce that frontend types match backend JSON shapes. Consider:
- Generating TS types from Go structs (via tool)
- Schema validation tests (read JSON, assert against TS interface)
- Documenting type sync checkpoints in rollout plans

**Plan:** `.squad/decisions/inbox/morpheus-option-a-replan.md`. Awaiting Ronnie approval before restart.

---

## Learnings (Plugin/Skill/MCP Loader Diagnosis — 2026-04-23)

**Requested by:** Ronnie. Symptom: "plugins and skills never seem to successfully load." Deliverable: diagnosis + work-item plan (no implementation). Plan saved to `.squad/decisions/inbox/morpheus-plugin-loader-plan.md`.

### Loader flow map (end-to-end)

```
YAML → config.ExpandPlugins (silent WARN on miss)
     → copilot.go buildSessionConfig:
         - EmitPluginResolutions (progress only, read-only)
         - EmitMCPResolutions (static field validation)
         - ResolveSkillsWithReporter (generator only; silent WARN + nil,nil on miss)
     → client.CreateSession → SDK indexes skills/spawns MCP
     → OnEvent(SessionSkillsLoaded/SessionMcpServersLoaded) → verifier records load
     → [DISABLED 4b593d3b] waitForToolVerification gate
     → SendAndWait
Reviewer skill path: cmd/run.go passes entry.Path raw to SetSkillDirectories —
  no ResolveSkills, no glob, no skill_dir expansion. Cross-config leakage
  (reviewer paths pooled across all matched configs, not scoped per-config).
```

### Failure modes observed in live runs

1. **Plugin refs like `azure-sdk-python@skills`** — marketplace naming, never present in local `plugins/*.yaml` (which defines only `azure-python`), and `~/.copilot/installed-plugins/` is empty on this machine. All 6 plugin refs across baseline-skills and python-pairwise configs silently skipped on every invocation.
2. **Generator skill `azure-sdk-for-rust-bestpractices`** loads on every eval regardless of prompt language. Python prompt with Rust skill → 0 `skill.invoked` events across 4 turns.
3. **Extra skill `customize-cloud-agent`** reported as loaded — SDK builtin leaking past the "isolated ConfigDir" (#21) comment. Not a config bug, but the verifier's loose name-matching hides it.
4. **Verification gate disabled** (commit 4b593d3b). The SDK emits `SessionSkillsLoaded`/`SessionMcpServersLoaded` only AFTER `SendAndWait` begins; the original gate blocked BEFORE `SendAndWait`, causing a 10s timeout deadlock on every eval. Neo disabled observationally; no post-session fallback was added. Today tool load failures are logged (WARN) but never block.
5. **Every `hyoka list` / `hyoka run`** reloads all 13 configs and re-walks `plugins/` 13 times. Warnings for configs not even selected fire on every invocation. Noisy, not broken.
6. **Reviewer skill resolution gap**: `cmd/run.go:378` loops over ALL configs (not the selected one) and passes raw `entry.Path` strings to `SetSkillDirectories`. `skill_dir: true` is silently ignored. Missing paths produce no error.

### Silent-failure inventory

10 distinct silent-WARN-continue points traced (F1–F10 in the plan). Every one must be either converted to a structured error OR stop claiming the load happened.

### Hard-fail decision rationale

Silent degradation turns an eval into a non-representative comparison — the run "passes" but measures a different experiment than the config declared. That's worse than a failed eval because it corrupts trend data. The prior gate was a bad implementation (timing mismatch with SDK events), not a bad idea. The fix is **static pre-session validation** (we know at config-load time whether a plugin exists, a skill dir has SKILL.md, an MCP server has its command) plus keep the post-session SDK-event verifier as secondary confirmation. Hard-fail by default; add `required: false` opt-out only if a real use case demands it.

### Plan structure (summary)

- **WU-1 (Neo)** — `tool.ValidateAndExpand` with structured `ToolLoadReport` + `ToolLoadError`, wired into `copilot.go buildSessionConfig` before `CreateSession`. Returns `error_category: "tool_load_failure"`.
- **WU-2 (Neo)** — Reviewer skill resolution parity; move resolution inside per-config `reviewerFactory` closure.
- **WU-3 (Tank)** — Expanded Tools display grouping leaves under parent plugin/skill_dir.
- **WU-4 (Switch)** — Table-driven tests for each failure mode + golden render tests.
- **WU-5 (Oracle)** — Docs for hard-fail semantics + Tools format + plugin-install guidance.

### Schema findings

- `plugins:` uses marketplace syntax (`name@skills`) but the local registry uses plain names. No adapter bridges them. Either install into `~/.copilot/installed-plugins/` or rename local YAMLs to match — neither is done today, so every `plugins:` entry silently skips.
- `plugins/azure-python.yaml` is defined but referenced by nothing.
- `skills/generator/` contains one Rust skill, fired for all languages. Content-library problem, not a loader problem (flagged as out-of-scope).

### Out of scope recorded

- Re-enabling the SDK-event post-`SendAndWait` gate (defer to Phase 2).
- `required: false` opt-out (defer until use case appears).
- Generator skill content (file separate issue).
- Plugin-registry re-walk performance nit.

## Learnings (Issue #305 Status Verification — 2026-04-23)

**Requested by:** Ronnie (Issue triage). **Task:** Verify if v0.3.1 was released and decide whether to close issue #305.

### Key Findings

**Release Status:** NO GitHub Release or git tag for v0.3.1 exists. Only one release in repo: `sdk-eval v0.1.0` (tag: `tool/v0.1.0`).

**Phase Completion:** All 7 phase sub-issues (#306–#312) are CLOSED. However:
- Phases 0–2 work merged to main via PRs #558–#560
- Phases 3–6 code exists on `ronniegeraghty/dev` branch only (NOT on main)
- CHANGELOG.md with v0.3.1 section exists only on dev, not main

**Issue Checklist:** All 7 phase items remain unchecked ([ ]) in issue body.

### Decision: KEEP OPEN ✓

Rationale: #305 is a legitimate tracking issue for **unreleased v0.3.1** work. The fact that phase sub-issues are closed does NOT mean the release is done — code must merge to main and be tagged. Current state:
1. Work is tracked and phase-closed ✓
2. Work is NOT released ✗
3. CHANGELOG drafted on dev ✓

### Prior Audit Correction

The morpheus-issue-audit.md (2026-04-22) flagged #305 as "probably stale — likely shipped but no git tag found." **This was incorrect.** Investigation reveals:
- Work genuinely exists and is phase-complete
- But it's never been merged to main or tagged
- Leaving it open is correct; closing without a release would lose accountability for the integration/release work

### Recommendation for Ronnie

1. Clarify release intent: Ship v0.3.1 now (main + tag) or defer with Phase 7?
2. If ship now: Merge dev→main + `git tag v0.3.1 && git push origin v0.3.1` + create GitHub Release
3. If defer: Keep #305 open; update phase checklist as phases merge to main
4. Once tagged/released: Close #305 with tag reference

### Time Spent
~20 minutes (release verification + phase state checks + branch analysis)

---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.

---

### 2026-04-23T18:52Z: Cross-agent update — Plugin schema BREAKING CHANGE (Neo, commit `2c1de1c0`)

Per Ronnie directive, Neo reversed his earlier `@marketplace` validator (commit `769dea69`) and removed the hardcoded `microsoft/skills` magic from `plugin.ResolveInstalled`. New rule: remote plugin entries MUST declare `repo:` explicitly. Names with `@` are now rejected at validation. Affects any config audits you do — the canonical form is now:

```yaml
- name: azure-sdk-python
  type: plugin
  source: remote
  repo: github.com/microsoft/skills
```

This repo's configs (`configs/python-pairwise.yaml`, `configs/baseline-sonnet-skills.yaml`) are already migrated. Any wild config not in this repo using `name@skills` will fail validation with a migration message. See `decisions.md` entry at 2026-04-23T18:50Z for full schema and validator contracts.

### 2026-04-23T19:42Z: Plugin-loading saga closed end-to-end

Neo shipped `4a8c4a0d` — container plugins now fan out into per-child `ToolLoadItem`s, verifier matches by child basename. Live verified: `hyoka run` against `python-pairwise` config goes from 3/3 errors to 0/3, all 41+ azure-sdk-python children load. Combined with `2c1de1c0` (explicit `repo:`) + `3b306c9` (canonical `owner/repo` form), the plugin schema/loader story is now coherent: locator + content shape both contracted. If any prior audit referenced lingering plugin loader issues, they're resolved.

## 2026-04-23 — Report data model review (fan-out + Points)

- **Discrepancy reproduced.** Run `reports/20260423-195948`: summary "12/12 passed", but run-detail rows show red `1/4` / `0/0`. Eval-detail headers show ✅ correctly (use `r.success`); only `run-detail-page.tsx:236-237` invents its own roll-up via `grader_results.filter(g => g.pass === true)`.
- **Root cause is `expandReviewGraderResult` (engine_eval.go:903-953):** one passing `ai_review` grader becomes 3 report-side entries with `Pass:nil` and `Score:0`. The engine's `agg.Pass=true` only survives as top-level `EvalReport.Success`. Per-grader truth is destroyed in the conversion.
- **`Success=true` with zero `grader_results`** (gpt-5.3-codex no-files case): grading block is skipped at `engine_eval.go:433` `if len(generatedFiles) > 0`. Worth flagging as a separate semantics question.
- **`report.SessionSetup.Skills` is flat** — no parent/child linkage. After plugin/skill_dir fan-out, can't express "plugin parent has no status, only children do." Need `Parent`, `ParentKind`, `Kind` fields on `ToolLoadResult` (additive, omitempty, no schema bump).
- **Phase 2 `Points` should NOT be additive-only**: stop the review expansion; emit one `GraderResult` with unified `Pass`+`Points`. Bump `SchemaVersion` to v3 (semantics change in entry-count-per-grader). Old v2 reports keep their expanded shape; v3 uses the new 1-entry-with-Points shape.
- **Single source of truth recommendation**: add `EvalReport.GradersPassed/GradersTotal` populated at engine time. Eliminates roll-up divergence by construction.
- **Site is mostly correct** — `run-detail-page.tsx:236` is the only bad roll-up; everywhere else trusts `EvalReport.Success`.
- **My boundary with Trinity**: I own data model + roll-up logic; she owns site presentation. The `run-detail-page.tsx:236` fix is mine to specify, hers to implement.
- Wrote full assessment to `.squad/decisions/inbox/morpheus-report-architecture-review.md`. Screenshots in `/tmp/morpheus-site-review/`.

## 2026-04-23 — Synthesized report+site reviews into plan
Read trinity-site-ux-review.md and my own morpheus-report-architecture-review.md, appended Phases 4–6 to the session plan.md (Phase 4: Trinity site quick wins, independent; Phase 5: report schema v3 — Tank/Neo, sibling to Phase 2; Phase 6: site Phase-2 alignment, depends on both). Inserted 20 new SQL todos (p4-* x6, p5-* x7, p6-* x7) with dep edges so anything consuming Points/parent-linkage waits on Phase 2 + Phase 5. Made the structural connection explicit: Phase 2's grader Points work IS the fix for the "1/4 red on every row" bug — once each grader emits one GraderResult with Points[], expandReviewGraderResult disappears and the site stops seeing 3 nil-pass rows per ai_review. Did not touch existing Phases 1–3 (Tank mid-execution).

## Learnings

- **#634:** Show generated files as diffs against starter project — Filed issue for eval-detail diff visualization feature (gen vs. starter file comparison).

## 2026-04-24 — Learnings: hardcoded taxonomy drift (issue #635)

- **Bug pattern.** `Valid*` slices in `internal/validate/validate.go:14-42` are duplicated as inline map/equality literals in `internal/validate/schema.go:176-213` (`isValidPlane`, `isValidLanguage`, `isValidService`, `isValidCategory`, `isValidDifficulty`). The two sets have already drifted: `validate.go:26` has `"test"` in `ValidLanguages`; `schema.go:181-189` does not. Adding any taxonomy value requires editing two files, and forgetting one half silently rejects valid inputs at `schema.go:63`.
- **Escape-hatch smell.** `isTestValue` (`schema.go:215-219`) and the appended `"test"` entries in services/languages/categories are workarounds for the same problem — should disappear once discovery lands.
- **Third-place duplication.** `planeAbbrev` map at `validate.go:53-56` is yet another hardcoded plane-axis location, used by the ID-prefix check at `schema.go:166`. Discovery design must accommodate per-axis structure (abbreviations) for plane.
- **Consumers are small.** Only `cmd/new_prompt.go:14-32` consumes the exported `Valid*` slices outside the validate package itself. Site (`internal/serve/dashboard.go:240-243`) takes filter values as opaque strings — no server-side allowlist, so taxonomy discovery is API-compatible for serve.
- **Proposed shape (filed as #635).** `internal/taxonomy` package walks `prompts/`/`criteria/`/`configs/` once per process, unions observed values, optionally augmented by a forward-declaration `taxonomy.yaml` at repo root. Validation switches to set-membership + Levenshtein "did you mean" suggestion. Single-PR migration; primary owner suggested as Neo.
- **Issue:** https://github.com/ronniegeraghty/hyoka/issues/635

## 2026-04-24 — Scoped: Prompt grader `checks:` field

- **Scope doc:** `.squad/decisions/inbox/morpheus-prompt-grader-checks-scope.md`. Two-collaborator split: Neo owns schema + bucket text rendering + YAML migration + grader-side log; Tank owns badge format + report Points verification + e2e smoke. Parallel-safe.
- **Backward compat call:** Hard-migrate the two affected files (`criteria/language/python.yaml`, `criteria/language/test.yaml`) — both currently smuggle `1. … 2. …` numbered checks inside a single `prompt:` blob. java.yaml/rust.yaml use one-grader-per-criterion already, no migration. Single-`prompt:`/no-`checks:` case continues to work unchanged.
- **Execution model:** ONE LLM call per grader with N checks rendered as N numbered criteria — NOT one call per check. The existing review prompt (`internal/review/prompt.go:60-75`) already says "Each criterion … MUST appear exactly once in the criteria array," and the parser at `prompt_review_grader.go:120-127`/`:199-206` already maps returned criteria 1:1 into `result.Points`.
- **Renderer state:** Multi-Point nested rendering ALREADY EXISTS in `display_interactive.go:1003-1062` (`renderGraderWithPoints`). The reason YAML graders look bad isn't missing renderer code — it's that each YAML entry produces only one Point because the LLM gets only one criterion. Fixing the bucket text rendering (Neo) makes the existing renderer (Tank just tweaks badge format) Just Work.
- **Path unification:** REJECTED introducing a shared `PromptCheck` struct. The prompt-file path's `ParsedCriteria` and the YAML path's `Checks` already converge at the rendered-bucket-text layer. No need to lift the abstraction higher.
- **Two open questions** with reasonable defaults documented in scope (truncation behavior, preamble visibility) — implementation should not block on them.

### Time spent
~25 minutes (read existing graders + bucket builder + renderer + criteria YAMLs + scope draft).

### Learnings
- The prompt-grader Points pipeline is fully wired end-to-end since Phase 2; the YAML-side gap is purely a bucket-text-rendering issue, not a missing-feature issue. Future scopes touching grader output should look at the bucket text first before assuming renderer work is needed.
- `internal/criteria/buckets.go:119` (`FormatUnifiedPromptEntries`) is the single render point for YAML prompt-grader → review-LLM input. Anything that needs to change how the LLM sees the criteria text passes through here.

---

## 2026-04-24: Cross-Agent Update — Feature Shipped & Merged ✅

**Session:** 2026-04-24T05:58:18Z  
**Status:** ✅ Complete

Team Feature: Skill Usage grader + intentionally-failing check shipped and merged to `ronniegeraghty/dev` (commit ff38a7ec). Coordinator (Tank) consolidated 4 commits. Grader display validates ✅/❌ per-check rendering. All downstream artifacts (orchestration logs, session logs, decisions) recorded.

**Directive:** Team now works directly on `ronniegeraghty/dev` with frequent commits (no transient feature branches for squad work).
- ## 2026-04-24 — Issue Triage: Open Issues vs. `ronniegeraghty/dev`: **Task:** Triage open GitHub issues against work completed on `ronniegeraghty/dev` branch to identify candidates for closure.

**Method:**
1. Listed 100 open issues via `gh issue list`
2. Analyzed commits on `ronniegeraghty/dev` vs `main` (100+ commits ahead)
3. Cross-referenced merged PRs (#607 Phase 6, #592 Phase 5)
4. Examined specific commit evidence for individual issues

**Findings:**

**HIGH CONFIDENCE — Completed on dev:**
- **#586** (Builtin skill leakage) — Fixed in commit `445fea76` "Fix user-level skills leaking into eval Copilot sessions"
- **#619** (Tool load guardrail) — Implemented via `ValidateAndExpand` + hard-fail logic in `tool_load_failure` path (commits `8c947c8a`, `5c75b47c`, `557bb83b`)

**MEDIUM CONFIDENCE — Evidence suggests completion:**
- Phase 5/6 epics and sub-issues (#311, #312, #364-#369, #597-#600, #580) — All **already closed** in GitHub (checked via `gh issue view`)

**DEFERRED / STILL VALID:**
- #635 (taxonomy discovery) — No implementation evidence found
- #634 (file diffs on site) — Site work in Phase 4-6 doesn't cover this
- #73 (embed Copilot CLI) — SDK is used but not embedded
- #77 (skill vs dir property) — No commits found
- #86 (summary tab) — Dashboard work exists but specific "summary tab" unclear
- #88 (build failure false pass) — No fix evidence
- #78 (Azure MCP tool investigation) — Investigative issue, likely ongoing
- #72, #71, #14 — SDK/cleanup issues, no clear resolution

**Recommendation:** Close #586 and #619 with clear commit references. Others require deeper investigation or remain open as valid future work.


## 2026-04-23 — Grader structural audit

Ronnie asked for a written audit of grader structure after noticing
the on-site report shows graders in vastly different shapes.

**Investigated:** 8 grader kinds in `hyoka/internal/criteria/graders/`,
the shared `GraderResult` / `GraderInput` types, the report-side
`report.GraderResult` marshalling layer, and `site/src/app/components/
GraderResultRow.tsx`.

**Key findings:**

1. The grader interface is healthy (3 methods, single input struct,
   single result struct).
2. Common surface: `Kind, Name, Score, Weight, Pass, Gate, Message,
   Points`. Phase 2 made `Points` the canonical sub-check channel,
   and the React renderer treats it as such — but the per-kind
   `*Details` structs still live alongside, half-redundant.
3. Three real structural problems:
   - `OutputCheckDetails` is dropped at the report-marshalling layer
     (`report/types.go` has no field for it). The agent's produced
     files + per-knob sub-checks die in transit.
   - `BehaviorGraderDetails` is a 14-field union shared by behavior /
     action_sequence / tool_constraint. Each grader sets a different
     subset, the renderer guesses, and `action_sequence`'s expected-
     vs-actual sequence is invisible on the site.
   - `prompt_review` flattens `OverallScore, MaxScore, Summary, Issues,
     Strengths, Scores, IsConsensus` to the top level of
     `report.GraderResult`. Every other grader leaves those fields
     zero-valued. One grader's shape is everyone else's noise.
4. Score semantics also drift: 0.0/1.0 for some, partial credit for
   others, "X/N points" for output_check. Renderer ad-hocs the display.

**Deliverable:** `.squad/decisions/inbox/morpheus-grader-structure-
audit.md` with full breakdown + three options:

- **Option A** — plumb missing data, no structural change. ~50 LOC,
  no migration. Doesn't fix root cause.
- **Option B** — promote `Points` to canonical, reduce per-kind
  `*Details` to a discriminated `Extras` union, split
  `BehaviorGraderDetails` into three typed extras, drop the flattened
  review fields from `report.GraderResult`. ~300–500 LOC, breaking
  change to report JSON v4. **Recommended.**
- **Option C** — single `SubCheck` model + free-form `Evidence map`.
  Re-opens DM4 (no `interface{}` in results). Not recommended.

Open questions for Ronnie noted in the inbox file: external report
consumers? prompt_review score semantics alignment? FileGrader's 0.5
partial-credit?

No code changed. Implementation work, if Ronnie picks B, would go
to Neo (engine + types) and Trinity (site renderer).

---

## 2026-04-24: Issue Triage False Positive — #586 Commit Analysis

**FLAGGED:** Morpheus's analysis of #586 ("Fixed by commit 445fea76") was a false positive.

**Evidence from Switch's empirical verification:**
- Commit `445fea76` addresses **user-level** skills leaking from `~/.config/github-copilot/`, NOT builtin skills
- Issue #586 explicitly names builtin skills from `~/.copilot/pkg/universal/{cli-version}/builtin-skills/`
- Live eval run shows `skills=customize-cloud-agent` still loading into sessions
- Builtin skill filtering requires `SessionConfig.DisabledSkills` population — not implemented

**Lesson learned:** Commit-evidence triage (reading commit messages/diffs) is insufficient for semantic distinctions. Empirical verification (running tests/evals) is required for confidence on "issue closed" claims.

**Action:** #586 remains open; requires session config implementation (Neo/Tank).



## 2026-04-23 — Grader Unification Plan (Option B greenlit)

Ronnie greenlit Option B from the structural audit and added concrete UI
requirements: row header shows ONE canonical score string (not "Passed"
AND "100%" AND "1/1 points" simultaneously), Points are the single
source of truth for sub-checks, every grader explains *why* each point
passed or failed.

**Three open questions decided (documented in plan):**
1. External report consumers? → No. Hard cutover to schema v4. Loader
   rejects v3 with explicit "regenerate" error. Reports are git-ignored.
2. `prompt_review` score semantics? → Fold into Points. Per-criterion
   Points carry `Weight = criterion max points` so weighted scoring
   survives. Drop `OverallScore`/`MaxScore` from report entirely.
3. FileGrader 0.5 partial credit? → Normalize. Two Points per file
   (`file present` + `pattern matches`) when Pattern is set. Drop 0.5.

**Plan delivered:** `.squad/decisions/inbox/morpheus-grader-unification-
plan.md` covers final `GraderResult` shape, `GraderPoint` contract
(label/pass/message/weight/evidence), per-kind mapping for all 8
graders, canonical score format (`N/M points` always), site rendering
rules, file-by-file change list for Neo and Trinity, schema migration,
test plan, and a phased Neo/Trinity work split with three sync points.

**Key invariant introduced:** `GraderResult.Pass` and `GraderResult.Score`
are derived from Points — never set independently. `NewResult`
constructor enforces it. Empty Points panics (config error).

**Behavior-family graders split:** `BehaviorGraderDetails` (14-field
union shared by behavior/action_sequence/tool_constraint) becomes three
single-purpose `*Extras` structs. action_sequence's expected-vs-actual
diff finally reaches the site. output_check's ProducedFiles finally
reaches the site (was being dropped at marshalling).

**Rendering simplification:** the 6-way `if (X_details)` cascade and
the 5-way `passed` derivation cascade in `GraderResultRow.tsx` both
collapse — passed becomes `r.pass`; body becomes `<PointsList>` (always)
+ `<KindExtras>` (single switch).

No code changed yet. Awaiting Neo (engine + Go types) and Trinity (site
+ renderer) to pick up. Suggested split documented in plan §9 with
parallelism opportunities — Trinity unblocked once Neo lands the type
definitions in `grader.go` (Sync 1).

## Session Complete: v4 Grader Unification (2026-04-24)

**Date:** 2026-04-24  
**Outcome:** ✅ SHIPPED

Audit and Option B plan merged into `.squad/decisions.md`. v4 schema decision locked: "N/M points" format, auto-derived Pass/Score from Points, discriminated Extras union, schema version 4 with hard v3 rejection on engine load.

Trinity and Neo implementation complete and verified. Site + engine integration on dev branch, ready for Ronnie's live testing.

**Reference:** Orchestration logs (morpheus-audit, morpheus-plan, trinity-impl, neo-impl, trinity-verify).


---

## 2026-04-24: 🚨 Team default model is now claude-opus-4.7

Per `.squad/config.json` (`defaultModel: claude-opus-4.7`) and the standing policy at the top of `.squad/decisions.md`:

- **Every agent spawn defaults to `claude-opus-4.7`.**
- **`claude-haiku-4.5` is FORBIDDEN.** Even if your charter says "preferred: claude-haiku-4.5", that line is overridden. No Haiku, ever.
- **`claude-sonnet-4.5`** (latest Sonnet) is allowed only for trivial mechanical work where opus-4.7 would be wasteful.
- This affects what every future spawn looks like — expect opus-4.7 as your model.

- **Windows filenames:** Never use `:` in any filename. For ISO 8601 timestamps, use hyphens: `2026-04-24T23-58-37Z` not `2026-04-24T23:58:37Z`. Commit 8148ba13 renamed 83 files. See `.squad/decisions.md` and `.squad/skills/windows-compatibility/SKILL.md`.

## Scoping: Prompt-page fractional grader scores (2026-04-25)

**Ask:** Ronnie wants the prompt-detail page graphs to show "5/7 grader points passed" instead of binary pass/fail, plus a broader site review post-grader/embed work.

**Outcome:** Decision drop at `.squad/decisions/inbox/morpheus-prompt-graph-grader-scores.md`. Trinity is unblocked — this is **site-only**, no engine plumbing needed. The full v4 grader_results (with Points) already ships through `/api/runs` because `summary.json` carries `[]*EvalReport`.

**Where Trinity starts:**
1. Widen `site/src/app/data/types.ts` → `RunSummary.results: EvalReport[]` (drops a stack of `as EvalReport` casts already scattered through the codebase).
2. Add `pointsPassRate` helper in `lib/evalPass.ts` (1 line on top of existing `evalPointTotals`).
3. Rewrite six binary spots in `prompt-detail-page.tsx` to use point totals — score column, summary cards, both charts, both correlation tables.
4. Lift the in-progress click-gate from `run-detail-page.tsx:287` onto the prompt-page entries table.

**Issues flagged:**
- ⚠️ `graders_passed`/`graders_total` JSON fields are misnamed — engine writes POINTS counts there. `eval-detail-page` works around it, `run-detail-page` trusts the lie. Logged as Neo follow-up (schema v4.1 or v5).
- Legacy `review.overall_score` still used in 4 components (prompt-detail, prompts-page, dashboard, eval-detail) — replace with point-rate.
- Embed refactor (5690a925, 7a3f421a, 3b84c62e) is clean — no Makefile drift.
- `summary.json` carrying full reports is fine today; flag for scale later.

**Time spent:** ~45m. Within the 30–60m budget. No code changes by me — pure scoping pass per the ask.

---

## 2026-04-25: Scoping — Prompt-page fractional grader scores (45m)

**Ask:** Ronnie wants prompt-detail-page graphs to show "5/7 grader points passed" instead of binary pass/fail, plus post-refactor site review.

**Analysis:**
- Data already on wire — `summary.json` carries full `EvalReport` with `grader_results` and `points`
- Site-only fix: type narrowing in `EvalResult` hides the data
- Six UI gaps identified in prompt-detail-page.tsx
- Post-refactor audit: embed cleanup is clean, no Makefile drift
- Flagged engine naming bug: `graders_passed`/`graders_total` actually count POINTS

**Deliverable:** Decision drop at `.squad/decisions/inbox/morpheus-prompt-graph-grader-scores.md`. Trinity is unblocked.

**Time:** 45m (within budget). Pure analysis, no code.

---

## 2026-04-25: Site post-refactor review conclusion

All six prompt-detail gaps implemented by Trinity. Type widening complete. Dashboard, prompts-page, run-detail all updated. Post-implementation Playwright walkthrough confirmed fractional rendering across all pages. 132/132 tests pass. Ready for Ronnie's live testing.

Follow-ups deferred: grader-by-grader reliability table (separate session), Neo's field rename (schema v4.1 vs v5 decision).

---

## 2026-04-25: Audit — tool-load consolidation (45m)

**Ask:** Map the current remote tool loading flow against Ronnie's target spec, identify gaps, propose a unified module shape, and break the work into items for Neo + Tank.

**Deliverable:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md`.

### Learnings — current tool-load architecture

- **Cache split-brain.** Skills cache at `<BaseDir>/.skills-cache/<v>/<owner>/<repo>/` (`fetcher.go:209`); plugins look at `~/.hyoka/cache/default/<owner>/<repo>/` (`installed.go:45`) plus `~/.copilot/installed-plugins/...`. Two trees, neither aware of the other. Same git repo cloned twice if you use both a skill and a plugin from microsoft/skills.
- **Root cause of `.skills-cache/` in cwd:** `cmd/run.go:403` passes `ConfigDir: ""` to `ValidateAndExpand` for the reviewer factory. Empty string flows to `FetchRequest.BaseDir = ""`. `filepath.Join("", ".skills-cache", ...)` returns a relative path, resolved against `os.Getwd()`. Generator path uses the isolated tmp config dir (which gets `RemoveAll`'d after each eval) — also wrong, but invisible because it's tmp. The cache is effectively destroyed every run.
- **Plugins cannot fetch.** `ResolveInstalled` is stat-only. There is no `pluginFetcher`. Users get a "checked: <paths>" error and instructions to run `/plugin install`. This is the asymmetry to fix — skills got `gitFetcher` in the npx-removal refactor, plugins didn't.
- **Version freshness is unimplemented.** `ensureRepoCloned` runs `git fetch --all --tags` on every call regardless of whether the version is pinned. Pinned-vs-latest distinction exists in YAML (`entry.go:33`) but not in the resolver. Default ("unpinned") fetches HEAD with no audit of what SHA was actually used.
- **Post-session verification gate is INTENTIONALLY disabled.** `copilot.go:643-655` has a TODO block citing #347 — the SDK only emits `SessionSkillsLoaded`/`SessionMcpServersLoaded` after the first message round-trip, so the original gate (placed before SendAndWait) deadlocked. The fix is obvious in hindsight: move the gate to AFTER SendAndWait. Today the verifier accumulates statuses and the renderer flips Loaded→Failed, but the engine ignores it — evals run graders against code generated without the tools the prompt required. False-positive evals.
- **Verifier semantics are correct.** `tool_verification.go` already waits for ALL configured kinds (`emitIfReady` lines 94-108) before producing a verdict. The "wait for all tools, then fail" requirement Ronnie stated is satisfied by the existing data — we just have to consume it.
- **Pre-session validation IS sequential** (`validate.go:282`) but doesn't short-circuit per-entry — every entry validates, only `FirstError()` returns early at the report level. So pre-session, the report is complete; the gap is that the user only sees one error in the EvalResult. Fix is `AllErrors()`.
- **`configDir` is overloaded** as both the isolated Copilot config root AND the cache base. Decoupling these (fetcher gets `CacheRoot()`, configDir stays the throwaway) is the cleanup.

### Decisions baked into the proposal

- New `internal/toolload` package owns cache root + Resolve.
- Cache layout: `${CacheRoot}/repos/<owner>/<repo>/<v>/` shared between skills and plugins; `meta/<owner>/<repo>.json` records SHA + last-fetch.
- Fan-out: A (cache root) + D (collect-all-errors) parallel first wave; B (plugin fetch) + C (freshness) after A; E (re-enable post-session gate) after D; F (path dedup) last.
- Six open questions flagged for Ronnie. Defaults proposed for each so he can rubber-stamp or pick.

**Time:** ~45m. Pure scoping — no code changes.


## Learnings (Guardrail Enforcement Bug — 2026-04-23)

**Date:** 2026-04-23

**Task:** Bug investigation — guardrail real-time enforcement uses stale CLI defaults instead of resolved per-config/per-prompt limits.

**Findings:**
- **Bug confirmed:** `CopilotPromptRunner` is constructed once at CLI startup with CLI defaults (`maxTurns`, `maxFiles`, `maxSessionActions`), before any per-config limits are loaded.
- Real-time enforcement (`copilot.go:303-308` for turns, `339-344` for files, `316-320` for actions) uses the runner's stale fields: `e.maxTurns`, `e.maxFiles`, `e.maxSessionActions`.
- The Engine's `resolveLimits()` correctly merges `CLI → config → prompt` and writes the result to the **report**, but never passes it back to the runner for real-time enforcement.
- **Impact:** A config with `max_turns: 100` will kill the session at 25 turns because the runner's `e.maxTurns` was `0` (CLI default), falling back to hardcoded `25` at `copilot.go:224-227`.

**Bug class extension:**
- `maxTurns`: ✅ AFFECTED — CLI default `0` falls back to `25`, ignoring config overrides
- `maxFiles`: ✅ AFFECTED — CLI default `50`, ignores config overrides
- `maxSessionActions`: ⚠️ PARTIALLY AFFECTED — Only if config sets a value **higher** than the CLI default (e.g., config says `250`, CLI default `50` wins)

**Recommended fix (Option A):**
- Add `SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int)` method to `CopilotPromptRunner`
- Call it from `engine_eval.go` after `resolveLimits()`, before `evaluator.Run()`
- Real-time enforcement checks per-eval fields first, falls back to CLI defaults if unset
- Minimal structural change, backward compatible, testable

**Other options rejected:**
- Option B (construct runner per-eval): breaks resource reuse, higher overhead
- Option C (per-eval context): requires interface change, breaks all test stubs

**Smoke test:**
```bash
# Create config with max_turns: 100, run a Python prompt, verify no premature cancellation at turn 25
hyoka run --prompt-id identity-dp-python-default-credential --config test-high-turns --log-level debug
grep "Turn limit reached" verify-turns.log  # Should be empty or show turn 100, not 25
```

**Decision file:** `.squad/decisions/inbox/morpheus-maxturns-enforcement-bug.md`

**Architectural insight:** The per-eval resolution pattern (CLI → config → prompt) is correctly implemented for the **report** (post-hoc guardrail check), but not threaded through to the **runner** (real-time enforcement). This is a classic state-sync bug: two consumers of the same logical limit, one gets the fresh value, one gets the stale value. The fix threads the fresh value through by adding a per-eval state update hook on the runner.

---

## 2026-04-27: Cross-Agent Note — OPTA MaxTurns Fix Shipped

**From:** Scribe (session close)

Your guardrail enforcement bug investigation (2026-04-23) recommended Option A. The full Option A fix has shipped:

- **Neo** implemented `SetLimitsForEval()` with RWMutex (commits d2f6e93b + def6b803)
- **Switch** added comprehensive test coverage (commits 7dda6358 + fe9a93c9)  
- **Oracle** updated documentation (commit 4a8cd9d0)
- **Coordinator** verified via live smoke test (`python-pairwise` config, turns 1-3 without premature cancellation)

All commits merged to origin/ronniegeraghty/dev. No pre-existing test failures introduced. Ready for production.

---

## CROSS-AGENT UPDATE (2026-04-28T00-54-38Z — Scribe: Tool-Load Gate Fix — Option A Shipped)

**Decision shipped:** Morpheus's investigation (morpheus-tool-load-gate-bug.md) approved. Neo implemented Option A. Switch tested (5/5 cases pass, including 22s slow-load proof). Oracle documented.

**Contribution:** Investigation + decision document recommending event-driven gate. Implementation by Neo now live in commits 8fc6d4be and fb5be186.

**Impact:** Eliminates false positives when 45+ skills take >30s to load. Primary gate is now `AssistantTurnStart`, fallback ceiling is 5 minutes (not primary).

---

## CROSS-AGENT UPDATE (2026-04-28T18-23-00Z — Scribe: Per-Reviewer Vote Display Feature Shipped)

**Decision shipped:** Morpheus's data audit (morpheus-grader-vote-display.md) confirmed: per-reviewer votes already in JSON, rendering gap in frontend. No engine work needed.

**Implementation by Trinity:** Component refactor (ExpandablePoint.tsx) + type extension + integration into GraderResultRow. Auto-expand + amber badge on split votes.

**Validation by Switch:** Test suite (31/31 pass). Reconciliation cycle completed. No regressions.

**Contribution:** Investigation + scoped decision document. Identified frontend-only scope, enabled Trinity to ship quickly. Data audit saved engineering time.

**Status:** Feature complete. Commits c155340f (impl), 5a165d63 (tests v1), e347e4d6 (reconcile). Pushed to ronniegeraghty/dev.


---

## Learnings (Rerun Command Bug — Multi-Model + Pairwise Configs)

**Date:** 2026-04-29  
**Task:** Investigate broken rerun commands for multi-model and pairwise-expanded configs; provide options document

**Bug Confirmed:**

Rerun commands displayed in the web UI and written to `report.json` fail for evaluations using:
1. **Multi-model configs** (e.g., `python-pairwise` with `models: [opus, sonnet, codex]`)
2. **Pairwise-expanded configs** (e.g., `--pairwise` flag creating `baseline`, `without-azure`, etc.)

**Root Cause:**

When configs fan out (multi-model or pairwise), the engine synthesizes **virtual config names** that don't exist in YAML files:
- Multi-model: `python-pairwise/claude-opus-4.6` (from base `python-pairwise`)
- Pairwise: `python-pairwise/baseline/gpt-5.3-codex`, `python-pairwise/without-azure/claude-opus-4.6`

These synthetic names are written to `report.config_name` and used in `buildRerunCommand()` (engine_eval.go:790), but `hyoka run --config <synthetic-name>` fails because config resolution happens in `cmd/run.go` by looking up YAML entries.

**Code Locations:**
- Synthetic name generation: `hyoka/internal/eval/engine.go:339` (multi-model), `hyoka/internal/pairwise/pairwise.go:29,35` (pairwise)
- Rerun command builder: `hyoka/internal/eval/engine_eval.go:790`
- Config lookup: `hyoka/cmd/run.go:244` (splits comma-separated names, calls `cfgFile.GetConfigs()`)

**Live Evidence:**
- Tested `python-pairwise/claude-opus-4.6` → **Error: configs not found**
- Tested `python-pairwise/baseline/gpt-5.3-codex` → **Error: configs not found**
- Tested `baseline/claude-opus-4.6` (single-model, non-pairwise) → **✅ Works**

**Options Delivered:**

File: `.squad/decisions/inbox/morpheus-rerun-command-pairwise-options.md`

Three options drafted as requested:
1. **Option A (Most Complete):** Store original CLI invocation, rerun commands replay the exact user command (re-runs all models if multi-model was used).
2. **Option B (Fastest):** Add `--raw-config` escape hatch flag to accept synthetic names.
3. **Option C (User-Expected):** Store base config name + model, rerun commands use `--config <base> --model <specific>`.
   - **C1 (simplified):** Pairwise variants rerun base config + model (no tool ablation preserved).
   - **C2 (extended):** Add `--pairwise-variant` flag to preserve pairwise variant suffixes.

**Recommendation:** **Option C1** — matches user mental model, reuses existing flags, ships fast (~1 hour), pairwise lossy but acceptable for comparative analysis.

**Process Notes:**
- When a "rerun command doesn't work" bug is reported, check:
  1. Does the config name in the report exist in YAML files?
  2. Are there fan-out transformations (multi-model, pairwise, future dimensions)?
  3. Is the rerun command using a synthetic name from post-transformation state?
- Fan-out is conceptually "compile-time" (happens in cmd/run.go before evals start), but synthetic names leak into "runtime" artifacts (reports). The fix is either:
  - Store pre-transformation state (original user command or base config name) for rerun commands.
  - Add a reverse-transformation layer (parse synthetic names back to user-facing flags).


---

## Learnings (Rerun v2 — Tool-Ablation Fidelity Miss)

**Date:** 2026-04-29 (revision)
**Task:** Redraft rerun-command options after Ronnie rejected v1's coverage of pairwise variants.

**What v1 missed:**
v1's three options (A reconstruct CLI, B `--model` flag, C1 base+model — shipped) all addressed multi-model fan-out but **none preserved tool-ablation state**. C1 explicitly accepted that pairwise variants like `/without-azure` get stripped on rerun, treating that as acceptable lossiness. Ronnie's correction: a "rerun" button that silently runs the *baseline* when the user clicked a *variant* is a lying button. Tool-ablation fidelity is a hard requirement, not a nice-to-have.

**Why I missed it:** I framed the problem as "make the synthetic config name resolvable" rather than "make the rerun command reproduce the exact eval the user is staring at." First framing → drop variant suffix is fine. Second framing → variant suffix is the whole point. **Lesson: when a button says "rerun X," the contract is X-identical reproduction, not X-similar.** Ask "what does the user expect this button to do?" before "what's the smallest schema change?"

**v2 deliverable:** `.squad/decisions/inbox/morpheus-rerun-options-v2-tool-ablation.md`

Four options, all with full tool-ablation fidelity:
- **D** — `--without-tool <name>` repeatable flag (replays the same `pairwise.removeTool()` transform)
- **E** — `--exclude-tool <name>` reusing existing `excluded_tools` config field
- **F** — `--pairwise-variant <name>` flag (selects the named variant via `ExpandPairwise`) — **recommended**
- **G** — Inline/sidecar config snapshot (max fidelity, opaque command)

**Recommendation:** F, layered on C1 (not deprecating it). Three orthogonal flags handle the three fan-out dimensions: `--config` (YAML), `--model` (multi-model), `--pairwise-variant` (tool ablation). Smallest surface (~80 LOC), reuses the same `ExpandPairwise` machinery that produced the variant — drift-proof by construction.

**Investigation findings worth keeping:**
- Pairwise variants are synthesized in `cmd/run.go:287-296` via `pairwise.ExpandPairwise` *before* the engine sees them.
- Naming: `{base}/baseline`, `{base}/without-{tool}`, `{base}/without-{mcp}/{tool}` (deep MCP).
- Variant identity is **not** stored as a structured field on `EvalReport` today — only embedded in the synthetic `ConfigName` string and recovered via `parsePairwiseConfigName` (`engine.go:962-981`) for impact computation.
- Promoting variant identity to a real schema field (`PairwiseVariant`, `RemovedTool`) is a small win regardless of which rerun option ships — eliminates string-suffix parsing.

**Process note:** Held the line on "options doc only, no implementation." Ronnie wants a checkbox to land on before any code moves.

---

## Learnings (Trends Opt-In Redesign)

**Date:** 2026-04-29
**Task:** Investigate the trends process in hyoka, determine how it's invoked and how long it takes, then file a GitHub issue to gate it behind an opt-in flag.

**What trends does:**
- Historical cross-run aggregation: scans past reports in `reports/` directory
- Computes time-series performance metrics (pass rates, duration trends, config comparisons)
- Calls `trends.Generate()` (report scanning) and `trends.AnalyzeTrends()` (AI-powered analysis)
- AI analysis spawns a **Copilot SDK session** — expensive LLM call for regression detection and insights
- Writes markdown report to `reports/trends/`

**When/where triggered:**
- Automatically at the end of **every `hyoka run`** invocation (cmd/run.go:537-573)
- Also available as standalone `hyoka trends` subcommand (cmd/trends.go)
- Dashboard has on-demand `/api/trends` endpoint that calls `trends.Generate()` (internal/serve/dashboard.go:170-176)

**Current control mechanism (opt-out):**
- Flag: `--skip-trends` (boolean, default: false)
- Logic: `if !f.skipTrends && !f.dryRun` → run trends
- Users must explicitly pass `--skip-trends` to opt out

**Cost & Impact:**
- `trends.AnalyzeTrends()` spawns a Copilot SDK session with isolated config (analysis.go:43-50)
- Adds measurable time overhead to every eval run
- Most CI/CD workflows and quick iteration cycles don't need trend data

**Dependencies & Degradation:**
- Dashboard (`hyoka serve`) already calls `trends.Generate()` on-demand via `/api/trends` API
- Dashboard gracefully degrades if pre-computed trends are missing (no hard requirement for pre-computed state)
- No other subsystems depend on auto-generated trends data

**Why current design is problematic:**
- Trends process is **opt-out**, not opt-in → most users pay the cost unnecessarily
- Every single eval run (including fast iteration loops) spawns a Copilot session
- Users must remember the `--skip-trends` flag to avoid overhead

**Proposed redesign (issue #638):**
- **Invert default:** trends skip by default (opt-in model)
- **New flag:** `--with-trends` (boolean, default: false)
- **New logic:** `if f.withTrends` → generate trends
- **User workflow:** `hyoka run --with-trends` to get trend analysis
- **Migration:** No breaking change if we keep `--skip-trends` as deprecated alias, but preferred path is `--with-trends` for positive intent

**Architecture rationale:**
- CLI flags should default to opt-out only for essential features (generator, reviewer, guardrails)
- Optional post-analysis (trends, dashboard enhancements) should default to opt-in
- This aligns with "fast iteration loop" use case (most common) vs. "end-of-day summary" use case (trends user)

**Key files involved:**
- `hyoka/cmd/run.go`: lines 21, 41, 96, 537 (flag definition + invocation logic)
- `hyoka/internal/trends/trends.go`: 573 lines (main aggregation logic)
- `hyoka/internal/trends/analysis.go`: AI session spawning
- `hyoka/cmd/trends.go`: standalone trends subcommand (unaffected by this change)
- `hyoka/internal/serve/dashboard.go`: on-demand API endpoint (unaffected)

**Investigation outcome:**
- Filed GitHub issue #638 (squad label for triage)
- Recommended flag: `--with-trends` (positive intent, matches typical opt-in UX)
- No blockers identified; architecture supports graceful degradation
- Recommendation: invert default in `runFlags` struct, flip condition logic, update help text, deprecate `--skip-trends` or keep as alias

---

## Learnings (CI Failure Diagnosis — 2026-04-30)

**Date:** 2026-04-30
**Task:** Diagnose reported periodic GitHub Actions failures; Ronnie receiving notifications.

**Investigation Finding:**
- **Scheduled workflows:** 100% healthy. `Squad Heartbeat (Ralph)` runs every 30 minutes with 60+ consecutive successes, 0 failures.
- **Non-periodic failures:** 21 total failures across 2 workflows, all on **push events** (not scheduled):
  - **CI**: 18 failures (go vet errors in tests)
  - **Site Bundle Freshness**: 3 failures (stale site/dist/ bundle)

**Root Causes Identified:**

1. **CI (go vet failures):**
   - Lock copy violation: `hyoka/internal/toolload/cacheroot.go:110,117` — sync.Once copied to local variables (no-copy violation)
   - Unknown field `Model` in 3 test files (GraderResult struct field removed; tests not updated)
   - Pointer/value type mismatch in 2 test files (passing `&pass` where `pass` expected)
   - **Likely cause:** Recent refactoring removed/renamed struct fields; tests not synced

2. **Site Bundle Freshness:**
   - Developers pushing `site/src/**` changes without rebuilding/committing `site/dist/`
   - Workflow gate detects stale bundle via `git diff --exit-code site/dist/`
   - **Pattern:** 3 failures in recent week, all push-triggered, reproducible by running `cd site && npm run build`

**Key Insight:**
These are **not** periodic/scheduled failures. They are **frequent push-event failures** (dev branch activity) caused by:
- Unfinished refactoring (struct field changes without test updates)
- Developer workflow misalignment (site bundle rebuild not part of commit process)

**Process Note:**
Ronnie may have conflated "notifications arriving frequently due to dev activity" with "periodic scheduled job failures." The scheduled workflows are solid; the issue is **test-code sync and build artifact management** on push events.

**Recommended Handoff:**
- **Neo (engine):** Fix struct field sync in tests + lock handling in cacheroot.go (blocking all CI runs)
- **Tank (build):** Add site bundle rebuild enforcement (pre-commit hook or CI gate with auto-commit)

**Files Requiring Fixes (proposed):**
- `hyoka/internal/toolload/cacheroot.go` (lines 110, 117)
- `hyoka/internal/report/generator_test.go:140`
- `hyoka/internal/serve/dashboard_test.go:46`
- `hyoka/internal/comparison/comparison_test.go:59`
- `hyoka/cmd/compare_test.go:129`

## Learnings

### 2026-04-30: Phantom Grader Point Investigation — Checks-based Graders

**Issue:** When a YAML grader has `checks:`, the parent grader line is numbered like a criterion (e.g., `1. **DefaultAzureCredential Authentication**`), causing LLMs to treat it as a scoreable criterion in addition to the actual checks.

**Root Cause:** `hyoka/internal/criteria/buckets.go:136` formats the parent line as:
```go
fmt.Fprintf(&b, "%d. **%s**\n", i+1, e.Name)
```

This numbered format makes LLMs interpret three items:
1. The parent line (grader name)
2. Check 1
3. Check 2

But only checks 1 and 2 should be scored.

**Evidence:**
- Report: `/home/rgeraghty/projects/hyoka/reports/20260430-041731/.../report.json`
- Grader: `criteria/language/python.yaml` — "DefaultAzureCredential Authentication" with 2 checks
- Output: 3 points (2 correct checks + 1 phantom parent line)
- Phantom label: `"DefaultAzureCredential Authentication\n   Check the following criteria:\n   1. Uses...\n   2. Uses..."`

**Proposed Fix:** Stop numbering the parent line when checks exist. Use a section header format instead (e.g., `### Name` or `**Name:**`) so LLMs don't treat it as a criterion.

**Affected Code:**
- `hyoka/internal/criteria/buckets.go:136` (FormatUnifiedPromptEntries)
- Test: `hyoka/internal/criteria/buckets_test.go` (TestFormatUnifiedPromptEntries_Shapes)

**Recommended Owner:** Neo (criteria/graders pipeline expert).


---

## 2026-04-30: CI Failure Diagnosis & Phantom Grader Investigation

**Session:** 2026-04-30T04:32:44Z  
**Role:** Lead CI Owner + Investigation Agent

### Task 1: CI Failure Diagnosis

Diagnosed periodic GitHub Actions failures reported by Ronnie. Findings:

- **Scheduled workflows:** Healthy (Ralph Heartbeat: 60+ consecutive runs, 0 failures)
- **Failing workflows:** 2 push-triggered (21 failures in 100 runs since 2026-04-29)
  1. **CI (build-and-test):** 18 failures — go vet errors (sync.Once copy in cacheroot.go + GraderResult schema mismatch in 8 test files)
  2. **Site Bundle Freshness:** 3 failures — stale `site/dist/` bundle not rebuilt after `site/src/**` changes

Ownership delegated:
- **Neo:** Fix CI/test sync + sync.Once handling (✅ DONE — commits 99a185ba, e007695e)
- **Tank:** Add site bundle guardrail (✅ DONE — commit 0de4468b, Husky pre-commit hook)

Output: `.squad/decisions.md` (merged) — Root cause analysis and handoff matrix

### Task 2: Phantom Grader Point Investigation

Diagnosed anomaly in multi-point grader output. Bug confirmed:

**Root cause:** `hyoka/internal/criteria/buckets.go:136` formats parent grader line as numbered criterion (`1. **{Name}**`), causing LLMs to treat it as scoreable. Result: N+1 points instead of N.

**Evidence:** reports/20260430-041731/.../python-pairwise/claude-opus-4.6/report.json — DefaultAzureCredential Authentication grader returned 3 points for 2-check entry.

**Fix strategy:** Remove number from parent line — render as `**{Name}**` or `### {Name}` (section header, not criterion).

**Owner:** Neo (queued after CI fixes)

Output: `.squad/decisions.md` (merged) — Bug description and fix strategy

### Status

✅ Both diagnoses landed clean. All handoff decisions recorded in `.squad/decisions.md`.

---

## Session — Grader Redesign Scope (2026-04-30)

**Tasked by:** Ronnie  
**Deliverable:** Comprehensive scope for multi-part grader redesign  
**Status:** ✓ Complete

### Work

Defined four-part redesign with full implementation details, migration impact, and design decisions:

1. **Part 1 — Prompt Grader Semantics:** BREAKING CHANGE — only `checks:` entries are scorable (breaking for YAML files using only `prompt:`). `name` and `prompt` are LLM judge context only.

2. **Part 2 — Execution Order:** Prompt-file eval criteria runs first, then criteria-file graders in YAML file order. Typed graders no longer partition separately; all graders execute in unified order.

3. **Part 3 — Output Format:** Reports group graders by source file (3 indentation levels). New fields: `SourceFile` and `SourceType` on both `graders.GraderResult` and `report.GraderResult`. Engine populates `SourceType = "prompt_file"` or `"criteria_file"` and `SourceFile = absolute path`.

4. **Part 4 — Tool Usage Grader:** New `tool_usage` grader type. Verifies declared MCP servers and skills were actually used during session. One point per rule. Skips rules silently where env doesn't contain the tool. Edge case: if all rules skipped, emit trivial "no_applicable_rules" point.

### Scope Doc

- **File:** `.squad/decisions/inbox/morpheus-grader-redesign-scope.md` (251 lines)
- **Content:** All file changes, code patterns, validation logic, and migration paths
- **Assigned:** Neo (engine/graders), Tank (CLI output/report/site)

### Decision

✅ Ready for implementation immediately. All design patterns clear, no blockers.

### Output

- Orchestration log: `.squad/orchestration-log/2026-04-30T18-29-54Z-morpheus.md`
- Session log: `.squad/log/2026-04-30T18-29-54Z-grader-redesign.md`
- Decisions merged into `.squad/decisions.md`

## Learnings (Prompt Grader Determinism — 2026-04-27)

- **Root cause of point-count flake:** `averageReview` (`reviewer.go:550`) keys reviewer judgments by exact-string `c.Name`, which is whatever the LLM echoed back into its `criterion` JSON field. Two paraphrases of the same logical check → two map buckets → vote split → score variance.
- **Both grader-source flows already converge on the same shape** `(Preamble, []Check)` — YAML via `UnifiedGraderEntry.Checks`, prompt-file via `ParseEvaluationCriteria` → `CriterionEntry.Checks`. The splitter for prompt-file form was already rewritten (per parser.go:156-167 comment) to a flat checks list. No new splitting work needed.
- **`mergeBucketResults` already prefixes non-`combined` bucket criterion names with `[bucket-name] `** for cross-bucket disambiguation in the vote. With stable IDs, the right pattern is `bucket::id` for the vote key and keep the human prefix for display only — separate the structured key from the display label.
- **The dead-code path is genuinely dead:** `PanelReviewer.consolidate` only refers to itself; `buildConsolidationPrompt` is only called by that dead method. Three `TestBuildConsolidationPrompt*` tests in `review_test.go` go with it. Deterministic vote replaced LLM consolidation a while back but the corpse was left in place.
- **Existing retry pattern in `runSingleReview` (lines 509-522) is the right hook for new ID validation** — it already retries on parse failure and validation failure with model-specific re-prompts. Adding "missing/extra ids" to the validation set is a tiny extension, not new infrastructure.
- **Pattern worth capturing as a skill:** "deterministic LLM panel via stable IDs" — any place where multiple LLMs vote on the same items needs server-controlled IDs flowing through the prompt → response → vote chain. Free-text echo is unreliable. Capture this as a reusable pattern.

## Completion Update (Determinism Shipped — 2026-05-01)

Morpheus scoping proposal approved and shipped by Neo. Pipeline completed:

- **Morpheus proposal:** `.squad/decisions/inbox/morpheus-prompt-grader-schema.md` (scoping doc with 11-commit rollout sequence)
- **Neo implementation:** 11 commits spanning foundation (types/parser/builder) → integration (reviewer switch/vote keying) → cleanup (dead consolidation) → verification (determinism regression tests)
- **Verification:** Two-run byte-identical smoke test ✅ — same point counts per grader, no label drift, no bucket splits
- **Merged:** Both decisions consolidated into `.squad/decisions.md`; inbox files deleted
- **Status:** Ready for merge to main

### Team Updates

- **Neo:** Shipped determinism implementation; all commits on ronniegeraghty/dev (99d32205..120d0db8)
- **Switch:** Testing implications documented in determinism regression tests + unit coverage
- **Scribe:** Merged decisions, updated orchestration logs, cross-agent history


## Session — Grader Overhaul Scoping (2026-05-01)

**Tasked by:** Ronnie  
**Deliverable:** `.squad/decisions/inbox/morpheus-grader-overhaul-plan.md` (15-commit phased plan covering Asks 1–6)  
**Status:** ✓ Scoped

### Learnings

- **Determinism fix shipped 2026-05-01 is incomplete.** Two surviving
  bugs cause X/Y drift even after stable IDs landed:
  1. `averageReview` (`reviewer.go:683`) seeds `criteriaOrder` from the
     UNION of observed reviewer votes, not from `expected []ReviewCheck`.
     Any check no reviewer voted on disappears from `MaxScore`.
  2. `runSingleReview` drops a whole reviewer after 2 validation failures
     (`reviewer.go:646`); `ReviewPanelBuckets` silently drops a
     (model, bucket) result on session error (`buckets.go:140`). Either
     path mutates panel size between runs → strict any-fail flips.
- The "byte-identical smoke test" Neo ran proved the simple-case ID flow,
  but never exercised the reviewer-drop or bucket-failure paths. **A
  determinism smoke must include a fault-injection variant** (mock
  reviewer that fails IDs N times) before any future "shipped" claim.
- **Ronnie's "uniform grader report shape" already exists** as
  `graders.GraderResult{ Kind, Name, Weight, Gate, Score, Pass, Message,
  Points[], Extras, SourceFile, SourceType }` — the work for Ask 2 is
  consolidating *kinds*, not introducing a new struct.
- **Three workspace-output graders** (`file`, `output_check`, `program`)
  and **four tool-perspective graders** (`behavior`, `tool_constraint`,
  `tool_usage`, `action_sequence`) overlap in scope. `output_check` and
  a new consolidated `tool` grader can absorb most of them; keep
  `program` and `action_sequence` standalone.
- **`gate` is a phantom field** since Phase 2 removed gate
  short-circuiting (`grader.go:347-353`). Should deprecate explicitly.
- **YAML `checks:` field already exists for prompt graders** — typed
  graders DON'T need it because their config IS the check spec
  (e.g. `output_check.require_files` produces one check per path).
  Renaming is mainly a *result-shape* concern, not a schema change.

### Output

- Plan: `.squad/decisions/inbox/morpheus-grader-overhaul-plan.md`
  (15 commits, owners assigned, sequencing graph, verification matrix)

## 2026-05-01 — Determinism Overhaul: Real Bug + Legitimate Non-Determinism (Grader Overhaul Part 5)

**Commits:** 40307c40 (docs skill), 64f653d2 (progress), 7e110b02, 1f3c9ec9 (review retry logic)

**Finding:** Original "26 vs 25 checks" inconsistency was *partially* a real determinism bug (now fixed) and *partially* legitimate non-determinism. LLM action_sequence and behavior graders judge **non-deterministic LLM action logs**—two runs of the same prompt produce different actions → different check outcomes. Updated deterministic-llm-panel skill docs to clarify: consensus panel ensures stable reviews but cannot eliminate variance in *what* the LLM does. Panel now retries 3 times on missing IDs before synthesizing (commit 1f3c9ec9).


---

## CROSS-AGENT UPDATE (2026-05-01T23:15:27Z — Team Handoff)

**Planning Shipped:** Morpheus scoping work (morpheus-grader-pairwise-redesign-plan.md) approved and landed in `.squad/decisions.md`. Implementation now underway:

- **Neo:** Pairwise deep fix (4f293e06) + tool grader redesign (24de2f26) — both sections A-B complete
- **Tank:** Workspace grader (1f461a50) + activity grader (0896ba53) — sections C-D complete
- **Switch:** Testing/verification — 3 pairwise eval runs confirm all redesigned graders working. Type registration fix included (ec3c9057).

All commits shipped to ronniegeraghty/dev. Ready for merge to main.


## 2026-05-02 — SDK Tool Naming Collision Clarification (CROSS-AGENT UPDATE)

**From:** Switch 🧪  
**Topic:** Copilot SDK tool naming collision enforcement rules  
**Relevance:** Multi-source tool loading design for Grader  

### Key Finding

The Copilot SDK **does not enforce tool name uniqueness** across:
- Skills ↔ built-ins
- MCP servers ↔ built-ins
- Skills ↔ MCP servers
- Skill ↔ skill
- MCP server ↔ MCP server

SDK only validates collisions for custom SDK tools (via `OverridesBuiltInTool` flag).

### Implication for Your Grader Design

When scoping how Grader loads tools from multiple sources (skills + MCP servers), you **must assume no SDK-side safety**. Implement one of:
1. **Namespacing** — prefix by source (`mcp:server/tool`, `skill:name/tool`)
2. **Load-time validation** — fail early if collisions detected
3. **Scoped registry** — maintain separate tool namespaces per source

**No published resolution rules** from the SDK on collision handling.

### Documentation

- Switch's research: `.squad/orchestration-log/2026-05-02T02-59-28Z-switch.md`
- Session log: `.squad/log/2026-05-02T02-59-28Z-sdk-tool-naming.md`
- Recorded decision: `.squad/decisions/decisions.md` (D-2026-05-02-001)

### Your Next Steps

Incorporate explicit namespacing rules into Grader multi-source tool loading specification. Recommend scoped registry pattern (matches GitHub CLI precedent).


---

## 2026-05-02: Tool Disambiguation Scoping (tool_used Grader)

**Task:** Scoping document for tool name collision disambiguation in `tool_used` grader checks.

**Context:**
- Current `tool_used` matches by bare name only (`ev.Tool` string) — no source/namespace
- Commit `adb46786` fixed skill name matching; `40058f7c` filters redundant skill wrapper events
- Switch's SDK research (D-2026-05-02-001): SDK does NOT prevent cross-source collisions

**Problem:** When multiple tools share the same name (e.g., `azure-mcp/list-resources` + `aws-mcp/list-resources`), the grader cannot distinguish them. MCP × MCP collision is highest-risk; skill × skill and skill × MCP are edge cases.

**Proposal: Option A (Optional Source + Server Fields)**
- Schema: Add optional `source: skill|mcp|builtin` and `server: <name>` to `ToolCheckRule`
- Bare `tool: foo` continues to match any source (legacy behavior)
- Explicit `tool: list-resources, source: mcp, server: azure-mcp` disambiguates collisions
- Load-time warning if collision exists but prompt uses bare name
- Backward compatible (no breaking changes)

**Options Evaluated:**
- **A (Recommended):** Optional fields, graceful degradation, pairwise-friendly
- **B (Rejected):** Fully-qualified syntax (`mcp:server/tool`) — ripple to runner + all graders + site display
- **C (Rejected):** Load-time uniqueness validation — blocks legitimate pairwise MCP comparison scenarios

**Open Questions for User:**
1. Derive source at match-time via `EnvironmentTools` lookup (no `ActionEvent` changes) vs. add `Source`/`Server` to `ActionEvent`? → Recommend: lookup, no runner changes
2. When to log collision warnings? → Recommend: once at config load
3. Should `source` accept a list? → Recommend: no, single-source is clearer

**Scope Estimate:**
- Files: `types.go`, `tool_grader.go`, `tool_grader_test.go`, validation
- No breaking changes
- ~5 hours team time (2-3 impl + 1 test + 1 docs)

**Deliverable:** `.squad/decisions/inbox/morpheus-tool-used-disambiguation.md` — awaiting user sign-off before issue filed for Neo.

**Key Learnings:**
- MCP × MCP collision is the real risk (multiple Azure MCP servers, org-internal + public MCP)
- Graceful degradation (optional fields) wins over rigid enforcement (fail-fast) for pairwise scenarios
- Namespacing at log-time (Option B) would ripple to runner + all graders — avoid unless unavoidable
- Load-time warnings should be once-per-eval (not per-check) to avoid log spam

---

## CROSS-AGENT UPDATE (2026-05-02T04:20:51Z — Session: Tool-Used Disambig + Docs Audit)

**Agents Involved:** Morpheus (scoping), Neo (implementation), Switch (testing), Oracle (docs)

**Morpheus's Work Impact:**
- Scoping document evaluated 3 design options for tool disambiguation:
  - Option A: Optional source + server fields (SELECTED ✅)
  - Option B: Fully-qualified names (rejected due to ripple effects)
  - Option C: Load-time validation (rejected due to blocking pairwise scenarios)
- **Status:** Option A was implemented by Neo and shipped successfully
- Scoping document now recorded in decisions.md as reference material (superseded by implementation)

**Deferred Work (Not Pursued This Session):**
- Skill path disambiguation via `session.skills_loaded` (requires SDK integration)
- Load-time collision warnings (future enhancement)
- Automatic source inference (future)

**Cross-team Impact:**
- Option A validated as correct design choice — simple surface (source + mcp_server), high precision
- Upstream scoping saved Neo implementation time (clear spec, validated options)

**Status:** Scoping complete. Implementation shipped by Neo. Option A proved correct choice.

---

---

## 2026-05-02: Config-Aware Grader When Clauses — Scoping

**Task:** Scope `when:` matching against eval-config attributes (generator model, MCP servers, skills, individual MCP tools), so graders can self-gate to configs that actually expose the tools they check.

**Deliverable:** `.squad/decisions/inbox/morpheus-config-aware-when.md`

### Learnings

- **`props` flow is single-source.** Built once in `internal/eval/engine_eval.go:38-50`, threaded through `e.matchedForEval(props)` → `criteria.MatchingUnifiedEntries` (`internal/criteria/buckets.go`), which handles 3-layer file/group/grader `when:` merge via `mergeUnifiedWhen` + `matchesUnifiedWhen` (`internal/criteria/config.go:240`). The `WhenMap.Matches` in `internal/criteria/graders/types.go:64` is the third match site (used by `ApplicableGraders`, only referenced in tests today). All three use identical AND, case-insensitive, empty-matches-all semantics. **Adding keys to `props` requires zero match-site changes.**
- **`internal/config/tool_filter.go:matchesWhen` is a fourth `matchesWhen` site** — but it filters tool *entries themselves* by prompt props, so it's the layer that *defines* what tools are loaded. Asymmetric: do not feed tool-derived props back into it (circular).
- **`EnvironmentTools` already exists.** `graders.GraderInput.EnvironmentTools` is populated in `engine_eval.go:597-612` with `{Name, Kind, Repo, Path}` per generator tool entry. Graders consume it for tool_used matching. The same loop is the natural site to also emit prefixed keys into `props`. No new data plumbing needed.
- **MCP tool wildcard limitation.** `mcp_tools: ["*"]` is the dominant pattern in real configs (azure-mcp-opus.yaml). `internal/pairwise/pairwise.go:347` documents the limitation: "Until the SDK exposes tool discovery, leave the wildcard in place." Individual-tool gating (`mcp_tool:server/tool`) cannot enumerate wildcard servers without SDK work — must ship Phase 2 with this caveat or wait for SDK discovery API.
- **`tool_used` grader fields (`tool`, `source`, `mcp_server`) map cleanly to prefixed `when:` keys** — `mcp_server:foo`, `skill:bar`, `mcp_tool:server/tool`. This gives users one mental model on both sides.
- **Prefixed-string-keys beat structured-list-fields.** Keeping `WhenMap` as `map[string]string` means zero schema migration, zero parser/validator changes, and natural AND-composition with existing prompt keys. The cost is a YAML quoting requirement (`"mcp_server:azure": "true"`) which is well-understood.

### Pattern Worth Remembering

When adding new dimensions to a `map[string]string`-based filter, **prefer prefixed string keys over structured fields** when (1) the existing match semantics (AND/case-insensitive/empty-matches-all) are correct, (2) values are membership flags (`"true"`), and (3) the schema is loaded by multiple call sites. Cost of widening props is O(1); cost of widening the type is O(call sites × tests).

## Learnings

### 2026-05-02 — Flat string-key maps don't scale to multi-criterion `when:` matchers

When the property surface for a gating clause is small and homogeneous (just prompt scalars: language/service/plane), `map[string]string` with flat keys is fine. The moment you start namespacing keys by prefix to encode structured data (`skill:<name>` = `"true"`, `mcp_server:<name>` = `"true"`), you've already lost the design. Symptoms:
- Forces YAML quoting (`"skill:foo": "true"`).
- The "value" carries no information (`"true"` is a flag, not a value).
- Doesn't compose — multiple lookups become a sequence of independent map keys instead of one filter list.
- The map keys mirror neither the data shape (a list of tool-identity tuples) nor the consumer shape (`tool_used` checks already use `{tool, source, mcp_server}`).

The right shape is a struct with named scalar fields plus typed lists for structured filters. Ship the structured form before more files adopt the flat form — migration cost grows linearly with consumer count.

### 2026-05-02 — Ronnie prefers tool-identity-aligned syntax across `when:` and `tool_used`

Ronnie's pushback on the Phase 1 form was specifically that `when:` should mirror the structure that `tool_used` checks already use: `{name|tool, source, mcp_server}`. Consistency between gating and checking is a hard preference. When designing any future filter/matcher pair, **align the field names with the canonical identity tuple** (in this codebase: `internal/criteria/graders/tool_grader.go` is the source of truth for tool identity). Don't invent parallel vocabularies for the same concept across gate vs. check surfaces.

### 2026-05-02 — Scalar-or-list polymorphism: "OR within field, AND across fields"

Ronnie flagged a gap in the §1 `WhenClause` redesign: every scalar field was pinned as a single string, but real criteria need OR-across-values for the same key (e.g., one grader applying to both Python and Java prompts). His framing nails the mental model: **OR within a field, AND across fields**. That's the same shape every CI matrix and k8s selector uses, and it's worth memorizing as the default for any future gating clause we design.

The Go implementation idiom is a custom `StringOrSlice []string` type with a polymorphic `UnmarshalYAML(node *yaml.Node)` that switches on `node.Kind` (`yaml.ScalarNode` vs `yaml.SequenceNode`) and normalizes both to `[]string`. Both YAML authoring forms — flow `[a, b]` and block `- a\n- b` — decode identically because that's a YAML primitive, not a design choice. Always serialize back as a list (don't try to round-trip single-element slices to scalars) — a single shape downstream beats minor YAML cosmetics.

Bonus side effect: making every `WhenClause` field slice-typed collapses the merge rule to one uniform "child non-nil replaces parent" line per field, eliminating the old scalar-vs-list two-rule split. And authors gain a way to *clear* an inherited constraint by writing `field: []` at the child level — a capability the scalar `!= ""` test couldn't express. Pattern: when you're already restructuring a type, prefer uniformity over minimum-diff — the merge code, the docs, and the mental model all get simpler at once.

---

## 2026-05-02 — Phase 2 `when:` Schema Design + Cross-Eval Visualization Scoping

**Session:** 2026-05-02T07:59:35Z  
**Contribution:** Strategic design for two major features:
1. **Phase 2 Config-Aware `when:` Schema** — Designed structured `WhenClause` with `StringOrSlice` type, `ToolFilter` array, `MatchContext` bundling. Hard cut on Phase 1 prefixed-key form. Spec adopted by Neo for implementation; shipped commit 9da48f32.
2. **Cross-Eval Visualization (4 Views)** — Scoped summary band, per-config rollup strip, evals×checks matrix (collapsed-by-default), per-grader-type stacked bars. Spec adopted by Trinity for React implementation; shipped commits b644bdea + 81e797e1.

**Outcome:** Both specs delivered; implementations verified end-to-end. No product code changes.

---

## CROSS-AGENT NOTE (2026-05-05 — PR #640: Action Semantics Clarification)

**From:** Scribe (via Ronnie's directive in PR #640 spawn manifest)  
**Impact:** Architectural: action counting and session limits

Morpheus should note Ronnie's clarification on what constitutes an "action": "anything the agent does is meant to be an action from reasoning to tool calls to bash commands, to responses." This means `assistant.reasoning` events are **counted as actions** alongside tool calls and responses. This principle is now baked into the eval engine (commit 703f638b via PR #640 port), which ensures session action limits respect both per-eval overrides AND the semantic that reasoning is a first-class action type.

**Architectural implication:** If designing session limits, LLM action quotas, or reasoning time budgets in future, keep this semantic in mind: reasoning is not free.

**No immediate action needed.** Context for architectural decisions involving action budgeting.

## 2026-05-02 — Inline Graders Proposal (investigation)

**Task:** INVESTIGATE & PROPOSE adding `graders:` to prompt files (both `.prompt.md` frontmatter and `.prompt.yaml`), removing `evaluation_criteria:` from YAML.

**Deliverable:** `.squad/decisions/inbox/morpheus-inline-graders-proposal.md`

### Learnings

**Key file paths:**
- `hyoka/internal/prompt/parser.go:33-34, 134-153` — YAML prompt parsing; `frontmatter` struct holds `prompt_text` and `evaluation_criteria` fields
- `hyoka/internal/prompt/parser.go:117-120` — markdown `## Evaluation Criteria` extraction → `Prompt.EvaluationCriteria` + `ParsedCriteria`
- `hyoka/internal/prompt/types.go:39-43` — `EvaluationCriteria string` and `ParsedCriteria []CriterionEntry` are the only criteria-shaped fields on Prompt today
- `hyoka/internal/criteria/config.go:27-64` — `UnifiedGraderEntry` is the shared schema; flat `type` discriminator, `Type/Name/Weight/When/Isolate/Prompt/Checks`
- `hyoka/internal/criteria/config.go:126-228` — `validateEntry` + `validateConfig` (file-level name uniqueness)
- `hyoka/internal/criteria/buckets.go:231-257` — `MergeUnifiedCriteria` / `MergeUnifiedCriteriaToChecks` already merge prompt-type entries with free-text criteria; reusable
- `hyoka/internal/eval/engine_eval.go:664-701` — exec-list builder; implicit "Criteria from prompt file" bucket at position 0, matched criteria-file graders after
- `hyoka/internal/eval/engine.go:260-330` — `matchedForEval`, `reviewBuckets`, `mergedCriteria`; the three call sites that need inline-grader awareness

**Architectural insights:**
- `internal/criteria` does NOT import `internal/prompt` today — adding the reverse dependency is safe (no cycle).
- yaml.v3 is used WITHOUT `KnownFields(true)` for prompt parsing (parser.go:106) — must keep it that way for back-compat. The "deprecated `evaluation_criteria:` → loud error" check needs to be an explicit pre-decode key probe, NOT strict-mode.
- 91 / 91 production `.prompt.md` files use `## Evaluation Criteria`; 0 / 0 production `.prompt.yaml` files use `evaluation_criteria:`. The YAML breaking change costs us exactly one example file (`examples/prompts/example.prompt.yaml`).
- The implicit "Criteria from prompt file" bucket name has been stable across reports for the entire grader-redesign rewrite — renaming would invalidate every on-disk report. Recommended to make it a reserved name instead.
- `criteria.UnifiedGraderConfig` has `groups:` for hierarchical when/isolate, but on a prompt file the prompt itself IS the scope, so groups add no value inline → recommended to reject.

**Proposal recommendation:** Allow inline `graders:` with the unmodified `UnifiedGraderEntry` schema. Markdown's `## Evaluation Criteria` stays forever. Ship in two PRs (additive, then breaking). Hard-error on name collisions across all sources. Allow `when:` on inline graders with a warn-on-redundant validator hint.

## CROSS-AGENT UPDATE (2026-05-06T01-02-53Z — Scribe: Inline Graders Ship Complete)

**From:** Scribe (orchestration capture)  
**Status:** Proposal accepted, implementation shipped, decisions merged

Morpheus's inline-graders proposal was accepted by Ronnie with **one override**: `when:` clauses are **FORBIDDEN** on inline graders (hard error, not soft warning).

Neo implemented and shipped across 5 commits (b290b848 → 2c76fab3) on `ronniegeraghty/dev`. All test passing; build green.

**Decisions merged:** `.squad/decisions.md` now documents both Morpheus's proposal acceptance and Neo's implementation. Inbox files deleted per Scribe protocol.

**Proposal reference:** Morpheus's full 26.4 KB architectural spec remains available in Scribe logs for future reference (no longer in inbox).

