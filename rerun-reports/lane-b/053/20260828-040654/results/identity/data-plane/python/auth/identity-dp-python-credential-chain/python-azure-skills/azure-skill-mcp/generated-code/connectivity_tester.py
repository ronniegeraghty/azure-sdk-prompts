"""Synchronous and asynchronous Azure credential connectivity tests."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timezone

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError


@dataclass(frozen=True)
class ConnectivityResult:
    succeeded: bool
    detail: str
    expires_at: datetime | None = None


_FAILURE_PATTERNS = (
    (
        ("aadsts7000222", "expired client secret", "secret has expired"),
        "The client secret has expired.",
    ),
    (
        ("aadsts700027", "certificate is expired", "expired certificate"),
        "The client certificate is expired or invalid.",
    ),
    (
        ("aadsts7000215", "invalid client secret"),
        "The client secret is incorrect (use the secret value, not its ID).",
    ),
    (
        (
            "aadsts50020",
            "aadsts50059",
            "tenant not found",
            "invalid tenant",
            "wrong tenant",
        ),
        "The configured tenant is wrong or unavailable.",
    ),
    (
        ("aadsts700016", "application with identifier"),
        "The client application was not found in the configured tenant.",
    ),
    (
        ("aadsts700024", "assertion is not within its valid time range"),
        "The federated credential assertion is expired or not yet valid.",
    ),
    (
        ("aadsts70025", "no configured federated identity"),
        "No matching federated identity credential is configured.",
    ),
    (
        ("aadsts65001", "consent required", "admin consent"),
        "The application lacks required tenant consent.",
    ),
    (
        ("certificate", "private key"),
        "The client certificate or private key could not be loaded.",
    ),
    (
        ("managed identity", "no response"),
        "The managed identity endpoint did not respond.",
    ),
    (
        ("identity not found", "no user assigned identities"),
        "The requested managed identity is not assigned to this resource.",
    ),
    (
        ("federated token file", "token file"),
        "The workload identity token file is missing or unreadable.",
    ),
)


def explain_authentication_failure(error: BaseException) -> str:
    """Translate common Azure authentication failures without hiding details."""

    message = str(error).strip()
    normalized = message.lower()
    unavailable_markers = (
        "credentialunavailableerror",
        "credential unavailable",
        "authentication unavailable",
        "not found on path",
        "not installed",
        "no managed identity endpoint",
        "no identity has been assigned",
        "did not attempt to retrieve a token",
    )
    if isinstance(error, CredentialUnavailableError) or any(
        marker in normalized for marker in unavailable_markers
    ):
        reason = "No configured credential source is available."
    else:
        reason = "Microsoft Entra ID rejected the authentication request."

    for needles, explanation in _FAILURE_PATTERNS:
        if any(needle in normalized for needle in needles):
            reason = explanation
            break

    if not message:
        return reason
    compact_message = " ".join(message.split())
    return f"{reason} Azure SDK detail: {compact_message}"


def _expiry(expires_on: int) -> datetime:
    return datetime.fromtimestamp(expires_on, tz=timezone.utc)


def test_credential(
    credential: TokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Request a token and print a diagnostic result."""

    print(f"[sync] CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        detail = explain_authentication_failure(error)
        print(f"[sync] Authentication failed: {detail}")
        return ConnectivityResult(False, detail)
    except (OSError, ValueError) as error:
        detail = f"Credential configuration error: {error}"
        print(f"[sync] Authentication failed: {detail}")
        return ConnectivityResult(False, detail)

    expires_at = _expiry(token.expires_on)
    detail = "Token acquired successfully."
    print(f"[sync] Success: {detail}")
    print(f"[sync] Token expires (UTC): {expires_at.isoformat()}")
    return ConnectivityResult(True, detail, expires_at)


async def test_credential_async(
    credential: AsyncTokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Asynchronously request a token and print a diagnostic result."""

    print(f"[async] CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        detail = explain_authentication_failure(error)
        print(f"[async] Authentication failed: {detail}")
        return ConnectivityResult(False, detail)
    except (OSError, ValueError) as error:
        detail = f"Credential configuration error: {error}"
        print(f"[async] Authentication failed: {detail}")
        return ConnectivityResult(False, detail)

    expires_at = _expiry(token.expires_on)
    detail = "Token acquired successfully."
    print(f"[async] Success: {detail}")
    print(f"[async] Token expires (UTC): {expires_at.isoformat()}")
    return ConnectivityResult(True, detail, expires_at)
