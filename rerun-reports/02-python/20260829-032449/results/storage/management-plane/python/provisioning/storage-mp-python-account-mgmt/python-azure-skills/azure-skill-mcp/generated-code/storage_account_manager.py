"""Manage an Azure Storage account through the Azure management plane."""

from __future__ import annotations

import argparse
import logging
import os
import sys
from dataclasses import dataclass

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    Sku,
    StorageAccountCreateParameters,
)

LOG = logging.getLogger("storage-account-manager")


@dataclass(frozen=True)
class Config:
    subscription_id: str
    resource_group_name: str
    storage_account_name: str
    location: str = "eastus"

    @classmethod
    def from_environment(cls, *, require_values: bool) -> "Config":
        values = {
            "subscription_id": os.getenv("AZURE_SUBSCRIPTION_ID", ""),
            "resource_group_name": os.getenv("AZURE_RESOURCE_GROUP", ""),
            "storage_account_name": os.getenv("AZURE_STORAGE_ACCOUNT_NAME", ""),
        }
        missing = [name for name, value in values.items() if not value]
        if require_values and missing:
            environment_names = {
                "subscription_id": "AZURE_SUBSCRIPTION_ID",
                "resource_group_name": "AZURE_RESOURCE_GROUP",
                "storage_account_name": "AZURE_STORAGE_ACCOUNT_NAME",
            }
            missing_names = ", ".join(environment_names[name] for name in missing)
            raise ValueError(f"Missing required environment variables: {missing_names}")

        return cls(
            subscription_id=values["subscription_id"] or "<subscription-id>",
            resource_group_name=values["resource_group_name"] or "<resource-group>",
            storage_account_name=values["storage_account_name"] or "<globally-unique-account-name>",
        )


def print_plan(config: Config) -> None:
    """Print the Azure operations without authenticating or making network calls."""
    print(
        "Dry run; no Azure operations were performed.\n"
        f"1. Create StorageV2 account '{config.storage_account_name}' in "
        f"'{config.resource_group_name}' ({config.location}, Standard_LRS).\n"
        "2. List storage accounts in the resource group.\n"
        "3. Read the created account's properties.\n"
        "4. Enable blob versioning on the account's default Blob service.\n"
        "5. Delete the created storage account.\n"
        "Pass --execute and set the required environment variables to run this plan."
    )


def run_lifecycle(client: StorageManagementClient, config: Config) -> None:
    """Create, inspect, update, and delete one storage account."""
    created = False
    try:
        LOG.info("Creating storage account %s", config.storage_account_name)
        create_poller = client.storage_accounts.begin_create(
            config.resource_group_name,
            config.storage_account_name,
            StorageAccountCreateParameters(
                location=config.location,
                kind="StorageV2",
                sku=Sku(name="Standard_LRS"),
                enable_https_traffic_only=True,
                minimum_tls_version="TLS1_2",
                allow_blob_public_access=False,
            ),
        )
        created_account = create_poller.result()
        created = True
        LOG.info(
            "Created %s with provisioning state %s",
            created_account.name,
            created_account.provisioning_state,
        )

        LOG.info("Storage accounts in resource group %s:", config.resource_group_name)
        for account in client.storage_accounts.list_by_resource_group(
            config.resource_group_name
        ):
            LOG.info("  %s (%s)", account.name, account.location)

        account = client.storage_accounts.get_properties(
            config.resource_group_name,
            config.storage_account_name,
        )
        LOG.info(
            "Properties: id=%s, kind=%s, sku=%s, primary_location=%s",
            account.id,
            account.kind,
            account.sku.name if account.sku else None,
            account.primary_location,
        )

        LOG.info("Enabling blob versioning")
        client.blob_services.set_service_properties(
            config.resource_group_name,
            config.storage_account_name,
            "default",
            BlobServiceProperties(is_versioning_enabled=True),
        )
        blob_properties = client.blob_services.get_service_properties(
            config.resource_group_name,
            config.storage_account_name,
            "default",
        )
        LOG.info(
            "Blob versioning enabled: %s",
            blob_properties.is_versioning_enabled,
        )
    finally:
        if created:
            LOG.info("Deleting storage account %s", config.storage_account_name)
            client.storage_accounts.delete(
                config.resource_group_name,
                config.storage_account_name,
            )
            LOG.info("Storage account deleted")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run an Azure Storage account management-plane lifecycle."
    )
    parser.add_argument(
        "--execute",
        action="store_true",
        help="Perform the Azure operations. Without this flag, only print the plan.",
    )
    return parser.parse_args()


def main() -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    args = parse_args()

    try:
        config = Config.from_environment(require_values=args.execute)
        if not args.execute:
            print_plan(config)
            return 0

        credential = DefaultAzureCredential()
        client = StorageManagementClient(credential, config.subscription_id)
        try:
            run_lifecycle(client, config)
        finally:
            client.close()
            credential.close()
        return 0
    except ValueError as error:
        LOG.error("Configuration error: %s", error)
    except CredentialUnavailableError as error:
        LOG.error("No usable DefaultAzureCredential was found: %s", error)
    except ClientAuthenticationError as error:
        LOG.error("Azure authentication failed: %s", error)
    except ResourceNotFoundError as error:
        LOG.error("An Azure resource was not found: %s", error)
    except ServiceRequestError as error:
        LOG.error("Azure could not be reached: %s", error)
    except HttpResponseError as error:
        LOG.error(
            "Azure rejected an operation (status %s): %s",
            error.status_code,
            error.message,
        )
    except AzureError as error:
        LOG.error("Azure SDK error: %s", error)

    return 1


if __name__ == "__main__":
    sys.exit(main())
