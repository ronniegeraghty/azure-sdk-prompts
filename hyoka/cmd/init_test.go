package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

func TestInitCmdCreatesProjectDir(t *testing.T) {
	dir := t.TempDir()

	// Save and restore CWD
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	cmd := initCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("init command failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "✓ Initialized") {
		t.Errorf("expected success message, got: %s", output)
	}

	// Verify structure
	for _, sub := range config.ProjectSubdirs {
		path := filepath.Join(dir, ".hyoka", sub)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("missing %s: %v", sub, err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("%s should be a directory", sub)
		}
	}

	// Verify .gitignore
	gi, err := os.ReadFile(filepath.Join(dir, ".hyoka", ".gitignore"))
	if err != nil {
		t.Fatalf("missing .gitignore: %v", err)
	}
	if !strings.Contains(string(gi), "reports/") {
		t.Error(".gitignore should contain reports/")
	}
}

func TestInitCmdIdempotent(t *testing.T) {
	dir := t.TempDir()

	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	cmd1 := initCmd()
	cmd1.SetOut(&bytes.Buffer{})
	if err := cmd1.Execute(); err != nil {
		t.Fatalf("first init: %v", err)
	}

	cmd2 := initCmd()
	cmd2.SetOut(&bytes.Buffer{})
	if err := cmd2.Execute(); err != nil {
		t.Fatalf("second init (idempotent): %v", err)
	}
}

func TestInitCmdRegistered(t *testing.T) {
	cmd := rootCmd()
	names := make(map[string]bool)
	for _, sub := range cmd.Commands() {
		names[sub.Name()] = true
	}
	if !names["init"] {
		t.Error("init subcommand should be registered on root")
	}
}
