# Prompt Authoring Guide

Prompts are Markdown files with YAML frontmatter that define evaluation scenarios for AI code generation.

## File Naming

Prompts use the `.prompt.md` or `.prompt.yaml` extension:

```
prompts/
  storage/
    data-plane/
      python/
        crud/
          blob-upload.prompt.md
          blob-download.prompt.yaml
```

## Frontmatter Schema (Markdown Format)

Prompts use YAML frontmatter with `id` and `tags` at the top level, and all other metadata nested under a `properties:` block:

```yaml
---
id: storage-dp-python-crud              # Unique identifier (required)
tags:                                   # Optional tags for categorization
  - blob
  - upload
  - getting-started
properties:
  service: storage                      # Azure service name (required)
  plane: data-plane                     # data-plane or management-plane (required)
  language: python                      # Programming language (required)
  category: crud                        # Use-case category (required)
  difficulty: basic                     # basic, intermediate, advanced (required)
  description: "Upload a blob..."       # Short description
  sdk_package: azure-storage-blob       # Primary SDK package
  doc_url: https://learn.microsoft.com/...  # Reference documentation
  created: "2026-01-15"                 # Creation date
  author: your-name                     # Original author
expected_packages:                      # Packages the code should use (optional)
  - azure-storage-blob
  - azure-identity
expected_tools:                         # Tools the agent should invoke (optional)
  - create_file
  - run_terminal_command
starter_project: ./starter/             # Optional starter project directory
reference_answer: ./reference/          # Optional reference answer directory
timeout: 300                            # Prompt-specific timeout in seconds (optional)
max_session_actions: 50                 # Max actions per Copilot session (optional)
max_turns: 25                           # Max conversation turns (optional)
project_context:                        # Project-level context variables (optional)
  variable_name: value
---
```

## YAML-Only Format (`.prompt.yaml`)

Prompts can also be authored as pure YAML files with a `.prompt.yaml` or `.prompt.yml` extension. This is useful for programmatic generation or when frontmatter-in-markdown feels awkward:

```yaml
# storage-dp-python-crud.prompt.yaml
id: storage-dp-python-crud
tags:
  - blob
  - upload
properties:
  service: storage
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: "Upload a blob to Azure Storage"
  sdk_package: azure-storage-blob
prompt_text: |
  Write a Python script that uploads a file to Azure Blob Storage
  using the azure-storage-blob SDK with DefaultAzureCredential.

# graders: Inline evaluation criteria using the unified grader schema
# See docs/graders/ for full schema reference (prompt, program, workspace, tool, activity types)
graders:
  - type: prompt
    name: Correctness
    weight: 1.0
    checks:
      - Uses BlobServiceClient with DefaultAzureCredential
      - Creates container if it doesn't exist
      - Uploads file with proper content type detection
```

The YAML format uses the same `id`, `tags`, and `properties:` structure as Markdown. The `prompt_text` field replaces the `## Prompt` section and the `graders:` field replaces the `## Evaluation Criteria` section.

### Required Fields

- `id` — Unique prompt identifier (used in reports and filtering)
- `service` — Azure service (e.g., `storage`, `key-vault`, `cosmos-db`)
- `plane` — `data-plane` or `management-plane`
- `language` — `python`, `java`, `dotnet`, `go`, `javascript`, `typescript`
- `category` — Use-case category (e.g., `crud`, `auth`, `encryption`)
- `difficulty` — `basic`, `intermediate`, `advanced`

## Prompt Body

After the frontmatter, use Markdown sections:

```markdown
# Storage Blob Upload (Python)

## Prompt

Write a Python script that uploads a file to Azure Blob Storage
using the azure-storage-blob SDK with DefaultAzureCredential.

## Evaluation Criteria

- Uses BlobServiceClient with DefaultAzureCredential
- Creates container if it doesn't exist
- Uploads file with proper content type detection
- Handles BlobAlreadyExistsError gracefully

## Notes

Optional notes for the prompt author (not sent to the agent).
```

### Sections

| Section | Purpose |
|---------|---------|
| `## Prompt` | The actual prompt sent to the AI agent (required) |
| `## Evaluation Criteria` | Prompt-specific criteria for the reviewer (Tier 3) |
| `## Notes` | Author notes (not used in evaluation) |

## Tiered Evaluation Criteria

hyoka supports three tiers of evaluation criteria:

1. **Tier 1 — General** (always applied): 5 general criteria from the rubric (Code Builds, Latest Package Versions, Best Practices, Error Handling, Code Quality)

2. **Tier 2 — Attribute-Matched** (conditional): YAML files in a `criteria/` directory that activate based on prompt language, service, or other metadata. For example, `criteria/language/java.yaml` adds Java-specific criteria to all Java prompts.

3. **Tier 3 — Prompt-Specific** (per-prompt): Either the `## Evaluation Criteria` section in each `.prompt.md` file, or the `graders:` field in `.prompt.yaml` files.

All tiers are merged and sent to the reviewer as a unified criteria list. Use `--criteria-dir criteria/` to enable Tier 2 criteria.

### Inline Graders in Prompts

Both `.prompt.md` (frontmatter) and `.prompt.yaml` (top-level) support inline `graders:` using the unified grader schema. Inline graders are evaluated after implicit criteria (from `## Evaluation Criteria` markdown sections) and before matched criteria-file graders.

**Example: Inline grader in `.prompt.md` frontmatter:**

```markdown
---
id: storage-dp-python-crud
tags:
  - blob
  - crud
graders:
  - type: prompt
    name: Correctness
    weight: 1.0
    checks:
      - Uses BlobServiceClient with DefaultAzureCredential
      - Proper error handling for network failures
---

# Storage Blob Upload (Python)

## Prompt

Write a Python script that uploads a file to Azure Blob Storage...

## Evaluation Criteria

- Uses BlobServiceClient with DefaultAzureCredential
- Creates container if it doesn't exist
```

**Important:** Inline graders **forbid `when:` clauses** (hard error). Inline graders apply only to their specific prompt file and always execute. See [docs/graders/](./docs/graders/) for full schema reference.

### Tags for Grader Matching

Prompts can define a `tags:` field at the top level of frontmatter (both `.prompt.md` and `.prompt.yaml`). Criteria-file graders can match against these tags using the `when: { tags: { is: [...], not: [...] } }` condition:

```yaml
# Prompt file
id: identity-dp-python-auth
tags:
  - msal
  - multi-tenant
```

```yaml
# criteria/language/python.yaml — matches only multi-tenant prompts
graders:
  - type: prompt
    name: MSAL Best Practices
    when:
      tags:
        is: [multi-tenant]
    checks:
      - Uses MSAL 1.0+ for multi-tenant flow
```

Tags are case-insensitive for matching purposes.

### Criteria YAML Format

```yaml
# criteria/language/python.yaml
match:
  language: python
criteria:
  - name: DefaultAzureCredential Usage
    description: Authentication uses DefaultAzureCredential.
  - name: Context Manager for Clients
    description: SDK clients are used with `with` statements.
```

For advanced grader configuration including conditions (`when:`) and complex matching, see [docs/graders/index.md](./docs/graders/index.md).

## Near-Miss Detection

If no prompts are found, hyoka scans for common naming mistakes:

- `auth-prompt.md` → should be `auth.prompt.md` (dot, not hyphen)
- `auth.prompt.txt` → should be `auth.prompt.md` (wrong extension)
- `*.md` files with YAML frontmatter → may need `.prompt.md` suffix

## Creating New Prompts

Use the interactive scaffolding command:

```bash
hyoka new-prompt
```

This asks for service, language, plane, category, and difficulty, then generates a properly structured prompt file.
