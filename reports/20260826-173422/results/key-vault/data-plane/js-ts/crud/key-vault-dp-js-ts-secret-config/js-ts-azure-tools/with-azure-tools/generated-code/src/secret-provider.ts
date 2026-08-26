import type { SecretStore } from "./secret-store.js";

export interface ResolvedSecret {
  name: string;
  value: string;
  found: boolean;
  version?: string;
  expiresOn?: Date;
}

export class SecretProvider {
  public constructor(private readonly client: SecretStore) {}

  public async getSecret(
    name: string,
    defaultValue: string,
    version?: string,
  ): Promise<ResolvedSecret> {
    try {
      const options = version === undefined ? undefined : { version };
      const secret = await this.client.getSecret(name, options);

      return {
        name,
        value: secret.value ?? defaultValue,
        found: secret.value !== undefined,
        ...(secret.properties.version === undefined
          ? {}
          : { version: secret.properties.version }),
        ...(secret.properties.expiresOn === undefined
          ? {}
          : { expiresOn: secret.properties.expiresOn }),
      };
    } catch (error: unknown) {
      if (!isSecretNotFound(error)) {
        throw error;
      }

      return { name, value: defaultValue, found: false };
    }
  }

  public isNearExpiry(
    secret: Pick<ResolvedSecret, "expiresOn">,
    warningWindowMs: number,
    now = new Date(),
  ): boolean {
    return (
      secret.expiresOn !== undefined &&
      secret.expiresOn.getTime() <= now.getTime() + warningWindowMs
    );
  }
}

function isSecretNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}
