"""Demonstrate synchronous and asynchronous blob management operations."""

from __future__ import annotations

import asyncio
import os
import tempfile
from pathlib import Path

from blob_manager import (
    AsyncBlobStorageService,
    BlobStorageService,
    OperationResult,
    StorageSettings,
)
from blob_manager.config import ConfigurationError

OPERATION_TIMEOUT = 120


def show(step: str, result: OperationResult[object]) -> None:
    status = "OK" if result.success else "ERROR"
    print(f"[{status}] {step}: {result.message}")


def run_sync_demo(
    settings: StorageSettings, container_name: str, workspace: Path
) -> None:
    print("\n=== Synchronous demo ===")
    source = workspace / "sync-sample.txt"
    download = workspace / "sync-downloaded.txt"
    blob_name = "demo/sync-sample.txt"
    source.write_text("Initial synchronous sample.\n", encoding="utf-8")

    with BlobStorageService(settings, container_name) as service:
        upload = service.upload(
            source,
            blob_name,
            metadata={"demo": "sync"},
            tags={"environment": "demo", "implementation": "sync"},
            timeout=OPERATION_TIMEOUT,
        )
        show("upload", upload)

        listed = service.list_blobs(prefix="demo/", timeout=OPERATION_TIMEOUT)
        show("list", listed)
        if listed.success:
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes), tags={blob.tags}")

        show(
            "download",
            service.download(blob_name, download, timeout=OPERATION_TIMEOUT),
        )

        lease = service.acquire_lease(blob_name, timeout=OPERATION_TIMEOUT)
        show("acquire lease", lease)
        if lease.success and lease.value:
            source.write_text(
                "Overwritten while holding the synchronous lease.\n",
                encoding="utf-8",
            )
            show(
                "lease-protected overwrite",
                service.upload(
                    source,
                    blob_name,
                    metadata={"demo": "sync", "version": "2"},
                    tags={"environment": "demo", "implementation": "sync"},
                    lease_id=lease.value,
                    timeout=OPERATION_TIMEOUT,
                ),
            )
            show(
                "release lease",
                service.release_lease(
                    blob_name, lease.value, timeout=OPERATION_TIMEOUT
                ),
            )

        show("delete", service.delete(blob_name, timeout=OPERATION_TIMEOUT))


async def run_async_demo(
    settings: StorageSettings, container_name: str, workspace: Path
) -> None:
    print("\n=== Asynchronous demo ===")
    source = workspace / "async-sample.txt"
    download = workspace / "async-downloaded.txt"
    blob_name = "demo/async-sample.txt"
    source.write_text("Initial asynchronous sample.\n", encoding="utf-8")

    async with AsyncBlobStorageService(settings, container_name) as service:
        upload = await service.upload(
            source,
            blob_name,
            metadata={"demo": "async"},
            tags={"environment": "demo", "implementation": "async"},
            timeout=OPERATION_TIMEOUT,
        )
        show("upload", upload)

        listed = await service.list_blobs(
            prefix="demo/", timeout=OPERATION_TIMEOUT
        )
        show("list", listed)
        if listed.success:
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes), tags={blob.tags}")

        show(
            "download",
            await service.download(
                blob_name, download, timeout=OPERATION_TIMEOUT
            ),
        )

        lease = await service.acquire_lease(
            blob_name, timeout=OPERATION_TIMEOUT
        )
        show("acquire lease", lease)
        if lease.success and lease.value:
            source.write_text(
                "Overwritten while holding the asynchronous lease.\n",
                encoding="utf-8",
            )
            show(
                "lease-protected overwrite",
                await service.upload(
                    source,
                    blob_name,
                    metadata={"demo": "async", "version": "2"},
                    tags={"environment": "demo", "implementation": "async"},
                    lease_id=lease.value,
                    timeout=OPERATION_TIMEOUT,
                ),
            )
            show(
                "release lease",
                await service.release_lease(
                    blob_name, lease.value, timeout=OPERATION_TIMEOUT
                ),
            )

        show(
            "delete",
            await service.delete(blob_name, timeout=OPERATION_TIMEOUT),
        )


def main() -> int:
    try:
        settings = StorageSettings.from_env()
    except ConfigurationError as error:
        print(f"Configuration error: {error}")
        return 2

    container_name = os.getenv("AZURE_STORAGE_CONTAINER", "blob-manager-demo")
    with tempfile.TemporaryDirectory(prefix="blob-manager-demo-") as temp_dir:
        workspace = Path(temp_dir)
        run_sync_demo(settings, container_name, workspace)
        asyncio.run(run_async_demo(settings, container_name, workspace))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
