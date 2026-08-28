from __future__ import annotations

import os

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from .provider import AsyncSecretProvider, SecretProvider

VAULT_URL_ENV = "AZURE_KEY_VAULT_URL"


def _vault_url() -> str:
    vault_url = os.environ.get(VAULT_URL_ENV)
    if not vault_url:
        raise RuntimeError(f"{VAULT_URL_ENV} must be set")
    if not vault_url.startswith("https://"):
        raise ValueError(f"{VAULT_URL_ENV} must use HTTPS")
    return vault_url


def create_sync_provider() -> SecretProvider:
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=_vault_url(), credential=credential)
    return SecretProvider(client, credential)


def create_async_provider() -> AsyncSecretProvider:
    credential = AsyncDefaultAzureCredential()
    client = AsyncSecretClient(vault_url=_vault_url(), credential=credential)
    return AsyncSecretProvider(client, credential)
