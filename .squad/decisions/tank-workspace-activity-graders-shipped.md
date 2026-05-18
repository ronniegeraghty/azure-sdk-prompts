# Tank — Workspace and Activity Graders Shipped

**Date:** 2026-05-01  
**Agent:** Tank 📡  
**Status:** COMPLETE

---

## Summary

Delivered commits A and B per Morpheus's grader taxonomy redesign plan:
- **Commit A (1f461a50):** Workspace grader (replaces `output_check`)
- **Commit B (0896ba53):** Activity grader (replaces `action_sequence`, drops `behavior`)

Both graders are canonical, production-ready, with table-driven tests and loud parse errors for deprecated types.

---

## Commit A: Workspace Grader (1f461a50)

### Changes
- Renamed `output_check` → `workspace` with six check kinds
- All checks driven from WorkspaceDelta (NewFiles / ModifiedFiles / DeletedFiles)
- Check kinds:
  - `require_to_create` — path must be in NewFiles
  - `forbidden_to_create` — path must NOT be in NewFiles
  - `required_to_update` — path must be in ModifiedFiles
  - `required_to_delete` — path must be in DeletedFiles
  - `forbidden_to_delete` — DeletedFiles empty (when files:["*"]) or specific paths absent
  - `file` — state present|absent with optional min_bytes, max_bytes, contains, excludes
- Loud parse error on `type: output_check` pointing to workspace + migration guide

### Files
- `hyoka/internal/criteria/graders/workspace_grader.go` (new, 340 lines)
- `hyoka/internal/criteria/graders/workspace_grader_test.go` (new, 17 test cases)
- Updated: config.go, grader.go, registry.go, types.go

### Tests
- 17 table-driven test cases covering all 6 check kinds + validation
- All tests pass with `-race` flag

---

## Commit B: Activity Grader (0896ba53)

### Changes
- Renamed `action_sequence` + dropped `behavior` → single `activity` grader
- Powered by ActionLog, ActionsSummary, TerminatedBy (from GeneratorArtifact)
- Seven check kinds:
  - `turn_limit` — max turn ≤ configured max (migrated from tool grader)
  - `action_count` — TotalActions in [min, max]
  - `tool_call_count` — ToolCalls in [min, max]
  - `contains_subsequence` — ordered subsequence of tool names
  - `contains_action` — specific tool with optional min/max call counts
  - `not_truncated` — ActionsSummary.Truncated == false
  - `terminated_by` — TerminatedBy matches expectation (equals or not_in)
- Loud parse errors on `type: action_sequence` and `type: behavior`

### Files
- `hyoka/internal/criteria/graders/activity_grader.go` (new, 370 lines)
- `hyoka/internal/criteria/graders/activity_grader_test.go` (new, 7 test cases)
- Updated: config.go, grader.go, registry.go, types.go

### Tests
- 7 focused test cases covering all 7 check kinds
- All tests pass with `-race` flag

---

## Coordination Notes

- Neo's pairwise fix (4f293e06) and tool grader redesign (24de2f26) landed between my commits
- Neo's tool_grader_test.go has pre-existing build failures (uses old ToolCheckRule fields) — not blocking
- Switch updating criteria fixtures in parallel
- No conflicts with Neo's work; shared files (config.go, types.go) touched different sections

---

## Verification

```bash
# Build
go build ./hyoka/...  ✅

# Workspace tests
go test -race ./hyoka/internal/criteria/graders/... -run TestWorkspace  ✅

# Activity tests
go test -race ./hyoka/internal/criteria/graders/... -run TestActivity  ✅
```

---

## Next Steps (not blocking)

1. **Site component renames** — OutputCheckExtras → WorkspaceExtras, ActionSequenceExtras → ActivityExtras
2. **Docs update** — docs/graders.md canonical reference with new check kinds
3. **Criteria fixtures** — Switch updating criteria/language/test.yaml + python.yaml

---

## Commit SHAs

- Workspace grader: **1f461a50**
- Activity grader: **0896ba53**

Both pushed to `ronniegeraghty/dev` and ready for review/merge.
