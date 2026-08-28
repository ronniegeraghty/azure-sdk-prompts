"""Create, inspect, tag, and delete an Azure resource group.

Authentication uses DefaultAzureCredential. Set AZURE_SUBSCRIPTION_ID,
AZURE_RESOURCE_GROUP, and AZURE_LOCATION, or pass their command-line options.
"""

from __future__ import annotations

import argparse
import logging
import os
import sys
from collections.abc import Sequence

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.mgmt.resource.resources import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable

LOGGER = logging.getLogger("resource-group-manager")


def create_resource_group(
    client: ResourceManagementClient, name: str, location: str
) -> ResourceGroup:
    """Create a resource group, refusing to overwrite an existing group."""
    try:
        existing = client.resource_groups.get(name)
    except ResourceNotFoundError:
        existing = None

    if existing is not None:
        raise ValueError(
            f"Resource group '{name}' already exists; refusing to modify or delete it."
        )

    resource_group = client.resource_groups.create_or_update(
        name, ResourceGroup(location=location)
    )
    LOGGER.info("Created resource group '%s' in '%s'.", name, location)
    return resource_group


def list_resource_groups(client: ResourceManagementClient) -> None:
    """List every resource group visible in the subscription."""
    print("\nResource groups in the subscription:")
    found = False
    for resource_group in client.resource_groups.list():
        found = True
        print(
            f"- {resource_group.name} "
            f"(location={resource_group.location}, tags={resource_group.tags or {}})"
        )
    if not found:
        print("- None")


def get_resource_group(
    client: ResourceManagementClient, name: str
) -> ResourceGroup:
    """Get and display the requested resource group."""
    resource_group = client.resource_groups.get(name)
    print("\nCreated resource group details:")
    print(f"  id: {resource_group.id}")
    print(f"  name: {resource_group.name}")
    print(f"  location: {resource_group.location}")
    print(f"  tags: {resource_group.tags or {}}")
    return resource_group


def add_tag(
    client: ResourceManagementClient, name: str, tag_key: str, tag_value: str
) -> ResourceGroup:
    """Add or replace one tag while preserving all other tags."""
    resource_group = client.resource_groups.get(name)
    tags = dict(resource_group.tags or {})
    tags[tag_key] = tag_value

    updated = client.resource_groups.update(
        name, ResourceGroupPatchable(tags=tags)
    )
    LOGGER.info(
        "Set tag '%s=%s' on resource group '%s'.", tag_key, tag_value, name
    )
    print(f"\nUpdated tags: {updated.tags or {}}")
    return updated


def delete_resource_group(client: ResourceManagementClient, name: str) -> None:
    """Delete a resource group and wait for the operation to finish."""
    LOGGER.info("Deleting resource group '%s'...", name)
    client.resource_groups.begin_delete(name).result()
    LOGGER.info("Deleted resource group '%s'.", name)


def parse_args(argv: Sequence[str] | None = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--subscription-id",
        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
        help="Azure subscription ID (or set AZURE_SUBSCRIPTION_ID).",
    )
    parser.add_argument(
        "--resource-group",
        default=os.getenv("AZURE_RESOURCE_GROUP"),
        help="New resource group name (or set AZURE_RESOURCE_GROUP).",
    )
    parser.add_argument(
        "--location",
        default=os.getenv("AZURE_LOCATION"),
        help="Azure region, such as eastus (or set AZURE_LOCATION).",
    )
    parser.add_argument("--tag-key", default="Environment")
    parser.add_argument("--tag-value", default="Demo")
    parser.add_argument(
        "--keep-on-error",
        action="store_true",
        help="Do not delete the newly created group if a later operation fails.",
    )
    args = parser.parse_args(argv)

    missing = [
        option
        for option, value in (
            ("--subscription-id", args.subscription_id),
            ("--resource-group", args.resource_group),
            ("--location", args.location),
        )
        if not value
    ]
    if missing:
        parser.error(f"missing required configuration: {', '.join(missing)}")
    return args


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv)
    credential = DefaultAzureCredential()
    client = ResourceManagementClient(credential, args.subscription_id)
    created = False

    try:
        create_resource_group(client, args.resource_group, args.location)
        created = True
        list_resource_groups(client)
        get_resource_group(client, args.resource_group)
        add_tag(
            client,
            args.resource_group,
            args.tag_key,
            args.tag_value,
        )
        delete_resource_group(client, args.resource_group)
        created = False
        return 0
    except ClientAuthenticationError as error:
        LOGGER.error("Azure authentication failed: %s", error.message)
    except ResourceNotFoundError as error:
        LOGGER.error("The requested Azure resource was not found: %s", error.message)
    except HttpResponseError as error:
        status = error.status_code if error.status_code is not None else "unknown"
        LOGGER.error("Azure request failed (HTTP %s): %s", status, error.message)
    except ValueError as error:
        LOGGER.error("%s", error)
    except AzureError as error:
        LOGGER.error("Azure SDK operation failed: %s", error)
    finally:
        if created and not args.keep_on_error:
            try:
                LOGGER.warning("Cleaning up the resource group after an error.")
                delete_resource_group(client, args.resource_group)
            except AzureError as cleanup_error:
                LOGGER.error("Cleanup failed: %s", cleanup_error)
        client.close()
        credential.close()

    return 1


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(message)s",
    )
    sys.exit(main())
