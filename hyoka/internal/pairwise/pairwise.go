// Package pairwise generates config variants for pairwise tool-ablation testing.
// Given a base config with N togglable tools, it produces N+1 variants:
// one baseline with all tools enabled, and N variants each with one tool disabled.
package pairwise

import (
	"fmt"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

// ExpandPairwise generates N+1 config variants from a base config.
//   - Variant 0 (baseline): all tools enabled, named "{base}/baseline"
//   - Variant 1..N: each disables one togglable tool, named "{base}/without-{tool}"
//
// Tools with AlwaysOn: true in Generator.Tools are never toggled.
// Both Generator.Tools ([]ToolEntry) and Generator.AvailableTools ([]string)
// are considered; duplicates across the two lists are unified.
func ExpandPairwise(base config.ToolConfig) []config.ToolConfig {
	togglable := collectTogglable(base)

	variants := make([]config.ToolConfig, 0, len(togglable)+1)

	// Baseline: deep copy with all tools intact.
	baseline := cloneToolConfig(base)
	baseline.Name = base.Name + "/baseline"
	variants = append(variants, baseline)

	// One variant per togglable tool, with that tool removed.
	for _, tool := range togglable {
		v := cloneToolConfig(base)
		v.Name = fmt.Sprintf("%s/without-%s", base.Name, tool)
		removeTool(&v, tool)
		variants = append(variants, v)
	}

	return variants
}

// collectTogglable returns deduplicated tool names eligible for toggling.
// Order: ToolEntry names first (preserving order), then AvailableTools.
func collectTogglable(cfg config.ToolConfig) []string {
	if cfg.Generator == nil {
		return nil
	}

	seen := make(map[string]bool)
	var tools []string

	for _, te := range cfg.Generator.Tools {
		if te.AlwaysOn || seen[te.Name] {
			continue
		}
		seen[te.Name] = true
		tools = append(tools, te.Name)
	}

	for _, name := range cfg.Generator.AvailableTools {
		if seen[name] {
			continue
		}
		seen[name] = true
		tools = append(tools, name)
	}

	return tools
}

// removeTool removes a named tool from both Generator.Tools and
// Generator.AvailableTools in the given config.
func removeTool(cfg *config.ToolConfig, name string) {
	if cfg.Generator == nil {
		return
	}

	var tools []config.ToolEntry
	for _, te := range cfg.Generator.Tools {
		if te.Name != name {
			tools = append(tools, te)
		}
	}
	cfg.Generator.Tools = tools

	var avail []string
	for _, t := range cfg.Generator.AvailableTools {
		if t != name {
			avail = append(avail, t)
		}
	}
	cfg.Generator.AvailableTools = avail
}

// cloneToolConfig returns a deep copy of a ToolConfig so mutations to the
// clone never affect the original.
func cloneToolConfig(src config.ToolConfig) config.ToolConfig {
	dst := src // copies value fields (Name, Description, Plugins header)

	if src.Generator != nil {
		gen := *src.Generator

		if len(src.Generator.Skills) > 0 {
			gen.Skills = make([]config.Skill, len(src.Generator.Skills))
			copy(gen.Skills, src.Generator.Skills)
		}

		if len(src.Generator.Tools) > 0 {
			gen.Tools = make([]config.ToolEntry, len(src.Generator.Tools))
			copy(gen.Tools, src.Generator.Tools)
			for i, te := range gen.Tools {
				if te.When != nil {
					m := make(map[string]string, len(te.When))
					for k, v := range te.When {
						m[k] = v
					}
					gen.Tools[i].When = m
				}
			}
		}

		if len(src.Generator.AvailableTools) > 0 {
			gen.AvailableTools = make([]string, len(src.Generator.AvailableTools))
			copy(gen.AvailableTools, src.Generator.AvailableTools)
		}

		if len(src.Generator.ExcludedTools) > 0 {
			gen.ExcludedTools = make([]string, len(src.Generator.ExcludedTools))
			copy(gen.ExcludedTools, src.Generator.ExcludedTools)
		}

		if len(src.Generator.MCPServers) > 0 {
			gen.MCPServers = make(map[string]*config.MCPServer, len(src.Generator.MCPServers))
			for k, v := range src.Generator.MCPServers {
				srv := *v
				if len(v.Args) > 0 {
					srv.Args = make([]string, len(v.Args))
					copy(srv.Args, v.Args)
				}
				if len(v.Tools) > 0 {
					srv.Tools = make([]string, len(v.Tools))
					copy(srv.Tools, v.Tools)
				}
				gen.MCPServers[k] = &srv
			}
		}

		dst.Generator = &gen
	}

	if src.Reviewer != nil {
		rev := *src.Reviewer
		if len(src.Reviewer.Skills) > 0 {
			rev.Skills = make([]config.Skill, len(src.Reviewer.Skills))
			copy(rev.Skills, src.Reviewer.Skills)
		}
		if len(src.Reviewer.Models) > 0 {
			rev.Models = make([]string, len(src.Reviewer.Models))
			copy(rev.Models, src.Reviewer.Models)
		}
		dst.Reviewer = &rev
	}

	if len(src.Plugins) > 0 {
		dst.Plugins = make([]string, len(src.Plugins))
		copy(dst.Plugins, src.Plugins)
	}

	return dst
}
