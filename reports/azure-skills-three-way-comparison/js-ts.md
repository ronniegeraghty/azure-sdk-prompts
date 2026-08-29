# JavaScript/TypeScript three-way comparison

The report includes **13 complete prompt triplets** and **39 selected evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 103/118 | 87.3% | - |
| Azure Skill + MCP | 102/118 | 86.4% | -0.9 pp |
| Azure Skill + MCP + Microsoft Skills | 104/118 | 88.1% | +0.8 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 0 | 1 | 12 |
| Azure Skill + MCP + Microsoft Skills vs baseline | 1 | 0 | 12 |

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 88/130 | 67.7% | - |
| Azure Skill + MCP | 92/130 | 70.8% | +3.1 pp |
| Azure Skill + MCP + Microsoft Skills | 91/130 | 70.0% | +2.3 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Excluded prompt triplets

| Prompt ID | Affected arm | Reason |
|---|---|---|
| `event-hubs-dp-js-ts-streaming` | `js-ts-azure-skills/azure-skill-mcp` | Both the primary and retry generated files but timed out waiting for session.idle. |

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| `app-configuration-dp-js-ts-crud` | 6/8 | 6/8 | 6/8 |
| `cosmos-db-dp-js-ts-crud` | 6/7 | 6/7 | 6/7 |
| `identity-dp-js-ts-default-credential` | 4/5 | 4/5 | 4/5 |
| `identity-dp-js-ts-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-js-ts-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-js-ts-crud` | 4/5 | 4/5 | 4/5 |
| `key-vault-dp-js-ts-secret-config` | 12/13 | 12/13 | 12/13 |
| `resource-manager-mp-js-ts-rg-crud` | 8/8 | 8/8 | 8/8 |
| `service-bus-dp-js-ts-crud` | 7/8 | 7/8 | 7/8 |
| `storage-dp-js-ts-blob-manager` | 9/12 | 9/12 | 10/12 |
| `storage-dp-js-ts-crud` | 8/8 | 7/8 | 8/8 |
| `storage-dp-js-ts-encrypted-uploader` | 23/25 | 23/25 | 23/25 |
| `storage-mp-js-ts-account-mgmt` | 6/8 | 6/8 | 6/8 |
