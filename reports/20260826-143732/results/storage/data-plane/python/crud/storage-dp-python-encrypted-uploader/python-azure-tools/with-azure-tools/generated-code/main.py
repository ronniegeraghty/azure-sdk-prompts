from __future__ import annotations

import asyncio
import base64
import os
import sys

from configuration import (
    ConfigurationError,
    Settings,
    create_async_clients,
    create_sync_clients,
)
from encrypted_blob import (
    AsyncEncryptedBlobClient,
    EncryptedBlobClient,
    EncryptedBlobError,
    UploadResult,
)
from key_management import AsyncKeyManager, KeyManagementError, KeyManager


def print_result(label: str, result: UploadResult, decrypted: bytes) -> None:
    print(f"{label} vault key ID: {result.key_id}")
    print(
        f"{label} wrapped DEK (base64): "
        f"{base64.b64encode(result.wrapped_key).decode('ascii')}"
    )
    print(f"{label} decrypted output: {decrypted.decode('utf-8')}")


def run_sync(settings: Settings, plaintext: bytes) -> None:
    blob_name = f"sync-{settings.blob_name}"
    with create_sync_clients(settings) as clients:
        key_manager = KeyManager(
            clients.key_client, clients.credential, settings.key_name
        )
        encrypted_blobs = EncryptedBlobClient(
            clients.blob_service, key_manager, settings.storage_container
        )
        result = encrypted_blobs.upload_bytes(blob_name, plaintext)
        decrypted = encrypted_blobs.download_bytes(blob_name)
    print_result("sync", result, decrypted)


async def run_async(settings: Settings, plaintext: bytes) -> None:
    blob_name = f"async-{settings.blob_name}"
    async with create_async_clients(settings) as clients:
        key_manager = AsyncKeyManager(
            clients.key_client, clients.credential, settings.key_name
        )
        encrypted_blobs = AsyncEncryptedBlobClient(
            clients.blob_service, key_manager, settings.storage_container
        )
        result = await encrypted_blobs.upload_bytes(blob_name, plaintext)
        decrypted = await encrypted_blobs.download_bytes(blob_name)
    print_result("async", result, decrypted)


def main() -> int:
    try:
        settings = Settings.from_environment()
        plaintext = os.getenv(
            "DEMO_PLAINTEXT", "Client-side encrypted with Azure Key Vault"
        ).encode("utf-8")
        run_sync(settings, plaintext)
        asyncio.run(run_async(settings, plaintext))
        return 0
    except (ConfigurationError, KeyManagementError, EncryptedBlobError) as error:
        print(f"Error: {error}", file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
