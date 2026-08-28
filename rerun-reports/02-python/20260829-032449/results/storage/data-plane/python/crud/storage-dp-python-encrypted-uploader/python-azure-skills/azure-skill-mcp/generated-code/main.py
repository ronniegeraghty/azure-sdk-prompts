from __future__ import annotations

import asyncio
import os
from pathlib import Path

from blob_crypto import AsyncEncryptedBlobClient, EncryptedBlobClient
from config import (
    Settings,
    build_async_connections,
    build_sync_connections,
)
from key_management import AsyncKeyManager, KeyManager


def run_sync(settings: Settings, source: Path, blob_name: str) -> None:
    with build_sync_connections(settings) as connections:
        key_manager = KeyManager(
            connections.key_client,
            connections.credential,
            settings.key_name,
        )
        client = EncryptedBlobClient(
            connections.blob_service_client,
            settings.storage_container_name,
            key_manager,
        )
        upload = client.upload_file(blob_name, source, overwrite=True)
        decrypted = client.download_bytes(blob_name)

    print("Sync implementation")
    print(f"Vault key ID: {upload.key_id}")
    print(f"Wrapped DEK (base64): {upload.wrapped_key_base64}")
    print(f"Decrypted output: {decrypted.decode('utf-8')}")


async def run_async(
    settings: Settings, source: Path, blob_name: str
) -> None:
    async with build_async_connections(settings) as connections:
        key_manager = AsyncKeyManager(
            connections.key_client,
            connections.credential,
            settings.key_name,
        )
        client = AsyncEncryptedBlobClient(
            connections.blob_service_client,
            settings.storage_container_name,
            key_manager,
        )
        upload = await client.upload_file(blob_name, source, overwrite=True)
        decrypted = await client.download_bytes(blob_name)

    print("Async implementation")
    print(f"Vault key ID: {upload.key_id}")
    print(f"Wrapped DEK (base64): {upload.wrapped_key_base64}")
    print(f"Decrypted output: {decrypted.decode('utf-8')}")


def main() -> None:
    settings = Settings.from_env()
    source = Path(os.getenv("DEMO_FILE_PATH", "demo-input.txt"))
    if not source.is_file():
        raise FileNotFoundError(
            f"Demo input file {source} does not exist; set DEMO_FILE_PATH"
        )

    run_sync(
        settings,
        source,
        os.getenv("DEMO_SYNC_BLOB_NAME", "encrypted-sync.bin"),
    )
    asyncio.run(
        run_async(
            settings,
            source,
            os.getenv("DEMO_ASYNC_BLOB_NAME", "encrypted-async.bin"),
        )
    )


if __name__ == "__main__":
    main()
