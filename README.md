# azure-sdk-prompts

A curated library of sample prompts for testing Azure SDK documentation quality using the [doc-review-agent](https://github.com/coreai-microsoft/doc-review-agent).

Each prompt targets a specific service, language, plane, and use-case category. An evaluation script runs `doc-agent evaluate` against all or a filtered subset of prompts, producing timestamped HTML reports.

## Quick Start

### Adding a New Prompt

```bash
# 1. Copy the template
cp templates/prompt-template.prompt.md \
   prompts/<service>/<plane>/<language>/<use-case>.prompt.md

# 2. Edit the file — fill in all frontmatter fields and write your prompt
#    Required: id, service, plane, language, category, difficulty,
#    description, created, author

# 3. Validate your prompt
python scripts/validate-prompts.py

# 4. Regenerate the manifest
python scripts/generate-manifest.py

# 5. Commit
git add prompts/ manifest.yaml
git commit -m "prompt: add <service> <plane> <language> <category>"
```

### Running Evaluations

```bash
# Run ALL prompts
python scripts/run-evals.py

# Filter by service
python scripts/run-evals.py --service storage

# Filter by language
python scripts/run-evals.py --language dotnet

# Combine filters (AND logic)
python scripts/run-evals.py --service storage --language dotnet

# More filter combinations
python scripts/run-evals.py --plane data-plane --category authentication
python scripts/run-evals.py --tags identity

# Single prompt by ID or path
python scripts/run-evals.py --prompt-id storage-dp-dotnet-auth
python scripts/run-evals.py --prompt prompts/storage/data-plane/dotnet/authentication.prompt.md

# Dry run — list matching prompts without executing
python scripts/run-evals.py --service storage --dry-run
```

### Viewing Reports

After a run, reports are saved in `reports/runs/<timestamp>/`:

```bash
# Open the latest report for a specific prompt
open reports/runs/latest/storage/data-plane/dotnet/authentication/report.html

# Check the run summary
cat reports/runs/latest/run-metadata.yaml
```

## Repo Structure

```
azure-sdk-prompts/
├── README.md
├── LICENSE
├── manifest.yaml                      # Auto-generated index of all prompts
├── prompts/                           # Prompt library
│   ├── storage/
│   │   ├── data-plane/
│   │   │   ├── dotnet/
│   │   │   │   └── authentication.prompt.md
│   │   │   └── python/
│   │   │       └── pagination-list-blobs.prompt.md
│   │   └── management-plane/
│   │       └── ...
│   └── key-vault/
│       └── data-plane/
│           └── python/
│               └── crud-secrets.prompt.md
├── scripts/
│   ├── run-evals.py                   # Evaluation runner with composable filters
│   ├── generate-manifest.py           # Regenerate manifest.yaml from prompts
│   └── validate-prompts.py            # Validate prompt frontmatter
├── reports/
│   └── runs/                          # Timestamped evaluation results
│       └── <timestamp>/
│           ├── run-metadata.yaml
│           └── <service>/<plane>/<language>/<prompt>/
│               ├── report.html
│               ├── task.md
│               ├── execution.log
│               └── observations.md
└── templates/
    └── prompt-template.prompt.md      # Template for new prompts
```

## Tagging System

Every prompt uses YAML frontmatter for filtering and indexing:

| Field | Required | Values |
|---|---|---|
| `id` | ✅ | `{service}-{dp\|mp}-{lang}-{category-slug}` |
| `service` | ✅ | `storage`, `key-vault`, `cosmos-db`, `event-hubs`, `app-configuration`, `purview`, `digital-twins` |
| `plane` | ✅ | `data-plane`, `management-plane` |
| `language` | ✅ | `dotnet`, `java`, `js-ts`, `python`, `go`, `rust`, `cpp` |
| `category` | ✅ | `authentication`, `pagination`, `polling`, `retries`, `error-handling`, `crud`, `batch`, `streaming` |
| `difficulty` | ✅ | `basic`, `intermediate`, `advanced` |
| `description` | ✅ | What this prompt tests (1-3 sentences) |
| `sdk_package` | ❌ | SDK package name |
| `doc_url` | ❌ | Link to the docs page being evaluated |
| `tags` | ❌ | Free-form tags for additional filtering |
| `created` | ✅ | Date (YYYY-MM-DD) |
| `author` | ✅ | GitHub username |

## Prerequisites

- Python 3.9+
- [doc-review-agent](https://github.com/coreai-microsoft/doc-review-agent) CLI (`doc-agent`) installed
- PyYAML: `pip install pyyaml`

## License

[MIT](LICENSE)
