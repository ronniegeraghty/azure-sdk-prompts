from __future__ import annotations

import asyncio
import unittest
from dataclasses import dataclass
from datetime import datetime, timedelta, timezone
from types import SimpleNamespace
from typing import Any

from azure.core.exceptions import ResourceNotFoundError

from secret_config.cache import AsyncSecretCache, SecretCache
from secret_config.providers import AsyncSecretProvider, SecretProvider
from secret_config.rotation import AsyncSecretRotator, SecretRotator


def not_found() -> ResourceNotFoundError:
    return ResourceNotFoundError("not found")


@dataclass
class StoredSecret:
    value: str
    version: str = "v1"
    expires_on: datetime | None = None

    @property
    def properties(self) -> Any:
        return SimpleNamespace(
            version=self.version,
            expires_on=self.expires_on,
        )


class FakeClient:
    def __init__(self, secrets: dict[tuple[str, str | None], StoredSecret]) -> None:
        self.secrets = secrets
        self.calls: list[tuple[str, str | None]] = []

    def get_secret(
        self,
        name: str,
        version: str | None = None,
    ) -> StoredSecret:
        self.calls.append((name, version))
        try:
            return self.secrets[(name, version)]
        except KeyError as error:
            raise not_found() from error


class AsyncFakeClient(FakeClient):
    async def get_secret(
        self,
        name: str,
        version: str | None = None,
    ) -> StoredSecret:
        return super().get_secret(name, version)


class Poller:
    def __init__(self, events: list[str]) -> None:
        self.events = events

    def result(self) -> None:
        self.events.append("delete-complete")


class RotationClient:
    def __init__(self) -> None:
        self.events: list[str] = []

    def begin_delete_secret(self, name: str) -> Poller:
        self.events.append("delete-started")
        return Poller(self.events)

    def purge_deleted_secret(self, name: str) -> None:
        self.events.append("purged")

    def get_deleted_secret(self, name: str) -> None:
        self.events.append("purge-checked")
        raise not_found()

    def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime,
    ) -> StoredSecret:
        self.events.append("created")
        return StoredSecret(value, expires_on=expires_on)


class AsyncRotationClient(RotationClient):
    async def delete_secret(self, name: str) -> None:
        self.events.append("delete-started")
        self.events.append("delete-complete")

    async def purge_deleted_secret(self, name: str) -> None:
        super().purge_deleted_secret(name)

    async def get_deleted_secret(self, name: str) -> None:
        super().get_deleted_secret(name)

    async def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime,
    ) -> StoredSecret:
        return super().set_secret(name, value, expires_on=expires_on)


class SecretConfigTests(unittest.TestCase):
    def test_provider_handles_missing_and_specific_version(self) -> None:
        client = FakeClient(
            {("setting", "v2"): StoredSecret("versioned", version="v2")}
        )
        provider = SecretProvider(client)

        self.assertEqual(provider.get_value("missing", "fallback"), "fallback")
        result = provider.get("setting", version="v2")

        self.assertEqual(result.value, "versioned")
        self.assertEqual(result.version, "v2")
        self.assertTrue(result.found)

    def test_cache_loads_refreshes_and_refetches_near_expiry(self) -> None:
        expires_soon = datetime.now(timezone.utc) + timedelta(days=1)
        client = FakeClient(
            {
                ("stable", None): StoredSecret("one"),
                ("expiring", None): StoredSecret(
                    "short-lived",
                    expires_on=expires_soon,
                ),
            }
        )
        cache = SecretCache(
            SecretProvider(client),
            expiry_warning_window=timedelta(days=7),
        )

        cache.load_required({"stable": None, "expiring": None})
        cache.get("stable")
        cache.get("expiring")
        cache.refresh("stable")

        self.assertEqual(client.calls.count(("stable", None)), 2)
        self.assertEqual(client.calls.count(("expiring", None)), 2)
        self.assertIn("expiring", cache.expiring())

    def test_rotation_waits_for_delete_and_purge(self) -> None:
        client = RotationClient()
        expiry = datetime.now(timezone.utc) + timedelta(days=90)

        SecretRotator(client).rotate("setting", "new", expiry)

        self.assertEqual(
            client.events,
            [
                "delete-started",
                "delete-complete",
                "purged",
                "purge-checked",
                "created",
            ],
        )


class AsyncSecretConfigTests(unittest.IsolatedAsyncioTestCase):
    async def test_async_provider_cache_and_rotation(self) -> None:
        expires_soon = datetime.now(timezone.utc) + timedelta(days=1)
        client = AsyncFakeClient(
            {
                ("setting", "v2"): StoredSecret("versioned", version="v2"),
                ("expiring", None): StoredSecret(
                    "short-lived",
                    expires_on=expires_soon,
                ),
            }
        )
        provider = AsyncSecretProvider(client)

        self.assertEqual(
            await provider.get_value("missing", "fallback"),
            "fallback",
        )
        self.assertEqual(
            (await provider.get("setting", version="v2")).value,
            "versioned",
        )

        cache = AsyncSecretCache(provider)
        await cache.load_required({"expiring": None})
        await cache.get("expiring")
        self.assertEqual(client.calls.count(("expiring", None)), 2)

        rotation_client = AsyncRotationClient()
        rotator = AsyncSecretRotator(rotation_client)
        expiry = datetime.now(timezone.utc) + timedelta(days=90)
        await rotator.rotate("setting", "new", expiry)
        self.assertEqual(rotation_client.events[-1], "created")


if __name__ == "__main__":
    unittest.main()
