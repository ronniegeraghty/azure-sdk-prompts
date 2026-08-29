import {
  createCipheriv,
  createDecipheriv,
  type CipherGCM,
  type DecipherGCM,
  randomBytes,
} from "node:crypto";
import { buffer as streamToBuffer } from "node:stream/consumers";

import type { ContainerClient, Metadata } from "@azure/storage-blob";

import {
  BlobOperationError,
  EncryptionMetadataError,
  getErrorMessage,
  getStatusCode,
} from "./errors.js";
import {
  KeyManagementClient,
  KEY_WRAP_ALGORITHM,
} from "./key-management.js";

const CONTENT_ALGORITHM = "AES-256-GCM";
const ENCRYPTION_VERSION = "1";
const IV_BYTES = 12;
const AUTH_TAG_BYTES = 16;

const METADATA = {
  aad: "aad",
  authenticationTag: "authtag",
  contentAlgorithm: "contentalgorithm",
  encryptionVersion: "encryptionversion",
  initializationVector: "iv",
  keyId: "keyid",
  keyWrapAlgorithm: "keywrapalgorithm",
  wrappedKey: "wrappedkey",
} as const;

export interface EncryptedUploadResult {
  keyId: string;
  wrappedKeyBase64: string;
}

interface EncryptionMetadata {
  authenticationTag: Buffer;
  initializationVector: Buffer;
  keyId: string;
  wrappedKey: Buffer;
}

export class EncryptedBlobClient {
  readonly #containerClient: ContainerClient;
  readonly #keyManagementClient: KeyManagementClient;

  public constructor(
    containerClient: ContainerClient,
    keyManagementClient: KeyManagementClient,
  ) {
    this.#containerClient = containerClient;
    this.#keyManagementClient = keyManagementClient;
  }

  public async upload(
    blobName: string,
    plaintext: Buffer | string,
  ): Promise<EncryptedUploadResult> {
    const data = typeof plaintext === "string" ? Buffer.from(plaintext) : plaintext;
    const dataKey = await this.#keyManagementClient.generateAndWrapDataKey();

    try {
      const initializationVector = randomBytes(IV_BYTES);
      const aad = this.#createAdditionalAuthenticatedData(blobName);
      const cipher = createCipheriv(
        "aes-256-gcm",
        dataKey.plaintextKey,
        initializationVector,
        { authTagLength: AUTH_TAG_BYTES },
      ) as CipherGCM;
      cipher.setAAD(aad);

      const ciphertext = Buffer.concat([cipher.update(data), cipher.final()]);
      const authenticationTag = cipher.getAuthTag();
      const wrappedKeyBase64 = dataKey.wrappedKey.toString("base64");
      const metadata: Metadata = {
        [METADATA.aad]: "blob-path-v1",
        [METADATA.authenticationTag]: authenticationTag.toString("base64"),
        [METADATA.contentAlgorithm]: CONTENT_ALGORITHM,
        [METADATA.encryptionVersion]: ENCRYPTION_VERSION,
        [METADATA.initializationVector]:
          initializationVector.toString("base64"),
        [METADATA.keyId]: dataKey.keyId,
        [METADATA.keyWrapAlgorithm]: KEY_WRAP_ALGORITHM,
        [METADATA.wrappedKey]: wrappedKeyBase64,
      };

      try {
        await this.#containerClient
          .getBlockBlobClient(blobName)
          .uploadData(ciphertext, {
            blobHTTPHeaders: {
              blobContentType: "application/octet-stream",
            },
            metadata,
          });
      } catch (error) {
        throw new BlobOperationError(
          "upload",
          getErrorMessage(error),
          { cause: error },
        );
      }

      return {
        keyId: dataKey.keyId,
        wrappedKeyBase64,
      };
    } finally {
      dataKey.plaintextKey.fill(0);
    }
  }

  public async download(blobName: string): Promise<Buffer> {
    let ciphertext: Buffer;
    let metadata: Metadata | undefined;

    try {
      const response = await this.#containerClient
        .getBlobClient(blobName)
        .download();
      metadata = response.metadata;

      if (!response.readableStreamBody) {
        throw new Error("The Blob service returned no response body.");
      }

      ciphertext = await streamToBuffer(response.readableStreamBody);
    } catch (error) {
      const detail =
        getStatusCode(error) === 404
          ? `blob "${blobName}" does not exist.`
          : getErrorMessage(error);
      throw new BlobOperationError("download", detail, { cause: error });
    }

    const parameters = this.#parseMetadata(metadata);
    const dataKey = await this.#keyManagementClient.unwrapDataKey(
      parameters.wrappedKey,
      parameters.keyId,
    );

    try {
      const decipher = createDecipheriv(
        "aes-256-gcm",
        dataKey,
        parameters.initializationVector,
        { authTagLength: AUTH_TAG_BYTES },
      ) as DecipherGCM;
      decipher.setAAD(this.#createAdditionalAuthenticatedData(blobName));
      decipher.setAuthTag(parameters.authenticationTag);

      return Buffer.concat([
        decipher.update(ciphertext),
        decipher.final(),
      ]);
    } catch (error) {
      throw new BlobOperationError(
        "decryption",
        `authentication failed or the ciphertext is corrupt: ${getErrorMessage(error)}`,
        { cause: error },
      );
    } finally {
      dataKey.fill(0);
    }
  }

  #createAdditionalAuthenticatedData(blobName: string): Buffer {
    return Buffer.from(
      `azure-blob:${this.#containerClient.containerName}/${blobName}:v1`,
      "utf8",
    );
  }

  #parseMetadata(metadata: Metadata | undefined): EncryptionMetadata {
    if (!metadata) {
      throw new EncryptionMetadataError("metadata is missing.");
    }

    this.#expectMetadata(metadata, METADATA.encryptionVersion, ENCRYPTION_VERSION);
    this.#expectMetadata(metadata, METADATA.contentAlgorithm, CONTENT_ALGORITHM);
    this.#expectMetadata(metadata, METADATA.keyWrapAlgorithm, KEY_WRAP_ALGORITHM);
    this.#expectMetadata(metadata, METADATA.aad, "blob-path-v1");

    const keyId = this.#requiredMetadata(metadata, METADATA.keyId);
    const wrappedKey = this.#decodeBase64(
      metadata,
      METADATA.wrappedKey,
    );
    const initializationVector = this.#decodeBase64(
      metadata,
      METADATA.initializationVector,
      IV_BYTES,
    );
    const authenticationTag = this.#decodeBase64(
      metadata,
      METADATA.authenticationTag,
      AUTH_TAG_BYTES,
    );

    return {
      authenticationTag,
      initializationVector,
      keyId,
      wrappedKey,
    };
  }

  #expectMetadata(
    metadata: Metadata,
    name: string,
    expected: string,
  ): void {
    const actual = this.#requiredMetadata(metadata, name);
    if (actual !== expected) {
      throw new EncryptionMetadataError(
        `${name} is "${actual}"; expected "${expected}".`,
      );
    }
  }

  #requiredMetadata(metadata: Metadata, name: string): string {
    const value = metadata[name];
    if (!value) {
      throw new EncryptionMetadataError(`${name} is missing.`);
    }

    return value;
  }

  #decodeBase64(
    metadata: Metadata,
    name: string,
    expectedLength?: number,
  ): Buffer {
    const encoded = this.#requiredMetadata(metadata, name);
    if (
      !/^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$/.test(
        encoded,
      )
    ) {
      throw new EncryptionMetadataError(`${name} is not valid base64.`);
    }

    const decoded = Buffer.from(encoded, "base64");
    if (decoded.length === 0) {
      throw new EncryptionMetadataError(`${name} decoded to an empty value.`);
    }
    if (expectedLength !== undefined && decoded.length !== expectedLength) {
      throw new EncryptionMetadataError(
        `${name} decoded to ${decoded.length} bytes; expected ${expectedLength}.`,
      );
    }

    return decoded;
  }
}
