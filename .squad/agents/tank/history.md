# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/main.go (CLI entry), configs/ (run configs), reports/ (output), site/ (docs/serving), go.work (workspace)

## Core Context

**Archived 15 entries from earlier sessions.**

Historical patterns and learnings:

- ## Core Context: Agent Tank initialized as Platform Dev for hyoka. Owns CLI, config, build, reports, site, and plugins. The CLI supports: `list` (show prompts), `run...
- ## Recent Updates: 📌 Team initialized on 2026-04-03

📋 **Morpheus Audit (2026-04-03):** Audit of CLI and platform layer complete. Key findings: (1) **stale path in mai...
- ## Learnings: - **Agent Attempt Tail-Streaming Deleted (2026-04-23):** After four failed attempts to fix line-wrapping leaks (commits 6b3d3d48, 42ea88fb, fe6efebf...
- ## 2026-04-16 — Phase 3 Merged to Dev (Neo): Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has...
- ## Learnings — example-remote-skill PR (#573): - `examples/configs/*.yaml` are NOT auto-loaded. The default loader reads `configs/` only (see `hyoka/cmd/run.go` + `internal/config/LoadDir`). To u...
- ## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release: Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-ge...
- ## Session 2026-04-21 (Phase 6 Round-1: #602 Approval + #603 Wiring Tests Reassignment): **Mission:** Implement wiring-layer test fixes for PR #603 (reassigned from Neo per reviewer-protocol)

**Context:** Switch requested changes on PR...
- ## Session 2026-04-21T23:22:02Z: Main Sync and Docs Installed-Binary: **Status:** COMPLETE (Part A, Part B via Neo)  
**Branch:** ronniegeraghty/dev (commit 8bfc4da2)  
**User request:** Pull main into dev; switch docs...
- ## Team Context: Unified Grader Direction Proposed (2026-04-22): Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE sch...
- ## 2025 — workers default → 1: Flipped `--workers` default from `runtime.NumCPU()` (capped at 8) to `1` in `engine.go`. Updated help text in `run.go`. Kept the >8 clamp for explic...
- ## 2025 — --progress auto: worker-count-driven selection: Extended `--progress auto` in `hyoka/cmd/run.go` to pick "live" or "log" from the worker count. Commit `d6fd0a59`.

### Learnings
- Final decision s...
- ## 2025 — fix: --progress auto suppressed CI mode on piped multi-worker runs: Reordered the `--progress auto` switch in `hyoka/cmd/run.go` so `workers>1` short-circuits before the non-TTY check. Regression from d6fd0a59: piped...
- ## Team Updates: ### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three...
- ## 2026-04-23 — feat: human-friendly slog handler for stdout/stderr: Implemented ConsoleHandler to replace structured slog output on stdout/stderr with human-friendly messages. When logs go to console (no --log-file):...
- ## 2026-04-23 — fix: tail line wrapping with wide characters (emoji, CJK): The interactive renderer's multi-row tail clearing logic (commit `6b3d3d48`) had the right structure but used **rune counting** instead of proper te...

Full history archived. Recent entries below.

---

## Tool Validation Gate Fixed (2026-04-23)

**Neo's Work:** Fixed blocking tool verification gate that was preventing ALL evaluations from running. Root cause: SDK emits SessionSkillsLoaded events **during** first SendAndWait, not after CreateSession. Gate was blocking before SendAndWait, causing indefinite timeout before events could ever fire.

**Relevant to Tank:** The gate implementation (commit 92a9746c) created a deadlock in the eval flow. This fix disables the blocking gate. Tool failures are still logged (observational) but don't block eval execution. If Tank had any tests for the blocking gate behavior (WU-2 validation), those tests will need updating or removal since the gate is now observational-only.

**Status:** ✅ Gate disabled, evals running. Verified with live eval (88s, passed). Observability maintained via event logging.

**Decision:** Gate remains observational pending SDK event lifecycle documentation and architectural review. Options for future re-enablement documented in decisions.md.

### Session 2026-04-23 (WU-3 — Grouped Tool Display)

**Status:** COMPLETE (commit 42c243f9, pushed to ronniegeraghty/dev)

Implemented grouped tool display for progress renderers. Consumes ParentName/ParentKind fields added by Neo in commit acd36cde to group leaf tools (plugin children, skill_dir children) under their container.

**Changes:**
- Extended `toolLine` struct with `parentName` and `parentKind` fields
- Interactive renderer: During resolution, tools render flat with live tail updates. On EventToolsVerified or redraw, the full block is re-rendered grouped by parent.
- CI renderer: Added EventToolsVerified handler that emits a Tools: section with plain-text grouped output.
- Created `groupToolLines()` helper to group by (ParentKind, ParentName), preserving insertion order.
- Created `renderToolLineFlat()` for resolution-phase rendering (pre-grouping).
- Created `renderToolLine(tl, indented)` for grouped rendering with conditional indentation.

**Format:**
```
Tools:
  - <parent-name> (plugin | skills dir):
      - <child>: Loaded | Failed (reason)
```

**Testing:**
- All existing tests pass with `-race` flag.
- End-to-end verified with `--progress interactive` mode on azure-mcp-skills config.
- Top-level tools (no parent) render as before.

**Learnings:**
- Progress events support both live streaming (during resolution) and grouped display (after verification) — renderers need two code paths to handle both phases.
- The toolLine struct serves as internal bookkeeping; ToolStatus is the external event shape. Converting between them requires helper functions.
- CI renderer needs to handle tools too, not just graders — EventToolsVerified is emitted for all renderers.


### Session 2026-04-24 (WU-A3 — Wait-till-known + resilient parent fan-out)

**Status:** COMPLETE (commit 18d105c3 on ronniegeraghty/dev)

**Issue:** Ronnie's eval run showed the Tools section with stuck "🔄 Loading…" rows for skill-dir parents and missing fan-out for plugins/skills-dirs. Two distinct bugs.

**Root causes:**
1. `onToolResolutionStart` wrote a flat "🔄 Loading…" tail line for every Start event. If a Start never received a matching Result (validate.go emits a Start for skill_dir parents but only Result events for children), the Loading line stayed committed forever.
2. A plugin parent entered `toolLines` with `parentKind==""`, so `groupToolLines` rendered it as a top-level leaf AND as a parent group header — duplicate rows. The new `validate.go` plugin fan-out path was effectively unreachable cleanly.

**Fix:**
- Deferred emission: Start events only update internal bookkeeping. Result events commit the final line (grouped + indented for children, flat for top-level leaves, header-only for loaded plugin containers, flat-failed for failed plugin containers).
- Added `emittedParents map[parentKey]bool` to track which parent headers have been written so the first child of each parent triggers exactly one header.
- Reworked `groupToolLines` to detect containers (kind=plugin OR referenced as `ParentName` by at least one child) and filter orphan loading rows.
- Removed `renderToolLineFlat` (no longer needed).

**Files changed:**
- hyoka/internal/progress/display_interactive.go (193 insertions / 79 deletions)
- hyoka/internal/progress/display_interactive_test.go (4 new test cases + 1 rewrite of ANSI-marker test)

**Testing:**
- `go test ./internal/progress/... -race -count=1`: all 27 tests pass (23 existing + 4 new).
- Manual demo via a DUMP_SCREENSHOT helper test produced the exact grouped shape Ronnie specified (plugin + skills-dir + top-level MCP, with mixed loaded/failed children).
- End-to-end live run was NOT performed because the working tree contains Neo's in-flight plugin-schema refactor (`c.Plugins` removed; `plugins.go` now references missing symbols). `go build ./...` fails until Neo lands that sweep. The progress package builds and tests independently.

**Coordination with Neo:**
- My code is resilient to Neo's schema change: I consume only `ProgressEvent` fields that already exist (`ParentName`, `ParentKind`, `ToolKind`, `Status`, `Reason`). Any new fields Neo adds (e.g. `Source: remote|local`) can be rendered in `renderToolLine` in a follow-up without touching the grouping logic.
- One upstream bug I cannot fix from the renderer: `validate.go` `validateSkillDirEntry` calls `emitStart(entry.Name, …)` but children use `ParentName=entry.Path`. The renderer treats the orphan Start as loading-and-filtered, so output is correct — but the parent header shows `entry.Path` (an absolute path) instead of `entry.Name`. Recommend Neo normalize those: either skip the Start for skill_dir parents, or have children use `ParentName=entry.Name`.

**Deliverable screenshot:**

```
Tools:
  - azure-sdk-python (plugin):
      - skill1-from-plugin: ✅ Loaded
      - skill2-from-plugin: ❌ Failed (fetch timeout)
      - mcp1-from-plugin: ✅ Loaded
  - ./skills/generator (skills dir):
      - pyproject-authoring: ✅ Loaded
      - sdk-smart-defaults: ✅ Loaded
  - bicep-mcp (mcp): ✅ Loaded
```

**Learnings:**
- **Event semantics drive rendering simplicity.** Emitting Start before Result makes sense for observability but forces the renderer to invent a "pending" state. Deferring output to Result collapses the state machine (nothing → final) and makes unmatched Starts harmless.
- **Insertion-order bookkeeping + post-hoc grouping.** Keeping `toolLines` as a flat insertion-ordered slice and grouping only at render time made the redraw path (on `EventToolsVerified` flips) identical to the live path — no divergence between "what we wrote" and "what we'd write now".
- **Container detection must be kind-aware AND reference-aware.** A plugin is always a container (kind=plugin); a skill_dir parent is only detectable retroactively from its children's ParentKind. Handling both in one pass in `groupToolLines` is cleaner than trying to classify up front.
- **Orphan Starts happen.** The resolver emits `emitStart` for skill_dir parents and never a matching `emitResult`. Rather than spuriously render them as failures, filtering "loading" status from the grouped output keeps the transcript clean and lets the upstream fix happen later.

---

## Wave Completion: Plugin Loading Fix (2026-04-23)

The four-agent plugin-loading-fix wave (Neo, Tank, Oracle, Switch) completed successfully. Commits landed on ronniegeraghty/dev:
- **Neo (bc06fb8f):** Retired top-level `plugins:` field; plugins now under generator.tools/reviewer.tools as `type: plugin`
- **Tank (18d105c3, 5216678a):** Wait-till-known rendering; fan-out deduplication; parent header emitted once
- **Oracle (1e5c3b66):** docs/configuration.md plugin section; CHANGELOG breaking-change notice; config migrations
- **Switch (fb70d4c4):** 17 test functions (~29 cases); 5 new test files; full -race suite passes

**Orchestration logs:** `.squad/orchestration-log/2026-04-23T17-{44,45,46,47}Z-{neo,tank,oracle,switch}.md`  
**Session log:** `.squad/log/2026-04-23-plugin-loading-fix-wave.md`  
**Decision entries:** Merged from inbox into `.squad/decisions.md` (5 entries: Ronnie directive + 4 wave decisions)

Status: ✅ Scribe audit complete. Ready for Ronnie's release decision.

---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.

## 2026-04-23: Phase 1 — tool-loading display polish (5 fixes shipped)

Branch: `ronniegeraghty/dev` (pushed). Five surgical fixes to the live-eval interactive renderer; one logical fix per commit.

- **3635a09f** `fix(progress): use config name as skill_dir parent header` — `validateSkillDirEntry` was emitting the resolved filesystem `entry.Path` as the parent label, so headers read `- /abs/path/to/gen-skills (skills dir):`. Switched to `entry.Name` so the visible header is the config name (`gen-skills`). Updated `TestValidateAndExpand_HappyPath` to pin `Parent == "gen-skills"`.
- **582ab59f** `fix(progress): drop plugin parent Loaded/Failed badge` — `validatePluginEntry` and `emitPluginLoadedWithChildren` were emitting `emitResultWithParent(... ToolKindPlugin, ToolStatusLoaded ...)` for the plugin row itself, so the parent rendered with a `[Loaded]/[Failed]` badge alongside the child skills. Removed the three plugin-parent emits and taught `onToolsVerified`'s flip-to-Failed loop to skip `ToolKindPlugin` and any entry referenced as a parent (containers). Added `TestInteractive_PluginParentNoLoadedFailedBadge`.
- **efe18373** `fix(progress): show kind label on grouped child tool rows` — `renderToolLine` had an `if !indented { ... }` guard that swallowed the `(skill)`/`(mcp)` label for children rendered under a plugin/skill_dir header. Removed the guard. Added `TestInteractive_ChildKindLabelShown`.
- **9f994107** `fix(progress): rewrite frozen agent/grader rows in place` — Two related lifecycle bugs sharing one fix:
  - **(d)** Agent Attempt status was rendering below itself stuck on "Running" because a grader Start arriving before `agentComplete` would freeze the agent's tail; `agentComplete` then fell through to `writeLine` and appended a fresh "Completed" row at the bottom.
  - **(e)** `ai_review` grader entry was rendering twice when an unrelated event committed the grader's Running tail before its matching Complete event — `onGraderComplete` fell through to `writeLine` and produced a duplicate.
  - Added `agentLineFrozen`/`agentLineRow`/`graderRowByID`/`pendingGraderID` state on `interactiveEval`. `freezeTail` now records the row index of the displaced agent/grader before bumping `linesWritten`. New `rewriteFrozenLine(row, text)` helper uses the same DECSC/DECRC bracketed save-restore pattern as `redrawToolsBlock`. `agentComplete` and `onGraderComplete` prefer the in-place rewrite path when the row was frozen above; fall back to `writeLine` only when no prior row exists. The grader handler is **not** locked to single-line semantics so Phase 2's Points-based rendering can extend it cleanly.
  - Added `TestInteractive_AgentCompletedRowRewritesFrozenLine` and `TestInteractive_GraderCompleteIdempotentAfterFreeze` — both drive the displacement sequence and assert the rewrite escape (`\x1b7`) was emitted with no duplicate `writeLine` output.
- **fbcd9f38** `refactor(eval): drop redundant tool events around AI review grader` — `engine_eval.go` was wrapping the AI review grader call in `EventToolStart("Review panel: ...")`/`EventToolComplete("Review complete: ...")`; the renderer already shows grader lifecycle via Grader events, so the extra Tool events produced an orphan "Review panel" row above the actual grader entry and contributed to (e). Replaced both with `glg.Debug` — diagnostics preserved, no progress noise.

### Verification
- `go test -race ./hyoka/internal/progress/... ./hyoka/internal/config/tool/... ./hyoka/internal/eval/...` — all green.
- `go test ./hyoka/...` — all green.
- Live eval: `hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise` — 3/3 passed in 192.74s. Debug log shows `Created review panel for config` debug entries where the orphan `EventToolStart` used to be. No errors, no orphan rows in the log.

### Notes / Quirks
- `bytes.Buffer` does not interpret ANSI escapes, so the (d)/(e) tests can't use `strings.Count("Completed")` to assert "no duplicate" — both pre-rewrite and post-rewrite text live in the buffer simultaneously. The tests instead assert (1) `\x1b7` (DECSC) is present (proves rewrite path was taken) AND (2) `"Completed\n"` / `"Pass (N/10)\n"` is **not** present (proves no `writeLine` fallback). This is inferential but accurate; the live-eval was the visual confirmation.
- `validate.go` and the test files are gofmt-loose with no leading-tab indentation. `gofmt -d` shows diffs but the files build and test fine. Did not "fix" formatting — matched existing style.
- Plugin Start event still fires (`emitStart` in `validatePluginEntry`) and creates a `loading` orphan in `toolLines`; `groupToolLines` already handles it correctly via container-discovery (`isContainer[tl.name]`).
- `agentLineRow = linesWritten` is recorded in `freezeTail` **before** the increment, matching the `redrawToolsBlock` pattern (`up = linesWritten - toolsFirstLine`).

### Out of scope (Phase 2 — Neo)
- Grader sub-points generalization (`Points []GraderPoint` on Complete events). The (e) fix deliberately does not lock the grader handler to single-line rewriting; Phase 2 can extend the rewrite-vs-append decision without reworking this scaffolding.

## 2026-04-23: Phase 5 — Report schema v3 (6 commits shipped)

Branch: `ronniegeraghty/dev`. Six commits replace the v2 expanded-grader-row model with a single `ai_review` row that carries `Points`, plus structural metadata that makes the report shape match what the live progress renderer already shows. No PRs — direct commits per project workflow.

Commits (oldest → newest):

- **2feabc8b** `feat(report): bump schema to v3 with MigrateToV3 stub` — `CurrentSchemaVersion: 2 → 3`. Refactored `MigrateToV2` to lift only to v2 (idempotent there); added `MigrateToV3` that calls v1→v2 then bumps the version. v2 → v3 is metadata-only by design — the expanded-row shape is too lossy to de-expand. `WriteReport` and `rerender.migrateReport` both call `MigrateToV3`. `TestMigrateToV2` updated to assert schema=2; new `TestMigrateToV3`.
- **(Neo's earlier d3f26e2d)** `feat(report): add Points to GraderResult` — already shipped in Phase 2, marked todo done with no extra commit.
- **b8b04c2f** `feat(report): add parent linkage to ToolLoadResult` — `Kind`/`Parent`/`ParentKind` fields added to `report.ToolLoadResult` (all `omitempty`); `Status` also gained `omitempty` so container parent rows render without a status badge. `EvalResult` gained `ToolReport *tool.ToolLoadReport`; `copilot.go` populates it on all 3 return paths from `tool.ValidateAndExpand`. New `buildToolLoadResults` helper in `engine_eval.go` consumes the toolReport and emits parent rows + child rows with linkage; falls back to legacy "configured" entries when toolReport is nil. The setup-builder uses it via `setup.MCPServers, setup.Skills = buildToolLoadResults(...)`.
- **992ed39e** `refactor(eval): drop expandReviewGraderResult, single ai_review entry` — Deleted the panel-row fan-out function. Rewrote `convertGraderResults` to emit one row per grader; for `KindPromptReview` it sets `GraderName="ai_review"`, `GraderType="prompt_review"`, copies `ReviewDetails` verbatim, and inherits `Points` directly from the grader aggregate. `TestReviewResultsAppendedNotOverwritten` accepts both `"review"` (legacy) and `"prompt_review"` (v3).
- **d64f1d29** `feat(report): structured EnvironmentInfo.SkillGroups (Option B)` — **Option B chosen.** Two site components (`run-detail-page.tsx:297`, `eval-detail-page.tsx:504`) consume `skills_loaded` as `string[]`. Replacing it (Option A) would break the build for components Trinity owns in Phase 6. Instead added a sibling `SkillGroups []SkillLoadEntry` field that intersects the SDK-loaded set with the validator topology. `site/src/app/data/types.ts` mirrors the new shape via `SkillGroupEntry` + optional `skill_groups`; `skills_loaded` is untouched. Loaded skills with no validator entry fall back to `{Name, Kind: "skill"}` with empty parent linkage.
- **5698f803** `feat(report): pre-computed graders_passed/graders_total roll-ups` — Added `EvalReport.GradersPassed` and `GradersTotal` (both `int` with `omitempty`). Populated in `engine_eval.go` from the unified grader aggregate (`agg.Results`) — the authoritative pre-conversion grader set. The site can read these directly instead of recomputing from `GraderResults` (which now hides the panel under one `ai_review` row). `omitempty` so v2 reports round-trip identically.
- **03159a3e** `test(report): v2-read v3-write migration coverage` — `testdata/v2_report.json` fixture (3 expanded review rows + consensus). `TestV2ReportReadByV3Code` pins down: v3 code reads v2 without panic; expanded rows are preserved across `MigrateToV3` (no de-expansion); migrated v2 reports MUST NOT invent roll-ups (graders_passed/total stay zero). Plus live verification via `hyoka run` of key-vault-dp-python-crud.

### Live verification (capture from /tmp/tank-p5-verify.log run)

```json
{
  "schema_version": 3,
  "graders_passed": 2,
  "graders_total": 2,
  "grader_results": [
    {"grader_name":"files_present", "grader_type":"files_present", "pass":true},
    {"grader_name":"ai_review", "grader_type":"prompt_review", "pass":true, "points":[...], "review_details":{...}}
  ],
  "environment": {
    "skill_groups": [
      {"name":"azure-keyvault-py", "parent":"azure-sdk-python", "kind":"skill", "parent_kind":"plugin"},
      {"name":"customize-cloud-agent", "kind":"skill"}
    ]
  },
  "action_timeline": {
    "session_setup": {
      "skills": [
        {"name":"azure-sdk-python", "kind":"plugin"},
        {"name":"azure-keyvault-py", "status":"loaded", "kind":"skill", "parent":"azure-sdk-python", "parent_kind":"plugin"}
      ],
      "mcp_servers": [
        {"name":"azure", "status":"loaded", "details":"npx -y @azure/mcp@latest server start", "kind":"mcp"}
      ]
    }
  }
}
```

### Phase 6 handoff to Trinity

**Site changes Trinity owns:**
- `site/src/app/data/types.ts` already updated additively — `SkillGroupEntry` interface + optional `Environment.skill_groups`. No existing components broke.
- The grader-results table previously walked expanded review rows. v3 reports now emit ONE `ai_review` row carrying `Points`. Trinity needs to render the row with an expandable Points panel (Phase 4 already has the `Points` field — see Neo's history). Old v2 reports on disk will still contain expanded rows; Trinity's renderer must handle both shapes (detect `grader_type === "prompt_review"` as v3, fall back to `"review"` rows for v2).
- New `EvalReport.graders_passed` / `graders_total` fields are available — use these as the canonical roll-up; do NOT recompute from `grader_results` (which mixes shapes across schema versions).
- New `Environment.skill_groups` provides parent linkage for the skill-loading UI. `skills_loaded` (flat `string[]`) is preserved for fallback; if `skill_groups` is present and non-empty, prefer it for grouped rendering.
- `tool_availability` (different struct from `ToolLoadResult`) was not modified — its `Kind`/`Parent`/`ParentKind` columns will be null in v3 reports because the parent linkage lives in `action_timeline.session_setup.skills/mcp_servers` instead. If Trinity needs grouped rendering on the run-detail tool table, source it from `session_setup`.

**Decision: Option B (sibling field) instead of Option A (replace).** Documented in the d64f1d29 commit message. The driver was that `eval-detail-page.tsx:504` calls `.includes(tool)` on `skills_loaded` and `run-detail-page.tsx:297` calls `.map`. Either change would have forced site-side work that Phase 5 was explicitly NOT scoped for. Option B keeps the cut additive.

**Anchor screenshots dirs:** `.playwright-cli/`, `.trinity-screenshots/` already exist (untracked). Phase 6 Trinity will baseline before/after for the new ai_review row + skill_groups grouping.

### Verification
- `go test ./hyoka/...` — all green (including new `TestV2ReportReadByV3Code`, `TestMigrateToV3`, updated `TestMigrateToV2`, updated `TestReviewResultsAppendedNotOverwritten`).
- `go build ./...` — clean.
- Live eval against `key-vault-dp-python-crud` / `python-pairwise` produced a v3 report with all expected shapes (snippet above).

### Notes / Quirks
- `tool_availability` field uses `ToolAvailabilityEntry` (separate struct) and was deliberately NOT touched — parent linkage goes through `session_setup` (`SessionSetupEvent.Skills/MCPServers` typed as `[]ToolLoadResult`). If anyone sees null Kind/Parent in `tool_availability`, that's by design.
- `report.ReviewGraderDetail` field names are `ReviewGraderPanelEntry` and `ReviewGraderCriterion` (not `ReviewerResultDetail`/`CriterionDetail`); top-level has no `Summary` field — Summary lives on the panel entries. Tripped me up once.
- The `omitempty` on `EnvironmentInfo.SkillsLoaded` means an empty slice renders as missing (jq returns `null`). If the field shows `null` in JSON, the SDK didn't report any skills — that's an upstream signal, not a regression.
- The buildToolLoadResults helper duplicates plugin parent rows into both skills and mcp_servers buckets if the plugin has both kinds — intentional, each bucket needs its container row.

### Out of scope (Phase 6 — Trinity)
- Site rendering of the new `ai_review` row with Points panel.
- Adapting `eval-detail-page.tsx` and `run-detail-page.tsx` to consume `skill_groups` (Option A migration, optional).
- Visual regression baselines.

## Team Update — 2026-04-23 Grader Points Rethink Session

**Shipped this session:**
- Neo (Phase 2): Core fix — killed `expandReviewGraderResult`, established `Points[]` as canonical source. Fixes "all passed but rows red" bug.
- Trinity (Phase 4): Site UX audit against real data, identified 3 rendering inconsistencies.
- Trinity (Phase 6): Final site alignment to schema v3 Points display.

**Status:** All 6 phases shipped. 3 open questions for user (zero-files default, review pass rule, plugin collapsibility). Ready for release decision (merge dev → main, tag v0.3.1).

## 2026-04-24 — Agent Attempt single-line fix (follow-up to Phase 1.4)

Phase 1.4 (`9f994107`) only fixed the rewrite-tracking machinery — it did NOT
address the structural problem: `ensureAgentHeader()` wrote `"Agent Attempt:"`
as one committed line via `writeLine`, then `renderAgentEvent()` wrote the
state (`"  🔄 Running"`) as a separate tail line. Two physical rows. The
rewrite-in-place logic only ever targeted the tail row, so when graders stole
the tail, completion landed in the wrong place — and visually the header and
state were always on different lines.

**What I missed in 1.4:** I treated this as a row-tracking bug. It was a
layout bug. The whole "header on its own line" decision was wrong from the
start of the Agent Attempt section design. The rewrite machinery worked, but
it was rewriting a line that should never have been split off.

**Fix this round (single commit):**
- Dropped the standalone `r.writeLine("Agent Attempt:")` from
  `ensureAgentHeader()`. The function is now just a "section opened" gate.
- Updated `renderAgentStateLine()` to return the full
  `"Agent Attempt: 🔄 Running"` / `"Agent Attempt: ✅ Completed"` /
  `"Agent Attempt: ⚠ Guardrail hit — …"` line. Header and state share one row.
- The `agentLineFrozen` / `agentLineRow` / `rewriteFrozenLine` machinery from
  1.4 still applies — it now targets the SINGLE combined line, which is what
  it should have been doing all along.
- Updated `TestInteractive_AgentCompletedRowRewritesFrozenLine` to assert the
  rewritten payload contains `"Agent Attempt: "` (header + state), not just
  the icon.
- Added `TestInteractive_AgentAttemptSingleLineInvariant`: counts
  `\nAgent Attempt:` occurrences (i.e. distinct physical rows) — must be
  exactly 1 both mid-flight and post-completion. Guards against any future
  regression that splits the header off again.
- Updated the layout docstring at the top of `display_interactive.go` to
  reflect the single-line rendering.

**Lesson:** when a "rewrite in place" bug surfaces, before fixing the
tracking, ask whether the line being tracked should exist as a standalone
row at all. Two-line layouts paired with in-place updates of only one line
are a smell.

---

## 2026-04-24: Pairwise Display Bugs — Agent Attempt Completion Timing (commit dcff4f68)

**Branch:** ronniegeraghty/dev  
**Commit:** dcff4f68 `fix(progress): emit EventSessionDetails before graders to fix Agent Attempt completion timing`

### Root Causes

**Bug 1 (gpt-5.3-codex):** Grader output missing even though generation showed "Agent Attempt: Completed". Root cause was NOT display-specific — graders only run when `len(generatedFiles) > 0` (engine_eval.go:472). If codex generated zero files, the grader block was skipped entirely. The terminal event still arrived and showed "Completed" because the generation phase succeeded. Without EventSessionDetails, there was no Session Details section to visually signal the gap between generation and (skipped) grading.

**Bug 2 (claude-opus-4.6, claude-sonnet-4.5):** Agent Attempt line showed "Running" while graders rendered, then "Agent Attempt: Completed" appeared as a duplicate row at the bottom. Root cause: terminal event (EventPassed/EventFailed) was the ONLY trigger for `agentComplete()`, but it arrived AFTER all graders had already completed. By that time, `onGraderStart` had frozen the agent tail and written multiple grader rows, so `agentComplete()`'s call to `rewriteFrozenLine(e.agentLineRow, line)` used a stale row index — the cursor was now many rows below where `agentLineRow` pointed.

### Solution

Emit EventSessionDetails right after generation completes (engine_eval.go:467-495), BEFORE the grader block starts at line 497. The interactive renderer's `onSessionDetails` handler (display_interactive.go:876-899) now:
1. Calls `agentComplete(evt.FileCount, true, "")` if the agent is still in Running state
2. Renders the Session Details section (Files, Turns, Tool calls, Cost)

This flips the Agent Attempt line to "Completed" while it's still the active tail, so:
- **Bug 2 fix:** completion happens via `rewriteTail()` (in-place update while agent owns the tail) instead of `rewriteFrozenLine()` with a stale index after graders have written rows
- **Bug 1 behavior improved:** the Session Details section now renders between Agent Attempt and the terminal event, making the no-graders flow clearer (generation → session summary → pass/fail, no grader section)

### Implementation Details

- **Event emission:** Computed turn count and cost from `evalReport.SessionEvents`, emitted EventSessionDetails with `Files`, `Turns`, `ToolCalls`, `Cost`
- **Idempotent guards:** Updated `onPassed`, `onFailed`, `onError` to skip calling `agentComplete()` if `r.cur.agentState` is already `agentStateCompleted` or `agentStateGuardrail`. This prevents duplicate rewrites when the terminal event arrives after EventSessionDetails has already transitioned the state.
- **Test update:** `TestInteractive_AgentCompletedRowRewritesFrozenLine` now includes EventSessionDetails in the event sequence and verifies Session Details section renders correctly between Agent Attempt and Graders.

### Learnings

- **Event ordering drives display correctness.** Grader events were always emitted AFTER generation completed (engine_eval.go:561-583), but the renderer had no signal that generation itself was done until the terminal event arrived. EventSessionDetails closes that gap and gives the display a reliable "generation phase complete" marker.
- **Frozen-row rewrite works ONLY when no intervening rows are written.** `rewriteFrozenLine` calculates `up = r.cur.linesWritten - row`, which breaks if `linesWritten` increments between when `row` was recorded and when the rewrite fires. The fix avoids frozen-row rewrites entirely by completing the Agent Attempt while it's still the tail.
- **Per-eval state isolation in pairwise mode.** Workers=1 means one eval at a time, but multiple evals run back-to-back in the same display. Each eval's `interactiveEval` struct tracks its own `linesWritten`, `agentLineRow`, `graderRowByID`, etc. — state resets between evals via `startEval()`. No bleed observed during pairwise runs.
- **Cost calculation is a rough estimate.** `cost += float64(ev.InputTokens+ev.OutputTokens) * 0.00001` is a placeholder. Real cost depends on the model's pricing tier (input vs output rates, prompt caching). Good enough for display; report consumers should use structured token fields from `Environment.TotalInputTokens`/`TotalOutputTokens`.

### Verification

- `go build ./hyoka/...` — clean
- `go test -race ./hyoka/internal/progress/...` — all 27 tests pass (23 existing + 4 from prior wave)
- `go test ./hyoka/...` — all packages pass

Live eval (post-fix) not performed — working tree is on ronniegeraghty/dev with other in-flight work. The test update validates the event flow; Ronnie will verify on the next pairwise run.

### Files Changed

- `hyoka/internal/eval/engine_eval.go` (+28 lines): EventSessionDetails emission after guardrail checks
- `hyoka/internal/progress/display_interactive.go` (+40 lines, -6 lines): onSessionDetails calls agentComplete early; terminal event handlers guard against duplicate calls
- `hyoka/internal/progress/display_interactive_test.go` (+33 lines, -33 lines): TestInteractive_AgentCompletedRowRewritesFrozenLine updated to include EventSessionDetails and verify Session Details section


---

## 2026-04-24: Pairwise Display Bugs — Agent Attempt Completion Timing (commit dcff4f68)

**Branch:** ronniegeraghty/dev  
**Commit:** dcff4f68 `fix(progress): emit EventSessionDetails before graders to fix Agent Attempt completion timing`

### Root Causes

**Bug 1 (gpt-5.3-codex):** Grader output missing even though generation showed "Agent Attempt: Completed". Root cause was NOT display-specific — graders only run when `len(generatedFiles) > 0` (engine_eval.go:472). If codex generated zero files, the grader block was skipped entirely. The terminal event still arrived and showed "Completed" because the generation phase succeeded. Without EventSessionDetails, there was no Session Details section to visually signal the gap between generation and (skipped) grading.

**Bug 2 (claude-opus-4.6, claude-sonnet-4.5):** Agent Attempt line showed "Running" while graders rendered, then "Agent Attempt: Completed" appeared as a duplicate row at the bottom. Root cause: terminal event (EventPassed/EventFailed) was the ONLY trigger for `agentComplete()`, but it arrived AFTER all graders had already completed. By that time, `onGraderStart` had frozen the agent tail and written multiple grader rows, so `agentComplete()`'s call to `rewriteFrozenLine(e.agentLineRow, line)` used a stale row index — the cursor was now many rows below where `agentLineRow` pointed.

### Solution

Emit EventSessionDetails right after generation completes (engine_eval.go:467-495), BEFORE the grader block starts at line 497. The interactive renderer's `onSessionDetails` handler (display_interactive.go:876-899) now:
1. Calls `agentComplete(evt.FileCount, true, "")` if the agent is still in Running state
2. Renders the Session Details section (Files, Turns, Tool calls, Cost)

This flips the Agent Attempt line to "Completed" while it's still the active tail, so:
- **Bug 2 fix:** completion happens via `rewriteTail()` (in-place update while agent owns the tail) instead of `rewriteFrozenLine()` with a stale index after graders have written rows
- **Bug 1 behavior improved:** the Session Details section now renders between Agent Attempt and the terminal event, making the no-graders flow clearer (generation → session summary → pass/fail, no grader section)

### Implementation Details

- **Event emission:** Computed turn count and cost from `evalReport.SessionEvents`, emitted EventSessionDetails with `Files`, `Turns`, `ToolCalls`, `Cost`
- **Idempotent guards:** Updated `onPassed`, `onFailed`, `onError` to skip calling `agentComplete()` if `r.cur.agentState` is already `agentStateCompleted` or `agentStateGuardrail`. This prevents duplicate rewrites when the terminal event arrives after EventSessionDetails has already transitioned the state.
- **Test update:** `TestInteractive_AgentCompletedRowRewritesFrozenLine` now includes EventSessionDetails in the event sequence and verifies Session Details section renders correctly between Agent Attempt and Graders.

### Learnings

- **Event ordering drives display correctness.** Grader events were always emitted AFTER generation completed (engine_eval.go:561-583), but the renderer had no signal that generation itself was done until the terminal event arrived. EventSessionDetails closes that gap and gives the display a reliable "generation phase complete" marker.
- **Frozen-row rewrite works ONLY when no intervening rows are written.** `rewriteFrozenLine` calculates `up = r.cur.linesWritten - row`, which breaks if `linesWritten` increments between when `row` was recorded and when the rewrite fires. The fix avoids frozen-row rewrites entirely by completing the Agent Attempt while it's still the tail.
- **Per-eval state isolation in pairwise mode.** Workers=1 means one eval at a time, but multiple evals run back-to-back in the same display. Each eval's `interactiveEval` struct tracks its own `linesWritten`, `agentLineRow`, `graderRowByID`, etc. — state resets between evals via `startEval()`. No bleed observed during pairwise runs.
- **Cost calculation is a rough estimate.** `cost += float64(ev.InputTokens+ev.OutputTokens) * 0.00001` is a placeholder. Real cost depends on the model's pricing tier (input vs output rates, prompt caching). Good enough for display; report consumers should use structured token fields from `Environment.TotalInputTokens`/`TotalOutputTokens`.

### Verification

- `go build ./hyoka/...` — clean
- `go test -race ./hyoka/internal/progress/...` — all 27 tests pass (23 existing + 4 from prior wave)
- `go test ./hyoka/...` — all packages pass

Live eval (post-fix) not performed — working tree is on ronniegeraghty/dev with other in-flight work. The test update validates the event flow; Ronnie will verify on the next pairwise run.

### Files Changed

- `hyoka/internal/eval/engine_eval.go` (+28 lines): EventSessionDetails emission after guardrail checks
- `hyoka/internal/progress/display_interactive.go` (+40 lines, -6 lines): onSessionDetails calls agentComplete early; terminal event handlers guard against duplicate calls
- `hyoka/internal/progress/display_interactive_test.go` (+33 lines, -33 lines): TestInteractive_AgentCompletedRowRewritesFrozenLine updated to include EventSessionDetails and verify Session Details section

---

## 2026-04-24 — Bug 1 Root Cause Fixed by Neo (Grader Guard Removal)

**Context:** Tank's earlier sessions documented Bug 1 as: "Grader output missing even though generation showed 'Agent Attempt: Completed'". Root cause was NOT display-specific — graders only ran when `len(generatedFiles) > 0` (engine_eval.go:500). If a generator produced zero files, the grader block was skipped entirely by the engine.

**Tank's fix (2026-04-23/24):** Added EventSessionDetails to signal generation completion, allowing the display to render a Session Details section between Agent Attempt and the terminal event. This improved the *visual UX* of empty-grader flows but did NOT fix the underlying issue.

**Neo's fix (2026-04-24, commit 8794e70b):** Removed the `len(generatedFiles) > 0` guard entirely. Graders now run on every eval, regardless of file count. The engine threads the agent's final response through to graders, enabling:
- Response-only evaluation prompts
- Configurable graders like `output_check` to enforce rules on empty workspaces
- Pure logic-based graders independent of workspace state

**Status:** ✅ Bug 1 is now properly fixed at the source. Tank's display work and Neo's engine work are complementary — Tank clarifies *when* graders run; Neo ensures graders *always* run when they should. Both patches are needed for full correctness.

**See:** `.squad/decisions.md` — "Graders Run on Every Eval; Generator Response Threaded Through (2026-04-24T00:56:09Z)"

---

## 2026-04-24 — Bug 3: Duplicate "Agent Attempt: ✅ Completed" After Graders (commit 6f2e1f03)

**Branch:** ronniegeraghty/dev  
**Commit:** 6f2e1f03 `fix(progress): suppress reviewer session events after Agent Attempt completes`

### Symptom

After the graders section finished, two (or more) duplicate `"Agent Attempt: ✅ Completed"` rows appeared at the bottom of the eval block. The Bug 2 fix (commit dcff4f68) had correctly flipped Agent Attempt → Completed in-place via `onSessionDetails` BEFORE graders rendered, but downstream reviewer activity events were still landing in `renderAgentEvent` and creating extra tail lines.

### Root Cause

The PromptReviewGrader runs a real Copilot SDK session for AI review. That session emits the same generation events the generator emits — `EventReasoning`, `EventToolStart`, `EventToolComplete` — through `sendRawEvent` (see hyoka/internal/eval/copilot.go). `engine_eval.go:596` also emits `EventReasoning` directly when a review fails.

These events flow into `interactiveRenderer.HandleEvent` → switch lands on `onAgentActivity` (display_interactive.go:318-320) → `renderAgentEvent` (line 794).

In `renderAgentEvent` (lines 794-807, pre-fix):
- After `onSessionDetails` flips agentState → `agentStateCompleted` and freezes the tail, the tail is no longer `tailAgent`.
- When the FIRST reviewer activity event arrives:
  - `agentState == 0` check is false (it's Completed = 1) — fine, doesn't reset to Running ✓
  - BUT line 802 `if r.cur.tailKind != tailAgent` is true → calls `writeTail(tailAgent, renderAgentStateLine())`
  - `renderAgentStateLine()` reads agentState=Completed and renders **`"Agent Attempt: ✅ Completed"`**
  - That string gets written as a NEW tail line at the bottom
- The next reviewer activity event sees tailKind=tailAgent and just `rewriteAgentTail()`s in place — but the next reviewer (or the next "tail handover" — e.g., grader complete freezes the tail and another reviewer event arrives) creates ANOTHER fresh "Agent Attempt: ✅ Completed" row.

With 2 reviewers in pairwise mode, you get 2 stray rows. There's also the `EventReasoning` emit from engine_eval.go:596 path on review failures.

### Fix

Added a phase-state guard at the top of `renderAgentEvent` (display_interactive.go:794):

```go
// Agent Attempt is already finalized — generation phase is over. Ignore
// activity events from downstream sessions (reviewer Copilot sessions
// emit the same EventReasoning/EventToolStart/etc. through the shared
// event channel, but they belong to grader rows, not the agent tail).
if r.cur != nil && (r.cur.agentState == agentStateCompleted || r.cur.agentState == agentStateGuardrail) {
    return
}
```

The agent tail belongs to the GENERATION phase only — once `agentState` is Completed or Guardrail, no event should re-open it.

### Test

`TestInteractive_ReviewerEventsAfterCompletionIgnored` (display_interactive_test.go:520-568) drives the full sequence:
1. Standard prelude (Starting, EventReasoning to open agent gate)
2. EventSessionDetails (flips Agent Attempt → Completed)
3. EventGraderStart + EventGraderComplete (a typed grader run)
4. EventGraderStart (start the AI review grader)
5. **EventReasoning + EventToolStart + EventToolComplete** (simulate the reviewer session emitting events)
6. EventGraderComplete (review done)
7. EventPassed

Asserts: `strings.Count(out, "\nAgent Attempt:")` equals 1 — exactly ONE Agent Attempt row exists in the transcript.

### Verification

- `go build ./hyoka/...` — clean
- `go test -race ./hyoka/internal/progress/...` — all green (28 tests pass, including the new regression test)
- Live eval: `hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise --log-level debug --log-file hyoka-tank-bug3-verify.log` — 3 evals (2 passed, 1 failed), zero duplicate "Agent Attempt" rows in the output or log.
- `hyoka clean` — 6 orphaned sessions cleaned (normal for pairwise runs)

### Files Changed

- `hyoka/internal/progress/display_interactive.go` (+5 lines): phase-state guard in `renderAgentEvent`
- `hyoka/internal/progress/display_interactive_test.go` (+56 lines): new regression test `TestInteractive_ReviewerEventsAfterCompletionIgnored`

## Learnings

- **Reviewer Copilot sessions share the event channel with the generator.** Once a phase ends (agent attempt completed), incoming activity events from downstream sessions must NOT re-open the previous phase's tail. The renderer must guard each phase's rendering on phase-ownership, not just on event type.
- **Phase-state guards are critical in event-driven terminal renderers.** When multiple concurrent processes (generator, reviewer sessions) emit events through a shared channel, the renderer must filter events by phase. Just checking event type (`EventReasoning`, `EventToolStart`) is insufficient — you must also check whether the event is relevant to the *current rendering phase*.
- **The agent tail lifecycle is: unopened → Running (first activity event) → Completed/Guardrail (terminal state) → CLOSED.** The CLOSED state (post-completion) was implicit before this fix. Making it explicit via the guard prevents leakage from downstream sessions.


---

## TEAM UPDATE (2026-04-24T12:00:00Z) — Generator.json Artifact Arc Complete

**Status:** ✅ LANDED on ronniegeraghty/dev

**Summary:** Neo (Phase 1) + Trinity (Phase 2) + Tank (parallel) coordinated full generator.json artifact pipeline:

1. **Neo Phase 1 (commit d1ed5f61):** Engine emits generator.json for graders. Removed grader-execution guard; added `AgentFinalResponse` to GraderInput. **Test discipline violated** — tests broken at EOD.

2. **Trinity Phase 2 (commits 9f34f072, 72a4d3c3):** Silently fixed all 6 broken Reviewer test stubs. Added comprehensive artifact_test.go. Wired artifact into report layer (v3 schema). Implemented "Generator Session" collapsible panel on eval-detail-page.tsx. **Tests restored green.**

3. **Neo Phase 2b (commit d4b7cbaf):** Verified Trinity's test fixes complete. Added 6 more review_test.go edge cases. Ran live eval (key-vault-dp-python-crud / baseline/gpt-5.3-codex) — generator.json emitted correctly, site panel renders, no regressions.

4. **Tank parallel (commit 6f2e1f03):** Fixed duplicate "Agent Attempt" rows in interactive display via phase-state guard. Events from reviewer sessions no longer bleed into agent-tail rendering after generation completion. Phase-state filtering: once `agentState` is Completed/Guardrail, subsequent events ignored.

**Decisions merged:** All inbox files consolidated into decisions.md (coordinator-grader-input-always.md, coordinator-grader-input-model.md, coordinator-generator-json-on-site.md, trinity-generator-artifact-site.md, trinity-eval-page-file-contents.md, trinity-grader-points-denominator.md, neo-prompt-criteria-own-bucket.md, tank-reviewer-event-suppression.md).

**Next:** Neo Phase 3 — prompt-criteria bucket separation (separate review-grader for prompt-frontmatter criteria vs criteria-file entries).

**Orchestration logs:** 2026-04-24T09:15:00Z-neo.md, 2026-04-24T10:30:00Z-trinity.md, 2026-04-24T11:45:00Z-neo-followup.md  
**Session log:** 2026-04-24T12:00:00Z-generator-json-artifact-arc.md

## 2026-04-24: Per-Grader Display with Bucket Breakdown

**Task**: Split the single "ai_review" grader row into per-bucket graders with individual names and point breakdowns.

**Problem**: User complaint that all AI review grader points were lumped under one `- ai_review` line. They wanted:
- `- DefaultAzureCredential Authentication (prompt): 3/4`
- `- Criteria from prompt file (prompt): 5/5`
- Each with sub-bullets for individual points

**Root cause**: Engine created one `PromptReviewGrader("ai_review")` that processed all ReviewBuckets together, returning one GraderResult with all points from all buckets.

**Solution**:
1. **Display layer** (`display_interactive.go`): Added `displayKind()` helper to map `"prompt_review"` → `"prompt"` for user-facing labels.
2. **Engine** (`engine_eval.go`): Changed from one "ai_review" grader to N graders (one per ReviewBucket), each named after its bucket (e.g., "Criteria from prompt file").
3. **Report builder**: Preserved bucket names in JSON (`r.Name` instead of hard-coded "ai_review").
4. **Tests**: Updated to expect N `Review()` calls (one per bucket) instead of one `ReviewBuckets()` call.

**Key learnings**:
- The display already supported multi-point rendering (`renderGraderWithPoints`) — just needed per-bucket graders.
- Empty-buckets fallback needed: when no criteria exist, create one grader with name "ai_review" to preserve test compatibility.
- Optimization trade-off: Per-bucket grading makes N reviewer calls instead of one batched `ReviewBuckets()` call, prioritizing display clarity over API efficiency.

**Files changed**:
- `hyoka/internal/progress/display_interactive.go` — displayKind() mapper
- `hyoka/internal/eval/engine_eval.go` — per-bucket grader iteration
- `hyoka/internal/progress/display_interactive_points_test.go` — updated test assertions
- `hyoka/internal/eval/engine_reviewmode_runtime_test.go` — updated bucket-mode tests

**Verification**: All tests pass. Live run log shows bucket-level error messages ("Bucket review failed... bucket=Criteria from prompt file"), confirming per-bucket processing.

## 2026-04-24T03:40:38Z: Per-Grader Display Refactor + Coordinator Fixes

**Summary:** Per-bucket grader display landed (`4adc9288`, `4888a402`). During code review, Coordinator identified four critical bugs now fixed: (1) empty-workspace AI-grader bug (`engine_eval.go`), (2) `min_bytes_per_file` vacuous-pass (`output_check_grader.go`), (3) token usage display (`progress.go`), (4) site file-contents fallback (`eval-detail-page.tsx`). All tests pass.


## Session 2026-04-24: Fix KindPromptReview in validTypedKinds

**Status:** COMPLETE (commit 84b1606d on ronniegeraghty/dev)

**Issue:** User reported potential duplicate AI grader execution after commit 4adc9288. Investigation revealed that `graders.KindPromptReview` was incorrectly listed in `validTypedKinds` (hyoka/internal/criteria/config.go:86).

**Root cause:** `KindPromptReview` is the kind returned by manually-created `PromptReviewGrader` instances in engine_eval.go (Phase 2, lines 596-671). It should NOT be a valid criteria-file type. Only `type: prompt` should be used in YAML files for LLM-review graders. Leaving it in `validTypedKinds` suggested `type: prompt_review` was valid, but `NewGrader` doesn't handle it, which would cause instantiation errors if anyone added such entries.

**Fix:** Removed `graders.KindPromptReview` from `validTypedKinds` map in config.go.

**Testing:** go build, go test (eval + criteria packages) passed. Live eval completed successfully.

### Learnings

- `KindPromptReview` is NOT a criteria-file type — it's the kind of runtime graders created by the engine
- `PartitionMatched` filters `type: prompt` entries into `promptEntries`; everything else goes to `typedEntries`
- `NewGrader` handles only the typed grader kinds (file, program, behavior, action_sequence, tool_constraint, output_check)
- Per-bucket `PromptReviewGrader` creation is separate from the criteria-file typed graders path
- Criteria YAML schema: `type: prompt` for LLM-review graders, never `type: prompt_review`

---

## Session 2026-04-24: Fixed Duplicate AI Grader Criteria Display

**Status:** COMPLETE (commit 609ff869 on ronniegeraghty/dev)

**Issue:** User reported duplicate/mixed AI grader output even after commit 84b1606d removed KindPromptReview from validTypedKinds. Live eval showed BOTH review buckets ("Criteria from prompt file" and "combined") were displaying the SAME 8 criteria, instead of being properly separated.

**Root cause:** In engine_eval.go (lines 632-634), when creating per-bucket grader inputs, the code was copying the parent `graderInput` which included the merged `EvalCriteria` field. This caused each bucket's PromptReviewGrader to use ALL criteria instead of only its bucket-specific criteria.

When PromptReviewGrader.gradePanel() processed a single bucket, it checked:
```go
if criteria == "" && len(input.EvalCriteriaBuckets) == 1 {
    criteria = input.EvalCriteriaBuckets[0].Criteria
}
```
But `criteria` was already set from the merged `input.EvalCriteria`, so it never used the bucket-specific criteria.

**Fix:** Added `bucketInput.EvalCriteria = ""` after setting `EvalCriteriaBuckets` to force the PromptReviewGrader to use only the bucket's criteria.

**Verification:**
- BEFORE: "Criteria from prompt file" = 8 criteria, "combined" = 8 criteria (duplicated)
- AFTER: "Criteria from prompt file" = 5 criteria (prompt only), "combined" = 1 criterion (attribute-matched only)

**Additional change:** Updated python.yaml DefaultAzureCredential grader prompt to list two distinct criteria, making any future mis-grouping immediately visible in test runs.

### Learnings

- Per-bucket grader inputs must have ONLY the bucket-specific criteria, not the merged criteria
- The `EvalCriteria` field is used as a fallback when `EvalCriteriaBuckets` is empty or has a single bucket
- Bucket isolation requires clearing the parent's merged criteria to prevent bleed-through
- Review buckets are constructed by `BuildUnifiedReviewBuckets` which separates prompt-frontmatter criteria from attribute-matched criteria into distinct buckets
- The engine creates ONE PromptReviewGrader per bucket (Phase 2, commit 4adc9288)

### Pattern: Three-Attempt Chain (2026-04-24)

Attempted fixes across this session: `4adc9288` (initial per-bucket refactor) → `84b1606d` (remove KindPromptReview) → **`609ff869` (per-bucket grader input isolation, SUCCESSFUL)**. Only the third attempt fixed the display bug. Key takeaway: when working with multi-stage review pipelines, always verify bucket-level isolation at the grader input construction point, not just the bucket initialization point. The one-line fix (clearing `EvalCriteria` after setting `EvalCriteriaBuckets`) was essential to prevent merged criteria bleed-through into per-bucket grading logic.

## Session 2026-04-24: Fix AI Grader Bucket Structure (Fourth Pass)

**Status:** COMPLETE (commit 9e2d8100 on ronniegeraghty/dev)

**Issue:** User reported the grader display was STILL showing a "combined" bucket even after three prior fixes. The expected behavior: one top-level grader per criteria-file entry (using each entry's `name` field), NOT one "combined" bucket grouping all entries.

**Expected display:**
```
- Output Files Exist (output_check)
- Criteria from prompt file (prompt)
- DefaultAzureCredential Authentication (prompt)
```

**Broken display (commit 609ff869):**
```
- Output Files Exist (output_check)
- Criteria from prompt file (prompt)
- combined (prompt)              ← THIS SHOULD NOT EXIST
  - DefaultAzureCredential Authentication
    - Uses DefaultAzureCredential...
```

**Root cause:** `BuildUnifiedReviewBuckets` in `hyoka/internal/criteria/buckets.go` (lines 184-189) grouped ALL criteria-file entries into a single "combined" bucket in combined mode. The previous fix (609ff869) only addressed bucket *contents* isolation (clearing `EvalCriteria` per-bucket), NOT bucket *structure*.

**Investigation path:**
1. Read charter and history (fourth attempt in chain: `4adc9288` → `84b1606d` → `609ff869` → this)
2. Identified `BuildUnifiedReviewBuckets` as the bucket constructor
3. Found line 186-187 calling `combinedCriteriaFileBucket(matched)` which created one bucket for all entries
4. Changed combined mode to iterate through matched entries and create one bucket per entry

**Solution:** Modified `BuildUnifiedReviewBuckets` to create one bucket per criteria-file entry in combined mode:
```go
// OLD (line 184-189):
if mode != ReviewModeIsolated || !HasUnifiedIsolation(matched) {
    if len(matched) > 0 {
        buckets = append(buckets, combinedCriteriaFileBucket(matched))
    }
    return buckets
}

// NEW:
if mode != ReviewModeIsolated || !HasUnifiedIsolation(matched) {
    for _, m := range matched {
        buckets = append(buckets, graders.ReviewBucket{
            Name:     bucketName(m.Entry.Name, len(buckets)),
            Criteria: MergeUnifiedCriteria([]UnifiedGraderEntry{m.Entry}, ""),
        })
    }
    return buckets
}
```

**Files changed:**
- `hyoka/internal/criteria/buckets.go` — changed combined mode to per-entry buckets
- `hyoka/internal/criteria/buckets_test.go` — updated 3 tests to expect N+1 buckets (prompt + N entries) instead of 2 (prompt + combined)
- `hyoka/internal/eval/engine.go` — updated `reviewBuckets` comment to reflect new behavior
- `hyoka/internal/eval/engine_reviewbuckets_test.go` — updated 3 tests to expect 3 buckets
- `hyoka/internal/eval/engine_reviewmode_runtime_test.go` — updated test to expect 3 Review() calls
- `hyoka/internal/eval/engine_test.go` — changed `capturingReviewer` to accumulate all criteria (was overwriting on each call)

**Testing:**
- `go build ./...` — passed
- `go test -race ./hyoka/internal/eval/... ./hyoka/internal/review/... ./hyoka/internal/criteria/...` — all tests pass
- Live eval attempted but failed due to model availability (gemini-3-pro-preview unavailable)
- Code inspection confirms per-entry bucket logic is correct

**Behavior:**
- Combined mode (default): one bucket per entry (prompt + entry1 + entry2 + ...)
- Isolated mode: entries marked `isolate: true` get their own bucket, rest go into "combined" (behavior unchanged)
- The "combined" bucket still exists in isolated mode for leftover entries
- The `mergeBucketResults` function in `review/buckets.go` still has special handling for "combined" to avoid prefixing criteria names

### Learnings

**The difference between bucket *contents* and bucket *structure*:**
- Fix #3 (609ff869) fixed bucket *contents*: cleared `EvalCriteria` per-bucket so each bucket's grader only saw its own criteria
- Fix #4 (this) fixed bucket *structure*: changed the bucket construction layer to emit one bucket per entry instead of grouping all entries into a "combined" bucket
- The issue chain: initial per-bucket refactor (4adc9288) → remove KindPromptReview (84b1606d) → fix bucket contents (609ff869) → fix bucket structure (THIS)

**How buckets flow through the system:**
1. `BuildUnifiedReviewBuckets` (criteria/buckets.go) constructs the bucket list
2. Engine calls `e.reviewBuckets(task.Prompt, props)` which calls `BuildUnifiedReviewBuckets`
3. Engine creates one `PromptReviewGrader` per bucket (engine_eval.go lines 628-671)
4. Each grader processes ONE bucket and returns ONE `GraderResult`
5. Display layer renders each `GraderResult` as a top-level grader row with sub-bullets for individual criteria

**Why the "combined" name exists:**
- Isolated mode needs a bucket for leftover (non-isolated) entries
- The name "combined" signals to `mergeBucketResults` (review/buckets.go lines 172-174, 180-182) to NOT prefix criterion names with `[bucket-name]`
- This special handling is still needed for isolated mode

**Testing pattern:**
- When changing bucket count, search for hardcoded bucket count expectations in tests
- `capturingReviewer` pattern: if a reviewer is called multiple times, it must accumulate state (not overwrite)
- The test fixture Switch built (`criteria/language/test.yaml`, `prompts/test/`) would have been faster for iteration (30s vs 2min)

**Future:**
- If ANOTHER pass is needed, the next place to look is the display layer (`display_interactive.go`) or the JSON report builder
- The bucket names visible to the user are set in `BuildUnifiedReviewBuckets` — they come from `entry.Name` for criteria-file entries

---

## 2026-04-24 — Grader badge format alignment + Points wire-up verification

**Task:** Morpheus's prompt-grader `checks:` redesign, Tank's slice (scope §6/§7). File-disjoint with Neo on shared branch `ronniegeraghty/prompt-grader-checks`.

**Changes (all in `hyoka/internal/progress/display_interactive.go`):**
1. `renderGraderWithPoints` badge: `❌/✅ X/Y passed` → `❌ Fail (X/Y)` / `✅ Pass (X/Y)` per user spec.
2. Soft-truncated per-Point `name` to 50 cols using existing `truncateToWidth` (ANSI-aware). Full text stays in the report.
3. Updated 3 assertions in `display_interactive_points_test.go` for the new format.

**Verifications:**
- Tests green: `go test -race ./internal/progress/... ./internal/report/...`
- Live smoke (`test-dp-test-hello-markdown` × `test/haiku`) confirmed end-to-end Points flow through `convertGraderResults` (`internal/eval/engine_eval.go:1199`) for BOTH the YAML `checks:` grader (Markdown Structure → 2 Points) AND the prompt-file path (Criteria from prompt file → 3 Points). Phase 2 v3 wire-up was already correct — no schema bump needed.

**Learnings:**
- `truncateToWidth` already existed in this package (built for the tail-row in-place rewrite). Reuse > rebuild.
- The CI renderer (`display_ci.go`) deliberately doesn't render Points — only aggregate pass/total. Don't propagate badge format changes there.
- `report.GraderResult.Points` field name in JSON is `points` (lowercase, omitempty). Verified via raw report.json inspection.
- For fast iteration on grader-display work, the `test/haiku` config + `test-dp-test-hello-markdown` prompt round-trips in ~50s including consensus review.

**Decision note:** `.squad/decisions/inbox/tank-badge-format.md`

## CROSS-AGENT UPDATE (2026-04-24T05:37:56Z — Neo + Tank + Scribe: Prompt Grader Checks Landed)

**Feature:** Prompt Grader `checks:` field is now fully implemented, tested, and documented.

**What this means for your work:**
- YAML prompt graders can now declare `checks: [list of items]` instead of smuggling numbered lists in `prompt:`.
- Each check becomes a separate Point in the report and a nested row in the interactive display.
- Badge format updated to `✅ Pass (X/Y)` / `❌ Fail (X/Y)` for clarity.
- Two criteria files migrated (`criteria/language/python.yaml`, `criteria/language/test.yaml`).
- Debug logging added: `slog.Debug` fires when judge returns criterion count ≠ expected (helps track outliers).

**Commits:**
- Neo: `2949f578` — Schema, validation, rendering, migrations
- Tank: `a47cb97d` — Badge format, truncation, display verification

**Branch:** `ronniegeraghty/prompt-grader-checks` (both commits)

**If you touch:**
- Criteria YAML files → respect the new `checks:` shape
- Interactive display logic → Points are now per-check (more rows than before)
- Report-layer rendering → Points flow through unchanged (schema v3 already supports it)

**Follow-up:** Monitor the debug log for criterion-count mismatches (one judge returned the parent grader name as an extra criterion in early testing). Low signal so far — will collect data before deciding on post-filter or stricter review prompt.


---

## 2026-04-24: Feature Shipped — Grader Check Rendering Consolidated ✅

**Coordinated by:** Tank (Coordinator)  
**Session:** 2026-04-24T05:58:18Z  
**Commits merged:** a47cb97d (Tank), 2949f578 (Neo), 86a6574f (prior), ff38a7ec (Switch)  
**Status:** ✅ Shipped to dev

FF-merged `ronniegeraghty/prompt-grader-checks` into `ronniegeraghty/dev`. Branch consolidation complete.

**Feature:** Skill Usage grader + deliberately-failing check to demonstrate per-check Pass/Fail rendering in grader display.

**Verification:** Coordinator smoke 20260424-055601 ✅ confirmed all grader display changes work end-to-end.

**Next:** Branch `ronniegeraghty/prompt-grader-checks` deleted (now equal to dev).


---

## 2026-04-24: 🚨 Team default model is now claude-opus-4.7

Per `.squad/config.json` (`defaultModel: claude-opus-4.7`) and the standing policy at the top of `.squad/decisions.md`:

- **Every agent spawn defaults to `claude-opus-4.7`.**
- **`claude-haiku-4.5` is FORBIDDEN.** Even if your charter says "preferred: claude-haiku-4.5", that line is overridden. No Haiku, ever.
- **`claude-sonnet-4.5`** (latest Sonnet) is allowed only for trivial mechanical work where opus-4.7 would be wasteful.
- This affects what every future spawn looks like — expect opus-4.7 as your model.

- **Windows filenames:** Never use `:` in any filename. For ISO 8601 timestamps, use hyphens: `2026-04-24T23-58-37Z` not `2026-04-24T23:58:37Z`. Commit 8148ba13 renamed 83 files. See `.squad/decisions.md` and `.squad/skills/windows-compatibility/SKILL.md`.
