"""Offline-safe command that demonstrates client configuration."""

import sys

from .clients import create_clients
from .config import Settings


def main() -> int:
    """Build Azure SDK clients and report the selected configuration."""
    try:
        settings = Settings.from_environment()
        create_clients(settings)
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2

    print(f"Identity mode: {settings.identity_mode.value}")
    print(f"Blob endpoint: {settings.storage_account_url}")
    print(f"Key Vault endpoint: {settings.key_vault_url}")
    print("Azure SDK clients created; no network request was made.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

