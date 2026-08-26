# .NET Azure tools comparison

Run `20260827-010715` evaluated all 20 .NET prompts once with each explicit arm:

- `dotnet-azure-tools/baseline`: neither Azure MCP nor the Microsoft .NET SDK plugin.
- `dotnet-azure-tools/with-azure-tools`: Azure MCP plus `azure-sdk-dotnet`.

This run did not use Hyoka's automatic pairwise mode. All 40 evaluations completed without errors.

## Results by check type

| Check type | Baseline | With Azure tools | Difference |
|------------|----------|------------------|------------|
| Prompt-specific | 109/135 (80.7%) | 112/135 (83.0%) | +3 |
| Generic .NET | Not configured | Not configured | N/A |
| Workspace | Not configured | Not configured | N/A |
| Azure MCP usage | Not configured | Not configured | N/A |

Across the 20 prompt pairs, the Azure-tools arm improved 2 prompts, regressed 2, and tied 16.

## Prompt-specific comparison

| Prompt | Baseline | With Azure tools | Difference |
|--------|----------|------------------|------------|
| `app-configuration-dp-dotnet-crud` | 6/7 | 6/7 | 0 |
| `cosmos-db-dp-dotnet-crud` | 7/7 | 7/7 | 0 |
| `cosmos-db-dp-dotnet-error-handling` | 6/8 | 6/8 | 0 |
| `cosmos-db-dp-dotnet-pagination` | 7/8 | 7/8 | 0 |
| `event-hubs-dp-dotnet-streaming` | 7/7 | 7/7 | 0 |
| `identity-dp-dotnet-default-credential` | 5/5 | 5/5 | 0 |
| `identity-dp-dotnet-managed-identity` | 5/6 | 4/6 | -1 |
| `identity-dp-dotnet-service-principal` | 5/5 | 5/5 | 0 |
| `key-vault-dp-dotnet-crud` | 5/5 | 5/5 | 0 |
| `key-vault-dp-dotnet-error-handling` | 5/7 | 5/7 | 0 |
| `key-vault-dp-dotnet-pagination` | 5/7 | 5/7 | 0 |
| `key-vault-mp-dotnet-polling` | 7/8 | 6/8 | -1 |
| `resource-manager-mp-dotnet-rg-crud` | 2/6 | 6/6 | +4 |
| `service-bus-dp-dotnet-crud` | 7/8 | 7/8 | 0 |
| `storage-dp-dotnet-auth` | 3/5 | 3/5 | 0 |
| `storage-dp-dotnet-batch` | 5/8 | 5/8 | 0 |
| `storage-dp-dotnet-error-handling` | 4/6 | 4/6 | 0 |
| `storage-dp-dotnet-retries` | 6/8 | 7/8 | +1 |
| `storage-mp-dotnet-account-mgmt` | 6/7 | 6/7 | 0 |
| `storage-mp-dotnet-polling` | 6/7 | 6/7 | 0 |

## Interpretation

.NET has no generic language criteria, workspace check, or Azure MCP usage check in this evaluation set. Across the report metadata, no generator tool calls were recorded, so there is no evidence that Azure MCP was invoked.

The .NET plugin also has weaker service coverage than the other language plugins, including no Cosmos DB data-plane skill. The net prompt gain is concentrated in Resource Manager (+4), while 16 of 20 prompt pairs were unchanged. The result should not be treated as equivalent coverage to Python, JavaScript/TypeScript, or Java.

The Azure-tools arm took 3,370.3 seconds versus 2,427.7 seconds for the baseline, about 39% longer. This is one trial per prompt, so repeated runs are required to distinguish stable effects from model variance.
