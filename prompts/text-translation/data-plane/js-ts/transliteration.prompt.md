---
id: text-translation-dp-js-ts-transliteration
properties:
  service: text-translation
  plane: data-plane
  language: js-ts
  category: translation
  difficulty: intermediate
  description: "Can an application transliterate supplied text between supplied scripts and report the returned script metadata?"
  sdk_package: "@azure-rest/ai-translation-text"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-translation-text-rest-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - transliteration
  - scripts
---

# Text transliteration (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application using
`@azure-rest/ai-translation-text` 2.0.0. Target Node.js 22. Accept exactly four
command-line arguments: language code, source-script code, target-script code,
and text. Read `AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Invoke the SDK's transliteration operation with the supplied language,
  source-script, target-script, and text values in that order.
- Print the transliterated text returned by the service.
- Print the actual script metadata returned with that text.
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Transliteration workflow

- Calls `client.path("/transliterate").post` with the supplied text in
  `body.inputs` and the supplied `language`, `fromScript`, and `toScript` in
  `queryParameters`, without swapping the script codes.
- Checks `isUnexpected` and reads an item from `response.body.value`.
- Prints both the returned item's `text` and `script`.

### Scenario-specific anti-patterns

- Does not call `"/translate"`, invent a convenience `transliterate` method, use
  the retired v1 shape, or use an obsolete package.
- Does not swap `fromScript` and `toScript`.
- Does not echo the input text or requested target-script code as if they were
  returned output.

## Context

The reference application is in
`reference-apps/text-translation/transliteration/js-ts`.
