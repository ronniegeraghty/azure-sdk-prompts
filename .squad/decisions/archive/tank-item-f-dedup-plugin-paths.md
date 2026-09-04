# Tank Item F — Dedup plugin path enumeration

**Author:** Tank 📡
**Date:** 2026-04-27
**Status:** ✅ Implemented
**Spec:** `.squad/decisions/inbox/morpheus-tool-load-consolidation.md` § Item F + § 2 gap #6
**Branch:** `ronniegeraghty/dev`

## What shipped

### 1. New file: `hyoka/internal/plugin/paths.go`

Single source of truth for plugin path enumeration:

- **`PluginCacheCandidates(repoDir, name string) []string`** — returns the
  ordered candidate list `.github/plugins/<name>` → `.github/skills/<name>`
  → `skills/<name>`. Returns `nil` for empty inputs.
- **`FindPluginInRepo(repoDir, name string) (string, error)`** — walks
  `PluginCacheCandidates` using the `IsPluginDir` predicate; on miss returns
  an error that enumerates every checked path (operator-friendly).
- **`IsPluginDir(p string) bool`** — exported wrapper over the existing
  internal `isPluginDir` predicate so other packages don't re-implement it.

### 2. Refactored call sites

- **`internal/plugin/installed.go::ResolveInstalled`** — now calls
  `PluginCacheCandidates(hyokaCache, name)`. Behavior unchanged.
- **`internal/config/tool/validate.go::pluginCheckedPaths`** — replaced the
  three inline `filepath.Join(...)` calls with a single
  `plugin.PluginCacheCandidates(repoCache, name)...` spread.
- **`internal/config/tool/plugin_fetcher.go`** — fully rewritten:
  - Calls `plugin.FindPluginInRepo` directly (deleted the local
    `findPluginInRepo` body).
  - Calls `plugin.SplitOwnerRepo` directly (deleted local `parsePluginRepo`
    body, kept as a one-line shim so existing tests still compile).
  - **Deleted local mirrors:** `isPluginDir`, `isSkillDir`, `hasChildSkills`
    are gone from this file (Neo's note in Item B specifically called these
    out as duplicates to delete in Item F).

### 3. Skill side: NOT deduplicated

`findSkillInRepo` in `internal/config/tool/fetcher.go` has its own 5-path
candidate list (`.github/skills`, `.github/plugins`, `.claude/skills`,
`.agent/skills`, `skills/`) plus manifest fallback plus single-skill
fallback. There is **only one copy** of this list in the codebase — no
duplication exists to dedup. Per Neo's explicit note, plugin precedence
intentionally differs from skill precedence; not unifying.

### 4. Lock-file (`.hyoka-lock`) safety

`acquireRepoLock` (Item C) writes `.hyoka-lock` to
`<CacheRoot>/repos/<owner>/<repo>/` — the **parent** of the version
subdir. `PluginCacheCandidates` produces paths under the version subdir
(e.g. `<repoDir>/.github/plugins/<name>`), so the lock file is one
directory level above any enumeration. No enumerator in the dedup pass
walks the parent dir, so `.hyoka-lock` cannot be mistaken for a plugin or
skill dir. Documented in the `paths.go` doc-comment.

### 5. Legacy `~/.copilot/installed-plugins/` fallback

**Decision:** **Kept for one release with a deprecation `slog.Warn`.**

This is the recommended default per Morpheus's spec (gap #5.4) and was
listed as the recommended option in the task brief. I could not interactively
ask Ronnie because the `ask_user` tool is not available in non-interactive
mode; I went with the recommended default and surfaced the choice here.

Both legacy paths now emit a `slog.Warn` when hit:

```go
slog.Warn("Resolved plugin via deprecated legacy path; will be removed in a future release",
    "path", legacy, "plugin", name, "repo", repo,
    "hint", "move plugin into the canonical hyoka cache layout (see docs/configuration.md)")
```

**Follow-up (next release):** Drop the two legacy `os.UserHomeDir() + ~/.copilot/installed-plugins/...`
branches in `ResolveInstalled`. Also drop the matching paths from
`pluginCheckedPaths` in `validate.go` (those are still hand-rolled in
`validate.go` because they're not in the canonical cache layout — they're
the legacy paths themselves, not part of the dedup'd candidate list).

## Tests

New file `hyoka/internal/plugin/paths_test.go`:

- **`TestPluginCacheCandidates`** — table-driven, 5 cases:
  happy path, empty repoDir, empty plugin name, weird repoDir with
  trailing slash, hyphenated plugin name. All cover precedence order.
- **`TestFindPluginInRepo_Precedence`** — seeds all three candidate
  locations with valid `SKILL.md`-bearing dirs, asserts `.github/plugins`
  wins.
- **`TestFindPluginInRepo_NotFound_EnumeratesAllChecked`** — asserts the
  error message contains all three checked paths (operator UX contract).

Existing tests in `internal/plugin` and `internal/config/tool` continue
to pass — no behavior change beyond the deprecation `slog.Warn`.

## Verification

```
go build ./...                                                        ✅
go test -race ./hyoka/internal/plugin/... ./hyoka/internal/config/tool/... -timeout 3m
  ok  github.com/ronniegeraghty/hyoka/hyoka/internal/plugin           1.073s
  ok  github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool      2.607s
```

Pre-existing failures unchanged and out of scope: `cmd`, `comparison`,
`report`, `serve`, `rerender` (all per Tank Item A baseline — `Model`
field and `boolPtr` issues unrelated to plugin paths).

## Acceptance check

> "Removing or adding a path requires editing one function." (Morpheus Item F)

✅ Adding a new candidate path requires a single edit to
`PluginCacheCandidates` in `hyoka/internal/plugin/paths.go`. All three
historical call sites (`ResolveInstalled`, `pluginCheckedPaths`,
plugin fetcher's `findPluginInRepo`) pick up the change automatically.

## Notes for downstream

- **Switch (release notes):** When the legacy `~/.copilot/installed-plugins/`
  paths are dropped in a future release, mention in CHANGELOG that users
  who saw the deprecation `slog.Warn` need to migrate.
- **Scribe:** No charter or skill changes needed — the dedup is a pure
  internal refactor.
- **Neo:** Your note in `neo-item-b-plugin-remote-fetcher.md` § "Notes for
  Items E / F" is satisfied — `plugin_fetcher.go`'s local
  `isPluginDir`/`isSkillDir`/`hasChildSkills` mirrors are gone, and
  `findPluginInRepo` is now a one-line shim over `plugin.FindPluginInRepo`.
