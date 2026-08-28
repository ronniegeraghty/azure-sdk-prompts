"""Demonstrate synchronous and asynchronous blob management."""

from __future__ import annotations

import asyncio
import os
import tempfile
from pathlib import Path

from azure_blob_manager import (
    AsyncBlobStorageService,
    BlobStorageService,
    BlobStorageSettings,
    LeaseHandle,
    OperationResult,
)
from azure_blob_manager.config import create_async_clients, create_sync_clients

OPERATION_TIMEOUT = 60.0


def _show(step: str, result: OperationResult[object]) -> bool:
    marker = "OK" if result.success else "ERROR"
    print(f"[{marker}] {step}: {result.message}")
    return result.success


def run_sync(
    settings: BlobStorageSettings,
    container_name: str,
    sample: Path,
    work_dir: Path,
) -> None:
    print("\n=== Synchronous demo ===")
    client, credential = create_sync_clients(settings)
    service = BlobStorageService(
        client, container_name, max_concurrency=settings.max_concurrency
    )
    blob_name = "blob-manager-demo/sync-sample.txt"
    lease_id: str | None = None
    try:
        uploaded = service.upload(
            sample,
            blob_name,
            metadata={"demo": "sync"},
            tags={"project": "blob-manager", "mode": "sync"},
            timeout=OPERATION_TIMEOUT,
        )
        if not _show("upload", uploaded):
            return

        listed = service.list_blobs(timeout=OPERATION_TIMEOUT)
        _show("list", listed)
        if listed.success:
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes, tags={dict(blob.tags)})")

        _show(
            "download",
            service.download(
                blob_name, work_dir / "sync-download.txt", timeout=OPERATION_TIMEOUT
            ),
        )

        lease = service.acquire_lease(blob_name, timeout=OPERATION_TIMEOUT)
        if not _show("acquire lease", lease) or lease.value is None:
            return
        lease_id = lease.value.lease_id

        sample.write_text("Overwritten by the synchronous demo.\n", encoding="utf-8")
        _show(
            "leased overwrite",
            service.upload(
                sample, blob_name, lease_id=lease_id, timeout=OPERATION_TIMEOUT
            ),
        )
        deleted = service.delete(
            blob_name, lease_id=lease_id, timeout=OPERATION_TIMEOUT
        )
        if _show("delete", deleted):
            lease_id = None
    finally:
        if lease_id:
            _show(
                "release lease after incomplete demo",
                service.release_lease(
                    LeaseHandle(blob_name, lease_id),
                    timeout=OPERATION_TIMEOUT,
                ),
            )
        client.close()
        credential.close()


async def run_async(
    settings: BlobStorageSettings,
    container_name: str,
    sample: Path,
    work_dir: Path,
) -> None:
    print("\n=== Asynchronous demo ===")
    client, credential = create_async_clients(settings)
    service = AsyncBlobStorageService(
        client,
        container_name,
        max_concurrency=settings.max_concurrency,
        chunk_size=settings.max_block_size,
    )
    blob_name = "blob-manager-demo/async-sample.txt"
    lease_id: str | None = None
    try:
        sample.write_text("Uploaded by the asynchronous demo.\n", encoding="utf-8")
        uploaded = await service.upload(
            sample,
            blob_name,
            metadata={"demo": "async"},
            tags={"project": "blob-manager", "mode": "async"},
            timeout=OPERATION_TIMEOUT,
        )
        if not _show("upload", uploaded):
            return

        listed = await service.list_blobs(timeout=OPERATION_TIMEOUT)
        _show("list", listed)
        if listed.success:
            for blob in listed.value or []:
                print(f"  - {blob.name} ({blob.size} bytes, tags={dict(blob.tags)})")

        _show(
            "download",
            await service.download(
                blob_name, work_dir / "async-download.txt", timeout=OPERATION_TIMEOUT
            ),
        )

        lease = await service.acquire_lease(blob_name, timeout=OPERATION_TIMEOUT)
        if not _show("acquire lease", lease) or lease.value is None:
            return
        lease_id = lease.value.lease_id

        sample.write_text("Overwritten by the asynchronous demo.\n", encoding="utf-8")
        _show(
            "leased overwrite",
            await service.upload(
                sample, blob_name, lease_id=lease_id, timeout=OPERATION_TIMEOUT
            ),
        )
        deleted = await service.delete(
            blob_name, lease_id=lease_id, timeout=OPERATION_TIMEOUT
        )
        if _show("delete", deleted):
            lease_id = None
    finally:
        if lease_id:
            _show(
                "release lease after incomplete demo",
                await service.release_lease(
                    LeaseHandle(blob_name, lease_id), timeout=OPERATION_TIMEOUT
                ),
            )
        await client.close()
        await credential.close()


async def main() -> None:
    try:
        settings = BlobStorageSettings.from_env()
        container_name = os.environ["AZURE_STORAGE_CONTAINER"]
    except (KeyError, ValueError) as exc:
        print(f"Configuration error: {exc}")
        return

    with tempfile.TemporaryDirectory(prefix="blob-manager-demo-") as directory:
        work_dir = Path(directory)
        sample = work_dir / "sample.txt"
        sample.write_text("Uploaded by the synchronous demo.\n", encoding="utf-8")
        run_sync(settings, container_name, sample, work_dir)
        await run_async(settings, container_name, sample, work_dir)


if __name__ == "__main__":
    asyncio.run(main())
