from __future__ import annotations

import argparse
import logging
import os
import sys
from dataclasses import dataclass

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import ClientSecretCredential, CredentialUnavailableError
from azure.storage.blob import BlobServiceClient
from dotenv import load_dotenv

LOGGER = logging.getLogger("service-principal-example")
BLOB_SCOPE = "https://storage.azure.com/.default"


class ConfigurationError(ValueError):
    """Raised when required environment configuration is missing."""


@dataclass(frozen=True)
class Settings:
    tenant_id: str
    client_id: str
    client_secret: str
    storage_account_url: str

    @classmethod
    def from_environment(cls) -> "Settings":
        names = (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_CLIENT_SECRET",
            "AZURE_STORAGE_ACCOUNT_URL",
        )
        values = {name: os.environ.get(name, "").strip() for name in names}
        missing = [name for name, value in values.items() if not value]
        if missing:
            raise ConfigurationError(
                "Missing required environment variables: " + ", ".join(missing)
            )

        account_url = values["AZURE_STORAGE_ACCOUNT_URL"].rstrip("/")
        if not account_url.startswith("https://"):
            raise ConfigurationError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS")

        return cls(
            tenant_id=values["AZURE_TENANT_ID"],
            client_id=values["AZURE_CLIENT_ID"],
            client_secret=values["AZURE_CLIENT_SECRET"],
            storage_account_url=account_url,
        )


def create_credential(settings: Settings) -> ClientSecretCredential:
    return ClientSecretCredential(
        tenant_id=settings.tenant_id,
        client_id=settings.client_id,
        client_secret=settings.client_secret,
    )


def check_azure_access(settings: Settings) -> None:
    with create_credential(settings) as credential:
        # Acquire a token explicitly so authentication errors are distinguishable
        # from authorization, networking, and service errors.
        credential.get_token(BLOB_SCOPE)

        with BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ) as blob_service_client:
            account_info = blob_service_client.get_account_information()

    LOGGER.info(
        "Authenticated successfully; storage account kind is %s",
        account_info.get("sku_name", "unknown"),
    )


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Configure Azure Blob Storage with a service principal."
    )
    parser.add_argument(
        "--check-auth",
        action="store_true",
        help="Contact Microsoft Entra ID and Azure Blob Storage to verify access.",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    load_dotenv(override=False)

    try:
        settings = Settings.from_environment()
        if not args.check_auth:
            with create_credential(settings) as credential:
                with BlobServiceClient(
                    account_url=settings.storage_account_url,
                    credential=credential,
                ):
                    LOGGER.info(
                        "Azure credential and BlobServiceClient configured. "
                        "Run with --check-auth to make a request."
                    )
            return 0

        check_azure_access(settings)
        return 0
    except ConfigurationError as error:
        LOGGER.error("Configuration error: %s", error)
        return 2
    except CredentialUnavailableError as error:
        LOGGER.error("The service principal credential is unavailable: %s", error)
        return 3
    except ClientAuthenticationError as error:
        LOGGER.error(
            "Microsoft Entra authentication failed. Verify the tenant ID, client ID, "
            "client secret, secret expiry, and service principal status. Details: %s",
            error,
        )
        return 3
    except ServiceRequestError as error:
        LOGGER.error("Could not reach Azure: %s", error)
        return 4
    except HttpResponseError as error:
        if error.status_code in (401, 403):
            LOGGER.error(
                "Azure rejected the request with HTTP %s. Authentication may have "
                "failed, or the service principal may lack a Blob Storage data role.",
                error.status_code,
            )
            return 3
        LOGGER.error(
            "Azure Blob Storage returned HTTP %s: %s",
            error.status_code,
            error,
        )
        return 5


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(main())
