# Terminal Cell Width vs. Rune Count for Tail Line Truncation

**Status:** ✅ IMPLEMENTED  
**Date:** 2026-04-23  
**Scope:** `hyoka/internal/progress/display_interactive.go`  
**Supersedes:** `trinity-tail-truncation-final.md` (partially incorrect — fixed here)  
**Commits:** `6b3d3d48` (added truncation + multi-row clear structure), `fe6efebf` (fixed cell width calculation)

## Context

The interactive eval progress renderer uses an in-place rewrite pattern: a "tail line" is updated via `\r\033[2K` (carriage return + clear line) + multi-row clearing. When the tail text exceeds terminal width, it wraps to multiple physical rows. The renderer must:
1. Truncate NEW tail text to fit terminal width
2. Track how many rows the PREVIOUS tail occupied
3. Clear ALL previous rows before writing the new tail

Commit `6b3d3d48` implemented this structure (multi-row clear, `TermWidth()` detection, `truncateToWidth()` helper) but used **rune counting** for width calculation. This was **wrong** for wide characters.

## The Bug (Why `6b3d3d48` Didn't Work)

### Rune vs. Cell Width

- **Rune:** Go's iteration unit (one per Unicode code point). `for range string` counts runes.
- **Terminal cell (column):** What the terminal counts for layout. Most runes = 1 cell, but:
  - Emoji (🔄, ✅, ❌) = **2 cells**
  - CJK ideographs = **2 cells**
  - ANSI escape sequences = **0 cells**

### The Buggy Code (from `6b3d3d48`)

```go
func visibleWidth(s string) int {
    stripped := ansiSeqRE.ReplaceAllString(s, "")
    count := 0
    for range stripped {
        count++  // ❌ Counts RUNES, not CELLS
    }
    return count
}

func truncateToWidth(s string, max int) string {
    // ...
    if visible+1 > max-1 { // visible = rune count, not cell count
        out.append([]byte("…")...)
        out.append([]byte("\x1b[0m")...)
        return string(out)
    }
    // ...
}
```

### What Happened

Example: tail line `"🔄 edit → main.py · 3 tool calls   (00:12)"`

- **Rune count:** 44 runes (🔄 = 1 rune)
- **Actual cell width:** 45 cells (🔄 = 2 cells)
- Terminal width: 45 cols
- `truncateToWidth(text, 45)` → no truncation (44 runes < 45 max)
- Actual render: 45 cells → fits EXACTLY on one row
- BUT if one more character added: 45 runes, 46 cells → wraps to 2 rows
- `tailRowCount = (45 + 45 - 1) / 45 = 1` (based on rune count)
- Next rewrite: clears 1 row, but the line occupied 2 → **leaked second row**

## The Fix (Commit `fe6efebf`)

**Added dependency:** `github.com/mattn/go-runewidth`

### Updated `visibleWidth()`

```go
func visibleWidth(s string) int {
    stripped := ansiSeqRE.ReplaceAllString(s, "")
    return runewidth.StringWidth(stripped)  // ✅ Proper cell width
}
```

`runewidth.StringWidth()` accounts for:
- Wide characters (emoji, CJK) = 2 cells
- ANSI sequences = 0 cells (stripped first)
- ASCII/Latin = 1 cell per rune

### Updated `truncateToWidth()`

```go
func truncateToWidth(s string, max int) string {
    // ...
    r, size := decodeRune(s[i:])
    w := runewidth.RuneWidth(r)  // ✅ Get cell width per rune
    if visible+w > max-1 {       // Check CELL width, not rune count
        out = append(out, []byte("…")...)
        out = append(out, []byte("\x1b[0m")...)
        return string(out)
    }
    out = append(out, []byte(string(r))...)
    visible += w  // ✅ Accumulate CELL width
    i += size
    // ...
}
```

### Why This Works

- `runewidth.RuneWidth('🔄')` → `2` (correct!)
- `runewidth.RuneWidth('a')` → `1`
- Truncation and row-count logic now use **terminal cell count**, not rune count
- `tailRowCount = (visible + w - 1) / w` uses the correct cell width
- Multi-row clear pattern clears the right number of physical rows

## Testing

### Unit Tests (Updated)

`hyoka/internal/progress/truncate_test.go`:

```go
{
    name:  "emoji",
    input: "🔄",
    want:  2, // ✅ 2 cells, not 1 rune
},
{
    name:  "mix of text and emoji with ANSI",
    input: "\x1b[32m✅ Loaded\x1b[0m",
    want:  9, // ✅ (2) + space(1) + Loaded(6) = 9 cells
},
{
    name:  "wide char emoji truncated correctly",
    input: "🔄 hello world",
    max:   10,
    want:  "🔄 hello …\x1b[0m", // 2 + 1 + 5 + 1 + 1 = 10 cells
},
```

All progress tests pass with `-race`.

### Live Verification

```bash
stty cols 60
go run . run --prompt-id identity-dp-python-default-credential \
  --config "baseline/claude-opus-4.6"
```

**Result:** Tail stayed on exactly one row throughout the entire evaluation. No leaked wrapped content.

## Pattern Summary (Reusable)

When implementing in-place rewrites for terminal tail lines that may wrap:

1. **Use proper cell width calculation:**
   - Import `github.com/mattn/go-runewidth`
   - `runewidth.StringWidth()` for full strings
   - `runewidth.RuneWidth(r)` per-rune during iteration
   - NEVER use `len(string)` (bytes) or `for range` count (runes) for terminal layout

2. **Track row count:** Compute `ceil(visibleWidth(text) / termWidth)` after each write.

3. **Clear all previous rows before rewrite:**
   - Move cursor UP `(rows - 1)` lines via `\033[nA`
   - Clear each row top-to-bottom: `\r\033[2K\n` per row (no `\n` on last row)

4. **Write new content** and update row count.

5. **Truncate NEW text to terminal width** so it never exceeds one row going forward.

6. **Fallback width:** When `term.GetSize` fails (stdout not a TTY), fall back to `COLUMNS` env var, then 80.

## Dependency Trade-Off

**Why not use stdlib only?**

- `golang.org/x/text/width` provides East Asian Width property, but you'd have to implement wcwidth logic yourself
- `github.com/mattn/go-runewidth` is the de-facto standard Go library for terminal cell width (used by many CLI tools)
- Adds 2 deps: `go-runewidth` + `uax29` (Unicode word segmentation)
- **Worth it:** Correct emoji handling is essential for a polished CLI tool

## Decision

**ADOPTED:** The cell-width-aware truncation and row counting is the correct approach. The fix is committed (`fe6efebf`) and verified. Any future interactive terminal renderers in hyoka (or other Go tools) should use `go-runewidth` for layout calculations.

**Key takeaway:** Rune count ≠ terminal cell count. Always use `runewidth.StringWidth()` for terminal layout math when Unicode is involved.

## References

- **Commit `6b3d3d48`:** Added truncation + multi-row clear structure (used rune count — buggy)
- **Commit `fe6efebf`:** Fixed cell width calculation using `go-runewidth`
- **Files:** `hyoka/internal/progress/display_interactive.go`, `hyoka/internal/progress/truncate_test.go`
- **Dependency:** `github.com/mattn/go-runewidth` v0.0.23
- **Related:** `ansi-terminal-output` skill (if extracted)
