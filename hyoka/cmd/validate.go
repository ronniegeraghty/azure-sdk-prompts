package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/validate"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func noPromptsFoundError(promptsDir string) error {
	nearMisses := prompt.ScanNearMisses(promptsDir)
	fmt.Printf("✗ No prompts found in %s\n", promptsDir)
	if len(nearMisses) > 0 {
		fmt.Println("\nDid you mean one of these?")
		for _, nm := range nearMisses {
			fmt.Printf("  %s\n", nm)
		}
	}
	return fmt.Errorf("no prompts found in %s", promptsDir)
}

func validateCmd() *cobra.Command {
	var promptsDir string

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate prompts, configs, and criteria",
		Long:  "Validate all prompt files against schema rules and naming conventions, validate config files, and validate criteria files.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Peek configs for a prompt_directory override before resolving
			// the prompts dir so `hyoka validate` honors the same override
			// as `hyoka run`. PeekPromptDirectory is best-effort and never
			// fails — full config validation happens later in this command.
			configDir := resolveConfigDir(cmd)
			configPromptDir := config.PeekPromptDirectory(configDir)
			promptsDir = resolvePromptsDirWithConfig(cmd, configPromptDir)
			allOK := true

			// ── Validate prompts ──────────────────────────────────
			result, err := validate.Validate(promptsDir)
			if err != nil {
				return noPromptsFoundError(promptsDir)
			}
			if result.TotalFiles == 0 {
				return noPromptsFoundError(promptsDir)
			}
			fmt.Println(validate.FormatResult(result))
			if !result.OK() {
				allOK = false
			}

			// Check for prompts using old flat format (no properties: key)
			oldFormatCount := detectOldFormatPrompts(promptsDir)
			if oldFormatCount > 0 {
				fmt.Printf("⚠ %d prompt(s) use the old flat frontmatter format (no properties: key)\n", oldFormatCount)
				fmt.Println("  Consider migrating to the properties: map format")
			}

			// ── Validate configs ──────────────────────────────────
			baseDir := filepath.Dir(promptsDir)
			if _, err := os.Stat(configDir); err != nil {
				configDir = filepath.Join(baseDir, "configs")
			}
			if entries, err := os.ReadDir(configDir); err == nil {
				configCount := 0
				configErrors := 0
				oldConfigCount := 0
				for _, e := range entries {
					if e.IsDir() || (filepath.Ext(e.Name()) != ".yaml" && filepath.Ext(e.Name()) != ".yml") {
						continue
					}
					cfgPath := filepath.Join(configDir, e.Name())
					_, cfgErr := config.Load(cfgPath)
					configCount++
					if cfgErr != nil {
						fmt.Printf("✗ Config %s: %v\n", e.Name(), cfgErr)
						configErrors++
						allOK = false
					}
					// Check for old mcp_servers top-level key
					if detectOldConfigFormat(cfgPath) {
						oldConfigCount++
					}
				}
				if configCount > 0 {
					if configErrors == 0 {
						fmt.Printf("✓ All %d config(s) are valid\n", configCount)
					} else {
						fmt.Printf("✗ %d of %d config(s) have errors\n", configErrors, configCount)
					}
					if oldConfigCount > 0 {
						fmt.Printf("⚠ %d config(s) use deprecated mcp_servers top-level key (use generator.tools instead)\n", oldConfigCount)
					}
				}
			}

			// ── Validate criteria ─────────────────────────────────
			criteriaDir := resolveCriteriaDir(cmd)
			if criteriaDir == "" {
				criteriaDir = filepath.Join(baseDir, "criteria")
			}
			if _, err := os.Stat(criteriaDir); err == nil {
				configs, loadErr := criteria.LoadDir(criteriaDir)
				if loadErr != nil {
					fmt.Printf("✗ Criteria: %v\n", loadErr)
					allOK = false
				} else if len(configs) == 0 {
					fmt.Printf("⚠ No criteria files found in %s\n", criteriaDir)
				} else {
					criteriaErrors := 0
					totalGraders := 0
					for _, gc := range configs {
						totalGraders += len(gc.Graders)
						for _, g := range gc.Graders {
							if g.Name == "" {
								fmt.Printf("✗ Criteria %s: grader missing name\n", gc.Source)
								criteriaErrors++
								allOK = false
							}
							if g.Prompt == "" {
								fmt.Printf("✗ Criteria %s: grader %q missing prompt\n", gc.Source, g.Name)
								criteriaErrors++
								allOK = false
							}
						}
					}
					if criteriaErrors == 0 {
						fmt.Printf("✓ All %d criteria file(s) valid (%d grader(s))\n", len(configs), totalGraders)
					} else {
						fmt.Printf("✗ %d error(s) in criteria files\n", criteriaErrors)
					}
				}
			}

			if !allOK {
				return fmt.Errorf("validation failed: one or more prompts, configs, or criteria have errors")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&promptsDir, "prompts", "./prompts", "Path to prompt library directory")
	return cmd
}

// detectOldFormatPrompts scans prompt files for those using the old flat
// frontmatter format (no properties: key in YAML frontmatter).
func detectOldFormatPrompts(dir string) int {
	count := 0
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".prompt.md") {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		text := string(data)
		if !strings.HasPrefix(text, "---") {
			return nil
		}
		parts := strings.SplitN(text[3:], "---", 2)
		if len(parts) < 2 {
			return nil
		}
		frontmatter := parts[0]
		// Check if frontmatter contains a properties: key at root level
		if !strings.Contains(frontmatter, "\nproperties:") && !strings.HasPrefix(frontmatter, "properties:") {
			count++
		}
		return nil
	})
	return count
}

// detectOldConfigFormat checks if a config YAML file uses the deprecated
// mcp_servers top-level key instead of generator.tools.
func detectOldConfigFormat(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	var raw map[string]interface{}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&raw); err != nil {
		return false
	}
	// Check for top-level mcp_servers key
	if _, ok := raw["mcp_servers"]; ok {
		return true
	}
	// Also check inside configs array entries
	if cfgs, ok := raw["configs"]; ok {
		if cfgList, ok := cfgs.([]interface{}); ok {
			for _, item := range cfgList {
				if m, ok := item.(map[string]interface{}); ok {
					if _, ok := m["mcp_servers"]; ok {
						return true
					}
				}
			}
		}
	}
	return false
}
