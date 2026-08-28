import logging
import os
import sys
from dataclasses import dataclass

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import ClientSecretCredential, CredentialUnavailableError
from azure.mgmt.resource.resources import ResourceManagementClient
from dotenv import load_dotenv


LOGGER = logging.getLogger("azure_service_principal_example")


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing."""


@dataclass(frozen=True)
class Settings:
    tenant_id: str
    client_id: str
    client_secret: str
    subscription_id: str

    @classmethod
    def from_environment(cls) -> "Settings":
        values = {
            "tenant_id": os.getenv("AZURE_TENANT_ID", "").strip(),
            "client_id": os.getenv("AZURE_CLIENT_ID", "").strip(),
            "client_secret": os.getenv("AZURE_CLIENT_SECRET", "").strip(),
            "subscription_id": os.getenv("AZURE_SUBSCRIPTION_ID", "").strip(),
        }
        missing = [
            name.upper()
            for name, value in values.items()
            if not value
        ]
        if missing:
            environment_names = ", ".join(f"AZURE_{name}" for name in missing)
            raise ConfigurationError(
                f"Missing required environment variables: {environment_names}"
            )
        return cls(**values)


def create_resource_client(settings: Settings) -> ResourceManagementClient:
    credential = ClientSecretCredential(
        tenant_id=settings.tenant_id,
        client_id=settings.client_id,
        client_secret=settings.client_secret,
    )
    return ResourceManagementClient(
        credential=credential,
        subscription_id=settings.subscription_id,
    )


def list_resource_groups(client: ResourceManagementClient) -> None:
    LOGGER.info("Resource groups:")
    found_resource_group = False
    for resource_group in client.resource_groups.list():
        found_resource_group = True
        LOGGER.info("  %s (%s)", resource_group.name, resource_group.location)

    if not found_resource_group:
        LOGGER.info("  No resource groups found.")


def run() -> int:
    # Existing environment variables take precedence over local .env values.
    load_dotenv(override=False)

    try:
        settings = Settings.from_environment()
        client = create_resource_client(settings)
        list_resource_groups(client)
    except ConfigurationError as error:
        LOGGER.error("Configuration error: %s", error)
        return 2
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        LOGGER.error(
            "Azure authentication failed. Verify the tenant ID, client ID, "
            "client secret value, and secret expiration. Details: %s",
            error,
        )
        return 3
    except HttpResponseError as error:
        if error.status_code == 403:
            LOGGER.error(
                "Azure denied access. Assign the service principal a least-privilege "
                "Azure RBAC role for subscription %s.",
                settings.subscription_id,
            )
        else:
            LOGGER.error(
                "The Azure Resource Manager request failed (HTTP %s): %s",
                error.status_code or "unknown",
                error,
            )
        return 4

    return 0


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(run())


if __name__ == "__main__":
    main()
