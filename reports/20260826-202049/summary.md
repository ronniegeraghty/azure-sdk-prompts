# Evaluation Summary: 20260826-202049

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-202049` |
| Timestamp | 2026-08-26T12:20:49Z |
| Total Prompts | 4 |
| Total Configs | 2 |
| Total Evaluations | 8 |
| Passed | 0 |
| Failed | 7 |
| Errors | 1 |
| Duration | 3593.8s |

## Comparison Matrix

| Prompt | js-ts-azure-tools/baseline | js-ts-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-agents-dp-js-ts-basic-agent-lifecycle | ❌ 16/20 | ❌ 16/20 |
| ai-agents-dp-js-ts-file-search | ❌ 15/19 | ❌ 16/19 |
| ai-agents-dp-js-ts-function-tool | ❌ 17/22 | ❌ 18/22 |
| ai-projects-dp-js-ts-project-resource-inventory | ❌ 16/19 | ❌ 16/19 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-agents-dp-js-ts-basic-agent-lifecycle](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 16/20 | 318.9s | 5 |
| [ai-agents-dp-js-ts-basic-agent-lifecycle](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 16/20 | 407.4s | 5 |
| [ai-agents-dp-js-ts-file-search](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 15/19 | 431.3s | 5 |
| [ai-agents-dp-js-ts-file-search](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 16/19 | 369.7s | 5 |
| [ai-agents-dp-js-ts-function-tool](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 17/22 | 811.7s | 4 |
| [ai-agents-dp-js-ts-function-tool](results/ai-agents/data-plane/js-ts/agents/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 18/22 | 460.9s | 5 |
| [ai-projects-dp-js-ts-project-resource-inventory](results/ai-projects/data-plane/js-ts/projects/js-ts-azure-tools/baseline/report.md) | js-ts-azure-tools/baseline | ❌ | 16/19 | 404.2s | 5 |
| [ai-projects-dp-js-ts-project-resource-inventory](results/ai-projects/data-plane/js-ts/projects/js-ts-azure-tools/with-azure-tools/report.md) | js-ts-azure-tools/with-azure-tools | ❌ | 16/19 | 373.4s | 5 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-agents-dp-js-ts-basic-agent-lifecycle | 318.9s (js-ts-azure-tools/baseline) | 363.1s | 407.4s (js-ts-azure-tools/with-azure-tools) |
| ai-agents-dp-js-ts-file-search | 369.7s (js-ts-azure-tools/with-azure-tools) | 400.5s | 431.3s (js-ts-azure-tools/baseline) |
| ai-agents-dp-js-ts-function-tool | 460.9s (js-ts-azure-tools/with-azure-tools) | 636.3s | 811.7s (js-ts-azure-tools/baseline) |
| ai-projects-dp-js-ts-project-resource-inventory | 373.4s (js-ts-azure-tools/with-azure-tools) | 388.8s | 404.2s (js-ts-azure-tools/baseline) |

⏱ **Slowest:** ai-agents-dp-js-ts-function-tool/js-ts-azure-tools/baseline · **Fastest:** ai-agents-dp-js-ts-basic-agent-lifecycle/js-ts-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-agents-dp-js-ts-basic-agent-lifecycle | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-js-ts-file-search | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-js-ts-function-tool | 2 | 0 | 2 | 0.0% |
| ai-projects-dp-js-ts-project-resource-inventory | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| js-ts-azure-tools/baseline | 4 | 0 | 4 | 0.0% |
| js-ts-azure-tools/with-azure-tools | 4 | 0 | 4 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| powershell | 35 | 35 | 0 | 100.0% |
| rg | 22 | 22 | 0 | 100.0% |
| glob | 19 | 19 | 0 | 100.0% |
| github-mcp-server-search_code | 18 | 18 | 0 | 100.0% |
| apply_patch | 17 | 17 | 0 | 100.0% |
| azure-documentation | 16 | 16 | 0 | 100.0% |
| github-mcp-server-get_file_contents | 16 | 16 | 0 | 100.0% |
| view | 16 | 16 | 0 | 100.0% |
| azure-get_azure_bestpractices | 8 | 8 | 0 | 100.0% |
| skill | 6 | 6 | 0 | 100.0% |
| read_powershell | 4 | 4 | 0 | 100.0% |
| web_search | 3 | 3 | 0 | 100.0% |
| web_fetch | 2 | 2 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### ai-agents-dp-js-ts-basic-agent-lifecycle

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-js-ts-file-search

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-js-ts-function-tool

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-projects-dp-js-ts-project-resource-inventory

Baseline: **js-ts-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

