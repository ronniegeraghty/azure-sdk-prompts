import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import type { ContainerClient, Metadata } from "@azure/storage-blob";
import {
  KEY_WRAP_ALGORITHM,
  type KeyManagement,
  type ProtectedDataKey,
} from "./keyManagement.js";

const CONTENT_ENCRYPTION_ALGORITHM = "A256GCM";
const ENVELOPE_VERSION = "1";
const IV_LENGTH_BYTES = 12;
const AUTH_TAG_LENGTH_BYTES = 16;

const METADATA = {
  version: "ceversion",
  algorithm: "cealgorithm",
  wrapAlgorithm: "cewrapalgorithm",
  keyId: "cekeyid",
  wrappedKey: "cewrappedkey",
  iv: "ceiv",
  authenticationTag: "ceauthtag",
} as const;

export interface EncryptedUploadResult {
  blobName: string;
  keyId: string;
  wrappedKeyBase64: string;
}

export class BlobNotFoundError extends Error {
  public constructor(blobName: string, cause: unknown) {
    super(`Encrypted blob "${blobName}" does not exist.`, { cause });
    this.name = "BlobNotFoundError";
  }
}

export class BlobStorageError extends Error {
  public constructor(operation: string, blobName: string, cause: unknown) {
    super(
      `Azure Blob Storage ${operation} failed for "${blobName}"${formatAzureErrorDetails(cause)}.`,
      { cause },
    );
    this.name = "BlobStorageError";
  }
}

export class EncryptedBlobFormatError extends Error {
  public constructor(message: string, cause?: unknown) {
    super(message, cause === undefined ? undefined : { cause });
    this.name = "EncryptedBlobFormatError";
  }
}

export class EncryptedBlobStorage {
  public constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  public async upload(
    blobName: string,
    plaintext: Uint8Array,
    contentType = "application/octet-stream",
  ): Promise<EncryptedUploadResult> {
    const blockBlobClient =
      this.containerClient.getBlockBlobClient(blobName);

    return this.keyManagement.withNewDataKey(
      async (dataKey, protectedDataKey) => {
        const iv = randomBytes(IV_LENGTH_BYTES);
        const cipher = createCipheriv("aes-256-gcm", dataKey, iv, {
          authTagLength: AUTH_TAG_LENGTH_BYTES,
        });
        const ciphertext = Buffer.concat([
          cipher.update(plaintext),
          cipher.final(),
        ]);
        const authenticationTag = cipher.getAuthTag();
        const metadata = this.createMetadata(
          protectedDataKey,
          iv,
          authenticationTag,
        );

        try {
          await blockBlobClient.upload(ciphertext, ciphertext.length, {
            metadata,
            blobHTTPHeaders: {
              blobContentType: contentType,
            },
          });
        } catch (error) {
          throw new BlobStorageError("upload", blobName, error);
        }

        return {
          blobName,
          keyId: protectedDataKey.keyId,
          wrappedKeyBase64: Buffer.from(
            protectedDataKey.wrappedKey,
          ).toString("base64"),
        };
      },
    );
  }

  public async download(blobName: string): Promise<Buffer> {
    const blockBlobClient =
      this.containerClient.getBlockBlobClient(blobName);

    let download;
    try {
      download = await blockBlobClient.download();
    } catch (error) {
      if (getStatusCode(error) === 404) {
        throw new BlobNotFoundError(blobName, error);
      }
      throw new BlobStorageError("download", blobName, error);
    }

    if (!download.readableStreamBody) {
      throw new BlobStorageError(
        "download",
        blobName,
        new Error("Blob Storage returned no response body."),
      );
    }

    const envelope = this.readMetadata(download.metadata);
    let ciphertext: Buffer;
    try {
      ciphertext = await streamToBuffer(download.readableStreamBody);
    } catch (error) {
      throw new BlobStorageError("read response body", blobName, error);
    }

    return this.keyManagement.withUnwrappedDataKey(
      envelope.protectedDataKey,
      (dataKey) => {
        try {
          const decipher = createDecipheriv(
            "aes-256-gcm",
            dataKey,
            envelope.iv,
            { authTagLength: AUTH_TAG_LENGTH_BYTES },
          );
          decipher.setAuthTag(envelope.authenticationTag);
          return Buffer.concat([
            decipher.update(ciphertext),
            decipher.final(),
          ]);
        } catch (error) {
          throw new EncryptedBlobFormatError(
            "Blob authentication failed; the ciphertext or encryption metadata may have been modified.",
            error,
          );
        }
      },
    );
  }

  private createMetadata(
    protectedDataKey: ProtectedDataKey,
    iv: Uint8Array,
    authenticationTag: Uint8Array,
  ): Metadata {
    return {
      [METADATA.version]: ENVELOPE_VERSION,
      [METADATA.algorithm]: CONTENT_ENCRYPTION_ALGORITHM,
      [METADATA.wrapAlgorithm]: protectedDataKey.wrapAlgorithm,
      [METADATA.keyId]: protectedDataKey.keyId,
      [METADATA.wrappedKey]: Buffer.from(
        protectedDataKey.wrappedKey,
      ).toString("base64"),
      [METADATA.iv]: Buffer.from(iv).toString("base64"),
      [METADATA.authenticationTag]:
        Buffer.from(authenticationTag).toString("base64"),
    };
  }

  private readMetadata(metadata: Metadata | undefined): {
    protectedDataKey: ProtectedDataKey;
    iv: Buffer;
    authenticationTag: Buffer;
  } {
    if (!metadata) {
      throw new EncryptedBlobFormatError(
        "Blob does not contain encryption metadata.",
      );
    }

    const version = requireMetadata(metadata, METADATA.version);
    const algorithm = requireMetadata(metadata, METADATA.algorithm);
    const wrapAlgorithm = requireMetadata(
      metadata,
      METADATA.wrapAlgorithm,
    );

    if (version !== ENVELOPE_VERSION) {
      throw new EncryptedBlobFormatError(
        `Unsupported envelope version "${version}".`,
      );
    }
    if (algorithm !== CONTENT_ENCRYPTION_ALGORITHM) {
      throw new EncryptedBlobFormatError(
        `Unsupported content encryption algorithm "${algorithm}".`,
      );
    }
    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new EncryptedBlobFormatError(
        `Unsupported key wrap algorithm "${wrapAlgorithm}".`,
      );
    }

    const iv = decodeBase64Metadata(
      metadata,
      METADATA.iv,
      IV_LENGTH_BYTES,
    );
    const authenticationTag = decodeBase64Metadata(
      metadata,
      METADATA.authenticationTag,
      AUTH_TAG_LENGTH_BYTES,
    );
    const wrappedKey = decodeBase64Metadata(
      metadata,
      METADATA.wrappedKey,
    );

    return {
      protectedDataKey: {
        keyId: requireMetadata(metadata, METADATA.keyId),
        wrappedKey,
        wrapAlgorithm: KEY_WRAP_ALGORITHM,
      },
      iv,
      authenticationTag,
    };
  }
}

function requireMetadata(metadata: Metadata, name: string): string {
  const value = metadata[name];
  if (!value) {
    throw new EncryptedBlobFormatError(
      `Blob encryption metadata "${name}" is missing.`,
    );
  }
  return value;
}

function decodeBase64Metadata(
  metadata: Metadata,
  name: string,
  expectedLength?: number,
): Buffer {
  const value = requireMetadata(metadata, name);
  if (
    value.length % 4 !== 0 ||
    !/^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$/.test(
      value,
    )
  ) {
    throw new EncryptedBlobFormatError(
      `Blob encryption metadata "${name}" is not valid base64.`,
    );
  }

  const decoded = Buffer.from(value, "base64");
  if (expectedLength !== undefined && decoded.length !== expectedLength) {
    throw new EncryptedBlobFormatError(
      `Blob encryption metadata "${name}" has ${decoded.length} bytes; expected ${expectedLength}.`,
    );
  }
  if (decoded.length === 0) {
    throw new EncryptedBlobFormatError(
      `Blob encryption metadata "${name}" must not be empty.`,
    );
  }
  return decoded;
}

function streamToBuffer(stream: NodeJS.ReadableStream): Promise<Buffer> {
  return new Promise((resolve, reject) => {
    const chunks: Buffer[] = [];
    stream.on("data", (chunk: Buffer | Uint8Array | string) => {
      chunks.push(Buffer.from(chunk));
    });
    stream.once("end", () => resolve(Buffer.concat(chunks)));
    stream.once("error", reject);
  });
}

function getStatusCode(error: unknown): number | undefined {
  if (!error || typeof error !== "object") {
    return undefined;
  }
  const statusCode = (error as { statusCode?: unknown }).statusCode;
  return typeof statusCode === "number" ? statusCode : undefined;
}

function formatAzureErrorDetails(error: unknown): string {
  if (!error || typeof error !== "object") {
    return "";
  }

  const candidate = error as {
    code?: unknown;
    statusCode?: unknown;
  };
  const details: string[] = [];

  if (typeof candidate.code === "string") {
    details.push(`code ${candidate.code}`);
  }
  if (typeof candidate.statusCode === "number") {
    details.push(`HTTP ${candidate.statusCode}`);
  }

  return details.length > 0 ? ` (${details.join(", ")})` : "";
}
