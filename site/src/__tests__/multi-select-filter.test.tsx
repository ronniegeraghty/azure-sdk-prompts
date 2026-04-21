import { describe, it, expect, vi } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { MultiSelectFilter } from "../app/components/ui/multi-select-filter";

const OPTIONS = [
  { value: "a", label: "Alpha" },
  { value: "b", label: "Beta" },
];

function openDropdown(label = "Config") {
  fireEvent.click(screen.getByRole("button", { name: `Filter by ${label}` }));
}

describe("MultiSelectFilter", () => {
  it("closes the dropdown when clicking outside", () => {
    render(
      <div>
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={[]}
          onChange={vi.fn()}
        />
        <button type="button">outside</button>
      </div>,
    );

    openDropdown();
    expect(screen.getByRole("listbox", { name: "Config" })).toBeInTheDocument();

    // The component listens to "mousedown" on document, so dispatch that.
    fireEvent.mouseDown(screen.getByText("outside"));

    expect(screen.queryByRole("listbox", { name: "Config" })).not.toBeInTheDocument();
  });

  it("closes the dropdown when pressing Escape", () => {
    render(
      <MultiSelectFilter
        label="Config"
        options={OPTIONS}
        selected={[]}
        onChange={vi.fn()}
      />,
    );

    openDropdown();
    expect(screen.getByRole("listbox", { name: "Config" })).toBeInTheDocument();

    fireEvent.keyDown(document, { key: "Escape" });

    expect(screen.queryByRole("listbox", { name: "Config" })).not.toBeInTheDocument();
  });

  it("renders a 'No options' message when options list is empty", () => {
    render(
      <MultiSelectFilter
        label="Config"
        options={[]}
        selected={[]}
        onChange={vi.fn()}
      />,
    );

    openDropdown();

    const listbox = screen.getByRole("listbox", { name: "Config" });
    expect(listbox).toHaveTextContent("No options");
    // Confirm no option rows render when the list is empty.
    expect(screen.queryAllByRole("option")).toHaveLength(0);
  });
});
