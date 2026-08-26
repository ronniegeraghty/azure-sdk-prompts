import TextTranslationClient, {
  isUnexpected,
} from "@azure-rest/ai-translation-text";
import { DefaultAzureCredential } from "@azure/identity";

const [text, ...extra] = process.argv.slice(2);
if (!text || extra.length !== 0) {
  throw new Error("Usage: multilingual-translation <text>");
}

const endpoint = process.env["AZURE_TEXT_TRANSLATION_ENDPOINT"];
if (!endpoint) {
  throw new Error("AZURE_TEXT_TRANSLATION_ENDPOINT is required.");
}

const client = TextTranslationClient(
  endpoint,
  new DefaultAzureCredential(),
);
const response = await client.path("/translate").post({
  body: {
    inputs: [
      {
        text,
        targets: [{ language: "fr" }, { language: "ja" }],
      },
    ],
  },
});

if (isUnexpected(response)) {
  throw new Error(
    `${response.body.error.code}: ${response.body.error.message}`,
  );
}

const result = response.body.value[0];
if (!result?.detectedLanguage) {
  throw new Error("The service did not return a detected source language.");
}

console.log(
  `Detected source language: ${result.detectedLanguage.language}`,
);
for (const translation of result.translations) {
  console.log(`Translation (${translation.language}): ${translation.text}`);
}
