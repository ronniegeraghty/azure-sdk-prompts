import logging
import os
import sys
from collections.abc import Mapping

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import ClientSecretCredential
from azure.mgmt.resource import ResourceManagementClient
from dotenv import load_dotenv


LOGGER = logging.getLogger(__name__)
REQUIRED_ENVIRONMENT_VARIABLES = (
    "AZURE_TENANT_ID",
    "AZURE_CLIENT_ID",
    "AZURE_CLIENT_SECRET",
    "AZURE_SUBSCRIPTION_ID",
)


def read_configuration(environment: Mapping[str, str]) -> dict[str, str]:
    missing = [name for name in REQUIRED_ENVIRONMENT_VARIABLES if not environment.get(name)]
    if missing:
        raise ValueError(
            "Missing required environment variables: " + ", ".join(missing)
        )

    return {name: environment[name] for name in REQUIRED_ENVIRONMENT_VARIABLES}


def create_credential(configuration: Mapping[str, str]) -> ClientSecretCredential:
    return ClientSecretCredential(
        tenant_id=configuration["AZURE_TENANT_ID"],
        client_id=configuration["AZURE_CLIENT_ID"],
        client_secret=configuration["AZURE_CLIENT_SECRET"],
    )


def list_resource_groups(
    credential: ClientSecretCredential, subscription_id: str
) -> None:
    with ResourceManagementClient(
        credential=credential,
        subscription_id=subscription_id,
    ) as resource_client:
        resource_groups = resource_client.resource_groups.list()
        found_resource_group = False

        for resource_group in resource_groups:
            found_resource_group = True
            print(resource_group.name)

        if not found_resource_group:
            print("No resource groups found.")


def main() -> int:
    # load_dotenv does not override values already supplied by the environment.
    load_dotenv()

    try:
        configuration = read_configuration(os.environ)
    except ValueError as error:
        LOGGER.error("%s", error)
        return 2

    credential = create_credential(configuration)
    try:
        list_resource_groups(
            credential,
            subscription_id=configuration["AZURE_SUBSCRIPTION_ID"],
        )
    except ClientAuthenticationError as error:
        LOGGER.error(
            "Azure authentication failed. Verify the tenant ID, client ID, "
            "client secret, and service principal permissions: %s",
            error.message,
        )
        return 3
    except HttpResponseError as error:
        LOGGER.error(
            "Azure rejected the resource request (status %s): %s",
            error.status_code,
            error.message,
        )
        return 4
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(main())
