---
name: "contributor-guide"
description: "How to contribute, build, test"
domain: "collaboration"
confidence: "high"
source: "hyoka/README.md, docs/contributing.md"
---

## Context

Hyoka is an open evaluation tool for Azure SDK code generation. Contributors should follow the setup and workflow documented here to ensure smooth development and testing.

## Development Setup

### Prerequisites

- **Go 1.24.5+** (check with `go version`)
- **Copilot SDK** installed and authenticated
- **git** configured with GitHub SSH access

### Clone and Build

```bash
# Clone the repo
git clone https://github.com/ronniegeraghty/hyoka.git
cd hyoka

# Install dependencies (uses go.work for workspace)
go mod download

# Build the CLI
go build ./hyoka/...

# Run the CLI
go run ./hyoka <command>
```

## Development Workflow

### Before Starting Work

1. **Create a branch** from `ronniegeraghty/dev`:
   ```bash
   git checkout ronniegeraghty/dev
   git pull origin ronniegeraghty/dev
   git checkout -b ronniegeraghty/issue-{N}-{description}
   ```

2. **Configure git identity:**
   ```bash
   git config user.name "ronniegeraghty"
   git config user.email "ronniegeraghty@users.noreply.github.com"
   ```

### Making Changes

1. **Edit code** in `hyoka/internal/` or `hyoka/cmd/`
2. **Run tests** to verify no regressions:
   ```bash
   go test -race ./hyoka/...
   ```
3. **Run linter** (if configured):
   ```bash
   go fmt ./hyoka/...
   ```
4. **Test manually** with a quick eval run:
   ```bash
   go run ./hyoka run --prompt-id key-vault-dp-python-crud \
     --config baseline/claude-opus-4.6 --log-level debug
   ```

### Before Committing

1. **Ensure tests pass** with race detector:
   ```bash
   go test -race ./hyoka/...
   ```

2. **Clean up orphaned sessions** from test runs:
   ```bash
   go run ./hyoka clean
   ```

3. **Include the Co-authored-by trailer** in commit message:
   ```
   git commit -m "Fix eval engine timeout handling

   - Add timeout context to generation phase
   - Capture timeout error in action timeline
   
   Closes #162
   
   Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
   ```

### Pushing and Creating a PR

1. **Push to your fork:**
   ```bash
   gh auth switch --user ronniegeraghty
   git push origin ronniegeraghty/issue-{N}-{description}
   ```

2. **Create a PR** against `ronniegeraghty/dev`:
   ```bash
   gh pr create --base ronniegeraghty/dev \
     --title "Fix: eval engine timeout" \
     --body "Closes #162"
   ```

3. **Link to board:** Edit issue on Azure/projects/424 and set Status → "In Progress"

## Testing During Development

### Quick Test Run

For iterating on changes, run a single prompt × single config (fastest):

```bash
go run ./hyoka run --prompt-id key-vault-dp-python-crud \
  --config baseline/claude-opus-4.6 --log-level debug
```

Python prompts finish quickest (5-10 minutes). After each run, clean up:

```bash
go run ./hyoka clean
```

### Check Logs

```bash
# View the debug log
cat hyoka-debug.log

# Search for role-prefixed output
grep "role=" hyoka-debug.log | head -20
```

### Browse Results

```bash
go run ./hyoka serve
# Open http://localhost:8080
```

## Common Development Tasks

### Adding a New Grader Type

1. **Create the grader:** `hyoka/internal/graders/my_grader.go`
2. **Implement the interface:**
   ```go
   type MyGrader struct { /* fields */ }
   func (g *MyGrader) Kind() string { return "my" }
   func (g *MyGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) { /* ... */ }
   ```
3. **Register in factory:** `hyoka/internal/graders/registry.go`
4. **Write tests:** `hyoka/internal/graders/my_grader_test.go`
5. **Update config examples:** `configs/*.yaml`

### Adding a New CLI Command

1. **Create command file:** `hyoka/cmd/mycommand.go`
2. **Implement Cobra command:**
   ```go
   var myCmd = &cobra.Command{
       Use: "mycommand",
       RunE: func(cmd *cobra.Command, args []string) error { /* ... */ },
   }
   func init() { rootCmd.AddCommand(myCmd) }
   ```
3. **Document in help text**
4. **Write tests:** `hyoka/cmd/mycommand_test.go`

### Updating Documentation

- **Architecture docs:** `docs/architecture.md`
- **CLI reference:** `docs/cli-reference.md`
- **Contributing guide:** `docs/contributing.md`
- **README:** `hyoka/README.md`

## Code Review Process

1. **Submit PR** with `Closes #{issue}` in description
2. **Address review comments** in follow-up commits
3. **Ensure tests pass** (CI runs on PR)
4. **Squash if requested** for history cleanliness
5. **Merge to dev** (maintainer handles merge)

## Release Workflow

(Handled by maintainers)

1. Tag release on `main` branch
2. Create GitHub Release with changelog
3. Upload binary artifacts

## Resources

- **Go stdlib docs:** https://pkg.go.dev/std
- **Cobra docs:** https://cobra.dev
- **slog guide:** https://pkg.go.dev/log/slog
- **GitHub Copilot CLI docs:** https://github.com/cli/cli/wiki/GitHub-Copilot-CLI

## Anti-Patterns

- Working directly on `main` branch (use feature branches)
- Committing without running `go test -race`
- Leaving orphaned Copilot sessions (always run `hyoka clean`)
- Skipping tests to speed up iteration (tests catch regressions)
- Hardcoding paths or config names in code
