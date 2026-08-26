import { readFile, unlink, writeFile } from "node:fs/promises";
import { BlobStorageService } from "./blob-storage-service.js";
import { createStorageConfiguration } from "./config.js";

const SAMPLE_FILE = "sample.txt";
const DOWNLOADED_FILE = "downloaded-sample.txt";
const BLOB_NAME = "sample.txt";

async function main(): Promise<void> {
  const { blobServiceClient, containerName } = createStorageConfiguration();
  const storage = new BlobStorageService(blobServiceClient, containerName);

  console.log(`Ensuring container "${containerName}" exists...`);
  await storage.ensureContainerExists();

  await writeFile(
    SAMPLE_FILE,
    "Hello from the reusable Azure Blob Storage manager!\n",
    "utf8",
  );

  console.log(`Uploading "${SAMPLE_FILE}" with metadata and index tags...`);
  await storage.uploadFile(BLOB_NAME, SAMPLE_FILE, {
    contentType: "text/plain; charset=utf-8",
    metadata: { source: "blob-manager-demo" },
    tags: { environment: "demo", documentType: "sample" },
    onProgress: (bytes) => console.log(`  Uploaded ${bytes} bytes`),
  });
  console.log("Upload complete.");

  console.log("Listing blobs:");
  for await (const blob of storage.listBlobs()) {
    console.log(
      `  ${blob.name} (${blob.contentLength ?? "unknown"} bytes, ${blob.contentType ?? "unknown type"})`,
    );
  }

  console.log(`Downloading "${BLOB_NAME}" to "${DOWNLOADED_FILE}"...`);
  await storage.downloadToFile(BLOB_NAME, DOWNLOADED_FILE);
  const content = await readFile(DOWNLOADED_FILE, "utf8");
  console.log(`Downloaded content: ${JSON.stringify(content)}`);

  await writeFile(
    SAMPLE_FILE,
    "This content was written while holding an Azure Blob lease.\n",
    "utf8",
  );
  console.log(`Acquiring a lease and overwriting "${BLOB_NAME}"...`);
  const overwrite = await storage.uploadFile(BLOB_NAME, SAMPLE_FILE, {
    contentType: "text/plain; charset=utf-8",
    metadata: { source: "blob-manager-demo", operation: "lease-overwrite" },
    tags: { environment: "demo", documentType: "sample", version: "2" },
  });
  if (!overwrite.usedLease) {
    throw new Error("Expected the overwrite to be protected by a blob lease.");
  }
  console.log("Lease-protected overwrite complete.");

  console.log(`Deleting "${BLOB_NAME}"...`);
  const deleted = await storage.deleteBlob(BLOB_NAME);
  console.log(deleted ? "Blob deleted." : "Blob did not exist.");

  await Promise.all([
    unlink(SAMPLE_FILE).catch(ignoreMissingFile),
    unlink(DOWNLOADED_FILE).catch(ignoreMissingFile),
  ]);
}

function ignoreMissingFile(error: unknown): void {
  if (
    error instanceof Error &&
    "code" in error &&
    error.code === "ENOENT"
  ) {
    return;
  }
  throw error;
}

main().catch((error: unknown) => {
  console.error("Blob Storage demo failed:", error);
  process.exitCode = 1;
});
