package tools

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// remoteRegistryYAML is a valid registry that conforms to the schema
// defined in registry.go (version "1", typed entries with mcp/skill blocks).
const remoteRegistryYAML = `
version: "1"
tools:
  - name: azure-identity-mcp
    type: mcp
    description: Azure Identity MCP server
    version: "1.0.0"
    mcp:
      command: npx
      args: ["-y", "@azure/identity-mcp"]
      tools: ["get-token"]
  - name: keyvault-skill
    type: skill
    description: Key Vault Python skill
    skill:
      source: remote
      repo: microsoft/skills
      name: azure-keyvault-py
`

func TestFetch_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/yaml")
		w.Write([]byte(remoteRegistryYAML))
	}))
	defer srv.Close()

	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)
	f.HTTPClient = srv.Client()

	reg, err := f.Fetch(srv.URL + "/registry.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reg.Version != "1" {
		t.Errorf("got version %q, want %q", reg.Version, "1")
	}
	if len(reg.Tools) != 2 {
		t.Fatalf("got %d tools, want 2", len(reg.Tools))
	}
	if reg.Tools[0].Name != "azure-identity-mcp" {
		t.Errorf("tool[0].Name = %q, want %q", reg.Tools[0].Name, "azure-identity-mcp")
	}
	if reg.Tools[0].Type != TypeMCP {
		t.Errorf("tool[0].Type = %q, want %q", reg.Tools[0].Type, TypeMCP)
	}
	if reg.Tools[1].Type != TypeSkill {
		t.Errorf("tool[1].Type = %q, want %q", reg.Tools[1].Type, TypeSkill)
	}
}

func TestFetch_CacheTTL(t *testing.T) {
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Write([]byte(remoteRegistryYAML))
	}))
	defer srv.Close()

	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)
	f.HTTPClient = srv.Client()
	f.CacheTTL = 1 * time.Hour

	url := srv.URL + "/registry.yaml"

	// First fetch — hits the server.
	if _, err := f.Fetch(url); err != nil {
		t.Fatalf("first fetch: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected 1 HTTP call, got %d", callCount)
	}

	// Second fetch within TTL — served from cache.
	if _, err := f.Fetch(url); err != nil {
		t.Fatalf("second fetch: %v", err)
	}
	if callCount != 1 {
		t.Fatalf("expected still 1 HTTP call after cached fetch, got %d", callCount)
	}
}

func TestFetch_CacheExpiry(t *testing.T) {
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Write([]byte(remoteRegistryYAML))
	}))
	defer srv.Close()

	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)
	f.HTTPClient = srv.Client()
	f.CacheTTL = 1 * time.Millisecond

	url := srv.URL + "/registry.yaml"

	if _, err := f.Fetch(url); err != nil {
		t.Fatalf("first fetch: %v", err)
	}

	time.Sleep(10 * time.Millisecond)

	if _, err := f.Fetch(url); err != nil {
		t.Fatalf("second fetch after expiry: %v", err)
	}
	if callCount != 2 {
		t.Fatalf("expected 2 HTTP calls after cache expiry, got %d", callCount)
	}
}

func TestFetch_GracefulDegradation(t *testing.T) {
	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)

	// Pre-populate stale cache with a valid registry.
	url := "http://unreachable.invalid/registry.yaml"
	cachePath := f.cachePath(NormalizeGitHubURL(url))
	staleReg := &ToolRegistry{
		Version: "1",
		Tools: []ToolEntry{
			{
				Name: "stale-mcp",
				Type: TypeMCP,
				MCP:  &MCPConfig{Command: "echo"},
			},
		},
	}
	if err := f.writeCache(cachePath, staleReg); err != nil {
		t.Fatalf("seeding cache: %v", err)
	}

	// Set TTL to 0 so cache is always "expired" but stale-readable.
	f.CacheTTL = 0

	reg, err := f.Fetch(url)
	if err != nil {
		t.Fatalf("expected graceful degradation, got error: %v", err)
	}
	if reg.Tools[0].Name != "stale-mcp" {
		t.Errorf("got name %q, want %q", reg.Tools[0].Name, "stale-mcp")
	}
}

func TestFetch_NetworkError_NoCache(t *testing.T) {
	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)
	f.HTTPClient = &http.Client{Timeout: 100 * time.Millisecond}

	_, err := f.Fetch("http://unreachable.invalid/registry.yaml")
	if err == nil {
		t.Fatal("expected error for unreachable URL with no cache")
	}
}

func TestFetch_InvalidYAML(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not: valid: yaml: ["))
	}))
	defer srv.Close()

	f := NewRemoteFetcher(t.TempDir())
	f.HTTPClient = srv.Client()

	_, err := f.Fetch(srv.URL + "/bad.yaml")
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestFetch_HTTP404(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer srv.Close()

	f := NewRemoteFetcher(t.TempDir())
	f.HTTPClient = srv.Client()

	_, err := f.Fetch(srv.URL + "/missing.yaml")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
}

func TestFetch_OversizedResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(make([]byte, remoteMaxSize+100))
	}))
	defer srv.Close()

	f := NewRemoteFetcher(t.TempDir())
	f.HTTPClient = srv.Client()

	_, err := f.Fetch(srv.URL + "/huge.yaml")
	if err == nil {
		t.Fatal("expected error for oversized response")
	}
}

func TestNormalizeGitHubURL(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "https://github.com/Azure/azure-sdk-tools/blob/main/tools-registry.yaml",
			want:  "https://raw.githubusercontent.com/Azure/azure-sdk-tools/main/tools-registry.yaml",
		},
		{
			input: "https://raw.githubusercontent.com/Azure/azure-sdk-tools/main/tools-registry.yaml",
			want:  "https://raw.githubusercontent.com/Azure/azure-sdk-tools/main/tools-registry.yaml",
		},
		{
			input: "https://example.com/registry.yaml",
			want:  "https://example.com/registry.yaml",
		},
		{
			input: "https://github.com/org/repo/blob/v2/path/to/registry.yaml",
			want:  "https://raw.githubusercontent.com/org/repo/v2/path/to/registry.yaml",
		},
	}

	for _, tc := range tests {
		got := NormalizeGitHubURL(tc.input)
		if got != tc.want {
			t.Errorf("NormalizeGitHubURL(%q)\n  got  %q\n  want %q", tc.input, got, tc.want)
		}
	}
}

func TestCachePath_Deterministic(t *testing.T) {
	f := NewRemoteFetcher(t.TempDir())

	p1 := f.cachePath("https://example.com/a.yaml")
	p2 := f.cachePath("https://example.com/a.yaml")
	p3 := f.cachePath("https://example.com/b.yaml")

	if p1 != p2 {
		t.Error("same URL should produce same cache path")
	}
	if p1 == p3 {
		t.Error("different URLs should produce different cache paths")
	}
	if ext := filepath.Ext(p1); ext != ".yaml" {
		t.Errorf("cache path extension = %q, want .yaml", ext)
	}
}

func TestWriteAndReadCache(t *testing.T) {
	cacheDir := t.TempDir()
	f := NewRemoteFetcher(cacheDir)
	f.CacheTTL = 1 * time.Hour

	path := filepath.Join(cacheDir, "test.yaml")
	reg := &ToolRegistry{
		Version: "1",
		Tools: []ToolEntry{
			{Name: "tool-a", Type: TypeMCP, MCP: &MCPConfig{Command: "echo"}},
		},
	}

	if err := f.writeCache(path, reg); err != nil {
		t.Fatalf("writeCache: %v", err)
	}

	loaded, err := f.loadCache(path)
	if err != nil {
		t.Fatalf("loadCache: %v", err)
	}
	if loaded.Version != "1" {
		t.Errorf("got version %q, want %q", loaded.Version, "1")
	}
	if len(loaded.Tools) != 1 {
		t.Errorf("got %d tools, want 1", len(loaded.Tools))
	}
}

func TestFetchRemoteRegistry_Convenience(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(remoteRegistryYAML))
	}))
	defer srv.Close()

	// Clean up the default cache dir after test.
	defer os.RemoveAll(".tools-cache")

	reg, err := FetchRemoteRegistry(srv.URL + "/registry.yaml")
	if err != nil {
		t.Fatalf("FetchRemoteRegistry: %v", err)
	}
	if reg.Version != "1" {
		t.Errorf("got version %q, want %q", reg.Version, "1")
	}
}

func TestLoadRemoteRegistry_DelegatesToFetcher(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(remoteRegistryYAML))
	}))
	defer srv.Close()

	defer os.RemoveAll(".tools-cache")

	reg, err := LoadRemoteRegistry(srv.URL + "/registry.yaml")
	if err != nil {
		t.Fatalf("LoadRemoteRegistry: %v", err)
	}
	if len(reg.Tools) != 2 {
		t.Errorf("got %d tools, want 2", len(reg.Tools))
	}
}
