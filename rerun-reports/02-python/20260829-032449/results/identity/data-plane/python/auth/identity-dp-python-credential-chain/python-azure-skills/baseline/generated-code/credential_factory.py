"""Build Azure credential chains suited to the current deployment environment."""

from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Generic, TypeVar

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.identity import (
    AzureCliCredential,
    AzureDeveloperCliCredential,
    AzurePipelinesCredential,
    AzurePowerShellCredential,
    ChainedTokenCredential,
    EnvironmentCredential,
    ManagedIdentityCredential,
    VisualStudioCodeCredential,
    WorkloadIdentityCredential,
)
from azure.identity.aio import (
    AzureCliCredential as AsyncAzureCliCredential,
    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
    ChainedTokenCredential as AsyncChainedTokenCredential,
    EnvironmentCredential as AsyncEnvironmentCredential,
    ManagedIdentityCredential as AsyncManagedIdentityCredential,
    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
)

from environment_detector import RuntimeEnvironment

CredentialT = TypeVar("CredentialT", TokenCredential, AsyncTokenCredential)

USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID"
AZURE_PIPELINES_SERVICE_CONNECTION_ID = "AZURE_PIPELINES_SERVICE_CONNECTION_ID"


@dataclass(frozen=True)
class CredentialSelection(Generic[CredentialT]):
    """A credential plus human-readable information about how it was built."""

    credential: CredentialT
    strategy: str
    enable_cae: bool


def _azure_pipelines_settings() -> tuple[str, str, str, str] | None:
    tenant_id = os.environ.get("AZURE_TENANT_ID", "").strip()
    client_id = os.environ.get("AZURE_CLIENT_ID", "").strip()
    service_connection_id = os.environ.get(
        AZURE_PIPELINES_SERVICE_CONNECTION_ID, ""
    ).strip()
    system_access_token = os.environ.get("SYSTEM_ACCESSTOKEN", "").strip()
    if all((tenant_id, client_id, service_connection_id, system_access_token)):
        return tenant_id, client_id, service_connection_id, system_access_token
    return None


def build_sync_credential(
    environment: RuntimeEnvironment, *, enable_cae: bool = False
) -> CredentialSelection[TokenCredential]:
    """Build a synchronous credential chain for an environment."""

    if environment is RuntimeEnvironment.DEV:
        credential = ChainedTokenCredential(
            AzureCliCredential(),
            AzureDeveloperCliCredential(),
            AzurePowerShellCredential(),
            VisualStudioCodeCredential(),
        )
        strategy = (
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code"
        )
    elif environment is RuntimeEnvironment.CI:
        credentials: list[TokenCredential] = []
        settings = _azure_pipelines_settings()
        if settings:
            tenant_id, client_id, service_connection_id, system_access_token = settings
            credentials.append(
                AzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=service_connection_id,
                    system_access_token=system_access_token,
                )
            )
        credentials.append(EnvironmentCredential())
        credential = ChainedTokenCredential(*credentials)
        strategy = (
            "Azure Pipelines service connection -> pipeline environment variables"
            if settings
            else "pipeline environment variables (EnvironmentCredential)"
        )
    else:
        managed_identity_client_id = os.environ.get(
            USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID
        )
        managed_identity = ManagedIdentityCredential(
            client_id=managed_identity_client_id
        )
        credentials = [managed_identity]
        strategy = (
            "user-assigned managed identity"
            if managed_identity_client_id
            else "system-assigned managed identity"
        )
        if os.environ.get("AZURE_FEDERATED_TOKEN_FILE"):
            credentials.append(WorkloadIdentityCredential())
            strategy += " -> workload identity"
        credential = ChainedTokenCredential(*credentials)

    return CredentialSelection(credential, strategy, enable_cae)


def build_async_credential(
    environment: RuntimeEnvironment, *, enable_cae: bool = False
) -> CredentialSelection[AsyncTokenCredential]:
    """Build an asynchronous credential chain for an environment."""

    if environment is RuntimeEnvironment.DEV:
        credential = AsyncChainedTokenCredential(
            AsyncAzureCliCredential(),
            AsyncAzureDeveloperCliCredential(),
            AsyncAzurePowerShellCredential(),
            AsyncVisualStudioCodeCredential(),
        )
        strategy = (
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code"
        )
    elif environment is RuntimeEnvironment.CI:
        credentials: list[AsyncTokenCredential] = []
        settings = _azure_pipelines_settings()
        if settings:
            tenant_id, client_id, service_connection_id, system_access_token = settings
            credentials.append(
                AsyncAzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=service_connection_id,
                    system_access_token=system_access_token,
                )
            )
        credentials.append(AsyncEnvironmentCredential())
        credential = AsyncChainedTokenCredential(*credentials)
        strategy = (
            "Azure Pipelines service connection -> pipeline environment variables"
            if settings
            else "pipeline environment variables (EnvironmentCredential)"
        )
    else:
        managed_identity_client_id = os.environ.get(
            USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID
        )
        managed_identity = AsyncManagedIdentityCredential(
            client_id=managed_identity_client_id
        )
        credentials = [managed_identity]
        strategy = (
            "user-assigned managed identity"
            if managed_identity_client_id
            else "system-assigned managed identity"
        )
        if os.environ.get("AZURE_FEDERATED_TOKEN_FILE"):
            credentials.append(AsyncWorkloadIdentityCredential())
            strategy += " -> workload identity"
        credential = AsyncChainedTokenCredential(*credentials)

    return CredentialSelection(credential, strategy, enable_cae)
