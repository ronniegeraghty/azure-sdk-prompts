import argparse
import logging
import os
import sys

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_logging() -> None:
    level_name = os.getenv("AZURE_SDK_LOG_LEVEL", "WARNING").upper()
    level = getattr(logging, level_name, None)
    if not isinstance(level, int):
        raise ValueError(
            "AZURE_SDK_LOG_LEVEL must be a Python logging level such as DEBUG, "
            "INFO, WARNING, or ERROR"
        )

    logging.basicConfig(
        level=logging.WARNING,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    logging.getLogger("azure.identity").setLevel(level)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create a BlobServiceClient with DefaultAzureCredential."
    )
    parser.add_argument(
        "--list-containers",
        action="store_true",
        help="Make an authenticated, read-only request to list blob containers.",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    configure_logging()

    account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        print(
            "Set AZURE_STORAGE_ACCOUNT_URL to "
            "https://<account>.blob.core.windows.net.",
            file=sys.stderr,
        )
        return 2

    try:
        with DefaultAzureCredential() as credential:
            with BlobServiceClient(
                account_url=account_url,
                credential=credential,
            ) as client:
                print(
                    f"Created BlobServiceClient for account '{client.account_name}'."
                )

                if not args.list_containers:
                    print(
                        "No request was sent. Add --list-containers to test "
                        "authentication and authorization."
                    )
                    return 0

                for container in client.list_containers():
                    print(container["name"])
    except CredentialUnavailableError as error:
        logging.getLogger(__name__).error(
            "No credential in the chain was available: %s", error
        )
        return 1
    except ClientAuthenticationError as error:
        logging.getLogger(__name__).error("Authentication failed: %s", error.message)
        return 1
    except HttpResponseError as error:
        logging.getLogger(__name__).error(
            "Azure Storage rejected the request: %s", error.message
        )
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())

