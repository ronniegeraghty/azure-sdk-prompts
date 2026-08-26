import { randomBytes } from "node:crypto";

import type { TokenCredential } from "@azure/core-auth";
import {
  CryptographyClient,
  KeyClient,
  KnownEncryptionAlgorithms,
} from "@azure/keyvault-keys";

const DATA_KEY_LENGTH_BYTES = 32;
const KEY_WRAP_ALGORITHM = KnownEncryptionAlgorithms.RSAOaep256;

export interface WrappedDataKey {
  readonly keyId: string;
  readonly wrappedKey: Uint8Array;
  readonly algorithm: typeof KEY_WRAP_ALGORITHM;
}

export interface EnvelopeDataKey extends WrappedDataKey {
  readonly plaintextKey: Buffer;
}

export class KeyManagement {
  public constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: TokenCredential,
    private readonly keyName: string,
  ) {}

  public async generateAndWrapDataKey(): Promise<EnvelopeDataKey> {
    const plaintextKey = randomBytes(DATA_KEY_LENGTH_BYTES);

    try {
      const key = await this.keyClient.getKey(this.keyName);
      if (!key.id) {
        throw new Error(`Key Vault key "${this.keyName}" has no key ID`);
      }

      const cryptographyClient = new CryptographyClient(
        key.id,
        this.credential,
      );
      const result = await cryptographyClient.wrapKey(
        KEY_WRAP_ALGORITHM,
        plaintextKey,
      );

      return {
        plaintextKey,
        wrappedKey: result.result,
        keyId: key.id,
        algorithm: KEY_WRAP_ALGORITHM,
      };
    } catch (error) {
      plaintextKey.fill(0);
      throw new Error(
        `Failed to generate and wrap a data key with Key Vault key "${this.keyName}"`,
        { cause: error },
      );
    }
  }

  public async unwrapDataKey(
    keyId: string,
    wrappedKey: Uint8Array,
    algorithm: string,
  ): Promise<Buffer> {
    if (algorithm !== KEY_WRAP_ALGORITHM) {
      throw new Error(`Unsupported key-wrap algorithm: ${algorithm}`);
    }

    try {
      const cryptographyClient = new CryptographyClient(
        keyId,
        this.credential,
      );
      const result = await cryptographyClient.unwrapKey(
        KEY_WRAP_ALGORITHM,
        wrappedKey,
      );
      const plaintextKey = Buffer.from(
        result.result.buffer,
        result.result.byteOffset,
        result.result.byteLength,
      );

      if (plaintextKey.length !== DATA_KEY_LENGTH_BYTES) {
        plaintextKey.fill(0);
        throw new Error(
          `Unwrapped data key has ${plaintextKey.length} bytes; expected ${DATA_KEY_LENGTH_BYTES}`,
        );
      }

      return plaintextKey;
    } catch (error) {
      throw new Error(`Failed to unwrap the data key with Key Vault`, {
        cause: error,
      });
    }
  }
}
