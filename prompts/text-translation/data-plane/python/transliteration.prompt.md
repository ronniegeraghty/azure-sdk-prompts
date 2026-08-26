---
id: text-translation-dp-python-transliteration
properties:
  service: text-translation
  plane: data-plane
  language: python
  category: translation
  difficulty: intermediate
  description: "Can an application transliterate supplied text between supplied scripts and report the returned script metadata?"
  sdk_package: azure-ai-translation-text
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-translation-text-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - transliteration
  - scripts
---

# Text transliteration (Python)

## Prompt

Create a complete, runnable Python console application using
`azure-ai-translation-text` 2.0.0. Accept exactly four command-line arguments:
language code, source-script code, target-script code, and text. Read
`AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Invoke the SDK's transliteration operation with the supplied language,
  source-script, target-script, and text values in that order.
- Print the transliterated text returned by the service.
- Print the actual script metadata returned with that text.
- Include the dependency manifest and concise install and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Transliteration workflow

- Calls `TextTranslationClient.transliterate` with the supplied text in `body`
  and supplied `language`, `from_script`, and `to_script` keyword arguments,
  without swapping the script codes.
- Reads a returned transliteration result rather than a translation result.
- Prints both the returned item's `text` and `script`.

### Scenario-specific anti-patterns

- Does not call `translate`, invent a transliteration operation group, or use an
  obsolete package/API.
- Does not swap `from_script` and `to_script`.
- Does not echo the input text or requested target-script code as if they were
  returned output.

## Context

The reference application is in
`reference-apps/text-translation/transliteration/python`.
