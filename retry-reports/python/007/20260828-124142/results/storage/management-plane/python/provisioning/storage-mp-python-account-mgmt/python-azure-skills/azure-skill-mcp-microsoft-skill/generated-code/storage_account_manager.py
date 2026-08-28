"""Create, inspect, configure, and delete an Azure Storage account."""

from __future__ import annotations

import os
import re
import sys
import uuid

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    BlobServicePropertiesProperties,
    Sku,
    StorageAccountCheckNameAvailabilityParameters,
    StorageAccountCreateParameters,
    StorageAccountPropertiesCreateParameters,
)

LOCATION = "eastus"
STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def required_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def get_storage_account_name() -> str:
    name = os.getenv("AZURE_STORAGE_ACCOUNT_NAME")
    if not name:
        name = f"st{uuid.uuid4().hex[:22]}"

    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(name):
        raise ValueError(
            "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters "
            "and numbers."
        )
    return name


def describe_http_error(error: HttpResponseError) -> str:
    status = f"HTTP {error.status_code}" if error.status_code else "HTTP error"
    code = getattr(getattr(error, "error", None), "code", None)
    return (
        f"{status} ({code}): {error.message}"
        if code
        else f"{status}: {error.message}"
    )


def manage_storage_account() -> int:
    try:
        subscription_id = required_environment_variable("AZURE_SUBSCRIPTION_ID")
        resource_group_name = required_environment_variable("AZURE_RESOURCE_GROUP")
        account_name = get_storage_account_name()
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2

    cleanup_required = False

    try:
        with DefaultAzureCredential() as credential:
            client = StorageManagementClient(
                credential=credential,
                subscription_id=subscription_id,
            )
            try:
                availability = client.storage_accounts.check_name_availability(
                    StorageAccountCheckNameAvailabilityParameters(
                        name=account_name,
                        type="Microsoft.Storage/storageAccounts",
                    )
                )
                if not availability.name_available:
                    print(
                        f"Storage account name {account_name!r} is unavailable: "
                        f"{availability.reason or availability.message}",
                        file=sys.stderr,
                    )
                    return 2

                print(f"Creating storage account {account_name!r} in {LOCATION}...")
                cleanup_required = True
                account = client.storage_accounts.begin_create(
                    resource_group_name=resource_group_name,
                    account_name=account_name,
                    parameters=StorageAccountCreateParameters(
                        sku=Sku(name="Standard_LRS"),
                        kind="StorageV2",
                        location=LOCATION,
                        properties=StorageAccountPropertiesCreateParameters(
                            enable_https_traffic_only=True,
                            minimum_tls_version="TLS1_2",
                            allow_blob_public_access=False,
                            allow_cross_tenant_replication=False,
                        ),
                    ),
                ).result()
                print(f"Created: {account.id}")

                print(f"\nStorage accounts in resource group {resource_group_name!r}:")
                accounts = client.storage_accounts.list_by_resource_group(
                    resource_group_name
                )
                for listed_account in accounts:
                    print(
                        f"- {listed_account.name} "
                        f"({listed_account.location}, {listed_account.sku.name})"
                    )

                properties = client.storage_accounts.get_properties(
                    resource_group_name=resource_group_name,
                    account_name=account_name,
                )
                print(
                    "\nCreated account properties:\n"
                    f"  name: {properties.name}\n"
                    f"  location: {properties.location}\n"
                    f"  sku: {properties.sku.name}\n"
                    f"  kind: {properties.kind}\n"
                    f"  provisioning state: {properties.provisioning_state}"
                )

                blob_properties = client.blob_services.set_service_properties(
                    resource_group_name=resource_group_name,
                    account_name=account_name,
                    parameters=BlobServiceProperties(
                        blob_service_properties=BlobServicePropertiesProperties(
                            is_versioning_enabled=True
                        )
                    ),
                )
                versioning_enabled = (
                    blob_properties.blob_service_properties.is_versioning_enabled
                )
                print(f"\nBlob versioning enabled: {versioning_enabled}")

                print(f"\nDeleting storage account {account_name!r}...")
                client.storage_accounts.delete(
                    resource_group_name=resource_group_name,
                    account_name=account_name,
                )
                cleanup_required = False
                print("Storage account deleted.")
                return 0
            finally:
                if cleanup_required:
                    print(
                        f"Cleaning up storage account {account_name!r} after failure...",
                        file=sys.stderr,
                    )
                    try:
                        client.storage_accounts.delete(
                            resource_group_name=resource_group_name,
                            account_name=account_name,
                        )
                    except HttpResponseError as cleanup_error:
                        print(
                            "Cleanup failed: "
                            f"{describe_http_error(cleanup_error)}",
                            file=sys.stderr,
                        )
                client.close()
    except CredentialUnavailableError as error:
        print(f"No usable Azure credential was found: {error}", file=sys.stderr)
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error.message}", file=sys.stderr)
    except HttpResponseError as error:
        print(f"Azure request failed: {describe_http_error(error)}", file=sys.stderr)
    except AzureError as error:
        print(f"Azure SDK error: {error}", file=sys.stderr)

    return 1


if __name__ == "__main__":
    raise SystemExit(manage_storage_account())
