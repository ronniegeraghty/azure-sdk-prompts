# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Core Context

Agent Scribe initialized as session logger for hyoka. Maintains decisions.md, orchestration logs, session logs, and cross-agent context.

## Recent Updates

📌 Team initialized on 2026-04-03

## Learnings

Initial setup complete.

### Inbox-to-Decisions Merge Patterns (2026-04-17)

**Handling Phase 4 Kickoff Brief:**
1. **Condensation:** Morpheus's full 181-line brief (detailed architecture, per-agent guidance, risk register rationale) was summarized to ~70 lines in decisions.md, capturing:
   - Phase 4 North Star + dependency graph (critical path + parallel tracks)
   - Key decisions (3 major items: #566 Option A, Trinity workload, architectural contract)
   - Per-agent launch (bullet points, 4 agents)
   - Risk register (5 items, condensed rows)
   - Go-live gates (7 criteria)
2. **Cross-reference:** Condensed entry references archived full brief at `.squad/decisions/archive/morpheus-phase4-kickoff.md` for readers needing full rationale.
3. **Chronological ordering:** New decision inserted at top of "## Active Decisions" section (newest-first convention).
4. **Session logging:** Phase 4 fan-out recorded to `.squad/sessions/2026-04-17.md` with:
   - Phase 3 completion details (PR #562 commit, hotfix integration merge chain)
   - Phase 4 kickoff context (Morpheus delivery, Ronnie approval)
   - Agent assignments and model selections
   - Risk profile and go-live criteria
5. **Archive pattern:** Full-text briefs stored in `.squad/decisions/archive/` with filename matching inbox document; condensed summary in main decisions.md acts as index.
6. **Commit discipline:** Squad changes committed directly to `ronniegeraghty/dev` (no PR) with full trailers. Session logs in `.squad/sessions/` are gitignored (local tracking only); decisions and archives are committed.

**Patterns observed from other inbox files:**
- Previous decisions deduped by checking decisions.md for existing entries before merging
- Archive convention established: inbox files moved to archive after merge into decisions
- Session logs use daily files (YYYY-MM-DD.md) to track events, not timestamped JSON

**Future application:**
- When merging large briefing documents, always condense to 60–80 lines in decisions.md
- Use 3–4 key decision bullets + per-agent bullets for clarity
- Archive full text for traceability
- Cross-link archived doc from decisions.md summary


## 2026-04-17: Phase 4 Verified — Ready for v0.3.1 Release

Morpheus 🕶️ completed Phase 4 dogfood verification (6/6 checks PASSED, zero blockers). All subsystems verified: build, live eval, comparison auto-generation, serve endpoints, hierarchical criteria, cleanup. Recommendation: **Promote dev → main and cut v0.3.1 tag.**

Decision: .squad/decisions.md | Orchestration Log: .squad/orchestration-log/2026-04-17T20:53:40Z-morpheus.md
