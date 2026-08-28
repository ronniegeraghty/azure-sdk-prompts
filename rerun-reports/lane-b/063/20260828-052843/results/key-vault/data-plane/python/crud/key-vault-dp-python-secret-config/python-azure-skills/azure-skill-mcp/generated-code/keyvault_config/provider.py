from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Any

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


@dataclass(frozen=True, slots=True)
class SecretSnapshot:
    name: str
    value: str | None
    version: str | None
    expires_on: datetime | None
    found: bool


class SecretProvider:
    def __init__(
        self,
        client: SecretClient,
        credential: Any | None = None,
    ) -> None:
        self._client = client
        self._credential = credential

    @property
    def client(self) -> SecretClient:
        return self._client

    def get_secret(
        self,
        name: str,
        default: str | None = None,
        version: str | None = None,
    ) -> str | None:
        return self.get_secret_with_metadata(name, default, version).value

    def get_secret_with_metadata(
        self,
        name: str,
        default: str | None = None,
        version: str | None = None,
    ) -> SecretSnapshot:
        try:
            secret = self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretSnapshot(
                name=name,
                value=default,
                version=version,
                expires_on=None,
                found=False,
            )

        return SecretSnapshot(
            name=name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    def close(self) -> None:
        self._client.close()
        if self._credential is not None:
            self._credential.close()

    def __enter__(self) -> SecretProvider:
        return self

    def __exit__(self, *_: object) -> None:
        self.close()


class AsyncSecretProvider:
    def __init__(
        self,
        client: AsyncSecretClient,
        credential: Any | None = None,
    ) -> None:
        self._client = client
        self._credential = credential

    @property
    def client(self) -> AsyncSecretClient:
        return self._client

    async def get_secret(
        self,
        name: str,
        default: str | None = None,
        version: str | None = None,
    ) -> str | None:
        snapshot = await self.get_secret_with_metadata(name, default, version)
        return snapshot.value

    async def get_secret_with_metadata(
        self,
        name: str,
        default: str | None = None,
        version: str | None = None,
    ) -> SecretSnapshot:
        try:
            secret = await self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretSnapshot(
                name=name,
                value=default,
                version=version,
                expires_on=None,
                found=False,
            )

        return SecretSnapshot(
            name=name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    async def close(self) -> None:
        await self._client.close()
        if self._credential is not None:
            await self._credential.close()

    async def __aenter__(self) -> AsyncSecretProvider:
        return self

    async def __aexit__(self, *_: object) -> None:
        await self.close()
