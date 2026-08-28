"""Demonstrate synchronous and asynchronous blob management operations."""

from __future__ import annotations

import asyncio
import logging
import os
from pathlib import Path

from blob_manager import (
    AsyncBlobStorageManager,
    BlobStorageManager,
    OperationResult,
    StorageSettings,
)

CONTAINER = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
SAMPLE_FILE = Path("sample.txt")
SYNC_DOWNLOAD = Path("downloads/sync-sample.txt")
ASYNC_DOWNLOAD = Path("downloads/async-sample.txt")


def print_result(step: str, result: OperationResult[object]) -> None:
    marker = "OK" if result.success else "ERROR"
    print(f"[{marker}] {step}: {result.message}")


def print_listing(result: OperationResult[object]) -> None:
    print_result("List", result)
    if result.success and isinstance(result.value, list):
        for blob in result.value:
            print(f"      {blob.name} ({blob.size} bytes, tags={blob.tags})")


def run_sync_demo(settings: StorageSettings) -> None:
    blob_name = "demo/sync-sample.txt"
    print("\n--- Synchronous demo ---")
    SAMPLE_FILE.write_text("Initial synchronous content.\n", encoding="utf-8")

    with BlobStorageManager(settings) as manager:
        print_result(
            "Upload",
            manager.upload(
                CONTAINER,
                blob_name,
                SAMPLE_FILE,
                metadata={"source": "sync-demo"},
                tags={"project": "blob-manager", "mode": "sync"},
                timeout=120,
            ),
        )
        print_listing(manager.list_blobs(CONTAINER, prefix="demo/", timeout=60))
        print_result(
            "Download",
            manager.download(CONTAINER, blob_name, SYNC_DOWNLOAD, timeout=120),
        )

        lease_result = manager.acquire_lease(
            CONTAINER, blob_name, duration=30, timeout=30
        )
        print_result("Acquire lease", lease_result)
        if lease_result.success and lease_result.value is not None:
            lease = lease_result.value
            try:
                SAMPLE_FILE.write_text(
                    "Synchronous content overwritten while leased.\n",
                    encoding="utf-8",
                )
                print_result(
                    "Leased overwrite",
                    manager.upload(
                        CONTAINER,
                        blob_name,
                        SAMPLE_FILE,
                        tags={"project": "blob-manager", "mode": "sync"},
                        lease=lease,
                        timeout=120,
                    ),
                )
            finally:
                print_result("Release lease", manager.release_lease(lease, timeout=30))

        print_result("Delete", manager.delete(CONTAINER, blob_name, timeout=60))


async def run_async_demo(settings: StorageSettings) -> None:
    blob_name = "demo/async-sample.txt"
    print("\n--- Asynchronous demo ---")
    await asyncio.to_thread(
        SAMPLE_FILE.write_text, "Initial asynchronous content.\n", encoding="utf-8"
    )

    async with AsyncBlobStorageManager(settings) as manager:
        print_result(
            "Upload",
            await manager.upload(
                CONTAINER,
                blob_name,
                SAMPLE_FILE,
                metadata={"source": "async-demo"},
                tags={"project": "blob-manager", "mode": "async"},
                timeout=120,
            ),
        )
        print_listing(
            await manager.list_blobs(CONTAINER, prefix="demo/", timeout=60)
        )
        print_result(
            "Download",
            await manager.download(
                CONTAINER, blob_name, ASYNC_DOWNLOAD, timeout=120
            ),
        )

        lease_result = await manager.acquire_lease(
            CONTAINER, blob_name, duration=30, timeout=30
        )
        print_result("Acquire lease", lease_result)
        if lease_result.success and lease_result.value is not None:
            lease = lease_result.value
            try:
                await asyncio.to_thread(
                    SAMPLE_FILE.write_text,
                    "Asynchronous content overwritten while leased.\n",
                    encoding="utf-8",
                )
                print_result(
                    "Leased overwrite",
                    await manager.upload(
                        CONTAINER,
                        blob_name,
                        SAMPLE_FILE,
                        tags={"project": "blob-manager", "mode": "async"},
                        lease=lease,
                        timeout=120,
                    ),
                )
            finally:
                print_result(
                    "Release lease", await manager.release_lease(lease, timeout=30)
                )

        print_result(
            "Delete", await manager.delete(CONTAINER, blob_name, timeout=60)
        )


def main() -> int:
    logging.basicConfig(
        level=os.getenv("APP_LOG_LEVEL", "INFO"),
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    try:
        settings = StorageSettings.from_env()
    except (TypeError, ValueError) as error:
        print(f"[ERROR] Configuration: {error}")
        return 2

    run_sync_demo(settings)
    asyncio.run(run_async_demo(settings))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
