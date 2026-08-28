"""Configuration and authenticated client factories for Azure Blob Storage."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass

from azure.core.pipeline.policies import RetryPolicy
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _positive_int(name: str, default: int) -> int:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer, got {raw_value!r}") from exc
    if value < 1:
        raise ValueError(f"{name} must be at least 1")
    return value


def _non_negative_float(name: str, default: float) -> float:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = float(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be a number, got {raw_value!r}") from exc
    if value < 0:
        raise ValueError(f"{name} cannot be negative")
    return value


@dataclass(frozen=True, slots=True)
class BlobStorageSettings:
    account_url: str
    container_name: str = "blob-manager-demo"
    max_retries: int = 5
    retry_delay: float = 1.0
    retry_max_delay: float = 30.0
    connection_timeout: int = 20
    read_timeout: int = 120
    max_concurrency: int = 4
    block_size: int = 8 * 1024 * 1024
    logging_level: str = "WARNING"

    @classmethod
    def from_env(cls) -> "BlobStorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required, for example "
                "https://<account>.blob.core.windows.net"
            )
        if not account_url.startswith("https://"):
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS")

        logging_level = os.getenv("AZURE_STORAGE_LOG_LEVEL", "WARNING").upper()
        if logging_level not in logging.getLevelNamesMapping():
            raise ValueError(f"Invalid AZURE_STORAGE_LOG_LEVEL: {logging_level!r}")

        container_name = os.getenv(
            "AZURE_STORAGE_CONTAINER", "blob-manager-demo"
        ).strip()
        if not container_name:
            raise ValueError("AZURE_STORAGE_CONTAINER cannot be empty")

        return cls(
            account_url=account_url,
            container_name=container_name,
            max_retries=_positive_int("AZURE_STORAGE_MAX_RETRIES", 5),
            retry_delay=_non_negative_float("AZURE_STORAGE_RETRY_DELAY", 1.0),
            retry_max_delay=_non_negative_float(
                "AZURE_STORAGE_RETRY_MAX_DELAY", 30.0
            ),
            connection_timeout=_positive_int(
                "AZURE_STORAGE_CONNECTION_TIMEOUT", 20
            ),
            read_timeout=_positive_int("AZURE_STORAGE_READ_TIMEOUT", 120),
            max_concurrency=_positive_int("AZURE_STORAGE_MAX_CONCURRENCY", 4),
            block_size=_positive_int(
                "AZURE_STORAGE_BLOCK_SIZE", 8 * 1024 * 1024
            ),
            logging_level=logging_level,
        )

    def configure_logging(self) -> None:
        level = logging.getLevelNamesMapping()[self.logging_level]
        logging.basicConfig(
            level=level,
            format="%(asctime)s %(levelname)s %(name)s: %(message)s",
        )
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            level
        )

    def retry_policy(self) -> RetryPolicy:
        return RetryPolicy(
            retry_total=self.max_retries,
            retry_connect=self.max_retries,
            retry_read=self.max_retries,
            retry_status=self.max_retries,
            retry_backoff_factor=self.retry_delay,
            retry_backoff_max=self.retry_max_delay,
            retry_mode="exponential",
        )


def create_sync_client(
    settings: BlobStorageSettings,
) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    credential = DefaultAzureCredential()
    client = BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=settings.retry_policy(),
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
        logging_enable=True,
        max_block_size=settings.block_size,
        max_single_put_size=settings.block_size,
    )
    return client, credential


def create_async_client(
    settings: BlobStorageSettings,
) -> tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:
    credential = AsyncDefaultAzureCredential()
    client = AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        retry_policy=settings.retry_policy(),
        connection_timeout=settings.connection_timeout,
        read_timeout=settings.read_timeout,
        logging_enable=True,
        max_block_size=settings.block_size,
        max_single_put_size=settings.block_size,
    )
    return client, credential
