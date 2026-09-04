package config

// ResolvePluginsDir returns the first existing plugins directory candidate,
// searching the project root and common relative fallbacks. Returns
// "./plugins" as a final default even when no candidate exists. Exposed
// so pre-session validators (internal/config/tool.ValidateAndExpand) can
// locate the same registry this package uses.
//
// The legacy `./plugins/` directory is still honored for back-compat with
// plugin YAMLs checked into this repo. The preferred layout for local
// plugins moving forward is `./.hyoka/plugins/<name>/plugin.yaml` (mirroring
// the rest of the `.hyoka/` project convention); ValidateAndExpand searches
// that directory directly in addition to the path returned here.
func ResolvePluginsDir() string {
proj := DiscoverFromCWD()
candidates := ResolveCandidates(proj, "plugins", "./plugins", "../plugins")
if len(candidates) > 0 {
return candidates[0]
}
return "./plugins"
}
