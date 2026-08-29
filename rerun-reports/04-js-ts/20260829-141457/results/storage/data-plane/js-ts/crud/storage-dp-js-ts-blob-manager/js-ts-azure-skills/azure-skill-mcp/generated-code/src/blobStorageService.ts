import { createReadStream, createWriteStream } from "node:fs";
import { rename, rm, stat } from "node:fs/promises";
import { randomUUID } from "node:crypto";
import { pipeline } from "node:stream/promises";
import type {
  BlobItem,
  BlobLeaseClient,
  BlobServiceClient,
  BlockBlobClient,
  ContainerClient,
} from "@azure/storage-blob";

export interface UploadFileOptions {
  metadata?: Record<string, string>;
  tags?: Record<string, string>;
  contentType?: string;
}

export interface UploadFileResult {
  etag?: string;
  lastModified?: Date;
  leaseProtected: boolean;
}

export interface BlobStorageServiceOptions {
  uploadBufferSize?: number;
  uploadConcurrency?: number;
  leaseDurationSeconds?: number;
  leaseRenewalIntervalMs?: number;
}

const DEFAULT_UPLOAD_BUFFER_SIZE = 8 * 1024 * 1024;
const DEFAULT_UPLOAD_CONCURRENCY = 5;
const DEFAULT_LEASE_DURATION_SECONDS = 60;
const DEFAULT_LEASE_RENEWAL_INTERVAL_MS = 30_000;

function hasStatusCode(error: unknown, statusCode: number): boolean {
  return (
    typeof error === "object" &&
    error !== null &&
    "statusCode" in error &&
    error.statusCode === statusCode
  );
}

function validatePositiveInteger(name: string, value: number): void {
  if (!Number.isSafeInteger(value) || value <= 0) {
    throw new Error(`${name} must be a positive integer.`);
  }
}

export class BlobStorageService {
  private readonly containerClient: ContainerClient;
  private readonly uploadBufferSize: number;
  private readonly uploadConcurrency: number;
  private readonly leaseDurationSeconds: number;
  private readonly leaseRenewalIntervalMs: number;

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
    this.leaseDurationSeconds =
      options.leaseDurationSeconds ?? DEFAULT_LEASE_DURATION_SECONDS;
    this.leaseRenewalIntervalMs =
      options.leaseRenewalIntervalMs ?? DEFAULT_LEASE_RENEWAL_INTERVAL_MS;

    validatePositiveInteger("uploadBufferSize", this.uploadBufferSize);
    validatePositiveInteger("uploadConcurrency", this.uploadConcurrency);
    validatePositiveInteger("leaseDurationSeconds", this.leaseDurationSeconds);
    validatePositiveInteger(
      "leaseRenewalIntervalMs",
      this.leaseRenewalIntervalMs,
    );
    if (
      this.leaseDurationSeconds < 15 ||
      this.leaseDurationSeconds > 60
    ) {
      throw new Error("leaseDurationSeconds must be between 15 and 60.");
    }
    if (
      this.leaseRenewalIntervalMs >= this.leaseDurationSeconds * 1_000
    ) {
      throw new Error(
        "leaseRenewalIntervalMs must be shorter than the lease duration.",
      );
    }
  }

  public async initialize(): Promise<void> {
    await this.containerClient.createIfNotExists();
  }

  public async uploadFile(
    localFilePath: string,
    blobName: string,
    options: UploadFileOptions = {},
  ): Promise<UploadFileResult> {
    const file = await stat(localFilePath);
    if (!file.isFile()) {
      throw new Error(`Upload source is not a file: ${localFilePath}`);
    }

    const blobClient = this.containerClient.getBlockBlobClient(blobName);
    if (await this.blobExists(blobClient)) {
      return this.uploadExistingBlobWithLease(
        blobClient,
        localFilePath,
        options,
      );
    }

    const response = await blobClient.uploadStream(
      createReadStream(localFilePath),
      this.uploadBufferSize,
      this.uploadConcurrency,
      {
        ...this.uploadOptions(options),
        conditions: { ifNoneMatch: "*" },
      },
    );

    return {
      ...(response.etag ? { etag: response.etag } : {}),
      ...(response.lastModified
        ? { lastModified: response.lastModified }
        : {}),
      leaseProtected: false,
    };
  }

  public async downloadFile(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    const blobClient = this.containerClient.getBlobClient(blobName);
    const response = await blobClient.download();
    if (!response.readableStreamBody) {
      throw new Error(`Blob download returned no content: ${blobName}`);
    }

    const temporaryPath = `${destinationPath}.${randomUUID()}.part`;
    try {
      await pipeline(
        response.readableStreamBody,
        createWriteStream(temporaryPath, { flags: "wx" }),
      );
      await rm(destinationPath, { force: true });
      await rename(temporaryPath, destinationPath);
    } catch (error) {
      await rm(temporaryPath, { force: true });
      throw error;
    }
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

  private async blobExists(blobClient: BlockBlobClient): Promise<boolean> {
    try {
      await blobClient.getProperties();
      return true;
    } catch (error) {
      if (hasStatusCode(error, 404)) {
        return false;
      }
      throw error;
    }
  }

  private uploadOptions(options: UploadFileOptions) {
    return {
      ...(options.metadata ? { metadata: options.metadata } : {}),
      ...(options.tags ? { tags: options.tags } : {}),
      ...(options.contentType
        ? { blobHTTPHeaders: { blobContentType: options.contentType } }
        : {}),
    };
  }

  private async uploadExistingBlobWithLease(
    blobClient: BlockBlobClient,
    localFilePath: string,
    options: UploadFileOptions,
  ): Promise<UploadFileResult> {
    const leaseClient = blobClient.getBlobLeaseClient();
    const lease = await leaseClient.acquireLease(this.leaseDurationSeconds);
    if (!lease.leaseId) {
      await leaseClient.releaseLease();
      throw new Error("Azure did not return an ID for the acquired blob lease.");
    }
    const uploadAbortController = new AbortController();
    const renewalStopController = new AbortController();
    let renewalFailure: unknown;
    let operationFailure: unknown;

    const renewalTask = this.renewLeaseUntilStopped(
      leaseClient,
      renewalStopController.signal,
    ).catch((error: unknown) => {
      renewalFailure = error;
      uploadAbortController.abort(error);
    });

    let response:
      | Awaited<ReturnType<BlockBlobClient["uploadStream"]>>
      | undefined;

    try {
      response = await blobClient.uploadStream(
        createReadStream(localFilePath),
        this.uploadBufferSize,
        this.uploadConcurrency,
        {
          ...this.uploadOptions(options),
          abortSignal: uploadAbortController.signal,
          conditions: { leaseId: lease.leaseId },
        },
      );
      if (renewalFailure) {
        throw renewalFailure;
      }
    } catch (error) {
      operationFailure = renewalFailure ?? error;
    } finally {
      renewalStopController.abort();
      await renewalTask;
      try {
        await leaseClient.releaseLease();
      } catch (releaseError) {
        operationFailure = operationFailure
          ? new AggregateError(
              [operationFailure, releaseError],
              "Blob upload and lease release both failed.",
            )
          : releaseError;
      }
    }

    if (operationFailure) {
      throw operationFailure;
    }
    if (!response) {
      throw new Error("Blob upload completed without a response.");
    }

    return {
      ...(response.etag ? { etag: response.etag } : {}),
      ...(response.lastModified
        ? { lastModified: response.lastModified }
        : {}),
      leaseProtected: true,
    };
  }

  private async renewLeaseUntilStopped(
    leaseClient: BlobLeaseClient,
    stopSignal: AbortSignal,
  ): Promise<void> {
    while (!stopSignal.aborted) {
      await new Promise<void>((resolve) => {
        const onAbort = (): void => {
          clearTimeout(timeout);
          resolve();
        };
        const timeout = setTimeout(() => {
          stopSignal.removeEventListener("abort", onAbort);
          resolve();
        }, this.leaseRenewalIntervalMs);
        stopSignal.addEventListener("abort", onAbort, { once: true });
      });

      if (!stopSignal.aborted) {
        await leaseClient.renewLease();
      }
    }
  }
}
