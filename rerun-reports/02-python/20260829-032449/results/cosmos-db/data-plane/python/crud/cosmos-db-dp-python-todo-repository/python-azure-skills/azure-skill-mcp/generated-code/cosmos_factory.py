from __future__ import annotations

import os
from dataclasses import dataclass
from types import TracebackType

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.cosmos.aio import ContainerProxy as AsyncContainerProxy
from azure.cosmos.container import ContainerProxy
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential


DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60
INDEXING_POLICY = {
    "indexingMode": "consistent",
    "automatic": True,
    "includedPaths": [{"path": "/*"}],
    "excludedPaths": [{"path": "/description/?"}],
}


@dataclass(frozen=True)
class CosmosSettings:
    endpoint: str
    database_name: str = "todo-db"
    container_name: str = "items"

    @classmethod
    def from_environment(cls) -> "CosmosSettings":
        endpoint = os.getenv("COSMOS_ENDPOINT")
        if not endpoint:
            raise RuntimeError(
                "COSMOS_ENDPOINT must contain the Azure Cosmos DB account endpoint."
            )
        return cls(
            endpoint=endpoint,
            database_name=os.getenv("COSMOS_DATABASE_NAME", "todo-db"),
            container_name=os.getenv("COSMOS_CONTAINER_NAME", "items"),
        )


@dataclass
class SyncCosmosResources:
    client: CosmosClient
    credential: DefaultAzureCredential
    container: ContainerProxy

    def __enter__(self) -> "SyncCosmosResources":
        return self

    def __exit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> None:
        self.client.close()
        self.credential.close()


@dataclass
class AsyncCosmosResources:
    client: AsyncCosmosClient
    credential: AsyncDefaultAzureCredential
    container: AsyncContainerProxy

    async def __aenter__(self) -> "AsyncCosmosResources":
        return self

    async def __aexit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> None:
        await self.client.close()
        await self.credential.close()


def create_sync_resources(
    settings: CosmosSettings | None = None,
) -> SyncCosmosResources:
    settings = settings or CosmosSettings.from_environment()
    credential = DefaultAzureCredential()
    client = CosmosClient(settings.endpoint, credential=credential)
    try:
        database = client.create_database_if_not_exists(id=settings.database_name)
        container = database.create_container_if_not_exists(
            id=settings.container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        client.close()
        credential.close()
        raise
    return SyncCosmosResources(client, credential, container)


async def create_async_resources(
    settings: CosmosSettings | None = None,
) -> AsyncCosmosResources:
    settings = settings or CosmosSettings.from_environment()
    credential = AsyncDefaultAzureCredential()
    client = AsyncCosmosClient(settings.endpoint, credential=credential)
    try:
        await client.__aenter__()
        database = await client.create_database_if_not_exists(
            id=settings.database_name
        )
        container = await database.create_container_if_not_exists(
            id=settings.container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        await client.close()
        await credential.close()
        raise
    return AsyncCosmosResources(client, credential, container)
