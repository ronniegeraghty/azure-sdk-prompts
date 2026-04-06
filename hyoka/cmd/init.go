package cmd

import (
	"fmt"
	"os"

	"github.com/ronniegeraghty/hyoka/internal/config"
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a .hyoka project directory in the current working directory",
		Long: `Scaffolds a .hyoka/ project directory with standard subdirectories:
  configs/   — evaluation config YAML files
  prompts/   — prompt library (.prompt.md files)
  criteria/  — attribute-matched grader criteria
  skills/    — Copilot skills (generator and reviewer)
  reports/   — evaluation output (git-ignored)

A .gitignore is created to exclude the reports/ directory.
Running init again on an existing .hyoka directory is safe (idempotent).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}

			root, err := config.InitProject(cwd)
			if err != nil {
				return fmt.Errorf("initializing project: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Initialized %s\n", root)
			fmt.Fprintln(cmd.OutOrStdout(), "  Created subdirectories: configs/, prompts/, criteria/, skills/, reports/")
			fmt.Fprintln(cmd.OutOrStdout(), "  Created .gitignore (reports/ excluded)")
			return nil
		},
	}
	return cmd
}
