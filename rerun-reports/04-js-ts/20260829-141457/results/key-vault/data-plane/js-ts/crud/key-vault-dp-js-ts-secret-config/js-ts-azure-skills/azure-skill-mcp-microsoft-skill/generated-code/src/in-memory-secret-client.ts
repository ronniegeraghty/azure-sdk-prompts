import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";
import type { DeleteSecretPoller, SecretClientLike } from "./types.js";

interface StoredVersion {
  value: string;
  version: string;
  createdOn: Date;
  options: SetSecretOptions;
}

function notFound(name: string): Error & { code: string; statusCode: number } {
  return Object.assign(new Error(`Secret '${name}' was not found`), {
    code: "SecretNotFound",
    statusCode: 404,
  });
}

export class InMemorySecretClient implements SecretClientLike {
  private readonly secrets = new Map<string, StoredVersion[]>();
  private readonly deleted = new Set<string>();
  private versionCounter = 0;
  private readonly requests = new Map<string, number>();
  public readonly operations: string[] = [];

  public async getSecret(
    name: string,
    options?: GetSecretOptions,
  ): Promise<KeyVaultSecret> {
    this.requests.set(name, (this.requests.get(name) ?? 0) + 1);
    const versions = this.secrets.get(name);
    const stored =
      options?.version === undefined
        ? versions?.at(-1)
        : versions?.find((candidate) => candidate.version === options.version);

    if (stored === undefined || this.deleted.has(name)) {
      throw notFound(name);
    }

    return this.toKeyVaultSecret(name, stored);
  }

  public async setSecret(
    name: string,
    value: string,
    options: SetSecretOptions = {},
  ): Promise<KeyVaultSecret> {
    if (this.deleted.has(name)) {
      throw Object.assign(
        new Error(`Secret '${name}' is soft-deleted and must be purged or recovered`),
        { code: "Conflict", statusCode: 409 },
      );
    }

    const stored: StoredVersion = {
      value,
      version: `v${++this.versionCounter}`,
      createdOn: new Date(),
      options,
    };
    const versions = this.secrets.get(name) ?? [];
    versions.push(stored);
    this.secrets.set(name, versions);
    this.operations.push(`set:${name}:${stored.version}`);
    return this.toKeyVaultSecret(name, stored);
  }

  public async beginDeleteSecret(name: string): Promise<DeleteSecretPoller> {
    if (!this.secrets.has(name)) {
      throw notFound(name);
    }

    this.operations.push(`begin-delete:${name}`);
    return {
      pollUntilDone: async () => {
        this.deleted.add(name);
        this.operations.push(`delete-complete:${name}`);
        return {};
      },
    };
  }

  public async purgeDeletedSecret(name: string): Promise<void> {
    if (!this.deleted.has(name)) {
      throw notFound(name);
    }

    this.operations.push(`purge:${name}`);
    this.deleted.delete(name);
    this.secrets.delete(name);
  }

  public getRequestCount(name: string): number {
    return this.requests.get(name) ?? 0;
  }

  private toKeyVaultSecret(
    name: string,
    stored: StoredVersion,
  ): KeyVaultSecret {
    return {
      name,
      value: stored.value,
      properties: {
        name,
        vaultUrl: "https://offline-demo.vault.azure.net",
        id: `https://offline-demo.vault.azure.net/secrets/${name}/${stored.version}`,
        version: stored.version,
        createdOn: stored.createdOn,
        ...stored.options,
      },
    };
  }
}
