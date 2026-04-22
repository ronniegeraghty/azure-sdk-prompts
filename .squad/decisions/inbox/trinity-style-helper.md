# Trinity → Style Helper API Reference

**Package:** `github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style`
**Status:** Landed on `ronniegeraghty/dev`. Ready for consumption by `display-interactive-renderer` and `display-ci-renderer`.
**Deps:** stdlib only.

## TL;DR

```go
import "github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"

st := style.New(os.Stdout)                // auto-detects TTY + NO_COLOR
fmt.Fprintln(os.Stdout, st.OK("✅ pass"))  // green if TTY, raw if piped
fmt.Fprintln(os.Stdout, st.Fail("❌ fail"))
```

## Types

```go
type Styler struct {
    Enabled bool  // public — flip for manual override
}
```

The zero value (`Styler{}`) is a valid, disabled styler. Nil pointer (`*Styler(nil)`) is also safe — all methods handle it and return the raw text / empty string.

## Constructors

| Function | Behavior |
|---|---|
| `New(w io.Writer) *Styler` | Enabled iff `w` is a `*os.File` referring to a character device **AND** `NO_COLOR` env is unset/empty. |
| `NewFromEnabled(enabled bool) *Styler` | Bypass detection. Use in tests or when policy is resolved elsewhere. |

**Detection rules** (in order):
1. `os.Getenv("NO_COLOR") != ""` → disabled (regardless of writer).
2. Writer is not a `*os.File` (e.g. `bytes.Buffer`, `*bufio.Writer`) → disabled.
3. `file.Stat().Mode() & os.ModeCharDevice == 0` (pipe/redirect/regular file) → disabled.
4. Otherwise → enabled.

## Color methods

All signatures: `func (s *Styler) X(text string) string`. When enabled, returns `"\x1b[Nm" + text + "\x1b[0m"`. When disabled, returns `text` unchanged.

| Method | ANSI code | Notes |
|---|---|---|
| `Green(text)` | `\x1b[32m` | |
| `Red(text)` | `\x1b[31m` | |
| `Yellow(text)` | `\x1b[33m` | |
| `Cyan(text)` | `\x1b[36m` | |
| `Blue(text)` | `\x1b[34m` | |
| `Dim(text)` | `\x1b[2m` | SGR 2 |
| `Bold(text)` | `\x1b[1m` | SGR 1 |
| `Reset()` | `\x1b[0m` or `""` | No args — returns the reset escape when enabled, empty string otherwise. Use for manual sequences. |

## Semantic helpers

Preferred over raw color methods in renderer code — they encode **intent**, so if the palette ever changes we only touch one place.

| Helper | Alias for | Use for |
|---|---|---|
| `OK(text)` | Green | ✅ Pass, Loaded, Complete |
| `Fail(text)` | Red | ❌ Fail, Error, missing |
| `Warn(text)` | Yellow | ⚠️ Skipped, partial, in-flight-with-warning |
| `Info(text)` | Cyan | Prompt/config headers, neutral metadata |
| `Muted(text)` | Dim | Timestamps, secondary stats, "00:42" elapsed markers |

## Usage patterns

### Pattern 1: renderer with injected styler

```go
type Renderer struct {
    w  io.Writer
    st *style.Styler
}

func NewRenderer(w io.Writer) *Renderer {
    return &Renderer{w: w, st: style.New(w)}
}

func (r *Renderer) WriteOK(line string) {
    fmt.Fprintln(r.w, r.st.OK("✅"), line)
}
```

### Pattern 2: testable renderer

```go
// In tests — force enabled regardless of buffer.
st := style.NewFromEnabled(true)
// Or force disabled for clean golden files.
st := style.NewFromEnabled(false)
```

### Pattern 3: conditional Reset in manual sequences

```go
// When composing a single line with multiple codes, use Reset() to terminate.
fmt.Fprintf(w, "%s%s%s\n", st.Green("["), st.Dim("info"), st.Green("]"))
// Or for a manual run:
fmt.Fprintf(w, "\x1b[1;34m%s%s\n", text, st.Reset()) // Reset returns "" when disabled
```

## What it does NOT do

- Does **not** strip ANSI from input strings. If you pass pre-styled text, it'll get double-wrapped.
- Does **not** handle 256-color / truecolor. If we need those later, extend via methods on `Styler`, not a new package.
- Does **not** expose writer helpers (`Fprintln`, etc.). Callers own their writer — keeps the package tiny.
- Does **not** auto-detect Windows legacy console. Modern Windows 10+ terminals handle ANSI natively; older hosts see escape codes. If that bites us, we'll add Windows VT mode enablement here.

## Testing checklist for consumers

When writing renderer tests, golden-file friendly patterns:

- Force `NewFromEnabled(false)` in tests that compare against plain-text goldens.
- Force `NewFromEnabled(true)` in tests that specifically assert ANSI escape presence.
- Never rely on `New(&bytes.Buffer{})` — it will always return disabled, which can mask bugs where you expected detection.

## Files

- `hyoka/internal/progress/style/style.go` — implementation (~135 LOC)
- `hyoka/internal/progress/style/style_test.go` — table-driven tests

## Verification

```bash
go build ./...
go test -race ./hyoka/internal/progress/style/...
go vet ./hyoka/...
```

All green as of commit on `ronniegeraghty/dev`.

— Trinity 🖤
