import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { DashboardPage } from "../app/components/dashboard-page";

// Mock the API module
vi.mock("../app/data/api", () => ({
  fetchRuns: vi.fn(),
}));

import { fetchRuns } from "../app/data/api";

// Mock data that produces the expected metrics:
// - Total: 1247 evals (distributed across runs)
// - Pass rate: 78.3% 
// - Avg duration: 9.8s
// - Models: 6 different configs
const mockRuns = [
  {
    run_id: "EVL-0042-abcdef",
    timestamp: "2026-04-20T18:00:00Z",
    total_evaluations: 400,
    passed: 320,
    duration_seconds: 9.5,
    avg_generation_duration_seconds: 5.0,
    avg_review_duration_seconds: 4.0,
    results: [
      {
        prompt_id: "identity-dp-python-default-credential",
        config_name: "baseline/claude-opus-4.6",
        success: true,
        review: { overall_score: 85 },
        duration_seconds: 8.5,
        generated_files: ["auth.py", "test_auth.py"],
        prompt_metadata: { service: "identity", language: "python" },
      },
      {
        prompt_id: "key-vault-dp-dotnet-crud",
        config_name: "baseline/claude-sonnet-4.5",
        success: true,
        review: { overall_score: 92 },
        duration_seconds: 9.0,
        generated_files: ["KeyVault.cs"],
        prompt_metadata: { service: "key-vault", language: "dotnet" },
      },
      {
        prompt_id: "storage-dp-java-blob",
        config_name: "azure-mcp/claude-opus-4.6",
        success: false,
        review: { overall_score: 45 },
        duration_seconds: 7.5,
        generated_files: ["BlobClient.java"],
        prompt_metadata: { service: "storage", language: "java" },
      },
    ],
  },
  {
    run_id: "EVL-0041-ghijkl",
    timestamp: "2026-04-20T17:00:00Z",
    total_evaluations: 420,
    passed: 330,
    duration_seconds: 10.1,
    avg_generation_duration_seconds: 5.5,
    avg_review_duration_seconds: 4.5,
    results: [
      {
        prompt_id: "cosmos-dp-python-query",
        config_name: "azure-mcp/claude-sonnet-4.5",
        success: true,
        review: { overall_score: 88 },
        duration_seconds: 9.8,
        generated_files: ["cosmos.py"],
        prompt_metadata: { service: "cosmos-db", language: "python" },
      },
      {
        prompt_id: "identity-mp-go-rbac",
        config_name: "baseline/gpt-5.3-codex",
        success: true,
        review: { overall_score: 78 },
        duration_seconds: 8.2,
        generated_files: ["rbac.go"],
        prompt_metadata: { service: "identity", language: "go" },
      },
      {
        prompt_id: "storage-mp-rust-containers",
        config_name: "azure-mcp/gpt-5.3-codex",
        success: false,
        review: { overall_score: 52 },
        duration_seconds: 9.5,
        generated_files: ["containers.rs"],
        prompt_metadata: { service: "storage", language: "rust" },
      },
    ],
  },
  {
    run_id: "EVL-0040-mnopqr",
    timestamp: "2026-04-20T16:00:00Z",
    total_evaluations: 427,
    passed: 327,
    duration_seconds: 9.8,
    avg_generation_duration_seconds: 5.2,
    avg_review_duration_seconds: 4.3,
    results: [],
  },
];

// DashboardPage uses real data from API, so we mock fetchRuns to return test data

describe("DashboardPage", () => {
  beforeEach(() => {
    vi.clearAllMocks();
    // Set up default mock data for all tests
    vi.mocked(fetchRuns).mockResolvedValue(mockRuns);
  });

  it("renders the page heading", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("Evaluation Dashboard")).toBeInTheDocument();
    });
  });

  it("renders stat cards", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("Total Evaluations")).toBeInTheDocument();
      expect(screen.getByText("Overall Pass Rate")).toBeInTheDocument();
      expect(screen.getByText("Avg Duration")).toBeInTheDocument();
      expect(screen.getByText("Models Tested")).toBeInTheDocument();
    });
  });

  it("renders stat values", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("1,247")).toBeInTheDocument();
      expect(screen.getByText("78.3%")).toBeInTheDocument();
      // Use getAllByText since "9.8s" appears in both stats and table
      const durationElements = screen.getAllByText("9.8s");
      expect(durationElements.length).toBeGreaterThan(0);
      expect(screen.getByText("6")).toBeInTheDocument();
    });
  });

  it("renders the recent evaluations table", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("Recent Evaluations")).toBeInTheDocument();
      // Multiple rows can have EVL-0042 and EVL-0041 in their IDs
      expect(screen.getAllByText(/EVL-0042/).length).toBeGreaterThan(0);
      expect(screen.getAllByText(/EVL-0041/).length).toBeGreaterThan(0);
    });
  });

  it("renders chart section titles", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("Pass Rate")).toBeInTheDocument();
      expect(screen.getByText("Duration Trends (seconds)")).toBeInTheDocument();
    });
  });

  it("renders the AI insights section", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("AI-Generated Insights")).toBeInTheDocument();
    });
  });

  it("renders toggle buttons for pass rate chart", async () => {
    render(
      <MemoryRouter>
        <DashboardPage />
      </MemoryRouter>
    );
    await waitFor(() => {
      expect(screen.getByText("service")).toBeInTheDocument();
      expect(screen.getByText("language")).toBeInTheDocument();
    });
  });
});
