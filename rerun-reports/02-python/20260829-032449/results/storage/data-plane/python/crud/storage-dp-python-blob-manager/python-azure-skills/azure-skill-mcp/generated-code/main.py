"""Demonstrate the synchronous and asynchronous blob management services."""

from __future__ import annotations

import asyncio
from pathlib import Path

from blob_manager import (
    AsyncBlobStorageService,
    BlobStorageService,
    BlobStorageSettings,
    create_async_client,
    create_sync_client,
)

SAMPLE_PATH = Path("sample.txt")
SYNC_DOWNLOAD_PATH = Path("downloads/sync-sample.txt")
ASYNC_DOWNLOAD_PATH = Path("downloads/async-sample.txt")
SYNC_BLOB_NAME = "demo/sync-sample.txt"
ASYNC_BLOB_NAME = "demo/async-sample.txt"
TIMEOUT_SECONDS = 120
TAGS = {"project": "blob-manager", "environment": "demo"}


def print_result(step: str, result: object) -> None:
    message = getattr(result, "message", str(result))
    succeeded = getattr(result, "succeeded", False)
    print(f"[{'OK' if succeeded else 'ERROR'}] {step}: {message}")


def run_sync(settings: BlobStorageSettings) -> None:
    print("\n=== Synchronous demo ===")
    client, credential = create_sync_client(settings)
    service = BlobStorageService(
        client,
        settings.container_name,
        max_concurrency=settings.max_concurrency,
    )
    try:
        upload = service.upload(
            SAMPLE_PATH,
            SYNC_BLOB_NAME,
            metadata={"source": "sync-demo"},
            tags=TAGS,
            timeout=TIMEOUT_SECONDS,
        )
        print_result("upload", upload)
        if not upload.succeeded:
            return

        listing = service.list_blobs(timeout=TIMEOUT_SECONDS)
        print_result("list", listing)
        if listing.value:
            for blob in listing.value:
                print(f"  - {blob.name} ({blob.size} bytes, tags={blob.tags})")

        download = service.download(
            SYNC_BLOB_NAME, SYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS
        )
        print_result("download", download)

        lease_result = service.acquire_lease(
            SYNC_BLOB_NAME, timeout=TIMEOUT_SECONDS
        )
        print_result("acquire lease", lease_result)
        if not lease_result.succeeded or lease_result.value is None:
            return

        lease = lease_result.value
        SAMPLE_PATH.write_text("Updated by the synchronous lease holder.\n")
        overwrite = service.upload(
            SAMPLE_PATH,
            SYNC_BLOB_NAME,
            metadata={"source": "sync-lease-demo"},
            tags=TAGS,
            lease=lease,
            timeout=TIMEOUT_SECONDS,
        )
        print_result("leased overwrite", overwrite)

        delete = service.delete(
            SYNC_BLOB_NAME, lease=lease, timeout=TIMEOUT_SECONDS
        )
        print_result("delete", delete)
    finally:
        client.close()
        credential.close()


async def run_async(settings: BlobStorageSettings) -> None:
    print("\n=== Asynchronous demo ===")
    client, credential = create_async_client(settings)
    service = AsyncBlobStorageService(
        client,
        settings.container_name,
        max_concurrency=settings.max_concurrency,
        block_size=settings.block_size,
    )
    try:
        SAMPLE_PATH.write_text("Hello from the asynchronous blob manager.\n")
        upload = await service.upload(
            SAMPLE_PATH,
            ASYNC_BLOB_NAME,
            metadata={"source": "async-demo"},
            tags=TAGS,
            timeout=TIMEOUT_SECONDS,
        )
        print_result("upload", upload)
        if not upload.succeeded:
            return

        listing = await service.list_blobs(timeout=TIMEOUT_SECONDS)
        print_result("list", listing)
        if listing.value:
            for blob in listing.value:
                print(f"  - {blob.name} ({blob.size} bytes, tags={blob.tags})")

        download = await service.download(
            ASYNC_BLOB_NAME, ASYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS
        )
        print_result("download", download)

        lease_result = await service.acquire_lease(
            ASYNC_BLOB_NAME, timeout=TIMEOUT_SECONDS
        )
        print_result("acquire lease", lease_result)
        if not lease_result.succeeded or lease_result.value is None:
            return

        lease = lease_result.value
        SAMPLE_PATH.write_text("Updated by the asynchronous lease holder.\n")
        overwrite = await service.upload(
            SAMPLE_PATH,
            ASYNC_BLOB_NAME,
            metadata={"source": "async-lease-demo"},
            tags=TAGS,
            lease=lease,
            timeout=TIMEOUT_SECONDS,
        )
        print_result("leased overwrite", overwrite)

        delete = await service.delete(
            ASYNC_BLOB_NAME, lease=lease, timeout=TIMEOUT_SECONDS
        )
        print_result("delete", delete)
    finally:
        await client.close()
        await credential.close()


def main() -> None:
    settings = BlobStorageSettings.from_env()
    settings.configure_logging()
    SAMPLE_PATH.write_text("Hello from the synchronous blob manager.\n")
    try:
        run_sync(settings)
        asyncio.run(run_async(settings))
    finally:
        SAMPLE_PATH.unlink(missing_ok=True)


if __name__ == "__main__":
    main()
