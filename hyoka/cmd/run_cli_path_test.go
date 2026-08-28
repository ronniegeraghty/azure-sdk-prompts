package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCmdCopilotCLIPathFlag(t *testing.T) {
	cmd := runCmd()
	flag := cmd.Flags().Lookup("copilot-cli-path")
	if flag == nil {
		t.Fatal("expected --copilot-cli-path flag")
	}
	if flag.DefValue != "" {
		t.Fatalf("expected empty default, got %q", flag.DefValue)
	}

	const path = `C:\tools\copilot.exe`
	if err := cmd.ParseFlags([]string{"--copilot-cli-path", path}); err != nil {
		t.Fatalf("parsing --copilot-cli-path: %v", err)
	}
	value, err := cmd.Flags().GetString("copilot-cli-path")
	if err != nil {
		t.Fatalf("reading --copilot-cli-path: %v", err)
	}
	if value != path {
		t.Fatalf("expected %q, got %q", path, value)
	}
}

func TestResolveCopilotCLIPath(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		path, err := resolveCopilotCLIPath("")
		if err != nil {
			t.Fatalf("resolveCopilotCLIPath: %v", err)
		}
		if path != "" {
			t.Fatalf("expected empty path, got %q", path)
		}
	})

	t.Run("file", func(t *testing.T) {
		file, err := os.CreateTemp(t.TempDir(), "copilot-*")
		if err != nil {
			t.Fatalf("creating temporary executable: %v", err)
		}
		if err := file.Close(); err != nil {
			t.Fatalf("closing temporary executable: %v", err)
		}

		path, err := resolveCopilotCLIPath(file.Name())
		if err != nil {
			t.Fatalf("resolveCopilotCLIPath: %v", err)
		}
		expected, err := filepath.Abs(file.Name())
		if err != nil {
			t.Fatalf("resolving expected path: %v", err)
		}
		if path != expected {
			t.Fatalf("expected %q, got %q", expected, path)
		}
	})

	t.Run("missing", func(t *testing.T) {
		missing := filepath.Join(t.TempDir(), "missing-copilot")
		_, err := resolveCopilotCLIPath(missing)
		if err == nil || !strings.Contains(err.Error(), "invalid --copilot-cli-path") {
			t.Fatalf("expected invalid path error, got %v", err)
		}
	})

	t.Run("directory", func(t *testing.T) {
		directory := t.TempDir()
		_, err := resolveCopilotCLIPath(directory)
		if err == nil || !strings.Contains(err.Error(), "path is a directory") {
			t.Fatalf("expected directory error, got %v", err)
		}
	})
}
