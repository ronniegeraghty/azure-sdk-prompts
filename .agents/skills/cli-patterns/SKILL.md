---
name: "cli-patterns"
description: "Cobra commands in cmd/ package, flag conventions"
domain: "architecture"
confidence: "high"
source: "hyoka/cmd/*.go"
---

## Context

Hyoka's CLI is built with Cobra, a popular Go CLI framework. Commands are organized in the `cmd/` package with consistent flag naming and help text patterns.

## Command Structure

Each command is a Cobra command object:

```go
// cmd/run.go
var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Run evaluation",
    Long:  "Run an evaluation with prompts and configs.",
    
    Args: cobra.NoArgs,  // Positional arguments check
    
    RunE: func(cmd *cobra.Command, args []string) error {
        // Retrieve flags
        promptID, _ := cmd.Flags().GetString("prompt-id")
        config, _ := cmd.Flags().GetString("config")
        
        // Validate
        if promptID == "" && config == "" {
            return fmt.Errorf("either --prompt-id or --config required")
        }
        
        // Execute
        return run(promptID, config)
    },
}

func init() {
    rootCmd.AddCommand(runCmd)
    runCmd.Flags().StringP("prompt-id", "", "", "Single prompt ID")
    runCmd.Flags().StringP("config", "", "", "Config (comma-sep for multiple)")
}
```

## Flag Naming Conventions

- **Kebab-case:** `--log-level`, `--max-files`, `--prompt-id` (not `--logLevel`, `--maxfiles`)
- **Short flags:** Single-letter for common flags (e.g., `-v` for `--verbose`)
- **Suffixes:**
  - `--*-timeout` for duration limits (e.g., `--gen-timeout`)
  - `--skip-*` for skipping phases (e.g., `--skip-review`)
  - `--*-file` for file paths (e.g., `--log-file`)

## Flag Types and Patterns

**String flags:**
```go
cmd.Flags().StringP("config", "", "", "Config name")
```

**Comma-separated lists:**
```go
cmd.Flags().String("tags", "", "Comma-separated tags (must quote: \"auth,crud\")")
// Parsed as: strings.Split(value, ",")
```

**Boolean flags:**
```go
cmd.Flags().Bool("dry-run", false, "Show what would run without executing")
```

**Numeric flags:**
```go
cmd.Flags().Int("max-files", 50, "Max files per eval")
cmd.Flags().Duration("gen-timeout", 600*time.Second, "Generation timeout")
```

## Help Text Guidelines

Keep help text brief but actionable:

```
// Good
--config          Config name or comma-sep list
--prompt-id       Single prompt ID (e.g., identity-dp-python-default-credential)
--service         Azure service (e.g., identity, key-vault)
--dry-run         List matching prompts without executing

// Bad (too verbose, no examples)
--prompt-id       The prompt identifier
--config          The configuration to use
```

## Command Execution Pattern

```go
RunE: func(cmd *cobra.Command, args []string) error {
    // 1. Get flags
    promptID, _ := cmd.Flags().GetString("prompt-id")
    config, _ := cmd.Flags().GetString("config")
    
    // 2. Validate
    if promptID == "" {
        return fmt.Errorf("--prompt-id required")
    }
    
    // 3. Execute
    ctx := context.Background()
    return executeEval(ctx, promptID, config)
}
```

## Common Commands in Hyoka

- `run` — Execute evaluation
- `list` — List available prompts
- `validate` — Validate prompt schema
- `check-env` — Check environment (Copilot SDK, Go version)
- `clean` — Clean orphaned sessions
- `serve` — Start local web server for results

## Error Handling in Commands

Return errors from `RunE`. Cobra catches them and prints with exit code 1:

```go
RunE: func(cmd *cobra.Command, args []string) error {
    result, err := runEval(...)
    if err != nil {
        slog.Error("Evaluation failed", "error", err)
        return err  // Cobra handles exit
    }
    return nil
}
```

## Output and Logging

- **User output (results, progress):** Write to stdout
- **Diagnostic logs:** Use slog with `--log-level` flag
- **Errors:** Return error from `RunE`, let Cobra print

## Flag Validation Example

```go
RunE: func(cmd *cobra.Command, args []string) error {
    config, _ := cmd.Flags().GetString("config")
    allConfigs, _ := cmd.Flags().GetBool("all-configs")
    
    if config == "" && !allConfigs {
        return fmt.Errorf("either --config or --all-configs required")
    }
    
    if config != "" && allConfigs {
        return fmt.Errorf("cannot use both --config and --all-configs")
    }
    
    return runWithConfig(config, allConfigs)
}
```

## Related Code Locations

- **Root command:** `hyoka/cmd/root.go`
- **Subcommands:** `hyoka/cmd/*.go`
- **Flag parsing utilities:** `hyoka/cmd/common.go` (if shared parsing logic)

## Anti-Patterns

- CamelCase flag names
- Positional arguments instead of flags
- Ambiguous short flags (avoid single `-a` if context unclear)
- Printing errors and returning them (let Cobra handle)
- Hardcoding defaults without CLI override
