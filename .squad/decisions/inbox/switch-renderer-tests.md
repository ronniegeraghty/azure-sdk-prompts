# Renderer Snapshot Test Coverage

**Author:** Switch 🤍
**Branch:** ronniegeraghty/dev
**Scope:** `hyoka/internal/progress/display_interactive_test.go` (extended),
`hyoka/internal/progress/display_ci_test.go` (new).

## Summary

Both progress renderers (interactive and CI) now have table-driven snapshot
tests covering the scenarios in the `tests-renderer-snapshots` task spec.
The interactive renderer's animation-heavy output is tested via a mix of
substring assertions and raw-ANSI-escape assertions (tail updates,
DECSC/DECRC redraw). The CI renderer is tested via substring matrices plus
one full-output snapshot with timestamp/duration normalization.

## Cases covered

### Interactive (`display_interactive_test.go`)

All six scenarios from the spec, plus dedicated ANSI-marker tests:

1. **`TestInteractive_Cases/happy_path_one_tool_two_graders`** — 1 tool loaded,
   agent completes, 2 graders pass.
2. **`TestInteractive_Cases/tool_load_failure_at_resolution`** — one tool
   emits Failed at resolution time.
3. **`TestInteractive_Cases/tools_verified_flip_loaded_to_failed`** —
   `ToolsVerified` flips a previously-Loaded tool to Failed.
4. **`TestInteractive_Cases/grader_fail_one_pass_one_fail`** — one grader
   passes, one fails (with reason surfaced).
5. **`TestInteractive_Cases/error_path_generation_error`** — `EventError`
   path: agent freezes as Failed, message surfaced, `1 errors` in summary.
6. **`TestInteractive_NoColorEnvDropsColor`** — `NO_COLOR=1` via `t.Setenv`:
   no SGR color codes (`\x1b[31m` etc.) emitted, plain text + glyphs remain.
7. **`TestInteractive_ANSIMarkers`** — (a) tail-update escape `\r\x1b[2K`
   appears on tool resolution flip; (b) DECSC `\x1b7` + DECRC `\x1b8` bracket
   the tools-block redraw, with save preceding restore.

Plus Neo's pre-existing three tests (`TestInteractive_HappyPath`,
`TestInteractive_NoToolsNoGraders`, `TestInteractive_ToolsVerifiedFlip`) are
retained unchanged.

### CI (`display_ci_test.go` — new file)

All five scenarios from the spec, plus one pinned full snapshot:

1. **`TestCIRenderer_Cases/happy_path_three_evals_all_pass`** — 3 evals all
   pass, summary table renders, `3/3 passed · report: …` footer.
2. **`TestCIRenderer_Cases/mixed_two_pass_one_fail_with_reason`** — 2 pass +
   1 fail; failure reason is a multi-line message, confirmed collapsed to
   one line by `oneLine()`.
3. **`TestCIRenderer_Cases/multi_eval_interleaved_graders`** — grader
   `Start`/`Complete` events from two evals interleave; each eval's
   `graderPass`/`graderTotal` counts are isolated (`2/2 graders` vs
   `1/2 graders`).
4. **`TestCIRenderer_Cases/no_color_drops_emoji_keeps_box_borders`** — buffer
   writer disables color; asserts `START`/`PASS` text (not `▶ start` /
   `✅ pass`) and that unicode box borders (`┌│└`) survive regardless.
5. **`TestCIRenderer_Cases/zero_evals_empty_summary_does_not_crash`** — zero
   evals reach the renderer; `writeTable` skipped (no box chars), summary
   header still printed, `0/0 passed` footer, no panic.
6. **`TestCIRenderer_HappyPathSnapshot`** — full-output golden for a 2-eval
   run (1 pass + 1 fail) after normalizing timestamps and durations. Locks
   intro line, start/finish lines, column widths, border glyphs, and footer
   shape.

## Test infrastructure

- `normalizeCI(s)` — regex-based stripper for `[HH:MM:SS]` timestamps,
  `(Ns, G/T graders)` durations, and table Duration-column cells. Three
  compiled regexes, placeholders `[HH:MM:SS]` / `DUR`.
- `feedInteractive(buf, events)` / `feedCI(buf, ...)` — one-line helpers
  that build a `Display` pointed at `buf`, drive the event stream, and call
  `Finish()`. Removes ~8 lines of boilerplate per case.
- `floatPtr(v)` — convenience for embedding `*float64` grader scores in
  table literals.

## What these tests protect

- Interactive: section headers (`Tools:`, `Agent Attempt:`, `Session
  Details:`, `Graders:`), tool status transitions, DECSC/DECRC bracket for
  block redraw, tail-update escape sequence, NO_COLOR compliance.
- CI: intro-line phrasing, event→line mapping per the Trinity memo, grader
  attribution across interleaved events, box-table rendering, NO_COLOR
  glyph/color suppression, empty-summary safety, full column layout.

## No testdata/ directory needed

Inline string literals with regex-normalized timestamps proved sufficient.
No `testdata/progress/` fixtures were added. If future tests need larger
golden outputs (e.g., 10-eval runs), drop them in
`hyoka/internal/progress/testdata/` per convention.

## Suggested doc mentions for Oracle

- "Renderer snapshot coverage" subsection under Testing in a future
  `docs/progress.md` or `docs/architecture.md` update — mention both
  `display_interactive_test.go` and `display_ci_test.go` as the canonical
  reference for the renderer contract.
- Note for contributors: when editing either renderer, expect the
  `TestCIRenderer_HappyPathSnapshot` golden to diff — update it only when
  the layout change is intentional.

— Switch 🤍
