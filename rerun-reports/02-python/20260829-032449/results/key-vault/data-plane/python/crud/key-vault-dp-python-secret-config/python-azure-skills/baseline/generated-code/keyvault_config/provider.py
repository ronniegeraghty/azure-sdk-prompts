"""Synchronous and asynchronous Azure Key Vault secret providers."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Any

from azure.core.exceptions import ResourceNotFoundError


@dataclass(frozen=True, slots=True)
class SecretValue:
    """A secret value and the metadata needed by configuration consumers."""

    name: str
    value: str | None
    version: str | None
    expires_on: datetime | None
    found: bool


class SecretProvider:
    """Retrieve Key Vault secrets without treating missing values as errors."""

    def __init__(self, client: Any) -> None:
        self._client = client

    def get_secret(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretValue:
        try:
            secret = self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretValue(name, default, version, None, False)

        return SecretValue(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    def get_value(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> str | None:
        return self.get_secret(name, default, version=version).value


class AsyncSecretProvider:
    """Asynchronously retrieve Key Vault secrets with missing-value defaults."""

    def __init__(self, client: Any) -> None:
        self._client = client

    async def get_secret(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretValue:
        try:
            secret = await self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretValue(name, default, version, None, False)

        return SecretValue(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    async def get_value(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> str | None:
        secret = await self.get_secret(name, default, version=version)
        return secret.value
