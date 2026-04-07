package cmd

import (
"fmt"
"os"
"path/filepath"

"github.com/ronniegeraghty/hyoka/internal/config"
"github.com/ronniegeraghty/hyoka/internal/prompt"
"github.com/ronniegeraghty/hyoka/internal/validate"
"github.com/spf13/cobra"
)

func noPromptsFoundError(promptsDir string) error {
	nearMisses := prompt.ScanNearMisses(promptsDir)
	fmt.Printf("\u2717 No prompts found in %s\n", promptsDir)
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
Short: "Validate prompts and configs",
Long:  "Validate all prompt files against schema rules and naming conventions, and validate config files.",
RunE: func(cmd *cobra.Command, args []string) error {
promptsDir = resolvePromptsDir(cmd)
allOK := true

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

configDir := filepath.Join(filepath.Dir(promptsDir), "configs")
if entries, err := os.ReadDir(configDir); err == nil {
configCount := 0
configErrors := 0
for _, e := range entries {
if e.IsDir() || filepath.Ext(e.Name()) != ".yaml" {
continue
}
cfgPath := filepath.Join(configDir, e.Name())
_, cfgErr := config.Load(cfgPath)
configCount++
if cfgErr != nil {
fmt.Printf("\u2717 Config %s: %v\n", e.Name(), cfgErr)
configErrors++
allOK = false
}
}
if configCount > 0 {
if configErrors == 0 {
fmt.Printf("\u2713 All %d config(s) are valid\n", configCount)
} else {
fmt.Printf("\u2717 %d of %d config(s) have errors\n", configErrors, configCount)
}
}
}

if !allOK {
return fmt.Errorf("validation failed: one or more prompts or configs have errors")
}
return nil
},
}

cmd.Flags().StringVar(&promptsDir, "prompts", "./prompts", "Path to prompt library directory")
return cmd
}
