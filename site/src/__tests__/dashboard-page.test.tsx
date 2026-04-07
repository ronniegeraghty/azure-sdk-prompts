import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { DashboardPage } from "../app/components/dashboard-page";

// DashboardPage uses hardcoded mock data, so it renders without API calls.

describe("DashboardPage", () => {
  it("renders the page heading", () => {
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
