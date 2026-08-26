# Evaluation Summary: 20260826-143732

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-143732` |
| Timestamp | 2026-08-26T06:37:32Z |
| Total Prompts | 19 |
| Total Configs | 2 |
| Total Evaluations | 38 |
| Passed | 0 |
| Failed | 38 |
| Errors | 0 |
| Duration | 9941.8s |

## Comparison Matrix

| Prompt | python-azure-tools/baseline | python-azure-tools/with-azure-tools |
|--------|--------|--------|
| app-configuration-dp-python-crud | ❌ 9/13 | ❌ 11/13 |
| app-configuration-dp-python-feature-flags | ❌ 13/16 | ❌ 13/16 |
| cosmos-db-dp-python-crud | ❌ 8/13 | ❌ 9/13 |
| cosmos-db-dp-python-todo-repository | ❌ 15/20 | ❌ 15/20 |
| event-hubs-dp-python-streaming | ❌ 10/14 | ❌ 12/14 |
| identity-dp-python-credential-chain | ❌ 18/21 | ❌ 19/21 |
| identity-dp-python-default-credential | ❌ 7/12 | ❌ 10/12 |
| identity-dp-python-managed-identity | ❌ 8/13 | ❌ 10/13 |
| identity-dp-python-service-principal | ❌ 10/12 | ❌ 10/12 |
| key-vault-dp-python-crud | ❌ 10/12 | ❌ 10/12 |
| key-vault-dp-python-secret-config | ❌ 17/19 | ❌ 14/19 |
| resource-manager-mp-python-rg-crud | ❌ 12/14 | ❌ 11/14 |
| service-bus-dp-python-crud | ❌ 8/14 | ❌ 10/14 |
| service-bus-dp-python-order-processor | ❌ 16/21 | ❌ 16/21 |
| storage-dp-python-blob-event-notifier | ❌ 15/18 | ❌ 16/18 |
| storage-dp-python-blob-manager | ❌ 15/17 | ❌ 13/17 |
| storage-dp-python-crud | ❌ 13/15 | ❌ 13/15 |
| storage-dp-python-encrypted-uploader | ❌ 23/26 | ❌ 23/26 |
| storage-mp-python-account-mgmt | ❌ 8/15 | ❌ 13/15 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-python-crud](results/app-configuration/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 9/13 | 139.9s | 2 |
| [app-configuration-dp-python-crud](results/app-configuration/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 11/13 | 343.6s | 2 |
| [app-configuration-dp-python-feature-flags](results/app-configuration/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 13/16 | 303.8s | 9 |
| [app-configuration-dp-python-feature-flags](results/app-configuration/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 13/16 | 331.9s | 8 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 8/13 | 127.0s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 9/13 | 150.6s | 2 |
| [cosmos-db-dp-python-todo-repository](results/cosmos-db/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 15/20 | 256.4s | 6 |
| [cosmos-db-dp-python-todo-repository](results/cosmos-db/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 15/20 | 338.7s | 0 |
| [event-hubs-dp-python-streaming](results/event-hubs/data-plane/python/streaming/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 10/14 | 129.8s | 2 |
| [event-hubs-dp-python-streaming](results/event-hubs/data-plane/python/streaming/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 12/14 | 245.0s | 3 |
| [identity-dp-python-credential-chain](results/identity/data-plane/python/auth/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 18/21 | 186.9s | 5 |
| [identity-dp-python-credential-chain](results/identity/data-plane/python/auth/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 19/21 | 290.4s | 9 |
| [identity-dp-python-default-credential](results/identity/data-plane/python/auth/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 7/12 | 150.0s | 3 |
| [identity-dp-python-default-credential](results/identity/data-plane/python/auth/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 10/12 | 334.2s | 3 |
| [identity-dp-python-managed-identity](results/identity/data-plane/python/auth/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 8/13 | 241.6s | 19 |
| [identity-dp-python-managed-identity](results/identity/data-plane/python/auth/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 10/13 | 318.6s | 19 |
| [identity-dp-python-service-principal](results/identity/data-plane/python/auth/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 10/12 | 205.7s | 4 |
| [identity-dp-python-service-principal](results/identity/data-plane/python/auth/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 10/12 | 240.0s | 4 |
| [key-vault-dp-python-crud](results/key-vault/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 10/12 | 130.9s | 2 |
| [key-vault-dp-python-crud](results/key-vault/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 10/12 | 184.0s | 2 |
| [key-vault-dp-python-secret-config](results/key-vault/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 17/19 | 290.4s | 11 |
| [key-vault-dp-python-secret-config](results/key-vault/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 14/19 | 322.6s | 10 |
| [resource-manager-mp-python-rg-crud](results/resource-manager/management-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 12/14 | 143.5s | 2 |
| [resource-manager-mp-python-rg-crud](results/resource-manager/management-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 11/14 | 222.5s | 3 |
| [service-bus-dp-python-crud](results/service-bus/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 8/14 | 162.6s | 3 |
| [service-bus-dp-python-crud](results/service-bus/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 10/14 | 361.8s | 4 |
| [service-bus-dp-python-order-processor](results/service-bus/data-plane/python/streaming/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 16/21 | 276.0s | 7 |
| [service-bus-dp-python-order-processor](results/service-bus/data-plane/python/streaming/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 16/21 | 376.0s | 6 |
| [storage-dp-python-blob-event-notifier](results/storage/data-plane/python/streaming/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 15/18 | 401.3s | 6 |
| [storage-dp-python-blob-event-notifier](results/storage/data-plane/python/streaming/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 16/18 | 378.1s | 7 |
| [storage-dp-python-blob-manager](results/storage/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 15/17 | 242.9s | 4 |
| [storage-dp-python-blob-manager](results/storage/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 13/17 | 410.6s | 9 |
| [storage-dp-python-crud](results/storage/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 13/15 | 119.0s | 2 |
| [storage-dp-python-crud](results/storage/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 13/15 | 199.6s | 2 |
| [storage-dp-python-encrypted-uploader](results/storage/data-plane/python/crud/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 23/26 | 282.8s | 5 |
| [storage-dp-python-encrypted-uploader](results/storage/data-plane/python/crud/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 23/26 | 498.1s | 6 |
| [storage-mp-python-account-mgmt](results/storage/management-plane/python/provisioning/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 8/15 | 263.2s | 2 |
| [storage-mp-python-account-mgmt](results/storage/management-plane/python/provisioning/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 13/15 | 340.2s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| identity-dp-python-service-principal | 205.7s (python-azure-tools/baseline) | 222.9s | 240.0s (python-azure-tools/with-azure-tools) |
| cosmos-db-dp-python-crud | 127.0s (python-azure-tools/baseline) | 138.8s | 150.6s (python-azure-tools/with-azure-tools) |
| event-hubs-dp-python-streaming | 129.8s (python-azure-tools/baseline) | 187.4s | 245.0s (python-azure-tools/with-azure-tools) |
| identity-dp-python-credential-chain | 186.9s (python-azure-tools/baseline) | 238.6s | 290.4s (python-azure-tools/with-azure-tools) |
| key-vault-dp-python-secret-config | 290.4s (python-azure-tools/baseline) | 306.5s | 322.6s (python-azure-tools/with-azure-tools) |
| storage-dp-python-crud | 119.0s (python-azure-tools/baseline) | 159.3s | 199.6s (python-azure-tools/with-azure-tools) |
| storage-mp-python-account-mgmt | 263.2s (python-azure-tools/baseline) | 301.7s | 340.2s (python-azure-tools/with-azure-tools) |
| app-configuration-dp-python-crud | 139.9s (python-azure-tools/baseline) | 241.8s | 343.6s (python-azure-tools/with-azure-tools) |
| cosmos-db-dp-python-todo-repository | 256.4s (python-azure-tools/baseline) | 297.6s | 338.7s (python-azure-tools/with-azure-tools) |
| identity-dp-python-managed-identity | 241.6s (python-azure-tools/baseline) | 280.1s | 318.6s (python-azure-tools/with-azure-tools) |
| key-vault-dp-python-crud | 130.9s (python-azure-tools/baseline) | 157.4s | 184.0s (python-azure-tools/with-azure-tools) |
| storage-dp-python-blob-event-notifier | 378.1s (python-azure-tools/with-azure-tools) | 389.7s | 401.3s (python-azure-tools/baseline) |
| resource-manager-mp-python-rg-crud | 143.5s (python-azure-tools/baseline) | 183.0s | 222.5s (python-azure-tools/with-azure-tools) |
| service-bus-dp-python-order-processor | 276.0s (python-azure-tools/baseline) | 326.0s | 376.0s (python-azure-tools/with-azure-tools) |
| app-configuration-dp-python-feature-flags | 303.8s (python-azure-tools/baseline) | 317.9s | 331.9s (python-azure-tools/with-azure-tools) |
| service-bus-dp-python-crud | 162.6s (python-azure-tools/baseline) | 262.2s | 361.8s (python-azure-tools/with-azure-tools) |
| storage-dp-python-blob-manager | 242.9s (python-azure-tools/baseline) | 326.8s | 410.6s (python-azure-tools/with-azure-tools) |
| identity-dp-python-default-credential | 150.0s (python-azure-tools/baseline) | 242.1s | 334.2s (python-azure-tools/with-azure-tools) |
| storage-dp-python-encrypted-uploader | 282.8s (python-azure-tools/baseline) | 390.4s | 498.1s (python-azure-tools/with-azure-tools) |

⏱ **Slowest:** storage-dp-python-encrypted-uploader/python-azure-tools/with-azure-tools · **Fastest:** storage-dp-python-crud/python-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-python-crud | 2 | 0 | 2 | 0.0% |
| app-configuration-dp-python-feature-flags | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-python-crud | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-python-todo-repository | 2 | 0 | 2 | 0.0% |
| event-hubs-dp-python-streaming | 2 | 0 | 2 | 0.0% |
| identity-dp-python-credential-chain | 2 | 0 | 2 | 0.0% |
| identity-dp-python-default-credential | 2 | 0 | 2 | 0.0% |
| identity-dp-python-managed-identity | 2 | 0 | 2 | 0.0% |
| identity-dp-python-service-principal | 2 | 0 | 2 | 0.0% |
| key-vault-dp-python-crud | 2 | 0 | 2 | 0.0% |
| key-vault-dp-python-secret-config | 2 | 0 | 2 | 0.0% |
| resource-manager-mp-python-rg-crud | 2 | 0 | 2 | 0.0% |
| service-bus-dp-python-crud | 2 | 0 | 2 | 0.0% |
| service-bus-dp-python-order-processor | 2 | 0 | 2 | 0.0% |
| storage-dp-python-blob-event-notifier | 2 | 0 | 2 | 0.0% |
| storage-dp-python-blob-manager | 2 | 0 | 2 | 0.0% |
| storage-dp-python-crud | 2 | 0 | 2 | 0.0% |
| storage-dp-python-encrypted-uploader | 2 | 0 | 2 | 0.0% |
| storage-mp-python-account-mgmt | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-tools/baseline | 19 | 0 | 19 | 0.0% |
| python-azure-tools/with-azure-tools | 19 | 0 | 19 | 0.0% |

## Pairwise Details (per Prompt)

### app-configuration-dp-python-crud

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### app-configuration-dp-python-feature-flags

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-python-crud

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-python-todo-repository

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-python-streaming

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-credential-chain

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-default-credential

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-managed-identity

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-python-service-principal

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-python-crud

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-python-secret-config

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-python-rg-crud

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-python-crud

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-python-order-processor

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-blob-event-notifier

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-blob-manager

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-crud

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-python-encrypted-uploader

Baseline: **python-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-python-account-mgmt

Baseline: **python-azure-tools/baseline** — 34/100

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

