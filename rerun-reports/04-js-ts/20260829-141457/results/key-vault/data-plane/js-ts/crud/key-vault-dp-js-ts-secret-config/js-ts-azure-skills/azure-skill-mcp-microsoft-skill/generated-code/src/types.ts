import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";

export interface DeleteSecretPoller {
  pollUntilDone(): Promise<unknown>;
}

export interface SecretClientLike {
  getSecret(name: string, options?: GetSecretOptions): Promise<KeyVaultSecret>;
  setSecret(
    name: string,
    value: string,
    options?: SetSecretOptions,
  ): Promise<KeyVaultSecret>;
  beginDeleteSecret(name: string): Promise<DeleteSecretPoller>;
  purgeDeletedSecret(name: string): Promise<void>;
}

export interface SecretSnapshot {
  name: string;
  value: string;
  found: boolean;
  version?: string;
  expiresOn?: Date;
}

export interface ExpiryInspection {
  name: string;
  expiresOn?: Date;
  isExpired: boolean;
  isNearExpiry: boolean;
  millisecondsRemaining?: number;
}
