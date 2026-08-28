export interface SecretPropertiesLike {
  name: string;
  version?: string;
  expiresOn?: Date;
}

export interface KeyVaultSecretLike {
  value?: string;
  properties: SecretPropertiesLike;
}

export interface GetSecretOptionsLike {
  version?: string;
}

export interface SetSecretOptionsLike {
  expiresOn?: Date;
}

export interface DeleteSecretPollerLike {
  pollUntilDone(): Promise<unknown>;
}

export interface SecretClientLike {
  getSecret(name: string, options?: GetSecretOptionsLike): Promise<KeyVaultSecretLike>;
  setSecret(
    name: string,
    value: string,
    options?: SetSecretOptionsLike,
  ): Promise<KeyVaultSecretLike>;
  beginDeleteSecret(name: string): Promise<DeleteSecretPollerLike>;
  purgeDeletedSecret(name: string): Promise<void>;
}
