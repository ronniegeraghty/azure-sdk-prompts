---
id: document-translation-dp-java-batch-container
properties:
  service: document-translation
  plane: data-plane
  language: java
  category: batch
  difficulty: advanced
  description: "Can an application start a container translation, wait for completion, and enumerate every document status and failure?"
  sdk_package: com.azure:azure-ai-translation-document
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - batch
  - polling
  - pagination
---

# Batch container document translation (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-translation-document` 2.0.1. Accept exactly three
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
- Include the project manifest and concise restore, build, and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Batch container workflow

- Constructs `DocumentTranslationClient`, not
  `SingleDocumentTranslationClient`.
- Creates `TranslationSource`, `TranslationTarget`, and
  `DocumentTranslationInput` from the supplied arguments, then calls
  `beginTranslation`.
- Calls the returned `SyncPoller.waitForCompletion()` and uses the resulting
  `TranslationStatusResult` ID and status.
- Iterates the `PagedIterable<DocumentStatusResult>` returned by
  `listDocumentStatuses(translationId)`, traversing all pages.
- Prints each document's `getId()` and `getStatus()`, and prints
  `getError().getCode()` and `getError().getMessage()` whenever error details
  are present.

### Scenario-specific anti-patterns

- Does not use the single-document client or multipart upload workflow.
- Does not omit long-running-operation polling or treat the start response as
  complete.
- Does not print only overall counts, select one document, or consume one page.
- Does not omit returned document failure code or message.

## Context

The reference application is in
`reference-apps/document-translation/batch-container/java`.
