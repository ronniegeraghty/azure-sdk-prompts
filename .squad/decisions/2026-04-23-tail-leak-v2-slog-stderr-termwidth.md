# Decision: Interactive Progress Tail Leak v2 — Root Cause & Fix

**Date:** 2026-04-23  
**Author:** Tank 📡  
**Status:** Implemented (commit 670c5dbf)  
**Affects:** CLI progress display (interactive mode)

## Context

The interactive progress renderer's multi-line tail clearing (introduced in fe6efebf for wide-char emoji support) was still exhibiting visual artifacts — old tail fragments leaking through on screen during `hyoka run` with interactive mode.

The user reported this issue persisted even after the runewidth-based cell counting fix. Investigation revealed TWO distinct root causes, not addressed by the previous fix.

## Root Causes

### 1. Foreign Writes from slog (Primary Cause)

**The Problem:**
- Interactive renderer writes to `os.Stdout` via `Display.w`
- slog ConsoleHandler (when no `--log-file`) writes to `os.Stderr`  
- Both stdout and stderr render to the same TTY screen
- Renderer tracks `tailRowCount` based on writes through `Display.w` only
- When slog emits warnings (e.g., plugin not found, process cleanup errors) BETWEEN tail writes, the terminal scrolls without the renderer knowing
- Result: `rewriteTail` tries to clear N rows, but cursor has moved M lines down → clears wrong rows, leaving old tail visible

**Example sequence:**
```
1. writeTail("🔄 Running… turn 3/25")   → cursor at row 10, tailRowCount=1
2. [slog.Warn emits 2 lines to stderr]  → cursor now at row 12  
3. rewriteTail("🔄 Running… turn 4/25") → moves up 1 row (thinks cursor at row 10+1=11), 
                                           clears row 11, writes at row 11
                                           BUT actual cursor was at row 12!
                                           Result: old tail at row 10 NOT cleared
```

**Why this wasn't caught earlier:**
- The previous fix (fe6efebf) addressed emoji cell width, not foreign writes
- Testing with `--log-file` specified → no console output → bug hidden
- Testing without triggering warnings → bug hidden
- Real-world runs often have warnings (missing plugins, model failures, cleanup notices)

### 2. Terminal Width Edge Case (Secondary Cause)

**The Problem:**
When tail text visible width EXACTLY equals terminal width, cursor wrapping behavior is terminal-dependent:
- **Some terminals:** Cursor wraps immediately to column 0 of next row after writing the Nth character
- **Other terminals:** Cursor stays at column N until the NEXT write, then wraps

This ambiguity means `tailRowCount` calculation can be off-by-one when `visibleWidth == termWidth`.

**Example:**
```
Terminal: 80 columns
Tail: "Status: 🔄 Running… turn 3/25, 10 tool calls  (00:12)"  → exactly 80 cells

Terminal A behavior:
  - After write: cursor at column 0 of row 2 (wrapped immediately)
  - tailRowCount=1 is WRONG (should be 2)
  
Terminal B behavior:
  - After write: cursor at column 80 of row 1 (delayed wrap)
  - tailRowCount=1 is CORRECT
```

Result: On rewriteTail, Terminal A clears 1 row (wrong), Terminal B clears 1 row (correct). Leak on Terminal A.

## Fix Applied

### 1. Suppress Console Logging During Interactive Mode

**Implementation:**
- Added `SuppressConsole bool` field to `logging.Options`
- When `SuppressConsole == true`, logging.Setup() creates a `slog.NewTextHandler(io.Discard, ...)` instead of ConsoleHandler
- In `run.go`, after resolving progress mode: if `mode == "interactive" && logFile == ""`, reconfigure logger with `SuppressConsole: true`

**Effect:**
- All slog.Warn/Error calls during the run write to io.Discard (no stderr output)
- When `--log-file` IS specified, warnings go to the file (unaffected)
- When progress mode is NOT interactive (ci/off), console warnings work normally (unaffected)

**Trade-off:**
- Users running `--progress interactive` without `--log-file` won't see warnings in real-time
- They can still use `--log-file` to capture warnings, or use `--progress ci` for append-only mode that tolerates interleaved output
- The existing downgrade logic (interactive→ci when logLevel=debug/info without --log-file) already established this pattern; we extended it to all log levels

### 2. Safety Margin on Tail Truncation

**Implementation:**
- Modified `writeTail` and `rewriteTail` to truncate to `(termWidth - 2)` instead of `termWidth`
- For very narrow terminals (< 12 cols), skip the margin (use full width)

**Effect:**
- Tail text never exactly equals terminal width → avoids cursor wrapping ambiguity
- 2-column buffer ensures cursor always stays on the same row after write
- Minimal visual impact (2 columns lost on typical 80-120 col terminals)

**Why 2 columns:**
- 1 column would be enough for ASCII-only, but wide chars (emoji) are 2 cells
- If the last rune before truncation is wide (2 cells) and we have 1-column margin, truncation might still hit exactly termWidth
- 2 columns guarantees room for any single rune + ellipsis

## Verification

1. **Tests:** All `go test -race ./hyoka/internal/progress/...` pass (including truncation, visibleWidth, interactive renderer snapshot tests)
2. **Manual repro:** Ran `hyoka run --prompt-id <id> --config <cfg> --log-level warn` (no --log-file) → warnings appeared BEFORE interactive mode started (during config load), but NO warnings during the run itself (suppression working)
3. **Narrow terminal:** `stty cols 60 && hyoka run ...` → no visual artifacts
4. **Cleanup:** `hyoka clean` terminated 6 orphaned processes from test runs

## Gaps / User Testing Required

**I could NOT fully verify the fix end-to-end because:**
- My test runs either timed out or didn't trigger enough warnings DURING the interactive phase (most warnings occur during config load, before renderer starts)
- The issue is **intermittent** and depends on timing (when slog emits warnings relative to tail rewrites)
- I don't have a way to force warnings during the generation/grading phase without modifying eval code

**User should:**
1. Run `hyoka run --prompt-id key-vault-dp-python-crud --config "baseline/claude-opus-4.6"` (no --log-file, no --log-level)
2. Watch for tail fragments leaking through during Agent Attempt / Graders sections
3. Try multiple prompts (vary length, emoji presence)
4. Try `stty cols 60`, `stty cols 80`, `stty cols 120` and repeat
5. If leak still occurs: check if warnings are appearing during the run (they shouldn't be with my fix)

## Decision

**Adopt this fix as the final resolution for the tail leak issue.**

If the leak persists after this fix, the next investigation should focus on:
- Race conditions in ticker goroutine (though mutex is present)
- Writes from eval/criteria code that bypass the renderer
- Terminal emulator bugs (some terminals may not handle ANSI cursor-up correctly)

## Related

- **fe6efebf**: Wide char fix (runewidth cell counting)
- **6b3d3d48**: Multi-row clearing base implementation
- **42ea88fb**: Truncation base implementation
- **047fed6b**: Scribe merge documenting tail leak postmortem (this decision supersedes that)
