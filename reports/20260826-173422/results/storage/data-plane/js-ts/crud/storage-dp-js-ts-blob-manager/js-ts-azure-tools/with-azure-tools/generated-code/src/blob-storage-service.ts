import { createReadStream, createWriteStream } from "node:fs";
import { access } from "node:fs/promises";
import { pipeline } from "node:stream/promises";
import {
  type BlobItem,
  type BlobServiceClient,
  type BlockBlobUploadStreamOptions,
  type BlockBlobClient,
  type ContainerClient,
  RestError,
} from "@azure/storage-blob";

const DEFAULT_BUFFER_SIZE = 8 * 1024 * 1024;
const DEFAULT_MAX_CONCURRENCY = 5;

export interface UploadFileOptions {
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
  contentType?: string;
  bufferSize?: number;
  maxConcurrency?: number;
  onProgress?: (bytesUploaded: number) => void;
}

export interface UploadResult {
  etag: string | undefined;
  versionId: string | undefined;
  usedLease: boolean;
}

export interface BlobSummary {
  name: string;
  contentLength: number | undefined;
  contentType: string | undefined;
  etag: string;
  lastModified: Date;
}

export class BlobStorageService {
  private readonly containerClient: ContainerClient;

  public constructor(
    blobServiceClient: BlobServiceClient,
    containerName: string,
  ) {
    this.containerClient = blobServiceClient.getContainerClient(containerName);
  }

  public async ensureContainerExists(): Promise<void> {
    await this.containerClient.createIfNotExists();
  }

  public async uploadFile(
    blobName: string,
    filePath: string,
    options: UploadFileOptions = {},
  ): Promise<UploadResult> {
    await access(filePath);
    const blobClient = this.containerClient.getBlockBlobClient(blobName);

    try {
      return await this.uploadWithConcurrencyProtection(
        blobClient,
        filePath,
        options,
      );
    } catch (error) {
      // Another writer may have created the blob after our existence check.
      if (isPreconditionFailure(error)) {
        return this.uploadWithConcurrencyProtection(
          blobClient,
          filePath,
          options,
        );
      }
      throw error;
    }
  }

  public async downloadToFile(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .download();

    if (!response.readableStreamBody) {
      throw new Error(`Blob ${blobName} returned no readable response body.`);
    }

    await pipeline(
      response.readableStreamBody,
      createWriteStream(destinationPath),
    );
  }

  public async *listBlobs(): AsyncGenerator<BlobSummary> {
    for await (const blob of this.containerClient.listBlobsFlat()) {
      yield toBlobSummary(blob);
    }
  }

  public async deleteBlob(blobName: string): Promise<boolean> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .deleteIfExists({ deleteSnapshots: "include" });
    return response.succeeded;
  }

  private async uploadWithConcurrencyProtection(
    blobClient: BlockBlobClient,
    filePath: string,
    options: UploadFileOptions,
  ): Promise<UploadResult> {
    const exists = await blobClient.exists();
    if (!exists) {
      const response = await this.uploadStream(blobClient, filePath, options, {
        ifNoneMatch: "*",
      });
      return {
        etag: response.etag,
        versionId: response.versionId,
        usedLease: false,
      };
    }

    const leaseClient = blobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(-1);

    try {
      if (!lease.leaseId) {
        throw new Error(
          `Azure did not return a lease ID for ${blobClient.name}.`,
        );
      }
      const response = await this.uploadStream(blobClient, filePath, options, {
        leaseId: lease.leaseId,
      });
      return {
        etag: response.etag,
        versionId: response.versionId,
        usedLease: true,
      };
    } finally {
      await leaseClient.releaseLease();
    }
  }

  private async uploadStream(
    blobClient: BlockBlobClient,
    filePath: string,
    options: UploadFileOptions,
    conditions: { ifNoneMatch: "*" } | { leaseId: string },
  ) {
    const bufferSize = options.bufferSize ?? DEFAULT_BUFFER_SIZE;
    const maxConcurrency =
      options.maxConcurrency ?? DEFAULT_MAX_CONCURRENCY;

    if (!Number.isSafeInteger(bufferSize) || bufferSize <= 0) {
      throw new Error("bufferSize must be a positive integer.");
    }
    if (!Number.isSafeInteger(maxConcurrency) || maxConcurrency <= 0) {
      throw new Error("maxConcurrency must be a positive integer.");
    }

    const onProgress = options.onProgress;
    const uploadOptions: BlockBlobUploadStreamOptions = {
      conditions,
      ...(options.metadata ? { metadata: options.metadata } : {}),
      ...(options.tags ? { tags: options.tags } : {}),
      ...(options.contentType
        ? { blobHTTPHeaders: { blobContentType: options.contentType } }
        : {}),
      ...(onProgress
        ? {
            onProgress: (progress) => onProgress(progress.loadedBytes),
          }
        : {}),
    };

    return blobClient.uploadStream(
      createReadStream(filePath),
      bufferSize,
      maxConcurrency,
      uploadOptions,
    );
  }
}

function isPreconditionFailure(error: unknown): boolean {
  return error instanceof RestError && error.statusCode === 412;
}

function toBlobSummary(blob: BlobItem): BlobSummary {
  return {
    name: blob.name,
    contentLength: blob.properties.contentLength,
    contentType: blob.properties.contentType,
    etag: blob.properties.etag,
    lastModified: blob.properties.lastModified,
  };
}
