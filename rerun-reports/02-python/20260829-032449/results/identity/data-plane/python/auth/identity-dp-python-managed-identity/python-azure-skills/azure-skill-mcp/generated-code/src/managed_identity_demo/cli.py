"""Command-line entry point for managed identity examples."""

from __future__ import annotations

import argparse
import logging
import os
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError

from managed_identity_demo.auth import AuthMode, create_credential
from managed_identity_demo.storage import (
    create_blob_service_client,
    list_container_names,
)

LOGGER = logging.getLogger("managed_identity_demo")
PLACEHOLDER_ACCOUNT_URL = "https://example.blob.core.windows.net"


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Authenticate an Azure Blob client with managed identity."
    )
    parser.add_argument(
        "--auth",
        type=AuthMode,
        choices=[mode.value for mode in AuthMode],
        default=AuthMode.SYSTEM,
        help=(
            "system/user for Azure-hosted production; "
            "local-default/local-cli for development (default: system)"
        ),
    )
    parser.add_argument(
        "--account-url",
        default=os.getenv("AZURE_STORAGE_ACCOUNT_URL", PLACEHOLDER_ACCOUNT_URL),
        help="Storage endpoint; defaults to AZURE_STORAGE_ACCOUNT_URL.",
    )
    parser.add_argument(
        "--client-id",
        default=os.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID"),
        help="User-assigned managed identity client ID.",
    )
    parser.add_argument(
        "--execute",
        action="store_true",
        help="Make the Azure request. Without this flag, only construct clients.",
    )
    parser.add_argument(
        "--verbose",
        action="store_true",
        help="Enable Azure Identity diagnostic logging (tokens are not logged).",
    )
    return parser


def configure_logging(verbose: bool) -> None:
    level = logging.DEBUG if verbose else logging.INFO
    logging.basicConfig(level=level, format="%(levelname)s %(name)s: %(message)s")
    if verbose:
        logging.getLogger("azure.identity").setLevel(logging.DEBUG)
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            logging.WARNING
        )


def run(argv: Sequence[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    configure_logging(args.verbose)

    try:
        credential = create_credential(
            args.auth,
            managed_identity_client_id=args.client_id,
        )
        client = create_blob_service_client(args.account_url, credential)
    except ValueError as error:
        LOGGER.error("Configuration error: %s", error)
        return 2

    try:
        if not args.execute:
            LOGGER.info(
                "Dry run complete: created %s credential and BlobServiceClient.",
                args.auth.value,
            )
            return 0

        if args.account_url == PLACEHOLDER_ACCOUNT_URL:
            LOGGER.error(
                "Set AZURE_STORAGE_ACCOUNT_URL or --account-url before using --execute."
            )
            return 2

        names = list(list_container_names(client))
        if names:
            for name in names:
                print(name)
        else:
            LOGGER.info("Authentication succeeded; no containers were returned.")
        return 0
    except CredentialUnavailableError as error:
        LOGGER.error(
            "The selected credential is unavailable: %s. "
            "For managed identity, verify the identity is attached to this Azure host. "
            "For local use, sign in with the selected developer tool.",
            error,
        )
        return 3
    except ClientAuthenticationError as error:
        LOGGER.error(
            "Microsoft Entra authentication failed: %s. "
            "Check the user-assigned client ID, tenant context, and identity endpoint.",
            error,
        )
        return 3
    except HttpResponseError as error:
        status = error.status_code or "unknown"
        LOGGER.error(
            "Azure Storage rejected the request (HTTP %s): %s. "
            "Verify the endpoint and assign a Blob data-plane role such as "
            "Storage Blob Data Reader at the narrowest required scope.",
            status,
            error.message,
        )
        return 4
    except ServiceRequestError as error:
        LOGGER.error(
            "Could not reach Azure Storage: %s. Check DNS, proxy, firewall, "
            "private endpoint routing, and the account URL.",
            error,
        )
        return 5
    finally:
        client.close()
        close_credential = getattr(credential, "close", None)
        if close_credential is not None:
            close_credential()


def main() -> None:
    sys.exit(run())


if __name__ == "__main__":
    main()
