# Decision: Run-level filter system (#600)

**Author:** Trinity 🖤
**Date:** 2026-04-21
**Status:** Proposed (in PR)

## Context

The runs page (`/runs`) listed every evaluation run as a flat scrolling list.
As the eval library grew, finding "all runs that touched Python" or "all runs
that used the azure-mcp config" became tedious. R146/R147 (#600) asked for a
run-level filter UI with URL persistence.

## Decision

Filter at the **run** level (not per-eval). A run matches when every active
filter dimension finds at least one matching eval inside that run. This
preserves the runs page's identity (a list of runs) while letting users narrow
to runs they care about.

### Semantics

- **Within a dimension:** OR (multi-select).
- **Across dimensions:** AND.
- **Empty dimension:** match-all (consistent with `group-based-comparison-ui`).
- **Status:** derived from run aggregate — `errors > 0` → `errors`, else
  `failed > 0` → `failing`, else `passing`. Errors take precedence so a run
  with both errors and failures only appears under `errors`.

### Module layout

- `site/src/app/lib/run-filters.ts` — pure model (catalog, matching, URL
  ser/deserialize). No React, fast unit tests.
- `site/src/app/components/ui/multi-select-filter.tsx` — reusable chip
  dropdown primitive (outside-click + Escape close, accessible labels).
- `site/src/app/components/runs-page.tsx` — `<FilterBar>` composes the three
  multi-selects, holds active count + reset.

### URL persistence

`useSearchParams` is the source of truth. Filter changes call
`setSearchParams(..., { replace: true })` so we don't pollute history. Param
keys are stable: `config`, `lang`, `status`, comma-joined.

## Alternatives considered

- **React Context for filter state** — rejected; URL-as-state means reload +
  share work for free and there's only one consumer.
- **Per-eval filtering on the runs page** — rejected; would require flattening
  runs into eval cards which changes the page identity. The run-detail page
  already does per-eval filtering for that purpose.
- **Server-side filtering** — overkill for current data sizes. Client-side
  filtering is instant on hundreds of runs. Revisit if list grows past ~5k.

## Consequences

- New reusable primitive (`MultiSelectFilter`) other pages can adopt
  (prompts page, dashboard).
- Pure-function lib pattern (extracted from `comparison-groups.ts`) is now the
  default for filter logic in the site. See `group-based-comparison-ui` skill.
