# Java three-way comparison

The report includes **19 complete prompt triplets** and **57 valid evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 155/167 | 92.8% | - |
| Azure Skill + MCP | 151/167 | 90.4% | -2.4 pp |
| Azure Skill + MCP + Microsoft skill | 149/167 | 89.2% | -3.6 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 1 | 3 | 15 |
| Microsoft skill vs baseline | 3 | 2 | 14 |
| Microsoft skill vs Azure Skill + MCP | 5 | 2 | 12 |

Adding Azure Skill + MCP changed the prompt-check rate by **-2.4 pp**. Adding the Microsoft language skill changed it by **-1.2 pp** relative to Azure Skill + MCP.

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 187/228 | 82% | - |
| Azure Skill + MCP | 186/228 | 81.6% | -0.4 pp |
| Azure Skill + MCP + Microsoft skill | 177/228 | 77.6% | -4.4 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| `app-configuration-dp-java-crud` | 7/7 | 7/7 | 7/7 |
| `app-configuration-dp-java-feature-flags` | 8/9 | 7/9 | 9/9 |
| `cosmos-db-dp-java-crud` | 6/7 | 6/7 | 7/7 |
| `cosmos-db-dp-java-todo-repository` | 11/14 | 12/14 | 11/14 |
| `event-hubs-dp-java-streaming` | 7/7 | 7/7 | 7/7 |
| `identity-dp-java-credential-chain` | 14/14 | 14/14 | 14/14 |
| `identity-dp-java-default-credential` | 5/5 | 5/5 | 5/5 |
| `identity-dp-java-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-java-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-java-crud` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-java-secret-config` | 10/10 | 10/10 | 10/10 |
| `resource-manager-mp-java-rg-crud` | 7/7 | 7/7 | 7/7 |
| `service-bus-dp-java-crud` | 7/7 | 7/7 | 7/7 |
| `service-bus-dp-java-order-processor` | 9/12 | 9/12 | 9/12 |
| `storage-dp-java-blob-event-notifier` | 9/10 | 6/10 | 8/10 |
| `storage-dp-java-blob-manager` | 6/7 | 6/7 | 7/7 |
| `storage-dp-java-crud` | 6/7 | 5/7 | 6/7 |
| `storage-dp-java-encrypted-uploader` | 20/20 | 20/20 | 20/20 |
| `storage-mp-java-account-mgmt` | 8/8 | 8/8 | 0/8 |
