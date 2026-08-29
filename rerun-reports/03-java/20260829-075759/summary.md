# Evaluation Summary: 20260829-075759

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260829-075759` |
| Timestamp | 2026-08-28T23:57:59Z |
| Total Prompts | 19 |
| Total Configs | 3 |
| Total Evaluations | 57 |
| Passed | 2 |
| Failed | 55 |
| Errors | 0 |
| Duration | 22069.4s |

## Comparison Matrix

| Prompt | java-azure-skills/azure-skill-mcp | java-azure-skills/azure-skill-mcp-microsoft-skill | java-azure-skills/baseline |
|--------|--------|--------|--------|
| app-configuration-dp-java-crud | ❌ 15/19 | ❌ 15/19 | ❌ 15/19 |
| app-configuration-dp-java-feature-flags | ❌ 18/21 | ❌ 17/21 | ❌ 17/21 |
| cosmos-db-dp-java-crud | ❌ 14/19 | ❌ 15/19 | ❌ 14/19 |
| cosmos-db-dp-java-todo-repository | ❌ 21/26 | ❌ 23/26 | ❌ 21/26 |
| event-hubs-dp-java-streaming | ❌ 14/19 | ❌ 14/19 | ❌ 14/19 |
| identity-dp-java-credential-chain | ❌ 24/26 | ❌ 24/26 | ❌ 24/26 |
| identity-dp-java-default-credential | ❌ 15/17 | ❌ 15/17 | ✅ 17/17 |
| identity-dp-java-managed-identity | ❌ 16/18 | ❌ 15/18 | ❌ 15/18 |
| identity-dp-java-service-principal | ❌ 16/17 | ✅ 17/17 | ❌ 16/17 |
| key-vault-dp-java-crud | ❌ 16/17 | ❌ 16/17 | ❌ 16/17 |
| key-vault-dp-java-secret-config | ❌ 19/22 | ❌ 20/22 | ❌ 19/22 |
| resource-manager-mp-java-rg-crud | ❌ 17/19 | ❌ 16/19 | ❌ 15/19 |
| service-bus-dp-java-crud | ❌ 16/19 | ❌ 16/19 | ❌ 14/19 |
| service-bus-dp-java-order-processor | ❌ 19/24 | ❌ 18/24 | ❌ 20/24 |
| storage-dp-java-blob-event-notifier | ❌ 21/22 | ❌ 20/22 | ❌ 18/22 |
| storage-dp-java-blob-manager | ❌ 15/19 | ❌ 17/19 | ❌ 1/19 |
| storage-dp-java-crud | ❌ 18/19 | ❌ 17/19 | ❌ 16/19 |
| storage-dp-java-encrypted-uploader | ❌ 31/32 | ❌ 31/32 | ❌ 31/32 |
| storage-mp-java-account-mgmt | ❌ 16/20 | ❌ 0/13 | ❌ 1/20 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-java-crud](results/app-configuration/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 15/19 | 339.2s | 3 |
| [app-configuration-dp-java-crud](results/app-configuration/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 15/19 | 290.3s | 2 |
| [app-configuration-dp-java-crud](results/app-configuration/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 15/19 | 380.4s | 2 |
| [app-configuration-dp-java-feature-flags](results/app-configuration/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 18/21 | 579.4s | 11 |
| [app-configuration-dp-java-feature-flags](results/app-configuration/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/21 | 432.0s | 14 |
| [app-configuration-dp-java-feature-flags](results/app-configuration/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 17/21 | 532.3s | 14 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 14/19 | 391.3s | 2 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 15/19 | 340.3s | 3 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 14/19 | 288.7s | 2 |
| [cosmos-db-dp-java-todo-repository](results/cosmos-db/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 21/26 | 417.7s | 8 |
| [cosmos-db-dp-java-todo-repository](results/cosmos-db/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 23/26 | 462.4s | 10 |
| [cosmos-db-dp-java-todo-repository](results/cosmos-db/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 21/26 | 400.8s | 10 |
| [event-hubs-dp-java-streaming](results/event-hubs/data-plane/java/streaming/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 14/19 | 379.8s | 3 |
| [event-hubs-dp-java-streaming](results/event-hubs/data-plane/java/streaming/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 14/19 | 339.2s | 3 |
| [event-hubs-dp-java-streaming](results/event-hubs/data-plane/java/streaming/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 14/19 | 313.9s | 2 |
| [identity-dp-java-credential-chain](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 24/26 | 432.0s | 10 |
| [identity-dp-java-credential-chain](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 24/26 | 456.2s | 15 |
| [identity-dp-java-credential-chain](results/identity/data-plane/java/auth/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 24/26 | 438.3s | 10 |
| [identity-dp-java-default-credential](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 15/17 | 351.7s | 0 |
| [identity-dp-java-default-credential](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 15/17 | 337.1s | 0 |
| [identity-dp-java-default-credential](results/identity/data-plane/java/auth/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ✅ | 17/17 | 414.6s | 0 |
| [identity-dp-java-managed-identity](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 16/18 | 306.8s | 0 |
| [identity-dp-java-managed-identity](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 15/18 | 393.1s | 0 |
| [identity-dp-java-managed-identity](results/identity/data-plane/java/auth/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 15/18 | 321.0s | 0 |
| [identity-dp-java-service-principal](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 16/17 | 443.5s | 0 |
| [identity-dp-java-service-principal](results/identity/data-plane/java/auth/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 17/17 | 386.9s | 0 |
| [identity-dp-java-service-principal](results/identity/data-plane/java/auth/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 16/17 | 330.9s | 0 |
| [key-vault-dp-java-crud](results/key-vault/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 16/17 | 355.1s | 3 |
| [key-vault-dp-java-crud](results/key-vault/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 16/17 | 341.3s | 3 |
| [key-vault-dp-java-crud](results/key-vault/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 16/17 | 296.6s | 3 |
| [key-vault-dp-java-secret-config](results/key-vault/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 19/22 | 483.3s | 12 |
| [key-vault-dp-java-secret-config](results/key-vault/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 20/22 | 386.2s | 11 |
| [key-vault-dp-java-secret-config](results/key-vault/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 19/22 | 424.2s | 13 |
| [resource-manager-mp-java-rg-crud](results/resource-manager/management-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 17/19 | 396.2s | 3 |
| [resource-manager-mp-java-rg-crud](results/resource-manager/management-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 16/19 | 378.4s | 3 |
| [resource-manager-mp-java-rg-crud](results/resource-manager/management-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 15/19 | 372.4s | 3 |
| [service-bus-dp-java-crud](results/service-bus/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 16/19 | 460.9s | 3 |
| [service-bus-dp-java-crud](results/service-bus/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 16/19 | 393.9s | 3 |
| [service-bus-dp-java-crud](results/service-bus/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 14/19 | 300.6s | 2 |
| [service-bus-dp-java-order-processor](results/service-bus/data-plane/java/streaming/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 19/24 | 494.6s | 12 |
| [service-bus-dp-java-order-processor](results/service-bus/data-plane/java/streaming/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 18/24 | 435.3s | 11 |
| [service-bus-dp-java-order-processor](results/service-bus/data-plane/java/streaming/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 20/24 | 425.1s | 11 |
| [storage-dp-java-blob-event-notifier](results/storage/data-plane/java/streaming/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 21/22 | 447.8s | 16 |
| [storage-dp-java-blob-event-notifier](results/storage/data-plane/java/streaming/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 20/22 | 404.3s | 17 |
| [storage-dp-java-blob-event-notifier](results/storage/data-plane/java/streaming/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 18/22 | 351.6s | 9 |
| [storage-dp-java-blob-manager](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 15/19 | 492.1s | 6 |
| [storage-dp-java-blob-manager](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/19 | 415.5s | 6 |
| [storage-dp-java-blob-manager](results/storage/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 1/19 | 225.4s | 0 |
| [storage-dp-java-crud](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 18/19 | 374.5s | 3 |
| [storage-dp-java-crud](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 17/19 | 297.9s | 4 |
| [storage-dp-java-crud](results/storage/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 16/19 | 232.9s | 0 |
| [storage-dp-java-encrypted-uploader](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 31/32 | 491.7s | 13 |
| [storage-dp-java-encrypted-uploader](results/storage/data-plane/java/crud/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 31/32 | 423.8s | 15 |
| [storage-dp-java-encrypted-uploader](results/storage/data-plane/java/crud/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 31/32 | 519.2s | 14 |
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 16/20 | 426.3s | 2 |
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | java-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 0/13 | 173.4s | 0 |
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-skills/baseline/report.md) | java-azure-skills/baseline | ❌ | 1/20 | 346.5s | 0 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| storage-dp-java-encrypted-uploader | 423.8s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 478.3s | 519.2s (java-azure-skills/baseline) |
| identity-dp-java-credential-chain | 432.0s (java-azure-skills/azure-skill-mcp) | 442.2s | 456.2s (java-azure-skills/azure-skill-mcp-microsoft-skill) |
| service-bus-dp-java-crud | 300.6s (java-azure-skills/baseline) | 385.1s | 460.9s (java-azure-skills/azure-skill-mcp) |
| event-hubs-dp-java-streaming | 313.9s (java-azure-skills/baseline) | 344.3s | 379.8s (java-azure-skills/azure-skill-mcp) |
| app-configuration-dp-java-crud | 290.3s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 336.6s | 380.4s (java-azure-skills/baseline) |
| storage-dp-java-crud | 232.9s (java-azure-skills/baseline) | 301.8s | 374.5s (java-azure-skills/azure-skill-mcp) |
| cosmos-db-dp-java-todo-repository | 400.8s (java-azure-skills/baseline) | 427.0s | 462.4s (java-azure-skills/azure-skill-mcp-microsoft-skill) |
| service-bus-dp-java-order-processor | 425.1s (java-azure-skills/baseline) | 451.7s | 494.6s (java-azure-skills/azure-skill-mcp) |
| cosmos-db-dp-java-crud | 288.7s (java-azure-skills/baseline) | 340.1s | 391.3s (java-azure-skills/azure-skill-mcp) |
| identity-dp-java-default-credential | 337.1s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 367.8s | 414.6s (java-azure-skills/baseline) |
| resource-manager-mp-java-rg-crud | 372.4s (java-azure-skills/baseline) | 382.3s | 396.2s (java-azure-skills/azure-skill-mcp) |
| identity-dp-java-managed-identity | 306.8s (java-azure-skills/azure-skill-mcp) | 340.3s | 393.1s (java-azure-skills/azure-skill-mcp-microsoft-skill) |
| key-vault-dp-java-secret-config | 386.2s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 431.2s | 483.3s (java-azure-skills/azure-skill-mcp) |
| app-configuration-dp-java-feature-flags | 432.0s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 514.6s | 579.4s (java-azure-skills/azure-skill-mcp) |
| storage-dp-java-blob-manager | 225.4s (java-azure-skills/baseline) | 377.7s | 492.1s (java-azure-skills/azure-skill-mcp) |
| identity-dp-java-service-principal | 330.9s (java-azure-skills/baseline) | 387.1s | 443.5s (java-azure-skills/azure-skill-mcp) |
| key-vault-dp-java-crud | 296.6s (java-azure-skills/baseline) | 331.0s | 355.1s (java-azure-skills/azure-skill-mcp) |
| storage-dp-java-blob-event-notifier | 351.6s (java-azure-skills/baseline) | 401.2s | 447.8s (java-azure-skills/azure-skill-mcp) |
| storage-mp-java-account-mgmt | 173.4s (java-azure-skills/azure-skill-mcp-microsoft-skill) | 315.4s | 426.3s (java-azure-skills/azure-skill-mcp) |

⏱ **Slowest:** app-configuration-dp-java-feature-flags/java-azure-skills/azure-skill-mcp · **Fastest:** storage-mp-java-account-mgmt/java-azure-skills/azure-skill-mcp-microsoft-skill

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-java-crud | 3 | 0 | 3 | 0.0% |
| app-configuration-dp-java-feature-flags | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-java-crud | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-java-todo-repository | 3 | 0 | 3 | 0.0% |
| event-hubs-dp-java-streaming | 3 | 0 | 3 | 0.0% |
| identity-dp-java-credential-chain | 3 | 0 | 3 | 0.0% |
| identity-dp-java-default-credential | 3 | 1 | 2 | 33.3% |
| identity-dp-java-managed-identity | 3 | 0 | 3 | 0.0% |
| identity-dp-java-service-principal | 3 | 1 | 2 | 33.3% |
| key-vault-dp-java-crud | 3 | 0 | 3 | 0.0% |
| key-vault-dp-java-secret-config | 3 | 0 | 3 | 0.0% |
| resource-manager-mp-java-rg-crud | 3 | 0 | 3 | 0.0% |
| service-bus-dp-java-crud | 3 | 0 | 3 | 0.0% |
| service-bus-dp-java-order-processor | 3 | 0 | 3 | 0.0% |
| storage-dp-java-blob-event-notifier | 3 | 0 | 3 | 0.0% |
| storage-dp-java-blob-manager | 3 | 0 | 3 | 0.0% |
| storage-dp-java-crud | 3 | 0 | 3 | 0.0% |
| storage-dp-java-encrypted-uploader | 3 | 0 | 3 | 0.0% |
| storage-mp-java-account-mgmt | 3 | 0 | 3 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| java-azure-skills/azure-skill-mcp | 19 | 0 | 19 | 0.0% |
| java-azure-skills/azure-skill-mcp-microsoft-skill | 19 | 1 | 18 | 5.3% |
| java-azure-skills/baseline | 19 | 1 | 18 | 5.3% |

## Prompt Deltas

| Prompt | Passes On | Fails On |
|--------|-----------|----------|
| identity-dp-java-default-credential | java-azure-skills/baseline | java-azure-skills/azure-skill-mcp |
| identity-dp-java-default-credential | java-azure-skills/baseline | java-azure-skills/azure-skill-mcp-microsoft-skill |
| identity-dp-java-service-principal | java-azure-skills/azure-skill-mcp-microsoft-skill | java-azure-skills/azure-skill-mcp |
| identity-dp-java-service-principal | java-azure-skills/azure-skill-mcp-microsoft-skill | java-azure-skills/baseline |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 163 | 163 | 0 | 100.0% |
| web_fetch | 111 | 109 | 2 | 98.2% |
| azure-documentation | 103 | 103 | 0 | 100.0% |
| apply_patch | 92 | 92 | 0 | 100.0% |
| view | 89 | 84 | 5 | 94.4% |
| glob | 86 | 86 | 0 | 100.0% |
| azure-get_azure_bestpractices | 74 | 74 | 0 | 100.0% |
| rg | 51 | 51 | 0 | 100.0% |
| skill | 37 | 35 | 2 | 94.6% |
| web_search | 37 | 37 | 0 | 100.0% |
| github-mcp-server-search_code | 26 | 26 | 0 | 100.0% |
| github-mcp-server-get_file_contents | 21 | 20 | 1 | 95.2% |
| task | 4 | 4 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### app-configuration-dp-java-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### app-configuration-dp-java-feature-flags

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-java-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-java-todo-repository

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-java-streaming

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-credential-chain

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-default-credential

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-managed-identity

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-service-principal

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-java-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-java-secret-config

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-java-rg-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-java-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-java-order-processor

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-blob-event-notifier

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-blob-manager

Baseline: **java-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-crud

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-encrypted-uploader

Baseline: **java-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-java-account-mgmt

Baseline: **java-azure-skills/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

