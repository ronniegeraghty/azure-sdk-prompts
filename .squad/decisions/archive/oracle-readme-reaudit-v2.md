# Oracle: README re-audit v2 (executed-command validation)

**Supersedes:** prior `oracle-readme-audit` decision (commit `9931af2c`), which verified commands by reading source rather than executing them.

## Branch-divergence finding

`origin/main` is 3 commits ahead of `phase-5` / `ronniegeraghty/dev`:

| Commit | Subject | Impact |
|---|---|---|
| `9f293cee` | Move main.go into hyoka/ (#569) | `main.go` now lives at `hyoka/main.go`. Correct invocation: `go run ./hyoka <cmd>` (NOT `go run .`). |
| `8e8ae1fc` | docs: fix outdated commands in README.md (#575) | Already established `--config baseline/claude-opus-4.6` on run examples, replaced `hyoka tools` → `hyoka plugins`, fixed structure path. |
| `a0a78426` | Add example remote-skill config (#573) | n/a for README. |

On `phase-5`, `main.go` is still at the repo root, so `go run .` works *locally*. But README must be correct for the **destination layout** (`main` post-merge), not the in-flight checkout. Validation was therefore performed in a worktree on `origin/main`:

```
git worktree add ../hyoka-worktrees/main-audit origin/main
```

(Used `../hyoka-worktrees/` instead of `/tmp/` per runtime policy that forbids `/tmp` writes.)

## Commands tested (all executed in `origin/main` worktree)

| # | Command | Exit | Result | Notes |
|---|---|---:|---|---|
| 1 | `go build ./hyoka/...` | 0 | ✅ pass | Clean build on `origin/main`. |
| 2 | `go build .` (repo root) | — | ❌ FAIL | "no Go files in /…/main-audit" — confirms `go build .` is wrong on main. |
| 3 | `go test -race ./...` | 0 | ✅ pass | All packages green (~30s). Recursive form covers everything. |
| 4 | `go run ./hyoka list` | 0 | ✅ pass | "Found 89 prompt(s)". |
| 5 | `go run ./hyoka list --service key-vault --language python` | 0 | ✅ pass | "Found 4 prompt(s)". |
| 6 | `go run ./hyoka run --service storage --config baseline/claude-opus-4.6 --dry-run` | 0 | ✅ pass | Plan: 24 evals × 1 config. |
| 7 | `go run ./hyoka run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6 --dry-run` | 0 | ✅ pass | Quick-start parity (dry-run substitute for live invocation). |
| 8 | `go run ./hyoka serve --help` | 0 | ✅ pass | Subcommand exists. |
| 9 | `go run ./hyoka version` | 0 | ✅ pass | "hyoka version 0.3.0". |
| 10 | `go run ./hyoka init --help` | 0 | ✅ pass | |
| 11 | `go run ./hyoka validate --help` | 0 | ✅ pass | |
| 12 | `go run ./hyoka check-env` | 0 | ✅ pass | |
| 13 | `go run ./hyoka clean --help` | 0 | ✅ pass | |
| 14 | `go run ./hyoka compare --help` | 0 | ✅ pass | |
| 15 | `go run ./hyoka plugins` | 0 | ✅ pass | Listed `azure-python` plugin. |

Doc-link existence (on `phase-5`, where the README ships from):

| Path | Exists? |
|---|---|
| `docs/prompt-authoring.md` | ✅ |
| `docs/configuration.md` | ✅ |
| `docs/grader-config-schema.md` | ✅ |
| `docs/cli-reference.md` | ✅ |
| `docs/architecture.md` | ✅ |
| `docs/guardrails.md` | ✅ |
| `docs/roadmap.md` | ✅ (added on phase-5; not on main) |
| `CONTRIBUTING.md` | ✅ (added on phase-5 by `ed62da30`; not on main) |
| `LICENSE` | ✅ |

> Note: `docs/roadmap.md` and `CONTRIBUTING.md` do not exist on `main` today, but they will exist when `phase-5` merges, since both are tracked on the branch. Acceptable — README ships *with* phase-5.

## Diff applied (phase-5 README.md)

All `go run .` → `go run ./hyoka`; `go build .` → `go build ./hyoka/...`. One copy-tweak for scope-guard (replaced "Generated output" with "Captured the agent's output" in the post-quick-start narrative).

```
-go run . list  # Verify installation
+go run ./hyoka list  # Verify installation

-go run . list --service key-vault --language python
+go run ./hyoka list --service key-vault --language python

-go run . run \
+go run ./hyoka run \

-go run . serve
+go run ./hyoka serve

-3. Generated output based on the prompt
+3. Captured the agent's output

-go build .
+go build ./hyoka/...

-go run . run --prompt-id key-vault-dp-python-crud \
+go run ./hyoka run --prompt-id key-vault-dp-python-crud \
```

PR #575's earlier fixes (presence of `--config baseline/claude-opus-4.6` on every run example, no references to `hyoka tools`, `hyoka/main.go` path) were confirmed already incorporated and were preserved.

## Scope guard

- Confirmed README contains no "code generation" / "code-gen" framing. Neutral phrasing ("AI agents", "agent's output") used throughout.
- No `site/` files touched (issue #364 lockout respected).

## README.backup

✅ Deleted (`git rm README.backup`). 22 KB leftover from `bcd4e541` (Trinity's pre-#368 backup). Did not wait for issue #593.

## Handoff

- Branch: `phase-5`
- Commit: see commit message `docs(#368): re-audit README with executed-command validation, delete README.backup`
- Prior `oracle-readme-audit` decision is **superseded** by this artifact.
