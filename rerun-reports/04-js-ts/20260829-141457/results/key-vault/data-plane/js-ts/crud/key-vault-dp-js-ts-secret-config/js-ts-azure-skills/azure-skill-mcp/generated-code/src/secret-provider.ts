import type { GetSecretOptions, KeyVaultSecret } from "@azure/keyvault-secrets";

export interface SecretReader {
  getSecret(name: string, options?: GetSecretOptions): Promise<KeyVaultSecret>;
}

export interface SecretRequest {
  defaultValue: string;
  version?: string;
}

export interface ResolvedSecret {
  name: string;
  value: string;
  version?: string;
  expiresOn?: Date;
  found: boolean;
}

function isSecretNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { statusCode?: unknown; code?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}

export class KeyVaultSecretProvider {
  public constructor(private readonly client: SecretReader) {}

  public async getSecret(name: string, request: SecretRequest): Promise<ResolvedSecret> {
    if (!name.trim()) {
      throw new Error("Secret name must not be empty.");
    }

    try {
      const options = request.version ? { version: request.version } : undefined;
      const secret = await this.client.getSecret(name, options);

      if (secret.value === undefined) {
        throw new Error(`Key Vault returned secret "${name}" without a value.`);
      }

      return {
        name: secret.name,
        value: secret.value,
        version: secret.properties.version,
        expiresOn: secret.properties.expiresOn,
        found: true
      };
    } catch (error: unknown) {
      if (!isSecretNotFound(error)) {
        throw error;
      }

      return {
        name,
        value: request.defaultValue,
        version: request.version,
        found: false
      };
    }
  }

  public isExpiringSoon(
    secret: Pick<ResolvedSecret, "expiresOn">,
    warningWindowMs: number,
    now = new Date()
  ): boolean {
    if (warningWindowMs < 0) {
      throw new Error("Expiry warning window must not be negative.");
    }

    return secret.expiresOn !== undefined
      && secret.expiresOn.getTime() - now.getTime() <= warningWindowMs;
  }
}
