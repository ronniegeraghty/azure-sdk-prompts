from __future__ import annotations

import asyncio
import base64

from blob_crypto import AsyncEncryptedBlobClient, EncryptedBlobClient
from config import AsyncAzureConnections, Settings, SyncAzureConnections
from key_management import AsyncKeyManager, KeyManager

SYNC_BLOB_NAME = "encrypted-demo-sync.bin"
ASYNC_BLOB_NAME = "encrypted-demo-async.bin"


def print_result(label: str, key_id: str, wrapped_key: bytes, plaintext: bytes) -> None:
    print(f"{label} vault key ID: {key_id}")
    print(
        f"{label} wrapped DEK (base64): "
        f"{base64.b64encode(wrapped_key).decode('ascii')}"
    )
    print(f"{label} decrypted output: {plaintext.decode('utf-8')}")


def run_sync(settings: Settings) -> None:
    with SyncAzureConnections(settings) as connections:
        key_manager = KeyManager(
            key_client=connections.key_client,
            credential=connections.credential,
            key_name=settings.key_name,
            key_version=settings.key_version,
        )
        encrypted_blobs = EncryptedBlobClient(
            connections.container_client,
            key_manager,
        )
        result = encrypted_blobs.upload_bytes(
            SYNC_BLOB_NAME,
            b"Hello from the synchronous encrypted uploader.",
            overwrite=True,
        )
        plaintext = encrypted_blobs.download_bytes(SYNC_BLOB_NAME)
        print_result("Sync", result.key_id, result.wrapped_data_key, plaintext)


async def run_async(settings: Settings) -> None:
    async with AsyncAzureConnections(settings) as connections:
        key_manager = AsyncKeyManager(
            key_client=connections.key_client,
            credential=connections.credential,
            key_name=settings.key_name,
            key_version=settings.key_version,
        )
        encrypted_blobs = AsyncEncryptedBlobClient(
            connections.container_client,
            key_manager,
        )
        result = await encrypted_blobs.upload_bytes(
            ASYNC_BLOB_NAME,
            b"Hello from the asynchronous encrypted uploader.",
            overwrite=True,
        )
        plaintext = await encrypted_blobs.download_bytes(ASYNC_BLOB_NAME)
        print_result("Async", result.key_id, result.wrapped_data_key, plaintext)


def main() -> None:
    settings = Settings.from_environment()
    run_sync(settings)
    asyncio.run(run_async(settings))


if __name__ == "__main__":
    main()
