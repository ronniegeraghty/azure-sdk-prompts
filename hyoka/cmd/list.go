package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	f := &runFlags{}
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List prompts, configs, and criteria",
		Long:    "List prompts matching the given filters alongside available configs and criteria.",
		RunE: func(cmd *cobra.Command, args []string) error {
			f.prompts = resolvePromptsDir(cmd)

			prompts, err := prompt.LoadPrompts(f.prompts)
			if err != nil {
				return fmt.Errorf("loading prompts: %w", err)
			}

			filter := buildFilter(f)
			filtered := prompt.FilterPrompts(prompts, filter)

			// Load configs
			configDir := resolveConfigDir(cmd)
			var configs []config.ToolConfig
			if cfgFile, err := config.LoadDir(configDir); err == nil {
				configs = cfgFile.Configs
			}

			// Load criteria
			baseDir := filepath.Dir(f.prompts)
			criteriaDir := resolveCriteriaDir(cmd)
			if criteriaDir == "" {
				criteriaDir = filepath.Join(baseDir, "criteria")
			}
			var criteriaConfigs []criteria.GraderConfig
			if _, err := os.Stat(criteriaDir); err == nil {
				criteriaConfigs, _ = criteria.LoadDir(criteriaDir)
			}

			if jsonOutput {
				return listJSON(filtered, configs, criteriaConfigs)
			}

			// ── Prompts ───────────────────────────────────────────
			if len(filtered) == 0 {
				fmt.Println("No prompts matched the given filters.")
			} else {
				fmt.Printf("Prompts (%d):\n\n", len(filtered))
				for _, p := range filtered {
					fmt.Printf("  %-40s %s/%s/%s [%s]\n", p.ID, p.Service(), p.Plane(), p.Language(), p.Category())
					if p.Description() != "" {
						fmt.Printf("  %-40s %s\n", "", p.Description())
					}
				}
			}

			// ── Configs ───────────────────────────────────────────
			if len(configs) > 0 {
				fmt.Printf("\nConfigs (%d):\n\n", len(configs))
				for _, c := range configs {
					model := ""
					toolCount := 0
					if c.Generator != nil {
						model = c.Generator.Model
						toolCount = len(c.Generator.Tools)
					}
					fmt.Printf("  %-40s model: %-25s tools: %d\n", c.Name, model, toolCount)
				}
			}

			// ── Criteria ──────────────────────────────────────────
			if len(criteriaConfigs) > 0 {
				fmt.Printf("\nCriteria (%d):\n\n", len(criteriaConfigs))
				for _, gc := range criteriaConfigs {
					source := gc.Source
					if rel, err := filepath.Rel(".", source); err == nil {
						source = rel
					}
					whenParts := make([]string, 0, len(gc.When))
					for k, v := range gc.When {
						whenParts = append(whenParts, k+"="+v)
					}
					whenStr := "*"
					if len(whenParts) > 0 {
						whenStr = strings.Join(whenParts, ", ")
					}
					fmt.Printf("  %-40s when: %-25s graders: %d\n", source, whenStr, len(gc.Graders))
				}
			}

			return nil
		},
	}

	addFilterFlags(cmd, f)
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output as JSON (prompts, configs, criteria)")
	return cmd
}

// listOutput is the JSON structure for `hyoka list --json`.
type listOutput struct {
	Prompts  []*prompt.Prompt    `json:"prompts"`
	Configs  []listConfigEntry   `json:"configs"`
	Criteria []listCriteriaEntry `json:"criteria"`
}

type listConfigEntry struct {
	Name      string `json:"name"`
	Model     string `json:"model"`
	ToolCount int    `json:"tool_count"`
}

type listCriteriaEntry struct {
	Source      string            `json:"source"`
	When        map[string]string `json:"when,omitempty"`
	GraderCount int               `json:"grader_count"`
}

func listJSON(prompts []*prompt.Prompt, configs []config.ToolConfig, criteriaConfigs []criteria.GraderConfig) error {
	out := listOutput{
		Prompts:  prompts,
		Configs:  make([]listConfigEntry, 0, len(configs)),
		Criteria: make([]listCriteriaEntry, 0, len(criteriaConfigs)),
	}
	if out.Prompts == nil {
		out.Prompts = []*prompt.Prompt{}
	}

	for _, c := range configs {
		model := ""
		toolCount := 0
		if c.Generator != nil {
			model = c.Generator.Model
			toolCount = len(c.Generator.Tools)
		}
		out.Configs = append(out.Configs, listConfigEntry{
			Name:      c.Name,
			Model:     model,
			ToolCount: toolCount,
		})
	}

	for _, gc := range criteriaConfigs {
		source := gc.Source
		if rel, err := filepath.Rel(".", source); err == nil {
			source = rel
		}
		out.Criteria = append(out.Criteria, listCriteriaEntry{
			Source:      source,
			When:        gc.When,
			GraderCount: len(gc.Graders),
		})
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling output: %w", err)
	}
	fmt.Println(string(data))
	return nil
}
