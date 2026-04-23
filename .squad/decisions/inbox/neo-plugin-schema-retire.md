# Decision: Retire top-level `plugins:` YAML field

**Author:** Neo
**Date:** 2026-04-24
**Status:** Proposed (implemented on `ronniegeraghty/dev`, awaiting Scribe merge into decisions/)
**Related:** `.squad/decisions/inbox/neo-plugin-loading-diagnosis.md`

## Context

The old config schema had two ways to attach a plugin:

1. Top-level `plugins: [name, ...]` on a config — opaque, role-less, auto-appended children to **both** generator and reviewer tool lists at `config.Load` time.
2. `type: plugin` under `generator.tools` / `reviewer.tools` — role-correct but **not actually implemented** in the tool-entry validator (it fell through to an unknown-type branch).

The duplication caused (a) plugins silently landing on reviewers that shouldn't have them, (b) "unknown tool type: plugin" confusion for users who tried the modern shape, and (c) two separate code paths (`ExpandPlugins` vs `ValidateAndExpand`) that disagreed about what counted as a success.

## Decision

Retire the top-level `plugins:` field. Plugins are tool entries.

### Schema

```yaml
configs:
  - name: example
    generator:
      model: claude-opus-4.6
      tools:
        - type: plugin
          name: azure-sdk-python
          source: local          # local | remote | (omit for inference)
        - type: skill
          name: azure-sdk-python-basics
```

- `type: plugin` is a first-class tool entry with `name` (required) and `source` (optional: `local` | `remote`).
- Plugins declared under `generator.tools` apply to the generator role only. Same for reviewer. No auto-dual-role.
- To share, list the plugin in both blocks.

### Resolver

- **Local:** `./.hyoka/plugins/<name>/plugin.yaml` → `./.hyoka/plugins/<name>.yaml` → legacy `./plugins/<name>.yaml`.
- **Remote:** hyoka cache (`~/.hyoka/cache/...`) → legacy `~/.copilot/installed-plugins/`.
- **Inference:** refs containing `@…` default to remote-first; bare names default to local-first.

### Errors

- **Legacy schema:** Parse fails with a migration-hint error the moment a top-level `plugins:` key is seen. Message tells the user to move it under `generator.tools` / `reviewer.tools` as `type: plugin`.
- **Plugin not found:** hard-fail `ToolLoadError` enumerating **every filesystem path inspected** plus an install hint. Surfaced as `tool_load_failure` **before** session creation.

### Fail-fast guarantee

The existing `eval/copilot.go` pre-session gate that promotes `*ToolLoadError` to a hard-fail already covers the new path. No eval runs against a misconfigured plugin.

## Rationale (Ronnie's firm decisions, mapped)

1. **Retire the legacy top-level field** — done. Hard rejection, not deprecation.
2. **No dual-role auto-append** — done. Role inherited from surrounding list.
3. **`source: local|remote`** — done. Local tier preferred for `local`, remote preferred for `remote`.
4. **Enumerated-paths errors** — done. Every candidate path appears in the error body.
5. **Fail-fast in validation phase** — done. Blocks before CreateSession.

## Migration shape

For any config that used the old schema:

```yaml
# before
plugins:
  - azure-sdk-python

# after  (choose role(s))
generator:
  tools:
    - type: plugin
      name: azure-sdk-python
reviewer:
  tools:
    - type: plugin
      name: azure-sdk-python
```

No configs in this repo needed migration — both `configs/` and `.hyoka/configs/` already used the tool-entry shape.

## Verification

- `go build ./...` — clean
- `go test ./...` — all packages pass
- `hyoka validate` — 13 configs, 89 prompts, 3 criteria files valid
- Legacy `plugins:` field rejected with migration hint (verified via synthetic config)
- Missing plugin produces enumerated-path error (verified via synthetic config)

## Follow-ups (not in this wave)

- Move the Azure-Python plugin YAML from `./plugins/` into `./.hyoka/plugins/` so the legacy registry path can be deleted.
- Consider adding a real git-fetch step for remote plugins (currently `ResolveInstalled` is filesystem-only, which is sufficient because cache misses still fail-fast).
