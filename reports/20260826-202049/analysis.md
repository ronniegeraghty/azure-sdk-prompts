# Foundry JavaScript/TypeScript evaluation

Run `20260826-202049` evaluated four Foundry prompts with
`gpt-5.6-sol` in two configurations:

- Baseline: no Azure MCP and no Microsoft SDK skills.
- Azure tools: Azure MCP plus the `azure-sdk-js` plugin.

## Prompt checks

These checks come only from each prompt's `## Evaluation Criteria` and measure
scenario-specific SDK behavior.

| Prompt | Baseline | Azure MCP + SDK skills | Difference |
|--------|----------|-------------------------|------------|
| Basic agent lifecycle | 10/10 (100%) | 10/10 (100%) | 0 pp |
| Function tool | 12/12 (100%) | 12/12 (100%) | 0 pp |
| File search | 9/9 (100%) | 9/9 (100%) | 0 pp |
| Project resource inventory | 9/9 (100%) | 9/9 (100%) | 0 pp |
| **Aggregate** | **40/40 (100%)** | **40/40 (100%)** | **0 pp** |

Both configurations satisfied every scenario-specific SDK check. The baseline
function-tool session reported a timeout while waiting for `session.idle`, but
its generated application was retained and passed all 12 prompt checks.

## Generic JavaScript/TypeScript checks

These checks come from `criteria/language/js-ts.yaml` and are reported
separately from prompt correctness.

| Check type | Baseline | Azure MCP + SDK skills | Difference |
|------------|----------|-------------------------|------------|
| Generic JavaScript/TypeScript | 24/40 (60%) | 26/40 (65%) | +5 pp |

## Tool and workspace checks

| Check | Baseline | Azure MCP + SDK skills |
|-------|----------|-------------------------|
| Workspace grader | 0/4 | 0/4 |
| Azure MCP tool grader | 0/4 | 0/4 |
| Observed Azure MCP use | 0 calls across 0/4 runs | 24 calls across 4/4 runs |

The workspace grader failed even though every evaluation retained four or five
generated files. The tool grader also failed every Azure-tools run despite the
timeline recording Azure MCP calls. These grader results don't reflect the
generated applications or actual tool usage and aren't folded into the SDK
comparison.

## Result links

| Prompt | Baseline | Azure MCP + SDK skills |
|--------|----------|-------------------------|
| Basic agent lifecycle | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-basic-agent-lifecycle/js-ts-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-basic-agent-lifecycle/js-ts-azure-tools/with-azure-tools/report.md) |
| Function tool | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-function-tool/js-ts-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-function-tool/js-ts-azure-tools/with-azure-tools/report.md) |
| File search | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-file-search/js-ts-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/js-ts/agents/ai-agents-dp-js-ts-file-search/js-ts-azure-tools/with-azure-tools/report.md) |
| Project resource inventory | [Report](results/ai-projects/data-plane/js-ts/projects/ai-projects-dp-js-ts-project-resource-inventory/js-ts-azure-tools/baseline/report.md) | [Report](results/ai-projects/data-plane/js-ts/projects/ai-projects-dp-js-ts-project-resource-inventory/js-ts-azure-tools/with-azure-tools/report.md) |
