"""Azure Storage exception translation."""

from __future__ import annotations

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceModifiedError,
    ResourceNotFoundError,
)


def storage_error_message(operation: str, error: AzureError) -> str:
    """Convert SDK exceptions into concise messages safe to show to callers."""

    if isinstance(error, ResourceNotFoundError):
        reason = "the container or blob was not found"
    elif isinstance(error, ClientAuthenticationError):
        reason = "authentication failed; check the managed identity configuration"
    elif isinstance(error, ResourceExistsError):
        reason = "the blob changed concurrently and was not overwritten"
    elif isinstance(error, ResourceModifiedError):
        reason = "the blob changed concurrently; retry with its current lease"
    elif isinstance(error, HttpResponseError):
        error_code = str(getattr(error, "error_code", "") or "")
        if error.status_code == 403:
            reason = "permission was denied by Azure Storage"
        elif error.status_code in (409, 412) and "lease" in error_code.lower():
            reason = (
                "the blob is leased by another client or the supplied lease is invalid"
            )
        else:
            detail = error_code or f"HTTP {error.status_code or 'error'}"
            reason = f"Azure Storage returned {detail}"
    else:
        reason = f"the Azure SDK reported {type(error).__name__}"

    return f"{operation} failed: {reason}."
