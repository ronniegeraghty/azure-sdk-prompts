from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timedelta, timezone
from typing import Protocol

from azure.core.exceptions import ResourceNotFoundError


@dataclass(frozen=True, slots=True)
class SecretResult:
    name: str
    value: str | None
    version: str | None
    expires_on: datetime | None
    found: bool

    def expires_within(
        self,
        warning_window: timedelta,
        *,
        now: datetime | None = None,
    ) -> bool:
        if self.expires_on is None:
            return False
        current_time = now or datetime.now(timezone.utc)
        expiry = self.expires_on
        if expiry.tzinfo is None:
            expiry = expiry.replace(tzinfo=timezone.utc)
        return expiry <= current_time + warning_window


class SecretProperties(Protocol):
    version: str | None
    expires_on: datetime | None


class KeyVaultSecret(Protocol):
    value: str | None
    properties: SecretProperties


class SyncSecretClient(Protocol):
    def get_secret(
        self,
        name: str,
        version: str | None = None,
    ) -> KeyVaultSecret: ...


class AsyncSecretClient(Protocol):
    async def get_secret(
        self,
        name: str,
        version: str | None = None,
    ) -> KeyVaultSecret: ...


class SecretProvider:
    def __init__(self, client: SyncSecretClient) -> None:
        self._client = client

    def get(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretResult:
        try:
            secret = self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretResult(name, default, version, None, False)

        return SecretResult(
            name=name,
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
        return self.get(name, default, version=version).value


class AsyncSecretProvider:
    def __init__(self, client: AsyncSecretClient) -> None:
        self._client = client

    async def get(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretResult:
        try:
            secret = await self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretResult(name, default, version, None, False)

        return SecretResult(
            name=name,
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
        return (await self.get(name, default, version=version)).value
