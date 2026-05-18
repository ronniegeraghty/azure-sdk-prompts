package tool

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// stubPluginCloneFailing replaces pluginCloneFn with a stub that always
// returns an error, keeping unit tests offline. Returns a restore function
// the caller defers (or registers via t.Cleanup).
func stubPluginCloneFailing(t *testing.T) func() {
	t.Helper()
	prev := pluginCloneFn
	pluginCloneFn = func(_ context.Context, owner, repo, _, _ string) error {
		return errors.New("stubbed clone failure for " + owner + "/" + repo)
	}
	return func() { pluginCloneFn = prev }
}

// stubPluginCloneSeed replaces pluginCloneFn with a stub that, instead of
// running git, lays down a fake repo at cacheDir using the supplied seed
// function. The seed receives the cache dir and is responsible for creating
// any plugin/skill files the test needs.
func stubPluginCloneSeed(t *testing.T, seed func(cacheDir string) error) func() {
	t.Helper()
	prev := pluginCloneFn
	pluginCloneFn = func(_ context.Context, _, _, _, cacheDir string) error {
		if err := os.MkdirAll(cacheDir, 0o755); err != nil {
			return err
		}
		return seed(cacheDir)
	}
	return func() { pluginCloneFn = prev }
}

func TestPluginFetcher_CanFetch(t *testing.T) {
	cases := []struct {
		name  string
		entry Entry
		want  bool
	}{
		{"explicit remote plugin", Entry{Type: TypePlugin, Source: SourceRemote, Repo: "o/r", Name: "p"}, true},
		{"inferred remote (repo set, no source)", Entry{Type: TypePlugin, Repo: "o/r", Name: "p"}, true},
		{"local plugin", Entry{Type: TypePlugin, Source: SourceLocal, Name: "p"}, false},
		{"plugin no source no repo", Entry{Type: TypePlugin, Name: "p"}, false},
		{"remote skill (gitFetcher's territory)", Entry{Type: TypeSkill, Source: SourceRemote, Repo: "o/r", Name: "s"}, false},
		{"local skill", Entry{Type: TypeSkill, Source: SourceLocal, Path: "/x"}, false},
	}
	f := pluginFetcher{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := f.CanFetch(tc.entry); got != tc.want {
				t.Errorf("CanFetch(%+v) = %v, want %v", tc.entry, got, tc.want)
			}
		})
	}
}

func TestPluginFetcher_DefaultRegistry_RoutesPluginEntries(t *testing.T) {
	f := LookupFetcher(Entry{Type: TypePlugin, Source: SourceRemote, Repo: "owner/repo", Name: "p"})
	if f == nil {
		t.Fatal("expected a fetcher to handle remote plugin entry")
	}
	if f.Name() != pluginFetcherName {
		t.Errorf("expected pluginFetcher (%q), got %q", pluginFetcherName, f.Name())
	}
}

func TestParsePluginRepo(t *testing.T) {
	cases := []struct {
		in            string
		owner, repo   string
	}{
		{"owner/repo", "owner", "repo"},
		{"github.com/owner/repo", "owner", "repo"},
		{"https://github.com/owner/repo", "owner", "repo"},
		{"https://github.com/owner/repo.git", "owner", "repo"},
		{"owner/repo/subpath", "owner", "repo"},
		{"", "", ""},
		{"justone", "", ""},
		{"/leading", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			o, r := parsePluginRepo(tc.in)
			if o != tc.owner || r != tc.repo {
				t.Errorf("parsePluginRepo(%q) = (%q,%q), want (%q,%q)", tc.in, o, r, tc.owner, tc.repo)
			}
		})
	}
}

// TestPluginFetcher_Fetch_FindsPluginInAllPrecedenceLocations stages a
// fake repo containing a single-skill plugin and verifies the fetcher
// locates it at each candidate location, in precedence order.
func TestPluginFetcher_Fetch_FindsPluginInAllPrecedenceLocations(t *testing.T) {
	cases := []struct {
		name     string
		layout   string // relative to cache dir
		wantTail string
	}{
		{
			name:     ".github/plugins layout",
			layout:   filepath.Join(".github", "plugins"),
			wantTail: filepath.Join(".github", "plugins", "my-plugin"),
		},
		{
			name:     ".github/skills layout",
			layout:   filepath.Join(".github", "skills"),
			wantTail: filepath.Join(".github", "skills", "my-plugin"),
		},
		{
			name:     "skills/ layout",
			layout:   "skills",
			wantTail: filepath.Join("skills", "my-plugin"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			restoreCache := toolload.SetTestRoot(filepath.Join(home, "cache"))
			defer restoreCache()
			restoreClone := stubPluginCloneSeed(t, func(cacheDir string) error {
				pluginDir := filepath.Join(cacheDir, tc.layout, "my-plugin")
				if err := os.MkdirAll(pluginDir, 0o755); err != nil {
					return err
				}
				return os.WriteFile(filepath.Join(pluginDir, "SKILL.md"), []byte("# my-plugin\n"), 0o644)
			})
			defer restoreClone()

			f := pluginFetcher{}
			res, err := f.Fetch(context.Background(), FetchRequest{
				Entry: Entry{Type: TypePlugin, Source: SourceRemote, Repo: "owner/repo", Name: "my-plugin"},
			})
			if err != nil {
				t.Fatalf("Fetch: %v", err)
			}
			if !strings.HasSuffix(res.Dir, tc.wantTail) {
				t.Errorf("Dir = %q, want suffix %q", res.Dir, tc.wantTail)
			}
			if res.Version != "default" {
				t.Errorf("Version = %q, want %q", res.Version, "default")
			}
		})
	}
}

// TestPluginFetcher_Fetch_PluginNotFoundInRepo verifies the lookup error
// when the repo clones successfully but the named plugin is absent.
func TestPluginFetcher_Fetch_PluginNotFoundInRepo(t *testing.T) {
	home := t.TempDir()
	restoreCache := toolload.SetTestRoot(filepath.Join(home, "cache"))
	defer restoreCache()
	restoreClone := stubPluginCloneSeed(t, func(cacheDir string) error {
		// Lay down a different plugin so the repo isn't empty.
		other := filepath.Join(cacheDir, ".github", "plugins", "other")
		if err := os.MkdirAll(other, 0o755); err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(other, "SKILL.md"), []byte("# other\n"), 0o644)
	})
	defer restoreClone()

	_, err := pluginFetcher{}.Fetch(context.Background(), FetchRequest{
		Entry: Entry{Type: TypePlugin, Source: SourceRemote, Repo: "owner/repo", Name: "missing"},
	})
	if err == nil {
		t.Fatal("expected error for missing plugin in repo")
	}
	if !strings.Contains(err.Error(), `plugin "missing" not found`) {
		t.Errorf("error %q missing plugin name + not-found phrase", err.Error())
	}
	// All three precedence locations should appear in the error.
	for _, frag := range []string{
		filepath.Join(".github", "plugins", "missing"),
		filepath.Join(".github", "skills", "missing"),
		filepath.Join("skills", "missing"),
	} {
		if !strings.Contains(err.Error(), frag) {
			t.Errorf("expected checked path %q in error: %s", frag, err.Error())
		}
	}
}

// TestPluginFetcher_Fetch_CloneFails surfaces the upstream clone error
// to the caller, which the validator wraps into the aggregated tool-load
// failure summary.
func TestPluginFetcher_Fetch_CloneFails(t *testing.T) {
	home := t.TempDir()
	restoreCache := toolload.SetTestRoot(filepath.Join(home, "cache"))
	defer restoreCache()
	restoreClone := stubPluginCloneFailing(t)
	defer restoreClone()

	_, err := pluginFetcher{}.Fetch(context.Background(), FetchRequest{
		Entry: Entry{Type: TypePlugin, Source: SourceRemote, Repo: "owner/repo", Name: "p"},
	})
	if err == nil {
		t.Fatal("expected clone failure to surface")
	}
	if !strings.Contains(err.Error(), "cloning owner/repo") {
		t.Errorf("expected wrapped clone error; got %q", err.Error())
	}
}

// TestPluginFetcher_Fetch_BadRepo rejects malformed repo locators.
func TestPluginFetcher_Fetch_BadRepo(t *testing.T) {
	cases := []struct {
		name string
		repo string
	}{
		{"empty", ""},
		{"single segment", "justone"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := pluginFetcher{}.Fetch(context.Background(), FetchRequest{
				Entry: Entry{Type: TypePlugin, Source: SourceRemote, Repo: tc.repo, Name: "p"},
			})
			if err == nil {
				t.Fatal("expected error for malformed repo")
			}
		})
	}
}

// TestFindPluginInRepo_PrecedenceOrder verifies that when multiple
// candidate locations exist, .github/plugins wins over .github/skills,
// which wins over skills/. Mirrors plugin.ResolveInstalled exactly.
func TestFindPluginInRepo_PrecedenceOrder(t *testing.T) {
	repo := t.TempDir()
	for _, sub := range []string{
		filepath.Join(".github", "plugins", "p"),
		filepath.Join(".github", "skills", "p"),
		filepath.Join("skills", "p"),
	} {
		dir := filepath.Join(repo, sub)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	got, err := findPluginInRepo(repo, "p")
	if err != nil {
		t.Fatalf("findPluginInRepo: %v", err)
	}
	want := filepath.Join(repo, ".github", "plugins", "p")
	if got != want {
		t.Errorf("got %q, want %q (.github/plugins wins precedence)", got, want)
	}
}

// TestFindPluginInRepo_ContainerLayout accepts container plugins
// (skills/<child>/SKILL.md) at the resolved location, mirroring
// plugin.isPluginDir.
func TestFindPluginInRepo_ContainerLayout(t *testing.T) {
	repo := t.TempDir()
	parent := filepath.Join(repo, ".github", "plugins", "container")
	child := filepath.Join(parent, "skills", "child-1")
	if err := os.MkdirAll(child, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(child, "SKILL.md"), []byte("# child-1"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := findPluginInRepo(repo, "container")
	if err != nil {
		t.Fatalf("findPluginInRepo: %v", err)
	}
	if got != parent {
		t.Errorf("got %q, want %q", got, parent)
	}
}

// TestValidatePluginEntry_FetchSucceedsThenResolvesChildren is the
// end-to-end flow for Item B: cache miss → fetch into the canonical
// cache → ResolveInstalled finds the plugin → child skills fan out
// into the report.
func TestValidatePluginEntry_FetchSucceedsThenResolvesChildren(t *testing.T) {
	home := t.TempDir()
	prevHome := os.Getenv("HOME")
	prevWD, _ := os.Getwd()
	t.Cleanup(func() {
		_ = os.Setenv("HOME", prevHome)
		_ = os.Chdir(prevWD)
	})
	_ = os.Setenv("HOME", home)
	restoreCache := toolload.SetTestRoot(filepath.Join(home, ".hyoka", "cache"))
	defer restoreCache()
	if err := os.Chdir(home); err != nil {
		t.Fatal(err)
	}

	// Container plugin with two child skills: matches the microsoft/skills
	// .github/plugins/<name>/skills/<child>/SKILL.md layout.
	restoreClone := stubPluginCloneSeed(t, func(cacheDir string) error {
		for _, child := range []string{"alpha", "beta"} {
			dir := filepath.Join(cacheDir, ".github", "plugins", "demo", "skills", child)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("# "+child), 0o644); err != nil {
				return err
			}
		}
		return nil
	})
	defer restoreClone()

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Source: SourceRemote, Repo: "owner/repo", Name: "demo"},
		},
		ConfigDir: home,
	})
	if err != nil {
		t.Fatalf("ValidateAndExpand: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected report to succeed; items: %+v", report.Items)
	}
	gotChildren := map[string]bool{}
	for _, it := range report.Items {
		if it.Parent == "demo" && it.Status == "loaded" {
			gotChildren[it.Name] = true
		}
	}
	for _, want := range []string{"alpha", "beta"} {
		if !gotChildren[want] {
			t.Errorf("expected child skill %q in report; items: %+v", want, report.Items)
		}
	}
}
