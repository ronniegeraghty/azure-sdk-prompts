"""Create, inspect, tag, and optionally delete an Azure resource group."""

import argparse
import logging
import os
from collections.abc import Sequence

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.mgmt.resource import ResourceManagementClient
from azure.mgmt.resource.resources.models import ResourceGroup

LOGGER = logging.getLogger("resource-group-manager")


def parse_tag(value: str) -> tuple[str, str]:
    """Parse a KEY=VALUE command-line tag."""
    key, separator, tag_value = value.partition("=")
    if not separator or not key.strip():
        raise argparse.ArgumentTypeError("tags must use the format KEY=VALUE")
    return key.strip(), tag_value


def create_resource_group(
    client: ResourceManagementClient, name: str, location: str
) -> ResourceGroup:
    LOGGER.info("Creating or updating resource group %s in %s", name, location)
    return client.resource_groups.create_or_update(name, {"location": location})


def list_resource_groups(client: ResourceManagementClient) -> None:
    LOGGER.info("Resource groups in the subscription:")
    found = False
    for group in client.resource_groups.list():
        found = True
        LOGGER.info("  %-40s %s", group.name, group.location)
    if not found:
        LOGGER.info("  No resource groups found")


def get_resource_group(
    client: ResourceManagementClient, name: str
) -> ResourceGroup:
    group = client.resource_groups.get(name)
    LOGGER.info(
        "Resource group details: name=%s location=%s id=%s tags=%s",
        group.name,
        group.location,
        group.id,
        group.tags or {},
    )
    return group


def add_tags(
    client: ResourceManagementClient,
    name: str,
    tags_to_add: dict[str, str],
) -> ResourceGroup:
    current = client.resource_groups.get(name)
    merged_tags = {**(current.tags or {}), **tags_to_add}
    updated = client.resource_groups.update(name, {"tags": merged_tags})
    LOGGER.info("Updated tags for %s: %s", name, updated.tags or {})
    return updated


def delete_resource_group(client: ResourceManagementClient, name: str) -> None:
    LOGGER.info("Deleting resource group %s and all resources it contains", name)
    client.resource_groups.begin_delete(name).result()
    LOGGER.info("Deleted resource group %s", name)


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description=(
            "Create, list, inspect, and tag an Azure resource group. "
            "Deletion only occurs when --delete is supplied."
        )
    )
    parser.add_argument("--name", required=True, help="Resource group name")
    parser.add_argument(
        "--location",
        default="eastus",
        help="Azure location used when creating the group (default: eastus)",
    )
    parser.add_argument(
        "--tag",
        action="append",
        type=parse_tag,
        default=[],
        metavar="KEY=VALUE",
        help="Tag to merge into the resource group; may be repeated",
    )
    parser.add_argument(
        "--delete",
        action="store_true",
        help="Delete the resource group after the other operations complete",
    )
    return parser


def run(argv: Sequence[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    subscription_id = os.getenv("AZURE_SUBSCRIPTION_ID")
    if not subscription_id:
        LOGGER.error("AZURE_SUBSCRIPTION_ID is not set")
        return 2

    tags = dict(args.tag) or {"managed-by": "python-sdk"}

    try:
        with DefaultAzureCredential() as credential:
            with ResourceManagementClient(credential, subscription_id) as client:
                created = create_resource_group(client, args.name, args.location)
                LOGGER.info("Resource group ready: %s", created.id)
                list_resource_groups(client)
                get_resource_group(client, args.name)
                add_tags(client, args.name, tags)

                if args.delete:
                    delete_resource_group(client, args.name)
                else:
                    LOGGER.info(
                        "Resource group retained; pass --delete to remove it"
                    )
    except CredentialUnavailableError:
        LOGGER.exception(
            "No Azure credential is available; configure managed identity, "
            "developer-tool authentication, or service-principal environment variables"
        )
        return 1
    except ClientAuthenticationError:
        LOGGER.exception("Azure authentication failed")
        return 1
    except ResourceNotFoundError:
        LOGGER.exception("The requested resource group was not found")
        return 1
    except AzureError:
        LOGGER.exception("An Azure Resource Manager operation failed")
        return 1

    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    raise SystemExit(run())
