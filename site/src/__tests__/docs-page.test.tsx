import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, waitFor, within } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";
import { DocsPage } from "../app/components/docs-page";
import * as api from "../app/data/api";

vi.mock("../app/data/api");

const mockDocs = [
  { slug: "getting-started", title: "Getting Started" },
  { slug: "cli-reference", title: "CLI Reference" },
  { slug: "configuration", title: "Configuration" },
  { slug: "grader-config-schema", title: "Grader Config Schema" },
  { slug: "tool-filter-schema", title: "Tool Filter Schema" },
  { slug: "starter-files", title: "Starter Files" },
  { slug: "guardrails", title: "Guardrails" },
  { slug: "prompt-authoring", title: "Prompt Authoring" },
  { slug: "roadmap", title: "Roadmap" },
];

const mockDocContent = {
  slug: "getting-started",
  title: "Getting Started",
  content: `# Getting Started\n\nThis is the getting started guide.\n\n## Installation\n\n\`\`\`bash\ngo install github.com/ronniegeraghty/hyoka/hyoka@latest\n\`\`\`\n\n## Usage\n\nRun your first evaluation.`,
};

describe("DocsPage", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders loading state initially", () => {
    vi.mocked(api.fetchDocs).mockReturnValue(new Promise(() => {}));
    render(<DocsPage />);
    // Check for spinner by its class or text content
    expect(screen.getByText((content, element) => {
      return element?.classList?.contains('animate-spin') === true;
    })).toBeInTheDocument();
  });

  it("renders error state when docs fail to load", async () => {
    vi.mocked(api.fetchDocs).mockRejectedValue(new Error("Network error"));
    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Failed to load documentation")).toBeInTheDocument();
      expect(screen.getByText("Network error")).toBeInTheDocument();
    });
  });

  it("groups docs into logical sections", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "h3" })).toBeInTheDocument();
      expect(screen.getByText("CLI Reference", { selector: "h3" })).toBeInTheDocument();
      expect(screen.getByText("Configuration", { selector: "h3" })).toBeInTheDocument();
      expect(screen.getByText("Concepts", { selector: "h3" })).toBeInTheDocument();
    });
  });

  it("renders all docs in their correct groups", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      // Getting Started group
      const gettingStartedSection = screen.getByText("Getting Started", { selector: "h3" }).closest("div");
      expect(within(gettingStartedSection!).getByText("Getting Started", { selector: "button" })).toBeInTheDocument();

      // CLI Reference group
      const cliSection = screen.getByText("CLI Reference", { selector: "h3" }).closest("div");
      expect(within(cliSection!).getByText("CLI Reference", { selector: "button" })).toBeInTheDocument();

      // Configuration group should have 4 docs
      const configSection = screen.getByText("Configuration", { selector: "h3" }).closest("div");
      expect(within(configSection!).getByText("Configuration", { selector: "button" })).toBeInTheDocument();
      expect(within(configSection!).getByText("Grader Config Schema")).toBeInTheDocument();
      expect(within(configSection!).getByText("Tool Filter Schema")).toBeInTheDocument();
      expect(within(configSection!).getByText("Starter Files")).toBeInTheDocument();

      // Concepts group should have 3 docs
      const conceptsSection = screen.getByText("Concepts", { selector: "h3" }).closest("div");
      expect(within(conceptsSection!).getByText("Guardrails")).toBeInTheDocument();
      expect(within(conceptsSection!).getByText("Prompt Authoring")).toBeInTheDocument();
      expect(within(conceptsSection!).getByText("Roadmap")).toBeInTheDocument();
    });
  });

  it("renders search input", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      const searchInput = screen.getByPlaceholderText("Search docs...");
      expect(searchInput).toBeInTheDocument();
    });
  });

  it("filters docs based on search query", async () => {
    const user = userEvent.setup();
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "button" })).toBeInTheDocument();
    });

    const searchInput = screen.getByPlaceholderText("Search docs...");
    await user.type(searchInput, "schema");

    await waitFor(() => {
      // Should show docs with "schema" in title or slug
      expect(screen.getByText("Grader Config Schema")).toBeInTheDocument();
      expect(screen.getByText("Tool Filter Schema")).toBeInTheDocument();

      // Should NOT show unrelated docs
      expect(screen.queryByText("Getting Started", { selector: "button" })).not.toBeInTheDocument();
      expect(screen.queryByText("CLI Reference", { selector: "button" })).not.toBeInTheDocument();
    });
  });

  it("shows 'No results found' when search has no matches", async () => {
    const user = userEvent.setup();
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "button" })).toBeInTheDocument();
    });

    const searchInput = screen.getByPlaceholderText("Search docs...");
    await user.type(searchInput, "nonexistent-doc-xyz");

    await waitFor(() => {
      expect(screen.getByText("No results found")).toBeInTheDocument();
    });
  });

  it("loads and displays doc content when a doc is selected", async () => {
    const user = userEvent.setup();
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "button" })).toBeInTheDocument();
    });

    // First doc should be auto-selected
    await waitFor(() => {
      expect(api.fetchDoc).toHaveBeenCalledWith("getting-started");
      expect(screen.getByText("This is the getting started guide.")).toBeInTheDocument();
    });
  });

  it("switches to different doc when clicked", async () => {
    const user = userEvent.setup();
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc)
      .mockResolvedValueOnce(mockDocContent)
      .mockResolvedValueOnce({
        slug: "cli-reference",
        title: "CLI Reference",
        content: "# CLI Reference\n\nCommand line interface documentation.",
      });

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "button" })).toBeInTheDocument();
    });

    const cliButton = screen.getByText("CLI Reference", { selector: "button" });
    await user.click(cliButton);

    await waitFor(() => {
      expect(api.fetchDoc).toHaveBeenCalledWith("cli-reference");
      expect(screen.getByText("Command line interface documentation.")).toBeInTheDocument();
    });
  });

  it("highlights the active doc in the sidebar", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      const activeButton = screen.getByText("Getting Started", { selector: "button" });
      expect(activeButton.className).toContain("bg-emerald-500/10");
      expect(activeButton.className).toContain("text-emerald-400");
    });
  });

  it("does not show developer docs (architecture)", async () => {
    // Backend should filter out architecture.md, so it won't be in the list
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      expect(screen.getByText("Getting Started", { selector: "button" })).toBeInTheDocument();
    });

    // Verify architecture.md is not in the sidebar
    expect(screen.queryByText("Architecture")).not.toBeInTheDocument();
  });

  it("renders code blocks with proper styling", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      const codeElement = screen.getByText("go install github.com/ronniegeraghty/hyoka/hyoka@latest");
      expect(codeElement).toBeInTheDocument();
      expect(codeElement.tagName).toBe("CODE");
    });
  });

  it("renders markdown headings correctly", async () => {
    vi.mocked(api.fetchDocs).mockResolvedValue(mockDocs);
    vi.mocked(api.fetchDoc).mockResolvedValue(mockDocContent);

    render(<DocsPage />);

    await waitFor(() => {
      // H2 heading from markdown
      expect(screen.getByText("Installation")).toBeInTheDocument();
      expect(screen.getByText("Usage")).toBeInTheDocument();
    });
  });
});
