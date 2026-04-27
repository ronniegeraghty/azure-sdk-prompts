package plugin

import (
	"fmt"
	"path/filepath"
	"strings"
)

// PluginCacheCandidates returns the ordered list of candidate directories to
// search when locating a plugin named `name` inside a freshly-cloned repo
// directory `repoDir`. Precedence (first match wins):
//
//  1. <repoDir>/.github/plugins/<name>
//  2. <repoDir>/.github/skills/<name>
//  3. <repoDir>/skills/<name>
//
// This is the single source of truth for plugin path enumeration. Callers:
//   - plugin.ResolveInstalled (warm-cache lookup)
//   - plugin.FindPluginInRepo (post-clone lookup)
//   - tool.pluginCheckedPaths (error-message enumeration in validate.go)
//
// Plugin precedence intentionally differs from skill precedence
// (see tool.findSkillInRepo) — do not unify the two.
//
// Returns nil if `repoDir` or `name` is empty.
func PluginCacheCandidates(repoDir, name string) []string {
	if repoDir == "" || name == "" {
		return nil
	}
	return []string{
		filepath.Join(repoDir, ".github", "plugins", name),
		filepath.Join(repoDir, ".github", "skills", name),
		filepath.Join(repoDir, "skills", name),
	}
}

// FindPluginInRepo locates a plugin directory within a repo directory using
// PluginCacheCandidates precedence. A directory qualifies as a plugin when
// IsPluginDir returns true (single-skill SKILL.md, or container with
// SKILL.md-bearing children under `skills/`).
//
// Returns the matching dir on success, or an error enumerating every checked
// path so operators can see exactly where the resolver looked. The lock-file
// sentinel `<CacheRoot>/repos/<owner>/<repo>/.hyoka-lock` is never enumerated
// by this function because it lives one level above repoDir.
func FindPluginInRepo(repoDir, name string) (string, error) {
	candidates := PluginCacheCandidates(repoDir, name)
	for _, dir := range candidates {
		if IsPluginDir(dir) {
			return dir, nil
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "plugin %q not found in %s. Checked:", name, repoDir)
	for _, c := range candidates {
		b.WriteString("\n  - ")
		b.WriteString(c)
	}
	return "", fmt.Errorf("%s", b.String())
}

// IsPluginDir is the exported predicate (delegates to the internal isPluginDir)
// so other packages can reuse the same definition without re-implementing it.
func IsPluginDir(p string) bool { return isPluginDir(p) }
