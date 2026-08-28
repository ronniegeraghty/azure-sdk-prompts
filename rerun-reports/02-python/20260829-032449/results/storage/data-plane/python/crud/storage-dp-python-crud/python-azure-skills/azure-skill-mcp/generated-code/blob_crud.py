import logging
import os
from pathlib import Path

from azure.core.exceptions import HttpResponseError, ResourceExistsError
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


CONTAINER_NAME = "my-container"
BLOB_NAME = "reports/report.csv"
SOURCE_FILE = Path("report.csv")
DOWNLOAD_FILE = Path("report-downloaded.csv")


def run_blob_crud() -> None:
    account_url = os.environ.get("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        raise ValueError(
            "AZURE_STORAGE_ACCOUNT_URL must be set, for example "
            "https://<storage-account>.blob.core.windows.net"
        )

    credential = DefaultAzureCredential()
    blob_service_client = BlobServiceClient(
        account_url=account_url,
        credential=credential,
    )

    try:
        container_client = blob_service_client.get_container_client(CONTAINER_NAME)
        try:
            container_client.create_container()
            logging.info("Created container %s.", CONTAINER_NAME)
        except ResourceExistsError:
            logging.info("Container %s already exists.", CONTAINER_NAME)

        blob_client = container_client.get_blob_client(BLOB_NAME)
        with SOURCE_FILE.open("rb") as source:
            blob_client.upload_blob(source, overwrite=True)
        logging.info("Uploaded %s as %s.", SOURCE_FILE, BLOB_NAME)

        for blob in container_client.list_blobs():
            print(f"{blob.name}: {blob.size} bytes")

        with DOWNLOAD_FILE.open("wb") as destination:
            blob_client.download_blob().readinto(destination)
        logging.info("Downloaded %s to %s.", BLOB_NAME, DOWNLOAD_FILE)

        blob_client.delete_blob()
        logging.info("Deleted blob %s.", BLOB_NAME)

        container_client.delete_container()
        logging.info("Deleted container %s.", CONTAINER_NAME)
    finally:
        blob_service_client.close()
        credential.close()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    try:
        run_blob_crud()
    except (OSError, ValueError) as error:
        logging.error("%s", error)
        raise SystemExit(1) from error
    except HttpResponseError as error:
        logging.error(
            "Azure request failed (status %s): %s",
            error.status_code,
            error.message,
        )
        raise SystemExit(1) from error
