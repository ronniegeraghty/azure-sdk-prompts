"""Create an Azure Blob Storage client authenticated by DefaultAzureCredential."""

from __future__ import annotations

import argparse
import logging
import os
import sys

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_identity_logging() -> None:
    """Show which credential succeeds or why each credential is unavailable."""
    logger = logging.getLogger("azure.identity")
    logger.setLevel(logging.DEBUG)
    if not logger.handlers:
        handler = logging.StreamHandler(sys.stdout)
        handler.setFormatter(logging.Formatter("%(levelname)s %(name)s: %(message)s"))
        logger.addHandler(handler)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Demonstrate DefaultAzureCredential with Azure Blob Storage."
    )
    parser.add_argument(
        "--list-containers",
        action="store_true",
        help="Request a token and list containers. Without this flag, no network call is made.",
    )
    parser.add_argument(
        "--debug-auth",
        action="store_true",
        help="Enable detailed azure.identity credential-chain logging.",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    if args.debug_auth:
        configure_identity_logging()

    account_url = os.getenv(
        "AZURE_STORAGE_ACCOUNT_URL",
        "https://your-storage-account.blob.core.windows.net",
    )

    credential = DefaultAzureCredential()
    try:
        with BlobServiceClient(
            account_url=account_url,
            credential=credential,
        ) as blob_service_client:
            print(f"Created BlobServiceClient for {account_url}")

            if not args.list_containers:
                print(
                    "No network request was made. Add --list-containers to authenticate "
                    "and perform a read-only request."
                )
                return 0

            for container in blob_service_client.list_containers():
                print(container["name"])
    except ClientAuthenticationError as error:
        print(f"Authentication failed: {error}", file=sys.stderr)
        print(
            "Run again with --debug-auth to see which credentials were attempted.",
            file=sys.stderr,
        )
        return 2
    except HttpResponseError as error:
        print(
            f"Azure rejected the request ({error.status_code}): {error.message}",
            file=sys.stderr,
        )
        print(
            "Authentication may have succeeded; verify the account URL and Azure RBAC role.",
            file=sys.stderr,
        )
        return 3
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())

