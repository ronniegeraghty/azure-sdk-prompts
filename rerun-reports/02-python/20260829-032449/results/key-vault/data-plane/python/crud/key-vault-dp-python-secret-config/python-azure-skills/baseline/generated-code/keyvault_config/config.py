"""Secure Key Vault client factories."""

from __future__ import annotations

import os
from dataclasses import dataclass
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

VAULT_URL_ENV = "AZURE_KEY_VAULT_URL"


def get_vault_url() -> str:
    vault_url = os.getenv(VAULT_URL_ENV)
    if not vault_url:
        raise RuntimeError(f"{VAULT_URL_ENV} must be set")
    parsed = urlparse(vault_url)
    if (
        parsed.scheme != "https"
        or not parsed.hostname
        or parsed.username
        or parsed.password
        or parsed.query
        or parsed.fragment
        or parsed.path not in ("", "/")
    ):
        raise ValueError(f"{VAULT_URL_ENV} must be a valid Key Vault HTTPS URL")
    return vault_url


@dataclass(slots=True)
class KeyVaultResources:
    client: SecretClient
    credential: DefaultAzureCredential

    def close(self) -> None:
        self.client.close()
        self.credential.close()


@dataclass(slots=True)
class AsyncKeyVaultResources:
    client: AsyncSecretClient
    credential: AsyncDefaultAzureCredential

    async def close(self) -> None:
        await self.client.close()
        await self.credential.close()


def create_key_vault_resources() -> KeyVaultResources:
    vault_url = get_vault_url()
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=vault_url, credential=credential)
    return KeyVaultResources(client, credential)


def create_async_key_vault_resources() -> AsyncKeyVaultResources:
    vault_url = get_vault_url()
    credential = AsyncDefaultAzureCredential()
    client = AsyncSecretClient(vault_url=vault_url, credential=credential)
    return AsyncKeyVaultResources(client, credential)
