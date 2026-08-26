"""Azure SDK client construction using one shared token credential."""

from dataclasses import dataclass

from azure.core.credentials import TokenCredential
from azure.keyvault.secrets import SecretClient
from azure.storage.blob import BlobServiceClient

from .auth import create_credential
from .config import Settings


@dataclass(frozen=True)
class AzureClients:
    blob_service: BlobServiceClient
    secrets: SecretClient


def create_clients(
    settings: Settings,
    *,
    credential: TokenCredential | None = None,
) -> AzureClients:
    """Construct clients without making a network request.

    Supplying a credential is useful for tests or applications that manage the
    credential lifetime in a dependency-injection container.
    """
    selected_credential = credential or create_credential(
        settings.identity_mode,
        managed_identity_client_id=settings.managed_identity_client_id,
    )
    return AzureClients(
        blob_service=BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=selected_credential,
        ),
        secrets=SecretClient(
            vault_url=settings.key_vault_url,
            credential=selected_credential,
        ),
    )

