from __future__ import annotations

import os
import unittest

from azure.core.exceptions import ResourceNotFoundError
from cryptography.hazmat.primitives.keywrap import aes_key_unwrap, aes_key_wrap

from encrypted_blob.blob_transfer import (
    AsyncEncryptedBlobClient,
    BlobEncryptionError,
    SyncEncryptedBlobClient,
)
from encrypted_blob.key_management import DataKeyMaterial


class _Downloader:
    def __init__(self, record: dict[str, object]) -> None:
        self._record = record
        self.properties = type(
            "Properties", (), {"metadata": record["metadata"]}
        )()

    def readall(self) -> bytes:
        return self._record["data"]  # type: ignore[return-value]


class _AsyncDownloader(_Downloader):
    async def readall(self) -> bytes:
        return super().readall()


class _BlobClient:
    def __init__(self, blobs: dict[str, dict[str, object]], name: str) -> None:
        self._blobs = blobs
        self._name = name

    def __enter__(self) -> "_BlobClient":
        return self

    def __exit__(self, *args: object) -> None:
        return None

    def upload_blob(
        self,
        data: bytes,
        *,
        overwrite: bool,
        metadata: dict[str, str],
    ) -> None:
        if not overwrite and self._name in self._blobs:
            raise AssertionError("Test blob already exists")
        self._blobs[self._name] = {"data": data, "metadata": metadata}

    def download_blob(self) -> _Downloader:
        try:
            return _Downloader(self._blobs[self._name])
        except KeyError as error:
            raise ResourceNotFoundError("missing") from error


class _AsyncBlobClient(_BlobClient):
    async def __aenter__(self) -> "_AsyncBlobClient":
        return self

    async def __aexit__(self, *args: object) -> None:
        return None

    async def upload_blob(
        self,
        data: bytes,
        *,
        overwrite: bool,
        metadata: dict[str, str],
    ) -> None:
        super().upload_blob(data, overwrite=overwrite, metadata=metadata)

    async def download_blob(self) -> _AsyncDownloader:
        try:
            return _AsyncDownloader(self._blobs[self._name])
        except KeyError as error:
            raise ResourceNotFoundError("missing") from error


class _BlobService:
    def __init__(self) -> None:
        self.blobs: dict[str, dict[str, object]] = {}

    def get_blob_client(self, *, container: str, blob: str) -> _BlobClient:
        del container
        return _BlobClient(self.blobs, blob)


class _AsyncBlobService(_BlobService):
    def get_blob_client(self, *, container: str, blob: str) -> _AsyncBlobClient:
        del container
        return _AsyncBlobClient(self.blobs, blob)


class _KeyManager:
    key_id = "https://example.vault.azure.net/keys/wrapping/version"

    def __init__(self) -> None:
        self._kek = os.urandom(32)

    def generate_and_wrap_data_key(self) -> DataKeyMaterial:
        key = bytearray(os.urandom(32))
        return DataKeyMaterial(key, aes_key_wrap(self._kek, bytes(key)), self.key_id)

    def unwrap_data_key(self, wrapped_key: bytes, key_id: str) -> bytearray:
        if key_id != self.key_id:
            raise AssertionError("Unexpected key ID")
        return bytearray(aes_key_unwrap(self._kek, wrapped_key))


class _AsyncKeyManager(_KeyManager):
    async def generate_and_wrap_data_key(self) -> DataKeyMaterial:
        return super().generate_and_wrap_data_key()

    async def unwrap_data_key(
        self, wrapped_key: bytes, key_id: str
    ) -> bytearray:
        return super().unwrap_data_key(wrapped_key, key_id)


class SyncEncryptedBlobClientTests(unittest.TestCase):
    def test_round_trip_and_ciphertext_storage(self) -> None:
        service = _BlobService()
        client = SyncEncryptedBlobClient(service, _KeyManager(), "files")  # type: ignore[arg-type]
        plaintext = b"authenticated plaintext"

        result = client.upload_bytes("sync.bin", plaintext)
        decrypted = client.download_bytes("sync.bin")

        self.assertEqual(plaintext, decrypted)
        self.assertNotEqual(plaintext, service.blobs["sync.bin"]["data"])
        self.assertEqual(_KeyManager.key_id, result.key_id)
        self.assertTrue(result.wrapped_key_base64)

    def test_missing_blob_has_clear_error(self) -> None:
        client = SyncEncryptedBlobClient(  # type: ignore[arg-type]
            _BlobService(), _KeyManager(), "files"
        )
        with self.assertRaisesRegex(BlobEncryptionError, "does not exist"):
            client.download_bytes("missing.bin")

    def test_modified_ciphertext_fails_authentication(self) -> None:
        service = _BlobService()
        client = SyncEncryptedBlobClient(service, _KeyManager(), "files")  # type: ignore[arg-type]
        client.upload_bytes("changed.bin", b"original")
        ciphertext = bytearray(service.blobs["changed.bin"]["data"])  # type: ignore[arg-type]
        ciphertext[0] ^= 1
        service.blobs["changed.bin"]["data"] = bytes(ciphertext)

        with self.assertRaisesRegex(BlobEncryptionError, "authentication failed"):
            client.download_bytes("changed.bin")


class AsyncEncryptedBlobClientTests(unittest.IsolatedAsyncioTestCase):
    async def test_round_trip_and_ciphertext_storage(self) -> None:
        service = _AsyncBlobService()
        client = AsyncEncryptedBlobClient(  # type: ignore[arg-type]
            service, _AsyncKeyManager(), "files"
        )
        plaintext = b"async authenticated plaintext"

        result = await client.upload_bytes("async.bin", plaintext)
        decrypted = await client.download_bytes("async.bin")

        self.assertEqual(plaintext, decrypted)
        self.assertNotEqual(plaintext, service.blobs["async.bin"]["data"])
        self.assertEqual(_KeyManager.key_id, result.key_id)


if __name__ == "__main__":
    unittest.main()
