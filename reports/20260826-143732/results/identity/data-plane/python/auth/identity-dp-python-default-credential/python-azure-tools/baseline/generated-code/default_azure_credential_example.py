import logging
import os

from azure.identity import DefaultAzureCredential
from azure.keyvault.secrets import SecretClient


def create_secret_client() -> SecretClient:
    vault_url = os.environ["AZURE_KEY_VAULT_URL"]
    credential = DefaultAzureCredential()
    return SecretClient(vault_url=vault_url, credential=credential)


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    logging.getLogger("azure.identity").setLevel(logging.DEBUG)

    client = create_secret_client()
    secret_name = os.environ.get("AZURE_KEY_VAULT_SECRET_NAME", "example-secret")
    secret = client.get_secret(secret_name)
    print(f"Retrieved secret {secret.name!r} (value intentionally not displayed).")


if __name__ == "__main__":
    main()
