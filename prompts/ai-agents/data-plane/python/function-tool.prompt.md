---
id: ai-agents-dp-python-function-tool
properties:
  service: ai-agents
  plane: data-plane
  language: python
  category: agents
  difficulty: advanced
  description: "Can an agent implement the Azure AI Agents required-action workflow for a local function tool?"
  sdk_package: azure-ai-agents
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-agents
  - function-calling
  - required-action
  - tool-outputs
---

# Azure AI agent function tool (Python)

## Prompt

Create a complete, runnable Python console application that uses `azure-ai-agents`
to answer a weather question through a local function tool.

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
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client throughout.

## Evaluation Criteria

### Function-tool workflow

- Defines `get_weather` and exposes it through `FunctionTool` with the required
  schema.
- Passes the function definition when calling `create_agent`.
- Creates the thread, exact user message, and run with the created IDs.
- Detects `requires_action` and `SubmitToolOutputsAction`.
- Processes every `RequiredFunctionToolCall`, checks its name, and parses `location`
  and `unit` from its arguments.
- Produces deterministic JSON by invoking local code rather than returning a fixed
  assistant response.
- Creates `ToolOutput` values with each originating tool-call ID and submits them
  through `runs.submit_tool_outputs`.
- Resumes polling after submission, requires completed status, and retrieves
  ascending assistant `text_messages`.
- Deletes the created thread and agent.

### Scenario-specific anti-patterns

- Does not call the weather function before the service requests it.
- Does not discard tool-call IDs or submit one output for unrelated calls.
- Does not stop at `requires_action` or print the local function result as though it
  were the final assistant response.

## Context

The reference application is in `reference-apps/ai-agents/function-tool/python`.

