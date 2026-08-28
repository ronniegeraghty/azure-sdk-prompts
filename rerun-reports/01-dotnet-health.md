# .NET suite health

Run ID: `20260829-004156`

Status: **Completed with one infrastructure issue**

## Final audit

- Reports: 60/60.
- Prompts: 20/20.
- Complete three-arm triplets: 20/20.
- Evaluation summary: 17 passed, 42 failed criteria, 1 infrastructure error.
- Test/grader timeouts: 0.
- Skill-load failures: 0.
- Azure MCP calls: 261 succeeded, 0 failed.
- Azure MCP average latency: 5.7 seconds.
- Azure MCP maximum latency: 46.5 seconds.
- Azure MCP calls at or above 60 seconds: 0.
- Azure MCP request timeouts: 0.
- Copilot cleanup: 5 completed sessions removed; no orphaned Hyoka processes found.

## Primary-check totals

| Variant | Checks passed | Checks total | Rate | Infrastructure errors |
|---|---:|---:|---:|---:|
| Baseline | 114 | 135 | 84.4% | 0 |
| Azure Skill + MCP | 105 | 135 | 77.8% | 0 |
| Azure Skill + MCP + Microsoft Skills | 109 | 130 | 83.8% | 1 |

## Interim checkpoint: 2026-08-29 02:05 +08:00

- Reports completed: 30/60.
- Azure MCP calls: 121 succeeded, 0 failed.
- Azure MCP average latency: 6.3 seconds.
- Azure MCP maximum latency: 37.0 seconds.
- MCP request timeouts: 0.
- Skill-load failures: 0.

### Recorded issue

| Prompt | Variant | Phase | Observation |
|---|---|---|---|
| `identity-dp-dotnet-managed-identity` | `dotnet-azure-skills/azure-skill-mcp-microsoft-skill` | Generation | Ran for 10m18s, produced no files, and ended with `context deadline exceeded` during session cleanup. Review then failed because no generated files or response were available. No MCP tool call was involved. |

The primary result is retained unchanged. Any targeted retry will be recorded separately after all four primary suites complete.
