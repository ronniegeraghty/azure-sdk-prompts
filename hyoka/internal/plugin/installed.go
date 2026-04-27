package plugin

import (
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// ResolveInstalled resolves an installed plugin to a local plugin directory
// using an EXPLICIT repository locator and plugin name. There are no
// shorthand aliases — callers must always pass the repo string from the
// user's config (e.g. "github.com/microsoft/skills" or "microsoft/skills").
//
// It looks for the plugin under the canonical hyoka cache (toolload.CacheRoot)
// at the "default" version segment — matching where gitFetcher writes:
//
//<CacheRoot>/repos/<owner>/<repo>/default/.github/plugins/<name>
//<CacheRoot>/repos/<owner>/<repo>/default/.github/skills/<name>
//<CacheRoot>/repos/<owner>/<repo>/default/skills/<name>
//
// As a transitional fallback for plugins installed via the Copilot CLI
// before hyoka existed, it also checks:
//
//~/.copilot/installed-plugins/<owner>-<repo>/<name>/skills
//~/.copilot/installed-plugins/<name>/skills (legacy, repo-less)
//
// (Item F of the tool-load consolidation plan retires those legacy paths.)
//
// A directory qualifies as a plugin if it is either:
//
//   - a single-skill plugin (top-level SKILL.md), or
//   - a multi-skill container (a `skills/` subdirectory with at least one
//     child directory containing SKILL.md — the standard Copilot
//     `.github/plugins/<name>/` layout used by microsoft/skills).
//
// Returns the absolute path to the plugin's directory, or "" when the
// plugin cannot be located. To enumerate the children of a container
// plugin, pass the returned dir to EnumerateChildSkills.
func ResolveInstalled(repo, name string) string {
	owner, repoName := SplitOwnerRepo(repo)
	if owner != "" && repoName != "" {
		// Canonical cache: <CacheRoot>/repos/<owner>/<repo>/default
		hyokaCache := toolload.RepoCacheDir(owner, repoName, "default")
		for _, dir := range PluginCacheCandidates(hyokaCache, name) {
			if isPluginDir(dir) {
				return dir
			}
		}

		// Copilot CLI legacy install layout, keyed by "<owner>-<repo>".
		// Deprecated — slated for removal one release after Item F lands;
		// see .squad/decisions/inbox/tank-item-f-dedup-plugin-paths.md.
		if home, err := os.UserHomeDir(); err == nil {
			legacy := filepath.Join(home, ".copilot", "installed-plugins", owner+"-"+repoName, name, "skills")
			if isDir(legacy) {
				slog.Warn("Resolved plugin via deprecated legacy path; will be removed in a future release",
					"path", legacy, "plugin", name, "repo", repo,
					"hint", "move plugin into the canonical hyoka cache layout (see docs/configuration.md)")
				return legacy
			}
		}
	}

	// Legacy ~/.copilot/installed-plugins/<name>/skills (repo-less). Kept
	// as a final fallback for users who installed before repo: was required.
	// Deprecated — see note above.
	if home, err := os.UserHomeDir(); err == nil {
		if dir := filepath.Join(home, ".copilot", "installed-plugins", name, "skills"); isDir(dir) {
			slog.Warn("Resolved plugin via deprecated legacy repo-less path; will be removed in a future release",
				"path", dir, "plugin", name,
				"hint", "set repo: in the plugin entry and move into the canonical hyoka cache layout")
			return dir
		}
	}
	return ""
}

// EnumerateChildSkills returns the absolute paths of every SKILL.md-bearing
// child directory directly under <dir>/skills/. Returns nil for single-skill
// plugins (top-level SKILL.md) or directories without a `skills/` subtree.
//
// This is the fan-out counterpart to ResolveInstalled: container plugins
// such as microsoft/skills' `azure-sdk-python` hold dozens of child skills
// under `<plugin>/skills/<child>/SKILL.md`. The validator uses this list
// to record one report row per child so the SDK loads — and the post-session
// verifier checks — the actual leaf skills, not the parent container
// directory (which has no SKILL.md and would never load).
//
// Output is sorted lexicographically for deterministic test/render output.
func EnumerateChildSkills(dir string) []string {
skillsDir := filepath.Join(dir, "skills")
entries, err := os.ReadDir(skillsDir)
if err != nil {
return nil
}
var out []string
for _, e := range entries {
if !e.IsDir() {
continue
}
child := filepath.Join(skillsDir, e.Name())
if _, err := os.Stat(filepath.Join(child, "SKILL.md")); err != nil {
continue
}
abs, err := filepath.Abs(child)
if err != nil {
abs = child
}
out = append(out, abs)
}
sort.Strings(out)
return out
}

// SplitOwnerRepo normalizes a repo locator into (owner, repo). Accepts:
//
//"github.com/owner/repo"  → ("owner", "repo")
//"owner/repo"             → ("owner", "repo")
//"https://github.com/owner/repo[.git]" → ("owner", "repo")
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

// isPluginDir reports whether p is a usable plugin directory: either a
// single-skill plugin (top-level SKILL.md) or a container with a `skills/`
// subdirectory holding at least one SKILL.md-bearing child.
func isPluginDir(p string) bool {
if isSkillDir(p) {
return true
}
return hasChildSkills(p)
}

// hasChildSkills reports whether p contains a `skills/` subdirectory with
// at least one immediate child directory containing SKILL.md.
func hasChildSkills(p string) bool {
skillsDir := filepath.Join(p, "skills")
entries, err := os.ReadDir(skillsDir)
if err != nil {
return false
}
for _, e := range entries {
if !e.IsDir() {
continue
}
if _, err := os.Stat(filepath.Join(skillsDir, e.Name(), "SKILL.md")); err == nil {
return true
}
}
return false
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
