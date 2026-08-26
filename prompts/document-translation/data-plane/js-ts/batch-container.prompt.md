---
id: document-translation-dp-js-ts-batch-container
properties:
  service: document-translation
  plane: data-plane
  language: js-ts
  category: batch
  difficulty: advanced
  description: "Can an application start a container translation, wait for completion, and enumerate every document status and failure?"
  sdk_package: "@azure/ai-translation-document"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - batch
  - polling
  - pagination
---

# Batch container document translation (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application using
`@azure/ai-translation-document` 1.0.0. Target Node.js 22. Accept exactly three
command-line arguments: source-container URI, target-container URI, and
target-language code. Read `AZURE_DOCUMENT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Start one batch document translation from the supplied source container to the
  supplied target container and language.
- Use the SDK long-running operation and wait for it to finish.
- Print the returned overall translation status.
- Page through every document status in that operation and print each document
  ID and status.
- For every document whose returned failure information is present, print its
  returned error code and message.
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Batch container workflow

- Imports and constructs `DocumentTranslationClient` from the released non-REST
  `@azure/ai-translation-document` package, not
  `SingleDocumentTranslationClient`.
- Calls `startTranslation` with one input whose `source.sourceUrl`,
  `targets[0].targetUrl`, and `targets[0].language` come from the supplied
  arguments.
- Calls `pollUntilDone()` on the returned poller and prints its returned status.
- Uses `for await...of` over
  `client.listDocumentStatuses(result.id)`, traversing every page.
- Prints each document's `id` and `status`, and prints `error.code` and
  `error.message` whenever error details are present.

### Scenario-specific anti-patterns

- Does not use the obsolete/planned
  `@azure-rest/ai-translation-document` package, the single-document client, or
  multipart upload workflow.
- Does not omit poller completion or treat `startTranslation` as a promise for
  the final result.
- Does not print only overall counts, select one document, or consume one page.
- Does not omit returned document failure code or message.

## Context

The reference application is in
`reference-apps/document-translation/batch-container/js-ts`.
