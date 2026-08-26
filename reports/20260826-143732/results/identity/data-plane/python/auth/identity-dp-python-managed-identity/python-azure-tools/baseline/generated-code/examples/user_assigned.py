"""Construct SDK clients with a particular user-assigned identity."""

import os

from managed_identity_demo import IdentityMode, Settings, create_clients

settings = Settings(
    identity_mode=IdentityMode.USER_ASSIGNED,
    managed_identity_client_id=os.environ["AZURE_CLIENT_ID"],
    storage_account_url="https://example.blob.core.windows.net",
    key_vault_url="https://example.vault.azure.net",
)
clients = create_clients(settings)

print(type(clients.blob_service).__name__)
print(type(clients.secrets).__name__)

