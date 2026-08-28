# JavaScript/TypeScript three-way comparison

The report includes **13 complete prompt triplets** and **39 valid evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 86/101 | 85.1% | - |
| Azure Skill + MCP | 88/101 | 87.1% | +2.0 pp |
| Azure Skill + MCP + Microsoft skill | 87/101 | 86.1% | +1.0 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 3 | 1 | 9 |
| Microsoft skill vs baseline | 2 | 1 | 10 |
| Microsoft skill vs Azure Skill + MCP | 2 | 3 | 8 |

Adding Azure Skill + MCP changed the prompt-check rate by **+2.0 pp**. Adding the Microsoft language skill changed it by **-1.0 pp** relative to Azure Skill + MCP.

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 88/130 | 67.7% | - |
| Azure Skill + MCP | 88/130 | 67.7% | 0.0 pp |
| Azure Skill + MCP + Microsoft skill | 94/130 | 72.3% | +4.6 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Execution exclusions

| Prompt ID | Timed-out config | Attempts | Reason |
|---|---|---:|---|
| `storage-dp-js-ts-encrypted-uploader` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | 3 | Copilot SDK `session.idle` timeout |

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| `app-configuration-dp-js-ts-crud` | 6/8 | 7/8 | 6/8 |
| `cosmos-db-dp-js-ts-crud` | 5/7 | 5/7 | 6/7 |
| `event-hubs-dp-js-ts-streaming` | 7/8 | 8/8 | 8/8 |
| `identity-dp-js-ts-default-credential` | 4/5 | 4/5 | 4/5 |
| `identity-dp-js-ts-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-js-ts-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-js-ts-crud` | 4/5 | 4/5 | 4/5 |
| `key-vault-dp-js-ts-secret-config` | 12/13 | 12/13 | 12/13 |
| `resource-manager-mp-js-ts-rg-crud` | 8/8 | 8/8 | 8/8 |
| `service-bus-dp-js-ts-crud` | 7/8 | 7/8 | 7/8 |
| `storage-dp-js-ts-blob-manager` | 9/12 | 8/12 | 9/12 |
| `storage-dp-js-ts-crud` | 7/8 | 8/8 | 7/8 |
| `storage-mp-js-ts-account-mgmt` | 7/8 | 7/8 | 6/8 |
