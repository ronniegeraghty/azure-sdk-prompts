from __future__ import annotations

import os
from contextlib import asynccontextmanager, contextmanager
from collections.abc import AsyncIterator, Iterator

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from secret_provider import AsyncSecretProvider, SyncSecretProvider

VAULT_URL_ENV = "AZURE_KEYVAULT_URL"


def get_vault_url() -> str:
    vault_url = os.getenv(VAULT_URL_ENV)
    if not vault_url:
        raise RuntimeError(f"{VAULT_URL_ENV} must contain the Key Vault URL")
    if not vault_url.startswith("https://"):
        raise ValueError(f"{VAULT_URL_ENV} must use HTTPS")
    return vault_url


@contextmanager
def create_sync_provider() -> Iterator[SyncSecretProvider]:
    credential = DefaultAzureCredential()
    try:
        with SecretClient(
            vault_url=get_vault_url(),
            credential=credential,
        ) as client:
            yield SyncSecretProvider(client)
    finally:
        credential.close()


@asynccontextmanager
async def create_async_provider() -> AsyncIterator[AsyncSecretProvider]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncSecretClient(
            vault_url=get_vault_url(),
            credential=credential,
        ) as client:
            yield AsyncSecretProvider(client)
