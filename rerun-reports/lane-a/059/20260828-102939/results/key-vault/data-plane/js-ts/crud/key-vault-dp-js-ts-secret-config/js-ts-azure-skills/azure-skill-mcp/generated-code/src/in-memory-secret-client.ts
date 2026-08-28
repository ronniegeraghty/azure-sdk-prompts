import type {
  GetSecretOptionsLike,
  KeyVaultSecretLike,
  SecretClientLike,
  SetSecretOptionsLike,
} from "./secret-client.js";

interface StoredVersion {
  value: string;
  version: string;
  expiresOn?: Date;
}

export class InMemorySecretClient implements SecretClientLike {
  private readonly active = new Map<string, StoredVersion[]>();
  private readonly deleted = new Map<string, StoredVersion[]>();
  private nextVersion = 1;

  public async getSecret(
    name: string,
    options: GetSecretOptionsLike = {},
  ): Promise<KeyVaultSecretLike> {
    const versions = this.active.get(name);
    const secret = options.version
      ? versions?.find(({ version }) => version === options.version)
      : versions?.at(-1);

    if (!secret) {
      throw Object.assign(new Error(`Secret "${name}" was not found.`), {
        code: "SecretNotFound",
        statusCode: 404,
      });
    }

    return this.toSecret(name, secret);
  }

  public async setSecret(
    name: string,
    value: string,
    options: SetSecretOptionsLike = {},
  ): Promise<KeyVaultSecretLike> {
    if (this.deleted.has(name)) {
      throw new Error(`Secret "${name}" is soft-deleted and must be recovered or purged first.`);
    }

    const stored: StoredVersion = {
      value,
      version: `v${this.nextVersion++}`,
      ...(options.expiresOn !== undefined ? { expiresOn: options.expiresOn } : {}),
    };
    const versions = this.active.get(name) ?? [];
    versions.push(stored);
    this.active.set(name, versions);
    return this.toSecret(name, stored);
  }

  public async beginDeleteSecret(name: string): Promise<{ pollUntilDone(): Promise<void> }> {
    if (!this.active.has(name)) {
      throw Object.assign(new Error(`Secret "${name}" was not found.`), {
        code: "SecretNotFound",
        statusCode: 404,
      });
    }

    return {
      pollUntilDone: async () => {
        const versions = this.active.get(name);
        if (versions) {
          this.active.delete(name);
          this.deleted.set(name, versions);
        }
      },
    };
  }

  public async purgeDeletedSecret(name: string): Promise<void> {
    if (!this.deleted.delete(name)) {
      throw new Error(`Deleted secret "${name}" was not found.`);
    }
  }

  private toSecret(name: string, secret: StoredVersion): KeyVaultSecretLike {
    return {
      value: secret.value,
      properties: {
        name,
        version: secret.version,
        ...(secret.expiresOn !== undefined ? { expiresOn: secret.expiresOn } : {}),
      },
    };
  }
}
