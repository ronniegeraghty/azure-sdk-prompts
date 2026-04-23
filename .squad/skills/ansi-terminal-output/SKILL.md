# ANSI Terminal Output & In-Place Rewrites

## Overview

Patterns and gotchas for rendering terminal output with ANSI escape sequences, in-place rewrites, and proper Unicode/wide-character handling. Covers:
- ANSI escape sequence basics
- In-place tail line updates (carriage return + clear line)
- Multi-row clearing for wrapped content
- Terminal cell width vs. rune count vs. byte length
- Wide character handling (emoji, CJK)

## ANSI Escape Sequences (Quick Reference)

```go
const (
    ansiCR        = "\r"           // Carriage return (move to column 0)
    ansiClearLine = "\x1b[2K"      // Clear entire line (cursor stays in place)
    ansiReset     = "\x1b[0m"      // Reset all attributes (color, bold, etc.)
    ansiRed       = "\x1b[31m"     // Red text
    ansiGreen     = "\x1b[32m"     // Green text
    ansiYellow    = "\x1b[33m"     // Yellow text
    ansiDim       = "\x1b[2m"      // Dim/faint text
    ansiBold      = "\x1b[1m"      // Bold text
)

// Cursor movement
fmt.Fprintf(w, "\x1b[%dA", n)  // Move cursor UP n lines
fmt.Fprintf(w, "\x1b[%dB", n)  // Move cursor DOWN n lines
```

**Key insight:** ANSI sequences have **zero terminal cell width**. They must be stripped when computing visible width for layout purposes.

## In-Place Tail Line Updates (Single Row)

**Use case:** Updating a status line without scrolling (e.g., "🔄 Running… turn 5/25 (00:12)")

**Pattern:**
```go
// First write: just print
fmt.Fprint(w, text)

// Updates: carriage return + clear line + new text
fmt.Fprintf(w, "\r\x1b[2K%s", newText)
```

**Limitation:** `\r\x1b[2K` only clears the **current physical row**. If `text` wraps to multiple rows, earlier rows are NOT cleared.

## Multi-Row Clearing (Wrapped Content)

**Problem:** When tail text exceeds terminal width, it wraps to multiple physical rows. On rewrite, `\r\x1b[2K` only clears the last row — earlier wrapped rows leak through.

**Solution:** Track how many rows the previous tail occupied, move cursor back to the first row, clear each row.

### Implementation

```go
type tailState struct {
    text     string // current tail text (for bookkeeping)
    rowCount int    // how many physical terminal rows this tail occupies
}

func writeTail(w io.Writer, text string, termWidth int, state *tailState) {
    // Truncate to terminal width (defense in depth — may still wrap if width changes)
    text = truncateToWidth(text, termWidth)
    
    // Write the new tail
    w.Write([]byte(text))
    
    // Track row count for next rewrite
    visible := visibleWidth(text)
    state.rowCount = (visible + termWidth - 1) / termWidth  // ceil division
    if state.rowCount < 1 {
        state.rowCount = 1
    }
    state.text = text
}

func rewriteTail(w io.Writer, newText string, termWidth int, state *tailState) {
    if state.rowCount == 0 {
        return // no previous tail
    }
    
    newText = truncateToWidth(newText, termWidth)
    
    var buf bytes.Buffer
    oldRows := state.rowCount
    
    // Move cursor UP to the first row of the tail (if multi-row)
    if oldRows > 1 {
        fmt.Fprintf(&buf, "\x1b[%dA", oldRows-1)
    }
    
    // Clear each row from top to bottom
    for i := 0; i < oldRows; i++ {
        buf.WriteString("\r")         // Move to column 0
        buf.WriteString("\x1b[2K")    // Clear line
        if i < oldRows-1 {
            buf.WriteString("\n")     // Move to next row (not on last row)
        }
    }
    
    // Write the new tail
    buf.WriteString(newText)
    w.Write(buf.Bytes())
    
    // Update row count
    visible := visibleWidth(newText)
    state.rowCount = (visible + termWidth - 1) / termWidth
    if state.rowCount < 1 {
        state.rowCount = 1
    }
    state.text = newText
}
```

### Key Insights

1. **Track row count AFTER writing:** Compute `ceil(visibleWidth / termWidth)` and store it for the next rewrite.
2. **Clear from top to bottom:** Move cursor UP to first row, then clear each row sequentially.
3. **Don't `\n` on the last row:** The cursor should end at column 0 of the last (bottom) row, ready for the new text.

## Terminal Cell Width vs. Rune Count

### The Problem

Go's `len(string)` returns **bytes** (UTF-8 encoding length).  
Go's `for range string` iterates **runes** (Unicode code points).  
But terminals count **cells** (display columns).

**Most characters:** 1 rune = 1 cell (ASCII, Latin, most symbols)  
**Wide characters:** 1 rune = **2 cells** (emoji, CJK ideographs, some symbols)  
**ANSI sequences:** N bytes = **0 cells**

### The Solution: `go-runewidth`

Use `github.com/mattn/go-runewidth` for proper terminal cell width calculation:

```go
import "github.com/mattn/go-runewidth"

// Strip ANSI sequences first (they have 0 width)
var ansiSeqRE = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

func visibleWidth(s string) int {
    stripped := ansiSeqRE.ReplaceAllString(s, "")
    return runewidth.StringWidth(stripped)
}
```

**Examples:**
```go
visibleWidth("hello")                       // 5 cells
visibleWidth("\x1b[31mhello\x1b[0m")       // 5 cells (ANSI codes stripped)
visibleWidth("🔄")                          // 2 cells (wide char)
visibleWidth("\x1b[32m✅ Loaded\x1b[0m")   // 9 cells (2 + 1 + 6)
```

### Truncation with Wide Characters

```go
func truncateToWidth(s string, max int) string {
    if max <= 0 {
        return s
    }
    
    visible := 0
    out := make([]byte, 0, len(s))
    i := 0
    
    for i < len(s) {
        // Match ANSI sequence at position i
        if loc := ansiSeqRE.FindStringIndex(s[i:]); loc != nil && loc[0] == 0 {
            out = append(out, s[i:i+loc[1]]...)
            i += loc[1]
            continue
        }
        
        // Decode the next rune
        r, size := utf8.DecodeRuneInString(s[i:])
        w := runewidth.RuneWidth(r)  // Get cell width (1 or 2)
        
        // Check if adding this rune would exceed the limit (leave 1 col for ellipsis)
        if visible+w > max-1 {
            out = append(out, []byte("…")...)
            out = append(out, []byte("\x1b[0m")...)  // Reset to avoid style bleed
            return string(out)
        }
        
        out = append(out, []byte(string(r))...)
        visible += w  // Accumulate CELL width, not rune count
        i += size
    }
    
    return string(out)
}
```

**Key points:**
- Use `runewidth.RuneWidth(r)` to get cell width per rune (1 or 2)
- Accumulate **cell width**, not rune count
- Leave 1 column for the ellipsis `…`
- Append `\x1b[0m` after truncation to reset styles (prevents bleed)

## Terminal Width Detection

Use `golang.org/x/term` to detect terminal width:

```go
import (
    "os"
    "strconv"
    "golang.org/x/term"
)

func TermWidth() int {
    // Try to get terminal size via stdout fd
    if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
        return w
    }
    
    // Fallback to COLUMNS env var
    if v := os.Getenv("COLUMNS"); v != "" {
        if n, err := strconv.Atoi(v); err == nil && n > 0 {
            return n
        }
    }
    
    // Hard fallback
    return 80
}
```

**When `TermWidth()` fails:**
- Stdout is piped/redirected (not a TTY)
- No `COLUMNS` env var set
- Falls back to 80 (or 120, depending on tool preference)

**Best practice:** Don't skip truncation when width detection fails — use the fallback width.

## NO_COLOR Support

Respect the `NO_COLOR` env var (see https://no-color.org/):

```go
func isColorEnabled() bool {
    // NO_COLOR env var set (even to empty string) disables color
    if _, exists := os.LookupEnv("NO_COLOR"); exists {
        return false
    }
    
    // Check if stdout is a TTY
    return isTerminal(os.Stdout)
}

func isTerminal(f *os.File) bool {
    fi, err := f.Stat()
    if err != nil {
        return false
    }
    return fi.Mode()&os.ModeCharDevice != 0
}
```

When color is disabled, strip ANSI codes before writing:
```go
if !isColorEnabled() {
    text = ansiSeqRE.ReplaceAllString(text, "")
}
```

## Common Pitfalls

1. **Using `len(string)` for width:** Returns bytes, not cells. WRONG.
2. **Using `for range` count for width:** Returns runes, not cells. WRONG for emoji/CJK.
3. **Forgetting to strip ANSI before width calculation:** ANSI codes have 0 width but count as bytes/runes.
4. **Only clearing one row with `\r\x1b[2K`:** Wrapped content requires multi-row clearing.
5. **Not tracking previous row count:** Can't clear the right number of rows without it.
6. **Not appending `\x1b[0m` after truncation:** Styles bleed past the ellipsis.

## Testing

### Unit Test Structure

```go
func TestVisibleWidth(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  int
    }{
        {"plain text", "hello", 5},
        {"ANSI codes", "\x1b[31mred\x1b[0m", 3},
        {"emoji", "🔄", 2},
        {"wide char + text", "✅ Loaded", 9}, // 2 + 1 + 6
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := visibleWidth(tt.input)
            if got != tt.want {
                t.Errorf("visibleWidth(%q) = %d, want %d", tt.input, got, tt.want)
            }
        })
    }
}
```

### Live Verification

Test with narrow terminal to verify wrapping behavior:
```bash
stty cols 60
go run . <command>
```

Observe that tail lines stay on one row (no leaked wrapped content).

## References

- **`golang.org/x/term`:** Terminal size detection
- **`github.com/mattn/go-runewidth`:** Proper cell width calculation for Unicode
- **ANSI escape codes:** https://en.wikipedia.org/wiki/ANSI_escape_code
- **NO_COLOR:** https://no-color.org/
- **hyoka implementation:** `hyoka/internal/progress/display_interactive.go`, commits `6b3d3d48` + `fe6efebf`

## When to Use This Skill

- Implementing CLI tools with live progress output (spinners, progress bars, status lines)
- Any in-place terminal updates (not line-buffered append-only output)
- Dealing with Unicode-heavy content (emoji, CJK languages)
- Ensuring clean rendering across different terminal widths and TTY vs. piped output

---

**Key Takeaway:** Terminal cell width ≠ rune count. Always use `runewidth.StringWidth()` for layout math when Unicode is involved. Track row counts for multi-row clearing.

## Isolating Terminal Output from Foreign Writes

### The Problem

When rendering interactive progress with in-place tail updates, ANY writes to the same TTY from other code paths (logging, error handlers, debug prints) will break the renderer's row-count tracking.

Example: renderer tracks "tail is at row N", but slog emits a warning to stderr (which renders to the same TTY), moving the cursor to row N+2. Renderer's next rewriteTail thinks cursor is at row N, clears the wrong rows.

### The Solution: Suppress or Redirect Foreign Writes

**Option 1: Suppress console output during interactive rendering**
```go
// In logging setup:
if interactiveModeActive && logFile == "" {
    handler = slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: level})
} else {
    handler = NewConsoleHandler(os.Stderr, level, styler)
}
```

**Option 2: Redirect logging through the renderer**
```go
// Pass the renderer's writer to the logger
handler = NewConsoleHandler(renderer.Writer(), level, styler)

// Renderer routes writes through its own tracking
func (r *Renderer) Write(p []byte) (n int, err error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.freezeTail() // commit tail before foreign write
    return r.w.Write(p)
}
```

**Option 3: Downgrade to append-only mode when logging is verbose**
```go
// Existing pattern in hyoka:
if mode == "interactive" && (logLevel == "debug" || logLevel == "info") && logFile == "" {
    mode = "ci" // downgrade to append-only, tolerates interleaved output
}
```

**Recommendation:** Use Option 1 (suppress) or Option 3 (downgrade) for simplicity. Option 2 (redirect) is complex and requires all foreign writes to go through the renderer.

### Terminal Width "Exactly Fits" Edge Case

**Problem:** When tail text visible width EXACTLY equals terminal width, cursor wrapping is terminal-dependent. Some terminals wrap immediately, others delay until the next write.

**Solution:** Always truncate to `termWidth - 2` (not `termWidth`) to leave a safety margin:

```go
func writeTail(w io.Writer, text string, termWidth int, state *tailState) {
    maxWidth := termWidth - 2
    if maxWidth < 10 {
        maxWidth = termWidth // skip margin for very narrow terminals
    }
    text = truncateToWidth(text, maxWidth)
    // ... rest of write logic
}
```

**Why 2 columns:**
- Wide chars (emoji) are 2 cells
- If truncation leaves a wide char at position `termWidth - 1`, it would occupy `termWidth - 1` and `termWidth`, hitting the edge
- 2-column margin ensures any single rune + ellipsis fits without ambiguity

---

**Key Takeaway:** Terminal renderers must OWN the output stream. Any code writing to the same TTY outside the renderer will corrupt tracking state. Either suppress, redirect, or downgrade to append-only mode.

## When to Give Up on Streaming UX (Decision Framework)

**Context:** hyoka Agent Attempt tail-streaming experienced four sequential fix attempts (commits 6b3d3d48, 42ea88fb, fe6efebf, 670c5dbf) addressing different facets of line-wrapping leaks (multi-row clearing, wide chars, foreign writes, terminal width edge cases). Despite each fix resolving a specific symptom, complete stability remained elusive. Final resolution: delete streaming, replace with three-state machine (commit b17f1ef5).

### Red Flags for Abandoning Streaming

1. **Multiple fix attempts on the same root cause**
   - If you've addressed cell width, truncation, multi-row clearing, and foreign writes and issues persist, the problem is fundamental — not fixable with better math.

2. **Complexity scales faster than value**
   - Streaming adds ~150 lines of logic (row tracking, truncation, ticker, per-event updates).
   - If the streaming content is "nice-to-have" (e.g., live duration counter, activity messages) rather than "essential" (e.g., "is it running?", "did it finish?"), the complexity isn't justified.

3. **Terminal environment is uncontrollable**
   - Foreign writers (slog stderr, other processes) inject lines between your tail updates.
   - Terminal width changes dynamically (user resizes window).
   - Wide character handling varies by terminal emulator.
   - If you can't control these, defensive truncation and margin logic become never-ending whack-a-mole.

4. **Test scenarios don't cover real-world variability**
   - Unit tests with fixed terminal widths pass, but real runs fail.
   - You can't easily reproduce edge cases (exact-width wrapping, emoji at truncation boundary, slog warning at exact moment of tail rewrite).

### Alternative: Bounded-State Displays

**Pattern:** Replace streaming variable-length content with a fixed set of short states.

**Example (hyoka Agent Attempt):**
- Old: `🔄 Running… turn X/Y, N tool calls (MM:SS)` (variable length, wraps on narrow terminals, rewritten per-second + per-event)
- New: `🔄 Running` | `✅ Completed` | `Guardrail hit — turn limit (25)` (bounded length, fits in any standard terminal, event-driven transitions)

**Benefits:**
- **Zero wrapping risk:** Content is short by design, no truncation math needed.
- **Simpler code:** Single-line state transitions vs. row tracking + per-event updates.
- **Future-proof:** If requirements change, you can add states without re-engineering the entire tail mechanism.

**Trade-offs:**
- Lose granular real-time feedback (activity messages, live counters).
- Gain reliability and maintainability.

### Decision Criteria

| Factor | Continue Streaming | Switch to Bounded States |
|--------|-------------------|--------------------------|
| **Essential info** | Live updates are semantically critical (e.g., "3/10 files processed") | States cover essential info (e.g., "Running" → "Completed") |
| **Fix attempts** | First or second attempt; clear path to resolution | Three or more attempts; root cause unclear |
| **Environment control** | Dedicated terminal, no foreign writes, predictable width | Shared TTY, foreign writers, dynamic width |
| **Complexity trend** | Fixes are additive, scoped | Fixes are multiplicative, cross-cutting |
| **Test coverage** | Edge cases reproducible in unit tests | Edge cases only emerge in production |

### Lessons from hyoka Agent Attempt

**What streaming cost:**
- Four separate fix commits addressing different facets of the same problem.
- Persistent UX bugs visible to users despite each fix passing tests.
- ~150 lines of truncation, row tracking, ticker, and per-event update logic.

**What bounded states bought:**
- Complete elimination of line-wrapping class of bugs (content is short → never wraps).
- Net deletion of ~100 lines (removed streaming logic, added state enum).
- Event-driven state transitions (no ticker needed for agent section).

**User-visible difference:**
- Lost: Live activity messages ("bash tool call", "writing file X"), live duration counter, real-time tool call count.
- Kept: Essential status ("Running" vs "Completed"), guardrail failure reasons, overall outcome.
- User feedback: Preferred loss of detail over persistent UX bug.

### When Streaming Is Still Worth It

- **Long-running operations** where users need granular progress (e.g., "downloading 3/10 files (25 MB/100 MB)").
- **Intermediate state changes are semantically important** (e.g., "compiling…" → "linking…" → "packaging…").
- **Controlled terminal environment** (dedicated CLI app, no foreign writers, you own the TTY).
- **Value >> complexity** (users explicitly request live updates, no other way to convey progress).

### Recommended Approach

1. **Start with bounded states** for new progress displays.
2. **Only add streaming** if users report that static states are insufficient.
3. **Stop after two fix attempts** if streaming proves fragile — revert to states.
4. **Document the trade-off** in team decisions so future contributors know why streaming was avoided.

**Golden rule:** If you're writing `truncateToWidth()`, `visibleWidth()`, and `tailRowCount` for the third time, you've already lost. Switch to states.

