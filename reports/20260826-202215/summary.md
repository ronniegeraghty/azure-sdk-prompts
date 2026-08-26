# Evaluation Summary: 20260826-202215

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-202215` |
| Timestamp | 2026-08-26T12:22:15Z |
| Total Prompts | 19 |
| Total Configs | 2 |
| Total Evaluations | 38 |
| Passed | 3 |
| Failed | 35 |
| Errors | 0 |
| Duration | 16888.4s |

## Comparison Matrix

| Prompt | java-azure-tools/baseline | java-azure-tools/with-azure-tools |
|--------|--------|--------|
| app-configuration-dp-java-crud | ❌ 15/19 | ❌ 15/19 |
| app-configuration-dp-java-feature-flags | ❌ 18/21 | ❌ 17/21 |
| cosmos-db-dp-java-crud | ❌ 14/19 | ❌ 14/19 |
| cosmos-db-dp-java-todo-repository | ❌ 17/26 | ❌ 21/26 |
| event-hubs-dp-java-streaming | ❌ 14/19 | ❌ 14/19 |
| identity-dp-java-credential-chain | ❌ 23/26 | ❌ 24/26 |
| identity-dp-java-default-credential | ✅ 17/17 | ❌ 16/17 |
| identity-dp-java-managed-identity | ❌ 16/18 | ❌ 16/18 |
| identity-dp-java-service-principal | ✅ 17/17 | ✅ 17/17 |
| key-vault-dp-java-crud | ❌ 16/17 | ❌ 16/17 |
| key-vault-dp-java-secret-config | ❌ 19/22 | ❌ 20/22 |
| resource-manager-mp-java-rg-crud | ❌ 16/19 | ❌ 15/19 |
| service-bus-dp-java-crud | ❌ 14/19 | ❌ 16/19 |
| service-bus-dp-java-order-processor | ❌ 20/24 | ❌ 19/24 |
| storage-dp-java-blob-event-notifier | ❌ 16/22 | ❌ 21/22 |
| storage-dp-java-blob-manager | ❌ 17/19 | ❌ 17/19 |
| storage-dp-java-crud | ❌ 16/19 | ❌ 17/19 |
| storage-dp-java-encrypted-uploader | ❌ 31/32 | ❌ 31/32 |
| storage-mp-java-account-mgmt | ❌ 17/20 | ❌ 17/20 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-java-crud](results/app-configuration/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 15/19 | 323.2s | 2 |
| [app-configuration-dp-java-crud](results/app-configuration/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 15/19 | 336.7s | 2 |
| [app-configuration-dp-java-feature-flags](results/app-configuration/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 18/21 | 646.0s | 11 |
| [app-configuration-dp-java-feature-flags](results/app-configuration/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 17/21 | 591.0s | 10 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 14/19 | 358.8s | 3 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 14/19 | 373.5s | 3 |
| [cosmos-db-dp-java-todo-repository](results/cosmos-db/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 17/26 | 321.6s | 9 |
| [cosmos-db-dp-java-todo-repository](results/cosmos-db/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 21/26 | 514.2s | 10 |
| [event-hubs-dp-java-streaming](results/event-hubs/data-plane/java/streaming/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 14/19 | 432.3s | 2 |
| [event-hubs-dp-java-streaming](results/event-hubs/data-plane/java/streaming/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 14/19 | 389.4s | 3 |
| [identity-dp-java-credential-chain](results/identity/data-plane/java/auth/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 23/26 | 567.1s | 14 |
| [identity-dp-java-credential-chain](results/identity/data-plane/java/auth/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 24/26 | 540.9s | 14 |
| [identity-dp-java-default-credential](results/identity/data-plane/java/auth/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ✅ | 17/17 | 362.4s | 0 |
| [identity-dp-java-default-credential](results/identity/data-plane/java/auth/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 16/17 | 410.6s | 0 |
| [identity-dp-java-managed-identity](results/identity/data-plane/java/auth/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 16/18 | 331.0s | 0 |
| [identity-dp-java-managed-identity](results/identity/data-plane/java/auth/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 16/18 | 338.6s | 0 |
| [identity-dp-java-service-principal](results/identity/data-plane/java/auth/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ✅ | 17/17 | 377.0s | 0 |
| [identity-dp-java-service-principal](results/identity/data-plane/java/auth/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ✅ | 17/17 | 360.8s | 0 |
| [key-vault-dp-java-crud](results/key-vault/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 16/17 | 332.1s | 2 |
| [key-vault-dp-java-crud](results/key-vault/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 16/17 | 437.6s | 3 |
| [key-vault-dp-java-secret-config](results/key-vault/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 19/22 | 450.1s | 13 |
| [key-vault-dp-java-secret-config](results/key-vault/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 20/22 | 597.4s | 14 |
| [resource-manager-mp-java-rg-crud](results/resource-manager/management-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 16/19 | 532.4s | 3 |
| [resource-manager-mp-java-rg-crud](results/resource-manager/management-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 15/19 | 509.8s | 3 |
| [service-bus-dp-java-crud](results/service-bus/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 14/19 | 374.8s | 3 |
| [service-bus-dp-java-crud](results/service-bus/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 16/19 | 345.5s | 3 |
| [service-bus-dp-java-order-processor](results/service-bus/data-plane/java/streaming/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 20/24 | 498.9s | 10 |
| [service-bus-dp-java-order-processor](results/service-bus/data-plane/java/streaming/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 19/24 | 602.2s | 15 |
| [storage-dp-java-blob-event-notifier](results/storage/data-plane/java/streaming/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 16/22 | 501.9s | 25 |
| [storage-dp-java-blob-event-notifier](results/storage/data-plane/java/streaming/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 21/22 | 542.1s | 17 |
| [storage-dp-java-blob-manager](results/storage/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 17/19 | 483.6s | 7 |
| [storage-dp-java-blob-manager](results/storage/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 17/19 | 505.5s | 7 |
| [storage-dp-java-crud](results/storage/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 16/19 | 312.8s | 2 |
| [storage-dp-java-crud](results/storage/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 17/19 | 364.8s | 4 |
| [storage-dp-java-encrypted-uploader](results/storage/data-plane/java/crud/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 31/32 | 412.2s | 12 |
| [storage-dp-java-encrypted-uploader](results/storage/data-plane/java/crud/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 31/32 | 552.4s | 16 |
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 17/20 | 443.7s | 2 |
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 17/20 | 511.6s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| service-bus-dp-java-crud | 345.5s (java-azure-tools/with-azure-tools) | 360.1s | 374.8s (java-azure-tools/baseline) |
| cosmos-db-dp-java-crud | 358.8s (java-azure-tools/baseline) | 366.2s | 373.5s (java-azure-tools/with-azure-tools) |
| service-bus-dp-java-order-processor | 498.9s (java-azure-tools/baseline) | 550.5s | 602.2s (java-azure-tools/with-azure-tools) |
| identity-dp-java-service-principal | 360.8s (java-azure-tools/with-azure-tools) | 368.9s | 377.0s (java-azure-tools/baseline) |
| app-configuration-dp-java-feature-flags | 591.0s (java-azure-tools/with-azure-tools) | 618.5s | 646.0s (java-azure-tools/baseline) |
| key-vault-dp-java-secret-config | 450.1s (java-azure-tools/baseline) | 523.7s | 597.4s (java-azure-tools/with-azure-tools) |
| storage-dp-java-blob-event-notifier | 501.9s (java-azure-tools/baseline) | 522.0s | 542.1s (java-azure-tools/with-azure-tools) |
| event-hubs-dp-java-streaming | 389.4s (java-azure-tools/with-azure-tools) | 410.9s | 432.3s (java-azure-tools/baseline) |
| identity-dp-java-managed-identity | 331.0s (java-azure-tools/baseline) | 334.8s | 338.6s (java-azure-tools/with-azure-tools) |
| key-vault-dp-java-crud | 332.1s (java-azure-tools/baseline) | 384.9s | 437.6s (java-azure-tools/with-azure-tools) |
| app-configuration-dp-java-crud | 323.2s (java-azure-tools/baseline) | 330.0s | 336.7s (java-azure-tools/with-azure-tools) |
| resource-manager-mp-java-rg-crud | 509.8s (java-azure-tools/with-azure-tools) | 521.1s | 532.4s (java-azure-tools/baseline) |
| identity-dp-java-credential-chain | 540.9s (java-azure-tools/with-azure-tools) | 554.0s | 567.1s (java-azure-tools/baseline) |
| storage-dp-java-blob-manager | 483.6s (java-azure-tools/baseline) | 494.5s | 505.5s (java-azure-tools/with-azure-tools) |
| storage-dp-java-crud | 312.8s (java-azure-tools/baseline) | 338.8s | 364.8s (java-azure-tools/with-azure-tools) |
| cosmos-db-dp-java-todo-repository | 321.6s (java-azure-tools/baseline) | 417.9s | 514.2s (java-azure-tools/with-azure-tools) |
| storage-dp-java-encrypted-uploader | 412.2s (java-azure-tools/baseline) | 482.3s | 552.4s (java-azure-tools/with-azure-tools) |
| storage-mp-java-account-mgmt | 443.7s (java-azure-tools/baseline) | 477.7s | 511.6s (java-azure-tools/with-azure-tools) |
| identity-dp-java-default-credential | 362.4s (java-azure-tools/baseline) | 386.5s | 410.6s (java-azure-tools/with-azure-tools) |

⏱ **Slowest:** app-configuration-dp-java-feature-flags/java-azure-tools/baseline · **Fastest:** storage-dp-java-crud/java-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-java-crud | 2 | 0 | 2 | 0.0% |
| app-configuration-dp-java-feature-flags | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-java-crud | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-java-todo-repository | 2 | 0 | 2 | 0.0% |
| event-hubs-dp-java-streaming | 2 | 0 | 2 | 0.0% |
| identity-dp-java-credential-chain | 2 | 0 | 2 | 0.0% |
| identity-dp-java-default-credential | 2 | 1 | 1 | 50.0% |
| identity-dp-java-managed-identity | 2 | 0 | 2 | 0.0% |
| identity-dp-java-service-principal | 2 | 2 | 0 | 100.0% |
| key-vault-dp-java-crud | 2 | 0 | 2 | 0.0% |
| key-vault-dp-java-secret-config | 2 | 0 | 2 | 0.0% |
| resource-manager-mp-java-rg-crud | 2 | 0 | 2 | 0.0% |
| service-bus-dp-java-crud | 2 | 0 | 2 | 0.0% |
| service-bus-dp-java-order-processor | 2 | 0 | 2 | 0.0% |
| storage-dp-java-blob-event-notifier | 2 | 0 | 2 | 0.0% |
| storage-dp-java-blob-manager | 2 | 0 | 2 | 0.0% |
| storage-dp-java-crud | 2 | 0 | 2 | 0.0% |
| storage-dp-java-encrypted-uploader | 2 | 0 | 2 | 0.0% |
| storage-mp-java-account-mgmt | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| java-azure-tools/baseline | 19 | 2 | 17 | 10.5% |
| java-azure-tools/with-azure-tools | 19 | 1 | 18 | 5.3% |

## Prompt Deltas

| Prompt | Passes On | Fails On |
|--------|-----------|----------|
| identity-dp-java-default-credential | java-azure-tools/baseline | java-azure-tools/with-azure-tools |

## Pairwise Details (per Prompt)

### app-configuration-dp-java-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### app-configuration-dp-java-feature-flags

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-java-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-java-todo-repository

Baseline: **java-azure-tools/baseline** — 52/100

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-java-streaming

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-credential-chain

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-default-credential

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-managed-identity

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-java-service-principal

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-java-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-java-secret-config

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-java-rg-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-java-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-java-order-processor

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-blob-event-notifier

Baseline: **java-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-blob-manager

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-crud

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-java-encrypted-uploader

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-java-account-mgmt

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

