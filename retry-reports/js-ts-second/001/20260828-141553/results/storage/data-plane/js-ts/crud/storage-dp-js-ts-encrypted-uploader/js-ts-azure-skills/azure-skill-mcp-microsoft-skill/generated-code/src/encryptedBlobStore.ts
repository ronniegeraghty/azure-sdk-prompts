import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import type { ContainerClient } from "@azure/storage-blob";
import { EncryptedBlobError } from "./errors.js";
import {
  KEY_WRAP_ALGORITHM,
  type KeyManagement,
  type ProtectedDataKey,
} from "./keyManagement.js";

const ENCRYPTION_VERSION = "1";
const CONTENT_ALGORITHM = "AES-256-GCM";
const IV_BYTES = 12;
const AUTH_TAG_BYTES = 16;

const METADATA = {
  version: "encryptionversion",
  contentAlgorithm: "contentalgorithm",
  wrapAlgorithm: "wrapalgorithm",
  keyId: "keyid",
  wrappedKey: "wrappedkey",
  iv: "iv",
  authenticationTag: "authenticationtag",
} as const;

export interface EncryptedUploadResult {
  blobUrl: string;
  keyId: string;
  wrappedKeyBase64: string;
}

interface EncryptedPayload {
  ciphertext: Buffer;
  iv: Buffer;
  authenticationTag: Buffer;
  protectedKey: ProtectedDataKey;
}

export class EncryptedBlobStore {
  constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  async upload(
    blobName: string,
    plaintext: Buffer,
  ): Promise<EncryptedUploadResult> {
    const encrypted = await this.keyManagement.withNewDataKey(
      async (dataKey, protectedKey): Promise<EncryptedPayload> => {
        const iv = randomBytes(IV_BYTES);
        const cipher = createCipheriv("aes-256-gcm", dataKey, iv, {
          authTagLength: AUTH_TAG_BYTES,
        });
        const ciphertext = Buffer.concat([
          cipher.update(plaintext),
          cipher.final(),
        ]);

        return {
          ciphertext,
          iv,
          authenticationTag: cipher.getAuthTag(),
          protectedKey,
        };
      },
    );

    const wrappedKeyBase64 = encrypted.protectedKey.wrappedKey.toString("base64");
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);

    try {
      await blockBlobClient.upload(
        encrypted.ciphertext,
        encrypted.ciphertext.length,
        {
          blobHTTPHeaders: {
            blobContentType: "application/octet-stream",
          },
          metadata: {
            [METADATA.version]: ENCRYPTION_VERSION,
            [METADATA.contentAlgorithm]: CONTENT_ALGORITHM,
            [METADATA.wrapAlgorithm]: encrypted.protectedKey.wrapAlgorithm,
            [METADATA.keyId]: encrypted.protectedKey.keyId,
            [METADATA.wrappedKey]: wrappedKeyBase64,
            [METADATA.iv]: encrypted.iv.toString("base64"),
            [METADATA.authenticationTag]:
              encrypted.authenticationTag.toString("base64"),
          },
        },
      );
    } catch (error) {
      throw new EncryptedBlobError(`upload of "${blobName}"`, error);
    }

    return {
      blobUrl: blockBlobClient.url,
      keyId: encrypted.protectedKey.keyId,
      wrappedKeyBase64,
    };
  }

  async download(blobName: string): Promise<Buffer> {
    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);

    let ciphertext: Buffer;
    let metadata: Record<string, string>;
    try {
      const response = await blockBlobClient.download();
      if (!response.readableStreamBody) {
        throw new Error("The blob download response did not contain a body.");
      }

      ciphertext = await streamToBuffer(response.readableStreamBody);
      metadata = requireEncryptionMetadata(response.metadata);
    } catch (error) {
      if (error instanceof EncryptedBlobError) {
        throw error;
      }
      throw new EncryptedBlobError(`download of "${blobName}"`, error);
    }

    const protectedKey: ProtectedDataKey = {
      keyId: metadata[METADATA.keyId]!,
      wrappedKey: decodeBase64Metadata(
        METADATA.wrappedKey,
        metadata[METADATA.wrappedKey]!,
      ),
      wrapAlgorithm: KEY_WRAP_ALGORITHM,
    };
    const iv = decodeBase64Metadata(METADATA.iv, metadata[METADATA.iv]!);
    const authenticationTag = decodeBase64Metadata(
      METADATA.authenticationTag,
      metadata[METADATA.authenticationTag]!,
    );

    if (iv.length !== IV_BYTES) {
      throw new EncryptedBlobError(
        `decryption of "${blobName}"`,
        new Error(`Invalid AES-GCM initialization vector length: ${iv.length}.`),
      );
    }
    if (authenticationTag.length !== AUTH_TAG_BYTES) {
      throw new EncryptedBlobError(
        `decryption of "${blobName}"`,
        new Error(
          `Invalid AES-GCM authentication tag length: ${authenticationTag.length}.`,
        ),
      );
    }

    try {
      return await this.keyManagement.withRecoveredDataKey(
        protectedKey,
        async (dataKey) => {
          const decipher = createDecipheriv("aes-256-gcm", dataKey, iv, {
            authTagLength: AUTH_TAG_BYTES,
          });
          decipher.setAuthTag(authenticationTag);
          return Buffer.concat([
            decipher.update(ciphertext),
            decipher.final(),
          ]);
        },
      );
    } catch (error) {
      throw new EncryptedBlobError(`decryption of "${blobName}"`, error);
    }
  }
}

function requireEncryptionMetadata(
  metadata: Record<string, string> | undefined,
): Record<string, string> {
  if (!metadata) {
    throw new Error("The blob has no encryption metadata.");
  }

  for (const metadataName of Object.values(METADATA)) {
    if (!metadata[metadataName]) {
      throw new Error(
        `The blob is missing encryption metadata "${metadataName}".`,
      );
    }
  }

  if (metadata[METADATA.version] !== ENCRYPTION_VERSION) {
    throw new Error(
      `Unsupported encryption metadata version: ${metadata[METADATA.version]}.`,
    );
  }
  if (metadata[METADATA.contentAlgorithm] !== CONTENT_ALGORITHM) {
    throw new Error(
      `Unsupported content encryption algorithm: ${metadata[METADATA.contentAlgorithm]}.`,
    );
  }
  if (metadata[METADATA.wrapAlgorithm] !== KEY_WRAP_ALGORITHM) {
    throw new Error(
      `Unsupported key wrap algorithm: ${metadata[METADATA.wrapAlgorithm]}.`,
    );
  }

  return metadata;
}

function decodeBase64Metadata(name: string, value: string): Buffer {
  const decoded = Buffer.from(value, "base64");
  if (decoded.length === 0 || decoded.toString("base64") !== value) {
    throw new EncryptedBlobError(
      "metadata validation",
      new Error(`Blob metadata "${name}" is not valid canonical base64.`),
    );
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
