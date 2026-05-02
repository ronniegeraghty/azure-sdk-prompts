# Decisions

## D-2026-05-02-001: SDK Tool Naming Collision Rules Clarified

**Date:** 2026-05-02T02:59:28Z  
**Decision Authority:** Switch (Research)  
**Status:** ✅ RECORDED  

### Issue

Grader design (Morpheus's multi-source tool loading) depends on understanding Copilot SDK tool naming collision rules. Can we rely on SDK to prevent or detect collisions across skills, MCP servers, and built-ins?

### Finding

**The Copilot SDK does NOT enforce tool name uniqueness across skills, MCP servers, or against built-ins.** 

SDK only validates collisions for custom SDK tools (via `OverridesBuiltInTool` flag). All other cross-source collisions are unvalidated by the SDK.

### Impact

- **Grader tool loading design** must not assume SDK-side safety
- **Explicit namespacing** or **load-time validation** required at CLI level
- **Scoped tool registry** recommended: `mcp:servername/toolname`, `skill:skillname/toolname`

### Implication for Morpheus

When Morpheus designs Grader multi-source tool loading, he must:
1. Assume no SDK-side collision guarantee
2. Implement one of: namespacing, load-time validation, or scoped registry
3. Document tool resolution order for hyoka graders

### Decision

Record finding. Morpheus owns namespacing design during Grader implementation. Precedent: GitHub CLI namespaces extensions.

---

**Sources:**
- Orchestration log: `2026-05-02T02-59-28Z-switch.md`
- Session log: `2026-05-02T02-59-28Z-sdk-tool-naming.md`
