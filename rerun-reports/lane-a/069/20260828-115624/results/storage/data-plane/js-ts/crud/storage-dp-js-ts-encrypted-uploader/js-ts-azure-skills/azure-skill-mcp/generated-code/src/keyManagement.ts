import { randomBytes } from "node:crypto";
import type { TokenCredential } from "@azure/core-auth";
import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";

const DATA_KEY_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface WrappedDataKey {
  keyId: string;
  wrappedKey: Uint8Array;
  algorithm: typeof KEY_WRAP_ALGORITHM;
}

type DataKeyOperation<T> = (
  dataKey: Buffer,
  wrappedDataKey: WrappedDataKey,
) => Promise<T>;

export class KeyManagementError extends Error {
  constructor(
    message: string,
    public readonly operation: "resolve" | "wrap" | "unwrap",
    options?: ErrorOptions,
  ) {
    super(message, options);
    this.name = "KeyManagementError";
  }
}

export class KeyManagement {
  private constructor(
    private readonly cryptographyClient: CryptographyClient,
    private readonly credential: TokenCredential,
    public readonly keyId: string,
    private readonly keyIdPrefix: string,
  ) {}

  static async create(
    keyClient: KeyClient,
    credential: TokenCredential,
    keyName: string,
    keyVersion?: string,
  ): Promise<KeyManagement> {
    try {
      const key = await keyClient.getKey(
        keyName,
        keyVersion ? { version: keyVersion } : {},
      );

      if (!key.id) {
        throw new Error("Key Vault returned a key without an ID.");
      }

      const keyUrl = new URL(key.id);
      const pathParts = keyUrl.pathname.split("/").filter(Boolean);
      if (
        pathParts.length !== 3 ||
        pathParts[0] !== "keys" ||
        !pathParts[1] ||
        !pathParts[2]
      ) {
        throw new Error(`Key Vault returned an invalid versioned key ID: ${key.id}`);
      }

      const keyIdPrefix = `${keyUrl.origin}/keys/${pathParts[1]}/`;
      return new KeyManagement(
        new CryptographyClient(key.id, credential),
        credential,
        key.id,
        keyIdPrefix,
      );
    } catch (error) {
      throw new KeyManagementError(
        `Unable to resolve Key Vault key "${keyName}". The key may not exist, may be disabled, or the managed identity may lack key permissions.`,
        "resolve",
        { cause: error },
      );
    }
  }

  async withGeneratedDataKey<T>(operation: DataKeyOperation<T>): Promise<T> {
    const dataKey = randomBytes(DATA_KEY_BYTES);

    try {
      let wrappedKey: Uint8Array;
      try {
        const result = await this.cryptographyClient.wrapKey(
          KEY_WRAP_ALGORITHM,
          dataKey,
        );
        wrappedKey = result.result;
      } catch (error) {
        throw new KeyManagementError(
          "Key Vault could not wrap the data encryption key. The key may be disabled or the managed identity may lack wrapKey permission.",
          "wrap",
          { cause: error },
        );
      }

      return await operation(dataKey, {
        keyId: this.keyId,
        wrappedKey,
        algorithm: KEY_WRAP_ALGORITHM,
      });
    } finally {
      dataKey.fill(0);
    }
  }

  async withUnwrappedDataKey<T>(
    wrappedDataKey: WrappedDataKey,
    operation: DataKeyOperation<T>,
  ): Promise<T> {
    if (!this.isAllowedKeyVersion(wrappedDataKey.keyId)) {
      throw new KeyManagementError(
        `Blob metadata references key "${wrappedDataKey.keyId}", which is not a version of the configured Key Vault key.`,
        "unwrap",
      );
    }

    let dataKey: Buffer | undefined;
    try {
      try {
        const client =
          wrappedDataKey.keyId === this.keyId
            ? this.cryptographyClient
            : new CryptographyClient(wrappedDataKey.keyId, this.credential);
        const result = await client.unwrapKey(
          wrappedDataKey.algorithm,
          wrappedDataKey.wrappedKey,
        );
        dataKey = Buffer.from(result.result);
      } catch (error) {
        throw new KeyManagementError(
          "Key Vault could not unwrap the data encryption key. The key version may be disabled or the managed identity may lack unwrapKey permission.",
          "unwrap",
          { cause: error },
        );
      }

      if (dataKey.length !== DATA_KEY_BYTES) {
        throw new KeyManagementError(
          `Key Vault returned an invalid ${dataKey.length}-byte data encryption key.`,
          "unwrap",
        );
      }

      return await operation(dataKey, wrappedDataKey);
    } finally {
      dataKey?.fill(0);
    }
  }

  private isAllowedKeyVersion(keyId: string): boolean {
    try {
      const url = new URL(keyId);
      if (url.search || url.hash) {
        return false;
      }

      const prefix = `${url.origin}${url.pathname.slice(
        0,
        url.pathname.lastIndexOf("/") + 1,
      )}`;
      const version = url.pathname.slice(url.pathname.lastIndexOf("/") + 1);
      return prefix === this.keyIdPrefix && version.length > 0;
    } catch {
      return false;
    }
  }
}
