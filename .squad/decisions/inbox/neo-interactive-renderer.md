# Interactive Progress Renderer — Architecture

**Author:** Neo
**Branch:** ronniegeraghty/dev
**Scope:** `hyoka/internal/progress/display_interactive.go` (new), minimal
dispatch edits to `display.go`, `cmd/run.go` auto-mode update.

## Summary

New renderer for single-eval, human-watched runs. Only the TAIL line is
updatable; older lines are immutable. One sanctioned exception: the
`EventToolsVerified` bulk event can redraw the Tools block using DECSC/DECRC
save-restore. All other sections are append-only event-driven, with a
1-second ticker refreshing only the Agent Attempt tail line's duration
counter.

## Mode wiring

- New `ProgressMode` constant `ModeInteractive = "interactive"`.
- `NewDisplay` dispatches to `newInteractiveRenderer` when Mode is
  interactive, mirroring Trinity's CI-renderer delegation pattern.
- `cmd/run.go` auto-mode now selects `"interactive"` for `workers==1` (was
  `"live"`) and `"ci"` for `workers>1` (was `"log"`). Explicit
  `--progress live|log|ci|off` still overrides.
- Debug/info log level without `--log-file` still downgrades `interactive` →
  `ci` to keep stderr slog output from corrupting cursor moves.

## Per-eval layout

```
Prompt: <id>
Config: <name>
Tools:                                  (header + block omitted if zero tools)
  - <name> (<kind>): 🔄 Loading…        (tail; flips in-place)
  - <name> (<kind>): ✅ Loaded
  - <name> (<kind>): ❌ Failed (reason)
Agent Attempt:
  🔄 <activity> · N tool calls  (MM:SS) (tail; per-second refresh)
  ✅ Complete — N files, M tool calls  (MM:SS)
Session Details:
  Files: a, b, … (N more)
  Turns: …   Tool calls: …   Cost: $X.XX
Graders:
  - <id> (<kind>): ✅ Pass (8/10)
  - <id> (<kind>): 🔄 Running…          (tail; next grader appends fresh line)
```

## Tail-update protocol

- `writeLine(text)` — commit `text + "\n"`, increment `linesWritten`. Used
  for header rows (Prompt, Config, "Tools:", "Agent Attempt:", etc.) and for
  non-tail section rows after their status is finalized.
- `writeTail(kind, text)` — freezes any existing tail, writes text without
  newline. Records `tailKind` so ticker knows whether to refresh.
- `rewriteTail(text)` — emits `"\r\x1b[2K" + text`. Same physical row; no
  change to `linesWritten`.
- `freezeTail()` — writes `"\n"`, increments `linesWritten`, clears
  `tailKind`. Tail content becomes immutable.

## The one exception: `redrawToolsBlock`

Triggered only from `onToolsVerified` when at least one tool status flips.
Sequence:

```
\x1b7                  (DECSC: save cursor)
\x1b[<N>A  \r          (move up N lines, column 0)
for each tool line:
  \x1b[2K + new text + "\n"
\x1b8                  (DECRC: restore cursor)
```

`N = linesWritten - toolsFirstLine`. Lines between the tool block and the
current tail are untouched by the redraw because we write exactly the same
number of lines as the original tool block occupies. The cursor restore
puts us back on the original tail line. `toolsVerified` flag guards against
double redraws.

## Multi-eval handling

Interactive mode is selected only when `workers==1`, so there is at most one
eval in flight. Queued evals are rendered sequentially: on receiving an
`EventStarting` with a new `EvalID`, the current eval's tail is frozen, a
blank separator line is printed, and a fresh per-eval block begins.

## Terminal events

- `EventPassed` — calls `agentComplete(..., true)`, increments counters,
  finalizes block.
- `EventFailed` — calls `agentComplete(..., false)`, appends a `❌` reason
  line, finalizes.
- `EventError` — same as Failed path but increments the `errors` counter.

## Color / style

All styled output goes through `style.Styler` (Trinity's package):
`sty.OK/Fail/Muted/Info`. `style.New(writer)` honors `NO_COLOR` and non-TTY
writers — tests capture into a `bytes.Buffer` and receive raw text.

## Counters

`Display.completed/passed/failed/errors` are incremented in the dispatch
shim so `CompletedEvalCount()` remains correct across all renderer modes.
The final Summary line is written by `interactiveRenderer.finish()`.

## Test seams for Switch

- Construct with `Writer: &bytes.Buffer{}` — disables color, disables ticker
  effects on output (ticker still runs but buffer doesn't show the tail
  updates as animation, only as final state).
- Key substrings to assert in golden tests: section headers (`"Tools:"`,
  `"Agent Attempt:"`, `"Session Details:"`, `"Graders:"`), state markers
  (`"✅ Loaded"`, `"❌ Failed"`, `"✅ Pass"`, `"✅ Complete"`), and
  `"Summary: X/Y passed"`.
- Omission tests: no `"Tools:"` if no resolution events; no `"Graders:"` if
  no grader events.
- Tool flip test: send ResolutionResult=loaded, then ToolsVerified with
  status=failed → assert failure reason appears in output.

## Files touched

- NEW `hyoka/internal/progress/display_interactive.go`
- NEW `hyoka/internal/progress/display_interactive_test.go` (3 tests)
- MOD `hyoka/internal/progress/display.go` — add `ModeInteractive` constant,
  `interactive` field on Display, dispatch in `NewDisplay`, `HandleEvent`,
  `Finish`.
- MOD `hyoka/cmd/run.go` — auto-mode picks `"interactive"` / `"ci"`; flag
  help text updated.

## Verified

- `go build ./...` ✅
- `go test -race -count=1 ./hyoka/internal/progress/...` ✅
- `go test -race -count=1 ./hyoka/...` ✅
- `go vet ./hyoka/...` ✅
