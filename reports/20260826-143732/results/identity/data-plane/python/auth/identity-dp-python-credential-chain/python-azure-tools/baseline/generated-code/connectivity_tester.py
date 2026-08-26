"""Connectivity diagnostics for synchronous and asynchronous Azure credentials."""

from __future__ import annotations

from datetime import datetime, timezone

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError


def _expiry_text(expires_on: int) -> str:
    return datetime.fromtimestamp(expires_on, timezone.utc).isoformat()


def _authentication_failure_reason(error: BaseException) -> tuple[str, str]:
    message = str(error).strip() or error.__class__.__name__
    normalized = message.lower()

    indicators = (
        (
            "expired certificate",
            ("certificate has expired", "certificate is expired", "expired cert"),
        ),
        (
            "wrong tenant or tenant not found",
            (
                "tenant not found",
                "invalid tenant",
                "unauthorized_client",
                "aadsts90002",
                "aadsts500011",
            ),
        ),
        (
            "client secret or certificate rejected",
            (
                "invalid_client",
                "aadsts7000215",
                "aadsts700027",
                "client secret",
            ),
        ),
        (
            "federated identity configuration rejected",
            ("aadsts70021", "federated identity credential", "assertion"),
        ),
        (
            "managed identity unavailable or not assigned",
            (
                "no identity",
                "identity not found",
                "managed identity",
                "imds endpoint",
                "msi endpoint",
            ),
        ),
        (
            "developer login required",
            (
                "az login",
                "azure developer cli",
                "visual studio code",
                "authentication required",
            ),
        ),
        (
            "network or authority endpoint unavailable",
            (
                "connection",
                "temporarily unavailable",
                "name resolution",
                "timed out",
                "timeout",
            ),
        ),
        (
            "insufficient permission or consent",
            (
                "access_denied",
                "insufficient",
                "consent",
                "aadsts65001",
            ),
        ),
    )
    for reason, phrases in indicators:
        if any(phrase in normalized for phrase in phrases):
            return reason, message

    if isinstance(error, CredentialUnavailableError):
        return "no credential in the chain is available", message
    if isinstance(error, ClientAuthenticationError):
        return "Azure rejected the authentication request", message
    return f"unexpected {error.__class__.__name__}", message


def _print_failure(label: str, error: BaseException, enable_cae: bool) -> None:
    reason, detail = _authentication_failure_reason(error)
    print(f"[{label}] FAILURE")
    print(f"[{label}] CAE requested: {enable_cae}")
    print(f"[{label}] Reason: {reason}")
    print(f"[{label}] Detail: {detail}")


def test_credential(
    credential: TokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> bool:
    """Request a token synchronously and print actionable diagnostics."""

    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except (ClientAuthenticationError, CredentialUnavailableError) as error:
        _print_failure("sync", error, enable_cae)
        return False
    except Exception as error:
        _print_failure("sync", error, enable_cae)
        return False

    print("[sync] SUCCESS")
    print(f"[sync] Token expires: {_expiry_text(token.expires_on)}")
    print(f"[sync] CAE requested: {enable_cae}")
    return True


async def test_credential_async(
    credential: AsyncTokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> bool:
    """Request a token asynchronously and print actionable diagnostics."""

    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except (ClientAuthenticationError, CredentialUnavailableError) as error:
        _print_failure("async", error, enable_cae)
        return False
    except Exception as error:
        _print_failure("async", error, enable_cae)
        return False

    print("[async] SUCCESS")
    print(f"[async] Token expires: {_expiry_text(token.expires_on)}")
    print(f"[async] CAE requested: {enable_cae}")
    return True
