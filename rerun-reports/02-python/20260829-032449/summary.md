# Evaluation Summary: 20260829-032449

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260829-032449` |
| Timestamp | 2026-08-28T19:24:49Z |
| Total Prompts | 19 |
| Total Configs | 3 |
| Total Evaluations | 57 |
| Passed | 0 |
| Failed | 56 |
| Errors | 1 |
| Duration | 15861.1s |

## Comparison Matrix

| Prompt | python-azure-skills/azure-skill-mcp | python-azure-skills/azure-skill-mcp-microsoft-skill | python-azure-skills/baseline |
|--------|--------|--------|--------|
| app-configuration-dp-python-crud | ❌ 12/13 | ❌ 12/13 | ❌ 8/13 |
| app-configuration-dp-python-feature-flags | ❌ 15/16 | ❌ 14/16 | ❌ 12/16 |
| cosmos-db-dp-python-crud | ❌ 9/13 | ❌ 10/13 | ❌ 8/13 |
| cosmos-db-dp-python-todo-repository | ❌ 16/20 | ❌ 17/20 | ❌ 15/20 |
| event-hubs-dp-python-streaming | ❌ 10/14 | ❌ 12/14 | ❌ 10/14 |
| identity-dp-python-credential-chain | ❌ 19/21 | ❌ 20/21 | ❌ 18/21 |
| identity-dp-python-default-credential | ❌ 11/12 | ❌ 8/12 | ❌ 8/12 |
| identity-dp-python-managed-identity | ❌ 11/13 | ❌ 10/13 | ❌ 9/13 |
| identity-dp-python-service-principal | ❌ 10/12 | ❌ 11/12 | ❌ 10/12 |
| key-vault-dp-python-crud | ❌ 11/12 | ❌ 11/12 | ❌ 9/12 |
| key-vault-dp-python-secret-config | ❌ 12/19 | ❌ 17/19 | ❌ 16/19 |
| resource-manager-mp-python-rg-crud | ❌ 13/14 | ❌ 12/14 | ❌ 10/14 |
| service-bus-dp-python-crud | ❌ 10/14 | ❌ 11/14 | ❌ 8/14 |
| service-bus-dp-python-order-processor | ❌ 17/21 | ❌ 16/21 | ❌ 15/21 |
| storage-dp-python-blob-event-notifier | ❌ 14/18 | ❌ 17/18 | ❌ 16/18 |
| storage-dp-python-blob-manager | ❌ 16/17 | ❌ 15/17 | ❌ 15/17 |
| storage-dp-python-crud | ❌ 14/15 | ❌ 14/15 | ❌ 13/15 |
| storage-dp-python-encrypted-uploader | ❌ 25/26 | ❌ 25/26 | ❌ 22/26 |
| storage-mp-python-account-mgmt | ❌ 14/15 | ❌ 13/15 | ❌ 5/15 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-python-crud](results/app-configuration/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 12/13 | 242.8s | 2 |
| [app-configuration-dp-python-crud](results/app-configuration/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 12/13 | 240.9s | 2 |
| [app-configuration-dp-python-crud](results/app-configuration/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 8/13 | 119.9s | 2 |
| [app-configuration-dp-python-feature-flags](results/app-configuration/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 15/16 | 350.8s | 5 |
| [app-configuration-dp-python-feature-flags](results/app-configuration/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 14/16 | 321.5s | 7 |
| [app-configuration-dp-python-feature-flags](results/app-configuration/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 12/16 | 366.1s | 6 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 9/13 | 226.3s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 10/13 | 152.5s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 8/13 | 146.9s | 2 |
| [cosmos-db-dp-python-todo-repository](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 16/20 | 414.8s | 9 |
| [cosmos-db-dp-python-todo-repository](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/20 | 343.3s | 0 |
| [cosmos-db-dp-python-todo-repository](results/cosmos-db/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 15/20 | 282.2s | 7 |
| [event-hubs-dp-python-streaming](results/event-hubs/data-plane/python/streaming/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 10/14 | 233.2s | 2 |
| [event-hubs-dp-python-streaming](results/event-hubs/data-plane/python/streaming/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 12/14 | 243.3s | 2 |
| [event-hubs-dp-python-streaming](results/event-hubs/data-plane/python/streaming/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 10/14 | 140.5s | 2 |
| [identity-dp-python-credential-chain](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 19/21 | 321.4s | 8 |
| [identity-dp-python-credential-chain](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 20/21 | 274.1s | 9 |
| [identity-dp-python-credential-chain](results/identity/data-plane/python/auth/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 18/21 | 232.5s | 7 |
| [identity-dp-python-default-credential](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 11/12 | 242.0s | 4 |
| [identity-dp-python-default-credential](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 8/12 | 253.1s | 3 |
| [identity-dp-python-default-credential](results/identity/data-plane/python/auth/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 8/12 | 257.9s | 3 |
| [identity-dp-python-managed-identity](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 11/13 | 358.7s | 15 |
| [identity-dp-python-managed-identity](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 10/13 | 218.5s | 14 |
| [identity-dp-python-managed-identity](results/identity/data-plane/python/auth/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 9/13 | 211.0s | 14 |
| [identity-dp-python-service-principal](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 10/12 | 275.5s | 3 |
| [identity-dp-python-service-principal](results/identity/data-plane/python/auth/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 11/12 | 201.2s | 4 |
| [identity-dp-python-service-principal](results/identity/data-plane/python/auth/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 10/12 | 248.0s | 3 |
| [key-vault-dp-python-crud](results/key-vault/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 11/12 | 225.7s | 2 |
| [key-vault-dp-python-crud](results/key-vault/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 11/12 | 184.2s | 2 |
| [key-vault-dp-python-crud](results/key-vault/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 9/12 | 163.2s | 2 |
| [key-vault-dp-python-secret-config](results/key-vault/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 12/19 | 325.1s | 9 |
| [key-vault-dp-python-secret-config](results/key-vault/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/19 | 316.1s | 7 |
| [key-vault-dp-python-secret-config](results/key-vault/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 16/19 | 229.3s | 8 |
| [resource-manager-mp-python-rg-crud](results/resource-manager/management-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 13/14 | 418.2s | 2 |
| [resource-manager-mp-python-rg-crud](results/resource-manager/management-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 12/14 | 263.6s | 3 |
| [resource-manager-mp-python-rg-crud](results/resource-manager/management-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 10/14 | 148.2s | 3 |
| [service-bus-dp-python-crud](results/service-bus/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 10/14 | 194.3s | 3 |
| [service-bus-dp-python-crud](results/service-bus/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 11/14 | 208.8s | 4 |
| [service-bus-dp-python-crud](results/service-bus/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 8/14 | 151.5s | 2 |
| [service-bus-dp-python-order-processor](results/service-bus/data-plane/python/streaming/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 17/21 | 420.6s | 6 |
| [service-bus-dp-python-order-processor](results/service-bus/data-plane/python/streaming/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 16/21 | 331.8s | 9 |
| [service-bus-dp-python-order-processor](results/service-bus/data-plane/python/streaming/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 15/21 | 316.9s | 11 |
| [storage-dp-python-blob-event-notifier](results/storage/data-plane/python/streaming/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 14/18 | 356.5s | 9 |
| [storage-dp-python-blob-event-notifier](results/storage/data-plane/python/streaming/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/18 | 366.9s | 6 |
| [storage-dp-python-blob-event-notifier](results/storage/data-plane/python/streaming/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 16/18 | 311.6s | 6 |
| [storage-dp-python-blob-manager](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 16/17 | 298.7s | 6 |
| [storage-dp-python-blob-manager](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 15/17 | 361.5s | 6 |
| [storage-dp-python-blob-manager](results/storage/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 15/17 | 367.4s | 8 |
| [storage-dp-python-crud](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 14/15 | 190.0s | 2 |
| [storage-dp-python-crud](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 14/15 | 178.6s | 2 |
| [storage-dp-python-crud](results/storage/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 13/15 | 132.6s | 2 |
| [storage-dp-python-encrypted-uploader](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 25/26 | 447.5s | 8 |
| [storage-dp-python-encrypted-uploader](results/storage/data-plane/python/crud/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 25/26 | 360.0s | 10 |
| [storage-dp-python-encrypted-uploader](results/storage/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 22/26 | 381.9s | 6 |
| [storage-mp-python-account-mgmt](results/storage/management-plane/python/provisioning/python-azure-skills/azure-skill-mcp/report.md) | python-azure-skills/azure-skill-mcp | ❌ | 14/15 | 235.5s | 2 |
| [storage-mp-python-account-mgmt](results/storage/management-plane/python/provisioning/python-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | python-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 13/15 | 814.7s | 2 |
| [storage-mp-python-account-mgmt](results/storage/management-plane/python/provisioning/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 5/15 | 237.0s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-python-todo-repository | 282.2s (python-azure-skills/baseline) | 346.8s | 414.8s (python-azure-skills/azure-skill-mcp) |
| storage-dp-python-crud | 132.6s (python-azure-skills/baseline) | 167.1s | 190.0s (python-azure-skills/azure-skill-mcp) |
| identity-dp-python-default-credential | 242.0s (python-azure-skills/azure-skill-mcp) | 251.0s | 257.9s (python-azure-skills/baseline) |
| storage-dp-python-blob-manager | 298.7s (python-azure-skills/azure-skill-mcp) | 342.5s | 367.4s (python-azure-skills/baseline) |
| identity-dp-python-managed-identity | 211.0s (python-azure-skills/baseline) | 262.8s | 358.7s (python-azure-skills/azure-skill-mcp) |
| key-vault-dp-python-secret-config | 229.3s (python-azure-skills/baseline) | 290.2s | 325.1s (python-azure-skills/azure-skill-mcp) |
| cosmos-db-dp-python-crud | 146.9s (python-azure-skills/baseline) | 175.2s | 226.3s (python-azure-skills/azure-skill-mcp) |
| identity-dp-python-service-principal | 201.2s (python-azure-skills/azure-skill-mcp-microsoft-skill) | 241.6s | 275.5s (python-azure-skills/azure-skill-mcp) |
| key-vault-dp-python-crud | 163.2s (python-azure-skills/baseline) | 191.0s | 225.7s (python-azure-skills/azure-skill-mcp) |
| service-bus-dp-python-order-processor | 316.9s (python-azure-skills/baseline) | 356.5s | 420.6s (python-azure-skills/azure-skill-mcp) |
| event-hubs-dp-python-streaming | 140.5s (python-azure-skills/baseline) | 205.7s | 243.3s (python-azure-skills/azure-skill-mcp-microsoft-skill) |
| storage-dp-python-blob-event-notifier | 311.6s (python-azure-skills/baseline) | 345.0s | 366.9s (python-azure-skills/azure-skill-mcp-microsoft-skill) |
| storage-mp-python-account-mgmt | 235.5s (python-azure-skills/azure-skill-mcp) | 429.1s | 814.7s (python-azure-skills/azure-skill-mcp-microsoft-skill) |
| identity-dp-python-credential-chain | 232.5s (python-azure-skills/baseline) | 276.0s | 321.4s (python-azure-skills/azure-skill-mcp) |
| service-bus-dp-python-crud | 151.5s (python-azure-skills/baseline) | 184.9s | 208.8s (python-azure-skills/azure-skill-mcp-microsoft-skill) |
| app-configuration-dp-python-crud | 119.9s (python-azure-skills/baseline) | 201.2s | 242.8s (python-azure-skills/azure-skill-mcp) |
| storage-dp-python-encrypted-uploader | 360.0s (python-azure-skills/azure-skill-mcp-microsoft-skill) | 396.4s | 447.5s (python-azure-skills/azure-skill-mcp) |
| app-configuration-dp-python-feature-flags | 321.5s (python-azure-skills/azure-skill-mcp-microsoft-skill) | 346.1s | 366.1s (python-azure-skills/baseline) |
| resource-manager-mp-python-rg-crud | 148.2s (python-azure-skills/baseline) | 276.7s | 418.2s (python-azure-skills/azure-skill-mcp) |

⏱ **Slowest:** storage-mp-python-account-mgmt/python-azure-skills/azure-skill-mcp-microsoft-skill · **Fastest:** app-configuration-dp-python-crud/python-azure-skills/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-python-crud | 3 | 0 | 3 | 0.0% |
| app-configuration-dp-python-feature-flags | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-python-crud | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-python-todo-repository | 3 | 0 | 3 | 0.0% |
| event-hubs-dp-python-streaming | 3 | 0 | 3 | 0.0% |
| identity-dp-python-credential-chain | 3 | 0 | 3 | 0.0% |
| identity-dp-python-default-credential | 3 | 0 | 3 | 0.0% |
| identity-dp-python-managed-identity | 3 | 0 | 3 | 0.0% |
| identity-dp-python-service-principal | 3 | 0 | 3 | 0.0% |
| key-vault-dp-python-crud | 3 | 0 | 3 | 0.0% |
| key-vault-dp-python-secret-config | 3 | 0 | 3 | 0.0% |
| resource-manager-mp-python-rg-crud | 3 | 0 | 3 | 0.0% |
| service-bus-dp-python-crud | 3 | 0 | 3 | 0.0% |
| service-bus-dp-python-order-processor | 3 | 0 | 3 | 0.0% |
| storage-dp-python-blob-event-notifier | 3 | 0 | 3 | 0.0% |
| storage-dp-python-blob-manager | 3 | 0 | 3 | 0.0% |
| storage-dp-python-crud | 3 | 0 | 3 | 0.0% |
| storage-dp-python-encrypted-uploader | 3 | 0 | 3 | 0.0% |
| storage-mp-python-account-mgmt | 3 | 0 | 3 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-skills/azure-skill-mcp | 19 | 0 | 19 | 0.0% |
| python-azure-skills/azure-skill-mcp-microsoft-skill | 19 | 0 | 19 | 0.0% |
| python-azure-skills/baseline | 19 | 0 | 19 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 207 | 207 | 0 | 100.0% |
| view | 115 | 115 | 0 | 100.0% |
| glob | 111 | 111 | 0 | 100.0% |
| apply_patch | 101 | 101 | 0 | 100.0% |
| azure-documentation | 83 | 83 | 0 | 100.0% |
| azure-get_azure_bestpractices | 72 | 72 | 0 | 100.0% |
| rg | 50 | 50 | 0 | 100.0% |
| skill | 40 | 39 | 1 | 97.5% |
| web_fetch | 33 | 33 | 0 | 100.0% |
| github-mcp-server-get_file_contents | 19 | 13 | 6 | 68.4% |
| web_search | 13 | 13 | 0 | 100.0% |
| github-mcp-server-search_code | 10 | 10 | 0 | 100.0% |
| azure-appconfig | 1 | 1 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### app-configuration-dp-python-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### app-configuration-dp-python-feature-flags

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-python-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-python-todo-repository

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-python-streaming

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-credential-chain

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-default-credential

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-managed-identity

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-service-principal

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-python-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-python-secret-config

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-python-rg-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-python-crud

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-python-order-processor

Baseline: **python-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-blob-event-notifier

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-blob-manager

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-encrypted-uploader

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-python-account-mgmt

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

