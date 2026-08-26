# Evaluation Summary: 20260826-111051

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-111051` |
| Timestamp | 2026-08-26T03:10:51Z |
| Total Prompts | 1 |
| Total Configs | 2 |
| Total Evaluations | 2 |
| Passed | 2 |
| Failed | 0 |
| Errors | 0 |
| Duration | 441.9s |

## Comparison Matrix

| Prompt | dotnet-cosmos-skill/baseline | dotnet-cosmos-skill/without-azure-sdk-dotnet |
|--------|--------|--------|
| cosmos-db-dp-dotnet-crud | ✅ 7/7 | ✅ 7/7 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-cosmos-skill/baseline/report.md) | dotnet-cosmos-skill/baseline | ✅ | 7/7 | 148.3s | 3 |
| [cosmos-db-dp-dotnet-crud](results/cosmos-db/data-plane/dotnet/crud/dotnet-cosmos-skill/without-azure-sdk-dotnet/report.md) | dotnet-cosmos-skill/without-azure-sdk-dotnet | ✅ | 7/7 | 293.4s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-dotnet-crud | 148.3s (dotnet-cosmos-skill/baseline) | 220.9s | 293.4s (dotnet-cosmos-skill/without-azure-sdk-dotnet) |

⏱ **Slowest:** cosmos-db-dp-dotnet-crud/dotnet-cosmos-skill/without-azure-sdk-dotnet · **Fastest:** cosmos-db-dp-dotnet-crud/dotnet-cosmos-skill/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| cosmos-db-dp-dotnet-crud | 2 | 2 | 0 | 100.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| dotnet-cosmos-skill/baseline | 1 | 1 | 0 | 100.0% |
| dotnet-cosmos-skill/without-azure-sdk-dotnet | 1 | 1 | 0 | 100.0% |

## Pairwise Tool Impact

Impact = baseline score − score without tool (positive = tool helps).

| Tool | Impact | Baseline | Without | Baseline Pass | Without Pass |
|------|--------|----------|---------|---------------|-------------|
| azure-sdk-dotnet | 0.0 | 100.0 | 100.0 | ✅ | ✅ |

## Pairwise Details (per Prompt)

### cosmos-db-dp-dotnet-crud

Baseline: **dotnet-cosmos-skill/baseline** — 7/7

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|
| azure-sdk-dotnet | 0.0 | 100.0 | ✅ |

