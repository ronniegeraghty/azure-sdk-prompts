import { createReadStream } from "node:fs";
import { stat } from "node:fs/promises";
import {
  type BlobItem,
  type BlobServiceClient,
  type BlockBlobClient,
} from "@azure/storage-blob";

export interface UploadOptions {
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
  contentType?: string;
  bufferSize?: number;
  maxConcurrency?: number;
}

interface StatusCodeError {
  statusCode?: number;
}

function hasStatusCode(error: unknown, ...statusCodes: number[]): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const statusCode = (error as StatusCodeError).statusCode;
  return statusCode !== undefined && statusCodes.includes(statusCode);
}

export class BlobStorageManager {
  public constructor(
    private readonly serviceClient: BlobServiceClient,
    private readonly containerName: string,
  ) {
    if (!containerName) {
      throw new Error("containerName must not be empty.");
    }
  }

  public async upload(
    filePath: string,
    blobName: string,
    options: UploadOptions = {},
  ): Promise<void> {
    const file = await stat(filePath);
    if (!file.isFile()) {
      throw new Error(`Upload source is not a file: ${filePath}`);
    }

    const blobClient = this.getBlockBlobClient(blobName);
    await this.ensureBlobExists(blobClient);

    const leaseClient = blobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(-1);
    if (!lease.leaseId) {
      throw new Error(`Azure did not return a lease ID for blob "${blobName}".`);
    }

    try {
      const bufferSize = options.bufferSize ?? 8 * 1024 * 1024;
      const maxConcurrency = options.maxConcurrency ?? 5;
      const source = createReadStream(filePath, { highWaterMark: bufferSize });

      await blobClient.uploadStream(source, bufferSize, maxConcurrency, {
        conditions: { leaseId: lease.leaseId },
        ...(options.metadata ? { metadata: options.metadata } : {}),
        ...(options.tags ? { tags: options.tags } : {}),
        ...(options.contentType
          ? { blobHTTPHeaders: { blobContentType: options.contentType } }
          : {}),
      });
    } finally {
      await leaseClient.releaseLease();
    }
  }

  public async download(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    await this.getBlockBlobClient(blobName).downloadToFile(destinationPath);
  }

  public async list(): Promise<BlobItem[]> {
    const blobs: BlobItem[] = [];
    const containerClient = this.serviceClient.getContainerClient(
      this.containerName,
    );

    for await (const blob of containerClient.listBlobsFlat({
      includeMetadata: true,
      includeTags: true,
    })) {
      blobs.push(blob);
    }

    return blobs;
  }

  public async delete(blobName: string): Promise<boolean> {
    const response = await this.getBlockBlobClient(blobName).deleteIfExists({
      deleteSnapshots: "include",
    });
    return response.succeeded;
  }

  private getBlockBlobClient(blobName: string): BlockBlobClient {
    if (!blobName) {
      throw new Error("blobName must not be empty.");
    }

    return this.serviceClient
      .getContainerClient(this.containerName)
      .getBlockBlobClient(blobName);
  }

  private async ensureBlobExists(blobClient: BlockBlobClient): Promise<void> {
    try {
      await blobClient.getProperties();
      return;
    } catch (error) {
      if (!hasStatusCode(error, 404)) {
        throw error;
      }
    }

    try {
      await blobClient.upload("", 0, {
        conditions: { ifNoneMatch: "*" },
      });
    } catch (error) {
      // Another writer may have created the blob after our properties check.
      if (!hasStatusCode(error, 409, 412)) {
        throw error;
      }
    }
  }
}
