"""Example Azure operations with actionable authentication errors."""

from collections.abc import Iterator

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.keyvault.secrets import KeyVaultSecret

from .clients import AzureClients


class AzureAccessError(RuntimeError):
    """An Azure request failed with context suitable for application logs."""


def list_container_names(clients: AzureClients) -> Iterator[str]:
    """Yield Blob container names after Azure authorizes the identity."""
    try:
        for container in clients.blob_service.list_containers():
            yield container["name"]
    except ClientAuthenticationError as error:
        raise AzureAccessError(
            "Azure rejected the credential. Confirm the identity is enabled, "
            "the correct user-assigned client ID is configured, and the host "
            "can reach the managed identity endpoint."
        ) from error
    except HttpResponseError as error:
        raise AzureAccessError(
            "Blob Storage rejected the request. Confirm the account URL and "
            "assign a Blob Data role to the managed identity."
        ) from error


def get_secret(clients: AzureClients, name: str) -> KeyVaultSecret:
    """Read a Key Vault secret after Azure authorizes the identity."""
    if not name.strip():
        raise ValueError("Secret name must not be empty")

    try:
        return clients.secrets.get_secret(name)
    except ClientAuthenticationError as error:
        raise AzureAccessError(
            "Azure rejected the credential. Confirm the identity configuration "
            "and managed identity endpoint connectivity."
        ) from error
    except HttpResponseError as error:
        raise AzureAccessError(
            "Key Vault rejected the request. Confirm the vault URL and grant "
            "the identity permission to read secrets."
        ) from error

