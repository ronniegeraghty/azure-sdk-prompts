"""Azure managed identity authentication examples."""

from .credentials import (
    CredentialMode,
    create_credential,
    system_assigned_credential,
    user_assigned_credential,
)

__all__ = [
    "CredentialMode",
    "create_credential",
    "system_assigned_credential",
    "user_assigned_credential",
]
