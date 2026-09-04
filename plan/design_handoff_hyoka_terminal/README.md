# Handoff: hyoka CLI terminal output

## Overview
This handoff covers the redesigned terminal output for the `hyoka` CLI — the interactive display shown while `hyoka run` is executing evals and the final summary printed when the run completes.

The design covers the full lifecycle:
1. **plan & queue** — printed once at startup
2. **running** — live, updates as evals progress; works identically for 1 worker or N workers
3. **failure detail** — surfaced inline when a grader fails
4. **final summary** — collapsed one-line recap of every eval, plus `next` action hints
5. **verbose mode** (`--verbose`) — tools tree expanded; per-turn generator agent stream

## About the design files
The files in this bundle are **design references created in HTML**. They are NOT production code to copy — they simulate the terminal output in a browser with React so the rhythm, colors, spacing and live states are visible and interactive.

Your job is to **implement this design in the hyoka Go CLI** using a real TUI library — Bubble Tea (`github.com/charmbracelet/bubbletea`) + Lipgloss (`github.com/charmbracelet/lipgloss`) are recommended. The HTML is the spec for how the output should look; use the tokens, symbols, and structure described in this README to drive the Go implementation.

## Fidelity
**High-fidelity.** All colors, glyphs, column widths, and live-update mechanics are specified exactly. Reproduce the layout as written. Minor adjustments to accommodate terminal width reflow are expected.

## Runtime environment
- **Language:** Go
- **CLI entrypoint:** `go run ./hyoka run …`
- **Recommended libraries:**
  - `github.com/charmbracelet/bubbletea` — Elm-style model/update/view loop
  - `github.com/charmbracelet/lipgloss` — styling and color
  - `github.com/charmbracelet/bubbles/spinner` — spinner component (braille frames match the design)
  - `github.com/charmbracelet/bubbles/progress` — progress bar
  - `github.com/mattn/go-isatty` — detect CI / non-TTY
- **Rendering mode:** inline (NOT alt-screen). Use `tea.Println` to emit completed blocks as permanent scrollback; keep only the "live tail" (progress footer + active-block spinners) as the Bubble Tea view. This is critical — completed blocks must remain in scrollback with full detail.

## Design Tokens

### Colors (ANSI 256 / truecolor)
All colors are lipgloss-compatible hex; fall back to the nearest ANSI-256 automatically.

| Token              | Hex       | Usage                                    |
|--------------------|-----------|------------------------------------------|
| `fg`               | `#d4d7dc` | default foreground, values, labels       |
| `fg-dim`           | `#8a8f98` | secondary text, descriptions             |
| `fg-dimmer`        | `#565b64` | chrome (rules, bullets, separators)      |
| `fg-bright`        | `#ffffff` | bold emphasis (`b` class)                |
| `green`            | `#7ec77a` | pass, ✓, filled bar segments             |
| `green-bg`         | `rgba(126,199,122,0.08)` | (N/A in terminal — ignore) |
| `red`              | `#e35f5f` | fail, ✗, failure blocks                  |
| `red-bg`           | `rgba(227,95,95,0.10)`   | failure block background (use lipgloss `Background` on red-tinted block) |
| `yellow`           | `#e0b457` | running, spinner, worker count, "generating"/"grading" phase words |
| `blue`             | `#6ca4d8` | (reserved)                               |
| `magenta`          | `#c381c0` | (reserved)                               |
| `cyan`             | `#5fb3b3` | file paths, actionable commands in `next` |
| `rule`             | `#22262d` | (mostly unused in terminal; use `fg-dimmer` for rules) |

**Respect `NO_COLOR` and `CLICOLOR=0`** — fall back to glyph-only output with no ANSI escapes.

### Glyphs

| Symbol | Meaning |
|---|---|
| `✓` | pass |
| `✗` | fail |
| `·` | queued / inactive / bullet |
| `▸` | turn/step marker (verbose) |
| `│` | vertical separator inside a line |
| `↳` | failure lead-in |
| `↗` | external link hint |
| `↓` `↑` | token direction (input/output) |

### Spinner frames (braille)
```
⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏
```
Interval: **~90ms per frame** (matches the mockup). The `bubbles/spinner` package ships `spinner.Dot` which matches.

### Bars
- **Progress bar (bottom of screen):** `█` filled, `░` empty. Width dynamically sized to terminal width minus label/padding (~36 chars is a good default). Three-color fill: `green` done, `yellow` running, `dimmer` remaining.
- **Per-grader pip bar:** `▰` filled, `▱` empty. One pip per criterion. Color: `green` if grader passed, `red` if failed.
- **Summary ratio bar:** same `█` as progress, but two-color only (`green` passed / `red` failed).

### Section rule
```
── label · trailing ─────────────────────────────
```
- Two em-dashes + space before the label
- Space + `·` + space + trailing text (dimmed) if present
- Padded with `─` to fill line to a reasonable width (~60–70 cols, or to `termWidth - 4`)
- `fg-dimmer` for all rule chars; `fg-bright` bold for the label; `fg-dim` for trailing

### ASCII fallback (`--ascii`)
When the user passes `--ascii` or the terminal doesn't support the glyphs, swap:
- `✓` → `[ok]`
- `✗` → `[x]`
- `·` → `.`
- `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏` → rotating `|/-\`
- `█▰` → `#`
- `░▱` → `-`
- `─│↳↗▸` → `-|->!>`

---

## Screens

### 1. Plan block (startup, printed once)

**Purpose:** Confirm what's about to run. Prints immediately after argument parsing.

**Layout:**
```
$ hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise --workers 3

── plan  ──────────────────────────────────────────────────────
  prompts 1   configs 3   evals 3   workers 3
  prompt  key-vault-dp-python-crud  python · CRUD · azure-keyvault

── queue  ─────────────────────────────────────────────────────
  ·  gpt-5.3-codex       queued
  ·  claude-opus-4.6     queued
  ·  claude-sonnet-4.5   queued

starting 3 workers…
```

**Rules:**
- First line: `$ ` in `fg-dimmer`, command in `fg`, args in `fg-dim`
- Key-value row: `prompts`/`configs`/`evals`/`workers` labels in `fg-dimmer`, values in `fg`. `evals` value is **bold** (`fg-bright`). `workers` value is `yellow`.
- Three spaces between k/v pairs.
- Queue rows: two-space indent, `·` in `fg-dimmer`, label in `fg`, status `queued` in `fg-dimmer`. Pad labels to a common width.

### 2. Live running view

**Purpose:** Shows every eval's progress. Completed evals stay **fully expanded** in scrollback (emitted via `tea.Println`); only actively-running blocks are re-rendered each tick.

**Block structure** (applies to every eval, queued → generating → grading → pass/fail):
```
<sym>  <config-label>   ·  w<N>   │  <phase>   ·   t <elapsed>   ·   files <name>
     turns 4  calls 4  tok 179.3k↓ 1,460↑  tools azure(mcp 1/1), azure-sdk-python(40/40)
     <grader-sym>  <grader-name>       [<pips>]  <passed>/<total>
     <grader-sym>  <grader-name>       [<pips>]  <passed>/<total>
     [optional failure block — see §3]
```

**Line 1 — header:**
- 2-space indent NOT used on line 1; the symbol is at column 0
- Symbol: ✓/✗/spinner/· (colored by phase)
- Two spaces, then label (bold, `fg-bright`)
- `   ·  ` separator (spaces for rhythm)
- `w<N>` in `fg-dim` (only when multi-worker)
- `   │  ` separator
- Phase word: `generating`/`grading`/`pass`/`fail`/`queued` (colored)
- `   ·   ` separator, then `t` key (`fg-dimmer`) + value (`fg`) showing elapsed
- `   ·   ` separator, then `files` key + comma-separated filenames in `cyan` (only if files produced)

**Line 2 — generator stats** (shown whenever phase != `queued`):
- 5-space indent
- `turns N  calls N  tok <in>↓ <out>↑  tools azure(mcp X/Y), azure-sdk-python(N/M)`
- Labels in `fg-dimmer`, values in `fg-dim` or `fg`
- Token counts formatted as `179.3k↓` — use `k`/`M` suffixes with 1 decimal place
- **Default (non-verbose):** tools are collapsed to counts like `azure(mcp 1/1), azure-sdk-python(40/40 skills)`. If a plugin exposes both skills and mcp, show both: `(40/40 skills · 1/1 mcp)`.

**Grader rows** (one per grader attached to the eval):
- 5-space indent
- Symbol: ✓ (green) / ✗ (red) / spinner (yellow) / · (dimmer)
- Two spaces, name padded to 18 chars
- Pip bar `[▰▰▱▱]` — one pip per criterion, color matches pass/fail state
- Two spaces, trailing text:
  - if `queued` → `queued` (dimmer)
  - if `running` → `grading…` (yellow)
  - if `pass` → `N/N` (dimmer)
  - if `fail` → `N/N` (red)

**Worker assignment display:**
- Each active block shows `w<N>` indicating which worker picked it up. When a worker frees up and grabs the next queued eval, a new block starts below.

### 3. Failure block

Rendered inline under a failed eval's graders (not as a separate section).

**Layout:**
```
     ↳ output_files  2 rules failed
  ✗ min_files             need ≥ 1 produced file, found 0
  ✗ min_bytes_per_file    no produced files to check (≥ 1 required)
  hint: agent completed 3 turns but never called `create`
```

**Rules:**
- Background-tinted with `red-bg` + left border of `red` (2 cols). Use lipgloss `Border` + `Background`.
- Lead line: `↳ <grader-name>` in bold-red, followed by dimmer `<N> rules failed`
- Each failed rule: `✗` red, name padded to 22 chars, reason in `fg-dim`
- Optional `hint:` line at the bottom in `fg-dimmer`

### 4. Progress footer

Always at the bottom of the live view; updates in place.

```
  ████████████░░░░░░░░░░░░░░░░░░░░░░░░  1/3   ·   workers 3   elapsed 2:10   eta ~1:20
```

**Rules:**
- Bar width scales with terminal width (target ~36 chars for standard 80-col)
- Three segments: `green` (done) + `yellow` (running) + `dimmer` (queued)
- `done/total` after the bar — `done` in `fg`, `/total` in `fg-dim`
- Remaining k/v pairs: label in `fg-dimmer`, value in `fg`
- `workers` value in `yellow`
- Elapsed format: `MM:SS` for < 1h, `H:MM:SS` for >= 1h

### 5. Final summary

Printed when all evals complete. Uses **collapsed one-line blocks** — this is the ONLY place collapsed form appears.

```
── summary · run 20260424-034027  ─────────────────────

✓  claude-opus-4.6      pass   10/10 criteria   1:58   4t · 179k tok   kv_secrets_crud.py
✓  claude-sonnet-4.5    pass   12/12 criteria   1:35   3t · 136k tok   README.md, keyvault_crud.py, …
✗  gpt-5.3-codex        fail   8/10 criteria    1:12   3t · 103k tok

  ████████████████████████████████████   2/3 passed   ·   elapsed 4:46

── next  ──────────────────────────────────────────────
  view    hyoka view 20260424-034027         ↗ open report in browser
  trend   hyoka trend --last 10              AI analysis across recent runs
  retry   hyoka run --rerun-failed 20260424-034027  rerun only the failed eval

  report  reports/20260424-034027/
```

**Collapsed block format:**
```
<sym>  <label padded 20>  <phase>   <N/N criteria>   <elapsed>   <turns>t · <tok>k tok   <files>
```

**`next` section:**
- Three one-line actions. Each: indent 2, label in `fg-dim`, command in `cyan`, description in `fg-dimmer`.
- The commands are literal shell commands the user can copy-paste.
- Final `report` line points to the on-disk report directory.

### 6. `--verbose` mode

Adds two expansions when `--verbose` is passed:

**6a. Tool tree** (replaces the collapsed `tools azure(mcp 1/1), azure-sdk-python(40/40)` line):
```
     tools
       ├─ azure                 mcp           ✓ loaded
       ├─ generator-skills      skills dir    1/1 skills
       │  └─ azure-sdk-for-rust-bestpractices
       └─ azure-sdk-python      plugin        40/40 skills  ·  1/1 mcp
          ├─ azure-keyvault-py              ├─ azure-identity-py
          ├─ azure-storage-blob-py          ├─ azure-cosmos-py
          ├─ azure-servicebus-py            ├─ azure-eventhub-py
          ├─ fastapi-router-py              ├─ pydantic-models-py
          └─ … 32 more  (hyoka show tools --expand)
```
- Tree-drawing chars in `fg-dimmer`
- Plugin/dir names in `fg`, kind (mcp/skills dir/plugin) in `fg-dimmer`, loaded-count in `green`
- Individual skills in `fg-dim`
- Two-column layout for large skill lists — pad left column to ~34 chars
- Truncate after 8 entries with `… N more (hyoka show tools --expand)`

**6b. Per-turn agent stream** (under the generator header):
```
     ▸ turn 1   skill        azure-keyvault-py              ok   6.2s
     ▸ turn 2   create       kv_secrets_crud.py             ok   2.1 KB
     ▸ turn 3   bash         python -m py_compile kv_secrets_crud.py  ok
     ⠋ turn 4   thinking…
```
- `▸` dimmer for completed turns; spinner for the active turn
- `turn N` in `fg-dim`
- Tool name in `cyan`, padded to ~12 chars
- Target/argument in `fg` for file paths, `fg-dim` for command strings
- Result (`ok`/elapsed/size) right-aligned in `green` or `fg-dimmer`

Below the turns, a `graders` block shows what's queued:
```
·  graders   │  waiting for generator
     ·  output_files       queued   checks min_files · min_bytes_per_file
     ·  prompt_criteria    queued   6 criteria from prompt file
```

---

## State machine (per eval)

```
queued → generating → grading → pass | fail
```

| State | Transition trigger |
|---|---|
| `queued` → `generating` | worker picks up the eval |
| `generating` → `grading` | generator agent session ends (success or agent error) |
| `grading` → `pass` | all graders return pass |
| `grading` → `fail` | any grader returns fail |

**Live updates:**
- While `generating`: generator stats line ticks (elapsed, turns, calls, tokens update on each agent event)
- While `grading`: the active grader row has the spinner symbol and `grading…` trailing; others in the same eval wait with `·` + `queued`
- On `pass`/`fail`: the block is frozen (emit via `tea.Println`) and a new active block for the next queued eval appears below

---

## Implementation notes for Bubble Tea

### Suggested file layout
```
internal/tui/
  tui.go           // top-level Model + Update + View; Program entrypoint
  eval.go          // per-eval sub-model (state, elapsed timer, stats)
  render.go        // pure rendering: plan, block (expanded + collapsed), failure, progress, summary, next
  styles.go        // lipgloss styles keyed off the Design Tokens table
  glyphs.go        // glyph constants + ASCII fallbacks
  verbose.go       // verbose-only rendering (tool tree, turn stream)
```

### Key patterns
- **Completed blocks as scrollback:** when an eval resolves, call `m.Println(renderBlock(eval))` on the returned `tea.Cmd` — this emits the block above the live view so it survives in terminal history. Then remove the eval from the active-model list.
- **Single ticker:** one global `tick` message at ~90ms drives all spinners. Don't give each eval its own ticker.
- **Width reflow:** on `tea.WindowSizeMsg`, recompute bar widths and pad columns. If width < 80, drop the `tools` column and shorten stat keys (`tok` → `t`, `calls` → `c`).
- **Non-TTY fallback:** check `isatty.IsTerminal(os.Stdout.Fd())` at startup. If false, bypass Bubble Tea entirely and use plain line-buffered printing: one line per state transition, no spinners, no ANSI.

### Event ingestion
The CLI already emits events from the eval runner (you can see the existing verbose log has session details, turn counts, tool calls). Wire those into Bubble Tea messages:
- `EvalStartedMsg{evalID, worker, label}`
- `GeneratorTurnMsg{evalID, turn, tool, target, result, duration}`
- `GeneratorDoneMsg{evalID, stats}`
- `GraderStartedMsg{evalID, graderName}`
- `GraderDoneMsg{evalID, graderName, passed, total, failures}`
- `EvalDoneMsg{evalID, phase}`

---

## Files in this bundle

- `Terminal B Developed.html` — the primary mockup. Open in any browser. Five artboards on a zoomable canvas: plan & queue, mid-flight, longer-queue, final summary, verbose.
- `terminal.css` — design tokens and utility classes used by the mockup.
- `spinner.jsx` — spinner hook used by the mockup's React components (reference for timing).
- `terminal-b-primitives.jsx` — reusable block/grader/progress primitives.
- `terminal-b-scenes.jsx` — the five scene components.
- `design-canvas.jsx` — canvas shell (not needed for the Go impl; just for viewing).

## Acceptance checklist

- [ ] Plan block prints once at startup with correct styling
- [ ] Running view shows one block per eval, per-worker `w<N>` label
- [ ] Spinner animates at ~90ms braille cycle
- [ ] Generator stats line updates in place while the eval is generating
- [ ] Grader rows transition queued → running → pass/fail correctly
- [ ] Failure block renders with 2-col red left-border and lists every failed rule
- [ ] Progress footer shows green/yellow/dimmer segments and accurate done/total
- [ ] Completed blocks stay fully expanded in scrollback (not collapsed mid-run)
- [ ] Final summary prints collapsed one-liners + ratio bar + `next` commands
- [ ] `--verbose` expands the tools tree and per-turn stream
- [ ] `--ascii` swaps all special glyphs to ASCII equivalents
- [ ] `NO_COLOR` / non-TTY disables all ANSI escapes
- [ ] Output reflows on terminal resize (columns dropped on narrow widths)
