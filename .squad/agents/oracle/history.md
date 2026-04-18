# Oracle — History

## Project Context

- **Project:** hyoka — Go evaluation tool for AI-generated Azure SDK code, powered by Copilot SDK and multi-model review panels.
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers
- **User:** Ronnie Geraghty
- **My domain:** Docs — docs/, README.md, AGENTS.md, CHANGELOG.md, inline documentation

## Learnings

### Prompt Frontmatter Schema (Nested Properties Format)

**Current format (all existing prompts):** `id` and `tags` are top-level keys; all other metadata (service, plane, language, category, difficulty, description, sdk_package, doc_url, created, author, etc.) is nested under `properties:` map.

```yaml
id: key-vault-dp-python-crud
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: '...'
  sdk_package: azure-keyvault-secrets
  doc_url: https://...
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- secrets
- crud
```

**Parser reference:** `hyoka/internal/prompt/parser.go` lines 16–34 define `frontmatter` struct with `ID`, `Tags`, and `Properties map[string]string`. Flat frontmatter compat was removed in Phase 3 (nested is the only supported format).

**Why this matters:** All doc examples must use this schema. When updating docs or helping authors, reference this nested format. The `.prompt.yaml` format also uses the same `id/tags/properties:` structure.

### Session 2026-04-04T00-05 (Morpheus Evolution Plan)

Evolution plan assigns you Go version doc fixes and documentation for all new features across phases. Read `.squad/decisions.md` for full plan.

### Session 2026-04-04T19:48 (Phase 0 Execution — Go Version Update)

**Status:** COMPLETE  
**PR:** #169

Updated Go version references from 1.24.5 to 1.26.1 across 15 files including go.mod, go.work, CI config, and documentation.

**Cross-agent notes:** No conflicts with other Phase 0 work (Neo's reviewer factory, Tank's CI/config, Switch's tests). All agents' code compatible with Go 1.26.1.

**Files:** go.mod, go.work, .github/workflows/ci.yml, docs, README, inline comments

### Session 2026-04-05T00-00 (Prompt Authoring Docs Update)

**Status:** COMPLETE
**PR:** ronniegeraghty/update-prompt-authoring-docs (branch created, push ready)

Updated `docs/prompt-authoring.md` to accurately reflect the current prompt frontmatter format (nested `properties:` schema) used by all existing prompts in the repository.

**Changes:**
- Reordered sections: Frontmatter Schema first (before YAML-only format)
- Updated all frontmatter examples to show `id` and `tags` at top level, with service/plane/language/category/difficulty/description etc. nested under `properties:`
- Added `max_session_actions`, `max_turns`, and `project_context` optional fields with descriptions
- Clarified field optionality (required vs optional)
- Updated `.prompt.yaml` format examples to match the nested structure
- Maintained existing doc tone, structure, and guidance

**Reference sources:** Surveyed 10+ real prompts across key-vault, cosmos-db, app-configuration, resource-manager, storage to verify format. Cross-referenced `parser.go` lines 16–34 to confirm schema structure.

### Session 2026-04-16T16:55 (Prompt Authoring Docs Refresh — Completion)

**Status:** COMPLETE  
**Branch:** ronniegeraghty/update-prompt-authoring-docs  
**Commit:** f6439bed  
**Orchestration Log:** `.squad/orchestration-log/2026-04-16T16:55:15Z-oracle.md`

Verified docs/prompt-authoring.md update is production-ready. Schema documentation now fully aligned with nested `properties:` format. All examples validated against parser source and 10+ production prompts. Ready for PR review and merge.


## 2026-04-16 — Phase 3 Merged to Dev (Neo)

Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has both Phase 3 features and starter-aware guardrail fix. All tests pass, CI green.

### Session 2026-04-16T22:45 (Phase 4 Kickoff: Examples Update #363)

**Status:** COMPLETE  
**PR:** #568

Updated `examples/configs/example-full.yaml` to reflect Phase 3 unified grading architecture (PR #562):

**Changes:**
- Unified tools list: MCP servers now specified as `type: mcp` within `tools:` array (not separate `mcp_servers:` section)
- MCP entry format: `type: mcp`, `command`, `args`, `mcp_tools` (formerly `tools:` in mcp_servers)
- Updated comments to explain Phase 3 architecture: generator→reviewer→graders are now unified pipeline (review is a grader type, not separate phase)
- Completed minimal config example with `reviewer:` section (was previously incomplete)
- Added `limits:` section with all available options (max_turns, max_files, max_output_size, max_session_actions)
- All examples pass `go run . validate` (12 configs, 89 prompts, 2 criteria files valid)

**Reference:** Phase 3 PR #562 implemented unified grading pipeline where review panel is now `PromptReviewGrader` running alongside pluggable graders. Phase 4 brief §3 (Morpheus) directed examples update to reflect this architecture.

**No deletions needed:** `examples/configs/example-registry.yaml` was already gone (dead feature from pre-Phase-3 tool registry system). Brief mentioned deleting it; already removed in prior phase.

### Session 2026-04-16T23:33 (README Command Audit)

**Status:** COMPLETE  
**Branch:** ronniegeraghty/fix-readme-commands  
**Commit:** f5715393  

Audited and fixed all outdated commands in README.md. Many examples failed because they omitted the required `--config` flag (hyoka now requires either `--config <name>` or `--all-configs` when multiple configs exist).

**Key fixes:**
- Added `--config baseline/claude-opus-4.6` to all example run commands in Quick Start and Filtering sections
- Replaced non-existent "hyoka tools" command with "hyoka plugins" (tools command never existed)
- Fixed repo structure: `hyoka/cmd/hyoka/main.go` → `hyoka/main.go` (moved in PR #569)
- Updated heading "Tool Configurations" → "Configurations" (clearer, no command named "tools")
- Added `--prompt-id` to config examples to make them actually executable
- Clarified that `--all-configs` is required when running without specific config filter

**Testing method:** Ran every command from the README in dry-run mode or with `--help` to verify syntax. Used `go run ./hyoka` from worktree. All commands now work.

**Commands verified working:**
- `go run ./hyoka list`, `validate`, `check-env`, `configs`, `plugins`, `compare`, `trends`, `serve`, `report`, `new-prompt`, `clean`, `version`
- `go run ./hyoka run` with all flag combinations: `--service`, `--language`, `--prompt-id`, `--config`, `--all-configs`, `--dry-run`
- Build: `go build ./hyoka/...`

**Pattern for docs audits:** Always test commands in a live environment (worktree or main checkout). Use `--dry-run` for expensive operations. Cross-reference AGENTS.md for authoritative command syntax.

### Session 2026-04-17T01:15 (Phase 3.5 + Wave 1 CHANGELOG & Drift Sweep)

**Status:** COMPLETE  
**PR:** #576

CHANGELOG.md created and populated with 6 merged PRs from Phase 3.5 + Wave 1:

**Added entries:**
- PR #562: Unified grading architecture (AI review is now a grader type, structured JSON responses, deterministic voting)
- PR #567: Starter-aware MaxOutputSize guardrails (only count delta output, not starter files)
- PR #568: Examples updated for Phase 3 unified grading
- PR #570: Site rebranding from Azure SDK code-gen to general-purpose AI agent evaluation
- PR #571: WorkspaceDelta feature (file-level tracking) + eval detail page rendering
- PR #572: GraderResultRow component + eval detail redesign

**Drift fixes:**
- README.md: Rebranded intro from "Azure SDK code" to "comprehensive evaluation tool for AI agents"
- README.md: Updated `review/` package comment from "Multi-model review panel + rubric" → "Multi-model review (PromptReviewGrader)"
- README.md: Updated guardrails table to reflect starter-aware counting (MaxOutputSize counts deltas only)
- AGENTS.md: Same branding refresh (removed "Azure SDK code" reference)
- AGENTS.md: Same `review/` package comment update

**Drift patterns found:**
1. **Branding lag** — README and AGENTS still referenced "Azure SDK code generation" despite site rebrand (PR #570). Fixed by rewriting overview sections.
2. **Architecture comment drift** — `review/` package comment mentioned "rubric" (pre-Phase-3) instead of unified grader pattern. Fixed.
3. **Guardrails documentation lag** — guardrails table didn't mention starter-aware calculation (PR #567 hotfix). Added clarification that MaxOutputSize/MaxFiles count deltas only.

**Files most out-of-sync:**
1. README.md (line 3) — branding + line 87 guardrails table
2. AGENTS.md (line 5, 31) — branding + architecture comment
3. CHANGELOG.md — didn't exist yet (created from scratch)

**Watch for next wave:**
- PR bodies should guide CHANGELOG entries (most accurate)
- Site rebrands need systematic README/AGENTS refresh
- Architecture package comments need updates when grader types change
- Guardrails documentation needs to stay current with limit-calculation logic changes


## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release

Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-generation, serve endpoints, hierarchical criteria, cleanup. Recommendation: **Promote dev → main and cut v0.3.1 tag.**

Decision: .squad/decisions.md | Orchestration Log: .squad/orchestration-log/2026-04-17T20:53:40Z-morpheus.md

**2026-04-18 — Morpheus Phase 4 Verification:** Playwright-cli now established as standard for UI verification. Oracle's playwright-cli skill complete and validated across full feature set. Skill ready for team use.
