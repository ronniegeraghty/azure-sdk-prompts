import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { Layout } from "../app/components/layout";

describe("Layout", () => {
  it("renders the navbar with hyoka brand", () => {
    render(
      <MemoryRouter>
        <Layout />
      </MemoryRouter>
    );
    const elements = screen.getAllByText("hyoka");
    expect(elements.length).toBeGreaterThanOrEqual(1);
  });

  it("renders navigation links", () => {
    render(
      <MemoryRouter>
        <Layout />
      </MemoryRouter>
    );
    expect(screen.getByText("Home")).toBeInTheDocument();
    expect(screen.getByText("Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Compare")).toBeInTheDocument();
    expect(screen.getByText("Runs")).toBeInTheDocument();
    expect(screen.getByText("Prompts")).toBeInTheDocument();
    expect(screen.getByText("Docs")).toBeInTheDocument();
  });

  it("renders the footer with copyright", () => {
    render(
      <MemoryRouter>
        <Layout />
      </MemoryRouter>
    );
    const year = new Date().getFullYear();
    expect(screen.getByText(new RegExp(`© ${year}`))).toBeInTheDocument();
  });
});
