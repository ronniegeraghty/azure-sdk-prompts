# Java suite health

Status: **Completed with one infrastructure issue**

Run ID: `20260829-075759`

## Final audit

- Reports: 57/57.
- Prompts: 19/19.
- Complete three-arm triplets: 19/19.
- Evaluation summary: 2 passed all criteria, 55 failed one or more criteria, 0 SDK errors.
- No-output evaluations: 1.
- Session/test/grader timeouts: 0.
- Skill-load failures: 0.
- Azure MCP calls: 177 succeeded, 0 failed.
- Azure MCP average latency: 6.8 seconds.
- Azure MCP maximum latency: 21.4 seconds.
- Azure MCP calls at or above 60 seconds: 0.
- Azure MCP request timeouts: 0.
- Copilot cleanup: no orphaned Hyoka processes found.
- Debug log: losslessly split into three GitHub-safe parts; original SHA-256 `94C70E5D2F0E1CFC2DB6F9FE2F37821E3D3D8285E0087F420C68DCF6679607AD`.
- Generated `summary.json`: losslessly stored as `summary.json.gz`; original SHA-256 `37255396C61696E3F1B5774EE21830C6B9FCACFF9A65EFD886313C32211FF5F1`.

## Primary-check totals

| Variant | Checks passed | Checks total | Rate | No-output evaluations |
|---|---:|---:|---:|---:|
| Baseline | 304 | 395 | 77.0% | 0 |
| Azure Skill + MCP | 341 | 395 | 86.3% | 0 |
| Azure Skill + MCP + Microsoft Skills | 326 | 388 | 84.0% | 1 |

## Interim checkpoint: 2026-08-29 08:18 +08:00

- Reports completed: 3/57.
- Azure MCP calls: 9 succeeded, 0 failed.
- Azure MCP average latency: 7.4 seconds.
- Azure MCP maximum latency: 16.7 seconds.
- MCP request timeouts: 0.
- Session/test timeouts: 0.
- Skill-load failures: 0.

### Recorded issue

| Prompt | Variant | Phase | Observation |
|---|---|---|---|
| `storage-mp-java-account-mgmt` | `java-azure-skills/azure-skill-mcp-microsoft-skill` | Generation | Generation completed after 172 seconds but produced no response, files, or tool calls. Thirteen configured graders then reported that no generated files or response were available. |

The primary result is retained unchanged and queued for a controlled retry after JS/TS completes.
