package eval

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
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

// buildToolIdentities constructs the resolved tool list for MatchContext,
// walking cfg.Generator.Tools and normalizing type via ResolvedType().
func buildToolIdentities(cfg config.ToolConfig) []criteria.ToolIdentity {
	if cfg.Generator == nil {
		return nil
	}
	var tools []criteria.ToolIdentity
	for _, te := range cfg.Generator.Tools {
		if te.Name == "" {
			continue
		}
		t := criteria.ToolIdentity{
			Name:   te.Name,
			Source: te.ResolvedType(),
		}
		// For MCP tools, populate MCPServer (the server name, not per-tool name).
		// Phase 2 doesn't yet implement per-MCP-tool gating — that's deferred.
		if t.Source == tool.TypeMCP {
			t.MCPServer = te.Name // top-level MCP server entry
		}
		tools = append(tools, t)
	}
	return tools
}
