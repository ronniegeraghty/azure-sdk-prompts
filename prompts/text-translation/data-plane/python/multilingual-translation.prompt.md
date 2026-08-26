---
id: text-translation-dp-python-multilingual-translation
properties:
  service: text-translation
  plane: data-plane
  language: python
  category: translation
  difficulty: intermediate
  description: "Can an application translate one supplied text into French and Japanese in one request and report the detected source language?"
  sdk_package: azure-ai-translation-text
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-translation-text-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - translation
  - language-detection
---

# Multilingual text translation (Python)

## Prompt

Create a complete, runnable Python console application using
`azure-ai-translation-text` 2.0.0. Accept one command-line argument containing
the text to translate and read `AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Send the supplied text in one SDK translation request whose targets are French
  (`fr`) and Japanese (`ja`).
- Omit the source language so the service detects it.
- Print the detected source language returned by the service.
- Traverse the returned translations and print each returned target-language code
  and translated text.
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Multilingual translation workflow

- Calls `TextTranslationClient.translate` exactly once with `body=[text]` and
  `to_language=["fr", "ja"]`, without passing `from_language`.
- Reads the first returned item's `detected_language.language`.
- Traverses every item in the returned `translations` collection.
- Prints each translation's returned `language` and `text` values.

### Scenario-specific anti-patterns

- Does not assume that version 2.0.0 requires two calls or invent a translations
  operation group.
- Does not make separate requests for French and Japanese or request only one
  target.
- Does not hardcode the detected language, target labels, or translated output.

## Context

The reference application is in
`reference-apps/text-translation/multilingual-translation/python`.
