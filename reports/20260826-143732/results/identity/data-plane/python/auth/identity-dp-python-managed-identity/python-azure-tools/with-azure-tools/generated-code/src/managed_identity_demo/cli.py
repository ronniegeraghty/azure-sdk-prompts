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

from .credentials import (
    CredentialConfigurationError,
    IdentityMode,
    create_credential,
)
from .storage import StorageConfigurationError, list_container_names

LOGGER = logging.getLogger(__name__)


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Authenticate an Azure Blob SDK client with managed identity."
    )
    parser.add_argument(
        "--identity",
        choices=[mode.value for mode in IdentityMode],
        default=os.getenv("IDENTITY_MODE", IdentityMode.AUTO.value),
        help="system, user, local, or auto (default: %(default)s)",
    )
    parser.add_argument(
        "--client-id",
        default=os.getenv("MANAGED_IDENTITY_CLIENT_ID"),
        help="Client ID of a user-assigned managed identity.",
    )
    parser.add_argument(
        "--account-url",
        default=os.getenv("AZURE_BLOB_ACCOUNT_URL"),
        help="Blob service URL, for example https://ACCOUNT.blob.core.windows.net.",
    )
    parser.add_argument(
        "--list-containers",
        action="store_true",
        help="Authenticate and list containers. Without this flag, no network call is made.",
    )
    parser.add_argument(
        "--debug",
        action="store_true",
        help="Enable Azure Identity debug logs; output can contain sensitive metadata.",
    )
    return parser


def _configure_logging(debug: bool) -> None:
    logging.basicConfig(
        level=logging.DEBUG if debug else logging.INFO,
        format="%(levelname)s %(name)s: %(message)s",
    )
    if debug:
        logging.getLogger("azure.identity").setLevel(logging.DEBUG)


def _print_dry_run(mode: IdentityMode, client_id: str | None, account_url: str | None) -> None:
    identity = mode.value
    if mode in (IdentityMode.USER, IdentityMode.AUTO) and client_id:
        identity += f" (user-assigned client ID ending in {client_id[-4:]})"
    print(f"Identity mode: {identity}")
    print(f"Blob account URL configured: {'yes' if account_url else 'no'}")
    print("Dry run complete; add --list-containers to contact Azure.")


def _troubleshooting_message(error: Exception) -> str:
    if isinstance(error, CredentialUnavailableError):
        return (
            "No selected credential is available. ManagedIdentityCredential only works on "
            "a supported Azure host with managed identity enabled; use --identity local "
            "when developing locally."
        )
    if isinstance(error, ClientAuthenticationError):
        return (
            "Microsoft Entra authentication failed. Confirm that the selected identity is "
            "enabled and, for a user-assigned identity, that its client ID is correct. "
            "Use --debug only in a secure terminal for credential-chain details."
        )
    if isinstance(error, ServiceRequestError):
        return (
            "The storage endpoint could not be reached. Check DNS, proxy, firewall, private "
            "endpoint routing, and AZURE_BLOB_ACCOUNT_URL."
        )
    if isinstance(error, HttpResponseError):
        if error.status_code == 403:
            return (
                "Azure Storage denied access. Assign the identity an appropriate data-plane "
                "role such as Storage Blob Data Reader at the narrowest required scope, then "
                "allow time for role assignment propagation."
            )
        return (
            f"Azure Storage returned HTTP {error.status_code or 'unknown'}. Check the account "
            "URL, service health, and the identity's data-plane permissions."
        )
    return str(error)


def main(argv: Sequence[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    _configure_logging(args.debug)
    mode = IdentityMode(args.identity)

    if not args.list_containers:
        if mode is IdentityMode.USER and not args.client_id:
            print(
                "Configuration error: user-assigned mode requires "
                "MANAGED_IDENTITY_CLIENT_ID or --client-id.",
                file=sys.stderr,
            )
            return 2
        _print_dry_run(mode, args.client_id, args.account_url)
        return 0

    if not args.account_url:
        print(
            "Configuration error: set AZURE_BLOB_ACCOUNT_URL or --account-url.",
            file=sys.stderr,
        )
        return 2

    try:
        with create_credential(mode, args.client_id) as credential:
            names = list_container_names(args.account_url, credential)
    except (CredentialConfigurationError, StorageConfigurationError) as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2
    except (
        CredentialUnavailableError,
        ClientAuthenticationError,
        ServiceRequestError,
        HttpResponseError,
    ) as error:
        LOGGER.debug("Azure SDK operation failed", exc_info=True)
        print(f"Azure operation failed: {_troubleshooting_message(error)}", file=sys.stderr)
        return 1

    if names:
        print("\n".join(names))
    else:
        print("No containers found.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
