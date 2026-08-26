---
id: document-translation-dp-python-single-document
properties:
  service: document-translation
  plane: data-plane
  language: python
  category: translation
  difficulty: intermediate
  description: "Can an application translate one local document with the synchronous document API and preserve the returned bytes?"
  sdk_package: azure-ai-translation-document
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-translation-document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - single-document
  - binary
---

# Single-document translation (Python)

## Prompt

Create a complete, runnable Python console application using
`azure-ai-translation-document` 2.0.0. Accept exactly three command-line
arguments: local document path, target-language code, and output path. Read
`AZURE_DOCUMENT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Open the local document as binary data and infer its media type from the file
  extension, rejecting unsupported extensions.
- Use the current single-document translation client and API, supplying the
  original filename, media type, file content, and target language.
- Persist the exact returned document bytes to the requested output path without
  decoding, parsing, or transforming them.
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Single-document workflow

- Constructs `SingleDocumentTranslationClient`, not
  `DocumentTranslationClient`.
- Creates `DocumentTranslateContent` with a document tuple containing input
  filename, binary file object or bytes, and a media type derived from the
  extension.
- Calls `translate(body=content, target_language=targetLanguage)`.
- Writes the returned bytes directly to a file opened in binary mode.

### Scenario-specific anti-patterns

- Does not use the batch/container client or invent a single-document method on
  it.
- Does not submit base64 JSON or omit multipart filename/content-type metadata.
- Does not decode, JSON-parse, or stringify the binary response before writing.
- Does not write a known or hardcoded translated output.

## Context

The reference application is in
`reference-apps/document-translation/single-document/python`.
