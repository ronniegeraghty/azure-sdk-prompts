import { readFile, rm, writeFile } from "node:fs/promises";
import { resolve } from "node:path";
import { BlobStorageService } from "./blob-storage-service.js";
import { createStorageConfiguration } from "./config.js";

const SAMPLE_BLOB_NAME = "blob-manager-sample.txt";
const SAMPLE_FILE = resolve("sample-upload.txt");
const DOWNLOADED_FILE = resolve("sample-download.txt");

async function main(): Promise<void> {
  const { blobServiceClient, containerName } = createStorageConfiguration();
  const storage = new BlobStorageService(blobServiceClient, containerName);

  console.log(`[setup] Ensuring container "${containerName}" exists...`);
  await storage.ensureContainer();

  await writeFile(SAMPLE_FILE, "Hello from Azure Blob Storage!\n", "utf8");

  console.log(`[1/5] Uploading "${SAMPLE_BLOB_NAME}" with index tags...`);
  await storage.upload(SAMPLE_FILE, SAMPLE_BLOB_NAME, {
    contentType: "text/plain; charset=utf-8",
    metadata: { source: "blob-manager-demo" },
    tags: { project: "blob-manager", environment: "demo" },
    onProgress: (bytes) => console.log(`      uploaded ${bytes} bytes`),
  });

  console.log("[2/5] Listing blobs...");
  const blobs = await storage.list();
  for (const blob of blobs) {
    console.log(
      `      ${blob.name} (${blob.contentLength ?? "unknown"} bytes) tags=${JSON.stringify(blob.tags ?? {})}`,
    );
  }

  console.log(`[3/5] Downloading "${SAMPLE_BLOB_NAME}"...`);
  await storage.download(SAMPLE_BLOB_NAME, DOWNLOADED_FILE);
  console.log(`      content: ${(await readFile(DOWNLOADED_FILE, "utf8")).trim()}`);

  console.log("[4/5] Acquiring a renewable lease and overwriting the blob...");
  await writeFile(
    SAMPLE_FILE,
    "This content was written while holding a blob lease.\n",
    "utf8",
  );
  await storage.upload(SAMPLE_FILE, SAMPLE_BLOB_NAME, {
    contentType: "text/plain; charset=utf-8",
    metadata: { source: "blob-manager-demo", update: "leased" },
    tags: { project: "blob-manager", environment: "demo" },
  });
  console.log("      leased overwrite complete");

  console.log(`[5/5] Deleting "${SAMPLE_BLOB_NAME}"...`);
  const deleted = await storage.delete(SAMPLE_BLOB_NAME);
  console.log(`      ${deleted ? "deleted" : "blob did not exist"}`);

  await Promise.all([
    rm(SAMPLE_FILE, { force: true }),
    rm(DOWNLOADED_FILE, { force: true }),
  ]);
  console.log("Demo complete.");
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
