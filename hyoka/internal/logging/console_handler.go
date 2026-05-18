// Package logging provides structured logging for hyoka using log/slog.
package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"
)

// ConsoleHandler is an slog.Handler that formats log records for human
// consumption on stdout/stderr. Unlike the structured TextHandler, it:
//   - Suppresses INFO and DEBUG entirely (diagnostics-only, noise on console).
//   - Formats WARN as: ⚠️  <msg> (key=val key=val) — attrs dim
//   - Formats ERROR as: ❌ <msg> (key=val key=val) — msg red, attrs dim
//   - Respects NO_COLOR and TTY detection via the style package.
//
// For diagnostic logging to a file, use slog.NewTextHandler instead.
type ConsoleHandler struct {
	w      io.Writer
	level  slog.Leveler
	styler *style.Styler
	groups []string  // group stack for WithGroup
	attrs  []slog.Attr // accumulated attrs for With
}

// NewConsoleHandler returns a handler that writes human-friendly log messages
// to w. The styler controls whether ANSI colors/dim are applied.
func NewConsoleHandler(w io.Writer, level slog.Leveler, styler *style.Styler) *ConsoleHandler {
	return &ConsoleHandler{
		w:      w,
		level:  level,
		styler: styler,
	}
}

// Enabled reports whether the handler will process records at the given level.
func (h *ConsoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	// Suppress INFO and DEBUG on console — they're diagnostic noise.
	// WARN and ERROR go through if they meet the configured level.
	if level < slog.LevelWarn {
		return false
	}
	minLevel := slog.LevelWarn
	if h.level != nil {
		minLevel = h.level.Level()
	}
	return level >= minLevel
}

// Handle formats and writes a log record to the underlying writer.
func (h *ConsoleHandler) Handle(ctx context.Context, r slog.Record) error {
	// Check enabled first — don't write INFO/DEBUG
	if !h.Enabled(ctx, r.Level) {
		return nil
	}

	var sb strings.Builder

	// Prefix with level emoji
	switch r.Level {
	case slog.LevelWarn:
		sb.WriteString("⚠️  ")
	case slog.LevelError:
		// Error message in red
		sb.WriteString("❌ ")
	default:
		// Should not reach here due to Enabled check, but be defensive.
		sb.WriteString(fmt.Sprintf("[%s] ", r.Level.String()))
	}

	// Write the message
	msg := r.Message
	if r.Level >= slog.LevelError {
		msg = h.styler.Red(msg)
	}
	sb.WriteString(msg)

	// Collect attrs (handler attrs + record attrs)
	attrPairs := make([]string, 0, len(h.attrs)+r.NumAttrs())
	for _, a := range h.attrs {
		attrPairs = append(attrPairs, formatAttr(a))
	}
	r.Attrs(func(a slog.Attr) bool {
		attrPairs = append(attrPairs, formatAttr(a))
		return true
	})

	// Append attrs in dim, if any
	if len(attrPairs) > 0 {
		attrsText := " (" + strings.Join(attrPairs, " ") + ")"
		sb.WriteString(h.styler.Dim(attrsText))
	}

	sb.WriteString("\n")
	_, err := h.w.Write([]byte(sb.String()))
	return err
}

// WithAttrs returns a new handler with the given attrs added.
func (h *ConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)
	return &ConsoleHandler{
		w:      h.w,
		level:  h.level,
		styler: h.styler,
		groups: h.groups,
		attrs:  newAttrs,
	}
}

// WithGroup returns a new handler with the given group name added.
// Groups are not currently rendered in console output; this is a no-op for
// forward compatibility.
func (h *ConsoleHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	newGroups := make([]string, len(h.groups)+1)
	copy(newGroups, h.groups)
	newGroups[len(h.groups)] = name
	return &ConsoleHandler{
		w:      h.w,
		level:  h.level,
		styler: h.styler,
		groups: newGroups,
		attrs:  h.attrs,
	}
}

// formatAttr formats a single attribute as key=value.
func formatAttr(a slog.Attr) string {
	if a.Key == "" {
		return ""
	}
	return fmt.Sprintf("%s=%v", a.Key, a.Value)
}
