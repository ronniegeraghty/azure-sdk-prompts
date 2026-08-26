# Foundry Python evaluation

Run `20260826-192543` evaluated four Foundry prompts with
`gpt-5.6-sol` in two configurations:

- Baseline: no Azure MCP and no Microsoft SDK skills.
- Azure tools: Azure MCP plus the `azure-sdk-python` plugin.

## Prompt checks

These checks come only from each prompt's `## Evaluation Criteria` and measure
scenario-specific SDK behavior.

| Prompt | Baseline | Azure MCP + SDK skills | Difference |
|--------|----------|-------------------------|------------|
| Basic agent lifecycle | 10/10 (100%) | 10/10 (100%) | 0 pp |
| Function tool | 12/12 (100%) | 0/12 (0%) | -100 pp |
| File search | 9/9 (100%) | 9/9 (100%) | 0 pp |
| Project resource inventory | 9/9 (100%) | 9/9 (100%) | 0 pp |
| **Aggregate** | **40/40 (100%)** | **28/40 (70%)** | **-30 pp** |

The Azure-tools function-tool run used all 19 turns for SDK research and wrote
no files. Its final response said it was still narrowing the official examples.
This is a generation failure rather than evidence that the SDK guidance was
incorrect.

## Generic Python checks

These checks come from `criteria/language/python.yaml` and are reported
separately from prompt correctness.

| Check type | Baseline | Azure MCP + SDK skills | Difference |
|------------|----------|-------------------------|------------|
| Generic Python | 15/20 (75%) | 14/20 (70%) | -5 pp |

The Azure-tools total is dominated by the function-tool run, which passed only
the async-applicability check because it generated no code. Across the other
three prompts, Azure tools passed 13/15 checks versus 11/15 for baseline.

## Tool and workspace checks

| Check | Baseline | Azure MCP + SDK skills |
|-------|----------|-------------------------|
| Workspace grader | 0/4 | 0/4 |
| Azure MCP tool grader | 0/4 | 0/4 |
| Observed Azure MCP use | 0 calls across 0/4 runs | 22 calls across 4/4 runs |

The workspace grader failed even when generated Python files were present. The
tool grader also failed every Azure-tools run despite the timeline recording
Azure MCP calls. These grader results don't reflect the generated applications
or actual tool usage and should not be folded into the SDK comparison.

## Result links

| Prompt | Baseline | Azure MCP + SDK skills |
|--------|----------|-------------------------|
| Basic agent lifecycle | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/with-azure-tools/report.md) |
| Function tool | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-function-tool/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-function-tool/python-azure-tools/with-azure-tools/report.md) |
| File search | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-file-search/python-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/python/agents/ai-agents-dp-python-file-search/python-azure-tools/with-azure-tools/report.md) |
| Project resource inventory | [Report](results/ai-projects/data-plane/python/projects/ai-projects-dp-python-project-resource-inventory/python-azure-tools/baseline/report.md) | [Report](results/ai-projects/data-plane/python/projects/ai-projects-dp-python-project-resource-inventory/python-azure-tools/with-azure-tools/report.md) |
