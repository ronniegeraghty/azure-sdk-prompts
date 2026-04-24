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

### Session 2026-04-20 (Phase 5 Coordination & Orchestration)

**Role:** Coordinator (Ronnie's session managing Phase 5)

**Phase 5 Mission:** Docs & Polish on new per-phase integration branch workflow

**Workflow Innovation:**
- Departed from per-issue PRs
- Introduced shared `phase-5` integration branch (off `ronniegeraghty/dev`)
- Owners work in parallel, merge directly (no per-issue PRs)
- Shared review on-branch (Switch)
- Live verification gate (Morpheus)
- Single rollup PR when all done

**Issues Scheduled:**
- Wave 1: #364, #367, #369 + TDD tests
- Wave 2: #366, #368
- Deferred: #365 (A/B compare, XL scope, Phase 6)

**Execution Summary:**
- All 5 issues delivered to `phase-5`
- One escalation (Trinity/Oracle locked out on #364, Morpheus rescued)
- Switch review cycle: 3 clean approvals, 2 re-reviews
- Morpheus live verification: All UI features working
- Rollup PR #592 opened

**Outcomes:**
- ✅ 5 issues approved and merged to `phase-5`
- ✅ Reviewer-protocol enforcement successful
- ✅ Live verification gate established
- ✅ Per-phase workflow validated
- ⏸️ #365 deferred to Phase 6 (scope management)

**Phase 5 Status:** COMPLETE. Rollup PR #592 open for Ronnie's review.

**Key Learning:** Per-phase workflow with shared integration branch + live verification gate significantly accelerates delivery compared to per-issue PRs. Reviewer-protocol escalation chain worked as designed: locked agents step back, escalation brings fresh eyes, root causes fixed instead of worked around.

## Team Updates

### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint closed at HEAD `2d38533f` on `ronniegeraghty/dev`. 15 commits total across three rounds, 48 new test cases.

**Your merges this sprint:** `60cc90bf` round-2 merge (schema + emitters consolidation, 5 inbox entries) · final wrap merge 2026-04-23T00:05:04Z (6 inbox entries including 2 bug reports: one resolved, one moved to Known Issues).

**Key reconciliation:** Updated the round-1/2 table in `.squad/decisions.md` to reflect that `82cd8590` (Neo's original ToolsVerified emission) never merged; behavior re-landed via Switch's `25ce00a7`. Added a Known Issues section containing the preexisting `hyoka clean` non-interactive-stdin bug.

See `.squad/orchestration-log/2026-04-23T00-05-04Z-sprint-wrap.md`.

---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.

### 2026-04-23T19:42Z: Decision-log discipline — extend, don't SUPERSEDE, when a follow-up adds a contract rather than reversing one

When commit `4a8c4a0d` landed (container plugin fan-out), the temptation was to mark the prior `repo:`-required entry as superseded. Wrong instinct. The earlier decision fixed the *locator* shape; this one fixes the *content* shape. Both contracts now hold simultaneously. Rule: only mark SUPERSEDED when the new decision *replaces* or *reverses* the earlier one. When it adds an orthogonal contract, the entry should explicitly say "extends, does not supersede" and cross-reference the prior commit. Readers a year later need to know both rules still apply.

---

## 2026-04-24: Session Complete — Orchestration, Decision Merge, History Updates ✅

**Session:** 2026-04-24T05:58:18Z  
**Status:** ✅ Complete

Merged feature-shipped state across team:

- ✅ Orchestration logs: Switch (ff38a7ec) + Coordinator merge (4 commits)
- ✅ Session log: prompt-grader-checks feature shipped
- ✅ Decision merge: Switch's tool-grader decision + work-on-dev directive → decisions.md (inbox deleted)
- ✅ Cross-agent history: Feature notes appended to Switch, Tank, Neo, Morpheus + work-on-dev directive to all active agents
- ✅ Ready for git commit of .squad/ state

**Work-on-dev directive:** All active agents notified of new branch strategy (no transient feature branches for squad work — merge to dev).


---

## 2026-04-24: 🚨 Team default model is now claude-opus-4.7

Per `.squad/config.json` (`defaultModel: claude-opus-4.7`) and the standing policy at the top of `.squad/decisions.md`:

- **Every agent spawn defaults to `claude-opus-4.7`.**
- **`claude-haiku-4.5` is FORBIDDEN.** Even if your charter says "preferred: claude-haiku-4.5", that line is overridden. No Haiku, ever.
- **`claude-sonnet-4.5`** (latest Sonnet) is allowed only for trivial mechanical work where opus-4.7 would be wasteful.
- This affects what every future spawn looks like — expect opus-4.7 as your model.

---

## 2026-04-24: Oracle Cleanup — Unimplemented Examples → Issues

**Decision:** Ronnie → Oracle: delete unimplemented doc sketches, file as GitHub issues  
**Outcome:** ✅ COMPLETED

**Execution:**
- Deleted `examples/prompts/graders-frontmatter-example.prompt.md`
- Removed "Option B" section from `docs/starter-files.md`
- Removed graders-frontmatter refs from `examples/README.md`
- Filed **#636**: graders frontmatter feature
- Filed **#637**: starter_files list feature

**Commits:** `09cf17f0` (cleanup) · `a41616fe` (Oracle history)  
**Branch:** ronniegeraghty/dev

**Rationale:** Documentation debt — unimplemented features belong in issues, not in shipped examples. Examples now match parser reality.

**Decision archive:** `.squad/decisions/archive/oracle-unimplemented-examples-to-issues.md`
