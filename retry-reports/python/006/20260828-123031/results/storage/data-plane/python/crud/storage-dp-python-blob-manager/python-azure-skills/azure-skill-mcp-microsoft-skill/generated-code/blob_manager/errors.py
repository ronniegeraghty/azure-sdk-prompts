"""Azure Storage error translation shared by both service implementations."""

from __future__ import annotations

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceModifiedError,
    ResourceNotFoundError,
    ServiceRequestError,
    ServiceResponseError,
)

HANDLED_AZURE_ERRORS = (
    ResourceNotFoundError,
    ResourceExistsError,
    ResourceModifiedError,
    ClientAuthenticationError,
    ServiceRequestError,
    ServiceResponseError,
    HttpResponseError,
)


def describe_storage_error(action: str, error: Exception) -> str:
    if isinstance(error, ResourceNotFoundError):
        detail = "the container or blob was not found"
    elif isinstance(error, ResourceExistsError):
        detail = "the blob already exists or a lease is already held"
    elif isinstance(error, ResourceModifiedError):
        detail = "the blob changed concurrently; retry with the latest version"
    elif isinstance(error, ClientAuthenticationError):
        detail = "Azure authentication failed"
    elif isinstance(error, (ServiceRequestError, ServiceResponseError)):
        detail = "Azure Storage could not be reached or returned an invalid response"
    elif isinstance(error, HttpResponseError):
        status = getattr(error, "status_code", None)
        error_code = getattr(error, "error_code", None)
        if status == 403:
            detail = "permission was denied"
        elif status == 409 and error_code:
            detail = f"the request conflicted with the blob state ({error_code})"
        else:
            suffix = f" ({error_code})" if error_code else ""
            detail = f"Azure Storage returned HTTP {status or 'error'}{suffix}"
    else:
        detail = str(error)
    return f"Could not {action}: {detail}."


def timeout_options(timeout: int | None) -> dict[str, int]:
    if timeout is None:
        return {}
    if timeout <= 0:
        raise ValueError("timeout must be greater than zero.")
    return {"timeout": timeout}
