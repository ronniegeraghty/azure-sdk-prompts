"""Environment-based Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient, ExponentialRetry
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

_HTTP_LOGGER = "azure.core.pipeline.policies.http_logging_policy"


def _env_bool(name: str, default: bool) -> bool:
    value = os.getenv(name)
    if value is None:
        return default
    normalized = value.strip().lower()
    if normalized in {"1", "true", "yes", "on"}:
        return True
    if normalized in {"0", "false", "no", "off"}:
        return False
    raise ValueError(f"{name} must be true or false, not {value!r}.")


def _env_int(name: str, default: int, minimum: int = 0) -> int:
    value = os.getenv(name)
    parsed = default if value is None else int(value)
    if parsed < minimum:
        raise ValueError(f"{name} must be at least {minimum}.")
    return parsed


@dataclass(frozen=True)
class StorageSettings:
    """Settings used by both synchronous and asynchronous clients."""

    account_url: str
    max_retries: int = 5
    retry_delay: int = 2
    retry_increment: int = 2
    retry_jitter: int = 1
    http_logging_enabled: bool = False
    http_log_level: str = "WARNING"
    max_block_size: int = 4 * 1024 * 1024
    max_single_put_size: int = 64 * 1024 * 1024
    max_concurrency: int = 4
    connection_timeout: int = 20
    read_timeout: int = 120

    @classmethod
    def from_env(cls) -> "StorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required "
                "(for example, https://<account>.blob.core.windows.net)."
            )
        if not account_url.lower().startswith("https://"):
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS.")

        log_level = os.getenv("AZURE_STORAGE_HTTP_LOG_LEVEL", "WARNING").upper()
        if not isinstance(getattr(logging, log_level, None), int):
            raise ValueError(
                f"AZURE_STORAGE_HTTP_LOG_LEVEL is not a valid logging level: {log_level!r}."
            )

        return cls(
            account_url=account_url,
            max_retries=_env_int("AZURE_STORAGE_MAX_RETRIES", 5),
            retry_delay=_env_int("AZURE_STORAGE_RETRY_DELAY", 2),
            retry_increment=_env_int("AZURE_STORAGE_RETRY_INCREMENT", 2),
            retry_jitter=_env_int("AZURE_STORAGE_RETRY_JITTER", 1),
            http_logging_enabled=_env_bool(
                "AZURE_STORAGE_HTTP_LOGGING_ENABLED", False
            ),
            http_log_level=log_level,
            max_block_size=_env_int(
                "AZURE_STORAGE_MAX_BLOCK_SIZE", 4 * 1024 * 1024, 1024
            ),
            max_single_put_size=_env_int(
                "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE", 64 * 1024 * 1024, 1024
            ),
            max_concurrency=_env_int("AZURE_STORAGE_MAX_CONCURRENCY", 4, 1),
            connection_timeout=_env_int(
                "AZURE_STORAGE_CONNECTION_TIMEOUT", 20, 1
            ),
            read_timeout=_env_int("AZURE_STORAGE_READ_TIMEOUT", 120, 1),
        )

    def retry_policy(self) -> ExponentialRetry:
        return ExponentialRetry(
            initial_backoff=self.retry_delay,
            increment_base=self.retry_increment,
            retry_total=self.max_retries,
            random_jitter_range=self.retry_jitter,
        )

    def configure_http_logging(self) -> None:
        logging.getLogger(_HTTP_LOGGER).setLevel(self.http_log_level)

    def client_options(self) -> dict[str, object]:
        return {
            "retry_policy": self.retry_policy(),
            "logging_enable": self.http_logging_enabled,
            "max_block_size": self.max_block_size,
            "max_single_put_size": self.max_single_put_size,
            "connection_timeout": self.connection_timeout,
            "read_timeout": self.read_timeout,
        }


def create_sync_client(
    settings: StorageSettings, credential: DefaultAzureCredential
) -> BlobServiceClient:
    settings.configure_http_logging()
    return BlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        **settings.client_options(),
    )


def create_async_client(
    settings: StorageSettings, credential: AsyncDefaultAzureCredential
) -> AsyncBlobServiceClient:
    settings.configure_http_logging()
    return AsyncBlobServiceClient(
        account_url=settings.account_url,
        credential=credential,
        **settings.client_options(),
    )
