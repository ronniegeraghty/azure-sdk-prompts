# Evaluation Summary: 20260829-004156

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260829-004156` |
| Timestamp | 2026-08-28T16:41:56Z |
| Total Prompts | 20 |
| Total Configs | 3 |
| Total Evaluations | 60 |
| Passed | 17 |
| Failed | 42 |
| Errors | 1 |
| Duration | 9310.5s |

## Comparison Matrix

| Prompt | dotnet-azure-skills/azure-skill-mcp | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | dotnet-azure-skills/baseline |
|--------|--------|--------|--------|
| app-configuration-dp-dotnet-crud | ❌ 6/7 | ❌ 6/7 | ❌ 6/7 |
| cosmos-db-dp-dotnet-crud | ✅ 7/7 | ✅ 7/7 | ✅ 7/7 |
| cosmos-db-dp-dotnet-error-handling | ❌ 6/8 | ❌ 6/8 | ❌ 7/8 |
| cosmos-db-dp-dotnet-pagination | ❌ 7/8 | ✅ 8/8 | ❌ 7/8 |
| event-hubs-dp-dotnet-streaming | ✅ 7/7 | ✅ 7/7 | ✅ 7/7 |
| identity-dp-dotnet-default-credential | ❌ 4/5 | ✅ 5/5 | ✅ 5/5 |
| identity-dp-dotnet-managed-identity | ❌ 4/6 | ❌ 0/1 | ❌ 5/6 |
| identity-dp-dotnet-service-principal | ✅ 5/5 | ✅ 5/5 | ✅ 5/5 |
| key-vault-dp-dotnet-crud | ✅ 5/5 | ❌ 4/5 | ✅ 5/5 |
| key-vault-dp-dotnet-error-handling | ❌ 4/7 | ❌ 4/7 | ❌ 4/7 |
| key-vault-dp-dotnet-pagination | ❌ 6/7 | ❌ 6/7 | ❌ 5/7 |
| key-vault-mp-dotnet-polling | ❌ 0/8 | ❌ 7/8 | ❌ 7/8 |
| resource-manager-mp-dotnet-rg-crud | ✅ 6/6 | ✅ 6/6 | ✅ 6/6 |
| service-bus-dp-dotnet-crud | ❌ 7/8 | ❌ 7/8 | ❌ 7/8 |
| storage-dp-dotnet-auth | ❌ 3/5 | ❌ 3/5 | ❌ 3/5 |
| storage-dp-dotnet-batch | ❌ 5/8 | ❌ 5/8 | ❌ 5/8 |
| storage-dp-dotnet-error-handling | ❌ 4/6 | ❌ 4/6 | ❌ 4/6 |
| storage-dp-dotnet-retries | ❌ 7/8 | ❌ 7/8 | ❌ 7/8 |
| storage-mp-dotnet-account-mgmt | ❌ 6/7 | ❌ 6/7 | ❌ 6/7 |
| storage-mp-dotnet-polling | ❌ 6/7 | ❌ 6/7 | ❌ 6/7 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [app-configuration-dp-dotnet-crud](results/app-configuration/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 6/7 | 163.2s | 2 |
| [app-configuration-dp-dotnet-crud](results/app-configuration/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 6/7 | 198.1s | 3 |
| [app-configuration-dp-dotnet-crud](results/app-configuration/data-plane/dotnet/crud/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 6/7 | 81.2s | 2 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ✅ | 7/7 | 200.3s | 3 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 7/7 | 166.9s | 3 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 7/7 | 87.9s | 3 |
| [cosmos-db-dp-dotnet-error-handling](results/cosmos-db/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 6/8 | 173.5s | 0 |
| [cosmos-db-dp-dotnet-error-handling](results/cosmos-db/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 6/8 | 143.5s | 0 |
| [cosmos-db-dp-dotnet-error-handling](results/cosmos-db/data-plane/dotnet/error-handling/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 7/8 | 146.9s | 0 |
| [cosmos-db-dp-dotnet-pagination](results/cosmos-db/data-plane/dotnet/pagination/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 7/8 | 206.3s | 3 |
| [cosmos-db-dp-dotnet-pagination](results/cosmos-db/data-plane/dotnet/pagination/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 8/8 | 187.3s | 3 |
| [cosmos-db-dp-dotnet-pagination](results/cosmos-db/data-plane/dotnet/pagination/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 7/8 | 104.3s | 3 |
| [event-hubs-dp-dotnet-streaming](results/event-hubs/data-plane/dotnet/streaming/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ✅ | 7/7 | 201.7s | 3 |
| [event-hubs-dp-dotnet-streaming](results/event-hubs/data-plane/dotnet/streaming/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 7/7 | 155.7s | 3 |
| [event-hubs-dp-dotnet-streaming](results/event-hubs/data-plane/dotnet/streaming/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 7/7 | 103.6s | 2 |
| [identity-dp-dotnet-default-credential](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 4/5 | 160.5s | 0 |
| [identity-dp-dotnet-default-credential](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 5/5 | 146.6s | 0 |
| [identity-dp-dotnet-default-credential](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 5/5 | 132.1s | 0 |
| [identity-dp-dotnet-managed-identity](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 4/6 | 168.6s | 0 |
| [identity-dp-dotnet-managed-identity](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 0/1 | 618.6s | 0 |
| [identity-dp-dotnet-managed-identity](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 5/6 | 137.5s | 0 |
| [identity-dp-dotnet-service-principal](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ✅ | 5/5 | 119.4s | 0 |
| [identity-dp-dotnet-service-principal](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 5/5 | 152.9s | 0 |
| [identity-dp-dotnet-service-principal](results/identity/data-plane/dotnet/auth/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 5/5 | 58.1s | 0 |
| [key-vault-dp-dotnet-crud](results/key-vault/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ✅ | 5/5 | 136.8s | 2 |
| [key-vault-dp-dotnet-crud](results/key-vault/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 4/5 | 182.1s | 3 |
| [key-vault-dp-dotnet-crud](results/key-vault/data-plane/dotnet/crud/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 5/5 | 100.3s | 3 |
| [key-vault-dp-dotnet-error-handling](results/key-vault/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 4/7 | 134.4s | 0 |
| [key-vault-dp-dotnet-error-handling](results/key-vault/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 4/7 | 162.1s | 0 |
| [key-vault-dp-dotnet-error-handling](results/key-vault/data-plane/dotnet/error-handling/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 4/7 | 61.3s | 0 |
| [key-vault-dp-dotnet-pagination](results/key-vault/data-plane/dotnet/pagination/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 6/7 | 229.2s | 3 |
| [key-vault-dp-dotnet-pagination](results/key-vault/data-plane/dotnet/pagination/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 6/7 | 149.4s | 2 |
| [key-vault-dp-dotnet-pagination](results/key-vault/data-plane/dotnet/pagination/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 5/7 | 85.6s | 3 |
| [key-vault-mp-dotnet-polling](results/key-vault/management-plane/dotnet/polling/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 0/8 | 206.0s | 2 |
| [key-vault-mp-dotnet-polling](results/key-vault/management-plane/dotnet/polling/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 7/8 | 169.0s | 3 |
| [key-vault-mp-dotnet-polling](results/key-vault/management-plane/dotnet/polling/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 7/8 | 225.1s | 3 |
| [resource-manager-mp-dotnet-rg-crud](results/resource-manager/management-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ✅ | 6/6 | 111.0s | 3 |
| [resource-manager-mp-dotnet-rg-crud](results/resource-manager/management-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ✅ | 6/6 | 119.2s | 3 |
| [resource-manager-mp-dotnet-rg-crud](results/resource-manager/management-plane/dotnet/crud/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ✅ | 6/6 | 133.4s | 3 |
| [service-bus-dp-dotnet-crud](results/service-bus/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 7/8 | 160.7s | 3 |
| [service-bus-dp-dotnet-crud](results/service-bus/data-plane/dotnet/crud/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 7/8 | 90.6s | 3 |
| [service-bus-dp-dotnet-crud](results/service-bus/data-plane/dotnet/crud/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 7/8 | 56.5s | 2 |
| [storage-dp-dotnet-auth](results/storage/data-plane/dotnet/authentication/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 3/5 | 119.1s | 0 |
| [storage-dp-dotnet-auth](results/storage/data-plane/dotnet/authentication/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 3/5 | 101.6s | 0 |
| [storage-dp-dotnet-auth](results/storage/data-plane/dotnet/authentication/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 3/5 | 97.2s | 0 |
| [storage-dp-dotnet-batch](results/storage/data-plane/dotnet/batch/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 5/8 | 141.6s | 0 |
| [storage-dp-dotnet-batch](results/storage/data-plane/dotnet/batch/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 5/8 | 201.0s | 0 |
| [storage-dp-dotnet-batch](results/storage/data-plane/dotnet/batch/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 5/8 | 111.8s | 0 |
| [storage-dp-dotnet-error-handling](results/storage/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 4/6 | 85.5s | 0 |
| [storage-dp-dotnet-error-handling](results/storage/data-plane/dotnet/error-handling/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 4/6 | 122.0s | 0 |
| [storage-dp-dotnet-error-handling](results/storage/data-plane/dotnet/error-handling/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 4/6 | 46.9s | 0 |
| [storage-dp-dotnet-retries](results/storage/data-plane/dotnet/retries/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 7/8 | 125.7s | 0 |
| [storage-dp-dotnet-retries](results/storage/data-plane/dotnet/retries/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 7/8 | 193.9s | 0 |
| [storage-dp-dotnet-retries](results/storage/data-plane/dotnet/retries/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 7/8 | 206.9s | 0 |
| [storage-mp-dotnet-account-mgmt](results/storage/management-plane/dotnet/provisioning/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 6/7 | 216.3s | 2 |
| [storage-mp-dotnet-account-mgmt](results/storage/management-plane/dotnet/provisioning/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 6/7 | 223.4s | 2 |
| [storage-mp-dotnet-account-mgmt](results/storage/management-plane/dotnet/provisioning/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 6/7 | 265.7s | 3 |
| [storage-mp-dotnet-polling](results/storage/management-plane/dotnet/polling/dotnet-azure-skills/azure-skill-mcp/report.md) | dotnet-azure-skills/azure-skill-mcp | ❌ | 6/7 | 190.2s | 3 |
| [storage-mp-dotnet-polling](results/storage/management-plane/dotnet/polling/dotnet-azure-skills/azure-skill-mcp-microsoft-skill/report.md) | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | ❌ | 6/7 | 248.7s | 2 |
| [storage-mp-dotnet-polling](results/storage/management-plane/dotnet/polling/dotnet-azure-skills/baseline/report.md) | dotnet-azure-skills/baseline | ❌ | 6/7 | 83.5s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| app-configuration-dp-dotnet-crud | 81.2s (dotnet-azure-skills/baseline) | 147.5s | 198.1s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| identity-dp-dotnet-default-credential | 132.1s (dotnet-azure-skills/baseline) | 146.4s | 160.5s (dotnet-azure-skills/azure-skill-mcp) |
| storage-mp-dotnet-account-mgmt | 216.3s (dotnet-azure-skills/azure-skill-mcp) | 235.2s | 265.7s (dotnet-azure-skills/baseline) |
| storage-dp-dotnet-auth | 97.2s (dotnet-azure-skills/baseline) | 106.0s | 119.1s (dotnet-azure-skills/azure-skill-mcp) |
| key-vault-dp-dotnet-error-handling | 61.3s (dotnet-azure-skills/baseline) | 119.3s | 162.1s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| storage-dp-dotnet-batch | 111.8s (dotnet-azure-skills/baseline) | 151.5s | 201.0s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| service-bus-dp-dotnet-crud | 56.5s (dotnet-azure-skills/baseline) | 102.6s | 160.7s (dotnet-azure-skills/azure-skill-mcp) |
| key-vault-dp-dotnet-crud | 100.3s (dotnet-azure-skills/baseline) | 139.7s | 182.1s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| storage-dp-dotnet-error-handling | 46.9s (dotnet-azure-skills/baseline) | 84.8s | 122.0s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| resource-manager-mp-dotnet-rg-crud | 111.0s (dotnet-azure-skills/azure-skill-mcp) | 121.2s | 133.4s (dotnet-azure-skills/baseline) |
| event-hubs-dp-dotnet-streaming | 103.6s (dotnet-azure-skills/baseline) | 153.7s | 201.7s (dotnet-azure-skills/azure-skill-mcp) |
| identity-dp-dotnet-managed-identity | 137.5s (dotnet-azure-skills/baseline) | 308.2s | 618.6s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| cosmos-db-dp-dotnet-pagination | 104.3s (dotnet-azure-skills/baseline) | 166.0s | 206.3s (dotnet-azure-skills/azure-skill-mcp) |
| cosmos-db-dp-dotnet-crud | 87.9s (dotnet-azure-skills/baseline) | 151.7s | 200.3s (dotnet-azure-skills/azure-skill-mcp) |
| cosmos-db-dp-dotnet-error-handling | 143.5s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) | 154.6s | 173.5s (dotnet-azure-skills/azure-skill-mcp) |
| identity-dp-dotnet-service-principal | 58.1s (dotnet-azure-skills/baseline) | 110.1s | 152.9s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| key-vault-mp-dotnet-polling | 169.0s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) | 200.0s | 225.1s (dotnet-azure-skills/baseline) |
| storage-mp-dotnet-polling | 83.5s (dotnet-azure-skills/baseline) | 174.1s | 248.7s (dotnet-azure-skills/azure-skill-mcp-microsoft-skill) |
| key-vault-dp-dotnet-pagination | 85.6s (dotnet-azure-skills/baseline) | 154.7s | 229.2s (dotnet-azure-skills/azure-skill-mcp) |
| storage-dp-dotnet-retries | 125.7s (dotnet-azure-skills/azure-skill-mcp) | 175.5s | 206.9s (dotnet-azure-skills/baseline) |

⏱ **Slowest:** identity-dp-dotnet-managed-identity/dotnet-azure-skills/azure-skill-mcp-microsoft-skill · **Fastest:** storage-dp-dotnet-error-handling/dotnet-azure-skills/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| app-configuration-dp-dotnet-crud | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-dotnet-crud | 3 | 3 | 0 | 100.0% |
| cosmos-db-dp-dotnet-error-handling | 3 | 0 | 3 | 0.0% |
| cosmos-db-dp-dotnet-pagination | 3 | 1 | 2 | 33.3% |
| event-hubs-dp-dotnet-streaming | 3 | 3 | 0 | 100.0% |
| identity-dp-dotnet-default-credential | 3 | 2 | 1 | 66.7% |
| identity-dp-dotnet-managed-identity | 3 | 0 | 3 | 0.0% |
| identity-dp-dotnet-service-principal | 3 | 3 | 0 | 100.0% |
| key-vault-dp-dotnet-crud | 3 | 2 | 1 | 66.7% |
| key-vault-dp-dotnet-error-handling | 3 | 0 | 3 | 0.0% |
| key-vault-dp-dotnet-pagination | 3 | 0 | 3 | 0.0% |
| key-vault-mp-dotnet-polling | 3 | 0 | 3 | 0.0% |
| resource-manager-mp-dotnet-rg-crud | 3 | 3 | 0 | 100.0% |
| service-bus-dp-dotnet-crud | 3 | 0 | 3 | 0.0% |
| storage-dp-dotnet-auth | 3 | 0 | 3 | 0.0% |
| storage-dp-dotnet-batch | 3 | 0 | 3 | 0.0% |
| storage-dp-dotnet-error-handling | 3 | 0 | 3 | 0.0% |
| storage-dp-dotnet-retries | 3 | 0 | 3 | 0.0% |
| storage-mp-dotnet-account-mgmt | 3 | 0 | 3 | 0.0% |
| storage-mp-dotnet-polling | 3 | 0 | 3 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| dotnet-azure-skills/azure-skill-mcp | 20 | 5 | 15 | 25.0% |
| dotnet-azure-skills/azure-skill-mcp-microsoft-skill | 20 | 6 | 14 | 30.0% |
| dotnet-azure-skills/baseline | 20 | 6 | 14 | 30.0% |

## Prompt Deltas

| Prompt | Passes On | Fails On |
|--------|-----------|----------|
| cosmos-db-dp-dotnet-pagination | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | dotnet-azure-skills/azure-skill-mcp |
| cosmos-db-dp-dotnet-pagination | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | dotnet-azure-skills/baseline |
| identity-dp-dotnet-default-credential | dotnet-azure-skills/azure-skill-mcp-microsoft-skill | dotnet-azure-skills/azure-skill-mcp |
| identity-dp-dotnet-default-credential | dotnet-azure-skills/baseline | dotnet-azure-skills/azure-skill-mcp |
| key-vault-dp-dotnet-crud | dotnet-azure-skills/azure-skill-mcp | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| key-vault-dp-dotnet-crud | dotnet-azure-skills/baseline | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| azure-documentation | 182 | 182 | 0 | 100.0% |
| powershell | 154 | 154 | 0 | 100.0% |
| view | 95 | 91 | 4 | 95.8% |
| web_fetch | 82 | 42 | 40 | 51.2% |
| glob | 82 | 82 | 0 | 100.0% |
| azure-get_azure_bestpractices | 78 | 78 | 0 | 100.0% |
| rg | 56 | 56 | 0 | 100.0% |
| apply_patch | 53 | 53 | 0 | 100.0% |
| skill | 32 | 31 | 1 | 96.9% |
| web_search | 26 | 26 | 0 | 100.0% |
| github-mcp-server-get_file_contents | 17 | 16 | 1 | 94.1% |
| github-mcp-server-search_code | 15 | 15 | 0 | 100.0% |
| azure-appconfig | 1 | 1 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### app-configuration-dp-dotnet-crud

Baseline: **dotnet-azure-skills/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-crud

Baseline: **dotnet-azure-skills/baseline** — 7/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-error-handling

Baseline: **dotnet-azure-skills/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### cosmos-db-dp-dotnet-pagination

Baseline: **dotnet-azure-skills/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### event-hubs-dp-dotnet-streaming

Baseline: **dotnet-azure-skills/baseline** — 7/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-default-credential

Baseline: **dotnet-azure-skills/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-managed-identity

Baseline: **dotnet-azure-skills/baseline** — 5/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### identity-dp-dotnet-service-principal

Baseline: **dotnet-azure-skills/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-crud

Baseline: **dotnet-azure-skills/baseline** — 5/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-error-handling

Baseline: **dotnet-azure-skills/baseline** — 4/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-dp-dotnet-pagination

Baseline: **dotnet-azure-skills/baseline** — 5/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### key-vault-mp-dotnet-polling

Baseline: **dotnet-azure-skills/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### resource-manager-mp-dotnet-rg-crud

Baseline: **dotnet-azure-skills/baseline** — 6/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### service-bus-dp-dotnet-crud

Baseline: **dotnet-azure-skills/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-auth

Baseline: **dotnet-azure-skills/baseline** — 3/5

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-batch

Baseline: **dotnet-azure-skills/baseline** — 5/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-error-handling

Baseline: **dotnet-azure-skills/baseline** — 4/6

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-dp-dotnet-retries

Baseline: **dotnet-azure-skills/baseline** — 7/8

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-dotnet-account-mgmt

Baseline: **dotnet-azure-skills/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### storage-mp-dotnet-polling

Baseline: **dotnet-azure-skills/baseline** — 6/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

