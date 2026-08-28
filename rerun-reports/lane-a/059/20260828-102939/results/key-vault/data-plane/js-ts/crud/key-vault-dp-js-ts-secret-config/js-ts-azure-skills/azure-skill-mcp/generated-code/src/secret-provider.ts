import type { SecretClientLike } from "./secret-client.js";

export interface SecretValue {
  name: string;
  value: string;
  version?: string;
  expiresOn?: Date;
  found: boolean;
}

interface ServiceErrorLike {
  code?: unknown;
  statusCode?: unknown;
}

function isServiceErrorLike(error: unknown): error is ServiceErrorLike {
  return typeof error === "object" && error !== null;
}

function isSecretNotFound(error: unknown): boolean {
  if (!isServiceErrorLike(error)) {
    return false;
  }

  return error.statusCode === 404 || error.code === "SecretNotFound";
}

export function expiresWithin(
  expiresOn: Date | undefined,
  warningWindowMs: number,
  now = new Date(),
): boolean {
  if (!expiresOn) {
    return false;
  }

  return expiresOn.getTime() - now.getTime() <= warningWindowMs;
}

export class KeyVaultSecretProvider {
  public constructor(private readonly client: SecretClientLike) {}

  public async getSecret(
    name: string,
    defaultValue: string,
    version?: string,
  ): Promise<SecretValue> {
    try {
      const secret = await this.client.getSecret(name, version ? { version } : undefined);

      return {
        name,
        value: secret.value ?? defaultValue,
        found: true,
        ...(secret.properties.version !== undefined
          ? { version: secret.properties.version }
          : {}),
        ...(secret.properties.expiresOn !== undefined
          ? { expiresOn: secret.properties.expiresOn }
          : {}),
      };
    } catch (error: unknown) {
      if (!isSecretNotFound(error)) {
        throw error;
      }

      return {
        name,
        value: defaultValue,
        found: false,
      };
    }
  }

  public async getSecretVersion(
    name: string,
    version: string,
    defaultValue: string,
  ): Promise<SecretValue> {
    return this.getSecret(name, defaultValue, version);
  }

  public async inspectExpiry(name: string, version?: string): Promise<Date | undefined> {
    const secret = await this.getSecret(name, "", version);
    return secret.expiresOn;
  }

  public isNearExpiry(
    secret: Pick<SecretValue, "expiresOn">,
    warningWindowMs: number,
    now = new Date(),
  ): boolean {
    return expiresWithin(secret.expiresOn, warningWindowMs, now);
  }
}
