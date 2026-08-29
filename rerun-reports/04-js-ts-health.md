# JS/TS suite health

Status: **In progress**

- Run ID: `20260829-141457`
- Progress: 4/42 reports
- Complete triplets: 0/14
- Azure MCP timeouts: 0 observed

## Interim issues

| Prompt | Arm | Issue | Generated output |
|---|---|---|---:|
| `event-hubs-dp-js-ts-streaming` | Baseline | Timed out waiting for `session.idle` after the 10-minute generation limit. | 4 files |
| `event-hubs-dp-js-ts-streaming` | Azure Skill + MCP | Timed out waiting for `session.idle` after the 10-minute generation limit. | 4 files |

Both evaluations are queued for controlled retry after all primary suites finish. The suite remains running.
