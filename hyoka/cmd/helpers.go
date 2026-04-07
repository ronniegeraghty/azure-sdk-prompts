package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/ronniegeraghty/hyoka/internal/config"
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
	proj := discoverProject()
	candidates := config.ResolveCandidates(proj, "prompts", "./prompts", "../prompts")
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
	return resolvePathFlag(cmd, "criteria-dir", candidates)
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

	resolveSkillRefs := func(refs []config.SkillRef) {
		for j := range refs {
			typ := refs[j].Type
			if typ == "" && refs[j].Path != "" {
				typ = "local"
			}
			if typ != "local" || refs[j].Path == "" {
				continue
			}
			if filepath.IsAbs(refs[j].Path) {
				continue
			}
			candidates := []string{
				refs[j].Path,
				filepath.Join(filepath.Dir(promptsDir), refs[j].Path),
			}
			for _, c := range candidates {
				if info, err := os.Stat(c); err == nil && info.IsDir() {
					abs, absErr := filepath.Abs(c)
					if absErr != nil {
						slog.Warn("Failed to resolve absolute skill path", "path", c, "error", absErr)
						continue
					}
					refs[j].Path = abs
					break
				}
			}
		}
	}

	resolveDirs := func(dirs []string) []string {
		for j := range dirs {
			if filepath.IsAbs(dirs[j]) {
				continue
			}
			candidates := []string{
				dirs[j],
				filepath.Join(filepath.Dir(promptsDir), dirs[j]),
			}
			for _, c := range candidates {
				if info, err := os.Stat(c); err == nil && info.IsDir() {
					abs, absErr := filepath.Abs(c)
					if absErr != nil {
						slog.Warn("Failed to resolve absolute skill path", "path", c, "error", absErr)
						continue
					}
					dirs[j] = abs
					break
				}
			}
		}
		return dirs
	}

	for i := range configs {
		if configs[i].Generator != nil {
			resolveSkills(configs[i].Generator.Tools)
		}
		if configs[i].Reviewer != nil {
			resolveSkills(configs[i].Reviewer.Tools)
			resolveSkillRefs(configs[i].Reviewer.Skills)
		}
		if len(configs[i].ReviewerSkillDirectories) > 0 {
			configs[i].ReviewerSkillDirectories = resolveDirs(configs[i].ReviewerSkillDirectories)
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

// parseByteSize parses a human-readable byte size string (e.g., "1MB", "512KB", "1048576").
func parseByteSize(s string) (int64, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	multipliers := map[string]int64{
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
	}
	for suffix, mult := range multipliers {
		if strings.HasSuffix(s, suffix) {
			numStr := strings.TrimSuffix(s, suffix)
			num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q", numStr)
			}
			return int64(num * float64(mult)), nil
		}
	}
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid size %q \u2014 use a number with optional KB/MB/GB suffix", s)
	}
	return num, nil
}
