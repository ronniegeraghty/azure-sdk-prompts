# .NET three-way comparison

The report includes **20 complete prompt triplets** and **60 valid evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 110/135 | 81.5% | - |
| Azure Skill + MCP | 108/135 | 80% | -1.5 pp |
| Azure Skill + MCP + Microsoft Skills | 104/135 | 77% | -4.5 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 2 | 4 | 14 |
| Azure Skill + MCP + Microsoft Skills vs baseline | 1 | 3 | 16 |

Compared with baseline, Azure Skill + MCP changed the prompt-check rate by **-1.5 pp**, while Azure Skill + MCP + Microsoft Skills changed it by **-4.5 pp**.

## Language checks

No generic .NET language criteria are configured.

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| `app-configuration-dp-dotnet-crud` | 6/7 | 6/7 | 6/7 |
| `cosmos-db-dp-dotnet-crud` | 7/7 | 7/7 | 7/7 |
| `cosmos-db-dp-dotnet-error-handling` | 6/8 | 6/8 | 6/8 |
| `cosmos-db-dp-dotnet-pagination` | 7/8 | 7/8 | 7/8 |
| `event-hubs-dp-dotnet-streaming` | 7/7 | 7/7 | 6/7 |
| `identity-dp-dotnet-default-credential` | 5/5 | 4/5 | 5/5 |
| `identity-dp-dotnet-managed-identity` | 3/6 | 3/6 | 3/6 |
| `identity-dp-dotnet-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-dotnet-crud` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-dotnet-error-handling` | 4/7 | 4/7 | 4/7 |
| `key-vault-dp-dotnet-pagination` | 6/7 | 5/7 | 5/7 |
| `key-vault-mp-dotnet-polling` | 7/8 | 7/8 | 7/8 |
| `resource-manager-mp-dotnet-rg-crud` | 3/6 | 4/6 | 6/6 |
| `service-bus-dp-dotnet-crud` | 7/8 | 7/8 | 7/8 |
| `storage-dp-dotnet-auth` | 3/5 | 3/5 | 3/5 |
| `storage-dp-dotnet-batch` | 6/8 | 5/8 | 6/8 |
| `storage-dp-dotnet-error-handling` | 4/6 | 3/6 | 4/6 |
| `storage-dp-dotnet-retries` | 7/8 | 8/8 | 0/8 |
| `storage-mp-dotnet-account-mgmt` | 6/7 | 6/7 | 6/7 |
| `storage-mp-dotnet-polling` | 6/7 | 6/7 | 6/7 |
