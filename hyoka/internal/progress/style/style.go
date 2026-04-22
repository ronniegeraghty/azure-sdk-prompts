// Package style provides small, dependency-free ANSI color and text-style
// primitives for the progress renderers.
//
// The package honors the NO_COLOR convention (https://no-color.org/): any
// non-empty value of the NO_COLOR environment variable disables color output,
// regardless of TTY status. Colors are also disabled when the target writer
// is not a character device (e.g., a pipe, file, or buffer).
//
// Stdlib only — no new dependencies.
package style

import (
	"io"
	"os"
)

// ANSI escape sequences. Kept unexported so callers go through the Styler
// methods, which respect the Enabled flag.
const (
	ansiReset = "\x1b[0m"
	ansiBold  = "\x1b[1m"
	ansiDim   = "\x1b[2m"

	ansiRed    = "\x1b[31m"
	ansiGreen  = "\x1b[32m"
	ansiYellow = "\x1b[33m"
	ansiBlue   = "\x1b[34m"
	ansiCyan   = "\x1b[36m"
)

// Styler wraps strings with ANSI style codes when Enabled is true; otherwise
// it returns the raw string unchanged. The zero value is a safe, disabled
// styler.
type Styler struct {
	// Enabled controls whether style methods emit ANSI escape codes.
	Enabled bool
}

// New returns a Styler that enables colors when all of the following hold:
//
//   - w is a *os.File
//   - that file refers to a character device (i.e., a terminal)
//   - the NO_COLOR environment variable is unset or empty
//
// For any other writer (bytes.Buffer, pipe, regular file), colors are
// disabled. Use NewFromEnabled to bypass detection in tests.
func New(w io.Writer) *Styler {
	return &Styler{Enabled: detectEnabled(w)}
}

// NewFromEnabled returns a Styler with the given enabled state. Intended for
// tests and callers that have already resolved TTY/NO_COLOR policy elsewhere.
func NewFromEnabled(enabled bool) *Styler {
	return &Styler{Enabled: enabled}
}

// detectEnabled implements the TTY + NO_COLOR check. Split out so it's easy
// to reason about independently from the constructor.
func detectEnabled(w io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	f, ok := w.(*os.File)
	if !ok || f == nil {
		return false
	}
	info, err := f.Stat()
	if err != nil {
		return false
	}
	// ModeCharDevice is set for terminals on both Unix and Windows. This is
	// the stdlib-only equivalent of isatty without pulling in x/term.
	return info.Mode()&os.ModeCharDevice != 0
}

// wrap applies the given ANSI code to text when Enabled, otherwise returns
// text unchanged. Kept as the single chokepoint so the Enabled gate lives in
// one place.
func (s *Styler) wrap(code, text string) string {
	if s == nil || !s.Enabled {
		return text
	}
	return code + text + ansiReset
}

// Green wraps text in green ANSI codes when Enabled.
func (s *Styler) Green(text string) string { return s.wrap(ansiGreen, text) }

// Red wraps text in red ANSI codes when Enabled.
func (s *Styler) Red(text string) string { return s.wrap(ansiRed, text) }

// Yellow wraps text in yellow ANSI codes when Enabled.
func (s *Styler) Yellow(text string) string { return s.wrap(ansiYellow, text) }

// Cyan wraps text in cyan ANSI codes when Enabled.
func (s *Styler) Cyan(text string) string { return s.wrap(ansiCyan, text) }

// Blue wraps text in blue ANSI codes when Enabled.
func (s *Styler) Blue(text string) string { return s.wrap(ansiBlue, text) }

// Dim wraps text in the dim SGR attribute when Enabled.
func (s *Styler) Dim(text string) string { return s.wrap(ansiDim, text) }

// Bold wraps text in the bold SGR attribute when Enabled.
func (s *Styler) Bold(text string) string { return s.wrap(ansiBold, text) }

// Reset returns the reset escape sequence when Enabled, or "" otherwise.
// Useful when callers manually interleave codes and need to terminate a run.
func (s *Styler) Reset() string {
	if s == nil || !s.Enabled {
		return ""
	}
	return ansiReset
}

// OK styles text as a success message (green).
func (s *Styler) OK(text string) string { return s.Green(text) }

// Fail styles text as a failure message (red).
func (s *Styler) Fail(text string) string { return s.Red(text) }

// Warn styles text as a warning message (yellow).
func (s *Styler) Warn(text string) string { return s.Yellow(text) }

// Info styles text as an informational message (cyan).
func (s *Styler) Info(text string) string { return s.Cyan(text) }

// Muted styles text as de-emphasized (dim).
func (s *Styler) Muted(text string) string { return s.Dim(text) }
