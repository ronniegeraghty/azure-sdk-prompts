import {
  createCipheriv,
  createDecipheriv,
  randomBytes,
} from "node:crypto";
import type {
  BlobDownloadResponseParsed,
  ContainerClient,
  Metadata,
} from "@azure/storage-blob";
import {
  KeyManagement,
  KEY_WRAP_ALGORITHM,
  type WrappedDataKey,
} from "./keyManagement.js";

const CONTENT_ALGORITHM = "AES-256-GCM";
const FORMAT_VERSION = "1";
const IV_BYTES = 12;
const AUTH_TAG_BYTES = 16;

interface EncryptionMetadata extends Record<string, string> {
  encryptionversion: string;
  contentalgorithm: string;
  wrapalgorithm: string;
  keyid: string;
  wrappedkey: string;
  iv: string;
  authtag: string;
}

export interface EncryptedUploadResult {
  blobName: string;
  keyId: string;
  wrappedKeyBase64: string;
  eTag?: string;
}

export class EncryptedBlobError extends Error {
  constructor(
    message: string,
    public readonly operation: "upload" | "download" | "metadata" | "decrypt",
    options?: ErrorOptions,
  ) {
    super(message, options);
    this.name = "EncryptedBlobError";
  }
}

function encodeMetadata(
  wrappedDataKey: WrappedDataKey,
  iv: Buffer,
  authenticationTag: Buffer,
): EncryptionMetadata {
  return {
    encryptionversion: FORMAT_VERSION,
    contentalgorithm: CONTENT_ALGORITHM,
    wrapalgorithm: wrappedDataKey.algorithm,
    keyid: wrappedDataKey.keyId,
    wrappedkey: Buffer.from(wrappedDataKey.wrappedKey).toString("base64"),
    iv: iv.toString("base64"),
    authtag: authenticationTag.toString("base64"),
  };
}

function requireMetadata(metadata: Metadata, name: keyof EncryptionMetadata): string {
  const value = metadata[name];
  if (!value) {
    throw new EncryptedBlobError(
      `Encrypted blob metadata is missing "${name}".`,
      "metadata",
    );
  }
  return value;
}

function decodeBase64Metadata(
  metadata: Metadata,
  name: keyof Pick<EncryptionMetadata, "wrappedkey" | "iv" | "authtag">,
): Buffer {
  const encoded = requireMetadata(metadata, name);
  const decoded = Buffer.from(encoded, "base64");

  if (decoded.length === 0 || decoded.toString("base64") !== encoded) {
    throw new EncryptedBlobError(
      `Encrypted blob metadata "${name}" is not canonical Base64.`,
      "metadata",
    );
  }

  return decoded;
}

async function streamToBuffer(
  stream: NodeJS.ReadableStream | undefined,
): Promise<Buffer> {
  if (!stream) {
    throw new EncryptedBlobError(
      "Blob Storage returned a download response without a body.",
      "download",
    );
  }

  return new Promise((resolve, reject) => {
    const chunks: Buffer[] = [];
    stream.on("data", (chunk: Buffer | Uint8Array | string) => {
      chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
    });
    stream.on("end", () => resolve(Buffer.concat(chunks)));
    stream.on("error", reject);
  });
}

function parseMetadata(response: BlobDownloadResponseParsed): {
  wrappedDataKey: WrappedDataKey;
  iv: Buffer;
  authenticationTag: Buffer;
} {
  const metadata = response.metadata ?? {};
  const version = requireMetadata(metadata, "encryptionversion");
  const contentAlgorithm = requireMetadata(metadata, "contentalgorithm");
  const wrapAlgorithm = requireMetadata(metadata, "wrapalgorithm");

  if (version !== FORMAT_VERSION) {
    throw new EncryptedBlobError(
      `Unsupported encrypted blob format version "${version}".`,
      "metadata",
    );
  }
  if (contentAlgorithm !== CONTENT_ALGORITHM) {
    throw new EncryptedBlobError(
      `Unsupported content encryption algorithm "${contentAlgorithm}".`,
      "metadata",
    );
  }
  if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
    throw new EncryptedBlobError(
      `Unsupported key wrap algorithm "${wrapAlgorithm}".`,
      "metadata",
    );
  }

  const iv = decodeBase64Metadata(metadata, "iv");
  const authenticationTag = decodeBase64Metadata(metadata, "authtag");
  if (iv.length !== IV_BYTES) {
    throw new EncryptedBlobError(
      `Expected a ${IV_BYTES}-byte initialization vector, received ${iv.length} bytes.`,
      "metadata",
    );
  }
  if (authenticationTag.length !== AUTH_TAG_BYTES) {
    throw new EncryptedBlobError(
      `Expected a ${AUTH_TAG_BYTES}-byte authentication tag, received ${authenticationTag.length} bytes.`,
      "metadata",
    );
  }

  return {
    wrappedDataKey: {
      keyId: requireMetadata(metadata, "keyid"),
      wrappedKey: decodeBase64Metadata(metadata, "wrappedkey"),
      algorithm: KEY_WRAP_ALGORITHM,
    },
    iv,
    authenticationTag,
  };
}

export class EncryptedBlobClient {
  constructor(
    private readonly containerClient: ContainerClient,
    private readonly keyManagement: KeyManagement,
  ) {}

  async upload(
    blobName: string,
    plaintext: Buffer | Uint8Array | string,
  ): Promise<EncryptedUploadResult> {
    const plaintextBuffer =
      typeof plaintext === "string" ? Buffer.from(plaintext, "utf8") : Buffer.from(plaintext);

    return this.keyManagement.withGeneratedDataKey(
      async (dataKey, wrappedDataKey) => {
        const iv = randomBytes(IV_BYTES);
        const cipher = createCipheriv("aes-256-gcm", dataKey, iv, {
          authTagLength: AUTH_TAG_BYTES,
        });
        const ciphertext = Buffer.concat([
          cipher.update(plaintextBuffer),
          cipher.final(),
        ]);
        const authenticationTag = cipher.getAuthTag();
        const metadata = encodeMetadata(wrappedDataKey, iv, authenticationTag);

        try {
          const response = await this.containerClient
            .getBlockBlobClient(blobName)
            .uploadData(ciphertext, {
              metadata,
              blobHTTPHeaders: {
                blobContentType: "application/octet-stream",
              },
            });

          return {
            blobName,
            keyId: wrappedDataKey.keyId,
            wrappedKeyBase64: metadata.wrappedkey,
            ...(response.etag ? { eTag: response.etag } : {}),
          };
        } catch (error) {
          throw new EncryptedBlobError(
            `Failed to upload encrypted blob "${blobName}". The container may not exist or the managed identity may lack Blob Data Contributor access.`,
            "upload",
            { cause: error },
          );
        }
      },
    );
  }

  async download(blobName: string): Promise<Buffer> {
    let response: BlobDownloadResponseParsed;
    try {
      response = await this.containerClient.getBlobClient(blobName).download();
    } catch (error) {
      throw new EncryptedBlobError(
        `Failed to download encrypted blob "${blobName}". It may not exist or the managed identity may lack Blob Data Reader access.`,
        "download",
        { cause: error },
      );
    }

    let ciphertext: Buffer;
    try {
      ciphertext = await streamToBuffer(response.readableStreamBody);
    } catch (error) {
      if (error instanceof EncryptedBlobError) {
        throw error;
      }
      throw new EncryptedBlobError(
        `The response stream for encrypted blob "${blobName}" failed while downloading.`,
        "download",
        { cause: error },
      );
    }
    const { wrappedDataKey, iv, authenticationTag } = parseMetadata(response);

    return this.keyManagement.withUnwrappedDataKey(
      wrappedDataKey,
      async (dataKey) => {
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
          throw new EncryptedBlobError(
            `Failed to authenticate or decrypt blob "${blobName}". Its ciphertext or cryptographic metadata may have been modified.`,
            "decrypt",
            { cause: error },
          );
        }
      },
    );
  }
}
