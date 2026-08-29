# JS/TS suite health

Status: **Completed with four infrastructure issues**

Run ID: `20260829-141457`

## Final audit

- Reports: 42/42.
- Prompts: 14/14.
- Complete three-arm triplets: 14/14.
- Evaluation summary: 0 passed all criteria, 39 failed one or more criteria, 3 SDK errors.
- No-output evaluations: 1.
- Session-idle timeouts: 3.
- Test/grader timeouts: 0.
- Skill-load failures: 0.
- Azure MCP calls: 103 succeeded, 0 failed.
- Azure MCP average latency: 8.3 seconds.
- Azure MCP maximum latency: 45.7 seconds.
- Azure MCP calls at or above 60 seconds: 0.
- Azure MCP request timeouts: 0.
- Copilot cleanup: no orphaned Hyoka processes; 0 sessions required cleanup.
- Debug log: 49,418,418 bytes; SHA-256 `0393243D3DF7350375D49F25B24036D43070BE41749DC776B7E407DDF890C867`.
- Generated `summary.json`: 48,344,179 bytes; SHA-256 `2645C42E7294178759BFACBE03D830C38C588C8284AFA15B7C08C1B33E418092`.

## Primary-check totals

| Variant | Checks passed | Checks total | Rate | No-output evaluations |
|---|---:|---:|---:|---:|
| Baseline | 202 | 266 | 75.9% | 1 |
| Azure Skill + MCP | 208 | 266 | 78.2% | 0 |
| Azure Skill + MCP + Microsoft Skills | 208 | 266 | 78.2% | 0 |

## Recorded issues

| Prompt | Variant | Phase | Observation |
|---|---|---|---|
| `event-hubs-dp-js-ts-streaming` | Baseline | Generation | Timed out waiting for `session.idle` after the 10-minute generation limit; four files were generated. |
| `event-hubs-dp-js-ts-streaming` | Azure Skill + MCP | Generation | Timed out waiting for `session.idle` after the 10-minute generation limit; four files were generated. |
| `event-hubs-dp-js-ts-streaming` | Azure Skill + MCP + Microsoft Skills | Generation | Timed out waiting for `session.idle` after the 10-minute generation limit; four files were generated. |
| `key-vault-dp-js-ts-crud` | Baseline | Generation | Completed with no generated files or executed tool calls after tool invocations were emitted as response text. |

The primary results are retained unchanged. All four affected evaluations are queued for controlled retry.
