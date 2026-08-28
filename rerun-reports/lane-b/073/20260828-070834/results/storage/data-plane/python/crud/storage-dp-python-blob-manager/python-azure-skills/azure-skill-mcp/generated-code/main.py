"""Demonstrate synchronous and asynchronous blob management operations."""

from __future__ import annotations

import asyncio
import os
from pathlib import Path

from blob_manager import (
    AsyncBlobStorageManager,
    BlobStorageManager,
    BlobStorageSettings,
    create_async_blob_service_client,
    create_blob_service_client,
)

CONTAINER_NAME = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
SAMPLE_FILE = Path("sample-upload.txt")
SYNC_DOWNLOAD = Path("downloads/sync-sample.txt")
ASYNC_DOWNLOAD = Path("downloads/async-sample.txt")
REQUEST_TIMEOUT = 60
TAGS = {"project": "blob-manager", "purpose": "demo"}
METADATA = {"created-by": "blob-manager-demo"}


def show(step: str, result: object) -> bool:
    succeeded = bool(getattr(result, "succeeded"))
    marker = "OK" if succeeded else "ERROR"
    print(f"[{marker}] {step}: {getattr(result, 'message')}")
    value = getattr(result, "value")
    if succeeded and isinstance(value, list):
        for name in value:
            print(f"  - {name}")
    return succeeded


def run_sync(settings: BlobStorageSettings) -> None:
    print("\n=== Synchronous demo ===")
    client, credential = create_blob_service_client(settings)
    manager = BlobStorageManager(
        client, CONTAINER_NAME, max_concurrency=settings.max_concurrency
    )
    blob_name = "sync/sample.txt"
    try:
        upload = manager.upload(
            SAMPLE_FILE,
            blob_name,
            metadata=METADATA,
            tags=TAGS,
            timeout=REQUEST_TIMEOUT,
        )
        if not show("Upload with index tags", upload):
            return

        show("List blobs", manager.list_blobs(timeout=REQUEST_TIMEOUT))
        show(
            "Download",
            manager.download(blob_name, SYNC_DOWNLOAD, timeout=REQUEST_TIMEOUT),
        )

        lease_result = manager.acquire_lease(blob_name, timeout=REQUEST_TIMEOUT)
        if not show("Acquire lease", lease_result) or lease_result.value is None:
            return

        lease = lease_result.value
        try:
            SAMPLE_FILE.write_text("Updated safely under a synchronous lease.\n")
            show(
                "Overwrite while holding lease",
                manager.upload(
                    SAMPLE_FILE,
                    blob_name,
                    metadata=METADATA,
                    tags=TAGS,
                    lease=lease,
                    timeout=REQUEST_TIMEOUT,
                ),
            )
            show(
                "Delete",
                manager.delete(blob_name, lease=lease, timeout=REQUEST_TIMEOUT),
            )
        finally:
            # Deleting a leased blob releases the lease; a failed delete still needs cleanup.
            manager.release_lease(lease, timeout=REQUEST_TIMEOUT)
    finally:
        client.close()
        credential.close()


async def run_async(settings: BlobStorageSettings) -> None:
    print("\n=== Asynchronous demo ===")
    client, credential = create_async_blob_service_client(settings)
    manager = AsyncBlobStorageManager(
        client,
        CONTAINER_NAME,
        max_concurrency=settings.max_concurrency,
        upload_chunk_size=settings.max_block_size,
    )
    blob_name = "async/sample.txt"
    try:
        SAMPLE_FILE.write_text("Initial asynchronous sample content.\n")
        upload = await manager.upload(
            SAMPLE_FILE,
            blob_name,
            metadata=METADATA,
            tags=TAGS,
            timeout=REQUEST_TIMEOUT,
        )
        if not show("Upload with index tags", upload):
            return

        show("List blobs", await manager.list_blobs(timeout=REQUEST_TIMEOUT))
        show(
            "Download",
            await manager.download(
                blob_name, ASYNC_DOWNLOAD, timeout=REQUEST_TIMEOUT
            ),
        )

        lease_result = await manager.acquire_lease(blob_name, timeout=REQUEST_TIMEOUT)
        if not show("Acquire lease", lease_result) or lease_result.value is None:
            return

        lease = lease_result.value
        try:
            SAMPLE_FILE.write_text("Updated safely under an asynchronous lease.\n")
            show(
                "Overwrite while holding lease",
                await manager.upload(
                    SAMPLE_FILE,
                    blob_name,
                    metadata=METADATA,
                    tags=TAGS,
                    lease=lease,
                    timeout=REQUEST_TIMEOUT,
                ),
            )
            show(
                "Delete",
                await manager.delete(
                    blob_name, lease=lease, timeout=REQUEST_TIMEOUT
                ),
            )
        finally:
            await manager.release_lease(lease, timeout=REQUEST_TIMEOUT)
    finally:
        await client.close()
        await credential.close()


async def main() -> None:
    try:
        settings = BlobStorageSettings.from_env()
    except (TypeError, ValueError) as error:
        print(f"Configuration error: {error}")
        return

    SAMPLE_FILE.write_text("Initial synchronous sample content.\n")
    run_sync(settings)
    await run_async(settings)


if __name__ == "__main__":
    asyncio.run(main())
