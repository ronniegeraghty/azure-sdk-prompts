import type { GetSecretOptions, KeyVaultSecret } from "@azure/keyvault-secrets";

const millisecondsPerDay = 24 * 60 * 60 * 1_000;

export interface SecretClientReader {
  getSecret(name: string, options?: GetSecretOptions): Promise<KeyVaultSecret>;
}

export interface SecretResult {
  name: string;
  value: string;
  found: boolean;
  usedDefault: boolean;
  version?: string;
  expiresOn?: Date;
}

export interface SecretExpiryStatus {
  name: string;
  found: boolean;
  expiresOn?: Date;
  isExpired: boolean;
  isNearExpiry: boolean;
}

export interface SecretReader {
  getSecret(name: string, defaultValue?: string, version?: string): Promise<SecretResult>;
}

export class KeyVaultSecretProvider implements SecretReader {
  public constructor(private readonly client: SecretClientReader) {}

  public async getSecret(
    name: string,
    defaultValue = "",
    version?: string,
  ): Promise<SecretResult> {
    if (name.trim().length === 0) {
      throw new Error("Secret name must not be empty.");
    }

    try {
      const options: GetSecretOptions = version === undefined ? {} : { version };
      const secret = await this.client.getSecret(name, options);
      const usedDefault = secret.value === undefined;

      return {
        name,
        value: secret.value ?? defaultValue,
        found: true,
        usedDefault,
        ...(secret.properties.version === undefined
          ? {}
          : { version: secret.properties.version }),
        ...(secret.properties.expiresOn === undefined
          ? {}
          : { expiresOn: secret.properties.expiresOn }),
      };
    } catch (error: unknown) {
      if (!isSecretNotFoundError(error)) {
        throw error;
      }

      return {
        name,
        value: defaultValue,
        found: false,
        usedDefault: true,
      };
    }
  }

  public async getExpiryStatus(
    name: string,
    warningWindowDays = 7,
    version?: string,
    now = new Date(),
  ): Promise<SecretExpiryStatus> {
    assertWarningWindow(warningWindowDays);
    const secret = await this.getSecret(name, "", version);
    return expiryStatusFromSecret(secret, warningWindowDays, now);
  }
}

export function expiryStatusFromSecret(
  secret: Pick<SecretResult, "name" | "found" | "expiresOn">,
  warningWindowDays: number,
  now = new Date(),
): SecretExpiryStatus {
  assertWarningWindow(warningWindowDays);
  const expiresOn = secret.expiresOn;

  return {
    name: secret.name,
    found: secret.found,
    ...(expiresOn === undefined ? {} : { expiresOn }),
    isExpired: expiresOn !== undefined && expiresOn.getTime() <= now.getTime(),
    isNearExpiry:
      expiresOn !== undefined &&
      expiresOn.getTime() <= now.getTime() + warningWindowDays * millisecondsPerDay,
  };
}

function assertWarningWindow(warningWindowDays: number): void {
  if (!Number.isFinite(warningWindowDays) || warningWindowDays < 0) {
    throw new Error("Expiry warning window must be a non-negative number of days.");
  }
}

function isSecretNotFoundError(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}
