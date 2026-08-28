from __future__ import annotations

import asyncio
import unittest
from datetime import datetime, timedelta, timezone
from types import SimpleNamespace
from unittest.mock import AsyncMock, Mock

from azure.core.exceptions import ResourceNotFoundError

from secret_cache import AsyncSecretCache, SyncSecretCache
from secret_provider import AsyncSecretProvider, SyncSecretProvider
from secret_rotation import rotate_secret, rotate_secret_async


def make_secret(
    name: str,
    value: str,
    *,
    version: str = "v1",
    expires_on: datetime | None = None,
) -> SimpleNamespace:
    return SimpleNamespace(
        name=name,
        value=value,
        properties=SimpleNamespace(version=version, expires_on=expires_on),
    )


class SyncSecretTests(unittest.TestCase):
    def test_specific_version_and_expiry_are_returned(self) -> None:
        expires_on = datetime.now(timezone.utc) + timedelta(days=2)
        client = Mock()
        client.get_secret.return_value = make_secret(
            "api-key",
            "value",
            version="v2",
            expires_on=expires_on,
        )
        provider = SyncSecretProvider(client)

        record = provider.get_secret_record("api-key", version="v2")

        self.assertEqual(record.version, "v2")
        self.assertEqual(record.expires_on, expires_on)
        self.assertTrue(record.expires_within(timedelta(days=7)))
        client.get_secret.assert_called_once_with("api-key", version="v2")

    def test_missing_secret_returns_default_and_version_is_forwarded(self) -> None:
        client = Mock()
        client.get_secret.side_effect = ResourceNotFoundError("missing")
        provider = SyncSecretProvider(client)

        self.assertEqual(provider.get_secret("missing", "fallback", version="v2"), "fallback")
        client.get_secret.assert_called_once_with("missing", version="v2")

    def test_cache_refreshes_near_expiry_entry(self) -> None:
        client = Mock()
        client.get_secret.side_effect = [
            make_secret(
                "api-key",
                "old",
                expires_on=datetime.now(timezone.utc) + timedelta(days=1),
            ),
            make_secret(
                "api-key",
                "new",
                expires_on=datetime.now(timezone.utc) + timedelta(days=30),
            ),
        ]
        cache = SyncSecretCache(SyncSecretProvider(client))

        cache.bulk_load({"api-key": None})

        self.assertEqual(cache.get("api-key"), "new")
        self.assertEqual(client.get_secret.call_count, 2)

    def test_rotation_waits_for_delete_and_purge_before_set(self) -> None:
        events: list[str] = []
        poller = Mock()
        poller.result.side_effect = lambda timeout: events.append("delete-complete")
        client = Mock()
        client.begin_delete_secret.side_effect = lambda name: events.append("delete") or poller
        client.purge_deleted_secret.side_effect = lambda name: events.append("purge")
        client.get_deleted_secret.side_effect = ResourceNotFoundError("gone")
        client.set_secret.side_effect = (
            lambda *args, **kwargs: events.append("set") or make_secret(args[0], args[1])
        )

        rotate_secret(
            client,
            "rotating",
            "new-value",
            datetime.now(timezone.utc) + timedelta(days=30),
        )

        self.assertEqual(events, ["delete", "delete-complete", "purge", "set"])


class AsyncSecretTests(unittest.IsolatedAsyncioTestCase):
    async def test_specific_version_and_expiry_are_returned(self) -> None:
        expires_on = datetime.now(timezone.utc) + timedelta(days=2)
        client = SimpleNamespace(
            get_secret=AsyncMock(
                return_value=make_secret(
                    "api-key",
                    "value",
                    version="v2",
                    expires_on=expires_on,
                )
            )
        )
        provider = AsyncSecretProvider(client)

        record = await provider.get_secret_record("api-key", version="v2")

        self.assertEqual(record.version, "v2")
        self.assertEqual(record.expires_on, expires_on)
        self.assertTrue(record.expires_within(timedelta(days=7)))
        client.get_secret.assert_awaited_once_with("api-key", version="v2")

    async def test_missing_secret_returns_default_and_version_is_forwarded(self) -> None:
        client = SimpleNamespace(
            get_secret=AsyncMock(side_effect=ResourceNotFoundError("missing"))
        )
        provider = AsyncSecretProvider(client)

        value = await provider.get_secret("missing", "fallback", version="v2")

        self.assertEqual(value, "fallback")
        client.get_secret.assert_awaited_once_with("missing", version="v2")

    async def test_cache_refreshes_near_expiry_entry(self) -> None:
        client = SimpleNamespace(
            get_secret=AsyncMock(
                side_effect=[
                    make_secret(
                        "api-key",
                        "old",
                        expires_on=datetime.now(timezone.utc) + timedelta(days=1),
                    ),
                    make_secret(
                        "api-key",
                        "new",
                        expires_on=datetime.now(timezone.utc) + timedelta(days=30),
                    ),
                ]
            )
        )
        cache = AsyncSecretCache(AsyncSecretProvider(client))

        await cache.bulk_load({"api-key": None})

        self.assertEqual(await cache.get("api-key"), "new")
        self.assertEqual(client.get_secret.await_count, 2)

    async def test_rotation_waits_for_delete_and_purge_before_set(self) -> None:
        events: list[str] = []

        class Poller:
            async def result(self) -> None:
                events.append("delete-complete")

        async def begin_delete(name: str) -> Poller:
            events.append("delete")
            return Poller()

        async def purge(name: str) -> None:
            events.append("purge")

        async def get_deleted(name: str) -> None:
            raise ResourceNotFoundError("gone")

        async def set_secret(*args: object, **kwargs: object) -> SimpleNamespace:
            events.append("set")
            return make_secret(str(args[0]), str(args[1]))

        client = SimpleNamespace(
            begin_delete_secret=begin_delete,
            purge_deleted_secret=purge,
            get_deleted_secret=get_deleted,
            set_secret=set_secret,
        )

        await rotate_secret_async(
            client,
            "rotating",
            "new-value",
            datetime.now(timezone.utc) + timedelta(days=30),
        )

        self.assertEqual(events, ["delete", "delete-complete", "purge", "set"])


if __name__ == "__main__":
    unittest.main()
