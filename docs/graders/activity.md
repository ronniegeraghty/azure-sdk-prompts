# Activity Grader

The `activity` grader validates session behavior and action patterns during code generation. It checks turn counts, tool invocation sequences, action types, and session termination status against configured expectations. The grader passes only if **all** checks pass (boolean semantics).

This canonical grader replaces the legacy `action_sequence` and `behavior` graders.

## When to Use

- Enforce maximum turn/action count limits for performance
- Verify specific action types occur (e.g., "must contain file_write action")
- Exclude undesirable action patterns (e.g., "must not generate errors")
- Validate tool invocation sequences (e.g., "bash then file_search")
- Verify session termination reason (e.g., "must complete normally, not timeout")
- Check for incomplete reasoning or truncation

## Configuration

The `activity` grader uses a top-level `checks:` array. Each check validates session activity:

```yaml
graders:
  - name: Session Activity
    type: activity
    weight: 0.3
    checks:
      - kind: turn_limit
        max: 25
      - kind: contains_action
        type: message
      - kind: terminated_by
        equals: completed
```

### `checks` Schema

Seven check kinds are supported:

#### 1. `turn_limit`

Maximum turn number must not exceed configured value.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"turn_limit"` |
| `max` | int | yes | Maximum allowed turn number (1-indexed) |

#### 2. `action_count`

Total action count must fall within optional min/max bounds.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"action_count"` |
| `min` | int | no | Minimum action count (inclusive) |
| `max` | int | no | Maximum action count (inclusive) |

#### 3. `tool_call_count`

Tool call count must fall within optional min/max bounds.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"tool_call_count"` |
| `min` | int | no | Minimum tool call count (inclusive) |
| `max` | int | no | Maximum tool call count (inclusive) |

#### 4. `contains_subsequence`

Action log must contain an ordered subsequence of named tools (actions not in the list are ignored).

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"contains_subsequence"` |
| `tools` | []string | yes | Ordered list of tool names to find in sequence |

#### 5. `contains_action`

At least one action matching the specified criteria must exist. Optional count bounds (default: min=1).

Filters action events by:
- `type`: Action event type (tool_call, file_read, file_write, bash, mcp_call, skill, message, reasoning, intent, warning, error, file_change, truncation, compaction, turn_start, turn_end, abort, other)
- `tool`: Tool/skill name (exact match)
- `contains`: Substring that must appear in event text (Output for messages/reasoning/intent/tool output; Error for error events; Path for file_change)
- `excludes`: Substring that must NOT appear in event text

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"contains_action"` |
| `type` | string | no | Action event type filter (empty = any type) |
| `tool` | string | no | Tool/skill name filter (empty = any tool) |
| `contains` | string | no | Substring required in event text |
| `excludes` | string | no | Substring forbidden in event text |
| `min` | int | no | Minimum matching actions (default: 1) |
| `max` | int | no | Maximum matching actions (default: unbounded) |

#### 6. `excludes_action`

No actions matching the specified criteria must exist (count must be 0).

Same filters as `contains_action`: type, tool, contains, excludes.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"excludes_action"` |
| `type` | string | no | Action event type filter |
| `tool` | string | no | Tool/skill name filter |
| `contains` | string | no | Substring forbidden in event text |
| `excludes` | string | no | Substring required to be absent in event text |

#### 7. `terminated_by`

Session termination reason must match expectation. Either an exact match (`equals`) or must not match any in a list (`not_in`).

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `kind` | string | yes | Must be `"terminated_by"` |
| `equals` | string | no | Exact termination reason required (one of: completed, max_actions, max_turns, guardrail, error) |
| `not_in` | []string | no | List of forbidden termination reasons |

> Note: `equals` and `not_in` are mutually exclusive. Specify exactly one.

## Valid Action Types

For `type` filters: `tool_call`, `file_read`, `file_write`, `bash`, `mcp_call`, `skill`, `message`, `reasoning`, `intent`, `warning`, `error`, `file_change`, `truncation`, `compaction`, `turn_start`, `turn_end`, `abort`, `other`.

## Valid Termination Reasons

- `completed` — session finished normally
- `max_actions` — hit max action limit
- `max_turns` — hit max turn limit
- `guardrail` — triggered guardrail (e.g., file limit, output size)
- `error` — session encountered error

## Examples

### Basic Limits

```yaml
graders:
  - name: Session Limits
    type: activity
    weight: 0.3
    checks:
      - kind: turn_limit
        max: 25
      - kind: action_count
        min: 5
        max: 100
      - kind: tool_call_count
        max: 20
```

### Action Sequence

```yaml
graders:
  - name: Tool Sequence
    type: activity
    weight: 0.2
    checks:
      # Bash must be called before file_search
      - kind: contains_subsequence
        tools: [bash, file_search]
```

### Verify Action Presence

```yaml
graders:
  - name: Must Generate Files
    type: activity
    weight: 0.3
    checks:
      # Must contain at least one file_write action
      - kind: contains_action
        type: file_write
      # Must contain a message (from the LLM)
      - kind: contains_action
        type: message
```

### Exclude Undesirable Patterns

```yaml
graders:
  - name: No Errors
    type: activity
    weight: 0.2
    checks:
      # Session must complete normally
      - kind: terminated_by
        equals: completed
      # Must not contain any error events
      - kind: excludes_action
        type: error
```

### Verify Content

```yaml
graders:
  - name: Content Validation
    type: activity
    weight: 0.4
    checks:
      # A message containing "TODO" must NOT appear (code incomplete)
      - kind: excludes_action
        type: message
        contains: "TODO"
      # At least one bash invocation must succeed (no errors)
      - kind: contains_action
        type: bash
        excludes: "error"
```

### Comprehensive Session Check

```yaml
graders:
  - name: Comprehensive Session Behavior
    type: activity
    weight: 0.5
    checks:
      - kind: turn_limit
        max: 25
      - kind: action_count
        min: 10
        max: 200
      - kind: tool_call_count
        min: 1
        max: 30
      - kind: contains_subsequence
        tools: [bash, file_search, bash]
      - kind: contains_action
        type: file_write
      - kind: excludes_action
        type: error
      - kind: terminated_by
        equals: completed
```

## Result Structure

Each activity grader produces:
- **Pass/Fail**: Binary result (true only if ALL checks pass)
- **Check results**: Individual pass/fail for each configured check
- **Summary**: Action count, turn count, termination reason
- **Score**: 1.0 if all checks pass, 0.0 if any check fails (boolean grader)

Results visible in evaluation reports under `grader_results`.

## Data Visible to Grader

Activity graders can access:
- **ActionLog**: Ordered list of all agent actions during the session
  - Action type (tool_call, file_read, file_write, bash, message, etc.)
  - Tool/skill name (if applicable)
  - Event text (Output, Error, Path, etc.)
  - Timestamp
- **ActionsSummary**: Aggregated counts
  - TotalActions: Total number of actions
  - ToolCalls: Total tool invocations
  - TurnsCompleted: Highest turn number reached
- **TerminatedBy**: Session termination reason

## Notes

- **All-or-nothing**: Grader passes only if **every** check passes. A single failed check fails the grader.
- **Action types**: Type values match ActionEvent.Type from the evaluation engine (e.g., `tool_call`, `file_write`, `error`).
- **Substring matching**: `contains` and `excludes` are case-sensitive substring matches, not regex patterns.
- **Sequence matching**: `contains_subsequence` finds tools in order but allows other actions in between (e.g., `[bash, file_search]` matches `bash → read → file_search`).
- **Default count**: `contains_action` defaults to `min=1` if not specified; use `max` to enforce upper limits.
- **Filtering**: All filters (`type`, `tool`, `contains`, `excludes`) are optional and combine with AND logic (all must match).

## Troubleshooting

- **No actions match**: Verify action type is valid and filters are correct. Check logs to see actual action types present.
- **Sequence never found**: Remember that other actions can appear between sequence items. Check tool names match exactly.
- **Termination check fails**: Verify the session actually completed with expected reason (look at evaluation logs).
- **Count bounds reversed**: Ensure `min` ≤ `max` and that actual counts are within bounds.

See [index.md](./index.md) for general grader concepts and [../configuration.md](../configuration.md) for config file structure.
