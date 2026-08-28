import argparse
import logging
import os

from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError, DefaultAzureCredential
from azure.storage.blob import BlobServiceClient


def configure_identity_logging(enabled: bool) -> None:
    if not enabled:
        return

    logging.basicConfig(
        level=logging.WARNING,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    logging.getLogger("azure.identity").setLevel(logging.DEBUG)


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Authenticate to Azure Blob Storage with DefaultAzureCredential."
    )
    parser.add_argument(
        "--debug-auth",
        action="store_true",
        help="Show Azure Identity credential-chain diagnostics.",
    )
    args = parser.parse_args()
    configure_identity_logging(args.debug_auth)

    account_url = os.environ.get("AZURE_STORAGE_ACCOUNT_URL")
    if not account_url:
        raise SystemExit(
            "Set AZURE_STORAGE_ACCOUNT_URL to "
            "https://<account-name>.blob.core.windows.net"
        )

    try:
        with DefaultAzureCredential() as credential:
            with BlobServiceClient(
                account_url=account_url,
                credential=credential,
            ) as client:
                print(f"Authenticated to {account_url}")
                for container in client.list_containers():
                    print(container["name"])
    except CredentialUnavailableError as error:
        logging.error(
            "No credential in DefaultAzureCredential was available:\n%s", error
        )
        raise SystemExit(1) from error
    except ClientAuthenticationError as error:
        logging.error(
            "A credential attempted authentication but failed:\n%s", error
        )
        raise SystemExit(1) from error


if __name__ == "__main__":
    main()
