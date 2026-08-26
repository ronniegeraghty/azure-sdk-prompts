"""Perform create, read, update, and delete operations on an Azure Key Vault secret."""

import os
import sys

from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.identity import DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


SECRET_NAME = os.getenv("AZURE_KEY_VAULT_SECRET_NAME", "crud-demo-secret")
INITIAL_VALUE = "initial-value"
UPDATED_VALUE = "updated-value"


def require_vault_url() -> str:
    vault_url = os.getenv("AZURE_KEY_VAULT_URL")
    if not vault_url:
        raise ValueError(
            "AZURE_KEY_VAULT_URL is required (for example, "
            "https://your-vault-name.vault.azure.net)."
        )
    return vault_url


def run_crud_operations() -> None:
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=require_vault_url(), credential=credential)

    try:
        created_secret = client.set_secret(SECRET_NAME, INITIAL_VALUE)
        print(f"Created secret '{created_secret.name}'.")

        read_secret = client.get_secret(SECRET_NAME)
        print(f"Read secret '{read_secret.name}' with value '{read_secret.value}'.")

        updated_secret = client.set_secret(SECRET_NAME, UPDATED_VALUE)
        print(
            f"Updated secret '{updated_secret.name}' "
            f"to value '{updated_secret.value}'."
        )

        delete_poller = client.begin_delete_secret(SECRET_NAME)
        delete_poller.wait()
        print(f"Deleted secret '{SECRET_NAME}'.")

        # Waiting for deletion ensures the secret is available in the deleted
        # secrets collection before the purge request is sent.
        client.get_deleted_secret(SECRET_NAME)
        client.purge_deleted_secret(SECRET_NAME)
        print(f"Purged secret '{SECRET_NAME}'.")
    finally:
        client.close()
        credential.close()


def main() -> int:
    try:
        run_crud_operations()
        return 0
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
    except ResourceNotFoundError as error:
        print(f"Secret was not found during the operation: {error}", file=sys.stderr)
    except AzureError as error:
        print(f"Azure Key Vault operation failed: {error}", file=sys.stderr)
    except KeyboardInterrupt:
        print("Operation cancelled.", file=sys.stderr)
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
