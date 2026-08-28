# Azure skills three-way rerun ledger

Branch: `weidongxu-microsoft/issue-656-three-way-rerun`

Progress comment: https://github.com/ronniegeraghty/hyoka/pull/656#issuecomment-5455143231

Overall status: **Running**

Current suite: **Java**

## Suite status

| Order | Suite | Prompts | Evaluations | Status | Run ID | Health |
|---:|---|---:|---:|---|---|---|
| 1 | .NET | 20 | 60 | Completed with 1 issue | `20260829-004156` | 60/60 reports; 20/20 triplets; MCP 261/261 |
| 2 | Python | 19 | 57 | Completed with 1 issue | `20260829-032449` | 57/57 reports; 19/19 triplets; MCP 156/156 |
| 3 | Java | 19 | 57 | In progress | - | Pending post-suite audit |
| 4 | JS/TS | 14 | 42 | Pending | - | - |

Completed: .NET, Python.

Remaining: Java, JS/TS.

## Frozen environment

- Copilot CLI: `copilot.exe.old-60692-1787802413382`
- Copilot CLI SHA-256: `72CA06C41930B83FC323D5C4F5FE97863557DB3F79DA5A198DA16C315577E4EF`
- Copilot SDK: v1
- Azure MCP: `@azure/mcp@3.0.0-beta.38`

## Crash-recovery protocol

1. Read `run-state.json`; it is the machine-readable source of truth.
2. Check the suite marked `in_progress` and inspect its output directory and log.
3. Do not rerun completed evaluations until the partial output has been audited.
4. Before starting a suite, mark it `in_progress`, update this ledger, commit, and push.
5. After a suite, record its run ID and health summary, mark it completed, commit all output, and push.
6. Edit the single progress comment linked above with the same checkpoint before starting the next suite. Do not add a new progress comment for each suite.

## Checkpoint history

| Time | Event | Commit |
|---|---|---|
| 2026-08-29 00:35 +08:00 | Initialized the four-suite rerun plan. | Pending |
| 2026-08-29 00:40 +08:00 | Froze the runtime versions and marked .NET in progress. | Pending |
| 2026-08-29 02:05 +08:00 | Recorded one isolated .NET full-arm generation timeout at 30/60 reports; 121/121 Azure MCP calls had succeeded. | Pending |
| 2026-08-29 03:18 +08:00 | Completed the .NET suite audit: 60/60 reports, 20/20 triplets, one generation timeout, and no MCP/test timeouts. | `5c74a62d` |
| 2026-08-29 03:25 +08:00 | Marked Python in progress after the .NET checkpoint was pushed. | Pending |
| 2026-08-29 04:06 +08:00 | Recorded one Python full-arm SDK generation timeout at 7/57 reports; 28/28 Azure MCP calls had succeeded. | Pending |
| 2026-08-29 07:49 +08:00 | Completed the Python suite audit: 57/57 reports, 19/19 triplets, one generation timeout, and no MCP/test timeouts. | `ddaeb546` |
| 2026-08-29 07:55 +08:00 | Marked Java in progress after the Python checkpoint was pushed. | Pending |
