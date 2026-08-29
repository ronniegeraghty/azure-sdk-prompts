# .NET three-way comparison

The report includes **19 complete prompt triplets** and **57 selected evaluations**.

## Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 109/129 | 84.5% | - |
| Azure Skill + MCP | 101/129 | 78.3% | -6.2 pp |
| Azure Skill + MCP + Microsoft Skills | 109/129 | 84.5% | 0.0 pp |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | 1 | 3 | 15 |
| Azure Skill + MCP + Microsoft Skills vs baseline | 2 | 2 | 15 |

## Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | 0/0 | Not configured | - |
| Azure Skill + MCP | 0/0 | Not configured | n/a |
| Azure Skill + MCP + Microsoft Skills | 0/0 | Not configured | n/a |

## Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| Workspace checks | Not configured | Not configured | Not configured |
| Azure MCP usage checks | Not configured | Not configured | Not configured |

Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.

## Excluded prompt triplets

| Prompt ID | Affected arm | Reason |
|---|---|---|
| `identity-dp-dotnet-managed-identity` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | The primary timed out with no output; the retry completed without a timeout but still produced no generated files. |

## Per-prompt prompt checks

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---:|---:|---:|
| `app-configuration-dp-dotnet-crud` | 6/7 | 6/7 | 6/7 |
| `cosmos-db-dp-dotnet-crud` | 7/7 | 7/7 | 7/7 |
| `cosmos-db-dp-dotnet-error-handling` | 7/8 | 6/8 | 6/8 |
| `cosmos-db-dp-dotnet-pagination` | 7/8 | 7/8 | 8/8 |
| `event-hubs-dp-dotnet-streaming` | 7/7 | 7/7 | 7/7 |
| `identity-dp-dotnet-default-credential` | 5/5 | 4/5 | 5/5 |
| `identity-dp-dotnet-service-principal` | 5/5 | 5/5 | 5/5 |
| `key-vault-dp-dotnet-crud` | 5/5 | 5/5 | 4/5 |
| `key-vault-dp-dotnet-error-handling` | 4/7 | 4/7 | 4/7 |
| `key-vault-dp-dotnet-pagination` | 5/7 | 6/7 | 6/7 |
| `key-vault-mp-dotnet-polling` | 7/8 | 0/8 | 7/8 |
| `resource-manager-mp-dotnet-rg-crud` | 6/6 | 6/6 | 6/6 |
| `service-bus-dp-dotnet-crud` | 7/8 | 7/8 | 7/8 |
| `storage-dp-dotnet-auth` | 3/5 | 3/5 | 3/5 |
| `storage-dp-dotnet-batch` | 5/8 | 5/8 | 5/8 |
| `storage-dp-dotnet-error-handling` | 4/6 | 4/6 | 4/6 |
| `storage-dp-dotnet-retries` | 7/8 | 7/8 | 7/8 |
| `storage-mp-dotnet-account-mgmt` | 6/7 | 6/7 | 6/7 |
| `storage-mp-dotnet-polling` | 6/7 | 6/7 | 6/7 |
