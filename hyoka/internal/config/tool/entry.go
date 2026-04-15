package tool

// Entry represents a tool, MCP server, or skill with optional conditions.
// When the When map is empty, the entry is unconditionally included.
// When the When map has entries, all key-value pairs must match the
// prompt's properties for the entry to be included.
type Entry struct {
	Name     string            `yaml:"name" json:"name"`
	Type     string            `yaml:"type,omitempty" json:"type,omitempty"` // "tool" (default), "mcp", "skill"
	When     map[string]string `yaml:"when,omitempty" json:"when,omitempty"`
	AlwaysOn bool              `yaml:"always_on,omitempty" json:"always_on,omitempty"`
	Pairwise string            `yaml:"pairwise,omitempty" json:"pairwise,omitempty"` // "off", "shallow" (default), "deep"
	// MCP-specific fields
	Command  string   `yaml:"command,omitempty" json:"command,omitempty"`
	Args     []string `yaml:"args,omitempty" json:"args,omitempty"`
	MCPTools []string `yaml:"mcp_tools,omitempty" json:"mcp_tools,omitempty"`
	// Skill-specific fields
	Source   string `yaml:"source,omitempty" json:"source,omitempty"` // "local" or "remote"
	Path     string `yaml:"path,omitempty" json:"path,omitempty"`
	Repo     string `yaml:"repo,omitempty" json:"repo,omitempty"`
	SkillDir bool   `yaml:"skill_dir,omitempty" json:"skill_dir,omitempty"` // true = path is a directory of skills, false = path is a single skill
}

// ResolvedType returns the normalized entry type ("tool" by default).
func (e Entry) ResolvedType() string {
	if e.Type == "" {
		return TypeTool
	}
	return e.Type
}

// ResolvedPairwise returns the effective pairwise setting ("shallow" default).
// AlwaysOn overrides any explicit setting to "off".
func (e Entry) ResolvedPairwise() string {
	if e.AlwaysOn {
		return "off"
	}
	if e.Pairwise == "" {
		return "shallow"
	}
	return e.Pairwise
}

// SkillSource returns the normalized skill source ("local" or "remote").
// If Source is unset, it infers from Path/Repo when possible.
func (e Entry) SkillSource() string {
	if e.Source != "" {
		return e.Source
	}
	if e.Path != "" {
		return SourceLocal
	}
	if e.Repo != "" {
		return SourceRemote
	}
	return ""
}
