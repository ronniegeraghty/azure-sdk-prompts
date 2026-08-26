import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from .provider import AsyncKeyVaultSecretProvider, KeyVaultSecretProvider

VAULT_URL_ENVIRONMENT_VARIABLE = "AZURE_KEY_VAULT_URL"


def _vault_url() -> str:
    vault_url = os.environ.get(VAULT_URL_ENVIRONMENT_VARIABLE)
    if not vault_url:
        raise RuntimeError(
            f"{VAULT_URL_ENVIRONMENT_VARIABLE} must contain the Key Vault URL"
        )
    return vault_url


@dataclass
class SyncKeyVaultConfiguration:
    credential: DefaultAzureCredential
    client: SecretClient
    provider: KeyVaultSecretProvider

    def close(self) -> None:
        self.client.close()
        self.credential.close()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, traceback) -> None:
        self.close()


@dataclass
class AsyncKeyVaultConfiguration:
    credential: AsyncDefaultAzureCredential
    client: AsyncSecretClient
    provider: AsyncKeyVaultSecretProvider

    async def close(self) -> None:
        await self.client.close()
        await self.credential.close()

    async def __aenter__(self):
        return self

    async def __aexit__(self, exc_type, exc_value, traceback) -> None:
        await self.close()


def create_sync_configuration() -> SyncKeyVaultConfiguration:
    vault_url = _vault_url()
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=vault_url, credential=credential)
    return SyncKeyVaultConfiguration(
        credential, client, KeyVaultSecretProvider(client)
    )


def create_async_configuration() -> AsyncKeyVaultConfiguration:
    vault_url = _vault_url()
    credential = AsyncDefaultAzureCredential()
    client = AsyncSecretClient(vault_url=vault_url, credential=credential)
    return AsyncKeyVaultConfiguration(
        credential, client, AsyncKeyVaultSecretProvider(client)
    )
