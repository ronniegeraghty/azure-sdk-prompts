# JS/TS Event Hubs full-arm retry health

Status: **Completed and selected**

- Prompt: `event-hubs-dp-js-ts-streaming`
- Variant: Azure Skill + MCP + Microsoft Skills
- Retry run ID: `20260829-191313`
- SDK errors: 0
- Generated files: 4 reportable files, plus `package-lock.json`
- Tool calls: 6
- Azure MCP calls: 2 succeeded, 0 failed
- Azure MCP average latency: 6.2 seconds
- Azure MCP maximum latency: 9.9 seconds
- Skill-load failures: 0
- Dependency installation and TypeScript compilation: succeeded
- Orphaned Copilot processes terminated by post-run cleanup: 0

The retry reached `session.idle`, completed without an SDK error, generated the TypeScript project, and successfully installed dependencies and compiled it. It meets the healthy-retry selection policy and replaces the primary full-arm result.

The complete Event Hubs prompt triplet remains excluded from final comparisons because the Azure Skill + MCP arm was unresolved after its retry.
