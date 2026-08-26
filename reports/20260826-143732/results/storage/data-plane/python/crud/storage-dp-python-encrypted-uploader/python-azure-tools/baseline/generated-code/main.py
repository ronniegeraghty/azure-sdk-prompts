"""Run synchronous and asynchronous encrypted blob round trips."""

from __future__ import annotations

import asyncio
import base64

from blob_crypto import AsyncEncryptedBlobClient, EncryptedBlobClient, UploadResult
from config import AsyncAzureConnections, AzureSettings, SyncAzureConnections
from key_management import AsyncKeyVaultKeyManager, KeyVaultKeyManager

SYNC_BLOB_NAME = "encrypted-demo-sync.bin"
ASYNC_BLOB_NAME = "encrypted-demo-async.bin"
SYNC_MESSAGE = b"Hello from the synchronous encrypted uploader."
ASYNC_MESSAGE = b"Hello from the asynchronous encrypted uploader."


def print_result(label: str, result: UploadResult, plaintext: bytes) -> None:
    print(f"{label} vault key ID: {result.protected_key.key_id}")
    print(
        f"{label} wrapped DEK (base64): "
        f"{base64.b64encode(result.protected_key.wrapped_key).decode('ascii')}"
    )
    print(f"{label} decrypted output: {plaintext.decode('utf-8')}")


def run_sync(settings: AzureSettings) -> None:
    with SyncAzureConnections(settings) as connections:
        key_manager = KeyVaultKeyManager(
            connections.key_client,
            connections.credential,
            settings.key_name,
        )
        encrypted_blobs = EncryptedBlobClient(
            connections.container, key_manager
        )
        result = encrypted_blobs.upload(SYNC_BLOB_NAME, SYNC_MESSAGE)
        decrypted = encrypted_blobs.download(SYNC_BLOB_NAME)
        print_result("Sync", result, decrypted)


async def run_async(settings: AzureSettings) -> None:
    async with AsyncAzureConnections(settings) as connections:
        key_manager = AsyncKeyVaultKeyManager(
            connections.key_client,
            connections.credential,
            settings.key_name,
        )
        encrypted_blobs = AsyncEncryptedBlobClient(
            connections.container, key_manager
        )
        result = await encrypted_blobs.upload(ASYNC_BLOB_NAME, ASYNC_MESSAGE)
        decrypted = await encrypted_blobs.download(ASYNC_BLOB_NAME)
        print_result("Async", result, decrypted)


def main() -> None:
    settings = AzureSettings.from_environment()
    run_sync(settings)
    asyncio.run(run_async(settings))


if __name__ == "__main__":
    main()
