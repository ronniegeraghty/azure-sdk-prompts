# Python MCP grader verification

Runs `20260827-100631` and `20260827-103815` repeated the original four
Python Foundry cases after applying #652 and #655.

## MCP server check

The criterion now checks for any tool call whose `mcp_server_name` is `azure`.

| Prompt | Baseline | Azure MCP and SDK skills |
|--------|----------|----------------------------|
| Basic agent lifecycle | 0 calls, expected failure | 5 calls, pass |
| File search | 0 calls, expected failure | 5 calls, pass |
| Function tool | 0 calls, expected failure | 7 calls, pass |
| Project resource inventory | 0 calls, expected failure | 9 calls, pass |
| **Aggregate** | **0 calls, 0/4 passed** | **26 calls, 4/4 passed** |

This confirms that the grader uses the recorded MCP server name instead of
looking for a nonexistent tool named `azure`. MCP-tagged tool execution events
also contribute to the action timeline's MCP call count.

## Tool durations

All 207 completed tool events across the eight evaluations recorded a positive
duration. This confirms the SDK v1 start/completion tracker restored per-tool
durations.

## Remaining workspace issue

The Python workspace check still failed all eight evaluations because
`name: "*.py"` is treated as a literal path rather than a glob. This issue is
independent of the MCP grader and SDK v1 fixes.

## Reports

| Prompt | Baseline | Azure MCP and SDK skills |
|--------|----------|----------------------------|
| Basic agent lifecycle | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/with-azure-tools/report.md) |
| File search | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-file-search/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-file-search/python-azure-tools/with-azure-tools/report.md) |
| Function tool | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-function-tool/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-function-tool/python-azure-tools/with-azure-tools/report.md) |
| Project resource inventory | [Baseline](../20260827-103815/results/ai-projects/data-plane/python/projects/ai-projects-dp-python-project-resource-inventory/python-azure-tools/baseline/report.md) | [Azure tools](../20260827-103815/results/ai-projects/data-plane/python/projects/ai-projects-dp-python-project-resource-inventory/python-azure-tools/with-azure-tools/report.md) |
