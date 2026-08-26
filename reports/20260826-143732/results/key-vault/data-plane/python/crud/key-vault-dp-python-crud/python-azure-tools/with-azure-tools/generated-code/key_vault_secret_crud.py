"""Create, read, update, delete, and purge an Azure Key Vault secret.

Install dependencies with:
    python -m pip install -r requirements.txt

Set AZURE_KEY_VAULT_URL to the vault URL, authenticate with any identity
supported by DefaultAzureCredential, and run:
    python key_vault_secret_crud.py
"""

from __future__ import annotations

import argparse
import logging
import os
import re
from collections.abc import Callable
from typing import TypeVar
from urllib.parse import urlparse

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


INITIAL_VALUE = "initial-value"
UPDATED_VALUE = "updated-value"
SECRET_NAME_PATTERN = re.compile(r"^[0-9A-Za-z-]{1,127}$")
T = TypeVar("T")

logger = logging.getLogger(__name__)


class SecretOperationError(RuntimeError):
    """An Azure Key Vault secret operation failed."""


def execute(operation_name: str, operation: Callable[[], T]) -> T:
    """Run one Key Vault operation and add actionable error context."""
    try:
        return operation()
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        raise SecretOperationError(
            f"{operation_name} failed: DefaultAzureCredential could not authenticate. "
            "Configure a supported local credential or managed identity."
        ) from error
    except ResourceNotFoundError as error:
        raise SecretOperationError(
            f"{operation_name} failed: the secret or vault was not found."
        ) from error
    except HttpResponseError as error:
        if error.status_code == 403:
            detail = (
                "access denied; grant the identity secret get, set, delete, and purge "
                "permissions (for example, the Key Vault Secrets Officer RBAC role)"
            )
        elif error.status_code == 409:
            detail = (
                "the request conflicts with the vault state; a secret with this name "
                "may already be soft-deleted"
            )
        else:
            detail = f"Azure returned HTTP status {error.status_code or 'unknown'}"
        raise SecretOperationError(f"{operation_name} failed: {detail}.") from error


def validate_vault_url(vault_url: str) -> str:
    """Validate and normalize the configured HTTPS vault URL."""
    parsed = urlparse(vault_url)
    if parsed.scheme != "https" or not parsed.netloc or parsed.path not in ("", "/"):
        raise ValueError(
            "AZURE_KEY_VAULT_URL must be an HTTPS vault URL such as "
            "https://<vault-name>.vault.azure.net/"
        )
    return vault_url.rstrip("/") + "/"


def perform_crud(vault_url: str, secret_name: str) -> None:
    """Perform the complete secret lifecycle, including permanent purge."""
    with DefaultAzureCredential() as credential:
        with SecretClient(vault_url=vault_url, credential=credential) as client:
            created = execute(
                "Create secret",
                lambda: client.set_secret(secret_name, INITIAL_VALUE),
            )
            logger.info(
                "Created secret %s (version %s).",
                created.name,
                created.properties.version,
            )

            retrieved = execute(
                "Read secret",
                lambda: client.get_secret(secret_name),
            )
            if retrieved.value != INITIAL_VALUE:
                raise SecretOperationError(
                    "Read secret failed: the retrieved value did not match the created value."
                )
            logger.info("Read and verified secret %s.", retrieved.name)

            updated = execute(
                "Update secret",
                lambda: client.set_secret(secret_name, UPDATED_VALUE),
            )
            logger.info(
                "Updated secret %s to a new version (%s).",
                updated.name,
                updated.properties.version,
            )

            updated_secret = execute(
                "Verify updated secret",
                lambda: client.get_secret(secret_name),
            )
            if updated_secret.value != UPDATED_VALUE:
                raise SecretOperationError(
                    "Update secret failed: the retrieved value was not 'updated-value'."
                )
            logger.info("Verified the updated value for secret %s.", secret_name)

            execute(
                "Delete secret",
                lambda: client.begin_delete_secret(secret_name).result(),
            )
            logger.info("Soft-deleted secret %s.", secret_name)

            execute(
                "Purge secret",
                lambda: client.purge_deleted_secret(secret_name),
            )
            logger.info("Permanently purged secret %s.", secret_name)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run all CRUD operations on an Azure Key Vault secret."
    )
    parser.add_argument(
        "--vault-url",
        default=os.getenv("AZURE_KEY_VAULT_URL"),
        help="Key Vault URL; defaults to AZURE_KEY_VAULT_URL.",
    )
    parser.add_argument(
        "--secret-name",
        default="crud-demo-secret",
        help="Secret name to use (default: crud-demo-secret).",
    )
    return parser.parse_args()


def main() -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    args = parse_args()

    try:
        if not args.vault_url:
            raise ValueError(
                "Set AZURE_KEY_VAULT_URL or provide the --vault-url argument."
            )
        vault_url = validate_vault_url(args.vault_url)
        if not SECRET_NAME_PATTERN.fullmatch(args.secret_name):
            raise ValueError(
                "The secret name must contain 1-127 letters, numbers, or hyphens."
            )
        perform_crud(vault_url, args.secret_name)
    except (ValueError, SecretOperationError) as error:
        logger.error("%s", error)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
