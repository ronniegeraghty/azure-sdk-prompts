package graders

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WorkspaceGrader evaluates workspace changes the agent made during a session
// (WorkspaceDelta: NewFiles, ModifiedFiles, DeletedFiles) against a set of
// configured checks.
//
// All checks are powered by WorkspaceDelta, not by manual workspace scanning.
// Six canonical check kinds:
//   - require_to_create: path must be in NewFiles
//   - forbidden_to_create: path must NOT be in NewFiles
//   - required_to_update: path must be in ModifiedFiles
//   - required_to_delete: path must be in DeletedFiles
//   - forbidden_to_delete: if files:["*"], DeletedFiles must be empty; else specific paths must not be deleted
//   - file: state present (exists on disk + optional size/content checks) or absent (not on disk)
//
// Grader passes iff ALL checks pass. Score is 1.0 on pass, 0.0 on fail (boolean grader).
type WorkspaceGrader struct {
	name string
	cfg  WorkspaceConfig
}

// NewWorkspaceGrader constructs a WorkspaceGrader from a parsed config.
// Returns an error only for structurally invalid config. An empty Checks slice
// is accepted; the grader will trivially pass at runtime.
func NewWorkspaceGrader(name string, cfg *WorkspaceConfig) (*WorkspaceGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("workspace grader %q: config is required", name)
	}
	// Validate check structure
	for i, check := range cfg.Checks {
		if err := validateWorkspaceCheck(check); err != nil {
			return nil, fmt.Errorf("workspace grader %q check %d: %w", name, i, err)
		}
	}
	return &WorkspaceGrader{name: name, cfg: *cfg}, nil
}

func validateWorkspaceCheck(check WorkspaceCheck) error {
	switch check.Kind {
	case "require_to_create", "forbidden_to_create", "required_to_update", "required_to_delete", "forbidden_to_delete":
		if len(check.Files) == 0 {
			return fmt.Errorf("kind %q requires non-empty files array", check.Kind)
		}
	case "file":
		if check.Name == "" {
			return fmt.Errorf("kind %q requires name field", check.Kind)
		}
		if check.State != "present" && check.State != "absent" {
			return fmt.Errorf("kind %q: state must be 'present' or 'absent'", check.Kind)
		}
		if check.State == "absent" {
			if check.MinBytes != nil || check.MaxBytes != nil || check.Contains != "" || check.Excludes != "" {
				return fmt.Errorf("kind %q: state=absent does not support min_bytes, max_bytes, contains, or excludes", check.Kind)
			}
		}
		if check.MinBytes != nil && check.MaxBytes != nil && *check.MinBytes > *check.MaxBytes {
			return fmt.Errorf("kind %q: min_bytes (%d) > max_bytes (%d)", check.Kind, *check.MinBytes, *check.MaxBytes)
		}
	default:
		return fmt.Errorf("unknown workspace check kind: %q", check.Kind)
	}
	return nil
}

// Kind returns the grader type identifier.
func (g *WorkspaceGrader) Kind() string { return "workspace" }

// Name returns the human-readable name.
func (g *WorkspaceGrader) Name() string { return g.name }

// Grade evaluates every configured check against input.WorkspaceDelta and
// returns a GraderResult whose Pass is the AND of every sub-check.
func (g *WorkspaceGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	if input.WorkspaceDelta == nil {
		return GraderResult{}, fmt.Errorf("workspace grader %q: WorkspaceDelta is nil", g.name)
	}

	// If no checks configured, trivially pass
	if len(g.cfg.Checks) == 0 {
		checks := []GraderCheck{{
			Label:   "no checks",
			Pass:    true,
			Message: "no workspace checks configured — trivially passed",
		}}
		return NewResult("workspace", g.name, input.Config, checks, "no checks configured", &GraderExtras{
			Workspace: &WorkspaceExtras{},
		}), nil
	}

	delta := input.WorkspaceDelta
	var graderChecks []GraderCheck

	// Build index for NewFiles, ModifiedFiles, DeletedFiles
	newFilesSet := make(map[string]bool)
	for _, f := range delta.NewFiles {
		newFilesSet[f.Path] = true
	}
	modifiedFilesSet := make(map[string]bool)
	for _, f := range delta.ModifiedFiles {
		modifiedFilesSet[f.Path] = true
	}
	deletedFilesSet := make(map[string]bool)
	for _, f := range delta.DeletedFiles {
		deletedFilesSet[f.Path] = true
	}

	for _, check := range g.cfg.Checks {
		switch check.Kind {
		case "require_to_create":
			for _, file := range check.Files {
				present := newFilesSet[file]
				label := fmt.Sprintf("require_to_create: %s", file)
				var msg string
				if !present {
					msg = fmt.Sprintf("file %q was not created", file)
				} else {
					msg = fmt.Sprintf("file %q was created", file)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    present,
					Message: msg,
				})
			}

		case "forbidden_to_create":
			for _, file := range check.Files {
				forbidden := newFilesSet[file]
				label := fmt.Sprintf("forbidden_to_create: %s", file)
				var msg string
				if forbidden {
					msg = fmt.Sprintf("file %q was created (forbidden)", file)
				} else {
					msg = fmt.Sprintf("file %q was not created", file)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    !forbidden,
					Message: msg,
				})
			}

		case "required_to_update":
			for _, file := range check.Files {
				updated := modifiedFilesSet[file]
				label := fmt.Sprintf("required_to_update: %s", file)
				var msg string
				if !updated {
					msg = fmt.Sprintf("file %q was not updated", file)
				} else {
					msg = fmt.Sprintf("file %q was updated", file)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    updated,
					Message: msg,
				})
			}

		case "required_to_delete":
			for _, file := range check.Files {
				deleted := deletedFilesSet[file]
				label := fmt.Sprintf("required_to_delete: %s", file)
				var msg string
				if !deleted {
					msg = fmt.Sprintf("file %q was not deleted", file)
				} else {
					msg = fmt.Sprintf("file %q was deleted", file)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    deleted,
					Message: msg,
				})
			}

		case "forbidden_to_delete":
			// Special case: files:["*"] means no deletions allowed at all
			if len(check.Files) == 1 && check.Files[0] == "*" {
				noDeletions := len(delta.DeletedFiles) == 0
				label := "forbidden_to_delete: *"
				var msg string
				if !noDeletions {
					msg = fmt.Sprintf("%d file(s) were deleted (forbidden)", len(delta.DeletedFiles))
				} else {
					msg = "no files were deleted"
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    noDeletions,
					Message: msg,
				})
			} else {
				// Specific files must not be deleted
				for _, file := range check.Files {
					deleted := deletedFilesSet[file]
					label := fmt.Sprintf("forbidden_to_delete: %s", file)
					var msg string
					if deleted {
						msg = fmt.Sprintf("file %q was deleted (forbidden)", file)
					} else {
						msg = fmt.Sprintf("file %q was not deleted", file)
					}
					graderChecks = append(graderChecks, GraderCheck{
						Label:   label,
						Pass:    !deleted,
						Message: msg,
					})
				}
			}

		case "file":
			// Check if file exists on disk
			filePath := filepath.Join(input.WorkspacePath, check.Name)
			fileInfo, err := os.Stat(filePath)
			exists := err == nil

			if check.State == "absent" {
				label := fmt.Sprintf("file: %s (state=absent)", check.Name)
				var msg string
				if exists {
					msg = fmt.Sprintf("file %q exists (expected absent)", check.Name)
				} else {
					msg = fmt.Sprintf("file %q does not exist", check.Name)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    !exists,
					Message: msg,
				})
			} else { // state == "present"
				label := fmt.Sprintf("file: %s (state=present)", check.Name)
				if !exists {
					graderChecks = append(graderChecks, GraderCheck{
						Label:   label,
						Pass:    false,
						Message: fmt.Sprintf("file %q does not exist", check.Name),
					})
					continue
				}

				// Size checks
				size := fileInfo.Size()
				sizeOK := true
				var sizeMsg string
				if check.MinBytes != nil && size < *check.MinBytes {
					sizeOK = false
					sizeMsg = fmt.Sprintf("size %d bytes < min %d bytes", size, *check.MinBytes)
				} else if check.MaxBytes != nil && size > *check.MaxBytes {
					sizeOK = false
					sizeMsg = fmt.Sprintf("size %d bytes > max %d bytes", size, *check.MaxBytes)
				}

				if !sizeOK {
					graderChecks = append(graderChecks, GraderCheck{
						Label:   label,
						Pass:    false,
						Message: sizeMsg,
					})
					continue
				}

				// Content checks (contains / excludes)
				if check.Contains != "" || check.Excludes != "" {
					content, err := os.ReadFile(filePath)
					if err != nil {
						graderChecks = append(graderChecks, GraderCheck{
							Label:   label,
							Pass:    false,
							Message: fmt.Sprintf("failed to read file: %v", err),
						})
						continue
					}
					contentStr := string(content)

					if check.Contains != "" && !strings.Contains(contentStr, check.Contains) {
						graderChecks = append(graderChecks, GraderCheck{
							Label:   label,
							Pass:    false,
							Message: fmt.Sprintf("file %q does not contain %q", check.Name, check.Contains),
						})
						continue
					}

					if check.Excludes != "" && strings.Contains(contentStr, check.Excludes) {
						graderChecks = append(graderChecks, GraderCheck{
							Label:   label,
							Pass:    false,
							Message: fmt.Sprintf("file %q contains forbidden text %q", check.Name, check.Excludes),
						})
						continue
					}
				}

				// All checks passed
				msg := fmt.Sprintf("file %q exists", check.Name)
				if check.MinBytes != nil || check.MaxBytes != nil {
					msg += fmt.Sprintf(" (size: %d bytes)", size)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    true,
					Message: msg,
				})
			}
		}
	}

	// Compute pass/fail count
	passed := 0
	for _, c := range graderChecks {
		if c.Pass {
			passed++
		}
	}

	msg := fmt.Sprintf("workspace checks: %d/%d passed", passed, len(graderChecks))

	// Build extras
	producedFiles := make([]FileEntry, 0, len(delta.NewFiles)+len(delta.ModifiedFiles))
	for _, f := range delta.NewFiles {
		producedFiles = append(producedFiles, FileEntry{Path: f.Path, Size: f.Size})
	}
	for _, f := range delta.ModifiedFiles {
		producedFiles = append(producedFiles, FileEntry{Path: f.Path, Size: f.SizeAfter})
	}

	extras := &GraderExtras{
		Workspace: &WorkspaceExtras{
			ProducedFiles: producedFiles,
		},
	}

	return NewResult("workspace", g.name, input.Config, graderChecks, msg, extras), nil
}
