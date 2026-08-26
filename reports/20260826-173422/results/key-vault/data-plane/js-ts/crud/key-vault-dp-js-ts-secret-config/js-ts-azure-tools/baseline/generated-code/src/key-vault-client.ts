import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";

export interface DeleteSecretPoller {
  pollUntilDone(): Promise<unknown>;
}

export interface KeyVaultClient {
  getSecret(name: string, options?: GetSecretOptions): Promise<KeyVaultSecret>;
  setSecret(
    name: string,
    value: string,
    options?: SetSecretOptions,
  ): Promise<KeyVaultSecret>;
  beginDeleteSecret(name: string): Promise<DeleteSecretPoller>;
  purgeDeletedSecret(name: string): Promise<void>;
}
