---
id: ai-projects-dp-js-ts-project-resource-inventory
properties:
  service: ai-projects
  plane: data-plane
  language: js-ts
  category: projects
  difficulty: intermediate
  description: "Can an application enumerate and retrieve typed Azure AI project connections and deployments?"
  sdk_package: "@azure/ai-projects"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - connections
  - deployments
  - pagination
---

# Azure AI project resource inventory (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application that uses
`@azure/ai-projects` to inspect connections and model deployments in a Microsoft
Foundry project.

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
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Project inventory workflow

- Creates `AIProjectClient` for `FOUNDRY_PROJECT_ENDPOINT`.
- Iterates `client.connections.list()` with `for await...of` and reads typed
  connection properties.
- Calls `client.connections.get` for `CONNECTION_NAME`, not `getWithCredentials`,
  and doesn't print credentials.
- Iterates `client.deployments.list()` with `for await...of`.
- Uses the `ModelDeployment` discriminator and properties before printing publisher,
  model name, and model version.
- Calls `client.deployments.get` for `DEPLOYMENT_NAME` and rejects a result whose
  type isn't `ModelDeployment`.

### Scenario-specific anti-patterns

- Does not replace async pageable iteration with a single assumed page.
- Does not select resources only from the locally enumerated results instead of
  exercising the SDK's name-based get operations.
- Does not request or display connection credentials.

## Context

The reference application is in
`reference-apps/ai-projects/project-resource-inventory/js-ts`.
