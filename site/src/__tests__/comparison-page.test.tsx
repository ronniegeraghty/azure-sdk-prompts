import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { render, screen, waitFor, within, fireEvent } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { ComparisonPage } from "../app/components/comparison-page";
import { STORAGE_KEY } from "../app/lib/comparison-groups";

vi.mock("../app/data/api", () => ({
  fetchRuns: vi.fn(),
}));

import { fetchRuns } from "../app/data/api";

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
      {
        prompt_id: "identity-dp-python-auth",
        config_name: "baseline/claude-opus-4.6",
        success: true,
        review: { overall_score: 9, max_score: 10, summary: "" },
        duration_seconds: 10,
        prompt_metadata: { service: "identity", plane: "data-plane", language: "python", category: "auth", difficulty: "basic" },
      },
      {
        prompt_id: "identity-dp-python-auth",
        config_name: "azure-mcp/claude-opus-4.6",
        success: true,
        review: { overall_score: 10, max_score: 10, summary: "" },
        duration_seconds: 12,
        prompt_metadata: { service: "identity", plane: "data-plane", language: "python", category: "auth", difficulty: "basic" },
      },
      {
        prompt_id: "storage-dp-dotnet-crud",
        config_name: "baseline/claude-opus-4.6",
        success: false,
        review: { overall_score: 4, max_score: 10, summary: "" },
        duration_seconds: 20,
        prompt_metadata: { service: "storage", plane: "data-plane", language: "dotnet", category: "crud", difficulty: "intermediate" },
      },
      {
        prompt_id: "storage-dp-dotnet-crud",
        config_name: "azure-mcp/claude-opus-4.6",
        success: true,
        review: { overall_score: 8, max_score: 10, summary: "" },
        duration_seconds: 18,
        prompt_metadata: { service: "storage", plane: "data-plane", language: "dotnet", category: "crud", difficulty: "intermediate" },
      },
    ],
  },
];

function installLocalStorageShim(): void {
  const store = new Map<string, string>();
  const shim: Storage = {
    getItem: (k) => store.get(k) ?? null,
    setItem: (k, v) => void store.set(k, String(v)),
    removeItem: (k) => void store.delete(k),
    clear: () => store.clear(),
    key: (i) => Array.from(store.keys())[i] ?? null,
    get length() {
      return store.size;
    },
  } as Storage;
  Object.defineProperty(window, "localStorage", { value: shim, configurable: true, writable: true });
}

beforeEach(() => {
  vi.clearAllMocks();
  installLocalStorageShim();
  vi.mocked(fetchRuns).mockResolvedValue(mockRuns as any);
});

afterEach(() => {
  // shim is reinstalled per-test; nothing else to do.
});

function renderPage() {
  return render(
    <MemoryRouter>
      <ComparisonPage />
    </MemoryRouter>
  );
}

describe("ComparisonPage", () => {
  it("renders the page heading after runs load", async () => {
    renderPage();
    await waitFor(() => {
      expect(screen.getByRole("heading", { name: "Compare" })).toBeInTheDocument();
    });
  });

  it("shows loading spinner while fetching runs", () => {
    vi.mocked(fetchRuns).mockReturnValue(new Promise(() => {}));
    renderPage();
    expect(document.querySelector(".animate-spin")).toBeInTheDocument();
  });

  it("shows empty state with create-first-group call to action when no groups", async () => {
    renderPage();
    await waitFor(() => {
      expect(screen.getByText(/no groups yet/i)).toBeInTheDocument();
      expect(screen.getByRole("button", { name: /create first group/i })).toBeInTheDocument();
    });
  });

  it("creates a group when 'Add group' is clicked and renders summary card", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));
    await waitFor(() => {
      expect(screen.getByText("Group 1")).toBeInTheDocument();
      const cards = screen.getAllByTestId("summary-card");
      expect(cards).toHaveLength(1);
    });
  });

  it("computes correct metrics for a group filtered by config", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));

    await waitFor(() => screen.getByText("baseline/claude-opus-4.6"));
    fireEvent.click(screen.getByText("baseline/claude-opus-4.6"));

    await waitFor(() => {
      const card = screen.getByTestId("summary-card");
      expect(within(card).getByText("50.0%")).toBeInTheDocument();
      expect(within(card).getByText("65.0%")).toBeInTheDocument();
      expect(within(card).getByText("2")).toBeInTheDocument();
    });
  });

  it("shows single-group hint when only one group is defined", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));
    await waitFor(() => {
      expect(screen.getByText(/add another to see side-by-side/i)).toBeInTheDocument();
    });
  });

  it("warns when a group matches no evals", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));

    await waitFor(() => screen.getByText("python"));
    fireEvent.click(screen.getByText("python"));
    fireEvent.click(screen.getByText("storage"));

    await waitFor(() => {
      expect(screen.getByText(/no evals match these filters/i)).toBeInTheDocument();
    });
  });

  it("toggles charts via the visualizations panel", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));

    await waitFor(() => expect(document.querySelector('[data-testid="chart-pass_rate"]')).toBeTruthy());

    fireEvent.click(screen.getByRole("checkbox", { name: /^pass rate$/i }));
    await waitFor(() => {
      expect(document.querySelector('[data-testid="chart-pass_rate"]')).toBeNull();
    });
  });

  it("persists groups to localStorage and rehydrates on reload", async () => {
    const { unmount } = renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));
    await waitFor(() => screen.getByText("Group 1"));

    await waitFor(() => {
      const raw = window.localStorage.getItem(STORAGE_KEY);
      expect(raw).toBeTruthy();
      expect(raw).toContain("Group 1");
    });

    unmount();

    renderPage();
    await waitFor(() => {
      expect(screen.getByText("Group 1")).toBeInTheDocument();
    });
  });

  it("removes a group when the Remove button is clicked", async () => {
    renderPage();
    await waitFor(() => screen.getByRole("button", { name: /add group/i }));
    fireEvent.click(screen.getByRole("button", { name: /add group/i }));
    await waitFor(() => screen.getByText("Group 1"));

    fireEvent.click(screen.getByRole("button", { name: /remove group/i }));
    await waitFor(() => {
      expect(screen.queryByText("Group 1")).not.toBeInTheDocument();
      expect(screen.getByText(/no groups yet/i)).toBeInTheDocument();
    });
  });
});
