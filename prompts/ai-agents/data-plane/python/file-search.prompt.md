---
id: ai-agents-dp-python-file-search
properties:
  service: ai-agents
  plane: data-plane
  language: python
  category: agents
  difficulty: advanced
  description: "Can an agent ground an Azure AI Agents response in an uploaded document through file search?"
  sdk_package: azure-ai-agents
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-agents
  - file-search
  - vector-store
  - grounding
---

# Azure AI agent file search (Python)

## Prompt

Create a complete, runnable Python console application that uses `azure-ai-agents`
to answer a question grounded in an uploaded document.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `PROJECT_ENDPOINT` and `MODEL_DEPLOYMENT_NAME`.
- Create a text document containing this exact fact:
  `The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.`
- Upload the document for agents, create a vector store containing it, and wait until
  the vector store is ready. Fail if indexing doesn't complete successfully.
- Create an agent named `hyoka-trail-guide-agent` with a file-search tool and tool
  resources that reference the created vector store.
- Create a thread with the exact user message
  `According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?`
- Create and process a run until terminal status, and require successful completion.
- List messages in chronological order and print assistant text.
- Delete the thread, agent, vector store, and uploaded file in dependency order.
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client throughout.

## Evaluation Criteria

### File-search workflow

- Uploads the document with `files.upload_and_poll` and `FilePurpose.AGENTS`.
- Creates the vector store with the uploaded file ID through `create_and_poll` and
  requires completed status.
- Creates `FileSearchTool` with the vector-store ID and passes both its definitions
  and resources when creating the agent.
- Creates the thread, exact grounded question, and run with the created IDs.
- Processes the run through `runs.create_and_process`, requires completed status,
  and retrieves ascending assistant `text_messages`.
- Deletes the thread before the agent, then deletes the vector store before the
  uploaded file.

### Scenario-specific anti-patterns

- Does not put a local file path or uploaded file ID directly in the user message as
  a substitute for file-search tool resources.
- Does not start the run before vector-store indexing completes.
- Does not print the known document fact directly as though it were the agent's
  grounded response.

## Context

The reference application is in `reference-apps/ai-agents/file-search/python`.
