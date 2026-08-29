import type { KeyVaultSecret } from "@azure/keyvault-secrets";

import type { SecretStore } from "./secret-store.js";

export interface RotateSecretOptions {
  expiresOn: Date;
  contentType?: string;
  tags?: Record<string, string>;
}

export interface RotationResult {
  secret: KeyVaultSecret;
  previousVersions: string[];
}

export class SecretRotationHelper {
  constructor(private readonly client: SecretStore) {}

  async rotate(
    name: string,
    value: string,
    options: RotateSecretOptions,
  ): Promise<RotationResult> {
    if (options.expiresOn.getTime() <= Date.now()) {
      throw new RangeError("The new secret expiry date must be in the future");
    }

    const previousVersions: string[] = [];
    for await (const properties of this.client.listPropertiesOfSecretVersions(
      name,
    )) {
      if (properties.version !== undefined) {
        previousVersions.push(properties.version);
      }
    }

    const setOptions = {
      expiresOn: options.expiresOn,
      ...(options.contentType === undefined
        ? {}
        : { contentType: options.contentType }),
      ...(options.tags === undefined ? {} : { tags: options.tags }),
    };
    const secret = await this.client.setSecret(name, value, setOptions);

    return { secret, previousVersions };
  }

  async deleteAndPurge(name: string): Promise<void> {
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }
}
