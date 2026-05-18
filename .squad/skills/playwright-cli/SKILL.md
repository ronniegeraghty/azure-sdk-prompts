---
name: playwright-cli
description: Automate browser interactions, test web pages and work with Playwright tests via the playwright-cli command.
allowed-tools: Bash(playwright-cli:*) Bash(npx:*) Bash(npm:*)
confidence: medium
source: https://github.com/microsoft/playwright-cli/blob/main/skills/playwright-cli/SKILL.md
---

# Browser Automation with playwright-cli

The `playwright-cli` is installed globally on this machine (`@playwright/cli`). Use it for live browser-based verification of the hyoka site (`go run . serve`).

## Quick start

```bash
# open browser at a URL
playwright-cli open http://localhost:8080
# get snapshot of accessibility tree (refs like e1, e2, e3...)
playwright-cli snapshot
# interact with the page using refs from the snapshot
playwright-cli click e15
playwright-cli fill e5 "user@example.com" --submit
playwright-cli press Enter
# take a screenshot
playwright-cli screenshot --filename=verify.png
# close browser
playwright-cli close
```

After EVERY command, playwright-cli prints a snapshot YAML file path. Read that file (or use `playwright-cli snapshot` again) to see the current page state and find element refs (`e1`, `e2`, ...) for the next interaction.

## Core commands

```bash
playwright-cli open [url]            # opens browser, optionally navigates
playwright-cli goto <url>            # navigate
playwright-cli snapshot              # get a11y snapshot with refs
playwright-cli snapshot --filename=foo.yml   # save to specific file
playwright-cli click <ref>           # click element by ref
playwright-cli fill <ref> "text"     # fill input
playwright-cli type "text"           # type at current focus
playwright-cli press Enter           # press a key
playwright-cli hover <ref>           # hover element
playwright-cli select <ref> "value"  # select option
playwright-cli eval "expr"           # eval JS, returns result
playwright-cli eval "el => el.textContent" e5    # eval on element
playwright-cli screenshot --filename=foo.png     # screenshot
playwright-cli screenshot e5 --filename=elem.png # element screenshot
playwright-cli resize 1920 1080      # resize viewport
playwright-cli go-back / go-forward / reload
playwright-cli close                 # close browser
```

## Targeting elements

Default: refs from snapshots (e15, e23). Also supports CSS, role locators, test ids:
```bash
playwright-cli click "#main > button.submit"
playwright-cli click "getByRole('button', { name: 'Submit' })"
playwright-cli click "getByTestId('submit-button')"
```

## Raw output (for piping/diffing)

```bash
playwright-cli --raw eval "document.title"
playwright-cli --raw snapshot > before.yml
playwright-cli click e5
playwright-cli --raw snapshot > after.yml
diff before.yml after.yml
```

## Tabs / sessions

```bash
playwright-cli tab-new https://example.com
playwright-cli tab-list
playwright-cli tab-select 0
playwright-cli tab-close

# Named browser sessions for parallel work
playwright-cli -s=mysession open url
playwright-cli -s=mysession click e6
playwright-cli -s=mysession close
playwright-cli list                  # list sessions
playwright-cli close-all             # close all
```

## DevTools

```bash
playwright-cli console               # browser console messages
playwright-cli console warning       # filter level
playwright-cli network               # network log
playwright-cli tracing-start / tracing-stop
```

## Verification recipe (hyoka site)

```bash
# Start server in another shell: cd /home/rgeraghty/projects/hyoka && go run . serve
# Default port is 8080 unless --port specified.

playwright-cli open http://localhost:8080
playwright-cli snapshot --filename=/tmp/home.yml
# Verify page title
playwright-cli --raw eval "document.title"
# Verify logo SVG present
playwright-cli --raw eval "document.querySelector('svg[aria-label]')?.getAttribute('aria-label')"
# Navigate to an eval detail page (find a link first via snapshot, then click)
playwright-cli click eN              # using ref from snapshot
playwright-cli screenshot --filename=eval-detail.png
playwright-cli close
```

## Installation (already done on this machine)

```bash
npm install -g @playwright/cli@latest    # provides playwright-cli command
```

If `playwright-cli` isn't on PATH, fall back to `npx playwright-cli`.
