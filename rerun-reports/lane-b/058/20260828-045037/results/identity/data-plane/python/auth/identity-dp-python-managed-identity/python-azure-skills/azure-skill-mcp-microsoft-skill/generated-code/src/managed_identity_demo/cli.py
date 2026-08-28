"""Command-line entry point for the managed identity examples."""

import argparse
import logging
import os
from typing import Optional, Sequence

from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError

from .credentials import CredentialMode, create_credential
from .storage import AzureOperationError, build_blob_service_client, list_container_names


def _account_url(explicit_url: Optional[str]) -> str:
    if explicit_url:
        return explicit_url
    if value := os.getenv("AZURE_STORAGE_ACCOUNT_URL"):
        return value
    if name := os.getenv("AZURE_STORAGE_ACCOUNT_NAME"):
        return f"https://{name}.blob.core.windows.net"
    raise ValueError(
        "Set AZURE_STORAGE_ACCOUNT_URL or AZURE_STORAGE_ACCOUNT_NAME, "
        "or pass --account-url."
    )


def _parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Authenticate an Azure Blob SDK client with managed identity."
    )
    parser.add_argument(
        "command",
        nargs="?",
        choices=("inspect", "list-containers"),
        default="inspect",
        help="'inspect' is offline-safe; 'list-containers' contacts Azure.",
    )
    parser.add_argument(
        "--mode",
        choices=[mode.value for mode in CredentialMode],
        default=CredentialMode.LOCAL.value,
        help="Strict managed identity, local-only, or local/production auto-detection.",
    )
    parser.add_argument("--account-url")
    parser.add_argument(
        "--client-id",
        default=os.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID"),
        help="Client ID of a user-assigned managed identity.",
    )
    parser.add_argument(
        "--allow-interactive-browser",
        action="store_true",
        help="Allow browser login as the last local development fallback.",
    )
    parser.add_argument(
        "--debug",
        action="store_true",
        help="Enable Azure Identity diagnostic logging. Logs can contain metadata.",
    )
    return parser


def run(argv: Optional[Sequence[str]] = None) -> int:
    args = _parser().parse_args(argv)
    if args.debug:
        logging.basicConfig(level=logging.DEBUG)
        logging.getLogger("azure.identity").setLevel(logging.DEBUG)

    credential = None
    try:
        account_url = _account_url(args.account_url)
        mode = CredentialMode(args.mode)
        credential = create_credential(
            mode,
            client_id=args.client_id,
            allow_interactive_browser=args.allow_interactive_browser,
        )

        if args.command == "inspect":
            with build_blob_service_client(account_url, credential):
                print(f"Configured BlobServiceClient for {account_url}")
                print(f"Credential mode: {mode.value}")
                print("No token or network request was made.")
            return 0

        names = list_container_names(account_url, credential)
        if names:
            for name in names:
                print(name)
        else:
            print("No containers were returned.")
        return 0
    except CredentialUnavailableError as exc:
        logging.error(
            "No credential is available. On Azure, enable/attach managed identity. "
            "Locally, sign in with a supported developer tool. Details: %s",
            exc,
        )
    except ClientAuthenticationError as exc:
        logging.error(
            "Credential authentication failed. Verify tenant selection, identity "
            "attachment, and token audience. Details: %s",
            exc,
        )
    except (AzureOperationError, ValueError) as exc:
        logging.error("%s", exc)
    finally:
        if credential is not None:
            credential.close()
    return 2


def main() -> None:
    raise SystemExit(run())
