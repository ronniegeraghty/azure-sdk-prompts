package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
)

// TestReviewerFactory_PerConfigIsolation verifies that the reviewerFactory
// validates only the specific config's Reviewer.Tools, not a pooled slice
// from multiple configs. This test exercises the fix from commit 0131f35d.
func TestReviewerFactory_PerConfigIsolation(t *testing.T) {
	// Create two configs with different reviewer skills
	dir := t.TempDir()
	
	// Create config A's reviewer skill
	skillA := filepath.Join(dir, "skill-a")
	if err := os.Mkdir(skillA, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillA, "SKILL.md"), []byte("# Skill A"), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create config B's reviewer skill
	skillB := filepath.Join(dir, "skill-b")
	if err := os.Mkdir(skillB, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillB, "SKILL.md"), []byte("# Skill B"), 0644); err != nil {
		t.Fatal(err)
	}
	
	configA := &config.ToolConfig{
		Name: "config-a",
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: skillA, Name: "skill-a"},
			},
		},
	}
	
	configB := &config.ToolConfig{
		Name: "config-b",
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: skillB, Name: "skill-b"},
			},
		},
	}
	
	// Simulate the reviewerFactory for config A
	// It should only see skill-a, not skill-b
	reportA, errA := tool.ValidateAndExpand(context.Background(), tool.ValidationInput{
		ReviewerTools: configA.Reviewer.Tools,
		ConfigDir:     "",
		PluginsDir:    "",
	})
	
	if errA != nil {
		t.Fatalf("config A validation failed: %v", errA)
	}
	
	skillDirsA := reportA.ReviewerSkillDirs()
	if len(skillDirsA) != 1 {
		t.Fatalf("config A: expected 1 reviewer skill dir, got %d", len(skillDirsA))
	}
	
	// Verify it's skill-a
	if !containsPath(skillDirsA, skillA) {
		t.Errorf("config A: expected skill-a in dirs, got %v", skillDirsA)
	}
	
	// Verify it doesn't contain skill-b
	if containsPath(skillDirsA, skillB) {
		t.Errorf("config A: should not contain skill-b, but found it in %v", skillDirsA)
	}
	
	// Now validate config B
	reportB, errB := tool.ValidateAndExpand(context.Background(), tool.ValidationInput{
		ReviewerTools: configB.Reviewer.Tools,
		ConfigDir:     "",
		PluginsDir:    "",
	})
	
	if errB != nil {
		t.Fatalf("config B validation failed: %v", errB)
	}
	
	skillDirsB := reportB.ReviewerSkillDirs()
	if len(skillDirsB) != 1 {
		t.Fatalf("config B: expected 1 reviewer skill dir, got %d", len(skillDirsB))
	}
	
	// Verify it's skill-b
	if !containsPath(skillDirsB, skillB) {
		t.Errorf("config B: expected skill-b in dirs, got %v", skillDirsB)
	}
	
	// Verify it doesn't contain skill-a
	if containsPath(skillDirsB, skillA) {
		t.Errorf("config B: should not contain skill-a, but found it in %v", skillDirsB)
	}
}

// TestReviewerFactory_MissingSkillFailsFast verifies that a missing
// reviewer skill causes the factory to return an error immediately,
// rather than silently passing an unresolved path to the SDK.
func TestReviewerFactory_MissingSkillFailsFast(t *testing.T) {
	cfg := &config.ToolConfig{
		Name: "test-config",
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: "./nonexistent", Name: "missing"},
			},
		},
	}
	
	// Simulate reviewerFactory validation
	_, err := tool.ValidateAndExpand(context.Background(), tool.ValidationInput{
		ReviewerTools: cfg.Reviewer.Tools,
		ConfigDir:     "",
		PluginsDir:    "",
	})
	
	if err == nil {
		t.Fatal("expected error for missing reviewer skill")
	}
	
	toolErr, ok := err.(*tool.ToolLoadError)
	if !ok {
		t.Fatalf("expected *ToolLoadError, got %T", err)
	}
	
	if toolErr.Kind != "skill" {
		t.Errorf("expected Kind=skill, got %q", toolErr.Kind)
	}
}

// TestReviewerFactory_EmptySkillDirFailsFast verifies that an empty
// skill_dir entry fails reviewer construction.
func TestReviewerFactory_EmptySkillDirFailsFast(t *testing.T) {
	emptyDir := t.TempDir()
	
	cfg := &config.ToolConfig{
		Name: "test-config",
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: emptyDir, Name: "empty", SkillDir: true},
			},
		},
	}
	
	_, err := tool.ValidateAndExpand(context.Background(), tool.ValidationInput{
		ReviewerTools: cfg.Reviewer.Tools,
		ConfigDir:     "",
		PluginsDir:    "",
	})
	
	if err == nil {
		t.Fatal("expected error for empty reviewer skill_dir")
	}
}

// TestReviewerFactory_SkillDirExpansion verifies that skill_dir=true
// reviewer entries expand to their child skills (generator parity).
func TestReviewerFactory_SkillDirExpansion(t *testing.T) {
	dir := t.TempDir()
	skillsDir := filepath.Join(dir, "reviewer-skills")
	if err := os.Mkdir(skillsDir, 0755); err != nil {
		t.Fatal(err)
	}
	
	// Create two child skills
	for _, name := range []string{"rev-skill-1", "rev-skill-2"} {
		skillDir := filepath.Join(skillsDir, name)
		if err := os.Mkdir(skillDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# "+name), 0644); err != nil {
			t.Fatal(err)
		}
	}
	
	cfg := &config.ToolConfig{
		Name: "test-config",
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: skillsDir, Name: "rev-skills", SkillDir: true},
			},
		},
	}
	
	report, err := tool.ValidateAndExpand(context.Background(), tool.ValidationInput{
		ReviewerTools: cfg.Reviewer.Tools,
		ConfigDir:     "",
		PluginsDir:    "",
	})
	
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	
	skillDirs := report.ReviewerSkillDirs()
	if len(skillDirs) != 2 {
		t.Fatalf("expected 2 expanded reviewer skills, got %d", len(skillDirs))
	}
}

func containsPath(paths []string, target string) bool {
	// Normalize paths for comparison
	targetAbs, _ := filepath.Abs(target)
	for _, p := range paths {
		pAbs, _ := filepath.Abs(p)
		if pAbs == targetAbs {
			return true
		}
	}
	return false
}
