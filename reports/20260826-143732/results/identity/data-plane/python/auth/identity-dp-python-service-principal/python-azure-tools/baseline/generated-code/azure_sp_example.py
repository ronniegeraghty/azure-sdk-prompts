"""Create Azure SDK clients with a service-principal client secret."""

from __future__ import annotations

import argparse
import os
import sys
from dataclasses import dataclass

from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError
from azure.identity import ClientSecretCredential, CredentialUnavailableError
from azure.mgmt.resource import ResourceManagementClient
from dotenv import load_dotenv


class ConfigurationError(ValueError):
    """Raised when required Azure configuration is absent."""


@dataclass(frozen=True)
class AzureSettings:
    tenant_id: str
    client_id: str
    client_secret: str
    subscription_id: str

    @classmethod
    def from_environment(cls) -> AzureSettings:
        load_dotenv(override=False)

        names = (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_CLIENT_SECRET",
            "AZURE_SUBSCRIPTION_ID",
        )
        values = {name: os.environ.get(name, "") for name in names}
        missing = [name for name, value in values.items() if not value.strip()]
        if missing:
            raise ConfigurationError(
                "Missing required environment variables: " + ", ".join(missing)
            )

        return cls(
            tenant_id=values["AZURE_TENANT_ID"].strip(),
            client_id=values["AZURE_CLIENT_ID"].strip(),
            client_secret=values["AZURE_CLIENT_SECRET"],
            subscription_id=values["AZURE_SUBSCRIPTION_ID"].strip(),
        )


def create_credential(settings: AzureSettings) -> ClientSecretCredential:
    return ClientSecretCredential(
        tenant_id=settings.tenant_id,
        client_id=settings.client_id,
        client_secret=settings.client_secret,
    )


def create_resource_client(
    settings: AzureSettings, credential: ClientSecretCredential
) -> ResourceManagementClient:
    return ResourceManagementClient(
        credential=credential,
        subscription_id=settings.subscription_id,
    )


def list_resource_groups(client: ResourceManagementClient) -> None:
    """Perform a read-only SDK call that causes the credential to authenticate."""
    print("Resource groups:")
    found = False
    for resource_group in client.resource_groups.list():
        found = True
        print(f"- {resource_group.name}")
    if not found:
        print("(none)")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Configure an Azure SDK client with a service principal."
    )
    parser.add_argument(
        "--list-resource-groups",
        action="store_true",
        help="authenticate and perform a read-only Azure Resource Manager request",
    )
    return parser.parse_args()


def run() -> int:
    args = parse_args()
    credential: ClientSecretCredential | None = None
    client: ResourceManagementClient | None = None

    try:
        settings = AzureSettings.from_environment()
        credential = create_credential(settings)
        client = create_resource_client(settings, credential)

        if args.list_resource_groups:
            list_resource_groups(client)
        else:
            print(
                "Azure credential and ResourceManagementClient configured. "
                "No network request was made."
            )
        return 0
    except ConfigurationError as exc:
        print(f"Configuration error: {exc}", file=sys.stderr)
        return 2
    except CredentialUnavailableError as exc:
        print(f"Azure credential unavailable: {exc}", file=sys.stderr)
        return 3
    except ClientAuthenticationError as exc:
        print(
            "Azure authentication failed. Verify the tenant ID, client ID, "
            f"client secret, and secret expiration. Details: {exc}",
            file=sys.stderr,
        )
        return 3
    except HttpResponseError as exc:
        print(
            f"Azure request failed with status {exc.status_code}: {exc.message}",
            file=sys.stderr,
        )
        return 4
    except AzureError as exc:
        print(f"Azure SDK error: {exc}", file=sys.stderr)
        return 4
    finally:
        if client is not None:
            client.close()
        if credential is not None:
            credential.close()


if __name__ == "__main__":
    raise SystemExit(run())
