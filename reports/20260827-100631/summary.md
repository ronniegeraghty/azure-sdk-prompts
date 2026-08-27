# Evaluation Summary: 20260827-100631

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260827-100631` |
| Timestamp | 2026-08-27T02:06:31Z |
| Total Prompts | 3 |
| Total Configs | 2 |
| Total Evaluations | 6 |
| Passed | 0 |
| Failed | 6 |
| Errors | 0 |
| Duration | 1899.4s |

## Comparison Matrix

| Prompt | python-azure-tools/baseline | python-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-agents-dp-python-basic-agent-lifecycle | ❌ 14/17 | ❌ 15/17 |
| ai-agents-dp-python-file-search | ❌ 13/16 | ❌ 14/16 |
| ai-agents-dp-python-function-tool | ❌ 14/19 | ❌ 2/19 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-agents-dp-python-basic-agent-lifecycle](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 14/17 | 384.1s | 11 |
| [ai-agents-dp-python-basic-agent-lifecycle](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 15/17 | 342.9s | 9 |
| [ai-agents-dp-python-file-search](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 13/16 | 273.6s | 3 |
| [ai-agents-dp-python-file-search](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 14/16 | 268.6s | 4 |
| [ai-agents-dp-python-function-tool](results/ai-agents/data-plane/python/agents/python-azure-tools/baseline/report.md) | python-azure-tools/baseline | ❌ | 14/19 | 363.7s | 3 |
| [ai-agents-dp-python-function-tool](results/ai-agents/data-plane/python/agents/python-azure-tools/with-azure-tools/report.md) | python-azure-tools/with-azure-tools | ❌ | 2/19 | 265.4s | 0 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-agents-dp-python-function-tool | 265.4s (python-azure-tools/with-azure-tools) | 314.6s | 363.7s (python-azure-tools/baseline) |
| ai-agents-dp-python-basic-agent-lifecycle | 342.9s (python-azure-tools/with-azure-tools) | 363.5s | 384.1s (python-azure-tools/baseline) |
| ai-agents-dp-python-file-search | 268.6s (python-azure-tools/with-azure-tools) | 271.1s | 273.6s (python-azure-tools/baseline) |

⏱ **Slowest:** ai-agents-dp-python-basic-agent-lifecycle/python-azure-tools/baseline · **Fastest:** ai-agents-dp-python-function-tool/python-azure-tools/with-azure-tools

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-agents-dp-python-basic-agent-lifecycle | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-python-file-search | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-python-function-tool | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| python-azure-tools/baseline | 3 | 0 | 3 | 0.0% |
| python-azure-tools/with-azure-tools | 3 | 0 | 3 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 29 | 29 | 0 | 100.0% |
| github-mcp-server-search_code | 26 | 26 | 0 | 100.0% |
| view | 20 | 20 | 0 | 100.0% |
| github-mcp-server-get_file_contents | 14 | 11 | 3 | 78.6% |
| rg | 12 | 12 | 0 | 100.0% |
| apply_patch | 12 | 12 | 0 | 100.0% |
| azure-documentation | 11 | 11 | 0 | 100.0% |
| glob | 10 | 10 | 0 | 100.0% |
| web_fetch | 10 | 9 | 1 | 90.0% |
| azure-get_azure_bestpractices | 6 | 6 | 0 | 100.0% |
| skill | 4 | 3 | 1 | 75.0% |
| web_search | 4 | 4 | 0 | 100.0% |

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

