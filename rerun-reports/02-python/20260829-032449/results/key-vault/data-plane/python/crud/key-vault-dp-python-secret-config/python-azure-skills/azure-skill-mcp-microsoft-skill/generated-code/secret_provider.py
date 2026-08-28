from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timedelta, timezone
from typing import Any

from azure.core.exceptions import ResourceNotFoundError


def _as_utc(value: datetime) -> datetime:
    if value.tzinfo is None:
        return value.replace(tzinfo=timezone.utc)
    return value.astimezone(timezone.utc)


@dataclass(frozen=True, slots=True)
class SecretRecord:
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
        current_time = _as_utc(now or datetime.now(timezone.utc))
        return _as_utc(self.expires_on) <= current_time + warning_window

    @property
    def is_expired(self) -> bool:
        return self.expires_within(timedelta(0))


class SyncSecretProvider:
    def __init__(self, client: Any) -> None:
        self._client = client

    @property
    def client(self) -> Any:
        return self._client

    def get_secret_record(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretRecord:
        try:
            secret = self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretRecord(name, default, version, None, False)

        return SecretRecord(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    def get_secret(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> str | None:
        return self.get_secret_record(name, default, version=version).value

    def get_expiry(
        self,
        name: str,
        *,
        version: str | None = None,
    ) -> datetime | None:
        return self.get_secret_record(name, version=version).expires_on


class AsyncSecretProvider:
    def __init__(self, client: Any) -> None:
        self._client = client

    @property
    def client(self) -> Any:
        return self._client

    async def get_secret_record(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretRecord:
        try:
            secret = await self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretRecord(name, default, version, None, False)

        return SecretRecord(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    async def get_secret(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> str | None:
        record = await self.get_secret_record(name, default, version=version)
        return record.value

    async def get_expiry(
        self,
        name: str,
        *,
        version: str | None = None,
    ) -> datetime | None:
        record = await self.get_secret_record(name, version=version)
        return record.expires_on
