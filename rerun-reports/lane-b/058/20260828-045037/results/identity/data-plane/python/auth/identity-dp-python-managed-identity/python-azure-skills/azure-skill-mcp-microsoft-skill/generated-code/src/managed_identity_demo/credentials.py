"""Credential selection for Azure-hosted and local environments."""

from enum import Enum
from typing import Optional, Union

from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

AzureCredential = Union[DefaultAzureCredential, ManagedIdentityCredential]


class CredentialMode(str, Enum):
    """Supported authentication strategies."""

    SYSTEM_ASSIGNED = "system"
    USER_ASSIGNED = "user"
    LOCAL = "local"
    AUTO_SYSTEM = "auto-system"
    AUTO_USER = "auto-user"


def system_assigned_credential() -> ManagedIdentityCredential:
    """Use the single system-assigned identity attached to the Azure host."""
    return ManagedIdentityCredential()


def user_assigned_credential(client_id: str) -> ManagedIdentityCredential:
    """Select one user-assigned identity attached to the Azure host."""
    if not client_id.strip():
        raise ValueError("A managed identity client ID is required for user-assigned identity.")
    return ManagedIdentityCredential(client_id=client_id)


def local_development_credential(
    *,
    allow_interactive_browser: bool = False,
) -> DefaultAzureCredential:
    """Use developer-tool credentials such as Azure CLI, VS Code, or Azure Developer CLI."""
    return DefaultAzureCredential(
        exclude_managed_identity_credential=True,
        exclude_interactive_browser_credential=not allow_interactive_browser,
    )


def create_credential(
    mode: CredentialMode,
    *,
    client_id: Optional[str] = None,
    allow_interactive_browser: bool = False,
) -> AzureCredential:
    """Create a credential with either strict or environment-adaptive behavior."""
    if mode is CredentialMode.SYSTEM_ASSIGNED:
        return system_assigned_credential()
    if mode is CredentialMode.USER_ASSIGNED:
        return user_assigned_credential(client_id or "")
    if mode is CredentialMode.LOCAL:
        return local_development_credential(
            allow_interactive_browser=allow_interactive_browser
        )
    if mode is CredentialMode.AUTO_SYSTEM:
        return DefaultAzureCredential(
            exclude_interactive_browser_credential=not allow_interactive_browser
        )
    if mode is CredentialMode.AUTO_USER:
        if not client_id or not client_id.strip():
            raise ValueError(
                "A managed identity client ID is required for auto-user mode."
            )
        return DefaultAzureCredential(
            managed_identity_client_id=client_id,
            exclude_interactive_browser_credential=not allow_interactive_browser,
        )
    raise ValueError(f"Unsupported credential mode: {mode}")
