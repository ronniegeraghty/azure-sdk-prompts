import type { KeyVaultSecret } from "@azure/keyvault-secrets";

import type { KeyVaultClient } from "./key-vault-client.js";

export interface RotateSecretOptions {
  expiresOn: Date;
  cleanupForFullNameReuse?: boolean;
}

export interface RotationResult {
  secret: KeyVaultSecret;
  deletedAndPurged: boolean;
}

export class SecretRotator {
  public constructor(private readonly client: KeyVaultClient) {}

  public async rotate(
    name: string,
    value: string,
    options: RotateSecretOptions,
  ): Promise<RotationResult> {
    if (options.expiresOn.getTime() <= Date.now()) {
      throw new RangeError("The rotated secret expiry must be in the future.");
    }

    let secret = await this.client.setSecret(name, value, {
      expiresOn: options.expiresOn,
    });

    if (options.cleanupForFullNameReuse !== true) {
      return { secret, deletedAndPurged: false };
    }

    // Key Vault deletion operates on the secret name and therefore removes every
    // version, including the one just created. Recreate that value after purge.
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
    secret = await this.client.setSecret(name, value, {
      expiresOn: options.expiresOn,
    });

    return { secret, deletedAndPurged: true };
  }
}
