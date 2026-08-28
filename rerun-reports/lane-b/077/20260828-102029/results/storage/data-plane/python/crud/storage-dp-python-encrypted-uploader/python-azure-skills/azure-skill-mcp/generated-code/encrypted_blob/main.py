"""Run synchronous and asynchronous encrypted blob round trips."""

from __future__ import annotations

import argparse
import asyncio
from pathlib import Path

from .blob_transfer import AsyncEncryptedBlobClient, EncryptedBlobClient
from .configuration import AsyncAzureClients, AzureSettings, SyncAzureClients
from .key_management import AsyncKeyManager, KeyManager

DEFAULT_PAYLOAD = b"Client-side encryption with Azure Key Vault and Blob Storage."


def _load_payload(source: Path | None) -> bytes:
    return source.read_bytes() if source else DEFAULT_PAYLOAD


def _display_decrypted(payload: bytes) -> str:
    return payload.decode("utf-8", errors="replace")


def run_sync(settings: AzureSettings, payload: bytes) -> None:
    with SyncAzureClients(settings) as clients:
        with KeyManager.from_key_client(
            clients.key_client,
            clients.credential,
            settings.key_name,
            settings.key_version,
        ) as key_manager:
            encrypted_blobs = EncryptedBlobClient(
                clients.container_client, key_manager
            )
            result = encrypted_blobs.upload_bytes(settings.blob_name, payload)
            decrypted = encrypted_blobs.download_bytes(settings.blob_name)

    print("Sync implementation")
    print(f"Vault key ID: {result.key_id}")
    print(f"Wrapped DEK (base64): {result.wrapped_data_key_base64}")
    print(f"Decrypted output: {_display_decrypted(decrypted)}")


async def run_async(settings: AzureSettings, payload: bytes) -> None:
    async with AsyncAzureClients(settings) as clients:
        async with await AsyncKeyManager.from_key_client(
            clients.key_client,
            clients.credential,
            settings.key_name,
            settings.key_version,
        ) as key_manager:
            encrypted_blobs = AsyncEncryptedBlobClient(
                clients.container_client, key_manager
            )
            result = await encrypted_blobs.upload_bytes(
                settings.blob_name, payload
            )
            decrypted = await encrypted_blobs.download_bytes(settings.blob_name)

    print("\nAsync implementation")
    print(f"Vault key ID: {result.key_id}")
    print(f"Wrapped DEK (base64): {result.wrapped_data_key_base64}")
    print(f"Decrypted output: {_display_decrypted(decrypted)}")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Upload and download an AES-GCM encrypted Azure blob."
    )
    parser.add_argument(
        "source",
        nargs="?",
        type=Path,
        help="Optional file to upload; otherwise a built-in UTF-8 message is used.",
    )
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    settings = AzureSettings.from_environment()
    payload = _load_payload(args.source)
    run_sync(settings, payload)
    asyncio.run(run_async(settings, payload))


if __name__ == "__main__":
    main()
