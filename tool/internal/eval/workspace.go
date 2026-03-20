package eval

import (
"fmt"
"os"
"path/filepath"
)

// Workspace manages a temporary directory for an evaluation run.
type Workspace struct {
BaseDir string
Dir     string
}

// NewWorkspace creates a new workspace directory under baseDir.
func NewWorkspace(baseDir, promptID, configName string) (*Workspace, error) {
dirName := fmt.Sprintf("%s_%s", promptID, configName)
dir := filepath.Join(baseDir, dirName)

if err := os.MkdirAll(dir, 0755); err != nil {
return nil, fmt.Errorf("creating workspace directory: %w", err)
}

return &Workspace{
BaseDir: baseDir,
Dir:     dir,
}, nil
}

// Cleanup removes the workspace directory.
func (w *Workspace) Cleanup() error {
return os.RemoveAll(w.Dir)
}

// ListFiles returns all files in the workspace.
func (w *Workspace) ListFiles() ([]string, error) {
var files []string
err := filepath.Walk(w.Dir, func(path string, info os.FileInfo, err error) error {
if err != nil {
return err
}
if !info.IsDir() {
rel, err := filepath.Rel(w.Dir, path)
if err != nil {
return err
}
files = append(files, rel)
}
return nil
})
return files, err
}
