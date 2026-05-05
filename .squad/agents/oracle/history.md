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
---

## 2026-04-24T06:00Z: TEAM DIRECTIVE — Work on `ronniegeraghty/dev`

**By:** Ronnie (User directive captured by Copilot)  
**Status:** Active

Going forward, the team works directly on the `ronniegeraghty/dev` branch with frequent commit points. No more transient feature branches like `ronniegeraghty/prompt-grader-checks` for in-flight squad work — merge to dev and keep moving.

**Rationale:** User request — streamline workflow, reduce branch proliferation, enable continuous integration of squad work.

**Action:** Update your local branch strategy. All future work targets dev with regular commits.


---

## 2026-04-24: 🚨 Team default model is now claude-opus-4.7

Per `.squad/config.json` (`defaultModel: claude-opus-4.7`) and the standing policy at the top of `.squad/decisions.md`:

- **Every agent spawn defaults to `claude-opus-4.7`.**
- **`claude-haiku-4.5` is FORBIDDEN.** Even if your charter says "preferred: claude-haiku-4.5", that line is overridden. No Haiku, ever.
- **`claude-sonnet-4.5`** (latest Sonnet) is allowed only for trivial mechanical work where opus-4.7 would be wasteful.
- This affects what every future spawn looks like — expect opus-4.7 as your model.

---

## Work Unit 5: Example-File Schema Audit (2026-04-24)

**User Request:** Audit every example prompt/config/criteria file against the
current schemas. Report drift; fix safe drift; flag ambiguous cases.

### Scope audited
- 90 production prompts under `prompts/` ✅ all valid (covered by `hyoka validate`)
- 14 production configs under `configs/` ✅ all valid
- 4 production criteria files under `criteria/` ✅ all valid
- 4 example prompts under `examples/prompts/` — 1 ambiguous, 3 ok
- 3 example configs under `examples/configs/` — all ok
- 9 example criteria under `examples/criteria/` — 8 needed explicit `type: prompt`
- 1 init-seed `exampleConfig` const in `hyoka/cmd/init.go` — was BROKEN
- ~10 docs/*.md with embedded schema snippets — 5 had drift

### Fixes applied (10 files)
1. `hyoka/cmd/init.go` — wrapped `exampleConfig` in `configs:` list (was unparseable)
2. `criteria/language/{java,rust}.yaml` — added explicit `type: prompt`
3. `examples/criteria/{language/*,service/*,hierarchical-when-example}.yaml`
   (8 files) — added explicit `type: prompt`
4. `docs/configuration.md` — fixed flat-frontmatter snippet to nested `properties:`
5. `docs/starter-files.md` — fixed 3 frontmatter snippets to nested `properties:`;
   added implementation-status banner for unsupported `starter_files:` field
6. `docs/graders/index.md` — corrected schema (top-level `prompt:`/`checks:`,
   not `details: { prompt: ... }`); added v4 invariant call-out;
   relabeled `prompt_review` as engine-internal
7. `docs/graders/prompt.md` — full rewrite of Configuration + Example sections
   with correct top-level `prompt:`/`checks:` shape and v4 invariant note
8. `docs/graders/prompt_review.md` — rewrote to clarify it is **not** a
   user-configurable YAML `type:` value (validation rejects it)
9. `docs/grader-config-schema.md` — replaced 353-line legacy doc with a
   stub redirect + migration cheat sheet (kind→type, config→details,
   `gate:` removed, `kind: prompt` + `config.rubric` → top-level prompt/checks)
10. `examples/prompts/graders-frontmatter-example.prompt.md` — added warning
    banner that the `graders:` frontmatter field is NOT consumed by the
    parser (proposed-only)

### Flagged for human review (decision memo written)
- **`graders:` frontmatter field** — example documents a feature the parser
  doesn't consume. Decision needed: implement or delete the example.
- **`starter_files:` (list)** — design-draft only; `docs/starter-files.md`
  Option B does not match the parser. Banner added.

### Validation evidence
Before fixes: `✓ All 90 prompt(s) are valid / ✓ All 14 config(s) are valid /
✓ All 4 criteria file(s) valid (32 grader(s))`.
After fixes: identical output. Examples (which `hyoka validate` skips) all
pass the loaders directly via a temporary in-module audit binary
(documented in the new `example-file-validation` skill).

### Note for Neo
The unified loader silently translates "no `type:` + has `prompt:` + no
`details:`" → `type: prompt` at `criteria/bundle.go:84`. This kept eight
example files validating without explicit `type:` and hid the drift. With
all in-tree examples now migrated, consider either logging a deprecation
warning when this translation fires or removing it.

### Skill extracted
`.squad/skills/example-file-validation/SKILL.md` — schema sources of truth,
where drift hides outside `hyoka validate`, the in-module audit-binary
trick for testing example/ files, and the v4 invariants to defend.

### Key Learnings
- **`hyoka validate` covers runtime paths, not example/seed paths.** Three
  major drift sites (the init seed, examples/, docs/* snippets) are
  invisible to the validate command. A docs owner needs to walk these
  manually on every schema change.
- **Hidden translations mask drift.** The legacy-prompt → type-prompt
  coercion in `bundle.go` made 8 stale example files look correct. When
  the team adds compatibility shims, surface them as deprecation warnings
  so doc audits catch the staleness.
- **A brand-new doc set can still ship wrong.** `docs/graders/*.md` was
  written DURING the v4 unification but encoded the prompt-grader shape
  incorrectly (`details: { prompt: ... }` instead of top-level
  `prompt:`/`checks:`). Always validate doc snippets by pasting them
  through the actual loader, not by inspection.

## 2026-04-24: Follow-up note — legacy `name` Point issue was historical only

Trinity's site audit found 744/838 historical Points used the legacy `name` field. We were worried this might mean current engine code still emitted `Name:` instead of `Label:`. **Neo verified on a fresh run that the engine is clean** — every grader on current tip emits ≥1 Point with non-empty `Label`. The b7611606 invariant fix holds.

So: no engine fix was needed. The 88% legacy-name reports came from local `reports/` (git-ignored ephemera) generated before the rename. Trinity's fallback chain in the renderer is the right disposition, since those old artifacts will never be regenerated.

(Bundle.go:84 silent coercion note still stands — that's a separate observation about loader leniency.)

## Session: Feature Cleanup — Unimplemented Examples to Issues

**Date:** 2026-04-24  
**Requested by:** Ronnie Geraghty  
**Status:** ✅ COMPLETE  
**Branch:** ronniegeraghty/dev  
**Commit:** 09cf17f0

### Task

Per Ronnie's directive, remove two example files/doc sections that documented features not yet implemented in the parser:

1. **`examples/prompts/graders-frontmatter-example.prompt.md`** — sketched `graders:` frontmatter field (Issue #636)
2. **`docs/starter-files.md` "Option B"** — sketched `starter_files:` list syntax (Issue #637)

Convert both into GitHub issues for future implementation.

### Work Log

1. **Deleted** `examples/prompts/graders-frontmatter-example.prompt.md` via `git rm -f`
2. **Updated** `examples/README.md`:
   - Removed reference to deleted example
   - Removed "Prompt-Level Graders" section (documented unimplemented feature)
3. **Updated** `docs/starter-files.md`:
   - Removed "Option B: Explicit File List" section (2.2 → 2.2 precedence rules)
   - Removed validation checks for `starter_files`
   - Removed struct field documentation for `StarterFiles`
   - Removed implementation roadmap
   - Kept Option A (directory reference) as the sole recommended approach
4. **Created two GitHub issues:**
   - Issue #636: "Feature: support `graders:` list in prompt frontmatter" — documents the design concept, parser gap, and implementation steps
   - Issue #637: "Feature: support `starter_files:` list (Option B) in prompt frontmatter" — documents the design concept, lightweight use case, and validation rules
5. **Committed** with co-author trailer: `09cf17f0`
6. **Pushed** to `ronniegeraghty/dev`

### Outcomes

- ✅ All unimplemented features removed from shipped examples/docs
- ✅ Features properly tracked in GitHub (Issues #636, #637) with full context
- ✅ Documentation now reflects only implemented functionality
- ✅ Examples directory cleaned; `hyoka validate` will find no false positives
- ✅ Cross-references removed from examples/README.md

### Design Pattern Applied

**Documentation Hygiene Principle:** If a feature is sketched but unimplemented, move it to an issue rather than leaving stale examples/docs that drift from the parser. This ensures:
- Single source of truth: issue tracker, not docs
- Examples always match current parser behavior
- Feature requests clearly scoped with implementation steps

### Learnings

Unimplemented examples are worse than no examples — they confuse users and pollute validation reports. Always defer unshippped features to the issue tracker.


- **Windows filenames:** Never use `:` in any filename. For ISO 8601 timestamps, use hyphens: `2026-04-24T23-58-37Z` not `2026-04-24T23:58:37Z`. Commit 8148ba13 renamed 83 files. See `.squad/decisions.md` and `.squad/skills/windows-compatibility/SKILL.md`.

## Session: Document Version Pinning on Remote Skills/Plugins

**Date:** 2026-04-25
**Task:** Document `version:` field on remote skills/plugins and `tool_version_override:` map (existing code, missing docs).

### Learnings

**Version Pinning Mechanics:**
- **Per-entry `version:`** — defined on every tool entry (skill or plugin) via the `Version` field in `hyoka/internal/config/tool/entry.go:28-33`. Accepts branch name, tag, or commit SHA. Empty/omitted = repo default branch.
- **Fetcher Resolution:** `hyoka/internal/config/tool/fetcher.go:275-313` handles both fresh clones (`git clone --branch <version>`) and cached updates (`git fetch --all --tags && git checkout <version>`), enabling branch/tag/SHA pinning.
- **Top-level `tool_version_override:`** — a map[string]string keyed by tool entry name. Defined in `hyoka/internal/config/config.go:95-130`, with `ApplyVersionOverrides()` walking every Generator/Reviewer tool entry.
- **Resolution Order:** Per-entry `version:` ALWAYS wins over override map, which wins over fetcher default.
- **Multi-file Loading:** `tool_version_override:` maps merge across `--config-dir` files. Same tool name with conflicting versions across files = hard error (determinism guarantee).

### Work Summary

**Files Changed:**
1. **docs/configuration.md:**
   - Updated Remote Skills section (lines 151-180) to add `version:` to example YAML and field reference table
   - Added note linking version pinning to "Tool Versioning & Custom Fetchers" section
   - Improved "Tool Versioning & Custom Fetchers" multi-file merge explanation

2. **CHANGELOG.md:**
   - Added entry under "### Changed" documenting the docs improvements

**Key Implementation Detail:**
Version pinning works identically for remote skills, plugins, and any tool entry that uses the git fetcher — no special handling needed. The `version:` field is the ONLY way to pin; there is no separate `ref:` or `branch:` field.

## Session: Redirect Documentation to Repo-Keyed Version Override

**Date:** 2026-04-25
**Requested by:** Ronnie Geraghty (autonomous mode — user is AFK)
**Status:** ✅ COMPLETE
**Task:** Update `docs/configuration.md` for the NEW repo-keyed `tool_version_override` shape (Morpheus' approved proposal). Previous work (`oracle-version-docs`) documented the OLD name-keyed shape and has been REVERTED; this work replaces it with repo-keyed documentation.

### Context

**Previous Oracle work:** `oracle-version-docs` spawned to document tool versioning using **name-keyed** `tool_version_override` (keyed by tool entry name, e.g. `azure-sdk-java: "v1.4.2"`).

**Schema change (approved):** Morpheus' proposal migrates to **repo-keyed** override (keyed by repo, e.g. `Azure/azure-sdk-skills: "v1.4.2"`). Rationale: repos are the atomic pinning unit in the fetcher, not individual skill names. Monorepos with N skills from one repo no longer require N redundant override entries.

**Migration strategy:** Hard cut with a clear error message (user rewrites keys from tool names to repo names). Zero shipped configs use the old shape, so blast radius is tiny.

### Work Summary

**Files Updated:**
1. **docs/configuration.md**
   - Rewrote "Tool Versioning & Custom Fetchers" section (lines 461–551)
   - Added per-entry `version:` field documentation with examples
   - Documented `tool_version_override:` keyed by `owner/repo` (not `name`)
   - Included monorepo example showing multiple entries from one repo picking up single override
   - Included per-entry vs. override precedence example
   - Documented `github.com/` prefix normalization (bare form preferred)
   - Documented multi-file merge semantics: maps merge; conflicting values for same repo = hard error; identical values merge silently
   - Added "Migrating from name-keyed overrides" callout with error message pattern and before/after YAML

   - Updated "Remote Skills" section (lines 151–178)
     - Added `version:` to the field reference table
     - Added inline comment in YAML example showing optional `version:` field
     - Added cross-link to "Tool Versioning" section

### Design Patterns Applied

**Documentation Precedence:**
- Proposal-first: Read the approved decision before writing docs (ensures no misinterpretation)
- Hard-cut migration messaging: Explicitly quote the error users will see, then provide the fix
- Multi-file merge semantics: Clearly explain conflict detection (determinism guarantee for splits configs)
- Reuse, don't duplicate: Remote Skills section links to Tool Versioning; doesn't re-explain the same fields

**Content Organization:**
- Per-entry pinning first (simplest case), then override map (most powerful case), then merge rules (infrastructure detail)
- Migration callout as a separate subsection (discoverable when users see the error)
- YAML examples show the "happy path" (things that work), not edge cases

### Learnings

**Proposal-to-Docs Workflow:**
- Always read the full proposal (including edge cases, conflict detection, and migration strategy) before writing docs
- The proposal's concrete examples (monorepos, per-entry override, prefix normalization) directly translate to doc YAML
- The proposal's error messages should be quoted in the docs so users recognize what they see
- Hard-cut schema migrations need a dedicated "migration" section, not buried in a footnote

**Schema Naming Stability:**
- When a Go field changes **meaning** (keying semantics) but keeps the same **name** (ToolVersionOverride), keep the YAML field name stable (tool_version_override) — users don't need to do a global find-replace
- The Go type stays `map[string]string`; only validation and lookup sites change
- This reduces user friction on configs that need updating

**Multi-file Merge Behavior:**
- Conflicting values = hard error is the right choice for version pinning (silent last-wins would hide misconfigurations)
- Identical values merge silently to allow configs to be split without duplicating pins
- Document both cases explicitly; users often split configs and need to know what's safe

### Files Staged (not committed)

- `docs/configuration.md` (updated per-entry + override + migration sections, plus Remote Skills cross-link)

Ready for Scribe to commit.


## Session: Limits Documentation Update (2026-04-24)

**Date:** 2026-04-24  
**Task ID:** opta-docs  
**Status:** ✅ COMPLETE  

### Work Summary

Updated `docs/configuration.md` to explicitly clarify that the limit resolution order (prompt frontmatter > config YAML > CLI flag > engine default) applies to **real-time enforcement** during code generation, not just post-hoc reporting.

### Files Changed

1. **docs/configuration.md** (line 418)
   - Added one clarifying sentence: "These resolved limits are enforced in real-time during code generation, ensuring that session actions, turns, and file creation respect the merged priority order."
   - Placement: immediately after resolution order statement in Limits section
   - Tone: terse, Microsoft Style Guide
   - No fearmongering or historical bug notes

### Commit

- **Commit:** 4a8cd9d0
- **Branch:** ronniegeraghty/dev
- **Message:** docs(configuration): clarify real-time enforcement of resolved limits

### Related

- **Decision file:** `.squad/decisions/inbox/morpheus-maxturns-enforcement-bug.md` — bug fix context (Option A implemented)
- **Team context:** `SetLimitsForEval()` now threads resolved limits into runner for real-time enforcement


## CROSS-AGENT UPDATE (2026-04-28T00-54-38Z — Scribe: Tool-Load Gate Fix — Option A Shipped)

**Decision shipped:** Morpheus investigated. Neo implemented. Switch tested. Oracle documented.

**Documentation:** Added "Post-Session Tool Verification" section to docs/configuration.md. Explains AssistantTurnStart as primary gate, per-kind failure reasons, 5-minute ceiling semantics.

**Commit:** f53eb3b1. User-facing guide updated.

**Result:** Full feature ship complete: Investigation → Implementation → Testing → Documentation.

---

## Session: Grader Documentation Audit & Update (2026-05-01)

**Date:** 2026-05-01  
**Task ID:** oracle-graders-docs (Charter)  
**Status:** ✅ COMPLETE  

Audited and updated docs/graders/ to match current canonical grader system. Recent grader redesign shipped with flattened envelope and new canonical types.

### Findings vs. Code

**Verified Current Schema (all 5 canonical types now ship with top-level `checks:`):**

1. **program** — single check kind: `command` (only kind supported)
   - Executes command in workspace; exit 0 = pass
   - Score = passed_checks / total_checks
   - Fields: kind, command, args, timeout

2. **workspace** — six check kinds (replaces legacy output_check + file):
   - `require_to_create`, `forbidden_to_create` — file path must/must-not be in NewFiles
   - `required_to_update`, `required_to_delete`, `forbidden_to_delete` — delta operations
   - `file` — on-disk state (present/absent) + optional size/content checks (min_bytes, max_bytes, contains, excludes)
   - Boolean grader (all checks must pass)

3. **tool** — four check kinds (replaces legacy tool_constraint, behavior, tool_usage):
   - `tool_used` (optional min_calls/max_calls), `tool_not_used`
   - `any_from_group` (optional except list), `none_from_group` (optional except list)
   - Boolean grader (all checks must pass)

4. **activity** — seven check kinds (replaces action_sequence, behavior):
   - `turn_limit` (max bounds)
   - `action_count`, `tool_call_count` (min/max bounds)
   - `contains_subsequence` (tools array)
   - `contains_action` — **NEW SHAPE TODAY**: type/tool/contains/excludes filters + min/max count bounds (default min=1)
   - `excludes_action` — negative form of contains_action (count must be 0)
   - `terminated_by` (equals or not_in list)
   - Boolean grader (all checks must pass)
   - **Note:** `not_truncated` was DELETED today; removed from docs

5. **prompt** — LLM-judged review:
   - Top-level `prompt:` (optional preamble) + `checks:` (list of strings, each = one Point)
   - One model per grader instance

### Files Updated

1. **docs/graders/index.md** — Complete rewrite
   - Documented all 5 canonical types with unified flat-checks schema
   - Added comprehensive schema overview with `checks:` at top level (not nested)
   - Added "All Five Canonical Grader Types" example showing prompt, workspace, tool, activity, program
   - Added deprecation table (file, output_check, behavior, action_sequence, tool_constraint, tool_usage)
   - Clarified engine-internal vs. user-configurable (prompt_review is internal)

2. **docs/graders/program.md** — Updated for flattened schema
   - Changed examples to use top-level `checks:` with `kind: command`
   - Documented that only `command` kind is supported
   - Added multiple examples showing sequential checks, build + test combinations
   - Added troubleshooting section

3. **docs/graders/prompt.md** — Fixed deprecation references
   - Updated references from deprecated `output_check`, `file`, `tool_constraint` to canonical `workspace`, `program`, `tool`

4. **docs/graders/workspace.md** — NEW (replaces output_check + file)
   - Full documentation of 6 check kinds with schema table for each
   - Examples: basic file creation, forbid secrets, verify content, prevent deletion, update requirements, comprehensive check
   - Comprehensive "Data Visible to Grader" section
   - Troubleshooting section

5. **docs/graders/tool.md** — NEW (replaces tool_constraint, behavior, tool_usage)
   - Full documentation of 4 check kinds with schema table for each
   - Examples: basic requirements, call bounds, forbid dangerous, group checks, comprehensive validation
   - Tool name resolution rules + group definition guidance
   - Troubleshooting section

6. **docs/graders/activity.md** — NEW (replaces action_sequence, behavior)
   - Full documentation of 7 check kinds with schema table for each
   - Valid action types enumerated
   - Valid termination reasons enumerated (completed, max_actions, max_turns, guardrail, error)
   - Examples: basic limits, action sequence, action presence, exclude patterns, content validation, comprehensive check
   - **Contains_action repurposing documented:** now supports type/tool/contains/excludes + min/max count
   - **Excludes_action documented:** negative form, count must be 0
   - Troubleshooting section

7. **docs/graders/output_check.md** — Deprecated
   - Added clear "DEPRECATED" header with deprecation notice
   - Migration guide showing old config → new workspace equivalent
   - Field mapping table (output_check fields → workspace check kinds)

8. **CHANGELOG.md** — Updated Unreleased section
   - Added comprehensive "Grader documentation audit and schema update" entry under Changed section
   - Documented new canonical graders, flattened schema, new docs created, deprecated docs deleted, activity grader notes

### Files Deleted

- `docs/graders/action_sequence.md` (legacy, replaced by activity)
- `docs/graders/behavior.md` (legacy, replaced by activity + tool)
- `docs/graders/file.md` (legacy, replaced by workspace)
- `docs/graders/tool_constraint.md` (legacy, replaced by tool)

### Verification

1. ✅ Matched code against docs for all 5 canonical types
   - Checked hyoka/internal/criteria/graders/{types,activity,workspace,tool,program,prompt}_grader.go
   - Verified check kinds, field names, validation rules

2. ✅ Verified flattened schema in criteria/language/test.yaml
   - All 5 canonical types use top-level `checks:` at example level
   - No nested `details:` object for canonical types

3. ✅ Verified check kind validators in code match docs
   - Activity grader: validActivityCheck() confirms 7 kinds + new contains_action/excludes_action shapes
   - Workspace grader: validateWorkspaceCheck() confirms 6 kinds
   - Tool grader: validateToolCheckRule() confirms 4 kinds
   - Program grader: only `command` kind allowed

4. ✅ Cross-reference linkage verified
   - index.md links to new markdown files (activity, workspace, tool)
   - Deprecated docs properly linked in migration sections
   - CHANGELOG references correct doc file paths

### Learnings

**Canonical Grader Inventory (for future doc updates):**

The 5 canonical graders are the authoritative types. Key code paths for future reference:

- **Schema container:** hyoka/internal/criteria/graders/types.go (GraderConfig + kind-specific *Config structs)
- **Validators:** Each grader's NewXxxGrader constructor validates config + checks
  - Program: Only `kind: command` supported; timeout defaults to 30s
  - Workspace: 6 kinds; file checks (state: present|absent) validate state-check field combinations
  - Tool: 4 kinds; except: list is optional for group checks
  - Activity: 7 kinds; contains_action supports type/tool/contains/excludes + min/max; min defaults to 1
  - Prompt: No validation of rubric/prompt content (LLM determines success)

- **Test reference:** criteria/language/test.yaml exercises all 5 kinds (canonical example)

**Documentation Pattern for Future Grader Changes:**

1. **If check kind is added/removed:** Update activity_grader.go (kinds list), then sync:
   - docs/graders/activity.md (check kind table)
   - docs/graders/index.md (check kinds reference)
   - CHANGELOG.md (Breaking Changes or Changed section)

2. **If field added to check:** Update schema table + example + troubleshooting

3. **If schema changes (e.g., another envelope flatten):** All 5 grader docs need updates + index.md + CHANGELOG

**Schema Pattern (no more legacy cruft):**

- All canonical graders: checks at top level, NO details wrapper
- Engine-internal graders (prompt_review): may use different internal structure (not user-documented)
- Deprecation notices: Always include migration path (old shape → new shape with field mapping)

**Code-vs-Doc Discrepancies Found:** None. All grader behavior matches documentation.

### Commit

Ready for commit on `ronniegeraghty/dev` with standard Co-authored-by trailer.

---

### 2026-05-02 — Tool Grader Breaking Change: `tool: skill` No Longer Matches (Neo)

**Sync:** Oracle must update docs/graders/tool.md in next iteration.

**What Changed:**
- Neo fixed tool_used grader double-counting bug
- Filter: redundant tool_call events with Tool="skill" removed from action log
- **Breaking:** Criteria using `tool: skill` (catch-all skill matching) no longer works

**Migration Path:**
- Replace `tool: skill` with individual skill name: `tool: markdown-headings`
- OR use future group matching: `any_from_group: <skill-group-name>` (pending implementation)

**Root Cause:**
skill events landed twice in toolCounts:
1. tool.execution_start → Tool="skill" (generic wrapper)
2. skill.invoked → Tool="markdown-headings" (individual name)

Removing (1) eliminates double-counting; (2) provides canonical match key.

**Code Reference:** hyoka/internal/eval/action.go, TestActionTimeline_ToGraderActionLog_SkillEvents

---

## Session: Comprehensive Graders Documentation Audit (2026-05-02)

**Date:** 2026-05-02  
**Status:** ✅ COMPLETE  
**Branch:** ronniegeraghty/dev  
**Decision Drop:** `.squad/decisions/inbox/oracle-grader-docs-audit.md`

### Audit Scope

Comprehensive review of all hyoka documentation touching graders, criteria, and review systems. Verified against current Go implementations and identified documentation gaps.

### Critical Issues Found & Fixed

**CRITICAL #1: Undocumented Tool Grader Fields**
- Issue: `source` and `mcp_server` fields in ToolCheckRule struct were implemented in code but not documented
- Evidence: tool_grader.go lines 76-82, 131, 155-158 actively use these fields
- Fix: Added field documentation to tool.md with table updates + "Filtering by Tool Source" example

**CRITICAL #2: Architecture.md Lists Non-existent Graders**
- Issue: Listed `output_check` and `action_sequence` as canonical (don't exist), missing `workspace` and `activity`
- Evidence: Schema flatten commit 7410ecf1 removed all deprecated kinds; types.go defines only 5 canonical + 1 internal
- Fix: Updated canonical grader list to accurate five types; moved legacy kinds to separate "deprecated" section

**CRITICAL #3: grader-config-schema.md Not Marked as Legacy**
- Issue: Document claimed to describe "current schema" but was pre-v4 with removed graders, `details:` wrapper, etc.
- Fix: Rewrote header to mark as OBSOLETE, added redirect to current docs, created grader removal table

**MEDIUM #1: WorkspaceDelta Nil Handling Undocumented**
- Issue: workspace.md didn't mention WorkspaceDelta is a pointer and may be nil in older reports
- Fix: Added "Important: WorkspaceDelta Availability" section explaining v0.4+ behavior vs legacy

### Files Modified

| File | Changes |
|------|---------|
| docs/architecture.md | Fixed canonical grader list (removed output_check/action_sequence, added workspace/activity) |
| docs/graders/tool.md | Documented source + mcp_server fields, added filtering examples |
| docs/graders/workspace.md | Added WorkspaceDelta nil handling section |
| docs/grader-config-schema.md | Rewrote as obsolete legacy reference, removed deprecated grader docs |

### Code-vs-Docs Verification

✓ Five canonical graders match types.go constants (program, prompt, workspace, tool, activity)  
✓ PromptReviewGrader confirmed as engine-internal (not user-configurable YAML)  
✓ Tool grader Source/MCPServer fields match active implementation  
✓ WorkspaceDelta pointer type verified in GraderInput struct  
✓ Removed graders confirmed via schema-flatten commit  
✓ Prompt architecture clarified: `type: prompt` criteria ≠ PromptGrader class  

### Testing & Validation

- Criteria files validate against current schema: ✓ test.yaml, python.yaml use documented structure
- Tool grader: ✓ All four check kinds (tool_used, tool_not_used, any_from_group, none_from_group) exercised
- Activity grader: ✓ All activity checks (turn_limit, action_count, tool_call_count, contains_action, etc.) shown
- Workspace grader: ✓ All six check kinds (require_to_create, forbidden_to_create, etc.) demonstrated

### Architectural Insights Documented

- `type: prompt` in criteria YAML flows to review panel (NOT through graders.NewGrader)
- Prompt criteria entries have different structure than PromptGrader class (design separation verified)
- Both architectures correct; docs now clarify the distinction in prompt_review.md

### Next Steps

- Monitor neo/issue-grader-redesign for any schema evolution
- If prompt:+checks: structure moves to runtime PromptGrader in future, update docs/graders/prompt.md
- Consider architecture diagram for "Criteria → Review Panel vs Graders → Engine" flow

**Status:** All critical documentation gaps resolved. Docs now match current hyoka codebase (commit 0a4d1fd9).

---

## CROSS-AGENT UPDATE (2026-05-02T04:20:51Z — Session: Tool-Used Disambig + Docs Audit)

**Agents Involved:** Neo (feature), Switch (testing), Oracle (docs), Morpheus (scoping)

**Oracle's Work Impact:**
- Conducted comprehensive grader documentation audit
- Fixed 4 critical doc issues:
  1. **Documented Neo's new tool_used fields:** Added `source` and `mcp_server` to `docs/graders/tool.md`
  2. **Fixed canonical grader list in architecture.md:** Removed 2 phantom entries (output_check, action_sequence), added 2 missing (workspace, activity)
  3. **Flagged obsolete legacy schema doc:** Marked `docs/grader-config-schema.md` as LEGACY (pre-v4)
  4. **Documented WorkspaceDelta nil handling:** Added availability section to `docs/graders/workspace.md`
- All documentation now consistent with current code implementation

**Cross-team Impact:**
- Architecture now has authoritative canonical grader list (5 graders: program, prompt, tool, workspace, activity)
- Users have clear guidance on tool source/server disambiguation via updated docs
- Legacy schema references no longer confuse new readers

**Status:** Documentation audit complete, all grader docs accurate and current.

---

---

### 2026-05-02 — Tool Grader Fields Audit (Oracle)

**Task:** Verify that `source` and `mcp_server` fields on `tool_used` and `tool_not_used` checks are fully and accurately documented.

**Findings:**

**Coverage: ALL REQUIREMENTS MET** ✅

1. ✅ **source field** — Documented with all 3 values (skill/mcp/builtin) in both tool_used and tool_not_used tables (lines 46, 57)
2. ✅ **mcp_server field** — Documented in both tables (lines 47, 58) with note: "only meaningful with `source: mcp`"
3. ✅ **Validation rule** — mcp_server requires source=mcp, documented in table descriptions
4. ✅ **Both check kinds** — Identical source/mcp_server fields on tool_used and tool_not_used (NOT on group checks)
5. ✅ **Group checks clarification** — Schema tables show any_from_group/none_from_group lack these fields
6. ✅ **YAML examples** — Section "Filtering by Tool Source" (lines 168-189) demonstrates all three source values + mcp_server usage
7. ✅ **test.yaml reference** — No prior mention; added new Reference section

**Gaps Found: 3 (all minor, all fixed)**

| Gap | Severity | Fix | Location |
|-----|----------|-----|----------|
| Validation rule not explicit | Low | Added Notes bullet: "If `mcp_server` is specified, `source` must be set to `mcp`" | Line 232 |
| Scope of source/mcp_server unclear | Low | Added Notes bullet: "These apply only to `tool_used` and `tool_not_used`. Group checks do not support..." | Line 231 |
| No canonical test reference | Low | Added Reference section linking criteria/language/test.yaml | Lines 245-247 |

**Documentation Changes:**

File: `docs/graders/tool.md`
- Added 2 new bullets to Notes section (lines 231-232):
  - Scope clarification: source/mcp_server fields apply only to tool_used and tool_not_used
  - Validation rule: mcp_server requires source=mcp (explicit error message)
- Added new Reference section (lines 245-247) linking to criteria/language/test.yaml with description

**Files Audited:**
1. docs/graders/tool.md — **UPDATED**
2. docs/graders/index.md — OK (high-level overview, not grader-specific)
3. docs/architecture.md — OK (doesn't detail grader fields)
4. README.md — OK (no grader field specifics)
5. criteria/language/test.yaml — OK (already canonical example)

**Code Verification:**
- Cross-referenced hyoka/internal/criteria/graders/tool_grader.go validation rules (lines 76-95)
- Confirmed validation error messages match code error text
- Confirmed all 4 check kinds (tool_used, tool_not_used, any_from_group, none_from_group) correctly represented

**Decision:** Minor edits to improve clarity and discoverability. No factual corrections needed—implementation was already complete and mostly well-documented.

**Learnings:**
- Tool grader documentation is comprehensive and accurate (Neo's prior work was thorough)
- Main gaps were clarity/discoverability, not factual errors
- Canonical test example (test.yaml) valuable reference but not previously mentioned in docs
- Validation rules benefit from explicit error-message phrasing, not just "only meaningful"


## Session: 2026-05-02 — Tool Grader Fields Documentation Audit

**Type:** Feature Verification  
**Partner:** Neo (parallel work updating criteria YAML)  
**Timestamp:** 2026-05-02T04-30-57Z

### Work Done

Audited `docs/graders/tool.md` to verify completeness of `source` and `mcp_server` field documentation following Neo's implementation. Found and fixed 3 low-severity gaps:
- Added explicit MCP server validation error rule
- Clarified scope (don't apply to group checks)
- Added reference section linking to canonical test.yaml example

### Partner Context

Neo updated canonical criteria YAML files to use new source/mcp_server fields, providing reference examples for tool source disambiguation. Both agents worked on same feature topic.

### Outcomes

✅ 7/7 documentation requirements verified  
✅ 3 gaps resolved  
✅ Code alignment verified against tool_grader.go  
✅ Documentation now complete and discoverable  

---

---

## CROSS-AGENT NOTE (2026-05-05 — PR #640: Action Semantics Clarification)

**From:** Scribe (via Ronnie's directive in PR #640 spawn manifest)  
**Impact:** How actions are counted in eval sessions

Oracle should be aware that per Ronnie's clarification ("anything the agent does is meant to be an action from reasoning to tool calls to bash commands, to responses"), **`assistant.reasoning` events count as actions** for the purpose of session limits (`maxSessionActionsLimit`). This is already implemented and verified in `hyoka/internal/eval/copilot.go:359-365`. The PR #640 port (commit 703f638b) ensured the action counter respects per-eval limit overrides at all sites.

**No immediate action needed.** This is context for future decisions involving action-counting or session-limit tuning.
