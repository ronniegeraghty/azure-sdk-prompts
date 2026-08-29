# Azure skills three-way rerun ledger

Branch: `weidongxu-microsoft/issue-656-three-way-rerun`

Progress comment: https://github.com/ronniegeraghty/hyoka/pull/656#issuecomment-5455143231

Overall status: **Running**

Current suite: **JS/TS**

## Suite status

| Order | Suite | Prompts | Evaluations | Status | Run ID | Health |
|---:|---|---:|---:|---|---|---|
| 1 | .NET | 20 | 60 | Completed with 1 issue | `20260829-004156` | 60/60 reports; 20/20 triplets; MCP 261/261 |
| 2 | Python | 19 | 57 | Completed with 1 issue | `20260829-032449` | 57/57 reports; 19/19 triplets; MCP 156/156 |
| 3 | Java | 19 | 57 | Completed with 1 issue | `20260829-075759` | 57/57 reports; 19/19 triplets; MCP 177/177 |
| 4 | JS/TS | 14 | 42 | In progress | - | Pending post-suite audit |

Completed: .NET, Python, Java.

Remaining: JS/TS.

## Retry and finalization policy

1. Finish all four primary language suites before running retries.
2. Retry every evaluation with an infrastructure timeout, missing response, or missing generated output using the same frozen runtime.
3. Store retries under `rerun-reports/05-retries` and keep the primary result unchanged.
4. If a retry is healthy, select it for that prompt/config in `selected-results.json`.
5. If a retry still has an infrastructure problem, retain both attempts and explicitly flag or exclude the unresolved triplet in the final report.
6. Generate final reports only from the selection manifest.
7. Commit and push every primary and retry artifact to `weidongxu-microsoft/issue-656-three-way-rerun`.
8. Replace #656's prior-run report data and analysis with the new selected results. Remove obsolete comments about the prior MCP timeout behavior and its usage/causal analysis instead of leaving contradictory results.
9. Replacing the original `weidongxu-microsoft/azure-skills-three-way-comparison` branch is deferred until the final rerun is reviewed.

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
| 2026-08-29 08:18 +08:00 | Recorded one Java full-arm no-output generation at 3/57 reports; 9/9 Azure MCP calls had succeeded. | Pending |
| 2026-08-29 12:03 +08:00 | Recorded the post-suite retry, result-selection, and #656 replacement policy. | Pending |
| 2026-08-29 14:06 +08:00 | Completed the Java suite audit: 57/57 reports, 19/19 triplets, one no-output generation, and no MCP/session/test timeouts. | `326de281` |
| 2026-08-29 14:12 +08:00 | Marked JS/TS in progress after the Java checkpoint was pushed. | Pending |
| 2026-08-29 14:56 +08:00 | At 4/42 JS/TS reports, recorded session-idle timeouts for the Event Hubs baseline and Azure Skill + MCP arms; both produced four files. No Azure MCP timeout was observed. | Pending |
