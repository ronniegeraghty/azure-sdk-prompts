"""Manage Azure Resource Groups with the Azure management-plane SDK."""

from __future__ import annotations

import argparse
import json
import logging
import os
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
    ServiceRequestError,
    ServiceResponseError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroupPatchable

LOGGER = logging.getLogger("resource_groups")


def resource_group_to_dict(resource_group: object) -> dict[str, object]:
    """Return the useful Resource Group fields as JSON-serializable data."""
    return {
        "id": getattr(resource_group, "id", None),
        "name": getattr(resource_group, "name", None),
        "location": getattr(resource_group, "location", None),
        "managed_by": getattr(resource_group, "managed_by", None),
        "provisioning_state": getattr(
            getattr(resource_group, "properties", None),
            "provisioning_state",
            None,
        ),
        "tags": getattr(resource_group, "tags", None) or {},
    }


def list_resource_groups(client: ResourceManagementClient) -> None:
    groups = [
        resource_group_to_dict(group)
        for group in client.resource_groups.list()
    ]
    print(json.dumps(groups, indent=2, sort_keys=True))


def get_resource_group(
    client: ResourceManagementClient, resource_group_name: str
) -> None:
    group = client.resource_groups.get(resource_group_name)
    print(json.dumps(resource_group_to_dict(group), indent=2, sort_keys=True))


def add_tags(
    client: ResourceManagementClient,
    resource_group_name: str,
    tags_to_add: dict[str, str],
) -> None:
    group = client.resource_groups.get(resource_group_name)
    merged_tags = dict(group.tags or {})
    merged_tags.update(tags_to_add)

    updated_group = client.resource_groups.update(
        resource_group_name,
        ResourceGroupPatchable(tags=merged_tags),
    )
    print(json.dumps(resource_group_to_dict(updated_group), indent=2, sort_keys=True))


def delete_resource_group(
    client: ResourceManagementClient, resource_group_name: str
) -> None:
    LOGGER.info("Deleting resource group %s", resource_group_name)
    client.resource_groups.begin_delete(resource_group_name).result()
    print(f"Deleted resource group '{resource_group_name}'.")


def parse_tag(value: str) -> tuple[str, str]:
    key, separator, tag_value = value.partition("=")
    if not separator or not key.strip():
        raise argparse.ArgumentTypeError("tags must use the format KEY=VALUE")
    return key.strip(), tag_value


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Manage Azure Resource Groups with the Azure Python SDK."
    )
    parser.add_argument(
        "--subscription-id",
        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
    )
    parser.add_argument(
        "--verbose",
        action="store_true",
        help="Enable informational logging.",
    )

    subparsers = parser.add_subparsers(dest="command", required=True)
    subparsers.add_parser("list", help="List every Resource Group.")

    get_parser = subparsers.add_parser("get", help="Get one Resource Group.")
    get_parser.add_argument("resource_group_name")

    tag_parser = subparsers.add_parser(
        "tag",
        help="Add or replace tags without removing other tags.",
    )
    tag_parser.add_argument("resource_group_name")
    tag_parser.add_argument(
        "--tag",
        action="append",
        required=True,
        type=parse_tag,
        metavar="KEY=VALUE",
        help="Tag to add. Repeat this option to add multiple tags.",
    )

    delete_parser = subparsers.add_parser(
        "delete",
        help="Delete a Resource Group and every resource it contains.",
    )
    delete_parser.add_argument("resource_group_name")
    delete_parser.add_argument(
        "--yes",
        action="store_true",
        help="Confirm the destructive operation.",
    )
    return parser


def execute(args: argparse.Namespace) -> None:
    with DefaultAzureCredential() as credential:
        with ResourceManagementClient(
            credential=credential,
            subscription_id=args.subscription_id,
        ) as client:
            if args.command == "list":
                list_resource_groups(client)
            elif args.command == "get":
                get_resource_group(client, args.resource_group_name)
            elif args.command == "tag":
                add_tags(client, args.resource_group_name, dict(args.tag))
            elif args.command == "delete":
                delete_resource_group(client, args.resource_group_name)


def describe_http_error(error: HttpResponseError) -> str:
    status = f"HTTP {error.status_code}: " if error.status_code else ""
    request_id = None
    if error.response is not None:
        request_id = error.response.headers.get("x-ms-request-id")
    suffix = f" (request ID: {request_id})" if request_id else ""
    return f"{status}{error.message or str(error)}{suffix}"


def main(argv: Sequence[str] | None = None) -> int:
    parser = build_parser()
    args = parser.parse_args(argv)

    logging.basicConfig(
        level=logging.INFO if args.verbose else logging.WARNING,
        format="%(levelname)s: %(message)s",
    )

    if not args.subscription_id:
        parser.error(
            "set AZURE_SUBSCRIPTION_ID or pass --subscription-id before the command"
        )
    if args.command == "delete" and not args.yes:
        parser.error("delete requires --yes to confirm deletion")

    try:
        execute(args)
    except CredentialUnavailableError as error:
        LOGGER.error(
            "No Azure credential is available. Configure managed identity, "
            "workload identity, or supported local developer credentials: %s",
            error,
        )
        return 2
    except ClientAuthenticationError as error:
        LOGGER.error("Azure authentication failed: %s", error.message or error)
        return 2
    except ResourceNotFoundError as error:
        LOGGER.error("Resource Group not found: %s", describe_http_error(error))
        return 3
    except (ServiceRequestError, ServiceResponseError) as error:
        LOGGER.error("Could not communicate with Azure: %s", error)
        return 4
    except HttpResponseError as error:
        LOGGER.error("Azure Resource Manager request failed: %s", describe_http_error(error))
        return 5
    except KeyboardInterrupt:
        LOGGER.error("Operation cancelled.")
        return 130
    return 0


if __name__ == "__main__":
    sys.exit(main())
