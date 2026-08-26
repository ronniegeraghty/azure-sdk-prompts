---
id: ai-agents-dp-python-basic-agent-lifecycle
properties:
  service: ai-agents
  plane: data-plane
  language: python
  category: agents
  difficulty: intermediate
  description: "Can an agent build a complete Azure AI Agents application that creates an agent conversation, runs it to completion, and reads the assistant response?"
  sdk_package: azure-ai-agents
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-agents
  - threads
  - messages
  - runs
  - polling
---

# Basic Azure AI agent lifecycle (Python)

## Prompt

Create a complete, runnable Python console application that uses the
`azure-ai-agents` package to run a basic Azure AI Agent conversation.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read the Azure AI project endpoint from `PROJECT_ENDPOINT` and the model deployment
  name from `MODEL_DEPLOYMENT_NAME`.
- Create the Azure AI Agents client using the project endpoint and
  `DefaultAzureCredential`.
- Create an agent named `hyoka-basic-agent` with the instructions
  `Answer the user's question clearly and concisely.`
- Create a thread and add this exact user message:
  `What is the capital of France?`
- Create a run for that thread and agent, then use the SDK to refresh the run until it
  reaches a terminal status.
- After a successful run, list the thread messages in chronological order and print
  the text content of every assistant message.
- Delete the thread and agent after the conversation finishes.
- Include the project manifest and concise commands for restoring, building, and
  running the application.

Use the synchronous Azure AI Agents client throughout the agent workflow.

## Evaluation Criteria

### Azure AI Agents workflow

- Creates `AgentsClient` from `azure.ai.agents` with the project endpoint.
- Calls `agents_client.create_agent` with the deployment from
  `MODEL_DEPLOYMENT_NAME`, the required name, and the required instructions.
- Creates a thread with `agents_client.threads.create`.
- Adds the exact user message with `agents_client.messages.create`, role `user`, and
  the created thread ID.
- Creates the run with the created thread ID and agent ID, then polls with
  `agents_client.runs.get` or the SDK's combined create-and-process operation until a
  terminal status is reached.
- Retrieves messages only after successful completion, requests
  `ListSortOrder.ASCENDING`, and extracts text values from assistant
  `text_messages`.
- Deletes both the created thread and the created agent through `AgentsClient`.

### Scenario-specific anti-patterns

- Does not print a hardcoded answer instead of retrieving agent messages.
- Does not treat the initially created run as already completed.
- Does not substitute an Azure OpenAI chat client or another non-agents API for the
  Azure AI Agents thread and run workflow.

## Context

This prompt is one member of the four-language Azure AI Agents parity set. Its
reference application is in
`reference-apps/ai-agents/basic-agent-lifecycle/python`.
