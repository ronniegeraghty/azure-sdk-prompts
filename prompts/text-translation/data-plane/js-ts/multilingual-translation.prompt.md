---
id: text-translation-dp-js-ts-multilingual-translation
properties:
  service: text-translation
  plane: data-plane
  language: js-ts
  category: translation
  difficulty: intermediate
  description: "Can an application translate one supplied text into French and Japanese in one request and report the detected source language?"
  sdk_package: "@azure-rest/ai-translation-text"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-translation-text-rest-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - translation
  - language-detection
---

# Multilingual text translation (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application using
`@azure-rest/ai-translation-text` 2.0.0. Target Node.js 22. Accept one
command-line argument containing the text to translate and read
`AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Send the supplied text in one SDK translation request whose targets are French
  (`fr`) and Japanese (`ja`).
- Omit the source language so the service detects it.
- Print the detected source language returned by the service.
- Traverse the returned translations and print each returned target-language code
  and translated text.
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Multilingual translation workflow

- Creates one v2 `TranslateInputItem` containing the supplied text and two target
  objects for `fr` and `ja`, without setting its `language`.
- Calls `client.path("/translate").post` exactly once with
  `body: { inputs: [input] }` and checks `isUnexpected`.
- Reads the returned item's `detectedLanguage.language` and traverses every item
  in its `translations` array.
- Prints each translation's returned `language` and `text` values.

### Scenario-specific anti-patterns

- Does not use the retired v1 query-parameter request shape, invent a convenience
  `translate` method, or use an obsolete package.
- Does not make separate requests for French and Japanese or request only one
  target.
- Does not hardcode the detected language, target labels, or translated output.

## Context

The reference application is in
`reference-apps/text-translation/multilingual-translation/js-ts`.
