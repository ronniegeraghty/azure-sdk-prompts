"""Consistent user-facing error conversion for Azure Storage operations."""

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

from .models import OperationResult

STORAGE_EXCEPTIONS = (
    ResourceNotFoundError,
    ResourceExistsError,
    ResourceModifiedError,
    ClientAuthenticationError,
    ServiceRequestError,
    ServiceResponseError,
    HttpResponseError,
)


def storage_failure(action: str, exc: Exception) -> OperationResult[None]:
    error_code = str(getattr(exc, "error_code", "") or "")
    if isinstance(exc, ResourceNotFoundError):
        detail = "the container or blob was not found"
    elif isinstance(exc, ResourceExistsError):
        detail = "the blob changed concurrently or already exists"
    elif isinstance(exc, ResourceModifiedError):
        detail = "the blob changed concurrently; reload it and retry"
    elif isinstance(exc, ClientAuthenticationError):
        detail = "authentication failed; verify the managed identity and RBAC role"
    elif isinstance(exc, (ServiceRequestError, ServiceResponseError)):
        detail = f"Azure Storage could not be reached: {exc}"
    elif isinstance(exc, HttpResponseError) and exc.status_code == 403:
        detail = "permission denied; verify the identity has the required data-plane role"
    elif isinstance(exc, HttpResponseError) and any(
        lease_error in error_code
        for lease_error in (
            "LeaseAlreadyPresent",
            "LeaseIdMismatchWithBlobOperation",
            "LeaseLost",
        )
    ):
        detail = f"the blob lease prevented the operation ({error_code})"
    else:
        detail = f"Azure Storage returned {error_code or type(exc).__name__}: {exc}"
    return OperationResult(success=False, message=f"Could not {action}: {detail}.")
