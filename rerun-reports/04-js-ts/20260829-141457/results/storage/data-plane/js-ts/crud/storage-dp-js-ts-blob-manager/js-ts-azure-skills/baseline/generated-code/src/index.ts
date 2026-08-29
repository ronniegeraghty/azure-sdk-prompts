import { mkdir, readFile, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { BlobStorageService } from "./blobStorageService.js";
import { createBlobStorageConfiguration } from "./config.js";

async function main(): Promise<void> {
  const { blobServiceClient, containerName } =
    createBlobStorageConfiguration();
  const storage = new BlobStorageService(blobServiceClient, containerName);

  const workingDirectory = join(process.cwd(), ".demo");
  const sourcePath = join(workingDirectory, "sample.txt");
  const downloadPath = join(workingDirectory, "sample.downloaded.txt");
  const blobName = "blob-manager-sample.txt";
  let uploaded = false;

  await mkdir(workingDirectory, { recursive: true });

  try {
    await writeFile(sourcePath, "Hello from the Azure Blob manager!\n");

    console.log(`Uploading "${blobName}" with metadata and index tags...`);
    await storage.upload(blobName, sourcePath, {
      metadata: { source: "typescript-demo" },
      tags: { environment: "demo", category: "sample" },
    });
    uploaded = true;
    console.log("Upload complete.");

    console.log(`Listing blobs in container "${containerName}"...`);
    const blobs = await storage.list();
    for (const blob of blobs) {
      console.log(`- ${blob.name}`);
    }

    console.log(`Downloading "${blobName}"...`);
    await storage.download(blobName, downloadPath);
    console.log("Downloaded content:");
    console.log(await readFile(downloadPath, "utf8"));

    console.log("Acquiring a lease and overwriting the blob...");
    await writeFile(sourcePath, "This content was written under a blob lease.\n");
    await storage.upload(blobName, sourcePath, {
      metadata: { source: "typescript-demo", revision: "2" },
      tags: { environment: "demo", category: "sample" },
    });
    console.log("Lease-protected overwrite complete.");
  } finally {
    if (uploaded) {
      console.log(`Deleting "${blobName}"...`);
      const deleted = await storage.delete(blobName);
      console.log(deleted ? "Delete complete." : "Blob was already absent.");
    }
    await rm(workingDirectory, { recursive: true, force: true });
  }
}

main().catch((error: unknown) => {
  console.error("Blob Storage demo failed:", error);
  process.exitCode = 1;
});
