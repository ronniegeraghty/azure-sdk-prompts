import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { HomePage } from "../app/components/home-page";

describe("HomePage", () => {
  it("renders the hero heading", () => {
    render(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
    expect(screen.getByText("Evaluate AI Agent")).toBeInTheDocument();
    expect(screen.getByText("Output Quality")).toBeInTheDocument();
  });

  it("renders the developer evaluation tool badge", () => {
    render(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
    expect(screen.getByText("Developer Evaluation Tool")).toBeInTheDocument();
  });

  it("renders feature cards", () => {
    render(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
    expect(screen.getByText("Side-by-Side Comparison")).toBeInTheDocument();
    expect(screen.getByText("Deep Metrics")).toBeInTheDocument();
    expect(screen.getByText("Multi-Grader Pipeline")).toBeInTheDocument();
    expect(screen.getByText("Polyglot Support")).toBeInTheDocument();
  });

  it("renders the how-it-works steps", () => {
    render(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
    expect(screen.getByText("Define Prompts")).toBeInTheDocument();
    expect(screen.getByText("Run Evaluations")).toBeInTheDocument();
    expect(screen.getByText("Review Results")).toBeInTheDocument();
  });

  it("renders supported services", () => {
    render(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
    expect(screen.getByText("Storage")).toBeInTheDocument();
    expect(screen.getByText("Key Vault")).toBeInTheDocument();
    expect(screen.getByText("Identity")).toBeInTheDocument();
  });
});
