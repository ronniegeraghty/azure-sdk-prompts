# Decision: Remote plugin source requires `@marketplace` locator in name

**Agent:** Neo 💊
**Date:** 2026-04-27
**Branch:** ronniegeraghty/dev
**Driver:** Ronnie asked: "How does `source: remote` without a URL tell hyoka where to pull the plugin from?"

## Finding

`configs/python-pairwise.yaml` declared `{name: azure-sdk-python, type: plugin, source: remote}` with no locator field. There is, in fact, no dedicated `repo:` / `url:` field on plugin entries. The locator is encoded as a `@marketplace` suffix on `name` (e.g. `azure-sdk-python@skills` → microsoft/skills marketplace cache), parsed by `plugin.ResolveInstalled`.

Without the suffix, the resolver falls through to bare-name lookups under `~/.hyoka/cache/default/<name>/skills` and `~/.copilot/installed-plugins/<name>/skills` — which only match if something already placed the plugin at that exact path. No pre-validation caught this; the failure surfaced at eval time with a noisy "Checked:" path dump, producing the confusing symptom Ronnie flagged.

Not a missing feature — a **schema gap** that silently produced bad configs.

## Decision

1. **Keep the `@marketplace` locator convention.** Remote plugin resolution by name is the existing design (matches `/plugin install name@skills` UX). Adding a full `repo:` field on plugin entries would duplicate the skill-fetcher pipeline without need — remote plugin fetch today is handled by the Copilot CLI, not hyoka.

2. **Reject bare `source: remote` at validation time.** `validatePluginEntry` now fails fast with a fix-it message when a remote plugin name has no `@marketplace` suffix, instead of falling through to the noisy resolver dump.

3. **Fix the two broken configs** to use `azure-sdk-<lang>@skills`:
   - `configs/python-pairwise.yaml` (1 entry)
   - `configs/baseline-sonnet-skills.yaml` (5 entries)

4. **Future work (not blocking):** If auto-fetch for plugins is ever wanted, mirror the skill flow with explicit `repo:` / `ref:` fields on plugin entries. Today, `@marketplace` is sufficient and is the only supported locator.

## Files Changed

- `hyoka/internal/config/tool/validate.go` — new guard at top of `validatePluginEntry`
- `hyoka/internal/config/tool/validate_test.go` — `TestValidateAndExpand_RemotePluginMissingLocator`
- `configs/python-pairwise.yaml` — rename + doc comment
- `configs/baseline-sonnet-skills.yaml` — rename 5 plugin entries

## Verification

- `go build ./...` — clean
- `go test ./...` — all green (eval package 8.9s, everything else cached)
- `hyoka validate` — 89 prompts, 13 configs, 3 criteria files all valid

## Reusable Pattern

Any tool entry that references remote content **must** carry an explicit locator. For skills that's `repo:` (+ optional `ref:` / `version:`). For plugins that's an `@marketplace` suffix on `name`. Validation should reject remote entries missing their locator rather than letting the failure surface downstream.
