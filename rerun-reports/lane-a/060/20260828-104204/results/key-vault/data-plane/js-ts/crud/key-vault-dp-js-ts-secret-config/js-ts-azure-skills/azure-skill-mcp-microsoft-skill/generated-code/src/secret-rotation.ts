import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";

export interface SecretDeletePoller {
  pollUntilDone(): Promise<unknown>;
}

export interface SecretRotationClient {
  getSecret(name: string, options?: GetSecretOptions): Promise<KeyVaultSecret>;
  setSecret(name: string, value: string, options?: SetSecretOptions): Promise<KeyVaultSecret>;
  beginDeleteSecret(name: string): Promise<SecretDeletePoller>;
  purgeDeletedSecret(name: string): Promise<void>;
}

export interface RotationResult {
  name: string;
  previousVersion?: string;
  newVersion?: string;
  expiresOn: Date;
}

export interface PermanentCleanupConfirmation {
  confirmPermanentDeletion: string;
}

export class SecretRotationHelper {
  public constructor(private readonly client: SecretRotationClient) {}

  public async rotateSecret(
    name: string,
    newValue: string,
    expiresOn: Date,
  ): Promise<RotationResult> {
    if (expiresOn.getTime() <= Date.now()) {
      throw new Error("The rotated secret expiry date must be in the future.");
    }

    const current = await this.getCurrentSecretIfPresent(name);
    const options: SetSecretOptions = {
      enabled: true,
      expiresOn,
      ...(current?.properties.contentType === undefined
        ? {}
        : { contentType: current.properties.contentType }),
      tags: {
        ...current?.properties.tags,
        rotatedOn: new Date().toISOString(),
      },
    };
    const rotated = await this.client.setSecret(name, newValue, options);

    return {
      name,
      expiresOn,
      ...(current?.properties.version === undefined
        ? {}
        : { previousVersion: current.properties.version }),
      ...(rotated.properties.version === undefined
        ? {}
        : { newVersion: rotated.properties.version }),
    };
  }

  public async deleteAndPurgeSecret(
    name: string,
    confirmation: PermanentCleanupConfirmation,
  ): Promise<void> {
    if (confirmation.confirmPermanentDeletion !== name) {
      throw new Error(
        `Permanent deletion was not confirmed. Set confirmPermanentDeletion to '${name}'.`,
      );
    }

    // Deletion is name-scoped and removes every version, so wait for soft-delete to finish.
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }

  private async getCurrentSecretIfPresent(name: string): Promise<KeyVaultSecret | undefined> {
    try {
      return await this.client.getSecret(name);
    } catch (error: unknown) {
      if (isSecretNotFoundError(error)) {
        return undefined;
      }
      throw error;
    }
  }
}

function isSecretNotFoundError(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}
