from __future__ import annotations

import asyncio
import unittest
from datetime import datetime, timedelta, timezone

from keyvault_config.async_cache import AsyncSecretCache
from keyvault_config.async_provider import AsyncSecretProvider
from keyvault_config.cache import SecretCache
from keyvault_config.demo_backend import (
    AsyncInMemorySecretClient,
    InMemorySecretClient,
)
from keyvault_config.provider import SecretProvider
from keyvault_config.rotation import AsyncSecretRotator, SecretRotator


class SyncTests(unittest.TestCase):
    def setUp(self) -> None:
        self.now = datetime(2026, 1, 1, tzinfo=timezone.utc)
        self.client = InMemorySecretClient()

    def test_default_version_and_expiry(self) -> None:
        first = self.client.set_secret(
            "token", "one", expires_on=self.now + timedelta(days=2)
        )
        self.client.set_secret("token", "two")
        provider = SecretProvider(self.client)

        self.assertEqual("fallback", provider.get("missing", "fallback"))
        self.assertEqual("one", provider.get("token", version=first.version))
        self.assertEqual(
            self.now + timedelta(days=2),
            provider.get_expiry("token", version=first.version),
        )

    def test_cache_refreshes_near_expiry(self) -> None:
        self.client.set_secret(
            "token", "one", expires_on=self.now + timedelta(days=2)
        )
        provider = SecretProvider(self.client)
        cache = SecretCache(provider, clock=lambda: self.now)
        cache.load_required(["token"])
        self.client.set_secret(
            "token", "two", expires_on=self.now + timedelta(days=30)
        )

        self.assertEqual("two", cache.get("token"))
        self.assertEqual([], cache.expiring_keys())

    def test_rotation_recreates_secret(self) -> None:
        self.client.set_secret("token", "old")
        expiry = self.now + timedelta(days=90)

        SecretRotator(self.client, poll_interval=0).rotate(
            "token", "new", expires_on=expiry
        )

        secret = self.client.get_secret("token")
        self.assertEqual("new", secret.value)
        self.assertEqual(expiry, secret.properties.expires_on)


class AsyncTests(unittest.IsolatedAsyncioTestCase):
    async def asyncSetUp(self) -> None:
        self.now = datetime(2026, 1, 1, tzinfo=timezone.utc)
        self.inner = InMemorySecretClient()
        self.client = AsyncInMemorySecretClient(self.inner)

    async def test_provider_cache_and_rotation(self) -> None:
        first = await self.client.set_secret(
            "token", "one", expires_on=self.now + timedelta(days=1)
        )
        provider = AsyncSecretProvider(self.client)
        self.assertEqual("fallback", await provider.get("missing", "fallback"))
        self.assertEqual(
            "one", await provider.get("token", version=first.version)
        )

        cache = AsyncSecretCache(provider, clock=lambda: self.now)
        await cache.load_required(["token"])
        await self.client.set_secret(
            "token", "two", expires_on=self.now + timedelta(days=30)
        )
        self.assertEqual("two", await cache.get("token"))

        expiry = self.now + timedelta(days=90)
        await AsyncSecretRotator(
            self.client, poll_interval=0
        ).rotate("token", "three", expires_on=expiry)
        self.assertEqual("three", await provider.get("token"))


if __name__ == "__main__":
    unittest.main()
