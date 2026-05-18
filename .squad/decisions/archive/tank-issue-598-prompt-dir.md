# Decision: Tank — Configurable Prompt Directory (#598)

**Author:** Tank 📡
**Date:** 2026-04-21
**Status:** Implemented (PR pending)
**Issue:** #598
**Branch:** `ronniegeraghty/issue-598-configurable-prompt-dir`
**Target:** `phase-6`

## What Changed

Added a top-level `prompt_directory:` field to config YAML so users can point hyoka at a prompt library outside `.hyoka/prompts/`. Fully backwards compatible — when the field is absent, discovery is unchanged.

## Surfaces

- **Config key (new):** `prompt_directory:` (string, optional). Lives at the **top level** of a config YAML file (sibling of `configs:`). Relative paths resolve against the config file's directory; absolute paths are honored as-is.
- **CLI flag (existing):** `--prompts` still works as the highest-priority override.
- **No new env var.** Kept scope minimal per issue.

## Resolution Priority

1. `--prompts` CLI flag (existing user override)
2. `prompt_directory:` from a loaded config YAML (new)
3. `.hyoka/prompts/` (default since project-dir feature)
4. `./prompts/` (legacy fallback)
5. `../prompts/` (legacy fallback)

## Migration Story

- **Existing users:** zero action required. Default discovery unchanged.
- **New users:** `hyoka init` still scaffolds `.hyoka/prompts/`. They can opt into a custom layout by adding `prompt_directory:` to any config YAML in `.hyoka/configs/`.
- **Conflict handling:** If two config files in the same `--config-dir` declare different `prompt_directory:` values, `LoadDir` errors with both filenames. Identical values are accepted (no-op).

## Files Touched (Tank scope)

- `hyoka/internal/config/config.go` — `ConfigFile.PromptDirectory` field, Load/LoadDir propagation + conflict check
- `hyoka/internal/config/discovery.go` — `ResolvePromptDirCandidates`, `PeekPromptDirectory`
- `hyoka/cmd/helpers.go` — `resolvePromptsDirWithConfig`
- `hyoka/cmd/run.go` — load configs first, then resolve prompts dir
- `hyoka/cmd/validate.go` — peek configs for override
- `hyoka/cmd/list.go` — same
- `hyoka/cmd/serve.go` — same
- `hyoka/internal/config/prompt_dir_test.go` — 11 new tests covering Load, LoadDir, ResolvePromptDirCandidates, PeekPromptDirectory
- `docs/configuration.md` — new "Custom Prompt Directory" section

## Verification

Manual end-to-end smoke test in `.scratch-598/`:

| Scenario | Expected | Result |
|---|---|---|
| `hyoka init --with-examples` then `hyoka list` | finds `.hyoka/prompts/example.prompt.md` | ✅ |
| Config sets `prompt_directory: ../../my-prompts`; decoy in `.hyoka/prompts/` | `list` shows custom prompt only | ✅ |
| `--prompts ./.hyoka/prompts` overrides config | shows decoy | ✅ |
| `hyoka validate` honors `prompt_directory` | validates custom dir | ✅ |

`go test -race ./hyoka/...` passes for all owned packages.

## Coordination Note

Neo is concurrently implementing #580 in the same worktree (no separate worktree was created for that branch), which dropped uncommitted `criteria/buckets.go` and `eval/engine.go` changes into my working tree. I scoped my commit to only my own files; Neo's WIP is left in place for them to commit on their branch.
