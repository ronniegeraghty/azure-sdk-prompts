"""Demonstrate the synchronous and asynchronous blob management services."""

from __future__ import annotations

import asyncio
import os
from pathlib import Path
from typing import Any

from azure_blob_manager.config import (
    StorageSettings,
    create_async_client,
    create_sync_client,
)
from azure_blob_manager.service import (
    AsyncBlobStorageService,
    BlobStorageService,
    OperationResult,
)

CONTAINER = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
BLOB_NAME = os.getenv("AZURE_STORAGE_BLOB_NAME", "blob-manager-sample.txt")
SAMPLE_FILE = Path("blob-manager-sample.txt")
SYNC_DOWNLOAD = Path("blob-manager-sync-download.txt")
ASYNC_DOWNLOAD = Path("blob-manager-async-download.txt")
OPERATION_TIMEOUT = 120.0


def _print_result(step: str, result: OperationResult[Any]) -> bool:
    status = "OK" if result.success else "ERROR"
    print(f"[{status}] {step}: {result.message}")
    return result.success


def _print_blob_names(result: OperationResult[Any]) -> None:
    if not result.success or not result.value:
        return
    for blob in result.value:
        print(f"       - {blob.name} ({blob.size} bytes), tags={dict(blob.tags)}")


def run_sync_demo(settings: StorageSettings) -> None:
    print("\n=== Synchronous demo ===")
    client, credential = create_sync_client(settings)
    service = BlobStorageService(client, max_concurrency=settings.max_concurrency)
    lease = None

    try:
        SAMPLE_FILE.write_text("Initial content from the synchronous demo.\n", encoding="utf-8")
        upload = service.upload_file(
            CONTAINER,
            BLOB_NAME,
            SAMPLE_FILE,
            metadata={"demo": "sync"},
            tags={"project": "blob-manager", "mode": "sync"},
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("upload", upload)
        if not upload.success:
            print("[ERROR] Sync demo stopped to avoid changing a blob it did not create.")
            return

        listed = service.list_blobs(CONTAINER, timeout=OPERATION_TIMEOUT)
        _print_result("list", listed)
        _print_blob_names(listed)

        downloaded = service.download_file(
            CONTAINER,
            BLOB_NAME,
            SYNC_DOWNLOAD,
            overwrite_local=True,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("download", downloaded)

        lease_result = service.acquire_lease(
            CONTAINER,
            BLOB_NAME,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("acquire lease", lease_result)
        lease = lease_result.value

        if lease is not None:
            SAMPLE_FILE.write_text(
                "Lease-protected overwrite from the synchronous demo.\n",
                encoding="utf-8",
            )
            overwritten = service.upload_file(
                CONTAINER,
                BLOB_NAME,
                SAMPLE_FILE,
                metadata={"demo": "sync", "updated": "true"},
                tags={"project": "blob-manager", "mode": "sync"},
                overwrite=True,
                lease=lease,
                timeout=OPERATION_TIMEOUT,
            )
            _print_result("lease-protected overwrite", overwritten)

        deleted = service.delete_blob(
            CONTAINER,
            BLOB_NAME,
            lease=lease,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("delete", deleted)
    finally:
        client.close()
        credential.close()


async def run_async_demo(settings: StorageSettings) -> None:
    print("\n=== Asynchronous demo ===")
    client, credential = create_async_client(settings)
    service = AsyncBlobStorageService(client, max_concurrency=settings.max_concurrency)
    lease = None

    try:
        SAMPLE_FILE.write_text("Initial content from the asynchronous demo.\n", encoding="utf-8")
        upload = await service.upload_file(
            CONTAINER,
            BLOB_NAME,
            SAMPLE_FILE,
            metadata={"demo": "async"},
            tags={"project": "blob-manager", "mode": "async"},
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("upload", upload)
        if not upload.success:
            print("[ERROR] Async demo stopped to avoid changing a blob it did not create.")
            return

        listed = await service.list_blobs(CONTAINER, timeout=OPERATION_TIMEOUT)
        _print_result("list", listed)
        _print_blob_names(listed)

        downloaded = await service.download_file(
            CONTAINER,
            BLOB_NAME,
            ASYNC_DOWNLOAD,
            overwrite_local=True,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("download", downloaded)

        lease_result = await service.acquire_lease(
            CONTAINER,
            BLOB_NAME,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("acquire lease", lease_result)
        lease = lease_result.value

        if lease is not None:
            SAMPLE_FILE.write_text(
                "Lease-protected overwrite from the asynchronous demo.\n",
                encoding="utf-8",
            )
            overwritten = await service.upload_file(
                CONTAINER,
                BLOB_NAME,
                SAMPLE_FILE,
                metadata={"demo": "async", "updated": "true"},
                tags={"project": "blob-manager", "mode": "async"},
                overwrite=True,
                lease=lease,
                timeout=OPERATION_TIMEOUT,
            )
            _print_result("lease-protected overwrite", overwritten)

        deleted = await service.delete_blob(
            CONTAINER,
            BLOB_NAME,
            lease=lease,
            timeout=OPERATION_TIMEOUT,
        )
        _print_result("delete", deleted)
    finally:
        await client.close()
        await credential.close()


def main() -> int:
    try:
        settings = StorageSettings.from_env()
    except ValueError as error:
        print(f"[ERROR] Configuration: {error}")
        return 2

    run_sync_demo(settings)
    asyncio.run(run_async_demo(settings))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
