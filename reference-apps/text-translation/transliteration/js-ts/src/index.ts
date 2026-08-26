import TextTranslationClient, {
  isUnexpected,
} from "@azure-rest/ai-translation-text";
import { DefaultAzureCredential } from "@azure/identity";

const [language, fromScript, toScript, text, ...extra] =
  process.argv.slice(2);
if (!language || !fromScript || !toScript || !text || extra.length !== 0) {
  throw new Error(
    "Usage: transliteration <language> <from-script> <to-script> <text>",
  );
}

const endpoint = process.env["AZURE_TEXT_TRANSLATION_ENDPOINT"];
if (!endpoint) {
  throw new Error("AZURE_TEXT_TRANSLATION_ENDPOINT is required.");
}

const client = TextTranslationClient(
  endpoint,
  new DefaultAzureCredential(),
);
const response = await client.path("/transliterate").post({
  body: { inputs: [{ text }] },
  queryParameters: { language, fromScript, toScript },
});

if (isUnexpected(response)) {
  throw new Error(
    `${response.body.error.code}: ${response.body.error.message}`,
  );
}

const result = response.body.value[0];
if (!result) {
  throw new Error("The service returned no transliteration.");
}

console.log(`Transliterated text: ${result.text}`);
console.log(`Returned script: ${result.script}`);
