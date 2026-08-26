# Evaluation Summary: 20260826-192543

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-192543` |
| Timestamp | 2026-08-26T11:25:43Z |
| Total Prompts | 4 |
| Total Configs | 2 |
| Total Evaluations | 8 |
| Passed | 0 |
| Failed | 8 |
| Errors | 0 |
| Duration | 2655.4s |

## Comparison Matrix

| Prompt | python-azure-tools/baseline | python-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-agents-dp-python-basic-agent-lifecycle | ❌ 14/17 | ❌ 14/17 |
| ai-agents-dp-python-file-search | ❌ 12/16 | ❌ 13/16 |
| ai-agents-dp-python-function-tool | ❌ 16/19 | ❌ 1/19 |
| ai-projects-dp-python-project-resource-inventory | ❌ 13/16 | ❌ 14/16 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-agents-dp-python-basic-agent-lifecycle](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 14/17 | 355.8s | 9 |
| [ai-agents-dp-python-basic-agent-lifecycle](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 14/17 | 471.3s | 4 |
| [ai-agents-dp-python-file-search](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 12/16 | 361.1s | 3 |
| [ai-agents-dp-python-file-search](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 13/16 | 348.6s | 4 |
| [ai-agents-dp-python-function-tool](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 16/19 | 328.3s | 3 |
| [ai-agents-dp-python-function-tool](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 1/19 | 248.6s | 0 |
| [ai-projects-dp-python-project-resource-inventory](results/ai-projects/data-plane/python/projects/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 13/16 | 222.3s | 3 |
| [ai-projects-dp-python-project-resource-inventory](results/ai-projects/data-plane/python/projects/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 14/16 | 318.4s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-agents-dp-python-basic-agent-lifecycle | 355.8s (python-azure-tools/baseline) | 413.6s | 471.3s (python-azure-tools/with-azure-tools) |
| ai-agents-dp-python-file-search | 348.6s (python-azure-tools/with-azure-tools) | 354.8s | 361.1s (python-azure-tools/baseline) |
| ai-agents-dp-python-function-tool | 248.6s (python-azure-tools/with-azure-tools) | 288.5s | 328.3s (python-azure-tools/baseline) |
| ai-projects-dp-python-project-resource-inventory | 222.3s (python-azure-tools/baseline) | 270.4s | 318.4s (python-azure-tools/with-azure-tools) |

⏱ **Slowest:** ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/with-azure-tools · **Fastest:** ai-projects-dp-python-project-resource-inventory/python-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-agents-dp-python-basic-agent-lifecycle | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-python-file-search | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-python-function-tool | 2 | 0 | 2 | 0.0% |
| ai-projects-dp-python-project-resource-inventory | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-tools/baseline | 4 | 0 | 4 | 0.0% |
| python-azure-tools/with-azure-tools | 4 | 0 | 4 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 44 | 44 | 0 | 100.0% |
| view | 25 | 25 | 0 | 100.0% |
| github-mcp-server-search_code | 22 | 22 | 0 | 100.0% |
| rg | 14 | 14 | 0 | 100.0% |
| azure-documentation | 14 | 14 | 0 | 100.0% |
| apply_patch | 13 | 13 | 0 | 100.0% |
| glob | 12 | 12 | 0 | 100.0% |
| web_fetch | 11 | 10 | 1 | 90.9% |
| github-mcp-server-get_file_contents | 9 | 9 | 0 | 100.0% |
| azure-get_azure_bestpractices | 8 | 8 | 0 | 100.0% |
| web_search | 6 | 6 | 0 | 100.0% |
| skill | 5 | 2 | 3 | 40.0% |

## Pairwise Details (per Prompt)

### ai-agents-dp-python-basic-agent-lifecycle

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-python-file-search

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-python-function-tool

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-projects-dp-python-project-resource-inventory

Baseline: **python-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

