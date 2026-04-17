import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import {
  fetchRuns,
  fetchRun,
  fetchEval,
  fetchDocs,
  fetchDoc,
  fetchPrompts,
  fetchPrompt,
  fetchCompareConfigs,
  fetchPairwise,
} from "../app/data/api";

describe("API module", () => {
  const originalFetch = globalThis.fetch;

  beforeEach(() => {
    globalThis.fetch = vi.fn();
  });

  afterEach(() => {
    globalThis.fetch = originalFetch;
  });

  function mockFetchOk(data: unknown) {
    vi.mocked(globalThis.fetch).mockResolvedValue({
      ok: true,
      json: () => Promise.resolve(data),
    } as Response);
  }

  function mockFetchError(status: number, statusText: string) {
    vi.mocked(globalThis.fetch).mockResolvedValue({
      ok: false,
      status,
      statusText,
    } as Response);
  }

  // ── fetchRuns ──────────────────────────────────────────────────────

  it("fetchRuns calls the correct endpoint", async () => {
    mockFetchOk([]);
    const result = await fetchRuns();
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/runs");
    expect(result).toEqual([]);
  });

  it("fetchRuns throws on non-ok response", async () => {
    mockFetchError(500, "Internal Server Error");
    await expect(fetchRuns()).rejects.toThrow("API error 500: Internal Server Error");
  });

  // ── fetchRun ───────────────────────────────────────────────────────

  it("fetchRun calls the correct endpoint with encoded id", async () => {
    const mock = { run_id: "run-1" };
    mockFetchOk(mock);
    const result = await fetchRun("run-1");
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/runs/run-1");
    expect(result).toEqual(mock);
  });

  it("fetchRun encodes special characters", async () => {
    mockFetchOk({});
    await fetchRun("run/with spaces");
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/runs/run%2Fwith%20spaces");
  });

  // ── fetchEval ──────────────────────────────────────────────────────

  it("fetchEval builds the correct URL with query params", async () => {
    mockFetchOk({ prompt_id: "test" });
    await fetchEval("run-1", "some/path.json");
    expect(globalThis.fetch).toHaveBeenCalledWith(
      "/api/runs/run-1/eval?path=some%2Fpath.json"
    );
  });

  // ── fetchDocs ──────────────────────────────────────────────────────

  it("fetchDocs calls the correct endpoint", async () => {
    mockFetchOk([{ slug: "intro", title: "Intro" }]);
    const result = await fetchDocs();
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/docs");
    expect(result).toHaveLength(1);
  });

  // ── fetchDoc ───────────────────────────────────────────────────────

  it("fetchDoc calls the correct endpoint with slug", async () => {
    mockFetchOk({ slug: "intro", title: "Intro", content: "# Hello" });
    const result = await fetchDoc("intro");
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/docs/intro");
    expect(result.content).toBe("# Hello");
  });

  // ── fetchPrompts ───────────────────────────────────────────────────

  it("fetchPrompts calls the correct endpoint", async () => {
    mockFetchOk([]);
    await fetchPrompts();
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/prompts");
  });

  // ── fetchPrompt ────────────────────────────────────────────────────

  it("fetchPrompt calls the correct endpoint with id", async () => {
    mockFetchOk({ id: "p1" });
    await fetchPrompt("p1");
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/prompts/p1");
  });

  // ── fetchCompareConfigs ────────────────────────────────────────────

  it("fetchCompareConfigs encodes both config names", async () => {
    mockFetchOk({ kind: "configs", label_a: "a", label_b: "b", per_prompt: [], summary: {} });
    await fetchCompareConfigs("baseline/opus", "azure-mcp/opus");
    expect(globalThis.fetch).toHaveBeenCalledWith(
      "/api/compare/configs?a=baseline%2Fopus&b=azure-mcp%2Fopus"
    );
  });

  it("fetchCompareConfigs throws on server error", async () => {
    mockFetchError(404, "Not Found");
    await expect(fetchCompareConfigs("a", "b")).rejects.toThrow("API error 404: Not Found");
  });

  // ── fetchPairwise ──────────────────────────────────────────────────

  it("fetchPairwise calls the correct endpoint", async () => {
    mockFetchOk({ run_id: "run-1", timestamp: "t", reports: [] });
    const result = await fetchPairwise("run-1");
    expect(globalThis.fetch).toHaveBeenCalledWith("/api/runs/run-1/pairwise");
    expect(result).toEqual({ run_id: "run-1", timestamp: "t", reports: [] });
  });

  it("fetchPairwise returns null on 404 (run has no pairwise data)", async () => {
    vi.mocked(globalThis.fetch).mockResolvedValue({
      ok: false,
      status: 404,
      statusText: "Not Found",
    } as Response);
    const result = await fetchPairwise("run-without-pairwise");
    expect(result).toBeNull();
  });

  it("fetchPairwise throws on non-404 server errors", async () => {
    mockFetchError(500, "Internal Server Error");
    await expect(fetchPairwise("run-1")).rejects.toThrow("API error 500: Internal Server Error");
  });
});
