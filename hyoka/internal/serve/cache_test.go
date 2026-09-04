package serve

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestFileCache_ReadThenCacheHit(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")
	if err := os.WriteFile(path, []byte(`{"a":1}`), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	c := newFileCache()
	first, err := c.ReadFile(path)
	if err != nil {
		t.Fatalf("first read: %v", err)
	}
	if string(first) != `{"a":1}` {
		t.Fatalf("unexpected bytes: %s", first)
	}
	if c.len() != 1 {
		t.Fatalf("expected 1 cache entry, got %d", c.len())
	}

	// Delete the underlying file — a cache hit should still succeed.
	if err := os.Remove(path); err != nil {
		t.Fatalf("remove: %v", err)
	}
	// Re-create with identical content AND restore the original mtime so the
	// cache fingerprint still matches. If mtime differs, the cache busts — which
	// is correct, but not what this test covers (see TestFileCache_MtimeChange).
	c.mu.RLock()
	entry := c.items[path]
	c.mu.RUnlock()
	if err := os.WriteFile(path, []byte(`{"a":1}`), 0644); err != nil {
		t.Fatalf("rewrite: %v", err)
	}
	if err := os.Chtimes(path, entry.mtime, entry.mtime); err != nil {
		t.Fatalf("chtimes: %v", err)
	}

	second, err := c.ReadFile(path)
	if err != nil {
		t.Fatalf("second read: %v", err)
	}
	if string(second) != `{"a":1}` {
		t.Fatalf("unexpected bytes on hit: %s", second)
	}
}

func TestFileCache_MtimeChangeBustsCache(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")
	if err := os.WriteFile(path, []byte(`{"v":1}`), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	c := newFileCache()
	if _, err := c.ReadFile(path); err != nil {
		t.Fatalf("first read: %v", err)
	}

	// Rewrite with new content and forced newer mtime.
	if err := os.WriteFile(path, []byte(`{"v":2}`), 0644); err != nil {
		t.Fatalf("rewrite: %v", err)
	}
	future := time.Now().Add(2 * time.Second)
	if err := os.Chtimes(path, future, future); err != nil {
		t.Fatalf("chtimes: %v", err)
	}

	data, err := c.ReadFile(path)
	if err != nil {
		t.Fatalf("second read: %v", err)
	}
	if string(data) != `{"v":2}` {
		t.Fatalf("expected fresh bytes, got %s", data)
	}
}

func TestFileCache_MissingFileEvictsEntry(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")
	if err := os.WriteFile(path, []byte(`{}`), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	c := newFileCache()
	if _, err := c.ReadFile(path); err != nil {
		t.Fatalf("first read: %v", err)
	}
	if c.len() != 1 {
		t.Fatalf("expected cache populated")
	}

	if err := os.Remove(path); err != nil {
		t.Fatalf("remove: %v", err)
	}
	if _, err := c.ReadFile(path); err == nil {
		t.Fatalf("expected error after file removed")
	}
	if c.len() != 0 {
		t.Fatalf("expected cache evicted on error, got %d entries", c.len())
	}
}

func TestFileCache_ConcurrentReaders(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")
	payload := []byte(`{"hello":"world"}`)
	if err := os.WriteFile(path, payload, 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	c := newFileCache()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				data, err := c.ReadFile(path)
				if err != nil {
					t.Errorf("read: %v", err)
					return
				}
				if string(data) != string(payload) {
					t.Errorf("unexpected bytes: %s", data)
					return
				}
			}
		}()
	}
	wg.Wait()
}
