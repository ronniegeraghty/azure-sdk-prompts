// Package config provides configuration loading and parsing for the evaluation tool.
package config

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// GeneratorConfig holds all configuration for the generator agent.
type GeneratorConfig struct {
	Model         string      `yaml:"model,omitempty" json:"model,omitempty"`
	Models        []string    `yaml:"models,omitempty" json:"models,omitempty"`
	SystemPrompt  string      `yaml:"system_prompt,omitempty" json:"system_prompt,omitempty"`
	Tools         []ToolEntry `yaml:"tools,omitempty" json:"tools,omitempty"`
	ExcludedTools []string    `yaml:"excluded_tools,omitempty" json:"excluded_tools,omitempty"`
}

// ResolveModels returns the generator models to use. Models takes precedence
// over Model when both are set.
func (g *GeneratorConfig) ResolveModels() []string {
	if g == nil {
		return nil
	}
	if len(g.Models) > 0 {
		return g.Models
	}
	if g.Model != "" {
		return []string{g.Model}
	}
	return nil
}

// ReviewerConfig holds all configuration for the review/grading plane.
type ReviewerConfig struct {
	Models       []string    `yaml:"models,omitempty" json:"models,omitempty"`
	SystemPrompt string      `yaml:"system_prompt,omitempty" json:"system_prompt,omitempty"`
	Tools        []ToolEntry `yaml:"tools,omitempty" json:"tools,omitempty"`
}

// SessionLimits configures per-config guardrail limits for evaluation sessions.
// Zero values are ignored and fall back to engine-level defaults.
type SessionLimits struct {
	MaxTurns          int `yaml:"max_turns,omitempty" json:"max_turns,omitempty"`
	MaxFiles          int `yaml:"max_files,omitempty" json:"max_files,omitempty"`
	MaxSessionActions int `yaml:"max_session_actions,omitempty" json:"max_session_actions,omitempty"`
}

// ToolConfig represents a single evaluation configuration.
type ToolConfig struct {
	Name        string           `yaml:"name" json:"name"`
	Description string           `yaml:"description" json:"description"`
	Generator   *GeneratorConfig `yaml:"generator,omitempty" json:"generator,omitempty"`
	Reviewer    *ReviewerConfig  `yaml:"reviewer,omitempty" json:"reviewer,omitempty"`
	Plugins     []string         `yaml:"plugins,omitempty" json:"plugins,omitempty"`
	Limits      *SessionLimits   `yaml:"limits,omitempty" json:"limits,omitempty"`
}

// Normalize resolves plugin references into generator skill directories.
// It is idempotent — safe to call multiple times.
func (tc *ToolConfig) Normalize() {
	if tc.Generator == nil {
		tc.Generator = &GeneratorConfig{}
	}
	if tc.Reviewer == nil {
		tc.Reviewer = &ReviewerConfig{}
	}

	// Resolve installed Copilot CLI plugins to generator skills.
	// Format: "plugin-name@marketplace" (e.g., "azure-sdk-java@skills")
	// Resolves to: ~/.hyoka/cache/{marketplace}/{plugin}/skills/ (preferred)
	// or ~/.copilot/installed-plugins/{marketplace}/{plugin}/skills/ (fallback)
	for _, p := range tc.Plugins {
		if dir := resolveInstalledPlugin(p); dir != "" {
			tc.Generator.Tools = append(tc.Generator.Tools, ToolEntry{
				Name:   p,
				Type:   "skill",
				Source: "local",
				Path:   dir,
			})
			slog.Info("Resolved plugin to skill directory", "plugin", p, "path", dir)
		} else {
			slog.Warn("Could not resolve installed plugin", "plugin", p)
		}
	}
}

// resolveInstalledPlugin resolves a plugin reference (e.g., "azure-sdk-java@skills")
// to the local skills directory. Checks the git-clone cache first (.hyoka/cache/),
// then falls back to ~/.copilot/installed-plugins/ for backward compatibility.
// The format is "plugin-name@marketplace" where marketplace is the source
// (e.g., "skills" from microsoft/skills repo).
// Returns the path to the plugin's skills directory, or empty string if not found.
func resolveInstalledPlugin(ref string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	// Parse "plugin@marketplace" format
	plugin, marketplace := ref, ""
	if idx := len(ref) - 1; idx > 0 {
		for i := idx; i >= 0; i-- {
			if ref[i] == '@' {
				plugin = ref[:i]
				marketplace = ref[i+1:]
				break
			}
		}
	}

	// Special case: "name@skills" is shorthand for microsoft/skills repo
	if marketplace == "skills" {
		// Check .hyoka/cache/default/microsoft/skills/.github/plugins/{name}/
		hyokaCache := filepath.Join(home, ".hyoka", "cache", "default", "microsoft", "skills")
		skillLocations := []string{
			filepath.Join(hyokaCache, ".github", "plugins", plugin),
			filepath.Join(hyokaCache, ".github", "skills", plugin),
			filepath.Join(hyokaCache, "skills", plugin),
		}
		for _, dir := range skillLocations {
			if info, err := os.Stat(dir); err == nil && info.IsDir() {
				if _, err := os.Stat(filepath.Join(dir, "SKILL.md")); err == nil {
					return dir
				}
			}
		}
	}

	// Check ~/.hyoka/cache/ for any version (prefer "default")
	hyokaCache := filepath.Join(home, ".hyoka", "cache", "default")
	if marketplace != "" {
		// Try marketplace as owner/repo pattern
		dir := filepath.Join(hyokaCache, marketplace, plugin, "skills")
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir
		}
	}
	dir := filepath.Join(hyokaCache, plugin, "skills")
	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		return dir
	}

	// Fallback: check ~/.copilot/installed-plugins/ for backwards compatibility.
	basePath := filepath.Join(home, ".copilot", "installed-plugins")
	if marketplace != "" {
		dir := filepath.Join(basePath, marketplace, plugin, "skills")
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir
		}
	}
	dir = filepath.Join(basePath, plugin, "skills")
	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		return dir
	}

	return ""
}

// EffectiveGeneratorSkills returns the generator's skill list from the normalized config.
func (tc *ToolConfig) EffectiveGeneratorSkills() []ToolEntry {
	if tc.Generator != nil {
		return tc.Generator.Tools
	}
	return nil
}

// ConfigFile represents the top-level config file structure.
type ConfigFile struct {
	// PromptDirectory is an optional path that overrides the default prompt
	// discovery (.hyoka/prompts → ./prompts → ../prompts). When set in a
	// config YAML loaded from disk, relative paths are resolved against the
	// containing config file's directory by Load/LoadDir.
	PromptDirectory string `yaml:"prompt_directory,omitempty" json:"prompt_directory,omitempty"`
	// ToolVersionOverride pins specific tool entries (matched by Entry.Name)
	// to a given version. The version is forwarded to the registered Fetcher
	// (for the default npx fetcher, it becomes a git ref). Per-entry
	// `version:` set directly on a tool entry takes precedence over this map.
	// Empty map (or absent field) means "no overrides" — fetcher defaults
	// are used everywhere.
	ToolVersionOverride map[string]string `yaml:"tool_version_override,omitempty" json:"tool_version_override,omitempty"`
	Configs             []ToolConfig      `yaml:"configs"`
}

// ApplyVersionOverrides applies cf.ToolVersionOverride to every tool entry
// in every config (Generator and Reviewer). Entries with a non-empty
// Version field are left untouched (per-entry pin wins). Idempotent.
func (cf *ConfigFile) ApplyVersionOverrides() {
	if cf == nil || len(cf.ToolVersionOverride) == 0 {
		return
	}
	apply := func(entries []ToolEntry) {
		for i := range entries {
			if entries[i].Version != "" {
				continue
			}
			if v, ok := cf.ToolVersionOverride[entries[i].Name]; ok && v != "" {
				entries[i].Version = v
			}
		}
	}
	for i := range cf.Configs {
		if cf.Configs[i].Generator != nil {
			apply(cf.Configs[i].Generator.Tools)
		}
		if cf.Configs[i].Reviewer != nil {
			apply(cf.Configs[i].Reviewer.Tools)
		}
	}
}

// Load reads and parses a configuration file from the given path.
func Load(path string) (*ConfigFile, error) {
	slog.Debug("Loading config file", "path", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	cfg, err := Parse(data)
	if err != nil {
		return nil, err
	}
	if err := cfg.ExpandPlugins(resolvePluginsDir()); err != nil {
		return nil, fmt.Errorf("expanding plugins: %w", err)
	}
	cfg.ApplyVersionOverrides()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	// Resolve a relative prompt_directory against the config file's directory
	// so a config that sits under .hyoka/configs/ can reference ../my-prompts.
	if cfg.PromptDirectory != "" && !filepath.IsAbs(cfg.PromptDirectory) {
		cfg.PromptDirectory = filepath.Join(filepath.Dir(path), cfg.PromptDirectory)
	}
	return cfg, nil
}

// LoadDir reads all .yaml files in a directory and merges their configs.
// This allows splitting configs across multiple files (e.g., baseline.yaml, azure-mcp.yaml).
func LoadDir(dir string) (*ConfigFile, error) {
	slog.Debug("Loading config directory", "dir", dir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading config directory %s: %w", dir, err)
	}

	merged := &ConfigFile{}
	nameSource := make(map[string]string) // config name → source filename
	var promptDirSource string             // filename that first set prompt_directory
	for _, e := range entries {
		if e.IsDir() || (filepath.Ext(e.Name()) != ".yaml" && filepath.Ext(e.Name()) != ".yml") {
			continue
		}
		cf, err := Load(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("loading %s: %w", e.Name(), err)
		}
		if cf.PromptDirectory != "" {
			if merged.PromptDirectory == "" {
				merged.PromptDirectory = cf.PromptDirectory
				promptDirSource = e.Name()
			} else if merged.PromptDirectory != cf.PromptDirectory {
				return nil, fmt.Errorf(
					"conflicting prompt_directory: %q in %s vs %q in %s",
					merged.PromptDirectory, promptDirSource,
					cf.PromptDirectory, e.Name(),
				)
			}
		}
		for _, c := range cf.Configs {
			if prev, ok := nameSource[c.Name]; ok {
				return nil, fmt.Errorf("duplicate config name %q found in files %s and %s", c.Name, prev, e.Name())
			}
			nameSource[c.Name] = e.Name()
		}
		merged.Configs = append(merged.Configs, cf.Configs...)
		// Merge tool_version_override maps. Conflicting values across files
		// are an error — silently last-write-wins would be a footgun.
		for k, v := range cf.ToolVersionOverride {
			if merged.ToolVersionOverride == nil {
				merged.ToolVersionOverride = make(map[string]string)
			}
			if existing, ok := merged.ToolVersionOverride[k]; ok && existing != v {
				return nil, fmt.Errorf("conflicting tool_version_override for %q: %q vs %q (in %s)", k, existing, v, e.Name())
			}
			merged.ToolVersionOverride[k] = v
		}
	}

	if len(merged.Configs) == 0 {
		return nil, fmt.Errorf("no configs found in %s", dir)
	}
	return merged, nil
}

// Parse parses configuration from YAML bytes.
func Parse(data []byte) (*ConfigFile, error) {
	var cfg ConfigFile
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config YAML: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	for _, c := range cfg.Configs {
		models := c.Generator.ResolveModels()
		slog.Info("Config loaded", "name", c.Name, "models", strings.Join(models, ","))
	}
	return &cfg, nil
}

// Validate checks all configs for required fields and constraint violations.
func (cf *ConfigFile) Validate() error {
	if len(cf.Configs) == 0 {
		return fmt.Errorf("no configs defined")
	}
	namesSeen := make(map[string]int, len(cf.Configs))
	for i, c := range cf.Configs {
		if c.Name == "" {
			return fmt.Errorf("config at index %d has no name", i)
		}
		// Generator with a model is required for every config.
		if c.Generator == nil {
			return fmt.Errorf("config %q: generator.model or generator.models is required", c.Name)
		}
		models := c.Generator.ResolveModels()
		if len(models) == 0 {
			return fmt.Errorf("config %q: generator.model or generator.models is required", c.Name)
		}
		seenModels := make(map[string]bool, len(models))
		for _, model := range models {
			if model == "" {
				return fmt.Errorf("config %q: generator model must not be empty", c.Name)
			}
			if seenModels[model] {
				return fmt.Errorf("config %q: duplicate generator model %q", c.Name, model)
			}
			seenModels[model] = true
		}
		if prev, ok := namesSeen[c.Name]; ok {
			return fmt.Errorf("duplicate config name %q at index %d and %d", c.Name, prev, i)
		}
		namesSeen[c.Name] = i
		if c.Generator != nil {
			for j, te := range c.Generator.Tools {
				if err := validateToolEntry(te, c.Name, j); err != nil {
					return err
				}
			}
		}
		if c.Reviewer != nil {
			for j, te := range c.Reviewer.Tools {
				if err := validateToolEntry(te, c.Name, j); err != nil {
					return err
				}
			}
			// Check for duplicate reviewer models
			seen := make(map[string]bool, len(c.Reviewer.Models))
			for _, rm := range c.Reviewer.Models {
				if seen[rm] {
					return fmt.Errorf("config %q: duplicate reviewer model %q", c.Name, rm)
				}
				seen[rm] = true
			}
		}
		// Reject negative session limits — they bypass guardrails.
		if c.Limits != nil {
			if c.Limits.MaxTurns < 0 {
				return fmt.Errorf("config %q: limits.max_turns must not be negative", c.Name)
			}
			if c.Limits.MaxFiles < 0 {
				return fmt.Errorf("config %q: limits.max_files must not be negative", c.Name)
			}
			if c.Limits.MaxSessionActions < 0 {
				return fmt.Errorf("config %q: limits.max_session_actions must not be negative", c.Name)
			}
		}
	}
	return nil
}

// GetConfig returns a config by name, or an error if not found.
func (cf *ConfigFile) GetConfig(name string) (*ToolConfig, error) {
	for i := range cf.Configs {
		if cf.Configs[i].Name == name {
			return &cf.Configs[i], nil
		}
	}
	return nil, fmt.Errorf("config %q not found", name)
}

// GetConfigs returns configs matching the given names. If names is empty, returns all.
func (cf *ConfigFile) GetConfigs(names []string) ([]ToolConfig, error) {
	if len(names) == 0 {
		return cf.Configs, nil
	}
	nameSet := make(map[string]bool, len(names))
	for _, n := range names {
		nameSet[n] = true
	}
	var result []ToolConfig
	for _, c := range cf.Configs {
		if nameSet[c.Name] {
			result = append(result, c)
			delete(nameSet, c.Name)
		}
	}
	if len(nameSet) > 0 {
		var missing []string
		for n := range nameSet {
			missing = append(missing, n)
		}
		return nil, fmt.Errorf("configs not found: %v", missing)
	}
	return result, nil
}

// InstallSkillsAndPlugins is a no-op as of the git-clone resolver implementation.
// Plugin resolution now happens lazily on first use via the gitFetcher, which
// clones repos to the per-eval .skills-cache/ directory. This function remains
// for backward compatibility but does nothing — plugins are resolved at eval time
// by ExpandPlugins and the git-clone fetcher handles the actual git operations
// without any stdout pollution.
func InstallSkillsAndPlugins(configs []ToolConfig) error {
	// No-op: git-clone resolver handles everything lazily
	return nil
}
