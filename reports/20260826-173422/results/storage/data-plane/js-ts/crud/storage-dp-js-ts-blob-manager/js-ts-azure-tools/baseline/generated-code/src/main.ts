import { mkdtemp, readFile, rm, writeFile } from "node:fs/promises";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { BlobStorageService } from "./blobStorageService.js";
import {
  createBlobServiceClient,
  loadBlobStorageConfig,
} from "./config.js";

async function main(): Promise<void> {
  const config = loadBlobStorageConfig();
  const serviceClient = createBlobServiceClient(config);
  const containerClient = serviceClient.getContainerClient(
    config.containerName,
  );
  const blobStorage = new BlobStorageService(containerClient);

  const workingDirectory = await mkdtemp(join(tmpdir(), "blob-manager-"));
  const sourcePath = join(workingDirectory, "sample.txt");
  const downloadPath = join(workingDirectory, "downloaded.txt");
  const blobName = `sample-${Date.now()}.txt`;
  let blobUploaded = false;
  let blobDeleted = false;

  try {
    console.log(`[upload] Uploading ${blobName}...`);
    await writeFile(
      sourcePath,
      "Hello from the Azure Blob Storage manager!\n",
      "utf8",
    );
    const upload = await blobStorage.uploadFile(blobName, sourcePath, {
      metadata: {
        source: "blob-manager-demo",
      },
      tags: {
        project: "blob-manager",
        environment: "demo",
      },
    });
    blobUploaded = true;
    console.log(`[upload] Complete. ETag: ${upload.etag ?? "not returned"}`);

    console.log("[list] Listing blobs in the container...");
    const blobs = await blobStorage.listBlobs();
    for (const blob of blobs) {
      console.log(
        `  - ${blob.name} (${blob.properties.contentLength ?? 0} bytes)`,
      );
    }
    console.log(`[list] Found ${blobs.length} blob(s).`);

    console.log(`[download] Downloading ${blobName}...`);
    await blobStorage.downloadFile(blobName, downloadPath);
    const downloadedContent = await readFile(downloadPath, "utf8");
    console.log(`[download] Content: ${JSON.stringify(downloadedContent)}`);

    console.log(`[overwrite] Acquiring a lease and overwriting ${blobName}...`);
    await writeFile(
      sourcePath,
      "This content was safely overwritten while holding a blob lease.\n",
      "utf8",
    );
    const overwrite = await blobStorage.uploadFile(blobName, sourcePath, {
      metadata: {
        source: "blob-manager-demo",
        revision: "2",
      },
      tags: {
        project: "blob-manager",
        environment: "demo",
        revision: "2",
      },
    });
    console.log(
      `[overwrite] Complete. Lease acquired: ${overwrite.leaseAcquired}.`,
    );

    console.log(`[delete] Deleting ${blobName}...`);
    blobDeleted = await blobStorage.deleteBlob(blobName);
    console.log(`[delete] ${blobDeleted ? "Complete." : "Blob did not exist."}`);
  } finally {
    if (blobUploaded && !blobDeleted) {
      console.log(`[cleanup] Deleting ${blobName} after an earlier failure...`);
      await blobStorage.deleteBlob(blobName);
    }
    await rm(workingDirectory, { recursive: true });
  }
}

main().catch((error: unknown) => {
  console.error("[error] Demo failed:", error);
  process.exitCode = 1;
});
