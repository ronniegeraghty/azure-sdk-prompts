import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";
import type { TokenCredential } from "@azure/core-auth";
import { randomBytes } from "node:crypto";

const DATA_KEY_LENGTH_BYTES = 32;
const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  dataKey: Buffer;
  wrappedDataKey: Uint8Array;
  keyId: string;
  wrappingAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

export class KeyVaultOperationError extends Error {
  constructor(
    message: string,
    options: ErrorOptions,
  ) {
    super(message, options);
    this.name = "KeyVaultOperationError";
  }
}

export class KeyManagement {
  constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: TokenCredential,
    private readonly keyName: string,
  ) {}

  async generateAndProtectDataKey(): Promise<ProtectedDataKey> {
    const dataKey = randomBytes(DATA_KEY_LENGTH_BYTES);

    try {
      const vaultKey = await this.keyClient.getKey(this.keyName);
      if (!vaultKey.id) {
        throw new Error(`Key Vault did not return an ID for key "${this.keyName}"`);
      }

      const cryptographyClient = new CryptographyClient(
        vaultKey.id,
        this.credential,
      );
      const result = await cryptographyClient.wrapKey(
        KEY_WRAP_ALGORITHM,
        dataKey,
      );

      return {
        dataKey,
        wrappedDataKey: result.result,
        keyId: vaultKey.id,
        wrappingAlgorithm: KEY_WRAP_ALGORITHM,
      };
    } catch (error) {
      dataKey.fill(0);
      throw new KeyVaultOperationError(
        `Failed to protect the data encryption key with Key Vault key "${this.keyName}"`,
        { cause: error },
      );
    }
  }

  async recoverDataKey(
    wrappedDataKey: Uint8Array,
    keyId: string,
    wrappingAlgorithm: string,
  ): Promise<Buffer> {
    if (wrappingAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new KeyVaultOperationError(
        `Unsupported key wrapping algorithm: ${wrappingAlgorithm}`,
        { cause: new Error("Invalid encrypted blob metadata") },
      );
    }

    try {
      const cryptographyClient = new CryptographyClient(keyId, this.credential);
      const result = await cryptographyClient.unwrapKey(
        KEY_WRAP_ALGORITHM,
        wrappedDataKey,
      );
      const dataKey = Buffer.from(result.result);
      if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {
        dataKey.fill(0);
        throw new Error(
          `Key Vault returned an invalid data key length: ${dataKey.length}`,
        );
      }
      return dataKey;
    } catch (error) {
      throw new KeyVaultOperationError(
        `Failed to recover the data encryption key with Key Vault key "${keyId}"`,
        { cause: error },
      );
    }
  }
}
