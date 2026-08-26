# Java Azure tools comparison

Run `20260826-202215` evaluated all 19 Java prompts once with each explicit arm:

- `java-azure-tools/baseline`: neither Azure MCP nor the Microsoft Java SDK plugin.
- `java-azure-tools/with-azure-tools`: Azure MCP plus `azure-sdk-java`.

This run did not use Hyoka's automatic pairwise mode. All 38 evaluations completed without errors.

## Results by check type

| Check type | Baseline | With Azure tools | Difference |
|------------|----------|------------------|------------|
| Prompt-specific | 152/167 (91.0%) | 155/167 (92.8%) | +3 |
| Generic Java | 181/228 (79.4%) | 188/228 (82.5%) | +7 |
| Workspace | Not configured | Not configured | N/A |
| Azure MCP usage | Not configured | Not configured | N/A |

Across the 19 prompt pairs, the Azure-tools arm improved 3 prompts, regressed 2, and tied 14.

The generic Java differences were:

| Criterion | Baseline | With Azure tools | Difference |
|-----------|----------|------------------|------------|
| Reactor async usage | 18/19 | 19/19 | +1 |
| Azure SDK BOM | 5/19 | 4/19 | -1 |
| Client builder pattern | 13/19 | 14/19 | +1 |
| Code compiles | 17/19 | 18/19 | +1 |
| Correct dependencies | 15/19 | 16/19 | +1 |
| Correct imports | 19/19 | 19/19 | 0 |
| `DefaultAzureCredential` | 15/19 | 16/19 | +1 |
| LRO pattern | 15/19 | 16/19 | +1 |
| No deprecated classes | 19/19 | 19/19 | 0 |
| Pagination | 17/19 | 16/19 | -1 |
| Service-specific exceptions | 11/19 | 12/19 | +1 |
| Try-with-resources | 17/19 | 19/19 | +2 |

## Prompt-specific comparison

| Prompt | Baseline | With Azure tools | Difference |
|--------|----------|------------------|------------|
| `app-configuration-dp-java-crud` | 7/7 | 7/7 | 0 |
| `app-configuration-dp-java-feature-flags` | 9/9 | 8/9 | -1 |
| `cosmos-db-dp-java-crud` | 6/7 | 6/7 | 0 |
| `cosmos-db-dp-java-todo-repository` | 11/14 | 11/14 | 0 |
| `event-hubs-dp-java-streaming` | 7/7 | 7/7 | 0 |
| `identity-dp-java-credential-chain` | 13/14 | 14/14 | +1 |
| `identity-dp-java-default-credential` | 5/5 | 5/5 | 0 |
| `identity-dp-java-managed-identity` | 5/6 | 5/6 | 0 |
| `identity-dp-java-service-principal` | 5/5 | 5/5 | 0 |
| `key-vault-dp-java-crud` | 5/5 | 5/5 | 0 |
| `key-vault-dp-java-secret-config` | 10/10 | 10/10 | 0 |
| `resource-manager-mp-java-rg-crud` | 7/7 | 7/7 | 0 |
| `service-bus-dp-java-crud` | 7/7 | 6/7 | -1 |
| `service-bus-dp-java-order-processor` | 9/12 | 9/12 | 0 |
| `storage-dp-java-blob-event-notifier` | 6/10 | 9/10 | +3 |
| `storage-dp-java-blob-manager` | 7/7 | 7/7 | 0 |
| `storage-dp-java-crud` | 5/7 | 6/7 | +1 |
| `storage-dp-java-encrypted-uploader` | 20/20 | 20/20 | 0 |
| `storage-mp-java-account-mgmt` | 8/8 | 8/8 | 0 |

## Interpretation

The Java language criteria do not include workspace or Azure MCP usage checks. Across the report metadata, no generator tool calls were recorded, so there is no evidence that Azure MCP was invoked. The observed differences are therefore primarily attributable to the `azure-sdk-java` skill.

The Azure-tools arm took 8,824.6 seconds versus 8,061.9 seconds for the baseline, about 9% longer. This is one trial per prompt, so repeated runs are required to distinguish stable effects from model variance.
