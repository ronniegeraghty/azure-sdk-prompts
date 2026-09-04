import { describe, it, expect } from "vitest";
import {
  buildCatalog,
  computeMetrics,
  evalMatchesGroup,
  filterGroup,
  flattenResults,
  loadState,
  saveState,
  describeFilters,
  newGroupId,
  nextGroupColor,
  GROUP_COLORS,
  STORAGE_KEY,
  type ComparisonGroup,
} from "../app/lib/comparison-groups";
import type { EvalResult, RunSummary } from "../app/data/types";

function evalRow(over: Partial<EvalResult> = {}): EvalResult {
  return {
    prompt_id: "identity-dp-python-auth",
    config_name: "baseline/claude-opus-4.6",
    success: true,
    review: { overall_score: 9, max_score: 10, summary: "" },
    duration_seconds: 10,
    prompt_metadata: {
      service: "identity",
      plane: "data-plane",
      language: "python",
      category: "auth",
      difficulty: "basic",
    },
    ...over,
  };
}

function makeRun(results: EvalResult[]): RunSummary {
  return {
    run_id: "r",
    timestamp: "2026-04-21T00:00:00Z",
    total_evaluations: results.length,
    passed: results.filter((r) => r.success).length,
    failed: results.filter((r) => !r.success).length,
    errors: 0,
    duration_seconds: 0,
    results,
  };
}

const sampleResults: EvalResult[] = [
  evalRow({
    prompt_id: "identity-dp-python-auth",
    config_name: "baseline/claude-opus-4.6",
    success: true,
    review: { overall_score: 9, max_score: 10, summary: "" },
    prompt_metadata: { service: "identity", plane: "data-plane", language: "python", category: "auth", difficulty: "basic" },
  }),
  evalRow({
    prompt_id: "identity-dp-python-auth",
    config_name: "azure-mcp/claude-opus-4.6",
    success: true,
    review: { overall_score: 10, max_score: 10, summary: "" },
    prompt_metadata: { service: "identity", plane: "data-plane", language: "python", category: "auth", difficulty: "basic" },
  }),
  evalRow({
    prompt_id: "storage-dp-dotnet-crud",
    config_name: "baseline/claude-opus-4.6",
    success: false,
    review: { overall_score: 4, max_score: 10, summary: "" },
    prompt_metadata: { service: "storage", plane: "data-plane", language: "dotnet", category: "crud", difficulty: "intermediate" },
  }),
  evalRow({
    prompt_id: "storage-dp-dotnet-crud",
    config_name: "azure-mcp/claude-opus-4.6",
    success: true,
    review: { overall_score: 8, max_score: 10, summary: "" },
    prompt_metadata: { service: "storage", plane: "data-plane", language: "dotnet", category: "crud", difficulty: "intermediate" },
  }),
];

describe("flattenResults / buildCatalog", () => {
  it("flattens nested run.results", () => {
    const runs = [makeRun(sampleResults.slice(0, 2)), makeRun(sampleResults.slice(2))];
    expect(flattenResults(runs)).toHaveLength(4);
  });

  it("ignores runs without results", () => {
    const runs = [makeRun(sampleResults), { ...makeRun([]), results: undefined as unknown as EvalResult[] }];
    expect(flattenResults(runs)).toHaveLength(4);
  });

  it("builds catalog of distinct values across all dimensions, sorted", () => {
    const cat = buildCatalog(sampleResults);
    expect(cat.configs).toEqual(["azure-mcp/claude-opus-4.6", "baseline/claude-opus-4.6"]);
    expect(cat.services).toEqual(["identity", "storage"]);
    expect(cat.languages).toEqual(["dotnet", "python"]);
    expect(cat.categories).toEqual(["auth", "crud"]);
    expect(cat.difficulties).toEqual(["basic", "intermediate"]);
    expect(cat.planes).toEqual(["data-plane"]);
  });
});

describe("evalMatchesGroup / filterGroup", () => {
  const baseGroup: ComparisonGroup = { id: "g1", name: "G1", color: "#10b981", filters: {} };

  it("empty filters match all evals", () => {
    expect(filterGroup(sampleResults, baseGroup)).toHaveLength(4);
  });

  it("filters by config_name", () => {
    const g = { ...baseGroup, filters: { configs: ["baseline/claude-opus-4.6"] } };
    const filtered = filterGroup(sampleResults, g);
    expect(filtered).toHaveLength(2);
    expect(filtered.every((r) => r.config_name === "baseline/claude-opus-4.6")).toBe(true);
  });

  it("filters by language", () => {
    const g = { ...baseGroup, filters: { languages: ["python"] } };
    expect(filterGroup(sampleResults, g)).toHaveLength(2);
  });

  it("AND-combines multiple dimensions", () => {
    const g = { ...baseGroup, filters: { languages: ["python"], configs: ["baseline/claude-opus-4.6"] } };
    expect(filterGroup(sampleResults, g)).toHaveLength(1);
  });

  it("OR-combines values within a dimension", () => {
    const g = { ...baseGroup, filters: { languages: ["python", "dotnet"] } };
    expect(filterGroup(sampleResults, g)).toHaveLength(4);
  });

  it("evalMatchesGroup returns false when prompt_metadata missing the filtered field", () => {
    const r = evalRow({ prompt_metadata: undefined as any });
    const g = { ...baseGroup, filters: { services: ["identity"] } };
    expect(evalMatchesGroup(r, g)).toBe(false);
  });
});

describe("computeMetrics", () => {
  it("returns zero metrics for empty input", () => {
    const m = computeMetrics([]);
    expect(m.count).toBe(0);
    expect(m.pass_rate).toBe(0);
    expect(m.avg_score).toBe(0);
    expect(m.avg_duration).toBe(0);
    expect(m.score_distribution).toHaveLength(10);
    expect(m.score_distribution.every((b) => b === 0)).toBe(true);
  });

  it("computes pass_rate, avg_score, breakdowns", () => {
    const m = computeMetrics(sampleResults);
    expect(m.count).toBe(4);
    expect(m.pass_count).toBe(3);
    expect(m.pass_rate).toBeCloseTo(0.75);
    // (0.9 + 1.0 + 0.4 + 0.8) / 4 = 0.775
    expect(m.avg_score).toBeCloseTo(0.775);
    expect(m.by_service.identity).toEqual({ total: 2, passed: 2 });
    expect(m.by_service.storage).toEqual({ total: 2, passed: 1 });
    expect(m.by_language.python).toEqual({ total: 2, passed: 2 });
    expect(m.by_language.dotnet).toEqual({ total: 2, passed: 1 });
  });

  it("score distribution puts 1.0 in the top bin (not overflow)", () => {
    const r = evalRow({ review: { overall_score: 10, max_score: 10, summary: "" } });
    const m = computeMetrics([r]);
    expect(m.score_distribution[9]).toBe(1);
  });

  it("handles missing review max_score gracefully", () => {
    const r = evalRow({ review: { overall_score: 0, max_score: 0, summary: "" } });
    const m = computeMetrics([r]);
    expect(m.avg_score).toBe(0);
  });
});

describe("persistence", () => {
  function memStorage(): Storage {
    const map = new Map<string, string>();
    return {
      getItem: (k) => map.get(k) ?? null,
      setItem: (k, v) => void map.set(k, v),
      removeItem: (k) => void map.delete(k),
      clear: () => map.clear(),
      key: (i) => Array.from(map.keys())[i] ?? null,
      get length() {
        return map.size;
      },
    } as Storage;
  }

  it("round-trips groups + charts through localStorage", () => {
    const storage = memStorage();
    const groups: ComparisonGroup[] = [
      { id: "a", name: "A", color: "#10b981", filters: { languages: ["python"] } },
    ];
    saveState({ groups, charts: ["pass_rate", "avg_score"] }, storage);
    const loaded = loadState(storage);
    expect(loaded?.groups).toEqual(groups);
    expect(loaded?.charts).toEqual(["pass_rate", "avg_score"]);
  });

  it("returns null for empty storage", () => {
    expect(loadState(memStorage())).toBeNull();
  });

  it("returns null on malformed JSON", () => {
    const s = memStorage();
    s.setItem(STORAGE_KEY, "{not json");
    expect(loadState(s)).toBeNull();
  });

  it("drops malformed group entries", () => {
    const s = memStorage();
    s.setItem(
      STORAGE_KEY,
      JSON.stringify({
        groups: [{ id: "ok", name: "OK", color: "#fff", filters: {} }, { bad: true }],
        charts: ["pass_rate"],
      })
    );
    const loaded = loadState(s);
    expect(loaded?.groups).toHaveLength(1);
    expect(loaded?.groups[0].id).toBe("ok");
  });
});

describe("misc helpers", () => {
  it("nextGroupColor cycles through palette", () => {
    expect(nextGroupColor([])).toBe(GROUP_COLORS[0]);
    expect(nextGroupColor([{} as ComparisonGroup])).toBe(GROUP_COLORS[1]);
  });

  it("newGroupId returns unique-looking strings", () => {
    const a = newGroupId();
    const b = newGroupId();
    expect(a).not.toEqual(b);
    expect(a).toMatch(/^g-/);
  });

  it("describeFilters summarizes selected dimensions", () => {
    const g: ComparisonGroup = {
      id: "x",
      name: "X",
      color: "#000",
      filters: { languages: ["python"], services: ["identity"] },
    };
    const desc = describeFilters(g);
    expect(desc).toContain("Language: python");
    expect(desc).toContain("Service: identity");
  });

  it("describeFilters reports no filters", () => {
    const g: ComparisonGroup = { id: "x", name: "X", color: "#000", filters: {} };
    expect(describeFilters(g)).toMatch(/no filters/i);
  });
});
