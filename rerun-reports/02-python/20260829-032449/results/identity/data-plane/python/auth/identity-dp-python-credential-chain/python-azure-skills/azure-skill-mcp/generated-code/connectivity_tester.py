"""Synchronous and asynchronous Azure credential connectivity tests."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timezone
import re

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError


@dataclass(frozen=True)
class ConnectivityResult:
    succeeded: bool
    scope: str
    cae_requested: bool
    expires_on: datetime | None = None
    failure_category: str | None = None
    failure_detail: str | None = None


_FAILURE_PATTERNS = (
    (
        r"AADSTS7000222|certificate.+expired|expired.+certificate|client secret.+expired",
        "expired client certificate or secret",
    ),
    (
        r"AADSTS90002|tenant.+(?:not found|invalid)|invalid tenant",
        "wrong or unavailable tenant",
    ),
    (
        r"AADSTS700016|application.+not found|unauthorized_client",
        "client ID is unknown in the selected tenant",
    ),
    (
        r"AADSTS7000215|invalid client secret|invalid_client",
        "invalid client secret or client credential",
    ),
    (
        r"AADSTS700024|client assertion.+(?:expired|not within)",
        "expired or not-yet-valid federated assertion",
    ),
    (
        r"AADSTS700027|certificate.+(?:invalid|not registered)|thumbprint",
        "client certificate is invalid or not registered",
    ),
    (
        r"AADSTS70021|federated identity credential|subject.+issuer.+audience",
        "workload identity federation configuration mismatch",
    ),
    (
        r"AADSTS70011|invalid_scope|scope.+invalid",
        "invalid or unauthorized token scope",
    ),
    (
        r"managed identity|IMDS|identity endpoint|no identity|unavailable",
        "no usable managed or workload identity is available",
    ),
    (
        r"Azure CLI|Azure Developer CLI|Azure PowerShell|Visual Studio Code",
        "no authenticated developer-tool account is available",
    ),
)


def _diagnose(error: BaseException) -> tuple[str, str]:
    message = getattr(error, "message", None) or str(error)
    compact_message = re.sub(r"\s+", " ", message).strip()
    if isinstance(error, CredentialUnavailableError):
        category = "credential unavailable"
    else:
        category = "authentication rejected"

    for pattern, specific_category in _FAILURE_PATTERNS:
        if re.search(pattern, compact_message, flags=re.IGNORECASE):
            category = specific_category
            break

    detail = compact_message[:800] or error.__class__.__name__
    return category, detail


def _print_result(label: str, result: ConnectivityResult) -> None:
    cae = "yes" if result.cae_requested else "no"
    if result.succeeded:
        expiry = result.expires_on.isoformat() if result.expires_on else "unknown"
        print(f"[{label}] SUCCESS")
        print(f"[{label}] Token expiry (UTC): {expiry}")
        print(f"[{label}] CAE requested: {cae}")
        return

    print(f"[{label}] FAILURE")
    print(f"[{label}] Reason: {result.failure_category}")
    print(f"[{label}] Detail: {result.failure_detail}")
    print(f"[{label}] CAE requested: {cae}")


def test_credential_sync(
    credential: TokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Request a token synchronously and report a diagnostic result."""
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
        result = ConnectivityResult(
            succeeded=True,
            scope=scope,
            cae_requested=enable_cae,
            expires_on=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),
        )
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        category, detail = _diagnose(error)
        result = ConnectivityResult(
            succeeded=False,
            scope=scope,
            cae_requested=enable_cae,
            failure_category=category,
            failure_detail=detail,
        )
    except Exception as error:
        category, detail = _diagnose(error)
        result = ConnectivityResult(
            succeeded=False,
            scope=scope,
            cae_requested=enable_cae,
            failure_category=f"unexpected {category}",
            failure_detail=detail,
        )

    _print_result("sync", result)
    return result


async def test_credential_async(
    credential: AsyncTokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> ConnectivityResult:
    """Request a token asynchronously and report a diagnostic result."""
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
        result = ConnectivityResult(
            succeeded=True,
            scope=scope,
            cae_requested=enable_cae,
            expires_on=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),
        )
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        category, detail = _diagnose(error)
        result = ConnectivityResult(
            succeeded=False,
            scope=scope,
            cae_requested=enable_cae,
            failure_category=category,
            failure_detail=detail,
        )
    except Exception as error:
        category, detail = _diagnose(error)
        result = ConnectivityResult(
            succeeded=False,
            scope=scope,
            cae_requested=enable_cae,
            failure_category=f"unexpected {category}",
            failure_detail=detail,
        )

    _print_result("async", result)
    return result
