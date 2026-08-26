import { RestError } from "@azure/core-rest-pipeline";
import { DefaultAzureCredential } from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const containerName = "my-container";
const blobName = "greeting.txt";
const content = "Hello Azure!";

async function main(): Promise<void> {
  const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;

  if (!accountName) {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_NAME must be set to your Azure Storage account name.",
    );
  }

  const credential = new DefaultAzureCredential();
  const serviceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );
  const containerClient = serviceClient.getContainerClient(containerName);

  const createResult = await containerClient.createIfNotExists();
  console.log(
    createResult.succeeded
      ? `Created container: ${containerName}`
      : `Container already exists: ${containerName}`,
  );

  const blockBlobClient = containerClient.getBlockBlobClient(blobName);
  await blockBlobClient.upload(content, Buffer.byteLength(content));
  console.log(`Uploaded blob: ${blobName}`);

  console.log("Blobs in container:");
  for await (const blob of containerClient.listBlobsFlat()) {
    console.log(`- ${blob.name}`);
  }

  const downloaded = await blockBlobClient.downloadToBuffer();
  console.log(`Downloaded content: ${downloaded.toString("utf8")}`);

  await blockBlobClient.delete();
  console.log(`Deleted blob: ${blobName}`);

  await containerClient.delete();
  console.log(`Deleted container: ${containerName}`);
}

main().catch((error: unknown) => {
  if (error instanceof RestError) {
    console.error("Azure Blob Storage request failed:", {
      name: error.name,
      code: error.code,
      statusCode: error.statusCode,
      message: error.message,
      requestId: error.request?.headers.get("x-ms-client-request-id"),
    });
  } else if (error instanceof Error) {
    console.error(`Error: ${error.message}`);
  } else {
    console.error("An unknown error occurred:", error);
  }

  process.exitCode = 1;
});
