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
