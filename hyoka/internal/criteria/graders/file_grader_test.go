package graders

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func boolPtr(b bool) *bool { return &b }

func TestNewFileGrader_Defaults(t *testing.T) {
	g, err := NewFileGrader("test", &FileConfig{Path: "main.py"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindFile {
		t.Errorf("Kind() = %q, want %q", g.Kind(), KindFile)
	}
	if g.Name() != "test" {
		t.Errorf("Name() = %q, want %q", g.Name(), "test")
	}
	if !g.mustExist {
		t.Error("mustExist should default to true")
	}
}

func TestNewFileGrader_MustExistExplicit(t *testing.T) {
	g, err := NewFileGrader("opt", &FileConfig{Path: "x.py", MustExist: boolPtr(false)})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.mustExist {
		t.Error("mustExist should be false when explicitly set")
	}
}

func TestNewFileGrader_MissingPath(t *testing.T) {
	_, err := NewFileGrader("bad", &FileConfig{})
	if err == nil {
		t.Fatal("expected error for missing path")
	}
}

func TestNewFileGrader_InvalidRegex(t *testing.T) {
	_, err := NewFileGrader("bad", &FileConfig{Path: "x.py", Pattern: "[invalid"})
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestFileGrader_FileExists_NoPattern(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "main.py"), []byte("print('hello')"), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("exists_check", &FileConfig{Path: "main.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}

	if !result.Pass {
		t.Error("expected Pass=true")
	}
	if result.Score != 1.0 {
		t.Errorf("Score = %f, want 1.0", result.Score)
	}
	if result.Extras == nil || result.Extras.File == nil || len(result.Extras.File.Files) != 1 {
		t.Fatal("expected exactly 1 checked file in Extras.File.Files")
	}
	if !result.Extras.File.Files[0].Exists {
		t.Error("file should be marked as existing")
	}
}

func TestFileGrader_FileMissing_MustExist(t *testing.T) {
	dir := t.TempDir()

	g, _ := NewFileGrader("missing", &FileConfig{Path: "nope.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}

	if result.Pass {
		t.Error("expected Pass=false for missing required file")
	}
	if result.Score != 0 {
		t.Errorf("Score = %f, want 0", result.Score)
	}
	if result.Extras.File.Files[0].Exists {
		t.Error("file should be marked as not existing")
	}
}

func TestFileGrader_FileMissing_MustExistFalse(t *testing.T) {
	dir := t.TempDir()

	g, _ := NewFileGrader("optional", &FileConfig{Path: "opt.py", MustExist: boolPtr(false)})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 0.5},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}

	if !result.Pass {
		t.Error("expected Pass=true for optional missing file")
	}
	if result.Score != 1.0 {
		t.Errorf("Score = %f, want 1.0", result.Score)
	}
}

func TestFileGrader_PatternMatches(t *testing.T) {
	dir := t.TempDir()
	content := "from azure.identity import DefaultAzureCredential\n"
	if err := os.WriteFile(filepath.Join(dir, "main.py"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("import_check", &FileConfig{
		Path:    "main.py",
		Pattern: `import\s+DefaultAzureCredential`,
	})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}

	if !result.Pass {
		t.Error("expected Pass=true when pattern matches")
	}
	if result.Score != 1.0 {
		t.Errorf("Score = %f, want 1.0", result.Score)
	}
	cf := result.Extras.File.Files[0]
	if !cf.PatternMatched {
		t.Error("expected PatternMatched=true")
	}
}

func TestFileGrader_PatternDoesNotMatch(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "main.py"), []byte("print('hello')"), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("no_match", &FileConfig{
		Path:    "main.py",
		Pattern: `import\s+DefaultAzureCredential`,
	})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}

	if result.Pass {
		t.Error("expected Pass=false when pattern doesn't match")
	}
	if result.Score != 0.5 {
		t.Errorf("Score = %f, want 0.5", result.Score)
	}
	cf := result.Extras.File.Files[0]
	if cf.PatternMatched {
		t.Error("expected PatternMatched=false")
	}
}

func TestFileGrader_SubdirectoryPath(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "src", "auth")
	if err := os.MkdirAll(sub, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "client.py"), []byte("class Client: pass"), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("subdir", &FileConfig{Path: "src/auth/client.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}
	if !result.Pass || result.Score != 1.0 {
		t.Errorf("expected pass=true, score=1.0; got pass=%v, score=%f", result.Pass, result.Score)
	}
}

func TestFileGrader_GateFlag(t *testing.T) {
	dir := t.TempDir()

	g, _ := NewFileGrader("gate_check", &FileConfig{Path: "required.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0, Gate: true},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}
	if !result.Gate {
		t.Error("expected Gate=true to propagate from config")
	}
	if result.Pass {
		t.Error("expected Pass=false for missing gated file")
	}
}

func TestFileGrader_WeightFromConfig(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "x.py"), []byte("ok"), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("w", &FileConfig{Path: "x.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 0.3},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}
	if result.Weight != 0.3 {
		t.Errorf("Weight = %f, want 0.3", result.Weight)
	}
}

func TestFileGrader_DefaultWeight(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "x.py"), []byte("ok"), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("dw", &FileConfig{Path: "x.py"})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{}, // Weight=0 → EffectiveWeight()=1.0
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}
	if result.Weight != 1.0 {
		t.Errorf("Weight = %f, want 1.0 (default)", result.Weight)
	}
}

func TestFileGrader_ImplementsInterface(t *testing.T) {
	g, _ := NewFileGrader("iface", &FileConfig{Path: "x.py"})
	var _ Grader = g // compile-time interface check
}

func TestFileGrader_MultilinePattern(t *testing.T) {
	dir := t.TempDir()
	content := `import os
import sys

def main():
    print("hello")
`
	if err := os.WriteFile(filepath.Join(dir, "app.py"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	g, _ := NewFileGrader("multi", &FileConfig{
		Path:    "app.py",
		Pattern: `(?s)import os.*def main`,
	})
	result, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: dir,
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade error: %v", err)
	}
	if !result.Pass {
		t.Errorf("expected multiline pattern to match, msg=%s", result.Message)
	}
}
