import { randomBytes } from "node:crypto";

import { ManagedIdentityCredential } from "@azure/identity";
import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";

import { AzureOperationError } from "./errors.js";

const DATA_KEY_LENGTH_BYTES = 32;
const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  keyId: string;
  wrappedKey: Buffer;
  wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

export class KeyManagement {
  public constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: ManagedIdentityCredential,
    private readonly keyName: string,
  ) {}

  public generateDataKey(): Buffer {
    return randomBytes(DATA_KEY_LENGTH_BYTES);
  }

  public async protectDataKey(dataKey: Buffer): Promise<ProtectedDataKey> {
    if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {
      throw new Error(
        `The data encryption key must be ${DATA_KEY_LENGTH_BYTES} bytes.`,
      );
    }

    try {
      const key = await this.keyClient.getKey(this.keyName);
      if (!key.id) {
        throw new Error(`Key Vault returned key "${this.keyName}" without an ID.`);
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
      throw new AzureOperationError(
        "Azure Key Vault",
        `wrapping a data key with "${this.keyName}"`,
        error,
      );
    }
  }

  public async recoverDataKey(
    keyId: string,
    wrappedKey: Buffer,
    wrapAlgorithm: string,
  ): Promise<Buffer> {
    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new Error(`Unsupported key-wrap algorithm: ${wrapAlgorithm}`);
    }

    try {
      // The stored, versioned key ID keeps old blobs decryptable after rotation.
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
      throw new AzureOperationError(
        "Azure Key Vault",
        `unwrapping the data key with "${keyId}"`,
        error,
      );
    }
  }
}
