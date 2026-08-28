"""Authenticate with an Azure service principal and list resource groups."""

from __future__ import annotations

import logging
import os
import sys
from dataclasses import dataclass
from typing import Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import ClientSecretCredential
from azure.mgmt.resource import ResourceManagementClient
from dotenv import load_dotenv

LOGGER = logging.getLogger(__name__)
ARM_SCOPE = "https://management.azure.com/.default"


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing."""


@dataclass(frozen=True)
class AzureSettings:
    tenant_id: str
    client_id: str
    client_secret: str
    subscription_id: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        values = {
            name: os.getenv(name, "").strip()
            for name in (
                "AZURE_TENANT_ID",
                "AZURE_CLIENT_ID",
                "AZURE_CLIENT_SECRET",
                "AZURE_SUBSCRIPTION_ID",
            )
        }
        missing = [name for name, value in values.items() if not value]
        if missing:
            raise ConfigurationError(
                "Missing required environment variables: " + ", ".join(missing)
            )

        return cls(
            tenant_id=values["AZURE_TENANT_ID"],
            client_id=values["AZURE_CLIENT_ID"],
            client_secret=values["AZURE_CLIENT_SECRET"],
            subscription_id=values["AZURE_SUBSCRIPTION_ID"],
        )


def create_credential(settings: AzureSettings) -> ClientSecretCredential:
    """Create the deterministic credential used by this service."""
    return ClientSecretCredential(
        tenant_id=settings.tenant_id,
        client_id=settings.client_id,
        client_secret=settings.client_secret,
    )


def list_resource_group_names(
    settings: AzureSettings,
    credential: ClientSecretCredential,
) -> list[str]:
    """Verify authentication and return resource group names."""
    credential.get_token(ARM_SCOPE)
    client = ResourceManagementClient(credential, settings.subscription_id)
    return [resource_group.name for resource_group in client.resource_groups.list()]


def run() -> int:
    load_dotenv(override=False)

    try:
        settings = AzureSettings.from_environment()
        credential = create_credential(settings)
        names = list_resource_group_names(settings, credential)
    except ConfigurationError as error:
        LOGGER.error("Configuration error: %s", error)
        return 2
    except ClientAuthenticationError:
        LOGGER.error(
            "Azure authentication failed. Verify the tenant ID, client ID, "
            "client secret, secret expiration, and service principal status."
        )
        return 3
    except ServiceRequestError as error:
        LOGGER.error("Could not reach Azure: %s", error)
        return 4
    except HttpResponseError as error:
        status_code = getattr(error, "status_code", None)
        LOGGER.error(
            "Azure rejected the Resource Manager request (status %s). Verify "
            "the subscription ID and least-privilege RBAC assignment.",
            status_code if status_code is not None else "unknown",
        )
        return 5

    if names:
        print("Resource groups:")
        for name in names:
            print(f"- {name}")
    else:
        print("No resource groups found.")
    return 0


def main(argv: Sequence[str] | None = None) -> int:
    del argv
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    return run()


if __name__ == "__main__":
    sys.exit(main())
