import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { DashboardPage } from "../app/components/dashboard-page";

vi.mock("../app/api", () => ({
  getEvaluations: vi.fn(),
  getPrompts: vi.fn(),
}));

import { getEvaluations, getPrompts } from "../app/api";

describe("DashboardPage", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders the page heading", () => {
    vi.mocked(getEvaluations).mockResolvedValue([]);
    vi.mocked(getPrompts).mockResolvedValue([]);

    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("Evaluation Dashboard")).toBeInTheDocument();
  });

  it("renders stat cards", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("Total Evaluations")).toBeInTheDocument();
    expect(screen.getByText("Overall Pass Rate")).toBeInTheDocument();
    expect(screen.getByText("Avg Duration")).toBeInTheDocument();
    expect(screen.getByText("Models Tested")).toBeInTheDocument();
  });

  it("renders stat values", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("1,247")).toBeInTheDocument();
    expect(screen.getByText("78.3%")).toBeInTheDocument();
    expect(screen.getByText("9.8s")).toBeInTheDocument();
    expect(screen.getByText("6")).toBeInTheDocument();
  });

  it("renders the recent evaluations table", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("Recent Evaluations")).toBeInTheDocument();
    expect(screen.getByText("EVL-0042")).toBeInTheDocument();
    expect(screen.getByText("EVL-0041")).toBeInTheDocument();
  });

  it("renders chart section titles", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("Pass Rate")).toBeInTheDocument();
    expect(screen.getByText("Duration Trends (seconds)")).toBeInTheDocument();
    expect(screen.getByText("Model Comparison by Criteria")).toBeInTheDocument();
  });

  it("renders the AI insights section", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("AI-Generated Insights")).toBeInTheDocument();
  });

  it("renders toggle buttons for pass rate chart", () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    expect(screen.getByText("service")).toBeInTheDocument();
    expect(screen.getByText("language")).toBeInTheDocument();
  });
});

  // R154: Dashboard needs real data
  describe("real data integration", () => {
    it("displays total evaluations from API data", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "eval-1", score: 85 },
        { id: "eval-2", score: 90 },
        { id: "eval-3", score: 78 },
        { id: "eval-4", score: 92 },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Total Evaluations")).toBeInTheDocument();
      expect(screen.getByText("4")).toBeInTheDocument();
    });

    it("calculates overall pass rate from real evaluation scores", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "eval-1", score: 90 },
        { id: "eval-2", score: 85 },
        { id: "eval-3", score: 70 },
        { id: "eval-4", score: 95 },
        { id: "eval-5", score: 60 },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Overall Pass Rate")).toBeInTheDocument();
      // Pass threshold is typically 80%, so 3 out of 5 = 60%
      expect(screen.getByText(/60.*%/)).toBeInTheDocument();
    });

    it("calculates average duration from real evaluation data", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "eval-1", duration: 5.2 },
        { id: "eval-2", duration: 8.5 },
        { id: "eval-3", duration: 6.3 },
        { id: "eval-4", duration: 12.0 },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Avg Duration")).toBeInTheDocument();
      // Average of [5.2, 8.5, 6.3, 12.0] = 8.0s
      expect(screen.getByText(/8\.0s/)).toBeInTheDocument();
    });

    it("counts unique models tested from evaluation data", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "eval-1", model: "claude-opus-4.6" },
        { id: "eval-2", model: "claude-sonnet-4.5" },
        { id: "eval-3", model: "gpt-5.3-codex" },
        { id: "eval-4", model: "claude-opus-4.6" },
        { id: "eval-5", model: "claude-haiku-4.5" },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Models Tested")).toBeInTheDocument();
      // 4 unique models
      expect(screen.getByText("4")).toBeInTheDocument();
    });

    it("does NOT display hardcoded mock values", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "eval-1", score: 85 },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      await waitFor(() => {
        expect(screen.getByText("Total Evaluations")).toBeInTheDocument();
      });

      // Should NOT show the old mock values
      expect(screen.queryByText("1,247")).not.toBeInTheDocument();
      expect(screen.queryByText("78.3%")).not.toBeInTheDocument();
      expect(screen.queryByText("9.8s")).not.toBeInTheDocument();
    });

    it("populates recent evaluations table from API data", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([
        { id: "EVL-0100", promptId: "identity-dp-python", score: 92, timestamp: "2025-04-18T10:00:00Z" },
        { id: "EVL-0099", promptId: "key-vault-dp-dotnet", score: 88, timestamp: "2025-04-17T15:00:00Z" },
      ]);
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "identity-dp-python", name: "Identity Default Credential" },
        { id: "key-vault-dp-dotnet", name: "Key Vault CRUD" },
      ]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Recent Evaluations")).toBeInTheDocument();
      expect(screen.getByText("EVL-0100")).toBeInTheDocument();
      expect(screen.getByText("EVL-0099")).toBeInTheDocument();
      expect(screen.getByText("Identity Default Credential")).toBeInTheDocument();
      expect(screen.getByText("Key Vault CRUD")).toBeInTheDocument();
    });
  });

  describe("empty state handling", () => {
    it("shows zero metrics when no evaluations exist", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText("Total Evaluations")).toBeInTheDocument();
      expect(screen.getByText("0")).toBeInTheDocument();
      expect(screen.getByText(/0.*%|N\/A/)).toBeInTheDocument();
    });

    it("displays empty state message in recent evaluations when no data", async () => {
      vi.mocked(getEvaluations).mockResolvedValue([]);
      vi.mocked(getPrompts).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <DashboardPage />
        </MemoryRouter>
      );

      expect(await screen.findByText(/no evaluations|no data available/i)).toBeInTheDocument();
    });
  });
});
