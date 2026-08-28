import { randomBytes } from "node:crypto";
import type { ManagedIdentityCredential } from "@azure/identity";
import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";
import { KeyManagementError } from "./errors.js";

const DATA_KEY_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  keyId: string;
  wrappedKey: Buffer;
  wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

export class KeyManagement {
  private readonly vaultOrigin: string;
  private readonly vaultPathPrefix: string;

  constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: ManagedIdentityCredential,
    keyVaultUrl: string,
    private readonly keyName: string,
  ) {
    const vaultUrl = new URL(keyVaultUrl);
    this.vaultOrigin = vaultUrl.origin;
    this.vaultPathPrefix = `/keys/${encodeURIComponent(keyName)}/`;
  }

  async withNewDataKey<T>(
    operation: (
      dataKey: Buffer,
      protectedKey: ProtectedDataKey,
    ) => Promise<T>,
  ): Promise<T> {
    const dataKey = randomBytes(DATA_KEY_BYTES);

    try {
      const protectedKey = await this.protectDataKey(dataKey);
      return await operation(dataKey, protectedKey);
    } finally {
      dataKey.fill(0);
    }
  }

  async withRecoveredDataKey<T>(
    protectedKey: ProtectedDataKey,
    operation: (dataKey: Buffer) => Promise<T>,
  ): Promise<T> {
    const dataKey = await this.recoverDataKey(protectedKey);

    try {
      if (dataKey.length !== DATA_KEY_BYTES) {
        throw new Error(
          `The recovered data key is ${dataKey.length} bytes; expected ${DATA_KEY_BYTES}.`,
        );
      }

      return await operation(dataKey);
    } finally {
      dataKey.fill(0);
    }
  }

  private async protectDataKey(dataKey: Buffer): Promise<ProtectedDataKey> {
    try {
      const key = await this.keyClient.getKey(this.keyName);
      if (!key.id) {
        throw new Error(`Key Vault key ${this.keyName} did not return a key ID.`);
      }

      const cryptoClient = new CryptographyClient(key.id, this.credential);
      const wrapped = await cryptoClient.wrapKey(KEY_WRAP_ALGORITHM, dataKey);

      return {
        keyId: key.id,
        wrappedKey: Buffer.from(wrapped.result),
        wrapAlgorithm: KEY_WRAP_ALGORITHM,
      };
    } catch (error) {
      throw new KeyManagementError("key wrapping", error);
    }
  }

  private async recoverDataKey(
    protectedKey: ProtectedDataKey,
  ): Promise<Buffer> {
    try {
      this.validateKeyId(protectedKey.keyId);

      if (protectedKey.wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
        throw new Error(
          `Unsupported key wrap algorithm: ${protectedKey.wrapAlgorithm}.`,
        );
      }

      const cryptoClient = new CryptographyClient(
        protectedKey.keyId,
        this.credential,
      );
      const unwrapped = await cryptoClient.unwrapKey(
        protectedKey.wrapAlgorithm,
        protectedKey.wrappedKey,
      );

      return Buffer.from(unwrapped.result);
    } catch (error) {
      throw new KeyManagementError("key unwrapping", error);
    }
  }

  private validateKeyId(keyId: string): void {
    const parsedKeyId = new URL(keyId);
    const isVersionedConfiguredKey =
      parsedKeyId.origin === this.vaultOrigin &&
      parsedKeyId.pathname.startsWith(this.vaultPathPrefix) &&
      parsedKeyId.pathname.slice(this.vaultPathPrefix.length).length > 0 &&
      !parsedKeyId.pathname
        .slice(this.vaultPathPrefix.length)
        .includes("/");

    if (!isVersionedConfiguredKey) {
      throw new Error(
        "The protected data key references an unexpected vault, key, or unversioned key ID.",
      );
    }
  }
}
