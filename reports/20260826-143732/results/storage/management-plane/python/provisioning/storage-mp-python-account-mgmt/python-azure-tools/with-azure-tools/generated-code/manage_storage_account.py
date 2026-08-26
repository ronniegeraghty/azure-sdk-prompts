"""Manage an Azure Storage account through the Azure management-plane SDK.

Required environment variables:
    AZURE_SUBSCRIPTION_ID
    AZURE_RESOURCE_GROUP
    AZURE_STORAGE_ACCOUNT_NAME

The identity used by DefaultAzureCredential must have permission to manage
storage accounts in the target resource group.
"""

import argparse
import logging
import os
import re
from collections.abc import Sequence

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    Kind,
    Sku,
    SkuName,
    StorageAccountCheckNameAvailabilityParameters,
    StorageAccountCreateParameters,
)

LOGGER = logging.getLogger("storage-account-manager")
ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def parse_args(argv: Sequence[str] | None = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create, inspect, configure, and delete an Azure Storage account."
    )
    parser.add_argument(
        "--subscription-id",
        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (default: AZURE_SUBSCRIPTION_ID).",
    )
    parser.add_argument(
        "--resource-group",
        default=os.getenv("AZURE_RESOURCE_GROUP"),
        help="Existing resource group name (default: AZURE_RESOURCE_GROUP).",
    )
    parser.add_argument(
        "--account-name",
        default=os.getenv("AZURE_STORAGE_ACCOUNT_NAME"),
        help="Globally unique storage account name (default: AZURE_STORAGE_ACCOUNT_NAME).",
    )
    args = parser.parse_args(argv)

    missing = [
        option
        for option, value in (
            ("--subscription-id", args.subscription_id),
            ("--resource-group", args.resource_group),
            ("--account-name", args.account_name),
        )
        if not value
    ]
    if missing:
        parser.error(f"missing required arguments: {', '.join(missing)}")
    if not ACCOUNT_NAME_PATTERN.fullmatch(args.account_name):
        parser.error(
            "--account-name must contain 3-24 lowercase letters and numbers only"
        )

    return args


def run_workflow(
    client: StorageManagementClient,
    resource_group: str,
    account_name: str,
) -> int:
    creation_started = False
    exit_code = 0

    try:
        availability = client.storage_accounts.check_name_availability(
            StorageAccountCheckNameAvailabilityParameters(name=account_name)
        )
        if not availability.name_available:
            reason = availability.message or availability.reason or "name unavailable"
            raise ValueError(f"Storage account name '{account_name}' is unavailable: {reason}")

        LOGGER.info("Creating storage account '%s' in eastus", account_name)
        poller = client.storage_accounts.begin_create(
            resource_group,
            account_name,
            StorageAccountCreateParameters(
                sku=Sku(name=SkuName.STANDARD_LRS),
                kind=Kind.STORAGE_V2,
                location="eastus",
                enable_https_traffic_only=True,
                minimum_tls_version="TLS1_2",
                allow_blob_public_access=False,
            ),
        )
        creation_started = True
        created_account = poller.result()
        LOGGER.info("Created %s (%s)", created_account.name, created_account.id)

        LOGGER.info("Storage accounts in resource group '%s':", resource_group)
        for account in client.storage_accounts.list_by_resource_group(resource_group):
            LOGGER.info("  %s - %s", account.name, account.location)

        properties = client.storage_accounts.get_properties(
            resource_group, account_name
        )
        LOGGER.info(
            "Account properties: name=%s, location=%s, kind=%s, provisioning_state=%s",
            properties.name,
            properties.location,
            properties.kind,
            properties.provisioning_state,
        )

        blob_properties = client.blob_services.get_service_properties(
            resource_group, account_name
        )
        blob_properties.is_versioning_enabled = True
        client.blob_services.set_service_properties(
            resource_group, account_name, blob_properties
        )
        LOGGER.info("Enabled blob versioning for '%s'", account_name)

    except ValueError as error:
        LOGGER.error("%s", error)
        exit_code = 1
    except CredentialUnavailableError:
        LOGGER.exception(
            "No credential is available. Configure workload identity, managed "
            "identity, Azure CLI authentication, or service-principal environment variables."
        )
        exit_code = 1
    except ClientAuthenticationError:
        LOGGER.exception(
            "Azure authentication failed. Verify the selected identity and tenant."
        )
        exit_code = 1
    except HttpResponseError as error:
        LOGGER.error(
            "Azure Storage management request failed (status %s): %s",
            error.status_code,
            error.message,
        )
        exit_code = 1
    except AzureError:
        LOGGER.exception("An Azure SDK error interrupted the workflow.")
        exit_code = 1
    finally:
        if creation_started:
            try:
                LOGGER.info("Deleting storage account '%s'", account_name)
                client.storage_accounts.delete(resource_group, account_name)
                LOGGER.info("Deleted storage account '%s'", account_name)
            except ResourceNotFoundError:
                LOGGER.info("Storage account '%s' was already deleted", account_name)
            except AzureError:
                LOGGER.exception(
                    "Failed to delete storage account '%s'; manual cleanup is required.",
                    account_name,
                )
                exit_code = 1

    return exit_code


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv)
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

    try:
        with DefaultAzureCredential() as credential:
            with StorageManagementClient(
                credential=credential,
                subscription_id=args.subscription_id,
            ) as client:
                return run_workflow(client, args.resource_group, args.account_name)
    except CredentialUnavailableError:
        LOGGER.exception("DefaultAzureCredential could not find an available credential.")
    except ClientAuthenticationError:
        LOGGER.exception("DefaultAzureCredential could not authenticate.")
    except AzureError:
        LOGGER.exception("Failed to initialize the Azure management client.")
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
