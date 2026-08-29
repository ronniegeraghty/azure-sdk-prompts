import type { KeyVaultSecret } from "@azure/keyvault-secrets";

import type { SecretStore } from "./secret-store.js";

export interface SecretResult {
  name: string;
  value: string;
  version?: string;
  expiresOn?: Date;
  found: boolean;
}

function isNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { statusCode?: unknown; code?: unknown };
  return (
    candidate.statusCode === 404 ||
    candidate.code === "SecretNotFound" ||
    candidate.code === "ResourceNotFound"
  );
}

function toResult(
  name: string,
  secret: KeyVaultSecret,
  defaultValue: string,
): SecretResult {
  const result: SecretResult = {
    name,
    value: secret.value ?? defaultValue,
    found: secret.value !== undefined,
  };

  if (secret.properties.version !== undefined) {
    result.version = secret.properties.version;
  }
  if (secret.properties.expiresOn !== undefined) {
    result.expiresOn = secret.properties.expiresOn;
  }

  return result;
}

export class KeyVaultSecretProvider {
  constructor(private readonly client: SecretStore) {}

  async getSecret(
    name: string,
    defaultValue = "",
    version?: string,
  ): Promise<SecretResult> {
    try {
      const secret =
        version === undefined
          ? await this.client.getSecret(name)
          : await this.client.getSecret(name, { version });
      return toResult(name, secret, defaultValue);
    } catch (error) {
      if (!isNotFound(error)) {
        throw error;
      }

      return { name, value: defaultValue, found: false };
    }
  }

  async getSecretVersion(
    name: string,
    version: string,
    defaultValue = "",
  ): Promise<SecretResult> {
    return this.getSecret(name, defaultValue, version);
  }

  isNearExpiry(
    secret: Pick<SecretResult, "expiresOn">,
    warningWindowMs: number,
    now = new Date(),
  ): boolean {
    return (
      secret.expiresOn !== undefined &&
      secret.expiresOn.getTime() <= now.getTime() + warningWindowMs
    );
  }
}
