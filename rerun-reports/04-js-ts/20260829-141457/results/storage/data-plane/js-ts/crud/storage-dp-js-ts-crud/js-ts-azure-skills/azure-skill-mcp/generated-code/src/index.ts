import { RestError } from "@azure/core-rest-pipeline";
import { DefaultAzureCredential } from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const containerName = "my-container";
const blobName = "greeting.txt";
const content = "Hello Azure!";

async function readableStreamToString(
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
      "Set AZURE_STORAGE_ACCOUNT_NAME to the target storage account name.",
    );
  }

  const credential = new DefaultAzureCredential();
  const serviceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );
  const containerClient = serviceClient.getContainerClient(containerName);
  const blockBlobClient = containerClient.getBlockBlobClient(blobName);
  let containerIsReady = false;

  try {
    const createResult = await containerClient.createIfNotExists();
    containerIsReady = true;
    console.log(
      createResult.succeeded
        ? `Created container: ${containerName}`
        : `Container already exists: ${containerName}`,
    );

    await blockBlobClient.upload(content, Buffer.byteLength(content), {
      blobHTTPHeaders: { blobContentType: "text/plain; charset=utf-8" },
    });
    console.log(`Uploaded blob: ${blobName}`);

    console.log("Blobs in container:");
    for await (const blob of containerClient.listBlobsFlat()) {
      console.log(`- ${blob.name}`);
    }

    const downloadResponse = await blockBlobClient.download();
    if (!downloadResponse.readableStreamBody) {
      throw new Error(`Blob download returned no content for ${blobName}.`);
    }

    const downloadedContent = await readableStreamToString(
      downloadResponse.readableStreamBody,
    );
    console.log(`Downloaded content: ${downloadedContent}`);
  } finally {
    if (containerIsReady) {
      const deleteBlobResult = await blockBlobClient.deleteIfExists();
      console.log(
        deleteBlobResult.succeeded
          ? `Deleted blob: ${blobName}`
          : `Blob did not exist: ${blobName}`,
      );

      const deleteContainerResult = await containerClient.deleteIfExists();
      console.log(
        deleteContainerResult.succeeded
          ? `Deleted container: ${containerName}`
          : `Container did not exist: ${containerName}`,
      );
    }
  }
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof RestError) {
    console.error("Azure Storage request failed.", {
      name: error.name,
      code: error.code,
      statusCode: error.statusCode,
      message: error.message,
      requestId: error.request?.requestId,
    });
  } else if (error instanceof Error) {
    console.error(`Application error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
}
