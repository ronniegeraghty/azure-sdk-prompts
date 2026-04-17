import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { RunsPage } from "../app/components/runs-page";

vi.mock("../app/data/api", () => ({
  fetchRuns: vi.fn(),
}));

import { fetchRuns } from "../app/data/api";

const mockRuns = [
  {
    run_id: "run-abc123",
    timestamp: "2026-03-29T14:00:00Z",
    total_evaluations: 10,
    passed: 8,
    failed: 1,
    errors: 1,
    duration_seconds: 120.5,
    results: [],
  },
  {
    run_id: "run-def456",
    timestamp: "2026-03-28T10:00:00Z",
    total_evaluations: 5,
    passed: 5,
    failed: 0,
    errors: 0,
    duration_seconds: 55.3,
    results: [],
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
});
