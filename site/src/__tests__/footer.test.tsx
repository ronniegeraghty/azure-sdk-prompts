import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { Footer } from "../app/components/footer";

describe("Footer", () => {
  it("renders the hyoka brand", () => {
    render(<Footer />);
    expect(screen.getByText("hyoka")).toBeInTheDocument();
  });

  it("renders the tagline", () => {
    render(<Footer />);
    expect(screen.getByText("Evaluate AI code generation quality for Azure SDKs.")).toBeInTheDocument();
  });

  it("renders the copyright with current year", () => {
    render(<Footer />);
    const year = new Date().getFullYear();
    expect(screen.getByText(new RegExp(`© ${year}`))).toBeInTheDocument();
  });

  it("includes MIT License text", () => {
    render(<Footer />);
    expect(screen.getByText(/MIT License/)).toBeInTheDocument();
  });
});
