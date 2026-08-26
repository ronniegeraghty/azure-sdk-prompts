---
id: document-translation-dp-dotnet-single-document
properties:
  service: document-translation
  plane: data-plane
  language: dotnet
  category: translation
  difficulty: intermediate
  description: "Can an application translate one local document with the synchronous document API and preserve the returned bytes?"
  sdk_package: Azure.AI.Translation.Document
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/ai.translation.document-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - document-translation
  - single-document
  - binary
---

# Single-document translation (.NET)

## Prompt

Create a complete, runnable .NET console application using
`Azure.AI.Translation.Document` 3.0.0. Accept exactly three command-line
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
- Include the project manifest and concise restore, build, and run commands.

Use asynchronous SDK operations.

## Evaluation Criteria

### Single-document workflow

- Constructs `SingleDocumentTranslationClient`, not
  `DocumentTranslationClient`.
- Wraps the input stream in `MultipartFormFileData` with the input filename and a
  media type derived from the extension, then creates `DocumentTranslateContent`.
- Calls `TranslateAsync(targetLanguage, content)`.
- Writes `response.Value` directly as bytes to the output stream.

### Scenario-specific anti-patterns

- Does not use the batch/container client or invent a single-document method on
  it.
- Does not submit base64 JSON or omit multipart filename/content-type metadata.
- Does not convert the binary response to text, JSON, or an object before writing.
- Does not write a known or hardcoded translated output.

## Context

The reference application is in
`reference-apps/document-translation/single-document/dotnet`.
