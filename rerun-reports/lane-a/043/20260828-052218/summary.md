# Evaluation Summary: 20260828-052218

## Run Statistics

| Metric | Value |
|--------|-------|
| Run ID | `20260828-052218` |
| Timestamp | 2026-08-27T21:22:18Z |
| Total Prompts | 1 |
| Total Configs | 1 |
| Total Evaluations | 1 |
| Passed | 0 |
| Failed | 1 |
| Errors | 0 |
| Duration | 780.6s |

## Comparison Matrix

| Prompt | java-azure-skills/azure-skill-mcp |
|--------|--------|
| storage-mp-java-account-mgmt | ❌ 17/20 |

## Detailed Results

| Prompt | Config | Result | Score | Duration | Files |
|--------|--------|--------|-------|----------|-------|
| [storage-mp-java-account-mgmt](results/storage/management-plane/java/provisioning/java-azure-skills/azure-skill-mcp/report.md) | java-azure-skills/azure-skill-mcp | ❌ | 17/20 | 780.6s | 3 |

## Duration Analysis (by Prompt)

| Prompt | Min | Avg | Max |
|--------|-----|-----|-----|
| storage-mp-java-account-mgmt | 780.6s (java-azure-skills/azure-skill-mcp) | 780.6s | 780.6s (java-azure-skills/azure-skill-mcp) |

⏱ **Slowest:** storage-mp-java-account-mgmt/java-azure-skills/azure-skill-mcp · **Fastest:** storage-mp-java-account-mgmt/java-azure-skills/azure-skill-mcp

## Prompt Comparison

| Prompt | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| storage-mp-java-account-mgmt | 1 | 0 | 1 | 0.0% |

## Config Comparison

| Config | Total | Passed | Failed | Pass Rate |
|--------|-------|--------|--------|----------|
| java-azure-skills/azure-skill-mcp | 1 | 0 | 1 | 0.0% |

## Tool Usage

| Tool | Calls | Successes | Failures | Success Rate |
|------|-------|-----------|----------|-------------|
| github-mcp-server-get_file_contents | 10 | 10 | 0 | 100.0% |
| github-mcp-server-search_code | 5 | 5 | 0 | 100.0% |
| rg | 2 | 2 | 0 | 100.0% |
| apply_patch | 2 | 2 | 0 | 100.0% |
| powershell | 2 | 2 | 0 | 100.0% |
| view | 2 | 2 | 0 | 100.0% |
| glob | 2 | 2 | 0 | 100.0% |
| azure-get_azure_bestpractices | 2 | 0 | 2 | 0.0% |
| azure-documentation | 1 | 0 | 1 | 0.0% |
| skill | 1 | 1 | 0 | 100.0% |

