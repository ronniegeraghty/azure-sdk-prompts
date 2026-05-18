# Session Log: YAML Multi-Document Issue #621

**Timestamp:** 2026-04-22T17:36:47Z  
**Agent:** Morpheus  
**Context:** PR #607 hierarchical-when investigation + formal architectural proposal

## Session Summary

Consolidated findings from PR #607 discovery into formal architectural proposal (issue #621) for YAML multi-document stream support across prompts, configs, and criteria.

## Work Done

1. **PR #607 comment reply:** Explained hierarchical-when schema to Ronnie via in-thread response (example uses `---` separators incorrectly; canonical shape uses `groups:` list from `GraderConfig`)
2. **Issue #621 filed:** Comprehensive architectural proposal with:
   - Current behavior analysis (silent truncation in all three loaders)
   - Proposed semantics per artifact type (merge for criteria/configs, reject for prompts)
   - Implementation sketch (decode loop with merge logic)
   - Acceptance criteria (checkable)
   - Loader locations for future reference
3. **History updated:** Appended investigation results (loader file paths, pattern analysis, learnings) to Morpheus history.md

## Inbox Decisions Closed

- `morpheus-examples-validate.md` (examples directory audit complete)
- `morpheus-pr607-when-followup.md` (follow-up comment posted to thread)

## References

- Issue: https://github.com/ronniegeraghty/hyoka/issues/621
- PR #607 comment: In-thread reply (hierarchical-when explanation)
- Commit 8b6fce91: Example rewrite (criteria-only, single `groups:` list schema)
- Morpheus history.md: Lines 554–587 (investigation results)
