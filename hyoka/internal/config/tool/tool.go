// Package tool provides unified resolution for skills, MCP servers, and plugins.
// It consolidates tool entry types and resolution logic that was previously
// spread across the config and skills packages.
package tool

// Tool type constants.
const (
	TypeTool  = "tool"
	TypeMCP   = "mcp"
	TypeSkill = "skill"
)

// Skill source constants.
const (
	SourceLocal  = "local"
	SourceRemote = "remote"
)
