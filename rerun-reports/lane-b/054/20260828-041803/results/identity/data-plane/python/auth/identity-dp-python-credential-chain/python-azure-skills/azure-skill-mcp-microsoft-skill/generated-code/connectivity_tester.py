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
    success: bool
    cae_requested: bool
    expires_at: datetime | None = None
    failure_reason: str | None = None


def test_credential(
    credential: TokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Request a token and print a detailed synchronous test result."""
    print("[sync] Requesting token...")
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        result = _failure_result(error, enable_cae)
    except Exception as error:
        result = ConnectivityResult(
            success=False,
            cae_requested=enable_cae,
            failure_reason=(
                f"Unexpected {type(error).__name__}: {str(error).strip() or 'no details'}"
            ),
        )
    else:
        result = ConnectivityResult(
            success=True,
            cae_requested=enable_cae,
            expires_at=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),
        )
    _print_result("sync", result)
    return result


async def test_credential_async(
    credential: AsyncTokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Request a token and print a detailed asynchronous test result."""
    print("[async] Requesting token...")
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        result = _failure_result(error, enable_cae)
    except Exception as error:
        result = ConnectivityResult(
            success=False,
            cae_requested=enable_cae,
            failure_reason=(
                f"Unexpected {type(error).__name__}: {str(error).strip() or 'no details'}"
            ),
        )
    else:
        result = ConnectivityResult(
            success=True,
            cae_requested=enable_cae,
            expires_at=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),
        )
    _print_result("async", result)
    return result


def _failure_result(
    error: CredentialUnavailableError | ClientAuthenticationError,
    enable_cae: bool,
) -> ConnectivityResult:
    details = str(error).strip()
    normalized = details.lower()

    if isinstance(error, CredentialUnavailableError):
        category = "No identity is available for this credential"
    elif any(
        phrase in normalized
        for phrase in (
            "certificate has expired",
            "expired certificate",
            "aadsts700027",
        )
    ):
        category = "The client certificate is expired or invalid"
    elif any(
        phrase in normalized
        for phrase in (
            "secret is expired",
            "client secret",
            "aadsts7000222",
        )
    ):
        category = "The client secret is expired or invalid"
    elif any(
        phrase in normalized
        for phrase in (
            "tenant not found",
            "invalid tenant",
            "aadsts90002",
        )
    ):
        category = "The configured tenant is invalid or cannot be found"
    elif any(
        phrase in normalized
        for phrase in (
            "managed identity",
            "identity not found",
            "no response from the imds endpoint",
        )
    ):
        category = "The requested managed identity is unavailable"
    elif any(
        phrase in normalized
        for phrase in (
            "federated identity credential",
            "federated token",
            "aadsts70021",
            "aadsts700212",
        )
    ):
        category = "The workload identity federation settings do not match"
    elif "unauthorized_client" in normalized or "aadsts700016" in normalized:
        category = "The client ID is invalid or not registered in this tenant"
    elif "interaction_required" in normalized:
        category = "User interaction or a fresh developer login is required"
    else:
        category = "Microsoft Entra ID rejected the authentication request"

    reason = f"{category}. SDK details: {details or type(error).__name__}"
    return ConnectivityResult(
        success=False,
        cae_requested=enable_cae,
        failure_reason=reason,
    )


def _print_result(label: str, result: ConnectivityResult) -> None:
    status = "SUCCESS" if result.success else "FAILURE"
    print(f"[{label}] Result: {status}")
    print(f"[{label}] CAE requested: {'yes' if result.cae_requested else 'no'}")
    if result.expires_at:
        print(f"[{label}] Token expires (UTC): {result.expires_at.isoformat()}")
    if result.failure_reason:
        print(f"[{label}] Failure reason: {result.failure_reason}")
