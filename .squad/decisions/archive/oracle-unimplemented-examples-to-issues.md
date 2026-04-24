# Decision: Move Unimplemented Examples to Issues

**Date:** 2026-04-24  
**Author:** Oracle 🔮  
**Status:** IMPLEMENTED  
**Branch:** ronniegeraghty/dev  
**Commit:** 09cf17f0

## Summary

Removed two example files and documentation sections that sketched unimplemented parser features, converting them into properly scoped GitHub issues instead.

## What Happened

Per Ronnie's directive, I identified and cleaned up documentation debt:

1. **`examples/prompts/graders-frontmatter-example.prompt.md`** — Demonstrated a `graders:` frontmatter field the parser does not consume. Used pre-v4 schema (`kind/config/gate`). **→ Deleted**

2. **`docs/starter-files.md` "Option B"** — Sketched `starter_files:` list syntax as a lightweight alternative to `starter_project:` directory reference. Never implemented in parser. **→ Section removed**

3. **`examples/README.md` "Prompt-Level Graders"** — Documented graders in frontmatter (unimplemented feature). **→ Section removed**

## Issues Created

### Issue #636: Feature: support `graders:` list in prompt frontmatter

**URL:** https://github.com/ronniegeraghty/hyoka/issues/636  
**Labels:** enhancement, squad

Enables prompts to declare per-prompt graders inline via `graders:` array in frontmatter. Design mirrored the YAML grader shape (name, type, weight, prompt, checks).

**Implementation steps documented:**
- Add `Graders []Grader` field to Prompt struct
- Update parser to unmarshal `graders:` field
- Add validation checks
- Integrate into evaluation pipeline
- Add tests and docs

### Issue #637: Feature: support `starter_files:` list (Option B) in prompt frontmatter

**URL:** https://github.com/ronniegeraghty/hyoka/issues/637  
**Labels:** enhancement, squad

Enables prompts to list specific starter files inline (`starter_files: [./main.py, ./requirements.txt]`) as lightweight alternative to directory-based `starter_project:`.

**Implementation steps documented:**
- Add `StarterFiles []string` field to Prompt struct
- Update eval/copilot.go to copy individual files
- Add validation (path existence, no escape, no duplicates, conflict checks)
- Add tests and docs

## Files Changed

| File | Change |
|------|--------|
| `examples/prompts/graders-frontmatter-example.prompt.md` | Deleted via `git rm -f` |
| `examples/README.md` | Removed graders-frontmatter reference and "Prompt-Level Graders" section |
| `docs/starter-files.md` | Removed "Option B: Explicit File List" section, validation checks, struct field docs, roadmap |

## Rationale

**Documentation Hygiene Principle:**
- Examples must match current parser behavior — unimplemented features create confusion
- Sketches belong in issues, not in shipped docs
- Single source of truth: issue tracker, not stale examples
- Ensures `hyoka validate` reports only actionable findings

## Learnings for Future

1. **Unimplemented examples are debt** — They drift from reality and confuse users. Move to issues.
2. **Sketch → Issue → Implementation** — Design docs flow: idea → issue (scoped) → code
3. **Documentation reflects shipped behavior** — Never document "future" features in examples/docs shipped with releases.

## Sign-Off

✅ All deletions committed to `ronniegeraghty/dev`  
✅ Both features tracked as GitHub issues with full context  
✅ Examples directory now clean and accurate  
✅ Cross-references removed from README  

Ready to merge to dev → phase-6 → main in normal release cycle.
