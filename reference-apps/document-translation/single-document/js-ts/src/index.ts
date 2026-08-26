import { SingleDocumentTranslationClient } from
  "@azure/ai-translation-document";
import { DefaultAzureCredential } from "@azure/identity";
import { createReadStream, createWriteStream } from "node:fs";
import { basename, extname } from "node:path";
import { pipeline } from "node:stream/promises";

const [inputPath, targetLanguage, outputPath, ...extra] =
  process.argv.slice(2);
if (
  !inputPath ||
  !targetLanguage ||
  !outputPath ||
  extra.length !== 0
) {
  throw new Error(
    "Usage: single-document "
      + "<input-path> <target-language> <output-path>",
  );
}

const endpoint = process.env["AZURE_DOCUMENT_TRANSLATION_ENDPOINT"];
if (!endpoint) {
  throw new Error("AZURE_DOCUMENT_TRANSLATION_ENDPOINT is required.");
}

const client = new SingleDocumentTranslationClient(
  endpoint,
  new DefaultAzureCredential(),
);
const response = await client.translate(targetLanguage, {
  document: {
    contents: createReadStream(inputPath),
    filename: basename(inputPath),
    contentType: getMediaType(inputPath),
  },
});
if (!response.readableStreamBody) {
  throw new Error("The service returned no document response stream.");
}

await pipeline(
  response.readableStreamBody,
  createWriteStream(outputPath),
);

function getMediaType(path: string): string {
  switch (extname(path).toLowerCase()) {
    case ".txt":
      return "text/plain";
    case ".html":
    case ".htm":
      return "text/html";
    case ".docx":
      return "application/vnd.openxmlformats-officedocument"
        + ".wordprocessingml.document";
    case ".pdf":
      return "application/pdf";
    default:
      throw new Error(
        `Unsupported document extension: ${extname(path)}`,
      );
  }
}
