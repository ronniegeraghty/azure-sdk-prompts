# Evaluation Summary: 20260826-173422

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-173422` |
| Timestamp | 2026-08-26T09:34:22Z |
| Total Prompts | 14 |
| Total Configs | 2 |
| Total Evaluations | 28 |
| Passed | 0 |
| Failed | 27 |
| Errors | 1 |
| Duration | 8640.4s |

## Comparison Matrix

| Prompt | js-ts-azure-tools/baseline | js-ts-azure-tools/with-azure-tools |
|--------|--------|--------|
| app-configuration-dp-js-ts-crud | ❌ 12/18 | ❌ 13/18 |
| cosmos-db-dp-js-ts-crud | ❌ 10/17 | ❌ 12/17 |
| event-hubs-dp-js-ts-streaming | ❌ 14/18 | ❌ 13/18 |
| identity-dp-js-ts-default-credential | ❌ 12/15 | ❌ 12/15 |
| identity-dp-js-ts-managed-identity | ❌ 11/16 | ❌ 13/16 |
| identity-dp-js-ts-service-principal | ❌ 12/15 | ❌ 12/15 |
| key-vault-dp-js-ts-crud | ❌ 12/15 | ❌ 12/15 |
| key-vault-dp-js-ts-secret-config | ❌ 20/23 | ❌ 21/23 |
| resource-manager-mp-js-ts-rg-crud | ❌ 14/18 | ❌ 14/18 |
| service-bus-dp-js-ts-crud | ❌ 12/18 | ❌ 13/18 |
| storage-dp-js-ts-blob-manager | ❌ 15/22 | ❌ 19/22 |
| storage-dp-js-ts-crud | ❌ 16/18 | ❌ 15/18 |
| storage-dp-js-ts-encrypted-uploader | ❌ 30/35 | ❌ 31/35 |
| storage-mp-js-ts-account-mgmt | ❌ 13/18 | ❌ 13/18 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-js-ts-crud](results/app-configuration/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 12/18 | 282.8s | 4 |
| [app-configuration-dp-js-ts-crud](results/app-configuration/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 13/18 | 279.0s | 5 |
| [cosmos-db-dp-js-ts-crud](results/cosmos-db/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 10/17 | 251.1s | 4 |
| [cosmos-db-dp-js-ts-crud](results/cosmos-db/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 12/17 | 317.7s | 5 |
| [event-hubs-dp-js-ts-streaming](results/event-hubs/data-plane/js-ts/streaming/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 14/18 | 281.9s | 4 |
| [event-hubs-dp-js-ts-streaming](results/event-hubs/data-plane/js-ts/streaming/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 13/18 | 786.3s | 4 |
| [identity-dp-js-ts-default-credential](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 12/15 | 209.9s | 4 |
| [identity-dp-js-ts-default-credential](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 12/15 | 286.3s | 4 |
| [identity-dp-js-ts-managed-identity](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 11/16 | 211.8s | 4 |
| [identity-dp-js-ts-managed-identity](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 13/16 | 256.6s | 4 |
| [identity-dp-js-ts-service-principal](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 12/15 | 261.1s | 4 |
| [identity-dp-js-ts-service-principal](results/identity/data-plane/js-ts/auth/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 12/15 | 248.8s | 4 |
| [key-vault-dp-js-ts-crud](results/key-vault/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 12/15 | 207.9s | 4 |
| [key-vault-dp-js-ts-crud](results/key-vault/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 12/15 | 308.8s | 5 |
| [key-vault-dp-js-ts-secret-config](results/key-vault/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 20/23 | 321.9s | 10 |
| [key-vault-dp-js-ts-secret-config](results/key-vault/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 21/23 | 441.2s | 12 |
| [resource-manager-mp-js-ts-rg-crud](results/resource-manager/management-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 14/18 | 222.8s | 5 |
| [resource-manager-mp-js-ts-rg-crud](results/resource-manager/management-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 14/18 | 279.3s | 5 |
| [service-bus-dp-js-ts-crud](results/service-bus/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 12/18 | 204.8s | 5 |
| [service-bus-dp-js-ts-crud](results/service-bus/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 13/18 | 264.1s | 5 |
| [storage-dp-js-ts-blob-manager](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 15/22 | 269.0s | 7 |
| [storage-dp-js-ts-blob-manager](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 19/22 | 430.3s | 7 |
| [storage-dp-js-ts-crud](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 16/18 | 310.7s | 4 |
| [storage-dp-js-ts-crud](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 15/18 | 385.0s | 5 |
| [storage-dp-js-ts-encrypted-uploader](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 30/35 | 311.1s | 7 |
| [storage-dp-js-ts-encrypted-uploader](results/storage/data-plane/js-ts/crud/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 31/35 | 408.2s | 9 |
| [storage-mp-js-ts-account-mgmt](results/storage/management-plane/js-ts/provisioning/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 13/18 | 242.5s | 4 |
| [storage-mp-js-ts-account-mgmt](results/storage/management-plane/js-ts/provisioning/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 13/18 | 309.5s | 4 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-js-ts-crud | 251.1s (js-ts-azure-tools/baseline) | 284.4s | 317.7s (js-ts-azure-tools/with-azure-tools) |
| event-hubs-dp-js-ts-streaming | 281.9s (js-ts-azure-tools/baseline) | 534.1s | 786.3s (js-ts-azure-tools/with-azure-tools) |
| identity-dp-js-ts-managed-identity | 211.8s (js-ts-azure-tools/baseline) | 234.2s | 256.6s (js-ts-azure-tools/with-azure-tools) |
| key-vault-dp-js-ts-crud | 207.9s (js-ts-azure-tools/baseline) | 258.4s | 308.8s (js-ts-azure-tools/with-azure-tools) |
| key-vault-dp-js-ts-secret-config | 321.9s (js-ts-azure-tools/baseline) | 381.5s | 441.2s (js-ts-azure-tools/with-azure-tools) |
| storage-dp-js-ts-crud | 310.7s (js-ts-azure-tools/baseline) | 347.9s | 385.0s (js-ts-azure-tools/with-azure-tools) |
| app-configuration-dp-js-ts-crud | 279.0s (js-ts-azure-tools/with-azure-tools) | 280.9s | 282.8s (js-ts-azure-tools/baseline) |
| resource-manager-mp-js-ts-rg-crud | 222.8s (js-ts-azure-tools/baseline) | 251.1s | 279.3s (js-ts-azure-tools/with-azure-tools) |
| storage-dp-js-ts-blob-manager | 269.0s (js-ts-azure-tools/baseline) | 349.6s | 430.3s (js-ts-azure-tools/with-azure-tools) |
| storage-dp-js-ts-encrypted-uploader | 311.1s (js-ts-azure-tools/baseline) | 359.7s | 408.2s (js-ts-azure-tools/with-azure-tools) |
| storage-mp-js-ts-account-mgmt | 242.5s (js-ts-azure-tools/baseline) | 276.0s | 309.5s (js-ts-azure-tools/with-azure-tools) |
| identity-dp-js-ts-default-credential | 209.9s (js-ts-azure-tools/baseline) | 248.1s | 286.3s (js-ts-azure-tools/with-azure-tools) |
| identity-dp-js-ts-service-principal | 248.8s (js-ts-azure-tools/with-azure-tools) | 254.9s | 261.1s (js-ts-azure-tools/baseline) |
| service-bus-dp-js-ts-crud | 204.8s (js-ts-azure-tools/baseline) | 234.5s | 264.1s (js-ts-azure-tools/with-azure-tools) |

⏱ **Slowest:** event-hubs-dp-js-ts-streaming/js-ts-azure-tools/with-azure-tools · **Fastest:** service-bus-dp-js-ts-crud/js-ts-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |
| event-hubs-dp-js-ts-streaming | 2 | 0 | 2 | 0.0% |
| identity-dp-js-ts-default-credential | 2 | 0 | 2 | 0.0% |
| identity-dp-js-ts-managed-identity | 2 | 0 | 2 | 0.0% |
| identity-dp-js-ts-service-principal | 2 | 0 | 2 | 0.0% |
| key-vault-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |
| key-vault-dp-js-ts-secret-config | 2 | 0 | 2 | 0.0% |
| resource-manager-mp-js-ts-rg-crud | 2 | 0 | 2 | 0.0% |
| service-bus-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |
| storage-dp-js-ts-blob-manager | 2 | 0 | 2 | 0.0% |
| storage-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |
| storage-dp-js-ts-encrypted-uploader | 2 | 0 | 2 | 0.0% |
| storage-mp-js-ts-account-mgmt | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| js-ts-azure-tools/baseline | 14 | 0 | 14 | 0.0% |
| js-ts-azure-tools/with-azure-tools | 14 | 0 | 14 | 0.0% |

## Pairwise Details (per Prompt)

### app-configuration-dp-js-ts-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-js-ts-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-js-ts-streaming

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-js-ts-default-credential

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-js-ts-managed-identity

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-js-ts-service-principal

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-js-ts-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-js-ts-secret-config

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-js-ts-rg-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-js-ts-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-js-ts-blob-manager

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-js-ts-crud

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-js-ts-encrypted-uploader

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-js-ts-account-mgmt

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

