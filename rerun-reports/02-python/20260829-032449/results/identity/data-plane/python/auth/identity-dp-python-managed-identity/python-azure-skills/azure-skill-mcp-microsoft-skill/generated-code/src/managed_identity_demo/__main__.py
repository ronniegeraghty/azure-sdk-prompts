"""Command-line entry point for the managed identity examples."""

from __future__ import annotations

import argparse
import logging
import os
import sys
from typing import Optional, Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError

from .auth import IdentityMode, create_credential
from .storage import list_blob_containers

LOGGER = logging.getLogger(__name__)


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="List Blob containers using Microsoft Entra authentication."
    )
    parser.add_argument(
        "--identity",
        choices=("system", "user", "default"),
        default=os.getenv("IDENTITY_MODE", "default"),
        help=(
            "system: system-assigned managed identity; "
            "user: user-assigned managed identity; "
            "default: local-development/Azure fallback chain (default)"
        ),
    )
    parser.add_argument(
        "--client-id",
        default=os.getenv("AZURE_CLIENT_ID"),
        help="Client ID of a user-assigned managed identity.",
    )
    parser.add_argument(
        "--account-url",
        default=os.getenv("AZURE_STORAGE_ACCOUNT_URL"),
        help="Blob endpoint, for example https://myaccount.blob.core.windows.net.",
    )
    parser.add_argument(
        "--debug",
        action="store_true",
        help="Enable Azure Identity diagnostic logging. Logs may contain metadata.",
    )
    return parser


def run(argv: Optional[Sequence[str]] = None) -> int:
    args = build_parser().parse_args(argv)
    if not args.account_url:
        LOGGER.error(
            "Missing storage endpoint. Set AZURE_STORAGE_ACCOUNT_URL or pass --account-url."
        )
        return 2

    if args.debug:
        logging.getLogger("azure.identity").setLevel(logging.DEBUG)
        logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
            logging.DEBUG
        )

    try:
        credential = create_credential(
            args.identity,
            user_assigned_client_id=args.client_id,
        )
        with credential:
            containers = list_blob_containers(args.account_url, credential)
    except ValueError as error:
        LOGGER.error("%s", error)
        return 2
    except CredentialUnavailableError as error:
        LOGGER.error(
            "No usable credential was found. Confirm managed identity is enabled, "
            "or sign in locally with a supported developer credential. Details: %s",
            error,
        )
        return 3
    except ClientAuthenticationError as error:
        LOGGER.error(
            "Microsoft Entra authentication failed. Verify the selected identity and "
            "tenant configuration. Details: %s",
            error,
        )
        return 4
    except HttpResponseError as error:
        LOGGER.error(
            "Azure Storage rejected the request. Verify the account URL and that the "
            "identity has a data-plane role such as Storage Blob Data Reader. "
            "Status: %s; details: %s",
            error.status_code,
            error,
        )
        return 5
    except ServiceRequestError as error:
        LOGGER.error(
            "Could not reach Azure Storage. Verify DNS, private endpoint, firewall, "
            "proxy, and outbound network configuration. Details: %s",
            error,
        )
        return 6

    if containers:
        print("\n".join(containers))
    else:
        print("No blob containers found.")
    return 0


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(run())


if __name__ == "__main__":
    main()
