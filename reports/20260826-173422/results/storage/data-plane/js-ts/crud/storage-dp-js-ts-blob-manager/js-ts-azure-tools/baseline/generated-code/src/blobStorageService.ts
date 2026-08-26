import { createReadStream } from "node:fs";
import { stat } from "node:fs/promises";
import type {
  BlobItem,
  BlobRequestConditions,
  ContainerClient,
  Metadata,
  Tags,
} from "@azure/storage-blob";

const DEFAULT_BUFFER_SIZE = 8 * 1024 * 1024;
const DEFAULT_MAX_CONCURRENCY = 5;

export interface UploadOptions {
  metadata?: Metadata;
  tags?: Tags;
  bufferSize?: number;
  maxConcurrency?: number;
}

export interface UploadResult {
  etag: string | undefined;
  lastModified: Date | undefined;
  leaseAcquired: boolean;
}

function isPreconditionFailure(error: unknown): boolean {
  return (
    typeof error === "object" &&
    error !== null &&
    "statusCode" in error &&
    error.statusCode === 412
  );
}

export class BlobStorageService {
  public constructor(private readonly containerClient: ContainerClient) {}

  public async uploadFile(
    blobName: string,
    filePath: string,
    options: UploadOptions = {},
  ): Promise<UploadResult> {
    const file = await stat(filePath);
    if (!file.isFile()) {
      throw new Error(`Upload source is not a file: ${filePath}`);
    }

    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);
    const bufferSize = options.bufferSize ?? DEFAULT_BUFFER_SIZE;
    const maxConcurrency =
      options.maxConcurrency ?? DEFAULT_MAX_CONCURRENCY;

    if (bufferSize <= 0 || maxConcurrency <= 0) {
      throw new Error("bufferSize and maxConcurrency must be positive.");
    }

    if (await blockBlobClient.exists()) {
      return this.uploadWithLease(
        blobName,
        filePath,
        bufferSize,
        maxConcurrency,
        options,
      );
    }

    try {
      const response = await blockBlobClient.uploadStream(
        createReadStream(filePath, { highWaterMark: bufferSize }),
        bufferSize,
        maxConcurrency,
        {
          metadata: options.metadata,
          tags: options.tags,
          conditions: { ifNoneMatch: "*" },
        },
      );
      return {
        etag: response.etag,
        lastModified: response.lastModified,
        leaseAcquired: false,
      };
    } catch (error) {
      // Another writer may have created the blob after the existence check.
      if (!isPreconditionFailure(error)) {
        throw error;
      }
      return this.uploadWithLease(
        blobName,
        filePath,
        bufferSize,
        maxConcurrency,
        options,
      );
    }
  }

  private async uploadWithLease(
    blobName: string,
    filePath: string,
    bufferSize: number,
    maxConcurrency: number,
    options: UploadOptions,
  ): Promise<UploadResult> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);
    const leaseClient = blockBlobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(-1);
    const conditions: BlobRequestConditions = {
      leaseId: lease.leaseId,
    };

    try {
      const response = await blockBlobClient.uploadStream(
        createReadStream(filePath, { highWaterMark: bufferSize }),
        bufferSize,
        maxConcurrency,
        {
          metadata: options.metadata,
          tags: options.tags,
          conditions,
        },
      );
      return {
        etag: response.etag,
        lastModified: response.lastModified,
        leaseAcquired: true,
      };
    } finally {
      await leaseClient.releaseLease();
    }
  }

  public async downloadFile(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    await this.containerClient
      .getBlobClient(blobName)
      .downloadToFile(destinationPath);
  }

  public async listBlobs(): Promise<BlobItem[]> {
    const blobs: BlobItem[] = [];
    for await (const blob of this.containerClient.listBlobsFlat({
      includeMetadata: true,
      includeTags: true,
    })) {
      blobs.push(blob);
    }
    return blobs;
  }

  public async deleteBlob(blobName: string): Promise<boolean> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .deleteIfExists({ deleteSnapshots: "include" });
    return response.succeeded;
  }
}
