import type {
  KeyVaultSecret,
  SecretClient,
  SetSecretOptions,
} from "@azure/keyvault-secrets";

export interface RotateSecretOptions {
  expiresOn: Date;
  notBefore?: Date;
  tags?: Record<string, string>;
}

export class SecretRotationHelper {
  public constructor(private readonly client: SecretClient) {}

  public async rotateSecret(
    name: string,
    value: string,
    options: RotateSecretOptions,
  ): Promise<KeyVaultSecret> {
    const setOptions: SetSecretOptions = {
      expiresOn: options.expiresOn,
      ...(options.notBefore === undefined
        ? {}
        : { notBefore: options.notBefore }),
      ...(options.tags === undefined ? {} : { tags: options.tags }),
    };

    return this.client.setSecret(name, value, setOptions);
  }

  public async deleteAndPurgeSecret(name: string): Promise<void> {
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }
}
