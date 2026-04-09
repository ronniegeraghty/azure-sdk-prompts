# Property-based tool filter schema

> Design document for issue [#106](https://github.com/ronniegeraghty/hyoka/issues/106) — task 1.3a

## Problem

Tool availability used to be a flat allowlist (`generator.available_tools`). To support
**pairwise testing** and language-specific tooling, the tool list now lives inside
`generator.tools` entries that can be conditionally enabled per prompt.

## Schema

### `tools` entries on `generator`

```yaml
generator:
  model: claude-opus-4.6
  tools:
    # No condition → always available
    - name: "bash"
      type: tool

    # Include only when the prompt's language is python
    - name: "azure-mcp"
      type: tool
      when:
        language: python

    # Multiple conditions are ANDed
    - name: "cosmosdb-mcp"
      type: tool
      when:
        language: python
        service: cosmos-db
```

Only entries with `type: tool` (or empty type) participate in tool filtering. `type: mcp`
entries configure MCP servers, and `type: skill` entries configure skills. Pairwise
ablation can be configured for any entry type via the `pairwise` field.

### `ToolEntry` (tool filtering fields)

| Field     | Type                | Required | Description |
|-----------|---------------------|----------|-------------|
| `name`    | `string`            | yes      | Tool name (must match SDK tool identifiers) |
| `type`    | `string`            | no       | Defaults to `tool`; only `tool` entries are filtered |
| `when`    | `map[string]string` | no       | Include this tool only when **all** key-value pairs match |
| `always_on` | `bool`            | no       | Never toggled during pairwise tool ablation |
| `pairwise`  | `string`          | no       | Pairwise toggle mode: `off`, `shallow` (default), `deep` |

A `ToolEntry` with no `when` clause is unconditional (always included).

### Matchable prompt properties

These are the `Prompt` struct fields available for matching:

| Key          | Example values                     | Prompt field     |
|--------------|------------------------------------|------------------|
| `language`   | `python`, `dotnet`, `go`, `java`   | `Language`       |
| `service`    | `identity`, `key-vault`, `storage` | `Service`        |
| `plane`      | `data-plane`, `management-plane`   | `Plane`          |
| `category`   | `auth`, `crud`, `pagination`       | `Category`       |
| `difficulty` | `easy`, `medium`, `hard`           | `Difficulty`     |

All comparisons are case-sensitive, exact string equality. Multiple conditions within a
single `when` clause are **ANDed**: every key-value pair must match for the condition
to apply.

## Resolution logic

`ResolveTools` accepts the tool entries and prompt properties and returns the final
`availableTools []string` slice passed to the Copilot session.

```
1. Filter entries to type "tool" (or empty type).
2. Include tools with no `when` clause.
3. Include tools whose `when` map matches the prompt properties.
4. Deduplicate tool names, preserving first-seen order.
```

### Validation rules

1. Every `ToolEntry` must have a non-empty `name`.
2. `type` must be `tool`, `mcp`, or `skill` (empty defaults to `tool`).
3. MCP entries must specify `command`.
4. Skill entries must specify `path` (local) or `repo` (remote).

## Go types

See [`hyoka/internal/config/tool_filter.go`](../hyoka/internal/config/tool_filter.go)
for the implementation.
