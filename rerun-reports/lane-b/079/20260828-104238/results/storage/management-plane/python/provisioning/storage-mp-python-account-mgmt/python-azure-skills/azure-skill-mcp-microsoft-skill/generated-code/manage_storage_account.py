"""Manage an Azure Storage account through the Azure management-plane SDK."""

from __future__ import annotations

import argparse
import logging
import os
import re
import sys
from collections.abc import Sequence

from azure.core.exceptions import AzureError, ClientAuthenticationError
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    BlobServicePropertiesProperties,
    Kind,
    MinimumTlsVersion,
    Sku,
    SkuName,
    StorageAccountCreateParameters,
    StorageAccountPropertiesCreateParameters,
)

LOCATION = "eastus"
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")
LOGGER = logging.getLogger("storage-account-manager")


def parse_args(argv: Sequence[str] | None = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create, inspect, update, and delete an Azure Storage account."
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
            "Globally unique, 3-24 character lowercase alphanumeric storage "
            "account name (defaults to AZURE_STORAGE_ACCOUNT_NAME)."
        ),
    )
    parser.add_argument(
        "--execute",
        action="store_true",
        help="Perform the Azure operations. Without this flag, only print the plan.",
    )
    return parser.parse_args(argv)


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


def print_plan(args: argparse.Namespace) -> None:
    print("Dry run; no Azure requests were made.")
    print(f"1. Create {args.account_name!r} in {LOCATION} with Standard_LRS.")
    print(f"2. List storage accounts in resource group {args.resource_group!r}.")
    print(f"3. Get properties for {args.account_name!r}.")
    print("4. Enable blob versioning through the Blob service properties.")
    print(f"5. Delete {args.account_name!r}.")
    print("Re-run with --execute to perform these operations.")


def describe_account(prefix: str, account: object) -> None:
    sku = getattr(getattr(account, "sku", None), "name", None)
    print(
        f"{prefix}: name={getattr(account, 'name', None)}, "
        f"location={getattr(account, 'location', None)}, "
        f"kind={getattr(account, 'kind', None)}, sku={sku}"
    )


def create_parameters() -> StorageAccountCreateParameters:
    return StorageAccountCreateParameters(
        location=LOCATION,
        sku=Sku(name=SkuName.STANDARD_LRS),
        kind=Kind.STORAGE_V2,
        properties=StorageAccountPropertiesCreateParameters(
            enable_https_traffic_only=True,
            minimum_tls_version=MinimumTlsVersion.TLS1_2,
            allow_blob_public_access=False,
        ),
    )


def enable_blob_versioning(
    client: StorageManagementClient,
    resource_group: str,
    account_name: str,
) -> BlobServiceProperties:
    current = client.blob_services.get_service_properties(
        resource_group, account_name
    )
    if current.blob_service_properties is None:
        current.blob_service_properties = BlobServicePropertiesProperties()

    current.blob_service_properties.is_versioning_enabled = True
    updated = client.blob_services.set_service_properties(
        resource_group, account_name, current
    )
    if (
        updated.blob_service_properties is None
        or updated.blob_service_properties.is_versioning_enabled is not True
    ):
        raise RuntimeError("Azure did not report blob versioning as enabled.")
    return updated


def execute(args: argparse.Namespace) -> int:
    account_created = False
    exit_code = 0

    try:
        with DefaultAzureCredential() as credential:
            with StorageManagementClient(
                credential=credential,
                subscription_id=args.subscription_id,
            ) as client:
                try:
                    LOGGER.info("Creating storage account %s", args.account_name)
                    created = client.storage_accounts.begin_create(
                        args.resource_group,
                        args.account_name,
                        create_parameters(),
                    ).result()
                    account_created = True
                    describe_account("Created", created)

                    print(
                        f"Storage accounts in resource group "
                        f"{args.resource_group!r}:"
                    )
                    for account in client.storage_accounts.list_by_resource_group(
                        args.resource_group
                    ):
                        describe_account(" - Account", account)

                    properties = client.storage_accounts.get_properties(
                        args.resource_group, args.account_name
                    )
                    describe_account("Properties", properties)

                    enable_blob_versioning(
                        client, args.resource_group, args.account_name
                    )
                    print("Blob versioning enabled: True")
                except (CredentialUnavailableError, ClientAuthenticationError) as exc:
                    LOGGER.error("Azure authentication failed: %s", exc)
                    exit_code = 1
                except AzureError as exc:
                    LOGGER.error("Azure Storage management operation failed: %s", exc)
                    exit_code = 1
                except RuntimeError as exc:
                    LOGGER.error("Storage account verification failed: %s", exc)
                    exit_code = 1
                finally:
                    if account_created:
                        try:
                            LOGGER.info(
                                "Deleting storage account %s", args.account_name
                            )
                            client.storage_accounts.delete(
                                args.resource_group, args.account_name
                            )
                            print(f"Deleted storage account {args.account_name!r}.")
                        except AzureError as exc:
                            LOGGER.error(
                                "Failed to delete storage account %s: %s",
                                args.account_name,
                                exc,
                            )
                            exit_code = 1
    except (CredentialUnavailableError, ClientAuthenticationError) as exc:
        LOGGER.error("Unable to initialize Azure authentication: %s", exc)
        return 1

    return exit_code


def main(argv: Sequence[str] | None = None) -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    args = parse_args(argv)

    try:
        validate_args(args)
    except ValueError as exc:
        LOGGER.error("%s", exc)
        return 2

    if not args.execute:
        print_plan(args)
        return 0

    return execute(args)


if __name__ == "__main__":
    sys.exit(main())
