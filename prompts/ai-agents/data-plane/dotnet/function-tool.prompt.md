---
id: ai-agents-dp-dotnet-function-tool
properties:
  service: ai-agents
  plane: data-plane
  language: dotnet
  category: agents
  difficulty: advanced
  description: "Can an agent implement the Azure AI Agents required-action workflow for a local function tool?"
  sdk_package: Azure.AI.Agents.Persistent
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/ai.agents.persistent-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-agents
  - function-calling
  - required-action
  - tool-outputs
---

# Azure AI agent function tool (.NET)

## Prompt

Create a complete, runnable .NET console application that uses
`Azure.AI.Agents.Persistent` to answer a weather question through a local function
tool.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `PROJECT_ENDPOINT` and `MODEL_DEPLOYMENT_NAME`.
- Define a function tool named `get_weather` with required string parameters
  `location` and `unit`; `unit` must allow only `c` or `f`.
- Implement the local function deterministically: for Seattle, return temperature
  `21` for `c` or `70` for `f`, including the location and unit in the JSON result.
- Create an agent named `hyoka-weather-agent` whose instructions require using the
  function for weather questions.
- Create a thread with the exact user message
  `What is the weather in Seattle in celsius?`, then create a run.
- Poll the run. When it requires action, iterate its function tool calls, decode the
  JSON arguments, invoke `get_weather`, create an output correlated with each tool
  call ID, and submit all outputs through the SDK.
- Continue polling until terminal status. After successful completion, list messages
  in chronological order and print assistant text.
- Delete the thread and agent.
- Include the project manifest and concise restore, build, and run commands.

Use asynchronous SDK operations throughout.

## Evaluation Criteria

### Function-tool workflow

- Defines `FunctionToolDefinition` for `get_weather` with the required JSON schema.
- Passes the tool definition when creating the agent.
- Creates the thread, exact user message, and run with the created IDs.
- Detects `RunStatus.RequiresAction` and `SubmitToolOutputsAction`.
- Processes every `RequiredFunctionToolCall`, checks the function name, and parses
  `location` and `unit` from its arguments.
- Produces deterministic JSON by invoking local code rather than returning a fixed
  assistant response.
- Creates `ToolOutput` values correlated to the originating tool calls and submits
  them with `SubmitToolOutputsToRunAsync`.
- Resumes polling after submission, requires completed status, and retrieves
  ascending assistant messages through `MessageTextContent`.
- Deletes the created thread and agent.

### Scenario-specific anti-patterns

- Does not call the weather function before the service requests it.
- Does not discard tool-call IDs or submit one output for unrelated calls.
- Does not stop at `RequiresAction` or print the local function result as though it
  were the final assistant response.

## Context

The reference application is in `reference-apps/ai-agents/function-tool/dotnet`.

