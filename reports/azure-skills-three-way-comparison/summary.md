# Azure skills three-way comparison

Prompt checks are the primary task-correctness measure. Language checks are supplemental. Workspace and tool/MCP checks are excluded from scored aggregates.

## Prompt checks

| Language | Complete triplets | Baseline | Azure Skill + MCP | Difference vs baseline | Azure Skill + MCP + Microsoft Skills | Difference vs baseline |
|---|---:|---:|---:|---:|---:|---:|
| [Python](./python.md) | 19 | 150/172 (87.2%) | 156/172 (90.7%) | +3.5 pp | 159/172 (92.4%) | +5.2 pp |
| [JavaScript/TypeScript](./js-ts.md) | 13 | 103/118 (87.3%) | 102/118 (86.4%) | -0.9 pp | 104/118 (88.1%) | +0.8 pp |
| [Java](./java.md) | 18 | 140/159 (88.1%) | 146/159 (91.8%) | +3.7 pp | 149/159 (93.7%) | +5.6 pp |
| [.NET](./dotnet.md) | 19 | 109/129 (84.5%) | 101/129 (78.3%) | -6.2 pp | 109/129 (84.5%) | 0.0 pp |
| **Informational rollup** | **69** | **502/578 (86.9%)** | **505/578 (87.4%)** | **+0.5 pp** | **521/578 (90.1%)** | **+3.2 pp** |

The informational rollup combines equivalent prompt checks only. It is not a language ranking.

## Language checks

| Language | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Python | 77/95 (81.1%) | 84/95 (88.4%) | 90/95 (94.7%) |
| JavaScript/TypeScript | 88/130 (67.7%) | 92/130 (70.8%) | 91/130 (70.0%) |
| Java | 163/216 (75.5%) | 179/216 (82.9%) | 177/216 (81.9%) |
| .NET | Not configured | Not configured | Not configured |

## Runtime health

- Primary suites: 216/216 evaluations and 72/72 complete raw triplets.
- Azure MCP: 697/697 primary-suite calls and 29/29 controlled-retry calls succeeded; no MCP tool timeout occurred.
- Controlled retries: 7 completed; 4 healthy retry results selected.
- Final comparison set: 69 complete triplets and 207 selected evaluations.

## Excluded prompt triplets

| Language | Prompt ID | Affected arm | Reason |
|---|---|---|---|
| JavaScript/TypeScript | `event-hubs-dp-js-ts-streaming` | `js-ts-azure-skills/azure-skill-mcp` | Both the primary and retry generated files but timed out waiting for session.idle. |
| Java | `storage-mp-java-account-mgmt` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | Both the primary and retry completed without generated files. |
| .NET | `identity-dp-dotnet-managed-identity` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | The primary timed out with no output; the retry completed without a timeout but still produced no generated files. |

## Interpretation limits

- This is a single trial per prompt/config combination and is subject to model variance.
- Cross-language scores are not directly comparable because prompt inventories and criteria differ.
- Loaded skills might not be invoked for every prompt.
- MCP invocation and workspace checks are diagnostics, not generated-code correctness checks.
- Complete prompt triplets are excluded when any arm remains unhealthy after its controlled retry.
