import { AIProjectClient } from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";
import { BlobClient } from "@azure/storage-blob";
import { dirname, resolve } from "node:path";
import { mkdir, readFile } from "node:fs/promises";

const endpoint = requiredEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
const datasetName = requiredEnvironmentVariable("DATASET_NAME");
const datasetVersion = requiredEnvironmentVariable("DATASET_VERSION");
const inputPath = requiredEnvironmentVariable("DATA_FILE_PATH");
const downloadPath = requiredEnvironmentVariable("DOWNLOAD_FILE_PATH");
const expectedBytes = await readFile(inputPath);

const client = new AIProjectClient(endpoint, new DefaultAzureCredential());
let uploaded = false;

try {
  await client.datasets.uploadFile(datasetName, datasetVersion, inputPath);
  uploaded = true;

  const dataset = await client.datasets.get(datasetName, datasetVersion);
  console.log(
    `Dataset: name=${dataset.name} version=${dataset.version} id=${dataset.id}`
      + ` type=${dataset.type} dataUri=${dataset.dataUri}`,
  );

  const credential = await client.datasets.getCredentials(
    datasetName,
    datasetVersion,
  );
  const blobUri = new URL(credential.blobReference.blobUri);
  const authorizedBlobUri = new URL(
    credential.blobReference.credential.sasUri,
  );
  authorizedBlobUri.pathname = blobUri.pathname;
  const blobClient = new BlobClient(authorizedBlobUri.toString());

  await mkdir(dirname(resolve(downloadPath)), { recursive: true });
  await blobClient.downloadToFile(downloadPath);

  const downloadedBytes = await readFile(downloadPath);
  if (!expectedBytes.equals(downloadedBytes)) {
    throw new Error("Downloaded bytes don't match the source file.");
  }
  console.log(`Downloaded bytes verified: ${downloadedBytes.length}`);
} finally {
  if (uploaded) {
    await client.datasets.delete(datasetName, datasetVersion);
  }
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}
