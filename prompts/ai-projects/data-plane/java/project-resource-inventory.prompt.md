---
id: ai-projects-dp-java-project-resource-inventory
properties:
  service: ai-projects
  plane: data-plane
  language: java
  category: projects
  difficulty: intermediate
  description: "Can an application enumerate and retrieve typed Azure AI project connections and deployments?"
  sdk_package: com.azure:azure-ai-projects
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - connections
  - deployments
  - pagination
---

# Azure AI project resource inventory (Java)

## Prompt

Create a complete, runnable Java console application that uses
`com.azure:azure-ai-projects` to inspect connections and model deployments in a
Microsoft Foundry project.

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

Use the synchronous SDK clients throughout.

## Evaluation Criteria

### Project inventory workflow

- Uses `AIProjectClientBuilder` to create `ConnectionsClient` and
  `DeploymentsClient` for `FOUNDRY_PROJECT_ENDPOINT`.
- Iterates every `Connection` from `ConnectionsClient.listConnections` and reads its
  typed name, type, target, and default properties.
- Calls `ConnectionsClient.getConnection` for `CONNECTION_NAME` with credentials
  disabled, and doesn't print credentials.
- Iterates every `Deployment` from `DeploymentsClient.listDeployments`.
- Narrows deployments to `ModelDeployment` before reading publisher, model name, and
  model version.
- Calls `DeploymentsClient.getDeployment` for `DEPLOYMENT_NAME` and rejects a result
  that isn't a `ModelDeployment`.

### Scenario-specific anti-patterns

- Does not replace `PagedIterable` traversal with a single assumed page.
- Does not select resources only from the locally enumerated results instead of
  exercising the SDK's name-based get operations.
- Does not request or display connection credentials.

## Context

The reference application is in
`reference-apps/ai-projects/project-resource-inventory/java`.
