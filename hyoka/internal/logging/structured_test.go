package logging

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetup_LogFile(t *testing.T) {
	dir := t.TempDir()
	logPath := filepath.Join(dir, "test.log")

	closer, err := Setup(Options{Level: "info", FilePath: logPath})
	if err != nil {
		t.Fatalf("Setup() error: %v", err)
	}

	slog.Info("test log message", "key", "value")
	closer()

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "test log message") {
		t.Errorf("expected log file to contain 'test log message', got %q", content)
	}
	if !strings.Contains(content, "key=value") {
		t.Errorf("expected log file to contain 'key=value', got %q", content)
	}
}

func TestSetup_InvalidPath(t *testing.T) {
	_, err := Setup(Options{FilePath: "/nonexistent/dir/test.log"})
	if err == nil {
		t.Fatal("expected error for invalid log file path")
	}
	if !strings.Contains(err.Error(), "opening log file") {
		t.Errorf("expected 'opening log file' in error, got %q", err.Error())
	}
}

func TestSetup_LevelFiltering(t *testing.T) {
	dir := t.TempDir()
	logPath := filepath.Join(dir, "level-test.log")

	closer, err := Setup(Options{Level: "error", FilePath: logPath})
	if err != nil {
		t.Fatalf("Setup() error: %v", err)
	}

	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Warn("warn msg")
	slog.Error("error msg")
	closer()

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	content := string(data)

	if strings.Contains(content, "debug msg") {
		t.Error("debug message should be filtered at error level")
	}
	if strings.Contains(content, "info msg") {
		t.Error("info message should be filtered at error level")
	}
	if strings.Contains(content, "warn msg") {
		t.Error("warn message should be filtered at error level")
	}
	if !strings.Contains(content, "error msg") {
		t.Error("error message should be present at error level")
	}
}

func TestEvalLogger_StructuredFields(t *testing.T) {
	// Capture log output to a buffer
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	l := EvalLogger("my-prompt", "my-config", "generation", 3)
	l.Info("test message")

	output := buf.String()

	fields := map[string]string{
		"prompt": "my-prompt",
		"config": "my-config",
		"phase":  "generation",
		"worker": "3",
	}
	for key, val := range fields {
		expected := key + "=" + val
		if !strings.Contains(output, expected) {
			t.Errorf("expected %q in log output, got %q", expected, output)
		}
	}
}

func TestWithPhase_UpdatesPhaseField(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	base := slog.New(handler)

	l := base.With("prompt", "p1", "phase", "generation")
	l2 := WithPhase(l, "review")
	l2.Info("phase changed")

	output := buf.String()
	if !strings.Contains(output, "phase=review") {
		t.Errorf("expected 'phase=review' in output, got %q", output)
	}
}

func TestWithTurn_AddsTurnField(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	base := slog.New(handler)

	l := base.With("prompt", "p1")
	l2 := WithTurn(l, 7)
	l2.Info("turn tracked")

	output := buf.String()
	if !strings.Contains(output, "turn=7") {
		t.Errorf("expected 'turn=7' in output, got %q", output)
	}
}

func TestResolveLevel_DebugFlagWithLevel(t *testing.T) {
	// When both Debug flag and explicit Level are set, Level takes precedence
	got := resolveLevel(Options{Debug: true, Level: "error"})
	if got != slog.LevelError {
		t.Errorf("expected LevelError when Level is set, got %v", got)
	}
}

func TestResolveLevel_CaseInsensitive(t *testing.T) {
	tests := []struct {
		input string
		want  slog.Level
	}{
		{"DEBUG", slog.LevelDebug},
		{"Info", slog.LevelInfo},
		{"WARN", slog.LevelWarn},
		{"Error", slog.LevelError},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := resolveLevel(Options{Level: tt.input})
			if got != tt.want {
				t.Errorf("resolveLevel(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
