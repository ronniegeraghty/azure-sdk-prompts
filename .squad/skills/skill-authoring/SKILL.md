---
name: skill-authoring
description: How to author SKILL.md files that conform to the agentskills.io spec and actually trigger when relevant. Use when creating, reviewing, or fixing skills under skills/, .squad/skills/, or anywhere a SKILL.md is being authored — especially when a skill "exists but never gets used" (almost always a description problem).
---

# Skill Authoring (agentskills.io spec)

Source of truth: <https://agentskills.io/specification.md> and <https://agentskills.io/skill-creation/optimizing-descriptions.md>.

## Frontmatter — what's actually in the spec

Only **two** fields are required:

| Field | Required | Notes |
|-------|----------|-------|
| `name` | Yes | 1-64 chars, lowercase `a-z`, digits, hyphens. No leading/trailing/consecutive hyphens. **Must match parent directory name.** |
| `description` | Yes | 1-1024 chars. Describes what the skill does AND when to use it. |
| `license` | No | License name or filename. |
| `compatibility` | No | Env requirements (≤500 chars). Most skills don't need it. |
| `metadata` | No | Arbitrary string→string map for vendor-specific data. |
| `allowed-tools` | No | Experimental. Space-separated tool allow-list. |

**Anything else is non-spec.** Common offenders seen in this repo:

- `applyTo: "**/*.md"` — VS Code / Copilot-Chat field, NOT in agentskills.io spec. Strip it or move it under `metadata:` if a specific client needs it. It does nothing for agentskills.io activation.
- Missing frontmatter entirely (e.g., `skills/reviewer/java-sdk-validation/SKILL.md`) — the skill cannot be discovered or activated.

## How activation actually works (progressive disclosure)

1. **Discovery:** agent loads only `name` + `description` from every available skill.
2. **Activation:** when a user task matches a `description`, the agent loads the full `SKILL.md` body.
3. **Execution:** agent follows the body's instructions.

The whole burden of triggering is on the `description`. A bad description = the body never loads.

Important nuance from the docs: **agents tend to skip skills for tasks they think they can already handle.** For "obvious" tasks (like "write a markdown file with a heading"), descriptions must be **pushy** — explicitly telling the agent to apply the skill even when it thinks it doesn't need help.

## Writing descriptions that trigger

- **Imperative phrasing.** "Use this skill when …" beats "This skill provides guidance on …".
- **Both halves: what + when.** State the capability AND name the trigger contexts.
- **List indirect triggers.** Include phrases the user is likely to use without naming the domain. E.g., for a markdown-list skill: "even when the user only says 'three items' or 'numbered steps'."
- **Be pushy on simple tasks.** Explicitly say "apply this even for one-shot tasks / single-element documents / short prompts" so the agent doesn't optimize the skill away.
- **Concise.** A few sentences to a short paragraph. Hard cap is 1024 chars.

### Pattern

```
description: <One-line capability statement>. Use this skill whenever <primary trigger>, including <indirect trigger phrases>, and apply it before <action> so <desirable outcome on first pass>.
```

### Before / after

Bad (vague, passive, no triggers):
```yaml
description: "Guidance on writing proper markdown headings. Use when generating markdown files to ensure correct heading hierarchy."
```

Good (specific, imperative, lists indirect triggers, pushy):
```yaml
description: Authoritative rules for writing markdown headings (H1/H2/H3) with correct hierarchy, exactly one H1 per document, no skipped levels, and blank lines around each heading. Use this skill whenever you are about to write or edit a markdown (.md) file that contains any heading — even single-heading documents, even one-shot tasks, and even when the user only says "add a title" or "level-1 heading". Apply it before writing the file.
```

## Body structure conventions

The body loads only after activation, so spend tokens on what the agent **wouldn't already know**:

1. Short opening line stating scope.
2. Hard rules (bulleted, unambiguous).
3. One or two correct examples in fenced code blocks.
4. A "common mistakes to avoid" section — failure modes the model is likely to produce by default.

Skip generic background ("PDFs are documents that…"). Skip vague advice ("handle errors appropriately"). Project-specific conventions, edge cases, and exact tool/API choices are where skills earn their tokens.

## Quick checklist before shipping a skill

- [ ] `name` matches the parent directory exactly.
- [ ] `description` uses imperative voice and names both the capability and trigger contexts.
- [ ] `description` includes indirect/pushy triggers for tasks the agent might otherwise handle without help.
- [ ] No non-spec frontmatter fields (e.g., `applyTo`) leaking in — move client-specific keys under `metadata:`.
- [ ] Body is rules + examples + anti-patterns, not background reading.
- [ ] If activation matters, write a few realistic prompts and mentally check whether they'd match the description.
