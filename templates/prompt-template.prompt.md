---
id: <service>-<dp|mp>-<language>-<category-slug>
service: # storage | key-vault | cosmos-db | event-hubs | app-configuration | purview | digital-twins
plane: # data-plane | management-plane
language: # dotnet | java | js-ts | python | go | rust | cpp
category: # authentication | pagination | polling | retries | error-handling | crud | batch | streaming
difficulty: # basic | intermediate | advanced
description: >
  One to three sentences describing what this prompt tests.
sdk_package: # e.g., Azure.Storage.Blobs
api_version: # e.g., "2024-11-04"
doc_url: # https://learn.microsoft.com/...
tags: []
created: # YYYY-MM-DD
author: # GitHub username
---

# <Title>: <Service> (<Language>)

## Prompt

Write the exact prompt text here. This is what gets passed to `doc-agent evaluate`.
Be specific about what you're asking the agent to accomplish using only the SDK documentation.

## Expected Coverage

The documentation should cover:
- Key concept or API the prompt tests
- Expected packages or imports
- Configuration or setup steps
- Error handling guidance

## Context

Why this prompt matters and what documentation gap or quality aspect it evaluates.
