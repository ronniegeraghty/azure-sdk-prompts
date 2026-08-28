"""Create, read, update, delete, and purge an Azure Key Vault secret."""

import logging
import os
import sys
from uuid import uuid4

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


def configure_logging() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(message)s",
    )


def require_key_vault_url() -> str:
    key_vault_url = os.getenv("AZURE_KEY_VAULT_URL")
    if not key_vault_url:
        raise ValueError(
            "AZURE_KEY_VAULT_URL is required, for example "
            "'https://my-vault.vault.azure.net'."
        )
    return key_vault_url


def run_crud(client: SecretClient, secret_name: str) -> None:
    logging.info("Creating secret %r.", secret_name)
    created_secret = client.set_secret(secret_name, INITIAL_VALUE)
    logging.info("Created secret version %s.", created_secret.properties.version)

    logging.info("Reading secret %r.", secret_name)
    retrieved_secret = client.get_secret(secret_name)
    logging.info(
        "Read secret version %s successfully.",
        retrieved_secret.properties.version,
    )

    logging.info("Updating secret %r.", secret_name)
    updated_secret = client.set_secret(secret_name, UPDATED_VALUE)
    logging.info(
        "Updated secret to version %s.",
        updated_secret.properties.version,
    )

    logging.info("Deleting secret %r.", secret_name)
    delete_poller = client.begin_delete_secret(secret_name)
    delete_poller.result()
    logging.info("Secret deletion completed.")

    logging.info("Purging deleted secret %r.", secret_name)
    client.purge_deleted_secret(secret_name)
    logging.info("Secret purge completed.")


def main() -> int:
    configure_logging()

    try:
        key_vault_url = require_key_vault_url()
    except ValueError as error:
        logging.error("%s", error)
        return 2

    secret_name = os.getenv("AZURE_KEY_VAULT_SECRET_NAME")
    if not secret_name:
        secret_name = f"python-crud-{uuid4().hex}"

    try:
        with (
            DefaultAzureCredential() as credential,
            SecretClient(vault_url=key_vault_url, credential=credential) as client,
        ):
            run_crud(client, secret_name)
    except ClientAuthenticationError:
        logging.exception(
            "Authentication failed. Sign in with a supported DefaultAzureCredential "
            "method or configure a managed identity."
        )
        return 1
    except ResourceNotFoundError:
        logging.exception(
            "The vault or secret was not found during the CRUD operation."
        )
        return 1
    except HttpResponseError:
        logging.exception(
            "Azure Key Vault rejected an operation. Verify the vault URL, RBAC "
            "permissions, firewall settings, and that purge protection is disabled."
        )
        return 1
    except AzureError:
        logging.exception("An Azure SDK error interrupted the CRUD operation.")
        return 1

    return 0


if __name__ == "__main__":
    sys.exit(main())
