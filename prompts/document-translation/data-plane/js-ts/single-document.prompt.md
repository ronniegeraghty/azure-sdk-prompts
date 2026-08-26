---
id: document-translation-dp-js-ts-single-document
properties:
  service: document-translation
  plane: data-plane
  language: js-ts
  category: translation
  difficulty: intermediate
  description: "Can an application translate one local document with the synchronous document API and preserve the returned bytes?"
  sdk_package: "@azure/ai-translation-document"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - single-document
  - binary
---

# Single-document translation (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application using
`@azure/ai-translation-document` 1.0.0. Target Node.js 22. Accept exactly three
command-line arguments: local document path, target-language code, and output
path. Read `AZURE_DOCUMENT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Open the local document as binary data and infer its media type from the file
  extension, rejecting unsupported extensions.
- Use the current single-document translation client and API, supplying the
  original filename, media type, file content, and target language.
- Persist the exact returned document bytes to the requested output path without
  decoding, parsing, or transforming them.
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Single-document workflow

- Imports and constructs `SingleDocumentTranslationClient` from the released
  non-REST `@azure/ai-translation-document` package.
- Calls `translate(targetLanguage, { document: { contents, filename,
  contentType } })`, where the media type is derived from the extension.
- Requires `readableStreamBody` and pipes that response stream directly to the
  output file with backpressure-aware stream handling.

### Scenario-specific anti-patterns

- Does not use the obsolete/planned
  `@azure-rest/ai-translation-document` package or the batch client.
- Does not submit base64 JSON or omit multipart filename/content-type metadata.
- Does not call `text()`, parse JSON, concatenate string chunks, or otherwise
  transform the binary response before writing.
- Does not write a known or hardcoded translated output.

## Context

The reference application is in
`reference-apps/document-translation/single-document/js-ts`.
