package logging

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"
)

func TestConsoleHandler_Enabled(t *testing.T) {
	tests := []struct {
		name      string
		level     slog.Level
		checkLvl  slog.Level
		wantEnabled bool
	}{
		{
			name:        "DEBUG suppressed",
			level:       slog.LevelDebug,
			checkLvl:    slog.LevelDebug,
			wantEnabled: false,
		},
		{
			name:        "INFO suppressed",
			level:       slog.LevelInfo,
			checkLvl:    slog.LevelInfo,
			wantEnabled: false,
		},
		{
			name:        "WARN enabled when handler level is WARN",
			level:       slog.LevelWarn,
			checkLvl:    slog.LevelWarn,
			wantEnabled: true,
		},
		{
			name:        "ERROR enabled when handler level is WARN",
			level:       slog.LevelWarn,
			checkLvl:    slog.LevelError,
			wantEnabled: true,
		},
		{
			name:        "WARN suppressed when handler level is ERROR",
			level:       slog.LevelError,
			checkLvl:    slog.LevelWarn,
			wantEnabled: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewConsoleHandler(&bytes.Buffer{}, tt.level, style.NewFromEnabled(false))
			got := h.Enabled(context.Background(), tt.checkLvl)
			if got != tt.wantEnabled {
				t.Errorf("Enabled(%v) = %v, want %v", tt.checkLvl, got, tt.wantEnabled)
			}
		})
	}
}

func TestConsoleHandler_Handle(t *testing.T) {
	tests := []struct {
		name       string
		level      slog.Level
		msg        string
		attrs      []slog.Attr
		handlerAttrs []slog.Attr
		colorsOn   bool
		want       string
	}{
		{
			name:  "WARN no attrs no color",
			level: slog.LevelWarn,
			msg:   "Plugin not found, skipping",
			want:  "⚠️  Plugin not found, skipping\n",
		},
		{
			name:  "WARN with attrs no color",
			level: slog.LevelWarn,
			msg:   "Plugin not found, skipping",
			attrs: []slog.Attr{
				slog.String("plugin", "azure-sdk-java"),
				slog.String("config", "baseline-skills/claude-sonnet-4.5"),
			},
			want: "⚠️  Plugin not found, skipping (plugin=azure-sdk-java config=baseline-skills/claude-sonnet-4.5)\n",
		},
		{
			name:  "ERROR no attrs no color",
			level: slog.LevelError,
			msg:   "Session failed",
			want:  "❌ Session failed\n",
		},
		{
			name:  "ERROR with attrs no color",
			level: slog.LevelError,
			msg:   "Session failed",
			attrs: []slog.Attr{
				slog.String("reason", "timeout"),
			},
			want: "❌ Session failed (reason=timeout)\n",
		},
		{
			name:  "WARN with color",
			level: slog.LevelWarn,
			msg:   "Warning message",
			attrs: []slog.Attr{
				slog.String("k", "v"),
			},
			colorsOn: true,
			want:     "⚠️  Warning message\x1b[2m (k=v)\x1b[0m\n",
		},
		{
			name:     "ERROR with color",
			level:    slog.LevelError,
			msg:      "Error message",
			attrs:    []slog.Attr{slog.String("code", "42")},
			colorsOn: true,
			want:     "❌ \x1b[31mError message\x1b[0m\x1b[2m (code=42)\x1b[0m\n",
		},
		{
			name:  "handler With attrs",
			level: slog.LevelWarn,
			msg:   "Test",
			handlerAttrs: []slog.Attr{
				slog.String("prompt", "test-prompt"),
				slog.String("config", "test-config"),
			},
			attrs: []slog.Attr{
				slog.String("phase", "generation"),
			},
			want: "⚠️  Test (prompt=test-prompt config=test-config phase=generation)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			styler := style.NewFromEnabled(tt.colorsOn)
			h := NewConsoleHandler(&buf, slog.LevelDebug, styler)

			// Apply handler-level attrs if present
			if len(tt.handlerAttrs) > 0 {
				h = h.WithAttrs(tt.handlerAttrs).(*ConsoleHandler)
			}

			r := slog.NewRecord(time.Now(), tt.level, tt.msg, 0)
			for _, a := range tt.attrs {
				r.AddAttrs(a)
			}

			if err := h.Handle(context.Background(), r); err != nil {
				t.Fatalf("Handle() error = %v", err)
			}

			got := buf.String()
			if got != tt.want {
				t.Errorf("Handle() output mismatch:\ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}

func TestConsoleHandler_InfoDebugSuppressed(t *testing.T) {
	var buf bytes.Buffer
	h := NewConsoleHandler(&buf, slog.LevelDebug, style.NewFromEnabled(false))

	// INFO should not be written
	r := slog.NewRecord(time.Now(), slog.LevelInfo, "Info message", 0)
	if h.Enabled(context.Background(), slog.LevelInfo) {
		t.Fatal("Expected INFO to be disabled")
	}
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle() error = %v", err)
	}
	if buf.Len() > 0 {
		t.Errorf("INFO should produce no output, got: %q", buf.String())
	}

	// DEBUG should not be written
	buf.Reset()
	r = slog.NewRecord(time.Now(), slog.LevelDebug, "Debug message", 0)
	if h.Enabled(context.Background(), slog.LevelDebug) {
		t.Fatal("Expected DEBUG to be disabled")
	}
	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle() error = %v", err)
	}
	if buf.Len() > 0 {
		t.Errorf("DEBUG should produce no output, got: %q", buf.String())
	}
}

func TestConsoleHandler_WithGroup(t *testing.T) {
	// WithGroup is a no-op for console output (groups not rendered),
	// but we verify it doesn't break the handler.
	var buf bytes.Buffer
	h := NewConsoleHandler(&buf, slog.LevelWarn, style.NewFromEnabled(false))
	h2 := h.WithGroup("mygroup")

	r := slog.NewRecord(time.Now(), slog.LevelWarn, "Test", 0)
	r.AddAttrs(slog.String("key", "val"))

	if err := h2.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle() error = %v", err)
	}

	// Group is ignored in console output
	want := "⚠️  Test (key=val)\n"
	got := buf.String()
	if got != want {
		t.Errorf("WithGroup output mismatch:\ngot:  %q\nwant: %q", got, want)
	}
}

// TestConsoleHandler_NO_COLOR verifies that NO_COLOR=1 disables ANSI codes.
func TestConsoleHandler_NO_COLOR(t *testing.T) {
	// Save and restore NO_COLOR
	oldVal := os.Getenv("NO_COLOR")
	defer func() {
		if oldVal == "" {
			os.Unsetenv("NO_COLOR")
		} else {
			os.Setenv("NO_COLOR", oldVal)
		}
	}()

	os.Setenv("NO_COLOR", "1")

	var buf bytes.Buffer
	// style.New will respect NO_COLOR
	styler := style.New(&buf)
	h := NewConsoleHandler(&buf, slog.LevelWarn, styler)

	r := slog.NewRecord(time.Now(), slog.LevelError, "Error", 0)
	r.AddAttrs(slog.String("code", "42"))

	if err := h.Handle(context.Background(), r); err != nil {
		t.Fatalf("Handle() error = %v", err)
	}

	got := buf.String()
	// Should have no ANSI codes when NO_COLOR is set
	if strings.Contains(got, "\x1b[") {
		t.Errorf("Expected no ANSI codes with NO_COLOR=1, got: %q", got)
	}
	// Should still have the emoji and message
	if !strings.Contains(got, "❌") || !strings.Contains(got, "Error") {
		t.Errorf("Expected emoji and message, got: %q", got)
	}
}
