package skills

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

func TestResolveLocal_SingleSkill(t *testing.T) {
	dir := t.TempDir()
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
		t.Fatal(err)
	}

	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: skillDir, Name: "local-skill"},
	}, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 1 {
		t.Fatalf("expected 1 dir, got %d", len(dirs))
	}
	if dirs[0] != skillDir {
		t.Errorf("expected %q, got %q", skillDir, dirs[0])
	}
}

func TestResolveLocal_SingleSkillMissingSKILLMD(t *testing.T) {
	dir := t.TempDir()
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}
	// No SKILL.md — should warn and return nothing

	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: skillDir, Name: "local-skill"},
	}, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 0 {
		t.Errorf("expected 0 dirs for single skill missing SKILL.md, got %d", len(dirs))
	}
}

func TestResolveLocal_SkillDir(t *testing.T) {
	dir := t.TempDir()
	skillsDir := filepath.Join(dir, "skills", "generator")
	if err := os.MkdirAll(skillsDir, 0755); err != nil {
		t.Fatal(err)
	}
	// Create two skill subdirectories and one non-skill
	for _, name := range []string{"skill-a", "skill-b", "not-a-skill"} {
		if err := os.Mkdir(filepath.Join(skillsDir, name), 0755); err != nil {
			t.Fatal(err)
		}
	}
	for _, name := range []string{"skill-a", "skill-b"} {
		if err := os.WriteFile(filepath.Join(skillsDir, name, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: "skills/generator", Name: "gen-skills", SkillDir: true},
	}, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 2 {
		t.Fatalf("expected 2 skill subdirs, got %d: %v", len(dirs), dirs)
	}
}

func TestResolveLocal_SkillDirEmpty(t *testing.T) {
	dir := t.TempDir()
	emptyDir := filepath.Join(dir, "skills", "generator")
	if err := os.MkdirAll(emptyDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(emptyDir, ".gitkeep"), nil, 0644); err != nil {
		t.Fatal(err)
	}

	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: emptyDir, Name: "empty-skills", SkillDir: true},
	}, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 0 {
		t.Errorf("expected 0 dirs for empty skill_dir, got %d: %v", len(dirs), dirs)
	}
}

func TestResolveLocal_GlobPattern(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"skill-a", "skill-b", "not-a-skill.txt"} {
		p := filepath.Join(dir, "skills", name)
		if err := os.MkdirAll(p, 0755); err != nil {
			t.Fatal(err)
		}
	}
	// Create a file (not a dir) that matches the glob
	f := filepath.Join(dir, "skills", "not-a-skill.txt", "dummy")
	os.Remove(filepath.Join(dir, "skills", "not-a-skill.txt"))
	os.WriteFile(filepath.Join(dir, "skills", "readme.txt"), []byte("hi"), 0644)
	_ = f

	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: filepath.Join(dir, "skills", "*"), Name: "local-skill"},
	}, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should get skill-a and skill-b (directories only)
	if len(dirs) != 2 {
		t.Fatalf("expected 2 dirs from glob, got %d: %v", len(dirs), dirs)
	}
}

func TestResolveLocal_EmptyEntries(t *testing.T) {
	dirs, err := ResolveSkillDirs(nil, ".")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 0 {
		t.Errorf("expected 0 dirs, got %d", len(dirs))
	}
}

func TestResolveLocal_NonExistentDir(t *testing.T) {
	dirs, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "local", Path: "/does/not/exist", Name: "missing"},
	}, ".")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dirs) != 0 {
		t.Errorf("expected 0 dirs for non-existent path, got %d: %v", len(dirs), dirs)
	}
}

func TestCountSkills(t *testing.T) {
	dir := t.TempDir()
	// After skill_dir expansion, each dir should directly contain SKILL.md
	skillA := filepath.Join(dir, "skill-a")
	skillB := filepath.Join(dir, "skill-b")
	noSkill := filepath.Join(dir, "not-a-skill")
	for _, d := range []string{skillA, skillB, noSkill} {
		if err := os.Mkdir(d, 0755); err != nil {
			t.Fatal(err)
		}
	}
	for _, d := range []string{skillA, skillB} {
		if err := os.WriteFile(filepath.Join(d, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	count := CountSkills([]string{skillA, skillB, noSkill})
	if count != 2 {
		t.Errorf("expected 2 skills, got %d", count)
	}
}

func TestCountSkills_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	count := CountSkills([]string{dir})
	if count != 0 {
		t.Errorf("expected 0 skills for empty dir, got %d", count)
	}
}

func TestCountSkills_NilDirs(t *testing.T) {
	count := CountSkills(nil)
	if count != 0 {
		t.Errorf("expected 0 skills for nil input, got %d", count)
	}
}

func TestResolveLocal_InvalidType(t *testing.T) {
	_, err := ResolveSkillDirs([]config.ToolEntry{
		{Type: "skill", Source: "unknown", Path: "/some/path", Name: "bad-skill"},
	}, ".")
	if err == nil {
		t.Fatal("expected error for unknown type")
	}
}
