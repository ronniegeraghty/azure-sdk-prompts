"""Synchronous and asynchronous Azure credential connectivity checks."""

from __future__ import annotations

from datetime import datetime, timezone

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError


def _expiry_text(expires_on: int) -> str:
    return datetime.fromtimestamp(expires_on, tz=timezone.utc).isoformat()


def _failure_reason(error: Exception) -> str:
    message = str(error).strip()
    lowered = message.lower()

    patterns = (
        (("expired", "certificate"), "the client certificate has expired"),
        (("aadsts7000222",), "the client secret has expired"),
        (("expired", "secret"), "the client secret has expired"),
        (("aadsts7000215",), "the client secret is invalid"),
        (("invalid client secret",), "the client secret is invalid"),
        (("aadsts90002",), "the tenant does not exist or is incorrect"),
        (("tenant", "not found"), "the tenant does not exist or is incorrect"),
        (("wrong tenant",), "the configured tenant is incorrect"),
        (("no managed identity",), "no managed identity is assigned to this host"),
        (("identity not found",), "the requested managed identity is not available"),
        (("credentialunavailableerror",), "no credential in the chain is available"),
        (("connection", "refused"), "the identity endpoint could not be reached"),
        (("name resolution",), "the identity service hostname could not be resolved"),
        (("timed out",), "the identity service request timed out"),
    )
    for terms, reason in patterns:
        if all(term in lowered for term in terms):
            return f"{reason}: {message}"
    return message or error.__class__.__name__


def test_credential_sync(
    credential: TokenCredential, scope: str, *, enable_cae: bool = False
) -> bool:
    """Request a token and print a diagnostic result."""

    print(f"[sync] Requesting token (CAE requested: {enable_cae})")
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except CredentialUnavailableError as error:
        print(f"[sync] FAILED - no identity available: {_failure_reason(error)}")
        return False
    except ClientAuthenticationError as error:
        print(f"[sync] FAILED - authentication rejected: {_failure_reason(error)}")
        return False
    except (OSError, TimeoutError) as error:
        print(f"[sync] FAILED - identity service unavailable: {_failure_reason(error)}")
        return False

    print(f"[sync] SUCCESS - token expires at {_expiry_text(token.expires_on)}")
    return True


async def test_credential_async(
    credential: AsyncTokenCredential, scope: str, *, enable_cae: bool = False
) -> bool:
    """Request a token asynchronously and print a diagnostic result."""

    print(f"[async] Requesting token (CAE requested: {enable_cae})")
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except CredentialUnavailableError as error:
        print(f"[async] FAILED - no identity available: {_failure_reason(error)}")
        return False
    except ClientAuthenticationError as error:
        print(f"[async] FAILED - authentication rejected: {_failure_reason(error)}")
        return False
    except (OSError, TimeoutError) as error:
        print(f"[async] FAILED - identity service unavailable: {_failure_reason(error)}")
        return False

    print(f"[async] SUCCESS - token expires at {_expiry_text(token.expires_on)}")
    return True
