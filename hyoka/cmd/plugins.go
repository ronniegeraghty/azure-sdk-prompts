package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
	"github.com/spf13/cobra"
)

func pluginsCmd() *cobra.Command {
	var pluginsDir string
	var skillsDir string

	cmd := &cobra.Command{
		Use:     "tools",
		Aliases: []string{"plugins"},
		Short:   "List available tools and plugins",
		Long: `Scans the plugins and skills directories and lists all available tool definitions
with their skills, MCP servers, and source locations.

The "plugins" alias is supported for backward compatibility.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Load from both plugins/ and skills/ directories
			type sourceEntry struct {
				plugin *plugin.Plugin
				source string // "plugins" or "skills"
			}
			var entries []sourceEntry

			for _, dir := range []struct {
				path  string
				label string
			}{
				{pluginsDir, "plugins"},
				{skillsDir, "skills"},
			} {
				if _, err := os.Stat(dir.path); os.IsNotExist(err) {
					continue
				}
				reg := plugin.NewRegistry()
				if err := reg.LoadDir(dir.path); err != nil {
					return fmt.Errorf("loading %s from %s: %w", dir.label, dir.path, err)
				}
				for _, p := range reg.All() {
					entries = append(entries, sourceEntry{plugin: p, source: dir.label})
				}
			}

			if len(entries) == 0 {
				fmt.Printf("No tools found in %s or %s\n", pluginsDir, skillsDir)
				return nil
			}

			fmt.Printf("Found %d tool(s):\n\n", len(entries))
			for _, e := range entries {
				p := e.plugin
				// Determine display path
				srcPath := p.Source
				if rel, err := filepath.Rel(".", srcPath); err == nil {
					srcPath = rel
				}

				fmt.Printf("  %s", p.Name)
				if p.Description != "" {
					fmt.Printf(" — %s", p.Description)
				}
				fmt.Println()
				fmt.Printf("    Source: %s (%s)\n", srcPath, e.source)
				if len(p.Skills) > 0 {
					fmt.Printf("    Skills: %d\n", len(p.Skills))
				}
				if len(p.MCPServers) > 0 {
					fmt.Printf("    MCP Servers: %d\n", len(p.MCPServers))
				}
				if p.Hooks != nil {
					hooks := len(p.Hooks.PreToolUse) + len(p.Hooks.PostToolUse)
					if hooks > 0 {
						fmt.Printf("    Hooks: %d\n", hooks)
					}
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&pluginsDir, "plugins-dir", "./plugins", "Directory containing plugin YAML files")
	cmd.Flags().StringVar(&skillsDir, "skills-dir", "./skills", "Directory containing skill YAML files")

	return cmd
}
