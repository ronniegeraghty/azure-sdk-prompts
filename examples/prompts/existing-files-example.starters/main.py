"""Azure Key Vault Secrets Client — missing error handling.

This script connects to an Azure Key Vault and performs basic secret
operations.  It works when the vault is reachable and the secrets
exist, but has NO error handling — any failure will crash the process.

The evaluation prompt asks the agent to fix this code in place.
"""

import os

from azure.identity import DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


def create_client() -> SecretClient:
    """Create an authenticated SecretClient."""
    vault_url = os.environ["AZURE_KEYVAULT_URL"]
    credential = DefaultAzureCredential()
    client = SecretClient(vault_url=vault_url, credential=credential)
    return client


def get_secret(client: SecretClient, name: str) -> str:
    """Retrieve a secret value by name."""
    secret = client.get_secret(name)
    return secret.value


def set_secret(client: SecretClient, name: str, value: str) -> None:
    """Create or update a secret."""
    client.set_secret(name, value)
    print(f"Secret '{name}' saved successfully.")


def main() -> None:
    client = create_client()

    # Store a secret
    set_secret(client, "my-database-password", "s3cret!value")

    # Read it back
    password = get_secret(client, "my-database-password")
    print(f"Retrieved secret value: {password}")

    # Try to read a secret that may not exist
    api_key = get_secret(client, "my-api-key")
    print(f"API key: {api_key}")


if __name__ == "__main__":
    main()
