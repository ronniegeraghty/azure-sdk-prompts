"""Create, inspect, configure, and optionally delete an Azure Storage account."""

from __future__ import annotations

import argparse
import logging
import os
import re
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import BlobServiceProperties

LOCATION = "eastus"
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def required_environment_variable(name: str) -> str:
    """Return a required, non-empty environment variable."""
    value = os.getenv(name, "").strip()
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def parse_args(argv: Sequence[str] | None = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Manage the lifecycle of one Azure Storage account."
    )
    parser.add_argument(
        "--delete",
        action="store_true",
        help="Delete the account after enabling and verifying blob versioning.",
    )
    return parser.parse_args(argv)


def validate_storage_account_name(name: str) -> None:
    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(name):
        raise ValueError(
            "STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters and numbers."
        )


def manage_storage_account(
    client: StorageManagementClient,
    resource_group_name: str,
    account_name: str,
    delete_account: bool,
) -> None:
    availability = client.storage_accounts.check_name_availability(
        {"name": account_name, "type": "Microsoft.Storage/storageAccounts"}
    )
    if not availability.name_available:
        reason = availability.reason or "the name is unavailable"
        message = availability.message or "Choose a globally unique account name."
        raise ValueError(f"Cannot create {account_name}: {reason}. {message}")

    logging.info("Creating storage account %s in %s", account_name, LOCATION)
    created_account = client.storage_accounts.begin_create(
        resource_group_name,
        account_name,
        {
            "location": LOCATION,
            "kind": "StorageV2",
            "sku": {"name": "Standard_LRS"},
            "enable_https_traffic_only": True,
            "minimum_tls_version": "TLS1_2",
            "allow_blob_public_access": False,
        },
    ).result()
    logging.info(
        "Created %s with SKU %s",
        created_account.name,
        created_account.sku.name if created_account.sku else "unknown",
    )

    logging.info("Storage accounts in resource group %s:", resource_group_name)
    for account in client.storage_accounts.list_by_resource_group(
        resource_group_name
    ):
        logging.info(
            "  %s (%s, %s)",
            account.name,
            account.location,
            account.sku.name if account.sku else "unknown SKU",
        )

    properties = client.storage_accounts.get_properties(
        resource_group_name, account_name
    )
    logging.info(
        "Properties for %s: id=%s, kind=%s, location=%s",
        properties.name,
        properties.id,
        properties.kind,
        properties.location,
    )

    logging.info("Enabling blob versioning for %s", account_name)
    blob_properties = client.blob_services.set_service_properties(
        resource_group_name,
        account_name,
        BlobServiceProperties(is_versioning_enabled=True),
    )
    if blob_properties.is_versioning_enabled is not True:
        raise RuntimeError(
            f"Azure did not confirm blob versioning for {account_name}."
        )
    logging.info("Blob versioning is enabled")

    if delete_account:
        logging.info("Deleting storage account %s", account_name)
        client.storage_accounts.delete(resource_group_name, account_name)
        logging.info("Deleted storage account %s", account_name)
    else:
        logging.warning(
            "Account %s was not deleted. Run with --delete to complete step 6.",
            account_name,
        )


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv)
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

    credential: DefaultAzureCredential | None = None
    client: StorageManagementClient | None = None
    try:
        subscription_id = required_environment_variable("AZURE_SUBSCRIPTION_ID")
        resource_group_name = required_environment_variable(
            "AZURE_RESOURCE_GROUP_NAME"
        )
        account_name = required_environment_variable("STORAGE_ACCOUNT_NAME")
        validate_storage_account_name(account_name)

        credential = DefaultAzureCredential()
        client = StorageManagementClient(credential, subscription_id)
        manage_storage_account(
            client, resource_group_name, account_name, args.delete
        )
        return 0
    except ClientAuthenticationError as error:
        logging.error("Azure authentication failed: %s", error.message)
    except ResourceNotFoundError as error:
        logging.error("Azure resource was not found: %s", error.message)
    except HttpResponseError as error:
        status = error.status_code if error.status_code is not None else "unknown"
        logging.error("Azure request failed (HTTP %s): %s", status, error.message)
    except (RuntimeError, ValueError) as error:
        logging.error("%s", error)
    except KeyboardInterrupt:
        logging.error("Operation cancelled")
    finally:
        if client is not None:
            client.close()
        if credential is not None:
            credential.close()
    return 1


if __name__ == "__main__":
    sys.exit(main())
