"""Create, inspect, configure, and delete an Azure Storage account."""

from __future__ import annotations

import os
import re
import sys

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    Kind,
    Sku,
    SkuName,
    StorageAccountCheckNameAvailabilityParameters,
    StorageAccountCreateParameters,
)


LOCATION = "eastus"
ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")


def required_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def main() -> int:
    created = False
    deleted = False
    client: StorageManagementClient | None = None
    resource_group = ""
    account_name = ""

    try:
        subscription_id = required_environment_variable("AZURE_SUBSCRIPTION_ID")
        resource_group = required_environment_variable("AZURE_RESOURCE_GROUP")
        account_name = required_environment_variable("AZURE_STORAGE_ACCOUNT_NAME")

        if not ACCOUNT_NAME_PATTERN.fullmatch(account_name):
            raise ValueError(
                "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters "
                "and numbers."
            )

        credential = DefaultAzureCredential()
        client = StorageManagementClient(credential, subscription_id)

        availability = client.storage_accounts.check_name_availability(
            StorageAccountCheckNameAvailabilityParameters(name=account_name)
        )
        if not availability.name_available:
            reason = availability.message or availability.reason or "unknown reason"
            raise ValueError(
                f"Storage account name {account_name!r} is unavailable: {reason}"
            )

        print(f"Creating storage account {account_name!r} in {LOCATION}...")
        account = client.storage_accounts.begin_create(
            resource_group,
            account_name,
            StorageAccountCreateParameters(
                sku=Sku(name=SkuName.STANDARD_LRS),
                kind=Kind.STORAGE_V2,
                location=LOCATION,
                enable_https_traffic_only=True,
                minimum_tls_version="TLS1_2",
                allow_blob_public_access=False,
            ),
        ).result()
        created = True
        print(f"Created: {account.id}")

        print(f"\nStorage accounts in resource group {resource_group!r}:")
        for listed_account in client.storage_accounts.list_by_resource_group(
            resource_group
        ):
            print(f"- {listed_account.name} ({listed_account.location})")

        account = client.storage_accounts.get_properties(
            resource_group, account_name
        )
        print("\nCreated account properties:")
        print(f"  Name: {account.name}")
        print(f"  Location: {account.location}")
        print(f"  SKU: {account.sku.name}")
        print(f"  Kind: {account.kind}")
        print(f"  Provisioning state: {account.provisioning_state}")

        print("\nEnabling blob versioning...")
        client.blob_services.set_service_properties(
            resource_group,
            account_name,
            BlobServiceProperties(is_versioning_enabled=True),
        )
        blob_properties = client.blob_services.get_service_properties(
            resource_group, account_name
        )
        if blob_properties.is_versioning_enabled is not True:
            raise RuntimeError("Blob versioning was not enabled successfully.")
        print("Blob versioning enabled.")

        print(f"\nDeleting storage account {account_name!r}...")
        client.storage_accounts.delete(resource_group, account_name)
        deleted = True
        print("Storage account deleted.")
        return 0

    except (ValueError, RuntimeError) as error:
        print(f"Configuration or operation error: {error}", file=sys.stderr)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        print(f"Azure authentication failed: {error}", file=sys.stderr)
    except HttpResponseError as error:
        status = error.status_code if error.status_code is not None else "unknown"
        print(
            f"Azure request failed (HTTP {status}): {error.message}",
            file=sys.stderr,
        )
    except KeyboardInterrupt:
        print("Operation cancelled.", file=sys.stderr)
    finally:
        # Avoid leaving a billable account behind if a post-create step fails.
        if created and not deleted and client is not None:
            try:
                print(
                    f"Cleaning up storage account {account_name!r}...",
                    file=sys.stderr,
                )
                client.storage_accounts.delete(resource_group, account_name)
                print("Cleanup completed.", file=sys.stderr)
            except HttpResponseError as cleanup_error:
                print(
                    "Cleanup failed; delete the storage account manually: "
                    f"{cleanup_error.message}",
                    file=sys.stderr,
                )

    return 1


if __name__ == "__main__":
    raise SystemExit(main())
