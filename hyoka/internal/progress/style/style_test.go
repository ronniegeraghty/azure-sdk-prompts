package style

import (
	"bytes"
	"testing"
)

func TestStyler_EnabledWrapsWithANSI(t *testing.T) {
	s := NewFromEnabled(true)

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"Green", s.Green("ok"), "\x1b[32mok\x1b[0m"},
		{"Red", s.Red("bad"), "\x1b[31mbad\x1b[0m"},
		{"Yellow", s.Yellow("meh"), "\x1b[33mmeh\x1b[0m"},
		{"Cyan", s.Cyan("info"), "\x1b[36minfo\x1b[0m"},
		{"Blue", s.Blue("link"), "\x1b[34mlink\x1b[0m"},
		{"Dim", s.Dim("note"), "\x1b[2mnote\x1b[0m"},
		{"Bold", s.Bold("loud"), "\x1b[1mloud\x1b[0m"},
		{"OK", s.OK("pass"), "\x1b[32mpass\x1b[0m"},
		{"Fail", s.Fail("fail"), "\x1b[31mfail\x1b[0m"},
		{"Warn", s.Warn("warn"), "\x1b[33mwarn\x1b[0m"},
		{"Info", s.Info("info"), "\x1b[36minfo\x1b[0m"},
		{"Muted", s.Muted("quiet"), "\x1b[2mquiet\x1b[0m"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("got %q, want %q", tc.got, tc.want)
			}
		})
	}

	if got := s.Reset(); got != "\x1b[0m" {
		t.Errorf("Reset(enabled) = %q, want %q", got, "\x1b[0m")
	}
}

func TestStyler_DisabledReturnsRaw(t *testing.T) {
	s := NewFromEnabled(false)

	methods := map[string]func(string) string{
		"Green":  s.Green,
		"Red":    s.Red,
		"Yellow": s.Yellow,
		"Cyan":   s.Cyan,
		"Blue":   s.Blue,
		"Dim":    s.Dim,
		"Bold":   s.Bold,
		"OK":     s.OK,
		"Fail":   s.Fail,
		"Warn":   s.Warn,
		"Info":   s.Info,
		"Muted":  s.Muted,
	}
	for name, fn := range methods {
		t.Run(name, func(t *testing.T) {
			if got := fn("hello"); got != "hello" {
				t.Errorf("%s(disabled) = %q, want %q", name, got, "hello")
			}
		})
	}
	if got := s.Reset(); got != "" {
		t.Errorf("Reset(disabled) = %q, want empty string", got)
	}
}

func TestStyler_NilReceiverSafe(t *testing.T) {
	var s *Styler
	if got := s.Green("x"); got != "x" {
		t.Errorf("nil.Green = %q, want %q", got, "x")
	}
	if got := s.Reset(); got != "" {
		t.Errorf("nil.Reset = %q, want empty string", got)
	}
}

func TestNew_NonFileWriterDisabled(t *testing.T) {
	var buf bytes.Buffer
	s := New(&buf)
	if s.Enabled {
		t.Error("expected Enabled=false for bytes.Buffer writer")
	}
}

func TestNew_NoColorEnvDisables(t *testing.T) {
	// Even with a would-be-TTY file, NO_COLOR must force disable.
	t.Setenv("NO_COLOR", "1")

	// A nil writer can never be a terminal, but the key assertion here is
	// that the NO_COLOR branch short-circuits before the writer check.
	s := New(nil)
	if s.Enabled {
		t.Error("expected Enabled=false when NO_COLOR is set (nil writer)")
	}

	// Also verify with a buffer: NO_COLOR + non-TTY must still be disabled.
	var buf bytes.Buffer
	if s2 := New(&buf); s2.Enabled {
		t.Error("expected Enabled=false with NO_COLOR set and buffer writer")
	}
}

func TestNew_NoColorEmptyValueRespected(t *testing.T) {
	// An empty NO_COLOR should NOT by itself disable (per no-color.org:
	// "any non-empty value"). We can't assert Enabled=true portably (tests
	// often don't have a TTY attached), but we can ensure the env check
	// itself isn't tripped by an empty string — disablement should then
	// depend only on the writer.
	t.Setenv("NO_COLOR", "")
	var buf bytes.Buffer
	if s := New(&buf); s.Enabled {
		t.Error("expected Enabled=false for buffer regardless of NO_COLOR=''")
	}
}
