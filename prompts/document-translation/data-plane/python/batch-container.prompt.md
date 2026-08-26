---
id: document-translation-dp-python-batch-container
properties:
  service: document-translation
  plane: data-plane
  language: python
  category: batch
  difficulty: advanced
  description: "Can an application start a container translation, wait for completion, and enumerate every document status and failure?"
  sdk_package: azure-ai-translation-document
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - batch
  - polling
  - pagination
---

# Batch container document translation (Python)

## Prompt

Create a complete, runnable Python console application using
`azure-ai-translation-document` 2.0.0. Accept exactly three command-line
arguments: source-container URI, target-container URI, and target-language code.
Read `AZURE_DOCUMENT_TRANSLATION_ENDPOINT`.

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
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Batch container workflow

- Constructs `DocumentTranslationClient`, not
  `SingleDocumentTranslationClient`.
- Calls `begin_translation(sourceUrl, targetUrl, targetLanguage)` with all three
  supplied arguments.
- Calls `poller.result()` before reporting `poller.status()`.
- Traverses the pageable result from
  `client.list_document_statuses(poller.id)` so every page and document is
  consumed.
- Prints each document's `id` and `status`, and prints `error.code` and
  `error.message` whenever error details are present.

### Scenario-specific anti-patterns

- Does not use the single-document client or multipart upload workflow.
- Does not omit poller completion or treat the begin-operation response as final.
- Does not print only overall counts, select one document, or consume one page.
- Does not omit returned document failure code or message.

## Context

The reference application is in
`reference-apps/document-translation/batch-container/python`.
