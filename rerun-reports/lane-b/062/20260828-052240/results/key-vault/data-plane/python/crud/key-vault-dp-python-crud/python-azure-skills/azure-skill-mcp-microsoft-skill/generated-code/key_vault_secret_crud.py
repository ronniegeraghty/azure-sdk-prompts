"""Create, read, update, delete, and purge an Azure Key Vault secret."""

import os
import sys

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


SECRET_NAME = os.getenv("AZURE_KEYVAULT_SECRET_NAME", "crud-demo-secret")
INITIAL_VALUE = "initial-value"
UPDATED_VALUE = "updated-value"


def get_vault_url() -> str:
    """Return and validate the Key Vault URL from the environment."""
    vault_url = os.getenv("AZURE_KEYVAULT_URL")
    if not vault_url:
        raise ValueError(
            "AZURE_KEYVAULT_URL is required, for example "
            "https://<vault-name>.vault.azure.net/"
        )

    if not vault_url.startswith("https://") or ".vault.azure.net" not in vault_url:
        raise ValueError("AZURE_KEYVAULT_URL must be a valid Azure Key Vault HTTPS URL")

    return vault_url


def run_crud_operations(vault_url: str) -> None:
    """Perform all CRUD operations on one secret."""
    with DefaultAzureCredential() as credential:
        with SecretClient(vault_url=vault_url, credential=credential) as client:
            created = client.set_secret(SECRET_NAME, INITIAL_VALUE)
            print(
                f"Created secret {created.name!r}, "
                f"version {created.properties.version!r}."
            )

            retrieved = client.get_secret(SECRET_NAME)
            if retrieved.value != INITIAL_VALUE:
                raise RuntimeError("The retrieved secret does not match the created value")
            print(f"Read secret {retrieved.name!r}; its value matched the created value.")

            updated = client.set_secret(SECRET_NAME, UPDATED_VALUE)
            print(
                f"Updated secret {updated.name!r}, "
                f"version {updated.properties.version!r}."
            )

            retrieved_updated = client.get_secret(SECRET_NAME)
            if retrieved_updated.value != UPDATED_VALUE:
                raise RuntimeError("The retrieved secret does not match the updated value")
            print("Read the updated secret; its value matched 'updated-value'.")

            client.begin_delete_secret(SECRET_NAME).result()
            print(f"Deleted secret {SECRET_NAME!r} (soft delete).")

            client.purge_deleted_secret(SECRET_NAME)
            print(f"Purged secret {SECRET_NAME!r}.")


def main() -> int:
    """Run the example and translate expected failures into useful messages."""
    try:
        run_crud_operations(get_vault_url())
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        print(
            "Authentication failed. Sign in with a supported developer credential "
            "(for example, Azure CLI) or configure a managed identity. "
            f"Details: {error}",
            file=sys.stderr,
        )
    except ResourceNotFoundError as error:
        print(
            "The vault or secret was not found. Check AZURE_KEYVAULT_URL and the "
            f"secret lifecycle state. Details: {error}",
            file=sys.stderr,
        )
    except HttpResponseError as error:
        if error.status_code == 403:
            message = (
                "Access denied. Grant secret get, set, delete, and purge permissions "
                "through Key Vault RBAC or an access policy."
            )
        elif error.status_code == 409:
            message = (
                "The operation conflicted with the secret's current state. A "
                "previously deleted secret with the same name may still be retained."
            )
        else:
            message = "Azure Key Vault request failed."
        print(f"{message} Details: {error}", file=sys.stderr)
    except RuntimeError as error:
        print(f"Secret verification failed: {error}", file=sys.stderr)
    else:
        return 0

    return 1


if __name__ == "__main__":
    sys.exit(main())
