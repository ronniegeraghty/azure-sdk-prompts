"""Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass
from typing import Any
from urllib.parse import urlparse

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient, ExponentialRetry
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


class CappedExponentialRetry(ExponentialRetry):
    """Storage retry policy with exponential backoff and a deterministic ceiling."""

    def __init__(self, *, max_backoff: int, **kwargs: Any) -> None:
        super().__init__(**kwargs)
        self.max_backoff = max_backoff

    def get_backoff_time(self, settings: dict[str, Any]) -> float:
        return min(super().get_backoff_time(settings), self.max_backoff)


def _read_int(name: str, default: int, minimum: int = 0) -> int:
    raw_value = os.getenv(name, str(default))
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer, got {raw_value!r}.") from exc
    if value < minimum:
        raise ValueError(f"{name} must be at least {minimum}, got {value}.")
    return value


def _read_bool(name: str, default: bool) -> bool:
    raw_value = os.getenv(name, str(default)).strip().lower()
    if raw_value in {"1", "true", "yes", "on"}:
        return True
    if raw_value in {"0", "false", "no", "off"}:
        return False
    raise ValueError(f"{name} must be true or false, got {raw_value!r}.")


@dataclass(frozen=True, slots=True)
class StorageSettings:
    account_url: str
    retry_total: int = 5
    retry_initial_backoff: int = 2
    retry_increment_base: int = 2
    retry_jitter: int = 1
    retry_max_backoff: int = 30
    connection_timeout: int = 20
    read_timeout: int = 120
    max_concurrency: int = 4
    max_block_size: int = 8 * 1024 * 1024
    max_single_put_size: int = 64 * 1024 * 1024
    http_logging_enabled: bool = False
    http_logging_level: str = "WARNING"

    @classmethod
    def from_env(cls) -> "StorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required, for example "
                "'https://myaccount.blob.core.windows.net'."
            )

        parsed_url = urlparse(account_url)
        if parsed_url.scheme != "https" or not parsed_url.netloc:
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must be an absolute HTTPS endpoint.")

        logging_level = os.getenv("AZURE_STORAGE_LOG_LEVEL", "WARNING").upper()
        if logging_level not in logging.getLevelNamesMapping():
            raise ValueError(
                "AZURE_STORAGE_LOG_LEVEL must be a Python logging level such as "
                "DEBUG, INFO, WARNING, ERROR, or CRITICAL."
            )

        return cls(
            account_url=account_url,
            retry_total=_read_int("AZURE_STORAGE_RETRY_TOTAL", 5),
            retry_initial_backoff=_read_int("AZURE_STORAGE_RETRY_INITIAL_BACKOFF", 2),
            retry_increment_base=_read_int("AZURE_STORAGE_RETRY_INCREMENT_BASE", 2),
            retry_jitter=_read_int("AZURE_STORAGE_RETRY_JITTER", 1),
            retry_max_backoff=_read_int("AZURE_STORAGE_RETRY_MAX_BACKOFF", 30, 1),
            connection_timeout=_read_int("AZURE_STORAGE_CONNECTION_TIMEOUT", 20, 1),
            read_timeout=_read_int("AZURE_STORAGE_READ_TIMEOUT", 120, 1),
            max_concurrency=_read_int("AZURE_STORAGE_MAX_CONCURRENCY", 4, 1),
            max_block_size=_read_int("AZURE_STORAGE_MAX_BLOCK_SIZE_MIB", 8, 1)
            * 1024
            * 1024,
            max_single_put_size=_read_int("AZURE_STORAGE_MAX_SINGLE_PUT_SIZE_MIB", 64, 1)
            * 1024
            * 1024,
            http_logging_enabled=_read_bool("AZURE_STORAGE_HTTP_LOGGING", False),
            http_logging_level=logging_level,
        )


def configure_http_logging(settings: StorageSettings) -> None:
    """Configure Azure SDK HTTP pipeline logging without changing application loggers."""
    azure_http_logger = logging.getLogger("azure.core.pipeline.policies.http_logging_policy")
    azure_http_logger.setLevel(settings.http_logging_level)
    if not azure_http_logger.handlers:
        handler = logging.StreamHandler()
        handler.setFormatter(
            logging.Formatter("%(asctime)s %(levelname)s %(name)s: %(message)s")
        )
        azure_http_logger.addHandler(handler)
    azure_http_logger.propagate = False


def _retry_policy(settings: StorageSettings) -> CappedExponentialRetry:
    return CappedExponentialRetry(
        retry_total=settings.retry_total,
        initial_backoff=settings.retry_initial_backoff,
        increment_base=settings.retry_increment_base,
        random_jitter_range=settings.retry_jitter,
        max_backoff=settings.retry_max_backoff,
    )


def create_sync_client(
    settings: StorageSettings,
) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    configure_http_logging(settings)
    credential = DefaultAzureCredential()
    client = BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=_retry_policy(settings),
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
        logging_enable=settings.http_logging_enabled,
    )
    return client, credential


def create_async_client(
    settings: StorageSettings,
) -> tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:
    configure_http_logging(settings)
    credential = AsyncDefaultAzureCredential()
    client = AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=_retry_policy(settings),
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
        max_block_size=settings.max_block_size,
        max_single_put_size=settings.max_single_put_size,
        logging_enable=settings.http_logging_enabled,
    )
    return client, credential
