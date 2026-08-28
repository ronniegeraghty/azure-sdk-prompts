from __future__ import annotations

import asyncio
import unittest
from datetime import datetime, timedelta, timezone
from types import SimpleNamespace
from unittest.mock import AsyncMock, Mock, patch

from azure.core.exceptions import ResourceNotFoundError

from azure_keyvault_config.cache import AsyncSecretCache, SecretCache
from azure_keyvault_config.provider import AsyncSecretProvider, SecretProvider
from azure_keyvault_config.rotation import AsyncSecretRotator, SecretRotator


def secret(
    name: str,
    value: str,
    expires_on: datetime | None = None,
    version: str = "v1",
) -> SimpleNamespace:
    return SimpleNamespace(
        name=name,
        value=value,
        properties=SimpleNamespace(version=version, expires_on=expires_on),
    )


def not_found() -> ResourceNotFoundError:
    return ResourceNotFoundError("not found")


class ProviderTests(unittest.TestCase):
    def test_sync_provider_returns_default_and_passes_version(self) -> None:
        client = Mock()
        client.get_secret.side_effect = not_found()
        provider = SecretProvider(client)

        self.assertEqual(provider.get_secret("missing", "fallback", "v2"), "fallback")
        client.get_secret.assert_called_once_with("missing", "v2")

    def test_sync_provider_exposes_expiry(self) -> None:
        expiry = datetime.now(timezone.utc) + timedelta(days=30)
        client = Mock()
        client.get_secret.return_value = secret("item", "value", expiry)

        self.assertEqual(SecretProvider(client).get_expiry("item"), expiry)

    def test_async_provider_returns_default_and_passes_version(self) -> None:
        async def run() -> None:
            client = Mock()
            client.get_secret = AsyncMock(side_effect=not_found())
            provider = AsyncSecretProvider(client)

            value = await provider.get_secret("missing", "fallback", "v2")

            self.assertEqual(value, "fallback")
            client.get_secret.assert_awaited_once_with("missing", "v2")

        asyncio.run(run())


class CacheTests(unittest.TestCase):
    def test_loads_required_keys_and_uses_cache(self) -> None:
        provider = Mock()
        provider.get_secret_details.side_effect = [
            SimpleNamespace(
                name="one",
                value="1",
                version="v1",
                expires_on=None,
                found=True,
            ),
            SimpleNamespace(
                name="two",
                value="2",
                version="v1",
                expires_on=None,
                found=True,
            ),
        ]
        cache = SecretCache(provider, required_keys=("one", "two"))

        self.assertEqual(cache.load_required(), {"one": "1", "two": "2"})
        self.assertEqual(cache.get("one"), "1")
        self.assertEqual(provider.get_secret_details.call_count, 2)

    def test_get_automatically_refreshes_near_expiry(self) -> None:
        soon = datetime.now(timezone.utc) + timedelta(days=1)
        later = datetime.now(timezone.utc) + timedelta(days=30)
        provider = Mock()
        provider.get_secret_details.side_effect = [
            SimpleNamespace(
                name="key",
                value="old",
                version="v1",
                expires_on=soon,
                found=True,
            ),
            SimpleNamespace(
                name="key",
                value="new",
                version="v2",
                expires_on=later,
                found=True,
            ),
        ]
        cache = SecretCache(provider, required_keys=("key",))
        cache.load_required()

        self.assertEqual(cache.get("key"), "new")
        self.assertEqual(provider.get_secret_details.call_count, 2)

    def test_async_cache_loads_and_refreshes(self) -> None:
        async def run() -> None:
            soon = datetime.now(timezone.utc) + timedelta(days=1)
            later = datetime.now(timezone.utc) + timedelta(days=30)
            provider = Mock()
            provider.get_secret_details = AsyncMock(
                side_effect=[
                    SimpleNamespace(
                        name="key",
                        value="old",
                        version="v1",
                        expires_on=soon,
                        found=True,
                    ),
                    SimpleNamespace(
                        name="key",
                        value="new",
                        version="v2",
                        expires_on=later,
                        found=True,
                    ),
                ]
            )
            cache = AsyncSecretCache(provider, required_keys=("key",))

            await cache.load_required()
            self.assertEqual(await cache.get("key"), "new")

        asyncio.run(run())


class RotationTests(unittest.TestCase):
    @patch("azure_keyvault_config.rotation.time.sleep", return_value=None)
    def test_sync_rotation_waits_purges_then_sets(self, _sleep: Mock) -> None:
        client = Mock()
        poller = Mock()
        deleted = SimpleNamespace(recovery_id="https://recovery")
        poller.result.return_value = deleted
        client.begin_delete_secret.return_value = poller
        client.get_deleted_secret.side_effect = [deleted, not_found()]
        client.set_secret.return_value = secret("key", "new")
        expiry = datetime.now(timezone.utc) + timedelta(days=30)

        SecretRotator(client).rotate("key", "new", expiry)

        poller.wait.assert_called_once_with()
        client.purge_deleted_secret.assert_called_once_with("key")
        client.set_secret.assert_called_once_with("key", "new", expires_on=expiry)

    def test_async_rotation_waits_purges_then_sets(self) -> None:
        async def run() -> None:
            client = Mock()
            deleted = SimpleNamespace(recovery_id="https://recovery")
            client.delete_secret = AsyncMock(return_value=deleted)
            client.purge_deleted_secret = AsyncMock()
            client.get_deleted_secret = AsyncMock(side_effect=not_found())
            client.set_secret = AsyncMock(return_value=secret("key", "new"))
            expiry = datetime.now(timezone.utc) + timedelta(days=30)

            await AsyncSecretRotator(client).rotate("key", "new", expiry)

            client.delete_secret.assert_awaited_once_with("key")
            client.purge_deleted_secret.assert_awaited_once_with("key")
            client.set_secret.assert_awaited_once_with(
                "key", "new", expires_on=expiry
            )

        asyncio.run(run())


if __name__ == "__main__":
    unittest.main()
