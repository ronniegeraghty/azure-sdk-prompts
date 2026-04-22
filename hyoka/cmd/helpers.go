package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/spf13/cobra"
)

// cachedProjectDir holds the lazily discovered .hyoka project directory.
var cachedProjectDir *config.ProjectDir

func discoverProject() *config.ProjectDir {
	if cachedProjectDir == nil {
		cachedProjectDir = config.DiscoverFromCWD()
		if cachedProjectDir.Found() {
			slog.Info("Using .hyoka project directory", "path", cachedProjectDir.Root)
		}
	}
	return cachedProjectDir
}

// resolvePathFlag returns the flag value if explicitly set by the user,
// otherwise tries the candidate paths in order, falling back to the default.
func resolvePathFlag(cmd *cobra.Command, flagName string, candidates []string) string {
	if cmd.Flags().Changed(flagName) {
		val, _ := cmd.Flags().GetString(flagName)
		return val
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	val, _ := cmd.Flags().GetString(flagName)
	return val
}

func resolvePromptsDir(cmd *cobra.Command) string {
	return resolvePromptsDirWithConfig(cmd, "")
}

// resolvePromptsDirWithConfig is the config-aware variant of resolvePromptsDir.
// configPromptDir, when non-empty, takes precedence over the .hyoka/prompts/
// default but is still overridden by an explicit --prompts flag.
func resolvePromptsDirWithConfig(cmd *cobra.Command, configPromptDir string) string {
	proj := discoverProject()
	candidates := config.ResolvePromptDirCandidates(proj, configPromptDir)
	return resolvePathFlag(cmd, "prompts", candidates)
}

func resolveConfigFile(cmd *cobra.Command) string {
	proj := discoverProject()
	candidates := config.ResolveCandidates(proj, "configs", "./configs", "../configs")
	return resolvePathFlag(cmd, "config-file", candidates)
}

func resolveConfigDir(cmd *cobra.Command) string {
	proj := discoverProject()
	candidates := config.ResolveCandidates(proj, "configs", "./configs", "../configs")
	return resolvePathFlag(cmd, "config-dir", candidates)
}

func resolveOutputDir(cmd *cobra.Command) string {
	proj := discoverProject()
	candidates := config.ResolveCandidates(proj, "reports", "./reports", "../reports")
	return resolvePathFlag(cmd, "output", candidates)
}

func resolveOutputFile(cmd *cobra.Command, candidates []string) string {
	return resolvePathFlag(cmd, "output", candidates)
}

func resolveCriteriaDir(cmd *cobra.Command) string {
	proj := discoverProject()
	candidates := config.ResolveCandidates(proj, "criteria", "./criteria", "../criteria")
	resolved := resolvePathFlag(cmd, "criteria-dir", candidates)
	if resolved != "" {
		slog.Debug("Resolved criteria directory", "dir", resolved, "candidates", candidates)
	} else if len(candidates) > 0 {
		slog.Warn("Criteria candidates found but none selected", "candidates", candidates)
	}
	return resolved
}

// resolveConfigSkillDirs resolves relative local skill paths in loaded configs
// to absolute paths so they work regardless of which directory the tool is invoked from.
func resolveConfigSkillDirs(configs []config.ToolConfig, promptsDir string) {

	resolveSkills := func(entries []config.ToolEntry) {
		for j := range entries {
			if entries[j].ResolvedType() != "skill" || entries[j].SkillSource() != "local" || entries[j].Path == "" {
				continue
			}
			if filepath.IsAbs(entries[j].Path) {
				continue
			}
			candidates := []string{
				entries[j].Path,
				filepath.Join(filepath.Dir(promptsDir), entries[j].Path),
			}
			for _, c := range candidates {
				if info, err := os.Stat(c); err == nil && info.IsDir() {
					abs, absErr := filepath.Abs(c)
					if absErr != nil {
						slog.Warn("Failed to resolve absolute skill path", "path", c, "error", absErr)
					}
					entries[j].Path = abs
					break
				}
			}
		}
	}

	for i := range configs {
		if configs[i].Generator != nil {
			resolveSkills(configs[i].Generator.Tools)
		}
		if configs[i].Reviewer != nil {
			resolveSkills(configs[i].Reviewer.Tools)
		}
	}
}

func humanSize(b int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
	)
	switch {
	case b >= gb:
		return fmt.Sprintf("%.1fGB", float64(b)/float64(gb))
	case b >= mb:
		return fmt.Sprintf("%.1fMB", float64(b)/float64(mb))
	case b >= kb:
		return fmt.Sprintf("%.1fKB", float64(b)/float64(kb))
	default:
		return fmt.Sprintf("%dB", b)
	}
}

// openInBrowser opens the given file path in the default browser.
func openInBrowser(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", path)
	default:
		fmt.Printf("Open the report manually: %s\n", path)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Could not open browser: %v\nOpen manually: %s\n", err, path)
	}
}
