// GroupBuilder — UI for creating/editing a single ComparisonGroup.
//
// Shows a name input, a color swatch, and a multi-select chip cluster for each
// filter dimension. Empty selection in a dimension means "match all".

import { useState } from "react";
import { ChevronDown, ChevronUp, X, Trash2 } from "lucide-react";
import {
  type ComparisonGroup,
  type FilterCatalog,
  type FilterDimension,
  FILTER_DIMENSIONS,
  GROUP_COLORS,
  describeFilters,
} from "../lib/comparison-groups";

interface Props {
  group: ComparisonGroup;
  catalog: FilterCatalog;
  matchCount: number;
  totalCount: number;
  onChange: (next: ComparisonGroup) => void;
  onRemove: () => void;
}

export function GroupBuilder({ group, catalog, matchCount, totalCount, onChange, onRemove }: Props) {
  const [expanded, setExpanded] = useState(true);

  function setName(name: string) {
    onChange({ ...group, name });
  }
  function setColor(color: string) {
    onChange({ ...group, color });
  }
  function toggleValue(dim: FilterDimension, value: string) {
    const current = group.filters[dim] ?? [];
    const next = current.includes(value) ? current.filter((v) => v !== value) : [...current, value];
    onChange({ ...group, filters: { ...group.filters, [dim]: next } });
  }
  function clearDim(dim: FilterDimension) {
    const next = { ...group.filters };
    delete next[dim];
    onChange({ ...group, filters: next });
  }

  const empty = matchCount === 0;

  return (
    <div
      className="rounded-xl border bg-white/[0.02] p-4"
      style={{ borderColor: empty ? "rgba(245, 158, 11, 0.4)" : "rgba(255,255,255,0.08)" }}
      data-testid="group-card"
    >
      {/* Header row */}
      <div className="flex items-center gap-3">
        <div
          className="h-3 w-3 shrink-0 rounded-full"
          style={{ background: group.color }}
          aria-label={`Group color ${group.color}`}
        />
        <input
          aria-label="Group name"
          value={group.name}
          onChange={(e) => setName(e.target.value)}
          className="flex-1 rounded-md border border-white/10 bg-white/[0.03] px-2 py-1 text-white outline-none focus:border-emerald-400/40"
          style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}
        />
        <span
          className={`rounded-md px-2 py-1 ${empty ? "bg-amber-500/10 text-amber-300" : "bg-white/5 text-white/50"}`}
          style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 11 }}
          title={empty ? "No evals match this group's filters" : "Matching evals out of total"}
        >
          {matchCount} / {totalCount}
        </span>
        <button
          onClick={() => setExpanded(!expanded)}
          aria-label={expanded ? "Collapse group" : "Expand group"}
          className="rounded-md p-1 text-white/40 hover:bg-white/5 hover:text-white/80"
        >
          {expanded ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
        </button>
        <button
          onClick={onRemove}
          aria-label="Remove group"
          className="rounded-md p-1 text-white/40 hover:bg-red-500/10 hover:text-red-400"
        >
          <Trash2 className="h-4 w-4" />
        </button>
      </div>

      {/* Filter summary line — visible when collapsed */}
      {!expanded && (
        <p className="mt-2 text-white/40" style={{ fontSize: 12 }}>
          {describeFilters(group)}
        </p>
      )}

      {/* Filter editors */}
      {expanded && (
        <div className="mt-4 space-y-3">
          {/* Color picker */}
          <div className="flex items-center gap-2">
            <span className="text-white/40" style={{ fontSize: 11 }}>
              Color
            </span>
            <div className="flex gap-1.5">
              {GROUP_COLORS.map((c) => (
                <button
                  key={c}
                  onClick={() => setColor(c)}
                  aria-label={`Pick color ${c}`}
                  className={`h-5 w-5 rounded-full transition-transform hover:scale-110 ${
                    c === group.color ? "ring-2 ring-white/50 ring-offset-2 ring-offset-[#0a0a0f]" : ""
                  }`}
                  style={{ background: c }}
                />
              ))}
            </div>
          </div>

          {FILTER_DIMENSIONS.map((dim) => {
            const values = catalog[dim.key];
            const selected = group.filters[dim.key] ?? [];
            if (values.length === 0) return null;
            return (
              <div key={dim.key}>
                <div className="mb-1.5 flex items-center justify-between">
                  <span className="text-white/50" style={{ fontSize: 12 }}>
                    {dim.label}
                    {selected.length > 0 && (
                      <span className="ml-1.5 text-white/30">({selected.length})</span>
                    )}
                  </span>
                  {selected.length > 0 && (
                    <button
                      onClick={() => clearDim(dim.key)}
                      className="text-white/30 transition-colors hover:text-white/60"
                      style={{ fontSize: 11 }}
                    >
                      <X className="inline h-3 w-3" /> Clear
                    </button>
                  )}
                </div>
                <div className="flex flex-wrap gap-1.5">
                  {values.map((v) => {
                    const isOn = selected.includes(v);
                    return (
                      <button
                        key={v}
                        onClick={() => toggleValue(dim.key, v)}
                        aria-pressed={isOn}
                        className={`rounded-md border px-2 py-0.5 transition-colors ${
                          isOn
                            ? "border-emerald-400/40 bg-emerald-500/10 text-emerald-200"
                            : "border-white/10 bg-white/[0.03] text-white/50 hover:border-white/20 hover:text-white/80"
                        }`}
                        style={{ fontSize: 11, fontFamily: "'JetBrains Mono', monospace" }}
                      >
                        {v}
                      </button>
                    );
                  })}
                </div>
              </div>
            );
          })}

          {empty && (
            <p className="rounded-md bg-amber-500/5 px-3 py-2 text-amber-300" style={{ fontSize: 12 }}>
              No evals match these filters. Loosen a dimension or pick different values.
            </p>
          )}
        </div>
      )}
    </div>
  );
}
