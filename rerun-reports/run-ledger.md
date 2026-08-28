# Azure skills three-way rerun ledger

Branch: `weidongxu-microsoft/issue-656-three-way-rerun`

Overall status: **Planned**

Current suite: **None**

## Suite status

| Order | Suite | Prompts | Evaluations | Status | Run ID | Health |
|---:|---|---:|---:|---|---|---|
| 1 | .NET | 20 | 60 | Pending | - | - |
| 2 | Python | 19 | 57 | Pending | - | - |
| 3 | Java | 19 | 57 | Pending | - | - |
| 4 | JS/TS | 14 | 42 | Pending | - | - |

Completed: none.

Remaining: .NET, Python, Java, JS/TS.

## Crash-recovery protocol

1. Read `run-state.json`; it is the machine-readable source of truth.
2. Check the suite marked `in_progress` and inspect its output directory and log.
3. Do not rerun completed evaluations until the partial output has been audited.
4. Before starting a suite, mark it `in_progress`, update this ledger, commit, and push.
5. After a suite, record its run ID and health summary, mark it completed, commit all output, and push.
6. Post the same checkpoint to #656 before starting the next suite.

## Checkpoint history

| Time | Event | Commit |
|---|---|---|
| 2026-08-29 00:35 +08:00 | Initialized the four-suite rerun plan. | Pending |
