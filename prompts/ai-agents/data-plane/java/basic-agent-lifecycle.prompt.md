---
id: ai-agents-dp-java-basic-agent-lifecycle
properties:
  service: ai-agents
  plane: data-plane
  language: java
  category: agents
  difficulty: intermediate
  description: "Can an agent build a complete Azure AI Agents application that creates an agent conversation, runs it to completion, and reads the assistant response?"
  sdk_package: com.azure:azure-ai-agents-persistent
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-agents-persistent-readme
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

# Basic Azure AI agent lifecycle (Java)

## Prompt

Create a complete, runnable Java console application that uses the
`com.azure:azure-ai-agents-persistent` package to run a basic Azure AI Agent
conversation.

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

- Builds `PersistentAgentsClient` with `PersistentAgentsClientBuilder`, then obtains
  the administration, threads, messages, and runs subclients.
- Creates the agent with `CreateAgentOptions`, passing the deployment from
  `MODEL_DEPLOYMENT_NAME`, the required name, and the required instructions.
- Creates a thread with `ThreadsClient.createThread`.
- Adds the exact user message with `MessagesClient.createMessage`,
  `MessageRole.USER`, and the created thread ID.
- Creates the run with `CreateRunOptions` containing the created thread ID and agent
  ID.
- Polls by calling `RunsClient.getRun` with the thread ID and run ID until the run
  leaves its nonterminal statuses.
- Retrieves messages only after successful completion, iterates the paged SDK result
  with `ListSortOrder.ASCENDING`, and extracts `MessageTextContent` values from
  `MessageRole.AGENT` messages.
- Deletes both the created thread and the created agent through their SDK clients.

### Scenario-specific anti-patterns

- Does not print a hardcoded answer instead of retrieving agent messages.
- Does not treat the initially created run as already completed.
- Does not substitute an Azure OpenAI chat client or another non-agents API for the
  Azure AI Agents thread and run workflow.

## Context

This prompt is one member of the four-language Azure AI Agents parity set. Its
reference application is in
`reference-apps/ai-agents/basic-agent-lifecycle/java`.
