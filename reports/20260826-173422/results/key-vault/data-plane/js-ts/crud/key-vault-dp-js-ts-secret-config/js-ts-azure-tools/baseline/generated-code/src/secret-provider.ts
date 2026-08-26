import { RestError } from "@azure/core-rest-pipeline";

import type { KeyVaultClient } from "./key-vault-client.js";

export interface SecretResult {
  name: string;
  value: string;
  version?: string;
  expiresOn?: Date;
  usedDefault: boolean;
}

export interface GetSecretRequest {
  defaultValue: string;
  version?: string;
}

export class KeyVaultSecretProvider {
  public constructor(private readonly client: KeyVaultClient) {}

  public async getSecret(
    name: string,
    request: GetSecretRequest,
  ): Promise<SecretResult> {
    try {
      const secret = await this.client.getSecret(name, {
        ...(request.version === undefined ? {} : { version: request.version }),
      });
      if (secret.value === undefined) {
        throw new Error(`Key Vault returned no value for secret "${name}".`);
      }

      return {
        name,
        value: secret.value,
        ...(secret.properties.version === undefined
          ? {}
          : { version: secret.properties.version }),
        ...(secret.properties.expiresOn === undefined
          ? {}
          : { expiresOn: secret.properties.expiresOn }),
        usedDefault: false,
      };
    } catch (error: unknown) {
      if (error instanceof RestError && error.statusCode === 404) {
        return {
          name,
          value: request.defaultValue,
          usedDefault: true,
        };
      }

      throw error;
    }
  }

  public expiresWithin(
    secret: Pick<SecretResult, "expiresOn">,
    warningWindowMs: number,
    now = new Date(),
  ): boolean {
    if (secret.expiresOn === undefined) {
      return false;
    }

    return secret.expiresOn.getTime() <= now.getTime() + warningWindowMs;
  }
}
