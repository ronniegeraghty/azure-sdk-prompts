import argparse
import logging
import os
import sys

from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_identity_logging() -> None:
    handler = logging.StreamHandler(sys.stderr)
    handler.setFormatter(logging.Formatter("%(levelname)s %(name)s: %(message)s"))

    identity_logger = logging.getLogger("azure.identity")
    identity_logger.setLevel(logging.DEBUG)
    identity_logger.addHandler(handler)
    identity_logger.propagate = False


def create_blob_client(account_url: str) -> tuple[BlobServiceClient, DefaultAzureCredential]:
    credential = DefaultAzureCredential()
    client = BlobServiceClient(account_url=account_url, credential=credential)
    return client, credential


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Construct an Azure Blob SDK client with DefaultAzureCredential."
    )
    parser.add_argument(
        "--debug-auth",
        action="store_true",
        help="Show which credentials DefaultAzureCredential attempts.",
    )
    args = parser.parse_args()

    if args.debug_auth:
        configure_identity_logging()

    account_url = os.environ.get(
        "AZURE_STORAGE_ACCOUNT_URL",
        "https://example.blob.core.windows.net",
    )
    client, credential = create_blob_client(account_url)

    try:
        print(f"BlobServiceClient configured for {client.url}")
        print("No network request was made; authentication starts on the first SDK operation.")
    finally:
        client.close()
        credential.close()


if __name__ == "__main__":
    main()
