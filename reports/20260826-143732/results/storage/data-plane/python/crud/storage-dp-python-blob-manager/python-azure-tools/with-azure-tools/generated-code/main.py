"""Demonstrate sync and async Azure Blob Storage management operations."""

from __future__ import annotations

import asyncio
import os
import sys
from pathlib import Path

from blob_manager import (
    AsyncBlobStorageManager,
    BlobStorageManager,
    BlobStorageSettings,
    OperationResult,
)

SAMPLE_PATH = Path("sample-upload.txt")
SYNC_DOWNLOAD_PATH = Path("downloads/sync-sample.txt")
ASYNC_DOWNLOAD_PATH = Path("downloads/async-sample.txt")
TIMEOUT_SECONDS = 60


def show(step: str, result: OperationResult[object]) -> bool:
    status = "OK" if result.success else "ERROR"
    print(f"[{status}] {step}: {result.message}")
    return result.success


def run_sync(settings: BlobStorageSettings, container: str) -> bool:
    print("\n=== Synchronous demo ===")
    blob_name = "blob-manager-sync-sample.txt"
    with BlobStorageManager(settings) as manager:
        result = manager.upload(
            container,
            blob_name,
            SAMPLE_PATH,
            metadata={"demo": "sync"},
            tags={"project": "blob-manager", "implementation": "sync"},
            timeout=TIMEOUT_SECONDS,
        )
        if not show("Upload", result):
            return False

        listed = manager.list_blobs(
            container, prefix="blob-manager-", timeout=TIMEOUT_SECONDS
        )
        if show("List", listed):
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes, tags={blob.tags})")

        if not show(
            "Download",
            manager.download(
                container, blob_name, SYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS
            ),
        ):
            return False

        lease_result = manager.acquire_lease(
            container, blob_name, timeout=TIMEOUT_SECONDS
        )
        if not show("Acquire lease", lease_result) or lease_result.value is None:
            return False

        SAMPLE_PATH.write_text("Updated by the synchronous lease holder.\n")
        overwrite = manager.upload(
            container,
            blob_name,
            SAMPLE_PATH,
            metadata={"demo": "sync", "state": "updated"},
            tags={"project": "blob-manager", "implementation": "sync"},
            lease_id=lease_result.value.lease_id,
            timeout=TIMEOUT_SECONDS,
        )
        show("Overwrite under lease", overwrite)
        released = manager.release_lease(
            lease_result.value, timeout=TIMEOUT_SECONDS
        )
        show("Release lease", released)
        deleted = manager.delete(container, blob_name, timeout=TIMEOUT_SECONDS)
        show("Delete", deleted)
        return overwrite.success and released.success and deleted.success


async def run_async(settings: BlobStorageSettings, container: str) -> bool:
    print("\n=== Asynchronous demo ===")
    blob_name = "blob-manager-async-sample.txt"
    async with AsyncBlobStorageManager(settings) as manager:
        result = await manager.upload(
            container,
            blob_name,
            SAMPLE_PATH,
            metadata={"demo": "async"},
            tags={"project": "blob-manager", "implementation": "async"},
            timeout=TIMEOUT_SECONDS,
        )
        if not show("Upload", result):
            return False

        listed = await manager.list_blobs(
            container, prefix="blob-manager-", timeout=TIMEOUT_SECONDS
        )
        if show("List", listed):
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes, tags={blob.tags})")

        if not show(
            "Download",
            await manager.download(
                container, blob_name, ASYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS
            ),
        ):
            return False

        lease_result = await manager.acquire_lease(
            container, blob_name, timeout=TIMEOUT_SECONDS
        )
        if not show("Acquire lease", lease_result) or lease_result.value is None:
            return False

        SAMPLE_PATH.write_text("Updated by the asynchronous lease holder.\n")
        overwrite = await manager.upload(
            container,
            blob_name,
            SAMPLE_PATH,
            metadata={"demo": "async", "state": "updated"},
            tags={"project": "blob-manager", "implementation": "async"},
            lease_id=lease_result.value.lease_id,
            timeout=TIMEOUT_SECONDS,
        )
        show("Overwrite under lease", overwrite)
        released = await manager.release_lease(
            lease_result.value, timeout=TIMEOUT_SECONDS
        )
        show("Release lease", released)
        deleted = await manager.delete(
            container, blob_name, timeout=TIMEOUT_SECONDS
        )
        show("Delete", deleted)
        return overwrite.success and released.success and deleted.success


async def main() -> int:
    try:
        settings = BlobStorageSettings.from_env()
    except ValueError as exc:
        print(f"Configuration error: {exc}", file=sys.stderr)
        return 2

    container = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
    SAMPLE_PATH.write_text("Azure Blob Storage manager demo.\n")
    try:
        if not run_sync(settings, container):
            return 1
        SAMPLE_PATH.write_text("Azure Blob Storage manager async demo.\n")
        return 0 if await run_async(settings, container) else 1
    finally:
        SAMPLE_PATH.unlink(missing_ok=True)


if __name__ == "__main__":
    raise SystemExit(asyncio.run(main()))
