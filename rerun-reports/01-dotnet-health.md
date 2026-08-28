# .NET suite health

Run ID: `20260829-004156`

Status: **In progress**

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
