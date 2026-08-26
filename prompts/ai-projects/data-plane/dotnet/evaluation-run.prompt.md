---
id: ai-projects-dp-dotnet-evaluation-run
properties:
  service: ai-projects
  plane: data-plane
  language: dotnet
  category: projects
  difficulty: advanced
  description: "Can an application use the project OpenAI Evals client to run a built-in evaluator and traverse metric results?"
  sdk_package: Azure.AI.Projects
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/ai.projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - evaluations
  - polling
  - metrics
---

# Azure AI project evaluation run (.NET)

## Prompt

Create a complete, runnable .NET console application using
`Azure.AI.Projects` 3.0.0-beta.1 to run an OpenAI-compatible evaluation in a
Microsoft Foundry project.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `FOUNDRY_PROJECT_ENDPOINT`.
- Obtain the project OpenAI client and its Evals client from `AIProjectClient`.
- Create an evaluation named `hyoka-f1-evaluation` with a custom item schema for
  string fields `query`, `response`, and `ground_truth`.
- Configure the Azure built-in evaluator `builtin.f1_score`, mapping `response` and
  `ground_truth` to the corresponding item fields.
- Create a run named `hyoka-f1-run` against an inline JSONL data source containing
  this exact item: query `What is the capital of France?`, response `Paris`, and
  ground truth `Paris`.
- Retrieve the run repeatedly until it reaches `completed` or `failed`, and fail the
  application unless it completes successfully.
- Traverse every page of output items. Print each output item ID and status, then
  print every returned metric's name, score, and passed value.
- Delete the created evaluation after the run.
- Include the project manifest and concise restore, build, and run commands.

Use asynchronous SDK operations throughout.

## Evaluation Criteria

### Evaluation workflow

- Gets `EvaluationClient` from
  `projectClient.ProjectOpenAIClient.GetEvaluationClient()` rather than inventing an
  AI Projects evaluations operation group.
- Creates the evaluation with `CreateEvaluationAsync`, a custom schema, and an
  `azure_ai_evaluator` criterion whose evaluator name is `builtin.f1_score` and
  whose data mapping references the required item fields.
- Calls `CreateEvaluationRunAsync` with the created evaluation ID and an inline
  `jsonl`/`file_content` source containing the exact item.
- Polls through `GetEvaluationRunAsync` until an explicit terminal status and
  rejects a failed run.
- Uses `GetEvaluationRunOutputItemsAsync` pagination and traverses each output
  item's returned results to print metric name, score, and passed state.
- Deletes the evaluation with its exact ID.

### Scenario-specific anti-patterns

- Does not invent `projectClient.Evaluations` or assume the create-run response is
  terminal.
- Does not hardcode a metric name, score, or pass result instead of traversing
  returned result objects.
- Does not omit output-item pagination or print only the run's aggregate counts.
- Does not substitute a chat completion call or locally calculate F1.

## Context

The reference application is in
`reference-apps/ai-projects/evaluation-run/dotnet`.
