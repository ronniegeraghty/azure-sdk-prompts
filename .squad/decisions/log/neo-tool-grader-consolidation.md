## Status: SHIPPED (2026-05-01)
# Tool Grader Consolidation

**Status**: Implemented  
**Date**: 2026-05-01  
**Commit**: dcc3dad3

## Decision

Consolidate behavior_grader, tool_constraint_grader, and tool_usage_grader into a single canonical `tool` grader. Legacy grader types remain functional but emit deprecation warnings at load time.

## New Tool Grader API

The unified tool grader (`kind: tool`) supports multiple check types via a `checks:` array:

```yaml
graders:
  - name: Tool Checks
    type: tool
    details:
      checks:
        - kind: specific_tool
          name: bash
        - kind: tool_not_used
          name: dangerous_tool
        - kind: any_of_group
          group: mcp
        - kind: turn_limit
          n: 10
        - kind: min_calls
          name: skill
          n: 2
        - kind: max_calls
          name: edit
          n: 50
```

### Supported Check Kinds

- `specific_tool`: Verify a named tool was used at least once
- `tool_not_used`: Verify a named tool was NOT used
- `any_of_group`: At least one tool from a group was used
- `group_not_used`: No tool from a group was used
- `turn_limit`: Turn count ≤ N
- `min_calls`: Tool called ≥ N times
- `max_calls`: Tool called ≤ N times

### Group Resolution

- `mcp`: All MCP server tools
- `skill_plugin`: All skill plugins
- `skill_repo:{repo}`: Skills from a specific GitHub repo
- `tool_name_glob:{pattern}`: Tools matching a glob pattern

Each check emits its own GraderCheck with a stable ID (`check_1`, `check_2`, ...). The grader's overall pass is the AND of all checks.

## Deprecation Warnings

At load time, the following warnings are emitted:

- `kind: behavior` → "use 'tool' kind instead"
- `kind: tool_constraint` → "use 'tool' kind instead"
- `kind: tool_usage` → "use 'tool' kind instead"
- `kind: file` → "use 'output_check' with require_files instead"

Existing YAML configurations continue to work unchanged. No migration is required immediately.

## Deprecation Horizon

The legacy grader types will be removed in the next major release (v1.0). Users should migrate to the canonical types:

- `behavior` → `tool`
- `tool_constraint` → `tool`
- `tool_usage` → `tool`
- `file` → `output_check`

## Implementation Notes

- Tool grader validates all checks at construction time
- Each check is independent and evaluated separately
- Group resolution happens at evaluation time using EnvironmentTools
- The grader emits one GraderCheck per configured check
- No Extras payload is emitted (all data is in Checks)

## Testing

Comprehensive unit tests cover:
- All check kinds (positive and negative cases)
- Group resolution (mcp, skill_plugin, skill_repo, tool_name_glob)
- Validation errors (missing fields, invalid kinds, negative thresholds)
- Grader interface compliance

## Related Work

- Tank's C10: Renamed GraderPoint → GraderCheck
- Morpheus's grader overhaul plan: Proposed consolidation strategy
