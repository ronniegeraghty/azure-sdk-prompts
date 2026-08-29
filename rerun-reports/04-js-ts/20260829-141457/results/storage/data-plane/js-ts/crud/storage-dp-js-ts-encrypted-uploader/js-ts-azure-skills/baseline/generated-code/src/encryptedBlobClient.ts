import type { ContainerClient, Metadata } from "@azure/storage-blob";
import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import {
  KeyManagement,
  KeyVaultOperationError,
  type ProtectedDataKey,
} from "./keyManagement.js";

const CONTENT_ENCRYPTION_ALGORITHM = "aes-256-gcm";
const IV_LENGTH_BYTES = 12;
const AUTH_TAG_LENGTH_BYTES = 16;

interface EncryptionMetadata {
  keyId: string;
  wrappedDataKey: Buffer;
  wrappingAlgorithm: string;
  iv: Buffer;
  authenticationTag: Buffer;
}

export interface UploadResult {
  keyId: string;
  wrappedDataKeyBase64: string;
}

export class BlobStorageOperationError extends Error {
  constructor(
    message: string,
    options: ErrorOptions,
  ) {
    super(message, options);
    this.name = "BlobStorageOperationError";
  }
}

export class EncryptedBlobClient {
  constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  async upload(blobName: string, plaintext: Buffer): Promise<UploadResult> {
    let protectedKey: ProtectedDataKey | undefined;

    try {
      protectedKey = await this.keyManagement.generateAndProtectDataKey();
      const iv = randomBytes(IV_LENGTH_BYTES);
      const cipher = createCipheriv(
        CONTENT_ENCRYPTION_ALGORITHM,
        protectedKey.dataKey,
        iv,
        { authTagLength: AUTH_TAG_LENGTH_BYTES },
      );
      const ciphertext = Buffer.concat([
        cipher.update(plaintext),
        cipher.final(),
      ]);
      const authenticationTag = cipher.getAuthTag();

      const wrappedDataKeyBase64 = Buffer.from(
        protectedKey.wrappedDataKey,
      ).toString("base64");
      const metadata: Metadata = {
        contentencryptionalgorithm: CONTENT_ENCRYPTION_ALGORITHM,
        keywrappingalgorithm: protectedKey.wrappingAlgorithm,
        keyid: protectedKey.keyId,
        wrappeddek: wrappedDataKeyBase64,
        iv: iv.toString("base64"),
        authenticationtag: authenticationTag.toString("base64"),
      };

      await this.containerClient
        .getBlockBlobClient(blobName)
        .uploadData(ciphertext, {
          metadata,
          blobHTTPHeaders: {
            blobContentType: "application/octet-stream",
          },
        });

      return {
        keyId: protectedKey.keyId,
        wrappedDataKeyBase64,
      };
    } catch (error) {
      if (error instanceof KeyVaultOperationError) {
        throw error;
      }
      throw new BlobStorageOperationError(
        `Failed to encrypt and upload blob "${blobName}"`,
        { cause: error },
      );
    } finally {
      protectedKey?.dataKey.fill(0);
    }
  }

  async download(blobName: string): Promise<Buffer> {
    let dataKey: Buffer | undefined;

    try {
      const response = await this.containerClient
        .getBlobClient(blobName)
        .download();
      if (!response.readableStreamBody) {
        throw new Error(`Blob "${blobName}" returned no content`);
      }

      const ciphertext = await streamToBuffer(response.readableStreamBody);
      const metadata = parseEncryptionMetadata(response.metadata);
      dataKey = await this.keyManagement.recoverDataKey(
        metadata.wrappedDataKey,
        metadata.keyId,
        metadata.wrappingAlgorithm,
      );

      const decipher = createDecipheriv(
        CONTENT_ENCRYPTION_ALGORITHM,
        dataKey,
        metadata.iv,
        { authTagLength: AUTH_TAG_LENGTH_BYTES },
      );
      decipher.setAuthTag(metadata.authenticationTag);
      return Buffer.concat([decipher.update(ciphertext), decipher.final()]);
    } catch (error) {
      if (error instanceof KeyVaultOperationError) {
        throw error;
      }
      throw new BlobStorageOperationError(
        `Failed to download and decrypt blob "${blobName}"`,
        { cause: error },
      );
    } finally {
      dataKey?.fill(0);
    }
  }
}

function parseEncryptionMetadata(
  metadata: Metadata | undefined,
): EncryptionMetadata {
  const keyId = requiredMetadata(metadata, "keyid");
  const wrappingAlgorithm = requiredMetadata(
    metadata,
    "keywrappingalgorithm",
  );
  const contentEncryptionAlgorithm = requiredMetadata(
    metadata,
    "contentencryptionalgorithm",
  );
  if (contentEncryptionAlgorithm !== CONTENT_ENCRYPTION_ALGORITHM) {
    throw new Error(
      `Unsupported content encryption algorithm: ${contentEncryptionAlgorithm}`,
    );
  }

  const wrappedDataKey = decodeBase64(
    requiredMetadata(metadata, "wrappeddek"),
    "wrappeddek",
  );
  const iv = decodeBase64(requiredMetadata(metadata, "iv"), "iv");
  const authenticationTag = decodeBase64(
    requiredMetadata(metadata, "authenticationtag"),
    "authenticationtag",
  );

  if (iv.length !== IV_LENGTH_BYTES) {
    throw new Error(`Invalid IV length: ${iv.length}`);
  }
  if (authenticationTag.length !== AUTH_TAG_LENGTH_BYTES) {
    throw new Error(
      `Invalid authentication tag length: ${authenticationTag.length}`,
    );
  }

  return {
    keyId,
    wrappedDataKey,
    wrappingAlgorithm,
    iv,
    authenticationTag,
  };
}

function requiredMetadata(
  metadata: Metadata | undefined,
  name: string,
): string {
  const value = metadata?.[name];
  if (!value) {
    throw new Error(`Encrypted blob is missing metadata field "${name}"`);
  }
  return value;
}

function decodeBase64(value: string, fieldName: string): Buffer {
  const decoded = Buffer.from(value, "base64");
  if (decoded.length === 0 || decoded.toString("base64") !== value) {
    throw new Error(`Encrypted blob metadata field "${fieldName}" is invalid`);
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
