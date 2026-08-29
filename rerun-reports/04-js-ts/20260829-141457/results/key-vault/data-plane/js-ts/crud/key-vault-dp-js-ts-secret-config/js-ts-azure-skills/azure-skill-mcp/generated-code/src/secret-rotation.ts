import type {
  KeyVaultSecret,
  SetSecretOptions
} from "@azure/keyvault-secrets";

export interface DeleteSecretPollerLike {
  pollUntilDone(): Promise<unknown>;
}

export interface SecretWriter {
  setSecret(name: string, value: string, options?: SetSecretOptions): Promise<KeyVaultSecret>;
  beginDeleteSecret(name: string): Promise<DeleteSecretPollerLike>;
  purgeDeletedSecret(name: string): Promise<void>;
}

export interface RotationResult {
  name: string;
  version?: string;
  expiresOn: Date;
}

export class SecretRotationHelper {
  public constructor(private readonly client: SecretWriter) {}

  public async rotateSecret(
    name: string,
    value: string,
    expiresOn: Date
  ): Promise<RotationResult> {
    if (!name.trim()) {
      throw new Error("Secret name must not be empty.");
    }
    if (!value) {
      throw new Error("Rotated secret value must not be empty.");
    }
    if (expiresOn.getTime() <= Date.now()) {
      throw new Error("Rotated secret expiry must be in the future.");
    }

    const created = await this.client.setSecret(name, value, { expiresOn });
    return {
      name: created.name,
      version: created.properties.version,
      expiresOn
    };
  }

  public async deleteAndPurgeSecret(
    name: string,
    confirmPermanentPurge: boolean
  ): Promise<void> {
    if (!confirmPermanentPurge) {
      throw new Error(
        "Permanent purge was not confirmed. Deletion would remove every version of this secret."
      );
    }

    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }
}
