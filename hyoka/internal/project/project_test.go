package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	content := `
prompts_dir: my-prompts
configs_dir: my-configs
skills_dir: my-skills
criteria_dir: my-criteria
reports_dir: my-reports
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cfgPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Paths should be resolved relative to the config file dir
	wantPrompts := filepath.Join(dir, "my-prompts")
	if cfg.PromptsDir != wantPrompts {
		t.Errorf("PromptsDir = %q, want %q", cfg.PromptsDir, wantPrompts)
	}
	wantConfigs := filepath.Join(dir, "my-configs")
	if cfg.ConfigsDir != wantConfigs {
		t.Errorf("ConfigsDir = %q, want %q", cfg.ConfigsDir, wantConfigs)
	}
	wantSkills := filepath.Join(dir, "my-skills")
	if cfg.SkillsDir != wantSkills {
		t.Errorf("SkillsDir = %q, want %q", cfg.SkillsDir, wantSkills)
	}
	wantCriteria := filepath.Join(dir, "my-criteria")
	if cfg.CriteriaDir != wantCriteria {
		t.Errorf("CriteriaDir = %q, want %q", cfg.CriteriaDir, wantCriteria)
	}
	wantReports := filepath.Join(dir, "my-reports")
	if cfg.ReportsDir != wantReports {
		t.Errorf("ReportsDir = %q, want %q", cfg.ReportsDir, wantReports)
	}
	if cfg.ConfigPath != cfgPath {
		t.Errorf("ConfigPath = %q, want %q", cfg.ConfigPath, cfgPath)
	}
}

func TestLoad_AbsolutePaths(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	content := `
prompts_dir: /absolute/prompts
configs_dir: /absolute/configs
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cfgPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.PromptsDir != "/absolute/prompts" {
		t.Errorf("PromptsDir = %q, want /absolute/prompts", cfg.PromptsDir)
	}
	if cfg.ConfigsDir != "/absolute/configs" {
		t.Errorf("ConfigsDir = %q, want /absolute/configs", cfg.ConfigsDir)
	}
}

func TestLoad_PartialConfig(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	content := `prompts_dir: my-prompts`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cfgPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.PromptsDir != filepath.Join(dir, "my-prompts") {
		t.Errorf("PromptsDir = %q, want resolved path", cfg.PromptsDir)
	}
	// Unset fields should remain empty
	if cfg.ConfigsDir != "" {
		t.Errorf("ConfigsDir = %q, want empty", cfg.ConfigsDir)
	}
	if cfg.ReportsDir != "" {
		t.Errorf("ReportsDir = %q, want empty", cfg.ReportsDir)
	}
}

func TestLoad_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	if err := os.WriteFile(cfgPath, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(cfgPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// All fields should be empty
	if cfg.PromptsDir != "" || cfg.ConfigsDir != "" || cfg.ReportsDir != "" {
		t.Error("expected all dirs to be empty for empty config")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	if err := os.WriteFile(cfgPath, []byte("{{not yaml"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(cfgPath)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/hyoka.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestDiscover_FindsHyokaYAML(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "hyoka.yaml")

	content := `prompts_dir: prompts`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Discover(dir)
	if err != nil {
		t.Fatalf("Discover() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("Discover() returned nil, expected config")
	}
	if cfg.PromptsDir != filepath.Join(dir, "prompts") {
		t.Errorf("PromptsDir = %q, want %q", cfg.PromptsDir, filepath.Join(dir, "prompts"))
	}
}

func TestDiscover_FindsDotHyokaYAML(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, ".hyoka.yaml")

	content := `prompts_dir: my-prompts`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Discover(dir)
	if err != nil {
		t.Fatalf("Discover() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("Discover() returned nil, expected config")
	}
	if cfg.PromptsDir != filepath.Join(dir, "my-prompts") {
		t.Errorf("PromptsDir = %q", cfg.PromptsDir)
	}
}

func TestDiscover_PrefersHyokaYAMLOverDot(t *testing.T) {
	dir := t.TempDir()

	// Both exist — hyoka.yaml should win
	if err := os.WriteFile(filepath.Join(dir, "hyoka.yaml"), []byte("prompts_dir: from-hyoka"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".hyoka.yaml"), []byte("prompts_dir: from-dot"), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Discover(dir)
	if err != nil {
		t.Fatalf("Discover() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected config")
	}
	want := filepath.Join(dir, "from-hyoka")
	if cfg.PromptsDir != want {
		t.Errorf("PromptsDir = %q, want %q (hyoka.yaml should take priority)", cfg.PromptsDir, want)
	}
}

func TestDiscover_WalksUpParent(t *testing.T) {
	root := t.TempDir()
	child := filepath.Join(root, "sub", "deep")
	if err := os.MkdirAll(child, 0755); err != nil {
		t.Fatal(err)
	}

	cfgPath := filepath.Join(root, "hyoka.yaml")
	if err := os.WriteFile(cfgPath, []byte("prompts_dir: prompts"), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Discover(child)
	if err != nil {
		t.Fatalf("Discover() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected config found by walking up")
	}
	want := filepath.Join(root, "prompts")
	if cfg.PromptsDir != want {
		t.Errorf("PromptsDir = %q, want %q", cfg.PromptsDir, want)
	}
}

func TestDiscover_NoConfigReturnsNil(t *testing.T) {
	dir := t.TempDir()
	// No config file anywhere

	cfg, err := Discover(dir)
	if err != nil {
		t.Fatalf("Discover() error: %v", err)
	}
	if cfg != nil {
		t.Error("expected nil config when no file found")
	}
}

func TestEffective_NilConfig(t *testing.T) {
	var cfg *Config

	if cfg.EffectivePromptsDir() != "" {
		t.Error("expected empty for nil config")
	}
	if cfg.EffectiveConfigsDir() != "" {
		t.Error("expected empty for nil config")
	}
	if cfg.EffectiveSkillsDir() != "" {
		t.Error("expected empty for nil config")
	}
	if cfg.EffectiveCriteriaDir() != "" {
		t.Error("expected empty for nil config")
	}
	if cfg.EffectiveReportsDir() != "" {
		t.Error("expected empty for nil config")
	}
}
