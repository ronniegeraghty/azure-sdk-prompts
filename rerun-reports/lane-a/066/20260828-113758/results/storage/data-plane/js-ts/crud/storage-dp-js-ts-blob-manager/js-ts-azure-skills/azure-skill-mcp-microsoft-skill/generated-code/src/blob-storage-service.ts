import { createReadStream, createWriteStream } from "node:fs";
import { pipeline } from "node:stream/promises";
import {
  type BlobItem,
  type BlobServiceClient,
  type BlockBlobParallelUploadOptions,
  type ContainerClient,
  RestError,
} from "@azure/storage-blob";

const DEFAULT_BUFFER_SIZE = 8 * 1024 * 1024;
const DEFAULT_CONCURRENCY = 5;
const LEASE_DURATION_SECONDS = 60;
const LEASE_RENEWAL_INTERVAL_MS = 40_000;

export interface UploadOptions {
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
  contentType?: string;
  bufferSize?: number;
  concurrency?: number;
  onProgress?: (uploadedBytes: number) => void;
}

export interface BlobSummary {
  name: string;
  contentLength?: number;
  contentType?: string;
  lastModified?: Date;
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
}

export class BlobStorageService {
  private readonly containerClient: ContainerClient;

  public constructor(
    blobServiceClient: BlobServiceClient,
    containerName: string,
  ) {
    this.containerClient = blobServiceClient.getContainerClient(containerName);
  }

  public async ensureContainer(): Promise<void> {
    await this.containerClient.createIfNotExists();
  }

  public async upload(
    localFilePath: string,
    blobName: string,
    options: UploadOptions = {},
  ): Promise<void> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);
    const uploadOptions = this.createUploadOptions(options);

    try {
      await blockBlobClient.getProperties();
    } catch (error: unknown) {
      if (!this.isStatusCode(error, 404)) {
        throw error;
      }

      try {
        await blockBlobClient.uploadStream(
          createReadStream(localFilePath),
          options.bufferSize ?? DEFAULT_BUFFER_SIZE,
          options.concurrency ?? DEFAULT_CONCURRENCY,
          {
            ...uploadOptions,
            conditions: { ifNoneMatch: "*" },
          },
        );
        return;
      } catch (creationError: unknown) {
        if (!this.isStatusCode(creationError, 412)) {
          throw creationError;
        }
        // Another writer created the blob first; continue through the leased path.
      }
    }

    await this.uploadWithRenewableLease(localFilePath, blobName, options);
  }

  public async download(
    blobName: string,
    destinationFilePath: string,
  ): Promise<void> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .download();
    if (!response.readableStreamBody) {
      throw new Error(`Blob "${blobName}" returned no readable stream.`);
    }

    await pipeline(
      response.readableStreamBody,
      createWriteStream(destinationFilePath),
    );
  }

  public async list(): Promise<BlobSummary[]> {
    const blobs: BlobSummary[] = [];
    for await (const item of this.containerClient.listBlobsFlat({
      includeMetadata: true,
      includeTags: true,
    })) {
      blobs.push(this.toBlobSummary(item));
    }
    return blobs;
  }

  public async delete(blobName: string): Promise<boolean> {
    const response = await this.containerClient
      .getBlobClient(blobName)
      .deleteIfExists({ deleteSnapshots: "include" });
    return response.succeeded;
  }

  private async uploadWithRenewableLease(
    localFilePath: string,
    blobName: string,
    options: UploadOptions,
  ): Promise<void> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);
    const leaseClient = blockBlobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(LEASE_DURATION_SECONDS);
    const abortController = new AbortController();
    let renewalError: unknown;

    const renewalTimer = setInterval(() => {
      void leaseClient.renewLease().catch((error: unknown) => {
        renewalError = error;
        abortController.abort();
      });
    }, LEASE_RENEWAL_INTERVAL_MS);
    renewalTimer.unref();

    let operationError: unknown;
    try {
      await blockBlobClient.uploadStream(
        createReadStream(localFilePath),
        options.bufferSize ?? DEFAULT_BUFFER_SIZE,
        options.concurrency ?? DEFAULT_CONCURRENCY,
        {
          ...this.createUploadOptions(options),
          abortSignal: abortController.signal,
          conditions: { leaseId: lease.leaseId },
        },
      );

      if (renewalError) {
        throw new Error("The blob lease could not be renewed during upload.", {
          cause: renewalError,
        });
      }
    } catch (error: unknown) {
      operationError = error;
    } finally {
      clearInterval(renewalTimer);
    }

    try {
      await leaseClient.releaseLease();
    } catch (releaseError: unknown) {
      if (operationError) {
        throw new AggregateError(
          [operationError, releaseError],
          `Upload of "${blobName}" failed and its lease could not be released.`,
        );
      }
      throw releaseError;
    }

    if (operationError) {
      throw operationError;
    }
  }

  private createUploadOptions(
    options: UploadOptions,
  ): BlockBlobParallelUploadOptions {
    return {
      metadata: options.metadata,
      tags: options.tags,
      blobHTTPHeaders: options.contentType
        ? { blobContentType: options.contentType }
        : undefined,
      onProgress: options.onProgress
        ? ({ loadedBytes }) => options.onProgress?.(loadedBytes)
        : undefined,
    };
  }

  private toBlobSummary(item: BlobItem): BlobSummary {
    return {
      name: item.name,
      contentLength: item.properties.contentLength,
      contentType: item.properties.contentType,
      lastModified: item.properties.lastModified,
      metadata: item.metadata,
      tags: item.tags,
    };
  }

  private isStatusCode(error: unknown, statusCode: number): boolean {
    return error instanceof RestError && error.statusCode === statusCode;
  }
}
