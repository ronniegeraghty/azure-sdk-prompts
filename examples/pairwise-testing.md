# Pairwise Tool Ablation Testing

## Overview

Pairwise testing is a tool-ablation feature that measures the individual impact of each tool in a config on code generation quality. When you run hyoka with `--pairwise`, it generates **N+1 config variants**:

1. **Baseline:** The original config with all tools enabled
2. **N variants:** One config per tool, with that tool removed

The evaluation runs all variants and reports the score difference between baseline and each "without-X" variant, showing you which tools provide the most value.

## When to Use Pairwise Testing

- **Measuring tool impact:** Determine which MCP servers or Copilot tools contribute most to code quality
- **Cost optimization:** Identify low-value tools that can be removed to reduce latency/cost
- **Tool validation:** Verify that a new tool actually improves generated code
- **Comparative analysis:** Understand tool interactions (does removing tool A affect tool B's utility?)

## Example Commands

### Run pairwise with a single prompt

```bash
# Evaluate one prompt with tool ablation
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config azure-mcp/claude-opus-4.6 --pairwise
```

This generates configs:
- `azure-mcp/claude-opus-4.6` (baseline)
- `azure-mcp/claude-opus-4.6-without-azure-mcp`
- `azure-mcp/claude-opus-4.6-without-web-fetch`
- etc.

### Run pairwise with filtered prompts

```bash
# Test tool impact across multiple identity+Python prompts
go run ./hyoka run --service identity --language python \
  --config azure-mcp/claude-opus-4.6 --pairwise
```

### Run pairwise with multiple configs

```bash
# Compare tool impact across baseline and azure-mcp configs
go run ./hyoka run --prompt-id identity-dp-python-default-credential \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6" --pairwise
```

Each config spawns its own N+1 variants.

## Interpreting Results

The HTML report shows:

- **Baseline score:** The score achieved with all tools enabled
- **Without-X scores:** Score for each variant with one tool removed
- **Impact:** `baseline_score - without_X_score`
  - **Positive impact:** Removing the tool decreased quality (tool is valuable)
  - **Negative impact:** Removing the tool increased quality (tool may be harmful or redundant)
  - **Zero impact:** Tool had no measurable effect on this prompt

Example:
```
Config: azure-mcp/claude-opus-4.6
Baseline: 85.2

azure-mcp/claude-opus-4.6-without-azure-mcp: 72.1  (impact: +13.1) ← azure-mcp is very valuable
azure-mcp/claude-opus-4.6-without-web-fetch: 84.8  (impact: +0.4)  ← web-fetch has minimal impact
```

## Important Notes

- **`always_on: true` tools are never toggled:** Tools marked as always-on (like `create`, `edit`) are excluded from pairwise variants since they're essential for code generation.
- **Pairwise multiplies run time:** If your config has 5 toggleable tools, you'll run 6 evaluations (1 baseline + 5 variants). Plan accordingly.
- **Use with focused prompts:** Pairwise is most useful on representative prompts that exercise the tools you're testing.

## Advanced: Combining with Other Flags

```bash
# Pairwise with resource monitoring and debug logging
go run ./hyoka run --service key-vault --language python \
  --config azure-mcp/claude-opus-4.6 --pairwise \
  --monitor-resources --log-level debug --log-file pairwise-debug.log

# Dry run to see what configs would be generated
go run ./hyoka run --prompt-id cosmos-db-dp-python-pagination \
  --config azure-mcp/claude-opus-4.6 --pairwise --dry-run
```

## See Also

- `examples/configs/example-conditional-tools.yaml` — Shows how to configure tools with `when` conditions and `always_on`
- `docs/configuration.md` — Full tool configuration reference
- `examples/filtering-and-flags.md` — All hyoka CLI flags and filtering options
