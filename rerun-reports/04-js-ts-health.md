# JS/TS suite health

Status: **In progress**

- Run ID: `20260829-141457`
- Progress: 21/42 reports
- Complete triplets: 6/14
- Azure MCP timeouts: 0 observed

## Interim issues

| Prompt | Arm | Issue | Generated output |
|---|---|---|---:|
| `event-hubs-dp-js-ts-streaming` | Baseline | Timed out waiting for `session.idle` after the 10-minute generation limit. | 4 files |
| `event-hubs-dp-js-ts-streaming` | Azure Skill + MCP | Timed out waiting for `session.idle` after the 10-minute generation limit. | 4 files |
| `event-hubs-dp-js-ts-streaming` | Azure Skill + MCP + Microsoft Skills | Timed out waiting for `session.idle` after the 10-minute generation limit. | 4 files |
| `key-vault-dp-js-ts-crud` | Baseline | Completed with no files or executed tool calls after emitting malformed tool-invocation text. | None |

All four evaluations are queued for controlled retry after all primary suites finish. The suite remains running.
