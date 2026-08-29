import { RestError } from "@azure/core-rest-pipeline";
import { DefaultAzureCredential } from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const containerName = "my-container";
const blobName = "greeting.txt";
const blobContents = "Hello Azure!";

async function streamToString(stream: NodeJS.ReadableStream): Promise<string> {
  return await new Promise<string>((resolve, reject) => {
    const chunks: Buffer[] = [];

    stream.on("data", (chunk: Buffer | string) => {
      chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
    });
    stream.on("end", () => resolve(Buffer.concat(chunks).toString("utf8")));
    stream.on("error", reject);
  });
}

async function main(): Promise<void> {
  const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;

  if (!accountName) {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_NAME must be set to the Azure Storage account name.",
    );
  }

  const credential = new DefaultAzureCredential();
  const serviceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );
  const containerClient = serviceClient.getContainerClient(containerName);

  await containerClient.createIfNotExists();
  console.log(`Container ready: ${containerName}`);

  const blockBlobClient = containerClient.getBlockBlobClient(blobName);
  await blockBlobClient.upload(blobContents, Buffer.byteLength(blobContents));
  console.log(`Uploaded: ${blobName}`);

  console.log("Blobs:");
  for await (const blob of containerClient.listBlobsFlat()) {
    console.log(`- ${blob.name}`);
  }

  const downloadResponse = await blockBlobClient.download();
  if (!downloadResponse.readableStreamBody) {
    throw new Error(`The download response for "${blobName}" had no body.`);
  }

  const downloadedContents = await streamToString(
    downloadResponse.readableStreamBody,
  );
  console.log(`Downloaded content: ${downloadedContents}`);

  await blockBlobClient.delete();
  console.log(`Deleted blob: ${blobName}`);

  await containerClient.delete();
  console.log(`Deleted container: ${containerName}`);
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof RestError) {
    console.error(
      `Azure request failed (${error.statusCode ?? "unknown status"}, ${
        error.code ?? "unknown code"
      }): ${error.message}`,
    );
  } else if (error instanceof Error) {
    console.error(`Error: ${error.message}`);
  } else {
    console.error("An unknown error occurred:", error);
  }

  process.exitCode = 1;
}
