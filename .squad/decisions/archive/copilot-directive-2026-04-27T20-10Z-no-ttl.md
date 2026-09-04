### 2026-04-27T20:10Z: User directive
**By:** Ronnie (via Copilot)
**What:** For unpinned (`default` version) remote skills/plugins, do NOT use TTL-based freshness — keep `git fetch` every run. Reason: people testing skills under active development need updates to land immediately, not after a TTL window.
**Why:** User request — captured for team memory.
