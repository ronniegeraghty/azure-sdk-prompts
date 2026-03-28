// Package project provides project-level configuration discovery for hyoka.
//
// When hyoka runs in a repo other than the hyoka repo itself, it looks for a
// hyoka.yaml or .hyoka.yaml file in the working directory (or up to the git
// root) to discover where prompts, configs, skills, criteria, and reports live.
//
// Priority order for each path:
//  1. Explicit CLI flag (e.g., --prompts, --config-dir)
//  2. Project config file value (hyoka.yaml / .hyoka.yaml)
//  3. Default candidates (./prompts, ../prompts, etc.)
package project

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds paths discovered from a hyoka.yaml / .hyoka.yaml project file.
// All paths are stored as-is from the YAML; call Resolve() to make them
// absolute relative to the config file's directory.
type Config struct {
	// PromptsDir is the directory containing .prompt.md files.
	PromptsDir string `yaml:"prompts_dir,omitempty"`
	// ConfigsDir is the directory containing evaluation config YAML files.
	ConfigsDir string `yaml:"configs_dir,omitempty"`
	// SkillsDir is the directory containing skill files.
	SkillsDir string `yaml:"skills_dir,omitempty"`
	// CriteriaDir is the directory containing criteria YAML files.
	CriteriaDir string `yaml:"criteria_dir,omitempty"`
	// ReportsDir is the directory where evaluation reports are written.
	ReportsDir string `yaml:"reports_dir,omitempty"`

	// ConfigPath is the absolute path to the loaded config file (not from YAML).
	ConfigPath string `yaml:"-"`
}

// configFileNames lists filenames to probe, in priority order.
var configFileNames = []string{"hyoka.yaml", ".hyoka.yaml"}

// Load reads a project config from the given explicit path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading project config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing project config %s: %w", path, err)
	}
	abs, _ := filepath.Abs(path)
	cfg.ConfigPath = abs
	cfg.resolve()
	slog.Info("Loaded project config", "path", abs)
	return &cfg, nil
}

// Discover searches for a project config file starting from startDir and
// walking up parent directories until it finds one or reaches the filesystem
// root. Returns nil (no error) if no config file is found — this is the
// normal case when running inside the hyoka repo itself.
func Discover(startDir string) (*Config, error) {
	abs, err := filepath.Abs(startDir)
	if err != nil {
		return nil, fmt.Errorf("resolving start dir: %w", err)
	}

	dir := abs
	for {
		for _, name := range configFileNames {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				slog.Debug("Found project config", "path", candidate)
				return Load(candidate)
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached filesystem root
		}
		dir = parent
	}

	slog.Debug("No project config found", "searched_from", abs)
	return nil, nil
}

// resolve makes all relative paths in the config absolute, based on the
// directory containing the config file.
func (c *Config) resolve() {
	if c.ConfigPath == "" {
		return
	}
	base := filepath.Dir(c.ConfigPath)
	c.PromptsDir = resolveRelative(c.PromptsDir, base)
	c.ConfigsDir = resolveRelative(c.ConfigsDir, base)
	c.SkillsDir = resolveRelative(c.SkillsDir, base)
	c.CriteriaDir = resolveRelative(c.CriteriaDir, base)
	c.ReportsDir = resolveRelative(c.ReportsDir, base)
}

// resolveRelative makes path absolute if it is relative, using base as the
// reference directory. Empty strings are returned as-is.
func resolveRelative(path, base string) string {
	if path == "" || filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(base, path)
}

// EffectivePromptsDir returns the prompts directory to use, considering the
// project config. Returns empty string if not set (caller should fall back to
// default candidates).
func (c *Config) EffectivePromptsDir() string {
	if c == nil {
		return ""
	}
	return c.PromptsDir
}

// EffectiveConfigsDir returns the configs directory to use.
func (c *Config) EffectiveConfigsDir() string {
	if c == nil {
		return ""
	}
	return c.ConfigsDir
}

// EffectiveSkillsDir returns the skills directory to use.
func (c *Config) EffectiveSkillsDir() string {
	if c == nil {
		return ""
	}
	return c.SkillsDir
}

// EffectiveCriteriaDir returns the criteria directory to use.
func (c *Config) EffectiveCriteriaDir() string {
	if c == nil {
		return ""
	}
	return c.CriteriaDir
}

// EffectiveReportsDir returns the reports directory to use.
func (c *Config) EffectiveReportsDir() string {
	if c == nil {
		return ""
	}
	return c.ReportsDir
}
