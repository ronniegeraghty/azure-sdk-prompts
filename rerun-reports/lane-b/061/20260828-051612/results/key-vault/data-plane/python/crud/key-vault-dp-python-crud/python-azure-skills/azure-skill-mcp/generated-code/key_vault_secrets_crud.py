"""Demonstrate create, read, update, delete, and purge for a Key Vault secret."""

from __future__ import annotations

import os
import sys
import uuid

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.keyvault.secrets import SecretClient

INITIAL_VALUE = "initial-value"
UPDATED_VALUE = "updated-value"


def run_crud(vault_url: str, secret_name: str) -> None:
    credential = DefaultAzureCredential()

    try:
        with SecretClient(vault_url=vault_url, credential=credential) as client:
            created = client.set_secret(secret_name, INITIAL_VALUE)
            print(f"Created secret {created.name!r} (version {created.properties.version}).")

            fetched = client.get_secret(secret_name)
            if fetched.value != INITIAL_VALUE:
                raise RuntimeError("The value read after creation did not match.")
            print(f"Read secret {fetched.name!r} successfully.")

            updated = client.set_secret(secret_name, UPDATED_VALUE)
            fetched_updated = client.get_secret(secret_name)
            if fetched_updated.value != UPDATED_VALUE:
                raise RuntimeError("The value read after the update did not match.")
            print(
                f"Updated secret {updated.name!r} "
                f"(version {updated.properties.version})."
            )

            delete_poller = client.begin_delete_secret(secret_name)
            delete_poller.wait()
            deleted = delete_poller.result()
            print(f"Deleted secret {deleted.name!r}; it is now soft-deleted.")

            client.purge_deleted_secret(secret_name)
            print(f"Purged secret {secret_name!r} permanently.")
    finally:
        credential.close()


def main() -> int:
    vault_url = os.getenv("AZURE_KEY_VAULT_URL")
    if not vault_url:
        print(
            "Error: set AZURE_KEY_VAULT_URL to the vault URI, for example "
            "https://my-vault.vault.azure.net.",
            file=sys.stderr,
        )
        return 2

    secret_name = os.getenv(
        "AZURE_KEY_VAULT_SECRET_NAME",
        f"crud-demo-secret-{uuid.uuid4().hex[:8]}",
    )

    try:
        run_crud(vault_url, secret_name)
    except CredentialUnavailableError as error:
        print(
            f"Authentication credential is unavailable: {error}",
            file=sys.stderr,
        )
        return 1
    except ClientAuthenticationError as error:
        print(f"Azure authentication failed: {error}", file=sys.stderr)
        return 1
    except ResourceNotFoundError as error:
        print(f"Key Vault resource was not found: {error}", file=sys.stderr)
        return 1
    except HttpResponseError as error:
        status = f" (HTTP {error.status_code})" if error.status_code else ""
        print(f"Azure Key Vault request failed{status}: {error}", file=sys.stderr)
        return 1
    except RuntimeError as error:
        print(f"CRUD verification failed: {error}", file=sys.stderr)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
