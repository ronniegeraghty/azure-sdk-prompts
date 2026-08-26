# Evaluation Summary: 20260826-212414

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-212414` |
| Timestamp | 2026-08-26T13:24:14Z |
| Total Prompts | 4 |
| Total Configs | 2 |
| Total Evaluations | 8 |
| Passed | 0 |
| Failed | 8 |
| Errors | 0 |
| Duration | 4125.9s |

## Comparison Matrix

| Prompt | java-azure-tools/baseline | java-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-agents-dp-java-basic-agent-lifecycle | ❌ 20/23 | ❌ 18/23 |
| ai-agents-dp-java-file-search | ❌ 5/21 | ❌ 17/21 |
| ai-agents-dp-java-function-tool | ❌ 20/24 | ❌ 19/24 |
| ai-projects-dp-java-project-resource-inventory | ❌ 19/21 | ❌ 18/21 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-agents-dp-java-basic-agent-lifecycle](results/ai-agents/data-plane/java/agents/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 20/23 | 502.0s | 3 |
| [ai-agents-dp-java-basic-agent-lifecycle](results/ai-agents/data-plane/java/agents/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 18/23 | 523.6s | 3 |
| [ai-agents-dp-java-file-search](results/ai-agents/data-plane/java/agents/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 5/21 | 375.3s | 0 |
| [ai-agents-dp-java-file-search](results/ai-agents/data-plane/java/agents/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 17/21 | 570.4s | 3 |
| [ai-agents-dp-java-function-tool](results/ai-agents/data-plane/java/agents/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 20/24 | 570.5s | 3 |
| [ai-agents-dp-java-function-tool](results/ai-agents/data-plane/java/agents/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 19/24 | 592.5s | 3 |
| [ai-projects-dp-java-project-resource-inventory](results/ai-projects/data-plane/java/projects/java-azure-tools/baseline/report.md) | java-azure-tools/baseline | ❌ | 19/21 | 556.4s | 3 |
| [ai-projects-dp-java-project-resource-inventory](results/ai-projects/data-plane/java/projects/java-azure-tools/with-azure-tools/report.md) | java-azure-tools/with-azure-tools | ❌ | 18/21 | 434.8s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-agents-dp-java-basic-agent-lifecycle | 502.0s (java-azure-tools/baseline) | 512.8s | 523.6s (java-azure-tools/with-azure-tools) |
| ai-agents-dp-java-file-search | 375.3s (java-azure-tools/baseline) | 472.8s | 570.4s (java-azure-tools/with-azure-tools) |
| ai-agents-dp-java-function-tool | 570.5s (java-azure-tools/baseline) | 581.5s | 592.5s (java-azure-tools/with-azure-tools) |
| ai-projects-dp-java-project-resource-inventory | 434.8s (java-azure-tools/with-azure-tools) | 495.6s | 556.4s (java-azure-tools/baseline) |

⏱ **Slowest:** ai-agents-dp-java-function-tool/java-azure-tools/with-azure-tools · **Fastest:** ai-agents-dp-java-file-search/java-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-agents-dp-java-basic-agent-lifecycle | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-java-file-search | 2 | 0 | 2 | 0.0% |
| ai-agents-dp-java-function-tool | 2 | 0 | 2 | 0.0% |
| ai-projects-dp-java-project-resource-inventory | 2 | 0 | 2 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| java-azure-tools/baseline | 4 | 0 | 4 | 0.0% |
| java-azure-tools/with-azure-tools | 4 | 0 | 4 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| github-mcp-server-get_file_contents | 73 | 71 | 2 | 97.3% |
| github-mcp-server-search_code | 41 | 41 | 0 | 100.0% |
| powershell | 28 | 28 | 0 | 100.0% |
| rg | 22 | 22 | 0 | 100.0% |
| apply_patch | 14 | 14 | 0 | 100.0% |
| view | 14 | 14 | 0 | 100.0% |
| glob | 12 | 12 | 0 | 100.0% |
| web_fetch | 12 | 12 | 0 | 100.0% |
| azure-get_azure_bestpractices | 8 | 8 | 0 | 100.0% |
| web_search | 4 | 4 | 0 | 100.0% |
| skill | 4 | 4 | 0 | 100.0% |
| azure-documentation | 3 | 3 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### ai-agents-dp-java-basic-agent-lifecycle

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-java-file-search

Baseline: **java-azure-tools/baseline** — 0/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-java-function-tool

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-projects-dp-java-project-resource-inventory

Baseline: **java-azure-tools/baseline** — 1/1

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

