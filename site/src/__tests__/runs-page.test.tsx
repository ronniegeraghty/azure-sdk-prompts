import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, waitFor, fireEvent, within } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { RunsPage } from "../app/components/runs-page";

vi.mock("../app/data/api", () => ({
  fetchRuns: vi.fn(),
}));

import { fetchRuns } from "../app/data/api";

function evalRow(opts: { config: string; language: string; success: boolean }) {
  return {
    prompt_id: "p",
    config_name: opts.config,
    success: opts.success,
    review: { overall_score: 0, max_score: 0, summary: "" },
    duration_seconds: 0,
    prompt_metadata: {
      service: "identity",
      plane: "data-plane",
      language: opts.language,
      category: "auth",
      difficulty: "easy",
    },
  };
}

const mockRuns = [
  {
    run_id: "run-abc123",
    timestamp: "2026-03-29T14:00:00Z",
    total_evaluations: 10,
    passed: 8,
    failed: 1,
    errors: 1,
    duration_seconds: 120.5,
    results: [evalRow({ config: "baseline/claude-opus-4.6", language: "python", success: true })],
  },
  {
    run_id: "run-def456",
    timestamp: "2026-03-28T10:00:00Z",
    total_evaluations: 5,
    passed: 5,
    failed: 0,
    errors: 0,
    duration_seconds: 55.3,
    results: [evalRow({ config: "azure-mcp/claude-opus-4.6", language: "go", success: true })],
  },
];

describe("RunsPage", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("shows loading state initially", () => {
    vi.mocked(fetchRuns).mockReturnValue(new Promise(() => {}));

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    expect(document.querySelector(".animate-spin")).toBeInTheDocument();
  });

  it("renders run cards after data loads", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      // Timestamps are now displayed as run names
      expect(screen.getByText(/Mar 29, 2026/)).toBeInTheDocument();
      expect(screen.getByText(/Mar 28, 2026/)).toBeInTheDocument();
      // Eval counts shown
      expect(screen.getByText(/10 evaluations/)).toBeInTheDocument();
      expect(screen.getByText(/5 evaluations/)).toBeInTheDocument();
    });
  });

  it("shows error state on fetch failure", async () => {
    vi.mocked(fetchRuns).mockRejectedValue(new Error("Network failure"));

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText(/Network failure/)).toBeInTheDocument();
    });
  });

  it("renders N/A for missing or invalid timestamp", async () => {
    const runsWithBadDate = [
      {
        run_id: "run-baddate",
        timestamp: "not-a-date",
        total_evaluations: 3,
        passed: 1,
        failed: 2,
        errors: 0,
        duration_seconds: 30,
        results: [],
      },
    ];
    vi.mocked(fetchRuns).mockResolvedValue(runsWithBadDate as any);

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText(/3 evaluations/)).toBeInTheDocument();
    });
    // Invalid timestamp renders "Unknown" instead of "Invalid Date"
    expect(screen.getByText("Unknown")).toBeInTheDocument();
  });

  it("handles missing duration and pass counts gracefully", async () => {
    const runsWithMissing = [
      {
        run_id: "run-missing",
        timestamp: "",
        total_evaluations: 0,
        errors: 0,
        duration_seconds: undefined,
        results: [],
      },
    ];
    vi.mocked(fetchRuns).mockResolvedValue(runsWithMissing as any);

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText(/0 evaluations/)).toBeInTheDocument();
    });
    // Missing timestamp renders "Unknown"
    const unknownElements = screen.getAllByText("Unknown");
    expect(unknownElements.length).toBeGreaterThanOrEqual(1);
    // Pass rate should show 0.0% (not NaN%)
    expect(screen.getByText("0.0%")).toBeInTheDocument();
  });

  it("renders filter bar populated from run data", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByLabelText("Filter by Config")).toBeInTheDocument();
    });
    expect(screen.getByLabelText("Filter by Language")).toBeInTheDocument();
    expect(screen.getByLabelText("Filter by Status")).toBeInTheDocument();
    // Initially shows total run count
    expect(screen.getByText("2 runs")).toBeInTheDocument();
  });

  it("filters runs when a language is selected and shows reset", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => expect(screen.getByText("2 runs")).toBeInTheDocument());

    // Open the language dropdown and select "go"
    fireEvent.click(screen.getByLabelText("Filter by Language"));
    const listbox = await screen.findByRole("listbox", { name: "Language" });
    fireEvent.click(within(listbox).getByRole("option", { name: "go" }));

    await waitFor(() => {
      expect(screen.getByText("1 of 2")).toBeInTheDocument();
    });
    // run-def456 is the go run; abc123 (python) should be hidden
    expect(screen.getByText(/Mar 28, 2026/)).toBeInTheDocument();
    expect(screen.queryByText(/Mar 29, 2026/)).not.toBeInTheDocument();

    // Reset button should clear filters
    fireEvent.click(screen.getByLabelText("Reset filters"));
    await waitFor(() => expect(screen.getByText("2 runs")).toBeInTheDocument());
  });

  it("shows no-matches empty state when filters exclude every run", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter initialEntries={["/runs?lang=rust"]}>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText(/No runs match the current filters/)).toBeInTheDocument();
    });
  });

  it("hydrates filters from URL query params", async () => {
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);

    render(
      <MemoryRouter initialEntries={["/runs?lang=python&status=failing"]}>
        <RunsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      // python + failing matches run-abc123 (has 1 fail + 1 error → status=errors? No, errors take precedence)
      // run-abc123 has errors=1, so its status is "errors", not "failing".
      // Therefore the filter should match nothing → empty state.
      expect(screen.getByText(/No runs match the current filters/)).toBeInTheDocument();
    });
  });
});
