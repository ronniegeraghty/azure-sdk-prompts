import { describe, it, expect } from "vitest";
import {
  applyFilters,
  applyFiltersToSearchParams,
  buildCatalog,
  deriveFacets,
  deriveStatus,
  EMPTY_FILTERS,
  filtersFromSearchParams,
  hasActiveFilters,
  runMatches,
  toggleValue,
} from "../app/lib/run-filters";
import type { RunSummary, EvalResult } from "../app/data/types";

function evalResult(opts: {
  config?: string;
  language?: string;
  success?: boolean;
}): EvalResult {
  return {
    prompt_id: "p",
    config_name: opts.config ?? "baseline/claude-opus-4.6",
    success: opts.success ?? true,
    review: { overall_score: 0, max_score: 0, summary: "" },
    duration_seconds: 0,
    prompt_metadata: {
      service: "identity",
      plane: "data-plane",
      language: opts.language ?? "python",
      category: "auth",
      difficulty: "easy",
    },
  };
}

function run(opts: {
  id: string;
  passed?: number;
  failed?: number;
  errors?: number;
  results?: EvalResult[];
}): RunSummary {
  const results = opts.results ?? [];
  return {
    run_id: opts.id,
    timestamp: "2026-04-01T00:00:00Z",
    total_evaluations: results.length,
    passed: opts.passed ?? results.length,
    failed: opts.failed ?? 0,
    errors: opts.errors ?? 0,
    duration_seconds: 1,
    results,
  };
}

describe("deriveStatus", () => {
  it("flags errors before failures", () => {
    expect(deriveStatus(run({ id: "a", failed: 1, errors: 1 }))).toBe("errors");
  });
  it("returns failing when any failed and no errors", () => {
    expect(deriveStatus(run({ id: "a", passed: 1, failed: 1 }))).toBe("failing");
  });
  it("returns passing when no failures or errors", () => {
    expect(deriveStatus(run({ id: "a", passed: 2 }))).toBe("passing");
  });
});

describe("deriveFacets", () => {
  it("collects unique configs and languages from results", () => {
    const r = run({
      id: "a",
      results: [
        evalResult({ config: "c1", language: "python" }),
        evalResult({ config: "c2", language: "go" }),
        evalResult({ config: "c1", language: "python" }),
      ],
    });
    const f = deriveFacets(r);
    expect([...f.configs].sort()).toEqual(["c1", "c2"]);
    expect([...f.languages].sort()).toEqual(["go", "python"]);
  });

  it("handles missing results array", () => {
    const r: RunSummary = {
      run_id: "x",
      timestamp: "",
      total_evaluations: 0,
      passed: 0,
      failed: 0,
      errors: 0,
      duration_seconds: 0,
      results: undefined as unknown as EvalResult[],
    };
    const f = deriveFacets(r);
    expect(f.configs.size).toBe(0);
    expect(f.languages.size).toBe(0);
  });
});

describe("buildCatalog", () => {
  it("returns sorted unique values across runs and only present statuses", () => {
    const runs = [
      run({
        id: "a",
        passed: 1,
        results: [evalResult({ config: "c2", language: "python" })],
      }),
      run({
        id: "b",
        passed: 0,
        failed: 1,
        results: [evalResult({ config: "c1", language: "go", success: false })],
      }),
    ];
    const cat = buildCatalog(runs);
    expect(cat.configs).toEqual(["c1", "c2"]);
    expect(cat.languages).toEqual(["go", "python"]);
    expect(cat.statuses).toEqual(["passing", "failing"]);
  });
});

describe("runMatches & applyFilters", () => {
  const runs = [
    run({
      id: "py-pass",
      results: [evalResult({ config: "baseline", language: "python" })],
    }),
    run({
      id: "go-fail",
      passed: 0,
      failed: 1,
      results: [evalResult({ config: "azure-mcp", language: "go", success: false })],
    }),
    run({
      id: "mixed",
      passed: 1,
      failed: 1,
      results: [
        evalResult({ config: "baseline", language: "python" }),
        evalResult({ config: "azure-mcp", language: "dotnet", success: false }),
      ],
    }),
  ];

  it("returns all runs when no filters set", () => {
    expect(applyFilters(runs, EMPTY_FILTERS)).toHaveLength(3);
  });

  it("filters by language with OR semantics within dimension", () => {
    const out = applyFilters(runs, { ...EMPTY_FILTERS, languages: ["go", "dotnet"] });
    expect(out.map((r) => r.run_id)).toEqual(["go-fail", "mixed"]);
  });

  it("ANDs across dimensions: config AND status", () => {
    const out = applyFilters(runs, {
      configs: ["baseline"],
      languages: [],
      statuses: ["passing"],
    });
    expect(out.map((r) => r.run_id)).toEqual(["py-pass"]);
  });

  it("matches when run contains at least one eval per dimension", () => {
    // mixed has both a baseline+python eval AND an azure-mcp+dotnet eval.
    // Filtering for baseline + dotnet matches because each dimension matches
    // at the run level (any-of), even though no single eval has both.
    expect(
      runMatches(runs[2], {
        configs: ["baseline"],
        languages: ["dotnet"],
        statuses: [],
      }),
    ).toBe(true);
  });

  it("excludes runs whose status doesn't match", () => {
    const out = applyFilters(runs, {
      configs: [],
      languages: [],
      statuses: ["failing"],
    });
    expect(out.map((r) => r.run_id)).toEqual(["go-fail", "mixed"]);
  });
});

describe("hasActiveFilters & toggleValue", () => {
  it("hasActiveFilters reports correctly", () => {
    expect(hasActiveFilters(EMPTY_FILTERS)).toBe(false);
    expect(hasActiveFilters({ ...EMPTY_FILTERS, configs: ["x"] })).toBe(true);
  });
  it("toggleValue adds and removes", () => {
    expect(toggleValue<string>([], "a")).toEqual(["a"]);
    expect(toggleValue<string>(["a", "b"], "a")).toEqual(["b"]);
  });
});

describe("URL serialization", () => {
  it("round-trips filters through URLSearchParams", () => {
    const original = {
      configs: ["baseline/claude-opus-4.6", "azure-mcp/claude-opus-4.6"],
      languages: ["python"],
      statuses: ["failing"] as const,
    };
    const params = applyFiltersToSearchParams(new URLSearchParams(), {
      ...original,
      statuses: [...original.statuses],
    });
    const decoded = filtersFromSearchParams(params);
    expect(decoded.configs).toEqual(original.configs);
    expect(decoded.languages).toEqual(original.languages);
    expect(decoded.statuses).toEqual(["failing"]);
  });

  it("removes empty dimensions from URL", () => {
    const params = new URLSearchParams("config=x&lang=python&status=passing");
    applyFiltersToSearchParams(params, EMPTY_FILTERS);
    expect(params.toString()).toBe("");
  });

  it("ignores unknown statuses", () => {
    const params = new URLSearchParams("status=failing,bogus");
    const decoded = filtersFromSearchParams(params);
    expect(decoded.statuses).toEqual(["failing"]);
  });
});
