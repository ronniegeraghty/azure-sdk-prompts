# Evaluation Summary: 20260827-103815

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260827-103815` |
| Timestamp | 2026-08-27T02:38:15Z |
| Total Prompts | 1 |
| Total Configs | 2 |
| Total Evaluations | 2 |
| Passed | 0 |
| Failed | 2 |
| Errors | 0 |
| Duration | 750.0s |

## Comparison Matrix

| Prompt | python-azure-tools/baseline | python-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-projects-dp-python-project-resource-inventory | ❌ 13/16 | ❌ 15/16 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-projects-dp-python-project-resource-inventory](results/ai-projects/data-plane/python/projects/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 13/16 | 309.4s | 3 |
| [ai-projects-dp-python-project-resource-inventory](results/ai-projects/data-plane/python/projects/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 15/16 | 440.6s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-projects-dp-python-project-resource-inventory | 309.4s (python-azure-tools/baseline) | 375.0s | 440.6s (python-azure-tools/with-azure-tools) |

⏱ **Slowest:** ai-projects-dp-python-project-resource-inventory/python-azure-tools/with-azure-tools · **Fastest:** ai-projects-dp-python-project-resource-inventory/python-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-projects-dp-python-project-resource-inventory | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-tools/baseline | 1 | 0 | 1 | 0.0% |
| python-azure-tools/with-azure-tools | 1 | 0 | 1 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 10 | 10 | 0 | 100.0% |
| azure-documentation | 7 | 7 | 0 | 100.0% |
| rg | 6 | 6 | 0 | 100.0% |
| glob | 5 | 4 | 1 | 80.0% |
| github-mcp-server-get_file_contents | 5 | 5 | 0 | 100.0% |
| view | 4 | 4 | 0 | 100.0% |
| web_search | 3 | 3 | 0 | 100.0% |
| github-mcp-server-search_code | 3 | 3 | 0 | 100.0% |
| apply_patch | 2 | 2 | 0 | 100.0% |
| azure-get_azure_bestpractices | 2 | 2 | 0 | 100.0% |
| skill | 1 | 1 | 0 | 100.0% |
| web_fetch | 1 | 1 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### ai-projects-dp-python-project-resource-inventory

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

