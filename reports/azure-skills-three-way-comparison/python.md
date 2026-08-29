# Python three-way comparison

The report includes **19 complete prompt triplets** and **57 selected evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 150/172 | 87.2% | - |
| Azure Skill + MCP | 156/172 | 90.7% | +3.5 pp |
| Azure Skill + MCP + Microsoft Skills | 159/172 | 92.4% | +5.2 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 7 | 2 | 10 |
| Azure Skill + MCP + Microsoft Skills vs baseline | 5 | 2 | 12 |

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 77/95 | 81.1% | - |
| Azure Skill + MCP | 84/95 | 88.4% | +7.3 pp |
| Azure Skill + MCP + Microsoft Skills | 90/95 | 94.7% | +13.6 pp |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Workspace checks | 0/19 (0.0%) | 0/19 (0.0%) | 0/19 (0.0%) |
| Azure MCP usage checks | 0/19 (0.0%) | 19/19 (100.0%) | 17/19 (89.5%) |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| `app-configuration-dp-python-crud` | 5/6 | 6/6 | 6/6 |
| `app-configuration-dp-python-feature-flags` | 8/9 | 9/9 | 9/9 |
| `cosmos-db-dp-python-crud` | 5/6 | 5/6 | 5/6 |
| `cosmos-db-dp-python-todo-repository` | 11/13 | 11/13 | 11/13 |
| `event-hubs-dp-python-streaming` | 7/7 | 7/7 | 7/7 |
| `identity-dp-python-credential-chain` | 14/14 | 14/14 | 14/14 |
| `identity-dp-python-default-credential` | 4/5 | 5/5 | 3/5 |
| `identity-dp-python-managed-identity` | 5/6 | 5/6 | 5/6 |
| `identity-dp-python-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-python-crud` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-python-secret-config` | 11/12 | 6/12 | 11/12 |
| `resource-manager-mp-python-rg-crud` | 6/7 | 7/7 | 6/7 |
| `service-bus-dp-python-crud` | 5/7 | 5/7 | 6/7 |
| `service-bus-dp-python-order-processor` | 11/14 | 12/14 | 11/14 |
| `storage-dp-python-blob-event-notifier` | 11/11 | 9/11 | 11/11 |
| `storage-dp-python-blob-manager` | 10/10 | 10/10 | 9/10 |
| `storage-dp-python-crud` | 8/8 | 8/8 | 8/8 |
| `storage-dp-python-encrypted-uploader` | 18/19 | 19/19 | 19/19 |
| `storage-mp-python-account-mgmt` | 1/8 | 8/8 | 8/8 |
