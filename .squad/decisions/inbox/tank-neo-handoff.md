# Tank → Neo Handoff

**Date:** 2026-04-30  
**From:** Tank 📡  
**To:** Neo 💊  
**Re:** Grader redesign Part 3 — render-side ready for data model wiring

---

## Status

Tank's render-side work is committed on `tank/issue-grader-output-redesign` (commit `f3208e67`).
All rendering changes degrade gracefully when `SourceFile`/`SourceType` are empty strings.

---

## What's Waiting for Neo

### Fields Tank expects Neo to populate

| Struct | Field | JSON | Values |
|--------|-------|------|--------|
| `report.GraderResult` | `SourceFile string` | `source_file,omitempty` | Absolute path to originating file |
| `report.GraderResult` | `SourceType string` | `source_type,omitempty` | `"prompt_file"` or `"criteria_file"` |
| `graders.GraderResult` | `SourceFile string` | — | Same |
| `graders.GraderResult` | `SourceType string` | — | Same |
| `progress.ProgressEvent` | `GraderSourceFile string` | — | Already added by Tank |
| `progress.ProgressEvent` | `GraderSourceType string` | — | Already added by Tank |

Tank already added `GraderSourceFile` and `GraderSourceType` to `ProgressEvent` in
`hyoka/internal/progress/events.go`. Neo needs to populate these when emitting
`EventGraderStart` and `EventGraderComplete` in `engine_eval.go`.

### Engine wiring Neo needs to do

In `convertGraderResults()` (engine_eval.go):
```go
// Copy SourceFile/SourceType from graders.GraderResult → report.GraderResult
reportGrader.SourceFile = graderResult.SourceFile
reportGrader.SourceType = graderResult.SourceType
```

When emitting progress events:
```go
// In the EventGraderStart/EventGraderComplete emit call:
evt.GraderSourceFile = graderResult.SourceFile
evt.GraderSourceType = graderResult.SourceType
```

---

## Merge Order

If Neo lands first: Tank will rebase `tank/issue-grader-output-redesign` onto Neo's branch.
If Tank lands first: render degrades gracefully to flat/ungrouped until Neo's data lands.

---

## No Blockers

Tank's PR can be reviewed and merged independently. The 3-level format activates
automatically once Neo's SourceFile fields are populated.
