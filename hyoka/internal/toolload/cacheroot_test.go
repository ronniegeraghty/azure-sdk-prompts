package toolload

import (
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
