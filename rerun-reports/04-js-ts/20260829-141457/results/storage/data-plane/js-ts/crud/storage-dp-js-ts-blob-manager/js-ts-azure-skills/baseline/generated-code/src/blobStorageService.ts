import { createReadStream, createWriteStream } from "node:fs";
import { pipeline } from "node:stream/promises";
import {
  BlobItem,
  BlobServiceClient,
  ContainerClient,
  Metadata,
  Tags,
} from "@azure/storage-blob";

const UPLOAD_BUFFER_SIZE = 8 * 1024 * 1024;
const UPLOAD_CONCURRENCY = 5;

export interface UploadBlobOptions {
  metadata?: Metadata;
  tags?: Tags;
}

function hasStatusCode(error: unknown, ...statusCodes: number[]): boolean {
  return (
    typeof error === "object" &&
    error !== null &&
    "statusCode" in error &&
    typeof error.statusCode === "number" &&
    statusCodes.includes(error.statusCode)
  );
}

export class BlobStorageService {
  private readonly containerClient: ContainerClient;

  public constructor(
    blobServiceClient: BlobServiceClient,
    containerName: string,
  ) {
    this.containerClient = blobServiceClient.getContainerClient(containerName);
  }

  public async upload(
    blobName: string,
    sourceFilePath: string,
    options: UploadBlobOptions = {},
  ): Promise<void> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);

    try {
      await blockBlobClient.upload("", 0, {
        conditions: { ifNoneMatch: "*" },
      });
    } catch (error: unknown) {
      // A competing writer may create the blob between lookup and creation.
      if (!hasStatusCode(error, 409, 412)) {
        throw error;
      }
    }

    const leaseClient = blockBlobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(-1);

    try {
      await blockBlobClient.uploadStream(
        createReadStream(sourceFilePath),
        UPLOAD_BUFFER_SIZE,
        UPLOAD_CONCURRENCY,
        {
          metadata: options.metadata,
          tags: options.tags,
          conditions: { leaseId: lease.leaseId },
        },
      );
    } finally {
      await leaseClient.releaseLease();
    }
  }

  public async download(
    blobName: string,
    destinationFilePath: string,
  ): Promise<void> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .download();

    if (!response.readableStreamBody) {
      throw new Error(`Blob "${blobName}" returned no downloadable content.`);
    }

    await pipeline(
      response.readableStreamBody,
      createWriteStream(destinationFilePath),
    );
  }

  public async list(): Promise<BlobItem[]> {
    const blobs: BlobItem[] = [];
    for await (const blob of this.containerClient.listBlobsFlat({
      includeMetadata: true,
      includeTags: true,
    })) {
      blobs.push(blob);
    }
    return blobs;
  }

  public async delete(blobName: string): Promise<boolean> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .deleteIfExists({ deleteSnapshots: "include" });
    return response.succeeded;
  }
}
