# Evaluation Summary: 20260827-104745

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260827-104745` |
| Timestamp | 2026-08-27T02:47:45Z |
| Total Prompts | 1 |
| Total Configs | 3 |
| Total Evaluations | 3 |
| Passed | 0 |
| Failed | 3 |
| Errors | 0 |
| Duration | 550.9s |

## Comparison Matrix

| Prompt | python-azure-skills/azure | python-azure-skills/azure-with-sdk | python-azure-skills/baseline |
|--------|--------|--------|--------|
| cosmos-db-dp-python-crud | ❌ 9/13 | ❌ 10/13 | ❌ 8/13 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure/report.md) | python-azure-skills/azure | ❌ | 9/13 | 183.8s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/azure-with-sdk/report.md) | python-azure-skills/azure-with-sdk | ❌ | 10/13 | 204.2s | 2 |
| [cosmos-db-dp-python-crud](results/cosmos-db/data-plane/python/crud/python-azure-skills/baseline/report.md) | python-azure-skills/baseline | ❌ | 8/13 | 162.8s | 2 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| cosmos-db-dp-python-crud | 162.8s (python-azure-skills/baseline) | 183.6s | 204.2s (python-azure-skills/azure-with-sdk) |

⏱ **Slowest:** cosmos-db-dp-python-crud/python-azure-skills/azure-with-sdk · **Fastest:** cosmos-db-dp-python-crud/python-azure-skills/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| cosmos-db-dp-python-crud | 3 | 0 | 3 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-skills/azure | 1 | 0 | 1 | 0.0% |
| python-azure-skills/azure-with-sdk | 1 | 0 | 1 | 0.0% |
| python-azure-skills/baseline | 1 | 0 | 1 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| view | 6 | 6 | 0 | 100.0% |
| glob | 5 | 5 | 0 | 100.0% |
| apply_patch | 4 | 4 | 0 | 100.0% |
| powershell | 4 | 4 | 0 | 100.0% |
| azure-documentation | 3 | 3 | 0 | 100.0% |
| azure-get_azure_bestpractices | 2 | 2 | 0 | 100.0% |
| skill | 1 | 1 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### cosmos-db-dp-python-crud

Baseline: **python-azure-skills/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

