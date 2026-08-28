import { mkdtemp, readFile, rm, writeFile } from "node:fs/promises";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { BlobStorageManager } from "./blobStorageManager.js";
import { createBlobServiceClient } from "./config.js";

async function main(): Promise<void> {
  const containerName = process.env.AZURE_STORAGE_CONTAINER;
  if (!containerName) {
    throw new Error("AZURE_STORAGE_CONTAINER is required.");
  }

  const serviceClient = createBlobServiceClient();
  const manager = new BlobStorageManager(serviceClient, containerName);
  const blobName = process.env.AZURE_STORAGE_DEMO_BLOB ?? "blob-manager-demo.txt";
  const workDirectory = await mkdtemp(join(tmpdir(), "blob-manager-demo-"));
  const uploadPath = join(workDirectory, "sample.txt");
  const downloadPath = join(workDirectory, "downloaded.txt");
  let uploaded = false;

  try {
    await writeFile(uploadPath, "Hello from Azure Blob Storage!\n", "utf8");

    console.log(`Uploading "${blobName}" with index tags...`);
    await manager.upload(uploadPath, blobName, {
      contentType: "text/plain; charset=utf-8",
      metadata: { source: "blob-manager-demo" },
      tags: { project: "blob-manager", environment: "demo" },
    });
    uploaded = true;
    console.log("Upload complete.");

    console.log(`Listing blobs in container "${containerName}"...`);
    const blobs = await manager.list();
    for (const blob of blobs) {
      console.log(`- ${blob.name} (${blob.properties.contentLength ?? 0} bytes)`);
    }

    console.log(`Downloading "${blobName}"...`);
    await manager.download(blobName, downloadPath);
    console.log(`Downloaded content: ${await readFile(downloadPath, "utf8")}`);

    await writeFile(
      uploadPath,
      "This content was written while holding an exclusive blob lease.\n",
      "utf8",
    );
    console.log(`Acquiring a lease and overwriting "${blobName}"...`);
    await manager.upload(uploadPath, blobName, {
      contentType: "text/plain; charset=utf-8",
      metadata: { source: "blob-manager-demo", revision: "2" },
      tags: { project: "blob-manager", environment: "demo" },
    });
    console.log("Lease-protected overwrite complete.");
  } finally {
    if (uploaded) {
      console.log(`Deleting "${blobName}"...`);
      const deleted = await manager.delete(blobName);
      console.log(deleted ? "Delete complete." : "Blob was already absent.");
    }

    await rm(workDirectory, { recursive: true, force: true });
  }
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
