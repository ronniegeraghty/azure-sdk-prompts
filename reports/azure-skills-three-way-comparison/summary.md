# Azure skills three-way comparison

Prompt checks are the primary task-correctness measure. Language checks are supplemental. Workspace and tool/MCP checks are excluded from scored aggregates.

## Prompt checks

| Language | Complete triplets | Baseline | Azure Skill + MCP | Difference vs baseline | Azure Skill + MCP + Microsoft Skills | Difference vs baseline |
|---|---:|---:|---:|---:|---:|---:|
| [Python](./python.md) | 19 | 157/172 (91.3%) | 156/172 (90.7%) | -0.6 pp | 163/172 (94.8%) | +3.5 pp |
| [JavaScript/TypeScript](./js-ts.md) | 13 | 86/101 (85.1%) | 88/101 (87.1%) | +2.0 pp | 87/101 (86.1%) | +1.0 pp |
| [Java](./java.md) | 19 | 155/167 (92.8%) | 151/167 (90.4%) | -2.4 pp | 149/167 (89.2%) | -3.6 pp |
| [.NET](./dotnet.md) | 20 | 110/135 (81.5%) | 108/135 (80%) | -1.5 pp | 104/135 (77%) | -4.5 pp |
| **Informational rollup** | **71** | **508/575 (88.3%)** | **503/575 (87.5%)** | **-0.8 pp** | **503/575 (87.5%)** | **-0.8 pp** |

The informational rollup combines equivalent prompt checks only. It is not a language ranking.

## Language checks

| Language | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Python | 78/95 (82.1%) | 78/95 (82.1%) | 83/95 (87.4%) |
| JavaScript/TypeScript | 88/130 (67.7%) | 88/130 (67.7%) | 94/130 (72.3%) |
| Java | 187/228 (82.0%) | 186/228 (81.6%) | 177/228 (77.6%) |
| .NET | Not configured | Not configured | Not configured |

## Interpretation limits

- This is a single trial per prompt/config combination and is subject to model variance.
- Cross-language scores are not directly comparable because prompt inventories and criteria differ.
- Loaded skills might not be invoked for every prompt.
- MCP invocation and workspace checks are diagnostics, not generated-code correctness checks.
- One JS/TS prompt triplet is excluded because the Azure Skill + MCP + Microsoft Skills arm repeatedly hit a Copilot SDK `session.idle` timeout.
