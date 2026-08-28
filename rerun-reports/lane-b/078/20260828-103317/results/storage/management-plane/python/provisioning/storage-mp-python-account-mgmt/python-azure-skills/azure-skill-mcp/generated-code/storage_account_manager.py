"""Create, inspect, update, and delete an Azure Storage account."""

from __future__ import annotations

import argparse
import logging
import os
import re
import sys
import uuid

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    Kind,
    Sku,
    SkuName,
    StorageAccountCreateParameters,
)

LOGGER = logging.getLogger("storage-account-manager")
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description=(
            "Create, list, inspect, enable blob versioning on, and delete an "
            "Azure Storage account."
        )
    )
    parser.add_argument(
        "--subscription-id",
        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
    )
    parser.add_argument(
        "--resource-group",
        default=os.getenv("AZURE_RESOURCE_GROUP"),
        help="Existing resource group (defaults to AZURE_RESOURCE_GROUP).",
    )
    parser.add_argument(
        "--account-name",
        default=os.getenv("AZURE_STORAGE_ACCOUNT_NAME"),
        help=(
            "Globally unique storage account name. If omitted, a valid name "
            "is generated."
        ),
    )
    return parser.parse_args()


def validate_args(args: argparse.Namespace) -> str:
    if not args.subscription_id:
        raise ValueError(
            "A subscription ID is required. Pass --subscription-id or set "
            "AZURE_SUBSCRIPTION_ID."
        )
    if not args.resource_group:
        raise ValueError(
            "A resource group is required. Pass --resource-group or set "
            "AZURE_RESOURCE_GROUP."
        )

    account_name = args.account_name or f"st{uuid.uuid4().hex[:22]}"
    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(account_name):
        raise ValueError(
            "The storage account name must contain 3-24 lowercase letters or "
            "digits only."
        )
    return account_name


def manage_storage_account(
    client: StorageManagementClient,
    resource_group: str,
    account_name: str,
) -> None:
    created = False

    try:
        LOGGER.info("Creating storage account %s in eastus...", account_name)
        create_parameters = StorageAccountCreateParameters(
            sku=Sku(name=SkuName.STANDARD_LRS),
            kind=Kind.STORAGE_V2,
            location="eastus",
        )
        account = client.storage_accounts.begin_create(
            resource_group,
            account_name,
            create_parameters,
        ).result()
        created = True
        LOGGER.info("Created %s (%s).", account.name, account.id)

        LOGGER.info("Storage accounts in resource group %s:", resource_group)
        accounts = list(
            client.storage_accounts.list_by_resource_group(resource_group)
        )
        if accounts:
            for listed_account in accounts:
                LOGGER.info(
                    "  %s | %s | %s",
                    listed_account.name,
                    listed_account.location,
                    listed_account.sku.name if listed_account.sku else "unknown",
                )
        else:
            LOGGER.info("  No storage accounts found.")

        properties = client.storage_accounts.get_properties(
            resource_group,
            account_name,
        )
        LOGGER.info(
            "Properties: name=%s, kind=%s, location=%s, provisioning_state=%s",
            properties.name,
            properties.kind,
            properties.location,
            properties.provisioning_state,
        )

        LOGGER.info("Enabling blob versioning...")
        client.blob_services.set_service_properties(
            resource_group,
            account_name,
            BlobServiceProperties(is_versioning_enabled=True),
        )
        blob_properties = client.blob_services.get_service_properties(
            resource_group,
            account_name,
        )
        if blob_properties.is_versioning_enabled is not True:
            raise RuntimeError("Blob versioning was not enabled as requested.")
        LOGGER.info("Blob versioning is enabled.")
    finally:
        if created:
            operation_failed = sys.exc_info()[0] is not None
            LOGGER.info("Deleting storage account %s...", account_name)
            try:
                client.storage_accounts.delete(resource_group, account_name)
                LOGGER.info("Deleted storage account %s.", account_name)
            except HttpResponseError:
                LOGGER.exception(
                    "Cleanup failed; storage account %s may still exist.",
                    account_name,
                )
                if not operation_failed:
                    raise


def main() -> int:
    args = parse_args()

    try:
        account_name = validate_args(args)
        credential = DefaultAzureCredential()
        client = StorageManagementClient(credential, args.subscription_id)
        manage_storage_account(client, args.resource_group, account_name)
        return 0
    except ValueError as error:
        LOGGER.error("%s", error)
    except ClientAuthenticationError:
        LOGGER.exception(
            "Authentication failed. Configure a supported "
            "DefaultAzureCredential identity."
        )
    except ResourceNotFoundError:
        LOGGER.exception(
            "An Azure resource was not found. Confirm the subscription and "
            "resource group."
        )
    except HttpResponseError:
        LOGGER.exception("An Azure management request failed.")
    except RuntimeError:
        LOGGER.exception("The storage account workflow failed.")

    return 1


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(main())
