import type { SecretClientLike } from "./secret-client.js";

export interface RotationResult {
  name: string;
  version?: string;
  expiresOn: Date;
}

export class SecretRotationHelper {
  public constructor(private readonly client: SecretClientLike) {}

  public async rotateSecret(
    name: string,
    newValue: string,
    expiresOn: Date,
    now = new Date(),
  ): Promise<RotationResult> {
    if (!name.trim()) {
      throw new Error("Secret name must not be empty.");
    }
    if (!newValue) {
      throw new Error("Secret value must not be empty.");
    }
    if (expiresOn.getTime() <= now.getTime()) {
      throw new Error("The new secret expiry date must be in the future.");
    }

    const secret = await this.client.setSecret(name, newValue, { expiresOn });
    return {
      name,
      expiresOn,
      ...(secret.properties.version !== undefined
        ? { version: secret.properties.version }
        : {}),
    };
  }

  /**
   * Permanently removes the secret name and every version under it.
   * Key Vault does not support deleting only one historical secret version.
   */
  public async deleteAndPurgeSecret(name: string, confirmation: string): Promise<void> {
    if (confirmation !== name) {
      throw new Error(`Permanent deletion requires confirmation equal to "${name}".`);
    }

    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }
}
