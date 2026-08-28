"""Secure Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass
from urllib.parse import urlparse

from azure.core.pipeline.policies import AsyncRetryPolicy, RetryPolicy
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _read_int(name: str, default: int, minimum: int = 0) -> int:
    raw = os.getenv(name, str(default))
    try:
        value = int(raw)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer, got {raw!r}") from exc
    if value < minimum:
        raise ValueError(f"{name} must be at least {minimum}")
    return value


def _read_float(name: str, default: float, minimum: float = 0.0) -> float:
    raw = os.getenv(name, str(default))
    try:
        value = float(raw)
    except ValueError as exc:
        raise ValueError(f"{name} must be a number, got {raw!r}") from exc
    if value < minimum:
        raise ValueError(f"{name} must be at least {minimum}")
    return value


def _read_bool(name: str, default: bool) -> bool:
    raw = os.getenv(name)
    if raw is None:
        return default
    normalized = raw.strip().lower()
    if normalized in {"1", "true", "yes", "on"}:
        return True
    if normalized in {"0", "false", "no", "off"}:
        return False
    raise ValueError(f"{name} must be true or false, got {raw!r}")


@dataclass(frozen=True, slots=True)
class BlobStorageSettings:
    """Settings loaded from environment variables."""

    account_url: str
    retry_total: int = 5
    retry_backoff_factor: float = 0.8
    retry_backoff_max: float = 30.0
    logging_enabled: bool = False
    logging_level: str = "WARNING"
    max_block_size: int = 8 * 1024 * 1024
    max_single_put_size: int = 8 * 1024 * 1024
    max_concurrency: int = 4
    connection_timeout: int = 20
    read_timeout: int = 120

    @classmethod
    def from_env(cls) -> "BlobStorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required "
                "(for example, https://myaccount.blob.core.windows.net)"
            )
        parsed = urlparse(account_url)
        if parsed.scheme != "https" or not parsed.netloc:
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must be a valid HTTPS endpoint")

        logging_level = os.getenv("AZURE_STORAGE_LOG_LEVEL", "WARNING").upper()
        if logging_level not in logging.getLevelNamesMapping():
            raise ValueError(f"Invalid AZURE_STORAGE_LOG_LEVEL: {logging_level!r}")

        return cls(
            account_url=account_url,
            retry_total=_read_int("AZURE_STORAGE_RETRY_TOTAL", 5),
            retry_backoff_factor=_read_float("AZURE_STORAGE_RETRY_BACKOFF_FACTOR", 0.8),
            retry_backoff_max=_read_float("AZURE_STORAGE_RETRY_BACKOFF_MAX", 30.0),
            logging_enabled=_read_bool("AZURE_STORAGE_HTTP_LOGGING", False),
            logging_level=logging_level,
            max_block_size=_read_int(
                "AZURE_STORAGE_MAX_BLOCK_SIZE", 8 * 1024 * 1024, 1024 * 1024
            ),
            max_single_put_size=_read_int(
                "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE", 8 * 1024 * 1024, 1024 * 1024
            ),
            max_concurrency=_read_int("AZURE_STORAGE_MAX_CONCURRENCY", 4, 1),
            connection_timeout=_read_int("AZURE_STORAGE_CONNECTION_TIMEOUT", 20, 1),
            read_timeout=_read_int("AZURE_STORAGE_READ_TIMEOUT", 120, 1),
        )

    def configure_logging(self) -> None:
        logging.basicConfig(
            level=getattr(logging, self.logging_level),
            format="%(asctime)s %(levelname)s %(name)s: %(message)s",
        )
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            self.logging_level
        )


def create_sync_clients(
    settings: BlobStorageSettings,
) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    """Create a sync service client and its credential owner."""
    settings.configure_logging()
    credential = DefaultAzureCredential()
    retry_policy = RetryPolicy(
        retry_total=settings.retry_total,
        retry_backoff_factor=settings.retry_backoff_factor,
        retry_backoff_max=settings.retry_backoff_max,
    )
    client = BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=retry_policy,
        logging_enable=settings.logging_enabled,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
    )
    return client, credential


def create_async_clients(
    settings: BlobStorageSettings,
) -> tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:
    """Create an async service client and its credential owner."""
    settings.configure_logging()
    credential = AsyncDefaultAzureCredential()
    retry_policy = AsyncRetryPolicy(
        retry_total=settings.retry_total,
        retry_backoff_factor=settings.retry_backoff_factor,
        retry_backoff_max=settings.retry_backoff_max,
    )
    client = AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=retry_policy,
        logging_enable=settings.logging_enabled,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
    )
    return client, credential
