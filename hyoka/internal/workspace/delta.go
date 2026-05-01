// Package workspace provides types and utilities for tracking file-level changes
// in evaluation workspaces. It captures what the agent created, modified, and
// deleted during a session — distinct from starter files the agent inherited.
package workspace

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/utils"
)

// WorkspaceDelta captures file-level changes made by the agent during an
// evaluation session. It distinguishes agent contributions from inherited
// starter files, enabling graders to reason about what the agent actually did
// and providing a principled basis for guardrail warnings (Issue #566).
type WorkspaceDelta struct {
	// Byte metrics
	BytesAdded   int64 `json:"bytes_added"`   // Total bytes in new files
	BytesRemoved int64 `json:"bytes_removed"` // Total bytes from deleted files
	BytesNet     int64 `json:"bytes_net"`     // BytesAdded - BytesRemoved + net changes to modified files

	// File counts
	NewFileCount      int `json:"new_file_count"`
	ModifiedFileCount int `json:"modified_file_count"`
	DeletedFileCount  int `json:"deleted_file_count"`

	// Detailed file lists
	NewFiles      []NewFile      `json:"new_files"`
	ModifiedFiles []ModifiedFile `json:"modified_files"`
	DeletedFiles  []DeletedFile  `json:"deleted_files"`
}

// NewFile describes a file created by the agent.
type NewFile struct {
	Path string `json:"path"` // Relative to workspace root
	Size int64  `json:"size"`
	Hash string `json:"hash"` // SHA-256 hex
}

// ModifiedFile describes a file the agent modified.
type ModifiedFile struct {
	Path       string `json:"path"`
	SizeBefore int64  `json:"size_before"`
	SizeAfter  int64  `json:"size_after"`
	HashAfter  string `json:"hash_after"` // SHA-256 hex of final content
}

// DeletedFile describes a file the agent deleted.
type DeletedFile struct {
	Path         string `json:"path"`
	OriginalSize int64  `json:"original_size"`
}

// Snapshot captures the workspace state at a point in time: file paths
// mapped to their sizes and content hashes. This is the "before" picture
// used to compute deltas.
type Snapshot struct {
	files map[string]fileInfo
}

type fileInfo struct {
	size int64
	hash string
}

// TakeSnapshot walks the workspace directory and records each file's size
// and SHA-256 hash. Hidden files and directories are excluded. Returns an
// error if the directory cannot be walked, but individual file read/hash
// failures are logged and skipped (the file is omitted from the snapshot).
func TakeSnapshot(workspaceDir string) (*Snapshot, error) {
	snap := &Snapshot{
		files: make(map[string]fileInfo),
	}

	err := filepath.Walk(workspaceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden entries
		if info.Name() != filepath.Base(workspaceDir) && len(info.Name()) > 0 && info.Name()[0] == '.' {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip well-known build artifact directories (target/, node_modules/, etc.)
		if info.IsDir() && utils.IsDefaultExcludedDir(info.Name()) {
			return filepath.SkipDir
		}

		// Only capture regular files
		if !info.Mode().IsRegular() {
			return nil
		}

		rel, err := filepath.Rel(workspaceDir, path)
		if err != nil {
			return err
		}

		hash, err := hashFile(path)
		if err != nil {
			// Log but don't fail the entire snapshot for one unreadable file
			return nil
		}

		snap.files[rel] = fileInfo{
			size: info.Size(),
			hash: hash,
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walking workspace: %w", err)
	}

	return snap, nil
}

// ComputeDelta compares the "after" snapshot to the "before" snapshot and
// returns a WorkspaceDelta describing what changed. Files present in "after"
// but not "before" are new; files in "before" but not "after" are deleted;
// files in both are checked for content changes (via hash comparison).
func ComputeDelta(before, after *Snapshot) *WorkspaceDelta {
	delta := &WorkspaceDelta{}

	// Track which files in "before" we've seen in "after"
	seen := make(map[string]bool)

	// Compare after to before: find new and modified files
	for path, afterInfo := range after.files {
		seen[path] = true

		beforeInfo, existed := before.files[path]
		if !existed {
			// New file
			delta.NewFiles = append(delta.NewFiles, NewFile{
				Path: path,
				Size: afterInfo.size,
				Hash: afterInfo.hash,
			})
			delta.NewFileCount++
			delta.BytesAdded += afterInfo.size
			continue
		}

		// File existed — check if modified
		if beforeInfo.hash != afterInfo.hash {
			delta.ModifiedFiles = append(delta.ModifiedFiles, ModifiedFile{
				Path:       path,
				SizeBefore: beforeInfo.size,
				SizeAfter:  afterInfo.size,
				HashAfter:  afterInfo.hash,
			})
			delta.ModifiedFileCount++
			// Net change to modified file
			sizeDelta := afterInfo.size - beforeInfo.size
			if sizeDelta > 0 {
				delta.BytesAdded += sizeDelta
			} else if sizeDelta < 0 {
				delta.BytesRemoved += -sizeDelta
			}
		}
	}

	// Find deleted files (in "before" but not "after")
	for path, beforeInfo := range before.files {
		if !seen[path] {
			delta.DeletedFiles = append(delta.DeletedFiles, DeletedFile{
				Path:         path,
				OriginalSize: beforeInfo.size,
			})
			delta.DeletedFileCount++
			delta.BytesRemoved += beforeInfo.size
		}
	}

	// Compute net bytes
	delta.BytesNet = int64(delta.BytesAdded) - int64(delta.BytesRemoved)

	return delta
}

// hashFile computes the SHA-256 hash of a file's contents.
func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
