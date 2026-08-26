from __future__ import annotations

import os
from contextlib import AbstractAsyncContextManager, AbstractContextManager
from datetime import timedelta
from types import TracebackType

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from .cache import AsyncSecretCache, SecretCache
from .providers import AsyncSecretProvider, SecretProvider
from .rotation import AsyncSecretRotator, SecretRotator

VAULT_URL_ENVIRONMENT_VARIABLE = "AZURE_KEYVAULT_URL"


def get_vault_url() -> str:
    try:
        vault_url = os.environ[VAULT_URL_ENVIRONMENT_VARIABLE]
    except KeyError as error:
        raise RuntimeError(
            f"{VAULT_URL_ENVIRONMENT_VARIABLE} must contain the Key Vault URL"
        ) from error
    if not vault_url.startswith("https://"):
        raise ValueError("The Key Vault URL must use HTTPS")
    return vault_url


class SyncConfiguration(
    AbstractContextManager["SyncConfiguration"],
):
    def __init__(self, *, warning_days: int = 7) -> None:
        self._warning_window = _warning_window(warning_days)
        self._credential: DefaultAzureCredential | None = None
        self._client: SecretClient | None = None
        self.provider: SecretProvider
        self.cache: SecretCache
        self.rotator: SecretRotator

    def __enter__(self) -> "SyncConfiguration":
        credential = DefaultAzureCredential()
        client = SecretClient(
            vault_url=get_vault_url(),
            credential=credential,
        )
        try:
            credential.__enter__()
            client.__enter__()
        except BaseException:
            client.close()
            credential.close()
            raise

        self._credential = credential
        self._client = client
        self.provider = SecretProvider(client)
        self.cache = SecretCache(
            self.provider,
            expiry_warning_window=self._warning_window,
        )
        self.rotator = SecretRotator(client)
        return self

    def __exit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> bool | None:
        try:
            if self._client is not None:
                self._client.close()
        finally:
            if self._credential is not None:
                self._credential.close()
        return None


class AsyncConfiguration(
    AbstractAsyncContextManager["AsyncConfiguration"],
):
    def __init__(self, *, warning_days: int = 7) -> None:
        self._warning_window = _warning_window(warning_days)
        self._credential: AsyncDefaultAzureCredential | None = None
        self._client: AsyncSecretClient | None = None
        self.provider: AsyncSecretProvider
        self.cache: AsyncSecretCache
        self.rotator: AsyncSecretRotator

    async def __aenter__(self) -> "AsyncConfiguration":
        credential = AsyncDefaultAzureCredential()
        client = AsyncSecretClient(
            vault_url=get_vault_url(),
            credential=credential,
        )
        try:
            await credential.__aenter__()
            await client.__aenter__()
        except BaseException:
            await client.close()
            await credential.close()
            raise

        self._credential = credential
        self._client = client
        self.provider = AsyncSecretProvider(client)
        self.cache = AsyncSecretCache(
            self.provider,
            expiry_warning_window=self._warning_window,
        )
        self.rotator = AsyncSecretRotator(client)
        return self

    async def __aexit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> bool | None:
        try:
            if self._client is not None:
                await self._client.close()
        finally:
            if self._credential is not None:
                await self._credential.close()
        return None


def _warning_window(warning_days: int) -> timedelta:
    if warning_days < 0:
        raise ValueError("warning_days cannot be negative")
    return timedelta(days=warning_days)
