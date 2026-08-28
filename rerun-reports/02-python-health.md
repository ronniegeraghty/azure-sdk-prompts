# Python suite health

Status: **In progress**

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
