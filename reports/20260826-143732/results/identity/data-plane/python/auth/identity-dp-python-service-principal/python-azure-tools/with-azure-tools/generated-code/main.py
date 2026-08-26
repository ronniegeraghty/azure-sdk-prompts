import logging
import os
import sys
from dataclasses import dataclass
from itertools import islice
from urllib.parse import urlparse

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import ClientSecretCredential, CredentialUnavailableError
from azure.storage.blob import BlobServiceClient
from dotenv import load_dotenv

STORAGE_SCOPE = "https://storage.azure.com/.default"
MAX_CONTAINERS = 10

logger = logging.getLogger(__name__)


class ConfigurationError(ValueError):
    pass


@dataclass(frozen=True)
class Settings:
    tenant_id: str
    client_id: str
    client_secret: str
    storage_account_url: str

    @classmethod
    def from_environment(cls) -> "Settings":
        variable_names = (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_CLIENT_SECRET",
            "AZURE_STORAGE_ACCOUNT_URL",
        )
        missing = [name for name in variable_names if not os.getenv(name)]
        if missing:
            raise ConfigurationError(
                f"Missing required environment variables: {', '.join(missing)}"
            )

        storage_account_url = os.environ["AZURE_STORAGE_ACCOUNT_URL"].rstrip("/")
        parsed_url = urlparse(storage_account_url)
        if parsed_url.scheme != "https" or not parsed_url.netloc:
            raise ConfigurationError(
                "AZURE_STORAGE_ACCOUNT_URL must be a valid HTTPS URL."
            )

        return cls(
            tenant_id=os.environ["AZURE_TENANT_ID"],
            client_id=os.environ["AZURE_CLIENT_ID"],
            client_secret=os.environ["AZURE_CLIENT_SECRET"],
            storage_account_url=storage_account_url,
        )


def create_credential(settings: Settings) -> ClientSecretCredential:
    return ClientSecretCredential(
        tenant_id=settings.tenant_id,
        client_id=settings.client_id,
        client_secret=settings.client_secret,
    )


def list_container_names(settings: Settings) -> list[str]:
    with create_credential(settings) as credential:
        # Acquire a token first so authentication failures are reported separately.
        credential.get_token(STORAGE_SCOPE)

        with BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ) as blob_client:
            containers = blob_client.list_containers()
            return [
                container["name"]
                for container in islice(containers, MAX_CONTAINERS)
            ]


def run() -> int:
    load_dotenv(override=False)

    try:
        settings = Settings.from_environment()
        container_names = list_container_names(settings)
    except ConfigurationError as error:
        logger.error("Configuration error: %s", error)
        return 2
    except (CredentialUnavailableError, ClientAuthenticationError):
        logger.error(
            "Azure authentication failed. Verify the tenant ID, client ID, "
            "client secret, secret expiration, and service principal status."
        )
        return 3
    except HttpResponseError as error:
        logger.error(
            "Azure Storage rejected the request (status=%s, reason=%s). "
            "Verify the account URL and assign the least-privileged Storage "
            "Blob Data Reader role.",
            error.status_code,
            error.reason,
        )
        return 4
    except ServiceRequestError as error:
        logger.error("Could not reach Azure Storage: %s", error)
        return 5

    if container_names:
        print("Containers (up to 10):")
        for name in container_names:
            print(f"- {name}")
    else:
        print("Authentication succeeded; no containers were found.")
    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(run())
