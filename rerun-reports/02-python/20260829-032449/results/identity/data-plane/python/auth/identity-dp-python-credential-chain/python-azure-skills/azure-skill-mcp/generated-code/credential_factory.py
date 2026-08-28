"""Environment-specific Azure credential chains."""

from __future__ import annotations

import os
from dataclasses import dataclass

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
)
from azure.identity.aio import (
    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
)
from azure.identity.aio import (
    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
)
from azure.identity.aio import (
    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
)
from azure.identity.aio import (
    ChainedTokenCredential as AsyncChainedTokenCredential,
)
from azure.identity.aio import (
    EnvironmentCredential as AsyncEnvironmentCredential,
)
from azure.identity.aio import (
    ManagedIdentityCredential as AsyncManagedIdentityCredential,
)
from azure.identity.aio import (
    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
)
from azure.identity.aio import (
    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
)

from environment_detector import RuntimeEnvironment


@dataclass(frozen=True)
class SyncCredentialSelection:
    credential: TokenCredential
    strategy: str
    enable_cae: bool


@dataclass(frozen=True)
class AsyncCredentialSelection:
    credential: AsyncTokenCredential
    strategy: str
    enable_cae: bool


def _pipeline_settings() -> dict[str, str] | None:
    names = {
        "tenant_id": "AZURE_TENANT_ID",
        "client_id": "AZURE_CLIENT_ID",
        "service_connection_id": "AZURE_SERVICE_CONNECTION_ID",
        "system_access_token": "SYSTEM_ACCESSTOKEN",
    }
    values = {argument: os.getenv(variable, "") for argument, variable in names.items()}
    return values if all(values.values()) else None


def _workload_identity_is_configured() -> bool:
    return all(
        os.getenv(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )


def _managed_identity_client_id() -> str | None:
    # Keep this separate from AZURE_CLIENT_ID, which workload identity also uses.
    return os.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID") or None


def build_sync_credential(
    environment: RuntimeEnvironment,
    *,
    enable_cae: bool = False,
) -> SyncCredentialSelection:
    """Build a synchronous credential chain for the selected environment."""
    if environment is RuntimeEnvironment.DEV:
        credential = ChainedTokenCredential(
            VisualStudioCodeCredential(),
            AzureCliCredential(),
            AzurePowerShellCredential(),
            AzureDeveloperCliCredential(),
        )
        strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
    elif environment is RuntimeEnvironment.CI:
        pipeline_settings = _pipeline_settings()
        if pipeline_settings:
            credential = ChainedTokenCredential(
                AzurePipelinesCredential(**pipeline_settings)
            )
            strategy = "Azure Pipelines workload identity service connection"
        else:
            credential = ChainedTokenCredential(EnvironmentCredential())
            strategy = "environment credential (service principal variables)"
    else:
        managed_identity_client_id = _managed_identity_client_id()
        credentials = [
            ManagedIdentityCredential(client_id=managed_identity_client_id)
        ]
        workload_configured = _workload_identity_is_configured()
        if workload_configured:
            credentials.append(WorkloadIdentityCredential())
        credential = ChainedTokenCredential(*credentials)
        identity_kind = (
            f"user-assigned managed identity ({managed_identity_client_id})"
            if managed_identity_client_id
            else "system-assigned managed identity"
        )
        strategy = (
            f"{identity_kind} -> Kubernetes workload identity"
            if workload_configured
            else identity_kind
        )

    return SyncCredentialSelection(credential, strategy, enable_cae)


def build_async_credential(
    environment: RuntimeEnvironment,
    *,
    enable_cae: bool = False,
) -> AsyncCredentialSelection:
    """Build an asynchronous credential chain for the selected environment."""
    if environment is RuntimeEnvironment.DEV:
        credential = AsyncChainedTokenCredential(
            AsyncVisualStudioCodeCredential(),
            AsyncAzureCliCredential(),
            AsyncAzurePowerShellCredential(),
            AsyncAzureDeveloperCliCredential(),
        )
        strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
    elif environment is RuntimeEnvironment.CI:
        pipeline_settings = _pipeline_settings()
        if pipeline_settings:
            credential = AsyncChainedTokenCredential(
                AsyncAzurePipelinesCredential(**pipeline_settings)
            )
            strategy = "Azure Pipelines workload identity service connection"
        else:
            credential = AsyncChainedTokenCredential(AsyncEnvironmentCredential())
            strategy = "environment credential (service principal variables)"
    else:
        managed_identity_client_id = _managed_identity_client_id()
        credentials = [
            AsyncManagedIdentityCredential(client_id=managed_identity_client_id)
        ]
        workload_configured = _workload_identity_is_configured()
        if workload_configured:
            credentials.append(AsyncWorkloadIdentityCredential())
        credential = AsyncChainedTokenCredential(*credentials)
        identity_kind = (
            f"user-assigned managed identity ({managed_identity_client_id})"
            if managed_identity_client_id
            else "system-assigned managed identity"
        )
        strategy = (
            f"{identity_kind} -> Kubernetes workload identity"
            if workload_configured
            else identity_kind
        )

    return AsyncCredentialSelection(credential, strategy, enable_cae)
