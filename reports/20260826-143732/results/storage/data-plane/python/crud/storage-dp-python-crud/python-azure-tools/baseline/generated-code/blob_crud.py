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


def main() -> int:
    account_url = os.environ.get("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        print(
            "AZURE_STORAGE_ACCOUNT_URL must be set, for example "
            "https://<storage-account>.blob.core.windows.net.",
            file=sys.stderr,
        )
        return 2

    if not SOURCE_FILE.is_file():
        print(f"Source file not found: {SOURCE_FILE}", file=sys.stderr)
        return 2

    credential = DefaultAzureCredential()
    service_client = BlobServiceClient(
        account_url=account_url,
        credential=credential,
    )
    container_client = service_client.get_container_client(CONTAINER_NAME)
    blob_client = container_client.get_blob_client(BLOB_NAME)

    try:
        try:
            container_client.create_container()
            print(f"Created container: {CONTAINER_NAME}")
        except ResourceExistsError:
            print(f"Container already exists: {CONTAINER_NAME}")

        with SOURCE_FILE.open("rb") as source:
            blob_client.upload_blob(source, overwrite=True)
        print(f"Uploaded {SOURCE_FILE} to {BLOB_NAME}")

        for blob in container_client.list_blobs():
            print(f"{blob.name}: {blob.size} bytes")

        with DOWNLOAD_FILE.open("wb") as destination:
            destination.write(blob_client.download_blob().readall())
        print(f"Downloaded {BLOB_NAME} to {DOWNLOAD_FILE}")

        blob_client.delete_blob()
        print(f"Deleted blob: {BLOB_NAME}")

        container_client.delete_container()
        print(f"Deleted container: {CONTAINER_NAME}")
    except HttpResponseError as error:
        print(f"Azure Blob Storage request failed: {error}", file=sys.stderr)
        return 1
    finally:
        service_client.close()
        credential.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
