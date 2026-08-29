# .NET managed-identity retry health

Status: **Completed but not selected**

- Prompt: `identity-dp-dotnet-managed-identity`
- Variant: Azure Skill + MCP + Microsoft Skills
- Retry run ID: `20260829-183008`
- SDK errors: 0
- Generated files: 0
- Tool calls: 7
- Azure MCP calls: 8 succeeded, 0 failed
- Azure MCP average latency: 5.5 seconds
- Azure MCP maximum latency: 9.1 seconds
- Skill-load failures: 0
- Orphaned Copilot processes terminated by post-run cleanup: 2

The retry removed the original session-idle timeout but still produced no generated files. It does not meet the healthy-retry selection policy, so the retry is retained but not selected. The full prompt triplet is marked for exclusion from final comparisons.
