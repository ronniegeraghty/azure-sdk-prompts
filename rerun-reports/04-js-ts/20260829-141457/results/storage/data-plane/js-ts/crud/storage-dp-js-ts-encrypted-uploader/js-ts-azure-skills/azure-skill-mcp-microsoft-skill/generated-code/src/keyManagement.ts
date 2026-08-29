import { randomBytes } from "node:crypto";
import type { TokenCredential } from "@azure/core-auth";
import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";
import {
  describeAzureFailure,
  EncryptedBlobError,
} from "./errors.js";

const DATA_KEY_LENGTH_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  readonly keyId: string;
  readonly wrappedKey: Buffer;
  readonly wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

export class KeyVaultKeyManager {
  public constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: TokenCredential,
    private readonly keyName: string,
  ) {}

  public generateDataKey(): Buffer {
    return randomBytes(DATA_KEY_LENGTH_BYTES);
  }

  public async protectDataKey(dataKey: Buffer): Promise<ProtectedDataKey> {
    if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {
      throw new EncryptedBlobError(
        "cryptography",
        "wrap data key",
        `The data encryption key must be ${DATA_KEY_LENGTH_BYTES} bytes.`,
      );
    }

    try {
      const key = await this.keyClient.getKey(this.keyName);
      if (!key.id) {
        throw new Error("Key Vault returned a key without a versioned key ID.");
      }

      const cryptographyClient = new CryptographyClient(
        key.id,
        this.credential,
      );
      const result = await cryptographyClient.wrapKey(
        KEY_WRAP_ALGORITHM,
        dataKey,
      );

      return {
        keyId: key.id,
        wrappedKey: Buffer.from(result.result),
        wrapAlgorithm: KEY_WRAP_ALGORITHM,
      };
    } catch (error) {
      throw new EncryptedBlobError(
        "key-vault",
        "wrap data key",
        `Azure Key Vault could not protect the data key: ${describeAzureFailure(error)}`,
        { cause: error },
      );
    }
  }

  public async recoverDataKey(
    keyId: string,
    wrappedKey: Buffer,
    wrapAlgorithm: string,
  ): Promise<Buffer> {
    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new EncryptedBlobError(
        "cryptography",
        "unwrap data key",
        `Unsupported key wrap algorithm: ${wrapAlgorithm}.`,
      );
    }

    this.validateKeyId(keyId);

    try {
      const cryptographyClient = new CryptographyClient(
        keyId,
        this.credential,
      );
      const result = await cryptographyClient.unwrapKey(
        KEY_WRAP_ALGORITHM,
        wrappedKey,
      );
      const dataKey = Buffer.from(result.result);

      if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {
        dataKey.fill(0);
        throw new Error(
          `Key Vault returned an invalid ${dataKey.length}-byte data key.`,
        );
      }

      return dataKey;
    } catch (error) {
      throw new EncryptedBlobError(
        "key-vault",
        "unwrap data key",
        `Azure Key Vault could not recover the data key: ${describeAzureFailure(error)}`,
        { cause: error },
      );
    }
  }

  private validateKeyId(keyId: string): void {
    let parsed: URL;

    try {
      parsed = new URL(keyId);
    } catch (error) {
      throw new EncryptedBlobError(
        "cryptography",
        "validate key ID",
        "Blob metadata contains an invalid Key Vault key ID.",
        { cause: error },
      );
    }

    if (
      parsed.protocol !== "https:" ||
      !/^\/keys\/[^/]+\/[^/]+$/.test(parsed.pathname)
    ) {
      throw new EncryptedBlobError(
        "cryptography",
        "validate key ID",
        "Blob metadata must reference a versioned HTTPS Key Vault key ID.",
      );
    }
  }
}
