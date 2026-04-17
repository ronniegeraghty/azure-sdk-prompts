# Wave 1 Review Follow-ups

**Author:** Morpheus 🕶️  
**Date:** 2026-04-18  
**Source:** Architectural reviews of PRs #571 and #572

---

## Immediate follow-up work required

### 1. Fix workspace_delta TS field names — BLOCKING before real delta data flows

**Who:** Neo (owns #566) or Trinity (owns types.ts)  
**Scope:** Small PR, ~20 lines

The TS `workspace_delta` inline type in `EvalReport` uses invented field names that don't match Go JSON output:

| TS field (wrong) | Go JSON tag (correct) |
|-------------------|-----------------------|
| `files_created` | `new_file_count` |
| `files_modified` | `modified_file_count` |
| `files_deleted` | `deleted_file_count` |
| `total_size_bytes` | `bytes_net` (or `bytes_added`) |

Additionally, Go emits `new_files`, `modified_files`, `deleted_files` arrays and `bytes_added`/`bytes_removed` — none reflected in TS.

**Action:** Extract a named `WorkspaceDelta` interface in `types.ts` matching Go exactly. Update `eval-detail-page.tsx` lines 461–464 to use correct field names.

### 2. Wire grader_results + workspace_delta into EvalReport TS type

**Who:** Neo (#571 scope) or Trinity (#572 follow-up)  
**Scope:** Small, ~5 lines on EvalReport interface

PR #572 added `grader_results?: GraderResult[]` and `workspace_delta` to `EvalReport` — but PR #571 (which was supposed to establish these types) did not. Verify #571's scope is correct and these additions land properly.

### 3. Resolve EvalResult vs EvalReport type confusion

**Who:** Trinity  
**Scope:** Medium, can batch with #359

`eval-detail-page.tsx` fetches `RunSummary` then accesses detail fields via 10+ type casts. A `fetchEval()` function returning `EvalReport` already exists in `api.ts`. Either:
- Switch to `fetchEval()` for detail pages, or
- Widen `EvalResult` to be a proper superset

If not resolved, #359 and #360 will inherit the same casting anti-pattern.

---

## Non-blocking items for tracking

4. **Add `UnchangedFileCount` to Go `WorkspaceDelta`** (Neo) — useful for rendering context and grader reasoning
5. **Populate `WorkspaceDelta` at runtime** — wire `TakeSnapshot`/`ComputeDelta` in `engine_eval.go` (Neo)
6. **Add `WorkspaceDelta` to `GraderInput`** — graders need delta context for future threshold work (Neo)
