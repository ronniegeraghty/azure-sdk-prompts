"""Factories for securely authenticated Key Vault providers."""

from __future__ import annotations

import os
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from .provider import AsyncSecretProvider, SecretProvider

VAULT_URL_ENV = "AZURE_KEY_VAULT_URL"


def _vault_url_from_environment() -> str:
    vault_url = os.environ.get(VAULT_URL_ENV)
    if not vault_url:
        raise RuntimeError(f"{VAULT_URL_ENV} must contain the Key Vault URL")

    parsed = urlparse(vault_url)
    if parsed.scheme != "https" or not parsed.netloc or parsed.username:
        raise ValueError(f"{VAULT_URL_ENV} must be a valid HTTPS URL")
    return vault_url


def create_sync_provider() -> SecretProvider:
    credential = DefaultAzureCredential()
    client = SecretClient(
        vault_url=_vault_url_from_environment(), credential=credential
    )
    return SecretProvider(client, credential)


def create_async_provider() -> AsyncSecretProvider:
    credential = AsyncDefaultAzureCredential()
    client = AsyncSecretClient(
        vault_url=_vault_url_from_environment(), credential=credential
    )
    return AsyncSecretProvider(client, credential)
