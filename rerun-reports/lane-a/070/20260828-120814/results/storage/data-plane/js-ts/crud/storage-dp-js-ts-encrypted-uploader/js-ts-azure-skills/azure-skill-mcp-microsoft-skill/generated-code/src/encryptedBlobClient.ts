import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import { readFile, writeFile } from "node:fs/promises";
import type {
  Metadata,
} from "@azure/storage-blob";
import {
  KEY_WRAP_ALGORITHM,
  KeyManagement,
} from "./keyManagement.js";

const CONTENT_ENCRYPTION_ALGORITHM = "AES-256-GCM";
const ENCRYPTION_FORMAT_VERSION = "1";
const IV_BYTES = 12;
const AUTH_TAG_BYTES = 16;

const METADATA = {
  authenticationTag: "authenticationtag",
  contentEncryptionAlgorithm: "contentencryptionalgorithm",
  encryptionFormatVersion: "encryptionformatversion",
  initializationVector: "initializationvector",
  keyId: "keyid",
  keyWrapAlgorithm: "keywrapalgorithm",
  wrappedDataKey: "wrappeddatakey",
} as const;

export interface UploadEncryptionResult {
  readonly keyId: string;
  readonly wrappedDataKeyBase64: string;
}

interface EncryptedBlockBlob {
  uploadData(
    data: Buffer,
    options: {
      blobHTTPHeaders: { blobContentType: string };
      metadata: Metadata;
    },
  ): Promise<unknown>;
  download(): Promise<{
    readonly metadata?: Metadata;
    readonly readableStreamBody?: NodeJS.ReadableStream;
  }>;
}

interface EncryptedBlobContainer {
  getBlockBlobClient(blobName: string): EncryptedBlockBlob;
}

type DataKeyProtector = Pick<
  KeyManagement,
  "createProtectedDataKey" | "recoverDataKey"
>;

interface AzureErrorDetails {
  readonly code?: string;
  readonly message?: string;
  readonly statusCode?: number;
}

function getAzureErrorDetails(error: unknown): AzureErrorDetails {
  if (typeof error !== "object" || error === null) {
    return {};
  }

  const candidate = error as Record<string, unknown>;
  return {
    ...(typeof candidate.code === "string" ? { code: candidate.code } : {}),
    ...(typeof candidate.message === "string"
      ? { message: candidate.message }
      : {}),
    ...(typeof candidate.statusCode === "number"
      ? { statusCode: candidate.statusCode }
      : {}),
  };
}

export class BlobTransferError extends Error {
  public constructor(operation: string, blobName: string, cause: unknown) {
    const details = getAzureErrorDetails(cause);
    const notFound =
      details.statusCode === 404 ||
      details.code === "BlobNotFound" ||
      details.code === "ContainerNotFound";
    const context = [
      details.statusCode === undefined
        ? undefined
        : `status ${details.statusCode}`,
      details.code,
    ]
      .filter((value): value is string => value !== undefined)
      .join(", ");
    const reason = details.message ? `: ${details.message}` : "";

    super(
      notFound
        ? `Encrypted blob "${blobName}" was not found.`
        : `Azure Blob Storage ${operation} failed for "${blobName}"${context ? ` (${context})` : ""}${reason}`,
      { cause },
    );
    this.name = "BlobTransferError";
  }
}

export class EncryptedBlobMetadataError extends Error {
  public constructor(message: string, options?: ErrorOptions) {
    super(`Invalid encrypted blob metadata: ${message}`, options);
    this.name = "EncryptedBlobMetadataError";
  }
}

function requireMetadata(metadata: Metadata, name: string): string {
  const value = metadata[name];
  if (!value) {
    throw new EncryptedBlobMetadataError(`missing "${name}".`);
  }
  return value;
}

function decodeBase64Metadata(
  metadata: Metadata,
  name: string,
  expectedLength?: number,
): Buffer {
  const encoded = requireMetadata(metadata, name);
  if (
    encoded.length % 4 !== 0 ||
    !/^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$/.test(
      encoded,
    )
  ) {
    throw new EncryptedBlobMetadataError(`"${name}" is not valid base64.`);
  }

  const decoded = Buffer.from(encoded, "base64");
  if (expectedLength !== undefined && decoded.length !== expectedLength) {
    throw new EncryptedBlobMetadataError(
      `"${name}" must decode to ${expectedLength} bytes.`,
    );
  }
  if (decoded.length === 0) {
    throw new EncryptedBlobMetadataError(`"${name}" must not be empty.`);
  }

  return decoded;
}

async function streamToBuffer(
  stream: NodeJS.ReadableStream,
): Promise<Buffer> {
  const chunks: Buffer[] = [];
  for await (const chunk of stream) {
    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
  }
  return Buffer.concat(chunks);
}

export class EncryptedBlobClient {
  public constructor(
    private readonly containerClient: EncryptedBlobContainer,
    private readonly keyManagement: DataKeyProtector,
  ) {}

  public async upload(
    blobName: string,
    plaintext: Buffer,
    contentType = "application/octet-stream",
  ): Promise<UploadEncryptionResult> {
    const protectedDataKey =
      await this.keyManagement.createProtectedDataKey();

    try {
      const iv = randomBytes(IV_BYTES);
      const cipher = createCipheriv(
        "aes-256-gcm",
        protectedDataKey.dataKey,
        iv,
        { authTagLength: AUTH_TAG_BYTES },
      );
      const ciphertext = Buffer.concat([
        cipher.update(plaintext),
        cipher.final(),
      ]);
      const authenticationTag = cipher.getAuthTag();
      const wrappedDataKeyBase64 =
        protectedDataKey.wrappedKey.toString("base64");
      const metadata: Metadata = {
        [METADATA.authenticationTag]:
          authenticationTag.toString("base64"),
        [METADATA.contentEncryptionAlgorithm]:
          CONTENT_ENCRYPTION_ALGORITHM,
        [METADATA.encryptionFormatVersion]: ENCRYPTION_FORMAT_VERSION,
        [METADATA.initializationVector]: iv.toString("base64"),
        [METADATA.keyId]: protectedDataKey.keyId,
        [METADATA.keyWrapAlgorithm]: protectedDataKey.wrapAlgorithm,
        [METADATA.wrappedDataKey]: wrappedDataKeyBase64,
      };

      await this.getBlobClient(blobName).uploadData(ciphertext, {
        blobHTTPHeaders: {
          blobContentType: contentType,
        },
        metadata,
      });

      return {
        keyId: protectedDataKey.keyId,
        wrappedDataKeyBase64,
      };
    } catch (error) {
      if (
        error instanceof EncryptedBlobMetadataError ||
        error instanceof BlobTransferError
      ) {
        throw error;
      }
      throw new BlobTransferError("upload", blobName, error);
    } finally {
      protectedDataKey.dataKey.fill(0);
    }
  }

  public async download(blobName: string): Promise<Buffer> {
    let ciphertext: Buffer;
    let metadata: Metadata;

    try {
      const response = await this.getBlobClient(blobName).download();
      if (!response.readableStreamBody) {
        throw new Error("Blob download returned no response body.");
      }

      ciphertext = await streamToBuffer(response.readableStreamBody);
      metadata = response.metadata ?? {};
    } catch (error) {
      if (error instanceof BlobTransferError) {
        throw error;
      }
      throw new BlobTransferError("download", blobName, error);
    }

    this.validateEncryptionMetadata(metadata);
    const keyId = requireMetadata(metadata, METADATA.keyId);
    const wrapAlgorithm = requireMetadata(
      metadata,
      METADATA.keyWrapAlgorithm,
    );
    const wrappedDataKey = decodeBase64Metadata(
      metadata,
      METADATA.wrappedDataKey,
    );
    const iv = decodeBase64Metadata(
      metadata,
      METADATA.initializationVector,
      IV_BYTES,
    );
    const authenticationTag = decodeBase64Metadata(
      metadata,
      METADATA.authenticationTag,
      AUTH_TAG_BYTES,
    );
    const dataKey = await this.keyManagement.recoverDataKey(
      keyId,
      wrappedDataKey,
      wrapAlgorithm,
    );

    try {
      const decipher = createDecipheriv("aes-256-gcm", dataKey, iv, {
        authTagLength: AUTH_TAG_BYTES,
      });
      decipher.setAuthTag(authenticationTag);
      return Buffer.concat([
        decipher.update(ciphertext),
        decipher.final(),
      ]);
    } catch (error) {
      throw new EncryptedBlobMetadataError(
        "AES-GCM authentication failed; the ciphertext, key, IV, or authentication tag may have been altered.",
        { cause: error },
      );
    } finally {
      dataKey.fill(0);
    }
  }

  public async uploadFile(
    blobName: string,
    sourcePath: string,
    contentType = "application/octet-stream",
  ): Promise<UploadEncryptionResult> {
    const plaintext = await readFile(sourcePath);
    return this.upload(blobName, plaintext, contentType);
  }

  public async downloadToFile(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    const plaintext = await this.download(blobName);
    await writeFile(destinationPath, plaintext);
  }

  private getBlobClient(blobName: string): EncryptedBlockBlob {
    if (!blobName.trim()) {
      throw new EncryptedBlobMetadataError("blob name must not be empty.");
    }
    return this.containerClient.getBlockBlobClient(blobName);
  }

  private validateEncryptionMetadata(metadata: Metadata): void {
    const formatVersion = requireMetadata(
      metadata,
      METADATA.encryptionFormatVersion,
    );
    if (formatVersion !== ENCRYPTION_FORMAT_VERSION) {
      throw new EncryptedBlobMetadataError(
        `unsupported encryption format version "${formatVersion}".`,
      );
    }

    const contentAlgorithm = requireMetadata(
      metadata,
      METADATA.contentEncryptionAlgorithm,
    );
    if (contentAlgorithm !== CONTENT_ENCRYPTION_ALGORITHM) {
      throw new EncryptedBlobMetadataError(
        `unsupported content encryption algorithm "${contentAlgorithm}".`,
      );
    }

    const wrapAlgorithm = requireMetadata(
      metadata,
      METADATA.keyWrapAlgorithm,
    );
    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new EncryptedBlobMetadataError(
        `unsupported key wrap algorithm "${wrapAlgorithm}".`,
      );
    }
  }
}
