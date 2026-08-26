"""Demonstrate the synchronous and asynchronous blob services."""

from __future__ import annotations

import asyncio
import os
from pathlib import Path
from typing import Any

from blob_service import AsyncBlobStorageService, BlobStorageService, OperationResult
from config import StorageSettings, create_async_client, create_sync_client

CONTAINER_NAME = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
SAMPLE_PATH = Path("sample.txt")
SYNC_DOWNLOAD_PATH = Path("downloads") / "sample-sync.txt"
ASYNC_DOWNLOAD_PATH = Path("downloads") / "sample-async.txt"
TIMEOUT = 60


def show(step: str, result: OperationResult[Any]) -> bool:
    marker = "OK" if result.success else "ERROR"
    print(f"[{marker}] {step}: {result.message}")
    if result.success and isinstance(result.value, list):
        for blob in result.value:
            print(
                f"     - {blob['name']} ({blob['size']} bytes), "
                f"tags={blob['tags'] or {}}"
            )
    return result.success


def run_sync(settings: StorageSettings) -> None:
    print("\n=== Synchronous demo ===")
    client, credential = create_sync_client(settings)
    service = BlobStorageService(
        client, CONTAINER_NAME, max_concurrency=settings.max_concurrency
    )
    blob_name = "sync-sample.txt"

    try:
        if not show(
            "upload",
            service.upload_file(
                SAMPLE_PATH,
                blob_name,
                metadata={"demo": "sync"},
                tags={"project": "blob-manager", "mode": "sync"},
                timeout=TIMEOUT,
            ),
        ):
            return
        show("list", service.list_blobs(timeout=TIMEOUT))
        show(
            "download",
            service.download_file(blob_name, SYNC_DOWNLOAD_PATH, timeout=TIMEOUT),
        )

        lease_result = service.acquire_lease(blob_name, timeout=TIMEOUT)
        if show("acquire lease", lease_result):
            lease = lease_result.value
            try:
                SAMPLE_PATH.write_text(
                    "Updated by the synchronous lease holder.\n", encoding="utf-8"
                )
                show(
                    "leased overwrite",
                    service.upload_file(
                        SAMPLE_PATH,
                        blob_name,
                        metadata={"demo": "sync", "version": "2"},
                        tags={"project": "blob-manager", "mode": "sync"},
                        overwrite=True,
                        lease=lease,
                        timeout=TIMEOUT,
                    ),
                )
            finally:
                show("release lease", service.release_lease(lease, timeout=TIMEOUT))
        show("delete", service.delete_blob(blob_name, timeout=TIMEOUT))
    finally:
        client.close()
        credential.close()


async def run_async(settings: StorageSettings) -> None:
    print("\n=== Asynchronous demo ===")
    client, credential = create_async_client(settings)
    service = AsyncBlobStorageService(
        client, CONTAINER_NAME, max_concurrency=settings.max_concurrency
    )
    blob_name = "async-sample.txt"

    try:
        SAMPLE_PATH.write_text("Azure Blob Storage async demo.\n", encoding="utf-8")
        if not show(
            "upload",
            await service.upload_file(
                SAMPLE_PATH,
                blob_name,
                metadata={"demo": "async"},
                tags={"project": "blob-manager", "mode": "async"},
                timeout=TIMEOUT,
            ),
        ):
            return
        show("list", await service.list_blobs(timeout=TIMEOUT))
        show(
            "download",
            await service.download_file(
                blob_name, ASYNC_DOWNLOAD_PATH, timeout=TIMEOUT
            ),
        )

        lease_result = await service.acquire_lease(blob_name, timeout=TIMEOUT)
        if show("acquire lease", lease_result):
            lease = lease_result.value
            try:
                SAMPLE_PATH.write_text(
                    "Updated by the asynchronous lease holder.\n", encoding="utf-8"
                )
                show(
                    "leased overwrite",
                    await service.upload_file(
                        SAMPLE_PATH,
                        blob_name,
                        metadata={"demo": "async", "version": "2"},
                        tags={"project": "blob-manager", "mode": "async"},
                        overwrite=True,
                        lease=lease,
                        timeout=TIMEOUT,
                    ),
                )
            finally:
                show(
                    "release lease",
                    await service.release_lease(lease, timeout=TIMEOUT),
                )
        show("delete", await service.delete_blob(blob_name, timeout=TIMEOUT))
    finally:
        await client.close()
        await credential.close()


def main() -> None:
    try:
        settings = StorageSettings.from_env()
    except ValueError as exc:
        print(f"Configuration error: {exc}")
        return

    SAMPLE_PATH.write_text("Azure Blob Storage sync demo.\n", encoding="utf-8")
    run_sync(settings)
    asyncio.run(run_async(settings))


if __name__ == "__main__":
    main()
