import { randomUUID } from "node:crypto";
import { mkdir, readFile, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { BlobStorageService } from "./blobStorageService.js";
import {
  createBlobServiceClient,
  loadBlobStorageConfig,
} from "./config.js";

async function main(): Promise<void> {
  const config = loadBlobStorageConfig();
  const blobServiceClient = createBlobServiceClient(config);
  const storage = new BlobStorageService(
    blobServiceClient,
    config.containerName,
  );

  const runId = randomUUID();
  const blobName = `demo/${runId}/sample.txt`;
  const workingDirectory = join(tmpdir(), `azure-blob-manager-${runId}`);
  const uploadPath = join(workingDirectory, "sample.txt");
  const downloadPath = join(workingDirectory, "downloaded.txt");
  let uploaded = false;

  await mkdir(workingDirectory, { recursive: true });

  try {
    console.log(`Ensuring container "${config.containerName}" exists...`);
    await storage.initialize();

    await writeFile(uploadPath, "Hello from Azure Blob Storage!\n", "utf8");
    console.log(`Uploading "${blobName}" with blob index tags...`);
    await storage.uploadFile(uploadPath, blobName, {
      contentType: "text/plain; charset=utf-8",
      metadata: { demoRunId: runId },
      tags: { project: "blob-manager", lifecycle: "demo" },
    });
    uploaded = true;
    console.log("Upload complete.");

    console.log("Listing blobs:");
    const blobs = await storage.listBlobs();
    for (const blob of blobs) {
      console.log(
        `- ${blob.name} (${blob.properties.contentLength ?? 0} bytes), tags=${JSON.stringify(blob.tags ?? {})}`,
      );
    }

    console.log(`Downloading "${blobName}"...`);
    await storage.downloadFile(blobName, downloadPath);
    console.log(`Downloaded content: ${JSON.stringify(
      await readFile(downloadPath, "utf8"),
    )}`);

    await writeFile(
      uploadPath,
      "This content was written while holding a renewable blob lease.\n",
      "utf8",
    );
    console.log("Acquiring a lease and overwriting the blob...");
    const overwrite = await storage.uploadFile(uploadPath, blobName, {
      contentType: "text/plain; charset=utf-8",
      metadata: { demoRunId: runId, updated: "true" },
      tags: { project: "blob-manager", lifecycle: "demo" },
    });
    if (!overwrite.leaseProtected) {
      throw new Error("Expected the overwrite to be protected by a blob lease.");
    }
    console.log("Lease-protected overwrite complete.");
  } finally {
    try {
      if (uploaded) {
        console.log(`Deleting "${blobName}"...`);
        const deleted = await storage.deleteBlob(blobName);
        console.log(deleted ? "Delete complete." : "Blob was already absent.");
      }
    } finally {
      await rm(workingDirectory, { recursive: true, force: true });
    }
  }
}

main().catch((error: unknown) => {
  console.error("Blob Storage demo failed:", error);
  process.exitCode = 1;
});
