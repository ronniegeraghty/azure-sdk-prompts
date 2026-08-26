import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import { readFile, writeFile } from "node:fs/promises";

import type { ContainerClient, Metadata } from "@azure/storage-blob";

import { AzureOperationError } from "./errors.js";
import { KeyManagement } from "./key-management.js";

const CONTENT_ENCRYPTION_ALGORITHM = "AES-256-GCM";
const CONTENT_ENCRYPTION_VERSION = "1";
const IV_LENGTH_BYTES = 12;
const AUTH_TAG_LENGTH_BYTES = 16;

const METADATA = {
  version: "encryptionversion",
  contentAlgorithm: "contentencryptionalgorithm",
  wrapAlgorithm: "keywrapalgorithm",
  keyId: "keyid",
  wrappedKey: "wrappedkey",
  iv: "iv",
  authenticationTag: "authenticationtag",
} as const;

export interface UploadResult {
  blobUrl: string;
  keyId: string;
  wrappedKeyBase64: string;
  etag?: string;
}

export class EncryptedBlobStorage {
  public constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  public async uploadBuffer(
    blobName: string,
    plaintext: Buffer,
    contentType = "application/octet-stream",
  ): Promise<UploadResult> {
    const dataKey = this.keyManagement.generateDataKey();

    try {
      const iv = randomBytes(IV_LENGTH_BYTES);
      const cipher = createCipheriv("aes-256-gcm", dataKey, iv, {
        authTagLength: AUTH_TAG_LENGTH_BYTES,
      });
      const ciphertext = Buffer.concat([
        cipher.update(plaintext),
        cipher.final(),
      ]);
      const authenticationTag = cipher.getAuthTag();
      const protectedKey = await this.keyManagement.protectDataKey(dataKey);
      const wrappedKeyBase64 = protectedKey.wrappedKey.toString("base64");

      const metadata: Metadata = {
        [METADATA.version]: CONTENT_ENCRYPTION_VERSION,
        [METADATA.contentAlgorithm]: CONTENT_ENCRYPTION_ALGORITHM,
        [METADATA.wrapAlgorithm]: protectedKey.wrapAlgorithm,
        [METADATA.keyId]: Buffer.from(protectedKey.keyId, "utf8").toString(
          "base64",
        ),
        [METADATA.wrappedKey]: wrappedKeyBase64,
        [METADATA.iv]: iv.toString("base64"),
        [METADATA.authenticationTag]: authenticationTag.toString("base64"),
      };

      try {
        const blobClient = this.containerClient.getBlockBlobClient(blobName);
        const response = await blobClient.upload(
          ciphertext,
          ciphertext.length,
          {
            metadata,
            blobHTTPHeaders: {
              blobContentType: contentType,
            },
          },
        );

        return {
          blobUrl: blobClient.url,
          keyId: protectedKey.keyId,
          wrappedKeyBase64,
          ...(response.etag ? { etag: response.etag } : {}),
        };
      } catch (error) {
        throw new AzureOperationError(
          "Azure Blob Storage",
          `uploading blob "${blobName}"`,
          error,
        );
      }
    } finally {
      dataKey.fill(0);
    }
  }

  public async uploadFile(
    blobName: string,
    filePath: string,
    contentType = "application/octet-stream",
  ): Promise<UploadResult> {
    const plaintext = await readFile(filePath);
    try {
      return await this.uploadBuffer(blobName, plaintext, contentType);
    } finally {
      plaintext.fill(0);
    }
  }

  public async downloadBuffer(blobName: string): Promise<Buffer> {
    let ciphertext: Buffer;
    let metadata: Metadata;

    try {
      const blobClient = this.containerClient.getBlockBlobClient(blobName);
      const response = await blobClient.download();
      if (!response.readableStreamBody) {
        throw new Error("The blob download returned no response body.");
      }

      ciphertext = await streamToBuffer(response.readableStreamBody);
      metadata = response.metadata ?? {};
    } catch (error) {
      const statusCode = getStatusCode(error);
      throw new AzureOperationError(
        "Azure Blob Storage",
        `downloading blob "${blobName}"`,
        error,
        statusCode === 404
          ? `Blob "${blobName}" does not exist.`
          : undefined,
      );
    }

    const envelope = parseEnvelopeMetadata(metadata);
    const dataKey = await this.keyManagement.recoverDataKey(
      envelope.keyId,
      envelope.wrappedKey,
      envelope.wrapAlgorithm,
    );

    try {
      const decipher = createDecipheriv("aes-256-gcm", dataKey, envelope.iv, {
        authTagLength: AUTH_TAG_LENGTH_BYTES,
      });
      decipher.setAuthTag(envelope.authenticationTag);

      try {
        return Buffer.concat([
          decipher.update(ciphertext),
          decipher.final(),
        ]);
      } catch (error) {
        throw new Error(
          `Blob "${blobName}" failed authentication and cannot be decrypted.`,
          { cause: error },
        );
      }
    } finally {
      dataKey.fill(0);
    }
  }

  public async downloadToFile(
    blobName: string,
    destinationPath: string,
  ): Promise<void> {
    const plaintext = await this.downloadBuffer(blobName);
    try {
      await writeFile(destinationPath, plaintext);
    } finally {
      plaintext.fill(0);
    }
  }
}

interface EnvelopeMetadata {
  keyId: string;
  wrappedKey: Buffer;
  wrapAlgorithm: string;
  iv: Buffer;
  authenticationTag: Buffer;
}

function parseEnvelopeMetadata(metadata: Metadata): EnvelopeMetadata {
  const version = requireMetadata(metadata, METADATA.version);
  const contentAlgorithm = requireMetadata(
    metadata,
    METADATA.contentAlgorithm,
  );

  if (version !== CONTENT_ENCRYPTION_VERSION) {
    throw new Error(`Unsupported encryption metadata version: ${version}`);
  }
  if (contentAlgorithm !== CONTENT_ENCRYPTION_ALGORITHM) {
    throw new Error(
      `Unsupported content-encryption algorithm: ${contentAlgorithm}`,
    );
  }

  const iv = decodeBase64Metadata(metadata, METADATA.iv);
  const authenticationTag = decodeBase64Metadata(
    metadata,
    METADATA.authenticationTag,
  );
  if (iv.length !== IV_LENGTH_BYTES) {
    throw new Error(`Invalid AES-GCM IV length: ${iv.length} bytes.`);
  }
  if (authenticationTag.length !== AUTH_TAG_LENGTH_BYTES) {
    throw new Error(
      `Invalid AES-GCM authentication tag length: ${authenticationTag.length} bytes.`,
    );
  }

  return {
    keyId: decodeBase64Metadata(metadata, METADATA.keyId).toString("utf8"),
    wrappedKey: decodeBase64Metadata(metadata, METADATA.wrappedKey),
    wrapAlgorithm: requireMetadata(metadata, METADATA.wrapAlgorithm),
    iv,
    authenticationTag,
  };
}

function requireMetadata(metadata: Metadata, name: string): string {
  const value = metadata[name];
  if (!value) {
    throw new Error(`Blob is missing required encryption metadata "${name}".`);
  }
  return value;
}

function decodeBase64Metadata(metadata: Metadata, name: string): Buffer {
  const value = requireMetadata(metadata, name);
  const decoded = Buffer.from(value, "base64");
  if (decoded.length === 0 && value.length > 0) {
    throw new Error(`Blob metadata "${name}" is not valid base64.`);
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

function getStatusCode(error: unknown): number | undefined {
  if (typeof error !== "object" || error === null) {
    return undefined;
  }
  const statusCode = (error as Record<string, unknown>)["statusCode"];
  return typeof statusCode === "number" ? statusCode : undefined;
}
