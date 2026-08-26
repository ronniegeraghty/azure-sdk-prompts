---
id: ai-projects-dp-java-evaluation-run
properties:
  service: ai-projects
  plane: data-plane
  language: java
  category: projects
  difficulty: advanced
  description: "Can an application use the project OpenAI Evals client to run a built-in evaluator and traverse metric results?"
  sdk_package: com.azure:azure-ai-projects
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - evaluations
  - polling
  - metrics
---

# Azure AI project evaluation run (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-projects` 2.4.0 to run an OpenAI-compatible evaluation in a
Microsoft Foundry project.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `FOUNDRY_PROJECT_ENDPOINT`.
- Obtain the project OpenAI client from `AIProjectClientBuilder`.
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

Use the synchronous SDK clients throughout.

## Evaluation Criteria

### Evaluation workflow

- Calls `AIProjectClientBuilder.buildOpenAIClient()` rather than inventing an AI
  Projects evaluations client.
- Builds `EvalCreateParams` with a custom schema and adapts
  `TestingCriterionAzureAIEvaluator` for `builtin.f1_score` through
  `EvaluationsHelper.toTestingCriterion`.
- Creates the run through `openAIClient.evals().runs().create` with the created
  evaluation ID and a typed `CreateEvalJsonlRunDataSource` file-content source
  containing the exact item.
- Polls through `runs().retrieve` until an explicit terminal status and rejects a
  failed run.
- Uses the output-items `autoPager`, iterates every
  `OutputItemListResponse.Result`, and prints its name, score, and passed state.
- Deletes the evaluation with its exact ID.

### Scenario-specific anti-patterns

- Does not invent `projectClient.evaluations()` or assume the create-run response is
  terminal.
- Does not hardcode a metric name, score, or pass result instead of traversing
  returned result objects.
- Does not omit output-item paging or print only the run's aggregate counts.
- Does not substitute a chat completion call or locally calculate F1.

## Context

The reference application is in
`reference-apps/ai-projects/evaluation-run/java`.
