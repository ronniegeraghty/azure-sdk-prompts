# JavaScript/TypeScript Azure tools comparison

Run `20260826-173422` evaluated 14 JavaScript/TypeScript prompts with each explicit arm:

- `js-ts-azure-tools/baseline`: neither Azure MCP nor the Microsoft TypeScript SDK plugin.
- `js-ts-azure-tools/with-azure-tools`: Azure MCP plus `azure-sdk-typescript`.

This run did not use Hyoka's automatic pairwise mode.

## Excluded case

`event-hubs-dp-js-ts-streaming` is excluded from the comparison. Its Azure-tools evaluation timed out while waiting for the Copilot SDK session to become idle. A targeted retry in run `20260826-200118` reached the same timeout. This may be related to local npm registry restrictions.

The generated summary retains the original 28 evaluations and reports one error. The comparison below uses the remaining 13 complete prompt pairs.

## Results by check type

| Check type | Baseline | With Azure tools | Difference |
|------------|----------|------------------|------------|
| Prompt-specific | 99/118 (83.9%) | 104/118 (88.1%) | +5 |
| Generic JS/TS | 90/130 (69.2%) | 96/130 (73.8%) | +6 |
| Workspace | Not configured | Not configured | N/A |
| Azure MCP usage | Not configured | Not configured | N/A |

Across the 13 included prompt pairs, the Azure-tools arm improved 4 prompts, regressed 0, and tied 9.

The generic JS/TS differences were:

| Criterion | Baseline | With Azure tools | Difference |
|-----------|----------|------------------|------------|
| `@azure/identity` authentication | 9/13 | 10/13 | +1 |
| Async/await pattern | 6/13 | 11/13 | +5 |
| Endpoint and credential constructor | 11/13 | 11/13 | 0 |
| Correct `@azure/*` packages | 13/13 | 13/13 | 0 |
| `@azure/logger` | 1/13 | 1/13 | 0 |
| LRO pattern | 11/13 | 11/13 | 0 |
| No deprecated packages | 13/13 | 13/13 | 0 |
| Correct `package.json` dependencies | 13/13 | 13/13 | 0 |
| Pagination with `for await...of` | 10/13 | 11/13 | +1 |
| `RestError` handling | 3/13 | 2/13 | -1 |

## Prompt-specific comparison

| Prompt | Baseline | With Azure tools | Difference |
|--------|----------|------------------|------------|
| `app-configuration-dp-js-ts-crud` | 6/8 | 6/8 | 0 |
| `cosmos-db-dp-js-ts-crud` | 5/7 | 6/7 | +1 |
| `identity-dp-js-ts-default-credential` | 4/5 | 4/5 | 0 |
| `identity-dp-js-ts-managed-identity` | 4/6 | 5/6 | +1 |
| `identity-dp-js-ts-service-principal` | 5/5 | 5/5 | 0 |
| `key-vault-dp-js-ts-crud` | 4/5 | 4/5 | 0 |
| `key-vault-dp-js-ts-secret-config` | 12/13 | 13/13 | +1 |
| `resource-manager-mp-js-ts-rg-crud` | 8/8 | 8/8 | 0 |
| `service-bus-dp-js-ts-crud` | 7/8 | 7/8 | 0 |
| `storage-dp-js-ts-blob-manager` | 8/12 | 10/12 | +2 |
| `storage-dp-js-ts-crud` | 7/8 | 7/8 | 0 |
| `storage-dp-js-ts-encrypted-uploader` | 23/25 | 23/25 | 0 |
| `storage-mp-js-ts-account-mgmt` | 6/8 | 6/8 | 0 |

## Interpretation

The JS/TS language criteria do not include workspace or Azure MCP usage checks. Across the report metadata, no generator tool calls were recorded, so there is no evidence that Azure MCP was invoked. The observed differences are therefore primarily attributable to the `azure-sdk-typescript` skill.

The Azure-tools arm took 4,214.8 seconds for the included cases versus 3,307.4 seconds for the baseline, about 27% longer. This is one trial per prompt, so repeated runs are required to distinguish stable effects from model variance.
