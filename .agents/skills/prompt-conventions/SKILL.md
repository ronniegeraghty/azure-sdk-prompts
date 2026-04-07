---
name: "prompt-conventions"
description: "Prompt frontmatter format, properties, migration"
domain: "content"
confidence: "high"
source: "hyoka/internal/prompt/prompt.go, prompts/*.md files"
---

## Context

Hyoka prompts are Markdown files with YAML frontmatter. The frontmatter defines metadata (service, language, category, difficulty) that's used for filtering, tool selection, and grading. Prompts are organized by language/service and must follow a consistent structure.

## Prompt File Structure

```markdown
---
id: identity-dp-python-default-credential
service: identity
language: python
plane: data-plane
category: auth
difficulty: beginner
tags: [authentication, patterns]
---

# Prompt Title

Your task is to generate Python code that uses the Azure Identity library...

## Requirements

- Code should handle authentication
- Include error handling
- Add type hints

## Expected Output

A Python script that demonstrates the DefaultAzureCredential pattern.
```

## Frontmatter Fields

### Required Fields

- **id:** Unique identifier (pattern: `{service}-{plane}-{language}-{short-name}`)
  - Examples: `identity-dp-python-default-credential`, `key-vault-mp-dotnet-crud`
  - Must be unique across all prompts
  - Snake_case, lowercase

- **service:** Azure service name
  - Values: `identity`, `key-vault`, `storage`, `cosmos-db`, etc.
  - Used for `--service` filtering

- **language:** Programming language
  - Values: `python`, `dotnet`, `java`, `js-ts`, `go`, `rust`, `cpp`
  - Used for `--language` filtering and tool selection

- **plane:** Deployment plane
  - Values: `data-plane` or `management-plane`
  - Used for `--plane` filtering

### Optional Fields

- **category:** Use-case category
  - Examples: `auth`, `crud`, `pagination`, `streaming`
  - Used for `--category` filtering

- **difficulty:** Skill level for generated code
  - Values: `beginner`, `intermediate`, `advanced`
  - Informational for report context

- **tags:** Additional keyword filters (array)
  - Examples: `[authentication, error-handling, async]`
  - Used for `--tags` filtering

## File Organization

Prompts are stored in `prompts/` organized by service/language:

```
prompts/
  identity/
    python/
      default-credential.md
      managed-identity.md
    dotnet/
      default-credential.md
  key-vault/
    python/
      crud-secrets.md
    go/
      get-secret.md
```

## Prompt ID Naming Convention

Pattern: `{service}-{plane-abbrev}-{language}-{short-name}`

- **service:** Service name (identity, key-vault, storage, etc.)
- **plane-abbrev:** `dp` (data-plane) or `mp` (management-plane)
- **language:** `python`, `dotnet`, `java`, `js-ts`, `go`, `rust`, `cpp`
- **short-name:** Hyphenated description (max 4 words)

Examples:
- ✓ `identity-dp-python-default-credential`
- ✓ `key-vault-dp-dotnet-crud-secrets`
- ✓ `storage-mp-java-list-containers`
- ✗ `DefaultCredentialPython` (wrong case, no plane/service prefix)
- ✗ `identity-python-auth-with-logging-and-retry` (too long)

## Prompt Content Guidelines

### Header
Brief, actionable description of the task.

### Requirements Section
List concrete requirements for generated code:
- Functionality (what should the code do?)
- Error handling expectations
- Language-specific conventions (e.g., type hints for Python)
- API patterns to follow

### Expected Output
Describe the form of generated code (script vs. module, entry point, etc.)

### Example (optional)
Show a snippet of expected behavior or output format.

## Properties for Tool Selection

Properties in tool config filters match prompt metadata:

```yaml
# Config
tools:
  - name: python_test_runner
    properties:
      language: python    # Matches prompt.language
      service: key-vault  # Matches prompt.service
```

Prompt metadata is compared against tool properties — tools only appear if all properties match.

## Prompt Validation

Validate prompts before running:

```bash
go run ./hyoka validate --prompt-id identity-dp-python-default-credential

# Or validate all prompts
go run ./hyoka validate
```

The validator checks:
- Required frontmatter fields present
- ID uniqueness
- Language/service/plane values valid
- Prompt file exists and is readable

## Property Migration (Advanced)

For changes to property schema or tool filtering logic:

1. **Update prompt frontmatter** in batch:
   ```bash
   # Script to add/rename fields across all prompts
   for f in prompts/*/*.md; do
     # Add new field or rename existing
     sed -i 's/old_field:/new_field:/' "$f"
   done
   ```

2. **Update tool filters** in config:
   ```yaml
   # Before
   properties:
     lang: python
   
   # After (new property name)
   properties:
     language: python
   ```

3. **Test with dry-run:**
   ```bash
   go run ./hyoka run --service identity --language python --dry-run
   ```

## Listing and Filtering Prompts

```bash
# List all prompts
go run ./hyoka list

# Filter by service
go run ./hyoka list --service key-vault

# Filter by language + service
go run ./hyoka list --service identity --language python

# Filter by tags
go run ./hyoka list --tags "auth,crud"
```

## Anti-Patterns

- IDs with uppercase or special characters (use snake_case)
- Inconsistent plane values (`data_plane` vs `data-plane`)
- Missing required frontmatter fields
- Prompts with same ID in different directories (validator should catch)
- Tool properties that don't match any prompt metadata
- Overly specific prompts (too narrow to be useful across models)

## Related Code Locations

- **Prompt structure:** `hyoka/internal/prompt/prompt.go`
- **Validation logic:** `hyoka/internal/validate/validate.go`
- **Example prompts:** `hyoka/prompts/*/`
- **List command:** `hyoka/cmd/list.go`
