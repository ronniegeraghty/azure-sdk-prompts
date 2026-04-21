import { useEffect, useRef, useState } from "react";
import { ChevronDown, Check } from "lucide-react";

interface Option {
  value: string;
  label: string;
}

interface MultiSelectFilterProps {
  label: string;
  options: Option[];
  selected: string[];
  onChange: (next: string[]) => void;
  emptyLabel?: string;
}

// Compact dropdown with checkbox-style multi-select. Closes on outside click
// and on Escape. Designed to sit in a horizontal filter bar on the runs page.
export function MultiSelectFilter({
  label,
  options,
  selected,
  onChange,
  emptyLabel = "Any",
}: MultiSelectFilterProps) {
  const [open, setOpen] = useState(false);
  const ref = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (!open) return;
    function onDocClick(e: MouseEvent) {
      if (!ref.current) return;
      if (!ref.current.contains(e.target as Node)) setOpen(false);
    }
    function onKey(e: KeyboardEvent) {
      if (e.key === "Escape") setOpen(false);
    }
    document.addEventListener("mousedown", onDocClick);
    document.addEventListener("keydown", onKey);
    return () => {
      document.removeEventListener("mousedown", onDocClick);
      document.removeEventListener("keydown", onKey);
    };
  }, [open]);

  function toggle(value: string) {
    onChange(
      selected.includes(value)
        ? selected.filter((v) => v !== value)
        : [...selected, value],
    );
  }

  const summary =
    selected.length === 0
      ? emptyLabel
      : selected.length === 1
        ? selected[0]
        : `${selected.length} selected`;

  return (
    <div ref={ref} className="relative">
      <button
        type="button"
        onClick={() => setOpen((v) => !v)}
        aria-haspopup="listbox"
        aria-expanded={open}
        aria-label={`Filter by ${label}`}
        className={`flex items-center gap-2 rounded-lg border px-3 py-1.5 transition ${
          selected.length > 0
            ? "border-emerald-500/40 bg-emerald-500/10 text-emerald-300"
            : "border-white/10 bg-white/5 text-white/70 hover:border-white/20"
        }`}
        style={{ fontSize: 12 }}
      >
        <span className="text-white/40" style={{ fontSize: 11 }}>
          {label}:
        </span>
        <span>{summary}</span>
        <ChevronDown className="h-3 w-3" />
      </button>

      {open && (
        <div
          role="listbox"
          aria-label={label}
          className="absolute left-0 z-20 mt-1 max-h-72 w-56 overflow-y-auto rounded-lg border border-white/10 bg-[#13131a] p-1 shadow-xl"
        >
          {options.length === 0 ? (
            <div className="px-3 py-2 text-white/40" style={{ fontSize: 12 }}>
              No options
            </div>
          ) : (
            options.map((opt) => {
              const isSelected = selected.includes(opt.value);
              return (
                <button
                  key={opt.value}
                  type="button"
                  role="option"
                  aria-selected={isSelected}
                  onClick={() => toggle(opt.value)}
                  className={`flex w-full items-center justify-between gap-2 rounded px-2 py-1.5 text-left transition ${
                    isSelected
                      ? "bg-emerald-500/15 text-emerald-300"
                      : "text-white/70 hover:bg-white/5"
                  }`}
                  style={{ fontSize: 12 }}
                >
                  <span className="truncate">{opt.label}</span>
                  {isSelected && <Check className="h-3 w-3 shrink-0" />}
                </button>
              );
            })
          )}
        </div>
      )}
    </div>
  );
}
