import { DefaultAzureCredential } from "@azure/identity";
import { BlobServiceClient, RestError } from "@azure/storage-blob";

const containerName = "my-container";
const blobName = "greeting.txt";
const content = "Hello Azure!";

async function main(): Promise<void> {
  const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;
  if (!accountName) {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_NAME must be set to the Azure Storage account name.",
    );
  }

  const credential = new DefaultAzureCredential();
  const blobServiceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );

  const containerClient =
    blobServiceClient.getContainerClient(containerName);

  const createResult = await containerClient.createIfNotExists();
  console.log(
    createResult.succeeded
      ? `Created container: ${containerName}`
      : `Container already exists: ${containerName}`,
  );

  const blockBlobClient = containerClient.getBlockBlobClient(blobName);
  await blockBlobClient.upload(content, Buffer.byteLength(content), {
    blobHTTPHeaders: { blobContentType: "text/plain; charset=utf-8" },
  });
  console.log(`Uploaded blob: ${blobName}`);

  console.log("Blobs in container:");
  for await (const blob of containerClient.listBlobsFlat()) {
    console.log(`- ${blob.name}`);
  }

  const downloadedContent = (
    await blockBlobClient.downloadToBuffer()
  ).toString("utf8");
  console.log(`Downloaded content: ${downloadedContent}`);

  await blockBlobClient.delete();
  console.log(`Deleted blob: ${blobName}`);

  await containerClient.delete();
  console.log(`Deleted container: ${containerName}`);
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof RestError) {
    console.error("Azure Blob Storage request failed.", {
      statusCode: error.statusCode,
      code: error.code,
      message: error.message,
      requestId: error.request?.headers.get("x-ms-request-id"),
    });
  } else if (error instanceof Error) {
    console.error(`Unexpected error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
}
