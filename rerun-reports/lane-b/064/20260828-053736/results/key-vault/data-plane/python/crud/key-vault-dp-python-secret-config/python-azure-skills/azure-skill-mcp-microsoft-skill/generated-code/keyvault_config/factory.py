from __future__ import annotations

import os
from collections.abc import AsyncIterator, Iterator
from contextlib import asynccontextmanager, contextmanager
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

VAULT_URL_ENV = "AZURE_KEYVAULT_URL"


def get_vault_url() -> str:
    vault_url = os.environ.get(VAULT_URL_ENV)
    if not vault_url:
        raise RuntimeError(f"{VAULT_URL_ENV} must be set")

    parsed = urlparse(vault_url)
    if parsed.scheme != "https" or not parsed.netloc:
        raise ValueError(f"{VAULT_URL_ENV} must be a valid HTTPS URL")
    return vault_url


@contextmanager
def open_secret_client() -> Iterator[SecretClient]:
    credential = DefaultAzureCredential()
    try:
        with SecretClient(
            vault_url=get_vault_url(), credential=credential
        ) as client:
            yield client
    finally:
        credential.close()


@asynccontextmanager
async def open_async_secret_client() -> AsyncIterator[AsyncSecretClient]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncSecretClient(
            vault_url=get_vault_url(), credential=credential
        ) as client:
            yield client
