import type {
  DeletedSecret,
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";
import { RestError } from "@azure/core-rest-pipeline";

import type {
  DeleteSecretPoller,
  KeyVaultClient,
} from "./key-vault-client.js";

interface StoredSecret {
  value: string;
  version: string;
  expiresOn?: Date;
}

export class InMemoryKeyVaultClient implements KeyVaultClient {
  private readonly secrets = new Map<string, StoredSecret[]>();
  private readonly deletedNames = new Set<string>();
  private nextVersion = 1;

  public async getSecret(
    name: string,
    options: GetSecretOptions = {},
  ): Promise<KeyVaultSecret> {
    const versions = this.secrets.get(name) ?? [];
    const stored =
      options.version === undefined
        ? versions.at(-1)
        : versions.find(({ version }) => version === options.version);

    if (stored === undefined) {
      throw new RestError(`Secret "${name}" was not found.`, {
        statusCode: 404,
      });
    }

    return this.toKeyVaultSecret(name, stored);
  }

  public async setSecret(
    name: string,
    value: string,
    options: SetSecretOptions = {},
  ): Promise<KeyVaultSecret> {
    if (this.deletedNames.has(name)) {
      throw new RestError(`Secret "${name}" is soft-deleted.`, {
        statusCode: 409,
      });
    }

    const stored: StoredSecret = {
      value,
      version: String(this.nextVersion++),
      ...(options.expiresOn === undefined
        ? {}
        : { expiresOn: options.expiresOn }),
    };
    const versions = this.secrets.get(name) ?? [];
    versions.push(stored);
    this.secrets.set(name, versions);
    return this.toKeyVaultSecret(name, stored);
  }

  public async beginDeleteSecret(name: string): Promise<DeleteSecretPoller> {
    if (!this.secrets.has(name)) {
      throw new RestError(`Secret "${name}" was not found.`, {
        statusCode: 404,
      });
    }

    this.secrets.delete(name);
    this.deletedNames.add(name);
    const deletedSecret = {
      name,
      properties: { name, vaultUrl: "https://local.vault.invalid" },
    } as DeletedSecret;

    return {
      pollUntilDone: async () => deletedSecret,
    };
  }

  public async purgeDeletedSecret(name: string): Promise<void> {
    if (!this.deletedNames.delete(name)) {
      throw new RestError(`Deleted secret "${name}" was not found.`, {
        statusCode: 404,
      });
    }
  }

  private toKeyVaultSecret(
    name: string,
    stored: StoredSecret,
  ): KeyVaultSecret {
    return {
      name,
      value: stored.value,
      properties: {
        name,
        vaultUrl: "https://local.vault.invalid",
        version: stored.version,
        ...(stored.expiresOn === undefined
          ? {}
          : { expiresOn: stored.expiresOn }),
      },
    };
  }
}
