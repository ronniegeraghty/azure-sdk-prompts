import type {
  GetSecretOptions,
  KeyVaultSecret,
  SecretProperties,
  SetSecretOptions,
} from "@azure/keyvault-secrets";

import type { DeleteSecretPoller, SecretStore } from "./secret-store.js";

interface StoredSecret {
  value: string;
  properties: SecretProperties;
}

function notFound(name: string): Error & { statusCode: number } {
  return Object.assign(new Error(`Secret "${name}" was not found`), {
    statusCode: 404,
  });
}

export class InMemorySecretStore implements SecretStore {
  private readonly secrets = new Map<string, StoredSecret[]>();
  private readonly deleted = new Set<string>();
  private nextVersion = 1;

  async getSecret(
    name: string,
    options: GetSecretOptions = {},
  ): Promise<KeyVaultSecret> {
    const versions = this.secrets.get(name);
    const stored =
      options.version === undefined
        ? versions?.at(-1)
        : versions?.find(
            (item) => item.properties.version === options.version,
          );
    if (stored === undefined) {
      throw notFound(name);
    }

    return {
      name,
      value: stored.value,
      properties: { ...stored.properties },
    };
  }

  async setSecret(
    name: string,
    value: string,
    options: SetSecretOptions = {},
  ): Promise<KeyVaultSecret> {
    const version = `local-${this.nextVersion++}`;
    const properties: SecretProperties = {
      name,
      vaultUrl: "https://offline-demo.vault.azure.net",
      id: `https://offline-demo.vault.azure.net/secrets/${name}/${version}`,
      version,
      createdOn: new Date(),
      updatedOn: new Date(),
      recoverableDays: 90,
      recoveryLevel: "Recoverable+Purgeable",
      ...(options.enabled === undefined ? {} : { enabled: options.enabled }),
      ...(options.expiresOn === undefined
        ? {}
        : { expiresOn: options.expiresOn }),
      ...(options.notBefore === undefined
        ? {}
        : { notBefore: options.notBefore }),
      ...(options.contentType === undefined
        ? {}
        : { contentType: options.contentType }),
      ...(options.tags === undefined ? {} : { tags: options.tags }),
    };
    const stored = { value, properties };
    const versions = this.secrets.get(name) ?? [];
    versions.push(stored);
    this.secrets.set(name, versions);
    this.deleted.delete(name);
    return { name, value, properties: { ...properties } };
  }

  async beginDeleteSecret(name: string): Promise<DeleteSecretPoller> {
    if (!this.secrets.has(name)) {
      throw notFound(name);
    }

    return {
      pollUntilDone: async () => {
        this.deleted.add(name);
      },
    };
  }

  async purgeDeletedSecret(name: string): Promise<void> {
    if (!this.deleted.has(name)) {
      throw new Error(`Secret "${name}" must be deleted before it is purged`);
    }
    this.secrets.delete(name);
    this.deleted.delete(name);
  }

  async *listPropertiesOfSecretVersions(
    name: string,
  ): AsyncIterable<SecretProperties> {
    for (const stored of this.secrets.get(name) ?? []) {
      yield { ...stored.properties };
    }
  }
}
