"""Manage an existing Azure resource group with the management-plane SDK."""

from __future__ import annotations

import argparse
import os
import sys

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroupPatchable


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="List, inspect, tag, and optionally delete an Azure resource group."
    )
    parser.add_argument(
        "--subscription-id",
        default=os.environ.get("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID. Defaults to AZURE_SUBSCRIPTION_ID.",
    )
    parser.add_argument(
        "--resource-group",
        required=True,
        help="Name of the existing resource group to inspect and tag.",
    )
    parser.add_argument("--tag-name", required=True, help="Tag name to add or replace.")
    parser.add_argument("--tag-value", required=True, help="Value for the tag.")
    parser.add_argument(
        "--delete",
        action="store_true",
        help="Delete the resource group after updating its tag.",
    )
    parser.add_argument(
        "--confirm-delete",
        metavar="RESOURCE_GROUP_NAME",
        help="Required with --delete; must exactly match --resource-group.",
    )
    args = parser.parse_args()

    if not args.subscription_id:
        parser.error(
            "--subscription-id is required when AZURE_SUBSCRIPTION_ID is not set"
        )
    if args.delete and args.confirm_delete != args.resource_group:
        parser.error(
            "--confirm-delete must exactly match --resource-group when --delete is used"
        )
    if args.confirm_delete and not args.delete:
        parser.error("--confirm-delete can only be used with --delete")

    return args


def print_resource_group(prefix: str, resource_group: object) -> None:
    print(
        f"{prefix}: name={resource_group.name}, "
        f"location={resource_group.location}, "
        f"provisioning_state={resource_group.properties.provisioning_state}, "
        f"tags={resource_group.tags or {}}"
    )


def manage_resource_group(args: argparse.Namespace) -> None:
    credential = DefaultAzureCredential()
    client = ResourceManagementClient(credential, args.subscription_id)

    print("Resource groups in subscription:")
    for resource_group in client.resource_groups.list():
        print_resource_group("-", resource_group)

    resource_group = client.resource_groups.get(args.resource_group)
    print_resource_group("Selected resource group", resource_group)

    tags = dict(resource_group.tags or {})
    tags[args.tag_name] = args.tag_value
    updated = client.resource_groups.update(
        args.resource_group,
        ResourceGroupPatchable(tags=tags),
    )
    print_resource_group("Updated resource group", updated)

    if args.delete:
        print(f"Deleting resource group '{args.resource_group}'...")
        client.resource_groups.begin_delete(args.resource_group).result()
        print(f"Deleted resource group '{args.resource_group}'.")
    else:
        print("Deletion skipped. Pass --delete with --confirm-delete to delete it.")


def main() -> int:
    args = parse_args()

    try:
        manage_resource_group(args)
    except ClientAuthenticationError as exc:
        print(f"Azure authentication failed: {exc}", file=sys.stderr)
        return 2
    except ResourceNotFoundError as exc:
        print(
            f"Resource group '{args.resource_group}' was not found: {exc}",
            file=sys.stderr,
        )
        return 3
    except HttpResponseError as exc:
        status = exc.status_code if exc.status_code is not None else "unknown"
        print(
            f"Azure Resource Manager request failed (HTTP {status}): {exc}",
            file=sys.stderr,
        )
        return 4
    except (OSError, ValueError) as exc:
        print(f"Configuration error: {exc}", file=sys.stderr)
        return 5

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
