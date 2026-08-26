---
name: prompt-authoring
description: "Expert guidance for authoring and maintaining .prompt.md evaluation files for hyoka. Covers frontmatter formats, file structure, filtering, and best practices."
domain: "content"
confidence: "high"
source: "hyoka/internal/prompt/types.go, hyoka/internal/prompt/parser.go, prompts/"
---

# Prompt Authoring Skill

Expert guidance for authoring evaluation prompts for **hyoka**, the Azure SDK Prompt Evaluation Tool. Covers `.prompt.md` file format, frontmatter conventions, and tool filtering.

## Prompt File Format

Each prompt is a Markdown file with YAML frontmatter. Files live at:
```
prompts/{service}/{plane}/{language}/{slug}.prompt.md
```

## Frontmatter Formats

The parser accepts two frontmatter layouts. **Nested `properties:` is the current convention.**

### Nested Format (preferred)

Metadata lives inside a `properties:` map. The `Prompt` struct stores all values in `Properties map[string]string`.

```yaml
---
id: identity-dp-python-default-credential
properties:
  service: identity
  plane: data-plane
  language: python
  category: auth
  difficulty: basic
  description: "Can a developer set up DefaultAzureCredential?"
  sdk_package: azure-identity
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme
  created: "2025-07-28"
  author: ronniegeraghty
tags:
  - authentication
  - getting-started
---
```

### Flat Format (legacy, still supported)

Older prompts may use top-level fields. The parser auto-migrates them into the `Properties` map at load time.

```yaml
---
id: storage-dp-dotnet-auth
service: storage
plane: data-plane
language: dotnet
category: authentication
difficulty: basic
description: "Authenticate to Azure Blob Storage"
sdk_package: Azure.Storage.Blobs
created: 2025-07-27
author: ronniegeraghty
tags:
  - identity
---
```

Both formats produce identical `Prompt` structs. The `properties:` format is preferred for new prompts.

### Frontmatter Fields

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `id` | ✅ | string | Unique ID: `{service}-{dp\|mp}-{language}-{slug}` |
| `service` | ✅ | string | Azure service name |
| `plane` | ✅ | string | `data-plane` or `management-plane` |
| `language` | ✅ | string | Target programming language |
| `category` | ✅ | string | Use-case category |
| `difficulty` | ✅ | string | `basic`, `intermediate`, or `advanced` |
| `description` | ✅ | string | 1–3 sentences: what this prompt tests |
| `created` | ✅ | string | Date in `YYYY-MM-DD` format |
| `author` | ✅ | string | GitHub username |
| `sdk_package` | ❌ | string | Expected SDK package |
| `doc_url` | ❌ | string | Library reference docs URL |
| `tags` | ❌ | list | Free-form tags for filtering |
| `expected_tools` | ❌ | list | Tool names the agent should use |
| `expected_packages` | ❌ | list | SDK packages the generated code should import |
| `starter_project` | ❌ | string | Path to starter project dir (relative to prompt file) |
| `project_context` | ❌ | string | `blank` (default) or `existing` |
| `reference_answer` | ❌ | string | Inline reference code for LLM-as-judge scoring |
| `timeout` | ❌ | int | Per-prompt timeout in seconds |

### Valid Values

**Services:** `storage`, `key-vault`, `cosmos-db`, `event-hubs`, `app-configuration`, `purview`, `digital-twins`, `identity`, `resource-manager`, `service-bus`, `ai-agents`, `ai-projects`, `text-translation`, `document-translation`

**Languages:** `dotnet`, `java`, `js-ts`, `python`, `go`, `rust`, `cpp`

**Planes:** `data-plane`, `management-plane`

**Categories:** `authentication`, `best-practices`, `pagination`, `polling`, `retries`, `error-handling`, `crud`, `batch`, `streaming`, `auth`, `provisioning`, `agents`, `projects`, `translation`

**Difficulties:** `basic`, `intermediate`, `advanced`

## ID & File Naming

**ID pattern:** `{service}-{dp|mp}-{language}-{slug}`
- `dp` = data-plane, `mp` = management-plane
- Slug is kebab-case, max ~4 words
- Examples: `storage-dp-dotnet-auth`, `cosmos-db-dp-python-crud-items`

**Filename:** `{slug}.prompt.md` — matches the slug portion of the ID.

## Tool Filtering with `When`

Config entries use a `When` map (not `properties` on `ToolEntry`) to conditionally include tools based on prompt metadata:

```yaml
tools:
  - name: python_test_runner
    when:
      language: python
  - name: key-vault-mcp
    when:
      service: key-vault
```

At runtime, `matchesWhen(entry.When, prompt.Properties)` checks that every key-value pair in `When` matches the prompt's `Properties` map. If `When` is empty, the tool is unconditionally included.

## Prompt Content Structure

```markdown
# Title: Service (Language)

## Prompt
The exact prompt text sent to the AI agent.

## Evaluation Criteria
Bullet list of what the generated code should demonstrate.

## Context
Why this prompt matters. (Human-readable only — not sent to the eval tool.)
```

**How sections flow:**
- **Generator** receives ONLY the `## Prompt` text
- **Reviewer** receives prompt text, generated code, the rubric, AND `## Evaluation Criteria`
- **`## Context`** is for human readers only

## Writing Good Prompts

### DO ✅
- Be specific: "Create a BlobServiceClient using DefaultAzureCredential" not "Use storage"
- Specify the SDK package by name
- Ask for complete, runnable code
- Test one concept per prompt
- Set clear expectations in Evaluation Criteria

### DON'T ❌
- Don't be vague — "Write some Azure code" is too broad to score
- Don't test multiple concepts in one prompt
- Don't assume context — agent starts from a blank workspace unless `project_context: existing`
- Don't skip the description — it's used in reports and filtering

## doc_url Convention

Point to the **library's API reference docs**, not quickstarts:

| Language | URL Pattern |
|----------|-------------|
| Python | `learn.microsoft.com/en-us/python/api/overview/azure/{pkg}-readme` |
| .NET | `learn.microsoft.com/en-us/dotnet/api/overview/azure/{pkg}-readme` |
| Java | `learn.microsoft.com/en-us/java/api/overview/azure/{pkg}-readme` |
| JS/TS | `learn.microsoft.com/en-us/javascript/api/overview/azure/{pkg}-readme` |
| Go | `pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/...` |
| Rust | `docs.rs/{crate}/latest/{crate}/` |

## CLI Commands

```bash
# Scaffold a new prompt interactively
go run ./hyoka new-prompt

# Validate all prompts
go run ./hyoka validate

# List and filter prompts
go run ./hyoka list --service key-vault --language python
go run ./hyoka list --tags "auth,crud"

# Dry run (show matching prompts without executing)
go run ./hyoka run --service storage --language dotnet --dry-run
```

## Anti-Patterns

- IDs with uppercase or special characters (use kebab-case)
- Inconsistent plane values (`data_plane` vs `data-plane`)
- Missing required frontmatter fields
- Prompts with the same ID in different directories
- Overly specific prompts (too narrow to be useful across models)
