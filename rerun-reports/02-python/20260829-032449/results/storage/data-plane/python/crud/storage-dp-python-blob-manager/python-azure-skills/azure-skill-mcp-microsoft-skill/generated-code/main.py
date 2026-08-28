"""Demonstrate synchronous and asynchronous blob management."""

from __future__ import annotations

import asyncio
import logging
import tempfile
from pathlib import Path

from azure.core.exceptions import AzureError

from blob_manager import (
    AsyncBlobStorageService,
    BlobOperationResult,
    BlobStorageService,
    StorageSettings,
    create_async_blob_service,
    create_sync_blob_service,
)

REQUEST_TIMEOUT = 120.0


def print_result(step: str, result: BlobOperationResult[object]) -> None:
    marker = "OK" if result.succeeded else "ERROR"
    print(f"[{marker}] {step}: {result.message}")
    if result.succeeded and isinstance(result.value, list):
        for blob in result.value:
            print(f"       - {blob.name} ({blob.size} bytes, etag={blob.etag})")


def run_sync_demo(settings: StorageSettings, workspace: Path) -> None:
    print("\n=== Synchronous demo ===")
    source = workspace / "sync-sample.txt"
    downloaded = workspace / "sync-downloaded.txt"
    source.write_text("Initial synchronous content.\n", encoding="utf-8")
    blob_name = "blob-manager-demo/sync-sample.txt"

    with create_sync_blob_service(settings) as client:
        service = BlobStorageService(
            client,
            settings.container_name,
            max_concurrency=settings.max_concurrency,
        )
        print_result(
            "upload",
            service.upload(
                source,
                blob_name,
                metadata={"demo": "sync"},
                tags={"Project": "BlobManager", "Mode": "Sync"},
                timeout=REQUEST_TIMEOUT,
            ),
        )
        print_result(
            "list", service.list_blobs(name_starts_with="blob-manager-demo/")
        )
        print_result(
            "download",
            service.download(blob_name, downloaded, timeout=REQUEST_TIMEOUT),
        )

        lease_result = service.acquire_lease(blob_name, timeout=REQUEST_TIMEOUT)
        print_result("acquire lease", lease_result)
        if lease_result.succeeded and lease_result.value is not None:
            lease = lease_result.value
            try:
                source.write_text(
                    "Synchronous content overwritten while leased.\n",
                    encoding="utf-8",
                )
                print_result(
                    "leased overwrite",
                    service.upload(
                        source,
                        blob_name,
                        metadata={"demo": "sync", "revision": "leased"},
                        tags={"Project": "BlobManager", "Mode": "Sync"},
                        lease=lease,
                        timeout=REQUEST_TIMEOUT,
                    ),
                )
            finally:
                try:
                    lease.release(timeout=int(REQUEST_TIMEOUT))
                    print("[OK] release lease")
                except AzureError as exc:
                    print(f"[ERROR] release lease: {exc}")

        print_result("delete", service.delete(blob_name, timeout=REQUEST_TIMEOUT))


async def run_async_demo(settings: StorageSettings, workspace: Path) -> None:
    print("\n=== Asynchronous demo ===")
    source = workspace / "async-sample.txt"
    downloaded = workspace / "async-downloaded.txt"
    await asyncio.to_thread(
        source.write_text, "Initial asynchronous content.\n", encoding="utf-8"
    )
    blob_name = "blob-manager-demo/async-sample.txt"

    async with create_async_blob_service(settings) as client:
        service = AsyncBlobStorageService(
            client,
            settings.container_name,
            max_concurrency=settings.max_concurrency,
        )
        print_result(
            "upload",
            await service.upload(
                source,
                blob_name,
                metadata={"demo": "async"},
                tags={"Project": "BlobManager", "Mode": "Async"},
                timeout=REQUEST_TIMEOUT,
            ),
        )
        print_result(
            "list",
            await service.list_blobs(
                name_starts_with="blob-manager-demo/",
                timeout=REQUEST_TIMEOUT,
            ),
        )
        print_result(
            "download",
            await service.download(blob_name, downloaded, timeout=REQUEST_TIMEOUT),
        )

        lease_result = await service.acquire_lease(
            blob_name, timeout=REQUEST_TIMEOUT
        )
        print_result("acquire lease", lease_result)
        if lease_result.succeeded and lease_result.value is not None:
            lease = lease_result.value
            try:
                await asyncio.to_thread(
                    source.write_text,
                    "Asynchronous content overwritten while leased.\n",
                    encoding="utf-8",
                )
                print_result(
                    "leased overwrite",
                    await service.upload(
                        source,
                        blob_name,
                        metadata={"demo": "async", "revision": "leased"},
                        tags={"Project": "BlobManager", "Mode": "Async"},
                        lease=lease,
                        timeout=REQUEST_TIMEOUT,
                    ),
                )
            finally:
                try:
                    await lease.release(timeout=int(REQUEST_TIMEOUT))
                    print("[OK] release lease")
                except AzureError as exc:
                    print(f"[ERROR] release lease: {exc}")

        print_result(
            "delete", await service.delete(blob_name, timeout=REQUEST_TIMEOUT)
        )


async def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    try:
        settings = StorageSettings.from_env()
    except ValueError as exc:
        print(f"[ERROR] configuration: {exc}")
        return

    with tempfile.TemporaryDirectory(prefix="blob-manager-demo-") as directory:
        workspace = Path(directory)
        run_sync_demo(settings, workspace)
        await run_async_demo(settings, workspace)


if __name__ == "__main__":
    asyncio.run(main())
