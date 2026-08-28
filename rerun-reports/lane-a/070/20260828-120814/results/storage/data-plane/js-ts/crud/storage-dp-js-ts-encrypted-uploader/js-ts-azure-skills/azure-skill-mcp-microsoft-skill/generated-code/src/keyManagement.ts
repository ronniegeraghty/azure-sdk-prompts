import { randomBytes } from "node:crypto";
import type { TokenCredential } from "@azure/core-auth";
import {
  CryptographyClient,
  type KeyClient,
} from "@azure/keyvault-keys";

const DATA_KEY_BYTES = 32;
export const KEY_WRAP_ALGORITHM = "RSA-OAEP-256" as const;

export interface ProtectedDataKey {
  readonly dataKey: Buffer;
  readonly keyId: string;
  readonly wrappedKey: Buffer;
  readonly wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;
}

interface AzureErrorDetails {
  readonly code?: string;
  readonly message?: string;
  readonly statusCode?: number;
}

function getAzureErrorDetails(error: unknown): AzureErrorDetails {
  if (typeof error !== "object" || error === null) {
    return {};
  }

  const candidate = error as Record<string, unknown>;
  return {
    ...(typeof candidate.code === "string" ? { code: candidate.code } : {}),
    ...(typeof candidate.message === "string"
      ? { message: candidate.message }
      : {}),
    ...(typeof candidate.statusCode === "number"
      ? { statusCode: candidate.statusCode }
      : {}),
  };
}

export class KeyManagementError extends Error {
  public constructor(operation: string, cause: unknown) {
    const details = getAzureErrorDetails(cause);
    const context = [
      details.statusCode === undefined
        ? undefined
        : `status ${details.statusCode}`,
      details.code,
    ]
      .filter((value): value is string => value !== undefined)
      .join(", ");
    const reason = details.message ? `: ${details.message}` : "";

    super(
      `Azure Key Vault ${operation} failed${context ? ` (${context})` : ""}${reason}`,
      { cause },
    );
    this.name = "KeyManagementError";
  }
}

export class KeyManagement {
  public constructor(
    private readonly keyClient: KeyClient,
    private readonly credential: TokenCredential,
    private readonly vaultUrl: string,
    private readonly keyName: string,
  ) {}

  public async createProtectedDataKey(): Promise<ProtectedDataKey> {
    const dataKey = randomBytes(DATA_KEY_BYTES);

    try {
      const key = await this.keyClient.getKey(this.keyName);
      if (!key.id) {
        throw new Error(`Key Vault returned no ID for key ${this.keyName}.`);
      }

      const cryptographyClient = new CryptographyClient(
        key.id,
        this.credential,
      );
      const wrapResult = await cryptographyClient.wrapKey(
        KEY_WRAP_ALGORITHM,
        dataKey,
      );

      return {
        dataKey,
        keyId: key.id,
        wrappedKey: Buffer.from(wrapResult.result),
        wrapAlgorithm: KEY_WRAP_ALGORITHM,
      };
    } catch (error) {
      dataKey.fill(0);
      if (error instanceof KeyManagementError) {
        throw error;
      }
      throw new KeyManagementError("key lookup or wrap", error);
    }
  }

  public async recoverDataKey(
    keyId: string,
    wrappedKey: Buffer,
    wrapAlgorithm: string,
  ): Promise<Buffer> {
    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {
      throw new KeyManagementError(
        "unwrap",
        new Error(`Unsupported key wrap algorithm: ${wrapAlgorithm}`),
      );
    }

    this.assertConfiguredKeyId(keyId);

    try {
      const cryptographyClient = new CryptographyClient(keyId, this.credential);
      const unwrapResult = await cryptographyClient.unwrapKey(
        KEY_WRAP_ALGORITHM,
        wrappedKey,
      );
      const dataKey = Buffer.from(unwrapResult.result);

      if (dataKey.length !== DATA_KEY_BYTES) {
        dataKey.fill(0);
        throw new Error(
          `Key Vault returned a ${dataKey.length}-byte data key; expected ${DATA_KEY_BYTES} bytes.`,
        );
      }

      return dataKey;
    } catch (error) {
      if (error instanceof KeyManagementError) {
        throw error;
      }
      throw new KeyManagementError("unwrap", error);
    }
  }

  private assertConfiguredKeyId(keyId: string): void {
    let configuredVault: URL;
    let candidate: URL;

    try {
      configuredVault = new URL(this.vaultUrl);
      candidate = new URL(keyId);
    } catch (error) {
      throw new KeyManagementError("key ID validation", error);
    }

    const pathParts = candidate.pathname.split("/").filter(Boolean);
    const hasExpectedPath =
      pathParts.length === 3 &&
      pathParts[0]?.toLowerCase() === "keys" &&
      pathParts[1] === this.keyName &&
      Boolean(pathParts[2]);

    if (
      candidate.protocol !== "https:" ||
      candidate.origin !== configuredVault.origin ||
      candidate.search ||
      candidate.hash ||
      !hasExpectedPath
    ) {
      throw new KeyManagementError(
        "key ID validation",
        new Error(
          "Blob metadata references a key outside the configured vault, key name, or key version.",
        ),
      );
    }
  }
}
