# Java suite health

Status: **In progress**

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
