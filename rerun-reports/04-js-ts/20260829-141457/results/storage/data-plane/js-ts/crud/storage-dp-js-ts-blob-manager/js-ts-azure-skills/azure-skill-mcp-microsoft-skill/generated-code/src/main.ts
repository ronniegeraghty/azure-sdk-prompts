import { readFile, writeFile } from "node:fs/promises";
import { resolve } from "node:path";
import { BlobStorageService } from "./blob-storage-service.js";
import {
  createBlobServiceClient,
  loadStorageConfig,
} from "./config.js";

async function main(): Promise<void> {
  const config = loadStorageConfig();
  const service = new BlobStorageService(
    createBlobServiceClient(config),
    config.containerName,
    {
      uploadBufferSize: config.uploadBufferSize,
      uploadConcurrency: config.uploadConcurrency,
      leaseWaitMs: config.leaseWaitMs,
      leasePollMs: config.leasePollMs,
    },
  );

  const blobName = "sample.txt";
  const sourcePath = resolve("sample.txt");
  const destinationPath = resolve("downloaded-sample.txt");

  console.log(`Ensuring container "${config.containerName}" exists...`);
  await service.ensureContainer();

  await writeFile(sourcePath, "Hello from the Azure Blob manager!\n", "utf8");
  console.log(`Uploading "${blobName}" with metadata and index tags...`);
  await service.upload(blobName, sourcePath, {
    contentType: "text/plain; charset=utf-8",
    metadata: { createdBy: "blob-manager-demo" },
    tags: { project: "blob-manager", environment: "demo" },
    onProgress: (bytes) => console.log(`  Uploaded ${bytes} bytes`),
  });
  console.log("Upload complete.");

  console.log("Listing blobs...");
  const blobs = await service.list();
  for (const blob of blobs) {
    console.log(
      `  ${blob.name} (${blob.properties.contentLength ?? 0} bytes)`,
    );
  }

  console.log(`Downloading "${blobName}"...`);
  await service.download(blobName, destinationPath);
  console.log("Downloaded content:");
  console.log(await readFile(destinationPath, "utf8"));

  await writeFile(
    sourcePath,
    "This content was written while holding a blob lease.\n",
    "utf8",
  );
  console.log(`Acquiring a lease and overwriting "${blobName}"...`);
  await service.upload(blobName, sourcePath, {
    contentType: "text/plain; charset=utf-8",
    metadata: { updatedBy: "blob-manager-demo" },
    tags: { project: "blob-manager", environment: "demo", version: "2" },
  });
  console.log("Lease-protected overwrite complete.");

  console.log(`Deleting "${blobName}"...`);
  const deleted = await service.delete(blobName);
  console.log(deleted ? "Delete complete." : "Blob was already absent.");
}

main().catch((error: unknown) => {
  console.error("Blob manager demo failed:", error);
  process.exitCode = 1;
});
