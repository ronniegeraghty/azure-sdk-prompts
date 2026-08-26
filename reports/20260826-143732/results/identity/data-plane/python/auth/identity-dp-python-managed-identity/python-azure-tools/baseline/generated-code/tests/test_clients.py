from unittest.mock import Mock, patch

from managed_identity_demo.auth import IdentityMode
from managed_identity_demo.clients import create_clients
from managed_identity_demo.config import Settings


@patch("managed_identity_demo.clients.SecretClient")
@patch("managed_identity_demo.clients.BlobServiceClient")
def test_clients_share_credential(blob_client, secret_client):
    credential = Mock()
    settings = Settings(
        identity_mode=IdentityMode.SYSTEM_ASSIGNED,
        storage_account_url="https://storage.example",
        key_vault_url="https://vault.example",
    )

    clients = create_clients(settings, credential=credential)

    blob_client.assert_called_once_with(
        account_url="https://storage.example",
        credential=credential,
    )
    secret_client.assert_called_once_with(
        vault_url="https://vault.example",
        credential=credential,
    )
    assert clients.blob_service is blob_client.return_value
    assert clients.secrets is secret_client.return_value

