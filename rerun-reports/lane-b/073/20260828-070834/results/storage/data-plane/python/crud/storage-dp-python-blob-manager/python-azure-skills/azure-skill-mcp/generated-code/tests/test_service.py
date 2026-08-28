from __future__ import annotations

import tempfile
import unittest
from pathlib import Path
from types import SimpleNamespace

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError

from blob_manager.service import AsyncBlobStorageManager, BlobStorageManager


class _Downloader:
    def __init__(self, data: bytes) -> None:
        self._data = data

    def readinto(self, stream: object) -> int:
        return stream.write(self._data)


class _AsyncDownloader:
    def __init__(self, data: bytes) -> None:
        self._data = data

    async def chunks(self):
        yield self._data


class _Blob:
    def __init__(self, container: "_Container", name: str) -> None:
        self._container = container
        self._name = name

    def get_blob_properties(self, **kwargs):
        if self._name not in self._container.data:
            raise ResourceNotFoundError("missing")
        return SimpleNamespace(etag=self._container.etags[self._name])

    def upload_blob(self, stream, **kwargs):
        self._container.last_upload = kwargs
        self._container.data[self._name] = stream.read()
        self._container.etags[self._name] = "etag-2"


class _AsyncBlob:
    def __init__(self, container: "_AsyncContainer", name: str) -> None:
        self._container = container
        self._name = name

    async def get_blob_properties(self, **kwargs):
        if self._name not in self._container.data:
            raise ResourceNotFoundError("missing")
        return SimpleNamespace(etag=self._container.etags[self._name])

    async def upload_blob(self, chunks, **kwargs):
        self._container.last_upload = kwargs
        data = bytearray()
        async for chunk in chunks:
            data.extend(chunk)
        self._container.data[self._name] = bytes(data)
        self._container.etags[self._name] = "etag-2"


class _Container:
    def __init__(self) -> None:
        self.data: dict[str, bytes] = {}
        self.etags: dict[str, str] = {}
        self.last_upload: dict[str, object] = {}

    def get_blob_client(self, name: str):
        return _Blob(self, name)

    def list_blobs(self, **kwargs):
        return [SimpleNamespace(name=name) for name in self.data]

    def download_blob(self, name: str, **kwargs):
        if name not in self.data:
            raise ResourceNotFoundError("missing")
        return _Downloader(self.data[name])

    def delete_blob(self, name: str, **kwargs):
        if name not in self.data:
            raise ResourceNotFoundError("missing")
        del self.data[name]


class _AsyncContainer:
    def __init__(self) -> None:
        self.data: dict[str, bytes] = {}
        self.etags: dict[str, str] = {}
        self.last_upload: dict[str, object] = {}

    def get_blob_client(self, name: str):
        return _AsyncBlob(self, name)

    def list_blobs(self, **kwargs):
        async def items():
            for name in self.data:
                yield SimpleNamespace(name=name)

        return items()

    async def download_blob(self, name: str, **kwargs):
        if name not in self.data:
            raise ResourceNotFoundError("missing")
        return _AsyncDownloader(self.data[name])

    async def delete_blob(self, name: str, **kwargs):
        if name not in self.data:
            raise ResourceNotFoundError("missing")
        del self.data[name]


class _Client:
    def __init__(self, container: object) -> None:
        self.container = container

    def get_container_client(self, name: str):
        return self.container


class BlobStorageManagerTests(unittest.TestCase):
    def test_sync_round_trip_and_conditional_update(self) -> None:
        container = _Container()
        manager = BlobStorageManager(_Client(container), "test")

        with tempfile.TemporaryDirectory() as directory:
            source = Path(directory, "source.bin")
            destination = Path(directory, "download.bin")
            source.write_bytes(b"first")

            self.assertTrue(manager.upload(source, "sample", timeout=10).succeeded)
            self.assertFalse(container.last_upload["overwrite"])

            source.write_bytes(b"second")
            self.assertTrue(manager.upload(source, "sample", timeout=10).succeeded)
            self.assertEqual(
                container.last_upload["match_condition"],
                MatchConditions.IfNotModified,
            )
            self.assertEqual(manager.list_blobs(timeout=10).value, ["sample"])
            self.assertTrue(
                manager.download("sample", destination, timeout=10).succeeded
            )
            self.assertEqual(destination.read_bytes(), b"second")
            self.assertTrue(manager.delete("sample", timeout=10).succeeded)

    def test_missing_blob_is_a_result_not_an_exception(self) -> None:
        manager = BlobStorageManager(_Client(_Container()), "test")
        result = manager.delete("missing", timeout=10)
        self.assertFalse(result.succeeded)
        self.assertIn("not found", result.message)


class AsyncBlobStorageManagerTests(unittest.IsolatedAsyncioTestCase):
    async def test_async_round_trip_and_conditional_update(self) -> None:
        container = _AsyncContainer()
        manager = AsyncBlobStorageManager(
            _Client(container), "test", upload_chunk_size=2
        )

        with tempfile.TemporaryDirectory() as directory:
            source = Path(directory, "source.bin")
            destination = Path(directory, "download.bin")
            source.write_bytes(b"first")

            self.assertTrue((await manager.upload(source, "sample")).succeeded)
            self.assertFalse(container.last_upload["overwrite"])

            source.write_bytes(b"second")
            self.assertTrue((await manager.upload(source, "sample")).succeeded)
            self.assertEqual(
                container.last_upload["match_condition"],
                MatchConditions.IfNotModified,
            )
            self.assertEqual((await manager.list_blobs()).value, ["sample"])
            self.assertTrue(
                (await manager.download("sample", destination)).succeeded
            )
            self.assertEqual(destination.read_bytes(), b"second")
            self.assertTrue((await manager.delete("sample")).succeeded)


if __name__ == "__main__":
    unittest.main()
