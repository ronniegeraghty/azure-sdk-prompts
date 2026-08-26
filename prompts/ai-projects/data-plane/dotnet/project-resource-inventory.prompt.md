---
id: ai-projects-dp-dotnet-project-resource-inventory
properties:
  service: ai-projects
  plane: data-plane
  language: dotnet
  category: projects
  difficulty: intermediate
  description: "Can an application enumerate and retrieve typed Azure AI project connections and deployments?"
  sdk_package: Azure.AI.Projects
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/ai.projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - connections
  - deployments
  - pagination
---

# Azure AI project resource inventory (.NET)

## Prompt

Create a complete, runnable .NET console application that uses `Azure.AI.Projects`
to inspect connections and model deployments in a Microsoft Foundry project.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `FOUNDRY_PROJECT_ENDPOINT`, `CONNECTION_NAME`, and `DEPLOYMENT_NAME`.
- Enumerate every project connection through the SDK's pageable API and print each
  connection's name, type, target, and default status.
- Retrieve `CONNECTION_NAME` without requesting credentials and print the same typed
  metadata.
- Enumerate every project deployment through the SDK's pageable API. For each model
  deployment, print its name, model publisher, model name, and model version.
- Retrieve `DEPLOYMENT_NAME`, require it to be a model deployment, and print the same
  typed model metadata.
- Include the project manifest and concise restore, build, and run commands.

Use asynchronous SDK operations throughout.

## Evaluation Criteria

### Project inventory workflow

- Creates `AIProjectClient` for `FOUNDRY_PROJECT_ENDPOINT`.
- Enumerates `Connections.GetConnectionsAsync` with `await foreach` and uses
  `AIProjectConnection` properties rather than raw JSON.
- Calls `Connections.GetConnectionAsync` for `CONNECTION_NAME` with
  `includeCredentials: false` and doesn't print credentials.
- Enumerates `Deployments.GetDeploymentsAsync` with `await foreach`.
- Narrows `AIProjectDeployment` values to `ModelDeployment` before reading
  `ModelPublisher`, `ModelName`, and `ModelVersion`.
- Calls `Deployments.GetDeploymentAsync` for `DEPLOYMENT_NAME` and rejects a result
  that isn't a `ModelDeployment`.

### Scenario-specific anti-patterns

- Does not replace pageable iteration with a single assumed page.
- Does not select resources only from the locally enumerated results instead of
  exercising the SDK's name-based get operations.
- Does not request or display connection credentials.

## Context

The reference application is in
`reference-apps/ai-projects/project-resource-inventory/dotnet`.
