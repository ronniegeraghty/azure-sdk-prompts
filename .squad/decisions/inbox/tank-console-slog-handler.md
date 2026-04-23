# Console-Friendly slog Handler

**Agent:** Tank  
**Date:** 2026-04-23  
**Branch:** ronniegeraghty/dev  
**Commit:** 82fc9750

## Problem

Ronnie ran a `--pairwise` eval and saw warnings printing to stdout in structured slog format:

```
time=2026-04-23T00:52:25.233Z level=WARN msg="Plugin not found, skipping" plugin=azure-sdk-java@skills config=baseline-skills/claude-sonnet-4.5 hint="Install with: /plugin install azure-sdk-java@skills"
```

The `time=... level=... msg=...` format is diagnostic-oriented and clutters the console. He wanted:
- **Console (stdout/stderr)**: human-friendly output like `⚠️  Plugin not found, skipping (plugin=azure-sdk-java@skills, ...)` with short, dim attrs.
- **Log file**: keep structured format for diagnostics.

## Solution

Implemented a custom `slog.Handler` (ConsoleHandler) used when log output goes to stdout/stderr:

### ConsoleHandler Behavior
- **WARN**: `⚠️  <msg> (key=val key=val ...)` — attrs in dim
- **ERROR**: `❌ <msg>` in red, attrs in dim
- **INFO/DEBUG**: suppressed entirely on console (diagnostic noise — renderer carries user-facing progress)
- Respects NO_COLOR and TTY detection via `progress/style` package

### logging.Setup() Logic
```go
if opts.FilePath != "" {
    // File destination: structured TextHandler with timestamps
    handler = slog.NewTextHandler(f, &slog.HandlerOptions{Level: level})
} else {
    // Console destination: human-friendly ConsoleHandler
    styler := style.New(os.Stderr)
    handler = NewConsoleHandler(os.Stderr, level, styler)
}
```

### Tests
Added `console_handler_test.go` with 9 table-driven tests:
- Enabled level checks (DEBUG/INFO suppressed, WARN/ERROR enabled)
- Formatting with/without attrs, with/without colors
- NO_COLOR environment variable support
- Handler-level attrs (via WithAttrs)
- WithGroup (no-op for console output)

All tests pass with `-race`.

## Impact

- **Console output**: now human-readable with emoji, colored errors, dim attrs
- **Log files**: still structured with timestamps for diagnostics
- **Backward compatible**: existing --log-level, --log-file flags work unchanged
- **NO_COLOR aware**: respects NO_COLOR=1 environment variable

## Files Changed

- `hyoka/internal/logging/console_handler.go` (new)
- `hyoka/internal/logging/console_handler_test.go` (new)
- `hyoka/internal/logging/logging.go` (updated Setup function)

## Verification

Manual test with real CLI:
```bash
# Console warnings are human-friendly
./hyoka list --config nonexistent 2>&1 | head -3
⚠️  Plugin not found, skipping (plugin=azure-sdk-java@skills config=baseline-skills/claude-sonnet-4.5 hint=Install with: /plugin install azure-sdk-java@skills)

# Log file output is structured
./hyoka list --log-file test.log --config nonexistent
cat test.log | head -1
time=2026-04-23T01:16:00.388Z level=WARN msg="Plugin not found, skipping" plugin=azure-sdk-java@skills config=baseline-skills/claude-sonnet-4.5 hint="Install with: /plugin install azure-sdk-java@skills"

# NO_COLOR is respected
NO_COLOR=1 ./hyoka list --config nonexistent 2>&1 | head -1
⚠️  Plugin not found, skipping (plugin=azure-sdk-java@skills config=baseline-skills/claude-sonnet-4.5 hint=Install with: /plugin install azure-sdk-java@skills)
# (no ANSI codes in output)
```

## Design Decisions

1. **Handler selection at Setup time**: Decided based on whether `--log-file` is set. Keeps the decision in one place, no runtime handler swapping.

2. **INFO/DEBUG suppression on console**: The interactive renderer already provides user-facing progress. INFO/DEBUG logs are diagnostic noise on stdout/stderr. They still go to log files when `--log-file` is set.

3. **Style package reuse**: Used existing `progress/style` package for color/dim helpers and NO_COLOR detection. No new dependencies.

4. **Attrs always dim**: Attributes are secondary context — always rendered dim regardless of level. Only ERROR messages get red text, attrs stay dim.

5. **Groups ignored in console output**: slog groups are for structured hierarchy; console output is flat. WithGroup is a no-op (returns handler with groups tracked but not rendered).

## Future Considerations

- If we ever want dual output (both console and file), we could implement a multi-handler that dispatches each record to both handlers. Not needed now.

- If we want to customize emoji or colors, ConsoleHandler could accept a config struct. Current hardcoded values (⚠️, ❌, red, dim) are sensible defaults.

## Related

- Issue: Ronnie's directive from CLI output UX sprint round 2
- Pattern: Similar to how progress renderers choose display mode (interactive vs CI vs off)
- Dependency: `progress/style` package for ANSI codes and NO_COLOR handling
