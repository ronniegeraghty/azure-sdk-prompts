---
id: text-translation-dp-java-multilingual-translation
properties:
  service: text-translation
  plane: data-plane
  language: java
  category: translation
  difficulty: intermediate
  description: "Can an application translate one supplied text into French and Japanese in one request and report the detected source language?"
  sdk_package: com.azure:azure-ai-translation-text
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-translation-text-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - translation
  - language-detection
---

# Multilingual text translation (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-translation-text` 2.0.2. Accept one command-line argument
containing the text to translate and read `AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Send the supplied text in one SDK translation request whose targets are French
  (`fr`) and Japanese (`ja`).
- Omit the source language so the service detects it.
- Print the detected source language returned by the service.
- Traverse the returned translations and print each returned target-language code
  and translated text.
- Include the project manifest and concise restore, build, and run commands.

Use the synchronous SDK client.

## Evaluation Criteria

### Multilingual translation workflow

- Creates one `TranslateInputItem` containing the supplied text and two
  `TranslationTarget` values for `fr` and `ja`, without setting its source
  language.
- Calls `TextTranslationClient.translate` exactly once with the input item.
- Reads `TranslatedTextItem.getDetectedLanguage().getLanguage()` and traverses
  every `TranslationText` returned by `getTranslations()`.
- Prints each translation's returned `getLanguage()` and `getText()` values.

### Scenario-specific anti-patterns

- Does not invent a translations operation group or use an obsolete package/API.
- Does not make separate requests for French and Japanese or request only one
  target.
- Does not hardcode the detected language, target labels, or translated output.

## Context

The reference application is in
`reference-apps/text-translation/multilingual-translation/java`.
