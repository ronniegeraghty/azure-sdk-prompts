"""Create, inspect, configure, and delete an Azure Storage Account."""

from __future__ import annotations

import argparse
import logging
import os
import re

from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServicePropertiesProperties,
    Kind,
    Sku,
    StorageAccountCreateParameters,
)

LOCATION = "eastus"
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")

logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
logger = logging.getLogger(__name__)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Manage the lifecycle of an Azure Storage Account."
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
            "Globally unique account name (defaults to "
            "AZURE_STORAGE_ACCOUNT_NAME)."
        ),
    )
    return parser.parse_args()


def validate_args(args: argparse.Namespace) -> None:
    missing = [
        name
        for name, value in (
            ("subscription ID", args.subscription_id),
            ("resource group", args.resource_group),
            ("account name", args.account_name),
        )
        if not value
    ]
    if missing:
        raise ValueError(f"Missing required configuration: {', '.join(missing)}")

    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(args.account_name):
        raise ValueError(
            "The account name must contain 3-24 lowercase letters and numbers."
        )


def delete_account(
    client: StorageManagementClient, resource_group: str, account_name: str
) -> None:
    client.storage_accounts.delete(resource_group, account_name)
    logger.info("Deleted Storage Account '%s'.", account_name)


def manage_storage_account(args: argparse.Namespace) -> int:
    validate_args(args)
    credential = DefaultAzureCredential()
    client = StorageManagementClient(credential, args.subscription_id)
    account_created = False

    try:
        availability = client.storage_accounts.check_name_availability(
            {"name": args.account_name, "type": "Microsoft.Storage/storageAccounts"}
        )
        if not availability.name_available:
            raise ValueError(
                f"Storage Account name '{args.account_name}' is unavailable: "
                f"{availability.reason or 'unknown reason'}"
            )

        logger.info(
            "Creating Storage Account '%s' in %s...", args.account_name, LOCATION
        )
        account = client.storage_accounts.begin_create(
            args.resource_group,
            args.account_name,
            StorageAccountCreateParameters(
                sku=Sku(name="Standard_LRS"),
                kind=Kind.STORAGE_V2,
                location=LOCATION,
            ),
        ).result()
        account_created = True
        logger.info("Created %s with resource ID %s.", account.name, account.id)

        logger.info("Storage Accounts in resource group '%s':", args.resource_group)
        for listed_account in client.storage_accounts.list_by_resource_group(
            args.resource_group
        ):
            logger.info(
                "  %s (%s, %s)",
                listed_account.name,
                listed_account.location,
                listed_account.sku.name,
            )

        properties = client.storage_accounts.get_properties(
            args.resource_group, args.account_name
        )
        logger.info(
            "Properties: name=%s, location=%s, sku=%s, kind=%s, id=%s",
            properties.name,
            properties.location,
            properties.sku.name,
            properties.kind,
            properties.id,
        )

        blob_service = client.blob_services.get_service_properties(
            args.resource_group, args.account_name
        )
        if blob_service.blob_service_properties is None:
            blob_service.blob_service_properties = BlobServicePropertiesProperties()
        blob_service.blob_service_properties.is_versioning_enabled = True

        updated_blob_service = client.blob_services.set_service_properties(
            args.resource_group,
            args.account_name,
            parameters=blob_service,
        )
        versioning_enabled = (
            updated_blob_service.blob_service_properties.is_versioning_enabled
        )
        logger.info("Blob versioning enabled: %s", versioning_enabled)

        delete_account(client, args.resource_group, args.account_name)
        account_created = False
        return 0
    except ClientAuthenticationError as error:
        logger.error("Azure authentication failed: %s", error)
        return 2
    except HttpResponseError as error:
        status_code = error.status_code or "unknown"
        logger.error("Azure request failed (HTTP %s): %s", status_code, error.message)
        return 3
    except AzureError as error:
        logger.error("Azure SDK operation failed: %s", error)
        return 4
    except ValueError as error:
        logger.error("%s", error)
        return 5
    finally:
        if account_created:
            logger.warning("Cleaning up the Storage Account after a failed operation.")
            try:
                delete_account(client, args.resource_group, args.account_name)
            except AzureError as cleanup_error:
                logger.error(
                    "Cleanup failed; delete '%s' manually: %s",
                    args.account_name,
                    cleanup_error,
                )
        credential.close()


def main() -> int:
    try:
        args = parse_args()
        return manage_storage_account(args)
    except ValueError as error:
        logger.error("%s", error)
        return 5


if __name__ == "__main__":
    raise SystemExit(main())
