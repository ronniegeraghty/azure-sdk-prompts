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
      expect(screen.getByText(/run-abc123/)).toBeInTheDocument();
      expect(screen.getByText(/run-def456/)).toBeInTheDocument();
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
});
