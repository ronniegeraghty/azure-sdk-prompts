import { createReadStream, createWriteStream } from "node:fs";
import { mkdir } from "node:fs/promises";
import { dirname } from "node:path";
import { pipeline } from "node:stream/promises";
import {
  type BlobItem,
  type BlobServiceClient,
  type BlockBlobClient,
  RestError,
} from "@azure/storage-blob";

export interface UploadOptions {
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
  contentType?: string;
  onProgress?: (uploadedBytes: number) => void;
}

export interface BlobStorageServiceOptions {
  uploadBufferSize?: number;
  uploadConcurrency?: number;
  leaseWaitMs?: number;
  leasePollMs?: number;
}

const DEFAULT_UPLOAD_BUFFER_SIZE = 8 * 1024 * 1024;
const DEFAULT_UPLOAD_CONCURRENCY = 5;
const DEFAULT_LEASE_WAIT_MS = 30_000;
const DEFAULT_LEASE_POLL_MS = 1_000;

export class BlobStorageService {
  private readonly containerClient;
  private readonly uploadBufferSize: number;
  private readonly uploadConcurrency: number;
  private readonly leaseWaitMs: number;
  private readonly leasePollMs: number;

  public constructor(
    blobServiceClient: BlobServiceClient,
    containerName: string,
    options: BlobStorageServiceOptions = {},
  ) {
    this.containerClient = blobServiceClient.getContainerClient(containerName);
    this.uploadBufferSize =
      options.uploadBufferSize ?? DEFAULT_UPLOAD_BUFFER_SIZE;
    this.uploadConcurrency =
      options.uploadConcurrency ?? DEFAULT_UPLOAD_CONCURRENCY;
    this.leaseWaitMs = options.leaseWaitMs ?? DEFAULT_LEASE_WAIT_MS;
    this.leasePollMs = options.leasePollMs ?? DEFAULT_LEASE_POLL_MS;
  }

  public async ensureContainer(): Promise<void> {
    await this.containerClient.createIfNotExists();
  }

  public async upload(
    blobName: string,
    filePath: string,
    options: UploadOptions = {},
  ): Promise<void> {
    const blobClient = this.containerClient.getBlockBlobClient(blobName);

    // An atomic create handles a missing blob; if another writer wins that
    // race, retry through the lease-protected existing-blob path.
    for (;;) {
      if (!(await this.blobExists(blobClient))) {
        try {
          await this.uploadStream(blobClient, filePath, options, undefined, "*");
          return;
        } catch (error: unknown) {
          if (!this.isCreateRace(error)) {
            throw error;
          }
        }
      }

      const leaseClient = blobClient.getBlobLeaseClient();
      await this.acquireLease(leaseClient);
      let operationError: unknown;
      try {
        await this.uploadStream(
          blobClient,
          filePath,
          options,
          leaseClient.leaseId,
        );
        return;
      } catch (error: unknown) {
        operationError = error;
        throw error;
      } finally {
        try {
          await leaseClient.releaseLease();
        } catch (releaseError: unknown) {
          if (operationError === undefined) {
            throw releaseError;
          }
          console.error("Failed to release the blob lease:", releaseError);
        }
      }
    }
  }

  public async download(blobName: string, destinationPath: string): Promise<void> {
    const blobClient = this.containerClient.getBlobClient(blobName);
    const response = await blobClient.download();
    if (!response.readableStreamBody) {
      throw new Error(`Blob ${blobName} did not return a readable stream.`);
    }

    await mkdir(dirname(destinationPath), { recursive: true });
    await pipeline(
      response.readableStreamBody,
      createWriteStream(destinationPath),
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

  private async uploadStream(
    blobClient: BlockBlobClient,
    filePath: string,
    options: UploadOptions,
    leaseId?: string,
    ifNoneMatch?: string,
  ): Promise<void> {
    const source = createReadStream(filePath);
    await blobClient.uploadStream(
      source,
      this.uploadBufferSize,
      this.uploadConcurrency,
      {
        metadata: options.metadata,
        tags: options.tags,
        blobHTTPHeaders: options.contentType
          ? { blobContentType: options.contentType }
          : undefined,
        conditions: {
          leaseId,
          ifNoneMatch,
        },
        onProgress: options.onProgress
          ? ({ loadedBytes }) => options.onProgress?.(loadedBytes)
          : undefined,
      },
    );
  }

  private async blobExists(blobClient: BlockBlobClient): Promise<boolean> {
    try {
      await blobClient.getProperties();
      return true;
    } catch (error: unknown) {
      if (error instanceof RestError && error.statusCode === 404) {
        return false;
      }
      throw error;
    }
  }

  private async acquireLease(
    leaseClient: ReturnType<BlockBlobClient["getBlobLeaseClient"]>,
  ): Promise<void> {
    const deadline = Date.now() + this.leaseWaitMs;
    for (;;) {
      try {
        await leaseClient.acquireLease(-1);
        return;
      } catch (error: unknown) {
        if (!this.isLeaseConflict(error) || Date.now() >= deadline) {
          throw error;
        }
        await new Promise((resolve) => setTimeout(resolve, this.leasePollMs));
      }
    }
  }

  private isCreateRace(error: unknown): boolean {
    return (
      error instanceof RestError &&
      (error.statusCode === 409 || error.statusCode === 412)
    );
  }

  private isLeaseConflict(error: unknown): boolean {
    return (
      error instanceof RestError &&
      error.statusCode === 409 &&
      (error.code === "LeaseAlreadyPresent" ||
        error.code === "LeaseIsBreakingAndCannotBeAcquired")
    );
  }
}
