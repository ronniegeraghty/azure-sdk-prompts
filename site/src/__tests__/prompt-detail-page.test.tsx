import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router";
import { PromptDetailPage } from "../app/components/prompt-detail-page";
import type { PromptInfo } from "../app/data/api";
import type { RunSummary } from "../app/data/types";

// Mock the API module (fixed: was ../app/api, now ../app/data/api)
vi.mock("../app/data/api", () => ({
  fetchPrompt: vi.fn(),
  fetchRuns: vi.fn(),
}));

import { fetchPrompt, fetchRuns } from "../app/data/api";

const mockPromptInfo = (id: string): PromptInfo => ({
  id,
  service: "identity",
  plane: "data-plane",
  language: "python",
  category: "auth",
  difficulty: "beginner",
  description: "Test",
  sdk_package: "azure-identity",
  doc_url: "",
  tags: [],
  created: "2025-01-01T00:00:00Z",
  author: "test",
  prompt_text: "Test prompt",
  evaluation_criteria: "Test criteria",
  file_path: "/test.md",
});

const mockRun = (promptId: string): RunSummary => ({
  run_id: "run-1",
  timestamp: "2025-04-15T10:00:00Z",
  total_evaluations: 1,
  passed: 1,
  failed: 0,
  errors: 0,
  duration_seconds: 10,
  results: [{
    prompt_id: promptId,
    config_name: "baseline/claude-opus-4.6",
    success: true,
    review: {
      overall_score: 85,
      max_score: 100,
      summary: "Test"
    },
    duration_seconds: 10,
    prompt_metadata: {
      service: "identity",
      plane: "data-plane",
      language: "python",
      category: "auth",
      difficulty: "beginner",
      tags: []
    }
  }]
});

describe("PromptDetailPage - R151", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("R151: renders page with mocked API data", async () => {
    vi.mocked(fetchPrompt).mockResolvedValue(mockPromptInfo("test"));
    vi.mocked(fetchRuns).mockResolvedValue([mockRun("test")]);

    render(
      <MemoryRouter initialEntries={["/prompts/test"]}>
        <Routes>
          <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
        </Routes>
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText("test")).toBeInTheDocument();
    });
  });
});
