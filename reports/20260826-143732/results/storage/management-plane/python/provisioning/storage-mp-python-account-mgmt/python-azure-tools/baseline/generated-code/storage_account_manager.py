"""Create, inspect, update, and delete an Azure Storage account."""

from __future__ import annotations

import os
import re
import sys
from typing import NoReturn

from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    Kind,
    Sku,
    SkuName,
    StorageAccountCreateParameters,
)


def required_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def validate_storage_account_name(name: str) -> None:
    if not re.fullmatch(r"[a-z0-9]{3,24}", name):
        raise ValueError(
            "STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters or digits."
        )


def print_http_error(error: HttpResponseError) -> None:
    request_id = (
        error.response.headers.get("x-ms-request-id") if error.response else None
    )
    details = f" Request ID: {request_id}." if request_id else ""
    print(f"Azure request failed: {error.message}.{details}", file=sys.stderr)


def fail(message: str, exit_code: int = 1) -> NoReturn:
    print(message, file=sys.stderr)
    raise SystemExit(exit_code)


def main() -> None:
    subscription_id = required_environment_variable("AZURE_SUBSCRIPTION_ID")
    resource_group_name = required_environment_variable("RESOURCE_GROUP_NAME")
    storage_account_name = required_environment_variable("STORAGE_ACCOUNT_NAME")
    validate_storage_account_name(storage_account_name)

    credential = DefaultAzureCredential()
    client = StorageManagementClient(credential, subscription_id)

    print(f"Creating storage account {storage_account_name!r}...")
    create_poller = client.storage_accounts.begin_create(
        resource_group_name,
        storage_account_name,
        StorageAccountCreateParameters(
            sku=Sku(name=SkuName.STANDARD_LRS),
            kind=Kind.STORAGE_V2,
            location="eastus",
        ),
    )
    created_account = create_poller.result()
    print(f"Created: {created_account.id}")

    print(f"\nStorage accounts in resource group {resource_group_name!r}:")
    for account in client.storage_accounts.list_by_resource_group(
        resource_group_name
    ):
        print(f"- {account.name} ({account.location}, {account.sku.name})")

    account = client.storage_accounts.get_properties(
        resource_group_name, storage_account_name
    )
    print("\nCreated account properties:")
    print(f"  ID: {account.id}")
    print(f"  Location: {account.location}")
    print(f"  SKU: {account.sku.name}")
    print(f"  Kind: {account.kind}")
    print(f"  Provisioning state: {account.provisioning_state}")

    # Blob versioning is configured through the storage management plane's
    # blob service resource, rather than StorageAccountsOperations.update.
    print("\nEnabling blob versioning...")
    client.blob_services.set_service_properties(
        resource_group_name,
        storage_account_name,
        "default",
        BlobServiceProperties(is_versioning_enabled=True),
    )
    blob_properties = client.blob_services.get_service_properties(
        resource_group_name, storage_account_name, "default"
    )
    print(f"Blob versioning enabled: {blob_properties.is_versioning_enabled}")

    print(f"\nDeleting storage account {storage_account_name!r}...")
    client.storage_accounts.delete(resource_group_name, storage_account_name)
    print("Storage account deleted.")


if __name__ == "__main__":
    try:
        main()
    except ValueError as error:
        fail(f"Configuration error: {error}", 2)
    except ClientAuthenticationError as error:
        fail(f"Azure authentication failed: {error}", 3)
    except HttpResponseError as error:
        print_http_error(error)
        raise SystemExit(4) from error
    except AzureError as error:
        fail(f"Azure SDK error: {error}", 5)
    except KeyboardInterrupt:
        fail("Operation cancelled.", 130)
