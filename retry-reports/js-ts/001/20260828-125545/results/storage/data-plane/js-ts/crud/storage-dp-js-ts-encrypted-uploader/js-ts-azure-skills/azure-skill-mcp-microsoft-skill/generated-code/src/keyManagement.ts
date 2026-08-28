import { randomBytes } from "node:crypto";
import type { TokenCredential } from "@azure/core-auth";
import {
  CryptographyClient,
  type KeyClient,
  type KeyVaultKey,
} from "@azure/keyvault-keys";

export const DATA_KEY_LENGTH_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  keyId: string;
  wrappedKey: Uint8Array;
  wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

export class KeyManagementError extends Error {
  constructor(operation: string, cause: unknown) {
    super(
      `Azure Key Vault ${operation} failed${formatAzureErrorDetails(cause)}.`,
      { cause },
    );
    this.name = "KeyManagementError";
  }
}

export class KeyManagement {
  private readonly vaultOrigin: string;

  public constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: TokenCredential,
    keyVaultUrl: string,
    private readonly keyName: string,
  ) {
    this.vaultOrigin = new URL(keyVaultUrl).origin.toLowerCase();
  }

  public async withNewDataKey<T>(
    useDataKey: (
      dataKey: Buffer,
      protectedDataKey: ProtectedDataKey,
    ) => Promise<T> | T,
  ): Promise<T> {
    const dataKey = randomBytes(DATA_KEY_LENGTH_BYTES);

    try {
      const key = await this.getWrappingKey();
      const keyId = this.requireVersionedKeyId(key);
      const cryptographyClient = new CryptographyClient(
        keyId,
        this.credential,
      );

      let wrappedKey: Uint8Array;
      try {
        const result = await cryptographyClient.wrapKey(
          KEY_WRAP_ALGORITHM,
          dataKey,
        );
        wrappedKey = result.result;
      } catch (error) {
        throw new KeyManagementError(
          `wrap operation with key "${this.keyName}"`,
          error,
        );
      }

      return await useDataKey(dataKey, {
        keyId,
        wrappedKey,
        wrapAlgorithm: KEY_WRAP_ALGORITHM,
      });
    } finally {
      dataKey.fill(0);
    }
  }

  public async withUnwrappedDataKey<T>(
    protectedDataKey: ProtectedDataKey,
    useDataKey: (dataKey: Buffer) => Promise<T> | T,
  ): Promise<T> {
    this.validateProtectedKey(protectedDataKey);

    const cryptographyClient = new CryptographyClient(
      protectedDataKey.keyId,
      this.credential,
    );

    let unwrappedKey: Uint8Array;
    try {
      const result = await cryptographyClient.unwrapKey(
        protectedDataKey.wrapAlgorithm,
        protectedDataKey.wrappedKey,
      );
      unwrappedKey = result.result;
    } catch (error) {
      throw new KeyManagementError(
        `unwrap operation with key "${protectedDataKey.keyId}"`,
        error,
      );
    }

    const dataKey = Buffer.from(unwrappedKey);
    unwrappedKey.fill(0);

    try {
      if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {
        throw new Error(
          `Unwrapped data key has ${dataKey.length} bytes; expected ${DATA_KEY_LENGTH_BYTES}.`,
        );
      }
      return await useDataKey(dataKey);
    } finally {
      dataKey.fill(0);
    }
  }

  private async getWrappingKey(): Promise<KeyVaultKey> {
    let key: KeyVaultKey;
    try {
      key = await this.keyClient.getKey(this.keyName);
    } catch (error) {
      throw new KeyManagementError(`get key "${this.keyName}"`, error);
    }

    if (key.properties.enabled === false) {
      throw new KeyManagementError(
        `get key "${this.keyName}"`,
        new Error("The configured key is disabled."),
      );
    }

    if (key.keyType !== "RSA" && key.keyType !== "RSA-HSM") {
      throw new KeyManagementError(
        `get key "${this.keyName}"`,
        new Error(
          `RSA-OAEP-256 requires an RSA or RSA-HSM key, not "${key.keyType}".`,
        ),
      );
    }

    return key;
  }

  private requireVersionedKeyId(key: KeyVaultKey): string {
    if (!key.id || !key.properties.version) {
      throw new KeyManagementError(
        `get key "${this.keyName}"`,
        new Error("Key Vault did not return a versioned key ID."),
      );
    }
    return key.id;
  }

  private validateProtectedKey(protectedDataKey: ProtectedDataKey): void {
    if (protectedDataKey.wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new Error(
        `Unsupported key wrap algorithm "${protectedDataKey.wrapAlgorithm}".`,
      );
    }

    const keyUrl = new URL(protectedDataKey.keyId);
    const segments = keyUrl.pathname.split("/").filter(Boolean);
    const expectedKeyName = this.keyName.toLowerCase();

    if (
      keyUrl.origin.toLowerCase() !== this.vaultOrigin ||
      segments.length !== 3 ||
      segments[0]?.toLowerCase() !== "keys" ||
      segments[1]?.toLowerCase() !== expectedKeyName ||
      !segments[2]
    ) {
      throw new Error(
        "The protected data key references an unexpected vault, key, or key version.",
      );
    }
  }
}

function formatAzureErrorDetails(error: unknown): string {
  if (!error || typeof error !== "object") {
    return "";
  }

  const candidate = error as {
    code?: unknown;
    statusCode?: unknown;
  };
  const details: string[] = [];

  if (typeof candidate.code === "string") {
    details.push(`code ${candidate.code}`);
  }
  if (typeof candidate.statusCode === "number") {
    details.push(`HTTP ${candidate.statusCode}`);
  }

  return details.length > 0 ? ` (${details.join(", ")})` : "";
}
