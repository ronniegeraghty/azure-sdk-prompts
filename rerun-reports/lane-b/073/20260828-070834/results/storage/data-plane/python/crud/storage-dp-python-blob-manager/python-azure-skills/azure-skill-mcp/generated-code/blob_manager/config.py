"""Secure Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient, ExponentialRetry
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

_LOG_LEVELS = {
    "CRITICAL": logging.CRITICAL,
    "ERROR": logging.ERROR,
    "WARNING": logging.WARNING,
    "INFO": logging.INFO,
    "DEBUG": logging.DEBUG,
}


def _env_bool(name: str, default: bool) -> bool:
    value = os.getenv(name)
    if value is None:
        return default
    return value.strip().lower() in {"1", "true", "yes", "on"}


@dataclass(frozen=True)
class BlobStorageSettings:
    """Settings loaded from environment variables."""

    account_url: str
    retry_total: int = 5
    retry_initial_backoff: int = 2
    retry_increment_base: int = 2
    http_logging_enabled: bool = False
    http_logging_level: str = "WARNING"
    max_block_size: int = 4 * 1024 * 1024
    max_single_put_size: int = 8 * 1024 * 1024
    max_concurrency: int = 4
    connection_timeout: int = 20
    read_timeout: int = 120

    @classmethod
    def from_env(cls) -> "BlobStorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required, for example "
                "'https://<account>.blob.core.windows.net'."
            )
        if not account_url.startswith("https://"):
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS.")

        log_level = os.getenv("AZURE_STORAGE_HTTP_LOG_LEVEL", "WARNING").upper()
        if log_level not in _LOG_LEVELS:
            raise ValueError(
                "AZURE_STORAGE_HTTP_LOG_LEVEL must be one of: "
                + ", ".join(_LOG_LEVELS)
            )

        return cls(
            account_url=account_url,
            retry_total=int(os.getenv("AZURE_STORAGE_RETRY_TOTAL", "5")),
            retry_initial_backoff=int(
                os.getenv("AZURE_STORAGE_RETRY_INITIAL_BACKOFF", "2")
            ),
            retry_increment_base=int(
                os.getenv("AZURE_STORAGE_RETRY_INCREMENT_BASE", "2")
            ),
            http_logging_enabled=_env_bool(
                "AZURE_STORAGE_HTTP_LOGGING_ENABLED", False
            ),
            http_logging_level=log_level,
            max_block_size=int(
                os.getenv("AZURE_STORAGE_MAX_BLOCK_SIZE", str(4 * 1024 * 1024))
            ),
            max_single_put_size=int(
                os.getenv(
                    "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE", str(8 * 1024 * 1024)
                )
            ),
            max_concurrency=int(os.getenv("AZURE_STORAGE_MAX_CONCURRENCY", "4")),
            connection_timeout=int(
                os.getenv("AZURE_STORAGE_CONNECTION_TIMEOUT", "20")
            ),
            read_timeout=int(os.getenv("AZURE_STORAGE_READ_TIMEOUT", "120")),
        )

    def retry_policy(self) -> ExponentialRetry:
        return ExponentialRetry(
            retry_total=self.retry_total,
            initial_backoff=self.retry_initial_backoff,
            increment_base=self.retry_increment_base,
        )

    def configure_logging(self) -> None:
        if not self.http_logging_enabled:
            return
        logging.basicConfig(
            level=_LOG_LEVELS[self.http_logging_level],
            format="%(asctime)s %(levelname)s %(name)s: %(message)s",
        )
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            _LOG_LEVELS[self.http_logging_level]
        )

    def client_options(self) -> dict[str, object]:
        return {
            "retry_policy": self.retry_policy(),
            "logging_enable": self.http_logging_enabled,
            "max_block_size": self.max_block_size,
            "max_single_put_size": self.max_single_put_size,
            "connection_timeout": self.connection_timeout,
            "read_timeout": self.read_timeout,
        }


def create_blob_service_client(
    settings: BlobStorageSettings,
) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    """Create a passwordless synchronous service client."""
    settings.configure_logging()
    credential = DefaultAzureCredential()
    client = BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        **settings.client_options(),
    )
    return client, credential


def create_async_blob_service_client(
    settings: BlobStorageSettings,
) -> tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:
    """Create a passwordless asynchronous service client."""
    settings.configure_logging()
    credential = AsyncDefaultAzureCredential()
    client = AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        **settings.client_options(),
    )
    return client, credential
