import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { ComparisonPage } from "../app/components/comparison-page";

// Mock the API module so ComparisonPage doesn't call the real backend
vi.mock("../app/data/api", () => ({
  fetchRuns: vi.fn(),
  fetchCompareConfigs: vi.fn(),
}));

import { fetchRuns, fetchCompareConfigs } from "../app/data/api";

const mockRuns = [
  {
    run_id: "run-1",
    timestamp: "2026-03-29T14:00:00Z",
    total_evaluations: 4,
    passed: 3,
    failed: 1,
    errors: 0,
    duration_seconds: 60,
    results: [
      { prompt_id: "identity-dp-python-auth", config_name: "baseline/claude-opus-4.6", success: true, review: { overall_score: 9, max_score: 10, summary: "" }, duration_seconds: 10, prompt_metadata: { service: "identity", plane: "data-plane", language: "python", category: "auth", difficulty: "basic" } },
      { prompt_id: "storage-dp-python-crud", config_name: "azure-mcp/claude-opus-4.6", success: true, review: { overall_score: 7, max_score: 10, summary: "" }, duration_seconds: 15, prompt_metadata: { service: "storage", plane: "data-plane", language: "python", category: "crud", difficulty: "intermediate" } },
    ],
  },
];

const mockComparison = {
  config_a: "baseline/claude-opus-4.6",
  config_b: "azure-mcp/claude-opus-4.6",
  per_prompt: [
    { prompt_id: "identity-dp-python-auth", score_a: 0.9, score_b: 0.95, delta: 0.05 },
    { prompt_id: "storage-dp-python-crud", score_a: 0.7, score_b: 0.65, delta: -0.05 },
  ],
  summary: { avg_delta: 0.0, improved: 1, regressed: 1, unchanged: 0 },
};

describe("ComparisonPage", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders the page heading after runs load", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <ComparisonPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText("Config Comparison")).toBeInTheDocument();
    });
  });

  it("shows loading spinner while fetching runs", () => {
    vi.mocked(fetchRuns).mockReturnValue(new Promise(() => {})); // never resolves

    render(
      <MemoryRouter>
        <ComparisonPage />
      </MemoryRouter>
    );

    // The loader is present while runs load
    expect(document.querySelector(".animate-spin")).toBeInTheDocument();
  });

  it("shows config selector labels after runs load", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <ComparisonPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText("Config A (baseline)")).toBeInTheDocument();
      expect(screen.getByText("Config B (comparison)")).toBeInTheDocument();
    });
  });

  it("shows empty state prompt before configs are selected", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <ComparisonPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(
        screen.getByText("Select two configurations above to see a side-by-side comparison.")
      ).toBeInTheDocument();
    });
  });
});
