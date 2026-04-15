package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	var withExamples bool

	cmd := &cobra.Command{
		Use:   "init [path]",
		Short: "Initialize a .hyoka project directory",
		Long: `Scaffolds a .hyoka/ project directory with standard subdirectories:
  configs/   — evaluation config YAML files
  prompts/   — prompt library (.prompt.md files)
  criteria/  — attribute-matched grader criteria
  skills/    — Copilot skills (generator and reviewer)
  plugins/   — composable plugin definitions
  reports/   — evaluation output (git-ignored)

An optional path argument specifies where to create the project directory.
If omitted, the current working directory is used.

Use --with-examples to populate configs/ and prompts/ with starter files.

A .gitignore is created to exclude the reports/ directory.
Running init again on an existing .hyoka directory is safe (idempotent).`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var parentDir string
			if len(args) > 0 {
				absPath, err := filepath.Abs(args[0])
				if err != nil {
					return fmt.Errorf("resolving path %q: %w", args[0], err)
				}
				if err := os.MkdirAll(absPath, 0755); err != nil {
					return fmt.Errorf("creating directory %q: %w", absPath, err)
				}
				parentDir = absPath
			} else {
				cwd, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("getting working directory: %w", err)
				}
				parentDir = cwd
			}

			root, err := config.InitProject(parentDir)
			if err != nil {
				return fmt.Errorf("initializing project: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Initialized %s\n", root)
			fmt.Fprintf(cmd.OutOrStdout(), "  Created subdirectories: %s\n", strings.Join(config.ProjectSubdirs, "/, ")+"/")
			fmt.Fprintln(cmd.OutOrStdout(), "  Created .gitignore (reports/ excluded)")

			if withExamples {
				if err := writeExampleFiles(root); err != nil {
					return fmt.Errorf("writing example files: %w", err)
				}
				fmt.Fprintln(cmd.OutOrStdout(), "  Created example config and prompt")
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&withExamples, "with-examples", false, "Create starter example config and prompt files")
	return cmd
}

// writeExampleFiles creates a minimal starter config and prompt in the project.
func writeExampleFiles(root string) error {
	configPath := filepath.Join(root, "configs", "example.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.WriteFile(configPath, []byte(exampleConfig), 0644); err != nil {
			return fmt.Errorf("writing example config: %w", err)
		}
	}

	promptPath := filepath.Join(root, "prompts", "example.prompt.md")
	if _, err := os.Stat(promptPath); os.IsNotExist(err) {
		if err := os.WriteFile(promptPath, []byte(examplePrompt), 0644); err != nil {
			return fmt.Errorf("writing example prompt: %w", err)
		}
	}

	return nil
}

const exampleConfig = `name: example/baseline
description: Example baseline config — replace with your own

generator:
  model: claude-sonnet-4.5

reviewer:
  models:
    - claude-sonnet-4.5
`

const examplePrompt = `---
id: example-dp-python-hello-world
properties:
  service: example
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: "Example prompt — replace with a real evaluation scenario"
  created: "2025-01-01"
  author: hyoka-init
---

## Prompt

Write a simple Python script that prints "Hello, World!".

## Evaluation Criteria

- The script should print exactly "Hello, World!" to stdout.
- The code should be syntactically valid Python 3.
`
