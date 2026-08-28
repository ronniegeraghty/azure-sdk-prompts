from __future__ import annotations

import logging
import os
import sys
from dataclasses import dataclass

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError, ServiceRequestError
from azure.identity import ClientSecretCredential, CredentialUnavailableError
from azure.storage.blob import BlobServiceClient
from dotenv import load_dotenv

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class AzureConfig:
    tenant_id: str
    client_id: str
    client_secret: str
    storage_account_url: str


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing or invalid."""


def load_config() -> AzureConfig:
    # Existing environment variables take precedence over values in a local .env file.
    load_dotenv(override=False)

    names = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_CLIENT_SECRET",
        "AZURE_STORAGE_ACCOUNT_URL",
    )
    values = {name: os.getenv(name, "").strip() for name in names}
    missing = [name for name, value in values.items() if not value]
    if missing:
        raise ConfigurationError(
            f"Missing required environment variables: {', '.join(missing)}"
        )

    account_url = values["AZURE_STORAGE_ACCOUNT_URL"].rstrip("/")
    if not account_url.startswith("https://"):
        raise ConfigurationError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS.")

    return AzureConfig(
        tenant_id=values["AZURE_TENANT_ID"],
        client_id=values["AZURE_CLIENT_ID"],
        client_secret=values["AZURE_CLIENT_SECRET"],
        storage_account_url=account_url,
    )


def list_storage_containers(config: AzureConfig) -> list[str]:
    with ClientSecretCredential(
        tenant_id=config.tenant_id,
        client_id=config.client_id,
        client_secret=config.client_secret,
    ) as credential:
        with BlobServiceClient(
            account_url=config.storage_account_url,
            credential=credential,
        ) as client:
            return [container["name"] for container in client.list_containers()]


def main() -> int:
    try:
        config = load_config()
        container_names = list_storage_containers(config)
    except ConfigurationError as error:
        logger.error("Configuration error: %s", error)
        return 2
    except CredentialUnavailableError as error:
        logger.error("The configured credential is unavailable: %s", error)
        return 3
    except ClientAuthenticationError as error:
        logger.error(
            "Azure authentication failed. Check the tenant ID, client ID, client "
            "secret, secret expiration, and service principal status: %s",
            error,
        )
        return 3
    except ServiceRequestError as error:
        logger.error("Could not connect to Azure Storage: %s", error)
        return 4
    except HttpResponseError as error:
        logger.error(
            "Azure Storage rejected the request (status %s): %s",
            error.status_code,
            error.message,
        )
        return 5

    if container_names:
        print("Containers:")
        for name in container_names:
            print(f"- {name}")
    else:
        print("No containers found.")
    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(main())
