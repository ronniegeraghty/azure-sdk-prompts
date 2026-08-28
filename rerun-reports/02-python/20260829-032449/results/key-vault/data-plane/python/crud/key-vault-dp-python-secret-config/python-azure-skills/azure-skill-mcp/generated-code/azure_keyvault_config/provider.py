"""Synchronous and asynchronous Key Vault secret providers."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Optional

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


@dataclass(frozen=True)
class SecretDetails:
    """A secret value and the metadata needed by the configuration cache."""

    name: str
    value: Optional[str]
    version: Optional[str]
    expires_on: Optional[datetime]
    found: bool


class SecretProvider:
    """Read secrets through the synchronous Azure Key Vault client."""

    def __init__(self, client: SecretClient, credential: object | None = None) -> None:
        self.client = client
        self._credential = credential

    def get_secret(
        self, name: str, default: Optional[str] = None, version: Optional[str] = None
    ) -> Optional[str]:
        return self.get_secret_details(name, default=default, version=version).value

    def get_secret_details(
        self, name: str, default: Optional[str] = None, version: Optional[str] = None
    ) -> SecretDetails:
        try:
            secret = self.client.get_secret(name, version)
        except ResourceNotFoundError:
            return SecretDetails(name, default, version, None, False)

        return SecretDetails(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    def get_expiry(
        self, name: str, version: Optional[str] = None
    ) -> Optional[datetime]:
        return self.get_secret_details(name, version=version).expires_on

    def close(self) -> None:
        self.client.close()
        close = getattr(self._credential, "close", None)
        if close is not None:
            close()

    def __enter__(self) -> "SecretProvider":
        return self

    def __exit__(self, exc_type: object, exc_value: object, traceback: object) -> None:
        self.close()


class AsyncSecretProvider:
    """Read secrets through the asynchronous Azure Key Vault client."""

    def __init__(
        self, client: AsyncSecretClient, credential: object | None = None
    ) -> None:
        self.client = client
        self._credential = credential

    async def get_secret(
        self, name: str, default: Optional[str] = None, version: Optional[str] = None
    ) -> Optional[str]:
        details = await self.get_secret_details(name, default=default, version=version)
        return details.value

    async def get_secret_details(
        self, name: str, default: Optional[str] = None, version: Optional[str] = None
    ) -> SecretDetails:
        try:
            secret = await self.client.get_secret(name, version)
        except ResourceNotFoundError:
            return SecretDetails(name, default, version, None, False)

        return SecretDetails(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    async def get_expiry(
        self, name: str, version: Optional[str] = None
    ) -> Optional[datetime]:
        details = await self.get_secret_details(name, version=version)
        return details.expires_on

    async def close(self) -> None:
        await self.client.close()
        close = getattr(self._credential, "close", None)
        if close is not None:
            await close()

    async def __aenter__(self) -> "AsyncSecretProvider":
        return self

    async def __aexit__(
        self, exc_type: object, exc_value: object, traceback: object
    ) -> None:
        await self.close()
