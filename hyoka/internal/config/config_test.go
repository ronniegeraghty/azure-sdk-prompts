package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseValidConfig(t *testing.T) {
	data := []byte(`
configs:
  - name: test-config
    description: "Test configuration"
    generator:
      model: "gpt-4"
  - name: test-config-2
    description: "Second test"
    generator:
      model: "claude-sonnet-4.5"
      tools:
        - name: azure
          type: mcp
          command: npx
          args: ["-y", "@azure/mcp@latest"]
          mcp_tools: ["*"]
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Configs) != 2 {
		t.Fatalf("expected 2 configs, got %d", len(cfg.Configs))
	}
	if cfg.Configs[0].Name != "test-config" {
		t.Errorf("expected name 'test-config', got %q", cfg.Configs[0].Name)
	}
	if cfg.Configs[0].Generator.Model != "gpt-4" {
		t.Errorf("expected model 'gpt-4', got %q", cfg.Configs[0].Generator.Model)
	}
	// Check MCP server on second config
	if cfg.Configs[1].Generator == nil {
		t.Fatal("expected generator on second config")
	}
	var azure ToolEntry
	found := false
	for _, entry := range cfg.Configs[1].Generator.Tools {
		if entry.ResolvedType() == "mcp" && entry.Name == "azure" {
			azure = entry
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected 'azure' MCP server")
	}
	if azure.Command != "npx" {
		t.Errorf("expected command 'npx', got %q", azure.Command)
	}
}

func TestParseEmptyConfig(t *testing.T) {
	data := []byte(`configs: []`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for empty configs")
	}
}

func TestParseConfigMissingName(t *testing.T) {
	data := []byte(`
configs:
  - description: "No name"
    generator:
      model: "gpt-4"
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for config missing name")
	}
}

func TestParseInvalidYAML(t *testing.T) {
	data := []byte(`not: valid: yaml: [`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestValidateSameModelAccepted(t *testing.T) {
	data := []byte(`
configs:
  - name: same-model-ok
    description: "Same model for generator and reviewer is allowed"
    generator:
      model: "claude-opus-4.6"
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gpt-4.1"
`)
	_, err := Parse(data)
	if err != nil {
		t.Fatalf("expected no error when reviewer model matches generator, got: %v", err)
	}
}

func TestValidateDifferentModelsAccepted(t *testing.T) {
	data := []byte(`
configs:
  - name: good-config
    description: "Different models"
    generator:
      model: "claude-sonnet-4.5"
    reviewer:
      models:
        - "gpt-4.1"
        - "gemini-3-pro-preview"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	models := cfg.Configs[0].Reviewer.Models
	if len(models) != 2 {
		t.Errorf("expected 2 reviewer models, got %d", len(models))
	}
}

func TestValidateNoReviewerModelAccepted(t *testing.T) {
	data := []byte(`
configs:
  - name: no-reviewer
    description: "No reviewer model specified"
    generator:
      model: "claude-sonnet-4.5"
`)
	_, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateDuplicateReviewerModelsRejected(t *testing.T) {
	data := []byte(`
configs:
  - name: dupes
    description: "Duplicate reviewer models"
    generator:
      model: "claude-sonnet-4.5"
    reviewer:
      models:
        - "gpt-4.1"
        - "gpt-4.1"
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for duplicate reviewer models")
	}
}

func TestReviewerSingleModel(t *testing.T) {
	data := []byte(`
configs:
  - name: single-reviewer
    description: "Single reviewer model"
    generator:
      model: "claude-sonnet-4.5"
    reviewer:
      models:
        - "gpt-4.1"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	models := cfg.Configs[0].Reviewer.Models
	if len(models) != 1 || models[0] != "gpt-4.1" {
		t.Errorf("expected [gpt-4.1], got %v", models)
	}
}

func TestGetConfig(t *testing.T) {
	data := []byte(`
configs:
  - name: alpha
    description: "Alpha"
    generator:
      model: "gpt-4"
  - name: beta
    description: "Beta"
    generator:
      model: "claude-sonnet-4.5"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tc, err := cfg.GetConfig("beta")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tc.Name != "beta" {
		t.Errorf("expected 'beta', got %q", tc.Name)
	}

	_, err = cfg.GetConfig("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent config")
	}
}

func TestGetConfigs(t *testing.T) {
	data := []byte(`
configs:
  - name: alpha
    description: "Alpha"
    generator:
      model: "gpt-4"
  - name: beta
    description: "Beta"
    generator:
      model: "claude-sonnet-4.5"
  - name: gamma
    description: "Gamma"
    generator:
      model: "gpt-4"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Empty names returns all
	all, err := cfg.GetConfigs(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(all) != 3 {
		t.Errorf("expected 3 configs, got %d", len(all))
	}

	// Specific names
	subset, err := cfg.GetConfigs([]string{"alpha", "gamma"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(subset) != 2 {
		t.Errorf("expected 2 configs, got %d", len(subset))
	}

	// Missing name
	_, err = cfg.GetConfigs([]string{"alpha", "missing"})
	if err == nil {
		t.Fatal("expected error for missing config name")
	}
}

func TestLoadFromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	content := []byte(`
configs:
  - name: file-test
    description: "From file"
    generator:
      model: "gpt-4"
`)
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Configs[0].Name != "file-test" {
		t.Errorf("expected 'file-test', got %q", cfg.Configs[0].Name)
	}

	// Non-existent file
	_, err = Load(filepath.Join(dir, "nonexistent.yaml"))
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

func TestParseGeneratorSkillsAndPlugins(t *testing.T) {
	data := []byte(`
configs:
  - name: with-skills
    description: "Config with skills and a plugin tool entry"
    generator:
      model: "gpt-4"
      tools:
        - name: local-skill
          type: skill
          source: local
          path: "./skills/tool-use"
        - name: org-skill
          type: skill
          source: remote
          repo: "github:org/repo"
        - name: azure-functions
          type: plugin
          source: remote
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]
	if len(c.Generator.Tools) != 3 {
		t.Fatalf("expected 3 tools, got %d", len(c.Generator.Tools))
	}
	if c.Generator.Tools[0].ResolvedType() != "skill" || c.Generator.Tools[0].Path != "./skills/tool-use" {
		t.Errorf("expected local skill './skills/tool-use', got %+v", c.Generator.Tools[0])
	}
	if c.Generator.Tools[1].ResolvedType() != "skill" || c.Generator.Tools[1].Repo != "github:org/repo" {
		t.Errorf("expected remote skill 'github:org/repo', got %+v", c.Generator.Tools[1])
	}
	if c.Generator.Tools[2].ResolvedType() != "plugin" || c.Generator.Tools[2].Name != "azure-functions" {
		t.Errorf("expected plugin tool entry 'azure-functions', got %+v", c.Generator.Tools[2])
	}
}

// TestParse_RejectsRetiredTopLevelPluginsField asserts that the retired
// top-level `plugins:` field produces a migration-hint error rather than
// being silently ignored or parsed under the old model.
func TestParse_RejectsRetiredTopLevelPluginsField(t *testing.T) {
	data := []byte(`
configs:
  - name: legacy
    description: "uses retired schema"
    generator:
      model: "gpt-4"
    plugins:
      - "@azure/functions"
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected migration error for top-level plugins: field")
	}
	if !strings.Contains(err.Error(), "retired") || !strings.Contains(err.Error(), "type: plugin") {
		t.Errorf("error should mention retirement and migration shape; got: %v", err)
	}
}

func TestParseNoSkillsOrPlugins(t *testing.T) {
	data := []byte(`
configs:
  - name: no-extras
    description: "Config without skills or plugins"
    generator:
      model: "gpt-4"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]
	if len(c.Generator.Tools) != 0 {
		t.Errorf("expected 0 tools, got %d", len(c.Generator.Tools))
	}
}

func TestInstallSkillsAndPluginsEmpty(t *testing.T) {
	configs := []ToolConfig{
		{Name: "empty", Description: "No skills", Generator: &GeneratorConfig{Model: "gpt-4"}},
	}
	if err := InstallSkillsAndPlugins(configs); err != nil {
		t.Fatalf("expected no error for empty skills/plugins, got: %v", err)
	}
}

func TestParseNewFormatGeneratorReviewer(t *testing.T) {
	data := []byte(`
configs:
  - name: new-format
    description: "New format with generator/reviewer"
    generator:
      model: "claude-sonnet-4.5"
      tools:
        - name: generator-skills
          type: skill
          source: local
          path: "./skills/generator"
        - name: azure
          type: mcp
          command: npx
          args: ["-y", "@azure/mcp@latest"]
          mcp_tools: ["*"]
        - name: create
        - name: edit
      excluded_tools: ["web_fetch"]
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gemini-3-pro-preview"
      tools:
        - name: reviewer-skills
          type: skill
          source: local
          path: "./skills/reviewer"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]

	if c.Generator.Model != "claude-sonnet-4.5" {
		t.Errorf("expected model 'claude-sonnet-4.5', got %q", c.Generator.Model)
	}
	models := c.Reviewer.Models
	if len(models) != 2 || models[0] != "claude-opus-4.6" {
		t.Errorf("expected reviewer models [claude-opus-4.6 gemini-3-pro-preview], got %v", models)
	}
	if len(c.Generator.Tools) != 4 {
		t.Errorf("expected 4 generator tools, got %d", len(c.Generator.Tools))
	}
	var hasSkill, hasMCP bool
	for _, entry := range c.Generator.Tools {
		if entry.ResolvedType() == "skill" && entry.Path == "./skills/generator" {
			hasSkill = true
		}
		if entry.ResolvedType() == "mcp" && entry.Name == "azure" {
			hasMCP = true
		}
	}
	if !hasSkill {
		t.Error("expected generator skill entry with ./skills/generator")
	}
	if !hasMCP {
		t.Error("expected generator MCP entry named azure")
	}
	if len(c.Reviewer.Tools) != 1 {
		t.Errorf("expected 1 reviewer tool, got %d", len(c.Reviewer.Tools))
	}
	if len(c.Generator.ExcludedTools) != 1 {
		t.Errorf("expected 1 excluded tool, got %d", len(c.Generator.ExcludedTools))
	}
}

func TestGeneratorReviewerFieldsPopulated(t *testing.T) {
	data := []byte(`
configs:
  - name: full-config
    description: "Full config with generator and reviewer"
    generator:
      model: "claude-opus-4.6"
      tools:
        - name: azure
          type: mcp
          command: npx
          args: ["-y", "@azure/mcp@latest"]
        - name: generator-skills
          type: skill
          source: local
          path: "./skills/generator"
        - name: create
      excluded_tools: ["bash"]
    reviewer:
      models:
        - "gpt-4.1"
      tools:
        - name: reviewer-skills
          type: skill
          source: local
          path: "./skills/reviewer"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]

	if c.Generator == nil {
		t.Fatal("Generator should not be nil")
	}
	if c.Reviewer == nil {
		t.Fatal("Reviewer should not be nil")
	}
	if c.Generator.Model != "claude-opus-4.6" {
		t.Errorf("expected Generator.Model 'claude-opus-4.6', got %q", c.Generator.Model)
	}
	if len(c.Reviewer.Models) != 1 || c.Reviewer.Models[0] != "gpt-4.1" {
		t.Errorf("expected Reviewer.Models [gpt-4.1], got %v", c.Reviewer.Models)
	}
	if len(c.Generator.Tools) != 3 {
		t.Errorf("expected 3 generator tools, got %d", len(c.Generator.Tools))
	}
	if len(c.Reviewer.Tools) != 1 {
		t.Errorf("expected 1 reviewer tool, got %d", len(c.Reviewer.Tools))
	}
}

func TestParseRemoteSkill(t *testing.T) {
	data := []byte(`
configs:
  - name: with-remote
    description: "Config with remote skill"
    generator:
      model: "gpt-4"
      tools:
        - name: azure-keyvault-py
          type: skill
          source: remote
          repo: microsoft/skills
        - name: local-skill
          type: skill
          source: local
          path: "./skills/local"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]
	tools := c.Generator.Tools
	if len(tools) != 2 {
		t.Fatalf("expected 2 tools, got %d", len(tools))
	}
	if tools[0].ResolvedType() != "skill" || tools[0].Name != "azure-keyvault-py" || tools[0].Repo != "microsoft/skills" {
		t.Errorf("unexpected remote skill: %+v", tools[0])
	}
	if tools[1].ResolvedType() != "skill" || tools[1].Path != "./skills/local" {
		t.Errorf("unexpected local skill: %+v", tools[1])
	}
}

func TestValidateRejectsInvalidToolType(t *testing.T) {
	data := []byte(`
configs:
  - name: bad-skill
    description: "Bad skill type"
    generator:
      model: "gpt-4"
      tools:
        - name: bad-tool
          type: invalid
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for invalid tool type")
	}
}

func TestValidateRejectsLocalSkillMissingPath(t *testing.T) {
	data := []byte(`
configs:
  - name: no-path
    description: "Local skill missing path"
    generator:
      model: "gpt-4"
      tools:
        - name: missing-path
          type: skill
          source: local
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for local skill without path")
	}
}

func TestValidateRejectsRemoteSkillMissingRepo(t *testing.T) {
	data := []byte(`
configs:
  - name: no-repo
    description: "Remote skill missing repo"
    generator:
      model: "gpt-4"
      tools:
        - name: some-skill
          type: skill
          source: remote
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for remote skill without repo")
	}
}

func TestParseDuplicateConfigNamesRejected(t *testing.T) {
	data := []byte(`
configs:
  - name: same-name
    description: "First"
    generator:
      model: "gpt-4"
  - name: same-name
    description: "Second"
    generator:
      model: "claude-sonnet-4.5"
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for duplicate config names within a file")
	}
	if got := err.Error(); !strings.Contains(got, "duplicate config name") {
		t.Errorf("expected duplicate config name error, got: %v", err)
	}
}

func TestLoadDirDuplicateConfigNamesAcrossFiles(t *testing.T) {
	dir := t.TempDir()
	file1 := []byte(`
configs:
  - name: shared-name
    description: "In file1"
    generator:
      model: "gpt-4"
`)
	file2 := []byte(`
configs:
  - name: shared-name
    description: "In file2"
    generator:
      model: "claude-sonnet-4.5"
`)
	if err := os.WriteFile(filepath.Join(dir, "a.yaml"), file1, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "b.yaml"), file2, 0644); err != nil {
		t.Fatal(err)
	}
	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("expected error for duplicate config names across files")
	}
	got := err.Error()
	if !strings.Contains(got, "duplicate config name") || !strings.Contains(got, "a.yaml") || !strings.Contains(got, "b.yaml") {
		t.Errorf("expected error mentioning duplicate name and both files, got: %v", err)
	}
}

func TestGeneratorModelDirectAccess(t *testing.T) {
	c := ToolConfig{
		Name: "test",
		Generator: &GeneratorConfig{
			Model: "new-model",
		},
	}
	if c.Generator.Model != "new-model" {
		t.Errorf("expected 'new-model', got %q", c.Generator.Model)
	}
}

func TestResolveModelsSingle(t *testing.T) {
	g := &GeneratorConfig{Model: "gpt-4"}
	if got := g.ResolveModels(); !reflect.DeepEqual(got, []string{"gpt-4"}) {
		t.Errorf("expected [gpt-4], got %v", got)
	}
}

func TestResolveModelsMultiple(t *testing.T) {
	g := &GeneratorConfig{Models: []string{"claude-opus-4.6", "claude-sonnet-4.5"}}
	want := []string{"claude-opus-4.6", "claude-sonnet-4.5"}
	if got := g.ResolveModels(); !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestResolveModelsBoth(t *testing.T) {
	g := &GeneratorConfig{Model: "ignored", Models: []string{"m1", "m2"}}
	want := []string{"m1", "m2"}
	if got := g.ResolveModels(); !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestResolveModelsNone(t *testing.T) {
	g := &GeneratorConfig{}
	if got := g.ResolveModels(); len(got) != 0 {
		t.Errorf("expected empty models, got %v", got)
	}
	cf := &ConfigFile{
		Configs: []ToolConfig{{Name: "missing", Generator: g}},
	}
	if err := cf.Validate(); err == nil {
		t.Fatal("expected error for missing generator model")
	}
}

func TestValidateRejectsNilGenerator(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "no-gen"},
		},
	}
	err := cf.Validate()
	if err == nil {
		t.Fatal("expected error for nil generator")
	}
	want := `config "no-gen": generator.model or generator.models is required`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestParseSystemPrompt(t *testing.T) {
	data := []byte(`
configs:
  - name: with-system-prompt
    description: "Config with system prompts"
    generator:
      model: "claude-opus-4.6"
      system_prompt: "You are an Azure SDK expert."
    reviewer:
      models:
        - "gpt-4.1"
      system_prompt: "Review code for Azure best practices."
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]
	if c.Generator.SystemPrompt != "You are an Azure SDK expert." {
		t.Errorf("expected generator system_prompt 'You are an Azure SDK expert.', got %q", c.Generator.SystemPrompt)
	}
	if c.Reviewer.SystemPrompt != "Review code for Azure best practices." {
		t.Errorf("expected reviewer system_prompt 'Review code for Azure best practices.', got %q", c.Reviewer.SystemPrompt)
	}
}

func TestParseSystemPromptOmitted(t *testing.T) {
	data := []byte(`
configs:
  - name: no-system-prompt
    description: "Config without system prompts"
    generator:
      model: "gpt-4"
    reviewer:
      models:
        - "gpt-4.1"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]
	if c.Generator.SystemPrompt != "" {
		t.Errorf("expected empty generator system_prompt, got %q", c.Generator.SystemPrompt)
	}
	if c.Reviewer.SystemPrompt != "" {
		t.Errorf("expected empty reviewer system_prompt, got %q", c.Reviewer.SystemPrompt)
	}
}

func TestValidateRejectsEmptyGeneratorModel(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "empty-model", Generator: &GeneratorConfig{Model: ""}},
		},
	}
	err := cf.Validate()
	if err == nil {
		t.Fatal("expected error for empty generator model")
	}
	want := `config "empty-model": generator.model or generator.models is required`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestParseSessionLimits(t *testing.T) {
	data := []byte(`
configs:
  - name: with-limits
    generator:
      model: "gpt-4"
    limits:
      max_turns: 10
      max_files: 20
      max_session_actions: 30
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lim := cfg.Configs[0].Limits
	if lim == nil {
		t.Fatal("expected Limits to be set")
	}
	if lim.MaxTurns != 10 {
		t.Errorf("MaxTurns: expected 10, got %d", lim.MaxTurns)
	}
	if lim.MaxFiles != 20 {
		t.Errorf("MaxFiles: expected 20, got %d", lim.MaxFiles)
	}
	if lim.MaxSessionActions != 30 {
		t.Errorf("MaxSessionActions: expected 30, got %d", lim.MaxSessionActions)
	}
}

func TestParseSessionLimitsOmitted(t *testing.T) {
	data := []byte(`
configs:
  - name: no-limits
    generator:
      model: "gpt-4"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Configs[0].Limits != nil {
		t.Error("expected Limits to be nil when omitted")
	}
}

func TestParseSessionLimitsPartial(t *testing.T) {
	data := []byte(`
configs:
  - name: partial-limits
    generator:
      model: "gpt-4"
    limits:
      max_turns: 15
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lim := cfg.Configs[0].Limits
	if lim == nil {
		t.Fatal("expected Limits to be set")
	}
	if lim.MaxTurns != 15 {
		t.Errorf("MaxTurns: expected 15, got %d", lim.MaxTurns)
	}
	if lim.MaxFiles != 0 {
		t.Errorf("MaxFiles: expected 0, got %d", lim.MaxFiles)
	}
	if lim.MaxSessionActions != 0 {
		t.Errorf("MaxSessionActions: expected 0, got %d", lim.MaxSessionActions)
	}
}

func TestValidateRejectsNegativeMaxTurns(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "neg", Generator: &GeneratorConfig{Model: "gpt-4"}, Limits: &SessionLimits{MaxTurns: -1}},
		},
	}
	err := cf.Validate()
	if err == nil {
		t.Fatal("expected error for negative max_turns")
	}
	want := `config "neg": limits.max_turns must not be negative`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestValidateRejectsNegativeMaxFiles(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "neg", Generator: &GeneratorConfig{Model: "gpt-4"}, Limits: &SessionLimits{MaxFiles: -5}},
		},
	}
	err := cf.Validate()
	if err == nil {
		t.Fatal("expected error for negative max_files")
	}
	want := `config "neg": limits.max_files must not be negative`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestValidateRejectsNegativeMaxSessionActions(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "neg", Generator: &GeneratorConfig{Model: "gpt-4"}, Limits: &SessionLimits{MaxSessionActions: -10}},
		},
	}
	err := cf.Validate()
	if err == nil {
		t.Fatal("expected error for negative max_session_actions")
	}
	want := `config "neg": limits.max_session_actions must not be negative`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestValidateAcceptsZeroLimits(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{
			{Name: "zero", Generator: &GeneratorConfig{Model: "gpt-4"}, Limits: &SessionLimits{MaxTurns: 0, MaxFiles: 0}},
		},
	}
	if err := cf.Validate(); err != nil {
		t.Fatalf("zero limits should be accepted, got: %v", err)
	}
}

func TestParseNilGeneratorDoesNotPanic(t *testing.T) {
	data := []byte(`
configs:
  - name: nil-gen
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for nil generator")
	}
	want := `config "nil-gen": generator.model or generator.models is required`
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}
