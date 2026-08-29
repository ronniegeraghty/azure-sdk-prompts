import { randomBytes } from "node:crypto";

import type { ManagedIdentityCredential } from "@azure/identity";
import { CryptographyClient, type KeyClient } from "@azure/keyvault-keys";

import {
  EncryptionMetadataError,
  getErrorMessage,
  KeyVaultOperationError,
} from "./errors.js";

const DATA_KEY_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface WrappedDataKey {
  keyId: string;
  plaintextKey: Buffer;
  wrappedKey: Buffer;
}

export class KeyManagementClient {
  readonly #credential: ManagedIdentityCredential;
  readonly #keyClient: KeyClient;
  readonly #keyName: string;
  readonly #keyIdPrefix: string;

  public constructor(
    keyClient: KeyClient,
    credential: ManagedIdentityCredential,
    keyVaultUrl: string,
    keyName: string,
  ) {
    this.#keyClient = keyClient;
    this.#credential = credential;
    this.#keyName = keyName;
    this.#keyIdPrefix = `${keyVaultUrl.replace(/\/+$/, "")}/keys/`;
  }

  public async generateAndWrapDataKey(): Promise<WrappedDataKey> {
    const plaintextKey = randomBytes(DATA_KEY_BYTES);

    try {
      const vaultKey = await this.#keyClient.getKey(this.#keyName);
      if (!vaultKey.id) {
        throw new Error(`Key "${this.#keyName}" did not include a key ID.`);
      }

      const cryptographyClient = new CryptographyClient(
        vaultKey.id,
        this.#credential,
      );
      const wrapResult = await cryptographyClient.wrapKey(
        KEY_WRAP_ALGORITHM,
        plaintextKey,
      );

      return {
        keyId: vaultKey.id,
        plaintextKey,
        wrappedKey: Buffer.from(wrapResult.result),
      };
    } catch (error) {
      plaintextKey.fill(0);
      throw new KeyVaultOperationError(
        "key wrapping",
        getErrorMessage(error),
        { cause: error },
      );
    }
  }

  public async unwrapDataKey(
    wrappedKey: Uint8Array,
    keyId: string,
  ): Promise<Buffer> {
    this.#assertKeyIdBelongsToConfiguredVault(keyId);

    try {
      const cryptographyClient = new CryptographyClient(
        keyId,
        this.#credential,
      );
      const unwrapResult = await cryptographyClient.unwrapKey(
        KEY_WRAP_ALGORITHM,
        wrappedKey,
      );
      const plaintextKey = Buffer.from(unwrapResult.result);

      if (plaintextKey.length !== DATA_KEY_BYTES) {
        plaintextKey.fill(0);
        throw new Error(
          `Unwrapped data key was ${plaintextKey.length} bytes; expected ${DATA_KEY_BYTES}.`,
        );
      }

      return plaintextKey;
    } catch (error) {
      throw new KeyVaultOperationError(
        "key unwrapping",
        getErrorMessage(error),
        { cause: error },
      );
    }
  }

  #assertKeyIdBelongsToConfiguredVault(keyId: string): void {
    if (!keyId.startsWith(this.#keyIdPrefix)) {
      throw new EncryptionMetadataError(
        "the key ID does not belong to the configured Key Vault.",
      );
    }
  }
}
