from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from types import SimpleNamespace
from uuid import uuid4

from azure.core.exceptions import ResourceNotFoundError


@dataclass
class _StoredSecret:
    name: str
    value: str
    version: str
    expires_on: datetime | None

    @property
    def properties(self) -> SimpleNamespace:
        return SimpleNamespace(
            version=self.version, expires_on=self.expires_on
        )


class _CompletedPoller:
    def result(self, timeout: float | None = None) -> None:
        del timeout


class InMemorySecretClient:
    def __init__(self) -> None:
        self._active: dict[str, list[_StoredSecret]] = {}
        self._deleted: dict[str, list[_StoredSecret]] = {}

    def get_secret(
        self, name: str, version: str | None = None
    ) -> _StoredSecret:
        versions = self._active.get(name)
        if not versions:
            raise ResourceNotFoundError("Secret not found")
        if version is None:
            return versions[-1]
        for secret in versions:
            if secret.version == version:
                return secret
        raise ResourceNotFoundError("Secret version not found")

    def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime | None = None,
    ) -> _StoredSecret:
        if name in self._deleted:
            raise RuntimeError("A deleted secret must be purged first")
        secret = _StoredSecret(
            name, value, uuid4().hex, expires_on
        )
        self._active.setdefault(name, []).append(secret)
        return secret

    def begin_delete_secret(self, name: str) -> _CompletedPoller:
        versions = self._active.pop(name, None)
        if not versions:
            raise ResourceNotFoundError("Secret not found")
        self._deleted[name] = versions
        return _CompletedPoller()

    def get_deleted_secret(self, name: str) -> _StoredSecret:
        versions = self._deleted.get(name)
        if not versions:
            raise ResourceNotFoundError("Deleted secret not found")
        return versions[-1]

    def purge_deleted_secret(self, name: str) -> None:
        if name not in self._deleted:
            raise ResourceNotFoundError("Deleted secret not found")
        del self._deleted[name]


class _AsyncCompletedPoller:
    async def result(self) -> None:
        return None


class AsyncInMemorySecretClient:
    def __init__(self, client: InMemorySecretClient) -> None:
        self._client = client

    async def get_secret(
        self, name: str, version: str | None = None
    ) -> _StoredSecret:
        return self._client.get_secret(name, version)

    async def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime | None = None,
    ) -> _StoredSecret:
        return self._client.set_secret(
            name, value, expires_on=expires_on
        )

    async def begin_delete_secret(
        self, name: str
    ) -> _AsyncCompletedPoller:
        self._client.begin_delete_secret(name)
        return _AsyncCompletedPoller()

    async def get_deleted_secret(self, name: str) -> _StoredSecret:
        return self._client.get_deleted_secret(name)

    async def purge_deleted_secret(self, name: str) -> None:
        self._client.purge_deleted_secret(name)
