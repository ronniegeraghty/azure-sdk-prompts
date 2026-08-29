import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import { readFile, writeFile } from "node:fs/promises";
import type { ContainerClient } from "@azure/storage-blob";
import {
  describeAzureFailure,
  EncryptedBlobError,
} from "./errors.js";
import {
  KEY_WRAP_ALGORITHM,
  type KeyVaultKeyManager,
} from "./keyManagement.js";

const ENCRYPTION_ALGORITHM = "aes-256-gcm";
const ENCRYPTION_VERSION = "1";
const IV_LENGTH_BYTES = 12;
const AUTH_TAG_LENGTH_BYTES = 16;

const METADATA = {
  version: "encryptionversion",
  algorithm: "encryptionalgorithm",
  keyId: "keyid",
  wrapAlgorithm: "keywrapalgorithm",
  wrappedKey: "wrappeddek",
  iv: "iv",
  authenticationTag: "authenticationtag",
} as const;

export interface UploadResult {
  readonly keyId: string;
  readonly wrappedDataKeyBase64: string;
  readonly blobUrl: string;
}

interface EncryptionMetadata {
  readonly keyId: string;
  readonly wrappedKey: Buffer;
  readonly iv: Buffer;
  readonly authenticationTag: Buffer;
}

export class EncryptedBlobStorage {
  public constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManager: KeyVaultKeyManager,
  ) {}

  public async upload(
    blobName: string,
    plaintext: Buffer | string,
  ): Promise<UploadResult> {
    const data = typeof plaintext === "string"
      ? Buffer.from(plaintext, "utf8")
      : plaintext;
    const dataKey = this.keyManager.generateDataKey();

    try {
      const iv = randomBytes(IV_LENGTH_BYTES);
      const cipher = createCipheriv(ENCRYPTION_ALGORITHM, dataKey, iv, {
        authTagLength: AUTH_TAG_LENGTH_BYTES,
      });
      const ciphertext = Buffer.concat([cipher.update(data), cipher.final()]);
      const authenticationTag = cipher.getAuthTag();
      const protectedKey = await this.keyManager.protectDataKey(dataKey);
      const wrappedDataKeyBase64 = protectedKey.wrappedKey.toString("base64");
      const blockBlobClient =
        this.containerClient.getBlockBlobClient(blobName);

      try {
        await blockBlobClient.uploadData(ciphertext, {
          blobHTTPHeaders: {
            blobContentType: "application/octet-stream",
          },
          metadata: {
            [METADATA.version]: ENCRYPTION_VERSION,
            [METADATA.algorithm]: ENCRYPTION_ALGORITHM,
            [METADATA.keyId]: protectedKey.keyId,
            [METADATA.wrapAlgorithm]: protectedKey.wrapAlgorithm,
            [METADATA.wrappedKey]: wrappedDataKeyBase64,
            [METADATA.iv]: iv.toString("base64"),
            [METADATA.authenticationTag]:
              authenticationTag.toString("base64"),
          },
        });
      } catch (error) {
        throw new EncryptedBlobError(
          "storage",
          "upload encrypted blob",
          `Azure Blob Storage upload failed: ${describeAzureFailure(error)}`,
          { cause: error },
        );
      }

      return {
        keyId: protectedKey.keyId,
        wrappedDataKeyBase64,
        blobUrl: blockBlobClient.url,
      };
    } catch (error) {
      if (error instanceof EncryptedBlobError) {
        throw error;
      }

      throw new EncryptedBlobError(
        "cryptography",
        "encrypt blob",
        error instanceof Error ? error.message : "Local encryption failed.",
        { cause: error },
      );
    } finally {
      dataKey.fill(0);
    }
  }

  public async uploadFile(
    blobName: string,
    localPath: string,
  ): Promise<UploadResult> {
    return this.upload(blobName, await readFile(localPath));
  }

  public async download(blobName: string): Promise<Buffer> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);
    let ciphertext: Buffer;
    let metadata: Record<string, string>;

    try {
      const response = await blockBlobClient.download();
      if (!response.readableStreamBody) {
        throw new Error("Blob download returned no readable response body.");
      }

      ciphertext = await this.streamToBuffer(response.readableStreamBody);
      metadata = response.metadata ?? {};
    } catch (error) {
      throw new EncryptedBlobError(
        "storage",
        "download encrypted blob",
        `Azure Blob Storage download failed: ${describeAzureFailure(error)}`,
        { cause: error },
      );
    }

    const encryptionMetadata = this.parseMetadata(metadata);
    const dataKey = await this.keyManager.recoverDataKey(
      encryptionMetadata.keyId,
      encryptionMetadata.wrappedKey,
      KEY_WRAP_ALGORITHM,
    );

    try {
      const decipher = createDecipheriv(
        ENCRYPTION_ALGORITHM,
        dataKey,
        encryptionMetadata.iv,
        { authTagLength: AUTH_TAG_LENGTH_BYTES },
      );
      decipher.setAuthTag(encryptionMetadata.authenticationTag);
      return Buffer.concat([
        decipher.update(ciphertext),
        decipher.final(),
      ]);
    } catch (error) {
      throw new EncryptedBlobError(
        "cryptography",
        "decrypt blob",
        "Local AES-GCM decryption failed. The ciphertext or encryption metadata may have been altered.",
        { cause: error },
      );
    } finally {
      dataKey.fill(0);
    }
  }

  public async downloadToFile(
    blobName: string,
    localPath: string,
  ): Promise<void> {
    await writeFile(localPath, await this.download(blobName));
  }

  private parseMetadata(
    metadata: Record<string, string>,
  ): EncryptionMetadata {
    const required = (name: string): string => {
      const value = metadata[name];
      if (!value) {
        throw new EncryptedBlobError(
          "cryptography",
          "read encryption metadata",
          `Blob metadata is missing ${name}.`,
        );
      }

      return value;
    };

    const version = required(METADATA.version);
    const algorithm = required(METADATA.algorithm);
    const wrapAlgorithm = required(METADATA.wrapAlgorithm);

    if (version !== ENCRYPTION_VERSION) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `Unsupported encryption metadata version: ${version}.`,
      );
    }

    if (algorithm !== ENCRYPTION_ALGORITHM) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `Unsupported content encryption algorithm: ${algorithm}.`,
      );
    }

    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `Unsupported key wrap algorithm: ${wrapAlgorithm}.`,
      );
    }

    const wrappedKey = this.decodeBase64(
      required(METADATA.wrappedKey),
      METADATA.wrappedKey,
    );
    const iv = this.decodeBase64(required(METADATA.iv), METADATA.iv);
    const authenticationTag = this.decodeBase64(
      required(METADATA.authenticationTag),
      METADATA.authenticationTag,
    );

    if (wrappedKey.length === 0) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        "The wrapped data key is empty.",
      );
    }

    if (iv.length !== IV_LENGTH_BYTES) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `The AES-GCM initialization vector must be ${IV_LENGTH_BYTES} bytes.`,
      );
    }

    if (authenticationTag.length !== AUTH_TAG_LENGTH_BYTES) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `The AES-GCM authentication tag must be ${AUTH_TAG_LENGTH_BYTES} bytes.`,
      );
    }

    return {
      keyId: required(METADATA.keyId),
      wrappedKey,
      iv,
      authenticationTag,
    };
  }

  private decodeBase64(value: string, fieldName: string): Buffer {
    if (
      value.length % 4 !== 0 ||
      !/^[A-Za-z0-9+/]*={0,2}$/.test(value)
    ) {
      throw new EncryptedBlobError(
        "cryptography",
        "read encryption metadata",
        `Blob metadata field ${fieldName} is not valid base64.`,
      );
    }

    return Buffer.from(value, "base64");
  }

  private async streamToBuffer(
    readable: NodeJS.ReadableStream,
  ): Promise<Buffer> {
    const chunks: Buffer[] = [];
    for await (const chunk of readable) {
      chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
    }

    return Buffer.concat(chunks);
  }
}
