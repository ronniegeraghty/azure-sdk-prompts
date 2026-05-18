package artifact

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// GeneratorArtifact captures the complete state of a generator session for
// consumption by graders and serve-time inspection. This is written as a
// stable JSON file in the eval's report directory after generation completes
// (success or failure) and before graders are invoked.
//
// Graders — both typed and AI review — receive a path to this file via
// GraderInput.GeneratorArtifactPath and a pre-parsed pointer via
// GraderInput.GeneratorArtifact.
type GeneratorArtifact struct {
	// Eval metadata
	PromptID       string `json:"prompt_id"`
	EvalID         string `json:"eval_id,omitempty"`
	ConfigName     string `json:"config_name"`
	GeneratorModel string `json:"generator_model"`

	// Prompt and agent response
	OriginalPrompt string `json:"original_prompt"`
	FinalResponse  string `json:"final_response"` // Last assistant message

	// Workspace changes (what the agent did)
	WorkspaceDelta ArtifactWorkspaceDelta `json:"workspace_delta"`

	// Session metrics
	ActionsSummary ActionsSummary `json:"actions_summary"`

	// Timing
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	DurationMs int64     `json:"duration_ms"`

	// Termination info
	TerminatedBy string `json:"terminated_by"` // "completed", "max_actions", "max_turns", "guardrail", "timeout", "error"
	Error        string `json:"error,omitempty"` // Populated when terminated_by indicates failure
}

// ArtifactWorkspaceDelta is a shallow struct mirroring workspace.WorkspaceDelta
// for JSON serialization in the generator artifact.
type ArtifactWorkspaceDelta struct {
	BytesAdded   int64 `json:"bytes_added"`
	BytesRemoved int64 `json:"bytes_removed"`
	BytesNet     int64 `json:"bytes_net"`

	NewFileCount      int `json:"new_file_count"`
	ModifiedFileCount int `json:"modified_file_count"`
	DeletedFileCount  int `json:"deleted_file_count"`

	CreatedFiles  []ArtifactFileInfo `json:"created_files"`
	ModifiedFiles []ArtifactFileInfo `json:"modified_files"`
	DeletedFiles  []ArtifactFileInfo `json:"deleted_files"`
}

// ArtifactFileInfo describes a file in the workspace delta artifact.
type ArtifactFileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// ActionsSummary captures high-level session activity metrics.
type ActionsSummary struct {
	TotalActions   int  `json:"total_actions"`
	ToolCalls      int  `json:"tool_calls"`
	ReasoningSteps int  `json:"reasoning_steps"`
	Truncated      bool `json:"truncated"` // True if action limit was hit
}

// WriteToFile marshals the artifact as JSON and writes it to the specified path.
// Parent directories are created if they don't exist.
func (a *GeneratorArtifact) WriteToFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating artifact directory: %w", err)
	}

	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling generator artifact: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing generator artifact: %w", err)
	}

	return nil
}

// LoadGeneratorArtifact reads and unmarshals a generator artifact from disk.
func LoadGeneratorArtifact(path string) (*GeneratorArtifact, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading generator artifact: %w", err)
	}

	var artifact GeneratorArtifact
	if err := json.Unmarshal(data, &artifact); err != nil {
		return nil, fmt.Errorf("unmarshaling generator artifact: %w", err)
	}

	return &artifact, nil
}

// TruncateField truncates a string field to maxSize bytes and appends a marker.
func TruncateField(s string, maxSize int) string {
	if len(s) <= maxSize {
		return s
	}
	return s[:maxSize] + "\n[truncated]"
}

// FromWorkspaceDelta converts a workspace.WorkspaceDelta to the artifact's
// embedded delta representation.
func FromWorkspaceDelta(wsDelta *workspace.WorkspaceDelta) ArtifactWorkspaceDelta {
	if wsDelta == nil {
		return ArtifactWorkspaceDelta{}
	}

	delta := ArtifactWorkspaceDelta{
		BytesAdded:        wsDelta.BytesAdded,
		BytesRemoved:      wsDelta.BytesRemoved,
		BytesNet:          wsDelta.BytesNet,
		NewFileCount:      wsDelta.NewFileCount,
		ModifiedFileCount: wsDelta.ModifiedFileCount,
		DeletedFileCount:  wsDelta.DeletedFileCount,
	}

	for _, nf := range wsDelta.NewFiles {
		delta.CreatedFiles = append(delta.CreatedFiles, ArtifactFileInfo{
			Path: nf.Path,
			Size: nf.Size,
		})
	}

	for _, mf := range wsDelta.ModifiedFiles {
		delta.ModifiedFiles = append(delta.ModifiedFiles, ArtifactFileInfo{
			Path: mf.Path,
			Size: mf.SizeAfter,
		})
	}

	for _, df := range wsDelta.DeletedFiles {
		delta.DeletedFiles = append(delta.DeletedFiles, ArtifactFileInfo{
			Path: df.Path,
			Size: df.OriginalSize,
		})
	}

	return delta
}
