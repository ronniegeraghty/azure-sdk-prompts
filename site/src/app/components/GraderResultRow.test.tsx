import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { GraderResultRow } from "./GraderResultRow";
import type { GraderResult } from "../data/types";

describe("GraderResultRow", () => {
  it("renders a passing grader result", () => {
    const result: GraderResult = {
      grader_name: "Build Checker",
      grader_type: "program",
      pass: true,
      score: 1.0,
      summary: "Build successful",
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("Build Checker")).toBeInTheDocument();
    expect(screen.getByText("PASS")).toBeInTheDocument();
    expect(screen.getByText("100%")).toBeInTheDocument();
  });

  it("renders a failing grader result", () => {
    const result: GraderResult = {
      grader_name: "File Check",
      grader_type: "file",
      pass: false,
      score: 0.0,
      summary: "Required file missing",
      issues: ["Missing main.py"],
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("File Check")).toBeInTheDocument();
    expect(screen.getByText("FAIL")).toBeInTheDocument();
    expect(screen.getByText("0%")).toBeInTheDocument();
  });

  it("handles missing score gracefully", () => {
    const result: GraderResult = {
      grader_name: "Simple Check",
      grader_type: "file",
      pass: true,
      summary: "All good",
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("Simple Check")).toBeInTheDocument();
    expect(screen.getByText("PASS")).toBeInTheDocument();
    expect(screen.queryByText("%")).not.toBeInTheDocument();
  });

  it("expands to show details when clicked", () => {
    const result: GraderResult = {
      grader_name: "Code Review",
      grader_type: "review",
      pass: true,
      score: 0.85,
      summary: "Code quality is good",
      issues: ["Minor style issues"],
      strengths: ["Good error handling", "Clean architecture"],
    };

    render(<GraderResultRow result={result} />);

    // Summary should not be visible initially
    expect(screen.queryByText("Code quality is good")).not.toBeInTheDocument();

    // Click to expand - use the grader name which is unique
    fireEvent.click(screen.getByText("Code Review"));

    // Now details should be visible
    expect(screen.getByText("Code quality is good")).toBeInTheDocument();
    expect(screen.getByText("Minor style issues")).toBeInTheDocument();
    expect(screen.getByText("Good error handling")).toBeInTheDocument();
    expect(screen.getByText("Clean architecture")).toBeInTheDocument();
  });

  it("displays grader type label properly", () => {
    const result: GraderResult = {
      grader_name: "Prompt Reviewer",
      grader_type: "prompt_review",
      pass: true,
      score: 0.9,
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("Prompt Reviewer")).toBeInTheDocument();
    // The grader type label is formatted to "Prompt Review"
    const promptReviewElements = screen.getAllByText(/Prompt Review/);
    expect(promptReviewElements.length).toBeGreaterThan(0);
  });

  it("shows gate indicator for gate graders", () => {
    const result: GraderResult = {
      grader_name: "Critical Check",
      grader_type: "file",
      pass: false,
      gate: true,
      score: 0.0,
      summary: "Gate failure",
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("GATE")).toBeInTheDocument();
  });

  it("renders overall_score/max_score when score is not present", () => {
    const result: GraderResult = {
      grader_name: "Legacy Reviewer",
      grader_type: "review",
      pass: true,
      overall_score: 85,
      max_score: 100,
      summary: "Good work",
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("85/100")).toBeInTheDocument();
  });

  it("handles null/undefined pass status", () => {
    const result: GraderResult = {
      grader_name: "Incomplete",
      grader_type: "behavior",
      pass: null,
      summary: "Not evaluated",
    };

    render(<GraderResultRow result={result} />);
    expect(screen.getByText("N/A")).toBeInTheDocument();
  });

  it("renders file details when present", () => {
    const result: GraderResult = {
      grader_name: "File Checker",
      grader_type: "file",
      pass: true,
      file_details: {
        checked_files: [
          { path: "src/main.py", exists: true },
          { path: "README.md", exists: false },
        ],
      },
    };

    render(<GraderResultRow result={result} defaultExpanded />);
    expect(screen.getByText("src/main.py")).toBeInTheDocument();
    expect(screen.getByText("README.md")).toBeInTheDocument();
  });

  it("renders program details when present", () => {
    const result: GraderResult = {
      grader_name: "Build Test",
      grader_type: "program",
      pass: false,
      program_details: {
        command: "npm run build",
        exit_code: 1,
        stdout: "Building...",
        stderr: "Error: Module not found",
      },
    };

    render(<GraderResultRow result={result} defaultExpanded />);
    expect(screen.getByText("npm run build")).toBeInTheDocument();
    expect(screen.getByText("1")).toBeInTheDocument();
    expect(screen.getByText("Error: Module not found")).toBeInTheDocument();
  });
});
