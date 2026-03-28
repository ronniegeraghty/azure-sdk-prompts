# Prompt Authoring Guide

This guide explains how to write evaluation prompts for hyoka — from frontmatter schema to evaluation criteria and directory structure conventions.

## What Is a Prompt?

A prompt is a Markdown file (`.prompt.md`) that describes a code generation task for an AI agent. It includes:

1. **YAML frontmatter** — metadata for filtering and identification
2. **Prompt text** — exact instructions sent to the AI agent
3. **Evaluation criteria** — what the generated code should demonstrate
4. **Context** — why this prompt matters

hyoka discovers prompts automatically by scanning the `prompts/` directory for files ending in `.prompt.md`.

## File Format

```markdown
---
id: storage-dp-python-crud
service: storage
plane: data-plane
language: python
category: crud
difficulty: basic
description: >
  Can a developer upload, download, list, and delete blobs
  in Azure Blob Storage using the Python SDK?
sdk_package: azure-storage-blob
doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/storage-blob-readme
tags:
  - blob
  - crud
  - getting-started
created: 2025-07-27
author: ronniegeraghty
---

# CRUD Blobs: Storage (Python)

## Prompt

Write a Python script that performs CRUD operations on Azure Blob Storage:
1. Create a BlobServiceClient using DefaultAzureCredential
2. Create a container named "my-container" if it doesn't exist
3. Upload a local file "report.csv" as a blob named "reports/report.csv"
4. List all blobs in the container and print each blob's name and content length
5. Download the blob and save it to "report-downloaded.csv"
6. Delete the blob and then delete the container

Show required pip packages and proper error handling with HttpResponseError.

## Evaluation Criteria

- Uses `azure-storage-blob` and `azure-identity` packages
- `BlobServiceClient` with `DefaultAzureCredential`
- `BlobClient.upload_blob()` with `overwrite` parameter
- `ContainerClient.list_blobs()` iteration
- `BlobClient.download_blob()` and `readall()` or `readinto()`
- `HttpResponseError` and `ResourceExistsError` handling

## Context

Tests whether the agent can produce a complete CRUD workflow using the
Azure Blob Storage Python SDK with proper authentication and error handling.
```

## Frontmatter Schema

### Required Fields

| Field | Type | Values | Description |
|-------|------|--------|-------------|
| `id` | string | `{service}-{dp\|mp}-{language}-{slug}` | Unique identifier. Must match naming convention. |
| `service` | string | `storage`, `key-vault`, `cosmos-db`, `event-hubs`, `app-configuration`, `purview`, `digital-twins`, `identity`, `resource-manager`, `service-bus` | Target Azure service |
| `plane` | string | `data-plane`, `management-plane` | API plane |
| `language` | string | `dotnet`, `java`, `js-ts`, `python`, `go`, `rust`, `cpp` | Target programming language |
| `category` | string | `authentication`, `pagination`, `polling`, `retries`, `error-handling`, `crud`, `batch`, `streaming`, `auth`, `provisioning` | Use-case category |
| `difficulty` | string | `basic`, `intermediate`, `advanced` | Complexity level |
| `description` | string | Free text | 1–3 sentences describing what this prompt tests |
| `created` | string | `YYYY-MM-DD` | Creation date |
| `author` | string | GitHub username | Prompt author |

### Optional Fields

| Field | Type | Description |
|-------|------|-------------|
| `sdk_package` | string | SDK package name (e.g., `azure-storage-blob`, `com.azure:azure-storage-blob`) |
| `doc_url` | string | URL to official SDK documentation |
| `tags` | list | Free-form tags for additional filtering (e.g., `[blob, crud, async]`) |
| `expected_packages` | list | Packages the generated code should use |
| `expected_tools` | list | Tools the agent should invoke during generation |
| `starter_project` | string | Path to starter project directory (relative to prompt file) |
| `project_context` | map | Project setup configuration (`blank` or `existing`) |
| `reference_answer` | string | Path to reference implementation for comparison |
| `timeout` | int | Session timeout in seconds |

## ID Naming Convention

Prompt IDs follow a strict pattern:

```
{service}-{dp|mp}-{language}-{slug}
```

Where:
- `dp` = data-plane
- `mp` = management-plane
- `slug` = short kebab-case descriptor

**Examples:**
- `storage-dp-python-crud`
- `key-vault-dp-java-encrypted-uploader`
- `storage-mp-go-account-mgmt`
- `cosmos-db-dp-dotnet-crud`

The `validate` command enforces this convention.

## Directory Structure

Prompts are organized hierarchically:

```
prompts/
├── storage/
│   ├── data-plane/
│   │   ├── dotnet/
│   │   │   └── authentication.prompt.md
│   │   └── python/
│   │       ├── crud-blobs.prompt.md
│   │       └── pagination-list-blobs.prompt.md
│   └── management-plane/
│       └── go/
│           └── storage-account-mgmt.prompt.md
├── key-vault/
│   └── data-plane/
│       └── java/
│           └── encrypted-uploader.prompt.md
└── cosmos-db/
    └── data-plane/
        └── python/
            └── crud.prompt.md
```

Convention: `prompts/{service}/{plane}/{language}/{slug}.prompt.md`

## Sections

### `## Prompt` (Required)

The exact text sent to the AI agent. Be specific and actionable:

- List numbered steps or clear requirements
- Mention authentication approach (e.g., `DefaultAzureCredential`)
- Specify expected package versions where relevant
- Describe error handling expectations
- Include any specific API methods the code should use

**Good example:**
```markdown
## Prompt

Write a Python script that performs CRUD operations on Azure Blob Storage:
1. Create a BlobServiceClient using DefaultAzureCredential
2. Create a container named "my-container" if it doesn't exist
3. Upload a local file "report.csv" as blob "reports/report.csv"
4. List all blobs and print name + content length
5. Download the blob and save as "report-downloaded.csv"
6. Delete the blob and the container

Show required pip packages and proper error handling with HttpResponseError.
```

**Bad example:**
```markdown
## Prompt

Write code that uses Azure Blob Storage.
```

### `## Evaluation Criteria` (Recommended)

A bullet list of specific, testable requirements. Each criterion is scored as pass/fail by the review panel.

Tips:
- Use concrete API names, class names, and method names
- Cover imports, configuration, error handling, and cleanup
- Each criterion should be independently verifiable
- Be specific about what "correct" looks like

**Good example:**
```markdown
## Evaluation Criteria

- Uses `azure-storage-blob` and `azure-identity` packages
- `BlobServiceClient` created with `DefaultAzureCredential`
- `BlobClient.upload_blob()` called with `overwrite` parameter
- `ContainerClient.list_blobs()` iteration to enumerate blobs
- `HttpResponseError` handling for API errors
```

### `## Context` (Recommended)

Explains why the prompt matters and what quality aspect it evaluates. Helps other contributors understand the testing motivation.

```markdown
## Context

Tests whether the agent can produce a complete CRUD workflow using the
Azure Blob Storage Python SDK with proper authentication and error handling.
This is a foundational scenario — most SDK users start here.
```

## Creating a New Prompt

### Option A: Interactive Scaffolder

```bash
go run ./hyoka new-prompt
```

Prompts for service, plane, language, category, difficulty, description, and slug. Generates the file with populated frontmatter.

### Option B: Copy the Template

```bash
cp templates/prompt-template.prompt.md \
   prompts/<service>/<plane>/<language>/<slug>.prompt.md
```

Edit the file to fill in frontmatter and prompt content.

### Option C: Copy an Existing Prompt

Find a similar prompt and copy it:

```bash
go run ./hyoka list --service storage --language python
cp prompts/storage/data-plane/python/crud-blobs.prompt.md \
   prompts/storage/data-plane/python/my-new-prompt.prompt.md
```

### After Creating

Always validate:

```bash
go run ./hyoka validate
```

## Best Practices

1. **Be specific in prompts.** Numbered steps, concrete API names, expected behavior.
2. **Test what matters.** Each evaluation criterion should test a distinct quality aspect.
3. **Use the right difficulty.** `basic` = single API call, `intermediate` = multi-step workflow, `advanced` = multi-service or complex patterns.
4. **Include `sdk_package` and `doc_url`.** Helps reviewers verify the agent used correct packages.
5. **Tag generously.** Tags enable cross-cutting analysis (e.g., all `async` prompts, all `multi-service` prompts).
6. **Add context.** Explain why the prompt exists — future contributors will thank you.
7. **Mention authentication.** Almost every Azure SDK scenario needs auth — be explicit about the expected approach.
8. **Cover error handling.** Include specific exception types in evaluation criteria.
