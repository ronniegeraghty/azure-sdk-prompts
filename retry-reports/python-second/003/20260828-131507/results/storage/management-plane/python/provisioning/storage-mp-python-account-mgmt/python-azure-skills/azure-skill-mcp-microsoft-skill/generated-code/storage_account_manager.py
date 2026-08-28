"""Manage an Azure Storage account with the Azure management-plane SDK."""

from __future__ import annotations

import argparse
import os
import re
import sys

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
    StorageAccountCreateParameters,
)

LOCATION = "eastus"
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create, inspect, update, and optionally delete an Azure Storage account."
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
            "Globally unique storage account name "
            "(defaults to AZURE_STORAGE_ACCOUNT_NAME)."
        ),
    )
    parser.add_argument(
        "--keep",
        action="store_true",
        help="Keep the storage account instead of deleting it after the example.",
    )
    return parser.parse_args()


def validate_args(args: argparse.Namespace) -> None:
    missing = [
        name
        for name, value in (
            ("subscription ID", args.subscription_id),
            ("resource group", args.resource_group),
            ("storage account name", args.account_name),
        )
        if not value
    ]
    if missing:
        raise ValueError(f"Missing required value(s): {', '.join(missing)}")

    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(args.account_name):
        raise ValueError(
            "Storage account name must contain 3-24 lowercase letters or digits."
        )


def manage_storage_account(
    client: StorageManagementClient,
    resource_group: str,
    account_name: str,
    keep_account: bool,
) -> None:
    print(f"Creating storage account '{account_name}' in {LOCATION}...")
    account = client.storage_accounts.begin_create(
        resource_group_name=resource_group,
        account_name=account_name,
        parameters=StorageAccountCreateParameters(
            sku=Sku(name="Standard_LRS"),
            kind=Kind.STORAGE_V2,
            location=LOCATION,
        ),
    ).result()
    print(f"Created: {account.id}")

    print(f"\nStorage accounts in resource group '{resource_group}':")
    for item in client.storage_accounts.list_by_resource_group(resource_group):
        print(f"- {item.name} ({item.location}, {item.sku.name})")

    properties = client.storage_accounts.get_properties(
        resource_group_name=resource_group,
        account_name=account_name,
    )
    print("\nCreated account properties:")
    print(f"- name: {properties.name}")
    print(f"- location: {properties.location}")
    print(f"- kind: {properties.kind}")
    print(f"- provisioning state: {properties.provisioning_state}")
    print(f"- primary blob endpoint: {properties.primary_endpoints.blob}")

    print("\nEnabling blob versioning...")
    blob_properties = client.blob_services.set_service_properties(
        resource_group_name=resource_group,
        account_name=account_name,
        parameters=BlobServiceProperties(is_versioning_enabled=True),
    )
    print(f"Blob versioning enabled: {blob_properties.is_versioning_enabled}")

    if keep_account:
        print("\nStorage account retained because --keep was specified.")
    else:
        print(f"\nDeleting storage account '{account_name}'...")
        client.storage_accounts.delete(
            resource_group_name=resource_group,
            account_name=account_name,
        )
        print("Storage account deleted.")


def main() -> int:
    args = parse_args()

    try:
        validate_args(args)
        credential = DefaultAzureCredential()
        client = StorageManagementClient(credential, args.subscription_id)
        manage_storage_account(
            client=client,
            resource_group=args.resource_group,
            account_name=args.account_name,
            keep_account=args.keep,
        )
        return 0
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error.message}", file=sys.stderr)
        return 3
    except ResourceNotFoundError as error:
        print(f"Azure resource was not found: {error.message}", file=sys.stderr)
        return 4
    except HttpResponseError as error:
        status = error.status_code if error.status_code is not None else "unknown"
        print(
            f"Azure request failed (HTTP {status}): {error.message}",
            file=sys.stderr,
        )
        return 5


if __name__ == "__main__":
    raise SystemExit(main())
