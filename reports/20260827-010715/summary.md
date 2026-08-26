# Evaluation Summary: 20260827-010715

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260827-010715` |
| Timestamp | 2026-08-26T17:07:15Z |
| Total Prompts | 20 |
| Total Configs | 2 |
| Total Evaluations | 40 |
| Passed | 11 |
| Failed | 29 |
| Errors | 0 |
| Duration | 5798.9s |

## Comparison Matrix

| Prompt | dotnet-azure-tools/baseline | dotnet-azure-tools/with-azure-tools |
|--------|--------|--------|
| app-configuration-dp-dotnet-crud | ❌ 6/7 | ❌ 6/7 |
| cosmos-db-dp-dotnet-crud | ✅ 7/7 | ✅ 7/7 |
| cosmos-db-dp-dotnet-error-handling | ❌ 6/8 | ❌ 6/8 |
| cosmos-db-dp-dotnet-pagination | ❌ 7/8 | ❌ 7/8 |
| event-hubs-dp-dotnet-streaming | ✅ 7/7 | ✅ 7/7 |
| identity-dp-dotnet-default-credential | ✅ 5/5 | ✅ 5/5 |
| identity-dp-dotnet-managed-identity | ❌ 5/6 | ❌ 4/6 |
| identity-dp-dotnet-service-principal | ✅ 5/5 | ✅ 5/5 |
| key-vault-dp-dotnet-crud | ✅ 5/5 | ✅ 5/5 |
| key-vault-dp-dotnet-error-handling | ❌ 5/7 | ❌ 5/7 |
| key-vault-dp-dotnet-pagination | ❌ 5/7 | ❌ 5/7 |
| key-vault-mp-dotnet-polling | ❌ 7/8 | ❌ 6/8 |
| resource-manager-mp-dotnet-rg-crud | ❌ 2/6 | ✅ 6/6 |
| service-bus-dp-dotnet-crud | ❌ 7/8 | ❌ 7/8 |
| storage-dp-dotnet-auth | ❌ 3/5 | ❌ 3/5 |
| storage-dp-dotnet-batch | ❌ 5/8 | ❌ 5/8 |
| storage-dp-dotnet-error-handling | ❌ 4/6 | ❌ 4/6 |
| storage-dp-dotnet-retries | ❌ 6/8 | ❌ 7/8 |
| storage-mp-dotnet-account-mgmt | ❌ 6/7 | ❌ 6/7 |
| storage-mp-dotnet-polling | ❌ 6/7 | ❌ 6/7 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-dotnet-crud](results/app-configuration/data-plane/dotnet/crud/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 6/7 | 122.0s | 3 |
| [app-configuration-dp-dotnet-crud](results/app-configuration/data-plane/dotnet/crud/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 6/7 | 167.8s | 2 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 7/7 | 103.6s | 2 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 7/7 | 164.5s | 2 |
| [cosmos-db-dp-dotnet-error-handling](results/cosmos-db/data-plane/dotnet/error-handling/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 6/8 | 147.0s | 0 |
| [cosmos-db-dp-dotnet-error-handling](results/cosmos-db/data-plane/dotnet/error-handling/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 6/8 | 138.7s | 0 |
| [cosmos-db-dp-dotnet-pagination](results/cosmos-db/data-plane/dotnet/pagination/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 7/8 | 101.1s | 3 |
| [cosmos-db-dp-dotnet-pagination](results/cosmos-db/data-plane/dotnet/pagination/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 7/8 | 152.0s | 3 |
| [event-hubs-dp-dotnet-streaming](results/event-hubs/data-plane/dotnet/streaming/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 7/7 | 99.4s | 3 |
| [event-hubs-dp-dotnet-streaming](results/event-hubs/data-plane/dotnet/streaming/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 7/7 | 147.0s | 3 |
| [identity-dp-dotnet-default-credential](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 5/5 | 197.8s | 0 |
| [identity-dp-dotnet-default-credential](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 5/5 | 152.3s | 0 |
| [identity-dp-dotnet-managed-identity](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 5/6 | 201.2s | 0 |
| [identity-dp-dotnet-managed-identity](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 4/6 | 153.0s | 0 |
| [identity-dp-dotnet-service-principal](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 5/5 | 54.5s | 0 |
| [identity-dp-dotnet-service-principal](results/identity/data-plane/dotnet/auth/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 5/5 | 149.1s | 0 |
| [key-vault-dp-dotnet-crud](results/key-vault/data-plane/dotnet/crud/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 5/5 | 86.5s | 3 |
| [key-vault-dp-dotnet-crud](results/key-vault/data-plane/dotnet/crud/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 5/5 | 170.6s | 3 |
| [key-vault-dp-dotnet-error-handling](results/key-vault/data-plane/dotnet/error-handling/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 5/7 | 166.9s | 0 |
| [key-vault-dp-dotnet-error-handling](results/key-vault/data-plane/dotnet/error-handling/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 5/7 | 131.8s | 0 |
| [key-vault-dp-dotnet-pagination](results/key-vault/data-plane/dotnet/pagination/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 5/7 | 113.4s | 3 |
| [key-vault-dp-dotnet-pagination](results/key-vault/data-plane/dotnet/pagination/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 5/7 | 179.5s | 3 |
| [key-vault-mp-dotnet-polling](results/key-vault/management-plane/dotnet/polling/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 7/8 | 191.7s | 3 |
| [key-vault-mp-dotnet-polling](results/key-vault/management-plane/dotnet/polling/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 6/8 | 244.5s | 3 |
| [resource-manager-mp-dotnet-rg-crud](results/resource-manager/management-plane/dotnet/crud/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 2/6 | 120.4s | 3 |
| [resource-manager-mp-dotnet-rg-crud](results/resource-manager/management-plane/dotnet/crud/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 6/6 | 183.0s | 2 |
| [service-bus-dp-dotnet-crud](results/service-bus/data-plane/dotnet/crud/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 7/8 | 87.1s | 3 |
| [service-bus-dp-dotnet-crud](results/service-bus/data-plane/dotnet/crud/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 7/8 | 158.7s | 3 |
| [storage-dp-dotnet-auth](results/storage/data-plane/dotnet/authentication/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 3/5 | 46.5s | 0 |
| [storage-dp-dotnet-auth](results/storage/data-plane/dotnet/authentication/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 3/5 | 152.9s | 0 |
| [storage-dp-dotnet-batch](results/storage/data-plane/dotnet/batch/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 5/8 | 124.2s | 0 |
| [storage-dp-dotnet-batch](results/storage/data-plane/dotnet/batch/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 5/8 | 211.0s | 0 |
| [storage-dp-dotnet-error-handling](results/storage/data-plane/dotnet/error-handling/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 4/6 | 57.1s | 0 |
| [storage-dp-dotnet-error-handling](results/storage/data-plane/dotnet/error-handling/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 4/6 | 177.6s | 0 |
| [storage-dp-dotnet-retries](results/storage/data-plane/dotnet/retries/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 6/8 | 171.1s | 0 |
| [storage-dp-dotnet-retries](results/storage/data-plane/dotnet/retries/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 7/8 | 183.2s | 0 |
| [storage-mp-dotnet-account-mgmt](results/storage/management-plane/dotnet/provisioning/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 6/7 | 132.7s | 2 |
| [storage-mp-dotnet-account-mgmt](results/storage/management-plane/dotnet/provisioning/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 6/7 | 171.7s | 3 |
| [storage-mp-dotnet-polling](results/storage/management-plane/dotnet/polling/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ❌ | 6/7 | 103.6s | 3 |
| [storage-mp-dotnet-polling](results/storage/management-plane/dotnet/polling/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ❌ | 6/7 | 181.4s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-dotnet-crud | 103.6s (dotnet-azure-tools/baseline) | 134.0s | 164.5s (dotnet-azure-tools/with-azure-tools) |
| cosmos-db-dp-dotnet-error-handling | 138.7s (dotnet-azure-tools/with-azure-tools) | 142.9s | 147.0s (dotnet-azure-tools/baseline) |
| storage-mp-dotnet-polling | 103.6s (dotnet-azure-tools/baseline) | 142.5s | 181.4s (dotnet-azure-tools/with-azure-tools) |
| key-vault-mp-dotnet-polling | 191.7s (dotnet-azure-tools/baseline) | 218.1s | 244.5s (dotnet-azure-tools/with-azure-tools) |
| identity-dp-dotnet-service-principal | 54.5s (dotnet-azure-tools/baseline) | 101.8s | 149.1s (dotnet-azure-tools/with-azure-tools) |
| identity-dp-dotnet-managed-identity | 153.0s (dotnet-azure-tools/with-azure-tools) | 177.1s | 201.2s (dotnet-azure-tools/baseline) |
| service-bus-dp-dotnet-crud | 87.1s (dotnet-azure-tools/baseline) | 122.9s | 158.7s (dotnet-azure-tools/with-azure-tools) |
| storage-dp-dotnet-batch | 124.2s (dotnet-azure-tools/baseline) | 167.6s | 211.0s (dotnet-azure-tools/with-azure-tools) |
| key-vault-dp-dotnet-crud | 86.5s (dotnet-azure-tools/baseline) | 128.5s | 170.6s (dotnet-azure-tools/with-azure-tools) |
| cosmos-db-dp-dotnet-pagination | 101.1s (dotnet-azure-tools/baseline) | 126.6s | 152.0s (dotnet-azure-tools/with-azure-tools) |
| key-vault-dp-dotnet-error-handling | 131.8s (dotnet-azure-tools/with-azure-tools) | 149.3s | 166.9s (dotnet-azure-tools/baseline) |
| storage-mp-dotnet-account-mgmt | 132.7s (dotnet-azure-tools/baseline) | 152.2s | 171.7s (dotnet-azure-tools/with-azure-tools) |
| identity-dp-dotnet-default-credential | 152.3s (dotnet-azure-tools/with-azure-tools) | 175.0s | 197.8s (dotnet-azure-tools/baseline) |
| event-hubs-dp-dotnet-streaming | 99.4s (dotnet-azure-tools/baseline) | 123.2s | 147.0s (dotnet-azure-tools/with-azure-tools) |
| key-vault-dp-dotnet-pagination | 113.4s (dotnet-azure-tools/baseline) | 146.4s | 179.5s (dotnet-azure-tools/with-azure-tools) |
| storage-dp-dotnet-error-handling | 57.1s (dotnet-azure-tools/baseline) | 117.4s | 177.6s (dotnet-azure-tools/with-azure-tools) |
| resource-manager-mp-dotnet-rg-crud | 120.4s (dotnet-azure-tools/baseline) | 151.7s | 183.0s (dotnet-azure-tools/with-azure-tools) |
| storage-dp-dotnet-auth | 46.5s (dotnet-azure-tools/baseline) | 99.7s | 152.9s (dotnet-azure-tools/with-azure-tools) |
| storage-dp-dotnet-retries | 171.1s (dotnet-azure-tools/baseline) | 177.2s | 183.2s (dotnet-azure-tools/with-azure-tools) |
| app-configuration-dp-dotnet-crud | 122.0s (dotnet-azure-tools/baseline) | 144.9s | 167.8s (dotnet-azure-tools/with-azure-tools) |

⏱ **Slowest:** key-vault-mp-dotnet-polling/dotnet-azure-tools/with-azure-tools · **Fastest:** storage-dp-dotnet-auth/dotnet-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-dotnet-crud | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-dotnet-crud | 2 | 2 | 0 | 100.0% |
| cosmos-db-dp-dotnet-error-handling | 2 | 0 | 2 | 0.0% |
| cosmos-db-dp-dotnet-pagination | 2 | 0 | 2 | 0.0% |
| event-hubs-dp-dotnet-streaming | 2 | 2 | 0 | 100.0% |
| identity-dp-dotnet-default-credential | 2 | 2 | 0 | 100.0% |
| identity-dp-dotnet-managed-identity | 2 | 0 | 2 | 0.0% |
| identity-dp-dotnet-service-principal | 2 | 2 | 0 | 100.0% |
| key-vault-dp-dotnet-crud | 2 | 2 | 0 | 100.0% |
| key-vault-dp-dotnet-error-handling | 2 | 0 | 2 | 0.0% |
| key-vault-dp-dotnet-pagination | 2 | 0 | 2 | 0.0% |
| key-vault-mp-dotnet-polling | 2 | 0 | 2 | 0.0% |
| resource-manager-mp-dotnet-rg-crud | 2 | 1 | 1 | 50.0% |
| service-bus-dp-dotnet-crud | 2 | 0 | 2 | 0.0% |
| storage-dp-dotnet-auth | 2 | 0 | 2 | 0.0% |
| storage-dp-dotnet-batch | 2 | 0 | 2 | 0.0% |
| storage-dp-dotnet-error-handling | 2 | 0 | 2 | 0.0% |
| storage-dp-dotnet-retries | 2 | 0 | 2 | 0.0% |
| storage-mp-dotnet-account-mgmt | 2 | 0 | 2 | 0.0% |
| storage-mp-dotnet-polling | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| dotnet-azure-tools/baseline | 20 | 5 | 15 | 25.0% |
| dotnet-azure-tools/with-azure-tools | 20 | 6 | 14 | 30.0% |

## Prompt Deltas

| Prompt | Passes On | Fails On |
|--------|-----------|----------|
| resource-manager-mp-dotnet-rg-crud | dotnet-azure-tools/with-azure-tools | dotnet-azure-tools/baseline |

## Pairwise Details (per Prompt)

### app-configuration-dp-dotnet-crud

Baseline: **dotnet-azure-tools/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-crud

Baseline: **dotnet-azure-tools/baseline** — 7/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-error-handling

Baseline: **dotnet-azure-tools/baseline** — 6/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-pagination

Baseline: **dotnet-azure-tools/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-dotnet-streaming

Baseline: **dotnet-azure-tools/baseline** — 7/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-default-credential

Baseline: **dotnet-azure-tools/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-managed-identity

Baseline: **dotnet-azure-tools/baseline** — 5/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-service-principal

Baseline: **dotnet-azure-tools/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-crud

Baseline: **dotnet-azure-tools/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-error-handling

Baseline: **dotnet-azure-tools/baseline** — 5/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-pagination

Baseline: **dotnet-azure-tools/baseline** — 5/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-mp-dotnet-polling

Baseline: **dotnet-azure-tools/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-dotnet-rg-crud

Baseline: **dotnet-azure-tools/baseline** — 2/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-dotnet-crud

Baseline: **dotnet-azure-tools/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-auth

Baseline: **dotnet-azure-tools/baseline** — 3/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-batch

Baseline: **dotnet-azure-tools/baseline** — 5/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-error-handling

Baseline: **dotnet-azure-tools/baseline** — 4/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-retries

Baseline: **dotnet-azure-tools/baseline** — 6/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-dotnet-account-mgmt

Baseline: **dotnet-azure-tools/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-dotnet-polling

Baseline: **dotnet-azure-tools/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

