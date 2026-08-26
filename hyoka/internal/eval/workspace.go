// Package eval provides the core evaluation engine, workspace management, and Copilot interaction logic.
package eval

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/utils"
)

// EvalWorkspacePrefix is the directory name prefix used for isolated evaluation
// workspaces. The clean command uses this to find orphan workspaces.
const EvalWorkspacePrefix = "hyoka-eval-"

// EvalWorkspace is an isolated per-session workspace directory. Each evaluation
// session gets its own empty workspace so the Copilot agent cannot see or
// modify files from the user's development environment. The workspace is
// created before the session starts and removed after it completes.
type EvalWorkspace struct {
	// Dir is the absolute path to the isolated workspace directory.
	Dir string
	// CreatedAt records when the workspace was created, used by the clean
	// command to identify stale orphan workspaces.
	CreatedAt time.Time
}

// NewEvalWorkspace creates an empty isolated workspace in the OS temp
// directory. The directory name begins with EvalWorkspacePrefix so the clean
// command can discover orphan workspaces after a crash.
func NewEvalWorkspace() (*EvalWorkspace, error) {
	dir, err := os.MkdirTemp("", EvalWorkspacePrefix)
	if err != nil {
		return nil, fmt.Errorf("creating isolated eval workspace: %w", err)
	}
	return &EvalWorkspace{
		Dir:       dir,
		CreatedAt: time.Now(),
	}, nil
}

// CopyStarterFiles copies a starter project into the workspace.
func (w *EvalWorkspace) CopyStarterFiles(starterDir string) error {
	return copyDir(starterDir, w.Dir)
}

// ListFiles returns all non-hidden files in the workspace, relative to its root.
func (w *EvalWorkspace) ListFiles() ([]string, error) {
	return listFiles(w.Dir)
}

// Cleanup removes the workspace directory. Safe to call multiple times.
func (w *EvalWorkspace) Cleanup() error {
	if w == nil || w.Dir == "" {
		return nil
	}
	return os.RemoveAll(w.Dir)
}

// Workspace manages a directory for an evaluation run.
type Workspace struct {
	BaseDir string
	Dir     string
	persist bool // if true, Cleanup is a no-op
}

// NewWorkspace creates a new workspace in the OS temp directory.
// The workspace is ephemeral — generated files should be copied out
// before calling Cleanup.
func NewWorkspace(promptID, configName string) (*Workspace, error) {
	safeConfig := strings.ReplaceAll(configName, "/", "-")
	prefix := fmt.Sprintf("hyoka-%s-%s-", promptID, safeConfig)
	dir, err := os.MkdirTemp("", prefix)
	if err != nil {
		return nil, fmt.Errorf("creating temp workspace: %w", err)
	}

	return &Workspace{
		BaseDir: os.TempDir(),
		Dir:     dir,
	}, nil
}

// NewWorkspaceAt creates a persistent workspace at the given directory.
// The directory is created if it doesn't exist. Cleanup is a no-op.
func NewWorkspaceAt(dir string) (*Workspace, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("creating workspace directory: %w", err)
	}
	return &Workspace{
		BaseDir: filepath.Dir(dir),
		Dir:     dir,
		persist: true,
	}, nil
}

// Cleanup removes the workspace directory (no-op for persistent workspaces).
func (w *Workspace) Cleanup() error {
	if w.persist {
		return nil
	}
	return os.RemoveAll(w.Dir)
}

// ListFiles returns all non-hidden files in the workspace, relative to its root.
func (w *Workspace) ListFiles() ([]string, error) {
	return listFiles(w.Dir)
}

// CopyFilesTo copies all non-hidden workspace files into destDir,
// preserving relative paths. Returns the list of files copied.
func (w *Workspace) CopyFilesTo(destDir string) ([]string, error) {
	files, err := w.ListFiles()
	if err != nil {
		return nil, fmt.Errorf("listing workspace files: %w", err)
	}
	if len(files) == 0 {
		return nil, nil
	}

	for _, rel := range files {
		src := filepath.Join(w.Dir, rel)
		dst := filepath.Join(destDir, rel)

		if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
			return nil, fmt.Errorf("creating dir for %s: %w", rel, err)
		}
		data, err := os.ReadFile(src)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", rel, err)
		}
		if err := os.WriteFile(dst, data, 0644); err != nil {
			return nil, fmt.Errorf("writing %s: %w", rel, err)
		}
	}

	return files, nil
}

// CopyStarterFiles copies the starter project declared in the prompt into
// the workspace directory. Only files from the declared starter directory are
// copied — hidden files, symlinks, and build artifacts are excluded by copyDir.
// Returns the list of files copied (relative to the workspace root) and a nil
// error on success. Returns (nil, nil) when the prompt has no starter project.
func (w *Workspace) CopyStarterFiles(p *prompt.Prompt) ([]string, error) {
	if p.StarterProject == "" {
		return nil, nil
	}
	starterDir := resolveStarterDir(p)
	info, err := os.Stat(starterDir)
	if err != nil {
		return nil, fmt.Errorf("starter project %q: %w", p.StarterProject, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("starter project %q is not a directory", p.StarterProject)
	}
	if err := copyDir(starterDir, w.Dir); err != nil {
		return nil, fmt.Errorf("copying starter project: %w", err)
	}
	files, err := listFiles(w.Dir)
	if err != nil {
		return nil, fmt.Errorf("listing starter files: %w", err)
	}
	if len(files) == 0 {
		slog.Warn("Starter project directory contains no files", "dir", starterDir)
	}
	return files, nil
}

// resolveStarterDir returns the absolute path to a prompt's starter project
// directory. If StarterProject is relative, it is resolved relative to the
// prompt file's directory.
func resolveStarterDir(p *prompt.Prompt) string {
	dir := p.StarterProject
	if !filepath.IsAbs(dir) && p.FilePath != "" {
		dir = filepath.Join(filepath.Dir(p.FilePath), dir)
	}
	return dir
}

// listFiles is a helper used by Workspace and PromptRunner implementations.
func listFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && path != dir {
				return filepath.SkipDir
			}
			if utils.IsDefaultExcludedDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		files = append(files, rel)
		return nil
	})
	return files, err
}

// filterExcludedDirs removes files whose path contains any of the
// excluded directory names as a segment (#63). For example, excluding
// "node_modules" matches "node_modules/foo/bar.js" AND
// "project/node_modules/bar.js".
func filterExcludedDirs(files []string, excludeDirs []string) []string {
	if len(excludeDirs) == 0 {
		return files
	}
	excludeSet := make(map[string]bool, len(excludeDirs))
	for _, d := range excludeDirs {
		excludeSet[strings.TrimRight(d, "/")] = true
	}
	var filtered []string
	for _, f := range files {
		excluded := false
		parts := strings.Split(filepath.ToSlash(f), "/")
		for _, seg := range parts {
			if excludeSet[seg] {
				excluded = true
				break
			}
		}
		if !excluded {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

// copyDir recursively copies src to dst, skipping symlinks, hidden dirs,
// and well-known build artifact directories to prevent escape and bloat.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip symlinks to prevent following links outside the source tree
		if info.Mode()&os.ModeSymlink != 0 {
			rel, _ := filepath.Rel(src, path)
			slog.Warn("Skipping symlink in starter project", "path", rel)
			return nil
		}
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") && path != src {
				return filepath.SkipDir
			}
			if utils.IsDefaultExcludedDir(name) {
				slog.Debug("Skipping excluded directory", "dir", name)
				return filepath.SkipDir
			}
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, info.Mode())
	})
}

// NewIsolatedConfigDir creates an empty temporary directory to serve as the
// Copilot CLI configuration directory. By pointing ConfigDirectory at this empty
// directory, user-level skills and settings from ~/.config/github-copilot/
// are excluded from eval sessions. Only skills explicitly listed in the eval
// config's SkillDirectories are loaded (fixes #21).
// The caller must defer os.RemoveAll on the returned path.
func NewIsolatedConfigDir() (string, error) {
	dir, err := os.MkdirTemp("", "hyoka-config-*")
	if err != nil {
		return "", fmt.Errorf("creating isolated config dir: %w", err)
	}
	return dir, nil
}

// IsolateGraderWorkspace creates an ephemeral copy of sourceDir and returns
// the path together with a cleanup function. Each grader is given its own
// isolated copy so mutating graders (e.g. program graders running `make` or
// `npm install`) cannot pollute the workspace seen by subsequent graders.
// The original sourceDir is the canonical snapshot and remains untouched.
func IsolateGraderWorkspace(sourceDir string) (string, func(), error) {
	dir, err := NewReviewerWorkspace(sourceDir)
	if err != nil {
		return "", func() {}, err
	}
	return dir, func() { _ = os.RemoveAll(dir) }, nil
}

// NewReviewerWorkspace creates an ephemeral temporary workspace and copies
// all files from sourceDir into it. Reviewers operate on this copy so they
// cannot modify the original generated output (fixes #26).
// The caller must defer os.RemoveAll on the returned path.
func NewReviewerWorkspace(sourceDir string) (string, error) {
	dir, err := os.MkdirTemp("", "hyoka-review-*")
	if err != nil {
		return "", fmt.Errorf("creating reviewer workspace: %w", err)
	}
	if err := copyDir(sourceDir, dir); err != nil {
		os.RemoveAll(dir)
		return "", fmt.Errorf("copying files to reviewer workspace: %w", err)
	}
	return dir, nil
}

// ValidateWorkspaceContainment checks whether any new items appeared in dir
// since the pre-eval snapshot. Returns the names of items that escaped the
// workspace boundary. Called after recoverMisplacedFiles as a safety net
// to catch anything recovery could not handle (fixes #26).
func ValidateWorkspaceContainment(dir string, preSnapshot map[string]bool) []string {
	if preSnapshot == nil {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var escaped []string
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		if !preSnapshot[e.Name()] {
			escaped = append(escaped, e.Name())
		}
	}
	return escaped
}

// codeFileExts lists extensions that indicate output files.
var codeFileExts = map[string]bool{
	".py": true, ".cs": true, ".java": true, ".go": true, ".rs": true,
	".ts": true, ".js": true, ".cpp": true, ".c": true, ".h": true,
	".csproj": true, ".sln": true, ".json": true, ".xml": true,
	".yaml": true, ".yml": true, ".toml": true, ".mod": true,
	".txt": true, ".md": true, ".gradle": true, ".kt": true,
	".swift": true, ".rb": true, ".sh": true, ".bat": true,
	".html": true, ".css": true, ".sum": true, ".lock": true,
	".cfg": true, ".ini": true, ".env": true, ".dockerfile": true,
}

// DefaultIgnoreDirs is the centralized list of directory names that contain
// build artifacts, installed dependencies, or runtime caches. These directories
// should be excluded from review prompts (to avoid context overflow) and
// deleted when recovering misplaced files. Covers all supported languages.
var DefaultIgnoreDirs = map[string]bool{
	// JavaScript / TypeScript
	"node_modules":     true,
	"bower_components": true,
	".next":            true,
	".nuxt":            true,
	// Python
	"__pycache__":   true,
	"venv":          true,
	".venv":         true,
	"env":           true,
	".tox":          true,
	".eggs":         true,
	"site-packages": true,
	// Rust
	"target": true,
	// Go
	"vendor": true,
	// Java / Kotlin
	".gradle": true,
	".m2":     true,
	// C# / .NET
	"bin":      true,
	"obj":      true,
	"packages": true,
	// General build artifacts
	"dist":   true,
	"build":  true,
	".cache": true,
	"tmp":    true,
	".tmp":   true,
}

// junkDirs is an alias for DefaultIgnoreDirs, used by recoverMisplacedFiles
// to decide which directories to delete rather than recover.
var junkDirs = DefaultIgnoreDirs

// snapshotDir returns a set of non-hidden entry names (files AND directories) in
// a directory (non-recursive). Capturing directories lets recoverMisplacedFiles
// detect new directories created during an eval run.
func snapshotDir(dir string) map[string]bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	names := make(map[string]bool, len(entries))
	for _, e := range entries {
		if !strings.HasPrefix(e.Name(), ".") {
			names[e.Name()] = true
		}
	}
	return names
}

// recoverMisplacedFiles moves files and directories that appeared in dir since
// the snapshot into destDir. Files with recognized code extensions (or no
// extension) are moved; new directories are either moved into the workspace or
// deleted if they match a known junk pattern. Returns the count of recovered
// items (files + directories moved or cleaned up).
func recoverMisplacedFiles(dir string, preSnapshot map[string]bool, destDir string, _ string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}
	recovered := 0
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		if preSnapshot[e.Name()] {
			continue // existed before eval
		}

		src := filepath.Join(dir, e.Name())

		if e.IsDir() {
			// Junk directories → just delete
			if junkDirs[e.Name()] {
				if err := os.RemoveAll(src); err == nil {
					recovered++
					slog.Debug("Deleted junk directory", "path", src)
				}
				continue
			}
			// Other new directories → move into workspace
			dst := filepath.Join(destDir, e.Name())
			if err := os.Rename(src, dst); err != nil {
				// Rename may fail across filesystems; fall back to copy+delete
				if err := copyDir(src, dst); err == nil {
					os.RemoveAll(src)
				} else {
					continue
				}
			}
			recovered++
			slog.Debug("Recovered misplaced directory", "src", src, "dst", dst)
			continue
		}

		// Regular file handling (unchanged logic)
		ext := strings.ToLower(filepath.Ext(e.Name()))
		// Also recover extensionless files like "Dockerfile", "Makefile"
		if !codeFileExts[ext] && ext != "" {
			continue
		}

		dst := filepath.Join(destDir, e.Name())
		data, err := os.ReadFile(src)
		if err != nil {
			continue
		}
		if err := os.WriteFile(dst, data, 0644); err != nil {
			continue
		}
		os.Remove(src) // clean up the misplaced file
		recovered++
		slog.Debug("Recovered misplaced file", "src", src, "dst", dst)
	}
	return recovered
}
