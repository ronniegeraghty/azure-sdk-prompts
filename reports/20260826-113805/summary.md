# Evaluation Summary: 20260826-113805

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-113805` |
| Timestamp | 2026-08-26T03:38:05Z |
| Total Prompts | 1 |
| Total Configs | 2 |
| Total Evaluations | 2 |
| Passed | 0 |
| Failed | 2 |
| Errors | 0 |
| Duration | 1176.3s |

## Comparison Matrix

| Prompt | js-ts-cosmos-skill/baseline | js-ts-cosmos-skill/without-azure-sdk-typescript |
|--------|--------|--------|
| cosmos-db-dp-js-ts-crud | ❌ 12/17 | ❌ 10/17 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [cosmos-db-dp-js-ts-crud](results/cosmos-db/data-plane/js-ts/crud/js-ts-cosmos-skill/baseline/report.md) | js-ts-cosmos-skill/baseline | ❌ | 12/17 | 584.2s | 4 |
| [cosmos-db-dp-js-ts-crud](results/cosmos-db/data-plane/js-ts/crud/js-ts-cosmos-skill/without-azure-sdk-typescript/report.md) | js-ts-cosmos-skill/without-azure-sdk-typescript | ❌ | 10/17 | 588.9s | 4 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-js-ts-crud | 584.2s (js-ts-cosmos-skill/baseline) | 586.5s | 588.9s (js-ts-cosmos-skill/without-azure-sdk-typescript) |

⏱ **Slowest:** cosmos-db-dp-js-ts-crud/js-ts-cosmos-skill/without-azure-sdk-typescript · **Fastest:** cosmos-db-dp-js-ts-crud/js-ts-cosmos-skill/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| cosmos-db-dp-js-ts-crud | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| js-ts-cosmos-skill/baseline | 1 | 0 | 1 | 0.0% |
| js-ts-cosmos-skill/without-azure-sdk-typescript | 1 | 0 | 1 | 0.0% |

## Pairwise Tool Impact

Impact = baseline score − score without tool (positive = tool helps).

| Tool | Impact | Baseline | Without | Baseline Pass | Without Pass |
|------|--------|----------|---------|---------------|-------------|
| azure-sdk-typescript | 0.0 | 100.0 | 100.0 | ❌ | ❌ |

## Pairwise Details (per Prompt)

### cosmos-db-dp-js-ts-crud

Baseline: **js-ts-cosmos-skill/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|
| azure-sdk-typescript | 0.0 | 100.0 | ❌ |

