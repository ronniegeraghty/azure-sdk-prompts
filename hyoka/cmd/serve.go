package cmd

import (
	"fmt"

	"github.com/ronniegeraghty/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/internal/serve"
	"github.com/spf13/cobra"
)

func serveCmd() *cobra.Command {
	var port int
	var reportsDir string
	var siteDir string
	var docsDir string
	var promptsDir string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start a local web server to browse evaluation reports",
		Long:  "Starts an HTTP server that provides a web UI for browsing past evaluation runs, viewing summaries, and individual report pages.",
		RunE: func(cmd *cobra.Command, args []string) error {
			proj := discoverProject()

			reportsDir = resolvePathFlag(cmd, "output",
				config.ResolveCandidates(proj, "reports", "./reports", "../reports"))
			// site-dir only used when explicitly passed — embedded site is the default
			if cmd.Flags().Changed("site-dir") {
				siteDir, _ = cmd.Flags().GetString("site-dir")
			}
			docsDir = resolvePathFlag(cmd, "docs-dir",
				config.ResolveCandidates(proj, "docs", "./docs", "../docs"))
			promptsDir = resolvePathFlag(cmd, "prompts",
				config.ResolveCandidates(proj, "prompts", "./prompts", "../prompts"))

			fmt.Printf("Listening on http://localhost:%d\n", port)
			return serve.Start(serve.Options{
				ReportsDir: reportsDir,
				SiteDir:    siteDir,
				DocsDir:    docsDir,
				PromptsDir: promptsDir,
				Port:       port,
			})
		},
	}

	cmd.Flags().IntVar(&port, "port", 8080, "Port to serve on")
	cmd.Flags().StringVar(&reportsDir, "output", "./reports", "Directory containing evaluation reports")
	cmd.Flags().StringVar(&siteDir, "site-dir", "", "Directory containing the built React site (default: embedded)")
	cmd.Flags().StringVar(&docsDir, "docs-dir", "", "Directory containing documentation markdown files")
	cmd.Flags().StringVar(&promptsDir, "prompts", "", "Directory containing evaluation prompts")

	return cmd
}
