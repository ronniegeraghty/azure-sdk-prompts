import logging
import os
import sys
from pathlib import Path

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
)
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


CONTAINER_NAME = "my-container"
BLOB_NAME = "reports/report.csv"
SOURCE_FILE = Path("report.csv")
DOWNLOAD_FILE = Path("report-downloaded.csv")

logger = logging.getLogger(__name__)


def run_blob_crud() -> int:
    account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        logger.error("Set AZURE_STORAGE_ACCOUNT_URL to the storage account blob endpoint.")
        return 1

    if not SOURCE_FILE.is_file():
        logger.error("Source file does not exist: %s", SOURCE_FILE)
        return 1

    try:
        with DefaultAzureCredential() as credential:
            with BlobServiceClient(
                account_url=account_url,
                credential=credential,
            ) as blob_service_client:
                container_client = blob_service_client.get_container_client(
                    CONTAINER_NAME
                )

                try:
                    container_client.create_container()
                    logger.info("Created container %s.", CONTAINER_NAME)
                except ResourceExistsError:
                    logger.info("Container %s already exists.", CONTAINER_NAME)

                blob_client = container_client.get_blob_client(BLOB_NAME)
                with SOURCE_FILE.open("rb") as source:
                    blob_client.upload_blob(source, overwrite=True)
                logger.info("Uploaded %s as %s.", SOURCE_FILE, BLOB_NAME)

                for blob in container_client.list_blobs():
                    print(f"{blob.name}: {blob.size} bytes")

                with DOWNLOAD_FILE.open("wb") as destination:
                    blob_client.download_blob().readinto(destination)
                logger.info("Downloaded %s to %s.", BLOB_NAME, DOWNLOAD_FILE)

                blob_client.delete_blob()
                logger.info("Deleted blob %s.", BLOB_NAME)

                container_client.delete_container()
                logger.info("Deleted container %s.", CONTAINER_NAME)
    except ClientAuthenticationError as error:
        logger.error("Azure authentication failed: %s", error)
        return 1
    except HttpResponseError as error:
        logger.error("Azure Blob Storage request failed: %s", error)
        return 1
    except OSError as error:
        logger.error("Local file operation failed: %s", error)
        return 1

    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(run_blob_crud())
