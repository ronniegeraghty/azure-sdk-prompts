# Oracle — History

## Core Context

**Archived 14 entries from earlier sessions.**

Historical patterns and learnings:

- ## Project Context: - **Project:** hyoka — Go evaluation tool for AI agent outputs, powered by Copilot SDK and multi-model review panels.
- **Stack:** Go 1.26.1+, GitHu...
- ## Learnings: ### Phase 5 README Audit (2026-04-20)

**Status:** COMPLETE  
**Branch:** phase-5  
**Commit:** 9931af2c

Audited README.md for command accuracy and...
- ## 2026-04-16 — Phase 3 Merged to Dev (Neo): Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has...
- ## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release: Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-ge...
- ## 2026-04-20 — #364 frontend test mock fix: **Task:** Fix 20 failing tests from Trinity's rejected #364 (per Switch rejection — API mocking mismatch).

**Problem:** Prompt page tests mock `../...
- ## Session 2026-04-21T23:22:02Z: User Directive — Docs Installed-Binary Command Form: **Status:** NOTED (routing guidance for future)  
**Date:** 2026-04-21

### Directive

User Ronnie requested that all examples in `docs/` use instal...
- ## Session: Fix PR #618 non-blocking nits: **Date:** 2026-04-22 (async follow-up to Morpheus review)
**Outcome:** ✅ All 3 nits addressed & merged

### Work log
- **N1** (cmd/helpers.go:167-16...
- ## PR #618 merged into phase-6: All nits addressed. Work reflects code hygiene principles: total removal (no partial patches), cross-doc consistency, comment/code sync.

**2026-04-...
- ## Session: Rewrite Hierarchical When Example (Phase 6 Polish): **Date:** 2026-04-22 (follow-up to Morpheus PR #607 insight)  
**Outcome:** ✅ Example rewritten, validate green, build green

### Work Summary

**Pr...
- ## Team Context: Unified Grader Direction Proposed (2026-04-22): Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE sch...
- ## Phase 4 Docs: Per-Grader Documentation Scaffolding (2026-04-22): **Status:** COMPLETE  
**Branch:** ronniegeraghty/dev  
**Commit:** 96f3f7e9

Implemented Phase 4 of grader unification docs per Issue #627.

### Ar...
- ## Session: Phase 4 Documentation Finalization: **Date:** 2026-04-22  
**Work:** Finalized `docs/graders/output_check.md` against Tank's shipped v1 API (commit ad2a8ce7)

### Changes Made

1. **Re...
- ## Session: Output UX Sprint Documentation: **Date:** 2026-04-23  
**Work:** Documented the new interactive / CI progress renderers and `--workers` default flip.

### Files touched

- `hyoka/R...
- ## Team Updates: ### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three...

Full history archived. Recent entries below.

---

## Session: Tool Load Validation Documentation (WU-4)

**Date:** 2026-04-24  
**Work:** Documented the new tool load validation behavior shipping with Neo's tool verification fixes.

### Task: WU-4 from Neo's Fix Plan

Neo's investigation identified that tool load failures were silent — if a configured skill or MCP server failed to load, the eval would continue without the tool, leading to misleading pass/fail results. The fix implements a 10-second validation gate that aborts the eval immediately if any configured tool fails to load.

My documentation task (WU-4) was to explain this new behavior to users.

### Files Changed

1. **docs/configuration.md**
   - Added new "Tool Load Validation" subsection under the Tools section (after MCP Servers, before Limits)
   - Explains that all configured tools in `generator.tools` and `reviewer.tools` are implicitly required
   - Documents the 10-second timeout and abort-before-generation behavior
   - Lists common causes: missing SKILL.md, incorrect paths, remote skill unavailable, MCP server not found, SDK timeout
   - Explains how to diagnose with `--log-level debug --log-file` and grep for tool errors
   - Forward-looking note about future `required: false` opt-out (Phase 2, not in this release)

2. **docs/troubleshooting.md** (new file)
   - Comprehensive "Tool Load Failures" section as primary content
   - Explains diagnosis: use debug logging and grep for "tool|skill|mcp|verifier|failed"
   - Dedicated subsections for each common cause with specific fixes:
     * **Skill not found or SKILL.md missing** — check directories, verify paths are relative to config file, test with absolute path
     * **Glob pattern produces no matches** — verify glob syntax and that all matched dirs contain SKILL.md
     * **Remote skill download fails** — test GitHub access (`gh repo view`), check auth (`gh auth status`), manually test `npx skills add`
     * **MCP server fails to start** — verify command in PATH, test manually, ensure `mcp_tools` field is set
     * **SDK timeout (10 seconds)** — verify network, pre-download remote skills, check system resources
   - Quick checklist for users to verify all tools before re-running
   - Additional troubleshooting sections for other issues (timeouts, model not found, session init failures)
   - Guidance on getting help: check logs, check reports via `hyoka serve`, open a GitHub issue

### Style & References

- All docs follow existing Microsoft Style Guide conventions (technical accuracy, clear task-focused language)
- Examples use real command structures from the codebase and CLI
- Diagnostic workflows match the actual log format and tool names from Neo's verifier implementation
- No invented design decisions — all content derives from Neo's investigation doc (.squad/decisions/inbox/neo-tool-skill-investigation-2026-04-23.md)

### Commit

- Commit: `6ca1a341` — "docs(oracle): add tool load validation documentation — WU-4"
- Includes Copilot co-author trailer
- Pushed to `origin/ronniegeraghty/dev` (no PR per instructions)

### Notes

- The actual validation gate implementation (WU-1 + WU-3) is Neo's responsibility
- Test cases (WU-2) will be Switch's responsibility
- My docs are written to match Neo's decision (Option A: validation in copilot.go after session start) and assume the error message format from the verifier: `"required {kind} '{name}' failed to load: {reason}"` where reason is "SDK did not report skill/mcp as loaded"
- If Neo modifies the error message format, these docs can be easily updated in a follow-up

**Status:** ✅ Complete. Tool load validation behavior is now documented for users.

## Tool Validation Gate Fixed (2026-04-23)

**Neo's Work:** Fixed blocking tool verification gate that was preventing ALL evaluations from running. Root cause: SDK emits SessionSkillsLoaded events **during** first SendAndWait, not after CreateSession. Gate was blocking before SendAndWait, causing indefinite timeout before events could ever fire.

**Relevant to Oracle:** Documentation for the tool validation gate feature (commit 2c3835ca, architecture docs) becomes partially obsolete. The gate itself is no longer blocking, though observability (event logging) remains. Oracle should: (1) Update any architecture docs describing the gate as a blocking mechanism, (2) Document the SDK event timing discovery (events fire during SendAndWait, not after CreateSession), (3) Consider writing SDK event lifecycle docs for future reference.

**Status:** ✅ Gate disabled, evals running. Verified with live eval (88s, passed). Observability maintained via event logging.

**Decision:** Gate remains observational pending SDK event lifecycle documentation and architectural review. Options for future re-enablement documented in decisions.md.

## Tool Load Hard-Fail & Grouped Tools Output Documentation (2026-04-24)

**Neo's Implementation (WU-1, WU-2):**
Neo shipped pre-session static tool validation with hard-fail semantics:
1. `tool.ValidateAndExpand()` resolves every declared plugin, skill dir, MCP server **before** session creation
2. Failure aborts eval with `error_category="tool_load_failure"` — generator never invoked
3. Plugins and skill dirs now expand to children in progress output (ParentName/ParentKind fields)
4. Reviewer tool validation moved to per-config closure in cmd/run.go — eliminates cross-config leakage bug

**Documentation Task (WU-5):**
Updated `docs/configuration.md` "Tool Load Validation" section to replace post-hoc SDK timeout docs with user-facing behavior:
- Hard-fail contract: what triggers tool_load_failure (plugin not found, skill path missing, missing SKILL.md, empty skill dir, MCP unavailable)
- Tools progress output example showing grouped expansion (parent with children indented/connected)
- Error reporting format in EvalReport JSON (error, error_category, error_details fields)
- Config scoping explanation (reviewer tools validated independently per config)
- Updated diagnostics workflow (check report JSON + grep logs for tool details)

Updated `CHANGELOG.md` Unreleased section:
- **Added:** tool_load_failure error category + grouped Tools output
- **Changed:** tool load validation now hard-fails before generation
- **Fixed:** reviewer skill resolution cross-config leakage

**Key Learning:** Neo's decision memo (neo-tool-load-hardfail.md) describes the implementation contract in implementation terms (ValidateAndExpand, ToolLoadReport, Role field filtering). User docs need a different framing: focus on observable behavior (hard-fail prevents silent failures, grouped output enables diagnostics, per-config validation prevents leakage). Translate internal abstractions (e.g., role-based filtering) to user-visible consequences.

**Commit:** `557bb83b` — "docs: tool-load hard-fail and grouped Tools output"

**Status:** ✅ Complete. User-facing documentation for tool-load validation behavior now accurately reflects Neo's WU-1/WU-2 implementation.

---

## Work Unit 3: Plugin Field Migration & Schema Documentation

**Date:** 2026-04-24  
**Status:** ✅ Complete  
**User Request:** "Make sure our docs are updated to reflect the new schema" (plugin field retirement)

**Context:** Neo is retiring the top-level `plugins:` field in pre-1.0. All configs migrate to `generator.tools: [{type: plugin, ...}]` in a single commit. No deprecation path.

**Deliverables Completed:**

1. **Config file migration:**
   - `configs/baseline-sonnet-skills.yaml` — migrated `plugins: [...]` entries to `generator.tools` with `type: plugin` and `source: remote`
   - `configs/python-pairwise.yaml` — migrated `azure-sdk-python@skills` to generator.tools entry

2. **Documentation updates (docs/configuration.md):**
   - Added new **Plugins** subsection (after MCP Servers) documenting:
     * Plugin declaration syntax: `type: plugin`, `source: local|remote`
     * Local plugin resolution: `.hyoka/plugins/{name}.yaml` default path
     * Remote plugin behavior: fetched + cached under `.hyoka/cache/plugins/`
     * Dual-role semantics: explicit declaration needed in both generator AND reviewer (no auto-share)
     * Hard-fail semantics: fetch errors fail before session; missing tools fail at validation time
   - Updated Tools Progress Output example to show grouped plugin + skill dir display with children
   - Updated final "Plugins" section to reference the new location in schema docs

3. **CHANGELOG.md:**
   - Added **Breaking Changes** section with:
     * Clear statement: "Retired top-level `plugins:` field — Pre-1.0, no deprecation path"
     * Before/after migration example (OLD `plugins:` vs NEW `generator.tools`)
     * List of affected config files

4. **Validation:**
   - Both migrated configs validated as correct YAML
   - No remaining `plugins:` field references in docs or config examples (only in CHANGELOG and plan)

**Key Learnings:**

- **Schema migration clarity:** Users need explicit examples of the old vs new shape, not just abstract descriptions. Before/after YAML snippets in CHANGELOG are critical for seamless adoption.
- **Dual-role surprise:** The fact that plugins are NOT auto-shared between generator/reviewer environments needs prominence in docs — it's a breaking change from the previous implicit dual-role behavior.
- **Tools output parity:** Updated example output to match Ronnie's requested format (plugin names with child tools indented), consistent with skill dir grouping.

**Commit:** `1e5c3b66` — "docs: migrate plugins: field to generator.tools with type: plugin"

**Files touched:**
- `configs/baseline-sonnet-skills.yaml` — config migration
- `configs/python-pairwise.yaml` — config migration
- `docs/configuration.md` — 70 new lines documenting plugin schema + updates to output examples + final section redirect
- `CHANGELOG.md` — Breaking Changes section with migration guidance

**Status:** ✅ Complete. All docs reflect new plugin schema. Configs migrated. Ronnie can push to ship.

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

---

### 2026-04-23T18:52Z: Cross-agent update — Plugin schema BREAKING CHANGE (Neo, commit `2c1de1c0`)

Neo rewrote `docs/configuration.md` plugin section. The `@skills` magic alias is gone — remote plugin entries now require an explicit `repo:` field (e.g. `repo: github.com/microsoft/skills`). Names containing `@` are rejected at validation. This BREAKING CHANGE reverses commit `769dea69`. Per Ronnie: *"I want to be explicit when configs are written."*

For your work: any docs, examples, or CHANGELOG entries that reference `name@skills` shorthand need to be migrated to the explicit `repo:` form. Worth a CHANGELOG callout for downstream users. See `decisions.md` 2026-04-23T18:50Z entry for the full schema, validator messages, and migration guidance.
