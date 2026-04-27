package toolload

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCacheRoot_EnvOverride(t *testing.T) {
	t.Setenv(EnvCacheDir, "/custom/hyoka-cache")
	t.Setenv(EnvXDGCacheHome, "/should/be/ignored")
	resetForTest()
	if got := CacheRoot(); got != "/custom/hyoka-cache" {
		t.Errorf("HYOKA_CACHE_DIR not honored: got %q", got)
	}
}

func TestCacheRoot_XDG(t *testing.T) {
	t.Setenv(EnvCacheDir, "")
	t.Setenv(EnvXDGCacheHome, "/xdg")
	resetForTest()
	if got, want := CacheRoot(), filepath.Join("/xdg", "hyoka"); got != want {
		t.Errorf("XDG_CACHE_HOME not honored: got %q want %q", got, want)
	}
}

func TestCacheRoot_HomeDefault(t *testing.T) {
	t.Setenv(EnvCacheDir, "")
	t.Setenv(EnvXDGCacheHome, "")
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("no home dir on this system")
	}
	resetForTest()
	want := filepath.Join(home, ".hyoka", "cache")
	if got := CacheRoot(); got != want {
		t.Errorf("home default not used: got %q want %q", got, want)
	}
}

func TestRepoCacheDir(t *testing.T) {
	t.Setenv(EnvCacheDir, "/c")
	resetForTest()
	got := RepoCacheDir("microsoft", "skills", "v1")
	want := filepath.Join("/c", "repos", "microsoft", "skills", "v1")
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	got = RepoCacheDir("acme", "widgets", "")
	want = filepath.Join("/c", "repos", "acme", "widgets", "default")
	if got != want {
		t.Errorf("empty version not normalized: got %q want %q", got, want)
	}
}

func TestVersionSegment(t *testing.T) {
	cases := map[string]string{
		"":        "default",
		"default": "default",
		"v1.2.3":  "v1.2.3",
		"main":    "main",
	}
	for in, want := range cases {
		if got := VersionSegment(in); got != want {
			t.Errorf("VersionSegment(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestCacheRoot_NoCwdPollution asserts CacheRoot never returns a path
// containing ".skills-cache/" — the historical bug.
func TestCacheRoot_NoCwdPollution(t *testing.T) {
	t.Setenv(EnvCacheDir, "")
	t.Setenv(EnvXDGCacheHome, "")
	resetForTest()
	if strings.Contains(CacheRoot(), ".skills-cache") {
		t.Errorf("CacheRoot leaked .skills-cache: %s", CacheRoot())
	}
}

// TestCacheRoot_LegacySkillsCacheWarningFiresOnce verifies that when a
// legacy .skills-cache/ directory exists in the current working directory,
// CacheRoot logs a single migration warning (not one-per-tool, not zero).
// Item A added the warning; this test guards against regressions that might
// drop the warning or fire it on every cache access.
func TestCacheRoot_LegacySkillsCacheWarningFiresOnce(t *testing.T) {
	tmp := t.TempDir()
	if err := os.MkdirAll(filepath.Join(tmp, ".skills-cache"), 0o755); err != nil {
		t.Fatalf("seed legacy dir: %v", err)
	}
	prevWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(prevWd) })

	var buf bytes.Buffer
	prevLogger := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prevLogger) })

	t.Setenv(EnvCacheDir, "/test/cache")
	t.Setenv(EnvXDGCacheHome, "")
	resetForTest()

	// First call: must fire the warning.
	_ = CacheRoot()
	// Subsequent calls: sync.Once gates everything inside cacheRootOnce.Do,
	// so the warning must NOT fire again. The fires-once contract is what
	// keeps the message out of per-tool log spam.
	_ = CacheRoot()
	_ = CacheRoot()

	out := buf.String()
	count := strings.Count(out, "legacy .skills-cache/ dir detected")
	if count != 1 {
		t.Errorf("expected legacy-cache warning to fire exactly once, got %d\nlog output:\n%s", count, out)
	}
	if !strings.Contains(out, filepath.Join(tmp, ".skills-cache")) {
		t.Errorf("warning should include the detected path %q\nlog output:\n%s",
			filepath.Join(tmp, ".skills-cache"), out)
	}
}

// TestCacheRoot_NoLegacySkillsCache_NoWarning is the negative case for
// TestCacheRoot_LegacySkillsCacheWarningFiresOnce: when no .skills-cache/
// dir exists, the warning must stay quiet.
func TestCacheRoot_NoLegacySkillsCache_NoWarning(t *testing.T) {
	tmp := t.TempDir() // contains nothing
	prevWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(prevWd) })

	var buf bytes.Buffer
	prevLogger := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prevLogger) })

	t.Setenv(EnvCacheDir, "/test/cache")
	t.Setenv(EnvXDGCacheHome, "")
	resetForTest()
	_ = CacheRoot()

	if strings.Contains(buf.String(), "legacy .skills-cache/ dir detected") {
		t.Errorf("legacy warning fired without a legacy dir present\nlog output:\n%s", buf.String())
	}
}
