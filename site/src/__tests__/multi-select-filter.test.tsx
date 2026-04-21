import { describe, it, expect, vi } from "vitest";
import { useState } from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { MultiSelectFilter } from "../app/components/ui/multi-select-filter";

const OPTIONS = [
  { value: "a", label: "Alpha" },
  { value: "b", label: "Beta" },
  { value: "c", label: "Gamma" },
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

  describe("toggle / onChange", () => {
    it("calls onChange with the value added when an unselected option is clicked", async () => {
      const user = userEvent.setup();
      const onChange = vi.fn();
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["a"]}
          onChange={onChange}
        />,
      );

      openDropdown();
      await user.click(screen.getByRole("option", { name: /Beta/ }));

      expect(onChange).toHaveBeenCalledTimes(1);
      expect(onChange).toHaveBeenCalledWith(["a", "b"]);
    });

    it("calls onChange with the value removed when a selected option is clicked", async () => {
      const user = userEvent.setup();
      const onChange = vi.fn();
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["a", "b"]}
          onChange={onChange}
        />,
      );

      openDropdown();
      await user.click(screen.getByRole("option", { name: /Alpha/ }));

      expect(onChange).toHaveBeenCalledTimes(1);
      expect(onChange).toHaveBeenCalledWith(["b"]);
    });

    it("preserves multi-select across multiple clicks (controlled wrapper)", async () => {
      const user = userEvent.setup();
      const onChangeSpy = vi.fn();

      function Wrapper() {
        const [sel, setSel] = useState<string[]>([]);
        return (
          <MultiSelectFilter
            label="Config"
            options={OPTIONS}
            selected={sel}
            onChange={(next) => {
              onChangeSpy(next);
              setSel(next);
            }}
          />
        );
      }

      render(<Wrapper />);
      openDropdown();

      await user.click(screen.getByRole("option", { name: /Alpha/ }));
      await user.click(screen.getByRole("option", { name: /Beta/ }));
      await user.click(screen.getByRole("option", { name: /Gamma/ }));

      expect(onChangeSpy).toHaveBeenNthCalledWith(1, ["a"]);
      expect(onChangeSpy).toHaveBeenNthCalledWith(2, ["a", "b"]);
      expect(onChangeSpy).toHaveBeenNthCalledWith(3, ["a", "b", "c"]);

      // After three clicks, all three options should report aria-selected=true.
      for (const name of [/Alpha/, /Beta/, /Gamma/]) {
        expect(screen.getByRole("option", { name })).toHaveAttribute(
          "aria-selected",
          "true",
        );
      }

      // Click Beta again to deselect; the surviving selection is preserved.
      await user.click(screen.getByRole("option", { name: /Beta/ }));
      expect(onChangeSpy).toHaveBeenNthCalledWith(4, ["a", "c"]);
    });
  });

  describe("summary text", () => {
    it("shows the empty label when nothing is selected", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={[]}
          onChange={vi.fn()}
          emptyLabel="Any"
        />,
      );

      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveTextContent(/Config:\s*Any/);
    });

    it("respects a custom empty label", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={[]}
          onChange={vi.fn()}
          emptyLabel="All configs"
        />,
      );

      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveTextContent(/Config:\s*All configs/);
    });

    it("shows the single selected value when exactly one is selected", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["a"]}
          onChange={vi.fn()}
        />,
      );

      // Note: the component renders selected[0] (the value), not the label.
      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveTextContent(/Config:\s*a/);
    });

    it("shows 'N selected' when more than one is selected", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["a", "b"]}
          onChange={vi.fn()}
        />,
      );

      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveTextContent(/Config:\s*2 selected/);
    });

    it("scales the 'N selected' summary as more values are added (overflow branch)", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["a", "b", "c"]}
          onChange={vi.fn()}
        />,
      );

      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveTextContent(/Config:\s*3 selected/);
    });
  });

  describe("ARIA", () => {
    it("toggles aria-expanded on the trigger when opened and closed", async () => {
      const user = userEvent.setup();
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={[]}
          onChange={vi.fn()}
        />,
      );

      const trigger = screen.getByRole("button", { name: "Filter by Config" });
      expect(trigger).toHaveAttribute("aria-expanded", "false");

      await user.click(trigger);
      expect(trigger).toHaveAttribute("aria-expanded", "true");

      await user.click(trigger);
      expect(trigger).toHaveAttribute("aria-expanded", "false");
    });

    it("sets aria-selected on each option to match the selected state", () => {
      render(
        <MultiSelectFilter
          label="Config"
          options={OPTIONS}
          selected={["b"]}
          onChange={vi.fn()}
        />,
      );

      openDropdown();

      expect(
        screen.getByRole("option", { name: /Alpha/ }),
      ).toHaveAttribute("aria-selected", "false");
      expect(
        screen.getByRole("option", { name: /Beta/ }),
      ).toHaveAttribute("aria-selected", "true");
      expect(
        screen.getByRole("option", { name: /Gamma/ }),
      ).toHaveAttribute("aria-selected", "false");
    });
  });

  it("keeps the dropdown open when clicking inside the listbox", async () => {
    const user = userEvent.setup();
    render(
      <MultiSelectFilter
        label="Config"
        options={OPTIONS}
        selected={[]}
        onChange={vi.fn()}
      />,
    );

    openDropdown();
    const listbox = screen.getByRole("listbox", { name: "Config" });
    expect(listbox).toBeInTheDocument();

    // Click an option inside the listbox — the dropdown must remain open
    // (counterpart to the outside-click test above). The mousedown handler
    // should treat clicks inside `ref` as not-outside.
    await user.click(screen.getByRole("option", { name: /Alpha/ }));

    expect(
      screen.getByRole("listbox", { name: "Config" }),
    ).toBeInTheDocument();
  });
});
