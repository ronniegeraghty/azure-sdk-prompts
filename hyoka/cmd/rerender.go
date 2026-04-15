package cmd

import (
	"fmt"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/rerender"
	"github.com/spf13/cobra"
)

func rerenderCmd() *cobra.Command {
	var reportsDir string
	var all bool

	cmd := &cobra.Command{
		Use:     "rerender [run-id]",
		Aliases: []string{"report"},
		Short:   "Re-render reports from saved JSON without re-running evaluations",
		Long:    `Re-renders report.html, report.md, summary.html, and summary.md from existing report.json data using current templates. No evaluations are re-run. Useful after template improvements. The old name "report" is kept as an alias for backward compatibility.`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reportsDir = resolvePathFlag(cmd, "reports-dir", []string{"../reports", "./reports"})

			var runID string
			if len(args) > 0 {
				runID = args[0]
			}

			if !all && runID == "" {
				return fmt.Errorf("specify a run ID or use --all to re-render all runs")
			}

			return rerender.Run(rerender.Options{
				ReportsDir: reportsDir,
				RunID:      runID,
				All:        all,
			})
		},
	}

	cmd.Flags().StringVar(&reportsDir, "reports-dir", "./reports", "Directory containing evaluation reports")
	cmd.Flags().BoolVar(&all, "all", false, "Re-render all runs")

	return cmd
}
