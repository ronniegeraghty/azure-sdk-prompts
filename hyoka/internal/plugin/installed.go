package plugin

import (
"os"
"path/filepath"
"strings"
)

// ResolveInstalled resolves a plugin reference (e.g. "azure-sdk-java@skills")
// to a local skills directory. It checks the hyoka git-clone cache first,
// then falls back to ~/.copilot/installed-plugins/.
//
// Returns the absolute path to the plugin's skills directory, or "" when
// the plugin cannot be located. This function is read-only and has no side
// effects beyond filesystem stats.
//
// Mirrors the resolveInstalledPlugin helper in internal/config so the tool
// package (which validates plugins pre-session) can resolve marketplace
// plugin refs without importing internal/config.
func ResolveInstalled(ref string) string {
home, err := os.UserHomeDir()
if err != nil {
return ""
}

name, marketplace := ref, ""
if idx := strings.LastIndex(ref, "@"); idx > 0 {
name = ref[:idx]
marketplace = ref[idx+1:]
}

// Shorthand: name@skills maps to microsoft/skills repo cache.
if marketplace == "skills" {
hyokaCache := filepath.Join(home, ".hyoka", "cache", "default", "microsoft", "skills")
for _, dir := range []string{
filepath.Join(hyokaCache, ".github", "plugins", name),
filepath.Join(hyokaCache, ".github", "skills", name),
filepath.Join(hyokaCache, "skills", name),
} {
if isSkillDir(dir) {
return dir
}
}
}

// hyoka cache (default slot).
hyokaCache := filepath.Join(home, ".hyoka", "cache", "default")
if marketplace != "" {
if dir := filepath.Join(hyokaCache, marketplace, name, "skills"); isDir(dir) {
return dir
}
}
if dir := filepath.Join(hyokaCache, name, "skills"); isDir(dir) {
return dir
}

// Legacy ~/.copilot/installed-plugins/.
base := filepath.Join(home, ".copilot", "installed-plugins")
if marketplace != "" {
if dir := filepath.Join(base, marketplace, name, "skills"); isDir(dir) {
return dir
}
}
if dir := filepath.Join(base, name, "skills"); isDir(dir) {
return dir
}
return ""
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
