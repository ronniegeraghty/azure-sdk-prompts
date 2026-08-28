import os
import sys
from pathlib import Path

from azure.core.exceptions import HttpResponseError, ResourceExistsError
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


CONTAINER_NAME = "my-container"
BLOB_NAME = "reports/report.csv"
SOURCE_FILE = Path("report.csv")
DOWNLOAD_FILE = Path("report-downloaded.csv")


def create_blob_service_client() -> BlobServiceClient:
    account_url = os.environ.get("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        raise ValueError(
            "Set AZURE_STORAGE_ACCOUNT_URL to the storage account Blob service URL "
            "(for example, https://<account-name>.blob.core.windows.net)."
        )

    return BlobServiceClient(
        account_url=account_url,
        credential=DefaultAzureCredential(),
    )


def run_blob_crud() -> None:
    if not SOURCE_FILE.is_file():
        raise FileNotFoundError(f"Local upload file not found: {SOURCE_FILE}")

    service_client = create_blob_service_client()
    container_client = service_client.get_container_client(CONTAINER_NAME)

    try:
        container_client.create_container()
        print(f"Created container: {CONTAINER_NAME}")
    except ResourceExistsError:
        print(f"Container already exists: {CONTAINER_NAME}")

    blob_client = container_client.get_blob_client(BLOB_NAME)

    with SOURCE_FILE.open("rb") as source:
        blob_client.upload_blob(source, overwrite=True)
    print(f"Uploaded blob: {BLOB_NAME}")

    print(f"Blobs in {CONTAINER_NAME}:")
    for blob in container_client.list_blobs():
        print(f"{blob.name} ({blob.size} bytes)")

    with DOWNLOAD_FILE.open("wb") as destination:
        blob_client.download_blob().readinto(destination)
    print(f"Downloaded blob to: {DOWNLOAD_FILE}")

    blob_client.delete_blob()
    print(f"Deleted blob: {BLOB_NAME}")

    container_client.delete_container()
    print(f"Deleted container: {CONTAINER_NAME}")


def main() -> int:
    try:
        run_blob_crud()
    except (FileNotFoundError, ValueError) as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 1
    except HttpResponseError as error:
        print(
            f"Azure Blob Storage request failed "
            f"(status {error.status_code or 'unknown'}): {error.message}",
            file=sys.stderr,
        )
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
