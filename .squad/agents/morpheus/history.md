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
