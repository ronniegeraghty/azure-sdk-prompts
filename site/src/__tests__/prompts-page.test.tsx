import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { PromptsPage } from "../app/components/prompts-page";
import type { PromptInfo } from "../app/data/api";
import type { RunSummary } from "../app/data/types";

// Mock the API module (fixed: was ../app/api, now ../app/data/api)
vi.mock("../app/data/api", () => ({
  fetchPrompts: vi.fn(),
  fetchRuns: vi.fn(),
}));

import { fetchPrompts, fetchRuns } from "../app/data/api";

// Mock data helpers
const mockPrompt = (id: string): PromptInfo => ({
  id,
  service: "identity",
  plane: "data-plane",
  language: "python",
  category: "auth",
  difficulty: "beginner",
  description: `Test ${id}`,
  sdk_package: "azure-identity",
  doc_url: "",
  tags: [],
  created: "2025-01-01T00:00:00Z",
  author: "test",
  prompt_text: "",
  evaluation_criteria: "",
  file_path: "",
});

const mockRun = (promptId: string, score: number): RunSummary => ({
  run_id: `run-${promptId}`,
  timestamp: "2025-04-15T10:00:00Z",
  total_evaluations: 1,
  passed: score >= 70 ? 1 : 0,
  failed: score < 70 ? 1 : 0,
  errors: 0,
  duration_seconds: 10,
  results: [{
    prompt_id: promptId,
    config_name: "baseline/claude-opus-4.6",
    success: score >= 70,
    review: {
      overall_score: score,
      max_score: 100,
      summary: "Test"
    },
    duration_seconds: 10,
    prompt_metadata: {
      service: "identity",
      plane: "data-plane",
      language: "python",
      category: "auth",
      difficulty: "beginner"
    }
  }]
});

describe("PromptsPage - R150", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("R150: renders page with mocked API data", async () => {
    vi.mocked(fetchPrompts).mockResolvedValue([mockPrompt("p1")]);
    vi.mocked(fetchRuns).mockResolvedValue([mockRun("p1", 85)]);

    render(
      <MemoryRouter>
        <PromptsPage />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText("Prompt Explorer")).toBeInTheDocument();
    });
  });
});
