"""Command-line demonstration of managed identity with Azure SDK clients."""

from __future__ import annotations

import argparse
import logging
import os
import sys
from collections.abc import Sequence

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import CredentialUnavailableError

from managed_identity_demo.auth import (
    AuthSettings,
    ConfigurationError,
    create_credential,
)
from managed_identity_demo.clients import create_clients

LOGGER = logging.getLogger("managed_identity_demo")
AZURE_MANAGEMENT_SCOPE = "https://management.azure.com/.default"


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Authenticate Azure SDK clients with managed identity."
    )
    parser.add_argument(
        "--check-auth",
        action="store_true",
        help="Request a token. Without this flag, no network request is made.",
    )
    parser.add_argument(
        "--list-resources",
        action="store_true",
        help="List Key Vault secret properties and Blob containers.",
    )
    parser.add_argument(
        "--verbose",
        action="store_true",
        help="Enable Azure Identity diagnostic logging. Logs can contain metadata.",
    )
    return parser


def run(argv: Sequence[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    logging.basicConfig(
        level=logging.DEBUG if args.verbose else logging.INFO,
        format="%(levelname)s: %(message)s",
    )

    try:
        settings = AuthSettings.from_environment()
        credential = create_credential(settings)
        clients = create_clients(
            credential,
            key_vault_url=os.getenv("AZURE_KEY_VAULT_URL"),
            storage_account_url=os.getenv("AZURE_STORAGE_ACCOUNT_URL"),
        )

        auth_description = (
            "local developer credential chain"
            if settings.environment == "local"
            else f"{settings.identity_type}-assigned managed identity"
        )
        LOGGER.info("Configured %s.", auth_description)
        LOGGER.info(
            "SDK clients: Key Vault=%s, Blob Storage=%s.",
            "configured" if clients.secret_client else "not configured",
            "configured" if clients.blob_service_client else "not configured",
        )

        if not args.check_auth and not args.list_resources:
            LOGGER.info("Dry run complete; no token or Azure resource was requested.")
            return 0

        if args.check_auth:
            token = credential.get_token(AZURE_MANAGEMENT_SCOPE)
            LOGGER.info(
                "Authentication succeeded; token expires at %s.", token.expires_on
            )

        if args.list_resources:
            if not clients.secret_client and not clients.blob_service_client:
                raise ConfigurationError(
                    "Set AZURE_KEY_VAULT_URL or AZURE_STORAGE_ACCOUNT_URL "
                    "before using --list-resources."
                )
            if clients.secret_client:
                names = [
                    item.name
                    for item in clients.secret_client.list_properties_of_secrets()
                ]
                LOGGER.info("Key Vault secrets: %s", names)
            if clients.blob_service_client:
                names = [
                    item.name
                    for item in clients.blob_service_client.list_containers()
                ]
                LOGGER.info("Blob containers: %s", names)
        return 0
    except ConfigurationError as error:
        LOGGER.error("Configuration error: %s", error)
        return 2
    except CredentialUnavailableError as error:
        LOGGER.error(
            "No credential is available: %s. In Azure, confirm managed identity "
            "is enabled. Locally, sign in with Azure CLI, Azure Developer CLI, "
            "Azure PowerShell, or VS Code.",
            error,
        )
        return 3
    except ClientAuthenticationError as error:
        LOGGER.error(
            "Azure rejected authentication: %s. Check the identity client ID, "
            "host identity assignment, tenant, and Conditional Access policies.",
            error,
        )
        return 4
    except HttpResponseError as error:
        LOGGER.error(
            "Azure service request failed (status %s): %s. Authentication may "
            "have succeeded; verify RBAC at the target resource scope.",
            error.status_code,
            error.message,
        )
        return 5


def main() -> None:
    sys.exit(run())


if __name__ == "__main__":
    main()
