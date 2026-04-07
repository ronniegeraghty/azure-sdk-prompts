// Package plugin implements a composable plugin system for hyoka configs.
//
// A plugin bundles skills, MCP servers, and hooks into a reusable unit
// that can be referenced by name from evaluation configurations.
package plugin

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// Plugin represents a reusable, composable unit of configuration that bundles
// skills, MCP servers, and hooks.
type Plugin struct {
	Name        string                `yaml:"name" json:"name"`
	Description string                `yaml:"description,omitempty" json:"description,omitempty"`
	Skills      []PluginSkill         `yaml:"skills,omitempty" json:"skills,omitempty"`
	MCPServers  map[string]*MCPServer `yaml:"mcp_servers,omitempty" json:"mcp_servers,omitempty"`
	Hooks       *HookConfig           `yaml:"hooks,omitempty" json:"hooks,omitempty"`
	Source      string                `yaml:"-" json:"source,omitempty"` // file path
}

// PluginSkill mirrors the skill fields supported in plugin definitions.
type PluginSkill struct {
	Type string `yaml:"type" json:"type"`
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Repo string `yaml:"repo,omitempty" json:"repo,omitempty"`
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
}

// MCPServer represents an MCP server configuration within a plugin.
type MCPServer struct {
	Type    string   `yaml:"type" json:"type"`
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args" json:"args"`
	Tools   []string `yaml:"tools,omitempty" json:"tools,omitempty"`
}

// HookConfig defines declarative pre/post tool-use hooks.
type HookConfig struct {
	PreToolUse  []string `yaml:"pre_tool_use,omitempty" json:"pre_tool_use,omitempty"`
	PostToolUse []string `yaml:"post_tool_use,omitempty" json:"post_tool_use,omitempty"`
}

// ToolEntry describes a unified tool entry produced from a plugin.
type ToolEntry struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Source   string   `json:"source,omitempty"`
	Path     string   `json:"path,omitempty"`
	Repo     string   `json:"repo,omitempty"`
	Command  string   `json:"command,omitempty"`
	Args     []string `json:"args,omitempty"`
	MCPTools []string `json:"mcp_tools,omitempty"`
}

// Registry holds loaded plugins indexed by name.
type Registry struct {
	plugins map[string]*Plugin
}

// NewRegistry creates an empty plugin registry.
func NewRegistry() *Registry {
	return &Registry{plugins: make(map[string]*Plugin)}
}

// LoadDir loads all plugin YAML files from a directory.
func (r *Registry) LoadDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		slog.Debug("Plugin directory does not exist, skipping", "dir", dir)
		return nil
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}

		p, err := loadPlugin(path)
		if err != nil {
			slog.Warn("Skipping invalid plugin file", "path", path, "error", err)
			return nil
		}
		if _, exists := r.plugins[p.Name]; exists {
			return fmt.Errorf("duplicate plugin name %q (first: %s, second: %s)",
				p.Name, r.plugins[p.Name].Source, path)
		}
		r.plugins[p.Name] = p
		slog.Debug("Loaded plugin", "name", p.Name, "path", path,
			"skills", len(p.Skills), "mcp_servers", len(p.MCPServers))
		return nil
	})
}

func loadPlugin(path string) (*Plugin, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var p Plugin
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&p); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	if p.Name == "" {
		return nil, fmt.Errorf("%s: plugin missing name", path)
	}
	p.Source = path
	return &p, nil
}

// Get returns a plugin by name, or an error if not found.
func (r *Registry) Get(name string) (*Plugin, error) {
	p, ok := r.plugins[name]
	if !ok {
		return nil, fmt.Errorf("plugin %q not found", name)
	}
	return p, nil
}

// List returns all loaded plugin names.
func (r *Registry) List() []string {
	names := make([]string, 0, len(r.plugins))
	for name := range r.plugins {
		names = append(names, name)
	}
	return names
}

// All returns all loaded plugins.
func (r *Registry) All() []*Plugin {
	plugins := make([]*Plugin, 0, len(r.plugins))
	for _, p := range r.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

// Count returns the number of registered plugins.
func (r *Registry) Count() int {
	return len(r.plugins)
}

// ToToolEntries converts the plugin's skills and MCP servers into unified tool entries.
func (p *Plugin) ToToolEntries() []ToolEntry {
	var entries []ToolEntry
	for _, s := range p.Skills {
		entries = append(entries, toolEntryFromPluginSkill(s))
	}
	if len(p.MCPServers) > 0 {
		names := make([]string, 0, len(p.MCPServers))
		for name := range p.MCPServers {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			server := p.MCPServers[name]
			entries = append(entries, ToolEntry{
				Name:     name,
				Type:     "mcp",
				Command:  server.Command,
				Args:     server.Args,
				MCPTools: server.Tools,
			})
		}
	}
	return entries
}

func toolEntryFromPluginSkill(s PluginSkill) ToolEntry {
	name := s.Name
	if name == "" {
		if s.Path != "" {
			name = s.Path
		} else {
			name = s.Repo
		}
	}
	return ToolEntry{
		Name:   name,
		Type:   "skill",
		Source: s.Type,
		Path:   s.Path,
		Repo:   s.Repo,
	}
}
