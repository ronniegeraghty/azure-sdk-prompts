---
id: document-translation-dp-java-single-document
properties:
  service: document-translation
  plane: data-plane
  language: java
  category: translation
  difficulty: intermediate
  description: "Can an application translate one local document with the synchronous document API and preserve the returned bytes?"
  sdk_package: com.azure:azure-ai-translation-document
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - single-document
  - binary
---

# Single-document translation (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-translation-document` 2.0.1. Accept exactly three
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
- Include the project manifest and concise restore, build, and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Single-document workflow

- Constructs `SingleDocumentTranslationClient`, not
  `DocumentTranslationClient`.
- Creates `DocumentFileDetails` from `BinaryData` with the input filename and a
  media type derived from the extension, then creates
  `DocumentTranslateContent`.
- Calls `translate(targetLanguage, content)`.
- Writes the returned `BinaryData.toBytes()` directly to the output file.

### Scenario-specific anti-patterns

- Does not use the batch/container client or invent a single-document method on
  it.
- Does not submit base64 JSON or omit multipart filename/content-type metadata.
- Does not call `BinaryData.toString()` or otherwise convert the binary response
  before writing.
- Does not write a known or hardcoded translated output.

## Context

The reference application is in
`reference-apps/document-translation/single-document/java`.
