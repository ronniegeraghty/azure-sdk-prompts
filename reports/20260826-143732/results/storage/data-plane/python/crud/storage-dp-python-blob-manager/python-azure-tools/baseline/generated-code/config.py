"""Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient, ExponentialRetry
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _read_int(name: str, default: int, minimum: int = 0) -> int:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer") from exc
    if value < minimum:
        raise ValueError(f"{name} must be at least {minimum}")
    return value


def _read_float(name: str, default: float, minimum: float = 0.0) -> float:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = float(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be a number") from exc
    if value < minimum:
        raise ValueError(f"{name} must be at least {minimum}")
    return value


@dataclass(frozen=True)
class StorageSettings:
    account_url: str
    retry_total: int = 5
    retry_delay: float = 1.0
    retry_increment: float = 2.0
    http_log_level: str = "WARNING"
    max_block_size: int = 4 * 1024 * 1024
    max_single_put_size: int = 8 * 1024 * 1024
    max_concurrency: int = 4

    @classmethod
    def from_env(cls) -> "StorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required, for example "
                "https://<account>.blob.core.windows.net"
            )

        parsed_url = urlparse(account_url)
        if parsed_url.scheme != "https" or not parsed_url.netloc:
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must be a valid HTTPS endpoint")

        log_level = os.getenv("AZURE_STORAGE_HTTP_LOG_LEVEL", "WARNING").upper()
        if log_level not in {"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL", "OFF"}:
            raise ValueError(
                "AZURE_STORAGE_HTTP_LOG_LEVEL must be DEBUG, INFO, WARNING, "
                "ERROR, CRITICAL, or OFF"
            )

        return cls(
            account_url=account_url,
            retry_total=_read_int("AZURE_STORAGE_RETRY_TOTAL", 5),
            retry_delay=_read_float("AZURE_STORAGE_RETRY_DELAY", 1.0),
            retry_increment=_read_float("AZURE_STORAGE_RETRY_INCREMENT", 2.0),
            http_log_level=log_level,
            max_block_size=_read_int(
                "AZURE_STORAGE_MAX_BLOCK_SIZE", 4 * 1024 * 1024, 1024 * 1024
            ),
            max_single_put_size=_read_int(
                "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE",
                8 * 1024 * 1024,
                1024 * 1024,
            ),
            max_concurrency=_read_int("AZURE_STORAGE_MAX_CONCURRENCY", 4, 1),
        )

    @property
    def logging_enabled(self) -> bool:
        return self.http_log_level != "OFF"

    def configure_logging(self) -> None:
        if not self.logging_enabled:
            return
        level = getattr(logging, self.http_log_level)
        logging.basicConfig(level=level)
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            level
        )

    def new_retry_policy(self) -> ExponentialRetry:
        return ExponentialRetry(
            initial_backoff=self.retry_delay,
            increment_base=self.retry_increment,
            retry_total=self.retry_total,
        )


def create_sync_client(
    settings: StorageSettings,
) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    """Create a synchronous client and its credential.

    The caller owns both returned objects and must close them.
    """
    settings.configure_logging()
    credential = DefaultAzureCredential()
    client = BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=settings.new_retry_policy(),
        logging_enable=settings.logging_enabled,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
    )
    return client, credential


def create_async_client(
    settings: StorageSettings,
) -> tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:
    """Create an asynchronous client and its credential.

    The caller owns both returned objects and must close them asynchronously.
    """
    settings.configure_logging()
    credential = AsyncDefaultAzureCredential()
    client = AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=settings.new_retry_policy(),
        logging_enable=settings.logging_enabled,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
    )
    return client, credential
