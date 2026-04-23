# Plugin tool-load assertions check leaves, not parents

**Date:** 2026-04-23
**By:** Neo

**Decision:** When a remote plugin has the standard Copilot container layout (`<plugin>/skills/<child>/SKILL.md`), the validator MUST fan out into one report row per child skill. Tool-load assertions then check that each child loaded — never the parent plugin directory, which has no SKILL.md of its own and will never appear in the SDK's `SessionSkillsLoaded` event.

**Why:** The microsoft/skills `azure-sdk-python` plugin is a directory of 41 child skills, not a single skill. Asserting "did the plugin load?" by looking for the plugin's name in the SDK's loaded-skills list would always fail: the SDK loads the children, by their child basenames. The fan-out is what makes the assertion meaningful.

**Scope:** Applies to `plugin.ResolveInstalled` + `validatePluginEntry`. Single-skill plugins (top-level SKILL.md) keep the one-row-per-plugin behavior. Container plugins fan out.

**Reference:** Fix commit, plus `TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren` for the regression guard.
