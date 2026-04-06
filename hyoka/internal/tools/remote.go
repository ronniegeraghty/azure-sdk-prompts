package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// DefaultCacheTTL is the default time-to-live for cached registries.
const DefaultCacheTTL = 1 * time.Hour

// DefaultHTTPTimeout is the default timeout for HTTP requests.
const DefaultHTTPTimeout = 30 * time.Second

// remoteMaxSize is the maximum allowed size of a remote registry (1 MB).
const remoteMaxSize = 1 << 20

// RemoteFetcher fetches and caches tool registries from remote URLs.
type RemoteFetcher struct {
	// CacheDir is the directory where cached registries are stored.
	CacheDir string

	// CacheTTL controls how long cached registries remain valid.
	CacheTTL time.Duration

	// HTTPClient is the HTTP client used for fetching. If nil, a default
	// client with DefaultHTTPTimeout is used.
	HTTPClient *http.Client

	mu sync.Mutex
}

// NewRemoteFetcher returns a RemoteFetcher with sensible defaults.
func NewRemoteFetcher(cacheDir string) *RemoteFetcher {
	if cacheDir == "" {
		cacheDir = ".tools-cache"
	}
	return &RemoteFetcher{
		CacheDir: cacheDir,
		CacheTTL: DefaultCacheTTL,
		HTTPClient: &http.Client{
			Timeout: DefaultHTTPTimeout,
		},
	}
}

// FetchRemoteRegistry fetches a tool registry YAML from url, parses it,
// and returns the resulting ToolRegistry. Results are cached locally;
// subsequent calls within the TTL window return the cached version.
// On network failure the stale cache is returned when available.
func FetchRemoteRegistry(url string) (*ToolRegistry, error) {
	f := NewRemoteFetcher("")
	return f.Fetch(url)
}

// Fetch retrieves and parses the registry at the given URL.
func (f *RemoteFetcher) Fetch(url string) (*ToolRegistry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	url = NormalizeGitHubURL(url)
	cachePath := f.cachePath(url)

	// Check cache freshness.
	if reg, err := f.loadCache(cachePath); err == nil {
		slog.Debug("tools: using cached registry", "url", url)
		return reg, nil
	}

	// Fetch from remote.
	reg, fetchErr := f.fetchHTTP(url)
	if fetchErr != nil {
		// Graceful degradation: return stale cache if available.
		if stale, cacheErr := f.loadStaleCache(cachePath); cacheErr == nil {
			slog.Warn("tools: network error, using stale cache",
				"url", url, "error", fetchErr)
			return stale, nil
		}
		return nil, fmt.Errorf("fetching remote registry %s: %w", url, fetchErr)
	}

	// Persist to cache.
	if err := f.writeCache(cachePath, reg); err != nil {
		slog.Warn("tools: failed to cache registry", "url", url, "error", err)
	}

	return reg, nil
}

// fetchHTTP performs the HTTP GET, size-limited read, and YAML parse+validate.
func (f *RemoteFetcher) fetchHTTP(url string) (*ToolRegistry, error) {
	resp, err := f.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, remoteMaxSize+1))
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	if len(body) > remoteMaxSize {
		return nil, fmt.Errorf("registry exceeds maximum size of %d bytes", remoteMaxSize)
	}

	return ParseRegistry(body)
}

// NormalizeGitHubURL converts common GitHub URL patterns to raw content URLs.
// For example:
//
//	https://github.com/owner/repo/blob/main/tools.yaml
//
// becomes:
//
//	https://raw.githubusercontent.com/owner/repo/main/tools.yaml
func NormalizeGitHubURL(url string) string {
	if strings.Contains(url, "raw.githubusercontent.com") {
		return url
	}

	const ghPrefix = "https://github.com/"
	if strings.HasPrefix(url, ghPrefix) {
		rest := strings.TrimPrefix(url, ghPrefix)
		if idx := strings.Index(rest, "/blob/"); idx >= 0 {
			ownerRepo := rest[:idx]
			refAndPath := rest[idx+len("/blob/"):]
			return "https://raw.githubusercontent.com/" + ownerRepo + "/" + refAndPath
		}
	}

	return url
}

// cachePath returns a deterministic file path for the given URL.
func (f *RemoteFetcher) cachePath(url string) string {
	h := sha256.Sum256([]byte(url))
	name := hex.EncodeToString(h[:16]) + ".yaml"
	return filepath.Join(f.CacheDir, name)
}

// loadCache loads a cached registry only if it is fresh (within TTL).
func (f *RemoteFetcher) loadCache(path string) (*ToolRegistry, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if time.Since(info.ModTime()) > f.CacheTTL {
		return nil, fmt.Errorf("cache expired")
	}
	return f.readCacheFile(path)
}

// loadStaleCache loads a cached registry regardless of TTL.
func (f *RemoteFetcher) loadStaleCache(path string) (*ToolRegistry, error) {
	return f.readCacheFile(path)
}

func (f *RemoteFetcher) readCacheFile(path string) (*ToolRegistry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var reg ToolRegistry
	if err := yaml.Unmarshal(data, &reg); err != nil {
		return nil, err
	}
	return &reg, nil
}

// writeCache serialises the registry to disk.
func (f *RemoteFetcher) writeCache(path string, reg *ToolRegistry) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := yaml.Marshal(reg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
