import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { MemoryRouter, Route, Routes } from "react-router";
import { PromptDetailPage } from "../app/components/prompt-detail-page";

vi.mock("../app/api", () => ({
  getPromptById: vi.fn(),
  getEvaluationsForPrompt: vi.fn(),
}));

import { getPromptById, getEvaluationsForPrompt } from "../app/api";

describe("PromptDetailPage - R151 Improvements", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe("prompt content collapsible section", () => {
    it("renders prompt content in a collapsible section", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "identity-dp-python-default-credential",
        name: "Identity Default Credential",
        content: "Generate a Python script that uses DefaultAzureCredential to authenticate...",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([]);

      render(
        <MemoryRouter initialEntries={["/prompts/identity-dp-python-default-credential"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      const collapsibleTrigger = await screen.findByRole("button", { name: /prompt content/i });
      expect(collapsibleTrigger).toBeInTheDocument();
    });

    it("expands to show prompt content when clicked", async () => {
      const user = userEvent.setup();
      const promptContent = "Generate a Python script that uses DefaultAzureCredential to authenticate with Azure Key Vault and retrieve a secret.";

      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
        content: promptContent,
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      const collapsibleTrigger = await screen.findByRole("button", { name: /prompt content/i });
      await user.click(collapsibleTrigger);

      expect(await screen.findByText(promptContent)).toBeInTheDocument();
    });
  });

  describe("labeled badges instead of colored dots", () => {
    it("displays labeled badges for prompt metadata", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
        service: "identity",
        language: "python",
        plane: "data-plane",
        difficulty: "beginner",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      expect(await screen.findByText(/identity/i)).toBeInTheDocument();
      expect(screen.getByText(/python/i)).toBeInTheDocument();
      expect(screen.getByText(/data-plane/i)).toBeInTheDocument();
      expect(screen.getByText(/beginner/i)).toBeInTheDocument();

      // Should be badges, not just colored dots
      const serviceBadge = screen.getByText(/identity/i);
      expect(serviceBadge).toHaveClass(/badge|pill|label/);
    });

    it("does not use colored dots for metadata display", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
        service: "key-vault",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([]);

      const { container } = render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      await waitFor(() => {
        expect(screen.getByText(/key-vault/i)).toBeInTheDocument();
      });

      // Should not have dot/circle elements for metadata
      const coloredDots = container.querySelectorAll('[class*="dot"], [class*="circle"]');
      expect(coloredDots.length).toBe(0);
    });
  });

  describe("score trend chart", () => {
    it("uses days on x-axis instead of evaluation count", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { timestamp: "2025-04-15T10:00:00Z", score: 70 },
        { timestamp: "2025-04-16T10:00:00Z", score: 75 },
        { timestamp: "2025-04-17T10:00:00Z", score: 80 },
        { timestamp: "2025-04-18T10:00:00Z", score: 85 },
      ]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      // X-axis should show dates, not eval numbers
      await waitFor(() => {
        expect(screen.getByText(/Apr 15|04\/15/)).toBeInTheDocument();
      });
    });

    it("shows only average score line", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { timestamp: "2025-04-15T10:00:00Z", score: 70 },
        { timestamp: "2025-04-16T10:00:00Z", score: 75 },
        { timestamp: "2025-04-17T10:00:00Z", score: 80 },
      ]);

      const { container } = render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      await waitFor(() => {
        const lines = container.querySelectorAll('[class*="recharts-line"]');
        // Should have exactly 1 line (average score), not multiple
        expect(lines.length).toBe(1);
      });
    });
  });

  describe("pass rate by model", () => {
    it("shows ALL models, not just top 3", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { model: "claude-opus-4.6", score: 95 },
        { model: "claude-sonnet-4.5", score: 88 },
        { model: "gpt-5.3-codex", score: 85 },
        { model: "claude-haiku-4.5", score: 80 },
        { model: "gpt-5.2", score: 78 },
        { model: "gpt-5-mini", score: 72 },
      ]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      // All 6 models should be visible
      expect(await screen.findByText(/claude-opus-4\.6/i)).toBeInTheDocument();
      expect(screen.getByText(/claude-sonnet-4\.5/i)).toBeInTheDocument();
      expect(screen.getByText(/gpt-5\.3-codex/i)).toBeInTheDocument();
      expect(screen.getByText(/claude-haiku-4\.5/i)).toBeInTheDocument();
      expect(screen.getByText(/gpt-5\.2/i)).toBeInTheDocument();
      expect(screen.getByText(/gpt-5-mini/i)).toBeInTheDocument();
    });

    it("does not truncate model list to top 3", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { model: "model-1", score: 95 },
        { model: "model-2", score: 88 },
        { model: "model-3", score: 85 },
        { model: "model-4", score: 80 },
        { model: "model-5", score: 78 },
      ]);

      const { container } = render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      await waitFor(() => {
        expect(screen.getByText(/model-4/i)).toBeInTheDocument();
        expect(screen.getByText(/model-5/i)).toBeInTheDocument();
      });

      // Should NOT have "Show more" or "View all" button
      expect(screen.queryByText(/show more|view all/i)).not.toBeInTheDocument();
    });
  });

  describe("pass rate by tool", () => {
    it("shows environment tools, not bash/view/create", async () => {
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { tools: ["bash", "view", "create", "azure-cli", "dotnet", "mcp-azure-resources"] },
      ]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      // Environment tools should be shown
      expect(await screen.findByText(/azure-cli/i)).toBeInTheDocument();
      expect(screen.getByText(/dotnet/i)).toBeInTheDocument();
      expect(screen.getByText(/mcp-azure-resources/i)).toBeInTheDocument();

      // Basic tools should NOT be shown in pass rate section
      const passRateSection = screen.getByTestId("pass-rate-by-tool");
      expect(passRateSection).not.toHaveTextContent(/bash/);
      expect(passRateSection).not.toHaveTextContent(/view/);
      expect(passRateSection).not.toHaveTextContent(/create/);
    });

    it("includes a toggle to show/hide tool usage", async () => {
      const user = userEvent.setup();
      vi.mocked(getPromptById).mockResolvedValue({
        id: "test-prompt",
        name: "Test Prompt",
      });
      vi.mocked(getEvaluationsForPrompt).mockResolvedValue([
        { tools: ["azure-cli", "gh"] },
      ]);

      render(
        <MemoryRouter initialEntries={["/prompts/test-prompt"]}>
          <Routes>
            <Route path="/prompts/:promptId" element={<PromptDetailPage />} />
          </Routes>
        </MemoryRouter>
      );

      const usageToggle = await screen.findByLabelText(/show tool usage/i);
      expect(usageToggle).toBeInTheDocument();

      await user.click(usageToggle);

      // Usage data should appear when toggled
      expect(screen.getByText(/used in \d+ evaluations/i)).toBeInTheDocument();
    });
  });
});
