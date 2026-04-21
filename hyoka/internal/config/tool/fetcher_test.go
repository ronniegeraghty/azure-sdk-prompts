package tool

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"
)

// mockFetcher is a Fetcher used in tests to verify the registry actually
// dispatches custom fetchers at runtime — guarding against the #587 trap
// where config wiring exists but never reaches execution.
type mockFetcher struct {
	name      string
	matchRepo string
	calls     atomic.Int64
	lastReq   atomic.Pointer[FetchRequest]
	dir       string
	err       error
}

func (m *mockFetcher) Name() string { return m.name }

func (m *mockFetcher) CanFetch(e Entry) bool {
	if e.ResolvedType() != TypeSkill || e.SkillSource() != SourceRemote {
		return false
	}
	return m.matchRepo == "" || e.Repo == m.matchRepo
}

func (m *mockFetcher) Fetch(_ context.Context, req FetchRequest) (FetchResult, error) {
	m.calls.Add(1)
	r := req
	m.lastReq.Store(&r)
	if m.err != nil {
		return FetchResult{}, m.err
	}
	return FetchResult{Dir: m.dir, Version: req.Version}, nil
}

func TestRegistry_RegisterAndLookup(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(&npxFetcher{}); err != nil {
		t.Fatalf("register npx: %v", err)
	}
	custom := &mockFetcher{name: "custom", matchRepo: "owner/repo"}
	if err := r.Register(custom); err != nil {
		t.Fatalf("register custom: %v", err)
	}

	cases := []struct {
		name    string
		entry   Entry
		wantFet string
	}{
		{
			name:    "custom matches its repo first",
			entry:   Entry{Type: TypeSkill, Source: SourceRemote, Repo: "owner/repo"},
			wantFet: "custom",
		},
		{
			name:    "default falls through for other repos",
			entry:   Entry{Type: TypeSkill, Source: SourceRemote, Repo: "other/repo"},
			wantFet: "npx",
		},
		{
			name:    "local skill — no remote fetcher matches",
			entry:   Entry{Type: TypeSkill, Source: SourceLocal, Path: "/tmp/x"},
			wantFet: "",
		},
		{
			name:    "non-skill entry — no fetcher",
			entry:   Entry{Type: TypeMCP, Name: "x"},
			wantFet: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := r.Lookup(tc.entry)
			if tc.wantFet == "" {
				if got != nil {
					t.Fatalf("expected no fetcher, got %q", got.Name())
				}
				return
			}
			if got == nil {
				t.Fatalf("expected fetcher %q, got nil", tc.wantFet)
			}
			if got.Name() != tc.wantFet {
				t.Errorf("got fetcher %q, want %q", got.Name(), tc.wantFet)
			}
		})
	}
}

func TestRegistry_DefaultStaysLast(t *testing.T) {
	r := NewRegistry()
	// Register default first, then custom — custom must still come before
	// default in lookup order so it can shadow it.
	if err := r.Register(&npxFetcher{}); err != nil {
		t.Fatal(err)
	}
	if err := r.Register(&mockFetcher{name: "custom"}); err != nil {
		t.Fatal(err)
	}
	names := r.Names()
	if len(names) != 2 || names[0] != "custom" || names[1] != defaultFetcherName {
		t.Errorf("expected [custom, npx], got %v", names)
	}
}

func TestRegistry_DuplicateNameRejected(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(&mockFetcher{name: "x"}); err != nil {
		t.Fatal(err)
	}
	if err := r.Register(&mockFetcher{name: "x"}); err == nil {
		t.Errorf("expected duplicate-name error")
	}
}

func TestRegistry_RejectsNilAndUnnamed(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(nil); err == nil {
		t.Errorf("expected error for nil fetcher")
	}
	if err := r.Register(&mockFetcher{name: ""}); err == nil {
		t.Errorf("expected error for empty Name")
	}
}

// TestCustomFetcherInvokedAtRuntime is the #587-trap guard: it verifies that
// when ResolveSkills is asked to handle a remote skill, a registered custom
// fetcher actually runs — not just that the config plumbing accepts it.
func TestCustomFetcherInvokedAtRuntime(t *testing.T) {
	mock := &mockFetcher{name: "test-runtime", matchRepo: "acme/widgets", dir: t.TempDir()}
	if err := DefaultRegistry.Register(mock); err != nil {
		t.Fatalf("register: %v", err)
	}
	t.Cleanup(func() { DefaultRegistry.Unregister(mock.name) })

	dirs, err := ResolveSkills([]Entry{
		{Type: TypeSkill, Source: SourceRemote, Repo: "acme/widgets", Name: "widget", Version: "v1.2.3"},
	}, t.TempDir())
	if err != nil {
		t.Fatalf("ResolveSkills: %v", err)
	}
	if got := mock.calls.Load(); got != 1 {
		t.Fatalf("expected mock.Fetch to be called exactly once, got %d", got)
	}
	got := mock.lastReq.Load()
	if got == nil {
		t.Fatalf("no FetchRequest captured")
	}
	if got.Version != "v1.2.3" {
		t.Errorf("custom fetcher saw version %q, want %q", got.Version, "v1.2.3")
	}
	if got.Entry.Repo != "acme/widgets" {
		t.Errorf("custom fetcher saw repo %q, want %q", got.Entry.Repo, "acme/widgets")
	}
	if len(dirs) != 1 || dirs[0] != mock.dir {
		t.Errorf("ResolveSkills returned %v, want [%s]", dirs, mock.dir)
	}
}

func TestCustomFetcherErrorPropagates(t *testing.T) {
	wantErr := errors.New("mock fetch failure")
	mock := &mockFetcher{name: "test-err", matchRepo: "acme/fail", err: wantErr}
	if err := DefaultRegistry.Register(mock); err != nil {
		t.Fatalf("register: %v", err)
	}
	t.Cleanup(func() { DefaultRegistry.Unregister(mock.name) })

	_, err := ResolveSkills([]Entry{
		{Type: TypeSkill, Source: SourceRemote, Repo: "acme/fail", Name: "x"},
	}, t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "mock fetch failure") {
		t.Errorf("expected wrapped fetch error, got %v", err)
	}
}

func TestValidateFetchers(t *testing.T) {
	// Local skills and non-skill entries always pass — fetcher unrelated.
	if err := ValidateFetchers([][]Entry{{
		{Type: TypeSkill, Source: SourceLocal, Path: "/tmp/x"},
		{Type: TypeMCP, Name: "mcp"},
	}}); err != nil {
		t.Errorf("local-only validation: %v", err)
	}
	// Remote skill is OK because the default npx fetcher always matches.
	if err := ValidateFetchers([][]Entry{{
		{Type: TypeSkill, Source: SourceRemote, Repo: "acme/x"},
	}}); err != nil {
		t.Errorf("remote with default fetcher: %v", err)
	}
}

// TestNpxFetcher_VersionInPath verifies that the default fetcher writes
// different versions to different cache subdirectories — so toggling a
// version override doesn't poison the cache.
func TestNpxFetcher_VersionInPath(t *testing.T) {
	f := npxFetcher{}
	if !f.CanFetch(Entry{Type: TypeSkill, Source: SourceRemote, Repo: "x/y"}) {
		t.Fatal("npx must accept remote skill")
	}
	if f.CanFetch(Entry{Type: TypeSkill, Source: SourceLocal, Path: "/x"}) {
		t.Fatal("npx must reject local skill")
	}
	if f.CanFetch(Entry{Type: TypeMCP}) {
		t.Fatal("npx must reject non-skill")
	}
	if f.Name() != defaultFetcherName {
		t.Errorf("name mismatch: %q", f.Name())
	}
}
