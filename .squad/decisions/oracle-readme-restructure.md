# Decision: README.md Restructure (#368)

**Date:** 2026-04-20  
**Agent:** Oracle 🔮  
**Issue:** #368 — WI-055: README.md Restructure  
**Status:** Implemented and merged to phase-5

## Context

README.md was a 540-line monolith containing duplicate content, inline config details, verbose flag tables, and stale sections (roadmap, repo tree). It served as both an entry point and a comprehensive manual, making it overwhelming for new users and hard to maintain.

## Decision

Restructured README.md into a focused 6-section document (229 lines) that serves as an **entry point** directing users to detailed docs rather than duplicating them.

### New structure:
1. **Hero section** — what hyoka is, installation (source + CLI), quick start (5-minute scenario)
2. **Examples** — sample prompt, config, criteria with links to authoring guides
3. **Commands** — brief table + filtering examples, link to CLI reference for full flag docs
4. **Safety & Guardrails** — condensed summary (generator guardrails, safety boundaries, protections), link to detailed guardrails doc
5. **Contributing** — points to CONTRIBUTING.md, architecture docs, quick dev loop
6. **License** — MIT

### Content moved:
- **Repo tree** → removed (already in AGENTS.md from #367)
- **Inline config details** → docs/configuration.md
- **Verbose flag tables** → docs/cli-reference.md
- **Tagging system** → docs/prompt-authoring.md
- **Roadmap** → docs/roadmap.md (new file)

### Command invocation fix:
- Changed all examples from `go run ./hyoka` to `go run .`
- Rationale: main.go is at repo root, not in hyoka/ subdirectory
- Verified all commands work as documented

## Rationale

1. **Reduces cognitive load** — new users see what hyoka does and how to run it in < 2 minutes, not 540 lines
2. **Single source of truth** — detailed docs live in docs/, README links to them (no duplication drift)
3. **Maintainability** — when flags change, only docs/cli-reference.md updates; README stays stable
4. **Examples are real** — uses actual prompt IDs (key-vault-dp-python-crud) and config names (baseline/claude-opus-4.6) that exist in the repo

## Alternatives considered

- **Option A: Keep comprehensive README, remove docs/** — Rejected: docs/ provides structured deep-dives (architecture.md, guardrails.md) that don't belong in a README
- **Option B: Split into README.md + QUICKSTART.md** — Rejected: adds cognitive overhead ("which file do I read first?")
- **Option C: Move everything to docs/, README is just a link** — Rejected: users expect README to show basic usage, not just redirect

## Impact

- **Positive:** Faster onboarding (verified 5-minute quick start), easier maintenance (single source of truth for flags/config)
- **Neutral:** Users seeking deep detail must follow links to docs/ (expected behavior for a README)
- **Risk mitigation:** All examples verified to work; all links verified to point to existing files

## Follow-up

- **Docs sync:** Ensure docs/cli-reference.md, docs/configuration.md, docs/prompt-authoring.md stay accurate as features evolve
- **Version sync:** When cutting releases, verify README examples still use available configs/prompts
