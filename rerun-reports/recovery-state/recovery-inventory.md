# Three-way evaluation recovery inventory

This inventory preserves reports with valid grader outcomes, including ordinary check failures. Reports with `tool_load_failure`, SDK timeouts, or no report are excluded and listed for rerun.

Score cells use `P passed/total; L passed/total`, where P is prompt-specific checks and L is language checks. `.NET` has no configured language checks.

## Coverage summary

| Language | Expected combinations | Valid results | Required reruns |
|---|---:|---:|---:|
| Python | 57 | 18 | 39 |
| JavaScript/TypeScript | 42 | 14 | 28 |
| Java | 57 | 13 | 44 |
| .NET | 60 | 20 | 40 |
| **Total** | **216** | **65** | **151** |

## Python valid-data matrix

Run: `20260827-143238`

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---|---|---|
| `app-configuration-dp-python-crud` | P 6/6; L 4/5 | RERUN (tool_load_failure) | RERUN (missing) |
| `app-configuration-dp-python-feature-flags` | RERUN (missing) | RERUN (missing) | RERUN (tool_load_failure) |
| `cosmos-db-dp-python-crud` | P 6/6; L 3/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-python-todo-repository` | P 11/13; L 4/5 | RERUN (missing) | RERUN (missing) |
| `event-hubs-dp-python-streaming` | RERUN (missing) | RERUN (missing) | RERUN (tool_load_failure) |
| `identity-dp-python-credential-chain` | P 14/14; L 4/5 | RERUN (SDK timeout) | RERUN (tool_load_failure) |
| `identity-dp-python-default-credential` | P 5/5; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-python-managed-identity` | P 5/6; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-python-service-principal` | P 5/5; L 5/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-python-crud` | P 5/5; L 5/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-python-secret-config` | P 11/12; L 5/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `resource-manager-mp-python-rg-crud` | P 5/7; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `service-bus-dp-python-crud` | P 6/7; L 3/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `service-bus-dp-python-order-processor` | P 10/14; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-python-blob-event-notifier` | P 11/11; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-python-blob-manager` | P 10/10; L 5/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-python-crud` | P 8/8; L 5/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-python-encrypted-uploader` | P 19/19; L 5/5 | RERUN (tool_load_failure) | P 19/19; L 4/5 |
| `storage-mp-python-account-mgmt` | P 6/8; L 4/5 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |

## JavaScript/TypeScript valid-data matrix

Run: `20260827-143332`

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---|---|---|
| `app-configuration-dp-js-ts-crud` | P 6/8; L 6/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-js-ts-crud` | P 5/7; L 5/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `event-hubs-dp-js-ts-streaming` | P 7/8; L 5/10 | RERUN (SDK timeout) | RERUN (SDK timeout) |
| `identity-dp-js-ts-default-credential` | P 4/5; L 8/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-js-ts-managed-identity` | P 5/6; L 8/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-js-ts-service-principal` | P 5/5; L 6/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-js-ts-crud` | P 4/5; L 8/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-js-ts-secret-config` | P 12/13; L 8/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `resource-manager-mp-js-ts-rg-crud` | P 8/8; L 7/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `service-bus-dp-js-ts-crud` | P 7/8; L 5/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-js-ts-blob-manager` | P 9/12; L 7/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-js-ts-crud` | P 7/8; L 8/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-js-ts-encrypted-uploader` | P 25/25; L 7/10 | RERUN (SDK timeout) | RERUN (tool_load_failure) |
| `storage-mp-js-ts-account-mgmt` | P 7/8; L 7/10 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |

## Java valid-data matrix

Run: `20260827-143433`

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---|---|---|
| `app-configuration-dp-java-crud` | P 7/7; L 8/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `app-configuration-dp-java-feature-flags` | P 8/9; L 9/12 | RERUN (SDK timeout) | RERUN (tool_load_failure) |
| `cosmos-db-dp-java-crud` | P 6/7; L 8/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-java-todo-repository` | P 11/14; L 10/12 | RERUN (SDK timeout) | RERUN (SDK timeout) |
| `event-hubs-dp-java-streaming` | P 7/7; L 7/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-java-credential-chain` | P 14/14; L 10/12 | RERUN (SDK timeout) | RERUN (tool_load_failure) |
| `identity-dp-java-default-credential` | P 5/5; L 12/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-java-managed-identity` | P 5/6; L 11/12 | RERUN (tool_load_failure) | RERUN (missing) |
| `identity-dp-java-service-principal` | RERUN (missing) | RERUN (missing) | RERUN (tool_load_failure) |
| `key-vault-dp-java-crud` | P 5/5; L 11/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-java-secret-config` | P 10/10; L 11/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `resource-manager-mp-java-rg-crud` | P 7/7; L 9/12 | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `service-bus-dp-java-crud` | RERUN (missing) | RERUN (missing) | RERUN (tool_load_failure) |
| `service-bus-dp-java-order-processor` | P 9/12; L 10/12 | RERUN (missing) | RERUN (missing) |
| `storage-dp-java-blob-event-notifier` | RERUN (missing) | RERUN (missing) | RERUN (missing) |
| `storage-dp-java-blob-manager` | RERUN (missing) | RERUN (missing) | RERUN (missing) |
| `storage-dp-java-crud` | RERUN (missing) | RERUN (missing) | RERUN (missing) |
| `storage-dp-java-encrypted-uploader` | RERUN (missing) | RERUN (missing) | RERUN (missing) |
| `storage-mp-java-account-mgmt` | RERUN (missing) | RERUN (missing) | P 0/8; L 3/12 |

## .NET valid-data matrix

Run: `20260827-143539`

| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft Skills |
|---|---|---|---|
| `app-configuration-dp-dotnet-crud` | P 6/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-dotnet-crud` | P 7/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-dotnet-error-handling` | P 6/8; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `cosmos-db-dp-dotnet-pagination` | P 7/8; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `event-hubs-dp-dotnet-streaming` | P 7/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-dotnet-default-credential` | P 5/5; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `identity-dp-dotnet-managed-identity` | P 3/6; L n/a | RERUN (tool_load_failure) | RERUN (SDK timeout) |
| `identity-dp-dotnet-service-principal` | P 5/5; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-dotnet-crud` | P 5/5; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-dotnet-error-handling` | P 4/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-dp-dotnet-pagination` | P 6/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `key-vault-mp-dotnet-polling` | P 7/8; L n/a | RERUN (SDK timeout) | RERUN (tool_load_failure) |
| `resource-manager-mp-dotnet-rg-crud` | P 3/6; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `service-bus-dp-dotnet-crud` | P 7/8; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-dotnet-auth` | P 3/5; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-dotnet-batch` | P 6/8; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-dotnet-error-handling` | P 4/6; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-dp-dotnet-retries` | P 7/8; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-mp-dotnet-account-mgmt` | P 6/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |
| `storage-mp-dotnet-polling` | P 6/7; L n/a | RERUN (tool_load_failure) | RERUN (tool_load_failure) |

## Required reruns

| Language | Prompt ID | Variant | Reason |
|---|---|---|---|
| .NET | `app-configuration-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `app-configuration-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-pagination` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `cosmos-db-dp-dotnet-pagination` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `event-hubs-dp-dotnet-streaming` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `event-hubs-dp-dotnet-streaming` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `identity-dp-dotnet-default-credential` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `identity-dp-dotnet-default-credential` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `identity-dp-dotnet-managed-identity` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `identity-dp-dotnet-managed-identity` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | SDK timeout |
| .NET | `identity-dp-dotnet-service-principal` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `identity-dp-dotnet-service-principal` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-pagination` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `key-vault-dp-dotnet-pagination` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `key-vault-mp-dotnet-polling` | `dotnet-azure-skills/azure-skill-mcp` | SDK timeout |
| .NET | `key-vault-mp-dotnet-polling` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `resource-manager-mp-dotnet-rg-crud` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `resource-manager-mp-dotnet-rg-crud` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `service-bus-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `service-bus-dp-dotnet-crud` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-dp-dotnet-auth` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-dp-dotnet-auth` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-dp-dotnet-batch` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-dp-dotnet-batch` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-dp-dotnet-error-handling` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-dp-dotnet-retries` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-dp-dotnet-retries` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-mp-dotnet-account-mgmt` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-mp-dotnet-account-mgmt` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| .NET | `storage-mp-dotnet-polling` | `dotnet-azure-skills/azure-skill-mcp` | tool_load_failure |
| .NET | `storage-mp-dotnet-polling` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `app-configuration-dp-java-crud` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `app-configuration-dp-java-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `app-configuration-dp-java-feature-flags` | `java-azure-skills/azure-skill-mcp` | SDK timeout |
| Java | `app-configuration-dp-java-feature-flags` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `cosmos-db-dp-java-crud` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `cosmos-db-dp-java-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `cosmos-db-dp-java-todo-repository` | `java-azure-skills/azure-skill-mcp` | SDK timeout |
| Java | `cosmos-db-dp-java-todo-repository` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | SDK timeout |
| Java | `event-hubs-dp-java-streaming` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `event-hubs-dp-java-streaming` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `identity-dp-java-credential-chain` | `java-azure-skills/azure-skill-mcp` | SDK timeout |
| Java | `identity-dp-java-credential-chain` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `identity-dp-java-default-credential` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `identity-dp-java-default-credential` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `identity-dp-java-managed-identity` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `identity-dp-java-managed-identity` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `identity-dp-java-service-principal` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `identity-dp-java-service-principal` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `identity-dp-java-service-principal` | `java-azure-skills/baseline` | missing report |
| Java | `key-vault-dp-java-crud` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `key-vault-dp-java-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `key-vault-dp-java-secret-config` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `key-vault-dp-java-secret-config` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `resource-manager-mp-java-rg-crud` | `java-azure-skills/azure-skill-mcp` | tool_load_failure |
| Java | `resource-manager-mp-java-rg-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `service-bus-dp-java-crud` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `service-bus-dp-java-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Java | `service-bus-dp-java-crud` | `java-azure-skills/baseline` | missing report |
| Java | `service-bus-dp-java-order-processor` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `service-bus-dp-java-order-processor` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `storage-dp-java-blob-event-notifier` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `storage-dp-java-blob-event-notifier` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `storage-dp-java-blob-event-notifier` | `java-azure-skills/baseline` | missing report |
| Java | `storage-dp-java-blob-manager` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `storage-dp-java-blob-manager` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `storage-dp-java-blob-manager` | `java-azure-skills/baseline` | missing report |
| Java | `storage-dp-java-crud` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `storage-dp-java-crud` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `storage-dp-java-crud` | `java-azure-skills/baseline` | missing report |
| Java | `storage-dp-java-encrypted-uploader` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `storage-dp-java-encrypted-uploader` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Java | `storage-dp-java-encrypted-uploader` | `java-azure-skills/baseline` | missing report |
| Java | `storage-mp-java-account-mgmt` | `java-azure-skills/azure-skill-mcp` | missing report |
| Java | `storage-mp-java-account-mgmt` | `java-azure-skills/baseline` | missing report |
| JavaScript/TypeScript | `app-configuration-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `app-configuration-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `cosmos-db-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `cosmos-db-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `event-hubs-dp-js-ts-streaming` | `js-ts-azure-skills/azure-skill-mcp` | SDK timeout |
| JavaScript/TypeScript | `event-hubs-dp-js-ts-streaming` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | SDK timeout |
| JavaScript/TypeScript | `identity-dp-js-ts-default-credential` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `identity-dp-js-ts-default-credential` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `identity-dp-js-ts-managed-identity` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `identity-dp-js-ts-managed-identity` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `identity-dp-js-ts-service-principal` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `identity-dp-js-ts-service-principal` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `key-vault-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `key-vault-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `key-vault-dp-js-ts-secret-config` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `key-vault-dp-js-ts-secret-config` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `resource-manager-mp-js-ts-rg-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `resource-manager-mp-js-ts-rg-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `service-bus-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `service-bus-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `storage-dp-js-ts-blob-manager` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `storage-dp-js-ts-blob-manager` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `storage-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `storage-dp-js-ts-crud` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `storage-dp-js-ts-encrypted-uploader` | `js-ts-azure-skills/azure-skill-mcp` | SDK timeout |
| JavaScript/TypeScript | `storage-dp-js-ts-encrypted-uploader` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| JavaScript/TypeScript | `storage-mp-js-ts-account-mgmt` | `js-ts-azure-skills/azure-skill-mcp` | tool_load_failure |
| JavaScript/TypeScript | `storage-mp-js-ts-account-mgmt` | `js-ts-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `app-configuration-dp-python-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `app-configuration-dp-python-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Python | `app-configuration-dp-python-feature-flags` | `python-azure-skills/azure-skill-mcp` | missing report |
| Python | `app-configuration-dp-python-feature-flags` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `app-configuration-dp-python-feature-flags` | `python-azure-skills/baseline` | missing report |
| Python | `cosmos-db-dp-python-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `cosmos-db-dp-python-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `cosmos-db-dp-python-todo-repository` | `python-azure-skills/azure-skill-mcp` | missing report |
| Python | `cosmos-db-dp-python-todo-repository` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | missing report |
| Python | `event-hubs-dp-python-streaming` | `python-azure-skills/azure-skill-mcp` | missing report |
| Python | `event-hubs-dp-python-streaming` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `event-hubs-dp-python-streaming` | `python-azure-skills/baseline` | missing report |
| Python | `identity-dp-python-credential-chain` | `python-azure-skills/azure-skill-mcp` | SDK timeout |
| Python | `identity-dp-python-credential-chain` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `identity-dp-python-default-credential` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `identity-dp-python-default-credential` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `identity-dp-python-managed-identity` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `identity-dp-python-managed-identity` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `identity-dp-python-service-principal` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `identity-dp-python-service-principal` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `key-vault-dp-python-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `key-vault-dp-python-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `key-vault-dp-python-secret-config` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `key-vault-dp-python-secret-config` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `resource-manager-mp-python-rg-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `resource-manager-mp-python-rg-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `service-bus-dp-python-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `service-bus-dp-python-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `service-bus-dp-python-order-processor` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `service-bus-dp-python-order-processor` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `storage-dp-python-blob-event-notifier` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `storage-dp-python-blob-event-notifier` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `storage-dp-python-blob-manager` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `storage-dp-python-blob-manager` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `storage-dp-python-crud` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `storage-dp-python-crud` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |
| Python | `storage-dp-python-encrypted-uploader` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `storage-mp-python-account-mgmt` | `python-azure-skills/azure-skill-mcp` | tool_load_failure |
| Python | `storage-mp-python-account-mgmt` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | tool_load_failure |

Total required reruns: **151**.
