"""Token-based connectivity checks for asynchronous Azure credentials."""

from __future__ import annotations

from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import AzureError, ClientAuthenticationError
from azure.identity import CredentialUnavailableError

from connectivity_tester import _print_authentication_failure, _print_success


async def test_connectivity_async(
    credential: AsyncTokenCredential, scope: str, *, enable_cae: bool = False
) -> bool:
    """Request a token asynchronously and print a diagnostic result."""
    print(f"  CAE requested: {'yes' if enable_cae else 'no'}")
    try:
        token = await credential.get_token(scope, enable_cae=enable_cae)
    except (CredentialUnavailableError, ClientAuthenticationError) as error:
        _print_authentication_failure(error)
        return False
    except AzureError as error:
        message = " ".join(str(error).split())[:1000] or type(error).__name__
        print(f"  Result: FAILED - Azure SDK error: {message}")
        return False

    _print_success(token)
    return True
