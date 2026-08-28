#!/usr/bin/env python3
"""Create, inspect, tag, and delete an Azure resource group."""

from __future__ import annotations

import argparse
import json
import os
import sys
from typing import Any

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Manage an Azure resource group with the management-plane SDK."
    )
    parser.add_argument(
        "--subscription-id",
        default=os.environ.get("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
    )
    parser.add_argument("--name", required=True, help="Resource group name.")
    parser.add_argument(
        "--location",
        default="eastus",
        help="Azure region used when creating the resource group (default: eastus).",
    )
    parser.add_argument(
        "--tag-key",
        default="managed-by",
        help="Tag key to add (default: managed-by).",
    )
    parser.add_argument(
        "--tag-value",
        default="python-sdk",
        help="Tag value to add (default: python-sdk).",
    )
    parser.add_argument(
        "--confirm-delete",
        action="store_true",
        help="Delete the resource group after the other operations complete.",
    )
    return parser.parse_args()


def resource_group_to_dict(resource_group: Any) -> dict[str, Any]:
    return {
        "id": resource_group.id,
        "name": resource_group.name,
        "location": resource_group.location,
        "managed_by": resource_group.managed_by,
        "properties": (
            resource_group.properties.as_dict()
            if resource_group.properties is not None
            else None
        ),
        "tags": resource_group.tags or {},
    }


def print_resource_group(resource_group: Any) -> None:
    print(json.dumps(resource_group_to_dict(resource_group), indent=2, default=str))


def run_lifecycle(
    client: ResourceManagementClient,
    resource_group_name: str,
    location: str,
    tag_key: str,
    tag_value: str,
    confirm_delete: bool,
) -> None:
    if client.resource_groups.check_existence(resource_group_name):
        raise ValueError(
            f"Resource group '{resource_group_name}' already exists. "
            "Choose a new name to avoid modifying or deleting an existing group."
        )

    print(f"Creating resource group '{resource_group_name}' in '{location}'...")
    created = client.resource_groups.create_or_update(
        resource_group_name,
        {"location": location},
    )
    print_resource_group(created)

    print("\nResource groups in the subscription:")
    resource_groups = list(client.resource_groups.list())
    if not resource_groups:
        print("(none)")
    for resource_group in resource_groups:
        print(f"- {resource_group.name} ({resource_group.location})")

    print(f"\nDetails for '{resource_group_name}':")
    current = client.resource_groups.get(resource_group_name)
    print_resource_group(current)

    updated_tags = dict(current.tags or {})
    updated_tags[tag_key] = tag_value
    print(f"\nAdding tag '{tag_key}={tag_value}'...")
    updated = client.resource_groups.update(
        resource_group_name,
        {"tags": updated_tags},
    )
    print_resource_group(updated)

    if not confirm_delete:
        print(
            "\nDeletion skipped. Re-run with --confirm-delete to delete "
            f"'{resource_group_name}'."
        )
        return

    print(f"\nDeleting resource group '{resource_group_name}'...")
    delete_operation = client.resource_groups.begin_delete(resource_group_name)
    delete_operation.result()
    print(f"Resource group '{resource_group_name}' was deleted.")


def main() -> int:
    args = parse_args()
    if not args.subscription_id:
        print(
            "Error: provide --subscription-id or set AZURE_SUBSCRIPTION_ID.",
            file=sys.stderr,
        )
        return 2

    credential = DefaultAzureCredential()
    client = ResourceManagementClient(credential, args.subscription_id)

    try:
        run_lifecycle(
            client=client,
            resource_group_name=args.name,
            location=args.location,
            tag_key=args.tag_key,
            tag_value=args.tag_value,
            confirm_delete=args.confirm_delete,
        )
    except ValueError as error:
        print(f"Error: {error}", file=sys.stderr)
        return 2
    except ClientAuthenticationError as error:
        print(
            "Authentication failed. Sign in with a supported credential or configure "
            f"service-principal environment variables. Details: {error}",
            file=sys.stderr,
        )
        return 3
    except ResourceNotFoundError as error:
        print(f"Resource group was not found: {error}", file=sys.stderr)
        return 4
    except HttpResponseError as error:
        status_code = error.status_code if error.status_code is not None else "unknown"
        print(
            f"Azure request failed (HTTP {status_code}): {error.message}",
            file=sys.stderr,
        )
        return 5
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
