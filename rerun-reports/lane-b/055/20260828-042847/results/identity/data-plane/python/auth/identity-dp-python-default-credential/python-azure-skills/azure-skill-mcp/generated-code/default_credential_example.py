from __future__ import annotations

import logging
import os
import sys

from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_identity_logging() -> None:
    level_name = os.getenv("AZURE_IDENTITY_LOG_LEVEL", "WARNING").upper()
    level = getattr(logging, level_name, None)
    if not isinstance(level, int):
        raise ValueError(
            "AZURE_IDENTITY_LOG_LEVEL must be DEBUG, INFO, WARNING, ERROR, or CRITICAL"
        )

    handler = logging.StreamHandler(sys.stderr)
    handler.setFormatter(logging.Formatter("%(levelname)s %(name)s: %(message)s"))

    logger = logging.getLogger("azure.identity")
    logger.handlers.clear()
    logger.addHandler(handler)
    logger.setLevel(level)
    logger.propagate = False


def create_blob_client() -> tuple[BlobServiceClient, DefaultAzureCredential]:
    account_url = os.getenv(
        "AZURE_STORAGE_ACCOUNT_URL",
        "https://example.blob.core.windows.net",
    )
    credential = DefaultAzureCredential()
    client = BlobServiceClient(account_url=account_url, credential=credential)
    return client, credential


def main() -> None:
    configure_identity_logging()
    client, credential = create_blob_client()

    try:
        print(f"BlobServiceClient created for {client.url}")
        print("No network request was made, so no access token was requested.")

        if os.getenv("AZURE_RUN_LIVE_REQUEST") == "1":
            account_info = client.get_account_information()
            print(f"Storage account kind: {account_info.get('account_kind', 'unknown')}")
    finally:
        client.close()
        credential.close()


if __name__ == "__main__":
    main()
