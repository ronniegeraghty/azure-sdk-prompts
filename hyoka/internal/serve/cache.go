package serve

import (
	"os"
	"sync"
	"time"
)

// fileCache is an in-memory cache for JSON report files keyed by absolute
// path. Entries are invalidated when the file's on-disk modification time
// or size changes, so writers (a new eval run landing on disk) transparently
// bust the cache without any explicit invalidation.
//
// The cache is safe for concurrent use by multiple HTTP handlers.
type fileCache struct {
	mu    sync.RWMutex
	items map[string]cacheEntry
}

// cacheEntry holds the cached payload plus the file stat fingerprint that
// was valid when the entry was populated.
type cacheEntry struct {
	mtime time.Time
	size  int64
	data  []byte
}

func newFileCache() *fileCache {
	return &fileCache{items: make(map[string]cacheEntry)}
}

// ReadFile returns the contents of path, serving from cache when the file's
// mtime and size match the cached fingerprint. On a miss or stale entry it
// re-reads from disk and updates the cache.
//
// Any I/O error is returned to the caller; the cache entry is evicted so the
// next call will re-attempt the read rather than serve stale bytes.
func (c *fileCache) ReadFile(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		c.invalidate(path)
		return nil, err
	}

	mtime := info.ModTime()
	size := info.Size()

	c.mu.RLock()
	entry, ok := c.items[path]
	c.mu.RUnlock()
	if ok && entry.mtime.Equal(mtime) && entry.size == size {
		return entry.data, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		c.invalidate(path)
		return nil, err
	}

	c.mu.Lock()
	c.items[path] = cacheEntry{mtime: mtime, size: size, data: data}
	c.mu.Unlock()
	return data, nil
}

// invalidate removes a single entry. Safe to call for paths that aren't cached.
func (c *fileCache) invalidate(path string) {
	c.mu.Lock()
	delete(c.items, path)
	c.mu.Unlock()
}

// len returns the number of cached entries. Intended for tests.
func (c *fileCache) len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
