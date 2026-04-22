package eval

import (
	"log/slog"
	"os"
	"path/filepath"
)

// snapshotStarterSizes records the on-disk size of each starter file, keyed by
// the same workspace-relative path format that Workspace.ListFiles returns.
// Files that cannot be stat'd are recorded with size 0 — the safe default,
// since any later growth will be counted as agent output (worst case: no change
// from today's behavior). baseDir is the workspace root.
func snapshotStarterSizes(baseDir string, starterFiles []string) map[string]int64 {
	snap := make(map[string]int64, len(starterFiles))
	for _, rel := range starterFiles {
		info, err := os.Stat(filepath.Join(baseDir, rel))
		if err != nil {
			slog.Warn("Could not stat starter file for guardrail snapshot; treating original size as 0",
				"file", rel, "error", err)
			snap[rel] = 0
			continue
		}
		snap[rel] = info.Size()
	}
	return snap
}

// computeAgentFileCount returns the number of files the agent is responsible
// for: new files (not in the snapshot) plus starter files the agent deleted
// (present in snapshot but missing from the current list).
func computeAgentFileCount(files []string, snapshot map[string]int64) int {
	present := make(map[string]bool, len(files))
	newFiles := 0
	for _, f := range files {
		present[f] = true
		if _, isStarter := snapshot[f]; !isStarter {
			newFiles++
		}
	}
	deleted := 0
	for starter := range snapshot {
		if !present[starter] {
			deleted++
		}
	}
	return newFiles + deleted
}
