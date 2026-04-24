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
});
