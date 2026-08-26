"""Environment-driven Azure Blob Storage client configuration."""

from __future__ import annotations

import logging
import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _positive_int(name: str, default: int) -> int:
    raw_value = os.getenv(name, str(default))
    try:
        value = int(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be an integer, got {raw_value!r}") from exc
    if value < 0:
        raise ValueError(f"{name} must be zero or greater")
    return value


def _strictly_positive_int(name: str, default: int) -> int:
    value = _positive_int(name, default)
    if value == 0:
        raise ValueError(f"{name} must be greater than zero")
    return value


def _positive_float(name: str, default: float) -> float:
    raw_value = os.getenv(name, str(default))
    try:
        value = float(raw_value)
    except ValueError as exc:
        raise ValueError(f"{name} must be a number, got {raw_value!r}") from exc
    if value < 0:
        raise ValueError(f"{name} must be zero or greater")
    return value


def _boolean(name: str, default: bool) -> bool:
    raw_value = os.getenv(name)
    if raw_value is None:
        return default
    normalized = raw_value.strip().lower()
    if normalized in {"1", "true", "yes", "on"}:
        return True
    if normalized in {"0", "false", "no", "off"}:
        return False
    raise ValueError(f"{name} must be true or false, got {raw_value!r}")


@dataclass(frozen=True, slots=True)
class BlobStorageSettings:
    """Configuration used by both sync and async Blob service clients."""

    account_url: str
    max_retries: int = 5
    retry_delay: float = 1.0
    retry_max_delay: float = 30.0
    http_logging_enabled: bool = False
    http_log_level: str = "WARNING"
    max_block_size: int = 4 * 1024 * 1024
    max_single_put_size: int = 8 * 1024 * 1024
    max_concurrency: int = 4

    @classmethod
    def from_env(cls) -> "BlobStorageSettings":
        account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL", "").strip().rstrip("/")
        if not account_url:
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_URL is required "
                "(for example, https://<account>.blob.core.windows.net)"
            )
        if not account_url.startswith("https://"):
            raise ValueError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS")

        log_level = os.getenv("AZURE_HTTP_LOG_LEVEL", "WARNING").strip().upper()
        if log_level not in logging.getLevelNamesMapping():
            raise ValueError(f"AZURE_HTTP_LOG_LEVEL is invalid: {log_level!r}")

        return cls(
            account_url=account_url,
            max_retries=_positive_int("AZURE_STORAGE_MAX_RETRIES", 5),
            retry_delay=_positive_float("AZURE_STORAGE_RETRY_DELAY", 1.0),
            retry_max_delay=_positive_float("AZURE_STORAGE_RETRY_MAX_DELAY", 30.0),
            http_logging_enabled=_boolean("AZURE_HTTP_LOGGING_ENABLED", False),
            http_log_level=log_level,
            max_block_size=_strictly_positive_int(
                "AZURE_STORAGE_MAX_BLOCK_SIZE", 4 * 1024 * 1024
            ),
            max_single_put_size=_strictly_positive_int(
                "AZURE_STORAGE_MAX_SINGLE_PUT_SIZE", 8 * 1024 * 1024
            ),
            max_concurrency=_strictly_positive_int(
                "AZURE_STORAGE_MAX_CONCURRENCY", 4
            ),
        )

    def configure_logging(self) -> None:
        if self.http_logging_enabled:
            logging.basicConfig(level=self.http_log_level)
            logging.getLogger(
                "azure.core.pipeline.policies.http_logging_policy"
            ).setLevel(self.http_log_level)

    def _client_options(self) -> dict[str, int | float | bool]:
        return {
            "retry_total": self.max_retries,
            "retry_connect": self.max_retries,
            "retry_read": self.max_retries,
            "retry_status": self.max_retries,
            "retry_backoff_factor": self.retry_delay,
            "retry_backoff_max": self.retry_max_delay,
            "logging_enable": self.http_logging_enabled,
            "max_block_size": self.max_block_size,
            "max_single_put_size": self.max_single_put_size,
        }

    def create_sync_client(
        self,
    ) -> tuple[DefaultAzureCredential, BlobServiceClient]:
        self.configure_logging()
        credential = DefaultAzureCredential()
        client = BlobServiceClient(
            account_url=self.account_url,
            credential=credential,
            **self._client_options(),
        )
        return credential, client

    def create_async_client(
        self,
    ) -> tuple[AsyncDefaultAzureCredential, AsyncBlobServiceClient]:
        self.configure_logging()
        credential = AsyncDefaultAzureCredential()
        client = AsyncBlobServiceClient(
            account_url=self.account_url,
            credential=credential,
            **self._client_options(),
        )
        return credential, client
