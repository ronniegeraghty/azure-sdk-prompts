export interface SecretPropertiesLike {
  version?: string;
  expiresOn?: Date;
}

export interface SecretLike {
  name: string;
  value?: string;
  properties: SecretPropertiesLike;
}

export interface SetSecretOptionsLike {
  expiresOn?: Date;
  enabled?: boolean;
  contentType?: string;
  tags?: Record<string, string>;
}

export interface DeletePollerLike {
  pollUntilDone(): Promise<unknown>;
}

export interface SecretStore {
  getSecret(name: string, options?: { version?: string }): Promise<SecretLike>;
  setSecret(
    name: string,
    value: string,
    options?: SetSecretOptionsLike,
  ): Promise<SecretLike>;
  beginDeleteSecret(name: string): Promise<DeletePollerLike>;
  purgeDeletedSecret(name: string): Promise<void>;
}
