package cmd

import (
	"fmt"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/spf13/cobra"
)

func configsCmd() *cobra.Command {
	var configFile string
	var configDir string

	cmd := &cobra.Command{
		Use:        "configs",
		Short:      "List available configurations",
		Deprecated: "use 'hyoka list' instead, which shows configs alongside prompts and criteria",
		RunE: func(cmd *cobra.Command, args []string) error {
			var cfgFile *config.ConfigFile
			if cmd.Flags().Changed("config-file") {
				configFile = resolveConfigFile(cmd)
				var err error
				cfgFile, err = config.Load(configFile)
				if err != nil {
					return fmt.Errorf("loading config: %w", err)
				}
			} else {
				configDir = resolveConfigDir(cmd)
				var err error
				cfgFile, err = config.LoadDir(configDir)
				if err != nil {
					return fmt.Errorf("loading configs from %s: %w", configDir, err)
				}
			}

			fmt.Printf("Available configurations (%d):\n\n", len(cfgFile.Configs))
			for _, c := range cfgFile.Configs {
				model := ""
				if c.Generator != nil {
					model = c.Generator.Model
				}
				fmt.Printf("  %-20s %s (model: %s)\n", c.Name, c.Description, model)
				var mcpNames []string
				if c.Generator != nil {
					for _, entry := range c.Generator.Tools {
						if entry.ResolvedType() == "mcp" {
							mcpNames = append(mcpNames, entry.Name)
						}
					}
				}
				if len(mcpNames) > 0 {
					fmt.Printf("  %-20s MCP servers: ", "")
					fmt.Println(strings.Join(mcpNames, ", "))
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&configFile, "config-file", "", "Path to a specific configuration YAML file")
	cmd.Flags().StringVar(&configDir, "config-dir", "./configs", "Directory containing configuration YAML files")
	return cmd
}
