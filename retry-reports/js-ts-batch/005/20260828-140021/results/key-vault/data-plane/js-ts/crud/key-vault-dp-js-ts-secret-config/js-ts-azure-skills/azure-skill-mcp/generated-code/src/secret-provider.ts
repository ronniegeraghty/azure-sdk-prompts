import type { KeyVaultSecret, SecretClient } from "@azure/keyvault-secrets";

export interface SecretValue {
  name: string;
  value: string;
  version?: string;
  expiresOn?: Date;
  usedDefault: boolean;
}

export interface SecretLookup {
  name: string;
  defaultValue: string;
  version?: string;
}

export class KeyVaultSecretProvider {
  public constructor(private readonly client: SecretClient) {}

  public async getSecret(
    name: string,
    defaultValue: string,
    version?: string,
  ): Promise<SecretValue> {
    try {
      const secret = await this.client.getSecret(
        name,
        version === undefined ? undefined : { version },
      );

      return this.toSecretValue(secret, defaultValue);
    } catch (error: unknown) {
      if (isNotFoundError(error)) {
        return {
          name,
          value: defaultValue,
          ...(version === undefined ? {} : { version }),
          usedDefault: true,
        };
      }

      throw error;
    }
  }

  public async getSecretVersion(
    name: string,
    version: string,
    defaultValue: string,
  ): Promise<SecretValue> {
    return this.getSecret(name, defaultValue, version);
  }

  public isExpiringWithin(
    secret: Pick<SecretValue, "expiresOn">,
    warningWindowMs: number,
    now = new Date(),
  ): boolean {
    if (secret.expiresOn === undefined) {
      return false;
    }

    const remainingMs = secret.expiresOn.getTime() - now.getTime();
    return remainingMs <= warningWindowMs;
  }

  private toSecretValue(
    secret: KeyVaultSecret,
    defaultValue: string,
  ): SecretValue {
    return {
      name: secret.name,
      value: secret.value ?? defaultValue,
      ...(secret.properties.version === undefined
        ? {}
        : { version: secret.properties.version }),
      ...(secret.properties.expiresOn === undefined
        ? {}
        : { expiresOn: secret.properties.expiresOn }),
      usedDefault: secret.value === undefined,
    };
  }
}

function isNotFoundError(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  return "statusCode" in error && error.statusCode === 404;
}
