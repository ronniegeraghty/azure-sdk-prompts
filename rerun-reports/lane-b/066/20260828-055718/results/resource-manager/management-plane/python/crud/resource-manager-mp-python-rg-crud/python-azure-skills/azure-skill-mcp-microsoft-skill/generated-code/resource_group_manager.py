"""Manage Azure Resource Groups with the Azure management plane SDK."""

from __future__ import annotations

import argparse
import json
import os
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.resource.resources import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable


def print_resource_group(resource_group: ResourceGroup) -> None:
    """Print a resource group as formatted JSON."""
    print(json.dumps(resource_group.as_dict(), indent=2, sort_keys=True, default=str))


def create_resource_group(
    client: ResourceManagementClient,
    name: str,
    location: str,
) -> ResourceGroup:
    resource_group = client.resource_groups.create_or_update(
        name,
        ResourceGroup(location=location),
    )
    print(f"Created or updated resource group '{name}'.")
    return resource_group


def list_resource_groups(client: ResourceManagementClient) -> None:
    groups = list(client.resource_groups.list())
    print(json.dumps([group.as_dict() for group in groups], indent=2, default=str))
    print(f"Found {len(groups)} resource group(s).", file=sys.stderr)


def get_resource_group(
    client: ResourceManagementClient,
    name: str,
) -> ResourceGroup:
    resource_group = client.resource_groups.get(name)
    print_resource_group(resource_group)
    return resource_group


def add_resource_group_tag(
    client: ResourceManagementClient,
    name: str,
    key: str,
    value: str,
) -> ResourceGroup:
    current = client.resource_groups.get(name)
    tags = dict(current.tags or {})
    tags[key] = value

    updated = client.resource_groups.update(
        name,
        ResourceGroupPatchable(tags=tags),
    )
    print(f"Set tag '{key}={value}' on resource group '{name}'.")
    print_resource_group(updated)
    return updated


def delete_resource_group(client: ResourceManagementClient, name: str) -> None:
    poller = client.resource_groups.begin_delete(name)
    poller.result()
    print(f"Deleted resource group '{name}'.")


def run_workflow(
    client: ResourceManagementClient,
    name: str,
    location: str,
    tag_key: str,
    tag_value: str,
) -> None:
    create_resource_group(client, name, location)
    list_resource_groups(client)
    get_resource_group(client, name)
    add_resource_group_tag(client, name, tag_key, tag_value)
    delete_resource_group(client, name)


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Create, list, inspect, tag, and delete Azure Resource Groups."
    )
    parser.add_argument(
        "--subscription-id",
        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
    )
    subparsers = parser.add_subparsers(dest="command", required=True)

    create_parser = subparsers.add_parser("create", help="Create a resource group.")
    create_parser.add_argument("--name", required=True)
    create_parser.add_argument("--location", required=True)

    subparsers.add_parser("list", help="List all resource groups.")

    get_parser = subparsers.add_parser("get", help="Get a resource group.")
    get_parser.add_argument("--name", required=True)

    tag_parser = subparsers.add_parser("tag", help="Add or replace a tag.")
    tag_parser.add_argument("--name", required=True)
    tag_parser.add_argument("--key", required=True)
    tag_parser.add_argument("--value", required=True)

    delete_parser = subparsers.add_parser("delete", help="Delete a resource group.")
    delete_parser.add_argument("--name", required=True)
    delete_parser.add_argument(
        "--yes",
        action="store_true",
        help="Confirm deletion of the resource group and all resources it contains.",
    )

    workflow_parser = subparsers.add_parser(
        "workflow",
        help="Create, list, get, tag, and delete a resource group.",
    )
    workflow_parser.add_argument("--name", required=True)
    workflow_parser.add_argument("--location", required=True)
    workflow_parser.add_argument("--tag-key", default="managed-by")
    workflow_parser.add_argument("--tag-value", default="python-sdk")
    workflow_parser.add_argument(
        "--yes",
        action="store_true",
        help="Confirm final deletion of the resource group.",
    )
    return parser


def execute(args: argparse.Namespace, client: ResourceManagementClient) -> None:
    if args.command == "create":
        print_resource_group(create_resource_group(client, args.name, args.location))
    elif args.command == "list":
        list_resource_groups(client)
    elif args.command == "get":
        get_resource_group(client, args.name)
    elif args.command == "tag":
        add_resource_group_tag(client, args.name, args.key, args.value)
    elif args.command == "delete":
        if not args.yes:
            raise ValueError("Deletion requires --yes.")
        delete_resource_group(client, args.name)
    elif args.command == "workflow":
        if not args.yes:
            raise ValueError("The workflow deletes the resource group; pass --yes.")
        run_workflow(
            client,
            args.name,
            args.location,
            args.tag_key,
            args.tag_value,
        )


def main(argv: Sequence[str] | None = None) -> int:
    parser = build_parser()
    args = parser.parse_args(argv)

    if not args.subscription_id:
        parser.error(
            "Provide --subscription-id or set the AZURE_SUBSCRIPTION_ID environment variable."
        )

    try:
        with DefaultAzureCredential() as credential:
            with ResourceManagementClient(
                credential,
                args.subscription_id,
            ) as client:
                execute(args, client)
        return 0
    except CredentialUnavailableError as error:
        print(
            "No Azure credential is available. Configure managed identity, service "
            "principal environment variables, or a supported developer login.",
            file=sys.stderr,
        )
        print(f"Details: {error}", file=sys.stderr)
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error}", file=sys.stderr)
    except ResourceNotFoundError as error:
        print(f"Resource group was not found: {error}", file=sys.stderr)
    except HttpResponseError as error:
        status = error.status_code if error.status_code is not None else "unknown"
        print(f"Azure request failed (HTTP {status}): {error}", file=sys.stderr)
    except ValueError as error:
        print(f"Invalid operation: {error}", file=sys.stderr)
    except AzureError as error:
        print(f"Azure SDK error: {error}", file=sys.stderr)
    except KeyboardInterrupt:
        print("Operation cancelled.", file=sys.stderr)

    return 1


if __name__ == "__main__":
    raise SystemExit(main())
