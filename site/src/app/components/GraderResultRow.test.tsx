import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { GraderResultRow } from "./GraderResultRow";
import type { GraderResult } from "../data/types";

function makeResult(overrides: Partial<GraderResult> = {}): GraderResult {
  return {
    grader_name: "Grader",
    grader_type: "file",
    score: 1.0,
    weight: 1,
    pass: true,
    message: "",
    points: [{ label: "check one", pass: true }],
    ...overrides,
  };
}

describe("GraderResultRow (v4)", () => {
  it("renders the canonical N/M points score (no PASS/100%)", () => {
    render(<GraderResultRow result={makeResult({ points: [
      { label: "a", pass: true },
      { label: "b", pass: true },
      { label: "c", pass: false },
    ], pass: false })} />);
    expect(screen.getByText("2/3 points")).toBeInTheDocument();
    expect(screen.queryByText(/^PASS$/)).not.toBeInTheDocument();
    expect(screen.queryByText(/^100%$/)).not.toBeInTheDocument();
    expect(screen.queryByText(/^FAIL$/)).not.toBeInTheDocument();
  });

  it("falls back to defensive '1/1 points' when Points array is empty and grader passes", () => {
    render(<GraderResultRow result={makeResult({ points: [], pass: true })} />);
    expect(screen.getByText("1/1 points")).toBeInTheDocument();
    expect(screen.queryByText(/^PASS$/)).not.toBeInTheDocument();
    expect(screen.queryByText(/^100%$/)).not.toBeInTheDocument();
  });

  it("falls back to defensive '0/1 points' when Points array is empty and grader fails", () => {
    render(<GraderResultRow result={makeResult({ points: [], pass: false })} />);
    expect(screen.getByText("0/1 points")).toBeInTheDocument();
  });

  it("is collapsed by default — point labels are not visible until expanded", () => {
    render(<GraderResultRow result={makeResult({ points: [
      { label: "passing point label", pass: true },
      { label: "failing point label", pass: false },
    ], pass: false })} />);
    expect(screen.queryByText("passing point label")).not.toBeInTheDocument();
    expect(screen.queryByText("failing point label")).not.toBeInTheDocument();
  });

  it("renders point labels (incl. passing points) once expanded", () => {
    render(<GraderResultRow result={makeResult({ points: [
      { label: "passing point label", pass: true },
      { label: "failing point label", pass: false, message: "boom" },
    ], pass: false })} />);
    fireEvent.click(screen.getByText("Grader"));
    expect(screen.getByText("passing point label")).toBeInTheDocument();
    expect(screen.getByText("failing point label")).toBeInTheDocument();
    expect(screen.getByText("boom")).toBeInTheDocument();
  });

  it("falls back to message or 'Check passed' when point label is empty", () => {
    render(<GraderResultRow result={makeResult({ points: [
      { label: "", pass: true, message: "hello from message" },
      { label: "", pass: true },
    ] })} />);
    fireEvent.click(screen.getByText("Grader"));
    expect(screen.getByText("hello from message")).toBeInTheDocument();
    expect(screen.getByText("Check passed")).toBeInTheDocument();
  });

  it("falls back to legacy 'name' field when label is missing (pre-v4 reports)", () => {
    render(<GraderResultRow result={makeResult({ points: [
      { label: "", name: "legacy name", pass: true } as any,
    ] })} />);
    fireEvent.click(screen.getByText("Grader"));
    expect(screen.getByText("legacy name")).toBeInTheDocument();
  });

  it("respects defaultExpanded=true", () => {
    render(<GraderResultRow result={makeResult({ points: [{ label: "shown immediately", pass: true }] })} defaultExpanded />);
    expect(screen.getByText("shown immediately")).toBeInTheDocument();
  });

  it("shows GATE indicator for gate graders", () => {
    render(<GraderResultRow result={makeResult({ gate: true, pass: false })} />);
    expect(screen.getByText("GATE")).toBeInTheDocument();
  });

  describe("Per-Reviewer Criteria Integration", () => {
    it("passes reviewer votes to ExpandablePoint when extras.review.panel_results has matching criteria", () => {
      const result = makeResult({
        points: [
          { label: "Uses DefaultAzureCredential", pass: true },
          { label: "Handles errors properly", pass: false },
        ],
        extras: {
          review: {
            model: "claude-opus-4.6",
            summary: "Overall review",
            panel_results: [
              {
                model: "opus",
                overall_score: 5,
                max_score: 5,
                summary: "Good",
                criteria: [
                  { name: "Uses DefaultAzureCredential", passed: true, reason: "Correct usage" },
                  { name: "Handles errors properly", passed: false, reason: "Missing try-catch" },
                ],
              },
              {
                model: "sonnet",
                overall_score: 4,
                max_score: 5,
                summary: "OK",
                criteria: [
                  { name: "Uses DefaultAzureCredential", passed: true, reason: "Looks good" },
                  { name: "Handles errors properly", passed: true, reason: "Try-catch present" },
                ],
              },
            ],
          },
        },
      });

      render(<GraderResultRow result={result} />);

      // Expand to see points
      fireEvent.click(screen.getByText("Grader"));

      // Points should be visible
      expect(screen.getByText("Uses DefaultAzureCredential")).toBeInTheDocument();
      expect(screen.getByText("Handles errors properly")).toBeInTheDocument();

      // The second point should auto-expand because of disagreement (false vs true)
      // and show the reviewer votes
      expect(screen.getByText(/opus:/)).toBeInTheDocument();
      expect(screen.getByText(/sonnet:/)).toBeInTheDocument();
    });

    it("matches criteria by exact string: point.label ↔ criterion.name", () => {
      const result = makeResult({
        points: [
          { label: "Exact Match", pass: true },
          { label: "No Match", pass: true },
        ],
        extras: {
          review: {
            model: "claude-opus-4.6",
            summary: "Overall review",
            panel_results: [
              {
                model: "opus",
                overall_score: 5,
                max_score: 5,
                summary: "Good",
                criteria: [
                  { name: "Exact Match", passed: true, reason: "Found" },
                  { name: "Different Name", passed: true, reason: "Also found" },
                ],
              },
            ],
          },
        },
      });

      render(<GraderResultRow result={result} />);
      fireEvent.click(screen.getByText("Grader"));

      // First point should have matching criterion
      const exactMatchPoint = screen.getByText("Exact Match").closest('[role="button"]');
      if (exactMatchPoint) {
        fireEvent.click(exactMatchPoint);
        // Should show the reviewer vote
        expect(screen.getByText(/Found/)).toBeInTheDocument();
      }

      // Second point should not have matching criterion (no role=button since no reviewers)
      const noMatchPoint = screen.getByText("No Match").parentElement;
      expect(noMatchPoint?.querySelector('[role="button"]')).not.toBeInTheDocument();
    });

    it("handles points with no matching criteria (reviewerVotes empty array)", () => {
      const result = makeResult({
        points: [
          { label: "Orphan Point", pass: true },
        ],
        extras: {
          review: {
            model: "claude-opus-4.6",
            summary: "Overall review",
            panel_results: [
              {
                model: "opus",
                overall_score: 5,
                max_score: 5,
                summary: "Good",
                criteria: [
                  { name: "Different Name", passed: true, reason: "Doesn't match" },
                ],
              },
            ],
          },
        },
      });

      render(<GraderResultRow result={result} />);
      fireEvent.click(screen.getByText("Grader"));

      // Point should render but not be expandable (no reviewer votes matched)
      expect(screen.getByText("Orphan Point")).toBeInTheDocument();
      
      // Should not have expand button
      const pointContainer = screen.getByText("Orphan Point").closest('div');
      expect(pointContainer?.querySelector('[role="button"]')).not.toBeInTheDocument();
    });

    it("works when panel_results has no criteria field", () => {
      const result = makeResult({
        points: [
          { label: "Check", pass: true },
        ],
        extras: {
          review: {
            model: "claude-opus-4.6",
            summary: "Overall review",
            panel_results: [
              {
                model: "opus",
                overall_score: 5,
                max_score: 5,
                summary: "Good",
                // No criteria field
              },
            ],
          },
        },
      });

      render(<GraderResultRow result={result} />);
      fireEvent.click(screen.getByText("Grader"));

      // Should render without crashing
      expect(screen.getByText("Check")).toBeInTheDocument();
    });
  });
});
