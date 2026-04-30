import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import type { GraderResult } from "../app/data/types";

// Import the grouping helpers via the component module.
// Since GraderResultsGrouped is a module-internal function, we test the
// rendered output through GraderResultRow (the leaf) to verify grouping.
import { GraderResultRow } from "../app/components/GraderResultRow";

function makeGrader(overrides: Partial<GraderResult> = {}): GraderResult {
  return {
    grader_name: "Test Grader",
    grader_type: "output_check",
    score: 1.0,
    weight: 1.0,
    pass: true,
    message: "",
    points: [{ label: "check", pass: true }],
    ...overrides,
  };
}

describe("GraderResult source_file / source_type fields", () => {
  it("GraderResult interface accepts source_file and source_type fields", () => {
    const result = makeGrader({
      source_file: "/criteria/python.yaml",
      source_type: "criteria_file",
    });
    // Just exercising that the type accepts these fields.
    expect(result.source_file).toBe("/criteria/python.yaml");
    expect(result.source_type).toBe("criteria_file");
  });

  it("GraderResultRow renders normally when source_file is present", () => {
    const result = makeGrader({
      grader_name: "Output Files Check",
      source_file: "/criteria/python.yaml",
      source_type: "criteria_file",
      points: [
        { label: "min_files (1)", pass: true },
        { label: "min_bytes_per_file (1)", pass: true },
      ],
    });
    render(<GraderResultRow result={result} />);
    expect(screen.getByText("Output Files Check")).toBeInTheDocument();
    expect(screen.getByText("2/2 points")).toBeInTheDocument();
  });

  it("GraderResultRow renders normally when source_file is absent (flat fallback)", () => {
    const result = makeGrader({
      grader_name: "Legacy Grader",
      // No source_file
    });
    render(<GraderResultRow result={result} />);
    expect(screen.getByText("Legacy Grader")).toBeInTheDocument();
  });

  it("source_file is optional — missing fields do not crash", () => {
    // Simulate a pre-Neo report that has no source_file/source_type.
    const result: GraderResult = {
      grader_name: "Old Grader",
      grader_type: "prompt",
      score: 0.5,
      weight: 1.0,
      pass: false,
      message: "some message",
      points: [{ label: "criterion A", pass: false }],
    };
    // Should not throw.
    expect(() => render(<GraderResultRow result={result} />)).not.toThrow();
    expect(screen.getByText("Old Grader")).toBeInTheDocument();
  });
});
