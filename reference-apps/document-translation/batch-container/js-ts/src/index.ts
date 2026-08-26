import { DocumentTranslationClient } from
  "@azure/ai-translation-document";
import { DefaultAzureCredential } from "@azure/identity";

const [sourceUrl, targetUrl, language, ...extra] = process.argv.slice(2);
if (!sourceUrl || !targetUrl || !language || extra.length !== 0) {
  throw new Error(
    "Usage: batch-container "
      + "<source-container-uri> <target-container-uri> <target-language>",
  );
}

const endpoint = process.env["AZURE_DOCUMENT_TRANSLATION_ENDPOINT"];
if (!endpoint) {
  throw new Error("AZURE_DOCUMENT_TRANSLATION_ENDPOINT is required.");
}

const client = new DocumentTranslationClient(
  endpoint,
  new DefaultAzureCredential(),
);
const poller = client.startTranslation({
  inputs: [
    {
      source: { sourceUrl },
      targets: [{ targetUrl, language }],
    },
  ],
});
const result = await poller.pollUntilDone();
console.log(`Overall status: ${result.status}`);

for await (const document of client.listDocumentStatuses(result.id)) {
  console.log(`Document: id=${document.id} status=${document.status}`);
  if (document.error) {
    console.log(
      `Failure: code=${document.error.code} `
        + `message=${document.error.message}`,
    );
  }
}
