"""Configuration and authenticated Azure Blob Storage client factories."""

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


class ConfigurationError(ValueError):
    """Raised when required storage configuration is missing or invalid."""


def _integer_from_env(name: str, default: int, minimum: int = 0) -> int:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ConfigurationError(f"{name} must be an integer.") from exc
    if value < minimum:
        raise ConfigurationError(f"{name} must be at least {minimum}.")
    return value


def _float_from_env(name: str, default: float, minimum: float = 0.0) -> float:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    try:
        value = float(raw_value)
    except ValueError as exc:
        raise ConfigurationError(f"{name} must be a number.") from exc
    if value < minimum:
        raise ConfigurationError(f"{name} must be at least {minimum}.")
    return value


@dataclass(frozen=True, slots=True)
class StorageSettings:
    """Settings shared by synchronous and asynchronous storage clients."""

    account_url: str
    max_retries: int = 5
    retry_delay: float = 1.0
    max_retry_delay: float = 30.0
    log_level: str = "INFO"
    max_concurrency: int = 4
    block_size: int = 8 * 1024 * 1024
    single_upload_threshold: int = 64 * 1024 * 1024
    connection_timeout: int = 20
    read_timeout: int = 60

    @classmethod
    def from_env(cls) -> "StorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ConfigurationError(
                "AZURE_STORAGE_ACCOUNT_URL is required "
                "(for example, https://<account>.blob.core.windows.net)."
            )

        parsed_url = urlparse(account_url)
        if parsed_url.scheme != "https" or not parsed_url.netloc:
            raise ConfigurationError(
                "AZURE_STORAGE_ACCOUNT_URL must be a valid HTTPS endpoint."
            )

        log_level = os.getenv("AZURE_STORAGE_LOG_LEVEL", "INFO").upper()
        if log_level not in logging.getLevelNamesMapping():
            raise ConfigurationError(
                "AZURE_STORAGE_LOG_LEVEL must be a standard Python logging level."
            )

        return cls(
            account_url=account_url,
            max_retries=_integer_from_env("AZURE_STORAGE_MAX_RETRIES", 5),
            retry_delay=_float_from_env("AZURE_STORAGE_RETRY_DELAY", 1.0),
            max_retry_delay=_float_from_env(
                "AZURE_STORAGE_MAX_RETRY_DELAY", 30.0
            ),
            log_level=log_level,
            max_concurrency=_integer_from_env(
                "AZURE_STORAGE_MAX_CONCURRENCY", 4, minimum=1
            ),
            block_size=_integer_from_env(
                "AZURE_STORAGE_BLOCK_SIZE", 8 * 1024 * 1024, minimum=1024
            ),
            single_upload_threshold=_integer_from_env(
                "AZURE_STORAGE_SINGLE_UPLOAD_THRESHOLD",
                64 * 1024 * 1024,
                minimum=1024,
            ),
            connection_timeout=_integer_from_env(
                "AZURE_STORAGE_CONNECTION_TIMEOUT", 20, minimum=1
            ),
            read_timeout=_integer_from_env(
                "AZURE_STORAGE_READ_TIMEOUT", 60, minimum=1
            ),
        )

    def configure_logging(self) -> None:
        level = logging.getLevelNamesMapping()[self.log_level]
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
            retry_backoff_factor=self.retry_delay,
            retry_backoff_max=self.max_retry_delay,
            retry_mode="exponential",
        )

    def async_retry_policy(self) -> AsyncRetryPolicy:
        return AsyncRetryPolicy(
            retry_total=self.max_retries,
            retry_backoff_factor=self.retry_delay,
            retry_backoff_max=self.max_retry_delay,
            retry_mode="exponential",
        )

    def create_credential(self) -> DefaultAzureCredential:
        return DefaultAzureCredential()

    def create_async_credential(self) -> AsyncDefaultAzureCredential:
        return AsyncDefaultAzureCredential()

    def create_client(
        self, credential: DefaultAzureCredential
    ) -> BlobServiceClient:
        return BlobServiceClient(
            account_url=self.account_url,
            credential=credential,
            retry_policy=self.retry_policy(),
            logging_enable=True,
            connection_timeout=self.connection_timeout,
            read_timeout=self.read_timeout,
            max_block_size=self.block_size,
            max_single_put_size=self.single_upload_threshold,
        )

    def create_async_client(
        self, credential: AsyncDefaultAzureCredential
    ) -> AsyncBlobServiceClient:
        return AsyncBlobServiceClient(
            account_url=self.account_url,
            credential=credential,
            retry_policy=self.async_retry_policy(),
            logging_enable=True,
            connection_timeout=self.connection_timeout,
            read_timeout=self.read_timeout,
            max_block_size=self.block_size,
            max_single_put_size=self.single_upload_threshold,
        )
