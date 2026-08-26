# Foundry Java evaluation

Run `20260826-212414` evaluated four Foundry prompts with
`gpt-5.6-sol` in two configurations:

- Baseline: no Azure MCP and no Microsoft SDK skills.
- Azure tools: Azure MCP plus the `azure-sdk-java` plugin.

## Prompt checks

These checks come only from each prompt's `## Evaluation Criteria` and measure
scenario-specific SDK behavior.

| Prompt | Baseline | Azure MCP + SDK skills | Difference |
|--------|----------|-------------------------|------------|
| Basic agent lifecycle | 11/11 (100%) | 11/11 (100%) | 0 pp |
| Function tool | 12/12 (100%) | 11/12 (91.7%) | -8.3 pp |
| File search | 3/9 (33.3%) | 9/9 (100%) | +66.7 pp |
| Project resource inventory | 9/9 (100%) | 9/9 (100%) | 0 pp |
| **Aggregate** | **35/41 (85.4%)** | **40/41 (97.6%)** | **+12.2 pp** |

The baseline file-search session wrote no files and passed only the three
guardrail checks that prohibit shortcuts. Its Azure-tools counterpart reached
the 50-action limit but retained three files and passed all nine SDK checks.

The Azure-tools function-tool application missed the required thread, exact
user message, and run creation sequence. Baseline passed all 12 checks for that
scenario.

## Generic Java checks

These checks come from `criteria/language/java.yaml` and are reported separately
from prompt correctness.

| Check type | Baseline | Azure MCP + SDK skills | Difference |
|------------|----------|-------------------------|------------|
| Generic Java | 29/48 (60.4%) | 32/48 (66.7%) | +6.3 pp |

## Tool and workspace checks

| Check | Baseline | Azure MCP + SDK skills |
|-------|----------|-------------------------|
| Workspace grader | 0/4 | 0/4 |
| Azure MCP tool grader | 0/4 | 0/4 |
| Observed Azure MCP use | 0 calls across 0/4 runs | 11 calls across 4/4 runs |

The workspace grader failed even for evaluations with three generated files.
The tool grader also failed every Azure-tools run despite recorded Azure MCP
calls. These grader results don't reflect the generated applications or actual
tool usage and aren't folded into the SDK comparison.

## Result links

| Prompt | Baseline | Azure MCP + SDK skills |
|--------|----------|-------------------------|
| Basic agent lifecycle | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-basic-agent-lifecycle/java-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-basic-agent-lifecycle/java-azure-tools/with-azure-tools/report.md) |
| Function tool | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-function-tool/java-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-function-tool/java-azure-tools/with-azure-tools/report.md) |
| File search | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-file-search/java-azure-tools/baseline/report.md) | [Report](results/ai-agents/data-plane/java/agents/ai-agents-dp-java-file-search/java-azure-tools/with-azure-tools/report.md) |
| Project resource inventory | [Report](results/ai-projects/data-plane/java/projects/ai-projects-dp-java-project-resource-inventory/java-azure-tools/baseline/report.md) | [Report](results/ai-projects/data-plane/java/projects/ai-projects-dp-java-project-resource-inventory/java-azure-tools/with-azure-tools/report.md) |
