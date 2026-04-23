# Decision: Replace npx/copilot-plugin-install with Git-Clone Resolver

**Date:** 2026-04-23  
**Author:** Neo 💊  
**Status:** ✅ Implemented (commit `727a67b0` on `ronniegeraghty/dev`)  
**Context:** CLI Output UX Sprint follow-up — interactive renderer stomped by `npx skills add` stdout

---

## Problem

First real `--pairwise` run revealed that the new interactive progress renderer gets stomped by stdout from the Copilot SDK's `npx skills add` plugin auto-install. The existing implementation in `InstallSkillsAndPlugins` shells out to:
- `npx skills add <package>` for npm-style plugins
- `copilot plugin install <org/repo>` for GitHub repo plugins

Both commands write to stdout uncontrollably, breaking the renderer's live display. Additionally, the install phase is synchronous and blocks eval startup, adding latency to every run.

## Decision

**Stop using `npx skills add` and `copilot plugin install` entirely. Implement a git-clone-based skill resolver that hyoka owns end-to-end.**

### Implementation

#### 1. New `gitFetcher` (replaces `npxFetcher`)

**File:** `hyoka/internal/config/tool/fetcher.go`

**Skill Spec Parsing:**
- `name@skills` → clone `microsoft/skills`, look for skill `name`
  - The `@skills` suffix is shorthand for the microsoft/skills repo (e.g., `azure-sdk-python@skills`)
- `name@owner/repo` → clone `owner/repo`, look for skill `name`
- Bare `owner/repo` (no `@`, currently used by `Entry.Repo`) → clone repo, return root if no name specified

**Cache Layout:**
- Path: `<baseDir>/.skills-cache/<version>/<owner>/<repo>/`
- Preserves existing cache structure for backward compatibility
- Version-pinned caches in separate directories prevent poisoning

**Cache Reuse:**
- Check for `.git` directory first
- If present: `git fetch --all --tags && git checkout <version>`
- If absent: `git clone --branch <version> https://github.com/<owner>/<repo>.git <cacheDir>`
- Default version (empty string or "default") → checkout HEAD (default branch)

**Skill Discovery Order (first match wins):**
1. `.github/skills/<name>/`
2. `.github/plugins/<name>/` (microsoft/skills uses this)
3. `.claude/skills/<name>/`
4. `.agent/skills/<name>/`
5. `skills/<name>/`
6. Top-level `plugin.yaml` or `marketplace.yaml` with custom `skills_dir` (reserved for future)
7. If only one valid skill directory exists, use it (auto-select)

**Validation:**
- A valid skill directory contains either `SKILL.md` or `plugin.yaml`

**Git Output Suppression:**
- All git commands run via `exec.CommandContext` with stdout/stderr captured to a buffer
- Only surface captured stderr if the command exits non-zero (logged at Debug level via slog)
- No direct stdout/stderr writes — the renderer owns the screen

#### 2. Updated `InstallSkillsAndPlugins`

**File:** `hyoka/internal/config/config.go`

**Old behavior:** Deduplicate plugins, shell out to `npx skills add` or `copilot plugin install`, print `"Installing plugin: ..."` to stdout.

**New behavior:** No-op function. Returns `nil` immediately.

**Rationale:** Plugin resolution now happens lazily on first use via the `gitFetcher`. The new `EventToolResolutionStart` / `EventToolResolutionResult` events (from the CLI Output UX sprint) carry tool resolution status to the renderer, so no stdout pollution needed.

#### 3. Updated `resolveInstalledPlugin`

**File:** `hyoka/internal/config/config.go`

**New lookup order:**
1. `.hyoka/cache/default/microsoft/skills/.github/plugins/{name}/` (for `name@skills` shorthand)
2. `.hyoka/cache/default/microsoft/skills/.github/skills/{name}/` (alternate location)
3. `.hyoka/cache/default/microsoft/skills/skills/{name}/` (top-level skills/ dir)
4. `.hyoka/cache/default/{marketplace}/{plugin}/skills/` (generic marketplace format)
5. `.hyoka/cache/default/{plugin}/skills/` (bare plugin name)
6. `~/.copilot/installed-plugins/{marketplace}/{plugin}/skills/` (backward compatibility)
7. `~/.copilot/installed-plugins/{plugin}/skills/` (backward compatibility)

**Special case:** `name@skills` is parsed as `microsoft/skills` repo with plugin name extracted before the `@`.

#### 4. Tests

**File:** `hyoka/internal/config/tool/fetcher_test.go`

**New tests:**
- `TestParseSkillSpec`: Validates all spec parsing rules
  - `azure-sdk-python@skills` → (microsoft, skills, azure-sdk-python)
  - `myskill@acme/widgets` → (acme, widgets, myskill)
  - `github/copilot-skills` + `python-helper` → (github, copilot-skills, python-helper)
  - Bare repo with no name
- `TestFindSkillInRepo`: Skill discovery in `.github/skills/`, `.github/plugins/`, `skills/`
- `TestFindSingleSkill`: Auto-select when exactly one skill exists
- `TestIsValidSkillDir`: Validation via `SKILL.md` or `plugin.yaml` presence

**Updated tests:**
- All existing fetcher tests (`TestRegistry_*`, `TestValidateFetchers`, etc.) updated to use `gitFetcher` instead of `npxFetcher`
- Default fetcher name changed from `"npx"` to `"git"`

---

## Why This Design

### Spec Parsing

**`name@skills` shorthand:**
- Common case for microsoft/skills plugins (e.g., `azure-sdk-python@skills`, `azure-sdk-dotnet@skills`)
- Avoids verbose `name@microsoft/skills` in every config YAML
- Aligns with existing Copilot CLI plugin marketplace conventions

**`name@owner/repo` format:**
- Enables arbitrary repo sources without expanding the Entry struct
- Name and repo are bundled in a single string (Entry.Name field)
- Backward compatible with existing `Entry.Repo` + `Entry.Name` split for standard repos

**Bare `owner/repo`:**
- Preserves existing behavior for repos that are a single skill at the root
- Falls back gracefully when no specific skill name is needed

### Cache Layout

**Per-version directories:**
- Prevents version conflicts when `tool_version_override` changes between evals
- Allows parallel runs with different version pins to use independent caches

**Reuse via `git fetch`:**
- Faster than re-cloning (network bandwidth savings)
- Preserves git history for debugging (can inspect what version was checked out)

**Location under `<baseDir>/.skills-cache/`:**
- Scoped per-eval working directory (default: `.hyoka/`)
- Avoids polluting global user cache (`~/.cache/`, `~/.hyoka/`)
- Survives `hyoka clean` (which only clears reports/, not .skills-cache/)

### Skill Discovery

**Priority order:**
- `.github/plugins/` before `skills/` because microsoft/skills repo uses `.github/plugins/` exclusively
- `.claude/skills/` and `.agent/skills/` for future-proofing other agent frameworks
- Top-level `plugin.yaml` / `marketplace.yaml` reserved for custom skills_dir (not implemented yet)

**Auto-select single skill:**
- Convenience for repos that ship one skill but don't follow standard naming
- Prevents "skill not found" errors when the intent is unambiguous

**Validation via file markers:**
- `SKILL.md` is the primary marker (Copilot CLI convention)
- `plugin.yaml` is secondary (microsoft/skills uses this for plugin metadata)
- Avoids false positives from empty directories

### Git Output Suppression

**Why capture stdout/stderr:**
- `git clone` writes progress bars to stderr (even with `--quiet`, some messages leak)
- `git fetch` writes "Fetching origin" to stderr
- Any of these would stomp the interactive renderer's live display

**Why only log stderr on failure:**
- Successful git operations are expected — no need to clutter logs
- Failed operations need diagnostics — stderr contains the error message
- Logged at Debug level so `--log-level info` stays clean

**Why no direct stdout writes:**
- The renderer owns the screen in interactive mode
- CI mode uses line-buffered output — no live display to preserve
- Progress events (`EventToolResolutionStart` / `Result`) are the official status channel

---

## Alternatives Considered

### 1. Wrap `npx skills add` with output redirection

**Rejected:** Still shells out to an external tool we don't control. If `npx` changes its output format or behavior, we're at its mercy. Also, `npx` requires Node.js to be installed, adding a runtime dependency.

### 2. Use `go-git` library for pure-Go git operations

**Rejected:** Adds a non-stdlib dependency for functionality that `exec.Command("git", ...)` already provides. The stdlib path is simpler and matches existing codebase patterns (we already shell out to `git` in other places). Performance is not a bottleneck here — clones are network-bound, not CPU-bound.

### 3. Keep `npx skills add` but pipe stdout to `/dev/null`

**Rejected:** Doesn't solve the problem — `npx` writes to stderr too (npm warnings, install progress). Also doesn't address the latency problem (synchronous install phase blocking eval startup).

### 4. Pre-install plugins globally via a setup script

**Rejected:** Requires manual setup step before first run. Evals with version overrides would still need per-run installs. Doesn't scale to CI environments where each run is ephemeral.

---

## Migration Path

**Backward compatibility:** Fully preserved.
- Existing configs using `plugins: ["azure-sdk-python@skills"]` work unchanged
- `resolveInstalledPlugin` still checks `~/.copilot/installed-plugins/` as a fallback
- Cache layout under `.skills-cache/<version>/<owner>/<repo>/` is new but doesn't conflict with old paths

**No action required:** Users running evals after this change will automatically use the git-clone resolver. First run will clone repos to the cache; subsequent runs reuse the cache.

**Cleanup (optional):** Users can delete `~/.copilot/installed-plugins/` if they no longer use the Copilot CLI's plugin system outside of hyoka. The git-clone cache lives under `.hyoka/.skills-cache/` and is independent.

---

## Success Criteria

✅ No stdout pollution from skill resolution (git output suppressed)  
✅ Interactive renderer displays tool resolution events cleanly (via `EventToolResolutionStart` / `Result`)  
✅ Cache reuse works (second fetch of same skill is near-instant)  
✅ All existing tests pass (no regressions in skill resolution)  
✅ New tests cover spec parsing, skill discovery, cache reuse  
✅ `InstallSkillsAndPlugins` is a no-op (no pre-eval latency)  
✅ Backward compatibility with `~/.copilot/installed-plugins/` preserved  

---

## Future Work

- **Top-level `plugin.yaml` / `marketplace.yaml` parsing:** If a repo ships a custom `skills_dir` field, honor it. Not implemented yet because no known repos use this pattern.
- **Parallel skill cloning:** Currently sequential. Could use a worker pool to clone multiple repos in parallel (network-bound, not CPU-bound, so parallelism helps). Defer until latency becomes a problem.
- **Smart version resolution:** If `tool_version_override` specifies a semver range (e.g., `^1.2.0`), resolve to the latest matching tag. Currently only exact refs (branches, tags, commits) are supported.
- **Cache expiry / cleanup:** `.skills-cache/` grows unbounded. Could add a `hyoka cache clean` command to prune unused versions.

---

## References

- **Commit:** `727a67b0` on `ronniegeraghty/dev`
- **Files changed:**
  - `hyoka/internal/config/config.go` (InstallSkillsAndPlugins, resolveInstalledPlugin)
  - `hyoka/internal/config/tool/fetcher.go` (gitFetcher replaces npxFetcher)
  - `hyoka/internal/config/tool/fetcher_test.go` (new tests + updated existing tests)
- **Related decisions:**
  - CLI Output UX Overhaul (`.squad/decisions.md` — EventToolResolutionStart / Result schema)
  - Tool Versioning & Custom Fetcher (#597 — Fetcher interface, tool_version_override)
- **microsoft/skills repo:** https://github.com/microsoft/skills/tree/main/.github/plugins/
