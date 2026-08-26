"""Manage an Azure Resource Group through the Azure management-plane SDK."""

from __future__ import annotations

import argparse
import json
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable


def parse_tag(value: str) -> tuple[str, str]:
    """Parse a KEY=VALUE command-line tag."""
    key, separator, tag_value = value.partition("=")
    if not separator or not key.strip():
        raise argparse.ArgumentTypeError("tags must use the format KEY=VALUE")
    return key.strip(), tag_value


def print_resource_group(resource_group: ResourceGroup) -> None:
    """Print selected resource-group details as JSON."""
    print(
        json.dumps(
            {
                "id": resource_group.id,
                "name": resource_group.name,
                "location": resource_group.location,
                "provisioning_state": resource_group.properties.provisioning_state
                if resource_group.properties
                else None,
                "tags": resource_group.tags or {},
            },
            indent=2,
            sort_keys=True,
        )
    )


def create_resource_group(
    client: ResourceManagementClient,
    name: str,
    location: str,
) -> ResourceGroup:
    resource_group = client.resource_groups.create_or_update(
        name,
        ResourceGroup(location=location),
    )
    print(f"Created resource group '{name}'.")
    return resource_group


def list_resource_groups(client: ResourceManagementClient) -> None:
    print("Resource groups in the subscription:")
    found = False
    for resource_group in client.resource_groups.list():
        found = True
        print(f"- {resource_group.name} ({resource_group.location})")
    if not found:
        print("- none")


def get_resource_group(
    client: ResourceManagementClient,
    name: str,
) -> ResourceGroup:
    resource_group = client.resource_groups.get(name)
    print(f"Details for resource group '{name}':")
    print_resource_group(resource_group)
    return resource_group


def add_tag(
    client: ResourceManagementClient,
    name: str,
    key: str,
    value: str,
) -> ResourceGroup:
    resource_group = client.resource_groups.get(name)
    tags = dict(resource_group.tags or {})
    tags[key] = value
    updated = client.resource_groups.update(
        name,
        ResourceGroupPatchable(tags=tags),
    )
    print(f"Added tag '{key}={value}' to resource group '{name}'.")
    print_resource_group(updated)
    return updated


def delete_resource_group(
    client: ResourceManagementClient,
    name: str,
) -> None:
    client.resource_groups.begin_delete(name).result()
    print(f"Deleted resource group '{name}'.")


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description=(
            "Create, list, inspect, tag, and delete an Azure Resource Group. "
            "No Azure request is made unless --execute is supplied."
        )
    )
    parser.add_argument(
        "--subscription-id",
        required=True,
        help="Azure subscription ID.",
    )
    parser.add_argument(
        "--resource-group",
        required=True,
        help="Resource group name.",
    )
    parser.add_argument(
        "--location",
        default="eastus",
        help="Azure region used when creating the group (default: eastus).",
    )
    parser.add_argument(
        "--tag",
        type=parse_tag,
        default=("managed-by", "python-sdk"),
        metavar="KEY=VALUE",
        help="Tag to add (default: managed-by=python-sdk).",
    )
    parser.add_argument(
        "--execute",
        action="store_true",
        help="Perform the Azure operations. Without this flag, only the plan is shown.",
    )
    parser.add_argument(
        "--confirm-delete",
        action="store_true",
        help="Delete the resource group after the other operations complete.",
    )
    return parser


def run(args: argparse.Namespace) -> int:
    tag_key, tag_value = args.tag
    if not args.execute:
        print("Dry run; no Azure requests were made.")
        print(
            f"Would create '{args.resource_group}' in '{args.location}', list all "
            f"resource groups, get its details, and add tag "
            f"'{tag_key}={tag_value}'."
        )
        if args.confirm_delete:
            print(f"Would then delete '{args.resource_group}'.")
        else:
            print("Would keep the resource group because --confirm-delete was omitted.")
        return 0

    credential = DefaultAzureCredential()
    client = ResourceManagementClient(credential, args.subscription_id)
    try:
        create_resource_group(client, args.resource_group, args.location)
        list_resource_groups(client)
        get_resource_group(client, args.resource_group)
        add_tag(client, args.resource_group, tag_key, tag_value)

        if args.confirm_delete:
            delete_resource_group(client, args.resource_group)
        else:
            print(
                "Resource group was not deleted. Supply --confirm-delete to delete it."
            )
        return 0
    except ResourceNotFoundError as error:
        print(f"Azure resource was not found: {error.message}", file=sys.stderr)
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error.message}", file=sys.stderr)
    except HttpResponseError as error:
        status = f" (HTTP {error.status_code})" if error.status_code else ""
        print(f"Azure request failed{status}: {error.message}", file=sys.stderr)
    except KeyboardInterrupt:
        print("Operation cancelled.", file=sys.stderr)
    finally:
        client.close()
        credential.close()
    return 1


def main(argv: Sequence[str] | None = None) -> int:
    return run(build_parser().parse_args(argv))


if __name__ == "__main__":
    raise SystemExit(main())
