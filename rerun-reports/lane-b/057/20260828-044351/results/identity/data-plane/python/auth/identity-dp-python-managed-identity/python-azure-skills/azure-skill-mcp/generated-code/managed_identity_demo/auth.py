"""Credential selection for Azure-hosted and local environments."""

from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Literal, Mapping

from azure.core.credentials import TokenCredential
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

Environment = Literal["local", "azure"]
IdentityType = Literal["system", "user"]


class ConfigurationError(ValueError):
    """Raised when authentication configuration is incomplete or invalid."""


@dataclass(frozen=True)
class AuthSettings:
    environment: Environment
    identity_type: IdentityType
    client_id: str | None = None

    @classmethod
    def from_environment(
        cls, environ: Mapping[str, str] | None = None
    ) -> "AuthSettings":
        values = os.environ if environ is None else environ
        environment = values.get("APP_ENV", "local").strip().lower()
        identity_type = values.get("MANAGED_IDENTITY_TYPE", "system").strip().lower()
        client_id = values.get("AZURE_CLIENT_ID")

        if environment not in {"local", "azure"}:
            raise ConfigurationError("APP_ENV must be 'local' or 'azure'.")
        if identity_type not in {"system", "user"}:
            raise ConfigurationError(
                "MANAGED_IDENTITY_TYPE must be 'system' or 'user'."
            )
        if environment == "azure" and identity_type == "user" and not client_id:
            raise ConfigurationError(
                "AZURE_CLIENT_ID is required for a user-assigned managed identity."
            )

        return cls(
            environment=environment,
            identity_type=identity_type,
            client_id=client_id,
        )


def create_system_assigned_credential() -> ManagedIdentityCredential:
    """Create a credential for the system-assigned identity of the host."""
    return ManagedIdentityCredential()


def create_user_assigned_credential(client_id: str) -> ManagedIdentityCredential:
    """Create a credential for a specific user-assigned identity."""
    if not client_id.strip():
        raise ConfigurationError(
            "A user-assigned managed identity client ID is required."
        )
    return ManagedIdentityCredential(client_id=client_id)


def create_credential(settings: AuthSettings) -> TokenCredential:
    """Create a deterministic Azure credential for the selected environment."""
    if settings.environment == "local":
        # Avoid probing the managed identity endpoint from a developer machine.
        return DefaultAzureCredential(exclude_managed_identity_credential=True)
    if settings.identity_type == "user":
        return create_user_assigned_credential(settings.client_id or "")
    return create_system_assigned_credential()

