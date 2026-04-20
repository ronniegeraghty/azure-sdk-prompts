import { describe, it, expect, beforeEach, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { MemoryRouter } from "react-router";
import { PromptsPage } from "../app/components/prompts-page";

// Mock the API module
vi.mock("../app/api", () => ({
  getPrompts: vi.fn(),
  getEvaluations: vi.fn(),
}));

import { getPrompts, getEvaluations } from "../app/api";

describe("PromptsPage - R150 Improvements", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe("'only show prompts with evals' filter", () => {
    it("renders the 'with evals' filter toggle", async () => {
      vi.mocked(getPrompts).mockResolvedValue([]);
      vi.mocked(getEvaluations).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      await waitFor(() => {
        expect(screen.getByLabelText(/only show prompts with evaluations/i)).toBeInTheDocument();
      });
    });

    it("filters to show only prompts with evaluations by default", async () => {
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Test Prompt 1", service: "identity", language: "python", plane: "data-plane" },
        { id: "prompt-2", name: "Test Prompt 2", service: "key-vault", language: "dotnet", plane: "data-plane" },
        { id: "prompt-3", name: "Test Prompt 3", service: "storage", language: "java", plane: "data-plane" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", configName: "baseline", score: 85 },
        { promptId: "prompt-2", configName: "baseline", score: 92 },
        // prompt-3 has no evaluations
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      await waitFor(() => {
        expect(screen.getByText("Test Prompt 1")).toBeInTheDocument();
        expect(screen.getByText("Test Prompt 2")).toBeInTheDocument();
        expect(screen.queryByText("Test Prompt 3")).not.toBeInTheDocument();
      });
    });

    it("shows all prompts when 'with evals' filter is toggled off", async () => {
      const user = userEvent.setup();
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Test Prompt 1" },
        { id: "prompt-2", name: "Test Prompt 2" },
        { id: "prompt-3", name: "Test Prompt 3" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", configName: "baseline", score: 85 },
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const toggle = await screen.findByLabelText(/only show prompts with evaluations/i);
      await user.click(toggle);

      await waitFor(() => {
        expect(screen.getByText("Test Prompt 1")).toBeInTheDocument();
        expect(screen.getByText("Test Prompt 2")).toBeInTheDocument();
        expect(screen.getByText("Test Prompt 3")).toBeInTheDocument();
      });
    });
  });

  describe("ordering controls", () => {
    it("renders ordering dropdown with all options", async () => {
      vi.mocked(getPrompts).mockResolvedValue([]);
      vi.mocked(getEvaluations).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const orderingControl = await screen.findByLabelText(/sort by/i);
      expect(orderingControl).toBeInTheDocument();

      // Open dropdown and check options
      await userEvent.setup().click(orderingControl);

      expect(screen.getByText("Most Recently Evaluated")).toBeInTheDocument();
      expect(screen.getByText("Alphabetically")).toBeInTheDocument();
      expect(screen.getByText("Best Performing")).toBeInTheDocument();
      expect(screen.getByText("Worst Performing")).toBeInTheDocument();
    });

    it("orders prompts by most recently evaluated when selected", async () => {
      const user = userEvent.setup();
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Alpha Prompt" },
        { id: "prompt-2", name: "Beta Prompt" },
        { id: "prompt-3", name: "Gamma Prompt" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", timestamp: "2025-04-15T10:00:00Z" },
        { promptId: "prompt-2", timestamp: "2025-04-18T14:00:00Z" },
        { promptId: "prompt-3", timestamp: "2025-04-17T09:00:00Z" },
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const orderingControl = await screen.findByLabelText(/sort by/i);
      await user.click(orderingControl);
      await user.click(screen.getByText("Most Recently Evaluated"));

      const promptCards = await screen.findAllByTestId(/prompt-card/i);
      expect(promptCards[0]).toHaveTextContent("Beta Prompt");
      expect(promptCards[1]).toHaveTextContent("Gamma Prompt");
      expect(promptCards[2]).toHaveTextContent("Alpha Prompt");
    });

    it("orders prompts alphabetically when selected", async () => {
      const user = userEvent.setup();
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-2", name: "Zulu Prompt" },
        { id: "prompt-1", name: "Alpha Prompt" },
        { id: "prompt-3", name: "Mike Prompt" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const orderingControl = await screen.findByLabelText(/sort by/i);
      await user.click(orderingControl);
      await user.click(screen.getByText("Alphabetically"));

      const promptCards = await screen.findAllByTestId(/prompt-card/i);
      expect(promptCards[0]).toHaveTextContent("Alpha Prompt");
      expect(promptCards[1]).toHaveTextContent("Mike Prompt");
      expect(promptCards[2]).toHaveTextContent("Zulu Prompt");
    });

    it("orders prompts by best performing when selected", async () => {
      const user = userEvent.setup();
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Low Score Prompt" },
        { id: "prompt-2", name: "High Score Prompt" },
        { id: "prompt-3", name: "Mid Score Prompt" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", score: 65 },
        { promptId: "prompt-2", score: 95 },
        { promptId: "prompt-3", score: 80 },
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const orderingControl = await screen.findByLabelText(/sort by/i);
      await user.click(orderingControl);
      await user.click(screen.getByText("Best Performing"));

      const promptCards = await screen.findAllByTestId(/prompt-card/i);
      expect(promptCards[0]).toHaveTextContent("High Score Prompt");
      expect(promptCards[1]).toHaveTextContent("Mid Score Prompt");
      expect(promptCards[2]).toHaveTextContent("Low Score Prompt");
    });
  });

  describe("sparkline charts", () => {
    it("renders readable sparkline charts for prompts with multiple evaluations", async () => {
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Test Prompt" },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", score: 70, timestamp: "2025-04-10T10:00:00Z" },
        { promptId: "prompt-1", score: 75, timestamp: "2025-04-11T10:00:00Z" },
        { promptId: "prompt-1", score: 80, timestamp: "2025-04-12T10:00:00Z" },
        { promptId: "prompt-1", score: 85, timestamp: "2025-04-13T10:00:00Z" },
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const sparkline = await screen.findByTestId("sparkline-prompt-1");
      expect(sparkline).toBeInTheDocument();
      // Sparkline should have proper viewBox, stroke width, and color
      expect(sparkline).toHaveAttribute("viewBox");
      expect(sparkline.querySelector("path")).toHaveAttribute("stroke-width");
    });
  });

  describe("eval count and tags prominence", () => {
    it("displays eval count prominently for each prompt", async () => {
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Test Prompt", tags: ["auth", "crud"] },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([
        { promptId: "prompt-1", score: 85 },
        { promptId: "prompt-1", score: 90 },
        { promptId: "prompt-1", score: 88 },
      ]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const evalCount = await screen.findByText(/3 evaluations/i);
      expect(evalCount).toBeInTheDocument();
      expect(evalCount).toHaveClass(/prominent|badge|font-medium/);
    });

    it("displays tags as prominent badges", async () => {
      vi.mocked(getPrompts).mockResolvedValue([
        { id: "prompt-1", name: "Test Prompt", tags: ["auth", "pagination", "crud"] },
      ]);
      vi.mocked(getEvaluations).mockResolvedValue([]);

      render(
        <MemoryRouter>
          <PromptsPage />
        </MemoryRouter>
      );

      const authTag = await screen.findByText("auth");
      const paginationTag = await screen.findByText("pagination");
      const crudTag = await screen.findByText("crud");

      expect(authTag).toBeInTheDocument();
      expect(paginationTag).toBeInTheDocument();
      expect(crudTag).toBeInTheDocument();

      // Tags should be styled as badges
      expect(authTag).toHaveClass(/badge|pill|tag/);
    });
  });
});
