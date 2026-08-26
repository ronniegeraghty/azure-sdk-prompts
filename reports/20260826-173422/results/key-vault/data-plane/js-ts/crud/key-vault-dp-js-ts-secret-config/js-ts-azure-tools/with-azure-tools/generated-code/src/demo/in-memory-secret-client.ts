import type {
  DeletePollerLike,
  SecretLike,
  SecretStore,
  SetSecretOptionsLike,
} from "../secret-store.js";

interface StoredSecret extends SecretLike {
  value: string;
}

export class InMemorySecretClient implements SecretStore {
  private readonly active = new Map<string, StoredSecret[]>();
  private readonly deleted = new Set<string>();
  private versionCounter = 0;

  public async getSecret(
    name: string,
    options?: { version?: string },
  ): Promise<StoredSecret> {
    const versions = this.active.get(name);
    const secret =
      options?.version === undefined
        ? versions?.at(-1)
        : versions?.find(
            ({ properties }) => properties.version === options.version,
          );

    if (secret === undefined) {
      throw notFound(name);
    }
    return clone(secret);
  }

  public async setSecret(
    name: string,
    value: string,
    options: SetSecretOptionsLike = {},
  ): Promise<StoredSecret> {
    if (this.deleted.has(name)) {
      throw new Error(`Secret '${name}' is soft-deleted and must be purged`);
    }

    const version = `version-${++this.versionCounter}`;
    const secret: StoredSecret = {
      name,
      value,
      properties: {
        version,
        ...(options.expiresOn === undefined
          ? {}
          : { expiresOn: new Date(options.expiresOn) }),
      },
    };
    const versions = this.active.get(name) ?? [];
    versions.push(secret);
    this.active.set(name, versions);
    return clone(secret);
  }

  public async beginDeleteSecret(name: string): Promise<DeletePollerLike> {
    if (!this.active.has(name)) {
      throw notFound(name);
    }

    return {
      pollUntilDone: async () => {
        this.active.delete(name);
        this.deleted.add(name);
      },
    };
  }

  public async purgeDeletedSecret(name: string): Promise<void> {
    if (!this.deleted.delete(name)) {
      throw notFound(name);
    }
  }
}

function clone(secret: StoredSecret): StoredSecret {
  return {
    ...secret,
    properties: {
      ...secret.properties,
      ...(secret.properties.expiresOn === undefined
        ? {}
        : { expiresOn: new Date(secret.properties.expiresOn) }),
    },
  };
}

function notFound(name: string): Error & { statusCode: number; code: string } {
  return Object.assign(new Error(`Secret '${name}' was not found`), {
    statusCode: 404,
    code: "SecretNotFound",
  });
}
