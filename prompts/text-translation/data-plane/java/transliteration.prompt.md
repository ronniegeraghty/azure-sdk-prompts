---
id: text-translation-dp-java-transliteration
properties:
  service: text-translation
  plane: data-plane
  language: java
  category: translation
  difficulty: intermediate
  description: "Can an application transliterate supplied text between supplied scripts and report the returned script metadata?"
  sdk_package: com.azure:azure-ai-translation-text
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-translation-text-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - transliteration
  - scripts
---

# Text transliteration (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-translation-text` 2.0.2. Accept exactly four command-line
arguments: language code, source-script code, target-script code, and text. Read
`AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Invoke the SDK's transliteration operation with the supplied language,
  source-script, target-script, and text values in that order.
- Print the transliterated text returned by the service.
- Print the actual script metadata returned with that text.
- Include the project manifest and concise restore, build, and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Transliteration workflow

- Passes the four supplied values to
  `TextTranslationClient.transliterate(language, fromScript, toScript, text)`
  without swapping the script codes.
- Reads a returned `TransliteratedText` rather than a translation result.
- Prints both `TransliteratedText.getText()` and
  `TransliteratedText.getScript()`.

### Scenario-specific anti-patterns

- Does not call `translate`, invent a transliteration operation group, or use an
  obsolete package/API.
- Does not swap `fromScript` and `toScript`.
- Does not echo the input text or requested target-script code as if they were
  returned output.

## Context

The reference application is in
`reference-apps/text-translation/transliteration/java`.
