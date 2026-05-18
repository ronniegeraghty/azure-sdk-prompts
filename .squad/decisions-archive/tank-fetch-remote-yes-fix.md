# Decision: Remote skill fetcher uses wrong flags — follow-up needed

**Author:** Tank 📡
**Date:** 2025-04-16
**Related:** PR #573

## Context

While fixing the `npx skills add` hang in PR #573 (added `--yes` to stop the interactive prompt), I verified the fix with a real dry-run and discovered that the acute hang bug is only the surface issue. `fetchRemote` and the plugin installer pass flags the skills CLI doesn't understand.

## What the skills CLI actually accepts

From `npx -y skills --help` (`skills` package, current version), the `add` subcommand's real flags are:

- `-s, --skill <skills>` — select specific skill(s) from the repo (accepts `*` for all)
- `-a, --agent <agents>`  — target specific agent dirs (`*` for all)
- `-g, --global`          — install user-level instead of project-level
- `-y, --yes`             — skip confirmation prompts
- `--copy`                — copy instead of symlink
- `--all`                 — shorthand for `-s '*' -a '*' -y`

There is **no** `--directory` and **no** `--name`. hyoka passes both (`fetchRemote` → `args := []string{"skills", "add", entry.Repo, "--directory", installDir}` and `args = append(args, "--name", entry.Name)`). The CLI silently ignores them.

## Observed behavior with `--yes` alone

Dry-run against `examples/configs/example-remote-skill.yaml` (repo `microsoft/skills`, name `copilot-sdk`):

- CLI completes non-interactively ✓ (hang fixed)
- CLI installs `copilot-sdk` **and ~10 other bundled skills** into `<cwd>/.agents/skills/<skill>/` (default agent dir), not our `--directory`
- Hyoka's expected cache path `.skills-cache/microsoft/skills/copilot-sdk/` remains empty
- Dry-run reports `0 skill(s) found` for the generator config
- Side effect: pollutes `.agents/skills/` with 10 unasked-for skills on every fetch

## Proposed follow-up

Rework `fetchRemote` (and audit `InstallSkillsAndPlugins`) to:

1. Use `-s <name>` instead of `--name <name>` so only the requested skill is installed.
2. Drop `--directory` (no-op). Either:
   - Resolve the actual install path under `.agents/skills/<name>/` after the call, or
   - Use `-g --global` + resolve global agent dir, or
   - Use `--agent <dedicated-dir>` if the CLI supports a custom target (doesn't appear to — agents are well-known names like `claude-code`, `cursor`)
3. Decide on `--copy` vs symlinks. For hyoka's caching model, `--copy` is probably safer (self-contained cache, no symlink rot).
4. Consider whether to clean up sibling skills the CLI installs as a side effect, or accept them as acceptable collateral.

## Why not do it in PR #573

- Changes the resolved install path contract returned by `fetchRemote` → ripples through `skills.Resolve` callers and plugin registry loading.
- Needs thought on the caching model (is `.skills-cache/` even the right place if the CLI insists on `.agents/skills/`?).
- PR #573 is scoped to "example config + stop the hang". Keeping it tight makes review easier.

## Recommendation

Open a new issue titled "Remote skill fetcher passes invalid flags to skills CLI; resolved paths are wrong" and assign to Tank. Block any new remote-skill features until this lands.
