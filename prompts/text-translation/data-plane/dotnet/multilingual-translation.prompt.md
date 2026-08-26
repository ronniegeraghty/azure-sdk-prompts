---
id: text-translation-dp-dotnet-multilingual-translation
properties:
  service: text-translation
  plane: data-plane
  language: dotnet
  category: translation
  difficulty: intermediate
  description: "Can an application translate one supplied text into French and Japanese in one request and report the detected source language?"
  sdk_package: Azure.AI.Translation.Text
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/ai.translation.text-readme
  created: "2026-08-27"
  author: weidongxu-microsoft
tags:
  - foundry
  - text-translation
  - translation
  - language-detection
---

# Multilingual text translation (.NET)

## Prompt

Create a complete, runnable .NET console application using
`Azure.AI.Translation.Text` 2.0.0. Accept one command-line argument containing
the text to translate and read `AZURE_TEXT_TRANSLATION_ENDPOINT`.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Send the supplied text in one SDK translation request whose targets are French
  (`fr`) and Japanese (`ja`).
- Omit the source language so the service detects it.
- Print the detected source language returned by the service.
- Traverse the returned translations and print each returned target-language code
  and translated text.
- Include the project manifest and concise restore, build, and run commands.

Use asynchronous SDK operations.

## Evaluation Criteria

### Multilingual translation workflow

- Creates one `TranslateInputItem` containing the supplied text and two
  `TranslationTarget` values for `fr` and `ja`, without setting its source
  language.
- Calls `TextTranslationClient.TranslateAsync` exactly once for the input item.
- Reads `TranslatedTextItem.DetectedLanguage.Language` and traverses every
  `TranslationText` in `Translations`.
- Prints each translation's returned `Language` and `Text` values.

### Scenario-specific anti-patterns

- Does not invent a translations operation group or use an obsolete package/API.
- Does not make separate requests for French and Japanese or request only one
  target.
- Does not hardcode the detected language, target labels, or translated output.

## Context

The reference application is in
`reference-apps/text-translation/multilingual-translation/dotnet`.
