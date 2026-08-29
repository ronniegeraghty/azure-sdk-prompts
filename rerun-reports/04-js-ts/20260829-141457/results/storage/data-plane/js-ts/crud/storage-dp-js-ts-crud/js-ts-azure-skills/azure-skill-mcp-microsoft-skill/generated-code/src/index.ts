import { DefaultAzureCredential } from "@azure/identity";
import { BlobServiceClient, RestError } from "@azure/storage-blob";

const containerName = "my-container";
const blobName = "greeting.txt";
const blobContent = "Hello Azure!";

async function streamToString(
  stream: NodeJS.ReadableStream,
): Promise<string> {
  const chunks: Buffer[] = [];

  for await (const chunk of stream) {
    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
  }

  return Buffer.concat(chunks).toString("utf8");
}

async function main(): Promise<void> {
  const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;
  if (!accountName) {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_NAME must contain the Azure Storage account name.",
    );
  }

  const credential = new DefaultAzureCredential();
  const blobServiceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );

  const containerClient =
    blobServiceClient.getContainerClient(containerName);
  await containerClient.createIfNotExists();
  console.log(`Container ready: ${containerName}`);

  const blockBlobClient = containerClient.getBlockBlobClient(blobName);
  await blockBlobClient.upload(
    blobContent,
    Buffer.byteLength(blobContent, "utf8"),
    {
      blobHTTPHeaders: {
        blobContentType: "text/plain; charset=utf-8",
      },
    },
  );
  console.log(`Uploaded: ${blobName}`);

  console.log("Blobs:");
  for await (const blob of containerClient.listBlobsFlat()) {
    console.log(`- ${blob.name}`);
  }

  const downloadResponse = await blockBlobClient.download();
  if (!downloadResponse.readableStreamBody) {
    throw new Error(`The download for "${blobName}" returned no content stream.`);
  }

  const downloadedContent = await streamToString(
    downloadResponse.readableStreamBody,
  );
  console.log(`Downloaded content: ${downloadedContent}`);

  await blockBlobClient.delete();
  console.log(`Deleted blob: ${blobName}`);

  await containerClient.delete();
  console.log(`Deleted container: ${containerName}`);
}

void main().catch((error: unknown) => {
  if (error instanceof RestError) {
    console.error("Azure Blob Storage request failed.", {
      message: error.message,
      statusCode: error.statusCode,
      code: error.code,
      requestId: error.request?.headers.get("x-ms-request-id"),
    });
  } else if (error instanceof Error) {
    console.error(`Application error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
});
