"""Create an Azure SDK client with DefaultAzureCredential.

This example is offline-safe: it constructs the credential and client but does
not send a request. Authentication occurs when a service operation is invoked.
"""

import logging
import os

from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_identity_logging() -> None:
    if os.getenv("AZURE_IDENTITY_DEBUG", "").lower() not in {"1", "true", "yes"}:
        return

    handler = logging.StreamHandler()
    handler.setFormatter(
        logging.Formatter("%(asctime)s %(levelname)s %(name)s: %(message)s")
    )

    identity_logger = logging.getLogger("azure.identity")
    identity_logger.setLevel(logging.DEBUG)
    identity_logger.addHandler(handler)
    identity_logger.propagate = False


def main() -> None:
    configure_identity_logging()

    account_name = os.getenv("AZURE_STORAGE_ACCOUNT", "exampleaccount")
    account_url = f"https://{account_name}.blob.core.windows.net"

    with DefaultAzureCredential() as credential:
        with BlobServiceClient(
            account_url=account_url,
            credential=credential,
        ) as blob_service_client:
            container_client = blob_service_client.get_container_client("example")
            print(f"Configured Blob client for {container_client.url}")
            print("No network request was sent; tokens are acquired on first service call.")


if __name__ == "__main__":
    main()
