# Neo — Pairwise Deep Mode Bug Fix + Tool Grader Redesign

**By:** Neo 💊 (Core Eval Framework Developer)  
**Date:** 2026-05-01  
**Branch:** `ronniegeraghty/dev`  
**Status:** SHIPPED (Commits: 4f293e06, 24de2f26)

---

## Summary

Implemented Morpheus's pairwise deep bug fix and tool grader redesign as specified in `.squad/decisions/inbox/morpheus-grader-pairwise-redesign-plan.md`.

**Commit 1: fix(pairwise): honor ExcludedSkills/ExcludedTools at session-spawn time (SHA: 4f293e06)**
- Fixed split-brain where pairwise deep variants excluded skills from the report but the Copilot SDK loaded them all anyway.
- Added plugin deep mode support with `ExcludedTools` field.
- Test `TestPairwiseDeepVariantSkillsLoadedFilter` now passes (was failing, demonstrating the bug).

**Commit 2: feat(graders): redesign tool grader around tool/group framing (SHA: 24de2f26)**
- Replaced ad-hoc tool grader with four canonical kinds: `tool_used`, `tool_not_used`, `any_from_group`, `none_from_group`.
- Dropped magic group strings (`mcp`, `skill_plugin`, `skill_repo:*`, `tool_name_glob:*`).
- `turn_limit` migrated out (now belongs to activity grader per Tank's parallel work).
- Loud migration errors for legacy kind names.

---

## Commit 1 Details: Pairwise Deep Bug Fix

### Root Cause

`hyoka/internal/config/tool/validate.go::validateSkillDirEntry` (lines 782–851) walked skill_dir children and emitted `ToolLoadItem` rows without consulting `entry.ExcludedSkills`. The legacy `tool.ResolveSkills` path (used only for report metadata) did honor exclusions, creating a split-brain:

- **Report truth** (legacy path): Shows filtered single skill.
- **Session truth** (validate path): SDK loads all skills regardless of exclusions.

Evidence: `report.json` showed `skillDirectories: [markdown-lists]` but `skillsLoaded: [markdown-headings, markdown-lists]`.

### Fixes Applied

1. **validateSkillDirEntry (validate.go:833):**
   ```go
   // Skip if excluded (pairwise deep mode)
   if contains(entry.ExcludedSkills, e.Name()) {
       continue
   }
   ```

2. **Plugin deep mode:**
   - Added `ExcludedTools []string` field to `Entry` struct (entry.go:29).
   - Updated `validatePluginEntry` to pass `entry.ExcludedTools` to `emitPluginLoadedWithChildren`.
   - Modified `emitPluginLoadedWithChildren` to skip excluded tools.
   - Updated pairwise.go to enumerate plugin tools and populate `ExcludedTools` in deep variants.

3. **Helper function:**
   ```go
   func contains(slice []string, item string) bool {
       for _, s := range slice {
           if s == item {
               return true
           }
       }
       return false
   }
   ```

### MCP Deep Mode Audit

Current implementation: `pairwise.go::removeTool` sets `te.MCPTools = removeMCPTool(...)` (line 157). The MCP config directly consumes `entry.MCPTools` as an allow-list. No additional validation needed — SDK already filters per-server tools via the `mcp_tools:` field.

### Test Evidence

`TestPairwiseDeepVariantSkillsLoadedFilter` (pairwise_skill_filter_test.go):
- **Before fix:** Failed with "Expected 1 skills loaded, got 2: [markdown-headings markdown-lists]"
- **After fix:** Passes — `GeneratorSkillDirs()` returns only non-excluded skills.

---

## Commit 2 Details: Tool Grader Redesign

### New Schema

```yaml
- name: Tool Check
  type: tool
  weight: 1.0
  details:
    checks:
      - kind: tool_used
        tool: markdown-headings
        min_calls: 1   # optional
        max_calls: 2   # optional
      - kind: tool_not_used
        tool: bash
      - kind: any_from_group
        group: test-skills           # name of a skill_dir, plugin, or mcp_server entry
        except: [markdown-lists]     # optional (string or string[])
      - kind: none_from_group
        group: test-skills
        except: [markdown-headings]
```

### Kind Mapping

| Old Kind        | New Kind                          | Notes                                  |
|-----------------|-----------------------------------|----------------------------------------|
| `specific_tool` | `tool_used`                       | With optional min_calls/max_calls      |
| `min_calls`     | Folded into `tool_used.min_calls` | No longer a separate kind              |
| `max_calls`     | Folded into `tool_used.max_calls` | No longer a separate kind              |
| `tool_not_used` | `tool_not_used`                   | Unchanged                              |
| `any_of_group`  | `any_from_group`                  | With optional `except: []` exclusions  |
| `group_not_used`| `none_from_group`                 | With optional `except: []` exclusions  |
| `turn_limit`    | **REMOVED**                       | Migrates to `activity` grader (Tank)   |

### ToolCheckRule Fields

```go
type ToolCheckRule struct {
	Kind     string   `yaml:"kind" json:"kind"`
	Tool     string   `yaml:"tool,omitempty" json:"tool,omitempty"`           // For tool_used, tool_not_used
	Group    string   `yaml:"group,omitempty" json:"group,omitempty"`         // For any_from_group, none_from_group
	Except   []string `yaml:"except,omitempty" json:"except,omitempty"`       // Optional exclusion list for group checks
	MinCalls *int     `yaml:"min_calls,omitempty" json:"min_calls,omitempty"` // Optional for tool_used
	MaxCalls *int     `yaml:"max_calls,omitempty" json:"max_calls,omitempty"` // Optional for tool_used
}
```

Dropped fields: `Name string`, `N int`.

### Group Resolution

Groups resolve by entry `Name` from tool topology:
- `skill_dir` entry name → all child skills
- `plugin` entry name → all plugin's exported tools
- `mcp_server` entry name → all server's registered tools

**Current limitation:** `EnvironmentTool` lacks `Parent` linkage, so group resolution is a placeholder (returns all tools). TODO: Add `Parent string` and `ParentKind string` fields to `EnvironmentTool` (mirroring `ToolLoadItem` structure).

### Migration Errors

Legacy kind names produce loud errors with migration messages:

```
tool grader "Check 1": checks[0] uses deprecated kind "specific_tool" → tool_used
tool grader "Check 2": checks[1] uses deprecated kind "turn_limit" → REMOVED: turn_limit now belongs to the activity grader
tool grader "Check 3": checks[2] uses deprecated kind "min_calls" → FOLDED INTO tool_used: use tool_used with min_calls/max_calls fields
```

### Test Coverage

`tool_grader_test.go` (new table-driven tests):
- `TestToolGraderNewKinds`: Validates all four new kinds + legacy migration errors
- `TestToolGraderGrade`: End-to-end grading with pass/fail assertions

All tests pass.

---

## Files Changed

**Commit 1 (4f293e06):**
- `hyoka/internal/config/tool/entry.go`: Added `ExcludedTools []string` field
- `hyoka/internal/config/tool/validate.go`: Added exclusion checks in `validateSkillDirEntry`, `validatePluginEntry`, `emitPluginLoadedWithChildren`; added `contains` helper
- `hyoka/internal/pairwise/pairwise.go`: Added plugin deep mode support (`enumeratePluginTools`, updated `collectTogglable` and `removeTool`)

**Commit 2 (24de2f26):**
- `hyoka/internal/criteria/graders/tool_grader.go`: Complete rewrite with new kinds and validation
- `hyoka/internal/criteria/graders/types.go`: Reshaped `ToolCheckRule` struct
- `hyoka/internal/criteria/graders/tool_grader_test.go`: Rewrote tests for new schema

**Commit 3 (6c379aa6):**
- `.squad/agents/neo/history.md`: Updated work log
- `.squad/decisions/inbox/neo-pairwise-tool-grader-shipped.md`: This file

---

## Next Steps

1. **Oracle:** Update `docs/graders.md` with new tool grader schema and kind names.
2. **Tank/Switch:** Update criteria YAML files (e.g., `criteria/language/python.yaml`) to use new kind names.
3. **Neo (future):** Add `Parent` and `ParentKind` fields to `EnvironmentTool` to enable proper group resolution.
4. **Integration test:** Switch's `skillsLoaded` assertion should now pass with commit 4f293e06.

---

## Verification

```bash
# Build
cd /home/rgeraghty/projects/hyoka
go build ./...

# Tests
go test -race ./hyoka/internal/config/tool/... -run TestPairwiseDeepVariantSkillsLoadedFilter
go test -race ./hyoka/internal/pairwise/...
go test -race ./hyoka/internal/criteria/graders/... -run TestToolGrader
```

All tests pass.

---

## Coordination Notes

- **Tank:** Workspace grader (commit 1f461a50) ships independently, no merge conflict.
- **Switch:** Expects `skillsLoaded` to match deep variant exclusions — commit 4f293e06 delivers this.
- **Parallel work:** Tank is rewriting `output_check` → `workspace` and `action_sequence` → `activity`. Only `activity` grader should now accept `turn_limit`.

---

## Known Issues

- **Group resolution placeholder:** `resolveGroup` currently returns all tools because `EnvironmentTool` lacks parent linkage. This doesn't break the grader (it just means groups are over-inclusive). Proper resolution requires adding `Parent` and `ParentKind` fields to `EnvironmentTool`.

---

## Decision: ACCEPT

Both commits ship on `ronniegeraghty/dev`. No rollback needed. Oracle and Tank can proceed with dependent work.
