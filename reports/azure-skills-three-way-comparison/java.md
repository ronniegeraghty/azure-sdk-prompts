# Java three-way comparison

The report includes **18 complete prompt triplets** and **54 selected evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 140/159 | 88.1% | - |
| Azure Skill + MCP | 146/159 | 91.8% | +3.7 pp |
| Azure Skill + MCP + Microsoft Skills | 149/159 | 93.7% | +5.6 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 2 | 3 | 13 |
| Azure Skill + MCP + Microsoft Skills vs baseline | 3 | 2 | 13 |

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 163/216 | 75.5% | - |
| Azure Skill + MCP | 179/216 | 82.9% | +7.4 pp |
| Azure Skill + MCP + Microsoft Skills | 177/216 | 81.9% | +6.4 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Excluded prompt triplets

| Prompt ID | Affected arm | Reason |
|---|---|---|
| `storage-mp-java-account-mgmt` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | Both the primary and retry completed without generated files. |

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| `app-configuration-dp-java-crud` | 7/7 | 7/7 | 7/7 |
| `app-configuration-dp-java-feature-flags` | 9/9 | 9/9 | 9/9 |
| `cosmos-db-dp-java-crud` | 6/7 | 6/7 | 7/7 |
| `cosmos-db-dp-java-todo-repository` | 11/14 | 11/14 | 11/14 |
| `event-hubs-dp-java-streaming` | 7/7 | 7/7 | 7/7 |
| `identity-dp-java-credential-chain` | 14/14 | 14/14 | 14/14 |
| `identity-dp-java-default-credential` | 5/5 | 4/5 | 5/5 |
| `identity-dp-java-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-java-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-java-crud` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-java-secret-config` | 10/10 | 10/10 | 10/10 |
| `resource-manager-mp-java-rg-crud` | 7/7 | 7/7 | 7/7 |
| `service-bus-dp-java-crud` | 7/7 | 6/7 | 6/7 |
| `service-bus-dp-java-order-processor` | 10/12 | 9/12 | 9/12 |
| `storage-dp-java-blob-event-notifier` | 6/10 | 9/10 | 9/10 |
| `storage-dp-java-blob-manager` | 0/7 | 6/7 | 7/7 |
| `storage-dp-java-crud` | 6/7 | 6/7 | 6/7 |
| `storage-dp-java-encrypted-uploader` | 20/20 | 20/20 | 20/20 |
