"""Perform create, read, update, and delete/purge operations on a Key Vault secret."""

import os
import sys

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


INITIAL_VALUE = "initial-value"
UPDATED_VALUE = "updated-value"


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def run_secret_crud(vault_url: str, secret_name: str) -> None:
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=vault_url, credential=credential)

    try:
        created_secret = client.set_secret(secret_name, INITIAL_VALUE)
        print(f"Created secret {created_secret.name!r}.")

        retrieved_secret = client.get_secret(secret_name)
        print(f"Read secret {retrieved_secret.name!r}.")

        updated_secret = client.set_secret(secret_name, UPDATED_VALUE)
        print(f"Updated secret {updated_secret.name!r} to the requested value.")

        delete_poller = client.begin_delete_secret(secret_name)
        delete_poller.wait()
        print(f"Deleted secret {secret_name!r}.")

        client.purge_deleted_secret(secret_name)
        print(f"Purged secret {secret_name!r}.")
    finally:
        credential.close()


def main() -> int:
    try:
        vault_url = require_environment_variable("AZURE_KEY_VAULT_URL")
        secret_name = os.environ.get("AZURE_KEY_VAULT_SECRET_NAME", "crud-demo-secret")
        run_secret_crud(vault_url, secret_name)
        return 0
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error}", file=sys.stderr)
    except ResourceNotFoundError as error:
        print(f"Key Vault resource was not found: {error}", file=sys.stderr)
    except HttpResponseError as error:
        print(f"Azure Key Vault request failed: {error}", file=sys.stderr)
    except AzureError as error:
        print(f"Azure SDK operation failed: {error}", file=sys.stderr)
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
