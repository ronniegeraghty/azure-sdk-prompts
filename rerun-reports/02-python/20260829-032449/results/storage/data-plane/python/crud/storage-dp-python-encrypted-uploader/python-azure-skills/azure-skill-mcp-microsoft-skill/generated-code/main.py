"""Demonstrate sync and async encrypted Blob Storage round trips."""

from __future__ import annotations

import asyncio
import sys

from azure.core.exceptions import AzureError

from encrypted_blob.blob_transfer import (
    AsyncEncryptedBlobClient,
    BlobEncryptionError,
    SyncEncryptedBlobClient,
)
from encrypted_blob.config import (
    Settings,
    create_async_clients,
    create_sync_clients,
)
from encrypted_blob.key_management import (
    AsyncKeyManager,
    KeyManagementError,
    SyncKeyManager,
)


def run_sync(settings: Settings) -> None:
    with create_sync_clients(settings) as clients:
        key_manager = SyncKeyManager(
            clients.key_client, clients.credential, settings.key_name
        )
        encrypted_blobs = SyncEncryptedBlobClient(
            clients.blob_service,
            key_manager,
            settings.storage_container_name,
        )
        upload = encrypted_blobs.upload_file(
            settings.input_file, settings.sync_blob_name
        )
        plaintext = encrypted_blobs.download_file(
            settings.sync_blob_name, settings.sync_output_file
        )

    print("Sync implementation")
    print(f"Vault key ID: {upload.key_id}")
    print(f"Wrapped DEK (base64): {upload.wrapped_key_base64}")
    print(f"Decrypted output: {plaintext.decode('utf-8')}")


async def run_async(settings: Settings) -> None:
    async with create_async_clients(settings) as clients:
        key_manager = AsyncKeyManager(
            clients.key_client, clients.credential, settings.key_name
        )
        encrypted_blobs = AsyncEncryptedBlobClient(
            clients.blob_service,
            key_manager,
            settings.storage_container_name,
        )
        upload = await encrypted_blobs.upload_file(
            settings.input_file, settings.async_blob_name
        )
        plaintext = await encrypted_blobs.download_file(
            settings.async_blob_name, settings.async_output_file
        )

    print("Async implementation")
    print(f"Vault key ID: {upload.key_id}")
    print(f"Wrapped DEK (base64): {upload.wrapped_key_base64}")
    print(f"Decrypted output: {plaintext.decode('utf-8')}")


def main() -> int:
    try:
        settings = Settings.from_environment()
        run_sync(settings)
        asyncio.run(run_async(settings))
    except (
        AzureError,
        BlobEncryptionError,
        KeyManagementError,
        OSError,
        UnicodeError,
        ValueError,
    ) as error:
        print(f"Error: {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
