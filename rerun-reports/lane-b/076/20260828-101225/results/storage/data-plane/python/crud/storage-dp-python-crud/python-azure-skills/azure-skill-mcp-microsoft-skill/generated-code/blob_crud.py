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


def get_account_url() -> str:
    account_url = os.getenv("AZURE_STORAGE_ACCOUNT_URL")
    if account_url:
        return account_url

    account_name = os.getenv("AZURE_STORAGE_ACCOUNT_NAME")
    if account_name:
        return f"https://{account_name}.blob.core.windows.net"

    raise ValueError(
        "Set AZURE_STORAGE_ACCOUNT_URL or AZURE_STORAGE_ACCOUNT_NAME."
    )


def main() -> int:
    try:
        account_url = get_account_url()
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 1

    if not SOURCE_FILE.is_file():
        print(f"Input file not found: {SOURCE_FILE}", file=sys.stderr)
        return 1

    credential = DefaultAzureCredential()
    try:
        with BlobServiceClient(
            account_url=account_url,
            credential=credential,
        ) as blob_service_client:
            container_client = blob_service_client.get_container_client(
                CONTAINER_NAME
            )

            try:
                container_client.create_container()
                print(f"Created container: {CONTAINER_NAME}")
            except ResourceExistsError:
                print(f"Container already exists: {CONTAINER_NAME}")

            blob_client = container_client.get_blob_client(BLOB_NAME)
            with SOURCE_FILE.open("rb") as source:
                blob_client.upload_blob(source, overwrite=True)
            print(f"Uploaded blob: {BLOB_NAME}")

            for blob in container_client.list_blobs():
                print(f"{blob.name}: {blob.size} bytes")

            with DOWNLOAD_FILE.open("wb") as destination:
                blob_client.download_blob().readinto(destination)
            print(f"Downloaded blob to: {DOWNLOAD_FILE}")

            blob_client.delete_blob()
            print(f"Deleted blob: {BLOB_NAME}")

            container_client.delete_container()
            print(f"Deleted container: {CONTAINER_NAME}")
    except HttpResponseError as error:
        print(f"Azure Blob Storage request failed: {error}", file=sys.stderr)
        return 1
    except OSError as error:
        print(f"File operation failed: {error}", file=sys.stderr)
        return 1
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
