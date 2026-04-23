package plugin

import (
"os"
"path/filepath"
"strings"
)

// ResolveInstalled resolves an installed plugin to a local skills directory
// using an EXPLICIT repository locator and plugin name. There are no
// shorthand aliases — callers must always pass the repo string from the
// user's config (e.g. "github.com/microsoft/skills" or "microsoft/skills").
//
// It looks for the plugin under the hyoka git-clone cache:
//
//	~/.hyoka/cache/default/<owner>/<repo>/.github/plugins/<name>
//	~/.hyoka/cache/default/<owner>/<repo>/.github/skills/<name>
//	~/.hyoka/cache/default/<owner>/<repo>/skills/<name>
//
// As a transitional fallback for plugins installed via the Copilot CLI
// before hyoka existed, it also checks:
//
//	~/.copilot/installed-plugins/<owner>-<repo>/<name>/skills
//	~/.copilot/installed-plugins/<name>/skills (legacy, repo-less)
//
// Returns the absolute path to the plugin's skills directory, or "" when
// the plugin cannot be located. This function is read-only and has no side
// effects beyond filesystem stats.
func ResolveInstalled(repo, name string) string {
home, err := os.UserHomeDir()
if err != nil {
return ""
}

owner, repoName := SplitOwnerRepo(repo)
if owner != "" && repoName != "" {
hyokaCache := filepath.Join(home, ".hyoka", "cache", "default", owner, repoName)
for _, dir := range []string{
filepath.Join(hyokaCache, ".github", "plugins", name),
filepath.Join(hyokaCache, ".github", "skills", name),
filepath.Join(hyokaCache, "skills", name),
} {
if isSkillDir(dir) {
return dir
}
}

// Copilot CLI legacy install layout, keyed by "<owner>-<repo>".
legacy := filepath.Join(home, ".copilot", "installed-plugins", owner+"-"+repoName, name, "skills")
if isDir(legacy) {
return legacy
}
}

// Legacy ~/.copilot/installed-plugins/<name>/skills (repo-less). Kept
// as a final fallback for users who installed before repo: was required.
if dir := filepath.Join(home, ".copilot", "installed-plugins", name, "skills"); isDir(dir) {
return dir
}
return ""
}

// SplitOwnerRepo normalizes a repo locator into (owner, repo). Accepts:
//
//	"github.com/owner/repo"  → ("owner", "repo")
//	"owner/repo"             → ("owner", "repo")
//	"https://github.com/owner/repo[.git]" → ("owner", "repo")
//
// Returns ("", "") for unparseable input.
func SplitOwnerRepo(repo string) (string, string) {
r := strings.TrimSpace(repo)
if r == "" {
return "", ""
}
r = strings.TrimPrefix(r, "https://")
r = strings.TrimPrefix(r, "http://")
r = strings.TrimPrefix(r, "github.com/")
r = strings.TrimSuffix(r, ".git")
parts := strings.SplitN(r, "/", 3)
if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
return "", ""
}
return parts[0], parts[1]
}

func isDir(p string) bool {
info, err := os.Stat(p)
return err == nil && info.IsDir()
}

func isSkillDir(p string) bool {
if !isDir(p) {
return false
}
_, err := os.Stat(filepath.Join(p, "SKILL.md"))
return err == nil
}
