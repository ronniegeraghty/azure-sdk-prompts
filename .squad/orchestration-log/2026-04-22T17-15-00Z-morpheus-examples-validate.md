# Orchestration Log: Morpheus — Examples Directory Validation Audit

**Date:** 2026-04-22T17:15:00Z
**Agent:** Morpheus (🏗️)
**Mission:** Validate all artifacts under `examples/` (prompts, configs, criteria) against schema.

## Scope

- 3 prompt files + 1 template skeleton
- 3 config files  
- 8 criteria files (hierarchical, language-specific, service-specific)

## Findings

**Result:** 8/8 *real* artifacts schema-valid. 1 file (`prompt-template.prompt.md`) intentionally invalid — it's a copy-and-fill skeleton with empty required fields.

| Kind | File | Status |
|---|---|---|
| prompt | `example.prompt.yaml` | ✅ |
| prompt | `graders-frontmatter-example.prompt.md` | ✅ |
| prompt | `existing-files-example.prompt.md` (+ `.starters/`) | ✅ |
| prompt | `prompt-template.prompt.md` | ❌ intentional skeleton |
| config | `example-full.yaml` | ✅ |
| config | `example-generator-skills.yaml` | ✅ |
| config | `example-remote-skill.yaml` | ✅ |
| criteria | `hierarchical-when-example.yaml` | ✅ |
| criteria | `language/*.yaml` (5 files) | ✅ |
| criteria | `service/*.yaml` (2 files) | ✅ |

## Recommended follow-up

Optional, low-priority: rename `prompt-template.prompt.md` → `prompt-template.md` to skip validation scope. Keeps the template human-readable; no schema fix needed.

## Learnings captured

Updated Morpheus history.md with staging quirk: `starter_project` paths resolved relative to prompt file's directory; when staging with prefix, starter dir symlink must exist under unprefixed name.

## Outcome

✅ Examples directory is healthy reference material. Validate schema working correctly.
