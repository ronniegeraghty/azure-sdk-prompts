# CI Progress Renderer — Modes, Event Mapping, Summary Table

**Author:** Trinity 🖤
**Branch:** ronniegeraghty/dev
**Scope:** `hyoka/internal/progress/display_ci.go` (new, ~320 LOC) +
mode-dispatch edits in `hyoka/internal/progress/display.go` + one-line
config-count plumbing in `hyoka/internal/eval/engine.go`.

## Summary

Replaces the legacy `log` mode renderer (which printed verbose per-phase
lines like `Prompt: … / generating…`) with a proper CI-oriented view: one
timestamped start line + one timestamped finish line per eval, followed by
an end-of-run summary table. Designed for redirected output / CI logs — no
ANSI cursor movement, every byte stays on the page.

## Mode strings

| Mode | Dispatch | Notes |
|---|---|---|
| `"ci"` (new, preferred) | CI renderer (`display_ci.go`) | The explicit, documented name. |
| `"log"` (legacy) | CI renderer | Kept as an alias — `--progress log` callers see the new output with no flag change. The old per-phase/inline behavior is gone. |
| `"interactive"` | Interactive renderer (Neo's) | Sibling dispatch case in the same `switch`. |
| `"live"` | Legacy ANSI region redraw | Unchanged. |
| `"auto"` / `""` | TTY → live, else disabled | Unchanged. Note: `--progress auto` selection logic in `cmd/run.go` may later swap between `"ci"` and `"interactive"` based on worker count (out of scope here). |
| `"off"` | disabled | Unchanged. |

## Event → output-line mapping (during-run, append-only)

| Event | Line emitted | Tracked state |
|---|---|---|
| `EventStarting` | `[HH:MM:SS] ▶ start  <promptID>  \|  <configName>` | Creates eval state; records `startTime`, appends `evalID` to `order` slice for summary row ordering. |
| `EventGraderStart` | (no line) | Increments `graderTotal` on the eval's state. |
| `EventGraderComplete` | (no line) | If `Result == GraderResultPass`, increments `graderPass`. |
| `EventPassed` | `[HH:MM:SS] ✅ pass  <promptID>  \|  <configName>  (<dur>, G/T graders)` | Marks eval `result="pass"`, finalizes `duration`. |
| `EventFailed` | `[HH:MM:SS] ❌ fail  <promptID>  \|  <configName>  (<dur>, G/T graders) — <reason>` | Marks eval `result="fail"`; reason from `evt.Message`, collapsed to one line, defaults to `"graders failed"` when empty. |
| `EventError` | same as `EventFailed` but reason defaults to `"eval errored"` | Marks eval `result="error"` (rendered as `FAIL` in the summary table — we don't distinguish the two visually, they're both red). |
| all other events (`EventReasoning`, `EventToolStart/Complete`, `EventWritingFile`, `EventSendingPrompt`, `EventPhaseChange`, `EventWaiting`, `EventSessionDetails`, `EventToolResolutionStart/Result`, `EventToolsVerified`) | (no line) | CI mode is deliberately quiet — per-event chatter belongs in interactive / debug logs. |

### Event attribution for parallel evals

Grader events in CI mode arrive **interleaved** across evals (grader
serialization only applies to interactive mode per Neo's
`grader-serialization` memo). All state lookup is keyed by `evt.EvalID`, so
interleaving is safe — each eval's `graderPass`/`graderTotal` counters are
isolated.

## Glyphs & NO_COLOR strategy

Tied to the `style.Styler.Enabled` flag (NO_COLOR or non-TTY disables both):

| State | Start | Pass | Fail | Separator | Timestamp |
|---|---|---|---|---|---|
| colors enabled | `▶ start` (cyan) | `✅ pass ` (green) | `❌ fail ` (red) | `\|` (dim) | `[HH:MM:SS]` (dim) |
| NO_COLOR / piped | `START` | `PASS ` | `FAIL ` | `\|` | `[HH:MM:SS]` |

Box-drawing characters (`┌ ─ ┐ │ ┼ ├` etc.) are used **unconditionally** —
they're valid UTF-8 and render correctly in every CI log viewer I care
about (GitHub Actions, Datadog, Splunk, plain `less`). If a future consumer
needs ASCII-only borders, add a flag; don't gate on NO_COLOR.

## Timestamp format

`[HH:MM:SS]` relative to renderer construction (`r.startTime`), computed as
`time.Since(r.startTime)` rounded down. Captures elapsed wall time cleanly
across runs that span hours. Styled with `Muted` (dim) when colors are on.

## Summary table

Rendered at `Finish()` after a blank line and a bold `Summary` header.

### Row ordering

Rows appear in **first-seen eval order** — the order their `EventStarting`
arrived. Matches the chronological order of the start lines above the table,
so the table reads top-down as a replay of the run. (Implementation uses an
`order []string` slice keyed on `evalID`; a formal `sort.SliceStable` with a
no-op less-func documents the "preserve order" intent.)

### Column sizing

Auto-sized to the max of header length and all cell contents using plain
`len()`. Rationale: content is ASCII-ish (prompt IDs, config names, "PASS"/
"FAIL", "N/M", "NNs"). If CJK or wide-char content ever enters the table,
swap in `runewidth` — the change is localized to `writeTable` +
`writeBorder` / `writeRow`.

Two-pass algorithm:
1. Walk headers to seed `widths[i] = len(headers[i])`.
2. Walk all rows, updating `widths[i] = max(widths[i], len(row[i]))`.
3. Render top/header/mid/rows/bottom using unicode box-drawing characters.

Padding: each cell is rendered as `" " + cell + (width-len(cell)) spaces + " "`
— one space left/right margin. Headers get the same padding shape but
wrapped in `Styler.Bold` when colors are enabled. Result column contains
literal `PASS`/`FAIL` text (no emoji, no ANSI) so snapshot tests don't
carry styling noise.

### Footer

```
N/M passed · report: <reportDir>
```

`N` = count of eval states with `result == "pass"`. `M` = `len(order)` (total
evals that reached either start or finish). `report:` is omitted when
`reportDir` is empty. Footer is always plain text — no color, no emoji.

## Snapshot test guidance for Switch

1. **Force `Styler` disabled** — pass `Writer: &bytes.Buffer{}` to
   `NewDisplay`. Buffer is not a `*os.File`, so `style.New` returns disabled,
   so both ANSI codes and emoji drop out. Golden files stay portable.
2. **Elapsed timestamps are not stable** — either assert substring matches
   (`"START  "`, `"PASS  "`, `"Summary"`, `"1/2 passed"`) rather than
   full-line golden files, or stub `r.startTime` via a test hook (not
   currently exposed; happy to add one on request).
3. **Grader counts are deterministic** — emit N `EventGraderStart` then N
   `EventGraderComplete` with known `Result`s; the `(P/T graders)` suffix
   is reliable.
4. **Event order matters** — `EventPassed`/`EventFailed` should always come
   *after* all grader events for that eval, same as the engine wires it.
   Re-ordering would lose grader-tally attribution.

## Files touched

- `hyoka/internal/progress/display_ci.go` — **new** (~320 LOC)
- `hyoka/internal/progress/display.go` — `ModeCI` const, `Configs` field,
  `ci *ciRenderer` field, dispatch in `NewDisplay` / `HandleEvent` /
  `Finish`.
- `hyoka/internal/progress/display_test.go` — rewrote `TestDisplay_LogMode`
  against new format, added `TestDisplay_CIMode`.
- `hyoka/internal/eval/engine.go` — compute `len(uniqueConfigs)` from
  `tasks[].Config.Name` and pass as `DisplayConfig.Configs`.

## Verification

- `go build ./...` — clean.
- `go test -race ./hyoka/internal/progress/...` — pass.
- `go vet ./hyoka/...` — clean.
- Manual demo (throwaway test): output matches the sprint plan's CI layout
  block exactly (intro line, start/pass/fail lines with timestamps, summary
  table with unicode borders, `N/M passed · report: …` footer).

— Trinity 🖤
