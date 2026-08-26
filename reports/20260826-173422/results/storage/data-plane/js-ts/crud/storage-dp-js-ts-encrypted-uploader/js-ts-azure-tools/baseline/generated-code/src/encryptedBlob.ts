import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";

import type {
  BlobDownloadResponseParsed,
  BlockBlobClient,
  ContainerClient,
} from "@azure/storage-blob";

import { KeyManagement } from "./keyManagement.js";

const METADATA_VERSION = "1";
const CONTENT_ENCRYPTION_ALGORITHM = "AES-256-GCM";
const IV_LENGTH_BYTES = 12;
const AUTH_TAG_LENGTH_BYTES = 16;

interface EncryptionMetadata {
  readonly version: string;
  readonly contentEncryptionAlgorithm: string;
  readonly keyWrapAlgorithm: string;
  readonly keyId: string;
  readonly wrappedKey: string;
  readonly iv: string;
  readonly authenticationTag: string;
}

export interface EncryptedUploadResult {
  readonly blobName: string;
  readonly keyId: string;
  readonly wrappedKeyBase64: string;
}

export class EncryptedBlobClient {
  public constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  public async upload(
    blobName: string,
    plaintext: Buffer,
  ): Promise<EncryptedUploadResult> {
    const envelopeKey = await this.keyManagement.generateAndWrapDataKey();
    const iv = randomBytes(IV_LENGTH_BYTES);

    try {
      const cipher = createCipheriv(
        "aes-256-gcm",
        envelopeKey.plaintextKey,
        iv,
        { authTagLength: AUTH_TAG_LENGTH_BYTES },
      );
      const ciphertext = Buffer.concat([
        cipher.update(plaintext),
        cipher.final(),
      ]);
      const authenticationTag = cipher.getAuthTag();
      const wrappedKeyBase64 = Buffer.from(envelopeKey.wrappedKey).toString(
        "base64",
      );

      const metadata: Record<string, string> = {
        encryptionversion: METADATA_VERSION,
        encryptionalgorithm: CONTENT_ENCRYPTION_ALGORITHM,
        keywrapalgorithm: envelopeKey.algorithm,
        keyid: envelopeKey.keyId,
        wrappedkey: wrappedKeyBase64,
        iv: iv.toString("base64"),
        authenticationtag: authenticationTag.toString("base64"),
      };

      try {
        await this.blockBlob(blobName).uploadData(ciphertext, {
          metadata,
          blobHTTPHeaders: {
            blobContentType: "application/octet-stream",
          },
        });
      } catch (error) {
        throw new Error(`Failed to upload encrypted blob "${blobName}"`, {
          cause: error,
        });
      }

      return {
        blobName,
        keyId: envelopeKey.keyId,
        wrappedKeyBase64,
      };
    } finally {
      envelopeKey.plaintextKey.fill(0);
    }
  }

  public async download(blobName: string): Promise<Buffer> {
    let response: BlobDownloadResponseParsed;
    try {
      response = await this.blockBlob(blobName).download();
    } catch (error) {
      throw new Error(`Failed to download encrypted blob "${blobName}"`, {
        cause: error,
      });
    }

    const metadata = this.parseMetadata(response.metadata, blobName);
    const wrappedKey = this.decodeBase64(
      metadata.wrappedKey,
      "wrapped key",
      blobName,
    );
    const iv = this.decodeBase64(metadata.iv, "initialization vector", blobName);
    const authenticationTag = this.decodeBase64(
      metadata.authenticationTag,
      "authentication tag",
      blobName,
    );

    if (iv.length !== IV_LENGTH_BYTES) {
      throw new Error(
        `Blob "${blobName}" has an invalid ${iv.length}-byte initialization vector`,
      );
    }
    if (authenticationTag.length !== AUTH_TAG_LENGTH_BYTES) {
      throw new Error(
        `Blob "${blobName}" has an invalid ${authenticationTag.length}-byte authentication tag`,
      );
    }

    const ciphertext = await this.readBody(response, blobName);
    const plaintextKey = await this.keyManagement.unwrapDataKey(
      metadata.keyId,
      wrappedKey,
      metadata.keyWrapAlgorithm,
    );

    try {
      const decipher = createDecipheriv("aes-256-gcm", plaintextKey, iv, {
        authTagLength: AUTH_TAG_LENGTH_BYTES,
      });
      decipher.setAuthTag(authenticationTag);
      return Buffer.concat([decipher.update(ciphertext), decipher.final()]);
    } catch (error) {
      throw new Error(
        `Failed to decrypt blob "${blobName}"; its ciphertext or encryption metadata may be invalid`,
        { cause: error },
      );
    } finally {
      plaintextKey.fill(0);
    }
  }

  private blockBlob(blobName: string): BlockBlobClient {
    return this.containerClient.getBlockBlobClient(blobName);
  }

  private parseMetadata(
    metadata: Record<string, string> | undefined,
    blobName: string,
  ): EncryptionMetadata {
    const required = (name: string): string => {
      const value = metadata?.[name];
      if (!value) {
        throw new Error(
          `Blob "${blobName}" is missing encryption metadata "${name}"`,
        );
      }
      return value;
    };

    const parsed: EncryptionMetadata = {
      version: required("encryptionversion"),
      contentEncryptionAlgorithm: required("encryptionalgorithm"),
      keyWrapAlgorithm: required("keywrapalgorithm"),
      keyId: required("keyid"),
      wrappedKey: required("wrappedkey"),
      iv: required("iv"),
      authenticationTag: required("authenticationtag"),
    };

    if (parsed.version !== METADATA_VERSION) {
      throw new Error(
        `Blob "${blobName}" uses unsupported encryption metadata version "${parsed.version}"`,
      );
    }
    if (
      parsed.contentEncryptionAlgorithm !== CONTENT_ENCRYPTION_ALGORITHM
    ) {
      throw new Error(
        `Blob "${blobName}" uses unsupported content-encryption algorithm "${parsed.contentEncryptionAlgorithm}"`,
      );
    }

    return parsed;
  }

  private decodeBase64(
    value: string,
    fieldName: string,
    blobName: string,
  ): Buffer {
    const decoded = Buffer.from(value, "base64");
    if (decoded.length === 0 || decoded.toString("base64") !== value) {
      throw new Error(
        `Blob "${blobName}" has invalid base64 in its ${fieldName} metadata`,
      );
    }
    return decoded;
  }

  private async readBody(
    response: BlobDownloadResponseParsed,
    blobName: string,
  ): Promise<Buffer> {
    if (!response.readableStreamBody) {
      throw new Error(`Blob "${blobName}" download returned no response body`);
    }

    const chunks: Buffer[] = [];
    try {
      for await (const chunk of response.readableStreamBody) {
        chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
      }
    } catch (error) {
      throw new Error(`Failed while reading encrypted blob "${blobName}"`, {
        cause: error,
      });
    }
    return Buffer.concat(chunks);
  }
}
