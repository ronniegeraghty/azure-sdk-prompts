# Evaluation Summary: 20260826-223641

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260826-223641` |
| Timestamp | 2026-08-26T14:36:41Z |
| Total Prompts | 4 |
| Total Configs | 2 |
| Total Evaluations | 8 |
| Passed | 8 |
| Failed | 0 |
| Errors | 0 |
| Duration | 1675.9s |

## Comparison Matrix

| Prompt | dotnet-azure-tools/baseline | dotnet-azure-tools/with-azure-tools |
|--------|--------|--------|
| ai-agents-dp-dotnet-basic-agent-lifecycle | ✅ 11/11 | ✅ 11/11 |
| ai-agents-dp-dotnet-file-search | ✅ 9/9 | ✅ 9/9 |
| ai-agents-dp-dotnet-function-tool | ✅ 12/12 | ✅ 12/12 |
| ai-projects-dp-dotnet-project-resource-inventory | ✅ 9/9 | ✅ 9/9 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [ai-agents-dp-dotnet-basic-agent-lifecycle](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 11/11 | 250.6s | 3 |
| [ai-agents-dp-dotnet-basic-agent-lifecycle](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 11/11 | 152.3s | 3 |
| [ai-agents-dp-dotnet-file-search](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 9/9 | 202.6s | 3 |
| [ai-agents-dp-dotnet-file-search](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 9/9 | 232.1s | 3 |
| [ai-agents-dp-dotnet-function-tool](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 12/12 | 141.7s | 3 |
| [ai-agents-dp-dotnet-function-tool](results/ai-agents/data-plane/dotnet/agents/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 12/12 | 170.6s | 3 |
| [ai-projects-dp-dotnet-project-resource-inventory](results/ai-projects/data-plane/dotnet/projects/dotnet-azure-tools/baseline/report.md) | dotnet-azure-tools/baseline | ✅ | 9/9 | 261.1s | 3 |
| [ai-projects-dp-dotnet-project-resource-inventory](results/ai-projects/data-plane/dotnet/projects/dotnet-azure-tools/with-azure-tools/report.md) | dotnet-azure-tools/with-azure-tools | ✅ | 9/9 | 264.5s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| ai-agents-dp-dotnet-basic-agent-lifecycle | 152.3s (dotnet-azure-tools/with-azure-tools) | 201.4s | 250.6s (dotnet-azure-tools/baseline) |
| ai-agents-dp-dotnet-file-search | 202.6s (dotnet-azure-tools/baseline) | 217.4s | 232.1s (dotnet-azure-tools/with-azure-tools) |
| ai-agents-dp-dotnet-function-tool | 141.7s (dotnet-azure-tools/baseline) | 156.1s | 170.6s (dotnet-azure-tools/with-azure-tools) |
| ai-projects-dp-dotnet-project-resource-inventory | 261.1s (dotnet-azure-tools/baseline) | 262.8s | 264.5s (dotnet-azure-tools/with-azure-tools) |

⏱ **Slowest:** ai-projects-dp-dotnet-project-resource-inventory/dotnet-azure-tools/with-azure-tools · **Fastest:** ai-agents-dp-dotnet-function-tool/dotnet-azure-tools/baseline

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| ai-agents-dp-dotnet-basic-agent-lifecycle | 2 | 2 | 0 | 100.0% |
| ai-agents-dp-dotnet-file-search | 2 | 2 | 0 | 100.0% |
| ai-agents-dp-dotnet-function-tool | 2 | 2 | 0 | 100.0% |
| ai-projects-dp-dotnet-project-resource-inventory | 2 | 2 | 0 | 100.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| dotnet-azure-tools/baseline | 4 | 4 | 0 | 100.0% |
| dotnet-azure-tools/with-azure-tools | 4 | 4 | 0 | 100.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| github-mcp-server-get_file_contents | 39 | 34 | 5 | 87.2% |
| github-mcp-server-search_code | 32 | 32 | 0 | 100.0% |
| powershell | 23 | 23 | 0 | 100.0% |
| apply_patch | 17 | 17 | 0 | 100.0% |
| view | 13 | 5 | 8 | 38.5% |
| web_fetch | 9 | 0 | 9 | 0.0% |
| azure-get_azure_bestpractices | 8 | 8 | 0 | 100.0% |
| glob | 8 | 8 | 0 | 100.0% |
| rg | 7 | 7 | 0 | 100.0% |
| skill | 4 | 4 | 0 | 100.0% |
| azure-documentation | 4 | 4 | 0 | 100.0% |
| web_search | 2 | 2 | 0 | 100.0% |

## Pairwise Details (per Prompt)

### ai-agents-dp-dotnet-basic-agent-lifecycle

Baseline: **dotnet-azure-tools/baseline** — 11/11

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-dotnet-file-search

Baseline: **dotnet-azure-tools/baseline** — 9/9

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-agents-dp-dotnet-function-tool

Baseline: **dotnet-azure-tools/baseline** — 12/12

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

### ai-projects-dp-dotnet-project-resource-inventory

Baseline: **dotnet-azure-tools/baseline** — 9/9

| Tool Removed | Impact | Without Score | Pass |
|-------------|--------|---------------|------|

