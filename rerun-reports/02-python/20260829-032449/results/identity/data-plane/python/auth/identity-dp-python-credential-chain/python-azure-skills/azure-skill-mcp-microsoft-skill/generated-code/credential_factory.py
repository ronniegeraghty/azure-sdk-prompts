"""Environment-specific Azure credential chains."""

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

CredentialT = TypeVar("CredentialT", TokenCredential, AsyncTokenCredential)


@dataclass(frozen=True)
class CredentialSelection(Generic[CredentialT]):
    """A credential plus a human-readable description of its policy."""

    credential: CredentialT
    strategy: str
    enable_cae: bool


def build_sync_credential(
    environment: RuntimeEnvironment, *, enable_cae: bool = False
) -> CredentialSelection[TokenCredential]:
    """Build a synchronous credential chain for the selected environment."""
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
        credentials: list[TokenCredential] = [EnvironmentCredential()]
        pipeline_config = _azure_pipelines_config()
        if pipeline_config:
            credentials.append(AzurePipelinesCredential(**pipeline_config))
        credential = ChainedTokenCredential(*credentials)
        strategy = _ci_strategy(pipeline_config is not None)
    else:
        credentials, strategy = _sync_production_credentials()
        credential = ChainedTokenCredential(*credentials)

    return CredentialSelection(credential, strategy, enable_cae)


def build_async_credential(
    environment: RuntimeEnvironment, *, enable_cae: bool = False
) -> CredentialSelection[AsyncTokenCredential]:
    """Build an asynchronous credential chain for the selected environment."""
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
        credentials: list[AsyncTokenCredential] = [AsyncEnvironmentCredential()]
        pipeline_config = _azure_pipelines_config()
        if pipeline_config:
            credentials.append(AsyncAzurePipelinesCredential(**pipeline_config))
        credential = AsyncChainedTokenCredential(*credentials)
        strategy = _ci_strategy(pipeline_config is not None)
    else:
        credentials, strategy = _async_production_credentials()
        credential = AsyncChainedTokenCredential(*credentials)

    return CredentialSelection(credential, strategy, enable_cae)


def _azure_pipelines_config() -> dict[str, str] | None:
    variable_map = {
        "tenant_id": "AZURE_TENANT_ID",
        "client_id": "AZURE_CLIENT_ID",
        "service_connection_id": "AZURE_SERVICE_CONNECTION_ID",
        "system_access_token": "SYSTEM_ACCESSTOKEN",
    }
    values = {argument: os.getenv(variable) for argument, variable in variable_map.items()}
    if all(values.values()):
        return {key: value for key, value in values.items() if value is not None}
    return None


def _ci_strategy(has_azure_pipelines_config: bool) -> str:
    strategy = "pipeline identity: environment credential"
    if has_azure_pipelines_config:
        strategy += " -> Azure Pipelines workload identity service connection"
    return strategy


def _sync_production_credentials() -> tuple[list[TokenCredential], str]:
    managed_identity_client_id = os.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID")
    managed_identity = ManagedIdentityCredential(client_id=managed_identity_client_id)
    credentials: list[TokenCredential] = [managed_identity]
    strategy = _managed_identity_strategy(managed_identity_client_id)

    if _workload_identity_is_configured():
        credentials.append(WorkloadIdentityCredential())
        strategy += " -> Kubernetes workload identity fallback"
    return credentials, strategy


def _async_production_credentials() -> tuple[list[AsyncTokenCredential], str]:
    managed_identity_client_id = os.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID")
    managed_identity = AsyncManagedIdentityCredential(
        client_id=managed_identity_client_id
    )
    credentials: list[AsyncTokenCredential] = [managed_identity]
    strategy = _managed_identity_strategy(managed_identity_client_id)

    if _workload_identity_is_configured():
        credentials.append(AsyncWorkloadIdentityCredential())
        strategy += " -> Kubernetes workload identity fallback"
    return credentials, strategy


def _managed_identity_strategy(client_id: str | None) -> str:
    if client_id:
        return "production identity: user-assigned managed identity"
    return "production identity: system-assigned managed identity"


def _workload_identity_is_configured() -> bool:
    return all(
        os.getenv(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )
