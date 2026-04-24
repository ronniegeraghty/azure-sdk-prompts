# Session Log: Examples Validation + PR #607 Comment Follow-up

**Date:** 2026-04-22T17:20:00Z
**Branch:** `phase-6`
**Agents:** Morpheus (🏗️) × 2 runs
**Mission:** 
1. Audit `examples/` directory artifacts against schema
2. Investigate PR #607 review comment 3125681737 on hierarchical criteria

## Run 1: Examples Validation

**Result:** ✅ 8/8 real artifacts valid; 1 intentional skeleton. No action needed.

**Deliverable:** Optional follow-up to rename `prompt-template.prompt.md` → `prompt-template.md` (low priority).

**Learnings:** Staging workaround documented in Morpheus history.

---

## Run 2: PR #607 Hierarchical-When Investigation

**Result:** ❌ Example misleading; loader has silent multi-doc truncation bug.

**Finding:** YAML `---` separator suggests two group-level `when` blocks valid, but `criteria.go:134-136` only decodes first doc. Rust block silently lost. Correct pattern is top-level `groups:` list (one `when` per group).

**Recommendation:** File 3 separate issues:
1. Fix example (use `groups:` list)
2. Fix loader (reject or merge multi-docs)
3. Validate coverage gap (detect trailing YAML docs)

**Outcome:** ✅ Issue surfaced; threaded comment posted on PR #607 (comment 3125721580).

**Learnings:** Neo and Oracle histories updated with loader bug + example audit findings.

---

## Decisions merged

Both Morpheus runs generated decision inbox entries; merged into canonical `decisions.md` (see below).
