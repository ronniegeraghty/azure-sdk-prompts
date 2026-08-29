# JS/TS Event Hubs Baseline retry health

Status: **Completed and selected**

- Prompt: `event-hubs-dp-js-ts-streaming`
- Variant: Baseline
- Retry run ID: `20260829-184916`
- SDK errors: 0
- Generated files: 4
- Tool calls: 3
- Azure MCP calls: not applicable to Baseline
- Skill-load failures: 0
- Orphaned Copilot processes terminated by post-run cleanup: 0

The retry reached `session.idle`, completed without an SDK error, and generated the expected TypeScript project with `package.json`, `tsconfig.json`, `README.md`, and `src/index.ts`. It meets the healthy-retry selection policy and replaces the primary Baseline result.

Dependency installation was blocked by an invalid npm authentication token, so the generated project was not compiled. This was not an SDK, session, test, or MCP timeout.
