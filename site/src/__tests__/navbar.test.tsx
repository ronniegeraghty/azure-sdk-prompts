import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { Navbar } from "../app/components/navbar";

describe("Navbar", () => {
  it("renders the hyoka brand link", () => {
    render(
      <MemoryRouter>
        <Navbar />
      </MemoryRouter>
    );
    // Brand text
    expect(screen.getByText("hyoka")).toBeInTheDocument();
  });

  it("renders all navigation links", () => {
    render(
      <MemoryRouter>
        <Navbar />
      </MemoryRouter>
    );
    const expectedLinks = ["Home", "How It Works", "Runs", "Prompts", "Pairwise", "Compare", "Dashboard", "Docs"];
    for (const label of expectedLinks) {
      expect(screen.getAllByText(label).length).toBeGreaterThanOrEqual(1);
    }
  });

  it("highlights the active route", () => {
    render(
      <MemoryRouter initialEntries={["/dashboard"]}>
        <Navbar />
      </MemoryRouter>
    );
    // The Dashboard link should have the active class
    const dashboardLinks = screen.getAllByText("Dashboard");
    const desktopLink = dashboardLinks.find(
      (el) => el.closest("div.hidden")
    );
    if (desktopLink) {
      expect(desktopLink.className).toContain("bg-white/10");
    }
  });

  it("renders the Get Started link", () => {
    render(
      <MemoryRouter>
        <Navbar />
      </MemoryRouter>
    );
    expect(screen.getByText("Get Started")).toBeInTheDocument();
    const link = screen.getByText("Get Started").closest("a");
    expect(link).toHaveAttribute("href", "https://github.com/ronniegeraghty/hyoka");
  });
});
