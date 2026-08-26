# Evaluation Summary: 20260826-102427

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-102427` |
| Timestamp | 2026-08-26T02:24:27Z |
| Total Prompts | 1 |
| Total Configs | 2 |
| Total Evaluations | 2 |
| Passed | 0 |
| Failed | 2 |
| Errors | 0 |
| Duration | 474.3s |

## Comparison Matrix

| Prompt | python-cosmos-skill/baseline | python-cosmos-skill/without-azure-sdk-python |
|--------|--------|--------|
| cosmos-db-dp-python-crud | ❌ 10/13 | ❌ 9/13 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-cosmos-skill/baseline/report.md) | python-cosmos-skill/baseline | ❌ | 10/13 | 243.5s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-cosmos-skill/without-azure-sdk-python/report.md) | python-cosmos-skill/without-azure-sdk-python | ❌ | 9/13 | 230.7s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-python-crud | 230.7s (python-cosmos-skill/without-azure-sdk-python) | 237.1s | 243.5s (python-cosmos-skill/baseline) |

⏱ **Slowest:** cosmos-db-dp-python-crud/python-cosmos-skill/baseline · **Fastest:** cosmos-db-dp-python-crud/python-cosmos-skill/without-azure-sdk-python

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| cosmos-db-dp-python-crud | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-cosmos-skill/baseline | 1 | 0 | 1 | 0.0% |
| python-cosmos-skill/without-azure-sdk-python | 1 | 0 | 1 | 0.0% |

## Pairwise Tool Impact

Impact = baseline score − score without tool (positive = tool helps).

| Tool | Impact | Baseline | Without | Baseline Pass | Without Pass |
|------|--------|----------|---------|---------------|-------------|
| azure-sdk-python | 0.0 | 100.0 | 100.0 | ❌ | ❌ |

## Pairwise Details (per Prompt)

### cosmos-db-dp-python-crud

Baseline: **python-cosmos-skill/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|
| azure-sdk-python | 0.0 | 100.0 | ❌ |

