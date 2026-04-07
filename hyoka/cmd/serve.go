package cmd

import (
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
			siteDir = resolvePathFlag(cmd, "site-dir",
				config.ResolveCandidates(proj, "site/dist", "./site/dist", "../site/dist"))
			docsDir = resolvePathFlag(cmd, "docs-dir",
				config.ResolveCandidates(proj, "docs", "./docs", "../docs"))
			promptsDir = resolvePathFlag(cmd, "prompts",
				config.ResolveCandidates(proj, "prompts", "./prompts", "../prompts"))

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
	cmd.Flags().StringVar(&siteDir, "site-dir", "", "Directory containing the built React site (default: auto-detect site/dist)")
	cmd.Flags().StringVar(&docsDir, "docs-dir", "", "Directory containing documentation markdown files")
	cmd.Flags().StringVar(&promptsDir, "prompts", "", "Directory containing evaluation prompts")

	return cmd
}
