"""Synchronous and asynchronous Azure credential connectivity tests."""

from __future__ import annotations

from datetime import datetime, timezone

from azure.core.credentials import AccessToken, TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import AzureError, ClientAuthenticationError
from azure.identity import CredentialUnavailableError


def _format_expiry(token: AccessToken) -> str:
    return datetime.fromtimestamp(token.expires_on, timezone.utc).isoformat()


def _authentication_failure_reason(error: BaseException) -> str:
    detail = str(error).strip() or type(error).__name__
    normalized = detail.lower()

    patterns = (
        (
            ("certificate has expired", "expired certificate", "aadsts700027"),
            "client certificate is expired or invalid",
        ),
        (
            ("aadsts7000222", "client secret is expired", "expired client secret"),
            "client secret has expired",
        ),
        (
            ("aadsts7000215", "invalid client secret"),
            "client secret is incorrect",
        ),
        (
            ("aadsts90002", "tenant not found"),
            "tenant ID is incorrect or the tenant is unavailable",
        ),
        (
            ("aadsts700016", "application with identifier"),
            "client ID is incorrect or the application is not registered in this tenant",
        ),
        (
            ("aadsts700024", "client assertion is not within its valid time range"),
            "federated identity token is expired or not yet valid",
        ),
        (
            ("federated identity credential", "no matching federated identity record"),
            "workload identity federation is not configured for this subject",
        ),
        (
            ("managedidentitycredential authentication unavailable", "no identity"),
            "no managed identity is available to this workload",
        ),
        (
            ("az login", "azure cli not found"),
            "Azure CLI is unavailable or is not signed in",
        ),
        (
            ("timed out", "connection error", "name resolution"),
            "the identity endpoint or Microsoft Entra ID could not be reached",
        ),
    )
    for needles, reason in patterns:
        if any(needle in normalized for needle in needles):
            return f"{reason}. Azure Identity detail: {detail}"
    return f"Azure Identity rejected authentication: {detail}"


def test_credential(
    credential: TokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> bool:
    """Request a token synchronously and print an actionable result."""
    print(f"  CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = credential.get_token(scope, enable_cae=enable_cae)
    except CredentialUnavailableError as error:
        print(f"  Result: FAILED - no credential is available. {error}")
        return False
    except ClientAuthenticationError as error:
        print(f"  Result: FAILED - {_authentication_failure_reason(error)}")
        return False
    except AzureError as error:
        print(
            "  Result: FAILED - Azure Identity transport or token service error: "
            f"{type(error).__name__}: {error}"
        )
        return False
    except (OSError, ValueError, TypeError) as error:
        print(
            "  Result: FAILED - token request could not be completed: "
            f"{type(error).__name__}: {error}"
        )
        return False

    print("  Result: SUCCESS")
    print(f"  Token expires (UTC): {_format_expiry(token)}")
    return True


async def test_credential_async(
    credential: AsyncTokenCredential,
    scope: str,
    *,
    enable_cae: bool = False,
) -> bool:
    """Request a token asynchronously and print an actionable result."""
    print(f"  CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except CredentialUnavailableError as error:
        print(f"  Result: FAILED - no credential is available. {error}")
        return False
    except ClientAuthenticationError as error:
        print(f"  Result: FAILED - {_authentication_failure_reason(error)}")
        return False
    except AzureError as error:
        print(
            "  Result: FAILED - Azure Identity transport or token service error: "
            f"{type(error).__name__}: {error}"
        )
        return False
    except (OSError, ValueError, TypeError) as error:
        print(
            "  Result: FAILED - token request could not be completed: "
            f"{type(error).__name__}: {error}"
        )
        return False

    print("  Result: SUCCESS")
    print(f"  Token expires (UTC): {_format_expiry(token)}")
    return True
