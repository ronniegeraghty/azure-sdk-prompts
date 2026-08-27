# Tool Grader

The `tool` grader validates tool usage patterns during code generation. It checks which tools (skills, plugins, MCP servers) were invoked, call counts, and group membership. The grader passes only if **all** checks pass (boolean semantics).

This canonical grader consolidates and replaces the legacy `tool_constraint`, `behavior`, and `tool_usage` graders.

## When to Use

- Verify required tools were used (e.g., bash, markdown-lists skill)
- Forbid use of specific tools (e.g., dangerous_tool)
- Enforce tool usage from groups (e.g., "must use at least one logging tool")
- Verify call counts for performance or safety (e.g., "bash called 2–10 times")
- Prevent overuse of single tools (e.g., "max 5 bash invocations")

## Configuration

The `tool` grader uses a top-level `checks:` array. Each check validates tool availability and usage:

```yaml
graders:
  - name: Tool Usage
    type: tool
    weight: 0.5
    checks:
      - kind: tool_used
        tool: bash
        min_calls: 1
      - kind: tool_not_used
        tool: dangerous_tool
```

### `checks` Schema

Four check kinds are supported:

#### 1. `tool_used`

A named tool must be invoked at least once. Optional `min_calls` and `max_calls` bounds. Use `source` to filter by tool type (skill/mcp/builtin). To accept any tool from one MCP server, omit `tool` and set both `source: mcp` and `mcp_server`.

| Field      | Type     | Required | Description                              |
|------------|----------|----------|------------------------------------------|
| `kind`     | string   | yes      | Must be `"tool_used"` |
| `tool`     | string   | conditional | Tool name. Optional only when `source: mcp` and `mcp_server` are set. |
| `min_calls` | int     | no       | Minimum number of invocations (default: 1 if present is required) |
| `max_calls` | int     | no       | Maximum number of invocations |
| `source`   | string   | no       | Filter by tool source: `skill`, `mcp`, or `builtin`. If omitted, matches any source. |
| `mcp_server` | string | no       | MCP server name (only meaningful with `source: mcp`). Filters to a specific MCP server. |

#### 2. `tool_not_used`

A named tool must NOT be invoked at all. Use `source` to filter by tool type. To forbid every tool from one MCP server, omit `tool` and set both `source: mcp` and `mcp_server`.

| Field      | Type     | Required | Description                              |
|------------|----------|----------|------------------------------------------|
| `kind`     | string   | yes      | Must be `"tool_not_used"` |
| `tool`     | string   | conditional | Tool name. Optional only when `source: mcp` and `mcp_server` are set. |
| `source`   | string   | no       | Filter by tool source: `skill`, `mcp`, or `builtin`. If omitted, matches any source. |
| `mcp_server` | string | no       | MCP server name (only meaningful with `source: mcp`). Filters to a specific MCP server. |

#### 3. `any_from_group`

At least one tool from a named group must be invoked. Optional `except:` list to exclude specific tools from the group.

| Field      | Type       | Required | Description                              |
|------------|------------|----------|------------------------------------------|
| `kind`     | string     | yes      | Must be `"any_from_group"` |
| `group`    | string     | yes      | Group name (defined in tool topology) |
| `except`   | []string   | no       | Tools to exclude from the group check |

#### 4. `none_from_group`

No tool from a named group must be invoked. Optional `except:` list to exclude specific tools from the group (i.e., those exceptions ARE allowed).

| Field      | Type       | Required | Description                              |
|------------|------------|----------|------------------------------------------|
| `kind`     | string     | yes      | Must be `"none_from_group"` |
| `group`    | string     | yes      | Group name (defined in tool topology) |
| `except`   | []string   | no       | Tools to exclude from the group check (i.e., these tools are allowed) |

## Examples

### Basic Tool Requirements

```yaml
graders:
  - name: Required Tools
    type: tool
    weight: 0.3
    checks:
      - kind: tool_used
        tool: bash
      - kind: tool_used
        tool: file_search
```

### Tool Call Bounds

```yaml
graders:
  - name: Efficient Tool Usage
    type: tool
    weight: 0.4
    when:
      language: go
    checks:
      - kind: tool_used
        tool: bash
        min_calls: 1
        max_calls: 5
      - kind: tool_used
        tool: file_search
        min_calls: 1
        max_calls: 10
```

### Any tool from an MCP server

```yaml
graders:
  - name: Azure MCP Used
    type: tool
    checks:
      - kind: tool_used
        source: mcp
        mcp_server: azure
```

### Forbid Dangerous Tools

```yaml
graders:
  - name: No Dangerous Tools
    type: tool
    weight: 0.2
    checks:
      - kind: tool_not_used
        tool: shell_escape
      - kind: tool_not_used
        tool: exec_raw
```

### Group-Based Checks

```yaml
graders:
  - name: Logging Tools
    type: tool
    weight: 0.2
    checks:
      # Must use at least one tool from the logging group
      - kind: any_from_group
        group: logging_tools
      # But not all forbidden tools
      - kind: none_from_group
        group: dangerous_operations
        except: [read_file]  # read_file is allowed even though it's in dangerous_operations
```

### Comprehensive Tool Validation

```yaml
graders:
  - name: Comprehensive Tool Usage
    type: tool
    weight: 0.5
    checks:
      - kind: tool_used
        tool: bash
        min_calls: 1
        max_calls: 5
      - kind: tool_not_used
        tool: deprecated_api
      - kind: any_from_group
        group: markdown_tools
      - kind: none_from_group
        group: dangerous_tools
        except: [bash]  # bash is dangerous but we're allowing it
```

### Filtering by Tool Source

```yaml
graders:
  - name: MCP Server Usage
    type: tool
    weight: 0.3
    checks:
      # Must use a specific skill (not MCP or builtin)
      - kind: tool_used
        tool: markdown-headings
        source: skill
      # Must use a specific MCP server
      - kind: tool_used
        tool: file-search
        source: mcp
        mcp_server: native-tools
      # Forbid any builtin tool
      - kind: tool_not_used
        tool: bash
        source: builtin
```

## Result Structure

Each tool grader produces:
- **Pass/Fail**: Binary result (true only if ALL checks pass)
- **Check results**: Individual pass/fail for each configured check
- **Tool usage summary**: Which tools were invoked and how many times
- **Score**: 1.0 if all checks pass, 0.0 if any check fails (boolean grader)

Results visible in evaluation reports under `grader_results`.

## Data Visible to Grader

Tool graders can access:
- **EnvironmentTools**: List of tools/skills/MCP servers available in the generator environment
- **Action log**: Tool invocations and their counts (from ActionLog)
- **SkillsInvoked**: Set of skill names actually invoked (from skill.invoked events)
- **MCPServersUsed**: Set of MCP server names that recorded at least one tool call

## Tool Name Resolution

Tool names in checks are matched against:
1. **Skill names** (from skill directories or remote skills)
2. **Plugin names** (from plugins with exported tools)
3. **MCP server names** (from MCP server configurations)

Matching is case-sensitive. Use the exact name as configured in your evaluation config.

## Groups

Groups are defined in the tool topology and represent collections of related tools. Examples:
- `logging_tools` — logging/instrumentation skills
- `markdown_tools` — markdown generation skills
- `dangerous_tools` — tools with security implications

Consult your evaluation configuration or skill directory to see available groups.

## Notes

- **All-or-nothing**: Grader passes only if **every** check passes. A single failed check fails the grader.
- **Call count semantics**: `min_calls` and `max_calls` apply only to `kind: tool_used`. If `min_calls` is not specified for a `tool_used` check, the default is to just verify the tool was used at least once.
- **Source and mcp_server fields**: These apply only to `tool_used` and `tool_not_used`. Omit `tool` to match any invocation carrying the specified `mcp_server_name`. Group checks (`any_from_group`, `none_from_group`) do not support source or mcp_server filtering.
- **MCP server validation**: If `mcp_server` is specified, `source` must be set to `mcp`. Specifying `mcp_server` without `source: mcp` will cause a validation error.
- **Group matching**: Groups are resolved from the tool topology at evaluation time. If a group is not defined, the check will fail with a descriptive error.
- **Exception handling**: The `except:` list in group checks allows you to exclude specific tools from group validation (e.g., allow "bash" even though it's in the "dangerous_tools" group).

## Troubleshooting

- **Tool not found**: Verify the tool name matches exactly (case-sensitive) and is available in your evaluation config.
- **Group not found**: Check that the group name is defined in your tool topology and matches exactly.
- **Call count failing**: Ensure `min_calls` ≤ `max_calls` and that the actual invocation count is within bounds.
- **Multiple tools with same prefix**: Use the full tool name, not a prefix, for exact matching.

See [index.md](./index.md) for general grader concepts and [../configuration.md](../configuration.md) for config file structure.

## Reference

For a comprehensive example of all four tool check kinds in action, see [`criteria/language/test.yaml`](../../criteria/language/test.yaml), which demonstrates `tool_used`, `tool_not_used`, `any_from_group`, and `none_from_group` in a single test criteria file.
