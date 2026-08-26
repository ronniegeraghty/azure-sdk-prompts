"""Construct SDK clients with the Azure host's system-assigned identity."""

from managed_identity_demo import IdentityMode, Settings, create_clients

settings = Settings(
    identity_mode=IdentityMode.SYSTEM_ASSIGNED,
    storage_account_url="https://example.blob.core.windows.net",
    key_vault_url="https://example.vault.azure.net",
)
clients = create_clients(settings)

print(type(clients.blob_service).__name__)
print(type(clients.secrets).__name__)

