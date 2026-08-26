# Foundry .NET evaluation

Run `20260826-223641` evaluated four Foundry prompts with
`gpt-5.6-sol` in two configurations:

- Baseline: no Azure MCP and no Microsoft SDK skills.
- Azure tools: Azure MCP plus the `azure-sdk-dotnet` plugin.

## Prompt checks

These checks come only from each prompt's `## Evaluation Criteria` and measure
scenario-specific SDK behavior.

| Prompt | Baseline | Azure MCP + SDK skills | Difference |
|--------|----------|-------------------------|------------|
| Basic agent lifecycle | 11/11 (100%) | 11/11 (100%) | 0 pp |
| Function tool | 12/12 (100%) | 12/12 (100%) | 0 pp |
| File search | 9/9 (100%) | 9/9 (100%) | 0 pp |
| Project resource inventory | 9/9 (100%) | 9/9 (100%) | 0 pp |
| **Aggregate** | **41/41 (100%)** | **41/41 (100%)** | **0 pp** |

Both configurations satisfied every scenario-specific SDK check. All eight
evaluations retained three generated files and passed overall.

## Generic .NET checks

No `.NET` language criteria matched these prompts, so the run produced no
generic language scores. This differs from the Python, JavaScript/TypeScript,
and Java runs and prevents a generic-language comparison.

## Tool and workspace checks

| Check | Baseline | Azure MCP + SDK skills |
|-------|----------|-------------------------|
| Workspace grader | Not configured | Not configured |
| Azure MCP tool grader | Not configured | Not configured |
| Observed Azure MCP use | 0 calls across 0/4 runs | 12 calls across 4/4 runs |

Unlike the other language runs, no workspace or tool graders were loaded for
.NET. Actual Azure MCP use is reported directly from the action timelines.

## Result links

| Prompt | Baseline | Azure MCP + SDK skills |
|--------|----------|-------------------------|
| Basic agent lifecycle | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-basic-agent-lifecycle/dotnet-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-basic-agent-lifecycle/dotnet-azure-tools/with-azure-tools/report.md) |
| Function tool | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-function-tool/dotnet-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-function-tool/dotnet-azure-tools/with-azure-tools/report.md) |
| File search | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-file-search/dotnet-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/dotnet/agents/ai-agents-dp-dotnet-file-search/dotnet-azure-tools/with-azure-tools/report.md) |
| Project resource inventory | [Report](results/ai-projects/data-plane/dotnet/projects/ai-projects-dp-dotnet-project-resource-inventory/dotnet-azure-tools/baseline/report.md) | [Report](results/ai-projects/data-plane/dotnet/projects/ai-projects-dp-dotnet-project-resource-inventory/dotnet-azure-tools/with-azure-tools/report.md) |
