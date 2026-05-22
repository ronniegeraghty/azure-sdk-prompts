import { useState, useEffect } from "react";
import { fetchRuns } from "../data/api";
import type { RunSummary } from "../data/types";

export function useRuns() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    async function load() {
      try {
        const data = await fetchRuns();
        if (cancelled) return;
        setRuns(data);
      } catch (e) {
        if (!cancelled) setError(e instanceof Error ? e.message : "Failed to load runs");
      } finally {
        if (!cancelled) setLoading(false);
      }
    }
    load();
    return () => { cancelled = true; };
  }, []);

  return { runs, loading, error };
}
