import asyncio
import unittest
from datetime import datetime, timedelta, timezone
from types import SimpleNamespace
from unittest.mock import AsyncMock, Mock

from azure.core.exceptions import ResourceNotFoundError

from keyvault_config.cache import AsyncSecretCache, SecretCache
from keyvault_config.provider import (
    AsyncKeyVaultSecretProvider,
    KeyVaultSecretProvider,
)
from keyvault_config.rotation import AsyncSecretRotator, SecretRotator


def _secret(name, value, version="1", expires_on=None):
    return SimpleNamespace(
        name=name,
        value=value,
        properties=SimpleNamespace(
            version=version,
            expires_on=expires_on,
        ),
    )


class ProviderAndCacheTests(unittest.TestCase):
    def test_missing_secret_returns_default_and_version_is_forwarded(self):
        client = Mock()
        client.get_secret.side_effect = ResourceNotFoundError()
        provider = KeyVaultSecretProvider(client)

        info = provider.get_secret_info("missing", "fallback", "v2")

        self.assertEqual(info.value, "fallback")
        self.assertFalse(info.found)
        client.get_secret.assert_called_once_with("missing", "v2")

    def test_near_expiry_secret_is_refetched(self):
        client = Mock()
        expiry = datetime.now(timezone.utc) + timedelta(days=1)
        client.get_secret.side_effect = [
            _secret("api-key", "old", expires_on=expiry),
            _secret("api-key", "new", "2", expiry + timedelta(days=30)),
        ]
        cache = SecretCache(KeyVaultSecretProvider(client))
        cache.load_required(["api-key"])

        self.assertEqual(cache.get("api-key"), "new")
        self.assertEqual(client.get_secret.call_count, 2)


class AsyncProviderAndCacheTests(unittest.IsolatedAsyncioTestCase):
    async def test_bulk_load_and_expiry_refresh(self):
        client = AsyncMock()
        near = datetime.now(timezone.utc) + timedelta(days=1)
        far = near + timedelta(days=30)
        client.get_secret.side_effect = [
            _secret("one", "1", expires_on=near),
            _secret("two", "2", expires_on=far),
            _secret("one", "updated", "2", far),
        ]
        cache = AsyncSecretCache(AsyncKeyVaultSecretProvider(client))

        await cache.load_required(["one", "two"])
        refreshed = await cache.refresh_expiring()

        self.assertEqual(refreshed, {"one": "updated"})


class RotationTests(unittest.TestCase):
    def test_sync_rotation_waits_then_purges_before_create(self):
        events = []
        poller = Mock()
        poller.wait.side_effect = lambda: events.append("wait")
        client = Mock()
        client.begin_delete_secret.side_effect = lambda name: (
            events.append("delete") or poller
        )
        client.purge_deleted_secret.side_effect = lambda name: events.append(
            "purge"
        )
        client.get_deleted_secret.side_effect = ResourceNotFoundError()
        client.set_secret.side_effect = (
            lambda *args, **kwargs: events.append("create") or "created"
        )

        result = SecretRotator(client).rotate(
            "secret", "new", datetime.now(timezone.utc)
        )

        self.assertEqual(result, "created")
        self.assertEqual(events, ["delete", "wait", "purge", "create"])

    def test_async_rotation_waits_then_purges_before_create(self):
        async def run():
            events = []
            poller = Mock()
            poller.wait = AsyncMock(
                side_effect=lambda: events.append("wait")
            )
            client = Mock()
            client.begin_delete_secret = AsyncMock(
                side_effect=lambda name: events.append("delete") or poller
            )
            client.purge_deleted_secret = AsyncMock(
                side_effect=lambda name: events.append("purge")
            )
            client.get_deleted_secret = AsyncMock(
                side_effect=ResourceNotFoundError()
            )
            client.set_secret = AsyncMock(
                side_effect=lambda *args, **kwargs: (
                    events.append("create") or "created"
                )
            )

            result = await AsyncSecretRotator(client).rotate(
                "secret", "new", datetime.now(timezone.utc)
            )
            self.assertEqual(result, "created")
            self.assertEqual(events, ["delete", "wait", "purge", "create"])

        asyncio.run(run())
