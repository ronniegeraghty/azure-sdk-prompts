import os
from pathlib import Path
from urllib.parse import urlsplit, urlunsplit

from azure.ai.projects import AIProjectClient
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobClient


def main() -> None:
    endpoint = require_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
    dataset_name = require_environment_variable("DATASET_NAME")
    dataset_version = require_environment_variable("DATASET_VERSION")
    input_path = Path(require_environment_variable("DATA_FILE_PATH"))
    download_path = Path(require_environment_variable("DOWNLOAD_FILE_PATH"))
    expected_bytes = input_path.read_bytes()

    uploaded = False
    with (
        DefaultAzureCredential() as credential,
        AIProjectClient(endpoint, credential) as project_client,
    ):
        try:
            project_client.datasets.upload_file(
                name=dataset_name,
                version=dataset_version,
                file_path=str(input_path),
            )
            uploaded = True

            dataset = project_client.datasets.get(
                name=dataset_name,
                version=dataset_version,
            )
            print(
                f"Dataset: name={dataset.name} version={dataset.version} "
                f"id={dataset.id} type={dataset.type} dataUri={dataset.data_uri}"
            )

            dataset_credential = project_client.datasets.get_credentials(
                name=dataset_name,
                version=dataset_version,
            )
            blob_reference = dataset_credential.blob_reference
            blob_uri = urlsplit(blob_reference.blob_uri)
            sas_uri = urlsplit(blob_reference.credential.sas_uri)
            authorized_blob_uri = urlunsplit(
                (
                    sas_uri.scheme,
                    sas_uri.netloc,
                    blob_uri.path,
                    sas_uri.query,
                    "",
                )
            )

            download_path.parent.mkdir(parents=True, exist_ok=True)
            with BlobClient.from_blob_url(authorized_blob_uri) as blob_client:
                download_path.write_bytes(blob_client.download_blob().readall())

            downloaded_bytes = download_path.read_bytes()
            if downloaded_bytes != expected_bytes:
                raise RuntimeError(
                    "Downloaded bytes don't match the source file."
                )
            print(f"Downloaded bytes verified: {len(downloaded_bytes)}")
        finally:
            if uploaded:
                project_client.datasets.delete(
                    name=dataset_name,
                    version=dataset_version,
                )


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
