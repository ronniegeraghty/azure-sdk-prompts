---
name: "group-based-comparison-ui"
description: "Pattern for building flexible N-way comparison UIs over data with arbitrary attributes — groups defined by composable filters, configurable visualizations, localStorage persistence."
domain: "frontend, ux, data-viz"
confidence: "high"
source: "earned (#365 / WI-047 — Compare page redesign, PR #601)"
---

## Context

Use this skill when designing any UI that lets users compare arbitrary subsets of a dataset (e.g. eval results, telemetry, A/B cohorts) and visualize differences. Beats the typical "pick A vs B" picker because:

- It scales to N groups, not just 2.
- It composes — a "group" can be filtered on many attributes at once.
- It survives schema growth — adding a new attribute means adding it to the catalog, not redesigning the page.

## Patterns

### 1. Separate the model from the renderer

Put the group model, filter logic, and metric computation in a pure-function lib (`lib/<feature>-groups.ts`). The page is a thin orchestrator that:
- Loads raw data once.
- Holds groups + chart selection in component state.
- Delegates filtering/metrics to the lib.

This shape pays off in test density (lib tests are fast and don't need DOM) and reusability (the model can be embedded in other surfaces).

### 2. Filter semantics: OR within, AND across

```ts
filters: { languages: ["python", "go"], configs: ["baseline"] }
// → (lang ∈ {python, go}) AND (config ∈ {baseline})
```

Empty array (or missing key) = "match all" for that dimension. Document this prominently — it's the part users get wrong.

### 3. Build the catalog from real data

Walk the loaded dataset once and extract the distinct values per dimension. Don't list values that don't appear in the data — that's how you get phantom filter chips with zero matches.

```ts
function buildCatalog(rows: Row[]): Catalog {
  const sets = { lang: new Set(), service: new Set(), ... };
  for (const r of rows) {
    if (r.lang) sets.lang.add(r.lang);
    // ...
  }
  return { lang: [...sets.lang].sort(), ... };
}
```

### 4. Group identity = name + color

Each group gets a stable color from a fixed palette. Reuse that color across summary cards, chart bars, and legend entries — it becomes the user's mental "shorthand" for the group. For single-metric bar charts in recharts, use `<Cell>` to color per-bar:

```tsx
<Bar dataKey="value">
  {data.map((d, i) => <Cell key={i} fill={d.color} />)}
</Bar>
```

For grouped breakdowns, use one `<Bar>` per group with `fill={group.color}` — that gives you a free legend.

### 5. Persist with versioned key + shape validation

```ts
const STORAGE_KEY = "myapp:comparison:v1";

function load(storage = window.localStorage): State | null {
  try {
    const raw = storage.getItem(STORAGE_KEY);
    if (!raw) return null;
    const parsed = JSON.parse(raw);
    // Filter out malformed entries — never throw.
    return { groups: parsed.groups.filter(isValidGroup), ... };
  } catch {
    return null;
  }
}
```

`saveState` should also `try/catch` (quota / privacy mode). The `:v1` suffix gives a migration path; don't skip it.

### 6. Surface every edge case explicitly

Don't render an empty chart and hope the user figures it out. Each of these deserves a banner:
- 0 groups → CTA to create one
- 1 group → hint: "add another to compare"
- All groups empty → warning: filters too narrow
- Some groups empty → warning + skip those bars
- Wildly uneven group sizes → caveat: interpret with care

### 7. Configurable visualizations, not a fixed layout

Let users toggle which charts they care about. Default to a sensible 2–3 chart subset. Each chart should describe itself (label + one-line description) — it doubles as a tooltip for unfamiliar metrics.

## Examples

- `site/src/app/lib/comparison-groups.ts` — model + filtering + metrics + persistence (~290 LOC)
- `site/src/app/components/group-builder.tsx` — single-group editor UI
- `site/src/app/components/comparison-page.tsx` — orchestrator + chart panels
- `site/src/__tests__/comparison-groups.test.ts` — 21 lib tests demonstrating the test discipline this pattern enables

## Anti-Patterns

- **Hardcoded N=2 comparison.** Don't. The moment a user wants 3 cohorts, you're rewriting.
- **Computed metrics inside React components.** Move them to the lib so they're testable without rendering.
- **Schema-driven UI without a catalog.** Letting users type filter values in a free-text box leads to typos and zero-match groups. Always offer chips from the catalog.
- **Silent localStorage failures.** Wrap save in try/catch. Wrap load in try/catch + shape validation. Never let one corrupt entry break the page.
- **Same chart color for every group.** Defeats the entire purpose of side-by-side viz.
