package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestParseByteSize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int64
		wantErr bool
	}{
		{"plain bytes", "1024", 1024, false},
		{"1KB", "1KB", 1024, false},
		{"1kb lowercase", "1kb", 1024, false},
		{"512KB", "512KB", 512 * 1024, false},
		{"1MB", "1MB", 1024 * 1024, false},
		{"2MB", "2MB", 2 * 1024 * 1024, false},
		{"1GB", "1GB", 1024 * 1024 * 1024, false},
		{"fractional MB", "1.5MB", int64(1.5 * 1024 * 1024), false},
		{"with spaces", " 1MB ", 1024 * 1024, false},
		{"zero", "0", 0, false},
		{"invalid suffix", "1TB", 0, true},
		{"empty string", "", 0, true},
		{"just letters", "abc", 0, true},
		{"negative", "-1", -1, false},
		{"invalid with suffix", "abcMB", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseByteSize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseByteSize(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("parseByteSize(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestBuildFilter(t *testing.T) {
	tests := []struct {
		name     string
		flags    runFlags
		wantSvc  string
		wantLang string
		wantTags int
	}{
		{
			name:     "empty flags",
			flags:    runFlags{},
			wantSvc:  "",
			wantLang: "",
			wantTags: 0,
		},
		{
			name:     "service and language",
			flags:    runFlags{service: "storage", language: "dotnet"},
			wantSvc:  "storage",
			wantLang: "dotnet",
			wantTags: 0,
		},
		{
			name:     "tags comma separated",
			flags:    runFlags{tags: "auth, blob, identity"},
			wantTags: 3,
		},
		{
			name:     "single tag",
			flags:    runFlags{tags: "auth"},
			wantTags: 1,
		},
		{
			name: "all filter fields",
			flags: runFlags{
				service:  "keyvault",
				language: "python",
				plane:    "data-plane",
				category: "encryption",
				promptID: "kv-encrypt-py",
				tags:     "keys,secrets",
			},
			wantSvc:  "keyvault",
			wantLang: "python",
			wantTags: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := buildFilter(&tt.flags)
			if f.Service != tt.wantSvc {
				t.Errorf("Service = %q, want %q", f.Service, tt.wantSvc)
			}
			if f.Language != tt.wantLang {
				t.Errorf("Language = %q, want %q", f.Language, tt.wantLang)
			}
			if len(f.Tags) != tt.wantTags {
				t.Errorf("Tags count = %d, want %d", len(f.Tags), tt.wantTags)
			}
		})
	}
}

func TestRootCmd_FlagDefaults(t *testing.T) {
	root := rootCmd()

	// Verify global flags exist with correct defaults
	logLevel, err := root.PersistentFlags().GetString("log-level")
	if err != nil {
		t.Fatalf("log-level flag not found: %v", err)
	}
	if logLevel != "warn" {
		t.Errorf("expected log-level default 'warn', got %q", logLevel)
	}

	logFile, err := root.PersistentFlags().GetString("log-file")
	if err != nil {
		t.Fatalf("log-file flag not found: %v", err)
	}
	if logFile != "" {
		t.Errorf("expected log-file default '', got %q", logFile)
	}
}

func TestRunCmd_FlagDefaults(t *testing.T) {
	root := rootCmd()

	// Find the run subcommand
	var run *cobra.Command
	for _, c := range root.Commands() {
		if c.Use == "run" {
			run = c
			break
		}
	}
	if run == nil {
		t.Fatal("run command not found")
	}

	tests := []struct {
		flag     string
		wantStr  string
		wantInt  int
		wantBool bool
		isInt    bool
		isBool   bool
	}{
		{flag: "prompts", wantStr: "./prompts"},
		{flag: "output", wantStr: "./reports"},
		{flag: "progress", wantStr: "auto"},
		{flag: "config-dir", wantStr: "./configs"},
		{flag: "max-output-size", wantStr: "1MB"},
		{flag: "workers", wantInt: 0, isInt: true},
		{flag: "max-sessions", wantInt: 0, isInt: true},
		{flag: "timeout", wantInt: 600, isInt: true},
		{flag: "build-timeout", wantInt: 300, isInt: true},
		{flag: "review-timeout", wantInt: 300, isInt: true},
		{flag: "max-turns", wantInt: 25, isInt: true},
		{flag: "max-files", wantInt: 50, isInt: true},
		{flag: "dry-run", wantBool: false, isBool: true},
		{flag: "skip-review", wantBool: false, isBool: true},
		{flag: "verify-build", wantBool: false, isBool: true},
		{flag: "stub", wantBool: false, isBool: true},
		{flag: "yes", wantBool: false, isBool: true},
		{flag: "all-configs", wantBool: false, isBool: true},
		{flag: "allow-cloud", wantBool: false, isBool: true},
	}

	for _, tt := range tests {
		t.Run(tt.flag, func(t *testing.T) {
			if tt.isInt {
				got, err := run.Flags().GetInt(tt.flag)
				if err != nil {
					t.Fatalf("flag %q not found: %v", tt.flag, err)
				}
				if got != tt.wantInt {
					t.Errorf("flag %q = %d, want %d", tt.flag, got, tt.wantInt)
				}
			} else if tt.isBool {
				got, err := run.Flags().GetBool(tt.flag)
				if err != nil {
					t.Fatalf("flag %q not found: %v", tt.flag, err)
				}
				if got != tt.wantBool {
					t.Errorf("flag %q = %v, want %v", tt.flag, got, tt.wantBool)
				}
			} else {
				got, err := run.Flags().GetString(tt.flag)
				if err != nil {
					t.Fatalf("flag %q not found: %v", tt.flag, err)
				}
				if got != tt.wantStr {
					t.Errorf("flag %q = %q, want %q", tt.flag, got, tt.wantStr)
				}
			}
		})
	}
}
