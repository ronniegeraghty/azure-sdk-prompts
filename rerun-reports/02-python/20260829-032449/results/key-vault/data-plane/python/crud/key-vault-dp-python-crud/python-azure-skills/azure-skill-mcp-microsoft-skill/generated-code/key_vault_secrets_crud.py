"""Demonstrate create, read, update, delete, and purge for a Key Vault secret."""

import logging
import os
import sys

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceNotFoundError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


LOGGER = logging.getLogger(__name__)


def run_crud(vault_url: str, secret_name: str) -> None:
    credential = DefaultAzureCredential()
    try:
        with SecretClient(vault_url=vault_url, credential=credential) as client:
            created_secret = client.set_secret(secret_name, "initial-value")
            LOGGER.info(
                "Created secret %r (version %s).",
                created_secret.name,
                created_secret.properties.version,
            )

            retrieved_secret = client.get_secret(secret_name)
            LOGGER.info(
                "Read secret %r with value %r.",
                retrieved_secret.name,
                retrieved_secret.value,
            )

            updated_secret = client.set_secret(secret_name, "updated-value")
            LOGGER.info(
                "Updated secret %r to value %r (version %s).",
                updated_secret.name,
                updated_secret.value,
                updated_secret.properties.version,
            )

            deleted_secret = client.begin_delete_secret(secret_name).result()
            LOGGER.info("Soft-deleted secret %r.", deleted_secret.name)

            client.purge_deleted_secret(secret_name)
            LOGGER.info("Purged secret %r permanently.", secret_name)
    finally:
        credential.close()


def main() -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

    vault_url = os.getenv("AZURE_KEYVAULT_URL")
    if not vault_url:
        LOGGER.error(
            "AZURE_KEYVAULT_URL is required, for example "
            "https://<vault-name>.vault.azure.net/."
        )
        return 2

    secret_name = os.getenv("AZURE_KEYVAULT_SECRET_NAME", "crud-demo-secret")

    try:
        run_crud(vault_url, secret_name)
    except CredentialUnavailableError:
        LOGGER.exception(
            "No credential was available. Sign in with a supported developer tool "
            "or configure a managed identity."
        )
    except ClientAuthenticationError:
        LOGGER.exception("Azure authentication failed.")
    except ResourceNotFoundError:
        LOGGER.exception(
            "The vault or secret was not found, or the deleted secret was unavailable."
        )
    except ServiceRequestError:
        LOGGER.exception("Could not connect to Azure Key Vault.")
    except HttpResponseError as error:
        if error.status_code == 403:
            LOGGER.error(
                "Access denied. Grant secret get, set, delete, and purge permissions "
                "to the authenticated identity."
            )
        elif error.status_code == 409:
            LOGGER.error(
                "The operation conflicted with the vault state. Purge protection may "
                "prevent permanent deletion."
            )
        else:
            LOGGER.exception(
                "Azure Key Vault returned HTTP status %s.", error.status_code
            )
    else:
        return 0

    return 1


if __name__ == "__main__":
    sys.exit(main())
