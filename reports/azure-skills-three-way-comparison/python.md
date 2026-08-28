# Python three-way comparison

The report includes **19 complete prompt triplets** and **57 valid evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 157/172 | 91.3% | - |
| Azure Skill + MCP | 156/172 | 90.7% | -0.6 pp |
| Azure Skill + MCP + Microsoft skill | 163/172 | 94.8% | +3.5 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 3 | 4 | 12 |
| Microsoft skill vs baseline | 4 | 2 | 13 |
| Microsoft skill vs Azure Skill + MCP | 5 | 0 | 14 |

Adding Azure Skill + MCP changed the prompt-check rate by **-0.6 pp**. Adding the Microsoft language skill changed it by **+4.1 pp** relative to Azure Skill + MCP.

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 78/95 | 82.1% | - |
| Azure Skill + MCP | 78/95 | 82.1% | 0.0 pp |
| Azure Skill + MCP + Microsoft skill | 83/95 | 87.4% | +5.3 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| Workspace checks | 0/19 (0.0%) | 0/19 (0.0%) | 0/19 (0.0%) |
| Azure MCP usage checks | 0/19 (0.0%) | 19/19 (100.0%) | 19/19 (100.0%) |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| `app-configuration-dp-python-crud` | 6/6 | 6/6 | 6/6 |
| `app-configuration-dp-python-feature-flags` | 7/9 | 8/9 | 9/9 |
| `cosmos-db-dp-python-crud` | 6/6 | 5/6 | 5/6 |
| `cosmos-db-dp-python-todo-repository` | 11/13 | 11/13 | 11/13 |
| `event-hubs-dp-python-streaming` | 7/7 | 7/7 | 7/7 |
| `identity-dp-python-credential-chain` | 14/14 | 14/14 | 14/14 |
| `identity-dp-python-default-credential` | 5/5 | 4/5 | 4/5 |
| `identity-dp-python-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-python-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-python-crud` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-python-secret-config` | 11/12 | 11/12 | 11/12 |
| `resource-manager-mp-python-rg-crud` | 5/7 | 6/7 | 7/7 |
| `service-bus-dp-python-crud` | 6/7 | 6/7 | 6/7 |
| `service-bus-dp-python-order-processor` | 10/14 | 10/14 | 12/14 |
| `storage-dp-python-blob-event-notifier` | 11/11 | 9/11 | 11/11 |
| `storage-dp-python-blob-manager` | 10/10 | 10/10 | 10/10 |
| `storage-dp-python-crud` | 8/8 | 8/8 | 8/8 |
| `storage-dp-python-encrypted-uploader` | 19/19 | 18/19 | 19/19 |
| `storage-mp-python-account-mgmt` | 6/8 | 8/8 | 8/8 |
