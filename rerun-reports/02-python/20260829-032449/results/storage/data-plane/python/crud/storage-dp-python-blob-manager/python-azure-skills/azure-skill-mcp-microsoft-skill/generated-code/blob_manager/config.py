"""Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from contextlib import asynccontextmanager, contextmanager
from dataclasses import dataclass
from typing import AsyncIterator, Iterator
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient, ExponentialRetry
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _positive_int(name: str, default: int) -> int:
    raw_value = os.getenv(name, str(default))
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer, got {raw_value!r}") from exc
    if value < 1:
        raise ValueError(f"{name} must be at least 1")
    return value


def _log_level(name: str, default: str) -> int:
    raw_value = os.getenv(name, default).upper()
    level = logging.getLevelNamesMapping().get(raw_value)
    if level is None:
        raise ValueError(f"{name} must be a valid Python logging level, got {raw_value!r}")
    return level


def _validate_account_url(account_url: str) -> str:
    parsed = urlparse(account_url)
    is_local_emulator = parsed.hostname in {"127.0.0.1", "localhost"}
    if parsed.scheme != "https" and not (parsed.scheme == "http" and is_local_emulator):
        raise ValueError(
            "AZURE_STORAGE_ACCOUNT_URL must use HTTPS "
            "(HTTP is allowed only for a local emulator)"
        )
    if not parsed.netloc:
        raise ValueError("AZURE_STORAGE_ACCOUNT_URL must be an absolute URL")
    return account_url.rstrip("/")


@dataclass(frozen=True, slots=True)
class StorageSettings:
    """Environment-driven client and transfer settings."""

    account_url: str
    container_name: str
    retry_total: int = 5
    retry_initial_delay: int = 2
    retry_increment: int = 2
    http_log_level: int = logging.WARNING
    max_concurrency: int = 4
    max_block_size: int = 8 * 1024 * 1024
    max_single_put_size: int = 64 * 1024 * 1024

    @property
    def http_logging_enabled(self) -> bool:
        return self.http_log_level <= logging.INFO

    @classmethod
    def from_env(cls) -> "StorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL")
        if not account_url:
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL is required")

        container_name = os.getenv("AZURE_STORAGE_CONTAINER")
        if not container_name:
            raise ValueError("AZURE_STORAGE_CONTAINER is required")

        return cls(
            account_url=_validate_account_url(account_url),
            container_name=container_name,
            retry_total=_positive_int("AZURE_STORAGE_RETRY_TOTAL", 5),
            retry_initial_delay=_positive_int("AZURE_STORAGE_RETRY_INITIAL_DELAY", 2),
            retry_increment=_positive_int("AZURE_STORAGE_RETRY_INCREMENT", 2),
            http_log_level=_log_level("AZURE_STORAGE_HTTP_LOG_LEVEL", "WARNING"),
            max_concurrency=_positive_int("AZURE_STORAGE_MAX_CONCURRENCY", 4),
            max_block_size=_positive_int(
                "AZURE_STORAGE_MAX_BLOCK_SIZE", 8 * 1024 * 1024
            ),
            max_single_put_size=_positive_int(
                "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE", 64 * 1024 * 1024
            ),
        )


def configure_logging(settings: StorageSettings) -> None:
    """Configure Azure HTTP pipeline logging without changing application logging."""
    logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
        settings.http_log_level
    )


def _retry_policy(settings: StorageSettings) -> ExponentialRetry:
    return ExponentialRetry(
        retry_total=settings.retry_total,
        initial_backoff=settings.retry_initial_delay,
        increment_base=settings.retry_increment,
    )


def _client_options(settings: StorageSettings) -> dict[str, object]:
    return {
        "retry_policy": _retry_policy(settings),
        "logging_enable": settings.http_logging_enabled,
        "max_block_size": settings.max_block_size,
        "max_single_put_size": settings.max_single_put_size,
    }


@contextmanager
def create_sync_blob_service(
    settings: StorageSettings,
) -> Iterator[BlobServiceClient]:
    """Create and deterministically close a synchronous service client."""
    configure_logging(settings)
    with DefaultAzureCredential() as credential:
        with BlobServiceClient(
            account_url=settings.account_url,
            credential=credential,
            **_client_options(settings),
        ) as client:
            yield client


@asynccontextmanager
async def create_async_blob_service(
    settings: StorageSettings,
) -> AsyncIterator[AsyncBlobServiceClient]:
    """Create and deterministically close an asynchronous service client."""
    configure_logging(settings)
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncBlobServiceClient(
            account_url=settings.account_url,
            credential=credential,
            **_client_options(settings),
        ) as client:
            yield client
