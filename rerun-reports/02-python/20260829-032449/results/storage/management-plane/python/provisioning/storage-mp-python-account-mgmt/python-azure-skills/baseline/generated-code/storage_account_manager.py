"""Offline example of the Azure Storage management-plane account lifecycle."""

from __future__ import annotations

import argparse
import re
import sys
from collections.abc import Iterable
from dataclasses import dataclass
from typing import Protocol

from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient
from azure.mgmt.storage.models import (
    BlobServiceProperties,
    Sku,
    SkuName,
    StorageAccount,
    StorageAccountCreateParameters,
    StorageAccountUpdateParameters,
)


class Poller(Protocol):
    def result(self) -> StorageAccount | None: ...


class StorageAccountsOperations(Protocol):
    def begin_create(
        self,
        resource_group_name: str,
        account_name: str,
        parameters: StorageAccountCreateParameters,
    ) -> Poller: ...

    def list_by_resource_group(
        self, resource_group_name: str
    ) -> Iterable[StorageAccount]: ...

    def get_properties(
        self, resource_group_name: str, account_name: str
    ) -> StorageAccount: ...

    def begin_update(
        self,
        resource_group_name: str,
        account_name: str,
        parameters: StorageAccountUpdateParameters,
    ) -> Poller: ...

    def begin_delete(self, resource_group_name: str, account_name: str) -> Poller: ...


class BlobServicesOperations(Protocol):
    def set_service_properties(
        self,
        resource_group_name: str,
        account_name: str,
        blob_services_name: str,
        parameters: BlobServiceProperties,
    ) -> BlobServiceProperties: ...


class StorageClient(Protocol):
    storage_accounts: StorageAccountsOperations
    blob_services: BlobServicesOperations


@dataclass
class _CompletedPoller:
    value: StorageAccount | None = None

    def result(self) -> StorageAccount | None:
        return self.value


class _OfflineStorageAccounts:
    def __init__(self) -> None:
        self._accounts: dict[tuple[str, str], StorageAccount] = {}

    def begin_create(
        self,
        resource_group_name: str,
        account_name: str,
        parameters: StorageAccountCreateParameters,
    ) -> _CompletedPoller:
        key = (resource_group_name, account_name)
        if key in self._accounts:
            raise ValueError(f"Storage account '{account_name}' already exists")

        account = StorageAccount.deserialize(
            {
                "id": (
                "/subscriptions/00000000-0000-0000-0000-000000000000"
                f"/resourceGroups/{resource_group_name}"
                f"/providers/Microsoft.Storage/storageAccounts/{account_name}"
                ),
                "name": account_name,
                "location": parameters.location,
                "sku": {"name": parameters.sku.name},
            }
        )
        self._accounts[key] = account
        return _CompletedPoller(account)

    def list_by_resource_group(
        self, resource_group_name: str
    ) -> Iterable[StorageAccount]:
        return [
            account
            for (group, _), account in self._accounts.items()
            if group == resource_group_name
        ]

    def get_properties(
        self, resource_group_name: str, account_name: str
    ) -> StorageAccount:
        try:
            return self._accounts[(resource_group_name, account_name)]
        except KeyError as exc:
            raise ValueError(f"Storage account '{account_name}' was not found") from exc

    def begin_update(
        self,
        resource_group_name: str,
        account_name: str,
        parameters: StorageAccountUpdateParameters,
    ) -> _CompletedPoller:
        account = self.get_properties(resource_group_name, account_name)
        if parameters.tags is not None:
            account.tags = parameters.tags
        return _CompletedPoller(account)

    def begin_delete(
        self, resource_group_name: str, account_name: str
    ) -> _CompletedPoller:
        try:
            del self._accounts[(resource_group_name, account_name)]
        except KeyError as exc:
            raise ValueError(f"Storage account '{account_name}' was not found") from exc
        return _CompletedPoller()


class _OfflineBlobServices:
    def __init__(self, accounts: _OfflineStorageAccounts) -> None:
        self._accounts = accounts
        self._properties: dict[tuple[str, str], BlobServiceProperties] = {}

    def set_service_properties(
        self,
        resource_group_name: str,
        account_name: str,
        blob_services_name: str,
        parameters: BlobServiceProperties,
    ) -> BlobServiceProperties:
        self._accounts.get_properties(resource_group_name, account_name)
        if blob_services_name != "default":
            raise ValueError("The blob service name must be 'default'")
        self._properties[(resource_group_name, account_name)] = parameters
        return parameters


class OfflineStorageManagementClient:
    """In-memory client with the subset of StorageManagementClient used below."""

    def __init__(self) -> None:
        self.storage_accounts = _OfflineStorageAccounts()
        self.blob_services = _OfflineBlobServices(self.storage_accounts)


def validate_account_name(name: str) -> None:
    if not re.fullmatch(r"[a-z0-9]{3,24}", name):
        raise ValueError(
            "Storage account name must contain 3-24 lowercase letters or digits"
        )


def manage_storage_account(
    client: StorageClient, resource_group: str, account_name: str
) -> None:
    print(f"Creating '{account_name}' in eastus with Standard_LRS...")
    created = client.storage_accounts.begin_create(
        resource_group,
        account_name,
        StorageAccountCreateParameters(
            sku=Sku(name=SkuName.STANDARD_LRS),
            kind="StorageV2",
            location="eastus",
        ),
    ).result()
    if created is None:
        raise RuntimeError("The create operation returned no storage account")

    print(f"Storage accounts in resource group '{resource_group}':")
    for account in client.storage_accounts.list_by_resource_group(resource_group):
        print(f"  - {account.name} ({account.location})")

    properties = client.storage_accounts.get_properties(
        resource_group, account_name
    )
    sku_name = properties.sku.name if properties.sku else "unknown"
    print(
        f"Properties: id={properties.id}, location={properties.location}, sku={sku_name}"
    )

    # Account metadata updates use begin_update; blob versioning is configured
    # separately on the account's default blob service.
    client.storage_accounts.begin_update(
        resource_group,
        account_name,
        StorageAccountUpdateParameters(tags={"managed-by": "python-sdk-example"}),
    ).result()
    blob_properties = client.blob_services.set_service_properties(
        resource_group,
        account_name,
        "default",
        BlobServiceProperties(is_versioning_enabled=True),
    )
    print(f"Blob versioning enabled: {blob_properties.is_versioning_enabled}")

    client.storage_accounts.begin_delete(resource_group, account_name).result()
    print(f"Deleted '{account_name}'")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run an offline Azure Storage management-plane lifecycle example."
    )
    parser.add_argument("--resource-group", default="example-resource-group")
    parser.add_argument("--account-name", default="examplestorageacct123")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    credential: DefaultAzureCredential | None = None
    try:
        validate_account_name(args.account_name)

        # DefaultAzureCredential is constructed to demonstrate the requested
        # authentication setup. Offline mode never asks it for a token.
        credential = DefaultAzureCredential()
        client: StorageClient = OfflineStorageManagementClient()
        manage_storage_account(client, args.resource_group, args.account_name)
        return 0
    except ClientAuthenticationError as exc:
        print(f"Azure authentication failed: {exc}", file=sys.stderr)
    except HttpResponseError as exc:
        status = exc.status_code if exc.status_code is not None else "unknown"
        print(f"Azure request failed (HTTP {status}): {exc.message}", file=sys.stderr)
    except AzureError as exc:
        print(f"Azure SDK error: {exc}", file=sys.stderr)
    except (RuntimeError, ValueError) as exc:
        print(f"Operation failed: {exc}", file=sys.stderr)
    finally:
        if credential is not None:
            credential.close()
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
