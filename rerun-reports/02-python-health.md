# Python suite health

Status: **Completed with one infrastructure issue**

Run ID: `20260829-032449`

## Final audit

- Reports: 57/57.
- Prompts: 19/19.
- Complete three-arm triplets: 19/19.
- Evaluation summary: 0 passed all criteria, 56 failed one or more criteria, 1 infrastructure error.
- Test/grader timeouts: 0.
- Skill-load failures: 0.
- Azure MCP calls: 156 succeeded, 0 failed.
- Azure MCP average latency: 6.5 seconds.
- Azure MCP maximum latency: 27.6 seconds.
- Azure MCP calls at or above 60 seconds: 0.
- Azure MCP request timeouts: 0.
- Copilot cleanup: no orphaned Hyoka processes found.
- Debug log: losslessly split into four GitHub-safe parts; original SHA-256 `B597C6117F2607AF0FC9CA788B7CFD33550A966D14EB1D07A5D36F5D4CB9ABE1`.
- Generated `summary.json`: losslessly stored as `summary.json.gz`; original SHA-256 `D06ACA1A24A158F438D8023E93D436EE05BCCC11ECB16C9FAFB27196F09D2C57`.

## Primary-check totals

| Variant | Checks passed | Checks total | Rate | Infrastructure errors |
|---|---:|---:|---:|---:|
| Baseline | 227 | 305 | 74.4% | 0 |
| Azure Skill + MCP | 259 | 305 | 84.9% | 0 |
| Azure Skill + MCP + Microsoft Skills | 265 | 305 | 86.9% | 1 |

## Interim checkpoint: 2026-08-29 04:06 +08:00

- Reports completed: 7/57.
- Azure MCP calls: 28 succeeded, 0 failed.
- Azure MCP average latency: 6.4 seconds.
- Azure MCP maximum latency: 14.9 seconds.
- MCP request timeouts: 0.
- Test/grader timeouts: 0.
- Skill-load failures: 0.

### Recorded issue

| Prompt | Variant | Phase | Observation |
|---|---|---|---|
| `storage-mp-python-account-mgmt` | `python-azure-skills/azure-skill-mcp-microsoft-skill` | Generation | Hit `waiting for session.idle: context deadline exceeded` after approximately 10m13s. Two files were generated, but the report was classified as an SDK evaluation error. No Azure MCP call timed out. |

The primary result is retained unchanged. Any targeted retry will be recorded separately after all four primary suites complete.
