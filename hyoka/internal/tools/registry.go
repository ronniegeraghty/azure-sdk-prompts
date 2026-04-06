// Package tools provides a typed registry for curated tool configurations
// (MCP servers, Copilot skills) that evaluation configs can reference by name.
package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

// Supported tool types.
const (
	TypeMCP   = "mcp"
	TypeSkill = "skill"
)

// Supported skill sources.
const (
	SourceLocal  = "local"
	SourceRemote = "remote"
)

// ToolRegistry is the top-level catalog of curated tool configurations.
type ToolRegistry struct {
	Version string      `yaml:"version" json:"version"`
	Tools   []ToolEntry `yaml:"tools"   json:"tools"`
}

// ToolEntry describes a single tool that configs can reference by name.
type ToolEntry struct {
	Name        string       `yaml:"name"                 json:"name"`
	Type        string       `yaml:"type"                 json:"type"`
	Description string       `yaml:"description,omitempty" json:"description,omitempty"`
	Version     string       `yaml:"version,omitempty"    json:"version,omitempty"`
	MCP         *MCPConfig   `yaml:"mcp,omitempty"        json:"mcp,omitempty"`
	Skill       *SkillConfig `yaml:"skill,omitempty"      json:"skill,omitempty"`
}

// MCPConfig holds connection details for an MCP server tool.
type MCPConfig struct {
	Command string            `yaml:"command"          json:"command"`
	Args    []string          `yaml:"args,omitempty"   json:"args,omitempty"`
	Env     map[string]string `yaml:"env,omitempty"    json:"env,omitempty"`
	Tools   []string          `yaml:"tools,omitempty"  json:"tools,omitempty"`
}

// SkillConfig holds resolution details for a Copilot skill.
type SkillConfig struct {
	Source string `yaml:"source"          json:"source"`
	Repo   string `yaml:"repo,omitempty"  json:"repo,omitempty"`
	Name   string `yaml:"name,omitempty"  json:"name,omitempty"`
	Path   string `yaml:"path,omitempty"  json:"path,omitempty"`
}

// Get returns the tool entry with the given name and true, or a zero value
// and false if no entry matches.
func (r *ToolRegistry) Get(name string) (ToolEntry, bool) {
	for _, t := range r.Tools {
		if t.Name == name {
			return t, true
		}
	}
	return ToolEntry{}, false
}

// ParseRegistry unmarshals YAML bytes into a ToolRegistry and validates it.
func ParseRegistry(data []byte) (*ToolRegistry, error) {
	var reg ToolRegistry
	if err := yaml.Unmarshal(data, &reg); err != nil {
		return nil, fmt.Errorf("parsing registry YAML: %w", err)
	}
	if err := validateRegistry(&reg); err != nil {
		return nil, err
	}
	return &reg, nil
}

// LoadRegistry reads a local YAML file and returns a validated ToolRegistry.
func LoadRegistry(path string) (*ToolRegistry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading registry file %s: %w", path, err)
	}
	return ParseRegistry(data)
}

// LoadRemoteRegistry fetches a YAML registry from a URL and returns a
// validated ToolRegistry.
func LoadRemoteRegistry(url string) (*ToolRegistry, error) {
	resp, err := http.Get(url) //nolint:gosec // URL is caller-provided
	if err != nil {
		return nil, fmt.Errorf("fetching remote registry %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote registry %s returned status %d", url, resp.StatusCode)
	}

	const maxSize = 2 * 1024 * 1024 // 2 MB
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxSize))
	if err != nil {
		return nil, fmt.Errorf("reading remote registry body: %w", err)
	}
	return ParseRegistry(data)
}

// validateRegistry checks structural invariants on a parsed registry.
func validateRegistry(reg *ToolRegistry) error {
	if reg.Version != "1" {
		return fmt.Errorf("unsupported registry version %q (expected \"1\")", reg.Version)
	}

	seen := make(map[string]bool, len(reg.Tools))
	for i, t := range reg.Tools {
		if t.Name == "" {
			return fmt.Errorf("tool at index %d: name is required", i)
		}
		if seen[t.Name] {
			return fmt.Errorf("duplicate tool name %q", t.Name)
		}
		seen[t.Name] = true

		switch t.Type {
		case TypeMCP:
			if err := validateMCP(t); err != nil {
				return fmt.Errorf("tool %q: %w", t.Name, err)
			}
		case TypeSkill:
			if err := validateSkill(t); err != nil {
				return fmt.Errorf("tool %q: %w", t.Name, err)
			}
		default:
			return fmt.Errorf("tool %q: invalid type %q (expected %q or %q)", t.Name, t.Type, TypeMCP, TypeSkill)
		}
	}
	return nil
}

func validateMCP(t ToolEntry) error {
	if t.MCP == nil {
		return fmt.Errorf("mcp block is required for type %q", TypeMCP)
	}
	if t.MCP.Command == "" {
		return fmt.Errorf("mcp.command is required")
	}
	return nil
}

func validateSkill(t ToolEntry) error {
	if t.Skill == nil {
		return fmt.Errorf("skill block is required for type %q", TypeSkill)
	}
	switch t.Skill.Source {
	case SourceRemote:
		if t.Skill.Repo == "" {
			return fmt.Errorf("skill.repo is required for remote skills")
		}
		if t.Skill.Name == "" {
			return fmt.Errorf("skill.name is required for remote skills")
		}
	case SourceLocal:
		if t.Skill.Path == "" {
			return fmt.Errorf("skill.path is required for local skills")
		}
	default:
		return fmt.Errorf("invalid skill.source %q (expected %q or %q)", t.Skill.Source, SourceLocal, SourceRemote)
	}
	return nil
}
