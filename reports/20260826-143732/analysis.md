# Python Azure tools comparison

Run `20260826-143732` evaluated all 19 Python prompts once with each explicit arm:

- `python-azure-tools/baseline`: neither Azure MCP nor the Microsoft Python SDK plugin.
- `python-azure-tools/with-azure-tools`: Azure MCP plus `azure-sdk-python`.

This run did not use Hyoka's automatic pairwise mode.

## Results by check type

| Check type | Baseline | With Azure tools | Difference |
|------------|----------|------------------|------------|
| Prompt-specific | 158/172 (91.9%) | 158/172 (91.9%) | 0 |
| Generic Python | 77/95 (81.1%) | 90/95 (94.7%) | +13 |
| Workspace | 0/19 | 0/19 | 0 |
| Azure MCP usage | 0/19 | 0/19 | 0 |

Prompt-specific checks are the best measure of task correctness. Across the 19 prompt pairs, the Azure-tools arm improved 5 prompts, regressed 4, and tied 10. The aggregate prompt-specific score was exactly equal.

The generic Python improvement came from:

| Criterion | Baseline | With Azure tools |
|-----------|----------|------------------|
| Correct package imports | 19/19 | 19/19 |
| `DefaultAzureCredential` usage | 15/19 | 19/19 |
| Client context management | 13/19 | 18/19 |
| Async client usage | 18/19 | 18/19 |
| Azure exception handling | 12/19 | 16/19 |

## Prompt-specific comparison

| Prompt | Baseline | With Azure tools | Difference |
|--------|----------|------------------|------------|
| `app-configuration-dp-python-crud` | 5/6 | 6/6 | +1 |
| `app-configuration-dp-python-feature-flags` | 9/9 | 9/9 | 0 |
| `cosmos-db-dp-python-crud` | 5/6 | 4/6 | -1 |
| `cosmos-db-dp-python-todo-repository` | 10/13 | 10/13 | 0 |
| `event-hubs-dp-python-streaming` | 7/7 | 7/7 | 0 |
| `identity-dp-python-credential-chain` | 14/14 | 14/14 | 0 |
| `identity-dp-python-default-credential` | 4/5 | 5/5 | +1 |
| `identity-dp-python-managed-identity` | 4/6 | 5/6 | +1 |
| `identity-dp-python-service-principal` | 5/5 | 5/5 | 0 |
| `key-vault-dp-python-crud` | 5/5 | 5/5 | 0 |
| `key-vault-dp-python-secret-config` | 12/12 | 10/12 | -2 |
| `resource-manager-mp-python-rg-crud` | 7/7 | 6/7 | -1 |
| `service-bus-dp-python-crud` | 5/7 | 6/7 | +1 |
| `service-bus-dp-python-order-processor` | 12/14 | 12/14 | 0 |
| `storage-dp-python-blob-event-notifier` | 11/11 | 11/11 | 0 |
| `storage-dp-python-blob-manager` | 10/10 | 8/10 | -2 |
| `storage-dp-python-crud` | 8/8 | 8/8 | 0 |
| `storage-dp-python-encrypted-uploader` | 19/19 | 19/19 | 0 |
| `storage-mp-python-account-mgmt` | 6/8 | 8/8 | +2 |

## Interpretation

Azure MCP was configured only in the Azure-tools arm, but no Azure MCP invocation was recorded in any of its 19 generations. The measured difference therefore reflects the combined configuration as assigned, but the observed behavior is primarily attributable to the `azure-sdk-python` skill.

The generated summary marks all 38 evaluations failed because every evaluation failed at least one generic check. Do not use that all-or-nothing result as the experiment conclusion:

- The baseline is expected to fail the Azure MCP usage criterion because it has no Azure MCP.
- The Azure-tools arm also failed that criterion because it did not invoke Azure MCP.
- The workspace check reported 0/19 for both arms even when generated Python files were present, so it appears unreliable for this run.

The Azure-tools arm took 5,886.6 seconds in aggregate versus 4,053.6 seconds for the baseline, about 45% longer. This is one trial per prompt, so repeated runs are required to distinguish stable effects from model variance.
