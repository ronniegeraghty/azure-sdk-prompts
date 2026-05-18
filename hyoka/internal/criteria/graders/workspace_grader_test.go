package graders

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

func TestWorkspaceGrader(t *testing.T) {
	tests := []struct {
		name        string
		checks      []WorkspaceCheck
		delta       *workspace.WorkspaceDelta
		setupWS     func(dir string) error
		wantPass    bool
		wantMessage string
	}{
		{
			name: "require_to_create pass",
			checks: []WorkspaceCheck{
				{Kind: "require_to_create", Files: []string{"hello.md"}},
			},
			delta: &workspace.WorkspaceDelta{
				NewFiles: []workspace.NewFile{
					{Path: "hello.md", Size: 10},
				},
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "require_to_create fail",
			checks: []WorkspaceCheck{
				{Kind: "require_to_create", Files: []string{"hello.md"}},
			},
			delta: &workspace.WorkspaceDelta{
				NewFiles: []workspace.NewFile{},
			},
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name: "forbidden_to_create pass",
			checks: []WorkspaceCheck{
				{Kind: "forbidden_to_create", Files: []string{".env"}},
			},
			delta: &workspace.WorkspaceDelta{
				NewFiles: []workspace.NewFile{
					{Path: "hello.md", Size: 10},
				},
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "forbidden_to_create fail",
			checks: []WorkspaceCheck{
				{Kind: "forbidden_to_create", Files: []string{".env"}},
			},
			delta: &workspace.WorkspaceDelta{
				NewFiles: []workspace.NewFile{
					{Path: ".env", Size: 10},
				},
			},
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name: "required_to_update pass",
			checks: []WorkspaceCheck{
				{Kind: "required_to_update", Files: []string{"README.md"}},
			},
			delta: &workspace.WorkspaceDelta{
				ModifiedFiles: []workspace.ModifiedFile{
					{Path: "README.md", SizeAfter: 100},
				},
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "required_to_delete pass",
			checks: []WorkspaceCheck{
				{Kind: "required_to_delete", Files: []string{"old.txt"}},
			},
			delta: &workspace.WorkspaceDelta{
				DeletedFiles: []workspace.DeletedFile{
					{Path: "old.txt", OriginalSize: 50},
				},
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "forbidden_to_delete with wildcard pass",
			checks: []WorkspaceCheck{
				{Kind: "forbidden_to_delete", Files: []string{"*"}},
			},
			delta: &workspace.WorkspaceDelta{
				DeletedFiles: []workspace.DeletedFile{},
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "forbidden_to_delete with wildcard fail",
			checks: []WorkspaceCheck{
				{Kind: "forbidden_to_delete", Files: []string{"*"}},
			},
			delta: &workspace.WorkspaceDelta{
				DeletedFiles: []workspace.DeletedFile{
					{Path: "something.txt", OriginalSize: 10},
				},
			},
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name: "file state=present pass",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present"},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("content"), 0644)
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "file state=present fail",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present"},
			},
			delta:       &workspace.WorkspaceDelta{},
			setupWS:     func(dir string) error { return nil },
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name: "file state=absent pass",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "absent"},
			},
			delta:       &workspace.WorkspaceDelta{},
			setupWS:     func(dir string) error { return nil },
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "file state=present with min_bytes pass",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", MinBytes: int64Ptr(5)},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("hello world"), 0644)
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "file state=present with max_bytes fail",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", MaxBytes: int64Ptr(5)},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("hello world"), 0644)
			},
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name: "file state=present with contains pass",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", Contains: "# Hello"},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("# Hello World\ncontent"), 0644)
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "file state=present with excludes pass",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", Excludes: "TODO"},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("# Hello World\ncontent"), 0644)
			},
			wantPass:    true,
			wantMessage: "workspace checks: 1/1 passed",
		},
		{
			name: "file state=present with excludes fail",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", Excludes: "TODO"},
			},
			delta: &workspace.WorkspaceDelta{},
			setupWS: func(dir string) error {
				return os.WriteFile(filepath.Join(dir, "hello.md"), []byte("# TODO\ncontent"), 0644)
			},
			wantPass:    false,
			wantMessage: "workspace checks: 0/1 passed",
		},
		{
			name:        "no checks configured",
			checks:      []WorkspaceCheck{},
			delta:       &workspace.WorkspaceDelta{},
			wantPass:    true,
			wantMessage: "no checks configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp workspace
			tmpDir := t.TempDir()
			if tt.setupWS != nil {
				if err := tt.setupWS(tmpDir); err != nil {
					t.Fatalf("setupWS failed: %v", err)
				}
			}

			// Create grader
			cfg := &WorkspaceConfig{Checks: tt.checks}
			grader, err := NewWorkspaceGrader("test-workspace", cfg)
			if err != nil {
				t.Fatalf("NewWorkspaceGrader failed: %v", err)
			}

			// Grade
			input := GraderInput{
				WorkspacePath:  tmpDir,
				WorkspaceDelta: tt.delta,
				Config:         GraderConfig{Name: "test-workspace", Kind: "workspace"},
			}
			result, err := grader.Grade(context.Background(), input)
			if err != nil {
				t.Fatalf("Grade failed: %v", err)
			}

			// Validate result
			if result.Pass != tt.wantPass {
				t.Errorf("Pass = %v, want %v", result.Pass, tt.wantPass)
			}
			if result.Message != tt.wantMessage {
				t.Errorf("Message = %q, want %q", result.Message, tt.wantMessage)
			}
		})
	}
}

func TestWorkspaceGraderValidation(t *testing.T) {
	tests := []struct {
		name    string
		checks  []WorkspaceCheck
		wantErr bool
	}{
		{
			name: "valid require_to_create",
			checks: []WorkspaceCheck{
				{Kind: "require_to_create", Files: []string{"hello.md"}},
			},
			wantErr: false,
		},
		{
			name: "invalid require_to_create (empty files)",
			checks: []WorkspaceCheck{
				{Kind: "require_to_create", Files: []string{}},
			},
			wantErr: true,
		},
		{
			name: "valid file state=present",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present"},
			},
			wantErr: false,
		},
		{
			name: "invalid file (empty name)",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "", State: "present"},
			},
			wantErr: true,
		},
		{
			name: "invalid file (invalid state)",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "invalid"},
			},
			wantErr: true,
		},
		{
			name: "invalid file (modifiers on state=absent)",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "absent", MinBytes: int64Ptr(10)},
			},
			wantErr: true,
		},
		{
			name: "invalid file (min_bytes > max_bytes)",
			checks: []WorkspaceCheck{
				{Kind: "file", Name: "hello.md", State: "present", MinBytes: int64Ptr(100), MaxBytes: int64Ptr(10)},
			},
			wantErr: true,
		},
		{
			name: "unknown check kind",
			checks: []WorkspaceCheck{
				{Kind: "unknown", Files: []string{"hello.md"}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &WorkspaceConfig{Checks: tt.checks}
			_, err := NewWorkspaceGrader("test-workspace", cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWorkspaceGrader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func int64Ptr(v int64) *int64 {
	return &v
}
