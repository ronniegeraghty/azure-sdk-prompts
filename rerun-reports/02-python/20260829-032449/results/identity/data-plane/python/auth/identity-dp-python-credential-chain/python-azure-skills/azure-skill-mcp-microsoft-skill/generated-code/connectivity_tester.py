"""Token-based connectivity checks for synchronous Azure credentials."""

from __future__ import annotations

from datetime import datetime, timezone

from azure.core.credentials import AccessToken, TokenCredential
from azure.core.exceptions import AzureError, ClientAuthenticationError
from azure.identity import CredentialUnavailableError


def test_connectivity(
    credential: TokenCredential, scope: str, *, enable_cae: bool = False
) -> bool:
    """Request a token and print a diagnostic result without exposing the token."""
    print(f"  CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        _print_authentication_failure(error)
        return False
    except AzureError as error:
        print(f"  Result: FAILED - Azure SDK error: {_safe_message(error)}")
        return False

    _print_success(token)
    return True


def _print_success(token: AccessToken) -> None:
    expires_at = datetime.fromtimestamp(token.expires_on, tz=timezone.utc)
    print("  Result: SUCCESS")
    print(f"  Token expires (UTC): {expires_at.isoformat()}")


def _print_authentication_failure(error: Exception) -> None:
    reason = classify_authentication_failure(error)
    print(f"  Result: FAILED - {reason}")
    print(f"  SDK details: {_safe_message(error)}")


def classify_authentication_failure(error: Exception) -> str:
    """Translate common Entra and Azure Identity failures into actionable reasons."""
    message = str(error).lower()
    if "expired" in message and any(
        word in message for word in ("certificate", "x509", "x.509")
    ):
        return "the service principal certificate has expired"

    patterns = (
        (
            ("aadsts7000222", "client secret keys for app", "secret has expired"),
            "the service principal client secret has expired",
        ),
        (
            ("aadsts7000215", "invalid client secret"),
            "the service principal client secret is invalid",
        ),
        (
            ("aadsts90002", "tenant not found", "invalid_tenant"),
            "the tenant ID or authority is incorrect",
        ),
        (
            ("aadsts700016", "application with identifier"),
            "the client ID is not registered in the configured tenant",
        ),
        (
            (
                "federated identity credential",
                "federated token file",
                "federation settings",
            ),
            "the workload identity token file or federation settings are invalid",
        ),
        (
            ("continuous access evaluation", "enable_cae"),
            "the selected credential does not support the requested CAE token",
        ),
        (
            ("invalid_resource", "aadsts500011"),
            "the requested scope or resource is invalid for this tenant",
        ),
        (
            (
                "no managed identity",
                "identity not found",
                "credential unavailable",
                "managed identity endpoint",
            ),
            "no usable identity is available in this environment",
        ),
    )
    for needles, reason in patterns:
        if any(needle in message for needle in needles):
            return reason

    if isinstance(error, CredentialUnavailableError):
        return "no credential in the configured chain could attempt authentication"
    return "Microsoft Entra ID rejected the authentication request"


def _safe_message(error: Exception) -> str:
    message = " ".join(str(error).split())
    return message[:1000] if message else type(error).__name__
