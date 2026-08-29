# JS/TS Event Hubs Azure Skill + MCP retry health

Status: **Completed but not selected**

- Prompt: `event-hubs-dp-js-ts-streaming`
- Variant: Azure Skill + MCP
- Retry run ID: `20260829-185622`
- SDK errors: 1 session-idle timeout
- Generated files: 4
- Tool calls: 9
- Azure MCP calls: 4 succeeded, 0 failed
- Azure MCP average latency: 8.7 seconds
- Azure MCP maximum latency: 12.1 seconds
- Skill-load failures: 0
- Orphaned Copilot processes terminated by post-run cleanup: 0

The retry generated the expected four-file TypeScript project and every Azure MCP call completed successfully, but generation again timed out waiting for `session.idle`. It does not meet the healthy-retry selection policy, so the retry is retained but not selected. The full Event Hubs prompt triplet is marked for exclusion from final comparisons.

A review grader later returned malformed JSON, but the generation session-idle timeout already made the retry ineligible for selection.
