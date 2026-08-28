import assert from "node:assert/strict";
import { randomBytes } from "node:crypto";
import { Readable } from "node:stream";
import test from "node:test";
import type { Metadata } from "@azure/storage-blob";
import {
  BlobTransferError,
  EncryptedBlobClient,
  EncryptedBlobMetadataError,
} from "../src/encryptedBlobClient.js";
import {
  KEY_WRAP_ALGORITHM,
  type ProtectedDataKey,
} from "../src/keyManagement.js";

class FakeBlockBlob {
  public ciphertext?: Buffer;
  public metadata?: Metadata;

  public async uploadData(
    data: Buffer,
    options: {
      blobHTTPHeaders: { blobContentType: string };
      metadata: Metadata;
    },
  ): Promise<void> {
    this.ciphertext = Buffer.from(data);
    this.metadata = { ...options.metadata };
  }

  public async download(): Promise<{
    metadata?: Metadata;
    readableStreamBody?: NodeJS.ReadableStream;
  }> {
    if (!this.ciphertext) {
      throw Object.assign(new Error("The specified blob does not exist."), {
        code: "BlobNotFound",
        statusCode: 404,
      });
    }

    return {
      metadata: this.metadata ? { ...this.metadata } : undefined,
      readableStreamBody: Readable.from([this.ciphertext]),
    };
  }
}

class FakeContainer {
  private readonly blobs = new Map<string, FakeBlockBlob>();

  public getBlockBlobClient(blobName: string): FakeBlockBlob {
    let blob = this.blobs.get(blobName);
    if (!blob) {
      blob = new FakeBlockBlob();
      this.blobs.set(blobName, blob);
    }
    return blob;
  }
}

class FakeKeyManagement {
  public async createProtectedDataKey(): Promise<ProtectedDataKey> {
    const dataKey = randomBytes(32);
    return {
      dataKey,
      keyId:
        "https://example.vault.azure.net/keys/blob-encryption-key/version-1",
      wrappedKey: this.transform(dataKey),
      wrapAlgorithm: KEY_WRAP_ALGORITHM,
    };
  }

  public async recoverDataKey(
    _keyId: string,
    wrappedKey: Buffer,
    _wrapAlgorithm: string,
  ): Promise<Buffer> {
    return this.transform(wrappedKey);
  }

  private transform(input: Buffer): Buffer {
    return Buffer.from(input, (_value, index) => input[index]! ^ 0xa5);
  }
}

test("encrypts, stores required metadata, and decrypts a round-trip", async () => {
  const container = new FakeContainer();
  const client = new EncryptedBlobClient(
    container,
    new FakeKeyManagement(),
  );
  const plaintext = Buffer.from("confidential sample", "utf8");

  const result = await client.upload("sample.txt", plaintext, "text/plain");
  const storedBlob = container.getBlockBlobClient("sample.txt");
  const decrypted = await client.download("sample.txt");

  assert.deepEqual(decrypted, plaintext);
  assert.notDeepEqual(storedBlob.ciphertext, plaintext);
  assert.equal(storedBlob.metadata?.keyid, result.keyId);
  assert.equal(
    storedBlob.metadata?.wrappeddatakey,
    result.wrappedDataKeyBase64,
  );
  assert.equal(storedBlob.metadata?.contentencryptionalgorithm, "AES-256-GCM");
  assert.ok(storedBlob.metadata?.initializationvector);
  assert.ok(storedBlob.metadata?.authenticationtag);
});

test("rejects a modified AES-GCM authentication tag", async () => {
  const container = new FakeContainer();
  const client = new EncryptedBlobClient(
    container,
    new FakeKeyManagement(),
  );
  await client.upload("tampered.txt", Buffer.from("protected"));

  const blob = container.getBlockBlobClient("tampered.txt");
  assert.ok(blob.metadata);
  blob.metadata.authenticationtag = Buffer.alloc(16, 0).toString("base64");

  await assert.rejects(
    client.download("tampered.txt"),
    EncryptedBlobMetadataError,
  );
});

test("reports a missing blob as a storage error", async () => {
  const client = new EncryptedBlobClient(
    new FakeContainer(),
    new FakeKeyManagement(),
  );

  await assert.rejects(
    client.download("missing.txt"),
    (error: unknown) =>
      error instanceof BlobTransferError &&
      error.message === 'Encrypted blob "missing.txt" was not found.',
  );
});
