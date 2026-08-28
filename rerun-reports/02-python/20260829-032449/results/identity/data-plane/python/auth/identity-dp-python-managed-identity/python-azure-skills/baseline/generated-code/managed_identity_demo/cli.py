"""Command-line entry point for the managed identity examples."""

from __future__ import annotations

import argparse
import os
import sys
from typing import Optional, Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError

from .auth import IdentityType, create_credential
from .storage import create_blob_service_client, list_container_names


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Authenticate to Azure Blob Storage without storing credentials."
    )
    parser.add_argument(
        "--identity",
        choices=[item.value for item in IdentityType],
        required=True,
        help="'system' or 'user' on Azure; 'local' for an opt-in developer fallback",
    )
    parser.add_argument(
        "--account-url",
        default=os.getenv("AZURE_STORAGE_ACCOUNT_URL"),
        help="Blob endpoint; defaults to AZURE_STORAGE_ACCOUNT_URL",
    )
    parser.add_argument(
        "--client-id",
        help="Client ID of a user-assigned identity; defaults to AZURE_CLIENT_ID",
    )
    return parser


def run(argv: Optional[Sequence[str]] = None) -> int:
    args = build_parser().parse_args(argv)
    if not args.account_url:
        print(
            "Configuration error: pass --account-url or set AZURE_STORAGE_ACCOUNT_URL.",
            file=sys.stderr,
        )
        return 2

    try:
        credential = create_credential(
            IdentityType(args.identity),
            client_id=args.client_id,
        )
        client = create_blob_service_client(args.account_url, credential)
        names = list(list_container_names(client))
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        print(
            "Authentication failed. Confirm the managed identity is enabled and "
            "assigned to this host; for a user-assigned identity, verify its client "
            f"ID. Azure detail: {error}",
            file=sys.stderr,
        )
        return 3
    except HttpResponseError as error:
        print(
            "Azure rejected the request. Confirm the identity has a Blob data-plane "
            "role (for example, Storage Blob Data Reader) at the required scope. "
            f"Azure detail: {error}",
            file=sys.stderr,
        )
        return 4
    except ServiceRequestError as error:
        print(
            "Azure could not be reached. Check DNS, proxy, firewall, private endpoint, "
            f"and TLS settings. Azure detail: {error}",
            file=sys.stderr,
        )
        return 5

    if names:
        print("\n".join(names))
    else:
        print("Authenticated successfully; no containers were returned.")
    return 0


def main() -> None:
    raise SystemExit(run())


if __name__ == "__main__":
    main()
