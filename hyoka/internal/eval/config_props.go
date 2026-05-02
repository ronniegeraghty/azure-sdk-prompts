package eval

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// injectConfigProps adds eval-config-derived properties to the props map used
// for grader `when:` matching. Phase 2 keys (config-aware structured `tool:`
// blocks are matched via MatchContext.Tools, not props):
//
//   - "generator"  → cfg.Generator.Model
//   - "config"     → cfg.Name
//
// Entries with an empty Name are skipped silently. Keys are written
// unconditionally — they overwrite any prompt frontmatter clash.
func injectConfigProps(props map[string]string, cfg config.ToolConfig) {
	if props == nil {
		return
	}
	props["config"] = cfg.Name
	if cfg.Generator != nil {
		props["generator"] = cfg.Generator.Model
	}
}

// buildToolIdentities constructs the resolved tool list for MatchContext.
//
// Prefers env-derived data when available (SkillsLoaded carries the
// authoritative leaf-skill list expanded from skill_dir entries and
// narrowed by pairwise variants). Falls back to walking cfg for tool
// types env doesn't track at all (built-in tool entries, plugins).
//
// Pass env=nil only for paths that have no session output (e.g., the
// pre-session bundle parse). Skill-aware `when:` filters won't fire
// in that mode, but file-error reporting still works.
func buildToolIdentities(cfg config.ToolConfig, env *report.EnvironmentInfo) []criteria.ToolIdentity {
	if cfg.Generator == nil && env == nil {
		return nil
	}
	var tools []criteria.ToolIdentity

	if env != nil {
		// Skills: env.SkillsLoaded is the authoritative leaf list
		// (skill_dir entries are expanded; pairwise exclusions applied).
		for _, name := range env.SkillsLoaded {
			if name == "" {
				continue
			}
			tools = append(tools, criteria.ToolIdentity{
				Name:   name,
				Source: tool.TypeSkill,
			})
		}
		// MCP servers: env.MCPServers carries the configured set.
		for _, name := range env.MCPServers {
			if name == "" {
				continue
			}
			tools = append(tools, criteria.ToolIdentity{
				Name:      name,
				Source:    tool.TypeMCP,
				MCPServer: name,
			})
		}
	}

	// Built-in tools and plugins live only on cfg — env doesn't surface
	// them as a flat list. Walk cfg for these types and skip skill/mcp
	// (already covered above when env was present).
	if cfg.Generator != nil {
		for _, te := range cfg.Generator.Tools {
			if te.Name == "" {
				continue
			}
			rt := te.ResolvedType()
			if env != nil && (rt == tool.TypeSkill || rt == tool.TypeMCP) {
				continue
			}
			id := criteria.ToolIdentity{
				Name:   te.Name,
				Source: rt,
			}
			if rt == tool.TypeMCP {
				id.MCPServer = te.Name
			}
			tools = append(tools, id)
		}
	}
	return tools
}
