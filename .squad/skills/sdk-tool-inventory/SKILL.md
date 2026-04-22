---
name: sdk-tool-inventory
description: How to introspect what tools/skills/MCP servers actually loaded into a Copilot SDK session, and how to compare against the declared config to detect silent load failures.
---

# SDK Tool Inventory Introspection

## When to use this skill

- Adding guardrails or validators that need to know what tools the agent **actually** had available (not just what the config declared).
- Building diagnostics for "the eval ran but the agent never used the MCP server" symptoms.
- Surfacing tool/skill/MCP load status in reports.
- Writing tests that need to assert tool availability.

## Where the SDK exposes it

The Copilot Go SDK emits structured events at session start, before the model receives the prompt. Two event types matter:

| Event type | Data field | Contains |
|---|---|---|
| `copilot.SessionEventTypeSessionSkillsLoaded` | `event.Data.Skills []Skill` | Each `Skill` has `Name` (string). One event per session. |
| `copilot.SessionEventTypeSessionMcpServersLoaded` | `event.Data.Servers []Server` | Each `Server` has `Name` (string). One event per session. |

These events are recorded into `report.SessionEventRecord` in `hyoka/internal/eval/copilot.go` (around line 343 for skills, 355 for MCP). The handler currently extracts `Name` and joins into `rec.Content` as a comma-separated string, then stores via `evalReport.SessionEvents`.

There is **no** dedicated "tools loaded" event — the SDK exposes available tool names indirectly via `copilot.SessionEventTypeSessionToolsUpdated` (which currently just logs and doesn't capture details).

## Reconstructing the loaded inventory after a session

For post-session analysis (graders, guardrails, trend reports), parse the `SessionEvents` slice on `EvalReport`:

```go
var loadedSkills, loadedMCP []string
for _, ev := range evalReport.SessionEvents {
    switch ev.Type {
    case "session.skills_loaded":
        if ev.Content != "" {
            loadedSkills = strings.Split(ev.Content, ", ")
        }
    case "session.mcp_servers_loaded":
        if ev.Content != "" {
            loadedMCP = strings.Split(ev.Content, ", ")
        }
    }
}
```

The exact string forms of `ev.Type` come from the SDK constants — prefer the constants over magic strings when in engine code (`copilot.SessionEventTypeSessionSkillsLoaded.String()` or equivalent), but the recorded `ev.Type` values in `SessionEvents` are the SDK's serialized form.

`engine_eval.go` line 327-329 already does this for `env.SkillsLoaded`:

```go
case "session.skills_loaded":
    env.SkillsLoaded = strings.Split(ev.Content, ", ")
```

## Building the "expected" set from config

The declared/expected set comes from `task.Config.Generator.Tools` after attribute resolution:

```go
toolEntries := config.ResolveTools(task.Config.Generator.Tools, promptProps)
expectedSkills := map[string]bool{}
expectedMCP := map[string]bool{}
for _, entry := range toolEntries {
    switch entry.ResolvedType() {
    case "skill":
        expectedSkills[entry.Name] = true  // or use Path/Repo when Name empty
    case "mcp":
        expectedMCP[entry.Name] = true
    case "plugin":
        // plugins resolve into multiple skills — see internal/plugin
        for _, skill := range entry.ResolvedPluginSkills() {
            expectedSkills[skill.Name] = true
        }
    }
}
```

`copilot.go` already builds `expectedMCPServers map[string]bool` this way at session start (search for the variable name) and uses it for a load-success warning at lines 366-370.

## Comparison pattern

```go
missing := []string{}
for name := range expectedSkills {
    if !contains(loadedSkills, name) { missing = append(missing, "skill:"+name) }
}
for name := range expectedMCP {
    if !contains(loadedMCP, name) { missing = append(missing, "mcp:"+name) }
}
```

Prefix entries with their kind (`skill:`, `mcp:`) so error messages and `MissingTools` field entries are unambiguous when both kinds fail in one run.

## Gotchas

- **Skill name vs path vs repo.** Local skills have `Name` set explicitly; remote skills may have `Name` derived from `Path` or `Repo`. The SDK reports the resolved name. Match using the same resolution `engine_eval.go:374-382` uses when building `setup.Skills`.
- **Plugin expansion happens before SDK sees skills.** Plugins are expanded into individual skills in the config layer (`internal/plugin`), so the `expectedSkills` set must include the plugin's resolved skills, not the plugin name itself. The SDK never sees "plugin" as a category.
- **MCP server may register but report zero tools.** A loaded MCP server is necessary but not sufficient — a server that loaded but failed handshake might emit zero tool descriptors. For strict enforcement, also check that the loaded server contributed at least one tool (parse `SessionToolsUpdated` events or check `event.Data.Servers[].Tools` length if exposed).
- **Event ordering is not guaranteed across SDK versions.** Skills-loaded and mcp-servers-loaded events typically fire before any `assistant.message`, but don't assume strict ordering for guardrail timing — wait until after generation completes for the comparison, or build a small state machine in the event handler.
- **Empty `event.Data.Skills`/`Servers` is meaningful.** When the slice is empty but config declared tools, that's the failure condition — `copilot.go:351-354` and `:372-374` already log warnings on this case.

## Related issues

- #219 — added `SessionSetup` field for visibility (configured tools, status="configured")
- #347 — added `expectedMCPServers` warning logic in `copilot.go`
- #348 — `ToolAvailability` summary (available vs used)
- #490 — visibility tooling in action timeline
- #619 — promotes the warnings to hard-fail enforcement (this skill is the primer for #619)
