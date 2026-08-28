from __future__ import annotations

import unittest
from dataclasses import dataclass
from types import SimpleNamespace
from unittest.mock import AsyncMock, Mock, call, patch

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError

from blob_crypto import (
    AsyncEncryptedBlobClient,
    CiphertextIntegrityError,
    EncryptedBlobClient,
    EncryptedBlobNotFoundError,
    EncryptionMetadataError,
)
from key_management import (
    AsyncKeyManager,
    KeyManagementError,
    KeyManager,
    ProtectedDataKey,
)


@dataclass
class _StoredBlob:
    data: bytes
    metadata: dict[str, str]


class _FakeDownload:
    def __init__(self, data: bytes) -> None:
        self._data = data

    def readall(self) -> bytes:
        return self._data


class _FakeBlob:
    def __init__(self, container: "_FakeContainer", name: str) -> None:
        self._container = container
        self._name = name

    def _stored(self) -> _StoredBlob:
        try:
            return self._container.blobs[self._name]
        except KeyError as exc:
            raise ResourceNotFoundError("missing") from exc

    def get_blob_properties(self) -> SimpleNamespace:
        return SimpleNamespace(metadata=self._stored().metadata)

    def download_blob(self) -> _FakeDownload:
        return _FakeDownload(self._stored().data)


class _FakeContainer:
    def __init__(self) -> None:
        self.blobs: dict[str, _StoredBlob] = {}

    def upload_blob(
        self,
        *,
        name: str,
        data: bytes,
        metadata: dict[str, str],
        overwrite: bool,
    ) -> None:
        self.blobs[name] = _StoredBlob(data=data, metadata=metadata)

    def get_blob_client(self, name: str) -> _FakeBlob:
        return _FakeBlob(self, name)


class _FakeBlobService:
    def __init__(self) -> None:
        self.container = _FakeContainer()

    def get_container_client(self, name: str) -> _FakeContainer:
        return self.container


class _FakeKeyManager:
    DATA_KEY = bytes(range(32))

    def generate_data_key(self) -> bytes:
        return self.DATA_KEY

    def protect_data_key(self, data_key: bytes) -> ProtectedDataKey:
        self.assert_key(data_key)
        return ProtectedDataKey("https://vault/keys/k/version", b"wrapped")

    def recover_data_key(self, protected_key: ProtectedDataKey) -> bytes:
        return self.DATA_KEY

    @staticmethod
    def assert_key(data_key: bytes) -> None:
        if data_key != _FakeKeyManager.DATA_KEY:
            raise AssertionError("unexpected key")


class _AsyncDownload:
    def __init__(self, data: bytes) -> None:
        self._data = data

    async def readall(self) -> bytes:
        return self._data


class _AsyncBlob:
    def __init__(self, container: "_AsyncContainer", name: str) -> None:
        self._container = container
        self._name = name

    def _stored(self) -> _StoredBlob:
        try:
            return self._container.blobs[self._name]
        except KeyError as exc:
            raise ResourceNotFoundError("missing") from exc

    async def get_blob_properties(self) -> SimpleNamespace:
        return SimpleNamespace(metadata=self._stored().metadata)

    async def download_blob(self) -> _AsyncDownload:
        return _AsyncDownload(self._stored().data)


class _AsyncContainer:
    def __init__(self) -> None:
        self.blobs: dict[str, _StoredBlob] = {}

    async def upload_blob(
        self,
        *,
        name: str,
        data: bytes,
        metadata: dict[str, str],
        overwrite: bool,
    ) -> None:
        self.blobs[name] = _StoredBlob(data=data, metadata=metadata)

    def get_blob_client(self, name: str) -> _AsyncBlob:
        return _AsyncBlob(self, name)


class _AsyncBlobService:
    def __init__(self) -> None:
        self.container = _AsyncContainer()

    def get_container_client(self, name: str) -> _AsyncContainer:
        return self.container


class _AsyncKeyManager(_FakeKeyManager):
    async def protect_data_key(
        self, data_key: bytes
    ) -> ProtectedDataKey:
        return super().protect_data_key(data_key)

    async def recover_data_key(
        self, protected_key: ProtectedDataKey
    ) -> bytes:
        return super().recover_data_key(protected_key)


class SyncEncryptedBlobTests(unittest.TestCase):
    def setUp(self) -> None:
        self.service = _FakeBlobService()
        self.client = EncryptedBlobClient(
            self.service, "container", _FakeKeyManager()
        )

    def test_round_trip(self) -> None:
        result = self.client.upload_bytes("blob", b"secret")
        self.assertEqual("https://vault/keys/k/version", result.key_id)
        self.assertEqual(b"secret", self.client.download_bytes("blob"))

    def test_tampering_is_detected(self) -> None:
        self.client.upload_bytes("blob", b"secret")
        stored = self.service.container.blobs["blob"]
        stored.data = bytes([stored.data[0] ^ 1]) + stored.data[1:]
        with self.assertRaises(CiphertextIntegrityError):
            self.client.download_bytes("blob")

    def test_missing_metadata_is_rejected(self) -> None:
        self.service.container.blobs["blob"] = _StoredBlob(b"x", {})
        with self.assertRaises(EncryptionMetadataError):
            self.client.download_bytes("blob")

    def test_missing_blob_has_specific_error(self) -> None:
        with self.assertRaises(EncryptedBlobNotFoundError):
            self.client.download_bytes("missing")


class AsyncEncryptedBlobTests(unittest.IsolatedAsyncioTestCase):
    async def test_round_trip(self) -> None:
        service = _AsyncBlobService()
        client = AsyncEncryptedBlobClient(
            service, "container", _AsyncKeyManager()
        )
        result = await client.upload_bytes("blob", b"async secret")
        self.assertEqual("https://vault/keys/k/version", result.key_id)
        self.assertEqual(
            b"async secret", await client.download_bytes("blob")
        )


class KeyManagerTests(unittest.TestCase):
    @patch("key_management.CryptographyClient")
    def test_wrap_and_unwrap_use_versioned_key_id(
        self, crypto_client_type: Mock
    ) -> None:
        key_client = Mock()
        key_client.get_key.return_value = SimpleNamespace(
            id="https://vault/keys/k/version"
        )
        crypto_client = crypto_client_type.return_value
        crypto_client.wrap_key.return_value = SimpleNamespace(
            encrypted_key=b"wrapped"
        )
        crypto_client.unwrap_key.return_value = SimpleNamespace(
            key=bytes(range(32))
        )
        manager = KeyManager(key_client, Mock(), "k")

        protected = manager.protect_data_key(bytes(range(32)))
        recovered = manager.recover_data_key(protected)

        self.assertEqual(bytes(range(32)), recovered)
        self.assertEqual(
            [
                call("https://vault/keys/k/version", manager._credential),
                call("https://vault/keys/k/version", manager._credential),
            ],
            crypto_client_type.call_args_list,
        )
        self.assertEqual(2, crypto_client.close.call_count)

    @patch("key_management.CryptographyClient")
    def test_disabled_key_error_is_contextual(
        self, crypto_client_type: Mock
    ) -> None:
        key_client = Mock()
        key_client.get_key.return_value = SimpleNamespace(
            id="https://vault/keys/k/version"
        )
        crypto_client_type.return_value.wrap_key.side_effect = (
            HttpResponseError(message="Key is disabled")
        )
        manager = KeyManager(key_client, Mock(), "k")

        with self.assertRaisesRegex(KeyManagementError, "could not wrap"):
            manager.protect_data_key(bytes(range(32)))


class AsyncKeyManagerTests(unittest.IsolatedAsyncioTestCase):
    @patch("key_management.AsyncCryptographyClient")
    async def test_async_wrap_and_unwrap(
        self, crypto_client_type: Mock
    ) -> None:
        key_client = Mock()
        key_client.get_key = AsyncMock(
            return_value=SimpleNamespace(
                id="https://vault/keys/k/version"
            )
        )
        crypto_client = crypto_client_type.return_value
        crypto_client.wrap_key = AsyncMock(
            return_value=SimpleNamespace(encrypted_key=b"wrapped")
        )
        crypto_client.unwrap_key = AsyncMock(
            return_value=SimpleNamespace(key=bytes(range(32)))
        )
        crypto_client.close = AsyncMock()
        manager = AsyncKeyManager(key_client, Mock(), "k")

        protected = await manager.protect_data_key(bytes(range(32)))
        recovered = await manager.recover_data_key(protected)

        self.assertEqual(bytes(range(32)), recovered)
        self.assertEqual(2, crypto_client.close.await_count)


if __name__ == "__main__":
    unittest.main()
