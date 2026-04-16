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

