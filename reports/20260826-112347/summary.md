# Evaluation Summary: 20260826-112347

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-112347` |
| Timestamp | 2026-08-26T03:23:47Z |
| Total Prompts | 1 |
| Total Configs | 2 |
| Total Evaluations | 2 |
| Passed | 0 |
| Failed | 2 |
| Errors | 0 |
| Duration | 822.1s |

## Comparison Matrix

| Prompt | java-cosmos-skill/baseline | java-cosmos-skill/without-azure-sdk-java |
|--------|--------|--------|
| cosmos-db-dp-java-crud | ❌ 16/19 | ❌ 14/19 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-cosmos-skill/baseline/report.md) | java-cosmos-skill/baseline | ❌ | 16/19 | 476.6s | 3 |
| [cosmos-db-dp-java-crud](results/cosmos-db/data-plane/java/crud/java-cosmos-skill/without-azure-sdk-java/report.md) | java-cosmos-skill/without-azure-sdk-java | ❌ | 14/19 | 345.4s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-java-crud | 345.4s (java-cosmos-skill/without-azure-sdk-java) | 411.0s | 476.6s (java-cosmos-skill/baseline) |

⏱ **Slowest:** cosmos-db-dp-java-crud/java-cosmos-skill/baseline · **Fastest:** cosmos-db-dp-java-crud/java-cosmos-skill/without-azure-sdk-java

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| cosmos-db-dp-java-crud | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| java-cosmos-skill/baseline | 1 | 0 | 1 | 0.0% |
| java-cosmos-skill/without-azure-sdk-java | 1 | 0 | 1 | 0.0% |

## Pairwise Tool Impact

Impact = baseline score − score without tool (positive = tool helps).

| Tool | Impact | Baseline | Without | Baseline Pass | Without Pass |
|------|--------|----------|---------|---------------|-------------|
| azure-sdk-java | 0.0 | 100.0 | 100.0 | ❌ | ❌ |

## Pairwise Details (per Prompt)

### cosmos-db-dp-java-crud

Baseline: **java-cosmos-skill/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|
| azure-sdk-java | 0.0 | 100.0 | ❌ |

